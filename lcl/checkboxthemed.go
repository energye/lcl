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

// ICheckBoxThemed Parent: ICustomCheckBoxThemed
type ICheckBoxThemed interface {
	ICustomCheckBoxThemed
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
	SetOnEditingDone(fn TNotifyEvent)              // property event
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

type TCheckBoxThemed struct {
	TCustomCheckBoxThemed
}

func (m *TCheckBoxThemed) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := checkBoxThemedAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCheckBoxThemed) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	checkBoxThemedAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCheckBoxThemed) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := checkBoxThemedAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TCheckBoxThemed) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	checkBoxThemedAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCheckBoxThemed) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := checkBoxThemedAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TCheckBoxThemed) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	checkBoxThemedAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCheckBoxThemed) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := checkBoxThemedAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckBoxThemed) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	checkBoxThemedAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckBoxThemed) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := checkBoxThemedAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckBoxThemed) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	checkBoxThemedAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckBoxThemed) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := checkBoxThemedAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckBoxThemed) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	checkBoxThemedAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckBoxThemed) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 8, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 9, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 11, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 12, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 15, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 16, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 17, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 18, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 19, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBoxThemed) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 20, checkBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

// NewCheckBoxThemed class constructor
func NewCheckBoxThemed(owner IComponent) ICheckBoxThemed {
	r := checkBoxThemedAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCheckBoxThemed(r)
}

func TCheckBoxThemedClass() types.TClass {
	r := checkBoxThemedAPI().SysCallN(21)
	return types.TClass(r)
}

var (
	checkBoxThemedOnce   base.Once
	checkBoxThemedImport *imports.Imports = nil
)

func checkBoxThemedAPI() *imports.Imports {
	checkBoxThemedOnce.Do(func() {
		checkBoxThemedImport = api.NewDefaultImports()
		checkBoxThemedImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCheckBoxThemed_Create", 0), // constructor NewCheckBoxThemed
			/* 1 */ imports.NewTable("TCheckBoxThemed_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TCheckBoxThemed_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TCheckBoxThemed_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TCheckBoxThemed_ParentColor", 0), // property ParentColor
			/* 5 */ imports.NewTable("TCheckBoxThemed_ParentFont", 0), // property ParentFont
			/* 6 */ imports.NewTable("TCheckBoxThemed_ParentShowHint", 0), // property ParentShowHint
			/* 7 */ imports.NewTable("TCheckBoxThemed_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TCheckBoxThemed_OnDragDrop", 0), // event OnDragDrop
			/* 9 */ imports.NewTable("TCheckBoxThemed_OnDragOver", 0), // event OnDragOver
			/* 10 */ imports.NewTable("TCheckBoxThemed_OnEditingDone", 0), // event OnEditingDone
			/* 11 */ imports.NewTable("TCheckBoxThemed_OnEndDrag", 0), // event OnEndDrag
			/* 12 */ imports.NewTable("TCheckBoxThemed_OnMouseDown", 0), // event OnMouseDown
			/* 13 */ imports.NewTable("TCheckBoxThemed_OnMouseEnter", 0), // event OnMouseEnter
			/* 14 */ imports.NewTable("TCheckBoxThemed_OnMouseLeave", 0), // event OnMouseLeave
			/* 15 */ imports.NewTable("TCheckBoxThemed_OnMouseMove", 0), // event OnMouseMove
			/* 16 */ imports.NewTable("TCheckBoxThemed_OnMouseUp", 0), // event OnMouseUp
			/* 17 */ imports.NewTable("TCheckBoxThemed_OnMouseWheel", 0), // event OnMouseWheel
			/* 18 */ imports.NewTable("TCheckBoxThemed_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 19 */ imports.NewTable("TCheckBoxThemed_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 20 */ imports.NewTable("TCheckBoxThemed_OnStartDrag", 0), // event OnStartDrag
			/* 21 */ imports.NewTable("TCheckBoxThemed_TClass", 0), // function TCheckBoxThemedClass
		}
	})
	return checkBoxThemedImport
}
