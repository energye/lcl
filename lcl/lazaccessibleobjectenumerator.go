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

// ILazAccessibleObjectEnumerator Parent: IAVLTreeNodeEnumerator
type ILazAccessibleObjectEnumerator interface {
	IAVLTreeNodeEnumerator
	CurrentForLazAccessibleObject() ILazAccessibleObject // property
}

// TLazAccessibleObjectEnumerator Parent: TAVLTreeNodeEnumerator
type TLazAccessibleObjectEnumerator struct {
	TAVLTreeNodeEnumerator
}

func NewLazAccessibleObjectEnumerator(Tree IAVLTree, aLowToHigh bool) ILazAccessibleObjectEnumerator {
	r1 := lazAccessibleObjectEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(Tree), PascalBool(aLowToHigh))
	return AsLazAccessibleObjectEnumerator(r1)
}

func (m *TLazAccessibleObjectEnumerator) CurrentForLazAccessibleObject() ILazAccessibleObject {
	r1 := lazAccessibleObjectEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsLazAccessibleObject(r1)
}

func LazAccessibleObjectEnumeratorClass() TClass {
	ret := lazAccessibleObjectEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	lazAccessibleObjectEnumeratorImport       *imports.Imports = nil
	lazAccessibleObjectEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazAccessibleObjectEnumerator_Class", 0),
		/*1*/ imports.NewTable("LazAccessibleObjectEnumerator_Create", 0),
		/*2*/ imports.NewTable("LazAccessibleObjectEnumerator_CurrentForLazAccessibleObject", 0),
	}
)

func lazAccessibleObjectEnumeratorImportAPI() *imports.Imports {
	if lazAccessibleObjectEnumeratorImport == nil {
		lazAccessibleObjectEnumeratorImport = NewDefaultImports()
		lazAccessibleObjectEnumeratorImport.SetImportTable(lazAccessibleObjectEnumeratorImportTables)
		lazAccessibleObjectEnumeratorImportTables = nil
	}
	return lazAccessibleObjectEnumeratorImport
}
