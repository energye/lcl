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

// IListControlItems Parent: IOwnedCollection
type IListControlItems interface {
	IOwnedCollection
	AddToListControlItem() IListControlItem                     // function
	Sort()                                                      // procedure
	ItemsWithIntToListControlItem(index int32) IListControlItem // property Items Getter
	CaseSensitive() bool                                        // property CaseSensitive Getter
	SetCaseSensitive(value bool)                                // property CaseSensitive Setter
	SortType() types.TListItemsSortType                         // property SortType Getter
	SetSortType(value types.TListItemsSortType)                 // property SortType Setter
	SetOnCompare(fn TListCompareEvent)                          // property event
}

type TListControlItems struct {
	TOwnedCollection
}

func (m *TListControlItems) AddToListControlItem() IListControlItem {
	if !m.IsValid() {
		return nil
	}
	r := listControlItemsAPI().SysCallN(1, m.Instance())
	return AsListControlItem(r)
}

func (m *TListControlItems) Sort() {
	if !m.IsValid() {
		return
	}
	listControlItemsAPI().SysCallN(2, m.Instance())
}

func (m *TListControlItems) ItemsWithIntToListControlItem(index int32) IListControlItem {
	if !m.IsValid() {
		return nil
	}
	r := listControlItemsAPI().SysCallN(3, m.Instance(), uintptr(index))
	return AsListControlItem(r)
}

func (m *TListControlItems) CaseSensitive() bool {
	if !m.IsValid() {
		return false
	}
	r := listControlItemsAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListControlItems) SetCaseSensitive(value bool) {
	if !m.IsValid() {
		return
	}
	listControlItemsAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TListControlItems) SortType() types.TListItemsSortType {
	if !m.IsValid() {
		return 0
	}
	r := listControlItemsAPI().SysCallN(5, 0, m.Instance())
	return types.TListItemsSortType(r)
}

func (m *TListControlItems) SetSortType(value types.TListItemsSortType) {
	if !m.IsValid() {
		return
	}
	listControlItemsAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TListControlItems) SetOnCompare(fn TListCompareEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTListCompareEvent(fn)
	base.SetEvent(m, 6, listControlItemsAPI(), api.MakeEventDataPtr(cb))
}

// NewListControlItems class constructor
func NewListControlItems(owner IPersistent, itemClass types.TCollectionItemClass) IListControlItems {
	r := listControlItemsAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(itemClass))
	return AsListControlItems(r)
}

var (
	listControlItemsOnce   base.Once
	listControlItemsImport *imports.Imports = nil
)

func listControlItemsAPI() *imports.Imports {
	listControlItemsOnce.Do(func() {
		listControlItemsImport = api.NewDefaultImports()
		listControlItemsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListControlItems_Create", 0), // constructor NewListControlItems
			/* 1 */ imports.NewTable("TListControlItems_AddToListControlItem", 0), // function AddToListControlItem
			/* 2 */ imports.NewTable("TListControlItems_Sort", 0), // procedure Sort
			/* 3 */ imports.NewTable("TListControlItems_ItemsWithIntToListControlItem", 0), // property ItemsWithIntToListControlItem
			/* 4 */ imports.NewTable("TListControlItems_CaseSensitive", 0), // property CaseSensitive
			/* 5 */ imports.NewTable("TListControlItems_SortType", 0), // property SortType
			/* 6 */ imports.NewTable("TListControlItems_OnCompare", 0), // event OnCompare
		}
	})
	return listControlItemsImport
}
