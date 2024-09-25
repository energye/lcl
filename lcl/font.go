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
	r1 := LCL().SysCallN(3113)
	return AsFont(r1)
}

func (m *TFont) FontData() (resultFontData TFontData) {
	r1 := LCL().SysCallN(3115, 0, m.Instance(), 0)
	return *(*TFontData)(getPointer(r1))
}

func (m *TFont) SetFontData(AValue *TFontData) {
	LCL().SysCallN(3115, 1, m.Instance(), uintptr(unsafePointer(AValue)))
}

func (m *TFont) Handle() HFONT {
	r1 := LCL().SysCallN(3116, 0, m.Instance(), 0)
	return HFONT(r1)
}

func (m *TFont) SetHandle(AValue HFONT) {
	LCL().SysCallN(3116, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) IsMonoSpace() bool {
	r1 := LCL().SysCallN(3121, m.Instance())
	return GoBool(r1)
}

func (m *TFont) PixelsPerInch() int32 {
	r1 := LCL().SysCallN(3123, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFont) SetPixelsPerInch(AValue int32) {
	LCL().SysCallN(3123, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) CharSet() TFontCharSet {
	r1 := LCL().SysCallN(3110, 0, m.Instance(), 0)
	return TFontCharSet(r1)
}

func (m *TFont) SetCharSet(AValue TFontCharSet) {
	LCL().SysCallN(3110, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Color() TColor {
	r1 := LCL().SysCallN(3112, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TFont) SetColor(AValue TColor) {
	LCL().SysCallN(3112, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Height() int32 {
	r1 := LCL().SysCallN(3118, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFont) SetHeight(AValue int32) {
	LCL().SysCallN(3118, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Pitch() TFontPitch {
	r1 := LCL().SysCallN(3122, 0, m.Instance(), 0)
	return TFontPitch(r1)
}

func (m *TFont) SetPitch(AValue TFontPitch) {
	LCL().SysCallN(3122, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Quality() TFontQuality {
	r1 := LCL().SysCallN(3124, 0, m.Instance(), 0)
	return TFontQuality(r1)
}

func (m *TFont) SetQuality(AValue TFontQuality) {
	LCL().SysCallN(3124, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) Style() TFontStyles {
	r1 := LCL().SysCallN(3126, 0, m.Instance(), 0)
	return TFontStyles(r1)
}

func (m *TFont) SetStyle(AValue TFontStyles) {
	LCL().SysCallN(3126, 1, m.Instance(), uintptr(AValue))
}

func (m *TFont) HandleAllocated() bool {
	r1 := LCL().SysCallN(3117, m.Instance())
	return GoBool(r1)
}

func (m *TFont) IsDefault() bool {
	r1 := LCL().SysCallN(3119, m.Instance())
	return GoBool(r1)
}

func (m *TFont) IsEqual(AFont IFont) bool {
	r1 := LCL().SysCallN(3120, m.Instance(), GetObjectUintptr(AFont))
	return GoBool(r1)
}

func FontClass() TClass {
	ret := LCL().SysCallN(3111)
	return TClass(ret)
}

func (m *TFont) BeginUpdate() {
	LCL().SysCallN(3109, m.Instance())
}

func (m *TFont) EndUpdate() {
	LCL().SysCallN(3114, m.Instance())
}

func (m *TFont) SetDefault() {
	LCL().SysCallN(3125, m.Instance())
}
