package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	upBase    = widgets.NewUpload()
	upProg    = widgets.NewUpload()
	upDragger = widgets.NewUpload()
	upInit    bool
)

func init() {
	upBase.Files = []widgets.UploadFile{
		{UID: "f1", Name: "site-design.png", Status: widgets.UploadDone},
		{UID: "f2", Name: "draft.psd", Status: widgets.UploadError},
	}
	upBase.Hint = "支持单次或批量上传"
	upProg.Files = []widgets.UploadFile{
		{UID: "p1", Name: "logo.mov", Status: widgets.UploadUploading, Percent: 60},
	}
	upDragger.Dragger = true
	upDragger.Hint = "单次不超过 10MB"
}

func main() {
	launcher.Main(doc.Meta{
		ID: "upload", Title: "Upload 上传", Group: doc.GroupControl,
		What: []string{
			"触发区（Button 组合复用+upload 图标+提示文字）+ 文件列表（name+状态+删除）",
			"状态机：uploading（项内嵌 Progress 线形进度条组合复用+百分比）→ done（绿勾）/ error（红叉红名）",
			"**Dragger 拖拽区形态**（antd type=drag）：虚线描边 104px 大框（upload 大图标+主行+副行），点击走 OnPick；OS 文件拖入事件引擎无原语（如实标注）",
			"引擎无网络层：UI 形态与状态机在组件内，文件列表由宿主 AddFile/Remove/SetPercent/FinishUpload 驱动（接平台文件对话框与真实上传）",
		},
		Demos: []doc.Demo{
			{Title: "基础", Note: "done（绿勾）/ error（红叉红名）两种状态项；点删除 X 移除", Build: func(ctx *ui.BuildContext) ui.Widget {
				return upBase
			}},
			{Title: "上传中", Note: "uploading 60%：项内嵌进度条+百分比文字", Build: func(ctx *ui.BuildContext) ui.Widget {
				if !upInit {
					upInit = true
				}
				return upProg
			}},
			{Title: "拖拽区", Note: "Dragger：虚线描边大框+大图标+「点击或拖拽文件到此区域上传」+副行提示；点击回显", Build: func(ctx *ui.BuildContext) ui.Widget {
				upDragger.OnPick = func() {
					upDragger.AddFile(widgets.UploadFile{UID: "d1", Name: "dropped.txt", Status: widgets.UploadDone})
				}
				return upDragger
			}},
		},
		API: []doc.APIRow{
			{Name: "NewUpload", Type: "func() *Upload", Desc: "构造（默认按钮\"点击上传\"）"},
			{Name: "Files", Type: "[]UploadFile", Desc: "文件列表（UID/Name/Status/Percent）"},
			{Name: "Dragger", Type: "bool", Default: "false", Desc: "拖拽区形态（antd type=drag——虚线描边大框）"},
			{Name: "ButtonText / Hint", Type: "string", Default: "点击上传 / 空", Desc: "按钮文案 / 提示文字"},
			{Name: "OnPick", Type: "func()", Desc: "触发按钮/拖拽区点击（宿主接平台文件对话框）"},
			{Name: "OnChange", Type: "func(files []UploadFile)", Desc: "列表变化回调（增删/进度/完成）"},
			{Name: "AddFile / Remove", Type: "func(...)", Desc: "追加 / 删除文件项"},
			{Name: "SetPercent / FinishUpload", Type: "func(uid, ...)", Desc: "推进进度 / 完成（done|error）"},
			{Name: "Width", Type: "float32", Default: "320", Desc: "行宽"},
		},
	})
}
