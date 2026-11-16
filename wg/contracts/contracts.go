// Package contracts 把架构文档里的硬约束写成可执行断言。
//
// 设计要点（05-API-ROADMAP.md §7.5「契约与断言同生共死」）：
//   - 每条契约一个编号（C1–C28），编号与 docs/wg-engine 三份文档严格一致；
//   - 每条契约声明它从哪个里程碑起必须为绿（Since）；
//   - 未到期的契约报 PENDING 而不是 FAIL——但到期的契约如果没有断言，
//     Validate 会直接报错，防止"文档写了、代码没管"。
//
// 断言分两类：
//   - 静态断言（依赖、格式、常量）可在 M0 无窗口运行；
//   - 动态断言（出图、性能、输入）需要对应的 diag 命令，写在 diag/ 里，
//     这里只保留登记与归口，避免把渲染依赖拉进 contracts 包。
package contracts

import (
	"fmt"
	"runtime"
	"sort"
	"strings"

	"github.com/energye/lcl/wg/internal/deps"
)

// Stage 是里程碑。
type Stage uint8

// 里程碑编号，与 05-API-ROADMAP.md §5 一致。
const (
	M0 Stage = iota
	M1
	M2
	M3
	M4
	M5
)

// String 返回里程碑名。
func (s Stage) String() string {
	if s > M5 {
		return "M?"
	}
	return fmt.Sprintf("M%d", uint8(s))
}

// Status 是断言结果状态。
type Status uint8

const (
	// Pass 通过。
	Pass Status = iota
	// Fail 失败——门禁必须红。
	Fail
	// Pending 尚未到期的里程碑，不计入失败。
	Pending
	// Missing 已到期但没有断言实现——视为失败（防止漏写）。
	Missing
)

// String 返回状态名。
func (s Status) String() string {
	switch s {
	case Pass:
		return "PASS"
	case Fail:
		return "FAIL"
	case Pending:
		return "PENDING"
	case Missing:
		return "MISSING"
	default:
		return "?"
	}
}

// Contract 是一条架构契约。
type Contract struct {
	// ID 形如 "C7"，与文档一致。
	ID string
	// Title 是一句话约束描述。
	Title string
	// Owner 是归属层标记（⓪①②③），与文档一致。
	Owner string
	// Gate 是验证它的 diag 命令。
	Gate string
	// Since 是"从哪个里程碑起必须为绿"。
	Since Stage
	// Assert 是断言实现；nil 表示尚未实现（到期的 nil 会被 Validate 拦下）。
	Assert func() error
}

// Result 是单条契约的执行结果。
type Result struct {
	Contract
	Status Status
	Detail string
}

// String 返回定宽单行输出（供 diag 表格使用）。
func (r Result) String() string {
	line := fmt.Sprintf("%-4s %-8s %-4s %-22s %s", r.ID, r.Status, r.Since, r.Gate, r.Title)
	if r.Detail != "" {
		line += "  [" + r.Detail + "]"
	}
	return line
}

// Current 是当前里程碑，决定哪些契约必须为绿。
//
// 跟着引擎实现走：到了 M3，C1–C28 全部到期，任何一条没有断言
// 都会被 CheckAll 报 MISSING（而不是静默跳过）。M3 收口后
// "未来门禁必须报红"由 TestMissingGateIsRed 用临时契约验证，
// 不再依赖某个真实契约恰好还没实现。
var Current = M4

// SetStage 设置当前里程碑（diag 用命令行参数覆盖）。
func SetStage(s Stage) { Current = s }

// All 返回全部契约（按编号排序）。
func All() []Contract {
	out := make([]Contract, len(table))
	copy(out, table)
	sort.Slice(out, func(i, j int) bool { return lessID(out[i].ID, out[j].ID) })
	return out
}

// Get 按编号取契约。
func Get(id string) (Contract, bool) {
	for _, c := range table {
		if strings.EqualFold(c.ID, id) {
			return c, true
		}
	}
	return Contract{}, false
}

// CheckAll 执行所有断言，返回逐条结果。
//
// 不会 panic：断言内部若 panic 会被捕获并记为 FAIL（附 panic 信息），
// 因为自检工具本身崩溃比一条断言失败更难排查。
func CheckAll() []Result {
	list := All()
	out := make([]Result, 0, len(list))
	for _, c := range list {
		out = append(out, run(c, Current))
	}
	return out
}

// Run 执行到指定里程碑的结果集（不修改全局 Current）。
func Run(stage Stage) []Result {
	list := All()
	out := make([]Result, 0, len(list))
	for _, c := range list {
		out = append(out, run(c, stage))
	}
	return out
}

func run(c Contract, stage Stage) Result {
	r := Result{Contract: c}
	if c.Assert == nil {
		if stage >= c.Since {
			r.Status = Missing
			r.Detail = "到期但未实现断言"
		} else {
			r.Status = Pending
			r.Detail = "自 " + c.Since.String() + " 起生效"
		}
		return r
	}
	if err := safeAssert(c.Assert); err != nil {
		r.Status = Fail
		r.Detail = err.Error()
		return r
	}
	r.Status = Pass
	return r
}

func safeAssert(fn func() error) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic: %v", p)
		}
	}()
	return fn()
}

// Summary 汇总结果计数。
type Summary struct {
	Pass, Fail, Pending, Missing int
}

// OK 表示没有失败项（Pending 不算失败）。
func (s Summary) OK() bool { return s.Fail == 0 && s.Missing == 0 }

// String 返回一行汇总。
func (s Summary) String() string {
	return fmt.Sprintf("pass=%d fail=%d pending=%d missing=%d", s.Pass, s.Fail, s.Pending, s.Missing)
}

// Summarize 统计结果集。
func Summarize(rs []Result) Summary {
	var s Summary
	for _, r := range rs {
		switch r.Status {
		case Pass:
			s.Pass++
		case Fail:
			s.Fail++
		case Pending:
			s.Pending++
		case Missing:
			s.Missing++
		}
	}
	return s
}

// Validate 检查契约表自身的一致性：
//   - 编号唯一且形如 C<数字>；
//   - 覆盖 C1–C28 无缺号；
//   - 每条必须标注 Owner 与 Gate；
//   - 已到 Current 的契约必须有断言（否则就是"文档写了、代码没管"）。
func Validate() error {
	if len(table) != 28 {
		return fmt.Errorf("契约数量 %d，应为 28（C1–C28）", len(table))
	}
	seen := make(map[string]bool, len(table))
	for _, c := range table {
		if _, dup := seen[c.ID]; dup {
			return fmt.Errorf("契约编号重复: %s", c.ID)
		}
		seen[c.ID] = true
		if !strings.HasPrefix(c.ID, "C") {
			return fmt.Errorf("契约编号格式错误: %s", c.ID)
		}
		if c.Title == "" {
			return fmt.Errorf("%s 缺少标题", c.ID)
		}
		if c.Owner == "" {
			return fmt.Errorf("%s 缺少归属层", c.ID)
		}
		if c.Gate == "" {
			return fmt.Errorf("%s 缺少验证命令", c.ID)
		}
		if c.Since > M5 {
			return fmt.Errorf("%s 的 Since 越界: %v", c.ID, c.Since)
		}
		if c.Assert == nil && c.Since <= Current {
			return fmt.Errorf("%s 已到 %s 但未实现断言", c.ID, c.Since)
		}
	}
	for i := 1; i <= 28; i++ {
		id := fmt.Sprintf("C%d", i)
		if !seen[id] {
			return fmt.Errorf("缺少契约 %s", id)
		}
	}
	return nil
}

// lessID 按编号数字排序（C2 在 C10 之前）。
func lessID(a, b string) bool {
	na, nb := idNum(a), idNum(b)
	if na != nb {
		return na < nb
	}
	return a < b
}

func idNum(id string) int {
	n := 0
	for _, ch := range id {
		if ch >= '0' && ch <= '9' {
			n = n*10 + int(ch-'0')
		}
	}
	return n
}

// WGRoot 返回 wg 包根目录。
func WGRoot() string {
	_, file, _, _ := runtime.Caller(0)
	return deps.WGRoot(file, 1)
}

// checkDeps 是 C17 的断言实现。
func checkDeps() error {
	v, err := deps.Check(WGRoot())
	if err != nil {
		return err
	}
	if len(v) == 0 {
		return nil
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%d 处违规:\n", len(v))
	for i, item := range v {
		if i >= 8 {
			fmt.Fprintf(&b, "  ... 另有 %d 处\n", len(v)-i)
			break
		}
		b.WriteString("  " + item.String() + "\n")
	}
	return fmt.Errorf("%s", strings.TrimRight(b.String(), "\n"))
}
