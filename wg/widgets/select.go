package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── Select：下拉选择器 ────────────────────────────────────────────

// Select 是单选下拉框。Options 为选项文本，Selected 为当前索引（-1 未选）。
//
// 展开的列表交给 overlay.Host 承载（与 Tooltip 同一条浮层链路）：
// 定位、OutsideClick/Esc 关闭、z 序都由浮层体系负责，组件只声明内容。
type Select struct {
	controlBase
	Options   []string
	Selected  int
	Width     float32
	MaxHeight float32
	Placeholder string
	Disabled  bool
	OnChange  func(int)

	handle overlay.Handle
	taps   []*ui.GestureDetector // 跨 Build 存活
}

// NewSelect 创建下拉框。
func NewSelect(options []string) *Select {
	return &Select{Options: options, Selected: -1, Width: 180, MaxHeight: 240, Placeholder: "请选择"}
}

// SetSelected 设置选中项并回调。
func (s *Select) SetSelected(i int) {
	if i < -1 || i >= len(s.Options) || i == s.Selected {
		return
	}
	s.Selected = i
	s.MarkDirty()
	if s.OnChange != nil {
		s.OnChange(i)
	}
}

// Close 关闭下拉（若已展开）。
func (s *Select) Close() {
	if s.handle != 0 {
		// handle 的关闭由最近一次 Build 的 ctx.Overlay 执行；这里只清标记，
		// 浮层的 OutsideClick/Esc 也会自行关闭。
		s.handle = 0
	}
}

// Build 实现 ui.Widget。
func (s *Select) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	h := float32(34)
	n := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.Surface), Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	}).SetSize2(s.Width, h)

	caption := s.Placeholder
	if s.Selected >= 0 && s.Selected < len(s.Options) {
		caption = s.Options[s.Selected]
	}
	fg := th.Color.Text
	if s.Selected < 0 {
		fg = th.Color.TextFaint
	}
	l := text.LayoutText(caption, text.DefaultFont(th.Font.SizeSM), text.Options{})
	tv := scene.Text(scene.TextSpec{Src: caption, Layout: l, Color: fg})
	tv.SetMargin(geom.InLTRB(10, (h-l.Size.H)/2, 0, 0))
	n.Add(tv)
	al := text.LayoutText("▾", text.DefaultFont(th.Font.SizeSM), text.Options{})
	at := scene.Text(scene.TextSpec{Src: "▾", Layout: al, Color: th.Color.TextMuted})
	at.SetPositionKeepSize(true).SetStackPosition(geom.Pt(s.Width-20, (h-al.Size.H)/2))
	n.Add(at)

	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "select", Cursor: scene.CursorPointer})
	g := &ui.GestureDetector{OnTap: func(geom.Point) {
		if !s.Disabled {
			s.show(ctx, n)
		}
	}}
	s.taps = append(s.taps[:0], g)
	s.attach(ctx, n, g.Handler(), false)
	s.node = n
	return n
}

func (s *Select) show(ctx *ui.BuildContext, anchor *scene.Node) {
	if ctx.Overlay == nil || s.handle != 0 || s.Disabled {
		return
	}
	th := ctx.ThemeOr()
	mh := s.MaxHeight
	if mh <= 0 {
		mh = 240
	}
	listH := float32(len(s.Options)) * 32
	if listH > mh {
		listH = mh
	}
	s.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: s.Width, H: listH},
		Dismiss:   overlay.Dismiss(0), // OutsideClick/Esc 关闭，不吞点击
		Tag:       "select",
		Build: func() *scene.Node {
			col := scene.Flex(scene.Vertical, 0).SetWidth(s.Width)
			col.SetOverflow(true)
			for i, opt := range s.Options {
				item := scene.Box(scene.BoxSpec{
					Fill: scene.Solid(th.Color.Surface),
				}).SetWidth(s.Width).SetHeight(32)
				fg := th.Color.Text
				if i == s.Selected {
					item.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Selection)})
					fg = th.Color.Primary
				}
				l := text.LayoutText(opt, text.DefaultFont(th.Font.SizeSM), text.Options{})
				t := scene.Text(scene.TextSpec{Src: opt, Layout: l, Color: fg})
				t.SetMargin(geom.InLTRB(10, (32-l.Size.H)/2, 0, 0))
				item.Add(t)
				item.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "select-option", Cursor: scene.CursorPointer})
				idx := i
				og := &ui.GestureDetector{OnTap: func(geom.Point) {
					s.SetSelected(idx)
					if ctx.Overlay != nil {
						ctx.Overlay.Close(s.handle)
					}
					s.handle = 0
				}}
				s.taps = append(s.taps, og)
				if ctx.Input != nil {
					ctx.Input.Set(item, og.Handler())
				}
				col.Add(item)
			}
			return col
		},
	})
}

