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

// ICollection Parent: IPersistent
type ICollection interface {
	IPersistent
	OwnerToPersistent() IPersistent                                     // function
	AddToCollectionItem() ICollectionItem                               // function
	GetEnumerator() ICollectionEnumerator                               // function
	Insert(index int32) ICollectionItem                                 // function
	FindItemID(iD int32) ICollectionItem                                // function
	BeginUpdate()                                                       // procedure
	Clear()                                                             // procedure
	EndUpdate()                                                         // procedure
	Delete(index int32)                                                 // procedure
	Exchange(index1 int32, index2 int32)                                // procedure
	Move(index1 int32, index2 int32)                                    // procedure
	Count() int32                                                       // property Count Getter
	ItemClass() types.TCollectionItemClass                              // property ItemClass Getter
	ItemsWithIntToCollectionItem(index int32) ICollectionItem           // property Items Getter
	SetItemsWithIntToCollectionItem(index int32, value ICollectionItem) // property Items Setter
}

type TCollection struct {
	TPersistent
}

func (m *TCollection) OwnerToPersistent() IPersistent {
	if !m.IsValid() {
		return nil
	}
	r := collectionAPI().SysCallN(1, m.Instance())
	return AsPersistent(r)
}

func (m *TCollection) AddToCollectionItem() ICollectionItem {
	if !m.IsValid() {
		return nil
	}
	r := collectionAPI().SysCallN(2, m.Instance())
	return AsCollectionItem(r)
}

func (m *TCollection) GetEnumerator() ICollectionEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := collectionAPI().SysCallN(3, m.Instance())
	return AsCollectionEnumerator(r)
}

func (m *TCollection) Insert(index int32) ICollectionItem {
	if !m.IsValid() {
		return nil
	}
	r := collectionAPI().SysCallN(4, m.Instance(), uintptr(index))
	return AsCollectionItem(r)
}

func (m *TCollection) FindItemID(iD int32) ICollectionItem {
	if !m.IsValid() {
		return nil
	}
	r := collectionAPI().SysCallN(5, m.Instance(), uintptr(iD))
	return AsCollectionItem(r)
}

func (m *TCollection) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	collectionAPI().SysCallN(6, m.Instance())
}

func (m *TCollection) Clear() {
	if !m.IsValid() {
		return
	}
	collectionAPI().SysCallN(7, m.Instance())
}

func (m *TCollection) EndUpdate() {
	if !m.IsValid() {
		return
	}
	collectionAPI().SysCallN(8, m.Instance())
}

func (m *TCollection) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	collectionAPI().SysCallN(9, m.Instance(), uintptr(index))
}

func (m *TCollection) Exchange(index1 int32, index2 int32) {
	if !m.IsValid() {
		return
	}
	collectionAPI().SysCallN(10, m.Instance(), uintptr(index1), uintptr(index2))
}

func (m *TCollection) Move(index1 int32, index2 int32) {
	if !m.IsValid() {
		return
	}
	collectionAPI().SysCallN(11, m.Instance(), uintptr(index1), uintptr(index2))
}

func (m *TCollection) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := collectionAPI().SysCallN(12, m.Instance())
	return int32(r)
}

func (m *TCollection) ItemClass() types.TCollectionItemClass {
	if !m.IsValid() {
		return 0
	}
	r := collectionAPI().SysCallN(13, m.Instance())
	return types.TCollectionItemClass(r)
}

func (m *TCollection) ItemsWithIntToCollectionItem(index int32) ICollectionItem {
	if !m.IsValid() {
		return nil
	}
	r := collectionAPI().SysCallN(14, 0, m.Instance(), uintptr(index))
	return AsCollectionItem(r)
}

func (m *TCollection) SetItemsWithIntToCollectionItem(index int32, value ICollectionItem) {
	if !m.IsValid() {
		return
	}
	collectionAPI().SysCallN(14, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

// NewCollection class constructor
func NewCollection(itemClass types.TCollectionItemClass) ICollection {
	r := collectionAPI().SysCallN(0, uintptr(itemClass))
	return AsCollection(r)
}

var (
	collectionOnce   base.Once
	collectionImport *imports.Imports = nil
)

func collectionAPI() *imports.Imports {
	collectionOnce.Do(func() {
		collectionImport = api.NewDefaultImports()
		collectionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCollection_Create", 0), // constructor NewCollection
			/* 1 */ imports.NewTable("TCollection_OwnerToPersistent", 0), // function OwnerToPersistent
			/* 2 */ imports.NewTable("TCollection_AddToCollectionItem", 0), // function AddToCollectionItem
			/* 3 */ imports.NewTable("TCollection_GetEnumerator", 0), // function GetEnumerator
			/* 4 */ imports.NewTable("TCollection_Insert", 0), // function Insert
			/* 5 */ imports.NewTable("TCollection_FindItemID", 0), // function FindItemID
			/* 6 */ imports.NewTable("TCollection_BeginUpdate", 0), // procedure BeginUpdate
			/* 7 */ imports.NewTable("TCollection_Clear", 0), // procedure Clear
			/* 8 */ imports.NewTable("TCollection_EndUpdate", 0), // procedure EndUpdate
			/* 9 */ imports.NewTable("TCollection_Delete", 0), // procedure Delete
			/* 10 */ imports.NewTable("TCollection_Exchange", 0), // procedure Exchange
			/* 11 */ imports.NewTable("TCollection_Move", 0), // procedure Move
			/* 12 */ imports.NewTable("TCollection_Count", 0), // property Count
			/* 13 */ imports.NewTable("TCollection_ItemClass", 0), // property ItemClass
			/* 14 */ imports.NewTable("TCollection_ItemsWithIntToCollectionItem", 0), // property ItemsWithIntToCollectionItem
		}
	})
	return collectionImport
}
