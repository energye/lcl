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
)

// IListItemsEnumerator Parent: IObject
type IListItemsEnumerator interface {
	IObject
	MoveNext() bool     // function
	Current() IListItem // property Current Getter
}

type TListItemsEnumerator struct {
	TObject
}

func (m *TListItemsEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := listItemsEnumeratorAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TListItemsEnumerator) Current() IListItem {
	if !m.IsValid() {
		return nil
	}
	r := listItemsEnumeratorAPI().SysCallN(2, m.Instance())
	return AsListItem(r)
}

// NewListItemsEnumerator class constructor
func NewListItemsEnumerator(items IListItems) IListItemsEnumerator {
	r := listItemsEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(items))
	return AsListItemsEnumerator(r)
}

var (
	listItemsEnumeratorOnce   base.Once
	listItemsEnumeratorImport *imports.Imports = nil
)

func listItemsEnumeratorAPI() *imports.Imports {
	listItemsEnumeratorOnce.Do(func() {
		listItemsEnumeratorImport = api.NewDefaultImports()
		listItemsEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListItemsEnumerator_Create", 0), // constructor NewListItemsEnumerator
			/* 1 */ imports.NewTable("TListItemsEnumerator_MoveNext", 0), // function MoveNext
			/* 2 */ imports.NewTable("TListItemsEnumerator_Current", 0), // property Current
		}
	})
	return listItemsEnumeratorImport
}
