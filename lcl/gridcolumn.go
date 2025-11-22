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

// IGridColumn Parent: ICollectionItem
type IGridColumn interface {
	ICollectionItem
	IsDefault() bool                               // function
	FillDefaultFont()                              // procedure
	FixDesignFontsPPI(designTimePPI int32)         // procedure
	ScaleFontsPPI(toPPI int32, proportion float64) // procedure
	Grid() ICustomGrid                             // property Grid Getter
	DefaultWidth() int32                           // property DefaultWidth Getter
	StoredWidth() int32                            // property StoredWidth Getter
	WidthChanged() bool                            // property WidthChanged Getter
	Alignment() types.TAlignment                   // property Alignment Getter
	SetAlignment(value types.TAlignment)           // property Alignment Setter
	ButtonStyle() types.TColumnButtonStyle         // property ButtonStyle Getter
	SetButtonStyle(value types.TColumnButtonStyle) // property ButtonStyle Setter
	Color() types.TColor                           // property Color Getter
	SetColor(value types.TColor)                   // property Color Setter
	DropDownRows() int32                           // property DropDownRows Getter
	SetDropDownRows(value int32)                   // property DropDownRows Setter
	Expanded() bool                                // property Expanded Getter
	SetExpanded(value bool)                        // property Expanded Setter
	Font() IFont                                   // property Font Getter
	SetFont(value IFont)                           // property Font Setter
	Layout() types.TTextLayout                     // property Layout Getter
	SetLayout(value types.TTextLayout)             // property Layout Setter
	MinSize() int32                                // property MinSize Getter
	SetMinSize(value int32)                        // property MinSize Setter
	MaxSize() int32                                // property MaxSize Getter
	SetMaxSize(value int32)                        // property MaxSize Setter
	PickList() IStrings                            // property PickList Getter
	SetPickList(value IStrings)                    // property PickList Setter
	ReadOnly() bool                                // property ReadOnly Getter
	SetReadOnly(value bool)                        // property ReadOnly Setter
	SizePriority() int32                           // property SizePriority Getter
	SetSizePriority(value int32)                   // property SizePriority Setter
	Tag() uintptr                                  // property Tag Getter
	SetTag(value uintptr)                          // property Tag Setter
	Title() IGridColumnTitle                       // property Title Getter
	SetTitle(value IGridColumnTitle)               // property Title Setter
	Width() int32                                  // property Width Getter
	SetWidth(value int32)                          // property Width Setter
	Visible() bool                                 // property Visible Getter
	SetVisible(value bool)                         // property Visible Setter
	ValueChecked() string                          // property ValueChecked Getter
	SetValueChecked(value string)                  // property ValueChecked Setter
	ValueUnchecked() string                        // property ValueUnchecked Getter
	SetValueUnchecked(value string)                // property ValueUnchecked Setter
}

type TGridColumn struct {
	TCollectionItem
}

func (m *TGridColumn) IsDefault() bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TGridColumn) FillDefaultFont() {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(2, m.Instance())
}

func (m *TGridColumn) FixDesignFontsPPI(designTimePPI int32) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(3, m.Instance(), uintptr(designTimePPI))
}

func (m *TGridColumn) ScaleFontsPPI(toPPI int32, proportion float64) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(4, m.Instance(), uintptr(toPPI), uintptr(base.UnsafePointer(&proportion)))
}

func (m *TGridColumn) Grid() ICustomGrid {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnAPI().SysCallN(5, m.Instance())
	return AsCustomGrid(r)
}

func (m *TGridColumn) DefaultWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TGridColumn) StoredWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TGridColumn) WidthChanged() bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TGridColumn) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(9, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TGridColumn) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) ButtonStyle() types.TColumnButtonStyle {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(10, 0, m.Instance())
	return types.TColumnButtonStyle(r)
}

func (m *TGridColumn) SetButtonStyle(value types.TColumnButtonStyle) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(11, 0, m.Instance())
	return types.TColor(r)
}

func (m *TGridColumn) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) DropDownRows() int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TGridColumn) SetDropDownRows(value int32) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) Expanded() bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGridColumn) SetExpanded(value bool) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TGridColumn) Font() IFont {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnAPI().SysCallN(14, 0, m.Instance())
	return AsFont(r)
}

func (m *TGridColumn) SetFont(value IFont) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TGridColumn) Layout() types.TTextLayout {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(15, 0, m.Instance())
	return types.TTextLayout(r)
}

func (m *TGridColumn) SetLayout(value types.TTextLayout) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) MinSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TGridColumn) SetMinSize(value int32) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) MaxSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TGridColumn) SetMaxSize(value int32) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) PickList() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnAPI().SysCallN(18, 0, m.Instance())
	return AsStrings(r)
}

func (m *TGridColumn) SetPickList(value IStrings) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(18, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TGridColumn) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGridColumn) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TGridColumn) SizePriority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TGridColumn) SetSizePriority(value int32) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) Tag() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(21, 0, m.Instance())
	return uintptr(r)
}

func (m *TGridColumn) SetTag(value uintptr) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) Title() IGridColumnTitle {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnAPI().SysCallN(22, 0, m.Instance())
	return AsGridColumnTitle(r)
}

func (m *TGridColumn) SetTitle(value IGridColumnTitle) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TGridColumn) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TGridColumn) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TGridColumn) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGridColumn) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TGridColumn) ValueChecked() string {
	if !m.IsValid() {
		return ""
	}
	r := gridColumnAPI().SysCallN(25, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TGridColumn) SetValueChecked(value string) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(25, 1, m.Instance(), api.PasStr(value))
}

func (m *TGridColumn) ValueUnchecked() string {
	if !m.IsValid() {
		return ""
	}
	r := gridColumnAPI().SysCallN(26, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TGridColumn) SetValueUnchecked(value string) {
	if !m.IsValid() {
		return
	}
	gridColumnAPI().SysCallN(26, 1, m.Instance(), api.PasStr(value))
}

// NewGridColumn class constructor
func NewGridColumn(collection ICollection) IGridColumn {
	r := gridColumnAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsGridColumn(r)
}

var (
	gridColumnOnce   base.Once
	gridColumnImport *imports.Imports = nil
)

func gridColumnAPI() *imports.Imports {
	gridColumnOnce.Do(func() {
		gridColumnImport = api.NewDefaultImports()
		gridColumnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGridColumn_Create", 0), // constructor NewGridColumn
			/* 1 */ imports.NewTable("TGridColumn_IsDefault", 0), // function IsDefault
			/* 2 */ imports.NewTable("TGridColumn_FillDefaultFont", 0), // procedure FillDefaultFont
			/* 3 */ imports.NewTable("TGridColumn_FixDesignFontsPPI", 0), // procedure FixDesignFontsPPI
			/* 4 */ imports.NewTable("TGridColumn_ScaleFontsPPI", 0), // procedure ScaleFontsPPI
			/* 5 */ imports.NewTable("TGridColumn_Grid", 0), // property Grid
			/* 6 */ imports.NewTable("TGridColumn_DefaultWidth", 0), // property DefaultWidth
			/* 7 */ imports.NewTable("TGridColumn_StoredWidth", 0), // property StoredWidth
			/* 8 */ imports.NewTable("TGridColumn_WidthChanged", 0), // property WidthChanged
			/* 9 */ imports.NewTable("TGridColumn_Alignment", 0), // property Alignment
			/* 10 */ imports.NewTable("TGridColumn_ButtonStyle", 0), // property ButtonStyle
			/* 11 */ imports.NewTable("TGridColumn_Color", 0), // property Color
			/* 12 */ imports.NewTable("TGridColumn_DropDownRows", 0), // property DropDownRows
			/* 13 */ imports.NewTable("TGridColumn_Expanded", 0), // property Expanded
			/* 14 */ imports.NewTable("TGridColumn_Font", 0), // property Font
			/* 15 */ imports.NewTable("TGridColumn_Layout", 0), // property Layout
			/* 16 */ imports.NewTable("TGridColumn_MinSize", 0), // property MinSize
			/* 17 */ imports.NewTable("TGridColumn_MaxSize", 0), // property MaxSize
			/* 18 */ imports.NewTable("TGridColumn_PickList", 0), // property PickList
			/* 19 */ imports.NewTable("TGridColumn_ReadOnly", 0), // property ReadOnly
			/* 20 */ imports.NewTable("TGridColumn_SizePriority", 0), // property SizePriority
			/* 21 */ imports.NewTable("TGridColumn_Tag", 0), // property Tag
			/* 22 */ imports.NewTable("TGridColumn_Title", 0), // property Title
			/* 23 */ imports.NewTable("TGridColumn_Width", 0), // property Width
			/* 24 */ imports.NewTable("TGridColumn_Visible", 0), // property Visible
			/* 25 */ imports.NewTable("TGridColumn_ValueChecked", 0), // property ValueChecked
			/* 26 */ imports.NewTable("TGridColumn_ValueUnchecked", 0), // property ValueUnchecked
		}
	})
	return gridColumnImport
}
