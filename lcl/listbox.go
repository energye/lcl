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

// IListBox Parent: ICustomListBox
type IListBox interface {
	ICustomListBox
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
	DragKind() types.TDragKind                      // property DragKind Getter
	SetDragKind(value types.TDragKind)              // property DragKind Setter
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

type TListBox struct {
	TCustomListBox
}

func (m *TListBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := listBoxAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TListBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	listBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TListBox) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := listBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TListBox) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	listBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TListBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := listBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TListBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	listBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TListBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 4, listBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 5, listBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 6, listBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 7, listBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 8, listBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 9, listBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 10, listBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 11, listBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewListBox class constructor
func NewListBox(theOwner IComponent) IListBox {
	r := listBoxAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsListBox(r)
}

func TListBoxClass() types.TClass {
	r := listBoxAPI().SysCallN(12)
	return types.TClass(r)
}

var (
	listBoxOnce   base.Once
	listBoxImport *imports.Imports = nil
)

func listBoxAPI() *imports.Imports {
	listBoxOnce.Do(func() {
		listBoxImport = api.NewDefaultImports()
		listBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListBox_Create", 0), // constructor NewListBox
			/* 1 */ imports.NewTable("TListBox_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TListBox_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TListBox_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TListBox_OnContextPopup", 0), // event OnContextPopup
			/* 5 */ imports.NewTable("TListBox_OnDragDrop", 0), // event OnDragDrop
			/* 6 */ imports.NewTable("TListBox_OnDragOver", 0), // event OnDragOver
			/* 7 */ imports.NewTable("TListBox_OnEndDrag", 0), // event OnEndDrag
			/* 8 */ imports.NewTable("TListBox_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 9 */ imports.NewTable("TListBox_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 10 */ imports.NewTable("TListBox_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 11 */ imports.NewTable("TListBox_OnStartDrag", 0), // event OnStartDrag
			/* 12 */ imports.NewTable("TListBox_TClass", 0), // function TListBoxClass
		}
	})
	return listBoxImport
}
