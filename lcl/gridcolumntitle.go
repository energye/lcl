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

// IGridColumnTitle Parent: IPersistent
type IGridColumnTitle interface {
	IPersistent
	Column() IGridColumn                             // property
	Alignment() TAlignment                           // property
	SetAlignment(AValue TAlignment)                  // property
	Caption() string                                 // property
	SetCaption(AValue string)                        // property
	Color() TColor                                   // property
	SetColor(AValue TColor)                          // property
	Font() IFont                                     // property
	SetFont(AValue IFont)                            // property
	ImageIndex() TImageIndex                         // property
	SetImageIndex(AValue TImageIndex)                // property
	ImageLayout() TButtonLayout                      // property
	SetImageLayout(AValue TButtonLayout)             // property
	Layout() TTextLayout                             // property
	SetLayout(AValue TTextLayout)                    // property
	MultiLine() bool                                 // property
	SetMultiLine(AValue bool)                        // property
	PrefixOption() TPrefixOption                     // property
	SetPrefixOption(AValue TPrefixOption)            // property
	IsDefault() bool                                 // function
	FillTitleDefaultFont()                           // procedure
	FixDesignFontsPPI(ADesignTimePPI int32)          // procedure
	ScaleFontsPPI(AToPPI int32, AProportion float64) // procedure
}

// TGridColumnTitle Parent: TPersistent
type TGridColumnTitle struct {
	TPersistent
}

func NewGridColumnTitle(TheColumn IGridColumn) IGridColumnTitle {
	r1 := gridColumnTitleImportAPI().SysCallN(5, GetObjectUintptr(TheColumn))
	return AsGridColumnTitle(r1)
}

func (m *TGridColumnTitle) Column() IGridColumn {
	r1 := gridColumnTitleImportAPI().SysCallN(4, m.Instance())
	return AsGridColumn(r1)
}

func (m *TGridColumnTitle) Alignment() TAlignment {
	r1 := gridColumnTitleImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TGridColumnTitle) SetAlignment(AValue TAlignment) {
	gridColumnTitleImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Caption() string {
	r1 := gridColumnTitleImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TGridColumnTitle) SetCaption(AValue string) {
	gridColumnTitleImportAPI().SysCallN(1, 1, m.Instance(), PascalStr(AValue))
}

func (m *TGridColumnTitle) Color() TColor {
	r1 := gridColumnTitleImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGridColumnTitle) SetColor(AValue TColor) {
	gridColumnTitleImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Font() IFont {
	r1 := gridColumnTitleImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TGridColumnTitle) SetFont(AValue IFont) {
	gridColumnTitleImportAPI().SysCallN(8, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TGridColumnTitle) ImageIndex() TImageIndex {
	r1 := gridColumnTitleImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TGridColumnTitle) SetImageIndex(AValue TImageIndex) {
	gridColumnTitleImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) ImageLayout() TButtonLayout {
	r1 := gridColumnTitleImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TButtonLayout(r1)
}

func (m *TGridColumnTitle) SetImageLayout(AValue TButtonLayout) {
	gridColumnTitleImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) Layout() TTextLayout {
	r1 := gridColumnTitleImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TGridColumnTitle) SetLayout(AValue TTextLayout) {
	gridColumnTitleImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) MultiLine() bool {
	r1 := gridColumnTitleImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGridColumnTitle) SetMultiLine(AValue bool) {
	gridColumnTitleImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGridColumnTitle) PrefixOption() TPrefixOption {
	r1 := gridColumnTitleImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TPrefixOption(r1)
}

func (m *TGridColumnTitle) SetPrefixOption(AValue TPrefixOption) {
	gridColumnTitleImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TGridColumnTitle) IsDefault() bool {
	r1 := gridColumnTitleImportAPI().SysCallN(11, m.Instance())
	return GoBool(r1)
}

func GridColumnTitleClass() TClass {
	ret := gridColumnTitleImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TGridColumnTitle) FillTitleDefaultFont() {
	gridColumnTitleImportAPI().SysCallN(6, m.Instance())
}

func (m *TGridColumnTitle) FixDesignFontsPPI(ADesignTimePPI int32) {
	gridColumnTitleImportAPI().SysCallN(7, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TGridColumnTitle) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	gridColumnTitleImportAPI().SysCallN(15, m.Instance(), uintptr(AToPPI), uintptr(unsafePointer(&AProportion)))
}

var (
	gridColumnTitleImport       *imports.Imports = nil
	gridColumnTitleImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("GridColumnTitle_Alignment", 0),
		/*1*/ imports.NewTable("GridColumnTitle_Caption", 0),
		/*2*/ imports.NewTable("GridColumnTitle_Class", 0),
		/*3*/ imports.NewTable("GridColumnTitle_Color", 0),
		/*4*/ imports.NewTable("GridColumnTitle_Column", 0),
		/*5*/ imports.NewTable("GridColumnTitle_Create", 0),
		/*6*/ imports.NewTable("GridColumnTitle_FillTitleDefaultFont", 0),
		/*7*/ imports.NewTable("GridColumnTitle_FixDesignFontsPPI", 0),
		/*8*/ imports.NewTable("GridColumnTitle_Font", 0),
		/*9*/ imports.NewTable("GridColumnTitle_ImageIndex", 0),
		/*10*/ imports.NewTable("GridColumnTitle_ImageLayout", 0),
		/*11*/ imports.NewTable("GridColumnTitle_IsDefault", 0),
		/*12*/ imports.NewTable("GridColumnTitle_Layout", 0),
		/*13*/ imports.NewTable("GridColumnTitle_MultiLine", 0),
		/*14*/ imports.NewTable("GridColumnTitle_PrefixOption", 0),
		/*15*/ imports.NewTable("GridColumnTitle_ScaleFontsPPI", 0),
	}
)

func gridColumnTitleImportAPI() *imports.Imports {
	if gridColumnTitleImport == nil {
		gridColumnTitleImport = NewDefaultImports()
		gridColumnTitleImport.SetImportTable(gridColumnTitleImportTables)
		gridColumnTitleImportTables = nil
	}
	return gridColumnTitleImport
}
