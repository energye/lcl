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

// IGridColumns Parent: ICollection
type IGridColumns interface {
	ICollection
	AddToGridColumn() IGridColumn                               // function
	ColumnByTitle(title string) IGridColumn                     // function
	RealIndex(index int32) int32                                // function
	IndexOf(column IGridColumn) int32                           // function
	IsDefault() bool                                            // function
	HasIndex(index int32) bool                                  // function
	VisibleIndex(index int32) int32                             // function
	Clear()                                                     // procedure
	Grid() ICustomGrid                                          // property Grid Getter
	ItemsWithIntToGridColumn(index int32) IGridColumn           // property Items Getter
	SetItemsWithIntToGridColumn(index int32, value IGridColumn) // property Items Setter
	VisibleCount() int32                                        // property VisibleCount Getter
	Enabled() bool                                              // property Enabled Getter
}

type TGridColumns struct {
	TCollection
}

func (m *TGridColumns) AddToGridColumn() IGridColumn {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnsAPI().SysCallN(1, m.Instance())
	return AsGridColumn(r)
}

func (m *TGridColumns) ColumnByTitle(title string) IGridColumn {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnsAPI().SysCallN(2, m.Instance(), api.PasStr(title))
	return AsGridColumn(r)
}

func (m *TGridColumns) RealIndex(index int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnsAPI().SysCallN(3, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TGridColumns) IndexOf(column IGridColumn) int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnsAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(column))
	return int32(r)
}

func (m *TGridColumns) IsDefault() bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnsAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TGridColumns) HasIndex(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnsAPI().SysCallN(6, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TGridColumns) VisibleIndex(index int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnsAPI().SysCallN(7, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TGridColumns) Clear() {
	if !m.IsValid() {
		return
	}
	gridColumnsAPI().SysCallN(8, m.Instance())
}

func (m *TGridColumns) Grid() ICustomGrid {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnsAPI().SysCallN(9, m.Instance())
	return AsCustomGrid(r)
}

func (m *TGridColumns) ItemsWithIntToGridColumn(index int32) IGridColumn {
	if !m.IsValid() {
		return nil
	}
	r := gridColumnsAPI().SysCallN(10, 0, m.Instance(), uintptr(index))
	return AsGridColumn(r)
}

func (m *TGridColumns) SetItemsWithIntToGridColumn(index int32, value IGridColumn) {
	if !m.IsValid() {
		return
	}
	gridColumnsAPI().SysCallN(10, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TGridColumns) VisibleCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := gridColumnsAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TGridColumns) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := gridColumnsAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

// NewGridColumns class constructor
func NewGridColumns(grid ICustomGrid, itemClass types.TCollectionItemClass) IGridColumns {
	r := gridColumnsAPI().SysCallN(0, base.GetObjectUintptr(grid), uintptr(itemClass))
	return AsGridColumns(r)
}

var (
	gridColumnsOnce   base.Once
	gridColumnsImport *imports.Imports = nil
)

func gridColumnsAPI() *imports.Imports {
	gridColumnsOnce.Do(func() {
		gridColumnsImport = api.NewDefaultImports()
		gridColumnsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGridColumns_Create", 0), // constructor NewGridColumns
			/* 1 */ imports.NewTable("TGridColumns_AddToGridColumn", 0), // function AddToGridColumn
			/* 2 */ imports.NewTable("TGridColumns_ColumnByTitle", 0), // function ColumnByTitle
			/* 3 */ imports.NewTable("TGridColumns_RealIndex", 0), // function RealIndex
			/* 4 */ imports.NewTable("TGridColumns_IndexOf", 0), // function IndexOf
			/* 5 */ imports.NewTable("TGridColumns_IsDefault", 0), // function IsDefault
			/* 6 */ imports.NewTable("TGridColumns_HasIndex", 0), // function HasIndex
			/* 7 */ imports.NewTable("TGridColumns_VisibleIndex", 0), // function VisibleIndex
			/* 8 */ imports.NewTable("TGridColumns_Clear", 0), // procedure Clear
			/* 9 */ imports.NewTable("TGridColumns_Grid", 0), // property Grid
			/* 10 */ imports.NewTable("TGridColumns_ItemsWithIntToGridColumn", 0), // property ItemsWithIntToGridColumn
			/* 11 */ imports.NewTable("TGridColumns_VisibleCount", 0), // property VisibleCount
			/* 12 */ imports.NewTable("TGridColumns_Enabled", 0), // property Enabled
		}
	})
	return gridColumnsImport
}
