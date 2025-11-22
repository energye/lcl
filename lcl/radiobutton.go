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

// IRadioButton Parent: ICustomCheckBox
type IRadioButton interface {
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

type TRadioButton struct {
	TCustomCheckBox
}

func (m *TRadioButton) Checked() bool {
	if !m.IsValid() {
		return false
	}
	r := radioButtonAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRadioButton) SetChecked(value bool) {
	if !m.IsValid() {
		return
	}
	radioButtonAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TRadioButton) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := radioButtonAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TRadioButton) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	radioButtonAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TRadioButton) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := radioButtonAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TRadioButton) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	radioButtonAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TRadioButton) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := radioButtonAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TRadioButton) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	radioButtonAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TRadioButton) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := radioButtonAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRadioButton) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	radioButtonAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TRadioButton) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := radioButtonAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRadioButton) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	radioButtonAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TRadioButton) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := radioButtonAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRadioButton) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	radioButtonAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TRadioButton) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 8, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 11, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 12, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 15, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 16, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 17, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 18, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 19, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioButton) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 20, radioButtonAPI(), api.MakeEventDataPtr(cb))
}

// NewRadioButton class constructor
func NewRadioButton(theOwner IComponent) IRadioButton {
	r := radioButtonAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsRadioButton(r)
}

func TRadioButtonClass() types.TClass {
	r := radioButtonAPI().SysCallN(21)
	return types.TClass(r)
}

var (
	radioButtonOnce   base.Once
	radioButtonImport *imports.Imports = nil
)

func radioButtonAPI() *imports.Imports {
	radioButtonOnce.Do(func() {
		radioButtonImport = api.NewDefaultImports()
		radioButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRadioButton_Create", 0), // constructor NewRadioButton
			/* 1 */ imports.NewTable("TRadioButton_Checked", 0), // property Checked
			/* 2 */ imports.NewTable("TRadioButton_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TRadioButton_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TRadioButton_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TRadioButton_ParentColor", 0), // property ParentColor
			/* 6 */ imports.NewTable("TRadioButton_ParentFont", 0), // property ParentFont
			/* 7 */ imports.NewTable("TRadioButton_ParentShowHint", 0), // property ParentShowHint
			/* 8 */ imports.NewTable("TRadioButton_OnContextPopup", 0), // event OnContextPopup
			/* 9 */ imports.NewTable("TRadioButton_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TRadioButton_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TRadioButton_OnEndDrag", 0), // event OnEndDrag
			/* 12 */ imports.NewTable("TRadioButton_OnMouseDown", 0), // event OnMouseDown
			/* 13 */ imports.NewTable("TRadioButton_OnMouseEnter", 0), // event OnMouseEnter
			/* 14 */ imports.NewTable("TRadioButton_OnMouseLeave", 0), // event OnMouseLeave
			/* 15 */ imports.NewTable("TRadioButton_OnMouseMove", 0), // event OnMouseMove
			/* 16 */ imports.NewTable("TRadioButton_OnMouseUp", 0), // event OnMouseUp
			/* 17 */ imports.NewTable("TRadioButton_OnMouseWheel", 0), // event OnMouseWheel
			/* 18 */ imports.NewTable("TRadioButton_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 19 */ imports.NewTable("TRadioButton_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 20 */ imports.NewTable("TRadioButton_OnStartDrag", 0), // event OnStartDrag
			/* 21 */ imports.NewTable("TRadioButton_TClass", 0), // function TRadioButtonClass
		}
	})
	return radioButtonImport
}
