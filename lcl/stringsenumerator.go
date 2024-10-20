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

// IStringsEnumerator Parent: IObject
type IStringsEnumerator interface {
	IObject
	Current() string    // property
	GetCurrent() string // function
	MoveNext() bool     // function
}

// TStringsEnumerator Parent: TObject
type TStringsEnumerator struct {
	TObject
}

func NewStringsEnumerator(AStrings IStrings) IStringsEnumerator {
	r1 := stringsEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(AStrings))
	return AsStringsEnumerator(r1)
}

func (m *TStringsEnumerator) Current() string {
	r1 := stringsEnumeratorImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TStringsEnumerator) GetCurrent() string {
	r1 := stringsEnumeratorImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TStringsEnumerator) MoveNext() bool {
	r1 := stringsEnumeratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func StringsEnumeratorClass() TClass {
	ret := stringsEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	stringsEnumeratorImport       *imports.Imports = nil
	stringsEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("StringsEnumerator_Class", 0),
		/*1*/ imports.NewTable("StringsEnumerator_Create", 0),
		/*2*/ imports.NewTable("StringsEnumerator_Current", 0),
		/*3*/ imports.NewTable("StringsEnumerator_GetCurrent", 0),
		/*4*/ imports.NewTable("StringsEnumerator_MoveNext", 0),
	}
)

func stringsEnumeratorImportAPI() *imports.Imports {
	if stringsEnumeratorImport == nil {
		stringsEnumeratorImport = NewDefaultImports()
		stringsEnumeratorImport.SetImportTable(stringsEnumeratorImportTables)
		stringsEnumeratorImportTables = nil
	}
	return stringsEnumeratorImport
}
