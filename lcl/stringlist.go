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

// IStringList Parent: IStrings
type IStringList interface {
	IStrings
	LoadFromFSFile(Filename string) error
	Duplicates() TDuplicates               // property
	SetDuplicates(AValue TDuplicates)      // property
	Sorted() bool                          // property
	SetSorted(AValue bool)                 // property
	CaseSensitive() bool                   // property
	SetCaseSensitive(AValue bool)          // property
	OwnsObjects() bool                     // property
	SetOwnsObjects(AValue bool)            // property
	SortStyle() TStringsSortStyle          // property
	SetSortStyle(AValue TStringsSortStyle) // property
	Find(S string, OutIndex *int32) bool   // function
	Sort()                                 // procedure
	CustomSort(fn TStringListSortCompare)  // procedure
	SetOnChange(fn TNotifyEvent)           // property event
	SetOnChanging(fn TNotifyEvent)         // property event
}

// TStringList Parent: TStrings
type TStringList struct {
	TStrings
	customSortPtr uintptr
	changePtr     uintptr
	changingPtr   uintptr
}

func NewStringList() IStringList {
	r1 := LCL().SysCallN(5269)
	return AsStringList(r1)
}

func (m *TStringList) Duplicates() TDuplicates {
	r1 := LCL().SysCallN(5271, 0, m.Instance(), 0)
	return TDuplicates(r1)
}

func (m *TStringList) SetDuplicates(AValue TDuplicates) {
	LCL().SysCallN(5271, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringList) Sorted() bool {
	r1 := LCL().SysCallN(5278, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringList) SetSorted(AValue bool) {
	LCL().SysCallN(5278, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringList) CaseSensitive() bool {
	r1 := LCL().SysCallN(5267, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringList) SetCaseSensitive(AValue bool) {
	LCL().SysCallN(5267, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringList) OwnsObjects() bool {
	r1 := LCL().SysCallN(5273, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringList) SetOwnsObjects(AValue bool) {
	LCL().SysCallN(5273, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringList) SortStyle() TStringsSortStyle {
	r1 := LCL().SysCallN(5277, 0, m.Instance(), 0)
	return TStringsSortStyle(r1)
}

func (m *TStringList) SetSortStyle(AValue TStringsSortStyle) {
	LCL().SysCallN(5277, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringList) Find(S string, OutIndex *int32) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(5272, m.Instance(), PascalStr(S), uintptr(unsafePointer(&result1)))
	*OutIndex = int32(result1)
	return GoBool(r1)
}

func StringListClass() TClass {
	ret := LCL().SysCallN(5268)
	return TClass(ret)
}

func (m *TStringList) Sort() {
	LCL().SysCallN(5276, m.Instance())
}

func (m *TStringList) CustomSort(fn TStringListSortCompare) {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5270, m.Instance(), m.customSortPtr)
}

func (m *TStringList) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5274, m.Instance(), m.changePtr)
}

func (m *TStringList) SetOnChanging(fn TNotifyEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5275, m.Instance(), m.changingPtr)
}
