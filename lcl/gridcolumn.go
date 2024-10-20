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

// IGridColumn Parent: ICollectionItem
type IGridColumn interface {
	ICollectionItem
	Grid() ICustomGrid                               // property
	DefaultWidth() int32                             // property
	StoredWidth() int32                              // property
	WidthChanged() bool                              // property
	Alignment() TAlignment                           // property
	SetAlignment(AValue TAlignment)                  // property
	ButtonStyle() TColumnButtonStyle                 // property
	SetButtonStyle(AValue TColumnButtonStyle)        // property
	Color() TColor                                   // property
	SetColor(AValue TColor)                          // property
	DropDownRows() int32                             // property
	SetDropDownRows(AValue int32)                    // property
	Expanded() bool                                  // property
	SetExpanded(AValue bool)                         // property
	Font() IFont                                     // property
	SetFont(AValue IFont)                            // property
	Layout() TTextLayout                             // property
	SetLayout(AValue TTextLayout)                    // property
	MinSize() int32                                  // property
	SetMinSize(AValue int32)                         // property
	MaxSize() int32                                  // property
	SetMaxSize(AValue int32)                         // property
	PickList() IStrings                              // property
	SetPickList(AValue IStrings)                     // property
	ReadOnly() bool                                  // property
	SetReadOnly(AValue bool)                         // property
	SizePriority() int32                             // property
	SetSizePriority(AValue int32)                    // property
	Tag() uint32                                     // property
	SetTag(AValue uint32)                            // property
	Title() IGridColumnTitle                         // property
	SetTitle(AValue IGridColumnTitle)                // property
	Width() int32                                    // property
	SetWidth(AValue int32)                           // property
	Visible() bool                                   // property
	SetVisible(AValue bool)                          // property
	ValueChecked() string                            // property
	SetValueChecked(AValue string)                   // property
	ValueUnchecked() string                          // property
	SetValueUnchecked(AValue string)                 // property
	IsDefault() bool                                 // function
	FillDefaultFont()                                // procedure
	FixDesignFontsPPI(ADesignTimePPI int32)          // procedure
	ScaleFontsPPI(AToPPI int32, AProportion float64) // procedure
}

// TGridColumn Parent: TCollectionItem
type TGridColumn struct {
	TCollectionItem
}

func NewGridColumn(ACollection ICollection) IGridColumn {
	r1 := gridColumnImportAPI().SysCallN(4, GetObjectUintptr(ACollection))
	return AsGridColumn(r1)
}

func (m *TGridColumn) Grid() ICustomGrid {
	r1 := gridColumnImportAPI().SysCallN(11, m.Instance())
	return AsCustomGrid(r1)
}

func (m *TGridColumn) DefaultWidth() int32 {
	r1 := gridColumnImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TGridColumn) StoredWidth() int32 {
	r1 := gridColumnImportAPI().SysCallN(20, m.Instance())
	return int32(r1)
}

func (m *TGridColumn) WidthChanged() bool {
	r1 := gridColumnImportAPI().SysCallN(27, m.Instance())
	return GoBool(r1)
}

func (m *TGridColumn) Alignment() TAlignment {
	r1 := gridColumnImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TGridColumn) SetAlignment(AValue TAlignment) {
	gridColumnImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) ButtonStyle() TColumnButtonStyle {
	r1 := gridColumnImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TColumnButtonStyle(r1)
}

func (m *TGridColumn) SetButtonStyle(AValue TColumnButtonStyle) {
	gridColumnImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Color() TColor {
	r1 := gridColumnImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGridColumn) SetColor(AValue TColor) {
	gridColumnImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) DropDownRows() int32 {
	r1 := gridColumnImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetDropDownRows(AValue int32) {
	gridColumnImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Expanded() bool {
	r1 := gridColumnImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumn) SetExpanded(AValue bool) {
	gridColumnImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumn) Font() IFont {
	r1 := gridColumnImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TGridColumn) SetFont(AValue IFont) {
	gridColumnImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumn) Layout() TTextLayout {
	r1 := gridColumnImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TGridColumn) SetLayout(AValue TTextLayout) {
	gridColumnImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) MinSize() int32 {
	r1 := gridColumnImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetMinSize(AValue int32) {
	gridColumnImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) MaxSize() int32 {
	r1 := gridColumnImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetMaxSize(AValue int32) {
	gridColumnImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) PickList() IStrings {
	r1 := gridColumnImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TGridColumn) SetPickList(AValue IStrings) {
	gridColumnImportAPI().SysCallN(16, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumn) ReadOnly() bool {
	r1 := gridColumnImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumn) SetReadOnly(AValue bool) {
	gridColumnImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumn) SizePriority() int32 {
	r1 := gridColumnImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetSizePriority(AValue int32) {
	gridColumnImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Tag() uint32 {
	r1 := gridColumnImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TGridColumn) SetTag(AValue uint32) {
	gridColumnImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Title() IGridColumnTitle {
	r1 := gridColumnImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return AsGridColumnTitle(r1)
}

func (m *TGridColumn) SetTitle(AValue IGridColumnTitle) {
	gridColumnImportAPI().SysCallN(22, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumn) Width() int32 {
	r1 := gridColumnImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGridColumn) SetWidth(AValue int32) {
	gridColumnImportAPI().SysCallN(26, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumn) Visible() bool {
	r1 := gridColumnImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumn) SetVisible(AValue bool) {
	gridColumnImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumn) ValueChecked() string {
	r1 := gridColumnImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumn) SetValueChecked(AValue string) {
	gridColumnImportAPI().SysCallN(23, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumn) ValueUnchecked() string {
	r1 := gridColumnImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumn) SetValueUnchecked(AValue string) {
	gridColumnImportAPI().SysCallN(24, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumn) IsDefault() bool {
	r1 := gridColumnImportAPI().SysCallN(12, m.Instance())
	return GoBool(r1)
}

func GridColumnClass() TClass {
	ret := gridColumnImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TGridColumn) FillDefaultFont() {
	gridColumnImportAPI().SysCallN(8, m.Instance())
}

func (m *TGridColumn) FixDesignFontsPPI(ADesignTimePPI int32) {
	gridColumnImportAPI().SysCallN(9, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TGridColumn) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	gridColumnImportAPI().SysCallN(18, m.Instance(), uintptr(AToPPI), uintptr(unsafePointer(&AProportion)))
}

var (
	gridColumnImport       *imports.Imports = nil
	gridColumnImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("GridColumn_Alignment", 0),
		/*1*/ imports.NewTable("GridColumn_ButtonStyle", 0),
		/*2*/ imports.NewTable("GridColumn_Class", 0),
		/*3*/ imports.NewTable("GridColumn_Color", 0),
		/*4*/ imports.NewTable("GridColumn_Create", 0),
		/*5*/ imports.NewTable("GridColumn_DefaultWidth", 0),
		/*6*/ imports.NewTable("GridColumn_DropDownRows", 0),
		/*7*/ imports.NewTable("GridColumn_Expanded", 0),
		/*8*/ imports.NewTable("GridColumn_FillDefaultFont", 0),
		/*9*/ imports.NewTable("GridColumn_FixDesignFontsPPI", 0),
		/*10*/ imports.NewTable("GridColumn_Font", 0),
		/*11*/ imports.NewTable("GridColumn_Grid", 0),
		/*12*/ imports.NewTable("GridColumn_IsDefault", 0),
		/*13*/ imports.NewTable("GridColumn_Layout", 0),
		/*14*/ imports.NewTable("GridColumn_MaxSize", 0),
		/*15*/ imports.NewTable("GridColumn_MinSize", 0),
		/*16*/ imports.NewTable("GridColumn_PickList", 0),
		/*17*/ imports.NewTable("GridColumn_ReadOnly", 0),
		/*18*/ imports.NewTable("GridColumn_ScaleFontsPPI", 0),
		/*19*/ imports.NewTable("GridColumn_SizePriority", 0),
		/*20*/ imports.NewTable("GridColumn_StoredWidth", 0),
		/*21*/ imports.NewTable("GridColumn_Tag", 0),
		/*22*/ imports.NewTable("GridColumn_Title", 0),
		/*23*/ imports.NewTable("GridColumn_ValueChecked", 0),
		/*24*/ imports.NewTable("GridColumn_ValueUnchecked", 0),
		/*25*/ imports.NewTable("GridColumn_Visible", 0),
		/*26*/ imports.NewTable("GridColumn_Width", 0),
		/*27*/ imports.NewTable("GridColumn_WidthChanged", 0),
	}
)

func gridColumnImportAPI() *imports.Imports {
	if gridColumnImport == nil {
		gridColumnImport = NewDefaultImports()
		gridColumnImport.SetImportTable(gridColumnImportTables)
		gridColumnImportTables = nil
	}
	return gridColumnImport
}
