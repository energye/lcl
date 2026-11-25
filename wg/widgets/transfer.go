package widgets

// ── Transfer 穿梭框（M12 高级录入）───────────────────────────────
//
// antd：<Transfer dataSource targetKeys onChange render>
// 第一轮：双栏列表（源/目标）+ 中间操作列（右移/左移/全右/全左）
// + 列表项 **Checkbox 组件勾选**（池化实例——Tree.Checkable 先例）
// + 标题行（标题+勾选/总数计数）+ 点行切换勾选。
// TargetKeys 驱动左右分栏；OnChange(targetKeys, direction, moveKeys)。
//
// 组合复用：Checkbox（项勾选，池化实例跨 Build 持有）、Button
// （操作列四钮——外包被动锚 transfer-btn-* 防撞全页 button）。
// showSearch 搜索框（两侧输入框按 title 过滤，filterOption/onSearch
// 可定制）；header 全选勾选框（showSelectAll，默认开启，勾选本栏
// 可见项）。未实现（如实标注）：拖拽排序、分页、footer 插槽、
// render 自定义行。

import (
	"fmt"
	"strings"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// TransferItem 是穿梭项（antd dataSource 项）。
type TransferItem struct {
	Key   string
	Title string
}

// TransferDirection 移动方向（OnChange 的 direction 参数）。
type TransferDirection uint8

const (
	TransferRight TransferDirection = iota // 左→右（选入）
	TransferLeft                           // 右→左（移出）
)

// Transfer 是穿梭框。
type Transfer struct {
	controlBase
	DataSource []TransferItem
	TargetKeys []string // 右栏的 Key 集
	OnChange   func(targetKeys []string, dir TransferDirection, moveKeys []string)
	Width      float32 // 单栏宽（默认 220）

	// ShowSearch 搜索框（antd showSearch）：两侧面板顶部输入框过滤列表。
	ShowSearch bool
	// SearchPlaceholder 搜索框占位（antd locale.searchPlaceholder，默认「请输入搜索内容」）。
	SearchPlaceholder string
	// FilterOption 自定义过滤（antd filterOption(inputValue, option, direction)——
	// direction 为 "left"/"right"；nil = 默认 title 不区分大小写子串匹配）。
	FilterOption func(inputValue string, item TransferItem, direction string) bool
	// OnSearch 搜索内容变化回调（antd onSearch(direction, value)）。
	OnSearch func(direction string, value string)
	// ShowSelectAll header 全选勾选框（antd showSelectAll，默认 true）。
	ShowSelectAll bool

	checked map[string]bool
	// checks 池化复选框实例（跨 Build 持有——hover 态不丢，Tree.Checkable 先例）。
	checks map[string]*Checkbox
	// opBtns 池化操作按钮（Build 内新建实例按压态随重建丢失——教训⑨，Upload 先例）。
	opBtns map[string]*Button
	// 搜索框文本与池化 Input/全选 checkbox（跨 Build 持有，编辑态/焦点不丢）。
	searchL, searchR string
	searches         map[string]*Input
	selAlls          map[string]*Checkbox
}

// NewTransfer 创建穿梭框。
func NewTransfer(items []TransferItem) *Transfer {
	return &Transfer{DataSource: items, Width: 220, ShowSelectAll: true,
		checked: map[string]bool{}, checks: map[string]*Checkbox{},
		opBtns: map[string]*Button{}, searches: map[string]*Input{},
		selAlls: map[string]*Checkbox{}}
}

// isTarget key 是否在右栏。
func (t *Transfer) isTarget(key string) bool {
	for _, k := range t.TargetKeys {
		if k == key {
			return true
		}
	}
	return false
}

// notify 汇报目标键集。
func (t *Transfer) notify(dir TransferDirection, moveKeys []string) {
	if t.OnChange != nil {
		t.OnChange(append([]string(nil), t.TargetKeys...), dir, moveKeys)
	}
}

// move 把勾选项从 dir 侧移到对侧（勾选态清空）。
func (t *Transfer) move(dir TransferDirection) {
	var moveKeys []string
	if dir == TransferRight {
		for _, it := range t.DataSource {
			if t.checked[it.Key] && !t.isTarget(it.Key) {
				t.TargetKeys = append(t.TargetKeys, it.Key)
				moveKeys = append(moveKeys, it.Key)
			}
		}
	} else {
		var keep []string
		for _, k := range t.TargetKeys {
			if t.checked[k] {
				moveKeys = append(moveKeys, k)
			} else {
				keep = append(keep, k)
			}
		}
		t.TargetKeys = keep
	}
	t.checked = map[string]bool{}
	for _, c := range t.checks { // 池化实例勾选态复位
		c.Checked = false
	}
	t.MarkDirty()
	if len(moveKeys) > 0 {
		t.notify(dir, moveKeys)
	}
}

// moveAll 全部移到对侧（不依赖勾选）。
func (t *Transfer) moveAll(dir TransferDirection) {
	var moveKeys []string
	if dir == TransferRight {
		for _, it := range t.DataSource {
			if !t.isTarget(it.Key) {
				t.TargetKeys = append(t.TargetKeys, it.Key)
				moveKeys = append(moveKeys, it.Key)
			}
		}
	} else {
		moveKeys = append(moveKeys, t.TargetKeys...)
		t.TargetKeys = nil
	}
	t.checked = map[string]bool{}
	for _, c := range t.checks {
		c.Checked = false
	}
	t.MarkDirty()
	if len(moveKeys) > 0 {
		t.notify(dir, moveKeys)
	}
}

// sideDir 面板方向（antd filterOption/onSearch 的 direction："left"/"right"）。
func (t *Transfer) sideDir(side string) string {
	if side == "右" {
		return "right"
	}
	return "left"
}

// search 某侧当前搜索文本。
func (t *Transfer) search(side string) string {
	if side == "右" {
		return t.searchR
	}
	return t.searchL
}

// setSearch 更新某侧搜索文本并重建列表。
func (t *Transfer) setSearch(side, value string) {
	if side == "右" {
		t.searchR = value
	} else {
		t.searchL = value
	}
	if t.OnSearch != nil {
		t.OnSearch(t.sideDir(side), value)
	}
	// 延迟重建列表（击键分发内直接重建扰动输入焦点——DatePicker
	// refreshPanel 同款教训）。
	time.AfterFunc(0, t.MarkDirty)
}

// filterItems 按搜索文本过滤列表（默认 title 不区分大小写子串匹配）。
func (t *Transfer) filterItems(items []TransferItem, direction, q string) []TransferItem {
	if q == "" {
		return items
	}
	out := make([]TransferItem, 0, len(items))
	for _, it := range items {
		if t.FilterOption != nil {
			if t.FilterOption(q, it, direction) {
				out = append(out, it)
			}
		} else if strings.Contains(strings.ToLower(it.Title), strings.ToLower(q)) {
			out = append(out, it)
		}
	}
	return out
}

// sideItems 某侧（左=源/右=目标）全部项。
func (t *Transfer) sideItems(side string) []TransferItem {
	var items []TransferItem
	for _, it := range t.DataSource {
		if t.isTarget(it.Key) {
			if side == "右" {
				items = append(items, it)
			}
		} else if side == "左" {
			items = append(items, it)
		}
	}
	return items
}

// selectAll 全选/反选某侧可见（过滤后）项。
func (t *Transfer) selectAll(side string, on bool) {
	for _, it := range t.filterItems(t.sideItems(side), t.sideDir(side), t.search(side)) {
		t.checked[it.Key] = on
		if cb := t.checks[it.Key]; cb != nil {
			cb.Checked = on
			cb.Indeterminate = false
		}
	}
	t.MarkDirty()
}

// selAllState 某侧全选勾选框状态（全可见勾选→checked；部分→indeterminate）。
func (t *Transfer) selAllState(side string) (bool, bool) {
	vis := t.filterItems(t.sideItems(side), t.sideDir(side), t.search(side))
	if len(vis) == 0 {
		return false, false
	}
	cnt := 0
	for _, it := range vis {
		if t.checked[it.Key] {
			cnt++
		}
	}
	if cnt == len(vis) {
		return true, false
	}
	if cnt > 0 {
		return false, true
	}
	return false, false
}

// Build 实现 ui.Widget：标题行+列表 ×2 + 中间操作列。
func (t *Transfer) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(t)
	th := ctx.ThemeOr()
	w := t.Width
	if w <= 0 {
		w = 220
	}

	var source, target []TransferItem
	for _, it := range t.DataSource {
		if t.isTarget(it.Key) {
			target = append(target, it)
		} else {
			source = append(source, it)
		}
	}

	row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)

	colL := t.buildSide(ctx, th, "左", "源", "transfer-item", source, w)
	colR := t.buildSide(ctx, th, "右", "目标", "transfer-item-r", target, w)
	row.Add(colL)
	row.Add(t.buildOps(ctx, th, len(source), len(target)))
	row.Add(colR)

	row.SetHit(scene.HitSpec{Label: "transfer"})
	return row
}

// buildOps 中间操作列（右移/左移/全右/全左——Button 组合复用+被动锚）。
func (t *Transfer) buildOps(ctx *ui.BuildContext, th *theme.Data, nSource, nTarget int) *scene.Node {
	col := scene.Flex(scene.Vertical, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	mk := func(label, text string, disabled bool, fn func()) {
		btn, ok := t.opBtns[label]
		if !ok {
			btn = &Button{Size: ui.SizeSmall}
			t.opBtns[label] = btn
		}
		btn.Text = text
		btn.Disabled = disabled
		btn.OnClick = fn
		wrap := scene.Flex(scene.Horizontal, 0)
		wrap.Add(btn.Build(ctx))
		wrap.SetHit(scene.HitSpec{Label: label})
		col.Add(wrap)
	}
	mk("transfer-btn-right", ">", nSource == 0, func() { t.move(TransferRight) })
	mk("transfer-btn-left", "<", nTarget == 0, func() { t.move(TransferLeft) })
	mk("transfer-btn-right-all", "»", nSource == 0, func() { t.moveAll(TransferRight) })
	mk("transfer-btn-left-all", "«", nTarget == 0, func() { t.moveAll(TransferLeft) })
	return col
}

// buildSide 构建一栏（标题行+搜索框+项列表）。
func (t *Transfer) buildSide(ctx *ui.BuildContext, th *theme.Data, side, title, itemLabel string,
	items []TransferItem, w float32) *scene.Node {
	direction := t.sideDir(side)
	vis := t.filterItems(items, direction, t.search(side))

	col := scene.Flex(scene.Vertical, 0)
	col.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	col.SetWidth(w)

	// 标题行：全选勾选框 + 勾选/总数计数（antd "x/y items"）。
	checkedN := 0
	for _, it := range items {
		if t.checked[it.Key] {
			checkedN++
		}
	}
	head := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	head.SetWidth(w)
	head.SetPadding(geom.InLTRB(10, 6, 10, 6))
	head.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.FillHover)})
	if t.ShowSelectAll {
		allCb, ok := t.selAlls[side]
		if !ok {
			allCb = NewCheckbox("", false)
			allCb.HitLabel = "transfer-selall-" + side
			t.selAlls[side] = allCb
		}
		allCb.OnChange = func(on bool) { t.selectAll(side, on) }
		checked, indeterminate := t.selAllState(side)
		allCb.Checked = checked
		allCb.Indeterminate = indeterminate
		head.Add(allCb.Build(ctx))
	}
	ht := fmt.Sprintf("%s %d / %d 项", title, checkedN, len(items))
	hl := text.LayoutText(ht, text.DefaultFont(th.Font.SizeBase), text.Options{})
	headText := scene.Text(scene.TextSpec{Src: ht, Layout: hl, Color: th.Color.Text})
	headText.SetHit(scene.HitSpec{Label: "transfer-head-" + side})
	head.Add(headText)
	col.Add(head)

	// 搜索框（antd showSearch）。
	body := scene.Flex(scene.Vertical, 0)
	if t.ShowSearch {
		ph := t.SearchPlaceholder
		if ph == "" {
			ph = "请输入搜索内容"
		}
		si, ok := t.searches[side]
		if !ok {
			si = &Input{HitLabel: "transfer-search-" + side}
			t.searches[side] = si
		}
		si.Placeholder = ph
		si.Value = t.search(side)
		si.OnChanged = func(v string) { t.setSearch(side, v) }
		body.Add(si.Build(ctx))
	}

	// 项列表（搜索过滤后）。
	for _, it := range vis {
		it := it
		body.Add(t.buildItem(ctx, th, it, itemLabel, w))
	}
	body.SetHit(scene.HitSpec{Label: "transfer-list-" + side})
	col.Add(body)
	return col
}

// buildItem 构建列表项：Checkbox 勾选 + 标题（点行整体切换）。
func (t *Transfer) buildItem(ctx *ui.BuildContext, th *theme.Data, it TransferItem,
	itemLabel string, w float32) *scene.Node {
	row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	row.SetWidth(w)
	row.SetHeight(32)
	row.SetPadding(geom.InLTRB(10, 0, 10, 0))

	checked := t.checked[it.Key]
	cb, ok := t.checks[it.Key]
	if !ok {
		cb = NewCheckbox("", false)
		t.checks[it.Key] = cb
	}
	cb.Checked = checked
	cb.OnChange = func(on bool) {
		t.checked[it.Key] = on
		t.MarkDirty()
	}
	cbWrap := scene.Flex(scene.Horizontal, 0)
	cbWrap.Add(cb.Build(ctx))
	row.Add(cbWrap)

	il := text.LayoutText(it.Title, text.DefaultFont(th.Font.SizeBase), text.Options{})
	nameTxt := scene.Text(scene.TextSpec{Src: it.Title, Layout: il, Color: th.Color.Text})
	nameTxt.SetMargin(geom.InLTRB(0, (32-il.Size.H)/2, 0, 0))
	nameWrap := scene.Box(scene.BoxSpec{})
	nameWrap.Add(nameTxt)
	nameWrap.SetFlex(1, scene.FlexTight)
	nameWrap.SetHit(scene.HitSpec{Label: "transfer-name"})
	row.Add(nameWrap)

	row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: itemLabel, Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		ctx.Input.Set(row, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
			if !ev.Down() {
				return false
			}
			// 点行切换勾选（down 即执行——与浮层项语义一致）。
			nc := !t.checked[it.Key]
			t.checked[it.Key] = nc
			t.checks[it.Key].Checked = nc
			t.MarkDirty()
			return true
		}})
	}
	return row
}
