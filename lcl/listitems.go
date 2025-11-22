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

// IListItems Parent: IPersistent
type IListItems interface {
	IPersistent
	Add() IListItem                                                                                                // function
	FindCaption(startIndex int32, value string, partial bool, inclusive bool, wrap bool, partStart bool) IListItem // function
	FindDataWithPointer(data uintptr) IListItem                                                                    // function
	FindDataWithIntPointerBoolX2(startIndex int32, value uintptr, inclusive bool, wrap bool) IListItem             // function
	GetEnumerator() IListItemsEnumerator                                                                           // function
	IndexOf(item IListItem) int32                                                                                  // function
	Insert(index int32) IListItem                                                                                  // function
	AddItem(item IListItem)                                                                                        // procedure
	BeginUpdate()                                                                                                  // procedure
	Clear()                                                                                                        // procedure
	Delete(index int32)                                                                                            // procedure
	EndUpdate()                                                                                                    // procedure
	Exchange(index1 int32, index2 int32)                                                                           // procedure
	Move(fromIndex int32, toIndex int32)                                                                           // procedure
	InsertItem(item IListItem, index int32)                                                                        // procedure
	Flags() types.TListItemsFlags                                                                                  // property Flags Getter
	Count() int32                                                                                                  // property Count Getter
	SetCount(value int32)                                                                                          // property Count Setter
	Item(index int32) IListItem                                                                                    // property Item Getter
	SetItem(index int32, value IListItem)                                                                          // property Item Setter
	Owner() ICustomListView                                                                                        // property Owner Getter
}

type TListItems struct {
	TPersistent
}

func (m *TListItems) Add() IListItem {
	if !m.IsValid() {
		return nil
	}
	r := listItemsAPI().SysCallN(1, m.Instance())
	return AsListItem(r)
}

func (m *TListItems) FindCaption(startIndex int32, value string, partial bool, inclusive bool, wrap bool, partStart bool) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := listItemsAPI().SysCallN(2, m.Instance(), uintptr(startIndex), api.PasStr(value), api.PasBool(partial), api.PasBool(inclusive), api.PasBool(wrap), api.PasBool(partStart))
	return AsListItem(r)
}

func (m *TListItems) FindDataWithPointer(data uintptr) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := listItemsAPI().SysCallN(3, m.Instance(), uintptr(data))
	return AsListItem(r)
}

func (m *TListItems) FindDataWithIntPointerBoolX2(startIndex int32, value uintptr, inclusive bool, wrap bool) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := listItemsAPI().SysCallN(4, m.Instance(), uintptr(startIndex), uintptr(value), api.PasBool(inclusive), api.PasBool(wrap))
	return AsListItem(r)
}

func (m *TListItems) GetEnumerator() IListItemsEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := listItemsAPI().SysCallN(5, m.Instance())
	return AsListItemsEnumerator(r)
}

func (m *TListItems) IndexOf(item IListItem) int32 {
	if !m.IsValid() {
		return 0
	}
	r := listItemsAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(item))
	return int32(r)
}

func (m *TListItems) Insert(index int32) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := listItemsAPI().SysCallN(7, m.Instance(), uintptr(index))
	return AsListItem(r)
}

func (m *TListItems) AddItem(item IListItem) {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(item))
}

func (m *TListItems) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(9, m.Instance())
}

func (m *TListItems) Clear() {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(10, m.Instance())
}

func (m *TListItems) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(11, m.Instance(), uintptr(index))
}

func (m *TListItems) EndUpdate() {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(12, m.Instance())
}

func (m *TListItems) Exchange(index1 int32, index2 int32) {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(13, m.Instance(), uintptr(index1), uintptr(index2))
}

func (m *TListItems) Move(fromIndex int32, toIndex int32) {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(14, m.Instance(), uintptr(fromIndex), uintptr(toIndex))
}

func (m *TListItems) InsertItem(item IListItem, index int32) {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(item), uintptr(index))
}

func (m *TListItems) Flags() types.TListItemsFlags {
	if !m.IsValid() {
		return 0
	}
	r := listItemsAPI().SysCallN(16, m.Instance())
	return types.TListItemsFlags(r)
}

func (m *TListItems) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listItemsAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TListItems) SetCount(value int32) {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TListItems) Item(index int32) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := listItemsAPI().SysCallN(18, 0, m.Instance(), uintptr(index))
	return AsListItem(r)
}

func (m *TListItems) SetItem(index int32, value IListItem) {
	if !m.IsValid() {
		return
	}
	listItemsAPI().SysCallN(18, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TListItems) Owner() ICustomListView {
	if !m.IsValid() {
		return nil
	}
	r := listItemsAPI().SysCallN(19, m.Instance())
	return AsCustomListView(r)
}

// NewListItems class constructor
func NewListItems(owner ICustomListView) IListItems {
	r := listItemsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsListItems(r)
}

var (
	listItemsOnce   base.Once
	listItemsImport *imports.Imports = nil
)

func listItemsAPI() *imports.Imports {
	listItemsOnce.Do(func() {
		listItemsImport = api.NewDefaultImports()
		listItemsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListItems_Create", 0), // constructor NewListItems
			/* 1 */ imports.NewTable("TListItems_Add", 0), // function Add
			/* 2 */ imports.NewTable("TListItems_FindCaption", 0), // function FindCaption
			/* 3 */ imports.NewTable("TListItems_FindDataWithPointer", 0), // function FindDataWithPointer
			/* 4 */ imports.NewTable("TListItems_FindDataWithIntPointerBoolX2", 0), // function FindDataWithIntPointerBoolX2
			/* 5 */ imports.NewTable("TListItems_GetEnumerator", 0), // function GetEnumerator
			/* 6 */ imports.NewTable("TListItems_IndexOf", 0), // function IndexOf
			/* 7 */ imports.NewTable("TListItems_Insert", 0), // function Insert
			/* 8 */ imports.NewTable("TListItems_AddItem", 0), // procedure AddItem
			/* 9 */ imports.NewTable("TListItems_BeginUpdate", 0), // procedure BeginUpdate
			/* 10 */ imports.NewTable("TListItems_Clear", 0), // procedure Clear
			/* 11 */ imports.NewTable("TListItems_Delete", 0), // procedure Delete
			/* 12 */ imports.NewTable("TListItems_EndUpdate", 0), // procedure EndUpdate
			/* 13 */ imports.NewTable("TListItems_Exchange", 0), // procedure Exchange
			/* 14 */ imports.NewTable("TListItems_Move", 0), // procedure Move
			/* 15 */ imports.NewTable("TListItems_InsertItem", 0), // procedure InsertItem
			/* 16 */ imports.NewTable("TListItems_Flags", 0), // property Flags
			/* 17 */ imports.NewTable("TListItems_Count", 0), // property Count
			/* 18 */ imports.NewTable("TListItems_Item", 0), // property Item
			/* 19 */ imports.NewTable("TListItems_Owner", 0), // property Owner
		}
	})
	return listItemsImport
}
