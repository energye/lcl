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

// IMemo Parent: ICustomMemo
type IMemo interface {
	ICustomMemo
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
	DragKind() types.TDragKind                      // property DragKind Getter
	SetDragKind(value types.TDragKind)              // property DragKind Setter
	DragMode() types.TDragMode                      // property DragMode Getter
	SetDragMode(value types.TDragMode)              // property DragMode Setter
	ParentColor() bool                              // property ParentColor Getter
	SetParentColor(value bool)                      // property ParentColor Setter
	ParentFont() bool                               // property ParentFont Getter
	SetParentFont(value bool)                       // property ParentFont Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEditingDone(fn TNotifyEvent)               // property event
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

type TMemo struct {
	TCustomMemo
}

func (m *TMemo) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := memoAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TMemo) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	memoAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TMemo) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := memoAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TMemo) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	memoAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TMemo) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := memoAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TMemo) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	memoAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TMemo) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := memoAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMemo) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	memoAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TMemo) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := memoAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMemo) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	memoAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TMemo) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := memoAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMemo) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	memoAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TMemo) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 12, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 13, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 16, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 17, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 18, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 19, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 21, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 22, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 23, memoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMemo) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 24, memoAPI(), api.MakeEventDataPtr(cb))
}

// NewMemo class constructor
func NewMemo(owner IComponent) IMemo {
	r := memoAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsMemo(r)
}

func TMemoClass() types.TClass {
	r := memoAPI().SysCallN(25)
	return types.TClass(r)
}

var (
	memoOnce   base.Once
	memoImport *imports.Imports = nil
)

func memoAPI() *imports.Imports {
	memoOnce.Do(func() {
		memoImport = api.NewDefaultImports()
		memoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMemo_Create", 0), // constructor NewMemo
			/* 1 */ imports.NewTable("TMemo_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TMemo_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TMemo_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TMemo_ParentColor", 0), // property ParentColor
			/* 5 */ imports.NewTable("TMemo_ParentFont", 0), // property ParentFont
			/* 6 */ imports.NewTable("TMemo_ParentShowHint", 0), // property ParentShowHint
			/* 7 */ imports.NewTable("TMemo_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TMemo_OnDblClick", 0), // event OnDblClick
			/* 9 */ imports.NewTable("TMemo_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TMemo_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TMemo_OnEditingDone", 0), // event OnEditingDone
			/* 12 */ imports.NewTable("TMemo_OnEndDrag", 0), // event OnEndDrag
			/* 13 */ imports.NewTable("TMemo_OnMouseDown", 0), // event OnMouseDown
			/* 14 */ imports.NewTable("TMemo_OnMouseEnter", 0), // event OnMouseEnter
			/* 15 */ imports.NewTable("TMemo_OnMouseLeave", 0), // event OnMouseLeave
			/* 16 */ imports.NewTable("TMemo_OnMouseMove", 0), // event OnMouseMove
			/* 17 */ imports.NewTable("TMemo_OnMouseUp", 0), // event OnMouseUp
			/* 18 */ imports.NewTable("TMemo_OnMouseWheel", 0), // event OnMouseWheel
			/* 19 */ imports.NewTable("TMemo_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 20 */ imports.NewTable("TMemo_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 21 */ imports.NewTable("TMemo_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 22 */ imports.NewTable("TMemo_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 23 */ imports.NewTable("TMemo_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 24 */ imports.NewTable("TMemo_OnStartDrag", 0), // event OnStartDrag
			/* 25 */ imports.NewTable("TMemo_TClass", 0), // function TMemoClass
		}
	})
	return memoImport
}
