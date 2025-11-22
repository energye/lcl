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

// ISplitter Parent: ICustomSplitter
type ISplitter interface {
	ICustomSplitter
	ParentColor() bool                              // property ParentColor Getter
	SetParentColor(value bool)                      // property ParentColor Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
}

type TSplitter struct {
	TCustomSplitter
}

func (m *TSplitter) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := splitterAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSplitter) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	splitterAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TSplitter) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := splitterAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSplitter) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	splitterAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TSplitter) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 3, splitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSplitter) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 4, splitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSplitter) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 5, splitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSplitter) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 6, splitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSplitter) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 7, splitterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSplitter) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 8, splitterAPI(), api.MakeEventDataPtr(cb))
}

// NewSplitter class constructor
func NewSplitter(theOwner IComponent) ISplitter {
	r := splitterAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsSplitter(r)
}

func TSplitterClass() types.TClass {
	r := splitterAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	splitterOnce   base.Once
	splitterImport *imports.Imports = nil
)

func splitterAPI() *imports.Imports {
	splitterOnce.Do(func() {
		splitterImport = api.NewDefaultImports()
		splitterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSplitter_Create", 0), // constructor NewSplitter
			/* 1 */ imports.NewTable("TSplitter_ParentColor", 0), // property ParentColor
			/* 2 */ imports.NewTable("TSplitter_ParentShowHint", 0), // property ParentShowHint
			/* 3 */ imports.NewTable("TSplitter_OnMouseWheel", 0), // event OnMouseWheel
			/* 4 */ imports.NewTable("TSplitter_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 5 */ imports.NewTable("TSplitter_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 6 */ imports.NewTable("TSplitter_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 7 */ imports.NewTable("TSplitter_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 8 */ imports.NewTable("TSplitter_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 9 */ imports.NewTable("TSplitter_TClass", 0), // function TSplitterClass
		}
	})
	return splitterImport
}
