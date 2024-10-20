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

// IListControlItems Parent: IOwnedCollection
type IListControlItems interface {
	IOwnedCollection
	ItemsForListControlItem(AIndex int32) IListControlItem // property
	CaseSensitive() bool                                   // property
	SetCaseSensitive(AValue bool)                          // property
	SortType() TListItemsSortType                          // property
	SetSortType(AValue TListItemsSortType)                 // property
	AddForListControlItem() IListControlItem               // function
	CustomSort(fn TListItemsCompare)                       // procedure
	SortForOverload()                                      // procedure
	SetOnCompare(fn TListCompareEvent)                     // property event
}

// TListControlItems Parent: TOwnedCollection
type TListControlItems struct {
	TOwnedCollection
	customSortPtr uintptr
	comparePtr    uintptr
}

func NewListControlItems(AOwner IPersistent, AItemClass TCollectionItemClass) IListControlItems {
	r1 := listControlItemsImportAPI().SysCallN(3, GetObjectUintptr(AOwner), uintptr(AItemClass))
	return AsListControlItems(r1)
}

func (m *TListControlItems) ItemsForListControlItem(AIndex int32) IListControlItem {
	r1 := listControlItemsImportAPI().SysCallN(5, m.Instance(), uintptr(AIndex))
	return AsListControlItem(r1)
}

func (m *TListControlItems) CaseSensitive() bool {
	r1 := listControlItemsImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListControlItems) SetCaseSensitive(AValue bool) {
	listControlItemsImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListControlItems) SortType() TListItemsSortType {
	r1 := listControlItemsImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TListItemsSortType(r1)
}

func (m *TListControlItems) SetSortType(AValue TListItemsSortType) {
	listControlItemsImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TListControlItems) AddForListControlItem() IListControlItem {
	r1 := listControlItemsImportAPI().SysCallN(0, m.Instance())
	return AsListControlItem(r1)
}

func ListControlItemsClass() TClass {
	ret := listControlItemsImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TListControlItems) CustomSort(fn TListItemsCompare) {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	listControlItemsImportAPI().SysCallN(4, m.Instance(), m.customSortPtr)
}

func (m *TListControlItems) SortForOverload() {
	listControlItemsImportAPI().SysCallN(7, m.Instance())
}

func (m *TListControlItems) SetOnCompare(fn TListCompareEvent) {
	if m.comparePtr != 0 {
		RemoveEventElement(m.comparePtr)
	}
	m.comparePtr = MakeEventDataPtr(fn)
	listControlItemsImportAPI().SysCallN(6, m.Instance(), m.comparePtr)
}

var (
	listControlItemsImport       *imports.Imports = nil
	listControlItemsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListControlItems_AddForListControlItem", 0),
		/*1*/ imports.NewTable("ListControlItems_CaseSensitive", 0),
		/*2*/ imports.NewTable("ListControlItems_Class", 0),
		/*3*/ imports.NewTable("ListControlItems_Create", 0),
		/*4*/ imports.NewTable("ListControlItems_CustomSort", 0),
		/*5*/ imports.NewTable("ListControlItems_ItemsForListControlItem", 0),
		/*6*/ imports.NewTable("ListControlItems_SetOnCompare", 0),
		/*7*/ imports.NewTable("ListControlItems_SortForOverload", 0),
		/*8*/ imports.NewTable("ListControlItems_SortType", 0),
	}
)

func listControlItemsImportAPI() *imports.Imports {
	if listControlItemsImport == nil {
		listControlItemsImport = NewDefaultImports()
		listControlItemsImport.SetImportTable(listControlItemsImportTables)
		listControlItemsImportTables = nil
	}
	return listControlItemsImport
}
