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

// ICustomTimer Parent: IComponent
type ICustomTimer interface {
	IComponent
	Enabled() bool                   // property Enabled Getter
	SetEnabled(value bool)           // property Enabled Setter
	Interval() uint32                // property Interval Getter
	SetInterval(value uint32)        // property Interval Setter
	SetOnTimer(fn TNotifyEvent)      // property event
	SetOnStartTimer(fn TNotifyEvent) // property event
	SetOnStopTimer(fn TNotifyEvent)  // property event
}

type TCustomTimer struct {
	TComponent
}

func (m *TCustomTimer) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := customTimerAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTimer) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	customTimerAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTimer) Interval() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customTimerAPI().SysCallN(2, 0, m.Instance())
	return uint32(r)
}

func (m *TCustomTimer) SetInterval(value uint32) {
	if !m.IsValid() {
		return
	}
	customTimerAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTimer) SetOnTimer(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 3, customTimerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTimer) SetOnStartTimer(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 4, customTimerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTimer) SetOnStopTimer(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 5, customTimerAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomTimer class constructor
func NewCustomTimer(owner IComponent) ICustomTimer {
	r := customTimerAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomTimer(r)
}

func TCustomTimerClass() types.TClass {
	r := customTimerAPI().SysCallN(6)
	return types.TClass(r)
}

var (
	customTimerOnce   base.Once
	customTimerImport *imports.Imports = nil
)

func customTimerAPI() *imports.Imports {
	customTimerOnce.Do(func() {
		customTimerImport = api.NewDefaultImports()
		customTimerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomTimer_Create", 0), // constructor NewCustomTimer
			/* 1 */ imports.NewTable("TCustomTimer_Enabled", 0), // property Enabled
			/* 2 */ imports.NewTable("TCustomTimer_Interval", 0), // property Interval
			/* 3 */ imports.NewTable("TCustomTimer_OnTimer", 0), // event OnTimer
			/* 4 */ imports.NewTable("TCustomTimer_OnStartTimer", 0), // event OnStartTimer
			/* 5 */ imports.NewTable("TCustomTimer_OnStopTimer", 0), // event OnStopTimer
			/* 6 */ imports.NewTable("TCustomTimer_TClass", 0), // function TCustomTimerClass
		}
	})
	return customTimerImport
}
