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

// IMergedMenuItems Parent: IObject
type IMergedMenuItems interface {
	IObject
	VisibleCount() int32                  // property
	VisibleItems(Index int32) IMenuItem   // property
	InvisibleCount() int32                // property
	InvisibleItems(Index int32) IMenuItem // property
}

// TMergedMenuItems Parent: TObject
type TMergedMenuItems struct {
	TObject
}

func NewMergedMenuItems(aParent IMenuItem) IMergedMenuItems {
	r1 := mergedMenuItemsImportAPI().SysCallN(1, GetObjectUintptr(aParent))
	return AsMergedMenuItems(r1)
}

func (m *TMergedMenuItems) VisibleCount() int32 {
	r1 := mergedMenuItemsImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TMergedMenuItems) VisibleItems(Index int32) IMenuItem {
	r1 := mergedMenuItemsImportAPI().SysCallN(5, m.Instance(), uintptr(Index))
	return AsMenuItem(r1)
}

func (m *TMergedMenuItems) InvisibleCount() int32 {
	r1 := mergedMenuItemsImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TMergedMenuItems) InvisibleItems(Index int32) IMenuItem {
	r1 := mergedMenuItemsImportAPI().SysCallN(3, m.Instance(), uintptr(Index))
	return AsMenuItem(r1)
}

func MergedMenuItemsClass() TClass {
	ret := mergedMenuItemsImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	mergedMenuItemsImport       *imports.Imports = nil
	mergedMenuItemsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("MergedMenuItems_Class", 0),
		/*1*/ imports.NewTable("MergedMenuItems_Create", 0),
		/*2*/ imports.NewTable("MergedMenuItems_InvisibleCount", 0),
		/*3*/ imports.NewTable("MergedMenuItems_InvisibleItems", 0),
		/*4*/ imports.NewTable("MergedMenuItems_VisibleCount", 0),
		/*5*/ imports.NewTable("MergedMenuItems_VisibleItems", 0),
	}
)

func mergedMenuItemsImportAPI() *imports.Imports {
	if mergedMenuItemsImport == nil {
		mergedMenuItemsImport = NewDefaultImports()
		mergedMenuItemsImport.SetImportTable(mergedMenuItemsImportTables)
		mergedMenuItemsImportTables = nil
	}
	return mergedMenuItemsImport
}
