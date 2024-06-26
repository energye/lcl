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

// IGraphicsObject Parent: IPersistent
type IGraphicsObject interface {
	IPersistent
	SetOnChanging(fn TNotifyEvent) // property event
	SetOnChange(fn TNotifyEvent)   // property event
}

// TGraphicsObject Parent: TPersistent
type TGraphicsObject struct {
	TPersistent
	changingPtr uintptr
	changePtr   uintptr
}

func NewGraphicsObject() IGraphicsObject {
	r1 := LCL().SysCallN(3203)
	return AsGraphicsObject(r1)
}

func GraphicsObjectClass() TClass {
	ret := LCL().SysCallN(3202)
	return TClass(ret)
}

func (m *TGraphicsObject) SetOnChanging(fn TNotifyEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3205, m.Instance(), m.changingPtr)
}

func (m *TGraphicsObject) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3204, m.Instance(), m.changePtr)
}
