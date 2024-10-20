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

// ICollectionEnumerator Parent: IObject
type ICollectionEnumerator interface {
	IObject
	Current() ICollectionItem    // property
	GetCurrent() ICollectionItem // function
	MoveNext() bool              // function
}

// TCollectionEnumerator Parent: TObject
type TCollectionEnumerator struct {
	TObject
}

func NewCollectionEnumerator(ACollection ICollection) ICollectionEnumerator {
	r1 := collectionEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(ACollection))
	return AsCollectionEnumerator(r1)
}

func (m *TCollectionEnumerator) Current() ICollectionItem {
	r1 := collectionEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsCollectionItem(r1)
}

func (m *TCollectionEnumerator) GetCurrent() ICollectionItem {
	r1 := collectionEnumeratorImportAPI().SysCallN(3, m.Instance())
	return AsCollectionItem(r1)
}

func (m *TCollectionEnumerator) MoveNext() bool {
	r1 := collectionEnumeratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func CollectionEnumeratorClass() TClass {
	ret := collectionEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	collectionEnumeratorImport       *imports.Imports = nil
	collectionEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CollectionEnumerator_Class", 0),
		/*1*/ imports.NewTable("CollectionEnumerator_Create", 0),
		/*2*/ imports.NewTable("CollectionEnumerator_Current", 0),
		/*3*/ imports.NewTable("CollectionEnumerator_GetCurrent", 0),
		/*4*/ imports.NewTable("CollectionEnumerator_MoveNext", 0),
	}
)

func collectionEnumeratorImportAPI() *imports.Imports {
	if collectionEnumeratorImport == nil {
		collectionEnumeratorImport = NewDefaultImports()
		collectionEnumeratorImport.SetImportTable(collectionEnumeratorImportTables)
		collectionEnumeratorImportTables = nil
	}
	return collectionEnumeratorImport
}
