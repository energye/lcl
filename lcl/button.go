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

// IButton Parent: ICustomButton
type IButton interface {
	ICustomButton
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

type TButton struct {
	TCustomButton
}

func (m *TButton) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := buttonAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TButton) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	buttonAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TButton) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := buttonAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TButton) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	buttonAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TButton) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := buttonAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TButton) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	buttonAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TButton) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := buttonAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TButton) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	buttonAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TButton) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := buttonAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TButton) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	buttonAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TButton) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 6, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 7, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 8, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 9, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 10, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 13, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 15, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, buttonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TButton) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 18, buttonAPI(), api.MakeEventDataPtr(cb))
}

// NewButton class constructor
func NewButton(theOwner IComponent) IButton {
	r := buttonAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsButton(r)
}

func TButtonClass() types.TClass {
	r := buttonAPI().SysCallN(19)
	return types.TClass(r)
}

var (
	buttonOnce   base.Once
	buttonImport *imports.Imports = nil
)

func buttonAPI() *imports.Imports {
	buttonOnce.Do(func() {
		buttonImport = api.NewDefaultImports()
		buttonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TButton_Create", 0), // constructor NewButton
			/* 1 */ imports.NewTable("TButton_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TButton_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TButton_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TButton_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TButton_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TButton_OnContextPopup", 0), // event OnContextPopup
			/* 7 */ imports.NewTable("TButton_OnDragDrop", 0), // event OnDragDrop
			/* 8 */ imports.NewTable("TButton_OnDragOver", 0), // event OnDragOver
			/* 9 */ imports.NewTable("TButton_OnEndDrag", 0), // event OnEndDrag
			/* 10 */ imports.NewTable("TButton_OnMouseDown", 0), // event OnMouseDown
			/* 11 */ imports.NewTable("TButton_OnMouseEnter", 0), // event OnMouseEnter
			/* 12 */ imports.NewTable("TButton_OnMouseLeave", 0), // event OnMouseLeave
			/* 13 */ imports.NewTable("TButton_OnMouseMove", 0), // event OnMouseMove
			/* 14 */ imports.NewTable("TButton_OnMouseUp", 0), // event OnMouseUp
			/* 15 */ imports.NewTable("TButton_OnMouseWheel", 0), // event OnMouseWheel
			/* 16 */ imports.NewTable("TButton_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 17 */ imports.NewTable("TButton_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 18 */ imports.NewTable("TButton_OnStartDrag", 0), // event OnStartDrag
			/* 19 */ imports.NewTable("TButton_TClass", 0), // function TButtonClass
		}
	})
	return buttonImport
}
