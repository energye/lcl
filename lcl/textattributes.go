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

// ITextAttributes Parent: IPersistent
type ITextAttributes interface {
	IPersistent
	Name() string                   // property
	SetName(AValue string)          // property
	Pitch() TFontPitch              // property
	SetPitch(AValue TFontPitch)     // property
	Charset() TFontCharSet          // property
	SetCharset(AValue TFontCharSet) // property
	Color() TColor                  // property
	SetColor(AValue TColor)         // property
	Size() int32                    // property
	SetSize(AValue int32)           // property
	Style() TFontStyles             // property
	SetStyle(AValue TFontStyles)    // property
	Height() int32                  // property
	SetHeight(AValue int32)         // property
}

// TTextAttributes Parent: TPersistent
type TTextAttributes struct {
	TPersistent
}

func NewTextAttributes(AOwner IRichMemo, AType TAttributeType) ITextAttributes {
	r1 := LCL().SysCallN(5413, GetObjectUintptr(AOwner), uintptr(AType))
	return AsTextAttributes(r1)
}

func (m *TTextAttributes) Name() string {
	r1 := LCL().SysCallN(5415, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TTextAttributes) SetName(AValue string) {
	LCL().SysCallN(5415, 1, m.Instance(), PascalStr(AValue))
}

func (m *TTextAttributes) Pitch() TFontPitch {
	r1 := LCL().SysCallN(5416, 0, m.Instance(), 0)
	return TFontPitch(r1)
}

func (m *TTextAttributes) SetPitch(AValue TFontPitch) {
	LCL().SysCallN(5416, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Charset() TFontCharSet {
	r1 := LCL().SysCallN(5410, 0, m.Instance(), 0)
	return TFontCharSet(r1)
}

func (m *TTextAttributes) SetCharset(AValue TFontCharSet) {
	LCL().SysCallN(5410, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Color() TColor {
	r1 := LCL().SysCallN(5412, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TTextAttributes) SetColor(AValue TColor) {
	LCL().SysCallN(5412, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Size() int32 {
	r1 := LCL().SysCallN(5417, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTextAttributes) SetSize(AValue int32) {
	LCL().SysCallN(5417, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Style() TFontStyles {
	r1 := LCL().SysCallN(5418, 0, m.Instance(), 0)
	return TFontStyles(r1)
}

func (m *TTextAttributes) SetStyle(AValue TFontStyles) {
	LCL().SysCallN(5418, 1, m.Instance(), uintptr(AValue))
}

func (m *TTextAttributes) Height() int32 {
	r1 := LCL().SysCallN(5414, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTextAttributes) SetHeight(AValue int32) {
	LCL().SysCallN(5414, 1, m.Instance(), uintptr(AValue))
}

func TextAttributesClass() TClass {
	ret := LCL().SysCallN(5411)
	return TClass(ret)
}
