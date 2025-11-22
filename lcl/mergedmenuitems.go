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
)

// IMergedMenuItems Parent: IObject
type IMergedMenuItems interface {
	IObject
	VisibleCount() int32                  // property VisibleCount Getter
	VisibleItems(index int32) IMenuItem   // property VisibleItems Getter
	InvisibleCount() int32                // property InvisibleCount Getter
	InvisibleItems(index int32) IMenuItem // property InvisibleItems Getter
}

type TMergedMenuItems struct {
	TObject
}

func (m *TMergedMenuItems) VisibleCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := mergedMenuItemsAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TMergedMenuItems) VisibleItems(index int32) IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := mergedMenuItemsAPI().SysCallN(3, m.Instance(), uintptr(index))
	return AsMenuItem(r)
}

func (m *TMergedMenuItems) InvisibleCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := mergedMenuItemsAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TMergedMenuItems) InvisibleItems(index int32) IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := mergedMenuItemsAPI().SysCallN(5, m.Instance(), uintptr(index))
	return AsMenuItem(r)
}

// MergedMenuItems  is static instance
var MergedMenuItems _MergedMenuItemsClass

// _MergedMenuItemsClass is class type defined by TMergedMenuItems
type _MergedMenuItemsClass uintptr

func (_MergedMenuItemsClass) DefaultSort(item1 uintptr, item2 uintptr, parentItem uintptr) int32 {
	r := mergedMenuItemsAPI().SysCallN(1, uintptr(item1), uintptr(item2), uintptr(parentItem))
	return int32(r)
}

// NewMergedMenuItems class constructor
func NewMergedMenuItems(parent IMenuItem) IMergedMenuItems {
	r := mergedMenuItemsAPI().SysCallN(0, base.GetObjectUintptr(parent))
	return AsMergedMenuItems(r)
}

var (
	mergedMenuItemsOnce   base.Once
	mergedMenuItemsImport *imports.Imports = nil
)

func mergedMenuItemsAPI() *imports.Imports {
	mergedMenuItemsOnce.Do(func() {
		mergedMenuItemsImport = api.NewDefaultImports()
		mergedMenuItemsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMergedMenuItems_Create", 0), // constructor NewMergedMenuItems
			/* 1 */ imports.NewTable("TMergedMenuItems_DefaultSort", 0), // static function DefaultSort
			/* 2 */ imports.NewTable("TMergedMenuItems_VisibleCount", 0), // property VisibleCount
			/* 3 */ imports.NewTable("TMergedMenuItems_VisibleItems", 0), // property VisibleItems
			/* 4 */ imports.NewTable("TMergedMenuItems_InvisibleCount", 0), // property InvisibleCount
			/* 5 */ imports.NewTable("TMergedMenuItems_InvisibleItems", 0), // property InvisibleItems
		}
	})
	return mergedMenuItemsImport
}
