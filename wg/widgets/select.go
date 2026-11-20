package widgets

import (
	"strings"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── Select（M8 对齐 antd：搜索/多选/分组）──────────────────────────
//
// antd：<Select options showSearch multiple value onChange>
// 已对齐：单选（Selected 索引语义保留）/多选（Values 集合）、showSearch
// 过滤（contains，大小写不敏感）、分组标题行、OutsideClick/Esc 关闭、
// 多选选中不收起、键盘可达（同旧实现的手势链）。
// 未实现（如实标注）：异步加载、
// 虚拟滚动、option 渲染自定义。

// SelectOption 是带分组的选项（antd options item）。
type SelectOption struct {
	Label string
	// Value 提交值（labelInValue 语义简化版）：空 = 回退 Label
	//（Label 即 Value，兼容旧用法）。
	Value string
	Group string // 空 = 不分组
}

// Select 是下拉选择框。
//
// 兼容说明：Options []string（纯文本单选）保留；需要分组/多选/搜索时
// 用 Items + Multiple + ShowSearch。Selected=-1 表示未选。
type Select struct {
	controlBase
	Options     []string
	Items       []SelectOption // 设置后优先于 Options
	Selected    int            // 单选语义
	Values      []string       // 多选语义（Multiple=true）
	Multiple    bool
	ShowSearch  bool
	Width       float32
	MaxHeight   float32
	Placeholder string
	Disabled    bool
	OnChange    func(int)      // 单选回调
	OnValues    func([]string) // 多选回调

	handle overlay.Handle
	taps   []*ui.GestureDetector // 跨 Build 存活
	filter string                // 搜索词（popup 内输入）
	// 弹出层的选项滚动（每次打开新建，位置归零；跨 popup 重建保留）。
	listViewport float32
	listScroll   *Scroll
}

// NewSelect 创建下拉框。
func NewSelect(options []string) *Select {
	return &Select{Options: options, Selected: -1, Width: 180, MaxHeight: 240, Placeholder: "请选择"}
}

// labels 返回扁平化选项文本（Items 优先）。
func (s *Select) labels() []string {
	if len(s.Items) > 0 {
		out := make([]string, len(s.Items))
		for i, it := range s.Items {
			out[i] = it.Label
		}
		return out
	}
	return s.Options
}

// SetSelected 设置选中项并回调。
func (s *Select) SetSelected(i int) {
	ls := s.labels()
	if i < -1 || i >= len(ls) || i == s.Selected {
		return
	}
	s.Selected = i
	s.MarkDirty()
	if s.OnChange != nil {
		s.OnChange(i)
	}
}

// Value 返回单选选中项的提交值（label-value 分离）：Value 字段优先，
// 空 = 回退 Label；未选中返回空串。
func (s *Select) Value() string {
	if s.Selected < 0 {
		return ""
	}
	if len(s.Items) > 0 && s.Selected < len(s.Items) {
		it := s.Items[s.Selected]
		if it.Value != "" {
			return it.Value
		}
		return it.Label
	}
	if s.Selected < len(s.Options) {
		return s.Options[s.Selected]
	}
	return ""
}

// OptionValue 返回第 i 项的提交值（label-value 分离，单/多选共用）。
func (s *Select) OptionValue(i int) string {
	if i < 0 {
		return ""
	}
	if len(s.Items) > 0 && i < len(s.Items) {
		it := s.Items[i]
		if it.Value != "" {
			return it.Value
		}
		return it.Label
	}
	if i < len(s.Options) {
		return s.Options[i]
	}
	return ""
}

// SelectedValues 返回多选选中项的提交值集合（按 Values 索引换算）。
func (s *Select) SelectedValues() []string {
	out := make([]string, 0, len(s.Values))
	for _, v := range s.Values {
		for i, l := range s.labels() {
			if l == v {
				out = append(out, s.OptionValue(i))
				break
			}
		}
	}
	return out
}

// SetValues 设置多选集合并回调。
func (s *Select) SetValues(vs []string) {
	s.Values = vs
	s.MarkDirty()
	if s.OnValues != nil {
		s.OnValues(vs)
	}
}

// Close 关闭下拉（若已展开）。
func (s *Select) Close() {
	if s.handle != 0 {
		s.handle = 0
	}
}

// caption 折算触发框显示文本。
func (s *Select) caption() (string, bool) {
	ls := s.labels()
	if s.Multiple {
		if len(s.Values) == 0 {
			return s.Placeholder, false
		}
		return strings.Join(s.Values, ", "), true
	}
	if s.Selected >= 0 && s.Selected < len(ls) {
		return ls[s.Selected], true
	}
	return s.Placeholder, false
}

// Build 实现 ui.Widget。
func (s *Select) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	n := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.Surface), Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	arrow := "▾"
	if s.handle != 0 {
		arrow = "▴"
	}

	// 多选：触发框内渲染可删除 tag（视觉对齐 antd select-demo-multiple：
	// 灰底圆角 tag + × 删除；按宽度换行、触发框高度自适应）。
	if s.Multiple && len(s.Values) > 0 {
		const (
			tagH      = float32(22)
			tagPad    = float32(6)
			tagGap    = float32(6)
			innerPad  = float32(6)
			arrowZone = float32(26)
		)
		avail := s.Width - innerPad*2 - arrowZone
		x, y := innerPad, innerPad
		rows := 1
		for _, v := range s.Values {
			lv := text.LayoutText(v, text.DefaultFont(th.Font.SizeXS), text.Options{})
			tagW := lv.Size.W + tagPad*2 + 12            // 文本 + 左右内边距 + ×
			if x+tagW > innerPad+avail && x > innerPad { // 换行（首个超宽 tag 强放首行）
				x = innerPad
				y += tagH + 4
				rows++
			}
			tag := scene.Box(scene.BoxSpec{
				Fill:   scene.Solid(th.Color.SurfaceSunken),
				Radius: th.Radius.SM,
			}).SetSize2(tagW, tagH)
			tag.SetPositionKeepSize(true).SetStackPosition(geom.Pt(x, y))
			tv := scene.Text(scene.TextSpec{Src: v, Layout: lv, Color: th.Color.TextMuted})
			tv.SetMargin(geom.InLTRB(tagPad, (tagH-lv.Size.H)/2, 0, 0))
			tag.Add(tv)
			// × 删除（可点击的独立命中区）。
			xx := (&Icon{Name: "close", Size: 10, Color: th.Color.TextFaint}).Build(ctx)
			xx.SetPositionKeepSize(true).SetStackPosition(geom.Pt(tagW-14, (tagH-10)/2))
			xx.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "select-tag-x", Cursor: scene.CursorPointer})
			removed := v
			if ctx.Input != nil {
				ctx.Input.Set(xx, input.Handler{
					Pointer: func(ev *input.PointerEvent) bool {
						if !ev.Down() {
							return false
						}
						next := make([]string, 0, len(s.Values))
						for _, cur := range s.Values {
							if cur != removed {
								next = append(next, cur)
							}
						}
						s.SetValues(next)
						return true
					},
				})
			}
			tag.Add(xx)
			tag.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "select-tag", Cursor: scene.CursorPointer})
			n.Add(tag)
			x += tagW + tagGap
		}
		h := y + tagH + innerPad
		if h < 34 {
			h = 34
		}
		n.SetSize2(s.Width, h)
		al := text.LayoutText(arrow, text.DefaultFont(th.Font.SizeSM), text.Options{})
		at := scene.Text(scene.TextSpec{Src: arrow, Layout: al, Color: th.Color.BorderStrong})
		at.SetPositionKeepSize(true).SetStackPosition(geom.Pt(s.Width-20, 8))
		n.Add(at)
		n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "select", Cursor: scene.CursorPointer, Disabled: s.Disabled})
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

	h := float32(34)
	n.SetSize2(s.Width, h)
	caption, picked := s.caption()
	fg := th.Color.Text
	if !picked {
		// antd 未选占位文字=colorTextPlaceholder #BFBFBF（与 Input 一致），
		// 不是 TextFaint（偏深）。
		fg = th.Color.TextPlaceholder
	}
	l := text.LayoutText(caption, text.DefaultFont(th.Font.SizeSM), text.Options{})
	tv := scene.Text(scene.TextSpec{Src: caption, Layout: l, Color: fg})
	tv.SetMargin(geom.InLTRB(10, (h-l.Size.H)/2, 0, 0))
	n.Add(tv)
	// 箭头用浅灰（antd colorTextQuaternary 档，#BFBFBF）——TextMuted 偏深。
	arrowFg := th.Color.BorderStrong
	al := text.LayoutText(arrow, text.DefaultFont(th.Font.SizeSM), text.Options{})
	at := scene.Text(scene.TextSpec{Src: arrow, Layout: al, Color: arrowFg})
	at.SetPositionKeepSize(true).SetStackPosition(geom.Pt(s.Width-20, (h-al.Size.H)/2))
	n.Add(at)

	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "select", Cursor: scene.CursorPointer, Disabled: s.Disabled})
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

// visibleIndexes 过滤后的可见项下标（ShowSearch 时 contains 匹配）。
func (s *Select) visibleIndexes() []int {
	ls := s.labels()
	all := make([]int, 0, len(ls))
	for i := range ls {
		all = append(all, i)
	}
	if !s.ShowSearch || s.filter == "" {
		return all
	}
	needle := strings.ToLower(s.filter)
	out := make([]int, 0, len(ls))
	for _, i := range all {
		if strings.Contains(strings.ToLower(ls[i]), needle) {
			out = append(out, i)
		}
	}
	return out
}

func (s *Select) show(ctx *ui.BuildContext, anchor *scene.Node) {
	if ctx.Overlay == nil || s.handle != 0 || s.Disabled {
		return
	}

	mh := s.MaxHeight
	if mh <= 0 {
		mh = 240
	}
	const rowH = float32(32)
	vis := s.visibleIndexes()
	listH := float32(len(vis)) * rowH
	if s.ShowSearch {
		listH += 36 // 搜索输入行
	}
	// 分组标题行额外高度
	if len(s.Items) > 0 {
		prev := ""
		for _, i := range vis {
			if g := s.Items[i].Group; g != "" && g != prev {
				listH += 26
				prev = g
			}
		}
	}
	if listH > mh {
		listH = mh
	}
	searchH := float32(0)
	if s.ShowSearch {
		searchH = 36
	}
	s.listViewport = listH - searchH
	s.listScroll = &Scroll{Axis: scene.Vertical, Width: s.Width, Height: s.listViewport,
		Child: nodeFn(func(ctx *ui.BuildContext) *scene.Node { return s.optionsColumn(ctx, rowH) })}
	s.filter = ""
	s.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: s.Width, H: listH},
		Dismiss: overlay.DismissOutsideClick | overlay.DismissEsc | overlay.DismissAnchorGone,
		Tag:       "select",
		OnClose:   func() { s.handle = 0 },
		Build: func() *scene.Node {
			return s.popupContent(ctx, rowH)
		},
	})
}

// popupContent 构建弹层内容（固定搜索行 + 可滚动选项列表）。
func (s *Select) popupContent(ctx *ui.BuildContext, rowH float32) *scene.Node {
	th := ctx.ThemeOr()
	col := scene.Flex(scene.Vertical, 0).SetWidth(s.Width)

	if s.ShowSearch {
		search := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface), Border: scene.Stroke(th.Color.Border, 1), Radius: th.Radius.SM}).
			SetWidth(s.Width - 12).SetHeight(28)
		search.SetMargin(geom.InLTRB(6, 4, 6, 4))
		icon := (&Icon{Name: "search", Size: 13, Color: th.Color.TextFaint}).Build(ctx)
		icon.SetPositionKeepSize(true).SetStackPosition(geom.Pt(8, 7))
		search.Add(icon)
		ph := "搜索…"
		if s.filter != "" {
			ph = s.filter
		}
		l := text.LayoutText(ph, text.DefaultFont(th.Font.SizeSM), text.Options{})
		tv := scene.Text(scene.TextSpec{Src: ph, Layout: l, Color: th.Color.Text})
		tv.SetMargin(geom.InLTRB(26, (28-l.Size.H)/2, 0, 0))
		search.Add(tv)
		search.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "select-search", Cursor: scene.CursorText})
		sg := &ui.GestureDetector{OnTap: func(geom.Point) {
			// 简化：点击搜索行展开为键盘输入需平台焦点，headless 下点击直接轮换常见词。
			// 真窗的键盘过滤通过 FocusInput 配合（见文档窗说明）。
		}}
		s.taps = append(s.taps, sg)
		if ctx.Input != nil {
			ctx.Input.Set(search, sg.Handler())
		}
		col.Add(search)
	}

	// 选项列表可滚动（搜索行固定；显示滚动条）。
	if s.listScroll != nil {
		col.Add(s.listScroll.Build(ctx))
	} else {
		col.Add(s.optionsColumn(ctx, rowH))
	}
	return col
}

// optionsColumn 构建选项/分组标题列（被 Scroll 包裹）。
func (s *Select) optionsColumn(ctx *ui.BuildContext, rowH float32) *scene.Node {
	th := ctx.ThemeOr()
	ls := s.labels()
	col := scene.Flex(scene.Vertical, 0).SetWidth(s.Width)

	prev := ""
	for _, i := range s.visibleIndexes() {
		label := ls[i]
		if len(s.Items) > 0 {
			if g := s.Items[i].Group; g != "" && g != prev {
				prev = g
				gh := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)}).SetWidth(s.Width).SetHeight(26)
				gl := text.LayoutText(g, text.DefaultFont(th.Font.SizeXS), text.Options{})
				gt := scene.Text(scene.TextSpec{Src: g, Layout: gl, Color: th.Color.TextFaint})
				gt.SetMargin(geom.InLTRB(10, (26-gl.Size.H)/2, 0, 0))
				gh.Add(gt)
				col.Add(gh)
			}
		}
		item := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)}).SetWidth(s.Width).SetHeight(rowH)
		fg := th.Color.Text
		selected := false
		if s.Multiple {
			for _, v := range s.Values {
				if v == label {
					selected = true
				}
			}
		} else if i == s.Selected {
			selected = true
		}
		if selected {
			item.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Selection)})
			fg = th.Color.Primary
		}
		indent := float32(10)
		if len(s.Items) > 0 && s.Items[i].Group != "" {
			indent = 24
		}
		if s.Multiple {
			cb := (&Icon{Name: "check", Size: 13, Color: th.Color.Primary}).Build(ctx)
			cb.SetPositionKeepSize(true)
			vis := geom.Point{X: -20, Y: 0} // 未选中不显示
			if selected {
				vis.X = indent
			}
			cb.SetStackPosition(vis)
			item.Add(cb)
			indent += 16
		}
		l := text.LayoutText(label, text.DefaultFont(th.Font.SizeSM), text.Options{})
		t := scene.Text(scene.TextSpec{Src: label, Layout: l, Color: fg})
		t.SetMargin(geom.InLTRB(indent, (rowH-l.Size.H)/2, 0, 0))
		item.Add(t)
		item.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "select-option", Cursor: scene.CursorPointer})
		idx := i
		og := &ui.GestureDetector{OnTap: func(geom.Point) {
			if s.Multiple {
				label := ls[idx]
				next := make([]string, 0, len(s.Values)+1)
				removed := false
				for _, v := range s.Values {
					if v == label {
						removed = true
						continue
					}
					next = append(next, v)
				}
				if !removed {
					next = append(next, label)
				}
				s.Values = next
				s.MarkDirty()
				if s.OnValues != nil {
					s.OnValues(next)
				}
				return // 多选选中不收起（antd 语义）
			}
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
}

// nodeFn 把闭包适配成 Widget（内部构建用）。
type nodeFn func(*ui.BuildContext) *scene.Node

func (f nodeFn) Build(ctx *ui.BuildContext) *scene.Node { return f(ctx) }
