package widgets

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── Tree：可展开树 ────────────────────────────────────────────────

// TreeNode 是树节点。Children 为空的节点是叶子。
type TreeNode struct {
	Key      string
	Label    string
	Children []TreeNode
}

// Tree 是可展开树。Expanded 存展开节点的 Key 集合；Selected 为选中 Key；
// SelectedSet 是多选集合（TreeSelect 多选 tags 用，与 Selected 独立读取）。
type Tree struct {
	controlBase
	Nodes    []TreeNode
	Expanded map[string]bool
	Selected string
	// SelectedSet 多行选中高亮（优先级低于 Selected——非空时多行同亮）。
	SelectedSet map[string]bool
	Indent      float32
	RowH        float32
	Width       float32
	OnSelect    func(key string)
	OnToggle    func(key string, expanded bool)

	// Checkable 复选模式（antd checkable）：每行复选框是 **Checkbox 组件
	// 实例**（组合复用——hover 主色边框/indeterminate 半选/Space 键/禁用态
	// 全部白拿）。点击勾选/取消，Checked 集合变化时触发 OnCheck。
	Checkable bool
	// CheckStrictly 父子勾选是否解除关联（antd checkStrictly）：
	// false（默认）= **级联**——勾父全选子孙、子孙全选时父自动勾、
	// 部分选时父半选（Checkbox.Indeterminate）；true = 独立勾选（旧行为）。
	CheckStrictly bool
	// Checked 是勾选节点的 Key 集合（Checkable 模式下读写；级联模式下
	// 勾选父节点时子孙 Key 也会写入）。
	Checked map[string]bool
	OnCheck func(checkedKeys []string)

	// TapDownExec 浮层内嵌模式（TreeSelect 同款）：行点击 **down 即执行**
	// 而非 GestureDetector tap 配对——浮层内 tap 配对不可靠（Dropdown/
	// DatePicker 轮结论）。独立 Tree 窗（内容树）保持默认 false。
	TapDownExec bool

	// LoadData 异步加载子节点注入（antd loadData 直译——wg 无网络层，
	// 宿主闭包从内存/文件/模拟源取数据）：**首次展开**未加载的无子节点
	// 时触发；同步模式返回子节点直接挂载，异步模式返回 nil、宿主取数后
	// 调 SetChildren(key, children) 挂载（AfterFunc 延迟——模拟网络）。
	// LoadedKeys 标记已触发节点不重复加载。
	LoadData   func(key string) []TreeNode
	LoadedKeys map[string]bool
	// loadingKeys 加载中集合（渲染 loading 态——箭头转 loading 旋转）。
	loadingKeys map[string]bool

	taps []*ui.GestureDetector
	// checks 跨 Build 持有复选框实例（按 Key 池化）：Build 内新建实例
	// 会让 hover 态随重建丢失（Pagination/Tabs 先例）。
	checks map[string]*Checkbox
}

// NewTree 创建树。
func NewTree(nodes []TreeNode) *Tree {
	return &Tree{
		Nodes: nodes, Expanded: make(map[string]bool),
		Checked: make(map[string]bool),
		Indent:  18, RowH: 28, Width: 240,
	}
}

// Toggle 展开/收起节点。
func (t *Tree) Toggle(key string) {
	next := !t.Expanded[key]
	t.Expanded[key] = next
	t.MarkDirty()
	if t.OnToggle != nil {
		t.OnToggle(key, next)
	}
}

// needLoad 判断节点是否需要异步加载（antd loadData 语义）：LoadData
// 注入、未加载过、当前无子节点——三者同时成立才触发（有静态 Children
// 的节点不触发）。
func (t *Tree) needLoad(key string) bool {
	if t.LoadData == nil || t.LoadedKeys[key] {
		return false
	}
	var walk func(ns []TreeNode) *TreeNode
	walk = func(ns []TreeNode) *TreeNode {
		for i := range ns {
			if ns[i].Key == key {
				return &ns[i]
			}
			if r := walk(ns[i].Children); r != nil {
				return r
			}
		}
		return nil
	}
	n := walk(t.Nodes)
	return n != nil && len(n.Children) == 0
}

// SetChildren 挂载异步加载的子节点（宿主取数完成后调用——直接改 Nodes
// 树上该节点的 Children；MarkDirty 重建渲染）。加载中态同时清除。
func (t *Tree) SetChildren(key string, children []TreeNode) {
	var set func(ns []TreeNode) bool
	set = func(ns []TreeNode) bool {
		for i := range ns {
			if ns[i].Key == key {
				ns[i].Children = children
				return true
			}
			if set(ns[i].Children) {
				return true
			}
		}
		return false
	}
	if set(t.Nodes) {
		if t.loadingKeys != nil {
			delete(t.loadingKeys, key)
		}
		t.LoadedKeys[key] = true
		t.MarkDirty()
	}
}

// triggerLoad 首次展开触发异步加载（同步返回数据直接挂载；异步由宿主
// AfterFunc 延迟 SetChildren——antd loading 态期间箭头转 loading）。
func (t *Tree) triggerLoad(key string) {
	if t.loadingKeys == nil {
		t.loadingKeys = map[string]bool{}
	}
	t.loadingKeys[key] = true
	t.MarkDirty() // 箭头转 loading 态
	if t.LoadedKeys == nil {
		t.LoadedKeys = map[string]bool{}
	}
	// LoadedKeys 在 triggerLoad 即置位（防重复触发；SetChildren 兜底重置
	// loading 态）。
	t.LoadedKeys[key] = true
	kids := t.LoadData(key)
	if kids != nil {
		// 同步模式：延迟一拍挂载（Toggle 的 MarkDirty 已入队，挂载走
		// 同一重建帧——直改会与本帧渲染竞争）。
		time.AfterFunc(0, func() { t.SetChildren(key, kids) })
	}
}

// SyncChecked 重算全部父链的勾选/半选态（级联模式下外部直接改 Checked
// map 后调用——预置状态、程序化勾选都需要，否则父节点半选不生效）。
func (t *Tree) SyncChecked() {
	for _, n := range t.Nodes {
		t.syncAncestors(t.Nodes, n.Key)
	}
	t.MarkDirty()
}

// collectDescendants 收集 key 节点的全部子孙 Key（含自身）。
func (t *Tree) collectDescendants(nodes []TreeNode, key string, out *[]string) bool {
	for _, n := range nodes {
		if n.Key == key {
			var walk func(ns []TreeNode)
			walk = func(ns []TreeNode) {
				for _, c := range ns {
					*out = append(*out, c.Key)
					walk(c.Children)
				}
			}
			*out = append(*out, n.Key)
			walk(n.Children)
			return true
		}
		if t.collectDescendants(n.Children, key, out) {
			return true
		}
	}
	return false
}

// findParent 返回 key 的父节点（根层返回 nil）。
func (t *Tree) findParent(nodes []TreeNode, key string) *TreeNode {
	for i := range nodes {
		n := &nodes[i]
		for _, c := range n.Children {
			if c.Key == key {
				return n
			}
		}
		if p := t.findParent(n.Children, key); p != nil {
			return p
		}
	}
	return nil
}

// syncAncestors 级联模式下的父链同步：全选→父勾、部分→父半选（写
// Indeterminate 进池化 Checkbox）、全不选→父清。
func (t *Tree) syncAncestors(nodes []TreeNode, key string) {
	p := t.findParent(nodes, key)
	if p == nil {
		return
	}
	total, on := 0, 0
	var walk func(ns []TreeNode)
	walk = func(ns []TreeNode) {
		for _, c := range ns {
			total++
			if t.Checked[c.Key] {
				on++
			}
			walk(c.Children)
		}
	}
	walk(p.Children)
	cb := t.checks[p.Key]
	switch {
	case on == 0:
		t.Checked[p.Key] = false
		if cb != nil {
			cb.Indeterminate = false
		}
	case on == total:
		t.Checked[p.Key] = true
		if cb != nil {
			cb.Indeterminate = false
		}
	default:
		t.Checked[p.Key] = false
		if cb != nil {
			cb.Indeterminate = true
		}
	}
	t.syncAncestors(nodes, p.Key)
}

// Build 实现 ui.Widget。
func (t *Tree) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(t)
	th := ctx.ThemeOr()
	col := scene.Flex(scene.Vertical, 0).SetWidth(t.Width)
	t.taps = t.taps[:0]

	var addNodes func(nodes []TreeNode, depth int)
	addNodes = func(nodes []TreeNode, depth int) {
		for i := range nodes {
			n := nodes[i]
			row := scene.Flex(scene.Horizontal, 4).SetHeight(t.RowH).SetWidth(t.Width)
			hasKids := len(n.Children) > 0
			// 异步加载（antd loadData）：可加载节点（LoadData 注入+未加载+
			// 无静态 Children）显示展开箭头——点击展开触发 LoadData；加载
			// 中箭头转 loading（旋转 Icon）。
			loadable := t.LoadData != nil && !t.LoadedKeys[n.Key] && !hasKids
			loading := t.loadingKeys != nil && t.loadingKeys[n.Key]
			arrow := ""
			if hasKids || loadable {
				if loading {
					arrow = "◐" // loading 占位（真实旋转走 Icon Spin——行内
				} else if t.Expanded[n.Key] {
					arrow = "▾"
				} else {
					arrow = "▸"
				}
			}
			// 缩进 + 展开箭头
			prefix := ""
			for k := 0; k < depth; k++ {
				prefix += "  "
			}
			prefix += arrow
			pl := text.LayoutText(prefix, text.DefaultFont(th.Font.SizeSM), text.Options{})
			pt := scene.Text(scene.TextSpec{Src: prefix, Layout: pl, Color: th.Color.TextMuted})
			pt.SetMargin(geom.InLTRB(float32(depth)*t.Indent+2, (t.RowH-pl.Size.H)/2, 0, 0))
			row.Add(pt)

			// 复选模式：缩进/箭头后嵌 Checkbox 组件实例（池化按 Key）。
			// CheckStrictly=true 独立勾选；false（默认）级联——勾父全选
			// 子孙、子孙联动父链（全选勾/部分半选/全不选清）。
			if t.Checkable {
				if t.checks == nil {
					t.checks = make(map[string]*Checkbox)
				}
				ckey := n.Key
				cb := t.checks[ckey]
				if cb == nil {
					cb = &Checkbox{}
					cb.OnChange = func(on bool) {
						if t.Checked == nil {
							t.Checked = make(map[string]bool)
						}
						if t.CheckStrictly {
							t.Checked[ckey] = on
						} else {
							var keys []string
							t.collectDescendants(t.Nodes, ckey, &keys)
							for _, k := range keys {
								t.Checked[k] = on
							}
							t.syncAncestors(t.Nodes, ckey)
						}
						t.MarkDirty()
						if t.OnCheck != nil {
							keys := make([]string, 0, len(t.Checked))
							for k, v := range t.Checked {
								if v {
									keys = append(keys, k)
								}
							}
							t.OnCheck(keys)
						}
					}
					t.checks[ckey] = cb
				}
				cb.Checked = t.Checked[n.Key] // map 是真相源，每次同步
				// 级联模式下半选态从 Checked 集合推导（外部直接改 map 也生效）：
				// 子孙部分勾选→半选；池化实例的 Indeterminate 是上次运行残留。
				if !t.CheckStrictly && hasKids {
					total, on := 0, 0
					var walk func(ns []TreeNode)
					walk = func(ns []TreeNode) {
						for _, c := range ns {
							total++
							if t.Checked[c.Key] {
								on++
							}
							walk(c.Children)
						}
					}
					walk(n.Children)
					cb.Indeterminate = on > 0 && on < total
				}
				cbn := cb.Build(ctx)
				cbn.SetPositionKeepSize(true)
				cbn.SetStackPosition(geom.Pt(float32(depth)*t.Indent+2+pl.Size.W+4, (t.RowH-16)/2))
				row.Add(cbn)
			}

			ll := text.LayoutText(n.Label, text.DefaultFont(th.Font.SizeBase), text.Options{})
			fg := th.Color.Text
			if n.Key == t.Selected || t.SelectedSet[n.Key] {
				fg = th.Color.Primary
				row.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Selection)})
			}
			lt := scene.Text(scene.TextSpec{Src: n.Label, Layout: ll, Color: fg})
			lt.SetMargin(geom.InLTRB(0, (t.RowH-ll.Size.H)/2, 0, 0))
			row.Add(lt)

			if ctx.Input != nil && (hasKids || loadable || t.OnSelect != nil) {
				row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "tree-node", Cursor: scene.CursorPointer})
				key, kids := n.Key, hasKids
				loadableKey, loadableRow := n.Key, loadable
				onTap := func(geom.Point) {
					if kids {
						t.Toggle(key)
					} else if loadableRow {
						// 首次展开可加载节点：展开态置位+触发 LoadData。
						if !t.Expanded[loadableKey] {
							t.Expanded[loadableKey] = true
						}
						t.triggerLoad(loadableKey)
						if t.OnToggle != nil {
							t.OnToggle(loadableKey, true)
						}
					}
					// 选中态与回调解耦：OnSelect 为 nil 时点击节点仍默认选中
					// （antd 点击节点即选中，不依赖调用方是否监听）。
					if t.Selected != key {
						t.Selected = key
						t.MarkDirty()
					}
					if t.OnSelect != nil {
						t.OnSelect(key)
					}
				}
				if t.TapDownExec {
					// 浮层内嵌：tap 配对不可靠，down 即执行。
					ctx.Input.Set(row, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
						if !ev.Down() {
							return false
						}
						onTap(geom.Point{})
						return true
					}})
				} else {
					g := &ui.GestureDetector{OnTap: onTap}
					t.taps = append(t.taps, g)
					ctx.Input.Set(row, g.Handler())
				}
			}
			col.Add(row)
			if hasKids && t.Expanded[n.Key] {
				addNodes(n.Children, depth+1)
			}
		}
	}
	addNodes(t.Nodes, 0)
	t.node = col
	return col
}

// ── Form / FormItem ───────────────────────────────────────────────

// FormItem 是表单项：标签 + 控件 + 校验提示。
// FormLayout 是表单项布局（antd Form layout）。
type FormLayout uint8

const (
	// FormHorizontal 水平布局（标签左、控件右，默认）。
	FormHorizontal FormLayout = iota
	// FormVertical 垂直布局（标签在上、控件在下）。
	FormVertical
	// FormInline 行内紧凑（标签与控件同行，无整行宽度）。
	FormInline
)

type FormItem struct {
	Label    string
	Child    ui.Widget
	Required bool
	Error    string
	LabelW   float32
	// Layout 表单项布局（antd Form layout）：horizontal 默认/vertical 标签在上/
	// inline 同行紧凑。零值 = horizontal。
	Layout FormLayout
	// LabelColSpan 标签占 24 栅格数（antd labelCol.span）：>0 时按
	// (span/24)×父宽 计算标签宽，覆盖 LabelW。
	LabelColSpan int
}

// Build 实现 ui.Widget。
func (f FormItem) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	// vertical：标签整行在上、控件在下（antd layout=vertical）。
	if f.Layout == FormVertical {
		return f.buildVertical(ctx, th)
	}
	row := scene.Flex(scene.Horizontal, th.Space.SM)
	labelW := f.LabelW
	if labelW <= 0 {
		labelW = 80
	}
	// labelCol.span：按 (span/24)×父宽 算标签宽（antd labelCol 语义）。
	if f.LabelColSpan > 0 {
		pw := ctx.ViewportSize().W - 80
		if pw > 0 {
			labelW = pw * float32(f.LabelColSpan) / 24
		}
	}
	// antd Form：必填星号红色（#ff4d4f）置于标签前；标签为正文色而非弱化灰。
	labelRow := scene.Flex(scene.Horizontal, 0)
	if f.Required {
		font := text.DefaultFont(th.Font.SizeBase)
		star := scene.Text(scene.TextSpec{
			Src: "*", Layout: text.LayoutText("*", font, text.Options{}),
			Color: th.Color.Danger,
		})
		star.SetMargin(geom.InLTRB(0, 6, 3, 0))
		labelRow.Add(star)
	}
	ll := text.LayoutText(f.Label, text.DefaultFont(th.Font.SizeBase), text.Options{})
	lt := scene.Text(scene.TextSpec{Src: f.Label, Layout: ll, Color: th.Color.Text})
	lt.SetWidth(labelW).SetMargin(geom.InLTRB(0, 6, 0, 0))
	labelRow.Add(lt)
	row.Add(labelRow)
	if f.Child != nil {
		row.Add(buildChild(ctx, f.Child))
	}
	col := scene.Flex(scene.Vertical, 2, row)
	if f.Error != "" {
		el := text.LayoutText(f.Error, text.DefaultFont(th.Font.SizeXS), text.Options{})
		et := scene.Text(scene.TextSpec{Src: f.Error, Layout: el, Color: th.Color.Danger})
		et.SetMargin(geom.InLTRB(labelW+th.Space.SM, 0, 0, 0))
		// 被动定位锚：错误文字色断言用（scene.Text 直建，无 Label 锚）。
		et.SetHit(scene.HitSpec{Label: "form-error"})
		col.Add(et)
	}
	return col
}

// buildVertical 垂直布局：标签（含必填星）在上、控件在下，整行宽度。
func (f FormItem) buildVertical(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	col := scene.Flex(scene.Vertical, 4)
	labelRow := scene.Flex(scene.Horizontal, 0)
	if f.Required {
		font := text.DefaultFont(th.Font.SizeBase)
		star := scene.Text(scene.TextSpec{Src: "*", Layout: text.LayoutText("*", font, text.Options{}), Color: th.Color.Danger})
		star.SetMargin(geom.InLTRB(0, 0, 3, 0))
		labelRow.Add(star)
	}
	if f.Label != "" {
		ll := text.LayoutText(f.Label, text.DefaultFont(th.Font.SizeBase), text.Options{})
		lt := scene.Text(scene.TextSpec{Src: f.Label, Layout: ll, Color: th.Color.Text})
		labelRow.Add(lt)
	}
	if len(labelRow.Children()) > 0 {
		col.Add(labelRow)
	}
	if f.Child != nil {
		col.Add(buildChild(ctx, f.Child))
	}
	if f.Error != "" {
		el := text.LayoutText(f.Error, text.DefaultFont(th.Font.SizeXS), text.Options{})
		et := scene.Text(scene.TextSpec{Src: f.Error, Layout: el, Color: th.Color.Danger})
		et.SetHit(scene.HitSpec{Label: "form-error"})
		col.Add(et)
	}
	return col
}

// Validator 校验一个值，返回错误信息（空串表示通过）。
type Validator func(value string) string

// RequiredRule 返回必填校验器。
func RequiredRule(msg string) Validator {
	if msg == "" {
		msg = "此项必填"
	}
	return func(v string) string {
		if v == "" {
			return msg
		}
		return ""
	}
}

// FormRule 是声明式校验规则（antd Rule 直译）。
// 校验顺序：required → whitespace → len/min/max → pattern → email/url
// → validator（首个失败即停，antd validateFirst 默认语义）。
type FormRule struct {
	// Required 必填（非空）。
	Required bool
	// Whitespace 仅空白也算空（配合 Required；antd whitespace）。
	Whitespace bool
	// Len 精确长度（>0 生效；antd len）。
	Len int
	// Min/MinRune 最小长度（>0 生效；antd min——按 rune 计，中文友好）。
	Min int
	// Max 最大长度（>0 生效；antd max——按 rune 计）。
	Max int
	// Pattern 正则（Go regexp 语法；antd pattern）。
	Pattern string
	// Email/URL 内置格式校验（antd type:"email"/"url"）。
	Email bool
	URL   bool
	// Validator 自定义校验（antd validator——wg 无 Promise，同步函数
	// 替代；返回空串=通过）。空值时跳过（required 单独管空）。
	Validator func(value string) string
	// Message 自定义报错（覆盖内置模板；antd message）。
	Message string
}

// FormRule 校验一个值，返回错误信息（空串=通过）。
func (r FormRule) check(value string) string {
	fail := func(defaultTpl string) string {
		if r.Message != "" {
			return r.Message
		}
		return defaultTpl
	}
	if r.Required && value == "" {
		return fail("此项必填")
	}
	if r.Whitespace && value != "" && isBlankString(value) {
		return fail("不能只有空白字符")
	}
	// len/min/max 按 rune 计（中文友好；antd 对 string 同语义）。
	runes := []rune(value)
	n := len(runes)
	if r.Len > 0 && n != r.Len {
		return fail(fmt.Sprintf("长度必须为 %d", r.Len))
	}
	if r.Min > 0 && n < r.Min {
		return fail(fmt.Sprintf("长度不能少于 %d", r.Min))
	}
	if r.Max > 0 && n > r.Max {
		return fail(fmt.Sprintf("长度不能超过 %d", r.Max))
	}
	if r.Pattern != "" && value != "" {
		re, err := regexp.Compile(r.Pattern)
		if err != nil {
			return fail("规则正则无效: " + r.Pattern)
		}
		if !re.MatchString(value) {
			return fail("格式不正确")
		}
	}
	if r.Email && value != "" && !isEmail(value) {
		return fail("邮箱格式不正确")
	}
	if r.URL && value != "" && !isURL(value) {
		return fail("网址格式不正确")
	}
	if r.Validator != nil && value != "" {
		if msg := r.Validator(value); msg != "" {
			return msg
		}
	}
	return ""
}

// isBlankString 仅空白字符（空格/制表/换行）。
func isBlankString(s string) bool {
	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			return false
		}
	}
	return true
}

// isEmail 宽松邮箱格式（x@y.z 形态，antd type:"email" 近似）。
func isEmail(s string) bool {
	at := strings.IndexByte(s, '@')
	if at <= 0 || at == len(s)-1 {
		return false
	}
	dot := strings.LastIndexByte(s[at+1:], '.')
	return dot > 0 && at+1+dot < len(s)-1
}

// isURL http(s):// 前缀 + 非空主机（antd type:"url" 近似）。
func isURL(s string) bool {
	if !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
		return false
	}
	host := strings.TrimPrefix(strings.TrimPrefix(s, "https://"), "http://")
	return host != "" && !strings.HasPrefix(host, "/")
}

// Form 是表单容器。校验器按字段名注册，Validate 汇总错误。
type Form struct {
	ui.Base
	Items      []FormItem
	Values     map[string]string
	Validators map[string][]Validator
	RuleSet    map[string][]FormRule
	Errors     map[string]string
}

// NewForm 创建表单。
func NewForm() *Form {
	return &Form{
		Values:     make(map[string]string),
		Validators: make(map[string][]Validator),
		RuleSet:    make(map[string][]FormRule),
		Errors:     make(map[string]string),
	}
}

// Rule 为字段注册过程式校验器。
func (f *Form) Rule(field string, v ...Validator) *Form {
	f.Validators[field] = append(f.Validators[field], v...)
	return f
}

// Rules 为字段注册声明式规则（antd rules 直译；先于 Validator 执行）。
func (f *Form) Rules(field string, rules ...FormRule) *Form {
	f.RuleSet[field] = append(f.RuleSet[field], rules...)
	return f
}

// Validate 校验所有字段，返回是否通过（同时写入 Errors）。
func (f *Form) Validate() bool {
	f.Errors = make(map[string]string)
	// 先同步字段当前值：Input.Value 由编辑器实时维护，但 Form.Values
	// 从未被填充过，不在这里回读会让校验永远读到空值（--audit 实测）。
	for _, it := range f.Items {
		if it.Label == "" {
			continue
		}
		switch ch := it.Child.(type) {
		case *Input:
			f.Values[it.Label] = ch.Value
		case *InputNumber:
			f.Values[it.Label] = fmt.Sprintf("%v", ch.Value)
		}
	}
	ok := true
	for field, rules := range f.RuleSet {
		for _, r := range rules {
			if msg := r.check(f.Values[field]); msg != "" {
				f.Errors[field] = msg
				ok = false
				break
			}
		}
	}
	for field, rules := range f.Validators {
		if f.Errors[field] != "" {
			continue // 声明式规则已报错，过程式校验器跳过（首错即停）
		}
		for _, v := range rules {
			if msg := v(f.Values[field]); msg != "" {
				f.Errors[field] = msg
				ok = false
				break
			}
		}
	}
	f.MarkDirty()
	return ok
}

// Build 实现 ui.Widget。
func (f *Form) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(f)
	th := ctx.ThemeOr()
	col := scene.Flex(scene.Vertical, th.Space.MD)
	for i := range f.Items {
		item := f.Items[i]
		// 校验失败的字段：错误文字 + 输入框红边（antd status="error"）。
		// Status 直接写到 Child 输入实例上（Validate 后 MarkDirty 重建生效）；
		// 调用方手动设置的 Status 会被下次 Validate 覆盖（表单驱动语义）。
		if msg := f.Errors[item.Label]; msg != "" {
			item.Error = msg
			if ch, ok := item.Child.(*Input); ok {
				ch.Status = InputStatusError
			}
		} else if ch, ok := item.Child.(*Input); ok && ch.Status == InputStatusError {
			// 上轮校验失败、本轮已通过：边框回常态。
			ch.Status = StatusOK
		}
		col.Add(item.Build(ctx))
	}
	return col
}
