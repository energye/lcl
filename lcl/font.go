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

// IFont Parent: IFPCustomFont
type IFont interface {
	IFPCustomFont
	FontData() (resultFontData TFontData) // property
	SetFontData(AValue *TFontData)        // property
	Handle() HFONT                        // property
	SetHandle(AValue HFONT)               // property
	IsMonoSpace() bool                    // property
	PixelsPerInch() int32                 // property
	SetPixelsPerInch(AValue int32)        // property
	CharSet() TFontCharSet                // property
	SetCharSet(AValue TFontCharSet)       // property
	Color() TColor                        // property
	SetColor(AValue TColor)               // property
	Height() int32                        // property
	SetHeight(AValue int32)               // property
	Pitch() TFontPitch                    // property
	SetPitch(AValue TFontPitch)           // property
	Quality() TFontQuality                // property
	SetQuality(AValue TFontQuality)       // property
	Style() TFontStyles                   // property
	SetStyle(AValue TFontStyles)          // property
	HandleAllocated() bool                // function
	IsDefault() bool                      // function
	IsEqual(AFont IFont) bool             // function
	BeginUpdate()                         // procedure
	EndUpdate()                           // procedure
	SetDefault()                          // procedure
}

// TFont Parent: TFPCustomFont
type TFont struct {
	TFPCustomFont
}

func NewFont() IFont {
	r1 := fontImportAPI().SysCallN(4)
	return AsFont(r1)
}

func (m *TFont) FontData() (resultFontData TFontData) {
	r1 := fontImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return *(*TFontData)(getPointer(r1))
}

func (m *TFont) SetFontData(AValue *TFontData) {
	fontImportAPI().SysCallN(6, 1, m.Instance(), uintptr(unsafePointer(AValue)))
}

func (m *TFont) Handle() HFONT {
	r1 := fontImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return HFONT(r1)
}

func (m *TFont) SetHandle(AValue HFONT) {
	fontImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) IsMonoSpace() bool {
	r1 := fontImportAPI().SysCallN(12, m.Instance())
	return GoBool(r1)
}

func (m *TFont) PixelsPerInch() int32 {
	r1 := fontImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFont) SetPixelsPerInch(AValue int32) {
	fontImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) CharSet() TFontCharSet {
	r1 := fontImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TFontCharSet(r1)
}

func (m *TFont) SetCharSet(AValue TFontCharSet) {
	fontImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Color() TColor {
	r1 := fontImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TFont) SetColor(AValue TColor) {
	fontImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Height() int32 {
	r1 := fontImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFont) SetHeight(AValue int32) {
	fontImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Pitch() TFontPitch {
	r1 := fontImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TFontPitch(r1)
}

func (m *TFont) SetPitch(AValue TFontPitch) {
	fontImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Quality() TFontQuality {
	r1 := fontImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TFontQuality(r1)
}

func (m *TFont) SetQuality(AValue TFontQuality) {
	fontImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Style() TFontStyles {
	r1 := fontImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return TFontStyles(r1)
}

func (m *TFont) SetStyle(AValue TFontStyles) {
	fontImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) HandleAllocated() bool {
	r1 := fontImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TFont) IsDefault() bool {
	r1 := fontImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TFont) IsEqual(AFont IFont) bool {
	r1 := fontImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(AFont))
	return GoBool(r1)
}

func FontClass() TClass {
	ret := fontImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TFont) BeginUpdate() {
	fontImportAPI().SysCallN(0, m.Instance())
}

func (m *TFont) EndUpdate() {
	fontImportAPI().SysCallN(5, m.Instance())
}

func (m *TFont) SetDefault() {
	fontImportAPI().SysCallN(16, m.Instance())
}

var (
	fontImport       *imports.Imports = nil
	fontImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Font_BeginUpdate", 0),
		/*1*/ imports.NewTable("Font_CharSet", 0),
		/*2*/ imports.NewTable("Font_Class", 0),
		/*3*/ imports.NewTable("Font_Color", 0),
		/*4*/ imports.NewTable("Font_Create", 0),
		/*5*/ imports.NewTable("Font_EndUpdate", 0),
		/*6*/ imports.NewTable("Font_FontData", 0),
		/*7*/ imports.NewTable("Font_Handle", 0),
		/*8*/ imports.NewTable("Font_HandleAllocated", 0),
		/*9*/ imports.NewTable("Font_Height", 0),
		/*10*/ imports.NewTable("Font_IsDefault", 0),
		/*11*/ imports.NewTable("Font_IsEqual", 0),
		/*12*/ imports.NewTable("Font_IsMonoSpace", 0),
		/*13*/ imports.NewTable("Font_Pitch", 0),
		/*14*/ imports.NewTable("Font_PixelsPerInch", 0),
		/*15*/ imports.NewTable("Font_Quality", 0),
		/*16*/ imports.NewTable("Font_SetDefault", 0),
		/*17*/ imports.NewTable("Font_Style", 0),
	}
)

func fontImportAPI() *imports.Imports {
	if fontImport == nil {
		fontImport = NewDefaultImports()
		fontImport.SetImportTable(fontImportTables)
		fontImportTables = nil
	}
	return fontImport
}
