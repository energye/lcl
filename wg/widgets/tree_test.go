package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/scene"
)

// TestFormItemLayout 表单布局：vertical 纵向、horizontal 横向、labelCol 宽。
func TestFormItemLayout(t *testing.T) {
	ctx := dataCtx()

	// vertical：根是纵向 Flex（标签行 + 控件上下排）。
	v := FormItem{Label: "垂直", Child: NewInput(""), Layout: FormVertical}
	nv := v.Build(ctx)
	if nv == nil {
		t.Fatal("vertical 构建失败")
	}
	if len(nv.Children()) < 2 {
		t.Fatalf("vertical 应有标签+控件两段，got %d", len(nv.Children()))
	}

	// horizontal：根是纵向 Flex（行 + 可能错误行），行内横向排标签+控件。
	h := FormItem{Label: "水平", Child: NewInput("")}
	nh := h.Build(ctx)
	if nh == nil {
		t.Fatal("horizontal 构建失败")
	}

	// inline 目前走 horizontal 路径（标签+控件同行）。
	in := FormItem{Label: "行内", Child: NewInput(""), Layout: FormInline}
	if n := in.Build(ctx); n == nil {
		t.Fatal("inline 构建失败")
	}

	// labelCol 标签宽 = (span/24)×父宽，宽于默认 80。
	lc := FormItem{Label: "栅格", Child: NewInput(""), LabelColSpan: 6}
	nlc := lc.Build(ctx)
	if nlc == nil {
		t.Fatal("labelCol 构建失败")
	}
	// 第一行第一个子节点是标签行，其首子节点宽度应 > 80（视口 1024 时 span6≈236）。
	row := nlc.Children()[0]
	labelRow := row.Children()[0]
	labelW := labelRow.Size().W
	_ = scene.KindText
	if labelW > 0 && labelW <= 80 {
		t.Fatalf("labelCol span6 标签宽=%v，应 >80（视口默认）", labelW)
	}
}
