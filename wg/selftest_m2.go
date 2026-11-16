package wg

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
)

// selfTestM2 追加 M2 的容器级自检项：内容树接线、浮层集成、静止零帧。
//
// 为什么这些不能只放在各包的单元测试里：它们测的是**接线**——
// 树挂上了没有、浮层在不在命中路径上、静止时到底还出不出帧。
// 每一件单独看都"各包全绿"，接错线之后截图也是对的（第一帧就是对的），
// 只有连续出帧的统计能暴露问题。
func selfTestM2(add func(group, name string, fn func() error)) {
	// ── 内容树 ───────────────────────────────────────────────────────
	add("container", "挂载 1000 节点：首帧提交、次帧静止", func() error {
		c, err := m2Container(400, 300)
		if err != nil {
			return err
		}
		defer c.Close()
		c.Mount(thousandNodes(40))

		st0, err := c.FrameOnce()
		if err != nil {
			return err
		}
		if !st0.Presented {
			return fmt.Errorf("挂载后的首帧没有提交（nodes=%d）", st0.Nodes)
		}
		if st0.Nodes < 1000 {
			return fmt.Errorf("统计到的节点数 = %d", st0.Nodes)
		}
		st1, err := c.FrameOnce()
		if err != nil {
			return err
		}
		if st1.Presented {
			return fmt.Errorf("无变化却再次提交（damage=%v）", st1.Damage)
		}
		return nil
	})
	add("container", "改一个节点只重画它的包围盒", func() error {
		c, err := m2Container(400, 300)
		if err != nil {
			return err
		}
		defer c.Close()
		grid := thousandNodes(40)
		c.Mount(grid)
		if _, err := c.FrameOnce(); err != nil {
			return err
		}

		target := grid.Children()[7]
		target.SetOpacity(0.5) // 任意一次视觉变更
		st, err := c.FrameOnce()
		if err != nil {
			return err
		}
		if !st.Presented {
			return fmt.Errorf("改动后没有提交")
		}
		full := float32(400 * 300)
		area := st.Damage.W() * st.Damage.H()
		if area <= 0 {
			return fmt.Errorf("脏区为空")
		}
		// 改一个 8×8 的方块，脏区不该超过整屏的 1/4。
		if area > full*0.25 {
			return fmt.Errorf("脏区 %.0f 占整屏 %.0f 的 %.0f%%（应只重画该节点）",
				area, full, area/full*100)
		}
		return nil
	})
	add("container", "1000 节点静止后 0 帧", func() error {
		c, err := m2Container(400, 300)
		if err != nil {
			return err
		}
		defer c.Close()
		c.Mount(thousandNodes(40))

		clk := frame.NewManualClock(0, 16*time.Millisecond)
		tk := c.RunManual(clk)
		if tk == nil {
			return fmt.Errorf("手动帧泵未启动")
		}
		// 跑到静止：最多 8 帧。
		for i := 0; i < 8 && tk.Armed(); i++ {
			tk.Advance(16 * time.Millisecond)
		}
		if tk.Armed() {
			return fmt.Errorf("8 帧后仍有待触发的帧（静止判定失效）")
		}
		settled := c.Frames()
		// 再推进 300ms 的墙钟：不该产生任何一帧。
		for i := 0; i < 19; i++ {
			tk.Advance(16 * time.Millisecond)
		}
		if got := c.Frames(); got != settled {
			return fmt.Errorf("静止后又出了 %d 帧", got-settled)
		}
		return nil
	})
	add("container", "换主题触发整树重排", func() error {
		c, err := m2Container(400, 300)
		if err != nil {
			return err
		}
		defer c.Close()
		c.Mount(thousandNodes(40))
		if _, err := c.FrameOnce(); err != nil {
			return err
		}
		c.SetTheme(theme.Dark())
		st, err := c.FrameOnce()
		if err != nil {
			return err
		}
		if !st.Presented {
			return fmt.Errorf("换主题后没有重画")
		}
		if got := c.Theme(); got != theme.Dark() && got.Color.Background != theme.Dark().Color.Background {
			return fmt.Errorf("主题未生效")
		}
		return nil
	})

	// ── 浮层集成 ─────────────────────────────────────────────────────
	add("container", "浮层参与命中：点遮罩不透传到内容树", func() error {
		c, err := m2Container(400, 300)
		if err != nil {
			return err
		}
		defer c.Close()
		var taps int
		btn := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(0, 0, 0))})
		btn.SetSize2(80, 30)
		btn.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "btn"})
		root := scene.Stack(btn).SetSize2(400, 300)
		c.Mount(root)
		c.SetHandler(btn, input.Handler{Pointer: func(*input.PointerEvent) bool { taps++; return true }})
		if _, err := c.FrameOnce(); err != nil {
			return err
		}

		// 先确认没有浮层时点得到按钮。
		if !c.HandlePointer(m2Down(10, 10)) {
			return fmt.Errorf("未弹浮层时按钮点不到（命中路径没接上）")
		}
		if taps != 1 {
			return fmt.Errorf("按钮回调 %d 次", taps)
		}
		// 弹一个模态：遮罩吞掉同一处点击。
		c.Overlay().Show(m2Dialog())
		c.HandlePointer(m2Down(10, 10))
		if taps != 1 {
			return fmt.Errorf("模态遮罩未拦住内容树的点击（回调涨到 %d）", taps)
		}
		return nil
	})
	add("container", "ShowToast 落 BandToast 且无尺寸时不弹", func() error {
		c := NewHeadless(DefaultOptions())
		if id := c.ShowToast("还没尺寸", 0); id != 0 {
			return fmt.Errorf("未就绪时返回了句柄 %d", id)
		}
		c.Resize(400, 300)
		id := c.ShowToast("已保存", 0)
		if id == 0 {
			return fmt.Errorf("就绪后 ShowToast 返回 0")
		}
		es := c.Overlay().Entries()
		if len(es) != 1 || es[0].Band != theme.BandToast {
			return fmt.Errorf("浮层集合 = %d 条，层带 = %v", len(es), es[0].Band)
		}
		if _, err := c.FrameOnce(); err != nil {
			return err
		}
		// 到时自动消失（时钟推进由 TickTime 负责）。
		c.TickTime(time.Now().Add(4 * time.Second))
		if c.Overlay().Count() != 0 {
			return fmt.Errorf("Toast 到时未消失")
		}
		return nil
	})
	add("container", "浮层变化会请求下一帧", func() error {
		c, err := m2Container(400, 300)
		if err != nil {
			return err
		}
		defer c.Close()
		c.Mount(thousandNodes(4))
		if _, err := c.FrameOnce(); err != nil {
			return err
		}
		st, err := c.FrameOnce()
		if err != nil {
			return err
		}
		if st.Presented {
			return fmt.Errorf("静止帧不该提交")
		}
		c.Overlay().Show(m2Dialog())
		st2, err := c.FrameOnce()
		if err != nil {
			return err
		}
		if !st2.Presented {
			return fmt.Errorf("弹出浮层后没有重画")
		}
		return nil
	})
}

// m2Container 造一个已 Resize 的无窗口容器。
func m2Container(w, h int) (*Container, error) {
	c := NewHeadless(DefaultOptions())
	c.Resize(w, h)
	if _, err := c.FrameOnce(); err != nil {
		c.Close()
		return nil, err
	}
	return c, nil
}

// thousandNodes 造 n×n 个 8px 方块（默认 40 列 = 1600 个）。
//
// 用 Grid 而不是 Stack + 绝对定位：Stack 的定位子节点会被紧约束拉满，
// 1625 个全屏方块叠在一起，脏区断言就变成"整屏"了。
func thousandNodes(columns int) *scene.Node {
	if columns < 1 {
		columns = 1
	}
	children := make([]*scene.Node, 0, columns*columns)
	for i := 0; i < columns*columns; i++ {
		b := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(200, 200, 200))})
		b.SetSize2(8, 8)
		children = append(children, b)
	}
	return scene.Grid(columns, 1, children...).SetSize2(float32(columns*9), float32(columns*9))
}

// m2Dialog 造一个居中的模态对话框（遮罩吞点击）。
func m2Dialog() overlay.Entry {
	return overlay.Entry{
		Band:      theme.BandModal,
		Placement: overlay.PlacementCenter,
		Size:      geom.Sz(200, 120),
		Modal:     true,
		Build: func() *scene.Node {
			n := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(255, 255, 255))})
			n.SetSize2(200, 120)
			return n
		},
		Tag: "dialog",
	}
}

// m2Down 造按下事件。
func m2Down(x, y float32) *input.PointerEvent {
	p := geom.Pt(x, y)
	return &input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Device: input.DeviceMouse, Pos: p, PosRoot: p, Time: time.Now()}
}
