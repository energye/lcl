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

// IRadioGroup Parent: ICustomRadioGroup
type IRadioGroup interface {
	ICustomRadioGroup
	DragCursor() types.TCursor                     // property DragCursor Getter
	SetDragCursor(value types.TCursor)             // property DragCursor Setter
	DragMode() types.TDragMode                     // property DragMode Getter
	SetDragMode(value types.TDragMode)             // property DragMode Setter
	ParentFont() bool                              // property ParentFont Getter
	SetParentFont(value bool)                      // property ParentFont Setter
	ParentColor() bool                             // property ParentColor Getter
	SetParentColor(value bool)                     // property ParentColor Setter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	SetOnDblClick(fn TNotifyEvent)                 // property event
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

type TRadioGroup struct {
	TCustomRadioGroup
}

func (m *TRadioGroup) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := radioGroupAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TRadioGroup) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	radioGroupAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TRadioGroup) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := radioGroupAPI().SysCallN(2, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TRadioGroup) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	radioGroupAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TRadioGroup) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := radioGroupAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRadioGroup) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	radioGroupAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TRadioGroup) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := radioGroupAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRadioGroup) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	radioGroupAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TRadioGroup) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := radioGroupAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRadioGroup) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	radioGroupAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TRadioGroup) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 7, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 8, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 9, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 10, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 13, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 15, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRadioGroup) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 18, radioGroupAPI(), api.MakeEventDataPtr(cb))
}

// NewRadioGroup class constructor
func NewRadioGroup(theOwner IComponent) IRadioGroup {
	r := radioGroupAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsRadioGroup(r)
}

func TRadioGroupClass() types.TClass {
	r := radioGroupAPI().SysCallN(19)
	return types.TClass(r)
}

var (
	radioGroupOnce   base.Once
	radioGroupImport *imports.Imports = nil
)

func radioGroupAPI() *imports.Imports {
	radioGroupOnce.Do(func() {
		radioGroupImport = api.NewDefaultImports()
		radioGroupImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRadioGroup_Create", 0), // constructor NewRadioGroup
			/* 1 */ imports.NewTable("TRadioGroup_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TRadioGroup_DragMode", 0), // property DragMode
			/* 3 */ imports.NewTable("TRadioGroup_ParentFont", 0), // property ParentFont
			/* 4 */ imports.NewTable("TRadioGroup_ParentColor", 0), // property ParentColor
			/* 5 */ imports.NewTable("TRadioGroup_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TRadioGroup_OnDblClick", 0), // event OnDblClick
			/* 7 */ imports.NewTable("TRadioGroup_OnDragDrop", 0), // event OnDragDrop
			/* 8 */ imports.NewTable("TRadioGroup_OnDragOver", 0), // event OnDragOver
			/* 9 */ imports.NewTable("TRadioGroup_OnEndDrag", 0), // event OnEndDrag
			/* 10 */ imports.NewTable("TRadioGroup_OnMouseDown", 0), // event OnMouseDown
			/* 11 */ imports.NewTable("TRadioGroup_OnMouseEnter", 0), // event OnMouseEnter
			/* 12 */ imports.NewTable("TRadioGroup_OnMouseLeave", 0), // event OnMouseLeave
			/* 13 */ imports.NewTable("TRadioGroup_OnMouseMove", 0), // event OnMouseMove
			/* 14 */ imports.NewTable("TRadioGroup_OnMouseUp", 0), // event OnMouseUp
			/* 15 */ imports.NewTable("TRadioGroup_OnMouseWheel", 0), // event OnMouseWheel
			/* 16 */ imports.NewTable("TRadioGroup_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 17 */ imports.NewTable("TRadioGroup_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 18 */ imports.NewTable("TRadioGroup_OnStartDrag", 0), // event OnStartDrag
			/* 19 */ imports.NewTable("TRadioGroup_TClass", 0), // function TRadioGroupClass
		}
	})
	return radioGroupImport
}
