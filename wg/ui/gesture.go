package ui

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
)

// ── 声明式手势 ──────────────────────────────────────────────────────
//
// 组件不该自己写"按下 → 判断位移 → 判断时长 → 互斥"这套逻辑：
// 写十次就有十份不一致的 slop，而且互斥必然出错（滑块 vs 横向滚动）。
// GestureDetector 把竞技场包成声明式接口，组件只声明"我要什么手势"。

// GestureDetector 是声明式手势集合。
//
// 零值可用；填了哪个回调就注册哪个识别器（没填的不占竞技场名额，
// 少一个竞争者对裁决稳定性有实际好处）。
type GestureDetector struct {
	// OnTap 单击。
	OnTap func(pos geom.Point)
	// OnTapEnd 点击结束（accepted 表示是否胜出，用于回滚按下高亮）。
	OnTapEnd func(accepted bool)
	// OnLongPress 长按。
	OnLongPress func(pos geom.Point)
	// OnPanStart 拖动开始（已越过 slop 并胜出）。
	OnPanStart func(pos geom.Point)
	// OnPan 拖动中（delta 为相对上一次的增量）。
	OnPan func(pos, delta geom.Point)
	// OnPanEnd 拖动结束（velocity 为释放速度 DIP/s）。
	OnPanEnd func(pos, velocity geom.Point)

	// Axis 是拖动主轴锁（Horizontal/Vertical 时只认该方向）。
	Axis input.Axis
	// Config 是识别参数（零值用 DefaultGestureConfig）。
	Config input.GestureConfig

	// 运行时状态
	arena   *input.Arena
	long    input.Tickable
	members []input.Recognizer
	active  bool
}

// cfg 返回参数（补默认值）。
func (g *GestureDetector) cfg() input.GestureConfig {
	c := g.Config
	d := input.DefaultGestureConfig()
	if c.Slop <= 0 {
		c.Slop = d.Slop
	}
	if c.LongPress <= 0 {
		c.LongPress = d.LongPress
	}
	if c.MaxTapDelay <= 0 {
		c.MaxTapDelay = d.MaxTapDelay
	}
	c.AxisLock = g.Axis
	return c
}

// Handler 返回可挂到 input.Dispatcher 的处理器。
//
// 指针按下时创建竞技场并让本次配置的所有识别器入场；抬起后清场。
// 竞技场是**每指针每次按下**一个：跨按下复用会把上次的败者状态带进来。
func (g *GestureDetector) Handler() input.Handler {
	return input.Handler{Pointer: g.handle}
}

// handle 处理指针事件。
func (g *GestureDetector) handle(ev *input.PointerEvent) bool {
	if ev == nil {
		return false
	}
	switch ev.Kind {
	case input.PointerDown:
		g.disarm()
		g.arm(ev)
		return true
	case input.PointerMove:
		if !g.active {
			return false
		}
		for _, m := range g.members {
			m.Handle(ev)
		}
		return true
	case input.PointerUp, input.PointerCancel:
		if !g.active {
			return false
		}
		for _, m := range g.members {
			m.Handle(ev)
		}
		if ev.Kind == input.PointerUp && g.arena.State() != input.ArenaResolved {
			// 抬起时无人宣布：由竞技场裁决（第一个仍在场者胜）。
			g.arena.Sweep()
		}
		if ev.Kind == input.PointerCancel && g.arena.State() != input.ArenaResolved {
			g.arena.RejectAll()
		}
		g.disarm()
		return true
	}
	return false
}

// arm 为本次按下建场。
func (g *GestureDetector) arm(ev *input.PointerEvent) {
	a := input.NewArena()
	cfg := g.cfg()
	var members []input.Recognizer
	if g.OnTap != nil || g.OnTapEnd != nil {
		members = append(members, input.Tap(a, cfg, g.OnTap, g.OnTapEnd))
	}
	if g.OnLongPress != nil {
		lp := input.LongPress(a, cfg, g.OnLongPress)
		g.long = lp
		members = append(members, lp)
	} else {
		g.long = nil
	}
	if g.OnPan != nil || g.OnPanStart != nil || g.OnPanEnd != nil {
		members = append(members, input.Drag(a, cfg, g.OnPanStart, g.OnPan, g.OnPanEnd))
	}
	for _, m := range members {
		a.AddNamed(m, "ui")
	}
	g.arena = a
	g.members = members
	g.active = len(members) > 0
	if g.active {
		for _, m := range members {
			m.AddPointer(ev)
		}
		a.Close()
	}
}

// disarm 清场。
func (g *GestureDetector) disarm() {
	g.active = false
	g.members = nil
	g.arena = nil
	g.long = nil
}

// Tick 推进长按计时（宿主每帧调用）。
//
// 只在按住且注册了长按时有意义；返回 true 表示本次判定出了长按。
func (g *GestureDetector) Tick(now time.Time) bool {
	if g.long == nil {
		return false
	}
	return g.long.Tick(now)
}

// Active 表示当前有进行中的手势。
func (g *GestureDetector) Active() bool { return g.active }

// State 返回竞技场状态（诊断用）。
func (g *GestureDetector) State() input.ArenaState {
	if g.arena == nil {
		return input.ArenaResolved
	}
	return g.arena.State()
}

// Cursor 返回建议光标（有拖动时用 Move）。
func (g *GestureDetector) Cursor() scene.Cursor {
	if g.active && g.OnPan != nil {
		return scene.CursorMove
	}
	if g.OnTap != nil {
		return scene.CursorPointer
	}
	return scene.CursorDefault
}
