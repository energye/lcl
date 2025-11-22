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

// IShape Parent: ICustomShape
type IShape interface {
	ICustomShape
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
	DragKind() types.TDragKind                      // property DragKind Getter
	SetDragKind(value types.TDragKind)              // property DragKind Setter
	DragMode() types.TDragMode                      // property DragMode Getter
	SetDragMode(value types.TDragMode)              // property DragMode Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDock(fn TEndDragEvent)                  // property event
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
	SetOnStartDock(fn TStartDockEvent)              // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

type TShape struct {
	TCustomShape
}

func (m *TShape) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := shapeAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TShape) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	shapeAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TShape) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := shapeAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TShape) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	shapeAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TShape) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := shapeAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TShape) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	shapeAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TShape) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := shapeAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TShape) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	shapeAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TShape) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 5, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 6, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 7, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 8, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 9, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 12, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 13, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 14, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 15, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 17, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 18, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 19, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 21, shapeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TShape) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 22, shapeAPI(), api.MakeEventDataPtr(cb))
}

// NewShape class constructor
func NewShape(theOwner IComponent) IShape {
	r := shapeAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsShape(r)
}

func TShapeClass() types.TClass {
	r := shapeAPI().SysCallN(23)
	return types.TClass(r)
}

var (
	shapeOnce   base.Once
	shapeImport *imports.Imports = nil
)

func shapeAPI() *imports.Imports {
	shapeOnce.Do(func() {
		shapeImport = api.NewDefaultImports()
		shapeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TShape_Create", 0), // constructor NewShape
			/* 1 */ imports.NewTable("TShape_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TShape_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TShape_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TShape_ParentShowHint", 0), // property ParentShowHint
			/* 5 */ imports.NewTable("TShape_OnDragDrop", 0), // event OnDragDrop
			/* 6 */ imports.NewTable("TShape_OnDragOver", 0), // event OnDragOver
			/* 7 */ imports.NewTable("TShape_OnEndDock", 0), // event OnEndDock
			/* 8 */ imports.NewTable("TShape_OnEndDrag", 0), // event OnEndDrag
			/* 9 */ imports.NewTable("TShape_OnMouseDown", 0), // event OnMouseDown
			/* 10 */ imports.NewTable("TShape_OnMouseEnter", 0), // event OnMouseEnter
			/* 11 */ imports.NewTable("TShape_OnMouseLeave", 0), // event OnMouseLeave
			/* 12 */ imports.NewTable("TShape_OnMouseMove", 0), // event OnMouseMove
			/* 13 */ imports.NewTable("TShape_OnMouseUp", 0), // event OnMouseUp
			/* 14 */ imports.NewTable("TShape_OnMouseWheel", 0), // event OnMouseWheel
			/* 15 */ imports.NewTable("TShape_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 16 */ imports.NewTable("TShape_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 17 */ imports.NewTable("TShape_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 18 */ imports.NewTable("TShape_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 19 */ imports.NewTable("TShape_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 20 */ imports.NewTable("TShape_OnPaint", 0), // event OnPaint
			/* 21 */ imports.NewTable("TShape_OnStartDock", 0), // event OnStartDock
			/* 22 */ imports.NewTable("TShape_OnStartDrag", 0), // event OnStartDrag
			/* 23 */ imports.NewTable("TShape_TClass", 0), // function TShapeClass
		}
	})
	return shapeImport
}
