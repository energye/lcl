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
)

// IFPCustomFont Parent: IFPCanvasHelper
type IFPCustomFont interface {
	IFPCanvasHelper
	CopyFont() IFPCustomFont         // function
	GetTextHeight(text string) int32 // function
	GetTextWidth(text string) int32  // function
	// GetTextSize
	//  Creates a copy of the font with all properties the same, but not allocated
	GetTextSize(text string, W *int32, H *int32) // procedure
	Name() string                                // property Name Getter
	SetName(value string)                        // property Name Setter
	Size() int32                                 // property Size Getter
	SetSize(value int32)                         // property Size Setter
	Bold() bool                                  // property Bold Getter
	SetBold(value bool)                          // property Bold Setter
	Italic() bool                                // property Italic Getter
	SetItalic(value bool)                        // property Italic Setter
	Underline() bool                             // property Underline Getter
	SetUnderline(value bool)                     // property Underline Setter
	StrikeThrough() bool                         // property StrikeThrough Getter
	SetStrikeThrough(value bool)                 // property StrikeThrough Setter
	Orientation() int32                          // property Orientation Getter
	SetOrientation(value int32)                  // property Orientation Setter
}

type TFPCustomFont struct {
	TFPCanvasHelper
}

func (m *TFPCustomFont) CopyFont() IFPCustomFont {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomFontAPI().SysCallN(1, m.Instance())
	return AsFPCustomFont(r)
}

func (m *TFPCustomFont) GetTextHeight(text string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomFontAPI().SysCallN(2, m.Instance(), api.PasStr(text))
	return int32(r)
}

func (m *TFPCustomFont) GetTextWidth(text string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomFontAPI().SysCallN(3, m.Instance(), api.PasStr(text))
	return int32(r)
}

func (m *TFPCustomFont) GetTextSize(text string, W *int32, H *int32) {
	if !m.IsValid() {
		return
	}
	WPtr := uintptr(*W)
	HPtr := uintptr(*H)
	fPCustomFontAPI().SysCallN(4, m.Instance(), api.PasStr(text), uintptr(base.UnsafePointer(&WPtr)), uintptr(base.UnsafePointer(&HPtr)))
	*W = int32(WPtr)
	*H = int32(HPtr)
}

func (m *TFPCustomFont) Name() string {
	if !m.IsValid() {
		return ""
	}
	r := fPCustomFontAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFPCustomFont) SetName(value string) {
	if !m.IsValid() {
		return
	}
	fPCustomFontAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TFPCustomFont) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomFontAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TFPCustomFont) SetSize(value int32) {
	if !m.IsValid() {
		return
	}
	fPCustomFontAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomFont) Bold() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomFontAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCustomFont) SetBold(value bool) {
	if !m.IsValid() {
		return
	}
	fPCustomFontAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TFPCustomFont) Italic() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomFontAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCustomFont) SetItalic(value bool) {
	if !m.IsValid() {
		return
	}
	fPCustomFontAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TFPCustomFont) Underline() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomFontAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCustomFont) SetUnderline(value bool) {
	if !m.IsValid() {
		return
	}
	fPCustomFontAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TFPCustomFont) StrikeThrough() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomFontAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCustomFont) SetStrikeThrough(value bool) {
	if !m.IsValid() {
		return
	}
	fPCustomFontAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TFPCustomFont) Orientation() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomFontAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TFPCustomFont) SetOrientation(value int32) {
	if !m.IsValid() {
		return
	}
	fPCustomFontAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

// NewFPCustomFont class constructor
func NewFPCustomFont() IFPCustomFont {
	r := fPCustomFontAPI().SysCallN(0)
	return AsFPCustomFont(r)
}

var (
	fPCustomFontOnce   base.Once
	fPCustomFontImport *imports.Imports = nil
)

func fPCustomFontAPI() *imports.Imports {
	fPCustomFontOnce.Do(func() {
		fPCustomFontImport = api.NewDefaultImports()
		fPCustomFontImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomFont_Create", 0), // constructor NewFPCustomFont
			/* 1 */ imports.NewTable("TFPCustomFont_CopyFont", 0), // function CopyFont
			/* 2 */ imports.NewTable("TFPCustomFont_GetTextHeight", 0), // function GetTextHeight
			/* 3 */ imports.NewTable("TFPCustomFont_GetTextWidth", 0), // function GetTextWidth
			/* 4 */ imports.NewTable("TFPCustomFont_GetTextSize", 0), // procedure GetTextSize
			/* 5 */ imports.NewTable("TFPCustomFont_Name", 0), // property Name
			/* 6 */ imports.NewTable("TFPCustomFont_Size", 0), // property Size
			/* 7 */ imports.NewTable("TFPCustomFont_Bold", 0), // property Bold
			/* 8 */ imports.NewTable("TFPCustomFont_Italic", 0), // property Italic
			/* 9 */ imports.NewTable("TFPCustomFont_Underline", 0), // property Underline
			/* 10 */ imports.NewTable("TFPCustomFont_StrikeThrough", 0), // property StrikeThrough
			/* 11 */ imports.NewTable("TFPCustomFont_Orientation", 0), // property Orientation
		}
	})
	return fPCustomFontImport
}
