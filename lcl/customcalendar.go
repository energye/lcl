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

// ICustomCalendar Parent: IWinControl
type ICustomCalendar interface {
	IWinControl
	Date() string                               // property
	SetDate(AValue string)                      // property
	DateTime() TDateTime                        // property
	SetDateTime(AValue TDateTime)               // property
	DisplaySettings() TDisplaySettings          // property
	SetDisplaySettings(AValue TDisplaySettings) // property
	FirstDayOfWeek() TCalDayOfWeek              // property
	SetFirstDayOfWeek(AValue TCalDayOfWeek)     // property
	MaxDate() TDateTime                         // property
	SetMaxDate(AValue TDateTime)                // property
	MinDate() TDateTime                         // property
	SetMinDate(AValue TDateTime)                // property
	HitTest(APoint *TPoint) TCalendarPart       // function
	GetCalendarView() TCalendarView             // function
	SetOnChange(fn TNotifyEvent)                // property event
	SetOnDayChanged(fn TNotifyEvent)            // property event
	SetOnMonthChanged(fn TNotifyEvent)          // property event
	SetOnYearChanged(fn TNotifyEvent)           // property event
}

// TCustomCalendar Parent: TWinControl
type TCustomCalendar struct {
	TWinControl
	changePtr       uintptr
	dayChangedPtr   uintptr
	monthChangedPtr uintptr
	yearChangedPtr  uintptr
}

func NewCustomCalendar(AOwner IComponent) ICustomCalendar {
	r1 := customCalendarImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomCalendar(r1)
}

func (m *TCustomCalendar) Date() string {
	r1 := customCalendarImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomCalendar) SetDate(AValue string) {
	customCalendarImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomCalendar) DateTime() TDateTime {
	r1 := customCalendarImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDateTime(r1)
}

func (m *TCustomCalendar) SetDateTime(AValue TDateTime) {
	customCalendarImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) DisplaySettings() TDisplaySettings {
	r1 := customCalendarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDisplaySettings(r1)
}

func (m *TCustomCalendar) SetDisplaySettings(AValue TDisplaySettings) {
	customCalendarImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) FirstDayOfWeek() TCalDayOfWeek {
	r1 := customCalendarImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TCalDayOfWeek(r1)
}

func (m *TCustomCalendar) SetFirstDayOfWeek(AValue TCalDayOfWeek) {
	customCalendarImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) MaxDate() TDateTime {
	r1 := customCalendarImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TDateTime(r1)
}

func (m *TCustomCalendar) SetMaxDate(AValue TDateTime) {
	customCalendarImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) MinDate() TDateTime {
	r1 := customCalendarImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TDateTime(r1)
}

func (m *TCustomCalendar) SetMinDate(AValue TDateTime) {
	customCalendarImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) HitTest(APoint *TPoint) TCalendarPart {
	r1 := customCalendarImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(APoint)))
	return TCalendarPart(r1)
}

func (m *TCustomCalendar) GetCalendarView() TCalendarView {
	r1 := customCalendarImportAPI().SysCallN(6, m.Instance())
	return TCalendarView(r1)
}

func CustomCalendarClass() TClass {
	ret := customCalendarImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomCalendar) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	customCalendarImportAPI().SysCallN(10, m.Instance(), m.changePtr)
}

func (m *TCustomCalendar) SetOnDayChanged(fn TNotifyEvent) {
	if m.dayChangedPtr != 0 {
		RemoveEventElement(m.dayChangedPtr)
	}
	m.dayChangedPtr = MakeEventDataPtr(fn)
	customCalendarImportAPI().SysCallN(11, m.Instance(), m.dayChangedPtr)
}

func (m *TCustomCalendar) SetOnMonthChanged(fn TNotifyEvent) {
	if m.monthChangedPtr != 0 {
		RemoveEventElement(m.monthChangedPtr)
	}
	m.monthChangedPtr = MakeEventDataPtr(fn)
	customCalendarImportAPI().SysCallN(12, m.Instance(), m.monthChangedPtr)
}

func (m *TCustomCalendar) SetOnYearChanged(fn TNotifyEvent) {
	if m.yearChangedPtr != 0 {
		RemoveEventElement(m.yearChangedPtr)
	}
	m.yearChangedPtr = MakeEventDataPtr(fn)
	customCalendarImportAPI().SysCallN(13, m.Instance(), m.yearChangedPtr)
}

var (
	customCalendarImport       *imports.Imports = nil
	customCalendarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCalendar_Class", 0),
		/*1*/ imports.NewTable("CustomCalendar_Create", 0),
		/*2*/ imports.NewTable("CustomCalendar_Date", 0),
		/*3*/ imports.NewTable("CustomCalendar_DateTime", 0),
		/*4*/ imports.NewTable("CustomCalendar_DisplaySettings", 0),
		/*5*/ imports.NewTable("CustomCalendar_FirstDayOfWeek", 0),
		/*6*/ imports.NewTable("CustomCalendar_GetCalendarView", 0),
		/*7*/ imports.NewTable("CustomCalendar_HitTest", 0),
		/*8*/ imports.NewTable("CustomCalendar_MaxDate", 0),
		/*9*/ imports.NewTable("CustomCalendar_MinDate", 0),
		/*10*/ imports.NewTable("CustomCalendar_SetOnChange", 0),
		/*11*/ imports.NewTable("CustomCalendar_SetOnDayChanged", 0),
		/*12*/ imports.NewTable("CustomCalendar_SetOnMonthChanged", 0),
		/*13*/ imports.NewTable("CustomCalendar_SetOnYearChanged", 0),
	}
)

func customCalendarImportAPI() *imports.Imports {
	if customCalendarImport == nil {
		customCalendarImport = NewDefaultImports()
		customCalendarImport.SetImportTable(customCalendarImportTables)
		customCalendarImportTables = nil
	}
	return customCalendarImport
}
