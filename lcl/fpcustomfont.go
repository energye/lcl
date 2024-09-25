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
	r1 := LCL().SysCallN(2928)
	return AsFPCustomFont(r1)
}

func (m *TFPCustomFont) Name() string {
	r1 := LCL().SysCallN(2933, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFPCustomFont) SetName(AValue string) {
	LCL().SysCallN(2933, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFPCustomFont) Size() int32 {
	r1 := LCL().SysCallN(2935, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomFont) SetSize(AValue int32) {
	LCL().SysCallN(2935, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomFont) Bold() bool {
	r1 := LCL().SysCallN(2925, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetBold(AValue bool) {
	LCL().SysCallN(2925, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) Italic() bool {
	r1 := LCL().SysCallN(2932, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetItalic(AValue bool) {
	LCL().SysCallN(2932, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) Underline() bool {
	r1 := LCL().SysCallN(2937, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetUnderline(AValue bool) {
	LCL().SysCallN(2937, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) StrikeThrough() bool {
	r1 := LCL().SysCallN(2936, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomFont) SetStrikeThrough(AValue bool) {
	LCL().SysCallN(2936, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomFont) Orientation() int32 {
	r1 := LCL().SysCallN(2934, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomFont) SetOrientation(AValue int32) {
	LCL().SysCallN(2934, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomFont) CopyFont() IFPCustomFont {
	r1 := LCL().SysCallN(2927, m.Instance())
	return AsFPCustomFont(r1)
}

func (m *TFPCustomFont) GetTextHeight(text string) int32 {
	r1 := LCL().SysCallN(2929, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomFont) GetTextWidth(text string) int32 {
	r1 := LCL().SysCallN(2931, m.Instance(), PascalStr(text))
	return int32(r1)
}

func FPCustomFontClass() TClass {
	ret := LCL().SysCallN(2926)
	return TClass(ret)
}

func (m *TFPCustomFont) GetTextSize(text string, w, h *int32) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(2930, m.Instance(), PascalStr(text), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*w = int32(result1)
	*h = int32(result2)
}
