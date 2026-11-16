// Package deps 是 wg 的静态依赖检查器（契约 C17：单向依赖）。
//
// 它只做一件事：解析 wg 目录下所有 .go 文件的 import，按下方规则表判定违规。
// 之所以用源码解析而不是 go list：
//   - 不触发编译，毫秒级，可以挂在 --selftest 里；
//   - 能精确到"文件 + import 行"，报错可定位。
//
// 规则来源：05-API-ROADMAP.md §1「依赖方向（C17）」与 02-CANVAS.md §1。
package deps

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// ModulePrefix 是 wg 的导入前缀。
const ModulePrefix = "github.com/energye/lcl/wg"

// LCLPackage 是 LCL 绑定包（只允许 host 与 raster 的 GDI 后端引用）。
const LCLPackage = "github.com/energye/lcl"

// Violation 描述一处依赖违规。
type Violation struct {
	File   string // 相对 wg 根目录
	Line   int
	Import string
	Rule   string
}

// String 返回单行可读信息。
func (v Violation) String() string {
	return fmt.Sprintf("%s:%d 禁止 import %q（%s）", v.File, v.Line, v.Import, v.Rule)
}

// ruleSet 描述某个包允许依赖哪些 wg 子包。
type ruleSet struct {
	// allow 是允许的 wg 子包（相对包名，根为 ""）。
	allow map[string]bool
	// allowAll 表示允许依赖任意 wg 子包（组合根/诊断/检查器）。
	allowAll bool
	// allowLCL 表示允许 import LCL 绑定。
	allowLCL bool
	// allowUnsafe 表示允许 import unsafe。
	allowUnsafe bool
}

func set(names ...string) map[string]bool {
	m := make(map[string]bool, len(names))
	for _, n := range names {
		m[n] = true
	}
	return m
}

// rules 是 C17 的唯一事实来源。
//
// 关键约束（改动前请先读 05-API-ROADMAP.md §1）：
//   - 只有 host 与 raster/gdi、raster/software 可以 import lcl；
//   - render 不依赖 lcl，也不依赖 scene（否则录制层被渲染树绑死）；
//   - host 可以 import input，但 input 不 import host（方向单向：平台层产生引擎事件）；
//   - 上层（ui/widgets）不得出现 unsafe / lcl。
var rules = map[string]ruleSet{
	"":      {allowAll: true, allowLCL: true}, // 组合根：门面需要 lcl.IEngForm 出现在签名里
	"geom":  {allow: set()},
	"theme": {allow: set("geom")},
	"text": {
		// 度量与字形句柄要走 GDI，因此允许 unsafe 操作 LOGFONT/TEXTMETRIC；
		// 但不允许 import lcl（文本子系统必须能在无窗口环境工作）。
		allow: set("geom"), allowUnsafe: true,
	},
	"render": {allow: set("geom", "text", "theme")},
	"raster": {
		// 帧缓冲（backing）需要 unsafe 做字节视图，这是 raster 的核心职责；
		// 但 raster 本体不得 import lcl（LCL 位图对接放在 host 与 raster/gdi）。
		allow: set("geom", "render"), allowUnsafe: true,
	},
	"raster/software": {
		// 软件光栅需要 unsafe 做像素视图；不得碰 lcl。
		allow: set("geom", "render", "raster"), allowUnsafe: true,
	},
	"raster/gdi": {
		// GDI 后端是唯一可以同时碰 lcl 与 unsafe 的光栅路径。
		// 它还需要 text 的字体描述（HFONT 缓存与排版结果），
		// 因此允许依赖 text —— 这是"文字回放"职责决定的，不是图省事。
		allow: set("geom", "render", "raster", "text"), allowLCL: true, allowUnsafe: true,
	},
	"scene":     {allow: set("geom", "render", "text", "theme")},
	"composite": {allow: set("geom", "render", "raster", "scene", "theme")},
	"frame":     {allow: set("geom")},
	"host": {
		// host 需要 input 的事件类型：它把平台消息归一化成引擎事件，
		// 不 import input 就只能再抄一份结构体，再写一层纯搬运的转换。
		// 反向依赖（input → host）绝不成立，这才是真正要防的环。
		allow:    set("geom", "frame", "render", "raster", "scene", "composite", "text", "theme", "input"),
		allowLCL: true, allowUnsafe: true,
	},
	"input":   {allow: set("geom", "frame", "render", "scene", "text")},
	"overlay": {allow: set("geom", "host", "input", "scene", "theme")},
	"ui":      {allow: set("geom", "frame", "input", "overlay", "render", "scene", "text", "theme")},
	// widgets 需要 text 排版文本标签（Label/Input 的占位符与选区几何），
	// 以及 overlay 的浮层入口（Tooltip/Popover）；render 用于 CustomPaintNode
	// 录制图形（Spin/图表），与 ui 层同级放行；不得碰 GDI 句柄。
	"widgets": {allow: set("geom", "input", "overlay", "render", "scene", "text", "theme", "ui")},
	// widgets/chart 与 widgets 同级：图表也是组件，只多一层数据映射。
	"widgets/chart": {allow: set("geom", "render", "scene", "text", "theme", "ui")},
	// compat 是旧 API 门面，转发到新组件。
	"compat": {allow: set("geom", "scene", "ui", "widgets")},
	// migration 是示例组合层：把三个旧示例重建为组件场景，允许依赖组件与渲染基础。
	"migration": {allow: set("geom", "input", "render", "scene", "theme", "ui", "widgets", "widgets/chart")},
	"diag":    {allowAll: true, allowLCL: true},
	"contracts": {
		// 契约包只做断言，不参与渲染；C17 的断言实现需要调用检查器本身，
		// 因此放行 internal/deps（同模块的静态分析工具，无运行时副作用）。
		allow: set("geom", "frame", "raster", "raster/software", "raster/gdi",
			"host", "render", "scene", "composite",
			"input", "overlay", "ui", "theme", "text", "internal/deps"),
	},
	"internal":      {allowAll: true, allowUnsafe: true},
	"internal/deps": {allowAll: true},
	"cmd":           {allowAll: true, allowLCL: true},
}

// exempt 目录整体跳过（不属于架构管辖范围）。
var exempt = map[string]bool{
	"testdata":     true,
	"vendor":       true,
	"node_modules": true,
}

// Check 扫描 root（wg 根目录）并返回违规列表。
func Check(root string) ([]Violation, error) {
	root = filepath.Clean(root)
	if st, err := os.Stat(root); err != nil || !st.IsDir() {
		return nil, fmt.Errorf("wg/deps: 目录不可用: %s", root)
	}

	var out []Violation
	fset := token.NewFileSet()

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if exempt[name] {
				return fs.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		// 测试文件允许放宽（会 import testing/反射/对比工具），不参与架构判定。
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return nil
		}
		relDir := filepath.ToSlash(filepath.Dir(rel))
		if relDir == "." {
			relDir = ""
		}
		rs, ok := lookupRule(relDir)
		if !ok {
			// 未声明规则的目录：默认严格（不允许 lcl/unsafe，也不允许跨包）。
			rs = ruleSet{allow: set()}
		}

		f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return fmt.Errorf("wg/deps: 解析失败 %s: %w", rel, err)
		}
		for _, spec := range f.Imports {
			imp, _ := strconv.Unquote(spec.Path.Value)
			line := fset.Position(spec.Pos()).Line
			if v, bad := judge(relDir, imp, rs); bad {
				out = append(out, Violation{File: filepath.ToSlash(rel), Line: line, Import: imp, Rule: v})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].File != out[j].File {
			return out[i].File < out[j].File
		}
		return out[i].Line < out[j].Line
	})
	return out, nil
}

// lookupRule 按最长前缀匹配规则表（"raster/gdi" 优先于 "raster"）。
func lookupRule(dir string) (ruleSet, bool) {
	for d := dir; ; {
		if rs, ok := rules[d]; ok {
			return rs, true
		}
		i := strings.LastIndexByte(d, '/')
		if i < 0 {
			if rs, ok := rules[""]; ok && d == "" {
				return rs, true
			}
			return ruleSet{}, false
		}
		d = d[:i]
	}
}

// judge 判定单条 import 是否违规，返回原因。
func judge(dir, imp string, rs ruleSet) (string, bool) {
	switch imp {
	case "unsafe":
		if !rs.allowUnsafe {
			return "C17: 上层不得使用 unsafe", true
		}
		return "", false
	}
	if imp == LCLPackage || strings.HasPrefix(imp, LCLPackage+"/") {
		// wg 自身的包不算 lcl 绑定
		if strings.HasPrefix(imp, ModulePrefix) {
			return judgeWG(dir, imp, rs)
		}
		if !rs.allowLCL {
			return "C17: 只有 host / raster 后端可以 import lcl", true
		}
		return "", false
	}
	if strings.HasPrefix(imp, ModulePrefix) {
		return judgeWG(dir, imp, rs)
	}
	// 标准库与第三方：不限制（第三方白名单在 M3 再收紧）。
	return "", false
}

func judgeWG(dir, imp string, rs ruleSet) (string, bool) {
	sub := strings.TrimPrefix(imp, ModulePrefix)
	sub = strings.TrimPrefix(sub, "/")
	if rs.allowAll || rs.allow[sub] {
		return "", false
	}
	// 允许依赖自己的子包（例如 raster/software 依赖 raster 已在 allow 里显式声明）。
	return fmt.Sprintf("C17: %q 不在 %q 的允许依赖内", sub, pkgName(dir)), true
}

func pkgName(dir string) string {
	if dir == "" {
		return "wg"
	}
	return dir
}

// WGRoot 返回 wg 包根目录（从调用点位置推导，用于 --selftest 与 checkdeps）。
//
// 调用方传入自身文件路径（runtime.Caller 的 file）与本文件期望的相对深度。
func WGRoot(fromFile string, up int) string {
	dir := filepath.Dir(fromFile)
	for i := 0; i < up; i++ {
		dir = filepath.Dir(dir)
	}
	return dir
}

// Rules 返回规则表的人类可读形式（按目录名排序），供 checkdeps -v 使用。
func Rules() []string {
	keys := make([]string, 0, len(rules))
	for k := range rules {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]string, 0, len(keys))
	for _, k := range keys {
		rs := rules[k]
		name := pkgName(k)
		var parts []string
		switch {
		case rs.allowAll:
			parts = append(parts, "允许任意 wg 包")
		default:
			al := make([]string, 0, len(rs.allow))
			for a := range rs.allow {
				al = append(al, pkgName(a))
			}
			sort.Strings(al)
			if len(al) == 0 {
				parts = append(parts, "无 wg 内部依赖")
			} else {
				parts = append(parts, "允许 "+strings.Join(al, ", "))
			}
		}
		if rs.allowLCL {
			parts = append(parts, "允许 lcl")
		}
		if rs.allowUnsafe {
			parts = append(parts, "允许 unsafe")
		}
		out = append(out, fmt.Sprintf("%-16s %s", name, strings.Join(parts, "；")))
	}
	return out
}
