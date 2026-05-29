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

// SwitchI18nLang
//
//	切换本地国际化语言
//	数据结构 : key=value
//		key: 组件名+属性名
//		value: 翻译值
//		panel1.caption=title text
//		submitbtn.caption=ok submit
func SwitchI18nLang(data string) bool {
	return api.SwitchLocalesI18n(data)
}
