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

// IBitBtn Parent: ICustomBitBtn
type IBitBtn interface {
	ICustomBitBtn
	DragCursor() types.TCursor                     // property DragCursor Getter
	SetDragCursor(value types.TCursor)             // property DragCursor Setter
	DragKind() types.TDragKind                     // property DragKind Getter
	SetDragKind(value types.TDragKind)             // property DragKind Setter
	DragMode() types.TDragMode                     // property DragMode Getter
	SetDragMode(value types.TDragMode)             // property DragMode Setter
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
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

type TBitBtn struct {
	TCustomBitBtn
}

func (m *TBitBtn) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := bitBtnAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TBitBtn) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	bitBtnAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TBitBtn) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := bitBtnAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TBitBtn) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	bitBtnAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TBitBtn) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := bitBtnAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TBitBtn) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	bitBtnAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TBitBtn) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := bitBtnAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBitBtn) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	bitBtnAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TBitBtn) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := bitBtnAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBitBtn) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	bitBtnAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TBitBtn) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 6, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 7, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 8, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 9, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 10, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 13, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 15, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBitBtn) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 18, bitBtnAPI(), api.MakeEventDataPtr(cb))
}

// NewBitBtn class constructor
func NewBitBtn(theOwner IComponent) IBitBtn {
	r := bitBtnAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsBitBtn(r)
}

func TBitBtnClass() types.TClass {
	r := bitBtnAPI().SysCallN(19)
	return types.TClass(r)
}

var (
	bitBtnOnce   base.Once
	bitBtnImport *imports.Imports = nil
)

func bitBtnAPI() *imports.Imports {
	bitBtnOnce.Do(func() {
		bitBtnImport = api.NewDefaultImports()
		bitBtnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBitBtn_Create", 0), // constructor NewBitBtn
			/* 1 */ imports.NewTable("TBitBtn_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TBitBtn_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TBitBtn_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TBitBtn_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TBitBtn_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TBitBtn_OnContextPopup", 0), // event OnContextPopup
			/* 7 */ imports.NewTable("TBitBtn_OnDragDrop", 0), // event OnDragDrop
			/* 8 */ imports.NewTable("TBitBtn_OnDragOver", 0), // event OnDragOver
			/* 9 */ imports.NewTable("TBitBtn_OnEndDrag", 0), // event OnEndDrag
			/* 10 */ imports.NewTable("TBitBtn_OnMouseDown", 0), // event OnMouseDown
			/* 11 */ imports.NewTable("TBitBtn_OnMouseEnter", 0), // event OnMouseEnter
			/* 12 */ imports.NewTable("TBitBtn_OnMouseLeave", 0), // event OnMouseLeave
			/* 13 */ imports.NewTable("TBitBtn_OnMouseMove", 0), // event OnMouseMove
			/* 14 */ imports.NewTable("TBitBtn_OnMouseUp", 0), // event OnMouseUp
			/* 15 */ imports.NewTable("TBitBtn_OnMouseWheel", 0), // event OnMouseWheel
			/* 16 */ imports.NewTable("TBitBtn_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 17 */ imports.NewTable("TBitBtn_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 18 */ imports.NewTable("TBitBtn_OnStartDrag", 0), // event OnStartDrag
			/* 19 */ imports.NewTable("TBitBtn_TClass", 0), // function TBitBtnClass
		}
	})
	return bitBtnImport
}
