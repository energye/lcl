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

// ICalendar Parent: ICustomCalendar
type ICalendar interface {
	ICustomCalendar
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
}

// TCalendar Parent: TCustomCalendar
type TCalendar struct {
	TCustomCalendar
	dblClickPtr        uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
}

func NewCalendar(AOwner IComponent) ICalendar {
	r1 := calendarImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCalendar(r1)
}

func CalendarClass() TClass {
	ret := calendarImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCalendar) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(2, m.Instance(), m.dblClickPtr)
}

func (m *TCalendar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(3, m.Instance(), m.mouseDownPtr)
}

func (m *TCalendar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(4, m.Instance(), m.mouseEnterPtr)
}

func (m *TCalendar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(5, m.Instance(), m.mouseLeavePtr)
}

func (m *TCalendar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(6, m.Instance(), m.mouseMovePtr)
}

func (m *TCalendar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(7, m.Instance(), m.mouseUpPtr)
}

func (m *TCalendar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(8, m.Instance(), m.mouseWheelPtr)
}

func (m *TCalendar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(9, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCalendar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(13, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCalendar) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(10, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TCalendar) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(11, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TCalendar) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	calendarImportAPI().SysCallN(12, m.Instance(), m.mouseWheelRightPtr)
}

var (
	calendarImport       *imports.Imports = nil
	calendarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Calendar_Class", 0),
		/*1*/ imports.NewTable("Calendar_Create", 0),
		/*2*/ imports.NewTable("Calendar_SetOnDblClick", 0),
		/*3*/ imports.NewTable("Calendar_SetOnMouseDown", 0),
		/*4*/ imports.NewTable("Calendar_SetOnMouseEnter", 0),
		/*5*/ imports.NewTable("Calendar_SetOnMouseLeave", 0),
		/*6*/ imports.NewTable("Calendar_SetOnMouseMove", 0),
		/*7*/ imports.NewTable("Calendar_SetOnMouseUp", 0),
		/*8*/ imports.NewTable("Calendar_SetOnMouseWheel", 0),
		/*9*/ imports.NewTable("Calendar_SetOnMouseWheelDown", 0),
		/*10*/ imports.NewTable("Calendar_SetOnMouseWheelHorz", 0),
		/*11*/ imports.NewTable("Calendar_SetOnMouseWheelLeft", 0),
		/*12*/ imports.NewTable("Calendar_SetOnMouseWheelRight", 0),
		/*13*/ imports.NewTable("Calendar_SetOnMouseWheelUp", 0),
	}
)

func calendarImportAPI() *imports.Imports {
	if calendarImport == nil {
		calendarImport = NewDefaultImports()
		calendarImport.SetImportTable(calendarImportTables)
		calendarImportTables = nil
	}
	return calendarImport
}
