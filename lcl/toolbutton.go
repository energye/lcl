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
	r1 := oolButtonImportAPI().SysCallN(5, GetObjectUintptr(TheOwner))
	return AsToolButton(r1)
}

func (m *TToolButton) Index() int32 {
	r1 := oolButtonImportAPI().SysCallN(15, m.Instance())
	return int32(r1)
}

func (m *TToolButton) AllowAllUp() bool {
	r1 := oolButtonImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetAllowAllUp(AValue bool) {
	oolButtonImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) Down() bool {
	r1 := oolButtonImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetDown(AValue bool) {
	oolButtonImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) DragCursor() TCursor {
	r1 := oolButtonImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TToolButton) SetDragCursor(AValue TCursor) {
	oolButtonImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) DragKind() TDragKind {
	r1 := oolButtonImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TToolButton) SetDragKind(AValue TDragKind) {
	oolButtonImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) DragMode() TDragMode {
	r1 := oolButtonImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TToolButton) SetDragMode(AValue TDragMode) {
	oolButtonImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) DropdownMenu() IPopupMenu {
	r1 := oolButtonImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsPopupMenu(r1)
}

func (m *TToolButton) SetDropdownMenu(AValue IPopupMenu) {
	oolButtonImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolButton) Grouped() bool {
	r1 := oolButtonImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetGrouped(AValue bool) {
	oolButtonImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) ImageIndex() TImageIndex {
	r1 := oolButtonImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TToolButton) SetImageIndex(AValue TImageIndex) {
	oolButtonImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) Indeterminate() bool {
	r1 := oolButtonImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetIndeterminate(AValue bool) {
	oolButtonImportAPI().SysCallN(14, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) Marked() bool {
	r1 := oolButtonImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetMarked(AValue bool) {
	oolButtonImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) MenuItem() IMenuItem {
	r1 := oolButtonImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return AsMenuItem(r1)
}

func (m *TToolButton) SetMenuItem(AValue IMenuItem) {
	oolButtonImportAPI().SysCallN(17, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolButton) ParentShowHint() bool {
	r1 := oolButtonImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetParentShowHint(AValue bool) {
	oolButtonImportAPI().SysCallN(18, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) ShowCaption() bool {
	r1 := oolButtonImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetShowCaption(AValue bool) {
	oolButtonImportAPI().SysCallN(36, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) Style() TToolButtonStyle {
	r1 := oolButtonImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return TToolButtonStyle(r1)
}

func (m *TToolButton) SetStyle(AValue TToolButtonStyle) {
	oolButtonImportAPI().SysCallN(37, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolButton) Wrap() bool {
	r1 := oolButtonImportAPI().SysCallN(38, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolButton) SetWrap(AValue bool) {
	oolButtonImportAPI().SysCallN(38, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolButton) CheckMenuDropdown() bool {
	r1 := oolButtonImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TToolButton) PointInArrow(X, Y int32) bool {
	r1 := oolButtonImportAPI().SysCallN(19, m.Instance(), uintptr(X), uintptr(Y))
	return GoBool(r1)
}

func ToolButtonClass() TClass {
	ret := oolButtonImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TToolButton) Click() {
	oolButtonImportAPI().SysCallN(4, m.Instance())
}

func (m *TToolButton) ArrowClick() {
	oolButtonImportAPI().SysCallN(1, m.Instance())
}

func (m *TToolButton) GetCurrentIcon(ImageList *ICustomImageList, TheIndex *int32, TheEffect *TGraphicsDrawEffect) {
	var result0 uintptr
	var result1 uintptr
	var result2 uintptr
	oolButtonImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*ImageList = AsCustomImageList(result0)
	*TheIndex = int32(result1)
	*TheEffect = TGraphicsDrawEffect(result2)
}

func (m *TToolButton) SetOnArrowClick(fn TNotifyEvent) {
	if m.arrowClickPtr != 0 {
		RemoveEventElement(m.arrowClickPtr)
	}
	m.arrowClickPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(20, m.Instance(), m.arrowClickPtr)
}

func (m *TToolButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(21, m.Instance(), m.contextPopupPtr)
}

func (m *TToolButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(22, m.Instance(), m.dragDropPtr)
}

func (m *TToolButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(23, m.Instance(), m.dragOverPtr)
}

func (m *TToolButton) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(24, m.Instance(), m.endDockPtr)
}

func (m *TToolButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(25, m.Instance(), m.endDragPtr)
}

func (m *TToolButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(26, m.Instance(), m.mouseDownPtr)
}

func (m *TToolButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(27, m.Instance(), m.mouseEnterPtr)
}

func (m *TToolButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(28, m.Instance(), m.mouseLeavePtr)
}

func (m *TToolButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(29, m.Instance(), m.mouseMovePtr)
}

func (m *TToolButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(30, m.Instance(), m.mouseUpPtr)
}

func (m *TToolButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(31, m.Instance(), m.mouseWheelPtr)
}

func (m *TToolButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(32, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TToolButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(33, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TToolButton) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(34, m.Instance(), m.startDockPtr)
}

func (m *TToolButton) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	oolButtonImportAPI().SysCallN(35, m.Instance(), m.startDragPtr)
}

var (
	oolButtonImport       *imports.Imports = nil
	oolButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ToolButton_AllowAllUp", 0),
		/*1*/ imports.NewTable("ToolButton_ArrowClick", 0),
		/*2*/ imports.NewTable("ToolButton_CheckMenuDropdown", 0),
		/*3*/ imports.NewTable("ToolButton_Class", 0),
		/*4*/ imports.NewTable("ToolButton_Click", 0),
		/*5*/ imports.NewTable("ToolButton_Create", 0),
		/*6*/ imports.NewTable("ToolButton_Down", 0),
		/*7*/ imports.NewTable("ToolButton_DragCursor", 0),
		/*8*/ imports.NewTable("ToolButton_DragKind", 0),
		/*9*/ imports.NewTable("ToolButton_DragMode", 0),
		/*10*/ imports.NewTable("ToolButton_DropdownMenu", 0),
		/*11*/ imports.NewTable("ToolButton_GetCurrentIcon", 0),
		/*12*/ imports.NewTable("ToolButton_Grouped", 0),
		/*13*/ imports.NewTable("ToolButton_ImageIndex", 0),
		/*14*/ imports.NewTable("ToolButton_Indeterminate", 0),
		/*15*/ imports.NewTable("ToolButton_Index", 0),
		/*16*/ imports.NewTable("ToolButton_Marked", 0),
		/*17*/ imports.NewTable("ToolButton_MenuItem", 0),
		/*18*/ imports.NewTable("ToolButton_ParentShowHint", 0),
		/*19*/ imports.NewTable("ToolButton_PointInArrow", 0),
		/*20*/ imports.NewTable("ToolButton_SetOnArrowClick", 0),
		/*21*/ imports.NewTable("ToolButton_SetOnContextPopup", 0),
		/*22*/ imports.NewTable("ToolButton_SetOnDragDrop", 0),
		/*23*/ imports.NewTable("ToolButton_SetOnDragOver", 0),
		/*24*/ imports.NewTable("ToolButton_SetOnEndDock", 0),
		/*25*/ imports.NewTable("ToolButton_SetOnEndDrag", 0),
		/*26*/ imports.NewTable("ToolButton_SetOnMouseDown", 0),
		/*27*/ imports.NewTable("ToolButton_SetOnMouseEnter", 0),
		/*28*/ imports.NewTable("ToolButton_SetOnMouseLeave", 0),
		/*29*/ imports.NewTable("ToolButton_SetOnMouseMove", 0),
		/*30*/ imports.NewTable("ToolButton_SetOnMouseUp", 0),
		/*31*/ imports.NewTable("ToolButton_SetOnMouseWheel", 0),
		/*32*/ imports.NewTable("ToolButton_SetOnMouseWheelDown", 0),
		/*33*/ imports.NewTable("ToolButton_SetOnMouseWheelUp", 0),
		/*34*/ imports.NewTable("ToolButton_SetOnStartDock", 0),
		/*35*/ imports.NewTable("ToolButton_SetOnStartDrag", 0),
		/*36*/ imports.NewTable("ToolButton_ShowCaption", 0),
		/*37*/ imports.NewTable("ToolButton_Style", 0),
		/*38*/ imports.NewTable("ToolButton_Wrap", 0),
	}
)

func oolButtonImportAPI() *imports.Imports {
	if oolButtonImport == nil {
		oolButtonImport = NewDefaultImports()
		oolButtonImport.SetImportTable(oolButtonImportTables)
		oolButtonImportTables = nil
	}
	return oolButtonImport
}
