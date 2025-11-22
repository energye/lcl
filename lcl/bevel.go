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

// IBevel Parent: IGraphicControl
type IBevel interface {
	IGraphicControl
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	Shape() types.TBevelShape                      // property Shape Getter
	SetShape(value types.TBevelShape)              // property Shape Setter
	Style() types.TBevelStyle                      // property Style Getter
	SetStyle(value types.TBevelStyle)              // property Style Setter
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnPaint(fn TNotifyEvent)                    // property event
}

type TBevel struct {
	TGraphicControl
}

func (m *TBevel) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := bevelAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBevel) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	bevelAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TBevel) Shape() types.TBevelShape {
	if !m.IsValid() {
		return 0
	}
	r := bevelAPI().SysCallN(2, 0, m.Instance())
	return types.TBevelShape(r)
}

func (m *TBevel) SetShape(value types.TBevelShape) {
	if !m.IsValid() {
		return
	}
	bevelAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TBevel) Style() types.TBevelStyle {
	if !m.IsValid() {
		return 0
	}
	r := bevelAPI().SysCallN(3, 0, m.Instance())
	return types.TBevelStyle(r)
}

func (m *TBevel) SetStyle(value types.TBevelStyle) {
	if !m.IsValid() {
		return
	}
	bevelAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TBevel) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 4, bevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBevel) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 5, bevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBevel) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, bevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBevel) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 7, bevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBevel) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 8, bevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBevel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 9, bevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBevel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 10, bevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBevel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 11, bevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBevel) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, bevelAPI(), api.MakeEventDataPtr(cb))
}

// NewBevel class constructor
func NewBevel(owner IComponent) IBevel {
	r := bevelAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsBevel(r)
}

func TBevelClass() types.TClass {
	r := bevelAPI().SysCallN(13)
	return types.TClass(r)
}

var (
	bevelOnce   base.Once
	bevelImport *imports.Imports = nil
)

func bevelAPI() *imports.Imports {
	bevelOnce.Do(func() {
		bevelImport = api.NewDefaultImports()
		bevelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBevel_Create", 0), // constructor NewBevel
			/* 1 */ imports.NewTable("TBevel_ParentShowHint", 0), // property ParentShowHint
			/* 2 */ imports.NewTable("TBevel_Shape", 0), // property Shape
			/* 3 */ imports.NewTable("TBevel_Style", 0), // property Style
			/* 4 */ imports.NewTable("TBevel_OnMouseDown", 0), // event OnMouseDown
			/* 5 */ imports.NewTable("TBevel_OnMouseEnter", 0), // event OnMouseEnter
			/* 6 */ imports.NewTable("TBevel_OnMouseLeave", 0), // event OnMouseLeave
			/* 7 */ imports.NewTable("TBevel_OnMouseMove", 0), // event OnMouseMove
			/* 8 */ imports.NewTable("TBevel_OnMouseUp", 0), // event OnMouseUp
			/* 9 */ imports.NewTable("TBevel_OnMouseWheel", 0), // event OnMouseWheel
			/* 10 */ imports.NewTable("TBevel_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 11 */ imports.NewTable("TBevel_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 12 */ imports.NewTable("TBevel_OnPaint", 0), // event OnPaint
			/* 13 */ imports.NewTable("TBevel_TClass", 0), // function TBevelClass
		}
	})
	return bevelImport
}
