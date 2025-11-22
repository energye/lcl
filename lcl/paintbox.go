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

// IPaintBox Parent: IGraphicControl
type IPaintBox interface {
	IGraphicControl
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
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
	SetOnPaint(fn TNotifyEvent)                     // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

type TPaintBox struct {
	TGraphicControl
}

func (m *TPaintBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := paintBoxAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TPaintBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	paintBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TPaintBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := paintBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TPaintBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	paintBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TPaintBox) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := paintBoxAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPaintBox) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	paintBoxAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TPaintBox) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := paintBoxAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPaintBox) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	paintBoxAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TPaintBox) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := paintBoxAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPaintBox) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	paintBoxAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TPaintBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 6, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 8, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 9, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 10, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 11, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 14, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 15, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 16, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 18, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 19, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 22, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPaintBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 23, paintBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewPaintBox class constructor
func NewPaintBox(owner IComponent) IPaintBox {
	r := paintBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsPaintBox(r)
}

func TPaintBoxClass() types.TClass {
	r := paintBoxAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	paintBoxOnce   base.Once
	paintBoxImport *imports.Imports = nil
)

func paintBoxAPI() *imports.Imports {
	paintBoxOnce.Do(func() {
		paintBoxImport = api.NewDefaultImports()
		paintBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPaintBox_Create", 0), // constructor NewPaintBox
			/* 1 */ imports.NewTable("TPaintBox_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TPaintBox_DragMode", 0), // property DragMode
			/* 3 */ imports.NewTable("TPaintBox_ParentColor", 0), // property ParentColor
			/* 4 */ imports.NewTable("TPaintBox_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TPaintBox_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TPaintBox_OnContextPopup", 0), // event OnContextPopup
			/* 7 */ imports.NewTable("TPaintBox_OnDblClick", 0), // event OnDblClick
			/* 8 */ imports.NewTable("TPaintBox_OnDragDrop", 0), // event OnDragDrop
			/* 9 */ imports.NewTable("TPaintBox_OnDragOver", 0), // event OnDragOver
			/* 10 */ imports.NewTable("TPaintBox_OnEndDrag", 0), // event OnEndDrag
			/* 11 */ imports.NewTable("TPaintBox_OnMouseDown", 0), // event OnMouseDown
			/* 12 */ imports.NewTable("TPaintBox_OnMouseEnter", 0), // event OnMouseEnter
			/* 13 */ imports.NewTable("TPaintBox_OnMouseLeave", 0), // event OnMouseLeave
			/* 14 */ imports.NewTable("TPaintBox_OnMouseMove", 0), // event OnMouseMove
			/* 15 */ imports.NewTable("TPaintBox_OnMouseUp", 0), // event OnMouseUp
			/* 16 */ imports.NewTable("TPaintBox_OnMouseWheel", 0), // event OnMouseWheel
			/* 17 */ imports.NewTable("TPaintBox_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 18 */ imports.NewTable("TPaintBox_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 19 */ imports.NewTable("TPaintBox_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 20 */ imports.NewTable("TPaintBox_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 21 */ imports.NewTable("TPaintBox_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 22 */ imports.NewTable("TPaintBox_OnPaint", 0), // event OnPaint
			/* 23 */ imports.NewTable("TPaintBox_OnStartDrag", 0), // event OnStartDrag
			/* 24 */ imports.NewTable("TPaintBox_TClass", 0), // function TPaintBoxClass
		}
	})
	return paintBoxImport
}
