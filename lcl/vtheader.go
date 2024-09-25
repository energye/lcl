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
	r1 := LCL().SysCallN(5905, GetObjectUintptr(AOwner))
	return AsVTHeader(r1)
}

func (m *TVTHeader) DragImage() IVTDragImage {
	r1 := LCL().SysCallN(5907, m.Instance())
	return AsVTDragImage(r1)
}

func (m *TVTHeader) States() THeaderStates {
	r1 := LCL().SysCallN(5929, m.Instance())
	return THeaderStates(r1)
}

func (m *TVTHeader) Treeview() IBaseVirtualTree {
	r1 := LCL().SysCallN(5931, m.Instance())
	return AsBaseVirtualTree(r1)
}

func (m *TVTHeader) UseColumns() bool {
	r1 := LCL().SysCallN(5932, m.Instance())
	return GoBool(r1)
}

func (m *TVTHeader) AutoSizeIndex() TColumnIndex {
	r1 := LCL().SysCallN(5901, 0, m.Instance(), 0)
	return TColumnIndex(r1)
}

func (m *TVTHeader) SetAutoSizeIndex(AValue TColumnIndex) {
	LCL().SysCallN(5901, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Background() TColor {
	r1 := LCL().SysCallN(5902, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVTHeader) SetBackground(AValue TColor) {
	LCL().SysCallN(5902, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Columns() IVirtualTreeColumns {
	r1 := LCL().SysCallN(5904, 0, m.Instance(), 0)
	return AsVirtualTreeColumns(r1)
}

func (m *TVTHeader) SetColumns(AValue IVirtualTreeColumns) {
	LCL().SysCallN(5904, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) DefaultHeight() int32 {
	r1 := LCL().SysCallN(5906, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetDefaultHeight(AValue int32) {
	LCL().SysCallN(5906, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Font() IFont {
	r1 := LCL().SysCallN(5910, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TVTHeader) SetFont(AValue IFont) {
	LCL().SysCallN(5910, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) FixedAreaConstraints() IVTFixedAreaConstraints {
	r1 := LCL().SysCallN(5909, 0, m.Instance(), 0)
	return AsVTFixedAreaConstraints(r1)
}

func (m *TVTHeader) SetFixedAreaConstraints(AValue IVTFixedAreaConstraints) {
	LCL().SysCallN(5909, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) Height() int32 {
	r1 := LCL().SysCallN(5911, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetHeight(AValue int32) {
	LCL().SysCallN(5911, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Images() ICustomImageList {
	r1 := LCL().SysCallN(5912, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TVTHeader) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(5912, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) ImagesWidth() int32 {
	r1 := LCL().SysCallN(5913, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetImagesWidth(AValue int32) {
	LCL().SysCallN(5913, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) MainColumn() TColumnIndex {
	r1 := LCL().SysCallN(5918, 0, m.Instance(), 0)
	return TColumnIndex(r1)
}

func (m *TVTHeader) SetMainColumn(AValue TColumnIndex) {
	LCL().SysCallN(5918, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) MaxHeight() int32 {
	r1 := LCL().SysCallN(5919, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetMaxHeight(AValue int32) {
	LCL().SysCallN(5919, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) MinHeight() int32 {
	r1 := LCL().SysCallN(5920, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVTHeader) SetMinHeight(AValue int32) {
	LCL().SysCallN(5920, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Options() TVTHeaderOptions {
	r1 := LCL().SysCallN(5921, 0, m.Instance(), 0)
	return TVTHeaderOptions(r1)
}

func (m *TVTHeader) SetOptions(AValue TVTHeaderOptions) {
	LCL().SysCallN(5921, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) ParentFont() bool {
	r1 := LCL().SysCallN(5922, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TVTHeader) SetParentFont(AValue bool) {
	LCL().SysCallN(5922, 1, m.Instance(), PascalBool(AValue))
}

func (m *TVTHeader) PopupMenu() IPopupMenu {
	r1 := LCL().SysCallN(5923, 0, m.Instance(), 0)
	return AsPopupMenu(r1)
}

func (m *TVTHeader) SetPopupMenu(AValue IPopupMenu) {
	LCL().SysCallN(5923, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TVTHeader) SortColumn() TColumnIndex {
	r1 := LCL().SysCallN(5927, 0, m.Instance(), 0)
	return TColumnIndex(r1)
}

func (m *TVTHeader) SetSortColumn(AValue TColumnIndex) {
	LCL().SysCallN(5927, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) SortDirection() TSortDirection {
	r1 := LCL().SysCallN(5928, 0, m.Instance(), 0)
	return TSortDirection(r1)
}

func (m *TVTHeader) SetSortDirection(AValue TSortDirection) {
	LCL().SysCallN(5928, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) Style() TVTHeaderStyle {
	r1 := LCL().SysCallN(5930, 0, m.Instance(), 0)
	return TVTHeaderStyle(r1)
}

func (m *TVTHeader) SetStyle(AValue TVTHeaderStyle) {
	LCL().SysCallN(5930, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTHeader) AllowFocus(ColumnIndex TColumnIndex) bool {
	r1 := LCL().SysCallN(5899, m.Instance(), uintptr(ColumnIndex))
	return GoBool(r1)
}

func (m *TVTHeader) InHeader(P *TPoint) bool {
	r1 := LCL().SysCallN(5914, m.Instance(), uintptr(unsafePointer(P)))
	return GoBool(r1)
}

func (m *TVTHeader) InHeaderSplitterArea(P *TPoint) bool {
	r1 := LCL().SysCallN(5915, m.Instance(), uintptr(unsafePointer(P)))
	return GoBool(r1)
}

func (m *TVTHeader) ResizeColumns(ChangeBy int32, RangeStartCol TColumnIndex, RangeEndCol TColumnIndex, Options TVTColumnOptions) int32 {
	r1 := LCL().SysCallN(5924, m.Instance(), uintptr(ChangeBy), uintptr(RangeStartCol), uintptr(RangeEndCol), uintptr(Options))
	return int32(r1)
}

func VTHeaderClass() TClass {
	ret := LCL().SysCallN(5903)
	return TClass(ret)
}

func (m *TVTHeader) AutoFitColumns(Animated bool, SmartAutoFitType TSmartAutoFitType, RangeStartCol int32, RangeEndCol int32) {
	LCL().SysCallN(5900, m.Instance(), PascalBool(Animated), uintptr(SmartAutoFitType), uintptr(RangeStartCol), uintptr(RangeEndCol))
}

func (m *TVTHeader) FixDesignFontsPPI(ADesignTimePPI int32) {
	LCL().SysCallN(5908, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TVTHeader) Invalidate(Column IVirtualTreeColumn, ExpandToBorder bool) {
	LCL().SysCallN(5916, m.Instance(), GetObjectUintptr(Column), PascalBool(ExpandToBorder))
}

func (m *TVTHeader) LoadFromStream(Stream IStream) {
	LCL().SysCallN(5917, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TVTHeader) RestoreColumns() {
	LCL().SysCallN(5925, m.Instance())
}

func (m *TVTHeader) SaveToStream(Stream IStream) {
	LCL().SysCallN(5926, m.Instance(), GetObjectUintptr(Stream))
}
