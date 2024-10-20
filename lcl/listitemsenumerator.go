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

// IListItemsEnumerator Parent: IObject
type IListItemsEnumerator interface {
	IObject
	Current() IListItem // property
	MoveNext() bool     // function
}

// TListItemsEnumerator Parent: TObject
type TListItemsEnumerator struct {
	TObject
}

func NewListItemsEnumerator(AItems IListItems) IListItemsEnumerator {
	r1 := listItemsEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(AItems))
	return AsListItemsEnumerator(r1)
}

func (m *TListItemsEnumerator) Current() IListItem {
	r1 := listItemsEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsListItem(r1)
}

func (m *TListItemsEnumerator) MoveNext() bool {
	r1 := listItemsEnumeratorImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func ListItemsEnumeratorClass() TClass {
	ret := listItemsEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	listItemsEnumeratorImport       *imports.Imports = nil
	listItemsEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListItemsEnumerator_Class", 0),
		/*1*/ imports.NewTable("ListItemsEnumerator_Create", 0),
		/*2*/ imports.NewTable("ListItemsEnumerator_Current", 0),
		/*3*/ imports.NewTable("ListItemsEnumerator_MoveNext", 0),
	}
)

func listItemsEnumeratorImportAPI() *imports.Imports {
	if listItemsEnumeratorImport == nil {
		listItemsEnumeratorImport = NewDefaultImports()
		listItemsEnumeratorImport.SetImportTable(listItemsEnumeratorImportTables)
		listItemsEnumeratorImportTables = nil
	}
	return listItemsEnumeratorImport
}
