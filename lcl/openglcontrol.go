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

// IOpenGLControl Parent: ICustomOpenGLControl
type IOpenGLControl interface {
	ICustomOpenGLControl
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
	SetOnDblClick(fn TNotifyEvent)                     // property event
	SetOnDragDrop(fn TDragDropEvent)                   // property event
	SetOnDragOver(fn TDragOverEvent)                   // property event
	SetOnMouseDown(fn TMouseEvent)                     // property event
	SetOnMouseEnter(fn TNotifyEvent)                   // property event
	SetOnMouseLeave(fn TNotifyEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                       // property event
	SetOnMouseWheel(fn TMouseWheelEvent)               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)       // property event
}

// TOpenGLControl Parent: TCustomOpenGLControl
type TOpenGLControl struct {
	TCustomOpenGLControl
	constrainedResizePtr uintptr
	dblClickPtr          uintptr
	dragDropPtr          uintptr
	dragOverPtr          uintptr
	mouseDownPtr         uintptr
	mouseEnterPtr        uintptr
	mouseLeavePtr        uintptr
	mouseMovePtr         uintptr
	mouseUpPtr           uintptr
	mouseWheelPtr        uintptr
	mouseWheelDownPtr    uintptr
	mouseWheelUpPtr      uintptr
}

func NewOpenGLControl(TheOwner IComponent) IOpenGLControl {
	r1 := openGLControlImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsOpenGLControl(r1)
}

func OpenGLControlClass() TClass {
	ret := openGLControlImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TOpenGLControl) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(2, m.Instance(), m.constrainedResizePtr)
}

func (m *TOpenGLControl) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(3, m.Instance(), m.dblClickPtr)
}

func (m *TOpenGLControl) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(4, m.Instance(), m.dragDropPtr)
}

func (m *TOpenGLControl) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(5, m.Instance(), m.dragOverPtr)
}

func (m *TOpenGLControl) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(6, m.Instance(), m.mouseDownPtr)
}

func (m *TOpenGLControl) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(7, m.Instance(), m.mouseEnterPtr)
}

func (m *TOpenGLControl) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(8, m.Instance(), m.mouseLeavePtr)
}

func (m *TOpenGLControl) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(9, m.Instance(), m.mouseMovePtr)
}

func (m *TOpenGLControl) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(10, m.Instance(), m.mouseUpPtr)
}

func (m *TOpenGLControl) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(11, m.Instance(), m.mouseWheelPtr)
}

func (m *TOpenGLControl) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(12, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TOpenGLControl) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	openGLControlImportAPI().SysCallN(13, m.Instance(), m.mouseWheelUpPtr)
}

var (
	openGLControlImport       *imports.Imports = nil
	openGLControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("OpenGLControl_Class", 0),
		/*1*/ imports.NewTable("OpenGLControl_Create", 0),
		/*2*/ imports.NewTable("OpenGLControl_SetOnConstrainedResize", 0),
		/*3*/ imports.NewTable("OpenGLControl_SetOnDblClick", 0),
		/*4*/ imports.NewTable("OpenGLControl_SetOnDragDrop", 0),
		/*5*/ imports.NewTable("OpenGLControl_SetOnDragOver", 0),
		/*6*/ imports.NewTable("OpenGLControl_SetOnMouseDown", 0),
		/*7*/ imports.NewTable("OpenGLControl_SetOnMouseEnter", 0),
		/*8*/ imports.NewTable("OpenGLControl_SetOnMouseLeave", 0),
		/*9*/ imports.NewTable("OpenGLControl_SetOnMouseMove", 0),
		/*10*/ imports.NewTable("OpenGLControl_SetOnMouseUp", 0),
		/*11*/ imports.NewTable("OpenGLControl_SetOnMouseWheel", 0),
		/*12*/ imports.NewTable("OpenGLControl_SetOnMouseWheelDown", 0),
		/*13*/ imports.NewTable("OpenGLControl_SetOnMouseWheelUp", 0),
	}
)

func openGLControlImportAPI() *imports.Imports {
	if openGLControlImport == nil {
		openGLControlImport = NewDefaultImports()
		openGLControlImport.SetImportTable(openGLControlImportTables)
		openGLControlImportTables = nil
	}
	return openGLControlImport
}
