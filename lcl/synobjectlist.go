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

// ISynObjectList Parent: IComponent
type ISynObjectList interface {
	IComponent
	Add(anItem ISynObjectListItem) int32                // function
	Count() int32                                       // function
	IndexOf(anItem ISynObjectListItem) int32            // function
	Delete(index int32)                                 // procedure
	Clear()                                             // procedure
	Move(old int32, new int32)                          // procedure
	Sort()                                              // procedure
	Sorted() bool                                       // property Sorted Getter
	SetSorted(value bool)                               // property Sorted Setter
	BaseItems(index int32) ISynObjectListItem           // property BaseItems Getter
	SetBaseItems(index int32, value ISynObjectListItem) // property BaseItems Setter
	SetOnChange(fn TNotifyEvent)                        // property event
}

type TSynObjectList struct {
	TComponent
}

func (m *TSynObjectList) Add(anItem ISynObjectListItem) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synObjectListAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(anItem))
	return int32(r)
}

func (m *TSynObjectList) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synObjectListAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TSynObjectList) IndexOf(anItem ISynObjectListItem) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synObjectListAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(anItem))
	return int32(r)
}

func (m *TSynObjectList) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	synObjectListAPI().SysCallN(4, m.Instance(), uintptr(index))
}

func (m *TSynObjectList) Clear() {
	if !m.IsValid() {
		return
	}
	synObjectListAPI().SysCallN(5, m.Instance())
}

func (m *TSynObjectList) Move(old int32, new int32) {
	if !m.IsValid() {
		return
	}
	synObjectListAPI().SysCallN(6, m.Instance(), uintptr(old), uintptr(new))
}

func (m *TSynObjectList) Sort() {
	if !m.IsValid() {
		return
	}
	synObjectListAPI().SysCallN(7, m.Instance())
}

func (m *TSynObjectList) Sorted() bool {
	if !m.IsValid() {
		return false
	}
	r := synObjectListAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynObjectList) SetSorted(value bool) {
	if !m.IsValid() {
		return
	}
	synObjectListAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynObjectList) BaseItems(index int32) ISynObjectListItem {
	if !m.IsValid() {
		return nil
	}
	r := synObjectListAPI().SysCallN(9, 0, m.Instance(), uintptr(index))
	return AsSynObjectListItem(r)
}

func (m *TSynObjectList) SetBaseItems(index int32, value ISynObjectListItem) {
	if !m.IsValid() {
		return
	}
	synObjectListAPI().SysCallN(9, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TSynObjectList) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, synObjectListAPI(), api.MakeEventDataPtr(cb))
}

// NewSynObjectList class constructor
func NewSynObjectList(owner IComponent) ISynObjectList {
	r := synObjectListAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynObjectList(r)
}

func TSynObjectListClass() types.TClass {
	r := synObjectListAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	synObjectListOnce   base.Once
	synObjectListImport *imports.Imports = nil
)

func synObjectListAPI() *imports.Imports {
	synObjectListOnce.Do(func() {
		synObjectListImport = api.NewDefaultImports()
		synObjectListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynObjectList_Create", 0), // constructor NewSynObjectList
			/* 1 */ imports.NewTable("TSynObjectList_Add", 0), // function Add
			/* 2 */ imports.NewTable("TSynObjectList_Count", 0), // function Count
			/* 3 */ imports.NewTable("TSynObjectList_IndexOf", 0), // function IndexOf
			/* 4 */ imports.NewTable("TSynObjectList_Delete", 0), // procedure Delete
			/* 5 */ imports.NewTable("TSynObjectList_Clear", 0), // procedure Clear
			/* 6 */ imports.NewTable("TSynObjectList_Move", 0), // procedure Move
			/* 7 */ imports.NewTable("TSynObjectList_Sort", 0), // procedure Sort
			/* 8 */ imports.NewTable("TSynObjectList_Sorted", 0), // property Sorted
			/* 9 */ imports.NewTable("TSynObjectList_BaseItems", 0), // property BaseItems
			/* 10 */ imports.NewTable("TSynObjectList_OnChange", 0), // event OnChange
			/* 11 */ imports.NewTable("TSynObjectList_TClass", 0), // function TSynObjectListClass
		}
	})
	return synObjectListImport
}
