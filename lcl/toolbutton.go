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

// IToolButton Parent: IGraphicControl
type IToolButton interface {
	IGraphicControl
	Index() int32                                                                                // property
	AllowAllUp() bool                                                                            // property
	SetAllowAllUp(AValue bool)                                                                   // property
	Down() bool                                                                                  // property
	SetDown(AValue bool)                                                                         // property
	DragCursor() TCursor                                                                         // property
	SetDragCursor(AValue TCursor)                                                                // property
	DragKind() TDragKind                                                                         // property
	SetDragKind(AValue TDragKind)                                                                // property
	DragMode() TDragMode                                                                         // property
	SetDragMode(AValue TDragMode)                                                                // property
	DropdownMenu() IPopupMenu                                                                    // property
	SetDropdownMenu(AValue IPopupMenu)                                                           // property
	Grouped() bool                                                                               // property
	SetGrouped(AValue bool)                                                                      // property
	ImageIndex() TImageIndex                                                                     // property
	SetImageIndex(AValue TImageIndex)                                                            // property
	Indeterminate() bool                                                                         // property
	SetIndeterminate(AValue bool)                                                                // property
	Marked() bool                                                                                // property
	SetMarked(AValue bool)                                                                       // property
	MenuItem() IMenuItem                                                                         // property
	SetMenuItem(AValue IMenuItem)                                                                // property
	ParentShowHint() bool                                                                        // property
	SetParentShowHint(AValue bool)                                                               // property
	ShowCaption() bool                                                                           // property
	SetShowCaption(AValue bool)                                                                  // property
	Style() TToolButtonStyle                                                                     // property
	SetStyle(AValue TToolButtonStyle)                                                            // property
	Wrap() bool                                                                                  // property
	SetWrap(AValue bool)                                                                         // property
	CheckMenuDropdown() bool                                                                     // function
	PointInArrow(X, Y int32) bool                                                                // function
	Click()                                                                                      // procedure
	ArrowClick()                                                                                 // procedure
	GetCurrentIcon(ImageList *ICustomImageList, TheIndex *int32, TheEffect *TGraphicsDrawEffect) // procedure
	SetOnArrowClick(fn TNotifyEvent)                                                             // property event
	SetOnContextPopup(fn TContextPopupEvent)                                                     // property event
	SetOnDragDrop(fn TDragDropEvent)                                                             // property event
	SetOnDragOver(fn TDragOverEvent)                                                             // property event
	SetOnEndDock(fn TEndDragEvent)                                                               // property event
	SetOnEndDrag(fn TEndDragEvent)                                                               // property event
	SetOnMouseDown(fn TMouseEvent)                                                               // property event
	SetOnMouseEnter(fn TNotifyEvent)                                                             // property event
	SetOnMouseLeave(fn TNotifyEvent)                                                             // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                                           // property event
	SetOnMouseUp(fn TMouseEvent)                                                                 // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                                         // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                                               // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                                                 // property event
	SetOnStartDock(fn TStartDockEvent)                                                           // property event
	SetOnStartDrag(fn TStartDragEvent)                                                           // property event
}

// TToolButton Parent: TGraphicControl
type TToolButton struct {
	TGraphicControl
	arrowClickPtr     uintptr
	contextPopupPtr   uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDockPtr        uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewToolButton(TheOwner IComponent) IToolButton {
	r1 := LCL().SysCallN(5549, GetObjectUintptr(TheOwner))
	return AsToolButton(r1)
}

func (m *TToolButton) Index() int32 {
	r1 := LCL().SysCallN(5559, m.Instance())
	return int32(r1)
}

func (m *TToolButton) AllowAllUp() bool {
	r1 := LCL().SysCallN(5544, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetAllowAllUp(AValue bool) {
	LCL().SysCallN(5544, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) Down() bool {
	r1 := LCL().SysCallN(5550, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetDown(AValue bool) {
	LCL().SysCallN(5550, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) DragCursor() TCursor {
	r1 := LCL().SysCallN(5551, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TToolButton) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5551, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) DragKind() TDragKind {
	r1 := LCL().SysCallN(5552, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TToolButton) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5552, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) DragMode() TDragMode {
	r1 := LCL().SysCallN(5553, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TToolButton) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5553, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) DropdownMenu() IPopupMenu {
	r1 := LCL().SysCallN(5554, 0, m.Instance(), 0)
	return AsPopupMenu(r1)
}

func (m *TToolButton) SetDropdownMenu(AValue IPopupMenu) {
	LCL().SysCallN(5554, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolButton) Grouped() bool {
	r1 := LCL().SysCallN(5556, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetGrouped(AValue bool) {
	LCL().SysCallN(5556, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(5557, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TToolButton) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(5557, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) Indeterminate() bool {
	r1 := LCL().SysCallN(5558, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetIndeterminate(AValue bool) {
	LCL().SysCallN(5558, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) Marked() bool {
	r1 := LCL().SysCallN(5560, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetMarked(AValue bool) {
	LCL().SysCallN(5560, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) MenuItem() IMenuItem {
	r1 := LCL().SysCallN(5561, 0, m.Instance(), 0)
	return AsMenuItem(r1)
}

func (m *TToolButton) SetMenuItem(AValue IMenuItem) {
	LCL().SysCallN(5561, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(5562, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5562, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) ShowCaption() bool {
	r1 := LCL().SysCallN(5580, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetShowCaption(AValue bool) {
	LCL().SysCallN(5580, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) Style() TToolButtonStyle {
	r1 := LCL().SysCallN(5581, 0, m.Instance(), 0)
	return TToolButtonStyle(r1)
}

func (m *TToolButton) SetStyle(AValue TToolButtonStyle) {
	LCL().SysCallN(5581, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) Wrap() bool {
	r1 := LCL().SysCallN(5582, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetWrap(AValue bool) {
	LCL().SysCallN(5582, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) CheckMenuDropdown() bool {
	r1 := LCL().SysCallN(5546, m.Instance())
	return GoBool(r1)
}

func (m *TToolButton) PointInArrow(X, Y int32) bool {
	r1 := LCL().SysCallN(5563, m.Instance(), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func ToolButtonClass() TClass {
	ret := LCL().SysCallN(5547)
	return TClass(ret)
}

func (m *TToolButton) Click() {
	LCL().SysCallN(5548, m.Instance())
}

func (m *TToolButton) ArrowClick() {
	LCL().SysCallN(5545, m.Instance())
}

func (m *TToolButton) GetCurrentIcon(ImageList *ICustomImageList, TheIndex *int32, TheEffect *TGraphicsDrawEffect) {
	var result0 uintptr
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(5555, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*ImageList = AsCustomImageList(result0)
	*TheIndex = int32(result1)
	*TheEffect = TGraphicsDrawEffect(result2)
}

func (m *TToolButton) SetOnArrowClick(fn TNotifyEvent) {
	if m.arrowClickPtr != 0 {
		RemoveEventElement(m.arrowClickPtr)
	}
	m.arrowClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5564, m.Instance(), m.arrowClickPtr)
}

func (m *TToolButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5565, m.Instance(), m.contextPopupPtr)
}

func (m *TToolButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5566, m.Instance(), m.dragDropPtr)
}

func (m *TToolButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5567, m.Instance(), m.dragOverPtr)
}

func (m *TToolButton) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5568, m.Instance(), m.endDockPtr)
}

func (m *TToolButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5569, m.Instance(), m.endDragPtr)
}

func (m *TToolButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5570, m.Instance(), m.mouseDownPtr)
}

func (m *TToolButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5571, m.Instance(), m.mouseEnterPtr)
}

func (m *TToolButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5572, m.Instance(), m.mouseLeavePtr)
}

func (m *TToolButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5573, m.Instance(), m.mouseMovePtr)
}

func (m *TToolButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5574, m.Instance(), m.mouseUpPtr)
}

func (m *TToolButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5575, m.Instance(), m.mouseWheelPtr)
}

func (m *TToolButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5576, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TToolButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5577, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TToolButton) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5578, m.Instance(), m.startDockPtr)
}

func (m *TToolButton) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5579, m.Instance(), m.startDragPtr)
}
