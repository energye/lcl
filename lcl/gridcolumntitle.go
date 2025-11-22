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

// IGridColumnTitle Parent: IPersistent
type IGridColumnTitle interface {
	IPersistent
	IsDefault() bool                               // function
	FillTitleDefaultFont()                         // procedure
	FixDesignFontsPPI(designTimePPI int32)         // procedure
	ScaleFontsPPI(toPPI int32, proportion float64) // procedure
	Column() IGridColumn                           // property Column Getter
	Alignment() types.TAlignment                   // property Alignment Getter
	SetAlignment(value types.TAlignment)           // property Alignment Setter
	Caption() string                               // property Caption Getter
	SetCaption(value string)                       // property Caption Setter
	Color() types.TColor                           // property Color Getter
	SetColor(value types.TColor)                   // property Color Setter
	Font() IFont                                   // property Font Getter
	SetFont(value IFont)                           // property Font Setter
	ImageIndex() int32                             // property ImageIndex Getter
	SetImageIndex(value int32)                     // property ImageIndex Setter
	ImageLayout() types.TButtonLayout              // property ImageLayout Getter
	SetImageLayout(value types.TButtonLayout)      // property ImageLayout Setter
	Layout() types.TTextLayout                     // property Layout Getter
	SetLayout(value types.TTextLayout)             // property Layout Setter
	MultiLine() bool                               // property MultiLine Getter
	SetMultiLine(value bool)                       // property MultiLine Setter
	PrefixOption() types.TPrefixOption             // property PrefixOption Getter
	SetPrefixOption(value types.TPrefixOption)     // property PrefixOption Setter
}

type TGridColumnTitle struct {
	TPersistent
}

func (m *TGridColumnTitle) IsDefault() bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnTitleAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TGridColumnTitle) FillTitleDefaultFont() {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(2, m.Instance())
}

func (m *TGridColumnTitle) FixDesignFontsPPI(designTimePPI int32) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(3, m.Instance(), uintptr(designTimePPI))
}

func (m *TGridColumnTitle) ScaleFontsPPI(toPPI int32, proportion float64) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(4, m.Instance(), uintptr(toPPI), uintptr(base.UnsafePointer(&proportion)))
}

func (m *TGridColumnTitle) Column() IGridColumn {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnTitleAPI().SysCallN(5, m.Instance())
	return AsGridColumn(r)
}

func (m *TGridColumnTitle) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnTitleAPI().SysCallN(6, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TGridColumnTitle) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumnTitle) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := gridColumnTitleAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TGridColumnTitle) SetCaption(value string) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TGridColumnTitle) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnTitleAPI().SysCallN(8, 0, m.Instance())
	return types.TColor(r)
}

func (m *TGridColumnTitle) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumnTitle) Font() IFont {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnTitleAPI().SysCallN(9, 0, m.Instance())
	return AsFont(r)
}

func (m *TGridColumnTitle) SetFont(value IFont) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(9, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TGridColumnTitle) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnTitleAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TGridColumnTitle) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumnTitle) ImageLayout() types.TButtonLayout {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnTitleAPI().SysCallN(11, 0, m.Instance())
	return types.TButtonLayout(r)
}

func (m *TGridColumnTitle) SetImageLayout(value types.TButtonLayout) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumnTitle) Layout() types.TTextLayout {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnTitleAPI().SysCallN(12, 0, m.Instance())
	return types.TTextLayout(r)
}

func (m *TGridColumnTitle) SetLayout(value types.TTextLayout) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumnTitle) MultiLine() bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnTitleAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGridColumnTitle) SetMultiLine(value bool) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TGridColumnTitle) PrefixOption() types.TPrefixOption {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnTitleAPI().SysCallN(14, 0, m.Instance())
	return types.TPrefixOption(r)
}

func (m *TGridColumnTitle) SetPrefixOption(value types.TPrefixOption) {
	if !m.IsValid() {
		return
	}
	gridColumnTitleAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

// NewGridColumnTitle class constructor
func NewGridColumnTitle(theColumn IGridColumn) IGridColumnTitle {
	r := gridColumnTitleAPI().SysCallN(0, base.GetObjectUintptr(theColumn))
	return AsGridColumnTitle(r)
}

var (
	gridColumnTitleOnce   base.Once
	gridColumnTitleImport *imports.Imports = nil
)

func gridColumnTitleAPI() *imports.Imports {
	gridColumnTitleOnce.Do(func() {
		gridColumnTitleImport = api.NewDefaultImports()
		gridColumnTitleImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGridColumnTitle_Create", 0), // constructor NewGridColumnTitle
			/* 1 */ imports.NewTable("TGridColumnTitle_IsDefault", 0), // function IsDefault
			/* 2 */ imports.NewTable("TGridColumnTitle_FillTitleDefaultFont", 0), // procedure FillTitleDefaultFont
			/* 3 */ imports.NewTable("TGridColumnTitle_FixDesignFontsPPI", 0), // procedure FixDesignFontsPPI
			/* 4 */ imports.NewTable("TGridColumnTitle_ScaleFontsPPI", 0), // procedure ScaleFontsPPI
			/* 5 */ imports.NewTable("TGridColumnTitle_Column", 0), // property Column
			/* 6 */ imports.NewTable("TGridColumnTitle_Alignment", 0), // property Alignment
			/* 7 */ imports.NewTable("TGridColumnTitle_Caption", 0), // property Caption
			/* 8 */ imports.NewTable("TGridColumnTitle_Color", 0), // property Color
			/* 9 */ imports.NewTable("TGridColumnTitle_Font", 0), // property Font
			/* 10 */ imports.NewTable("TGridColumnTitle_ImageIndex", 0), // property ImageIndex
			/* 11 */ imports.NewTable("TGridColumnTitle_ImageLayout", 0), // property ImageLayout
			/* 12 */ imports.NewTable("TGridColumnTitle_Layout", 0), // property Layout
			/* 13 */ imports.NewTable("TGridColumnTitle_MultiLine", 0), // property MultiLine
			/* 14 */ imports.NewTable("TGridColumnTitle_PrefixOption", 0), // property PrefixOption
		}
	})
	return gridColumnTitleImport
}
