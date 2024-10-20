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
	r1 := customTimerImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomTimer(r1)
}

func (m *TCustomTimer) Enabled() bool {
	r1 := customTimerImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTimer) SetEnabled(AValue bool) {
	customTimerImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTimer) Interval() uint32 {
	r1 := customTimerImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomTimer) SetInterval(AValue uint32) {
	customTimerImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func CustomTimerClass() TClass {
	ret := customTimerImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomTimer) SetOnTimer(fn TNotifyEvent) {
	if m.timerPtr != 0 {
		RemoveEventElement(m.timerPtr)
	}
	m.timerPtr = MakeEventDataPtr(fn)
	customTimerImportAPI().SysCallN(6, m.Instance(), m.timerPtr)
}

func (m *TCustomTimer) SetOnStartTimer(fn TNotifyEvent) {
	if m.startTimerPtr != 0 {
		RemoveEventElement(m.startTimerPtr)
	}
	m.startTimerPtr = MakeEventDataPtr(fn)
	customTimerImportAPI().SysCallN(4, m.Instance(), m.startTimerPtr)
}

func (m *TCustomTimer) SetOnStopTimer(fn TNotifyEvent) {
	if m.stopTimerPtr != 0 {
		RemoveEventElement(m.stopTimerPtr)
	}
	m.stopTimerPtr = MakeEventDataPtr(fn)
	customTimerImportAPI().SysCallN(5, m.Instance(), m.stopTimerPtr)
}

var (
	customTimerImport       *imports.Imports = nil
	customTimerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomTimer_Class", 0),
		/*1*/ imports.NewTable("CustomTimer_Create", 0),
		/*2*/ imports.NewTable("CustomTimer_Enabled", 0),
		/*3*/ imports.NewTable("CustomTimer_Interval", 0),
		/*4*/ imports.NewTable("CustomTimer_SetOnStartTimer", 0),
		/*5*/ imports.NewTable("CustomTimer_SetOnStopTimer", 0),
		/*6*/ imports.NewTable("CustomTimer_SetOnTimer", 0),
	}
)

func customTimerImportAPI() *imports.Imports {
	if customTimerImport == nil {
		customTimerImport = NewDefaultImports()
		customTimerImport.SetImportTable(customTimerImportTables)
		customTimerImportTables = nil
	}
	return customTimerImport
}
