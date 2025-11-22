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

// ICustomCalendar Parent: IWinControl
type ICustomCalendar interface {
	IWinControl
	HitTest(point types.TPoint) types.TCalendarPart  // function
	GetCalendarView() types.TCalendarView            // function
	Date() string                                    // property Date Getter
	SetDate(value string)                            // property Date Setter
	DateTime() types.TDateTime                       // property DateTime Getter
	SetDateTime(value types.TDateTime)               // property DateTime Setter
	DisplaySettings() types.TDisplaySettings         // property DisplaySettings Getter
	SetDisplaySettings(value types.TDisplaySettings) // property DisplaySettings Setter
	FirstDayOfWeek() types.TCalDayOfWeek             // property FirstDayOfWeek Getter
	SetFirstDayOfWeek(value types.TCalDayOfWeek)     // property FirstDayOfWeek Setter
	MaxDate() types.TDateTime                        // property MaxDate Getter
	SetMaxDate(value types.TDateTime)                // property MaxDate Setter
	MinDate() types.TDateTime                        // property MinDate Getter
	SetMinDate(value types.TDateTime)                // property MinDate Setter
	SetOnChange(fn TNotifyEvent)                     // property event
	SetOnDayChanged(fn TNotifyEvent)                 // property event
	SetOnMonthChanged(fn TNotifyEvent)               // property event
	SetOnYearChanged(fn TNotifyEvent)                // property event
}

type TCustomCalendar struct {
	TWinControl
}

func (m *TCustomCalendar) HitTest(point types.TPoint) types.TCalendarPart {
	if !m.IsValid() {
		return 0
	}
	r := customCalendarAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&point)))
	return types.TCalendarPart(r)
}

func (m *TCustomCalendar) GetCalendarView() types.TCalendarView {
	if !m.IsValid() {
		return 0
	}
	r := customCalendarAPI().SysCallN(2, m.Instance())
	return types.TCalendarView(r)
}

func (m *TCustomCalendar) Date() string {
	if !m.IsValid() {
		return ""
	}
	r := customCalendarAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomCalendar) SetDate(value string) {
	if !m.IsValid() {
		return
	}
	customCalendarAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomCalendar) DateTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customCalendarAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomCalendar) SetDateTime(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customCalendarAPI().SysCallN(4, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomCalendar) DisplaySettings() types.TDisplaySettings {
	if !m.IsValid() {
		return 0
	}
	r := customCalendarAPI().SysCallN(5, 0, m.Instance())
	return types.TDisplaySettings(r)
}

func (m *TCustomCalendar) SetDisplaySettings(value types.TDisplaySettings) {
	if !m.IsValid() {
		return
	}
	customCalendarAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCalendar) FirstDayOfWeek() types.TCalDayOfWeek {
	if !m.IsValid() {
		return 0
	}
	r := customCalendarAPI().SysCallN(6, 0, m.Instance())
	return types.TCalDayOfWeek(r)
}

func (m *TCustomCalendar) SetFirstDayOfWeek(value types.TCalDayOfWeek) {
	if !m.IsValid() {
		return
	}
	customCalendarAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCalendar) MaxDate() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customCalendarAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomCalendar) SetMaxDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customCalendarAPI().SysCallN(7, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomCalendar) MinDate() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customCalendarAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomCalendar) SetMinDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	customCalendarAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomCalendar) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, customCalendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomCalendar) SetOnDayChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, customCalendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomCalendar) SetOnMonthChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, customCalendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomCalendar) SetOnYearChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, customCalendarAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomCalendar class constructor
func NewCustomCalendar(owner IComponent) ICustomCalendar {
	r := customCalendarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomCalendar(r)
}

func TCustomCalendarClass() types.TClass {
	r := customCalendarAPI().SysCallN(13)
	return types.TClass(r)
}

var (
	customCalendarOnce   base.Once
	customCalendarImport *imports.Imports = nil
)

func customCalendarAPI() *imports.Imports {
	customCalendarOnce.Do(func() {
		customCalendarImport = api.NewDefaultImports()
		customCalendarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCalendar_Create", 0), // constructor NewCustomCalendar
			/* 1 */ imports.NewTable("TCustomCalendar_HitTest", 0), // function HitTest
			/* 2 */ imports.NewTable("TCustomCalendar_GetCalendarView", 0), // function GetCalendarView
			/* 3 */ imports.NewTable("TCustomCalendar_Date", 0), // property Date
			/* 4 */ imports.NewTable("TCustomCalendar_DateTime", 0), // property DateTime
			/* 5 */ imports.NewTable("TCustomCalendar_DisplaySettings", 0), // property DisplaySettings
			/* 6 */ imports.NewTable("TCustomCalendar_FirstDayOfWeek", 0), // property FirstDayOfWeek
			/* 7 */ imports.NewTable("TCustomCalendar_MaxDate", 0), // property MaxDate
			/* 8 */ imports.NewTable("TCustomCalendar_MinDate", 0), // property MinDate
			/* 9 */ imports.NewTable("TCustomCalendar_OnChange", 0), // event OnChange
			/* 10 */ imports.NewTable("TCustomCalendar_OnDayChanged", 0), // event OnDayChanged
			/* 11 */ imports.NewTable("TCustomCalendar_OnMonthChanged", 0), // event OnMonthChanged
			/* 12 */ imports.NewTable("TCustomCalendar_OnYearChanged", 0), // event OnYearChanged
			/* 13 */ imports.NewTable("TCustomCalendar_TClass", 0), // function TCustomCalendarClass
		}
	})
	return customCalendarImport
}
