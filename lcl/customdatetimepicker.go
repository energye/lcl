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

// ICustomDateTimePicker Parent: ICustomControl
type ICustomDateTimePicker interface {
	ICustomControl
	DateIsNull() bool             // function
	SelectDate()                  // procedure
	SelectTime()                  // procedure
	SendExternalKey(aKey Char)    // procedure
	SendExternalKeyCode(Key Word) // procedure
	Paint()                       // procedure
}

// TCustomDateTimePicker Parent: TCustomControl
type TCustomDateTimePicker struct {
	TCustomControl
}

func NewCustomDateTimePicker(AOwner IComponent) ICustomDateTimePicker {
	r1 := customDateTimePickerImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomDateTimePicker(r1)
}

func (m *TCustomDateTimePicker) DateIsNull() bool {
	r1 := customDateTimePickerImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func CustomDateTimePickerClass() TClass {
	ret := customDateTimePickerImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomDateTimePicker) SelectDate() {
	customDateTimePickerImportAPI().SysCallN(4, m.Instance())
}

func (m *TCustomDateTimePicker) SelectTime() {
	customDateTimePickerImportAPI().SysCallN(5, m.Instance())
}

func (m *TCustomDateTimePicker) SendExternalKey(aKey Char) {
	customDateTimePickerImportAPI().SysCallN(6, m.Instance(), uintptr(aKey))
}

func (m *TCustomDateTimePicker) SendExternalKeyCode(Key Word) {
	customDateTimePickerImportAPI().SysCallN(7, m.Instance(), uintptr(Key))
}

func (m *TCustomDateTimePicker) Paint() {
	customDateTimePickerImportAPI().SysCallN(3, m.Instance())
}

var (
	customDateTimePickerImport       *imports.Imports = nil
	customDateTimePickerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomDateTimePicker_Class", 0),
		/*1*/ imports.NewTable("CustomDateTimePicker_Create", 0),
		/*2*/ imports.NewTable("CustomDateTimePicker_DateIsNull", 0),
		/*3*/ imports.NewTable("CustomDateTimePicker_Paint", 0),
		/*4*/ imports.NewTable("CustomDateTimePicker_SelectDate", 0),
		/*5*/ imports.NewTable("CustomDateTimePicker_SelectTime", 0),
		/*6*/ imports.NewTable("CustomDateTimePicker_SendExternalKey", 0),
		/*7*/ imports.NewTable("CustomDateTimePicker_SendExternalKeyCode", 0),
	}
)

func customDateTimePickerImportAPI() *imports.Imports {
	if customDateTimePickerImport == nil {
		customDateTimePickerImport = NewDefaultImports()
		customDateTimePickerImport.SetImportTable(customDateTimePickerImportTables)
		customDateTimePickerImportTables = nil
	}
	return customDateTimePickerImport
}
