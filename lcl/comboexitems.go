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

// IComboExItems Parent: IListControlItems
type IComboExItems interface {
	IListControlItems
	AddToComboExItem() IComboExItem                                                                                                                                                     // function
	AddItem(caption string, imageIndex types.SmallInt, overlayImageIndex types.SmallInt, selectedImageIndex types.SmallInt, indent types.SmallInt, data types.TCustomData) IComboExItem // function
	InsertWithInt(index int32) IComboExItem                                                                                                                                             // function
	ComboItems(index int32) IComboExItem                                                                                                                                                // property ComboItems Getter
}

type TComboExItems struct {
	TListControlItems
}

func (m *TComboExItems) AddToComboExItem() IComboExItem {
	if !m.IsValid() {
		return nil
	}
	r := comboExItemsAPI().SysCallN(1, m.Instance())
	return AsComboExItem(r)
}

func (m *TComboExItems) AddItem(caption string, imageIndex types.SmallInt, overlayImageIndex types.SmallInt, selectedImageIndex types.SmallInt, indent types.SmallInt, data types.TCustomData) IComboExItem {
	if !m.IsValid() {
		return nil
	}
	r := comboExItemsAPI().SysCallN(2, m.Instance(), api.PasStr(caption), uintptr(imageIndex), uintptr(overlayImageIndex), uintptr(selectedImageIndex), uintptr(indent), uintptr(data))
	return AsComboExItem(r)
}

func (m *TComboExItems) InsertWithInt(index int32) IComboExItem {
	if !m.IsValid() {
		return nil
	}
	r := comboExItemsAPI().SysCallN(3, m.Instance(), uintptr(index))
	return AsComboExItem(r)
}

func (m *TComboExItems) ComboItems(index int32) IComboExItem {
	if !m.IsValid() {
		return nil
	}
	r := comboExItemsAPI().SysCallN(4, m.Instance(), uintptr(index))
	return AsComboExItem(r)
}

// NewComboExItems class constructor
func NewComboExItems(owner IPersistent, itemClass types.TCollectionItemClass) IComboExItems {
	r := comboExItemsAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(itemClass))
	return AsComboExItems(r)
}

var (
	comboExItemsOnce   base.Once
	comboExItemsImport *imports.Imports = nil
)

func comboExItemsAPI() *imports.Imports {
	comboExItemsOnce.Do(func() {
		comboExItemsImport = api.NewDefaultImports()
		comboExItemsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TComboExItems_Create", 0), // constructor NewComboExItems
			/* 1 */ imports.NewTable("TComboExItems_AddToComboExItem", 0), // function AddToComboExItem
			/* 2 */ imports.NewTable("TComboExItems_AddItem", 0), // function AddItem
			/* 3 */ imports.NewTable("TComboExItems_InsertWithInt", 0), // function InsertWithInt
			/* 4 */ imports.NewTable("TComboExItems_ComboItems", 0), // property ComboItems
		}
	})
	return comboExItemsImport
}
