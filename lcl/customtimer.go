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

// ICustomTimer Parent: IComponent
type ICustomTimer interface {
	IComponent
	Enabled() bool                   // property
	SetEnabled(AValue bool)          // property
	Interval() uint32                // property
	SetInterval(AValue uint32)       // property
	SetOnTimer(fn TNotifyEvent)      // property event
	SetOnStartTimer(fn TNotifyEvent) // property event
	SetOnStopTimer(fn TNotifyEvent)  // property event
}

// TCustomTimer Parent: TComponent
type TCustomTimer struct {
	TComponent
	timerPtr      uintptr
	startTimerPtr uintptr
	stopTimerPtr  uintptr
}

func NewCustomTimer(AOwner IComponent) ICustomTimer {
	r1 := LCL().SysCallN(2380, GetObjectUintptr(AOwner))
	return AsCustomTimer(r1)
}

func (m *TCustomTimer) Enabled() bool {
	r1 := LCL().SysCallN(2381, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTimer) SetEnabled(AValue bool) {
	LCL().SysCallN(2381, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTimer) Interval() uint32 {
	r1 := LCL().SysCallN(2382, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomTimer) SetInterval(AValue uint32) {
	LCL().SysCallN(2382, 1, m.Instance(), uintptr(AValue))
}

func CustomTimerClass() TClass {
	ret := LCL().SysCallN(2379)
	return TClass(ret)
}

func (m *TCustomTimer) SetOnTimer(fn TNotifyEvent) {
	if m.timerPtr != 0 {
		RemoveEventElement(m.timerPtr)
	}
	m.timerPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2385, m.Instance(), m.timerPtr)
}

func (m *TCustomTimer) SetOnStartTimer(fn TNotifyEvent) {
	if m.startTimerPtr != 0 {
		RemoveEventElement(m.startTimerPtr)
	}
	m.startTimerPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2383, m.Instance(), m.startTimerPtr)
}

func (m *TCustomTimer) SetOnStopTimer(fn TNotifyEvent) {
	if m.stopTimerPtr != 0 {
		RemoveEventElement(m.stopTimerPtr)
	}
	m.stopTimerPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2384, m.Instance(), m.stopTimerPtr)
}
