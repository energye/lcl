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

// ICollection Parent: IPersistent
type ICollection interface {
	IPersistent
	Count() int32                                 // property
	ItemClass() TCollectionItemClass              // property
	Items(Index int32) ICollectionItem            // property
	SetItems(Index int32, AValue ICollectionItem) // property
	Owner() IPersistent                           // function
	Add() ICollectionItem                         // function
	GetEnumerator() ICollectionEnumerator         // function
	Insert(Index int32) ICollectionItem           // function
	FindItemID(ID int32) ICollectionItem          // function
	BeginUpdate()                                 // procedure
	Clear()                                       // procedure
	EndUpdate()                                   // procedure
	Delete(Index int32)                           // procedure
	Exchange(Index1, index2 int32)                // procedure
	Move(Index1, index2 int32)                    // procedure
	Sort(fn TCollectionSortCompare)               // procedure
}

// TCollection Parent: TPersistent
type TCollection struct {
	TPersistent
	sortPtr uintptr
}

func NewCollection(AItemClass TCollectionItemClass) ICollection {
	r1 := collectionImportAPI().SysCallN(5, uintptr(AItemClass))
	return AsCollection(r1)
}

func (m *TCollection) Count() int32 {
	r1 := collectionImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TCollection) ItemClass() TCollectionItemClass {
	r1 := collectionImportAPI().SysCallN(12, m.Instance())
	return TCollectionItemClass(r1)
}

func (m *TCollection) Items(Index int32) ICollectionItem {
	r1 := collectionImportAPI().SysCallN(13, 0, m.Instance(), uintptr(Index))
	return AsCollectionItem(r1)
}

func (m *TCollection) SetItems(Index int32, AValue ICollectionItem) {
	collectionImportAPI().SysCallN(13, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TCollection) Owner() IPersistent {
	r1 := collectionImportAPI().SysCallN(15, m.Instance())
	return AsPersistent(r1)
}

func (m *TCollection) Add() ICollectionItem {
	r1 := collectionImportAPI().SysCallN(0, m.Instance())
	return AsCollectionItem(r1)
}

func (m *TCollection) GetEnumerator() ICollectionEnumerator {
	r1 := collectionImportAPI().SysCallN(10, m.Instance())
	return AsCollectionEnumerator(r1)
}

func (m *TCollection) Insert(Index int32) ICollectionItem {
	r1 := collectionImportAPI().SysCallN(11, m.Instance(), uintptr(Index))
	return AsCollectionItem(r1)
}

func (m *TCollection) FindItemID(ID int32) ICollectionItem {
	r1 := collectionImportAPI().SysCallN(9, m.Instance(), uintptr(ID))
	return AsCollectionItem(r1)
}

func CollectionClass() TClass {
	ret := collectionImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCollection) BeginUpdate() {
	collectionImportAPI().SysCallN(1, m.Instance())
}

func (m *TCollection) Clear() {
	collectionImportAPI().SysCallN(3, m.Instance())
}

func (m *TCollection) EndUpdate() {
	collectionImportAPI().SysCallN(7, m.Instance())
}

func (m *TCollection) Delete(Index int32) {
	collectionImportAPI().SysCallN(6, m.Instance(), uintptr(Index))
}

func (m *TCollection) Exchange(Index1, index2 int32) {
	collectionImportAPI().SysCallN(8, m.Instance(), uintptr(Index1), uintptr(index2))
}

func (m *TCollection) Move(Index1, index2 int32) {
	collectionImportAPI().SysCallN(14, m.Instance(), uintptr(Index1), uintptr(index2))
}

func (m *TCollection) Sort(fn TCollectionSortCompare) {
	if m.sortPtr != 0 {
		RemoveEventElement(m.sortPtr)
	}
	m.sortPtr = MakeEventDataPtr(fn)
	collectionImportAPI().SysCallN(16, m.Instance(), m.sortPtr)
}

var (
	collectionImport       *imports.Imports = nil
	collectionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Collection_Add", 0),
		/*1*/ imports.NewTable("Collection_BeginUpdate", 0),
		/*2*/ imports.NewTable("Collection_Class", 0),
		/*3*/ imports.NewTable("Collection_Clear", 0),
		/*4*/ imports.NewTable("Collection_Count", 0),
		/*5*/ imports.NewTable("Collection_Create", 0),
		/*6*/ imports.NewTable("Collection_Delete", 0),
		/*7*/ imports.NewTable("Collection_EndUpdate", 0),
		/*8*/ imports.NewTable("Collection_Exchange", 0),
		/*9*/ imports.NewTable("Collection_FindItemID", 0),
		/*10*/ imports.NewTable("Collection_GetEnumerator", 0),
		/*11*/ imports.NewTable("Collection_Insert", 0),
		/*12*/ imports.NewTable("Collection_ItemClass", 0),
		/*13*/ imports.NewTable("Collection_Items", 0),
		/*14*/ imports.NewTable("Collection_Move", 0),
		/*15*/ imports.NewTable("Collection_Owner", 0),
		/*16*/ imports.NewTable("Collection_Sort", 0),
	}
)

func collectionImportAPI() *imports.Imports {
	if collectionImport == nil {
		collectionImport = NewDefaultImports()
		collectionImport.SetImportTable(collectionImportTables)
		collectionImportTables = nil
	}
	return collectionImport
}
