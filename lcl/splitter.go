//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// ISplitter Parent: ICustomSplitter
type ISplitter interface {
	ICustomSplitter
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
}

// TSplitter Parent: TCustomSplitter
type TSplitter struct {
	TCustomSplitter
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
}

func NewSplitter(TheOwner IComponent) ISplitter {
	r1 := splitterImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsSplitter(r1)
}

func (m *TSplitter) ParentColor() bool {
	r1 := splitterImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSplitter) SetParentColor(AValue bool) {
	splitterImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSplitter) ParentShowHint() bool {
	r1 := splitterImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSplitter) SetParentShowHint(AValue bool) {
	splitterImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func SplitterClass() TClass {
	ret := splitterImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TSplitter) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	splitterImportAPI().SysCallN(4, m.Instance(), m.mouseWheelPtr)
}

func (m *TSplitter) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	splitterImportAPI().SysCallN(5, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TSplitter) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	splitterImportAPI().SysCallN(9, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TSplitter) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	splitterImportAPI().SysCallN(6, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TSplitter) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	splitterImportAPI().SysCallN(7, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TSplitter) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	splitterImportAPI().SysCallN(8, m.Instance(), m.mouseWheelRightPtr)
}

var (
	splitterImport       *imports.Imports = nil
	splitterImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Splitter_Class", 0),
		/*1*/ imports.NewTable("Splitter_Create", 0),
		/*2*/ imports.NewTable("Splitter_ParentColor", 0),
		/*3*/ imports.NewTable("Splitter_ParentShowHint", 0),
		/*4*/ imports.NewTable("Splitter_SetOnMouseWheel", 0),
		/*5*/ imports.NewTable("Splitter_SetOnMouseWheelDown", 0),
		/*6*/ imports.NewTable("Splitter_SetOnMouseWheelHorz", 0),
		/*7*/ imports.NewTable("Splitter_SetOnMouseWheelLeft", 0),
		/*8*/ imports.NewTable("Splitter_SetOnMouseWheelRight", 0),
		/*9*/ imports.NewTable("Splitter_SetOnMouseWheelUp", 0),
	}
)

func splitterImportAPI() *imports.Imports {
	if splitterImport == nil {
		splitterImport = NewDefaultImports()
		splitterImport.SetImportTable(splitterImportTables)
		splitterImportTables = nil
	}
	return splitterImport
}
