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
	r1 := LCL().SysCallN(5238, GetObjectUintptr(AOwner))
	return AsStringGrid(r1)
}

func (m *TStringGrid) Modified() bool {
	r1 := LCL().SysCallN(5247, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetModified(AValue bool) {
	LCL().SysCallN(5247, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) InplaceEditor() IWinControl {
	r1 := LCL().SysCallN(5246, m.Instance())
	return AsWinControl(r1)
}

func (m *TStringGrid) AlternateColor() TColor {
	r1 := LCL().SysCallN(5230, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TStringGrid) SetAlternateColor(AValue TColor) {
	LCL().SysCallN(5230, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) AutoEdit() bool {
	r1 := LCL().SysCallN(5231, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetAutoEdit(AValue bool) {
	LCL().SysCallN(5231, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) CellHintPriority() TCellHintPriority {
	r1 := LCL().SysCallN(5232, 0, m.Instance(), 0)
	return TCellHintPriority(r1)
}

func (m *TStringGrid) SetCellHintPriority(AValue TCellHintPriority) {
	LCL().SysCallN(5232, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ColRowDraggingCursor() TCursor {
	r1 := LCL().SysCallN(5235, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStringGrid) SetColRowDraggingCursor(AValue TCursor) {
	LCL().SysCallN(5235, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ColRowDragIndicatorColor() TColor {
	r1 := LCL().SysCallN(5234, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TStringGrid) SetColRowDragIndicatorColor(AValue TColor) {
	LCL().SysCallN(5234, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ColSizingCursor() TCursor {
	r1 := LCL().SysCallN(5236, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStringGrid) SetColSizingCursor(AValue TCursor) {
	LCL().SysCallN(5236, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ColumnClickSorts() bool {
	r1 := LCL().SysCallN(5237, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetColumnClickSorts(AValue bool) {
	LCL().SysCallN(5237, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) DragCursor() TCursor {
	r1 := LCL().SysCallN(5239, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStringGrid) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5239, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) DragKind() TDragKind {
	r1 := LCL().SysCallN(5240, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TStringGrid) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5240, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) DragMode() TDragMode {
	r1 := LCL().SysCallN(5241, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TStringGrid) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5241, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) HeaderHotZones() TGridZoneSet {
	r1 := LCL().SysCallN(5242, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TStringGrid) SetHeaderHotZones(AValue TGridZoneSet) {
	LCL().SysCallN(5242, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) HeaderPushZones() TGridZoneSet {
	r1 := LCL().SysCallN(5243, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TStringGrid) SetHeaderPushZones(AValue TGridZoneSet) {
	LCL().SysCallN(5243, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ImageIndexSortAsc() TImageIndex {
	r1 := LCL().SysCallN(5244, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TStringGrid) SetImageIndexSortAsc(AValue TImageIndex) {
	LCL().SysCallN(5244, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ImageIndexSortDesc() TImageIndex {
	r1 := LCL().SysCallN(5245, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TStringGrid) SetImageIndexSortDesc(AValue TImageIndex) {
	LCL().SysCallN(5245, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) MouseWheelOption() TMouseWheelOption {
	r1 := LCL().SysCallN(5248, 0, m.Instance(), 0)
	return TMouseWheelOption(r1)
}

func (m *TStringGrid) SetMouseWheelOption(AValue TMouseWheelOption) {
	LCL().SysCallN(5248, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) ParentColor() bool {
	r1 := LCL().SysCallN(5249, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetParentColor(AValue bool) {
	LCL().SysCallN(5249, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) ParentFont() bool {
	r1 := LCL().SysCallN(5250, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringGrid) SetParentFont(AValue bool) {
	LCL().SysCallN(5250, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringGrid) RangeSelectMode() TRangeSelectMode {
	r1 := LCL().SysCallN(5251, 0, m.Instance(), 0)
	return TRangeSelectMode(r1)
}

func (m *TStringGrid) SetRangeSelectMode(AValue TRangeSelectMode) {
	LCL().SysCallN(5251, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) RowSizingCursor() TCursor {
	r1 := LCL().SysCallN(5252, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStringGrid) SetRowSizingCursor(AValue TCursor) {
	LCL().SysCallN(5252, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringGrid) TitleFont() IFont {
	r1 := LCL().SysCallN(5264, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TStringGrid) SetTitleFont(AValue IFont) {
	LCL().SysCallN(5264, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TStringGrid) TitleImageList() IImageList {
	r1 := LCL().SysCallN(5265, 0, m.Instance(), 0)
	return AsImageList(r1)
}

func (m *TStringGrid) SetTitleImageList(AValue IImageList) {
	LCL().SysCallN(5265, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TStringGrid) TitleStyle() TTitleStyle {
	r1 := LCL().SysCallN(5266, 0, m.Instance(), 0)
	return TTitleStyle(r1)
}

func (m *TStringGrid) SetTitleStyle(AValue TTitleStyle) {
	LCL().SysCallN(5266, 1, m.Instance(), uintptr(AValue))
}

func StringGridClass() TClass {
	ret := LCL().SysCallN(5233)
	return TClass(ret)
}

func (m *TStringGrid) SetOnCellProcess(fn TCellProcessEvent) {
	if m.cellProcessPtr != 0 {
		RemoveEventElement(m.cellProcessPtr)
	}
	m.cellProcessPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5253, m.Instance(), m.cellProcessPtr)
}

func (m *TStringGrid) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if m.checkboxToggledPtr != 0 {
		RemoveEventElement(m.checkboxToggledPtr)
	}
	m.checkboxToggledPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5254, m.Instance(), m.checkboxToggledPtr)
}

func (m *TStringGrid) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5255, m.Instance(), m.editingDonePtr)
}

func (m *TStringGrid) SetOnGetCellHint(fn TGetCellHintEvent) {
	if m.getCellHintPtr != 0 {
		RemoveEventElement(m.getCellHintPtr)
	}
	m.getCellHintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5256, m.Instance(), m.getCellHintPtr)
}

func (m *TStringGrid) SetOnGetCheckboxState(fn TGetCheckboxStateEvent) {
	if m.getCheckboxStatePtr != 0 {
		RemoveEventElement(m.getCheckboxStatePtr)
	}
	m.getCheckboxStatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5257, m.Instance(), m.getCheckboxStatePtr)
}

func (m *TStringGrid) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5258, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TStringGrid) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5259, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TStringGrid) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5260, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TStringGrid) SetOnSetCheckboxState(fn TSetCheckboxStateEvent) {
	if m.setCheckboxStatePtr != 0 {
		RemoveEventElement(m.setCheckboxStatePtr)
	}
	m.setCheckboxStatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5261, m.Instance(), m.setCheckboxStatePtr)
}

func (m *TStringGrid) SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) {
	if m.userCheckboxBitmapPtr != 0 {
		RemoveEventElement(m.userCheckboxBitmapPtr)
	}
	m.userCheckboxBitmapPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5262, m.Instance(), m.userCheckboxBitmapPtr)
}

func (m *TStringGrid) SetOnUserCheckboxImage(fn TUserCheckBoxImageEvent) {
	if m.userCheckboxImagePtr != 0 {
		RemoveEventElement(m.userCheckboxImagePtr)
	}
	m.userCheckboxImagePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5263, m.Instance(), m.userCheckboxImagePtr)
}
