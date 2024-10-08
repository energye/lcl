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

// IListControlItems Parent: IOwnedCollection
type IListControlItems interface {
	IOwnedCollection
	ItemsForListControlItem(AIndex int32) IListControlItem // property
	CaseSensitive() bool                                   // property
	SetCaseSensitive(AValue bool)                          // property
	SortType() TListItemsSortType                          // property
	SetSortType(AValue TListItemsSortType)                 // property
	AddForListControlItem() IListControlItem               // function
	CustomSort(fn TListItemsCompare)                       // procedure
	SortForOverload()                                      // procedure
	SetOnCompare(fn TListCompareEvent)                     // property event
}

// TListControlItems Parent: TOwnedCollection
type TListControlItems struct {
	TOwnedCollection
	customSortPtr uintptr
	comparePtr    uintptr
}

func NewListControlItems(AOwner IPersistent, AItemClass TCollectionItemClass) IListControlItems {
	r1 := LCL().SysCallN(4041, GetObjectUintptr(AOwner), uintptr(AItemClass))
	return AsListControlItems(r1)
}

func (m *TListControlItems) ItemsForListControlItem(AIndex int32) IListControlItem {
	r1 := LCL().SysCallN(4043, m.Instance(), uintptr(AIndex))
	return AsListControlItem(r1)
}

func (m *TListControlItems) CaseSensitive() bool {
	r1 := LCL().SysCallN(4039, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListControlItems) SetCaseSensitive(AValue bool) {
	LCL().SysCallN(4039, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListControlItems) SortType() TListItemsSortType {
	r1 := LCL().SysCallN(4046, 0, m.Instance(), 0)
	return TListItemsSortType(r1)
}

func (m *TListControlItems) SetSortType(AValue TListItemsSortType) {
	LCL().SysCallN(4046, 1, m.Instance(), uintptr(AValue))
}

func (m *TListControlItems) AddForListControlItem() IListControlItem {
	r1 := LCL().SysCallN(4038, m.Instance())
	return AsListControlItem(r1)
}

func ListControlItemsClass() TClass {
	ret := LCL().SysCallN(4040)
	return TClass(ret)
}

func (m *TListControlItems) CustomSort(fn TListItemsCompare) {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4042, m.Instance(), m.customSortPtr)
}

func (m *TListControlItems) SortForOverload() {
	LCL().SysCallN(4045, m.Instance())
}

func (m *TListControlItems) SetOnCompare(fn TListCompareEvent) {
	if m.comparePtr != 0 {
		RemoveEventElement(m.comparePtr)
	}
	m.comparePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4044, m.Instance(), m.comparePtr)
}
