package contracts

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
)

// 本文件是 M2 到期契约的断言实现。
//
// M2 的契约几乎全是"交互状态机"性质的：捕获会不会漏 Up、竞技场会不会
// 出现双胜者、焦点会不会跑到模态外面。这类断言不需要像素，只需要
// 把事件序列喂进去看状态机是否守约——所以它们能跑在无窗口自检里。

// m2Clock 是 M2 断言用的基准时刻（非零，避免被当成"没有时间戳"）。
func m2Clock(ms int) time.Time {
	return time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC).Add(time.Duration(ms) * time.Millisecond)
}

// m2Down 造按下事件（Move/Up 由 m2Kind 派生，见 C26 断言）。
func m2Down(x, y float32, ms int) *input.PointerEvent {
	p := geom.Pt(x, y)
	return &input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Device: input.DeviceMouse, Pos: p, PosRoot: p, Time: m2Clock(ms)}
}

// arenaProbe 是只记账不宣布的竞技场成员。
type arenaProbe struct {
	accept int
	reject int
	events int
	state  int // 模拟组件视觉状态：1 = 按下高亮
}

func (p *arenaProbe) AddPointer(*input.PointerEvent) { p.state = 1 }
func (p *arenaProbe) Handle(*input.PointerEvent)     { p.events++ }
func (p *arenaProbe) Accept()                        { p.accept++ }
func (p *arenaProbe) Reject()                        { p.reject++; p.state = 0 }

// ── C27 单指针单胜者，败者 Reject 且不改状态 ─────────────────────

func assertSingleWinner() error {
	a := input.NewArena()
	first, second, third := &arenaProbe{}, &arenaProbe{}, &arenaProbe{}
	a.AddNamed(first, "first")
	a.AddNamed(second, "second")
	a.AddNamed(third, "third")
	a.Close()

	// 第二个成员抢先宣布：它必须是唯一胜者，其余全部被拒。
	if a.Resolve(second) != true {
		return fmt.Errorf("首次宣布未被接受")
	}
	if a.Winner() != input.Recognizer(second) {
		return fmt.Errorf("胜者不是宣布者")
	}
	if second.accept != 1 {
		return fmt.Errorf("胜者 Accept 次数 = %d", second.accept)
	}
	if second.reject != 0 {
		return fmt.Errorf("胜者被拒绝 %d 次", second.reject)
	}

	// 败者：各被拒一次，且视觉状态被回滚（state 必须归零）。
	for i, p := range []*arenaProbe{first, third} {
		if p.reject != 1 {
			return fmt.Errorf("败者 #%d Reject 次数 = %d", i, p.reject)
		}
		if p.state != 0 {
			return fmt.Errorf("败者 #%d 未回滚视觉状态", i)
		}
		if p.accept != 0 {
			return fmt.Errorf("败者 #%d 收到 Accept", i)
		}
	}

	// 再宣布一次必须失败，且胜者不被改写（单指针单胜者）。
	if a.Resolve(third) {
		return fmt.Errorf("第二个宣布者被接受，出现双胜者")
	}
	if a.Winner() != input.Recognizer(second) {
		return fmt.Errorf("重复宣布改写了胜者")
	}

	// 已裁决后再 sweep 不得重开、不得二次通知。
	if got := a.Sweep(); got != input.Recognizer(second) {
		return fmt.Errorf("sweep 返回了别的成员")
	}
	if third.reject != 1 || second.accept != 1 {
		return fmt.Errorf("重复裁决：third.reject=%d second.accept=%d", third.reject, second.accept)
	}
	return nil
}
