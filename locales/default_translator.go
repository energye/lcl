//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package locales

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
// 切换默认语言
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
