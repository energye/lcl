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

// ICheckGroup Parent: ICustomCheckGroup
type ICheckGroup interface {
	ICustomCheckGroup
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TCheckGroup Parent: TCustomCheckGroup
type TCheckGroup struct {
	TCustomCheckGroup
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
}

func NewCheckGroup(TheOwner IComponent) ICheckGroup {
	r1 := checkGroupImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsCheckGroup(r1)
}

func (m *TCheckGroup) DragCursor() TCursor {
	r1 := checkGroupImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCheckGroup) SetDragCursor(AValue TCursor) {
	checkGroupImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckGroup) DragMode() TDragMode {
	r1 := checkGroupImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCheckGroup) SetDragMode(AValue TDragMode) {
	checkGroupImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckGroup) ParentFont() bool {
	r1 := checkGroupImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckGroup) SetParentFont(AValue bool) {
	checkGroupImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckGroup) ParentColor() bool {
	r1 := checkGroupImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckGroup) SetParentColor(AValue bool) {
	checkGroupImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckGroup) ParentShowHint() bool {
	r1 := checkGroupImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckGroup) SetParentShowHint(AValue bool) {
	checkGroupImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func CheckGroupClass() TClass {
	ret := checkGroupImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCheckGroup) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(7, m.Instance(), m.dblClickPtr)
}

func (m *TCheckGroup) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(8, m.Instance(), m.dragDropPtr)
}

func (m *TCheckGroup) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(9, m.Instance(), m.dragOverPtr)
}

func (m *TCheckGroup) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(10, m.Instance(), m.endDragPtr)
}

func (m *TCheckGroup) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(11, m.Instance(), m.mouseDownPtr)
}

func (m *TCheckGroup) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(12, m.Instance(), m.mouseEnterPtr)
}

func (m *TCheckGroup) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(13, m.Instance(), m.mouseLeavePtr)
}

func (m *TCheckGroup) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(14, m.Instance(), m.mouseMovePtr)
}

func (m *TCheckGroup) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(15, m.Instance(), m.mouseUpPtr)
}

func (m *TCheckGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(16, m.Instance(), m.mouseWheelPtr)
}

func (m *TCheckGroup) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(17, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCheckGroup) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(18, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCheckGroup) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	checkGroupImportAPI().SysCallN(19, m.Instance(), m.startDragPtr)
}

var (
	checkGroupImport       *imports.Imports = nil
	checkGroupImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CheckGroup_Class", 0),
		/*1*/ imports.NewTable("CheckGroup_Create", 0),
		/*2*/ imports.NewTable("CheckGroup_DragCursor", 0),
		/*3*/ imports.NewTable("CheckGroup_DragMode", 0),
		/*4*/ imports.NewTable("CheckGroup_ParentColor", 0),
		/*5*/ imports.NewTable("CheckGroup_ParentFont", 0),
		/*6*/ imports.NewTable("CheckGroup_ParentShowHint", 0),
		/*7*/ imports.NewTable("CheckGroup_SetOnDblClick", 0),
		/*8*/ imports.NewTable("CheckGroup_SetOnDragDrop", 0),
		/*9*/ imports.NewTable("CheckGroup_SetOnDragOver", 0),
		/*10*/ imports.NewTable("CheckGroup_SetOnEndDrag", 0),
		/*11*/ imports.NewTable("CheckGroup_SetOnMouseDown", 0),
		/*12*/ imports.NewTable("CheckGroup_SetOnMouseEnter", 0),
		/*13*/ imports.NewTable("CheckGroup_SetOnMouseLeave", 0),
		/*14*/ imports.NewTable("CheckGroup_SetOnMouseMove", 0),
		/*15*/ imports.NewTable("CheckGroup_SetOnMouseUp", 0),
		/*16*/ imports.NewTable("CheckGroup_SetOnMouseWheel", 0),
		/*17*/ imports.NewTable("CheckGroup_SetOnMouseWheelDown", 0),
		/*18*/ imports.NewTable("CheckGroup_SetOnMouseWheelUp", 0),
		/*19*/ imports.NewTable("CheckGroup_SetOnStartDrag", 0),
	}
)

func checkGroupImportAPI() *imports.Imports {
	if checkGroupImport == nil {
		checkGroupImport = NewDefaultImports()
		checkGroupImport.SetImportTable(checkGroupImportTables)
		checkGroupImportTables = nil
	}
	return checkGroupImport
}
