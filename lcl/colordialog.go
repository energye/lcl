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

// IColorDialog Parent: ICommonDialog
type IColorDialog interface {
	ICommonDialog
	Color() types.TColor         // property Color Getter
	SetColor(value types.TColor) // property Color Setter
	// CustomColors
	//  entry looks like ColorA = FFFF00 ... ColorX = C0C0C0
	CustomColors() IStrings                     // property CustomColors Getter
	SetCustomColors(value IStrings)             // property CustomColors Setter
	Options() types.TColorDialogOptions         // property Options Getter
	SetOptions(value types.TColorDialogOptions) // property Options Setter
}

type TColorDialog struct {
	TCommonDialog
}

func (m *TColorDialog) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := colorDialogAPI().SysCallN(1, 0, m.Instance())
	return types.TColor(r)
}

func (m *TColorDialog) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	colorDialogAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TColorDialog) CustomColors() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := colorDialogAPI().SysCallN(2, 0, m.Instance())
	return AsStrings(r)
}

func (m *TColorDialog) SetCustomColors(value IStrings) {
	if !m.IsValid() {
		return
	}
	colorDialogAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TColorDialog) Options() types.TColorDialogOptions {
	if !m.IsValid() {
		return 0
	}
	r := colorDialogAPI().SysCallN(3, 0, m.Instance())
	return types.TColorDialogOptions(r)
}

func (m *TColorDialog) SetOptions(value types.TColorDialogOptions) {
	if !m.IsValid() {
		return
	}
	colorDialogAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

// NewColorDialog class constructor
func NewColorDialog(theOwner IComponent) IColorDialog {
	r := colorDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsColorDialog(r)
}

func TColorDialogClass() types.TClass {
	r := colorDialogAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	colorDialogOnce   base.Once
	colorDialogImport *imports.Imports = nil
)

func colorDialogAPI() *imports.Imports {
	colorDialogOnce.Do(func() {
		colorDialogImport = api.NewDefaultImports()
		colorDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TColorDialog_Create", 0), // constructor NewColorDialog
			/* 1 */ imports.NewTable("TColorDialog_Color", 0), // property Color
			/* 2 */ imports.NewTable("TColorDialog_CustomColors", 0), // property CustomColors
			/* 3 */ imports.NewTable("TColorDialog_Options", 0), // property Options
			/* 4 */ imports.NewTable("TColorDialog_TClass", 0), // function TColorDialogClass
		}
	})
	return colorDialogImport
}
