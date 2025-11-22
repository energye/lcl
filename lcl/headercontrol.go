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

// IHeaderControl Parent: ICustomHeaderControl
type IHeaderControl interface {
	ICustomHeaderControl
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
	DragKind() types.TDragKind                      // property DragKind Getter
	SetDragKind(value types.TDragKind)              // property DragKind Setter
	DragMode() types.TDragMode                      // property DragMode Getter
	SetDragMode(value types.TDragMode)              // property DragMode Setter
	ParentFont() bool                               // property ParentFont Getter
	SetParentFont(value bool)                       // property ParentFont Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDock(fn TEndDragEvent)                  // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
}

type THeaderControl struct {
	TCustomHeaderControl
}

func (m *THeaderControl) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := headerControlAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *THeaderControl) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	headerControlAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *THeaderControl) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := headerControlAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *THeaderControl) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	headerControlAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *THeaderControl) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := headerControlAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *THeaderControl) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	headerControlAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *THeaderControl) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := headerControlAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *THeaderControl) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	headerControlAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *THeaderControl) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := headerControlAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *THeaderControl) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	headerControlAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *THeaderControl) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 6, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 7, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 8, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 9, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 10, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 11, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 14, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 15, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 16, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 18, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 19, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, headerControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *THeaderControl) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, headerControlAPI(), api.MakeEventDataPtr(cb))
}

// NewHeaderControl class constructor
func NewHeaderControl(owner IComponent) IHeaderControl {
	r := headerControlAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsHeaderControl(r)
}

func THeaderControlClass() types.TClass {
	r := headerControlAPI().SysCallN(22)
	return types.TClass(r)
}

var (
	headerControlOnce   base.Once
	headerControlImport *imports.Imports = nil
)

func headerControlAPI() *imports.Imports {
	headerControlOnce.Do(func() {
		headerControlImport = api.NewDefaultImports()
		headerControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("THeaderControl_Create", 0), // constructor NewHeaderControl
			/* 1 */ imports.NewTable("THeaderControl_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("THeaderControl_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("THeaderControl_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("THeaderControl_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("THeaderControl_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("THeaderControl_OnContextPopup", 0), // event OnContextPopup
			/* 7 */ imports.NewTable("THeaderControl_OnDragDrop", 0), // event OnDragDrop
			/* 8 */ imports.NewTable("THeaderControl_OnDragOver", 0), // event OnDragOver
			/* 9 */ imports.NewTable("THeaderControl_OnEndDock", 0), // event OnEndDock
			/* 10 */ imports.NewTable("THeaderControl_OnEndDrag", 0), // event OnEndDrag
			/* 11 */ imports.NewTable("THeaderControl_OnMouseDown", 0), // event OnMouseDown
			/* 12 */ imports.NewTable("THeaderControl_OnMouseEnter", 0), // event OnMouseEnter
			/* 13 */ imports.NewTable("THeaderControl_OnMouseLeave", 0), // event OnMouseLeave
			/* 14 */ imports.NewTable("THeaderControl_OnMouseMove", 0), // event OnMouseMove
			/* 15 */ imports.NewTable("THeaderControl_OnMouseUp", 0), // event OnMouseUp
			/* 16 */ imports.NewTable("THeaderControl_OnMouseWheel", 0), // event OnMouseWheel
			/* 17 */ imports.NewTable("THeaderControl_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 18 */ imports.NewTable("THeaderControl_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 19 */ imports.NewTable("THeaderControl_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 20 */ imports.NewTable("THeaderControl_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 21 */ imports.NewTable("THeaderControl_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 22 */ imports.NewTable("THeaderControl_TClass", 0), // function THeaderControlClass
		}
	})
	return headerControlImport
}
