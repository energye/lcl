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

// IFPCustomFont Parent: IFPCanvasHelper
type IFPCustomFont interface {
	IFPCanvasHelper
	Name() string                         // property
	SetName(AValue string)                // property
	Size() int32                          // property
	SetSize(AValue int32)                 // property
	Bold() bool                           // property
	SetBold(AValue bool)                  // property
	Italic() bool                         // property
	SetItalic(AValue bool)                // property
	Underline() bool                      // property
	SetUnderline(AValue bool)             // property
	StrikeThrough() bool                  // property
	SetStrikeThrough(AValue bool)         // property
	Orientation() int32                   // property
	SetOrientation(AValue int32)          // property
	CopyFont() IFPCustomFont              // function
	GetTextHeight(text string) int32      // function
	GetTextWidth(text string) int32       // function
	GetTextSize(text string, w, h *int32) // procedure
}

// TFPCustomFont Parent: TFPCanvasHelper
type TFPCustomFont struct {
	TFPCanvasHelper
}

func NewFPCustomFont() IFPCustomFont {
	r1 := fPCustomFontImportAPI().SysCallN(3)
	return AsFPCustomFont(r1)
}

func (m *TFPCustomFont) Name() string {
	r1 := fPCustomFontImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFPCustomFont) SetName(AValue string) {
	fPCustomFontImportAPI().SysCallN(8, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFPCustomFont) Size() int32 {
	r1 := fPCustomFontImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomFont) SetSize(AValue int32) {
	fPCustomFontImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomFont) Bold() bool {
	r1 := fPCustomFontImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetBold(AValue bool) {
	fPCustomFontImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) Italic() bool {
	r1 := fPCustomFontImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetItalic(AValue bool) {
	fPCustomFontImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) Underline() bool {
	r1 := fPCustomFontImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetUnderline(AValue bool) {
	fPCustomFontImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) StrikeThrough() bool {
	r1 := fPCustomFontImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetStrikeThrough(AValue bool) {
	fPCustomFontImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) Orientation() int32 {
	r1 := fPCustomFontImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomFont) SetOrientation(AValue int32) {
	fPCustomFontImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomFont) CopyFont() IFPCustomFont {
	r1 := fPCustomFontImportAPI().SysCallN(2, m.Instance())
	return AsFPCustomFont(r1)
}

func (m *TFPCustomFont) GetTextHeight(text string) int32 {
	r1 := fPCustomFontImportAPI().SysCallN(4, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomFont) GetTextWidth(text string) int32 {
	r1 := fPCustomFontImportAPI().SysCallN(6, m.Instance(), PascalStr(text))
	return int32(r1)
}

func FPCustomFontClass() TClass {
	ret := fPCustomFontImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TFPCustomFont) GetTextSize(text string, w, h *int32) {
	var result1 uintptr
	var result2 uintptr
	fPCustomFontImportAPI().SysCallN(5, m.Instance(), PascalStr(text), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*w = int32(result1)
	*h = int32(result2)
}

var (
	fPCustomFontImport       *imports.Imports = nil
	fPCustomFontImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCustomFont_Bold", 0),
		/*1*/ imports.NewTable("FPCustomFont_Class", 0),
		/*2*/ imports.NewTable("FPCustomFont_CopyFont", 0),
		/*3*/ imports.NewTable("FPCustomFont_Create", 0),
		/*4*/ imports.NewTable("FPCustomFont_GetTextHeight", 0),
		/*5*/ imports.NewTable("FPCustomFont_GetTextSize", 0),
		/*6*/ imports.NewTable("FPCustomFont_GetTextWidth", 0),
		/*7*/ imports.NewTable("FPCustomFont_Italic", 0),
		/*8*/ imports.NewTable("FPCustomFont_Name", 0),
		/*9*/ imports.NewTable("FPCustomFont_Orientation", 0),
		/*10*/ imports.NewTable("FPCustomFont_Size", 0),
		/*11*/ imports.NewTable("FPCustomFont_StrikeThrough", 0),
		/*12*/ imports.NewTable("FPCustomFont_Underline", 0),
	}
)

func fPCustomFontImportAPI() *imports.Imports {
	if fPCustomFontImport == nil {
		fPCustomFontImport = NewDefaultImports()
		fPCustomFontImport.SetImportTable(fPCustomFontImportTables)
		fPCustomFontImportTables = nil
	}
	return fPCustomFontImport
}
