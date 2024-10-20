//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// ICustomStaticText Parent: IWinControl
type ICustomStaticText interface {
	IWinControl
	Alignment() TAlignment                    // property
	SetAlignment(AValue TAlignment)           // property
	BorderStyle() TStaticBorderStyle          // property
	SetBorderStyle(AValue TStaticBorderStyle) // property
	FocusControl() IWinControl                // property
	SetFocusControl(AValue IWinControl)       // property
	ShowAccelChar() bool                      // property
	SetShowAccelChar(AValue bool)             // property
	Transparent() bool                        // property
	SetTransparent(AValue bool)               // property
}

// TCustomStaticText Parent: TWinControl
type TCustomStaticText struct {
	TWinControl
}

func NewCustomStaticText(AOwner IComponent) ICustomStaticText {
	r1 := customStaticTextImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsCustomStaticText(r1)
}

func (m *TCustomStaticText) Alignment() TAlignment {
	r1 := customStaticTextImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomStaticText) SetAlignment(AValue TAlignment) {
	customStaticTextImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomStaticText) BorderStyle() TStaticBorderStyle {
	r1 := customStaticTextImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TStaticBorderStyle(r1)
}

func (m *TCustomStaticText) SetBorderStyle(AValue TStaticBorderStyle) {
	customStaticTextImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomStaticText) FocusControl() IWinControl {
	r1 := customStaticTextImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TCustomStaticText) SetFocusControl(AValue IWinControl) {
	customStaticTextImportAPI().SysCallN(4, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomStaticText) ShowAccelChar() bool {
	r1 := customStaticTextImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStaticText) SetShowAccelChar(AValue bool) {
	customStaticTextImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomStaticText) Transparent() bool {
	r1 := customStaticTextImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStaticText) SetTransparent(AValue bool) {
	customStaticTextImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func CustomStaticTextClass() TClass {
	ret := customStaticTextImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	customStaticTextImport       *imports.Imports = nil
	customStaticTextImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomStaticText_Alignment", 0),
		/*1*/ imports.NewTable("CustomStaticText_BorderStyle", 0),
		/*2*/ imports.NewTable("CustomStaticText_Class", 0),
		/*3*/ imports.NewTable("CustomStaticText_Create", 0),
		/*4*/ imports.NewTable("CustomStaticText_FocusControl", 0),
		/*5*/ imports.NewTable("CustomStaticText_ShowAccelChar", 0),
		/*6*/ imports.NewTable("CustomStaticText_Transparent", 0),
	}
)

func customStaticTextImportAPI() *imports.Imports {
	if customStaticTextImport == nil {
		customStaticTextImport = NewDefaultImports()
		customStaticTextImport.SetImportTable(customStaticTextImportTables)
		customStaticTextImportTables = nil
	}
	return customStaticTextImport
}
