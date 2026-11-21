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
// 未实现（如实标注）：showSearch 搜索框、拖拽排序、分页、
// footer 插槽、render 自定义行、header 全选 checkbox（计数行）。

import (
	"fmt"

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
	TransferLeft                            // 右→左（移出）
)

// Transfer 是穿梭框。
type Transfer struct {
	controlBase
	DataSource []TransferItem
	TargetKeys []string // 右栏的 Key 集
	OnChange   func(targetKeys []string, dir TransferDirection, moveKeys []string)
	Width      float32 // 单栏宽（默认 220）

	checked map[string]bool
	// checks 池化复选框实例（跨 Build 持有——hover 态不丢，Tree.Checkable 先例）。
	checks map[string]*Checkbox
	// opBtns 池化操作按钮（Build 内新建实例按压态随重建丢失——教训⑨，Upload 先例）。
	opBtns map[string]*Button
}

// NewTransfer 创建穿梭框。
func NewTransfer(items []TransferItem) *Transfer {
	return &Transfer{DataSource: items, Width: 220,
		checked: map[string]bool{}, checks: map[string]*Checkbox{},
		opBtns: map[string]*Button{}}
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

// buildSide 构建一栏（标题行+项列表）。
func (t *Transfer) buildSide(ctx *ui.BuildContext, th *theme.Data, side, title, itemLabel string,
	items []TransferItem, w float32) *scene.Node {
	col := scene.Flex(scene.Vertical, 0)
	col.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	col.SetWidth(w)

	// 标题行：勾选/总数计数（antd "x/y items"）。
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
	ht := fmt.Sprintf("%s %d / %d 项", title, checkedN, len(items))
	hl := text.LayoutText(ht, text.DefaultFont(th.Font.SizeSM), text.Options{})
	headText := scene.Text(scene.TextSpec{Src: ht, Layout: hl, Color: th.Color.Text})
	headText.SetHit(scene.HitSpec{Label: "transfer-head-" + side})
	head.Add(headText)
	col.Add(head)

	// 项列表。
	body := scene.Flex(scene.Vertical, 0)
	for _, it := range items {
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
