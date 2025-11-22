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

// ITrackBar Parent: ICustomTrackBar
type ITrackBar interface {
	ICustomTrackBar
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
	DragMode() types.TDragMode                      // property DragMode Getter
	SetDragMode(value types.TDragMode)              // property DragMode Setter
	ParentColor() bool                              // property ParentColor Getter
	SetParentColor(value bool)                      // property ParentColor Setter
	ParentFont() bool                               // property ParentFont Getter
	SetParentFont(value bool)                       // property ParentFont Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
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
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

type TTrackBar struct {
	TCustomTrackBar
}

func (m *TTrackBar) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := trackBarAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TTrackBar) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	trackBarAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TTrackBar) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := trackBarAPI().SysCallN(2, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TTrackBar) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	trackBarAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TTrackBar) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := trackBarAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTrackBar) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	trackBarAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TTrackBar) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := trackBarAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTrackBar) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	trackBarAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TTrackBar) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := trackBarAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTrackBar) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	trackBarAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TTrackBar) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 6, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 7, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 8, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 9, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 10, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 13, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 15, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 18, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 19, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, trackBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTrackBar) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 21, trackBarAPI(), api.MakeEventDataPtr(cb))
}

// NewTrackBar class constructor
func NewTrackBar(owner IComponent) ITrackBar {
	r := trackBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsTrackBar(r)
}

func TTrackBarClass() types.TClass {
	r := trackBarAPI().SysCallN(22)
	return types.TClass(r)
}

var (
	trackBarOnce   base.Once
	trackBarImport *imports.Imports = nil
)

func trackBarAPI() *imports.Imports {
	trackBarOnce.Do(func() {
		trackBarImport = api.NewDefaultImports()
		trackBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTrackBar_Create", 0), // constructor NewTrackBar
			/* 1 */ imports.NewTable("TTrackBar_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TTrackBar_DragMode", 0), // property DragMode
			/* 3 */ imports.NewTable("TTrackBar_ParentColor", 0), // property ParentColor
			/* 4 */ imports.NewTable("TTrackBar_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TTrackBar_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TTrackBar_OnContextPopup", 0), // event OnContextPopup
			/* 7 */ imports.NewTable("TTrackBar_OnDragDrop", 0), // event OnDragDrop
			/* 8 */ imports.NewTable("TTrackBar_OnDragOver", 0), // event OnDragOver
			/* 9 */ imports.NewTable("TTrackBar_OnEndDrag", 0), // event OnEndDrag
			/* 10 */ imports.NewTable("TTrackBar_OnMouseDown", 0), // event OnMouseDown
			/* 11 */ imports.NewTable("TTrackBar_OnMouseEnter", 0), // event OnMouseEnter
			/* 12 */ imports.NewTable("TTrackBar_OnMouseLeave", 0), // event OnMouseLeave
			/* 13 */ imports.NewTable("TTrackBar_OnMouseMove", 0), // event OnMouseMove
			/* 14 */ imports.NewTable("TTrackBar_OnMouseUp", 0), // event OnMouseUp
			/* 15 */ imports.NewTable("TTrackBar_OnMouseWheel", 0), // event OnMouseWheel
			/* 16 */ imports.NewTable("TTrackBar_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 17 */ imports.NewTable("TTrackBar_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 18 */ imports.NewTable("TTrackBar_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 19 */ imports.NewTable("TTrackBar_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 20 */ imports.NewTable("TTrackBar_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 21 */ imports.NewTable("TTrackBar_OnStartDrag", 0), // event OnStartDrag
			/* 22 */ imports.NewTable("TTrackBar_TClass", 0), // function TTrackBarClass
		}
	})
	return trackBarImport
}
