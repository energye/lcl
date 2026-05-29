//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package locales

import "C"
import (
	"embed"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/lcl"
)

//如果您需要标准翻译，只需在项目中使用此单元并启用
//项目选项中的i18n。它将自动翻译您的项目。
//
//如果你想自己设置翻译语言，请使用LCLTranslator单元
//并在程序中手动调用SetDefaultLang。

//go:embed lcl
var lclLang embed.FS

func LCLLangConsts(lang string) []byte {
	data, err := lclLang.ReadFile("lcl/strconsts." + lang + ".po")
	if err != nil {
		return nil
	}
	return data
}

// SwitchLCLLang
// 切换LCL常量语言
// lang: de, fr, ja, ko, ru, zh_CN
func SwitchLCLLang(lang string) bool {
	data := LCLLangConsts(lang)
	if data == nil {
		return false
	}
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.Write(mem, data)
	mem.SetPosition(0)
	return api.LocalesUnitResourceStringsFormStream(mem.Instance(), "lclstrconsts")
}

// SwitchDefaultLang
// 切换默认语言（使用嵌入的LCL翻译文件）
// lang: de, fr, ja, ko, ru, zh_CN
// forceUpdate: 强制更新组件语言
func SwitchDefaultLang(lang string, forceUpdate bool) bool {
	if lang == "" {
		api.LocalesSetDefaultLang("", "", "", false)
		return true
	}
	data := LCLLangConsts(lang)
	if data == nil {
		return false
	}
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.Write(mem, data)
	mem.SetPosition(0)
	return api.LocalesSetDefaultLangFormStream(mem.Instance(), "", forceUpdate)
}

func SwitchAppLang(lang string, dir string, localeFileName string, forceUpdate bool) string {
	return api.LocalesSetDefaultLang(lang, dir, localeFileName, forceUpdate)
}

func SwitchAppLangFromStream(data []byte, forceUpdate bool) bool {
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.Write(mem, data)
	mem.SetPosition(0)
	return api.LocalesSetDefaultLangFormStream(mem.Instance(), "", forceUpdate)
}

// SwitchI18nLang 切换应用程序的国际化语言
//
// 通过解析 key=value 格式的翻译数据，自动将翻译值应用到匹配的组件属性上。                                                                                                                                                              //
// 数据格式：
//
//	每行一个翻译条目，格式为 key=value
//	以 # 开头的行为注释，会被忽略
//	空行会被忽略
//	key 大小写不敏感
//
// Key 格式：
//
//  1. 基本格式：组件名.属性名
//     - button1.caption=确定
//     - mainform.hint=这是主窗口
//
//  2. TCollection 集合项：父组件名.集合属性[索引].属性名
//     - listview.columns[0].caption=名称
//     - statusbar.panels[0].text=就绪
//     - headercontrol.sections[0].text=标题
//
//  3. TStrings 集合项：组件名.items[索引]
//     - combobox.items[0]=选项一
//     - radiogroup.items[0]=男
//     - checkgroup.items[0]=自动保存
//
//  4. SubComponent 子组件：父组件名.子组件属性名.属性名
//     - labelededit.editlabel.caption=用户名：
//     - buttonpanel.okbutton.caption=确定
//
// 支持的属性：
//
//   - Caption: 标题/文本（按钮、标签、菜单、窗体等）
//   - Hint: 提示信息（所有可视组件）
//   - Text: 文本内容（编辑框、备忘录）
//   - TextHint: 占位符提示（编辑框、下拉框）
//   - Title: 对话框标题（TDialog、TTaskDialog 等）
//   - SimpleText: 状态栏简单文本（TStatusBar）
//   - Filter: 文件对话框过滤器（TOpenDialog、TSaveDialog）
//   - ButtonCaption/ButtonHint: 按钮标题/提示（TEditButton）
//   - DialogTitle: 对话框标题（TFileNameEdit、TDirectoryEdit）
//   - FindText/ReplaceText: 查找/替换文本（TFindDialog、TReplaceDialog）
//   - OKCaption/CancelCaption: 按钮文本（TCalendarDialog）
//   - BalloonHint/BalloonTitle: 气泡提示（TTrayIcon）
//
// 示例：
//
//	data := `
//	# 窗体
//	mainform.caption=我的应用
//	# 按钮
//	submitbtn.caption=提交
//	submitbtn.hint=点击提交
//	# 下拉框选项
//	languagecombo.items[0]=中文
//	languagecombo.items[1]=英文
//	# 列表头
//	listview.columns[0].caption=名称
//	# 子组件标签
//	usernameedit.editlabel.caption=用户名：
//	`
//	locales.SwitchI18nLang(data)
//
// 参数：
//
//	data - 翻译数据，key=value 格式的多行文本
//
// 返回值：
//
//	true - 成功应用翻译
//	false - 解析失败
func SwitchI18nLang(data string) bool {
	return api.SwitchLocalesI18n(data)
}
