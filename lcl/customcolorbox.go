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

// ICustomColorBox Parent: ICustomComboBox
type ICustomColorBox interface {
	ICustomComboBox
	ColorRectWidth() int32                          // property
	SetColorRectWidth(AValue int32)                 // property
	ColorRectOffset() int32                         // property
	SetColorRectOffset(AValue int32)                // property
	StyleForColorBoxStyle() TColorBoxStyle          // property
	SetStyleForColorBoxStyle(AValue TColorBoxStyle) // property
	Colors(Index int32) TColor                      // property
	ColorNames(Index int32) string                  // property
	Selected() TColor                               // property
	SetSelected(AValue TColor)                      // property
	DefaultColorColor() TColor                      // property
	SetDefaultColorColor(AValue TColor)             // property
	NoneColorColor() TColor                         // property
	SetNoneColorColor(AValue TColor)                // property
	ColorDialog() IColorDialog                      // property
	SetColorDialog(AValue IColorDialog)             // property
	SetOnGetColors(fn TGetColorsEvent)              // property event
}

// TCustomColorBox Parent: TCustomComboBox
type TCustomColorBox struct {
	TCustomComboBox
	getColorsPtr uintptr
}

func NewCustomColorBox(AOwner IComponent) ICustomColorBox {
	r1 := customColorBoxImportAPI().SysCallN(6, GetObjectUintptr(AOwner))
	return AsCustomColorBox(r1)
}

func (m *TCustomColorBox) ColorRectWidth() int32 {
	r1 := customColorBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomColorBox) SetColorRectWidth(AValue int32) {
	customColorBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) ColorRectOffset() int32 {
	r1 := customColorBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomColorBox) SetColorRectOffset(AValue int32) {
	customColorBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) StyleForColorBoxStyle() TColorBoxStyle {
	r1 := customColorBoxImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TColorBoxStyle(r1)
}

func (m *TCustomColorBox) SetStyleForColorBoxStyle(AValue TColorBoxStyle) {
	customColorBoxImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) Colors(Index int32) TColor {
	r1 := customColorBoxImportAPI().SysCallN(5, m.Instance(), uintptr(Index))
	return TColor(r1)
}

func (m *TCustomColorBox) ColorNames(Index int32) string {
	r1 := customColorBoxImportAPI().SysCallN(2, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TCustomColorBox) Selected() TColor {
	r1 := customColorBoxImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorBox) SetSelected(AValue TColor) {
	customColorBoxImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) DefaultColorColor() TColor {
	r1 := customColorBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorBox) SetDefaultColorColor(AValue TColor) {
	customColorBoxImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) NoneColorColor() TColor {
	r1 := customColorBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorBox) SetNoneColorColor(AValue TColor) {
	customColorBoxImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) ColorDialog() IColorDialog {
	r1 := customColorBoxImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return AsColorDialog(r1)
}

func (m *TCustomColorBox) SetColorDialog(AValue IColorDialog) {
	customColorBoxImportAPI().SysCallN(1, 1, m.Instance(), GetObjectUintptr(AValue))
}

func CustomColorBoxClass() TClass {
	ret := customColorBoxImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomColorBox) SetOnGetColors(fn TGetColorsEvent) {
	if m.getColorsPtr != 0 {
		RemoveEventElement(m.getColorsPtr)
	}
	m.getColorsPtr = MakeEventDataPtr(fn)
	customColorBoxImportAPI().SysCallN(10, m.Instance(), m.getColorsPtr)
}

var (
	customColorBoxImport       *imports.Imports = nil
	customColorBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomColorBox_Class", 0),
		/*1*/ imports.NewTable("CustomColorBox_ColorDialog", 0),
		/*2*/ imports.NewTable("CustomColorBox_ColorNames", 0),
		/*3*/ imports.NewTable("CustomColorBox_ColorRectOffset", 0),
		/*4*/ imports.NewTable("CustomColorBox_ColorRectWidth", 0),
		/*5*/ imports.NewTable("CustomColorBox_Colors", 0),
		/*6*/ imports.NewTable("CustomColorBox_Create", 0),
		/*7*/ imports.NewTable("CustomColorBox_DefaultColorColor", 0),
		/*8*/ imports.NewTable("CustomColorBox_NoneColorColor", 0),
		/*9*/ imports.NewTable("CustomColorBox_Selected", 0),
		/*10*/ imports.NewTable("CustomColorBox_SetOnGetColors", 0),
		/*11*/ imports.NewTable("CustomColorBox_StyleForColorBoxStyle", 0),
	}
)

func customColorBoxImportAPI() *imports.Imports {
	if customColorBoxImport == nil {
		customColorBoxImport = NewDefaultImports()
		customColorBoxImport.SetImportTable(customColorBoxImportTables)
		customColorBoxImportTables = nil
	}
	return customColorBoxImport
}
