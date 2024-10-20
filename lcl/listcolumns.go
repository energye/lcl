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

// IListColumns Parent: ICollection
type IListColumns interface {
	ICollection
	OwnerForCustomListView() ICustomListView                // property
	ItemsForListColumn(AIndex int32) IListColumn            // property
	SetItemsForListColumn(AIndex int32, AValue IListColumn) // property
	AddForListColumn() IListColumn                          // function
	Update(Item ICollectionItem)                            // procedure
}

// TListColumns Parent: TCollection
type TListColumns struct {
	TCollection
}

func NewListColumns(AOwner ICustomListView) IListColumns {
	r1 := listColumnsImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsListColumns(r1)
}

func (m *TListColumns) OwnerForCustomListView() ICustomListView {
	r1 := listColumnsImportAPI().SysCallN(4, m.Instance())
	return AsCustomListView(r1)
}

func (m *TListColumns) ItemsForListColumn(AIndex int32) IListColumn {
	r1 := listColumnsImportAPI().SysCallN(3, 0, m.Instance(), uintptr(AIndex))
	return AsListColumn(r1)
}

func (m *TListColumns) SetItemsForListColumn(AIndex int32, AValue IListColumn) {
	listColumnsImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AIndex), GetObjectUintptr(AValue))
}

func (m *TListColumns) AddForListColumn() IListColumn {
	r1 := listColumnsImportAPI().SysCallN(0, m.Instance())
	return AsListColumn(r1)
}

func ListColumnsClass() TClass {
	ret := listColumnsImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TListColumns) Update(Item ICollectionItem) {
	listColumnsImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(Item))
}

var (
	listColumnsImport       *imports.Imports = nil
	listColumnsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListColumns_AddForListColumn", 0),
		/*1*/ imports.NewTable("ListColumns_Class", 0),
		/*2*/ imports.NewTable("ListColumns_Create", 0),
		/*3*/ imports.NewTable("ListColumns_ItemsForListColumn", 0),
		/*4*/ imports.NewTable("ListColumns_OwnerForCustomListView", 0),
		/*5*/ imports.NewTable("ListColumns_Update", 0),
	}
)

func listColumnsImportAPI() *imports.Imports {
	if listColumnsImport == nil {
		listColumnsImport = NewDefaultImports()
		listColumnsImport.SetImportTable(listColumnsImportTables)
		listColumnsImportTables = nil
	}
	return listColumnsImport
}
