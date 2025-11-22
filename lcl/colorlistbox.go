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

// IColorListBox Parent: ICustomColorListBox
type IColorListBox interface {
	ICustomColorListBox
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

type TColorListBox struct {
	TCustomColorListBox
}

func (m *TColorListBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := colorListBoxAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TColorListBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	colorListBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TColorListBox) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := colorListBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TColorListBox) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	colorListBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TColorListBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := colorListBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TColorListBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	colorListBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TColorListBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 4, colorListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorListBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 5, colorListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorListBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 6, colorListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorListBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 7, colorListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorListBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 8, colorListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorListBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 9, colorListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorListBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 10, colorListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorListBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 11, colorListBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewColorListBox class constructor
func NewColorListBox(owner IComponent) IColorListBox {
	r := colorListBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsColorListBox(r)
}

func TColorListBoxClass() types.TClass {
	r := colorListBoxAPI().SysCallN(12)
	return types.TClass(r)
}

var (
	colorListBoxOnce   base.Once
	colorListBoxImport *imports.Imports = nil
)

func colorListBoxAPI() *imports.Imports {
	colorListBoxOnce.Do(func() {
		colorListBoxImport = api.NewDefaultImports()
		colorListBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TColorListBox_Create", 0), // constructor NewColorListBox
			/* 1 */ imports.NewTable("TColorListBox_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TColorListBox_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TColorListBox_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TColorListBox_OnContextPopup", 0), // event OnContextPopup
			/* 5 */ imports.NewTable("TColorListBox_OnDragDrop", 0), // event OnDragDrop
			/* 6 */ imports.NewTable("TColorListBox_OnDragOver", 0), // event OnDragOver
			/* 7 */ imports.NewTable("TColorListBox_OnEndDrag", 0), // event OnEndDrag
			/* 8 */ imports.NewTable("TColorListBox_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 9 */ imports.NewTable("TColorListBox_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 10 */ imports.NewTable("TColorListBox_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 11 */ imports.NewTable("TColorListBox_OnStartDrag", 0), // event OnStartDrag
			/* 12 */ imports.NewTable("TColorListBox_TClass", 0), // function TColorListBoxClass
		}
	})
	return colorListBoxImport
}
