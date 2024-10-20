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

// IList Parent: IObject
type IList interface {
	IObject
	Capacity() int32                                                                   // property
	SetCapacity(AValue int32)                                                          // property
	Count() int32                                                                      // property
	SetCount(AValue int32)                                                             // property
	Items(Index int32) uintptr                                                         // property
	SetItems(Index int32, AValue uintptr)                                              // property
	List() uintptr                                                                     // property
	Add(Item uintptr) int32                                                            // function
	Expand() IList                                                                     // function
	Extract(item uintptr) uintptr                                                      // function
	First() uintptr                                                                    // function
	GetEnumerator() IListEnumerator                                                    // function
	IndexOf(Item uintptr) int32                                                        // function
	Last() uintptr                                                                     // function
	Remove(Item uintptr) int32                                                         // function
	FPOAttachObserver(AObserver IObject)                                               // procedure
	FPODetachObserver(AObserver IObject)                                               // procedure
	FPONotifyObservers(ASender IObject, AOperation TFPObservedOperation, Data uintptr) // procedure
	AddList(AList IList)                                                               // procedure
	Clear()                                                                            // procedure
	Delete(Index int32)                                                                // procedure
	Exchange(Index1, Index2 int32)                                                     // procedure
	Insert(Index int32, Item uintptr)                                                  // procedure
	Move(CurIndex, NewIndex int32)                                                     // procedure
	Pack()                                                                             // procedure
	Sort(fn TListSortCompare)                                                          // procedure
}

// TList Parent: TObject
type TList struct {
	TObject
	sortPtr uintptr
}

func NewList() IList {
	r1 := listImportAPI().SysCallN(6)
	return AsList(r1)
}

func (m *TList) Capacity() int32 {
	r1 := listImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TList) SetCapacity(AValue int32) {
	listImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TList) Count() int32 {
	r1 := listImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TList) SetCount(AValue int32) {
	listImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TList) Items(Index int32) uintptr {
	r1 := listImportAPI().SysCallN(18, 0, m.Instance(), uintptr(Index))
	return uintptr(r1)
}

func (m *TList) SetItems(Index int32, AValue uintptr) {
	listImportAPI().SysCallN(18, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TList) List() uintptr {
	r1 := listImportAPI().SysCallN(20, m.Instance())
	return uintptr(r1)
}

func (m *TList) Add(Item uintptr) int32 {
	r1 := listImportAPI().SysCallN(0, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TList) Expand() IList {
	r1 := listImportAPI().SysCallN(9, m.Instance())
	return AsList(r1)
}

func (m *TList) Extract(item uintptr) uintptr {
	r1 := listImportAPI().SysCallN(10, m.Instance(), uintptr(item))
	return uintptr(r1)
}

func (m *TList) First() uintptr {
	r1 := listImportAPI().SysCallN(14, m.Instance())
	return uintptr(r1)
}

func (m *TList) GetEnumerator() IListEnumerator {
	r1 := listImportAPI().SysCallN(15, m.Instance())
	return AsListEnumerator(r1)
}

func (m *TList) IndexOf(Item uintptr) int32 {
	r1 := listImportAPI().SysCallN(16, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TList) Last() uintptr {
	r1 := listImportAPI().SysCallN(19, m.Instance())
	return uintptr(r1)
}

func (m *TList) Remove(Item uintptr) int32 {
	r1 := listImportAPI().SysCallN(23, m.Instance(), uintptr(Item))
	return int32(r1)
}

func ListClass() TClass {
	ret := listImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TList) FPOAttachObserver(AObserver IObject) {
	listImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TList) FPODetachObserver(AObserver IObject) {
	listImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TList) FPONotifyObservers(ASender IObject, AOperation TFPObservedOperation, Data uintptr) {
	listImportAPI().SysCallN(13, m.Instance(), GetObjectUintptr(ASender), uintptr(AOperation), uintptr(Data))
}

func (m *TList) AddList(AList IList) {
	listImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(AList))
}

func (m *TList) Clear() {
	listImportAPI().SysCallN(4, m.Instance())
}

func (m *TList) Delete(Index int32) {
	listImportAPI().SysCallN(7, m.Instance(), uintptr(Index))
}

func (m *TList) Exchange(Index1, Index2 int32) {
	listImportAPI().SysCallN(8, m.Instance(), uintptr(Index1), uintptr(Index2))
}

func (m *TList) Insert(Index int32, Item uintptr) {
	listImportAPI().SysCallN(17, m.Instance(), uintptr(Index), uintptr(Item))
}

func (m *TList) Move(CurIndex, NewIndex int32) {
	listImportAPI().SysCallN(21, m.Instance(), uintptr(CurIndex), uintptr(NewIndex))
}

func (m *TList) Pack() {
	listImportAPI().SysCallN(22, m.Instance())
}

func (m *TList) Sort(fn TListSortCompare) {
	if m.sortPtr != 0 {
		RemoveEventElement(m.sortPtr)
	}
	m.sortPtr = MakeEventDataPtr(fn)
	listImportAPI().SysCallN(24, m.Instance(), m.sortPtr)
}

var (
	listImport       *imports.Imports = nil
	listImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("List_Add", 0),
		/*1*/ imports.NewTable("List_AddList", 0),
		/*2*/ imports.NewTable("List_Capacity", 0),
		/*3*/ imports.NewTable("List_Class", 0),
		/*4*/ imports.NewTable("List_Clear", 0),
		/*5*/ imports.NewTable("List_Count", 0),
		/*6*/ imports.NewTable("List_Create", 0),
		/*7*/ imports.NewTable("List_Delete", 0),
		/*8*/ imports.NewTable("List_Exchange", 0),
		/*9*/ imports.NewTable("List_Expand", 0),
		/*10*/ imports.NewTable("List_Extract", 0),
		/*11*/ imports.NewTable("List_FPOAttachObserver", 0),
		/*12*/ imports.NewTable("List_FPODetachObserver", 0),
		/*13*/ imports.NewTable("List_FPONotifyObservers", 0),
		/*14*/ imports.NewTable("List_First", 0),
		/*15*/ imports.NewTable("List_GetEnumerator", 0),
		/*16*/ imports.NewTable("List_IndexOf", 0),
		/*17*/ imports.NewTable("List_Insert", 0),
		/*18*/ imports.NewTable("List_Items", 0),
		/*19*/ imports.NewTable("List_Last", 0),
		/*20*/ imports.NewTable("List_List", 0),
		/*21*/ imports.NewTable("List_Move", 0),
		/*22*/ imports.NewTable("List_Pack", 0),
		/*23*/ imports.NewTable("List_Remove", 0),
		/*24*/ imports.NewTable("List_Sort", 0),
	}
)

func listImportAPI() *imports.Imports {
	if listImport == nil {
		listImport = NewDefaultImports()
		listImport.SetImportTable(listImportTables)
		listImportTables = nil
	}
	return listImport
}
