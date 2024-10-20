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

// IVTHeader Parent: IPersistent
type IVTHeader interface {
	IPersistent
	DragImage() IVTDragImage                                                                                            // property
	States() THeaderStates                                                                                              // property
	Treeview() IBaseVirtualTree                                                                                         // property
	UseColumns() bool                                                                                                   // property
	AutoSizeIndex() TColumnIndex                                                                                        // property
	SetAutoSizeIndex(AValue TColumnIndex)                                                                               // property
	Background() TColor                                                                                                 // property
	SetBackground(AValue TColor)                                                                                        // property
	Columns() IVirtualTreeColumns                                                                                       // property
	SetColumns(AValue IVirtualTreeColumns)                                                                              // property
	DefaultHeight() int32                                                                                               // property
	SetDefaultHeight(AValue int32)                                                                                      // property
	Font() IFont                                                                                                        // property
	SetFont(AValue IFont)                                                                                               // property
	FixedAreaConstraints() IVTFixedAreaConstraints                                                                      // property
	SetFixedAreaConstraints(AValue IVTFixedAreaConstraints)                                                             // property
	Height() int32                                                                                                      // property
	SetHeight(AValue int32)                                                                                             // property
	Images() ICustomImageList                                                                                           // property
	SetImages(AValue ICustomImageList)                                                                                  // property
	ImagesWidth() int32                                                                                                 // property
	SetImagesWidth(AValue int32)                                                                                        // property
	MainColumn() TColumnIndex                                                                                           // property
	SetMainColumn(AValue TColumnIndex)                                                                                  // property
	MaxHeight() int32                                                                                                   // property
	SetMaxHeight(AValue int32)                                                                                          // property
	MinHeight() int32                                                                                                   // property
	SetMinHeight(AValue int32)                                                                                          // property
	Options() TVTHeaderOptions                                                                                          // property
	SetOptions(AValue TVTHeaderOptions)                                                                                 // property
	ParentFont() bool                                                                                                   // property
	SetParentFont(AValue bool)                                                                                          // property
	PopupMenu() IPopupMenu                                                                                              // property
	SetPopupMenu(AValue IPopupMenu)                                                                                     // property
	SortColumn() TColumnIndex                                                                                           // property
	SetSortColumn(AValue TColumnIndex)                                                                                  // property
	SortDirection() TSortDirection                                                                                      // property
	SetSortDirection(AValue TSortDirection)                                                                             // property
	Style() TVTHeaderStyle                                                                                              // property
	SetStyle(AValue TVTHeaderStyle)                                                                                     // property
	AllowFocus(ColumnIndex TColumnIndex) bool                                                                           // function
	InHeader(P *TPoint) bool                                                                                            // function
	InHeaderSplitterArea(P *TPoint) bool                                                                                // function
	ResizeColumns(ChangeBy int32, RangeStartCol TColumnIndex, RangeEndCol TColumnIndex, Options TVTColumnOptions) int32 // function
	AutoFitColumns(Animated bool, SmartAutoFitType TSmartAutoFitType, RangeStartCol int32, RangeEndCol int32)           // procedure
	FixDesignFontsPPI(ADesignTimePPI int32)                                                                             // procedure
	Invalidate(Column IVirtualTreeColumn, ExpandToBorder bool)                                                          // procedure
	LoadFromStream(Stream IStream)                                                                                      // procedure
	RestoreColumns()                                                                                                    // procedure
	SaveToStream(Stream IStream)                                                                                        // procedure
}

// TVTHeader Parent: TPersistent
type TVTHeader struct {
	TPersistent
}

func NewVTHeader(AOwner IBaseVirtualTree) IVTHeader {
	r1 := vTHeaderImportAPI().SysCallN(6, GetObjectUintptr(AOwner))
	return AsVTHeader(r1)
}

func (m *TVTHeader) DragImage() IVTDragImage {
	r1 := vTHeaderImportAPI().SysCallN(8, m.Instance())
	return AsVTDragImage(r1)
}

func (m *TVTHeader) States() THeaderStates {
	r1 := vTHeaderImportAPI().SysCallN(30, m.Instance())
	return THeaderStates(r1)
}

func (m *TVTHeader) Treeview() IBaseVirtualTree {
	r1 := vTHeaderImportAPI().SysCallN(32, m.Instance())
	return AsBaseVirtualTree(r1)
}

func (m *TVTHeader) UseColumns() bool {
	r1 := vTHeaderImportAPI().SysCallN(33, m.Instance())
	return GoBool(r1)
}

func (m *TVTHeader) AutoSizeIndex() TColumnIndex {
	r1 := vTHeaderImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TColumnIndex(r1)
}

func (m *TVTHeader) SetAutoSizeIndex(AValue TColumnIndex) {
	vTHeaderImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Background() TColor {
	r1 := vTHeaderImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTHeader) SetBackground(AValue TColor) {
	vTHeaderImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Columns() IVirtualTreeColumns {
	r1 := vTHeaderImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsVirtualTreeColumns(r1)
}

func (m *TVTHeader) SetColumns(AValue IVirtualTreeColumns) {
	vTHeaderImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) DefaultHeight() int32 {
	r1 := vTHeaderImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetDefaultHeight(AValue int32) {
	vTHeaderImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Font() IFont {
	r1 := vTHeaderImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TVTHeader) SetFont(AValue IFont) {
	vTHeaderImportAPI().SysCallN(11, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) FixedAreaConstraints() IVTFixedAreaConstraints {
	r1 := vTHeaderImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsVTFixedAreaConstraints(r1)
}

func (m *TVTHeader) SetFixedAreaConstraints(AValue IVTFixedAreaConstraints) {
	vTHeaderImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) Height() int32 {
	r1 := vTHeaderImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetHeight(AValue int32) {
	vTHeaderImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Images() ICustomImageList {
	r1 := vTHeaderImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TVTHeader) SetImages(AValue ICustomImageList) {
	vTHeaderImportAPI().SysCallN(13, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) ImagesWidth() int32 {
	r1 := vTHeaderImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetImagesWidth(AValue int32) {
	vTHeaderImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) MainColumn() TColumnIndex {
	r1 := vTHeaderImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TColumnIndex(r1)
}

func (m *TVTHeader) SetMainColumn(AValue TColumnIndex) {
	vTHeaderImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) MaxHeight() int32 {
	r1 := vTHeaderImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetMaxHeight(AValue int32) {
	vTHeaderImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) MinHeight() int32 {
	r1 := vTHeaderImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetMinHeight(AValue int32) {
	vTHeaderImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Options() TVTHeaderOptions {
	r1 := vTHeaderImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return TVTHeaderOptions(r1)
}

func (m *TVTHeader) SetOptions(AValue TVTHeaderOptions) {
	vTHeaderImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) ParentFont() bool {
	r1 := vTHeaderImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TVTHeader) SetParentFont(AValue bool) {
	vTHeaderImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TVTHeader) PopupMenu() IPopupMenu {
	r1 := vTHeaderImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return AsPopupMenu(r1)
}

func (m *TVTHeader) SetPopupMenu(AValue IPopupMenu) {
	vTHeaderImportAPI().SysCallN(24, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) SortColumn() TColumnIndex {
	r1 := vTHeaderImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return TColumnIndex(r1)
}

func (m *TVTHeader) SetSortColumn(AValue TColumnIndex) {
	vTHeaderImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) SortDirection() TSortDirection {
	r1 := vTHeaderImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return TSortDirection(r1)
}

func (m *TVTHeader) SetSortDirection(AValue TSortDirection) {
	vTHeaderImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Style() TVTHeaderStyle {
	r1 := vTHeaderImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return TVTHeaderStyle(r1)
}

func (m *TVTHeader) SetStyle(AValue TVTHeaderStyle) {
	vTHeaderImportAPI().SysCallN(31, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) AllowFocus(ColumnIndex TColumnIndex) bool {
	r1 := vTHeaderImportAPI().SysCallN(0, m.Instance(), uintptr(ColumnIndex))
	return GoBool(r1)
}

func (m *TVTHeader) InHeader(P *TPoint) bool {
	r1 := vTHeaderImportAPI().SysCallN(15, m.Instance(), uintptr(unsafePointer(P)))
	return GoBool(r1)
}

func (m *TVTHeader) InHeaderSplitterArea(P *TPoint) bool {
	r1 := vTHeaderImportAPI().SysCallN(16, m.Instance(), uintptr(unsafePointer(P)))
	return GoBool(r1)
}

func (m *TVTHeader) ResizeColumns(ChangeBy int32, RangeStartCol TColumnIndex, RangeEndCol TColumnIndex, Options TVTColumnOptions) int32 {
	r1 := vTHeaderImportAPI().SysCallN(25, m.Instance(), uintptr(ChangeBy), uintptr(RangeStartCol), uintptr(RangeEndCol), uintptr(Options))
	return int32(r1)
}

func VTHeaderClass() TClass {
	ret := vTHeaderImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TVTHeader) AutoFitColumns(Animated bool, SmartAutoFitType TSmartAutoFitType, RangeStartCol int32, RangeEndCol int32) {
	vTHeaderImportAPI().SysCallN(1, m.Instance(), PascalBool(Animated), uintptr(SmartAutoFitType), uintptr(RangeStartCol), uintptr(RangeEndCol))
}

func (m *TVTHeader) FixDesignFontsPPI(ADesignTimePPI int32) {
	vTHeaderImportAPI().SysCallN(9, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TVTHeader) Invalidate(Column IVirtualTreeColumn, ExpandToBorder bool) {
	vTHeaderImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(Column), PascalBool(ExpandToBorder))
}

func (m *TVTHeader) LoadFromStream(Stream IStream) {
	vTHeaderImportAPI().SysCallN(18, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TVTHeader) RestoreColumns() {
	vTHeaderImportAPI().SysCallN(26, m.Instance())
}

func (m *TVTHeader) SaveToStream(Stream IStream) {
	vTHeaderImportAPI().SysCallN(27, m.Instance(), GetObjectUintptr(Stream))
}

var (
	vTHeaderImport       *imports.Imports = nil
	vTHeaderImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("VTHeader_AllowFocus", 0),
		/*1*/ imports.NewTable("VTHeader_AutoFitColumns", 0),
		/*2*/ imports.NewTable("VTHeader_AutoSizeIndex", 0),
		/*3*/ imports.NewTable("VTHeader_Background", 0),
		/*4*/ imports.NewTable("VTHeader_Class", 0),
		/*5*/ imports.NewTable("VTHeader_Columns", 0),
		/*6*/ imports.NewTable("VTHeader_Create", 0),
		/*7*/ imports.NewTable("VTHeader_DefaultHeight", 0),
		/*8*/ imports.NewTable("VTHeader_DragImage", 0),
		/*9*/ imports.NewTable("VTHeader_FixDesignFontsPPI", 0),
		/*10*/ imports.NewTable("VTHeader_FixedAreaConstraints", 0),
		/*11*/ imports.NewTable("VTHeader_Font", 0),
		/*12*/ imports.NewTable("VTHeader_Height", 0),
		/*13*/ imports.NewTable("VTHeader_Images", 0),
		/*14*/ imports.NewTable("VTHeader_ImagesWidth", 0),
		/*15*/ imports.NewTable("VTHeader_InHeader", 0),
		/*16*/ imports.NewTable("VTHeader_InHeaderSplitterArea", 0),
		/*17*/ imports.NewTable("VTHeader_Invalidate", 0),
		/*18*/ imports.NewTable("VTHeader_LoadFromStream", 0),
		/*19*/ imports.NewTable("VTHeader_MainColumn", 0),
		/*20*/ imports.NewTable("VTHeader_MaxHeight", 0),
		/*21*/ imports.NewTable("VTHeader_MinHeight", 0),
		/*22*/ imports.NewTable("VTHeader_Options", 0),
		/*23*/ imports.NewTable("VTHeader_ParentFont", 0),
		/*24*/ imports.NewTable("VTHeader_PopupMenu", 0),
		/*25*/ imports.NewTable("VTHeader_ResizeColumns", 0),
		/*26*/ imports.NewTable("VTHeader_RestoreColumns", 0),
		/*27*/ imports.NewTable("VTHeader_SaveToStream", 0),
		/*28*/ imports.NewTable("VTHeader_SortColumn", 0),
		/*29*/ imports.NewTable("VTHeader_SortDirection", 0),
		/*30*/ imports.NewTable("VTHeader_States", 0),
		/*31*/ imports.NewTable("VTHeader_Style", 0),
		/*32*/ imports.NewTable("VTHeader_Treeview", 0),
		/*33*/ imports.NewTable("VTHeader_UseColumns", 0),
	}
)

func vTHeaderImportAPI() *imports.Imports {
	if vTHeaderImport == nil {
		vTHeaderImport = NewDefaultImports()
		vTHeaderImport.SetImportTable(vTHeaderImportTables)
		vTHeaderImportTables = nil
	}
	return vTHeaderImport
}
