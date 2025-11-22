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

type TCalendar struct {
	TCustomCalendar
}

func (m *TCalendar) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 1, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 2, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 3, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 4, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 5, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 6, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 7, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 8, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 9, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 10, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 11, calendarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCalendar) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 12, calendarAPI(), api.MakeEventDataPtr(cb))
}

// NewCalendar class constructor
func NewCalendar(owner IComponent) ICalendar {
	r := calendarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCalendar(r)
}

func TCalendarClass() types.TClass {
	r := calendarAPI().SysCallN(13)
	return types.TClass(r)
}

var (
	calendarOnce   base.Once
	calendarImport *imports.Imports = nil
)

func calendarAPI() *imports.Imports {
	calendarOnce.Do(func() {
		calendarImport = api.NewDefaultImports()
		calendarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCalendar_Create", 0), // constructor NewCalendar
			/* 1 */ imports.NewTable("TCalendar_OnDblClick", 0), // event OnDblClick
			/* 2 */ imports.NewTable("TCalendar_OnMouseDown", 0), // event OnMouseDown
			/* 3 */ imports.NewTable("TCalendar_OnMouseEnter", 0), // event OnMouseEnter
			/* 4 */ imports.NewTable("TCalendar_OnMouseLeave", 0), // event OnMouseLeave
			/* 5 */ imports.NewTable("TCalendar_OnMouseMove", 0), // event OnMouseMove
			/* 6 */ imports.NewTable("TCalendar_OnMouseUp", 0), // event OnMouseUp
			/* 7 */ imports.NewTable("TCalendar_OnMouseWheel", 0), // event OnMouseWheel
			/* 8 */ imports.NewTable("TCalendar_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 9 */ imports.NewTable("TCalendar_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 10 */ imports.NewTable("TCalendar_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 11 */ imports.NewTable("TCalendar_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 12 */ imports.NewTable("TCalendar_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 13 */ imports.NewTable("TCalendar_TClass", 0), // function TCalendarClass
		}
	})
	return calendarImport
}
