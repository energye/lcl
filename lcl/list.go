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

// IList Parent: IObject
type IList interface {
	IObject
	Add(item uintptr) int32                                                                // function
	Expand() IList                                                                         // function
	Extract(item uintptr) uintptr                                                          // function
	First() uintptr                                                                        // function
	GetEnumerator() IListEnumerator                                                        // function
	IndexOf(item uintptr) int32                                                            // function
	Last() uintptr                                                                         // function
	Remove(item uintptr) int32                                                             // function
	FPOAttachObserver(observer IObject)                                                    // procedure
	FPODetachObserver(observer IObject)                                                    // procedure
	FPONotifyObservers(sender IObject, operation types.TFPObservedOperation, data uintptr) // procedure
	AddList(list IList)                                                                    // procedure
	Clear()                                                                                // procedure
	Delete(index int32)                                                                    // procedure
	Exchange(index1 int32, index2 int32)                                                   // procedure
	Insert(index int32, item uintptr)                                                      // procedure
	Move(curIndex int32, newIndex int32)                                                   // procedure
	Assign(listA IList, operator types.TListAssignOp, listB IList)                         // procedure
	Pack()                                                                                 // procedure
	Capacity() int32                                                                       // property Capacity Getter
	SetCapacity(value int32)                                                               // property Capacity Setter
	Count() int32                                                                          // property Count Getter
	SetCount(value int32)                                                                  // property Count Setter
	Items(index int32) uintptr                                                             // property Items Getter
	SetItems(index int32, value uintptr)                                                   // property Items Setter
}

type TList struct {
	TObject
}

func (m *TList) Add(item uintptr) int32 {
	if !m.IsValid() {
		return 0
	}
	r := listAPI().SysCallN(1, m.Instance(), uintptr(item))
	return int32(r)
}

func (m *TList) Expand() IList {
	if !m.IsValid() {
		return nil
	}
	r := listAPI().SysCallN(2, m.Instance())
	return AsList(r)
}

func (m *TList) Extract(item uintptr) uintptr {
	if !m.IsValid() {
		return 0
	}
	r := listAPI().SysCallN(3, m.Instance(), uintptr(item))
	return uintptr(r)
}

func (m *TList) First() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := listAPI().SysCallN(4, m.Instance())
	return uintptr(r)
}

func (m *TList) GetEnumerator() IListEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := listAPI().SysCallN(5, m.Instance())
	return AsListEnumerator(r)
}

func (m *TList) IndexOf(item uintptr) int32 {
	if !m.IsValid() {
		return 0
	}
	r := listAPI().SysCallN(6, m.Instance(), uintptr(item))
	return int32(r)
}

func (m *TList) Last() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := listAPI().SysCallN(7, m.Instance())
	return uintptr(r)
}

func (m *TList) Remove(item uintptr) int32 {
	if !m.IsValid() {
		return 0
	}
	r := listAPI().SysCallN(8, m.Instance(), uintptr(item))
	return int32(r)
}

func (m *TList) FPOAttachObserver(observer IObject) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(observer))
}

func (m *TList) FPODetachObserver(observer IObject) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(observer))
}

func (m *TList) FPONotifyObservers(sender IObject, operation types.TFPObservedOperation, data uintptr) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(sender), uintptr(operation), uintptr(data))
}

func (m *TList) AddList(list IList) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(list))
}

func (m *TList) Clear() {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(13, m.Instance())
}

func (m *TList) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(14, m.Instance(), uintptr(index))
}

func (m *TList) Exchange(index1 int32, index2 int32) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(16, m.Instance(), uintptr(index1), uintptr(index2))
}

func (m *TList) Insert(index int32, item uintptr) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(17, m.Instance(), uintptr(index), uintptr(item))
}

func (m *TList) Move(curIndex int32, newIndex int32) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(18, m.Instance(), uintptr(curIndex), uintptr(newIndex))
}

func (m *TList) Assign(listA IList, operator types.TListAssignOp, listB IList) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(listA), uintptr(operator), base.GetObjectUintptr(listB))
}

func (m *TList) Pack() {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(20, m.Instance())
}

func (m *TList) Capacity() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listAPI().SysCallN(21, 0, m.Instance())
	return int32(r)
}

func (m *TList) SetCapacity(value int32) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TList) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TList) SetCount(value int32) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TList) Items(index int32) uintptr {
	if !m.IsValid() {
		return 0
	}
	r := listAPI().SysCallN(23, 0, m.Instance(), uintptr(index))
	return uintptr(r)
}

func (m *TList) SetItems(index int32, value uintptr) {
	if !m.IsValid() {
		return
	}
	listAPI().SysCallN(23, 1, m.Instance(), uintptr(index), uintptr(value))
}

// List  is static instance
var List _ListClass

// _ListClass is class type defined by TList
type _ListClass uintptr

func (_ListClass) Error(msg string, data uintptr) {
	listAPI().SysCallN(15, api.PasStr(msg), uintptr(data))
}

// NewList class constructor
func NewList() IList {
	r := listAPI().SysCallN(0)
	return AsList(r)
}

var (
	listOnce   base.Once
	listImport *imports.Imports = nil
)

func listAPI() *imports.Imports {
	listOnce.Do(func() {
		listImport = api.NewDefaultImports()
		listImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TList_Create", 0), // constructor NewList
			/* 1 */ imports.NewTable("TList_Add", 0), // function Add
			/* 2 */ imports.NewTable("TList_Expand", 0), // function Expand
			/* 3 */ imports.NewTable("TList_Extract", 0), // function Extract
			/* 4 */ imports.NewTable("TList_First", 0), // function First
			/* 5 */ imports.NewTable("TList_GetEnumerator", 0), // function GetEnumerator
			/* 6 */ imports.NewTable("TList_IndexOf", 0), // function IndexOf
			/* 7 */ imports.NewTable("TList_Last", 0), // function Last
			/* 8 */ imports.NewTable("TList_Remove", 0), // function Remove
			/* 9 */ imports.NewTable("TList_FPOAttachObserver", 0), // procedure FPOAttachObserver
			/* 10 */ imports.NewTable("TList_FPODetachObserver", 0), // procedure FPODetachObserver
			/* 11 */ imports.NewTable("TList_FPONotifyObservers", 0), // procedure FPONotifyObservers
			/* 12 */ imports.NewTable("TList_AddList", 0), // procedure AddList
			/* 13 */ imports.NewTable("TList_Clear", 0), // procedure Clear
			/* 14 */ imports.NewTable("TList_Delete", 0), // procedure Delete
			/* 15 */ imports.NewTable("TList_Error", 0), // static procedure Error
			/* 16 */ imports.NewTable("TList_Exchange", 0), // procedure Exchange
			/* 17 */ imports.NewTable("TList_Insert", 0), // procedure Insert
			/* 18 */ imports.NewTable("TList_Move", 0), // procedure Move
			/* 19 */ imports.NewTable("TList_Assign", 0), // procedure Assign
			/* 20 */ imports.NewTable("TList_Pack", 0), // procedure Pack
			/* 21 */ imports.NewTable("TList_Capacity", 0), // property Capacity
			/* 22 */ imports.NewTable("TList_Count", 0), // property Count
			/* 23 */ imports.NewTable("TList_Items", 0), // property Items
		}
	})
	return listImport
}
