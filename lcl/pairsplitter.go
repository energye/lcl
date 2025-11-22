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

// IPairSplitter Parent: ICustomPairSplitter
type IPairSplitter interface {
	ICustomPairSplitter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
}

type TPairSplitter struct {
	TCustomPairSplitter
}

func (m *TPairSplitter) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := pairSplitterAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPairSplitter) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	pairSplitterAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TPairSplitter) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 2, pairSplitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitter) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 3, pairSplitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitter) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 4, pairSplitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitter) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 5, pairSplitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitter) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 6, pairSplitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitter) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 7, pairSplitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitter) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 8, pairSplitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitter) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 9, pairSplitterAPI(), api.MakeEventDataPtr(cb))
}

// NewPairSplitter class constructor
func NewPairSplitter(theOwner IComponent) IPairSplitter {
	r := pairSplitterAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsPairSplitter(r)
}

func TPairSplitterClass() types.TClass {
	r := pairSplitterAPI().SysCallN(10)
	return types.TClass(r)
}

var (
	pairSplitterOnce   base.Once
	pairSplitterImport *imports.Imports = nil
)

func pairSplitterAPI() *imports.Imports {
	pairSplitterOnce.Do(func() {
		pairSplitterImport = api.NewDefaultImports()
		pairSplitterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPairSplitter_Create", 0), // constructor NewPairSplitter
			/* 1 */ imports.NewTable("TPairSplitter_ParentShowHint", 0), // property ParentShowHint
			/* 2 */ imports.NewTable("TPairSplitter_OnMouseDown", 0), // event OnMouseDown
			/* 3 */ imports.NewTable("TPairSplitter_OnMouseEnter", 0), // event OnMouseEnter
			/* 4 */ imports.NewTable("TPairSplitter_OnMouseLeave", 0), // event OnMouseLeave
			/* 5 */ imports.NewTable("TPairSplitter_OnMouseMove", 0), // event OnMouseMove
			/* 6 */ imports.NewTable("TPairSplitter_OnMouseUp", 0), // event OnMouseUp
			/* 7 */ imports.NewTable("TPairSplitter_OnMouseWheel", 0), // event OnMouseWheel
			/* 8 */ imports.NewTable("TPairSplitter_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 9 */ imports.NewTable("TPairSplitter_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 10 */ imports.NewTable("TPairSplitter_TClass", 0), // function TPairSplitterClass
		}
	})
	return pairSplitterImport
}
