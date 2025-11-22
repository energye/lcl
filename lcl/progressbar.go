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

// IProgressBar Parent: ICustomProgressBar
type IProgressBar interface {
	ICustomProgressBar
	DragCursor() types.TCursor                     // property DragCursor Getter
	SetDragCursor(value types.TCursor)             // property DragCursor Setter
	DragKind() types.TDragKind                     // property DragKind Getter
	SetDragKind(value types.TDragKind)             // property DragKind Setter
	DragMode() types.TDragMode                     // property DragMode Getter
	SetDragMode(value types.TDragMode)             // property DragMode Setter
	ParentColor() bool                             // property ParentColor Getter
	SetParentColor(value bool)                     // property ParentColor Setter
	ParentFont() bool                              // property ParentFont Getter
	SetParentFont(value bool)                      // property ParentFont Setter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent)       // property event
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
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

type TProgressBar struct {
	TCustomProgressBar
}

func (m *TProgressBar) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := progressBarAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TProgressBar) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	progressBarAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TProgressBar) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := progressBarAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TProgressBar) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	progressBarAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TProgressBar) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := progressBarAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TProgressBar) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	progressBarAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TProgressBar) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := progressBarAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TProgressBar) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	progressBarAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TProgressBar) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := progressBarAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TProgressBar) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	progressBarAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TProgressBar) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := progressBarAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TProgressBar) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	progressBarAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TProgressBar) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 8, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 9, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 10, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 11, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 14, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 15, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 16, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 18, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 19, progressBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TProgressBar) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 20, progressBarAPI(), api.MakeEventDataPtr(cb))
}

// NewProgressBar class constructor
func NewProgressBar(owner IComponent) IProgressBar {
	r := progressBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsProgressBar(r)
}

func TProgressBarClass() types.TClass {
	r := progressBarAPI().SysCallN(21)
	return types.TClass(r)
}

var (
	progressBarOnce   base.Once
	progressBarImport *imports.Imports = nil
)

func progressBarAPI() *imports.Imports {
	progressBarOnce.Do(func() {
		progressBarImport = api.NewDefaultImports()
		progressBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TProgressBar_Create", 0), // constructor NewProgressBar
			/* 1 */ imports.NewTable("TProgressBar_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TProgressBar_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TProgressBar_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TProgressBar_ParentColor", 0), // property ParentColor
			/* 5 */ imports.NewTable("TProgressBar_ParentFont", 0), // property ParentFont
			/* 6 */ imports.NewTable("TProgressBar_ParentShowHint", 0), // property ParentShowHint
			/* 7 */ imports.NewTable("TProgressBar_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TProgressBar_OnDragDrop", 0), // event OnDragDrop
			/* 9 */ imports.NewTable("TProgressBar_OnDragOver", 0), // event OnDragOver
			/* 10 */ imports.NewTable("TProgressBar_OnEndDrag", 0), // event OnEndDrag
			/* 11 */ imports.NewTable("TProgressBar_OnMouseDown", 0), // event OnMouseDown
			/* 12 */ imports.NewTable("TProgressBar_OnMouseEnter", 0), // event OnMouseEnter
			/* 13 */ imports.NewTable("TProgressBar_OnMouseLeave", 0), // event OnMouseLeave
			/* 14 */ imports.NewTable("TProgressBar_OnMouseMove", 0), // event OnMouseMove
			/* 15 */ imports.NewTable("TProgressBar_OnMouseUp", 0), // event OnMouseUp
			/* 16 */ imports.NewTable("TProgressBar_OnMouseWheel", 0), // event OnMouseWheel
			/* 17 */ imports.NewTable("TProgressBar_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 18 */ imports.NewTable("TProgressBar_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 19 */ imports.NewTable("TProgressBar_OnStartDock", 0), // event OnStartDock
			/* 20 */ imports.NewTable("TProgressBar_OnStartDrag", 0), // event OnStartDrag
			/* 21 */ imports.NewTable("TProgressBar_TClass", 0), // function TProgressBarClass
		}
	})
	return progressBarImport
}
