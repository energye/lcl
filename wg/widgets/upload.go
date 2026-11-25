package widgets

// ── Upload 上传（M12 高级录入）────────────────────────────────────
//
// antd：<Upload action fileList onChange> 默认形态：
// [上传按钮 icon+] + 提示文字 + 文件列表（name + 状态 + 删除）。
// 引擎无网络层/文件对话框：组件只做 **UI 形态与状态机**——文件列表
// 由宿主经 AddFile/Remove/SetPercent/FinishUpload 驱动（接平台文件
// 对话框与真实上传）；OnPick 在触发按钮点击时回调（宿主接对话框）。
//
// 组合复用：Button（触发按钮）、Icon（upload/delete/check-circle/
// close-circle）、Progress（uploading 项内嵌线形进度条）。
// 状态机：uploading（进度条+百分比）→ done（绿勾）/ error（红叉）。
// Dragger（第一百三十一轮，antd type="drag"）：触发区渲染为虚线描边
// 圆角大框（antd dragger 令牌：104px 高、虚线边框、内含大图标+主行
// 文字+副行提示），点击仍走 OnPick（宿主接文件对话框）。OS 文件拖放
// 事件引擎无原语（如实标注）——drop 注入待引擎层支持。
// 未实现（如实标注）：OS 文件拖入事件、picture-list 卡片形态、
// beforeUpload/自定义 request、目录上传、maxCount 截断、
// showUploadList 开关、受控 fileList 之外的内部上传引擎。

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// UploadFileStatus 文件项状态（antd status）。
const (
	UploadUploading = "uploading"
	UploadDone      = "done"
	UploadError     = "error"
)

// UploadFile 是文件列表项（antd fileList 项的 UI 子集）。
type UploadFile struct {
	UID     string
	Name    string
	Status  string // UploadUploading / UploadDone / UploadError
	Percent int    // uploading 时 0-100
}

// Upload 是上传组件（触发区+文件列表状态机）。
type Upload struct {
	controlBase
	Files      []UploadFile
	ButtonText string // 默认"点击上传"
	Hint       string // 按钮旁提示文字（antd 默认空）
	Disabled   bool
	// Dragger 拖拽区形态（antd type="drag"）：触发区渲染为虚线描边大框
	//（inbox 大图标+主行文字+Hint 副行），点击仍走 OnPick。
	Dragger bool
	// OnPick 触发按钮点击（宿主接平台文件对话框；nil 则无操作）。
	OnPick func()
	// OnChange 文件列表变化（Add/Remove/进度/完成都触发）。
	OnChange func(files []UploadFile)
	Width    float32 // 列表/行宽（默认 320）

	// removeTaps 保留 GestureDetector 引用（防 GC 与重建失联——先例）。
	removeTaps []*ui.GestureDetector
	// trigger 触发按钮持久实例（Build 内新建实例会让按压态随重建丢失
	// ——教训⑨第四次现身，真窗 heuristic press 反馈为 0 实证）。
	trigger *Button
	// draggerTaps Dragger 框点击手势（跨 Build 持有）。
	draggerTaps []*ui.GestureDetector
}

// NewUpload 创建上传组件。
func NewUpload() *Upload {
	return &Upload{ButtonText: "点击上传", Width: 320}
}

// notify 列表快照回调。
func (u *Upload) notify() {
	if u.OnChange != nil {
		u.OnChange(u.Files)
	}
}

// AddFile 追加文件项并回调。
func (u *Upload) AddFile(f UploadFile) {
	u.Files = append(u.Files, f)
	u.MarkDirty()
	u.notify()
}

// Remove 删除指定项并回调。
func (u *Upload) Remove(uid string) {
	for i, f := range u.Files {
		if f.UID == uid {
			u.Files = append(u.Files[:i], u.Files[i+1:]...)
			break
		}
	}
	u.MarkDirty()
	u.notify()
}

// SetPercent 推进上传进度（0-100）并回调。
func (u *Upload) SetPercent(uid string, percent int) {
	for i := range u.Files {
		if u.Files[i].UID == uid {
			u.Files[i].Status = UploadUploading
			u.Files[i].Percent = percent
			break
		}
	}
	u.MarkDirty()
	u.notify()
}

// FinishUpload 完成/失败某项（status=done|error）并回调。
func (u *Upload) FinishUpload(uid string, status string) {
	for i := range u.Files {
		if u.Files[i].UID == uid {
			u.Files[i].Status = status
			break
		}
	}
	u.MarkDirty()
	u.notify()
}

// Build 实现 ui.Widget：触发区（按钮形态/Dragger 拖拽区形态）+ 文件列表。
func (u *Upload) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(u)
	th := ctx.ThemeOr()
	w := u.Width
	if w <= 0 {
		w = 320
	}
	col := scene.Flex(scene.Vertical, th.Space.XS)

	// Dragger 拖拽区形态（antd type="drag"）：虚线描边圆角大框——
	// inbox 大图标+主行文字+Hint 副行；点击走 OnPick。
	if u.Dragger {
		col.Add(u.buildDragger(ctx, th, w))
	} else {
		// 触发行：上传按钮（组合复用 Button+upload 图标）+ 提示文字。
		row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
		if u.trigger == nil {
			u.trigger = &Button{Icon: &Icon{Name: "upload", Size: 14}}
		}
		// 字段同步（持久实例跨 Build 复用——按压态不丢）。
		u.trigger.Text = u.ButtonText
		u.trigger.Disabled = u.Disabled
		u.trigger.OnClick = func() {
			if u.OnPick != nil {
				u.OnPick()
			}
		}
		btnWrap := scene.Flex(scene.Horizontal, 0)
		btnWrap.Add(u.trigger.Build(ctx))
		// 被动锚：Button 自带 hit label 会撞全页 button（锚唯一化）。
		btnWrap.SetHit(scene.HitSpec{Label: "upload-trigger"})
		row.Add(btnWrap)
		if u.Hint != "" {
			hl := text.LayoutText(u.Hint, text.DefaultFont(th.Font.SizeBase), text.Options{})
			hint := scene.Text(scene.TextSpec{Src: u.Hint, Layout: hl, Color: th.Color.TextFaint})
			hint.SetHit(scene.HitSpec{Label: "upload-hint"})
			row.Add(hint)
		}
		col.Add(row)
	}

	// 文件列表。
	for _, f := range u.Files {
		f := f
		col.Add(u.buildItem(ctx, th, f, w))
	}
	return col
}

// buildDragger 构建拖拽区大框（antd dragger 令牌：高 104、虚线边框、
// hover 边框主色——hover 原地迁移；内含 inbox 图标+主行+副行）。
func (u *Upload) buildDragger(ctx *ui.BuildContext, th *theme.Data, w float32) *scene.Node {
	main := u.ButtonText
	if main == "" || main == "点击上传" {
		main = "点击或拖拽文件到此区域上传"
	}
	// 虚线描边（antd dragger 形态——render.Dashed 成对实/空段）。
	box := scene.Box(scene.BoxSpec{
		Border: render.Dashed(th.Color.Border, 1, 4, 4),
		Radius: th.Radius.Base,
	}).SetSize2(w, 104)
	iconCol := scene.Flex(scene.Vertical, 4).SetCrossAlign(scene.CrossCenter)
	ic := (&Icon{Name: "upload", Size: 28, Color: th.Color.TextFaint}).Build(ctx)
	iconCol.Add(ic)
	ml := text.LayoutText(main, text.DefaultFont(th.Font.SizeBase), text.Options{})
	mt := scene.Text(scene.TextSpec{Src: main, Layout: ml, Color: th.Color.Text})
	iconCol.Add(mt)
	if u.Hint != "" {
		hl := text.LayoutText(u.Hint, text.DefaultFont(th.Font.SizeXS), text.Options{})
		ht := scene.Text(scene.TextSpec{Src: u.Hint, Layout: hl, Color: th.Color.TextFaint})
		iconCol.Add(ht)
	}
	// 居中内容：图标列 Stack 定位框中心。
	iconCol.SetPositionKeepSize(true)
	cw, ch := iconCol.Size().W, iconCol.Size().H
	iconCol.SetStackPosition(geom.Pt((w-cw)/2, (104-ch)/2))
	box.Add(iconCol)
	box.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "upload-dragger", Cursor: scene.CursorPointer, Disabled: u.Disabled})
	u.draggerTaps = u.draggerTaps[:0]
	if ctx.Input != nil {
		g := &ui.GestureDetector{OnTap: func(geom.Point) {
			if !u.Disabled && u.OnPick != nil {
				u.OnPick()
			}
		}}
		u.draggerTaps = append(u.draggerTaps, g)
		ctx.Input.Set(box, g.Handler())
	}
	return box
}

// buildItem 构建单个文件项（上行 name+状态图标+删除；下行 uploading 进度条）。
func (u *Upload) buildItem(ctx *ui.BuildContext, th *theme.Data, f UploadFile, w float32) *scene.Node {
	itemCol := scene.Flex(scene.Vertical, 2)
	row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	row.SetWidth(w)

	nameColor := th.Color.Text
	if f.Status == UploadError {
		nameColor = th.Color.Danger
	}
	nl := text.LayoutText(f.Name, text.DefaultFont(th.Font.SizeBase), text.Options{})
	nameTxt := scene.Text(scene.TextSpec{Src: f.Name, Layout: nl, Color: nameColor})
	nameTxt.SetHit(scene.HitSpec{Label: "upload-name"})
	nameBox := scene.Box(scene.BoxSpec{})
	nameBox.Add(nameTxt)
	nameBox.SetFlex(1, scene.FlexTight)
	row.Add(nameBox)

	switch f.Status {
	case UploadDone:
		ic := (&Icon{Name: "check-circle", Size: 14, Color: th.Color.Success}).Build(ctx)
		row.Add(ic)
	case UploadError:
		ic := (&Icon{Name: "close-circle", Size: 14, Color: th.Color.Danger}).Build(ctx)
		row.Add(ic)
	case UploadUploading:
		pl := text.LayoutText(fmt.Sprintf("%d%%", f.Percent), text.DefaultFont(th.Font.SizeBase), text.Options{})
		row.Add(scene.Text(scene.TextSpec{Src: fmt.Sprintf("%d%%", f.Percent), Layout: pl, Color: th.Color.TextMuted}))
	}

	// 删除 X（down 即执行——本组件列表在内容树，tap 配对可靠，但统一
	// down 语义与浮层项一致；此处用 GestureDetector 与 Button 同款）。
	del := (&Icon{Name: "delete", Size: 13, Color: th.Color.TextFaint}).Build(ctx)
	del.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "upload-remove", Cursor: scene.CursorPointer})
	g := &ui.GestureDetector{OnTap: func(geom.Point) {
		u.Remove(f.UID)
	}}
	u.removeTaps = append(u.removeTaps, g)
	if ctx.Input != nil {
		ctx.Input.Set(del, g.Handler())
	}
	row.Add(del)
	itemCol.Add(row)

	if f.Status == UploadUploading {
		pb := Progress{Value: float32(f.Percent) / 100, Width: w - 24, Height: 4, ShowInfo: false}
		itemCol.Add(pb.Build(ctx))
	}
	itemCol.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "upload-item"})
	return itemCol
}
