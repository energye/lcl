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

// IToggleBox Parent: ICustomCheckBox
type IToggleBox interface {
	ICustomCheckBox
	Checked() bool                                 // property Checked Getter
	SetChecked(value bool)                         // property Checked Setter
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

type TToggleBox struct {
	TCustomCheckBox
}

func (m *TToggleBox) Checked() bool {
	if !m.IsValid() {
		return false
	}
	r := toggleBoxAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToggleBox) SetChecked(value bool) {
	if !m.IsValid() {
		return
	}
	toggleBoxAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TToggleBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := toggleBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TToggleBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	toggleBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TToggleBox) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := toggleBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TToggleBox) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	toggleBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TToggleBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := toggleBoxAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TToggleBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	toggleBoxAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TToggleBox) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := toggleBoxAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToggleBox) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	toggleBoxAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TToggleBox) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := toggleBoxAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToggleBox) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	toggleBoxAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TToggleBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 8, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 9, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 10, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 11, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 14, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 15, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 16, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 18, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToggleBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 19, toggleBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewToggleBox class constructor
func NewToggleBox(theOwner IComponent) IToggleBox {
	r := toggleBoxAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsToggleBox(r)
}

func TToggleBoxClass() types.TClass {
	r := toggleBoxAPI().SysCallN(20)
	return types.TClass(r)
}

var (
	toggleBoxOnce   base.Once
	toggleBoxImport *imports.Imports = nil
)

func toggleBoxAPI() *imports.Imports {
	toggleBoxOnce.Do(func() {
		toggleBoxImport = api.NewDefaultImports()
		toggleBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TToggleBox_Create", 0), // constructor NewToggleBox
			/* 1 */ imports.NewTable("TToggleBox_Checked", 0), // property Checked
			/* 2 */ imports.NewTable("TToggleBox_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TToggleBox_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TToggleBox_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TToggleBox_ParentFont", 0), // property ParentFont
			/* 6 */ imports.NewTable("TToggleBox_ParentShowHint", 0), // property ParentShowHint
			/* 7 */ imports.NewTable("TToggleBox_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TToggleBox_OnDragDrop", 0), // event OnDragDrop
			/* 9 */ imports.NewTable("TToggleBox_OnDragOver", 0), // event OnDragOver
			/* 10 */ imports.NewTable("TToggleBox_OnEndDrag", 0), // event OnEndDrag
			/* 11 */ imports.NewTable("TToggleBox_OnMouseDown", 0), // event OnMouseDown
			/* 12 */ imports.NewTable("TToggleBox_OnMouseEnter", 0), // event OnMouseEnter
			/* 13 */ imports.NewTable("TToggleBox_OnMouseLeave", 0), // event OnMouseLeave
			/* 14 */ imports.NewTable("TToggleBox_OnMouseMove", 0), // event OnMouseMove
			/* 15 */ imports.NewTable("TToggleBox_OnMouseUp", 0), // event OnMouseUp
			/* 16 */ imports.NewTable("TToggleBox_OnMouseWheel", 0), // event OnMouseWheel
			/* 17 */ imports.NewTable("TToggleBox_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 18 */ imports.NewTable("TToggleBox_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 19 */ imports.NewTable("TToggleBox_OnStartDrag", 0), // event OnStartDrag
			/* 20 */ imports.NewTable("TToggleBox_TClass", 0), // function TToggleBoxClass
		}
	})
	return toggleBoxImport
}
