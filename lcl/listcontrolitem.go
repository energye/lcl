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

// IListControlItem Parent: ICollectionItem
type IListControlItem interface {
	ICollectionItem
	Data() TCustomData                // property
	SetData(AValue TCustomData)       // property
	Caption() string                  // property
	SetCaption(AValue string)         // property
	ImageIndex() TImageIndex          // property
	SetImageIndex(AValue TImageIndex) // property
}

// TListControlItem Parent: TCollectionItem
type TListControlItem struct {
	TCollectionItem
}

func NewListControlItem(ACollection ICollection) IListControlItem {
	r1 := listControlItemImportAPI().SysCallN(2, GetObjectUintptr(ACollection))
	return AsListControlItem(r1)
}

func (m *TListControlItem) Data() TCustomData {
	r1 := listControlItemImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCustomData(r1)
}

func (m *TListControlItem) SetData(AValue TCustomData) {
	listControlItemImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TListControlItem) Caption() string {
	r1 := listControlItemImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TListControlItem) SetCaption(AValue string) {
	listControlItemImportAPI().SysCallN(0, 1, m.Instance(), PascalStr(AValue))
}

func (m *TListControlItem) ImageIndex() TImageIndex {
	r1 := listControlItemImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TListControlItem) SetImageIndex(AValue TImageIndex) {
	listControlItemImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func ListControlItemClass() TClass {
	ret := listControlItemImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	listControlItemImport       *imports.Imports = nil
	listControlItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListControlItem_Caption", 0),
		/*1*/ imports.NewTable("ListControlItem_Class", 0),
		/*2*/ imports.NewTable("ListControlItem_Create", 0),
		/*3*/ imports.NewTable("ListControlItem_Data", 0),
		/*4*/ imports.NewTable("ListControlItem_ImageIndex", 0),
	}
)

func listControlItemImportAPI() *imports.Imports {
	if listControlItemImport == nil {
		listControlItemImport = NewDefaultImports()
		listControlItemImport.SetImportTable(listControlItemImportTables)
		listControlItemImportTables = nil
	}
	return listControlItemImport
}
