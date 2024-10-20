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

// ICustomColorListBox Parent: ICustomListBox
type ICustomColorListBox interface {
	ICustomListBox
	ColorRectWidth() int32                          // property
	SetColorRectWidth(AValue int32)                 // property
	ColorRectOffset() int32                         // property
	SetColorRectOffset(AValue int32)                // property
	StyleForColorBoxStyle() TColorBoxStyle          // property
	SetStyleForColorBoxStyle(AValue TColorBoxStyle) // property
	Colors(Index int32) TColor                      // property
	SetColors(Index int32, AValue TColor)           // property
	ColorNames(Index int32) string                  // property
	SelectedForColor() TColor                       // property
	SetSelectedForColor(AValue TColor)              // property
	DefaultColorColor() TColor                      // property
	SetDefaultColorColor(AValue TColor)             // property
	NoneColorColor() TColor                         // property
	SetNoneColorColor(AValue TColor)                // property
	ColorDialog() IColorDialog                      // property
	SetColorDialog(AValue IColorDialog)             // property
	SetOnGetColors(fn TLBGetColorsEvent)            // property event
}

// TCustomColorListBox Parent: TCustomListBox
type TCustomColorListBox struct {
	TCustomListBox
	getColorsPtr uintptr
}

func NewCustomColorListBox(AOwner IComponent) ICustomColorListBox {
	r1 := customColorListBoxImportAPI().SysCallN(6, GetObjectUintptr(AOwner))
	return AsCustomColorListBox(r1)
}

func (m *TCustomColorListBox) ColorRectWidth() int32 {
	r1 := customColorListBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomColorListBox) SetColorRectWidth(AValue int32) {
	customColorListBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) ColorRectOffset() int32 {
	r1 := customColorListBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomColorListBox) SetColorRectOffset(AValue int32) {
	customColorListBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) StyleForColorBoxStyle() TColorBoxStyle {
	r1 := customColorListBoxImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TColorBoxStyle(r1)
}

func (m *TCustomColorListBox) SetStyleForColorBoxStyle(AValue TColorBoxStyle) {
	customColorListBoxImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) Colors(Index int32) TColor {
	r1 := customColorListBoxImportAPI().SysCallN(5, 0, m.Instance(), uintptr(Index))
	return TColor(r1)
}

func (m *TCustomColorListBox) SetColors(Index int32, AValue TColor) {
	customColorListBoxImportAPI().SysCallN(5, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TCustomColorListBox) ColorNames(Index int32) string {
	r1 := customColorListBoxImportAPI().SysCallN(2, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TCustomColorListBox) SelectedForColor() TColor {
	r1 := customColorListBoxImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorListBox) SetSelectedForColor(AValue TColor) {
	customColorListBoxImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) DefaultColorColor() TColor {
	r1 := customColorListBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorListBox) SetDefaultColorColor(AValue TColor) {
	customColorListBoxImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) NoneColorColor() TColor {
	r1 := customColorListBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorListBox) SetNoneColorColor(AValue TColor) {
	customColorListBoxImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) ColorDialog() IColorDialog {
	r1 := customColorListBoxImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return AsColorDialog(r1)
}

func (m *TCustomColorListBox) SetColorDialog(AValue IColorDialog) {
	customColorListBoxImportAPI().SysCallN(1, 1, m.Instance(), GetObjectUintptr(AValue))
}

func CustomColorListBoxClass() TClass {
	ret := customColorListBoxImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomColorListBox) SetOnGetColors(fn TLBGetColorsEvent) {
	if m.getColorsPtr != 0 {
		RemoveEventElement(m.getColorsPtr)
	}
	m.getColorsPtr = MakeEventDataPtr(fn)
	customColorListBoxImportAPI().SysCallN(10, m.Instance(), m.getColorsPtr)
}

var (
	customColorListBoxImport       *imports.Imports = nil
	customColorListBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomColorListBox_Class", 0),
		/*1*/ imports.NewTable("CustomColorListBox_ColorDialog", 0),
		/*2*/ imports.NewTable("CustomColorListBox_ColorNames", 0),
		/*3*/ imports.NewTable("CustomColorListBox_ColorRectOffset", 0),
		/*4*/ imports.NewTable("CustomColorListBox_ColorRectWidth", 0),
		/*5*/ imports.NewTable("CustomColorListBox_Colors", 0),
		/*6*/ imports.NewTable("CustomColorListBox_Create", 0),
		/*7*/ imports.NewTable("CustomColorListBox_DefaultColorColor", 0),
		/*8*/ imports.NewTable("CustomColorListBox_NoneColorColor", 0),
		/*9*/ imports.NewTable("CustomColorListBox_SelectedForColor", 0),
		/*10*/ imports.NewTable("CustomColorListBox_SetOnGetColors", 0),
		/*11*/ imports.NewTable("CustomColorListBox_StyleForColorBoxStyle", 0),
	}
)

func customColorListBoxImportAPI() *imports.Imports {
	if customColorListBoxImport == nil {
		customColorListBoxImport = NewDefaultImports()
		customColorListBoxImport.SetImportTable(customColorListBoxImportTables)
		customColorListBoxImportTables = nil
	}
	return customColorListBoxImport
}
