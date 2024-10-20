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

// IOwnedCollection Parent: ICollection
type IOwnedCollection interface {
	ICollection
}

// TOwnedCollection Parent: TCollection
type TOwnedCollection struct {
	TCollection
}

func NewOwnedCollection(AOwner IPersistent, AItemClass TCollectionItemClass) IOwnedCollection {
	r1 := ownedCollectionImportAPI().SysCallN(1, GetObjectUintptr(AOwner), uintptr(AItemClass))
	return AsOwnedCollection(r1)
}

func OwnedCollectionClass() TClass {
	ret := ownedCollectionImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	ownedCollectionImport       *imports.Imports = nil
	ownedCollectionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("OwnedCollection_Class", 0),
		/*1*/ imports.NewTable("OwnedCollection_Create", 0),
	}
)

func ownedCollectionImportAPI() *imports.Imports {
	if ownedCollectionImport == nil {
		ownedCollectionImport = NewDefaultImports()
		ownedCollectionImport.SetImportTable(ownedCollectionImportTables)
		ownedCollectionImportTables = nil
	}
	return ownedCollectionImport
}
