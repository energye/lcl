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
	"github.com/energye/lcl/types"
)

// IOwnedCollection Parent: ICollection
type IOwnedCollection interface {
	ICollection
}

type TOwnedCollection struct {
	TCollection
}

// NewOwnedCollection class constructor
func NewOwnedCollection(owner IPersistent, itemClass types.TCollectionItemClass) IOwnedCollection {
	r := ownedCollectionAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(itemClass))
	return AsOwnedCollection(r)
}

var (
	ownedCollectionOnce   base.Once
	ownedCollectionImport *imports.Imports = nil
)

func ownedCollectionAPI() *imports.Imports {
	ownedCollectionOnce.Do(func() {
		ownedCollectionImport = api.NewDefaultImports()
		ownedCollectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TOwnedCollection_Create", 0), // constructor NewOwnedCollection
		}
	})
	return ownedCollectionImport
}
