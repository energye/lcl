//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
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

type TOpenGLControl struct {
	TCustomOpenGLControl
}

func (m *TOpenGLControl) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTConstrainedResizeEvent(fn)
	base.SetEvent(m, 1, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 2, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 3, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 4, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 5, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 8, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 9, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 10, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 11, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenGLControl) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 12, openGLControlAPI(), api.MakeEventDataPtr(cb))
}

// NewOpenGLControl class constructor
func NewOpenGLControl(theOwner IComponent) IOpenGLControl {
	r := openGLControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsOpenGLControl(r)
}

func TOpenGLControlClass() types.TClass {
	r := openGLControlAPI().SysCallN(13)
	return types.TClass(r)
}

var (
	openGLControlOnce   base.Once
	openGLControlImport *imports.Imports = nil
)

func openGLControlAPI() *imports.Imports {
	openGLControlOnce.Do(func() {
		openGLControlImport = api.NewDefaultImports()
		openGLControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TOpenGLControl_Create", 0), // constructor NewOpenGLControl
			/* 1 */ imports.NewTable("TOpenGLControl_OnConstrainedResize", 0), // event OnConstrainedResize
			/* 2 */ imports.NewTable("TOpenGLControl_OnDblClick", 0), // event OnDblClick
			/* 3 */ imports.NewTable("TOpenGLControl_OnDragDrop", 0), // event OnDragDrop
			/* 4 */ imports.NewTable("TOpenGLControl_OnDragOver", 0), // event OnDragOver
			/* 5 */ imports.NewTable("TOpenGLControl_OnMouseDown", 0), // event OnMouseDown
			/* 6 */ imports.NewTable("TOpenGLControl_OnMouseEnter", 0), // event OnMouseEnter
			/* 7 */ imports.NewTable("TOpenGLControl_OnMouseLeave", 0), // event OnMouseLeave
			/* 8 */ imports.NewTable("TOpenGLControl_OnMouseMove", 0), // event OnMouseMove
			/* 9 */ imports.NewTable("TOpenGLControl_OnMouseUp", 0), // event OnMouseUp
			/* 10 */ imports.NewTable("TOpenGLControl_OnMouseWheel", 0), // event OnMouseWheel
			/* 11 */ imports.NewTable("TOpenGLControl_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 12 */ imports.NewTable("TOpenGLControl_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 13 */ imports.NewTable("TOpenGLControl_TClass", 0), // function TOpenGLControlClass
		}
	})
	return openGLControlImport
}
