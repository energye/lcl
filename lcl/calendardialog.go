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

// ICalendarDialog Parent: IExtCommonDialog
type ICalendarDialog interface {
	IExtCommonDialog
	Date() types.TDateTime                           // property Date Getter
	SetDate(value types.TDateTime)                   // property Date Setter
	DisplaySettings() types.TDisplaySettings         // property DisplaySettings Getter
	SetDisplaySettings(value types.TDisplaySettings) // property DisplaySettings Setter
	FirstDayOfWeek() types.TCalDayOfWeek             // property FirstDayOfWeek Getter
	SetFirstDayOfWeek(value types.TCalDayOfWeek)     // property FirstDayOfWeek Setter
	OKCaption() string                               // property OKCaption Getter
	SetOKCaption(value string)                       // property OKCaption Setter
	CancelCaption() string                           // property CancelCaption Getter
	SetCancelCaption(value string)                   // property CancelCaption Setter
	SetOnDayChanged(fn TNotifyEvent)                 // property event
	SetOnMonthChanged(fn TNotifyEvent)               // property event
	SetOnYearChanged(fn TNotifyEvent)                // property event
	SetOnChange(fn TNotifyEvent)                     // property event
}

type TCalendarDialog struct {
	TExtCommonDialog
}

func (m *TCalendarDialog) Date() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	calendarDialogAPI().SysCallN(1, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCalendarDialog) SetDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	calendarDialogAPI().SysCallN(1, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCalendarDialog) DisplaySettings() types.TDisplaySettings {
	if !m.IsValid() {
		return 0
	}
	r := calendarDialogAPI().SysCallN(2, 0, m.Instance())
	return types.TDisplaySettings(r)
}

func (m *TCalendarDialog) SetDisplaySettings(value types.TDisplaySettings) {
	if !m.IsValid() {
		return
	}
	calendarDialogAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCalendarDialog) FirstDayOfWeek() types.TCalDayOfWeek {
	if !m.IsValid() {
		return 0
	}
	r := calendarDialogAPI().SysCallN(3, 0, m.Instance())
	return types.TCalDayOfWeek(r)
}

func (m *TCalendarDialog) SetFirstDayOfWeek(value types.TCalDayOfWeek) {
	if !m.IsValid() {
		return
	}
	calendarDialogAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCalendarDialog) OKCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := calendarDialogAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCalendarDialog) SetOKCaption(value string) {
	if !m.IsValid() {
		return
	}
	calendarDialogAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TCalendarDialog) CancelCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := calendarDialogAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCalendarDialog) SetCancelCaption(value string) {
	if !m.IsValid() {
		return
	}
	calendarDialogAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TCalendarDialog) SetOnDayChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, calendarDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendarDialog) SetOnMonthChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, calendarDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendarDialog) SetOnYearChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, calendarDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendarDialog) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, calendarDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewCalendarDialog class constructor
func NewCalendarDialog(owner IComponent) ICalendarDialog {
	r := calendarDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCalendarDialog(r)
}

func TCalendarDialogClass() types.TClass {
	r := calendarDialogAPI().SysCallN(10)
	return types.TClass(r)
}

var (
	calendarDialogOnce   base.Once
	calendarDialogImport *imports.Imports = nil
)

func calendarDialogAPI() *imports.Imports {
	calendarDialogOnce.Do(func() {
		calendarDialogImport = api.NewDefaultImports()
		calendarDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCalendarDialog_Create", 0), // constructor NewCalendarDialog
			/* 1 */ imports.NewTable("TCalendarDialog_Date", 0), // property Date
			/* 2 */ imports.NewTable("TCalendarDialog_DisplaySettings", 0), // property DisplaySettings
			/* 3 */ imports.NewTable("TCalendarDialog_FirstDayOfWeek", 0), // property FirstDayOfWeek
			/* 4 */ imports.NewTable("TCalendarDialog_OKCaption", 0), // property OKCaption
			/* 5 */ imports.NewTable("TCalendarDialog_CancelCaption", 0), // property CancelCaption
			/* 6 */ imports.NewTable("TCalendarDialog_OnDayChanged", 0), // event OnDayChanged
			/* 7 */ imports.NewTable("TCalendarDialog_OnMonthChanged", 0), // event OnMonthChanged
			/* 8 */ imports.NewTable("TCalendarDialog_OnYearChanged", 0), // event OnYearChanged
			/* 9 */ imports.NewTable("TCalendarDialog_OnChange", 0), // event OnChange
			/* 10 */ imports.NewTable("TCalendarDialog_TClass", 0), // function TCalendarDialogClass
		}
	})
	return calendarDialogImport
}
