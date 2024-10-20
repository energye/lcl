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

// IComboExItems Parent: IListControlItems
type IComboExItems interface {
	IListControlItems
	ComboItems(AIndex int32) IComboExItem                                                                                                                       // property
	AddForComboExItem() IComboExItem                                                                                                                            // function
	AddItem(ACaption string, AImageIndex SmallInt, AOverlayImageIndex SmallInt, ASelectedImageIndex SmallInt, AIndent SmallInt, AData TCustomData) IComboExItem // function
	InsertForComboExItem(AIndex int32) IComboExItem                                                                                                             // function
}

// TComboExItems Parent: TListControlItems
type TComboExItems struct {
	TListControlItems
}

func NewComboExItems(AOwner IPersistent, AItemClass TCollectionItemClass) IComboExItems {
	r1 := comboExItemsImportAPI().SysCallN(4, GetObjectUintptr(AOwner), uintptr(AItemClass))
	return AsComboExItems(r1)
}

func (m *TComboExItems) ComboItems(AIndex int32) IComboExItem {
	r1 := comboExItemsImportAPI().SysCallN(3, m.Instance(), uintptr(AIndex))
	return AsComboExItem(r1)
}

func (m *TComboExItems) AddForComboExItem() IComboExItem {
	r1 := comboExItemsImportAPI().SysCallN(0, m.Instance())
	return AsComboExItem(r1)
}

func (m *TComboExItems) AddItem(ACaption string, AImageIndex SmallInt, AOverlayImageIndex SmallInt, ASelectedImageIndex SmallInt, AIndent SmallInt, AData TCustomData) IComboExItem {
	r1 := comboExItemsImportAPI().SysCallN(1, m.Instance(), PascalStr(ACaption), uintptr(AImageIndex), uintptr(AOverlayImageIndex), uintptr(ASelectedImageIndex), uintptr(AIndent), uintptr(AData))
	return AsComboExItem(r1)
}

func (m *TComboExItems) InsertForComboExItem(AIndex int32) IComboExItem {
	r1 := comboExItemsImportAPI().SysCallN(5, m.Instance(), uintptr(AIndex))
	return AsComboExItem(r1)
}

func ComboExItemsClass() TClass {
	ret := comboExItemsImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	comboExItemsImport       *imports.Imports = nil
	comboExItemsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ComboExItems_AddForComboExItem", 0),
		/*1*/ imports.NewTable("ComboExItems_AddItem", 0),
		/*2*/ imports.NewTable("ComboExItems_Class", 0),
		/*3*/ imports.NewTable("ComboExItems_ComboItems", 0),
		/*4*/ imports.NewTable("ComboExItems_Create", 0),
		/*5*/ imports.NewTable("ComboExItems_InsertForComboExItem", 0),
	}
)

func comboExItemsImportAPI() *imports.Imports {
	if comboExItemsImport == nil {
		comboExItemsImport = NewDefaultImports()
		comboExItemsImport.SetImportTable(comboExItemsImportTables)
		comboExItemsImportTables = nil
	}
	return comboExItemsImport
}
