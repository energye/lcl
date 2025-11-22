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

// IFont Parent: IFPCustomFont
type IFont interface {
	IFPCustomFont
	HandleAllocated() bool               // function
	IsDefault() bool                     // function
	IsEqual(font IFont) bool             // function
	BeginUpdate()                        // procedure
	EndUpdate()                          // procedure
	SetDefault()                         // procedure
	FontData() TFontData                 // property FontData Getter
	SetFontData(value TFontData)         // property FontData Setter
	Handle() types.HFONT                 // property Handle Getter
	SetHandle(value types.HFONT)         // property Handle Setter
	IsMonoSpace() bool                   // property IsMonoSpace Getter
	PixelsPerInch() int32                // property PixelsPerInch Getter
	SetPixelsPerInch(value int32)        // property PixelsPerInch Setter
	CharSet() types.TFontCharSet         // property CharSet Getter
	SetCharSet(value types.TFontCharSet) // property CharSet Setter
	Color() types.TColor                 // property Color Getter
	SetColor(value types.TColor)         // property Color Setter
	Height() int32                       // property Height Getter
	SetHeight(value int32)               // property Height Setter
	Pitch() types.TFontPitch             // property Pitch Getter
	SetPitch(value types.TFontPitch)     // property Pitch Setter
	Quality() types.TFontQuality         // property Quality Getter
	SetQuality(value types.TFontQuality) // property Quality Setter
	Style() types.TFontStyles            // property Style Getter
	SetStyle(value types.TFontStyles)    // property Style Setter
}

type TFont struct {
	TFPCustomFont
}

func (m *TFont) HandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := fontAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TFont) IsDefault() bool {
	if !m.IsValid() {
		return false
	}
	r := fontAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TFont) IsEqual(font IFont) bool {
	if !m.IsValid() {
		return false
	}
	r := fontAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(font))
	return api.GoBool(r)
}

func (m *TFont) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(4, m.Instance())
}

func (m *TFont) EndUpdate() {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(5, m.Instance())
}

func (m *TFont) SetDefault() {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(6, m.Instance())
}

func (m *TFont) FontData() (result TFontData) {
	if !m.IsValid() {
		return
	}
	resultPtr := result.ToPas()
	fontAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(resultPtr)))
	result = resultPtr.ToGo()
	return
}

func (m *TFont) SetFontData(value TFontData) {
	if !m.IsValid() {
		return
	}
	valuePtr := value.ToPas()
	fontAPI().SysCallN(7, 1, m.Instance(), uintptr(base.UnsafePointer(valuePtr)))
}

func (m *TFont) Handle() types.HFONT {
	if !m.IsValid() {
		return 0
	}
	r := fontAPI().SysCallN(8, 0, m.Instance())
	return types.HFONT(r)
}

func (m *TFont) SetHandle(value types.HFONT) {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TFont) IsMonoSpace() bool {
	if !m.IsValid() {
		return false
	}
	r := fontAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TFont) PixelsPerInch() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fontAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TFont) SetPixelsPerInch(value int32) {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TFont) CharSet() types.TFontCharSet {
	if !m.IsValid() {
		return 0
	}
	r := fontAPI().SysCallN(11, 0, m.Instance())
	return types.TFontCharSet(r)
}

func (m *TFont) SetCharSet(value types.TFontCharSet) {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TFont) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := fontAPI().SysCallN(12, 0, m.Instance())
	return types.TColor(r)
}

func (m *TFont) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TFont) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fontAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TFont) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TFont) Pitch() types.TFontPitch {
	if !m.IsValid() {
		return 0
	}
	r := fontAPI().SysCallN(14, 0, m.Instance())
	return types.TFontPitch(r)
}

func (m *TFont) SetPitch(value types.TFontPitch) {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TFont) Quality() types.TFontQuality {
	if !m.IsValid() {
		return 0
	}
	r := fontAPI().SysCallN(15, 0, m.Instance())
	return types.TFontQuality(r)
}

func (m *TFont) SetQuality(value types.TFontQuality) {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TFont) Style() types.TFontStyles {
	if !m.IsValid() {
		return 0
	}
	r := fontAPI().SysCallN(16, 0, m.Instance())
	return types.TFontStyles(r)
}

func (m *TFont) SetStyle(value types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	fontAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

// NewFont class constructor
func NewFont() IFont {
	r := fontAPI().SysCallN(0)
	return AsFont(r)
}

var (
	fontOnce   base.Once
	fontImport *imports.Imports = nil
)

func fontAPI() *imports.Imports {
	fontOnce.Do(func() {
		fontImport = api.NewDefaultImports()
		fontImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFont_Create", 0), // constructor NewFont
			/* 1 */ imports.NewTable("TFont_HandleAllocated", 0), // function HandleAllocated
			/* 2 */ imports.NewTable("TFont_IsDefault", 0), // function IsDefault
			/* 3 */ imports.NewTable("TFont_IsEqual", 0), // function IsEqual
			/* 4 */ imports.NewTable("TFont_BeginUpdate", 0), // procedure BeginUpdate
			/* 5 */ imports.NewTable("TFont_EndUpdate", 0), // procedure EndUpdate
			/* 6 */ imports.NewTable("TFont_SetDefault", 0), // procedure SetDefault
			/* 7 */ imports.NewTable("TFont_FontData", 0), // property FontData
			/* 8 */ imports.NewTable("TFont_Handle", 0), // property Handle
			/* 9 */ imports.NewTable("TFont_IsMonoSpace", 0), // property IsMonoSpace
			/* 10 */ imports.NewTable("TFont_PixelsPerInch", 0), // property PixelsPerInch
			/* 11 */ imports.NewTable("TFont_CharSet", 0), // property CharSet
			/* 12 */ imports.NewTable("TFont_Color", 0), // property Color
			/* 13 */ imports.NewTable("TFont_Height", 0), // property Height
			/* 14 */ imports.NewTable("TFont_Pitch", 0), // property Pitch
			/* 15 */ imports.NewTable("TFont_Quality", 0), // property Quality
			/* 16 */ imports.NewTable("TFont_Style", 0), // property Style
		}
	})
	return fontImport
}
