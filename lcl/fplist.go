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

// IFPList Parent: IObject
type IFPList interface {
	IObject
	Capacity() int32                      // property
	SetCapacity(AValue int32)             // property
	Count() int32                         // property
	SetCount(AValue int32)                // property
	Items(Index int32) uintptr            // property
	SetItems(Index int32, AValue uintptr) // property
	List() uintptr                        // property
	Add(Item uintptr) int32               // function
	Expand() IFPList                      // function
	Extract(Item uintptr) uintptr         // function
	First() uintptr                       // function
	GetEnumerator() IFPListEnumerator     // function
	IndexOf(Item uintptr) int32           // function
	Last() uintptr                        // function
	Remove(Item uintptr) int32            // function
	AddList(AList IFPList)                // procedure
	Clear()                               // procedure
	Delete(Index int32)                   // procedure
	Exchange(Index1, Index2 int32)        // procedure
	Insert(Index int32, Item uintptr)     // procedure
	Move(CurIndex, NewIndex int32)        // procedure
	Pack()                                // procedure
	Sort(fn TListSortCompare)             // procedure
}

// TFPList Parent: TObject
type TFPList struct {
	TObject
	sortPtr uintptr
}

func NewFPList() IFPList {
	r1 := fPListImportAPI().SysCallN(6)
	return AsFPList(r1)
}

func (m *TFPList) Capacity() int32 {
	r1 := fPListImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPList) SetCapacity(AValue int32) {
	fPListImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPList) Count() int32 {
	r1 := fPListImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPList) SetCount(AValue int32) {
	fPListImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPList) Items(Index int32) uintptr {
	r1 := fPListImportAPI().SysCallN(15, 0, m.Instance(), uintptr(Index))
	return uintptr(r1)
}

func (m *TFPList) SetItems(Index int32, AValue uintptr) {
	fPListImportAPI().SysCallN(15, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TFPList) List() uintptr {
	r1 := fPListImportAPI().SysCallN(17, m.Instance())
	return uintptr(r1)
}

func (m *TFPList) Add(Item uintptr) int32 {
	r1 := fPListImportAPI().SysCallN(0, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TFPList) Expand() IFPList {
	r1 := fPListImportAPI().SysCallN(9, m.Instance())
	return AsFPList(r1)
}

func (m *TFPList) Extract(Item uintptr) uintptr {
	r1 := fPListImportAPI().SysCallN(10, m.Instance(), uintptr(Item))
	return uintptr(r1)
}

func (m *TFPList) First() uintptr {
	r1 := fPListImportAPI().SysCallN(11, m.Instance())
	return uintptr(r1)
}

func (m *TFPList) GetEnumerator() IFPListEnumerator {
	r1 := fPListImportAPI().SysCallN(12, m.Instance())
	return AsFPListEnumerator(r1)
}

func (m *TFPList) IndexOf(Item uintptr) int32 {
	r1 := fPListImportAPI().SysCallN(13, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TFPList) Last() uintptr {
	r1 := fPListImportAPI().SysCallN(16, m.Instance())
	return uintptr(r1)
}

func (m *TFPList) Remove(Item uintptr) int32 {
	r1 := fPListImportAPI().SysCallN(20, m.Instance(), uintptr(Item))
	return int32(r1)
}

func FPListClass() TClass {
	ret := fPListImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TFPList) AddList(AList IFPList) {
	fPListImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(AList))
}

func (m *TFPList) Clear() {
	fPListImportAPI().SysCallN(4, m.Instance())
}

func (m *TFPList) Delete(Index int32) {
	fPListImportAPI().SysCallN(7, m.Instance(), uintptr(Index))
}

func (m *TFPList) Exchange(Index1, Index2 int32) {
	fPListImportAPI().SysCallN(8, m.Instance(), uintptr(Index1), uintptr(Index2))
}

func (m *TFPList) Insert(Index int32, Item uintptr) {
	fPListImportAPI().SysCallN(14, m.Instance(), uintptr(Index), uintptr(Item))
}

func (m *TFPList) Move(CurIndex, NewIndex int32) {
	fPListImportAPI().SysCallN(18, m.Instance(), uintptr(CurIndex), uintptr(NewIndex))
}

func (m *TFPList) Pack() {
	fPListImportAPI().SysCallN(19, m.Instance())
}

func (m *TFPList) Sort(fn TListSortCompare) {
	if m.sortPtr != 0 {
		RemoveEventElement(m.sortPtr)
	}
	m.sortPtr = MakeEventDataPtr(fn)
	fPListImportAPI().SysCallN(21, m.Instance(), m.sortPtr)
}

var (
	fPListImport       *imports.Imports = nil
	fPListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPList_Add", 0),
		/*1*/ imports.NewTable("FPList_AddList", 0),
		/*2*/ imports.NewTable("FPList_Capacity", 0),
		/*3*/ imports.NewTable("FPList_Class", 0),
		/*4*/ imports.NewTable("FPList_Clear", 0),
		/*5*/ imports.NewTable("FPList_Count", 0),
		/*6*/ imports.NewTable("FPList_Create", 0),
		/*7*/ imports.NewTable("FPList_Delete", 0),
		/*8*/ imports.NewTable("FPList_Exchange", 0),
		/*9*/ imports.NewTable("FPList_Expand", 0),
		/*10*/ imports.NewTable("FPList_Extract", 0),
		/*11*/ imports.NewTable("FPList_First", 0),
		/*12*/ imports.NewTable("FPList_GetEnumerator", 0),
		/*13*/ imports.NewTable("FPList_IndexOf", 0),
		/*14*/ imports.NewTable("FPList_Insert", 0),
		/*15*/ imports.NewTable("FPList_Items", 0),
		/*16*/ imports.NewTable("FPList_Last", 0),
		/*17*/ imports.NewTable("FPList_List", 0),
		/*18*/ imports.NewTable("FPList_Move", 0),
		/*19*/ imports.NewTable("FPList_Pack", 0),
		/*20*/ imports.NewTable("FPList_Remove", 0),
		/*21*/ imports.NewTable("FPList_Sort", 0),
	}
)

func fPListImportAPI() *imports.Imports {
	if fPListImport == nil {
		fPListImport = NewDefaultImports()
		fPListImport.SetImportTable(fPListImportTables)
		fPListImportTables = nil
	}
	return fPListImport
}
