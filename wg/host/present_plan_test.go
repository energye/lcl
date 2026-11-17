package host

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
)

// TestExactPresentPlan 锁定 resize 过渡几何：内容 1:1，永不缩放。
func TestExactPresentPlan(t *testing.T) {
	// 尺寸吻合：全幅 1:1，不填色。
	p := exactPresentPlan(400, 300, 400, 300)
	if p.Copy != geom.RectXYWH(0, 0, 400, 300) || len(p.Fill) != 0 {
		t.Fatalf("吻合: copy=%v fill=%v", p.Copy, p.Fill)
	}

	// 变宽：拷贝保持原尺寸，右侧条带填色。
	p = exactPresentPlan(400, 300, 520, 300)
	if p.Copy != geom.RectXYWH(0, 0, 400, 300) {
		t.Fatalf("变宽 copy=%v（应 1:1，不缩放）", p.Copy)
	}
	if len(p.Fill) != 1 || p.Fill[0] != geom.RectXYWH(400, 0, 120, 300) {
		t.Fatalf("变宽 fill=%v", p.Fill)
	}

	// 变高：底部条带填色。
	p = exactPresentPlan(400, 300, 400, 380)
	if len(p.Fill) != 1 || p.Fill[0] != geom.RectXYWH(0, 300, 400, 80) {
		t.Fatalf("变高 fill=%v", p.Fill)
	}

	// 双向变大：右 + 下两条（右下角只填一次，条带互不重叠）。
	p = exactPresentPlan(400, 300, 520, 380)
	if p.Copy != geom.RectXYWH(0, 0, 400, 300) || len(p.Fill) != 2 {
		t.Fatalf("双向 copy=%v fill=%v", p.Copy, p.Fill)
	}

	// 缩小：只画交集（裁剪），不填色、不缩放。
	p = exactPresentPlan(400, 300, 300, 240)
	if p.Copy != geom.RectXYWH(0, 0, 300, 240) || len(p.Fill) != 0 {
		t.Fatalf("缩小 copy=%v fill=%v", p.Copy, p.Fill)
	}

	// 只缩窄：交集宽=客户区、高=帧高。
	p = exactPresentPlan(400, 300, 300, 300)
	if p.Copy != geom.RectXYWH(0, 0, 300, 300) || len(p.Fill) != 0 {
		t.Fatalf("缩窄 copy=%v fill=%v", p.Copy, p.Fill)
	}

	// 非法客户区：空方案。
	if p = exactPresentPlan(400, 300, 0, 0); !p.Copy.IsEmpty() || len(p.Fill) != 0 {
		t.Fatalf("零客户区 copy=%v fill=%v", p.Copy, p.Fill)
	}
}
