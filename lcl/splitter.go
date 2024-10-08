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
	r1 := LCL().SysCallN(5111, GetObjectUintptr(TheOwner))
	return AsSplitter(r1)
}

func (m *TSplitter) ParentColor() bool {
	r1 := LCL().SysCallN(5112, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSplitter) SetParentColor(AValue bool) {
	LCL().SysCallN(5112, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSplitter) ParentShowHint() bool {
	r1 := LCL().SysCallN(5113, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSplitter) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5113, 1, m.Instance(), PascalBool(AValue))
}

func SplitterClass() TClass {
	ret := LCL().SysCallN(5110)
	return TClass(ret)
}

func (m *TSplitter) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5114, m.Instance(), m.mouseWheelPtr)
}

func (m *TSplitter) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5115, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TSplitter) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5119, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TSplitter) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5116, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TSplitter) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5117, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TSplitter) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5118, m.Instance(), m.mouseWheelRightPtr)
}
