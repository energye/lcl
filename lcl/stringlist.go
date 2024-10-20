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
	r1 := stringListImportAPI().SysCallN(2)
	return AsStringList(r1)
}

func (m *TStringList) Duplicates() TDuplicates {
	r1 := stringListImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDuplicates(r1)
}

func (m *TStringList) SetDuplicates(AValue TDuplicates) {
	stringListImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringList) Sorted() bool {
	r1 := stringListImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringList) SetSorted(AValue bool) {
	stringListImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringList) CaseSensitive() bool {
	r1 := stringListImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringList) SetCaseSensitive(AValue bool) {
	stringListImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringList) OwnsObjects() bool {
	r1 := stringListImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStringList) SetOwnsObjects(AValue bool) {
	stringListImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStringList) SortStyle() TStringsSortStyle {
	r1 := stringListImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TStringsSortStyle(r1)
}

func (m *TStringList) SetSortStyle(AValue TStringsSortStyle) {
	stringListImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TStringList) Find(S string, OutIndex *int32) bool {
	var result1 uintptr
	r1 := stringListImportAPI().SysCallN(5, m.Instance(), PascalStr(S), uintptr(unsafePointer(&result1)))
	*OutIndex = int32(result1)
	return GoBool(r1)
}

func StringListClass() TClass {
	ret := stringListImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TStringList) Sort() {
	stringListImportAPI().SysCallN(9, m.Instance())
}

func (m *TStringList) CustomSort(fn TStringListSortCompare) {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	stringListImportAPI().SysCallN(3, m.Instance(), m.customSortPtr)
}

func (m *TStringList) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	stringListImportAPI().SysCallN(7, m.Instance(), m.changePtr)
}

func (m *TStringList) SetOnChanging(fn TNotifyEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	stringListImportAPI().SysCallN(8, m.Instance(), m.changingPtr)
}

var (
	stringListImport       *imports.Imports = nil
	stringListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("StringList_CaseSensitive", 0),
		/*1*/ imports.NewTable("StringList_Class", 0),
		/*2*/ imports.NewTable("StringList_Create", 0),
		/*3*/ imports.NewTable("StringList_CustomSort", 0),
		/*4*/ imports.NewTable("StringList_Duplicates", 0),
		/*5*/ imports.NewTable("StringList_Find", 0),
		/*6*/ imports.NewTable("StringList_OwnsObjects", 0),
		/*7*/ imports.NewTable("StringList_SetOnChange", 0),
		/*8*/ imports.NewTable("StringList_SetOnChanging", 0),
		/*9*/ imports.NewTable("StringList_Sort", 0),
		/*10*/ imports.NewTable("StringList_SortStyle", 0),
		/*11*/ imports.NewTable("StringList_Sorted", 0),
	}
)

func stringListImportAPI() *imports.Imports {
	if stringListImport == nil {
		stringListImport = NewDefaultImports()
		stringListImport.SetImportTable(stringListImportTables)
		stringListImportTables = nil
	}
	return stringListImport
}
