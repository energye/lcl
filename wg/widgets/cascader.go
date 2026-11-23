package widgets

// ── Cascader 级联选择（M12 高级录入）──────────────────────────────
//
// antd：<Cascader options value onChange placeholder changeOnSelect>
// 第一轮：触发框（Input 形态+down 图标+路径文本 " / " 分隔）+ 级联
// 面板（多列并排，点含子级的选项展开下一列，antd 令牌：列宽 180、
// item 高 28、列视口高 240）+ 叶子选择关面板（默认只能选叶子）+
// 路径上/选中项文字主色（antd 级联高亮是文字色不是实底）。
//
// 追加列 = **同帧 close+reopen**（状态驱动重开：overlay Entry 宽度
// 在 Show 时定死，列追加必须重开面板让宽度随列数走——同帧内完成
// 无视觉闪烁；重开后 expandPath 保持展开状态）。
//
// 组合复用：Icon（down/right）、Scroll（列视口）、BandPopup+
// PlacementBottomStart（Select/DatePicker/TimePicker 同通道）。
// showSearch 搜索：触发框变输入框，输入按路径 label 过滤，搜索态
// 扁平列表（显示完整路径，点选即选中整链）；filter/onSearch 可定制。
// 多选（antd multiple）：面板项嵌 Checkbox（Tree checkable 同款组合
// 复用），勾选完整路径链入集合（antd value 形态 string[][]），点选
// **不关面板**；触发框 tag 形态显示叶子 Label（Select/TreeSelect 多选
// 同款视觉+MaxTagCount 折叠）。
// 未实现（如实标注）：异步加载、changeOnSelect（任意级可选）、
// 键盘导航、禁用项、自定义 displayRender。

import (
	"fmt"
	"strings"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// 面板尺寸令牌（antd：列宽 180、item 高 28、列视口高 240）。
const (
	cascaderColW  = float32(180)
	cascaderItemH = float32(28)
	cascaderColH  = float32(240)
	cascaderGapX  = float32(4) // 列间视觉留白（antd 列间 border 分隔近似）
)

// CascaderOption 级联选项（antd options 项）。
type CascaderOption struct {
	Value    string
	Label    string
	Children []*CascaderOption
}

// Cascader 是级联选择器。
type Cascader struct {
	controlBase
	Options     []*CascaderOption
	Value       []string // 选中路径（叶子链的 Value 链）
	Placeholder string
	Width       float32
	OnChange    func(path []string)

	// ShowSearch 搜索框（antd showSearch）：触发框变输入框，输入按
	// 路径 label 过滤，搜索态扁平列表（显示完整路径）。
	ShowSearch bool
	// OnSearch 搜索内容变化回调（antd onSearch(search)）。
	OnSearch func(search string)
	// Filter 自定义过滤（antd showSearch.filter(inputValue, path)——
	// path 为完整路径；nil = 默认路径任意 label 含搜索子串（不区分大小写）。
	Filter func(inputValue string, path []*CascaderOption) bool

	// Multiple 多选（antd multiple）：面板项尾嵌 Checkbox 显示勾选态，
	// 点行 toggle 完整路径，不关面板。
	Multiple bool
	// Values 多选值集合（antd value 形态 string[][]——每项完整路径链）。
	Values [][]string
	// MaxTagCount 触发框最多显示 tag 数，超出折叠 "+N"（0=不限制）。
	MaxTagCount int
	// OnValuesChange 多选值变化回调（antd onChange multiple 返回数组）。
	OnValuesChange func(paths [][]string)

	// checks 跨面板开合持有的 Checkbox 实例（按路径 Key 池化——hover
	// 态随重建丢失教训）。
	checks map[string]*Checkbox

	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	panelOpen bool
	anchor    *scene.Node
	// expandPath 面板打开期间的展开路径（每级下标；长度=已展开列数-1）。
	expandPath []int
	// 搜索文本与池化输入框（跨 Build 持有，编辑态/焦点不丢）。
	search      string
	searchInput *Input
}

// cascaderPathKey 路径链的唯一键（Value join；多选池化/集合判定标识）。
func cascaderPathKey(path []string) string {
	return strings.Join(path, "/")
}

// NewCascader 创建级联选择器（默认占位"请选择"）。
func NewCascader() *Cascader {
	return &Cascader{Placeholder: "请选择"}
}

// displayText 触发框显示文本：antd 默认 displayRender=Label 链 join
// （链对不上时兜底 Value join）。
func (c *Cascader) displayText() string {
	if len(c.Value) == 0 {
		return c.Placeholder
	}
	labels := make([]string, 0, len(c.Value))
	node := c.Options
	for _, v := range c.Value {
		found := ""
		for _, o := range node {
			if o.Value == v {
				found = o.Label
				node = o.Children
				break
			}
		}
		if found == "" {
			return strings.Join(c.Value, " / ")
		}
		labels = append(labels, found)
	}
	return strings.Join(labels, " / ")
}

// Build 实现 ui.Widget：触发框。
func (c *Cascader) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(c)
	th := ctx.ThemeOr()
	w := c.Width
	if w <= 0 {
		w = 220
	}
	h := ui.MetricsFor(th, ui.SizeMedium).Height

	// 搜索触发框：Input 承载（可编辑），down 图标贴右。
	if c.ShowSearch {
		if c.searchInput == nil {
			c.searchInput = &Input{HitLabel: "cascader"}
		}
		c.searchInput.Placeholder = c.Placeholder
		c.searchInput.Value = c.search
		c.searchInput.Width = w
		c.searchInput.Height = h
		c.searchInput.OnChanged = func(v string) {
			c.search = v
			if c.OnSearch != nil {
				c.OnSearch(v)
			}
			c.expandPath = nil
			// 延迟重开面板（击键分发内直接重开扰动输入焦点——
			// DatePicker refreshPanel 同款教训）。
			time.AfterFunc(0, func() { c.reopenPanel(ctx, th) })
		}
		n := c.searchInput.Build(ctx)
		ic := (&Icon{Name: "down", Size: 12, Color: th.Color.TextFaint}).Build(ctx)
		ic.SetPositionKeepSize(true)
		ic.SetStackPosition(geom.Pt(w-11-12, (h-12)/2))
		n.Add(ic)
		c.anchor = n
		c.childNode = n
		return n
	}

	display := c.displayText()
	txtColor := th.Color.Text
	if len(c.Value) == 0 {
		display, txtColor = c.Placeholder, th.Color.TextPlaceholder
	}
	font := text.DefaultFont(th.Font.SizeBase)
	layout := text.LayoutText(display, font, text.Options{})

	// 多选触发框：tag 形态（Select/TreeSelect 同款视觉——灰底圆角+×
	// 删除、换行高度自适应、MaxTagCount 折叠 +N；显示叶子 Label）。
	if c.Multiple && len(c.Values) > 0 {
		n := c.buildMultiTrigger(ctx, th, w)
		n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "cascader", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			anchor := n
			c.anchor = n
			g := &ui.GestureDetector{OnTap: func(geom.Point) {
				if c.panelOpen {
					c.closePanel(ctx)
					return
				}
				c.expandPath = nil
				c.openPanel(ctx, anchor, th)
			}}
			ctx.Input.Set(n, g.Handler())
		}
		c.childNode = n
		return n
	}

	n := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	}).SetSize2(w, h)
	label := scene.Text(scene.TextSpec{Src: display, Layout: layout, Color: txtColor})
	label.SetMargin(geom.InLTRB(11, (h-layout.Size.H)/2, 0, 0))
	n.Add(label)
	ic := (&Icon{Name: "down", Size: 12, Color: th.Color.TextFaint}).Build(ctx)
	ic.SetPositionKeepSize(true)
	ic.SetStackPosition(geom.Pt(w-11-12, (h-12)/2))
	n.Add(ic)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "cascader", Cursor: scene.CursorPointer})

	if ctx.Input != nil {
		anchor := n
		c.anchor = n
		g := &ui.GestureDetector{OnTap: func(geom.Point) {
			if c.panelOpen {
				c.closePanel(ctx)
				return
			}
			c.expandPath = nil
			c.openPanel(ctx, anchor, th)
		}}
		ctx.Input.Set(n, g.Handler())
	}
	c.childNode = n
	return n
}

// buildMultiTrigger 多选触发框：可删除 tag 按宽度换行（Select 多选同款
// 视觉），MaxTagCount 超出折叠 "+N"，tag × down 即删；显示叶子 Label。
func (c *Cascader) buildMultiTrigger(ctx *ui.BuildContext, th *theme.Data, w float32) *scene.Node {
	const (
		tagH      = float32(22)
		tagPad    = float32(6)
		tagGap    = float32(6)
		innerPad  = float32(6)
		arrowZone = float32(26)
	)
	h := ui.MetricsFor(th, ui.SizeMedium).Height
	n := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	avail := w - innerPad*2 - arrowZone
	x, y := innerPad, innerPad
	shown, overflow := c.Values, 0
	if c.MaxTagCount > 0 && len(c.Values) > c.MaxTagCount {
		shown, overflow = c.Values[:c.MaxTagCount], len(c.Values)-c.MaxTagCount
	}
	for _, path := range shown {
		label := c.leafLabel(path)
		lv := text.LayoutText(label, text.DefaultFont(th.Font.SizeXS), text.Options{})
		tagW := lv.Size.W + tagPad*2 + 12
		if x+tagW > innerPad+avail && x > innerPad {
			x = innerPad
			y += tagH + 4
		}
		tag := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.SurfaceSunken),
			Radius: th.Radius.SM,
		}).SetSize2(tagW, tagH)
		tag.SetPositionKeepSize(true).SetStackPosition(geom.Pt(x, y))
		tv := scene.Text(scene.TextSpec{Src: label, Layout: lv, Color: th.Color.TextMuted})
		tv.SetMargin(geom.InLTRB(tagPad, (tagH-lv.Size.H)/2, 0, 0))
		tag.Add(tv)
		xx := (&Icon{Name: "close", Size: 10, Color: th.Color.TextFaint}).Build(ctx)
		xx.SetPositionKeepSize(true).SetStackPosition(geom.Pt(tagW-14, (tagH-10)/2))
		xx.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "cascader-tag-x", Cursor: scene.CursorPointer})
		removed := append([]string(nil), path...)
		if ctx.Input != nil {
			ctx.Input.Set(xx, input.Handler{
				Pointer: func(ev *input.PointerEvent) bool {
					if !ev.Down() {
						return false
					}
					c.togglePath(ctx, removed)
					return true
				},
			})
		}
		tag.Add(xx)
		// tag 本体**被动锚**（不拦截触发框点击——antd 点 tag 本体=开
		// 面板；× 热区独立 opaque 可删）。
		tag.SetHit(scene.HitSpec{Label: "cascader-tag"})
		n.Add(tag)
		x += tagW + tagGap
	}
	if overflow > 0 {
		chipText := fmt.Sprintf("+%d", overflow)
		cl := text.LayoutText(chipText, text.DefaultFont(th.Font.SizeXS), text.Options{})
		chipW := cl.Size.W + tagPad*2
		chip := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.SurfaceSunken),
			Radius: th.Radius.SM,
		}).SetSize2(chipW, tagH)
		chip.SetPositionKeepSize(true).SetStackPosition(geom.Pt(x, y))
		ct := scene.Text(scene.TextSpec{Src: chipText, Layout: cl, Color: th.Color.TextMuted})
		ct.SetMargin(geom.InLTRB(tagPad, (tagH-cl.Size.H)/2, 0, 0))
		chip.Add(ct)
		chip.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "cascader-tag-more"})
		n.Add(chip)
	}
	hh := y + tagH + innerPad
	if hh < h {
		hh = h
	}
	n.SetSize2(w, hh)
	ic := (&Icon{Name: "down", Size: 12, Color: th.Color.TextFaint}).Build(ctx)
	ic.SetPositionKeepSize(true)
	ic.SetStackPosition(geom.Pt(w-11-12, (hh-12)/2))
	n.Add(ic)
	return n
}

// leafLabel 路径链的叶子显示 Label（链对不上兜底末位 Value）。
func (c *Cascader) leafLabel(path []string) string {
	node := c.Options
	for k, v := range path {
		found := ""
		for _, o := range node {
			if o.Value == v {
				found = o.Label
				node = o.Children
				break
			}
		}
		if found == "" {
			return path[len(path)-1]
		}
		if k == len(path)-1 {
			return found
		}
	}
	return path[len(path)-1]
}

// Show 程序化打开面板（锚=最近构建的触发框节点；测试/自动化用）。
func (c *Cascader) Show(ctx *ui.BuildContext) {
	if c.childNode == nil {
		return
	}
	c.expandPath = nil
	c.openPanel(ctx, c.childNode, ctx.ThemeOr())
}

// IsOpen 返回面板是否打开。
func (c *Cascader) IsOpen() bool { return c.panelOpen }

// Close 关闭面板。
func (c *Cascader) Close() {
	if c.fb.host != nil {
		c.closePanel(nil)
	}
}

// openPanel 打开级联面板。
func (c *Cascader) openPanel(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data) {
	if c.handle != 0 {
		return
	}
	c.panelOpen = true
	depth := c.openDepth()
	panel := c.buildPanel(ctx, th, depth)
	w := float32(depth+1)*cascaderColW + float32(depth)*cascaderGapX
	if c.ShowSearch && c.search != "" {
		w = cascaderColW // 搜索态扁平列表单列
	}
	c.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: w, H: 0}, // 高自适应（列 240）
		Dismiss:   overlay.DismissOutsideClick | overlay.DismissEsc,
		Tag:       "cascader",
		OnClose: func() {
			c.handle = 0
			c.panelOpen = false
		},
		Build: func() *scene.Node { return panel },
	})
}

// closePanel 关闭面板。
func (c *Cascader) closePanel(ctx *ui.BuildContext) {
	if c.handle != 0 && c.fb.host != nil {
		c.fb.host.Close(c.handle)
	}
	c.handle = 0
	c.panelOpen = false
}

// openDepth 当前展开列数（1=仅第一列；叶子选中后路径长）。
func (c *Cascader) openDepth() int {
	depth := 1
	node := c.Options
	for i, idx := range c.expandPath {
		if idx < 0 || idx >= len(node) {
			break
		}
		if len(node[idx].Children) == 0 {
			break
		}
		node = node[idx].Children
		depth = i + 2
	}
	return depth
}

// nodesAtDepth 返回第 d 列（0 起）的选项集合；越界返回 nil。
func (c *Cascader) nodesAtDepth(d int) []*CascaderOption {
	node := c.Options
	for i := 0; i < d && i < len(c.expandPath); i++ {
		idx := c.expandPath[i]
		if idx < 0 || idx >= len(node) || len(node[idx].Children) == 0 {
			return nil
		}
		node = node[idx].Children
	}
	return node
}

// reopenPanel 同帧关+开（展开状态保持）——列追加时面板宽度随列数走。
func (c *Cascader) reopenPanel(ctx *ui.BuildContext, th *theme.Data) {
	if ctx == nil || ctx.Overlay == nil {
		return
	}
	if c.handle != 0 {
		ctx.Overlay.Close(c.handle)
		c.handle = 0
	}
	c.panelOpen = false
	c.openPanel(ctx, c.anchor, th)
}

// buildPanel 构建面板：搜索态扁平列表；否则 openDepth 列并排。
func (c *Cascader) buildPanel(ctx *ui.BuildContext, th *theme.Data, depth int) *scene.Node {
	if c.ShowSearch && c.search != "" {
		return c.buildSearchPanel(ctx, th)
	}
	panel := scene.Flex(scene.Horizontal, 0)
	panel.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	panel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "cascader-panel"})
	for d := 0; d < depth; d++ {
		col := c.buildColumn(ctx, th, d)
		scroll := &Scroll{Axis: scene.Vertical, Width: cascaderColW, Height: cascaderColH,
			Child: nodeFn(func(ctx *ui.BuildContext) *scene.Node { return col })}
		wrap := scene.Flex(scene.Vertical, 0)
		wrap.Add(scroll.Build(ctx))
		wrap.SetHit(scene.HitSpec{Label: "cascader-col"})
		panel.Add(wrap)
		if d < depth-1 {
			panel.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetSize2(1, cascaderColH))
		}
	}
	return panel
}

// buildColumn 构建一列的选项列表。
func (c *Cascader) buildColumn(ctx *ui.BuildContext, th *theme.Data, d int) *scene.Node {
	items := c.nodesAtDepth(d)
	col := scene.Flex(scene.Vertical, 0)
	for i, opt := range items {
		i, opt := i, opt
		// 高亮：路径上/已选链的项文字主色（antd 级联文字色高亮）。
		onPath := d < len(c.expandPath) && c.expandPath[d] == i
		font := text.DefaultFont(th.Font.SizeBase)
		l := text.LayoutText(opt.Label, font, text.Options{})
		color := th.Color.Text
		if onPath {
			color = th.Color.Primary
		}
		item := scene.Box(scene.BoxSpec{})
		item.SetSize2(cascaderColW, cascaderItemH)
		txt := scene.Text(scene.TextSpec{Src: opt.Label, Layout: l, Color: color})
		txt.SetMargin(geom.InLTRB(12, (cascaderItemH-l.Size.H)/2, 0, 0))
		item.Add(txt)
		if c.Multiple {
			// 多选：项尾嵌 Checkbox 组件实例（Tree checkable 同款组合
			// 复用），池化按路径 Key、只作显示态——行级点击走 pick
			//（toggle 路径+含子级深化），Checkbox 不抢点击。
			path := c.pathPrefix(d, i, opt)
			key := cascaderPathKey(path)
			if c.checks == nil {
				c.checks = map[string]*Checkbox{}
			}
			ck := c.checks[key]
			if ck == nil {
				ck = &Checkbox{}
				c.checks[key] = ck
			}
			ck.Checked = c.hasPath(path)
			cbn := ck.Build(ctx)
			cbn.SetHit(scene.HitSpec{Label: "cascader-check"})
			cbn.SetPositionKeepSize(true)
			cbn.SetStackPosition(geom.Pt(cascaderColW-12-16, (cascaderItemH-16)/2))
			item.Add(cbn)
		} else if len(opt.Children) > 0 {
			ic := (&Icon{Name: "right", Size: 11, Color: th.Color.TextFaint}).Build(ctx)
			ic.SetPositionKeepSize(true)
			ic.SetStackPosition(geom.Pt(cascaderColW-12-11, (cascaderItemH-11)/2))
			item.Add(ic)
		}
		item.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "cascader-item", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			ctx.Input.Set(item, input.Handler{
				Enter: func(*input.PointerEvent) {
					item.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.FillHover)})
					item.MarkNeedsPaint()
				},
				Leave: func(*input.PointerEvent) {
					item.SetBox(scene.BoxSpec{})
					item.MarkNeedsPaint()
				},
				Pointer: func(ev *input.PointerEvent) bool {
					if !ev.Down() {
						return false
					}
					// down 即执行（浮层内 tap 配对不可靠结论复用）。
					c.pick(ctx, th, d, i, opt)
					return true
				},
			})
		}
		col.Add(item)
	}
	return col
}

// searchMatches 搜索匹配项（完整路径列表，扁平化所有层级）。
func (c *Cascader) searchMatches() [][]*CascaderOption {
	var matches [][]*CascaderOption
	var walk func(node []*CascaderOption, path []*CascaderOption)
	walk = func(node []*CascaderOption, path []*CascaderOption) {
		for _, o := range node {
			p := make([]*CascaderOption, len(path), len(path)+1)
			copy(p, path)
			p = append(p, o)
			if c.matchSearch(p) {
				matches = append(matches, p)
			}
			if len(o.Children) > 0 {
				walk(o.Children, p)
			}
		}
	}
	walk(c.Options, nil)
	return matches
}

// matchSearch 判断路径是否匹配搜索（默认：路径任意 label 含搜索子串，不区分大小写）。
func (c *Cascader) matchSearch(path []*CascaderOption) bool {
	if c.search == "" {
		return true
	}
	if c.Filter != nil {
		return c.Filter(c.search, path)
	}
	needle := strings.ToLower(c.search)
	for _, o := range path {
		if strings.Contains(strings.ToLower(o.Label), needle) {
			return true
		}
	}
	return false
}

// cascaderPathLabel 路径的显示文本（label 链 " / " 分隔）。
func cascaderPathLabel(path []*CascaderOption) string {
	labels := make([]string, 0, len(path))
	for _, o := range path {
		labels = append(labels, o.Label)
	}
	return strings.Join(labels, " / ")
}

// buildSearchPanel 搜索态扁平列表（每项显示完整路径，点选即选中整链）。
func (c *Cascader) buildSearchPanel(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	panel := scene.Flex(scene.Vertical, 0)
	panel.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	panel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "cascader-panel"})
	for _, path := range c.searchMatches() {
		path := path
		label := cascaderPathLabel(path)
		font := text.DefaultFont(th.Font.SizeBase)
		l := text.LayoutText(label, font, text.Options{})
		item := scene.Box(scene.BoxSpec{})
		item.SetSize2(cascaderColW, cascaderItemH)
		txt := scene.Text(scene.TextSpec{Src: label, Layout: l, Color: th.Color.Text})
		txt.SetMargin(geom.InLTRB(12, (cascaderItemH-l.Size.H)/2, 0, 0))
		item.Add(txt)
		item.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "cascader-search-item", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			ctx.Input.Set(item, input.Handler{
				Enter: func(*input.PointerEvent) {
					item.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.FillHover)})
					item.MarkNeedsPaint()
				},
				Leave: func(*input.PointerEvent) {
					item.SetBox(scene.BoxSpec{})
					item.MarkNeedsPaint()
				},
				Pointer: func(ev *input.PointerEvent) bool {
					if !ev.Down() {
						return false
					}
					c.selectPath(ctx, path)
					return true
				},
			})
		}
		panel.Add(item)
	}
	return panel
}

// selectPath 选中搜索匹配的完整路径：设 Value+清搜索+关面板+回调。
func (c *Cascader) selectPath(ctx *ui.BuildContext, path []*CascaderOption) {
	values := make([]string, 0, len(path))
	for _, o := range path {
		values = append(values, o.Value)
	}
	c.Value = values
	c.search = ""
	if c.searchInput != nil {
		c.searchInput.Value = ""
	}
	c.MarkDirty()
	if ctx != nil && ctx.Overlay != nil && c.handle != 0 {
		ctx.Overlay.Close(c.handle)
	}
	c.handle = 0
	c.panelOpen = false
	if c.OnChange != nil {
		c.OnChange(values)
	}
}

// pick 选中某列某项：含子级→同帧重开面板深化；叶子→选中+关面板。
// Multiple 时点行走 togglePath（勾选不关面板——antd 多选语义；含子级
// 先深化展开再 toggle）。
func (c *Cascader) pick(ctx *ui.BuildContext, th *theme.Data, d, i int, opt *CascaderOption) {
	if c.Multiple {
		if len(opt.Children) > 0 {
			for len(c.expandPath) > d {
				c.expandPath = c.expandPath[:len(c.expandPath)-1]
			}
			for len(c.expandPath) < d {
				c.expandPath = append(c.expandPath, -1)
			}
			c.expandPath = append(c.expandPath, i)
			c.reopenPanel(ctx, th)
		}
		c.togglePath(ctx, c.pathPrefix(d, i, opt))
		return
	}
	if len(opt.Children) > 0 {
		// 深化展开路径到 d+1。
		for len(c.expandPath) > d {
			c.expandPath = c.expandPath[:len(c.expandPath)-1]
		}
		for len(c.expandPath) < d {
			c.expandPath = append(c.expandPath, -1) // 防御：不应发生
		}
		c.expandPath = append(c.expandPath, i)
		c.reopenPanel(ctx, th)
		return
	}
	// 叶子：路径 = 前缀 expandPath 的 Value 链 + 本项 Value。
	path := make([]string, 0, d+1)
	node := c.Options
	for k := 0; k < d && k < len(c.expandPath); k++ {
		idx := c.expandPath[k]
		if idx < 0 || idx >= len(node) {
			break
		}
		path = append(path, node[idx].Value)
		node = node[idx].Children
	}
	path = append(path, opt.Value)
	c.Value = path
	c.MarkDirty() // 触发框路径文本随根重建刷新
	if ctx != nil && ctx.Overlay != nil && c.handle != 0 {
		ctx.Overlay.Close(c.handle)
	}
	c.handle = 0
	c.panelOpen = false
	if c.OnChange != nil {
		c.OnChange(path)
	}
}

// togglePath 多选：路径勾选/取消（antd 多选点选不关面板）。down 分发内
// 改集合——MarkDirty 刷触发框；面板开着时 AfterFunc(0) 延迟 reopen 刷新
// 浮层勾选态（down 分发内重开扰动——Cascader 搜索轮同款教训）。
func (c *Cascader) togglePath(ctx *ui.BuildContext, path []string) {
	key := cascaderPathKey(path)
	removed := false
	next := make([][]string, 0, len(c.Values))
	for _, p := range c.Values {
		if cascaderPathKey(p) == key {
			removed = true
			continue
		}
		next = append(next, p)
	}
	if !removed {
		next = append(next, append([]string(nil), path...))
	}
	c.Values = next
	c.MarkDirty()
	if c.OnValuesChange != nil {
		c.OnValuesChange(c.Values)
	}
	if c.panelOpen {
		time.AfterFunc(0, func() { c.refreshPanelAfterToggle(ctx) })
	}
}

// refreshPanelAfterToggle 多选勾选后同帧 close+reopen 刷新浮层勾选态
// （expandPath 不动——展开状态保持）。
func (c *Cascader) refreshPanelAfterToggle(ctx *ui.BuildContext) {
	if ctx == nil || ctx.Overlay == nil || !c.panelOpen || c.handle == 0 {
		return
	}
	c.reopenPanel(ctx, ctx.ThemeOr())
}

// pathPrefix 第 d 列第 i 项的完整路径 Value 链（expandPath 前缀+本项）。
func (c *Cascader) pathPrefix(d, i int, opt *CascaderOption) []string {
	path := make([]string, 0, d+1)
	node := c.Options
	for k := 0; k < d && k < len(c.expandPath); k++ {
		idx := c.expandPath[k]
		if idx < 0 || idx >= len(node) {
			break
		}
		path = append(path, node[idx].Value)
		node = node[idx].Children
	}
	return append(path, opt.Value)
}

// hasPath 多选集合是否含该路径。
func (c *Cascader) hasPath(path []string) bool {
	key := cascaderPathKey(path)
	for _, p := range c.Values {
		if cascaderPathKey(p) == key {
			return true
		}
	}
	return false
}
