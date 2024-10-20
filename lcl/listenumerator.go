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

// IListEnumerator Parent: IObject
type IListEnumerator interface {
	IObject
	Current() uintptr    // property
	GetCurrent() uintptr // function
	MoveNext() bool      // function
}

// TListEnumerator Parent: TObject
type TListEnumerator struct {
	TObject
}

func NewListEnumerator(AList IList) IListEnumerator {
	r1 := listEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(AList))
	return AsListEnumerator(r1)
}

func (m *TListEnumerator) Current() uintptr {
	r1 := listEnumeratorImportAPI().SysCallN(2, m.Instance())
	return uintptr(r1)
}

func (m *TListEnumerator) GetCurrent() uintptr {
	r1 := listEnumeratorImportAPI().SysCallN(3, m.Instance())
	return uintptr(r1)
}

func (m *TListEnumerator) MoveNext() bool {
	r1 := listEnumeratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func ListEnumeratorClass() TClass {
	ret := listEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	listEnumeratorImport       *imports.Imports = nil
	listEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListEnumerator_Class", 0),
		/*1*/ imports.NewTable("ListEnumerator_Create", 0),
		/*2*/ imports.NewTable("ListEnumerator_Current", 0),
		/*3*/ imports.NewTable("ListEnumerator_GetCurrent", 0),
		/*4*/ imports.NewTable("ListEnumerator_MoveNext", 0),
	}
)

func listEnumeratorImportAPI() *imports.Imports {
	if listEnumeratorImport == nil {
		listEnumeratorImport = NewDefaultImports()
		listEnumeratorImport.SetImportTable(listEnumeratorImportTables)
		listEnumeratorImportTables = nil
	}
	return listEnumeratorImport
}
