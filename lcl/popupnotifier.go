//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// IPopupNotifier Parent: IComponent
type IPopupNotifier interface {
	IComponent
	Hide()                       // procedure
	Show()                       // procedure
	ShowAtPos(X int32, Y int32)  // procedure
	Color() types.TColor         // property Color Getter
	SetColor(value types.TColor) // property Color Setter
	Icon() IPicture              // property Icon Getter
	SetIcon(value IPicture)      // property Icon Setter
	Text() string                // property Text Getter
	SetText(value string)        // property Text Setter
	TextFont() IFont             // property TextFont Getter
	SetTextFont(value IFont)     // property TextFont Setter
	Title() string               // property Title Getter
	SetTitle(value string)       // property Title Setter
	TitleFont() IFont            // property TitleFont Getter
	SetTitleFont(value IFont)    // property TitleFont Setter
	Visible() bool               // property Visible Getter
	SetVisible(value bool)       // property Visible Setter
	SetOnClose(fn TCloseEvent)   // property event
}

type TPopupNotifier struct {
	TComponent
}

func (m *TPopupNotifier) Hide() {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(1, m.Instance())
}

func (m *TPopupNotifier) Show() {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(2, m.Instance())
}

func (m *TPopupNotifier) ShowAtPos(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(3, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TPopupNotifier) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := popupNotifierAPI().SysCallN(4, 0, m.Instance())
	return types.TColor(r)
}

func (m *TPopupNotifier) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TPopupNotifier) Icon() IPicture {
	if !m.IsValid() {
		return nil
	}
	r := popupNotifierAPI().SysCallN(5, 0, m.Instance())
	return AsPicture(r)
}

func (m *TPopupNotifier) SetIcon(value IPicture) {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPopupNotifier) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := popupNotifierAPI().SysCallN(6, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TPopupNotifier) SetText(value string) {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(6, 1, m.Instance(), api.PasStr(value))
}

func (m *TPopupNotifier) TextFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := popupNotifierAPI().SysCallN(7, 0, m.Instance())
	return AsFont(r)
}

func (m *TPopupNotifier) SetTextFont(value IFont) {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPopupNotifier) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := popupNotifierAPI().SysCallN(8, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TPopupNotifier) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(8, 1, m.Instance(), api.PasStr(value))
}

func (m *TPopupNotifier) TitleFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := popupNotifierAPI().SysCallN(9, 0, m.Instance())
	return AsFont(r)
}

func (m *TPopupNotifier) SetTitleFont(value IFont) {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPopupNotifier) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := popupNotifierAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPopupNotifier) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	popupNotifierAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TPopupNotifier) SetOnClose(fn TCloseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCloseEvent(fn)
	base.SetEvent(m, 11, popupNotifierAPI(), api.MakeEventDataPtr(cb))
}

// NewPopupNotifier class constructor
func NewPopupNotifier(owner IComponent) IPopupNotifier {
	r := popupNotifierAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsPopupNotifier(r)
}

func TPopupNotifierClass() types.TClass {
	r := popupNotifierAPI().SysCallN(12)
	return types.TClass(r)
}

var (
	popupNotifierOnce   base.Once
	popupNotifierImport *imports.Imports = nil
)

func popupNotifierAPI() *imports.Imports {
	popupNotifierOnce.Do(func() {
		popupNotifierImport = api.NewDefaultImports()
		popupNotifierImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPopupNotifier_Create", 0), // constructor NewPopupNotifier
			/* 1 */ imports.NewTable("TPopupNotifier_Hide", 0), // procedure Hide
			/* 2 */ imports.NewTable("TPopupNotifier_Show", 0), // procedure Show
			/* 3 */ imports.NewTable("TPopupNotifier_ShowAtPos", 0), // procedure ShowAtPos
			/* 4 */ imports.NewTable("TPopupNotifier_Color", 0), // property Color
			/* 5 */ imports.NewTable("TPopupNotifier_Icon", 0), // property Icon
			/* 6 */ imports.NewTable("TPopupNotifier_Text", 0), // property Text
			/* 7 */ imports.NewTable("TPopupNotifier_TextFont", 0), // property TextFont
			/* 8 */ imports.NewTable("TPopupNotifier_Title", 0), // property Title
			/* 9 */ imports.NewTable("TPopupNotifier_TitleFont", 0), // property TitleFont
			/* 10 */ imports.NewTable("TPopupNotifier_Visible", 0), // property Visible
			/* 11 */ imports.NewTable("TPopupNotifier_OnClose", 0), // event OnClose
			/* 12 */ imports.NewTable("TPopupNotifier_TClass", 0), // function TPopupNotifierClass
		}
	})
	return popupNotifierImport
}
