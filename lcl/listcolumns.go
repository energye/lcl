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

// IListColumns Parent: ICollection
type IListColumns interface {
	ICollection
	AddToListColumn() IListColumn                               // function
	Update(item ICollectionItem)                                // procedure
	OwnerToCustomListView() ICustomListView                     // property Owner Getter
	ItemsWithIntToListColumn(index int32) IListColumn           // property Items Getter
	SetItemsWithIntToListColumn(index int32, value IListColumn) // property Items Setter
}

type TListColumns struct {
	TCollection
}

func (m *TListColumns) AddToListColumn() IListColumn {
	if !m.IsValid() {
		return nil
	}
	r := listColumnsAPI().SysCallN(1, m.Instance())
	return AsListColumn(r)
}

func (m *TListColumns) Update(item ICollectionItem) {
	if !m.IsValid() {
		return
	}
	listColumnsAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(item))
}

func (m *TListColumns) OwnerToCustomListView() ICustomListView {
	if !m.IsValid() {
		return nil
	}
	r := listColumnsAPI().SysCallN(3, m.Instance())
	return AsCustomListView(r)
}

func (m *TListColumns) ItemsWithIntToListColumn(index int32) IListColumn {
	if !m.IsValid() {
		return nil
	}
	r := listColumnsAPI().SysCallN(4, 0, m.Instance(), uintptr(index))
	return AsListColumn(r)
}

func (m *TListColumns) SetItemsWithIntToListColumn(index int32, value IListColumn) {
	if !m.IsValid() {
		return
	}
	listColumnsAPI().SysCallN(4, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

// NewListColumns class constructor
func NewListColumns(owner ICustomListView) IListColumns {
	r := listColumnsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsListColumns(r)
}

var (
	listColumnsOnce   base.Once
	listColumnsImport *imports.Imports = nil
)

func listColumnsAPI() *imports.Imports {
	listColumnsOnce.Do(func() {
		listColumnsImport = api.NewDefaultImports()
		listColumnsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListColumns_Create", 0), // constructor NewListColumns
			/* 1 */ imports.NewTable("TListColumns_AddToListColumn", 0), // function AddToListColumn
			/* 2 */ imports.NewTable("TListColumns_Update", 0), // procedure Update
			/* 3 */ imports.NewTable("TListColumns_OwnerToCustomListView", 0), // property OwnerToCustomListView
			/* 4 */ imports.NewTable("TListColumns_ItemsWithIntToListColumn", 0), // property ItemsWithIntToListColumn
		}
	})
	return listColumnsImport
}
