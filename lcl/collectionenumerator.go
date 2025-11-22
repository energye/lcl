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

// ICollectionEnumerator Parent: IObject
type ICollectionEnumerator interface {
	IObject
	GetCurrent() ICollectionItem // function
	MoveNext() bool              // function
	Current() ICollectionItem    // property Current Getter
}

type TCollectionEnumerator struct {
	TObject
}

func (m *TCollectionEnumerator) GetCurrent() ICollectionItem {
	if !m.IsValid() {
		return nil
	}
	r := collectionEnumeratorAPI().SysCallN(1, m.Instance())
	return AsCollectionItem(r)
}

func (m *TCollectionEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := collectionEnumeratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCollectionEnumerator) Current() ICollectionItem {
	if !m.IsValid() {
		return nil
	}
	r := collectionEnumeratorAPI().SysCallN(3, m.Instance())
	return AsCollectionItem(r)
}

// NewCollectionEnumerator class constructor
func NewCollectionEnumerator(collection ICollection) ICollectionEnumerator {
	r := collectionEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsCollectionEnumerator(r)
}

var (
	collectionEnumeratorOnce   base.Once
	collectionEnumeratorImport *imports.Imports = nil
)

func collectionEnumeratorAPI() *imports.Imports {
	collectionEnumeratorOnce.Do(func() {
		collectionEnumeratorImport = api.NewDefaultImports()
		collectionEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCollectionEnumerator_Create", 0), // constructor NewCollectionEnumerator
			/* 1 */ imports.NewTable("TCollectionEnumerator_GetCurrent", 0), // function GetCurrent
			/* 2 */ imports.NewTable("TCollectionEnumerator_MoveNext", 0), // function MoveNext
			/* 3 */ imports.NewTable("TCollectionEnumerator_Current", 0), // property Current
		}
	})
	return collectionEnumeratorImport
}
