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

// ILazSynCustomTextAttributes Parent: IPersistent
type ILazSynCustomTextAttributes interface {
	IPersistent
	Clear()                                               // procedure
	SetAllPriorities(priority int32)                      // procedure
	BeginUpdate()                                         // procedure
	EndUpdate()                                           // procedure
	Background() types.TColor                             // property Background Getter
	SetBackground(value types.TColor)                     // property Background Setter
	Foreground() types.TColor                             // property Foreground Getter
	SetForeground(value types.TColor)                     // property Foreground Setter
	FrameColor() types.TColor                             // property FrameColor Getter
	SetFrameColor(value types.TColor)                     // property FrameColor Setter
	FrameStyle() types.TSynLineStyle                      // property FrameStyle Getter
	SetFrameStyle(value types.TSynLineStyle)              // property FrameStyle Setter
	FrameEdges() types.TSynFrameEdges                     // property FrameEdges Getter
	SetFrameEdges(value types.TSynFrameEdges)             // property FrameEdges Setter
	Style() types.TFontStyles                             // property Style Getter
	SetStyle(value types.TFontStyles)                     // property Style Setter
	StyleMask() types.TFontStyles                         // property StyleMask Getter
	SetStyleMask(value types.TFontStyles)                 // property StyleMask Setter
	StylePriority(index types.TFontStyle) int32           // property StylePriority Getter
	SetStylePriority(index types.TFontStyle, value int32) // property StylePriority Setter
	BackPriority() int32                                  // property BackPriority Getter
	SetBackPriority(value int32)                          // property BackPriority Setter
	ForePriority() int32                                  // property ForePriority Getter
	SetForePriority(value int32)                          // property ForePriority Setter
	FramePriority() int32                                 // property FramePriority Getter
	SetFramePriority(value int32)                         // property FramePriority Setter
	BoldPriority() int32                                  // property BoldPriority Getter
	SetBoldPriority(value int32)                          // property BoldPriority Setter
	ItalicPriority() int32                                // property ItalicPriority Getter
	SetItalicPriority(value int32)                        // property ItalicPriority Setter
	UnderlinePriority() int32                             // property UnderlinePriority Getter
	SetUnderlinePriority(value int32)                     // property UnderlinePriority Setter
	StrikeOutPriority() int32                             // property StrikeOutPriority Getter
	SetStrikeOutPriority(value int32)                     // property StrikeOutPriority Setter
}

type TLazSynCustomTextAttributes struct {
	TPersistent
}

func (m *TLazSynCustomTextAttributes) Clear() {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(1, m.Instance())
}

func (m *TLazSynCustomTextAttributes) SetAllPriorities(priority int32) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(2, m.Instance(), uintptr(priority))
}

func (m *TLazSynCustomTextAttributes) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(3, m.Instance())
}

func (m *TLazSynCustomTextAttributes) EndUpdate() {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(4, m.Instance())
}

func (m *TLazSynCustomTextAttributes) Background() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(5, 0, m.Instance())
	return types.TColor(r)
}

func (m *TLazSynCustomTextAttributes) SetBackground(value types.TColor) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) Foreground() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(6, 0, m.Instance())
	return types.TColor(r)
}

func (m *TLazSynCustomTextAttributes) SetForeground(value types.TColor) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) FrameColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(7, 0, m.Instance())
	return types.TColor(r)
}

func (m *TLazSynCustomTextAttributes) SetFrameColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) FrameStyle() types.TSynLineStyle {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(8, 0, m.Instance())
	return types.TSynLineStyle(r)
}

func (m *TLazSynCustomTextAttributes) SetFrameStyle(value types.TSynLineStyle) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) FrameEdges() types.TSynFrameEdges {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(9, 0, m.Instance())
	return types.TSynFrameEdges(r)
}

func (m *TLazSynCustomTextAttributes) SetFrameEdges(value types.TSynFrameEdges) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) Style() types.TFontStyles {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(10, 0, m.Instance())
	return types.TFontStyles(r)
}

func (m *TLazSynCustomTextAttributes) SetStyle(value types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) StyleMask() types.TFontStyles {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(11, 0, m.Instance())
	return types.TFontStyles(r)
}

func (m *TLazSynCustomTextAttributes) SetStyleMask(value types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) StylePriority(index types.TFontStyle) int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(12, 0, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TLazSynCustomTextAttributes) SetStylePriority(index types.TFontStyle, value int32) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(12, 1, m.Instance(), uintptr(index), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) BackPriority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynCustomTextAttributes) SetBackPriority(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) ForePriority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynCustomTextAttributes) SetForePriority(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) FramePriority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynCustomTextAttributes) SetFramePriority(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) BoldPriority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynCustomTextAttributes) SetBoldPriority(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) ItalicPriority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynCustomTextAttributes) SetItalicPriority(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) UnderlinePriority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynCustomTextAttributes) SetUnderlinePriority(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynCustomTextAttributes) StrikeOutPriority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynCustomTextAttributesAPI().SysCallN(19, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynCustomTextAttributes) SetStrikeOutPriority(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynCustomTextAttributesAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

// NewLazSynCustomTextAttributes class constructor
func NewLazSynCustomTextAttributes() ILazSynCustomTextAttributes {
	r := lazSynCustomTextAttributesAPI().SysCallN(0)
	return AsLazSynCustomTextAttributes(r)
}

var (
	lazSynCustomTextAttributesOnce   base.Once
	lazSynCustomTextAttributesImport *imports.Imports = nil
)

func lazSynCustomTextAttributesAPI() *imports.Imports {
	lazSynCustomTextAttributesOnce.Do(func() {
		lazSynCustomTextAttributesImport = api.NewDefaultImports()
		lazSynCustomTextAttributesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazSynCustomTextAttributes_Create", 0), // constructor NewLazSynCustomTextAttributes
			/* 1 */ imports.NewTable("TLazSynCustomTextAttributes_Clear", 0), // procedure Clear
			/* 2 */ imports.NewTable("TLazSynCustomTextAttributes_SetAllPriorities", 0), // procedure SetAllPriorities
			/* 3 */ imports.NewTable("TLazSynCustomTextAttributes_BeginUpdate", 0), // procedure BeginUpdate
			/* 4 */ imports.NewTable("TLazSynCustomTextAttributes_EndUpdate", 0), // procedure EndUpdate
			/* 5 */ imports.NewTable("TLazSynCustomTextAttributes_Background", 0), // property Background
			/* 6 */ imports.NewTable("TLazSynCustomTextAttributes_Foreground", 0), // property Foreground
			/* 7 */ imports.NewTable("TLazSynCustomTextAttributes_FrameColor", 0), // property FrameColor
			/* 8 */ imports.NewTable("TLazSynCustomTextAttributes_FrameStyle", 0), // property FrameStyle
			/* 9 */ imports.NewTable("TLazSynCustomTextAttributes_FrameEdges", 0), // property FrameEdges
			/* 10 */ imports.NewTable("TLazSynCustomTextAttributes_Style", 0), // property Style
			/* 11 */ imports.NewTable("TLazSynCustomTextAttributes_StyleMask", 0), // property StyleMask
			/* 12 */ imports.NewTable("TLazSynCustomTextAttributes_StylePriority", 0), // property StylePriority
			/* 13 */ imports.NewTable("TLazSynCustomTextAttributes_BackPriority", 0), // property BackPriority
			/* 14 */ imports.NewTable("TLazSynCustomTextAttributes_ForePriority", 0), // property ForePriority
			/* 15 */ imports.NewTable("TLazSynCustomTextAttributes_FramePriority", 0), // property FramePriority
			/* 16 */ imports.NewTable("TLazSynCustomTextAttributes_BoldPriority", 0), // property BoldPriority
			/* 17 */ imports.NewTable("TLazSynCustomTextAttributes_ItalicPriority", 0), // property ItalicPriority
			/* 18 */ imports.NewTable("TLazSynCustomTextAttributes_UnderlinePriority", 0), // property UnderlinePriority
			/* 19 */ imports.NewTable("TLazSynCustomTextAttributes_StrikeOutPriority", 0), // property StrikeOutPriority
		}
	})
	return lazSynCustomTextAttributesImport
}
