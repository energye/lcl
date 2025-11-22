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

// ITextAttributes Parent: IPersistent
type ITextAttributes interface {
	IPersistent
	Name() string                        // property Name Getter
	SetName(value string)                // property Name Setter
	Pitch() types.TFontPitch             // property Pitch Getter
	SetPitch(value types.TFontPitch)     // property Pitch Setter
	Charset() types.TFontCharSet         // property Charset Getter
	SetCharset(value types.TFontCharSet) // property Charset Setter
	Color() types.TColor                 // property Color Getter
	SetColor(value types.TColor)         // property Color Setter
	Size() int32                         // property Size Getter
	SetSize(value int32)                 // property Size Setter
	Style() types.TFontStyles            // property Style Getter
	SetStyle(value types.TFontStyles)    // property Style Setter
	Height() int32                       // property Height Getter
	SetHeight(value int32)               // property Height Setter
}

type TTextAttributes struct {
	TPersistent
}

func (m *TTextAttributes) Name() string {
	if !m.IsValid() {
		return ""
	}
	r := textAttributesAPI().SysCallN(1, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTextAttributes) SetName(value string) {
	if !m.IsValid() {
		return
	}
	textAttributesAPI().SysCallN(1, 1, m.Instance(), api.PasStr(value))
}

func (m *TTextAttributes) Pitch() types.TFontPitch {
	if !m.IsValid() {
		return 0
	}
	r := textAttributesAPI().SysCallN(2, 0, m.Instance())
	return types.TFontPitch(r)
}

func (m *TTextAttributes) SetPitch(value types.TFontPitch) {
	if !m.IsValid() {
		return
	}
	textAttributesAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TTextAttributes) Charset() types.TFontCharSet {
	if !m.IsValid() {
		return 0
	}
	r := textAttributesAPI().SysCallN(3, 0, m.Instance())
	return types.TFontCharSet(r)
}

func (m *TTextAttributes) SetCharset(value types.TFontCharSet) {
	if !m.IsValid() {
		return
	}
	textAttributesAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TTextAttributes) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := textAttributesAPI().SysCallN(4, 0, m.Instance())
	return types.TColor(r)
}

func (m *TTextAttributes) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	textAttributesAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TTextAttributes) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := textAttributesAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TTextAttributes) SetSize(value int32) {
	if !m.IsValid() {
		return
	}
	textAttributesAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TTextAttributes) Style() types.TFontStyles {
	if !m.IsValid() {
		return 0
	}
	r := textAttributesAPI().SysCallN(6, 0, m.Instance())
	return types.TFontStyles(r)
}

func (m *TTextAttributes) SetStyle(value types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	textAttributesAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TTextAttributes) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := textAttributesAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TTextAttributes) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	textAttributesAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

// NewTextAttributes class constructor
func NewTextAttributes(owner IRichMemo, type_ types.TAttributeType) ITextAttributes {
	r := textAttributesAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(type_))
	return AsTextAttributes(r)
}

var (
	textAttributesOnce   base.Once
	textAttributesImport *imports.Imports = nil
)

func textAttributesAPI() *imports.Imports {
	textAttributesOnce.Do(func() {
		textAttributesImport = api.NewDefaultImports()
		textAttributesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTextAttributes_Create", 0), // constructor NewTextAttributes
			/* 1 */ imports.NewTable("TTextAttributes_Name", 0), // property Name
			/* 2 */ imports.NewTable("TTextAttributes_Pitch", 0), // property Pitch
			/* 3 */ imports.NewTable("TTextAttributes_Charset", 0), // property Charset
			/* 4 */ imports.NewTable("TTextAttributes_Color", 0), // property Color
			/* 5 */ imports.NewTable("TTextAttributes_Size", 0), // property Size
			/* 6 */ imports.NewTable("TTextAttributes_Style", 0), // property Style
			/* 7 */ imports.NewTable("TTextAttributes_Height", 0), // property Height
		}
	})
	return textAttributesImport
}
