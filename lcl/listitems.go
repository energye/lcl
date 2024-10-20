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

// IListItems Parent: IPersistent
type IListItems interface {
	IPersistent
	Flags() TListItemsFlags                                                                              // property
	Count() int32                                                                                        // property
	SetCount(AValue int32)                                                                               // property
	Item(AIndex int32) IListItem                                                                         // property
	SetItem(AIndex int32, AValue IListItem)                                                              // property
	Owner() ICustomListView                                                                              // property
	Add() IListItem                                                                                      // function
	FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem // function
	FindData(AData uintptr) IListItem                                                                    // function
	FindData1(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem                           // function
	GetEnumerator() IListItemsEnumerator                                                                 // function
	IndexOf(AItem IListItem) int32                                                                       // function
	Insert(AIndex int32) IListItem                                                                       // function
	AddItem(AItem IListItem)                                                                             // procedure
	BeginUpdate()                                                                                        // procedure
	Clear()                                                                                              // procedure
	Delete(AIndex int32)                                                                                 // procedure
	EndUpdate()                                                                                          // procedure
	Exchange(AIndex1, AIndex2 int32)                                                                     // procedure
	Move(AFromIndex, AToIndex int32)                                                                     // procedure
	InsertItem(AItem IListItem, AIndex int32)                                                            // procedure
}

// TListItems Parent: TPersistent
type TListItems struct {
	TPersistent
}

func NewListItems(AOwner ICustomListView) IListItems {
	r1 := listItemsImportAPI().SysCallN(6, GetObjectUintptr(AOwner))
	return AsListItems(r1)
}

func (m *TListItems) Flags() TListItemsFlags {
	r1 := listItemsImportAPI().SysCallN(13, m.Instance())
	return TListItemsFlags(r1)
}

func (m *TListItems) Count() int32 {
	r1 := listItemsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListItems) SetCount(AValue int32) {
	listItemsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItems) Item(AIndex int32) IListItem {
	r1 := listItemsImportAPI().SysCallN(18, 0, m.Instance(), uintptr(AIndex))
	return AsListItem(r1)
}

func (m *TListItems) SetItem(AIndex int32, AValue IListItem) {
	listItemsImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AIndex), GetObjectUintptr(AValue))
}

func (m *TListItems) Owner() ICustomListView {
	r1 := listItemsImportAPI().SysCallN(20, m.Instance())
	return AsCustomListView(r1)
}

func (m *TListItems) Add() IListItem {
	r1 := listItemsImportAPI().SysCallN(0, m.Instance())
	return AsListItem(r1)
}

func (m *TListItems) FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem {
	r1 := listItemsImportAPI().SysCallN(10, m.Instance(), uintptr(StartIndex), PascalStr(Value), PascalBool(Partial), PascalBool(Inclusive), PascalBool(Wrap), PascalBool(PartStart))
	return AsListItem(r1)
}

func (m *TListItems) FindData(AData uintptr) IListItem {
	r1 := listItemsImportAPI().SysCallN(11, m.Instance(), uintptr(AData))
	return AsListItem(r1)
}

func (m *TListItems) FindData1(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem {
	r1 := listItemsImportAPI().SysCallN(12, m.Instance(), uintptr(StartIndex), uintptr(Value), PascalBool(Inclusive), PascalBool(Wrap))
	return AsListItem(r1)
}

func (m *TListItems) GetEnumerator() IListItemsEnumerator {
	r1 := listItemsImportAPI().SysCallN(14, m.Instance())
	return AsListItemsEnumerator(r1)
}

func (m *TListItems) IndexOf(AItem IListItem) int32 {
	r1 := listItemsImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(AItem))
	return int32(r1)
}

func (m *TListItems) Insert(AIndex int32) IListItem {
	r1 := listItemsImportAPI().SysCallN(16, m.Instance(), uintptr(AIndex))
	return AsListItem(r1)
}

func ListItemsClass() TClass {
	ret := listItemsImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TListItems) AddItem(AItem IListItem) {
	listItemsImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(AItem))
}

func (m *TListItems) BeginUpdate() {
	listItemsImportAPI().SysCallN(2, m.Instance())
}

func (m *TListItems) Clear() {
	listItemsImportAPI().SysCallN(4, m.Instance())
}

func (m *TListItems) Delete(AIndex int32) {
	listItemsImportAPI().SysCallN(7, m.Instance(), uintptr(AIndex))
}

func (m *TListItems) EndUpdate() {
	listItemsImportAPI().SysCallN(8, m.Instance())
}

func (m *TListItems) Exchange(AIndex1, AIndex2 int32) {
	listItemsImportAPI().SysCallN(9, m.Instance(), uintptr(AIndex1), uintptr(AIndex2))
}

func (m *TListItems) Move(AFromIndex, AToIndex int32) {
	listItemsImportAPI().SysCallN(19, m.Instance(), uintptr(AFromIndex), uintptr(AToIndex))
}

func (m *TListItems) InsertItem(AItem IListItem, AIndex int32) {
	listItemsImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(AItem), uintptr(AIndex))
}

var (
	listItemsImport       *imports.Imports = nil
	listItemsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListItems_Add", 0),
		/*1*/ imports.NewTable("ListItems_AddItem", 0),
		/*2*/ imports.NewTable("ListItems_BeginUpdate", 0),
		/*3*/ imports.NewTable("ListItems_Class", 0),
		/*4*/ imports.NewTable("ListItems_Clear", 0),
		/*5*/ imports.NewTable("ListItems_Count", 0),
		/*6*/ imports.NewTable("ListItems_Create", 0),
		/*7*/ imports.NewTable("ListItems_Delete", 0),
		/*8*/ imports.NewTable("ListItems_EndUpdate", 0),
		/*9*/ imports.NewTable("ListItems_Exchange", 0),
		/*10*/ imports.NewTable("ListItems_FindCaption", 0),
		/*11*/ imports.NewTable("ListItems_FindData", 0),
		/*12*/ imports.NewTable("ListItems_FindData1", 0),
		/*13*/ imports.NewTable("ListItems_Flags", 0),
		/*14*/ imports.NewTable("ListItems_GetEnumerator", 0),
		/*15*/ imports.NewTable("ListItems_IndexOf", 0),
		/*16*/ imports.NewTable("ListItems_Insert", 0),
		/*17*/ imports.NewTable("ListItems_InsertItem", 0),
		/*18*/ imports.NewTable("ListItems_Item", 0),
		/*19*/ imports.NewTable("ListItems_Move", 0),
		/*20*/ imports.NewTable("ListItems_Owner", 0),
	}
)

func listItemsImportAPI() *imports.Imports {
	if listItemsImport == nil {
		listItemsImport = NewDefaultImports()
		listItemsImport.SetImportTable(listItemsImportTables)
		listItemsImportTables = nil
	}
	return listItemsImport
}
