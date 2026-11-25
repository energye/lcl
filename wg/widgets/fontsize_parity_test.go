package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// findTextBySrc 在子树里找 Src 完全匹配的文本节点。
func findTextBySrc(n *scene.Node, src string) *scene.Node {
	if n == nil {
		return nil
	}
	if n.Kind() == scene.KindText {
		if sp := n.Text(); sp != nil && sp.Src == src {
			return n
		}
	}
	for _, c := range n.Children() {
		if got := findTextBySrc(c, src); got != nil {
			return got
		}
	}
	return nil
}

// TestFontSizeParity 固定 antd 5.29.3 的字号语义：
// 正文一律 fontSize=14，只有图标/角标/标签用 fontSizeSM=12。
// 主题 token 早在 M7 已对齐（14/16/20/24），本轮修的是组件层误用 SizeSM。
func TestFontSizeParity(t *testing.T) {
	th := theme.Default()
	if th.Font.SizeBase != 14 || th.Font.SizeLG != 16 || th.Font.SizeSM != 12 {
		t.Fatalf("主题字号 token 偏离 antd: Base=%v LG=%v SM=%v",
			th.Font.SizeBase, th.Font.SizeLG, th.Font.SizeSM)
	}
	ctx := &ui.BuildContext{Theme: th}

	assertText := func(name string, root *scene.Node, src string, want float32) {
		t.Helper()
		n := findTextBySrc(root, src)
		if n == nil {
			t.Fatalf("%s: 找不到文本 %q", name, src)
		}
		sp := n.Text()
		got := sp.LayoutFont.Size
		if got == 0 && sp.Layout != nil {
			got = sp.Layout.Font.Size
		}
		if got != want {
			t.Fatalf("%s: %q 字号=%v，want %v", name, src, got, want)
		}
	}

	// 数据展示：Descriptions label/content、Statistic title 均继承 fontSize(14)。
	desc := &Descriptions{Items: []DescriptionItem{{Label: "用户名", Content: "张三"}}}
	assertText("Descriptions label", desc.Build(ctx), "用户名:", th.Font.SizeBase)
	assertText("Descriptions content", desc.Build(ctx), "张三", th.Font.SizeBase)

	stat := &Statistic{Title: "总销售额", Value: 1234}
	assertText("Statistic title", stat.Build(ctx), "总销售额", th.Font.SizeBase)

	// Steps 标题是 fontSizeLG(16)，不是正文。
	steps := Steps{Labels: []string{"已完成"}}
	assertText("Steps title", steps.Build(ctx), "已完成", th.Font.SizeLG)

	// Tag 是 tagFontSize=fontSizeSM(12)——唯一的合法 12 场景。
	tag := &Tag{Text: "标签"}
	assertText("Tag", tag.Build(ctx), "标签", th.Font.SizeSM)

	// small 控件正文字号与 medium 一致（antd contentFontSizeSM=fontSize）。
	small := ui.MetricsFor(th, ui.SizeSmall)
	medium := ui.MetricsFor(th, ui.SizeMedium)
	if small.Font != th.Font.SizeBase || medium.Font != th.Font.SizeBase {
		t.Fatalf("控件正文字号 small=%v medium=%v，want 均为 %v",
			small.Font, medium.Font, th.Font.SizeBase)
	}

	// Progress 普通档文字继承 fontSize(14)；Switch 内部标签是 fontSizeSM(12)。
	prog := &Progress{Value: 0.5, ShowInfo: true}
	assertText("Progress text", prog.Build(ctx), "50%", th.Font.SizeBase)

	sw := &Switch{On: true, CheckedChildren: "开"}
	assertText("Switch inner label", sw.Build(ctx), "开", th.Font.SizeSM)
}
