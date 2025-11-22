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

// IFPList Parent: IObject
type IFPList interface {
	IObject
	Add(item uintptr) int32                                            // function
	Expand() IFPList                                                   // function
	Extract(item uintptr) uintptr                                      // function
	First() uintptr                                                    // function
	GetEnumerator() IFPListEnumerator                                  // function
	IndexOf(item uintptr) int32                                        // function
	IndexOfItem(item uintptr, direction types.TDirection) int32        // function
	Last() uintptr                                                     // function
	Remove(item uintptr) int32                                         // function
	AddList(list IFPList)                                              // procedure
	Clear()                                                            // procedure
	Delete(index int32)                                                // procedure
	Exchange(index1 int32, index2 int32)                               // procedure
	Insert(index int32, item uintptr)                                  // procedure
	Move(curIndex int32, newIndex int32)                               // procedure
	Assign(listA IFPList, operator types.TListAssignOp, listB IFPList) // procedure
	Pack()                                                             // procedure
	Capacity() int32                                                   // property Capacity Getter
	SetCapacity(value int32)                                           // property Capacity Setter
	Count() int32                                                      // property Count Getter
	SetCount(value int32)                                              // property Count Setter
	Items(index int32) uintptr                                         // property Items Getter
	SetItems(index int32, value uintptr)                               // property Items Setter
}

type TFPList struct {
	TObject
}

func (m *TFPList) Add(item uintptr) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(1, m.Instance(), uintptr(item))
	return int32(r)
}

func (m *TFPList) Expand() IFPList {
	if !m.IsValid() {
		return nil
	}
	r := fPListAPI().SysCallN(2, m.Instance())
	return AsFPList(r)
}

func (m *TFPList) Extract(item uintptr) uintptr {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(3, m.Instance(), uintptr(item))
	return uintptr(r)
}

func (m *TFPList) First() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(4, m.Instance())
	return uintptr(r)
}

func (m *TFPList) GetEnumerator() IFPListEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := fPListAPI().SysCallN(5, m.Instance())
	return AsFPListEnumerator(r)
}

func (m *TFPList) IndexOf(item uintptr) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(6, m.Instance(), uintptr(item))
	return int32(r)
}

func (m *TFPList) IndexOfItem(item uintptr, direction types.TDirection) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(7, m.Instance(), uintptr(item), uintptr(direction))
	return int32(r)
}

func (m *TFPList) Last() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(8, m.Instance())
	return uintptr(r)
}

func (m *TFPList) Remove(item uintptr) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(9, m.Instance(), uintptr(item))
	return int32(r)
}

func (m *TFPList) AddList(list IFPList) {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(list))
}

func (m *TFPList) Clear() {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(11, m.Instance())
}

func (m *TFPList) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(12, m.Instance(), uintptr(index))
}

func (m *TFPList) Exchange(index1 int32, index2 int32) {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(14, m.Instance(), uintptr(index1), uintptr(index2))
}

func (m *TFPList) Insert(index int32, item uintptr) {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(15, m.Instance(), uintptr(index), uintptr(item))
}

func (m *TFPList) Move(curIndex int32, newIndex int32) {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(16, m.Instance(), uintptr(curIndex), uintptr(newIndex))
}

func (m *TFPList) Assign(listA IFPList, operator types.TListAssignOp, listB IFPList) {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(17, m.Instance(), base.GetObjectUintptr(listA), uintptr(operator), base.GetObjectUintptr(listB))
}

func (m *TFPList) Pack() {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(18, m.Instance())
}

func (m *TFPList) Capacity() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(19, 0, m.Instance())
	return int32(r)
}

func (m *TFPList) SetCapacity(value int32) {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TFPList) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TFPList) SetCount(value int32) {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TFPList) Items(index int32) uintptr {
	if !m.IsValid() {
		return 0
	}
	r := fPListAPI().SysCallN(21, 0, m.Instance(), uintptr(index))
	return uintptr(r)
}

func (m *TFPList) SetItems(index int32, value uintptr) {
	if !m.IsValid() {
		return
	}
	fPListAPI().SysCallN(21, 1, m.Instance(), uintptr(index), uintptr(value))
}

// FPList  is static instance
var FPList _FPListClass

// _FPListClass is class type defined by TFPList
type _FPListClass uintptr

func (_FPListClass) Error(msg string, data uintptr) {
	fPListAPI().SysCallN(13, api.PasStr(msg), uintptr(data))
}

// NewFPList class constructor
func NewFPList() IFPList {
	r := fPListAPI().SysCallN(0)
	return AsFPList(r)
}

var (
	fPListOnce   base.Once
	fPListImport *imports.Imports = nil
)

func fPListAPI() *imports.Imports {
	fPListOnce.Do(func() {
		fPListImport = api.NewDefaultImports()
		fPListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPList_Create", 0), // constructor NewFPList
			/* 1 */ imports.NewTable("TFPList_Add", 0), // function Add
			/* 2 */ imports.NewTable("TFPList_Expand", 0), // function Expand
			/* 3 */ imports.NewTable("TFPList_Extract", 0), // function Extract
			/* 4 */ imports.NewTable("TFPList_First", 0), // function First
			/* 5 */ imports.NewTable("TFPList_GetEnumerator", 0), // function GetEnumerator
			/* 6 */ imports.NewTable("TFPList_IndexOf", 0), // function IndexOf
			/* 7 */ imports.NewTable("TFPList_IndexOfItem", 0), // function IndexOfItem
			/* 8 */ imports.NewTable("TFPList_Last", 0), // function Last
			/* 9 */ imports.NewTable("TFPList_Remove", 0), // function Remove
			/* 10 */ imports.NewTable("TFPList_AddList", 0), // procedure AddList
			/* 11 */ imports.NewTable("TFPList_Clear", 0), // procedure Clear
			/* 12 */ imports.NewTable("TFPList_Delete", 0), // procedure Delete
			/* 13 */ imports.NewTable("TFPList_Error", 0), // static procedure Error
			/* 14 */ imports.NewTable("TFPList_Exchange", 0), // procedure Exchange
			/* 15 */ imports.NewTable("TFPList_Insert", 0), // procedure Insert
			/* 16 */ imports.NewTable("TFPList_Move", 0), // procedure Move
			/* 17 */ imports.NewTable("TFPList_Assign", 0), // procedure Assign
			/* 18 */ imports.NewTable("TFPList_Pack", 0), // procedure Pack
			/* 19 */ imports.NewTable("TFPList_Capacity", 0), // property Capacity
			/* 20 */ imports.NewTable("TFPList_Count", 0), // property Count
			/* 21 */ imports.NewTable("TFPList_Items", 0), // property Items
		}
	})
	return fPListImport
}
