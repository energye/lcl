package widgets

// ── Typography copyable（antd Typography copyable）──────────────────
//
// 文本右侧显示复制图标（hover 可见），点击后短暂变勾（已复制态）。
// **wg 无剪贴板通道**：组件只负责 UI 与"复制动作"语义，实际写入剪贴板
// 由宿主在 OnCopy 回调里完成（与 Upload.OnPick 接文件对话框同款职责切分）。
//
// 未实现（如实标注）：tooltips 双态文案（Copy/Copied 气泡）、format
// text/html、自定义 icon 对、actions placement=start。

import (
	"time"

	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// CopyableText 是可复制的内联文本（antd Typography copyable）。
type CopyableText struct {
	controlBase
	Body string
	// Type 语义色（同 Text.Type）。
	Type TextType
	// CopyText 复制内容（空 = 用 Body）。
	CopyText string
	// OnCopy 复制回调（宿主在此写入剪贴板）。
	OnCopy func(text string)

	// copied 已复制态（点击后图标变 check，resetDelay 后复位）。
	copied bool
	icon   *scene.Node
}

// NewCopyableText 创建可复制文本。
func NewCopyableText(body string) *CopyableText { return &CopyableText{Body: body} }

// Build 实现 ui.Widget：文本 + 右侧 copy 图标。
func (c *CopyableText) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(c)
	th := ctx.ThemeOr()
	color := Text{Type: c.Type}.color(th)
	font := text.DefaultFont(th.Font.SizeBase)
	layout := text.LayoutText(c.Body, font, text.Options{LineHeight: 1.57})
	txt := scene.Text(scene.TextSpec{Src: c.Body, Layout: layout, Color: color})
	txt.SetHit(scene.HitSpec{Label: "typography"})

	row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	row.Add(txt)

	iconName := "copy"
	iconColor := th.Color.TextMuted
	if c.copied {
		iconName = "check"
		iconColor = th.Color.Success
	}
	ic := (&Icon{Name: iconName, Size: 14, Color: iconColor}).Build(ctx)
	c.icon = ic
	ic.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "typography-copy", Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		ctx.Input.Set(ic, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
			if !ev.Down() {
				return false
			}
			c.doCopy()
			return true
		}})
	}
	row.Add(ic)
	row.SetHit(scene.HitSpec{Label: "typography-copyable"})
	return row
}

// doCopy 执行复制：触发 OnCopy、进入已复制态、短暂后复位。
func (c *CopyableText) doCopy() {
	payload := c.CopyText
	if payload == "" {
		payload = c.Body
	}
	if c.OnCopy != nil {
		c.OnCopy(payload)
	}
	c.copied = true
	c.MarkDirty()
	time.AfterFunc(1200*time.Millisecond, func() {
		c.copied = false
		c.MarkDirty()
	})
}

// Copied 返回当前是否处于已复制态（诊断/断言用）。
func (c *CopyableText) Copied() bool { return c.copied }
