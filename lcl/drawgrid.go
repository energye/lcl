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

// IDrawGrid Parent: ICustomDrawGrid
type IDrawGrid interface {
	ICustomDrawGrid
	InplaceEditor() IWinControl                          // property
	AlternateColor() TColor                              // property
	SetAlternateColor(AValue TColor)                     // property
	AutoEdit() bool                                      // property
	SetAutoEdit(AValue bool)                             // property
	ColRowDraggingCursor() TCursor                       // property
	SetColRowDraggingCursor(AValue TCursor)              // property
	ColRowDragIndicatorColor() TColor                    // property
	SetColRowDragIndicatorColor(AValue TColor)           // property
	ColSizingCursor() TCursor                            // property
	SetColSizingCursor(AValue TCursor)                   // property
	ColumnClickSorts() bool                              // property
	SetColumnClickSorts(AValue bool)                     // property
	DragCursor() TCursor                                 // property
	SetDragCursor(AValue TCursor)                        // property
	DragKind() TDragKind                                 // property
	SetDragKind(AValue TDragKind)                        // property
	DragMode() TDragMode                                 // property
	SetDragMode(AValue TDragMode)                        // property
	ExtendedSelect() bool                                // property
	SetExtendedSelect(AValue bool)                       // property
	HeaderHotZones() TGridZoneSet                        // property
	SetHeaderHotZones(AValue TGridZoneSet)               // property
	HeaderPushZones() TGridZoneSet                       // property
	SetHeaderPushZones(AValue TGridZoneSet)              // property
	ImageIndexSortAsc() TImageIndex                      // property
	SetImageIndexSortAsc(AValue TImageIndex)             // property
	ImageIndexSortDesc() TImageIndex                     // property
	SetImageIndexSortDesc(AValue TImageIndex)            // property
	MouseWheelOption() TMouseWheelOption                 // property
	SetMouseWheelOption(AValue TMouseWheelOption)        // property
	ParentColor() bool                                   // property
	SetParentColor(AValue bool)                          // property
	ParentFont() bool                                    // property
	SetParentFont(AValue bool)                           // property
	RangeSelectMode() TRangeSelectMode                   // property
	SetRangeSelectMode(AValue TRangeSelectMode)          // property
	RowSizingCursor() TCursor                            // property
	SetRowSizingCursor(AValue TCursor)                   // property
	TitleFont() IFont                                    // property
	SetTitleFont(AValue IFont)                           // property
	TitleImageList() IImageList                          // property
	SetTitleImageList(AValue IImageList)                 // property
	TitleImageListWidth() int32                          // property
	SetTitleImageListWidth(AValue int32)                 // property
	TitleStyle() TTitleStyle                             // property
	SetTitleStyle(AValue TTitleStyle)                    // property
	SetOnCheckboxToggled(fn TToggledCheckboxEvent)       // property event
	SetOnEditingDone(fn TNotifyEvent)                    // property event
	SetOnGetCellHint(fn TGetCellHintEvent)               // property event
	SetOnGetCheckboxState(fn TGetCheckboxStateEvent)     // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)             // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)      // property event
	SetOnSetCheckboxState(fn TSetCheckboxStateEvent)     // property event
	SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) // property event
	SetOnUserCheckboxImage(fn TUserCheckBoxImageEvent)   // property event
}

// TDrawGrid Parent: TCustomDrawGrid
type TDrawGrid struct {
	TCustomDrawGrid
	checkboxToggledPtr    uintptr
	editingDonePtr        uintptr
	getCellHintPtr        uintptr
	getCheckboxStatePtr   uintptr
	mouseWheelHorzPtr     uintptr
	mouseWheelLeftPtr     uintptr
	mouseWheelRightPtr    uintptr
	setCheckboxStatePtr   uintptr
	userCheckboxBitmapPtr uintptr
	userCheckboxImagePtr  uintptr
}

func NewDrawGrid(AOwner IComponent) IDrawGrid {
	r1 := drawGridImportAPI().SysCallN(7, GetObjectUintptr(AOwner))
	return AsDrawGrid(r1)
}

func (m *TDrawGrid) InplaceEditor() IWinControl {
	r1 := drawGridImportAPI().SysCallN(16, m.Instance())
	return AsWinControl(r1)
}

func (m *TDrawGrid) AlternateColor() TColor {
	r1 := drawGridImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TDrawGrid) SetAlternateColor(AValue TColor) {
	drawGridImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) AutoEdit() bool {
	r1 := drawGridImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetAutoEdit(AValue bool) {
	drawGridImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) ColRowDraggingCursor() TCursor {
	r1 := drawGridImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDrawGrid) SetColRowDraggingCursor(AValue TCursor) {
	drawGridImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ColRowDragIndicatorColor() TColor {
	r1 := drawGridImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TDrawGrid) SetColRowDragIndicatorColor(AValue TColor) {
	drawGridImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ColSizingCursor() TCursor {
	r1 := drawGridImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDrawGrid) SetColSizingCursor(AValue TCursor) {
	drawGridImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ColumnClickSorts() bool {
	r1 := drawGridImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetColumnClickSorts(AValue bool) {
	drawGridImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) DragCursor() TCursor {
	r1 := drawGridImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDrawGrid) SetDragCursor(AValue TCursor) {
	drawGridImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) DragKind() TDragKind {
	r1 := drawGridImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TDrawGrid) SetDragKind(AValue TDragKind) {
	drawGridImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) DragMode() TDragMode {
	r1 := drawGridImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TDrawGrid) SetDragMode(AValue TDragMode) {
	drawGridImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ExtendedSelect() bool {
	r1 := drawGridImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetExtendedSelect(AValue bool) {
	drawGridImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) HeaderHotZones() TGridZoneSet {
	r1 := drawGridImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TDrawGrid) SetHeaderHotZones(AValue TGridZoneSet) {
	drawGridImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) HeaderPushZones() TGridZoneSet {
	r1 := drawGridImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TDrawGrid) SetHeaderPushZones(AValue TGridZoneSet) {
	drawGridImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ImageIndexSortAsc() TImageIndex {
	r1 := drawGridImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TDrawGrid) SetImageIndexSortAsc(AValue TImageIndex) {
	drawGridImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ImageIndexSortDesc() TImageIndex {
	r1 := drawGridImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TDrawGrid) SetImageIndexSortDesc(AValue TImageIndex) {
	drawGridImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) MouseWheelOption() TMouseWheelOption {
	r1 := drawGridImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return TMouseWheelOption(r1)
}

func (m *TDrawGrid) SetMouseWheelOption(AValue TMouseWheelOption) {
	drawGridImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ParentColor() bool {
	r1 := drawGridImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetParentColor(AValue bool) {
	drawGridImportAPI().SysCallN(18, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) ParentFont() bool {
	r1 := drawGridImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetParentFont(AValue bool) {
	drawGridImportAPI().SysCallN(19, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) RangeSelectMode() TRangeSelectMode {
	r1 := drawGridImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return TRangeSelectMode(r1)
}

func (m *TDrawGrid) SetRangeSelectMode(AValue TRangeSelectMode) {
	drawGridImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) RowSizingCursor() TCursor {
	r1 := drawGridImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDrawGrid) SetRowSizingCursor(AValue TCursor) {
	drawGridImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) TitleFont() IFont {
	r1 := drawGridImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TDrawGrid) SetTitleFont(AValue IFont) {
	drawGridImportAPI().SysCallN(32, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDrawGrid) TitleImageList() IImageList {
	r1 := drawGridImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return AsImageList(r1)
}

func (m *TDrawGrid) SetTitleImageList(AValue IImageList) {
	drawGridImportAPI().SysCallN(33, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDrawGrid) TitleImageListWidth() int32 {
	r1 := drawGridImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDrawGrid) SetTitleImageListWidth(AValue int32) {
	drawGridImportAPI().SysCallN(34, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) TitleStyle() TTitleStyle {
	r1 := drawGridImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return TTitleStyle(r1)
}

func (m *TDrawGrid) SetTitleStyle(AValue TTitleStyle) {
	drawGridImportAPI().SysCallN(35, 1, m.Instance(), uintptr(AValue))
}

func DrawGridClass() TClass {
	ret := drawGridImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TDrawGrid) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if m.checkboxToggledPtr != 0 {
		RemoveEventElement(m.checkboxToggledPtr)
	}
	m.checkboxToggledPtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(22, m.Instance(), m.checkboxToggledPtr)
}

func (m *TDrawGrid) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(23, m.Instance(), m.editingDonePtr)
}

func (m *TDrawGrid) SetOnGetCellHint(fn TGetCellHintEvent) {
	if m.getCellHintPtr != 0 {
		RemoveEventElement(m.getCellHintPtr)
	}
	m.getCellHintPtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(24, m.Instance(), m.getCellHintPtr)
}

func (m *TDrawGrid) SetOnGetCheckboxState(fn TGetCheckboxStateEvent) {
	if m.getCheckboxStatePtr != 0 {
		RemoveEventElement(m.getCheckboxStatePtr)
	}
	m.getCheckboxStatePtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(25, m.Instance(), m.getCheckboxStatePtr)
}

func (m *TDrawGrid) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(26, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TDrawGrid) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(27, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TDrawGrid) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(28, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TDrawGrid) SetOnSetCheckboxState(fn TSetCheckboxStateEvent) {
	if m.setCheckboxStatePtr != 0 {
		RemoveEventElement(m.setCheckboxStatePtr)
	}
	m.setCheckboxStatePtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(29, m.Instance(), m.setCheckboxStatePtr)
}

func (m *TDrawGrid) SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) {
	if m.userCheckboxBitmapPtr != 0 {
		RemoveEventElement(m.userCheckboxBitmapPtr)
	}
	m.userCheckboxBitmapPtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(30, m.Instance(), m.userCheckboxBitmapPtr)
}

func (m *TDrawGrid) SetOnUserCheckboxImage(fn TUserCheckBoxImageEvent) {
	if m.userCheckboxImagePtr != 0 {
		RemoveEventElement(m.userCheckboxImagePtr)
	}
	m.userCheckboxImagePtr = MakeEventDataPtr(fn)
	drawGridImportAPI().SysCallN(31, m.Instance(), m.userCheckboxImagePtr)
}

var (
	drawGridImport       *imports.Imports = nil
	drawGridImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DrawGrid_AlternateColor", 0),
		/*1*/ imports.NewTable("DrawGrid_AutoEdit", 0),
		/*2*/ imports.NewTable("DrawGrid_Class", 0),
		/*3*/ imports.NewTable("DrawGrid_ColRowDragIndicatorColor", 0),
		/*4*/ imports.NewTable("DrawGrid_ColRowDraggingCursor", 0),
		/*5*/ imports.NewTable("DrawGrid_ColSizingCursor", 0),
		/*6*/ imports.NewTable("DrawGrid_ColumnClickSorts", 0),
		/*7*/ imports.NewTable("DrawGrid_Create", 0),
		/*8*/ imports.NewTable("DrawGrid_DragCursor", 0),
		/*9*/ imports.NewTable("DrawGrid_DragKind", 0),
		/*10*/ imports.NewTable("DrawGrid_DragMode", 0),
		/*11*/ imports.NewTable("DrawGrid_ExtendedSelect", 0),
		/*12*/ imports.NewTable("DrawGrid_HeaderHotZones", 0),
		/*13*/ imports.NewTable("DrawGrid_HeaderPushZones", 0),
		/*14*/ imports.NewTable("DrawGrid_ImageIndexSortAsc", 0),
		/*15*/ imports.NewTable("DrawGrid_ImageIndexSortDesc", 0),
		/*16*/ imports.NewTable("DrawGrid_InplaceEditor", 0),
		/*17*/ imports.NewTable("DrawGrid_MouseWheelOption", 0),
		/*18*/ imports.NewTable("DrawGrid_ParentColor", 0),
		/*19*/ imports.NewTable("DrawGrid_ParentFont", 0),
		/*20*/ imports.NewTable("DrawGrid_RangeSelectMode", 0),
		/*21*/ imports.NewTable("DrawGrid_RowSizingCursor", 0),
		/*22*/ imports.NewTable("DrawGrid_SetOnCheckboxToggled", 0),
		/*23*/ imports.NewTable("DrawGrid_SetOnEditingDone", 0),
		/*24*/ imports.NewTable("DrawGrid_SetOnGetCellHint", 0),
		/*25*/ imports.NewTable("DrawGrid_SetOnGetCheckboxState", 0),
		/*26*/ imports.NewTable("DrawGrid_SetOnMouseWheelHorz", 0),
		/*27*/ imports.NewTable("DrawGrid_SetOnMouseWheelLeft", 0),
		/*28*/ imports.NewTable("DrawGrid_SetOnMouseWheelRight", 0),
		/*29*/ imports.NewTable("DrawGrid_SetOnSetCheckboxState", 0),
		/*30*/ imports.NewTable("DrawGrid_SetOnUserCheckboxBitmap", 0),
		/*31*/ imports.NewTable("DrawGrid_SetOnUserCheckboxImage", 0),
		/*32*/ imports.NewTable("DrawGrid_TitleFont", 0),
		/*33*/ imports.NewTable("DrawGrid_TitleImageList", 0),
		/*34*/ imports.NewTable("DrawGrid_TitleImageListWidth", 0),
		/*35*/ imports.NewTable("DrawGrid_TitleStyle", 0),
	}
)

func drawGridImportAPI() *imports.Imports {
	if drawGridImport == nil {
		drawGridImport = NewDefaultImports()
		drawGridImport.SetImportTable(drawGridImportTables)
		drawGridImportTables = nil
	}
	return drawGridImport
}
