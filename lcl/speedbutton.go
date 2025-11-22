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

// ISpeedButton Parent: ICustomSpeedButton
type ISpeedButton interface {
	ICustomSpeedButton
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
	SetOnPaint(fn TNotifyEvent)                    // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

type TSpeedButton struct {
	TCustomSpeedButton
}

func (m *TSpeedButton) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := speedButtonAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TSpeedButton) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	speedButtonAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TSpeedButton) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := speedButtonAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TSpeedButton) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	speedButtonAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TSpeedButton) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := speedButtonAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TSpeedButton) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	speedButtonAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TSpeedButton) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := speedButtonAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSpeedButton) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	speedButtonAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TSpeedButton) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := speedButtonAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSpeedButton) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	speedButtonAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TSpeedButton) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 6, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 8, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 9, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 10, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 11, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 14, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 15, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 16, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 18, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpeedButton) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 20, speedButtonAPI(), api.MakeEventDataPtr(cb))
}

// NewSpeedButton class constructor
func NewSpeedButton(owner IComponent) ISpeedButton {
	r := speedButtonAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSpeedButton(r)
}

func TSpeedButtonClass() types.TClass {
	r := speedButtonAPI().SysCallN(21)
	return types.TClass(r)
}

var (
	speedButtonOnce   base.Once
	speedButtonImport *imports.Imports = nil
)

func speedButtonAPI() *imports.Imports {
	speedButtonOnce.Do(func() {
		speedButtonImport = api.NewDefaultImports()
		speedButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSpeedButton_Create", 0), // constructor NewSpeedButton
			/* 1 */ imports.NewTable("TSpeedButton_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TSpeedButton_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TSpeedButton_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TSpeedButton_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TSpeedButton_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TSpeedButton_OnContextPopup", 0), // event OnContextPopup
			/* 7 */ imports.NewTable("TSpeedButton_OnDblClick", 0), // event OnDblClick
			/* 8 */ imports.NewTable("TSpeedButton_OnDragDrop", 0), // event OnDragDrop
			/* 9 */ imports.NewTable("TSpeedButton_OnDragOver", 0), // event OnDragOver
			/* 10 */ imports.NewTable("TSpeedButton_OnEndDrag", 0), // event OnEndDrag
			/* 11 */ imports.NewTable("TSpeedButton_OnMouseDown", 0), // event OnMouseDown
			/* 12 */ imports.NewTable("TSpeedButton_OnMouseEnter", 0), // event OnMouseEnter
			/* 13 */ imports.NewTable("TSpeedButton_OnMouseLeave", 0), // event OnMouseLeave
			/* 14 */ imports.NewTable("TSpeedButton_OnMouseMove", 0), // event OnMouseMove
			/* 15 */ imports.NewTable("TSpeedButton_OnMouseUp", 0), // event OnMouseUp
			/* 16 */ imports.NewTable("TSpeedButton_OnMouseWheel", 0), // event OnMouseWheel
			/* 17 */ imports.NewTable("TSpeedButton_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 18 */ imports.NewTable("TSpeedButton_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 19 */ imports.NewTable("TSpeedButton_OnPaint", 0), // event OnPaint
			/* 20 */ imports.NewTable("TSpeedButton_OnStartDrag", 0), // event OnStartDrag
			/* 21 */ imports.NewTable("TSpeedButton_TClass", 0), // function TSpeedButtonClass
		}
	})
	return speedButtonImport
}
