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

// ICheckListBox Parent: ICustomCheckListBox
type ICheckListBox interface {
	ICustomCheckListBox
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
	DragMode() types.TDragMode                      // property DragMode Getter
	SetDragMode(value types.TDragMode)              // property DragMode Setter
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

type TCheckListBox struct {
	TCustomCheckListBox
}

func (m *TCheckListBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := checkListBoxAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCheckListBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	checkListBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCheckListBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := checkListBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TCheckListBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	checkListBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCheckListBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 3, checkListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckListBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 4, checkListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckListBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 5, checkListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckListBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 6, checkListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckListBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 7, checkListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckListBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 8, checkListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckListBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 9, checkListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckListBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 10, checkListBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewCheckListBox class constructor
func NewCheckListBox(owner IComponent) ICheckListBox {
	r := checkListBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCheckListBox(r)
}

func TCheckListBoxClass() types.TClass {
	r := checkListBoxAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	checkListBoxOnce   base.Once
	checkListBoxImport *imports.Imports = nil
)

func checkListBoxAPI() *imports.Imports {
	checkListBoxOnce.Do(func() {
		checkListBoxImport = api.NewDefaultImports()
		checkListBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCheckListBox_Create", 0), // constructor NewCheckListBox
			/* 1 */ imports.NewTable("TCheckListBox_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TCheckListBox_DragMode", 0), // property DragMode
			/* 3 */ imports.NewTable("TCheckListBox_OnContextPopup", 0), // event OnContextPopup
			/* 4 */ imports.NewTable("TCheckListBox_OnDragDrop", 0), // event OnDragDrop
			/* 5 */ imports.NewTable("TCheckListBox_OnDragOver", 0), // event OnDragOver
			/* 6 */ imports.NewTable("TCheckListBox_OnEndDrag", 0), // event OnEndDrag
			/* 7 */ imports.NewTable("TCheckListBox_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 8 */ imports.NewTable("TCheckListBox_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 9 */ imports.NewTable("TCheckListBox_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 10 */ imports.NewTable("TCheckListBox_OnStartDrag", 0), // event OnStartDrag
			/* 11 */ imports.NewTable("TCheckListBox_TClass", 0), // function TCheckListBoxClass
		}
	})
	return checkListBoxImport
}
