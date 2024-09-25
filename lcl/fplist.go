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
	r1 := LCL().SysCallN(2993)
	return AsFPList(r1)
}

func (m *TFPList) Capacity() int32 {
	r1 := LCL().SysCallN(2989, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPList) SetCapacity(AValue int32) {
	LCL().SysCallN(2989, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPList) Count() int32 {
	r1 := LCL().SysCallN(2992, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPList) SetCount(AValue int32) {
	LCL().SysCallN(2992, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPList) Items(Index int32) uintptr {
	r1 := LCL().SysCallN(3002, 0, m.Instance(), uintptr(Index))
	return uintptr(r1)
}

func (m *TFPList) SetItems(Index int32, AValue uintptr) {
	LCL().SysCallN(3002, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TFPList) List() uintptr {
	r1 := LCL().SysCallN(3004, m.Instance())
	return uintptr(r1)
}

func (m *TFPList) Add(Item uintptr) int32 {
	r1 := LCL().SysCallN(2987, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TFPList) Expand() IFPList {
	r1 := LCL().SysCallN(2996, m.Instance())
	return AsFPList(r1)
}

func (m *TFPList) Extract(Item uintptr) uintptr {
	r1 := LCL().SysCallN(2997, m.Instance(), uintptr(Item))
	return uintptr(r1)
}

func (m *TFPList) First() uintptr {
	r1 := LCL().SysCallN(2998, m.Instance())
	return uintptr(r1)
}

func (m *TFPList) GetEnumerator() IFPListEnumerator {
	r1 := LCL().SysCallN(2999, m.Instance())
	return AsFPListEnumerator(r1)
}

func (m *TFPList) IndexOf(Item uintptr) int32 {
	r1 := LCL().SysCallN(3000, m.Instance(), uintptr(Item))
	return int32(r1)
}

func (m *TFPList) Last() uintptr {
	r1 := LCL().SysCallN(3003, m.Instance())
	return uintptr(r1)
}

func (m *TFPList) Remove(Item uintptr) int32 {
	r1 := LCL().SysCallN(3007, m.Instance(), uintptr(Item))
	return int32(r1)
}

func FPListClass() TClass {
	ret := LCL().SysCallN(2990)
	return TClass(ret)
}

func (m *TFPList) AddList(AList IFPList) {
	LCL().SysCallN(2988, m.Instance(), GetObjectUintptr(AList))
}

func (m *TFPList) Clear() {
	LCL().SysCallN(2991, m.Instance())
}

func (m *TFPList) Delete(Index int32) {
	LCL().SysCallN(2994, m.Instance(), uintptr(Index))
}

func (m *TFPList) Exchange(Index1, Index2 int32) {
	LCL().SysCallN(2995, m.Instance(), uintptr(Index1), uintptr(Index2))
}

func (m *TFPList) Insert(Index int32, Item uintptr) {
	LCL().SysCallN(3001, m.Instance(), uintptr(Index), uintptr(Item))
}

func (m *TFPList) Move(CurIndex, NewIndex int32) {
	LCL().SysCallN(3005, m.Instance(), uintptr(CurIndex), uintptr(NewIndex))
}

func (m *TFPList) Pack() {
	LCL().SysCallN(3006, m.Instance())
}

func (m *TFPList) Sort(fn TListSortCompare) {
	if m.sortPtr != 0 {
		RemoveEventElement(m.sortPtr)
	}
	m.sortPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3008, m.Instance(), m.sortPtr)
}
