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

// IColorDialog Parent: ICommonDialog
type IColorDialog interface {
	ICommonDialog
	Color() TColor                   // property
	SetColor(AValue TColor)          // property
	CustomColors() IStrings          // property
	SetCustomColors(AValue IStrings) // property
}

// TColorDialog Parent: TCommonDialog
type TColorDialog struct {
	TCommonDialog
}

func NewColorDialog(TheOwner IComponent) IColorDialog {
	r1 := colorDialogImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsColorDialog(r1)
}

func (m *TColorDialog) Color() TColor {
	r1 := colorDialogImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TColorDialog) SetColor(AValue TColor) {
	colorDialogImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorDialog) CustomColors() IStrings {
	r1 := colorDialogImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TColorDialog) SetCustomColors(AValue IStrings) {
	colorDialogImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func ColorDialogClass() TClass {
	ret := colorDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	colorDialogImport       *imports.Imports = nil
	colorDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ColorDialog_Class", 0),
		/*1*/ imports.NewTable("ColorDialog_Color", 0),
		/*2*/ imports.NewTable("ColorDialog_Create", 0),
		/*3*/ imports.NewTable("ColorDialog_CustomColors", 0),
	}
)

func colorDialogImportAPI() *imports.Imports {
	if colorDialogImport == nil {
		colorDialogImport = NewDefaultImports()
		colorDialogImport.SetImportTable(colorDialogImportTables)
		colorDialogImportTables = nil
	}
	return colorDialogImport
}
