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

// IListEnumerator Parent: IObject
type IListEnumerator interface {
	IObject
	GetCurrent() uintptr // function
	MoveNext() bool      // function
	Current() uintptr    // property Current Getter
}

type TListEnumerator struct {
	TObject
}

func (m *TListEnumerator) GetCurrent() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := listEnumeratorAPI().SysCallN(1, m.Instance())
	return uintptr(r)
}

func (m *TListEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := listEnumeratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TListEnumerator) Current() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := listEnumeratorAPI().SysCallN(3, m.Instance())
	return uintptr(r)
}

// NewListEnumerator class constructor
func NewListEnumerator(list IList) IListEnumerator {
	r := listEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(list))
	return AsListEnumerator(r)
}

var (
	listEnumeratorOnce   base.Once
	listEnumeratorImport *imports.Imports = nil
)

func listEnumeratorAPI() *imports.Imports {
	listEnumeratorOnce.Do(func() {
		listEnumeratorImport = api.NewDefaultImports()
		listEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListEnumerator_Create", 0), // constructor NewListEnumerator
			/* 1 */ imports.NewTable("TListEnumerator_GetCurrent", 0), // function GetCurrent
			/* 2 */ imports.NewTable("TListEnumerator_MoveNext", 0), // function MoveNext
			/* 3 */ imports.NewTable("TListEnumerator_Current", 0), // property Current
		}
	})
	return listEnumeratorImport
}
