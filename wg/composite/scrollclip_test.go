package composite

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
)

// TestScrollClipsContent 锁定滚动视口的绘制裁剪：平移出视口的内容不得
// 溢出绘制。此前 KindScroll 只平移不裁剪，溢出内容会盖住视口外的其它
// UI（scroll 示例真窗可复现：行内容画满了整个窗口）。
func TestScrollClipsContent(t *testing.T) {
	c := newComp(t, 40, 200)

	// 内容 30×200，视口 40×100：下半 100px 内容必须被裁掉。
	box := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(255, 0, 0))})
	box.SetSize2(30, 200)
	view := scene.Scroll(scene.Vertical, box)
	view.SetSize2(40, 100)
	root := scene.Stack()
	root.Add(view)
	root.Layout(scene.TightConstraints(geom.Sz(40, 200)))

	st, err := c.DrawTree(Build(root, BuildOptions{}))
	if err != nil {
		t.Fatal(err)
	}
	if !st.Presented {
		t.Fatal("首帧应提交")
	}
	pix := c.Backing().Pix()
	at := func(x, y int) uint32 { return pix[y*40+x] }
	// 视口内（y=50）应是红（B 通道为 0）；视口外（y=150）必须是背景白。
	if p := at(10, 50); (p>>16)&0xFF != 255 || p&0xFF != 0 {
		t.Fatalf("视口内像素 = %#08x，应为红", p)
	}
	if p := at(10, 150); (p>>16)&0xFF == 255 && p&0xFF == 0 {
		t.Fatalf("视口外像素 = %#08x，内容溢出了（应为背景白）", p)
	}
}
