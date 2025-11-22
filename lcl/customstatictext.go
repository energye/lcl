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

// ICustomStaticText Parent: IWinControl
type ICustomStaticText interface {
	IWinControl
	Alignment() types.TAlignment                   // property Alignment Getter
	SetAlignment(value types.TAlignment)           // property Alignment Setter
	BorderStyle() types.TStaticBorderStyle         // property BorderStyle Getter
	SetBorderStyle(value types.TStaticBorderStyle) // property BorderStyle Setter
	FocusControl() IWinControl                     // property FocusControl Getter
	SetFocusControl(value IWinControl)             // property FocusControl Setter
	ShowAccelChar() bool                           // property ShowAccelChar Getter
	SetShowAccelChar(value bool)                   // property ShowAccelChar Setter
	Transparent() bool                             // property Transparent Getter
	SetTransparent(value bool)                     // property Transparent Setter
}

type TCustomStaticText struct {
	TWinControl
}

func (m *TCustomStaticText) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := customStaticTextAPI().SysCallN(1, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TCustomStaticText) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	customStaticTextAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCustomStaticText) BorderStyle() types.TStaticBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := customStaticTextAPI().SysCallN(2, 0, m.Instance())
	return types.TStaticBorderStyle(r)
}

func (m *TCustomStaticText) SetBorderStyle(value types.TStaticBorderStyle) {
	if !m.IsValid() {
		return
	}
	customStaticTextAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomStaticText) FocusControl() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := customStaticTextAPI().SysCallN(3, 0, m.Instance())
	return AsWinControl(r)
}

func (m *TCustomStaticText) SetFocusControl(value IWinControl) {
	if !m.IsValid() {
		return
	}
	customStaticTextAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomStaticText) ShowAccelChar() bool {
	if !m.IsValid() {
		return false
	}
	r := customStaticTextAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomStaticText) SetShowAccelChar(value bool) {
	if !m.IsValid() {
		return
	}
	customStaticTextAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomStaticText) Transparent() bool {
	if !m.IsValid() {
		return false
	}
	r := customStaticTextAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomStaticText) SetTransparent(value bool) {
	if !m.IsValid() {
		return
	}
	customStaticTextAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

// NewCustomStaticText class constructor
func NewCustomStaticText(owner IComponent) ICustomStaticText {
	r := customStaticTextAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomStaticText(r)
}

func TCustomStaticTextClass() types.TClass {
	r := customStaticTextAPI().SysCallN(6)
	return types.TClass(r)
}

var (
	customStaticTextOnce   base.Once
	customStaticTextImport *imports.Imports = nil
)

func customStaticTextAPI() *imports.Imports {
	customStaticTextOnce.Do(func() {
		customStaticTextImport = api.NewDefaultImports()
		customStaticTextImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomStaticText_Create", 0), // constructor NewCustomStaticText
			/* 1 */ imports.NewTable("TCustomStaticText_Alignment", 0), // property Alignment
			/* 2 */ imports.NewTable("TCustomStaticText_BorderStyle", 0), // property BorderStyle
			/* 3 */ imports.NewTable("TCustomStaticText_FocusControl", 0), // property FocusControl
			/* 4 */ imports.NewTable("TCustomStaticText_ShowAccelChar", 0), // property ShowAccelChar
			/* 5 */ imports.NewTable("TCustomStaticText_Transparent", 0), // property Transparent
			/* 6 */ imports.NewTable("TCustomStaticText_TClass", 0), // function TCustomStaticTextClass
		}
	})
	return customStaticTextImport
}
