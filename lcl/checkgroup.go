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

// ICheckGroup Parent: ICustomCheckGroup
type ICheckGroup interface {
	ICustomCheckGroup
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

type TCheckGroup struct {
	TCustomCheckGroup
}

func (m *TCheckGroup) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := checkGroupAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCheckGroup) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	checkGroupAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCheckGroup) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := checkGroupAPI().SysCallN(2, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TCheckGroup) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	checkGroupAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCheckGroup) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := checkGroupAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckGroup) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	checkGroupAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckGroup) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := checkGroupAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckGroup) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	checkGroupAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckGroup) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := checkGroupAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckGroup) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	checkGroupAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckGroup) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 7, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 8, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 9, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 10, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 13, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 15, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckGroup) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 18, checkGroupAPI(), api.MakeEventDataPtr(cb))
}

// NewCheckGroup class constructor
func NewCheckGroup(theOwner IComponent) ICheckGroup {
	r := checkGroupAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCheckGroup(r)
}

func TCheckGroupClass() types.TClass {
	r := checkGroupAPI().SysCallN(19)
	return types.TClass(r)
}

var (
	checkGroupOnce   base.Once
	checkGroupImport *imports.Imports = nil
)

func checkGroupAPI() *imports.Imports {
	checkGroupOnce.Do(func() {
		checkGroupImport = api.NewDefaultImports()
		checkGroupImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCheckGroup_Create", 0), // constructor NewCheckGroup
			/* 1 */ imports.NewTable("TCheckGroup_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TCheckGroup_DragMode", 0), // property DragMode
			/* 3 */ imports.NewTable("TCheckGroup_ParentFont", 0), // property ParentFont
			/* 4 */ imports.NewTable("TCheckGroup_ParentColor", 0), // property ParentColor
			/* 5 */ imports.NewTable("TCheckGroup_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TCheckGroup_OnDblClick", 0), // event OnDblClick
			/* 7 */ imports.NewTable("TCheckGroup_OnDragDrop", 0), // event OnDragDrop
			/* 8 */ imports.NewTable("TCheckGroup_OnDragOver", 0), // event OnDragOver
			/* 9 */ imports.NewTable("TCheckGroup_OnEndDrag", 0), // event OnEndDrag
			/* 10 */ imports.NewTable("TCheckGroup_OnMouseDown", 0), // event OnMouseDown
			/* 11 */ imports.NewTable("TCheckGroup_OnMouseEnter", 0), // event OnMouseEnter
			/* 12 */ imports.NewTable("TCheckGroup_OnMouseLeave", 0), // event OnMouseLeave
			/* 13 */ imports.NewTable("TCheckGroup_OnMouseMove", 0), // event OnMouseMove
			/* 14 */ imports.NewTable("TCheckGroup_OnMouseUp", 0), // event OnMouseUp
			/* 15 */ imports.NewTable("TCheckGroup_OnMouseWheel", 0), // event OnMouseWheel
			/* 16 */ imports.NewTable("TCheckGroup_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 17 */ imports.NewTable("TCheckGroup_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 18 */ imports.NewTable("TCheckGroup_OnStartDrag", 0), // event OnStartDrag
			/* 19 */ imports.NewTable("TCheckGroup_TClass", 0), // function TCheckGroupClass
		}
	})
	return checkGroupImport
}
