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

// IStringGrid Parent: ICustomStringGrid
type IStringGrid interface {
	ICustomStringGrid
	Modified() bool                                      // property
	SetModified(AValue bool)                             // property
	InplaceEditor() IWinControl                          // property
	AlternateColor() TColor                              // property
	SetAlternateColor(AValue TColor)                     // property
	AutoEdit() bool                                      // property
	SetAutoEdit(AValue bool)                             // property
	CellHintPriority() TCellHintPriority                 // property
	SetCellHintPriority(AValue TCellHintPriority)        // property
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
	TitleStyle() TTitleStyle                             // property
	SetTitleStyle(AValue TTitleStyle)                    // property
	SetOnCellProcess(fn TCellProcessEvent)               // property event
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

// TStringGrid Parent: TCustomStringGrid
type TStringGrid struct {
	TCustomStringGrid
	cellProcessPtr        uintptr
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

func NewStringGrid(AOwner IComponent) IStringGrid {
	r1 := stringGridImportAPI().SysCallN(8, GetObjectUintptr(AOwner))
	return AsStringGrid(r1)
}

func (m *TStringGrid) Modified() bool {
	r1 := stringGridImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetModified(AValue bool) {
	stringGridImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) InplaceEditor() IWinControl {
	r1 := stringGridImportAPI().SysCallN(16, m.Instance())
	return AsWinControl(r1)
}

func (m *TStringGrid) AlternateColor() TColor {
	r1 := stringGridImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TStringGrid) SetAlternateColor(AValue TColor) {
	stringGridImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) AutoEdit() bool {
	r1 := stringGridImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetAutoEdit(AValue bool) {
	stringGridImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) CellHintPriority() TCellHintPriority {
	r1 := stringGridImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCellHintPriority(r1)
}

func (m *TStringGrid) SetCellHintPriority(AValue TCellHintPriority) {
	stringGridImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ColRowDraggingCursor() TCursor {
	r1 := stringGridImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStringGrid) SetColRowDraggingCursor(AValue TCursor) {
	stringGridImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ColRowDragIndicatorColor() TColor {
	r1 := stringGridImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TStringGrid) SetColRowDragIndicatorColor(AValue TColor) {
	stringGridImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ColSizingCursor() TCursor {
	r1 := stringGridImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStringGrid) SetColSizingCursor(AValue TCursor) {
	stringGridImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ColumnClickSorts() bool {
	r1 := stringGridImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetColumnClickSorts(AValue bool) {
	stringGridImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) DragCursor() TCursor {
	r1 := stringGridImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStringGrid) SetDragCursor(AValue TCursor) {
	stringGridImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) DragKind() TDragKind {
	r1 := stringGridImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TStringGrid) SetDragKind(AValue TDragKind) {
	stringGridImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) DragMode() TDragMode {
	r1 := stringGridImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TStringGrid) SetDragMode(AValue TDragMode) {
	stringGridImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) HeaderHotZones() TGridZoneSet {
	r1 := stringGridImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TStringGrid) SetHeaderHotZones(AValue TGridZoneSet) {
	stringGridImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) HeaderPushZones() TGridZoneSet {
	r1 := stringGridImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TStringGrid) SetHeaderPushZones(AValue TGridZoneSet) {
	stringGridImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ImageIndexSortAsc() TImageIndex {
	r1 := stringGridImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TStringGrid) SetImageIndexSortAsc(AValue TImageIndex) {
	stringGridImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ImageIndexSortDesc() TImageIndex {
	r1 := stringGridImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TStringGrid) SetImageIndexSortDesc(AValue TImageIndex) {
	stringGridImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) MouseWheelOption() TMouseWheelOption {
	r1 := stringGridImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TMouseWheelOption(r1)
}

func (m *TStringGrid) SetMouseWheelOption(AValue TMouseWheelOption) {
	stringGridImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ParentColor() bool {
	r1 := stringGridImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetParentColor(AValue bool) {
	stringGridImportAPI().SysCallN(19, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) ParentFont() bool {
	r1 := stringGridImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetParentFont(AValue bool) {
	stringGridImportAPI().SysCallN(20, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) RangeSelectMode() TRangeSelectMode {
	r1 := stringGridImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return TRangeSelectMode(r1)
}

func (m *TStringGrid) SetRangeSelectMode(AValue TRangeSelectMode) {
	stringGridImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) RowSizingCursor() TCursor {
	r1 := stringGridImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStringGrid) SetRowSizingCursor(AValue TCursor) {
	stringGridImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) TitleFont() IFont {
	r1 := stringGridImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TStringGrid) SetTitleFont(AValue IFont) {
	stringGridImportAPI().SysCallN(34, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TStringGrid) TitleImageList() IImageList {
	r1 := stringGridImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return AsImageList(r1)
}

func (m *TStringGrid) SetTitleImageList(AValue IImageList) {
	stringGridImportAPI().SysCallN(35, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TStringGrid) TitleStyle() TTitleStyle {
	r1 := stringGridImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return TTitleStyle(r1)
}

func (m *TStringGrid) SetTitleStyle(AValue TTitleStyle) {
	stringGridImportAPI().SysCallN(36, 1, m.Instance(), uintptr(AValue))
}

func StringGridClass() TClass {
	ret := stringGridImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TStringGrid) SetOnCellProcess(fn TCellProcessEvent) {
	if m.cellProcessPtr != 0 {
		RemoveEventElement(m.cellProcessPtr)
	}
	m.cellProcessPtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(23, m.Instance(), m.cellProcessPtr)
}

func (m *TStringGrid) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if m.checkboxToggledPtr != 0 {
		RemoveEventElement(m.checkboxToggledPtr)
	}
	m.checkboxToggledPtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(24, m.Instance(), m.checkboxToggledPtr)
}

func (m *TStringGrid) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(25, m.Instance(), m.editingDonePtr)
}

func (m *TStringGrid) SetOnGetCellHint(fn TGetCellHintEvent) {
	if m.getCellHintPtr != 0 {
		RemoveEventElement(m.getCellHintPtr)
	}
	m.getCellHintPtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(26, m.Instance(), m.getCellHintPtr)
}

func (m *TStringGrid) SetOnGetCheckboxState(fn TGetCheckboxStateEvent) {
	if m.getCheckboxStatePtr != 0 {
		RemoveEventElement(m.getCheckboxStatePtr)
	}
	m.getCheckboxStatePtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(27, m.Instance(), m.getCheckboxStatePtr)
}

func (m *TStringGrid) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(28, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TStringGrid) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(29, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TStringGrid) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(30, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TStringGrid) SetOnSetCheckboxState(fn TSetCheckboxStateEvent) {
	if m.setCheckboxStatePtr != 0 {
		RemoveEventElement(m.setCheckboxStatePtr)
	}
	m.setCheckboxStatePtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(31, m.Instance(), m.setCheckboxStatePtr)
}

func (m *TStringGrid) SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) {
	if m.userCheckboxBitmapPtr != 0 {
		RemoveEventElement(m.userCheckboxBitmapPtr)
	}
	m.userCheckboxBitmapPtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(32, m.Instance(), m.userCheckboxBitmapPtr)
}

func (m *TStringGrid) SetOnUserCheckboxImage(fn TUserCheckBoxImageEvent) {
	if m.userCheckboxImagePtr != 0 {
		RemoveEventElement(m.userCheckboxImagePtr)
	}
	m.userCheckboxImagePtr = MakeEventDataPtr(fn)
	stringGridImportAPI().SysCallN(33, m.Instance(), m.userCheckboxImagePtr)
}

var (
	stringGridImport       *imports.Imports = nil
	stringGridImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("StringGrid_AlternateColor", 0),
		/*1*/ imports.NewTable("StringGrid_AutoEdit", 0),
		/*2*/ imports.NewTable("StringGrid_CellHintPriority", 0),
		/*3*/ imports.NewTable("StringGrid_Class", 0),
		/*4*/ imports.NewTable("StringGrid_ColRowDragIndicatorColor", 0),
		/*5*/ imports.NewTable("StringGrid_ColRowDraggingCursor", 0),
		/*6*/ imports.NewTable("StringGrid_ColSizingCursor", 0),
		/*7*/ imports.NewTable("StringGrid_ColumnClickSorts", 0),
		/*8*/ imports.NewTable("StringGrid_Create", 0),
		/*9*/ imports.NewTable("StringGrid_DragCursor", 0),
		/*10*/ imports.NewTable("StringGrid_DragKind", 0),
		/*11*/ imports.NewTable("StringGrid_DragMode", 0),
		/*12*/ imports.NewTable("StringGrid_HeaderHotZones", 0),
		/*13*/ imports.NewTable("StringGrid_HeaderPushZones", 0),
		/*14*/ imports.NewTable("StringGrid_ImageIndexSortAsc", 0),
		/*15*/ imports.NewTable("StringGrid_ImageIndexSortDesc", 0),
		/*16*/ imports.NewTable("StringGrid_InplaceEditor", 0),
		/*17*/ imports.NewTable("StringGrid_Modified", 0),
		/*18*/ imports.NewTable("StringGrid_MouseWheelOption", 0),
		/*19*/ imports.NewTable("StringGrid_ParentColor", 0),
		/*20*/ imports.NewTable("StringGrid_ParentFont", 0),
		/*21*/ imports.NewTable("StringGrid_RangeSelectMode", 0),
		/*22*/ imports.NewTable("StringGrid_RowSizingCursor", 0),
		/*23*/ imports.NewTable("StringGrid_SetOnCellProcess", 0),
		/*24*/ imports.NewTable("StringGrid_SetOnCheckboxToggled", 0),
		/*25*/ imports.NewTable("StringGrid_SetOnEditingDone", 0),
		/*26*/ imports.NewTable("StringGrid_SetOnGetCellHint", 0),
		/*27*/ imports.NewTable("StringGrid_SetOnGetCheckboxState", 0),
		/*28*/ imports.NewTable("StringGrid_SetOnMouseWheelHorz", 0),
		/*29*/ imports.NewTable("StringGrid_SetOnMouseWheelLeft", 0),
		/*30*/ imports.NewTable("StringGrid_SetOnMouseWheelRight", 0),
		/*31*/ imports.NewTable("StringGrid_SetOnSetCheckboxState", 0),
		/*32*/ imports.NewTable("StringGrid_SetOnUserCheckboxBitmap", 0),
		/*33*/ imports.NewTable("StringGrid_SetOnUserCheckboxImage", 0),
		/*34*/ imports.NewTable("StringGrid_TitleFont", 0),
		/*35*/ imports.NewTable("StringGrid_TitleImageList", 0),
		/*36*/ imports.NewTable("StringGrid_TitleStyle", 0),
	}
)

func stringGridImportAPI() *imports.Imports {
	if stringGridImport == nil {
		stringGridImport = NewDefaultImports()
		stringGridImport.SetImportTable(stringGridImportTables)
		stringGridImportTables = nil
	}
	return stringGridImport
}
