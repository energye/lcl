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

// IPairSplitterSide Parent: IWinControl
type IPairSplitterSide interface {
	IWinControl
	Splitter() ICustomPairSplitter                 // property Splitter Getter
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

type TPairSplitterSide struct {
	TWinControl
}

func (m *TPairSplitterSide) Splitter() ICustomPairSplitter {
	if !m.IsValid() {
		return nil
	}
	r := pairSplitterSideAPI().SysCallN(1, m.Instance())
	return AsCustomPairSplitter(r)
}

func (m *TPairSplitterSide) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := pairSplitterSideAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPairSplitterSide) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	pairSplitterSideAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TPairSplitterSide) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 3, pairSplitterSideAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitterSide) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 4, pairSplitterSideAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitterSide) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 5, pairSplitterSideAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitterSide) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 6, pairSplitterSideAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitterSide) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 7, pairSplitterSideAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitterSide) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 8, pairSplitterSideAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitterSide) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 9, pairSplitterSideAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPairSplitterSide) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 10, pairSplitterSideAPI(), api.MakeEventDataPtr(cb))
}

// NewPairSplitterSide class constructor
func NewPairSplitterSide(theOwner IComponent) IPairSplitterSide {
	r := pairSplitterSideAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsPairSplitterSide(r)
}

func TPairSplitterSideClass() types.TClass {
	r := pairSplitterSideAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	pairSplitterSideOnce   base.Once
	pairSplitterSideImport *imports.Imports = nil
)

func pairSplitterSideAPI() *imports.Imports {
	pairSplitterSideOnce.Do(func() {
		pairSplitterSideImport = api.NewDefaultImports()
		pairSplitterSideImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPairSplitterSide_Create", 0), // constructor NewPairSplitterSide
			/* 1 */ imports.NewTable("TPairSplitterSide_Splitter", 0), // property Splitter
			/* 2 */ imports.NewTable("TPairSplitterSide_ParentShowHint", 0), // property ParentShowHint
			/* 3 */ imports.NewTable("TPairSplitterSide_OnMouseDown", 0), // event OnMouseDown
			/* 4 */ imports.NewTable("TPairSplitterSide_OnMouseEnter", 0), // event OnMouseEnter
			/* 5 */ imports.NewTable("TPairSplitterSide_OnMouseLeave", 0), // event OnMouseLeave
			/* 6 */ imports.NewTable("TPairSplitterSide_OnMouseMove", 0), // event OnMouseMove
			/* 7 */ imports.NewTable("TPairSplitterSide_OnMouseUp", 0), // event OnMouseUp
			/* 8 */ imports.NewTable("TPairSplitterSide_OnMouseWheel", 0), // event OnMouseWheel
			/* 9 */ imports.NewTable("TPairSplitterSide_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 10 */ imports.NewTable("TPairSplitterSide_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 11 */ imports.NewTable("TPairSplitterSide_TClass", 0), // function TPairSplitterSideClass
		}
	})
	return pairSplitterSideImport
}
