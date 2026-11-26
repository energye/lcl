package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/scene"
)

// findLabelNode 深度优先找命中锚为 label 的节点。
func findLabelNode(n *scene.Node, label string) *scene.Node {
	if n == nil {
		return nil
	}
	if n.HitSpec().Label == label {
		return n
	}
	for _, c := range n.Children() {
		if got := findLabelNode(c, label); got != nil {
			return got
		}
	}
	return nil
}

// TestWatermarkMultiline 多行水印：单行回退、多行优先、水印层锚存在。
func TestWatermarkMultiline(t *testing.T) {
	// 单行回退。
	w1 := NewWatermark(nil, "单行")
	w1.Width, w1.Height = 200, 100
	n1 := w1.Build(dataCtx())
	if findLabelNode(n1, "watermark-layer") == nil {
		t.Fatal("单行水印缺少 watermark-layer")
	}

	// 多行：Contents 优先（Content 被忽略），水印层仍存在。
	w2 := NewWatermark(nil, "被忽略")
	w2.Contents = []string{"第一行", "第二行", "第三行"}
	w2.Width, w2.Height = 200, 100
	n2 := w2.Build(dataCtx())
	if findLabelNode(n2, "watermark-layer") == nil {
		t.Fatal("多行水印缺少 watermark-layer")
	}

	// 空 Contents 回退到 Content。
	w3 := NewWatermark(nil, "仅单行")
	w3.Width, w3.Height = 200, 100
	if n := w3.Build(dataCtx()); findLabelNode(n, "watermark-layer") == nil {
		t.Fatal("空 Contents 回退单行失败")
	}
}
