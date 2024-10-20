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

// ICollectionItem Parent: IPersistent
type ICollectionItem interface {
	IPersistent
	Collection() ICollection          // property
	SetCollection(AValue ICollection) // property
	ID() int32                        // property
	Index() int32                     // property
	SetIndex(AValue int32)            // property
	DisplayName() string              // property
	SetDisplayName(AValue string)     // property
}

// TCollectionItem Parent: TPersistent
type TCollectionItem struct {
	TPersistent
}

func NewCollectionItem(ACollection ICollection) ICollectionItem {
	r1 := collectionItemImportAPI().SysCallN(2, GetObjectUintptr(ACollection))
	return AsCollectionItem(r1)
}

func (m *TCollectionItem) Collection() ICollection {
	r1 := collectionItemImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return AsCollection(r1)
}

func (m *TCollectionItem) SetCollection(AValue ICollection) {
	collectionItemImportAPI().SysCallN(1, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCollectionItem) ID() int32 {
	r1 := collectionItemImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TCollectionItem) Index() int32 {
	r1 := collectionItemImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCollectionItem) SetIndex(AValue int32) {
	collectionItemImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCollectionItem) DisplayName() string {
	r1 := collectionItemImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCollectionItem) SetDisplayName(AValue string) {
	collectionItemImportAPI().SysCallN(3, 1, m.Instance(), PascalStr(AValue))
}

func CollectionItemClass() TClass {
	ret := collectionItemImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	collectionItemImport       *imports.Imports = nil
	collectionItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CollectionItem_Class", 0),
		/*1*/ imports.NewTable("CollectionItem_Collection", 0),
		/*2*/ imports.NewTable("CollectionItem_Create", 0),
		/*3*/ imports.NewTable("CollectionItem_DisplayName", 0),
		/*4*/ imports.NewTable("CollectionItem_ID", 0),
		/*5*/ imports.NewTable("CollectionItem_Index", 0),
	}
)

func collectionItemImportAPI() *imports.Imports {
	if collectionItemImport == nil {
		collectionItemImport = NewDefaultImports()
		collectionItemImport.SetImportTable(collectionItemImportTables)
		collectionItemImportTables = nil
	}
	return collectionItemImport
}
