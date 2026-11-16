package launcher

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/examples/doc"
)

// SelfCheckReport 是 --selfcheck 的 JSON 摘要，也是枚举门禁的判定输入。
type SelfCheckReport struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Group     string `json:"group"`
	Demos     int    `json:"demos"`
	APIRows   int    `json:"apiRows"`
	Presented bool   `json:"presented"`
	Nodes     int    `json:"nodes"`
	Degraded  int    `json:"degraded"`
	NonBg     int    `json:"nonBg"`
	OK        bool   `json:"ok"`
	Reason    string `json:"reason,omitempty"`
}

// selfcheck 以 headless 方式出帧并判定：文档页必须真的画出像素、
// 不能有任何静默降级。打印 SelfCheckReport JSON，返回进程退出码。
//
// 这条路径不访问任何 LCL 符号（NewHeadless 无窗口），因此 go test 可以
// 以子进程方式枚举全部功能窗——真窗的像素证据仍由 --shot 提供。
func selfcheck(meta doc.Meta, o options) int {
	rep := SelfCheckReport{
		ID: meta.ID, Title: meta.Title, Group: meta.Group,
		Demos: len(meta.Demos), APIRows: len(meta.API),
	}
	fail := func(reason string) int {
		rep.OK, rep.Reason = false, reason
		printReport(rep)
		return 1
	}
	if len(meta.Demos) == 0 {
		return fail("功能窗至少要有一个演示块")
	}

	c := wg.NewHeadless(wg.DefaultOptions())
	c.Resize(o.w, o.h)
	c.MountWidget(doc.Page(meta))
	st, err := c.FrameOnce()
	if err != nil {
		return fail("出帧失败: " + err.Error())
	}
	rep.Presented = st.Presented
	rep.Nodes = st.Nodes
	rep.Degraded = st.Degraded
	if !st.Presented {
		return fail("首帧未提交（文档页没有可画内容）")
	}
	if st.Degraded > 0 {
		return fail(fmt.Sprintf("文档页有 %d 条指令被降级（能力未实现或后端不支持）", st.Degraded))
	}

	// 像素证据：合成缓冲按 4px 采样，非白像素必须足量（防止整页空白）。
	b := c.Backing()
	if b == nil {
		return fail("合成缓冲不可用")
	}
	px := b.Pix()
	nonBg := 0
	for i := 0; i < len(px); i += 16 { // 每 4 像素采 1 个
		if px[i] != 0xFFFFFFFF {
			nonBg++
		}
	}
	rep.NonBg = nonBg
	if nonBg < 100 {
		return fail(fmt.Sprintf("非背景采样点 %d < 100（页面几乎空白）", nonBg))
	}
	rep.OK = true
	printReport(rep)
	return 0
}

func printReport(rep SelfCheckReport) {
	out, err := json.Marshal(rep)
	if err != nil {
		fmt.Fprintln(os.Stderr, "selfcheck 序列化失败:", err)
		return
	}
	fmt.Println(string(out))
}
