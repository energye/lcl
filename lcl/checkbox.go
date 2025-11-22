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

// ICheckBox Parent: ICustomCheckBox
type ICheckBox interface {
	ICustomCheckBox
	Checked() bool                                 // property Checked Getter
	SetChecked(value bool)                         // property Checked Setter
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

type TCheckBox struct {
	TCustomCheckBox
}

func (m *TCheckBox) Checked() bool {
	if !m.IsValid() {
		return false
	}
	r := checkBoxAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckBox) SetChecked(value bool) {
	if !m.IsValid() {
		return
	}
	checkBoxAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := checkBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCheckBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	checkBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCheckBox) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := checkBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TCheckBox) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	checkBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCheckBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := checkBoxAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TCheckBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	checkBoxAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCheckBox) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := checkBoxAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckBox) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	checkBoxAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckBox) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := checkBoxAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckBox) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	checkBoxAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckBox) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := checkBoxAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckBox) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	checkBoxAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 8, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 12, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 13, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 16, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 17, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 18, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 19, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 21, checkBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewCheckBox class constructor
func NewCheckBox(theOwner IComponent) ICheckBox {
	r := checkBoxAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCheckBox(r)
}

func TCheckBoxClass() types.TClass {
	r := checkBoxAPI().SysCallN(22)
	return types.TClass(r)
}

var (
	checkBoxOnce   base.Once
	checkBoxImport *imports.Imports = nil
)

func checkBoxAPI() *imports.Imports {
	checkBoxOnce.Do(func() {
		checkBoxImport = api.NewDefaultImports()
		checkBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCheckBox_Create", 0), // constructor NewCheckBox
			/* 1 */ imports.NewTable("TCheckBox_Checked", 0), // property Checked
			/* 2 */ imports.NewTable("TCheckBox_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TCheckBox_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TCheckBox_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TCheckBox_ParentColor", 0), // property ParentColor
			/* 6 */ imports.NewTable("TCheckBox_ParentFont", 0), // property ParentFont
			/* 7 */ imports.NewTable("TCheckBox_ParentShowHint", 0), // property ParentShowHint
			/* 8 */ imports.NewTable("TCheckBox_OnContextPopup", 0), // event OnContextPopup
			/* 9 */ imports.NewTable("TCheckBox_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TCheckBox_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TCheckBox_OnEditingDone", 0), // event OnEditingDone
			/* 12 */ imports.NewTable("TCheckBox_OnEndDrag", 0), // event OnEndDrag
			/* 13 */ imports.NewTable("TCheckBox_OnMouseDown", 0), // event OnMouseDown
			/* 14 */ imports.NewTable("TCheckBox_OnMouseEnter", 0), // event OnMouseEnter
			/* 15 */ imports.NewTable("TCheckBox_OnMouseLeave", 0), // event OnMouseLeave
			/* 16 */ imports.NewTable("TCheckBox_OnMouseMove", 0), // event OnMouseMove
			/* 17 */ imports.NewTable("TCheckBox_OnMouseUp", 0), // event OnMouseUp
			/* 18 */ imports.NewTable("TCheckBox_OnMouseWheel", 0), // event OnMouseWheel
			/* 19 */ imports.NewTable("TCheckBox_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 20 */ imports.NewTable("TCheckBox_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 21 */ imports.NewTable("TCheckBox_OnStartDrag", 0), // event OnStartDrag
			/* 22 */ imports.NewTable("TCheckBox_TClass", 0), // function TCheckBoxClass
		}
	})
	return checkBoxImport
}
