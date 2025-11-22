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

// IButtonPanel Parent: ICustomButtonPanel
type IButtonPanel interface {
	ICustomButtonPanel
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
}

type TButtonPanel struct {
	TCustomButtonPanel
}

func (m *TButtonPanel) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 1, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButtonPanel) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 2, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButtonPanel) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 3, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButtonPanel) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 4, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButtonPanel) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 5, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButtonPanel) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 6, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButtonPanel) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 7, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButtonPanel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 8, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButtonPanel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 9, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButtonPanel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 10, buttonPanelAPI(), api.MakeEventDataPtr(cb))
}

// NewButtonPanel class constructor
func NewButtonPanel(owner IComponent) IButtonPanel {
	r := buttonPanelAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsButtonPanel(r)
}

func TButtonPanelClass() types.TClass {
	r := buttonPanelAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	buttonPanelOnce   base.Once
	buttonPanelImport *imports.Imports = nil
)

func buttonPanelAPI() *imports.Imports {
	buttonPanelOnce.Do(func() {
		buttonPanelImport = api.NewDefaultImports()
		buttonPanelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TButtonPanel_Create", 0), // constructor NewButtonPanel
			/* 1 */ imports.NewTable("TButtonPanel_OnDblClick", 0), // event OnDblClick
			/* 2 */ imports.NewTable("TButtonPanel_OnDragDrop", 0), // event OnDragDrop
			/* 3 */ imports.NewTable("TButtonPanel_OnMouseDown", 0), // event OnMouseDown
			/* 4 */ imports.NewTable("TButtonPanel_OnMouseEnter", 0), // event OnMouseEnter
			/* 5 */ imports.NewTable("TButtonPanel_OnMouseLeave", 0), // event OnMouseLeave
			/* 6 */ imports.NewTable("TButtonPanel_OnMouseMove", 0), // event OnMouseMove
			/* 7 */ imports.NewTable("TButtonPanel_OnMouseUp", 0), // event OnMouseUp
			/* 8 */ imports.NewTable("TButtonPanel_OnMouseWheel", 0), // event OnMouseWheel
			/* 9 */ imports.NewTable("TButtonPanel_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 10 */ imports.NewTable("TButtonPanel_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 11 */ imports.NewTable("TButtonPanel_TClass", 0), // function TButtonPanelClass
		}
	})
	return buttonPanelImport
}
