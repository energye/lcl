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

// ICustomDateTimePicker Parent: ICustomControl
type ICustomDateTimePicker interface {
	ICustomControl
	DateIsNull() bool               // function
	SelectDate()                    // procedure
	SelectTime()                    // procedure
	SendExternalKey(key uint16)     // procedure
	SendExternalKeyCode(key uint16) // procedure
	Paint()                         // procedure
}

type TCustomDateTimePicker struct {
	TCustomControl
}

func (m *TCustomDateTimePicker) DateIsNull() bool {
	if !m.IsValid() {
		return false
	}
	r := customDateTimePickerAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomDateTimePicker) SelectDate() {
	if !m.IsValid() {
		return
	}
	customDateTimePickerAPI().SysCallN(2, m.Instance())
}

func (m *TCustomDateTimePicker) SelectTime() {
	if !m.IsValid() {
		return
	}
	customDateTimePickerAPI().SysCallN(3, m.Instance())
}

func (m *TCustomDateTimePicker) SendExternalKey(key uint16) {
	if !m.IsValid() {
		return
	}
	customDateTimePickerAPI().SysCallN(4, m.Instance(), uintptr(key))
}

func (m *TCustomDateTimePicker) SendExternalKeyCode(key uint16) {
	if !m.IsValid() {
		return
	}
	customDateTimePickerAPI().SysCallN(5, m.Instance(), uintptr(key))
}

func (m *TCustomDateTimePicker) Paint() {
	if !m.IsValid() {
		return
	}
	customDateTimePickerAPI().SysCallN(6, m.Instance())
}

// NewCustomDateTimePicker class constructor
func NewCustomDateTimePicker(owner IComponent) ICustomDateTimePicker {
	r := customDateTimePickerAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomDateTimePicker(r)
}

func TCustomDateTimePickerClass() types.TClass {
	r := customDateTimePickerAPI().SysCallN(7)
	return types.TClass(r)
}

var (
	customDateTimePickerOnce   base.Once
	customDateTimePickerImport *imports.Imports = nil
)

func customDateTimePickerAPI() *imports.Imports {
	customDateTimePickerOnce.Do(func() {
		customDateTimePickerImport = api.NewDefaultImports()
		customDateTimePickerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomDateTimePicker_Create", 0), // constructor NewCustomDateTimePicker
			/* 1 */ imports.NewTable("TCustomDateTimePicker_DateIsNull", 0), // function DateIsNull
			/* 2 */ imports.NewTable("TCustomDateTimePicker_SelectDate", 0), // procedure SelectDate
			/* 3 */ imports.NewTable("TCustomDateTimePicker_SelectTime", 0), // procedure SelectTime
			/* 4 */ imports.NewTable("TCustomDateTimePicker_SendExternalKey", 0), // procedure SendExternalKey
			/* 5 */ imports.NewTable("TCustomDateTimePicker_SendExternalKeyCode", 0), // procedure SendExternalKeyCode
			/* 6 */ imports.NewTable("TCustomDateTimePicker_Paint", 0), // procedure Paint
			/* 7 */ imports.NewTable("TCustomDateTimePicker_TClass", 0), // function TCustomDateTimePickerClass
		}
	})
	return customDateTimePickerImport
}
