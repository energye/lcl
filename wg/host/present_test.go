package host

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
)

// Presenter 把"合成"与"提交到窗口"解耦：Frame 记账、Paint 提交。
// 这组测试在无窗口环境验证记账链路，确保帧不会漏记/重复记。
func TestPresenterFrameAccounting(t *testing.T) {
	p := NewPresenter(0, nil)
	if p == nil {
		t.Fatal("Presenter 为 nil")
	}
	if s := p.Stats(); s.Buffered || s.Frames != 0 || s.Painted != 0 {
		t.Fatalf("初始统计错误: %+v", s)
	}

	b := raster.NewBacking(8, 6)
	if b == nil {
		t.Fatal("Backing 分配失败")
	}
	dmg := geom.RectXYWH(1, 2, 4, 3)
	p.Frame(b, dmg)

	s := p.Stats()
	if !s.Buffered || s.Frames != 1 || s.Painted != 0 {
		t.Fatalf("Frame 后统计错误: %+v", s)
	}
	if s.LastDamage != dmg {
		t.Fatalf("脏区 = %v，应为 %v", s.LastDamage, dmg)
	}

	// headless（hwnd=0）下 Paint 不提交，Painted 不应增加。
	if p.Paint() {
		t.Fatal("headless Presenter 不应报告提交成功")
	}
	if s = p.Stats(); s.Painted != 0 {
		t.Fatalf("headless Paint 后 Painted = %d", s.Painted)
	}

	p.Close()
	if s = p.Stats(); s.Buffered {
		t.Fatal("Close 后不应仍缓存帧")
	}
}

func TestPresenterFrameOverwrite(t *testing.T) {
	p := NewPresenter(0, nil)
	b := raster.NewBacking(4, 4)
	p.Frame(b, geom.RectXYWH(0, 0, 2, 2))
	p.Frame(b, geom.RectXYWH(1, 1, 2, 2))
	s := p.Stats()
	if s.Frames != 2 {
		t.Fatalf("Frames = %d，应累计为 2", s.Frames)
	}
	if s.LastDamage != geom.RectXYWH(1, 1, 2, 2) {
		t.Fatalf("LastDamage 未更新为最后一帧: %v", s.LastDamage)
	}
}
