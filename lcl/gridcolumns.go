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

// IGridColumns Parent: ICollection
type IGridColumns interface {
	ICollection
	Grid() ICustomGrid                                     // property
	ItemsForGridColumn(Index int32) IGridColumn            // property
	SetItemsForGridColumn(Index int32, AValue IGridColumn) // property
	VisibleCount() int32                                   // property
	Enabled() bool                                         // property
	AddForGridColumn() IGridColumn                         // function
	ColumnByTitle(aTitle string) IGridColumn               // function
	RealIndex(Index int32) int32                           // function
	IndexOf(Column IGridColumn) int32                      // function
	IsDefault() bool                                       // function
	HasIndex(Index int32) bool                             // function
	VisibleIndex(Index int32) int32                        // function
}

// TGridColumns Parent: TCollection
type TGridColumns struct {
	TCollection
}

func NewGridColumns(AGrid ICustomGrid, aItemClass TCollectionItemClass) IGridColumns {
	r1 := gridColumnsImportAPI().SysCallN(3, GetObjectUintptr(AGrid), uintptr(aItemClass))
	return AsGridColumns(r1)
}

func (m *TGridColumns) Grid() ICustomGrid {
	r1 := gridColumnsImportAPI().SysCallN(5, m.Instance())
	return AsCustomGrid(r1)
}

func (m *TGridColumns) ItemsForGridColumn(Index int32) IGridColumn {
	r1 := gridColumnsImportAPI().SysCallN(9, 0, m.Instance(), uintptr(Index))
	return AsGridColumn(r1)
}

func (m *TGridColumns) SetItemsForGridColumn(Index int32, AValue IGridColumn) {
	gridColumnsImportAPI().SysCallN(9, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TGridColumns) VisibleCount() int32 {
	r1 := gridColumnsImportAPI().SysCallN(11, m.Instance())
	return int32(r1)
}

func (m *TGridColumns) Enabled() bool {
	r1 := gridColumnsImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TGridColumns) AddForGridColumn() IGridColumn {
	r1 := gridColumnsImportAPI().SysCallN(0, m.Instance())
	return AsGridColumn(r1)
}

func (m *TGridColumns) ColumnByTitle(aTitle string) IGridColumn {
	r1 := gridColumnsImportAPI().SysCallN(2, m.Instance(), PascalStr(aTitle))
	return AsGridColumn(r1)
}

func (m *TGridColumns) RealIndex(Index int32) int32 {
	r1 := gridColumnsImportAPI().SysCallN(10, m.Instance(), uintptr(Index))
	return int32(r1)
}

func (m *TGridColumns) IndexOf(Column IGridColumn) int32 {
	r1 := gridColumnsImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(Column))
	return int32(r1)
}

func (m *TGridColumns) IsDefault() bool {
	r1 := gridColumnsImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TGridColumns) HasIndex(Index int32) bool {
	r1 := gridColumnsImportAPI().SysCallN(6, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TGridColumns) VisibleIndex(Index int32) int32 {
	r1 := gridColumnsImportAPI().SysCallN(12, m.Instance(), uintptr(Index))
	return int32(r1)
}

func GridColumnsClass() TClass {
	ret := gridColumnsImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	gridColumnsImport       *imports.Imports = nil
	gridColumnsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("GridColumns_AddForGridColumn", 0),
		/*1*/ imports.NewTable("GridColumns_Class", 0),
		/*2*/ imports.NewTable("GridColumns_ColumnByTitle", 0),
		/*3*/ imports.NewTable("GridColumns_Create", 0),
		/*4*/ imports.NewTable("GridColumns_Enabled", 0),
		/*5*/ imports.NewTable("GridColumns_Grid", 0),
		/*6*/ imports.NewTable("GridColumns_HasIndex", 0),
		/*7*/ imports.NewTable("GridColumns_IndexOf", 0),
		/*8*/ imports.NewTable("GridColumns_IsDefault", 0),
		/*9*/ imports.NewTable("GridColumns_ItemsForGridColumn", 0),
		/*10*/ imports.NewTable("GridColumns_RealIndex", 0),
		/*11*/ imports.NewTable("GridColumns_VisibleCount", 0),
		/*12*/ imports.NewTable("GridColumns_VisibleIndex", 0),
	}
)

func gridColumnsImportAPI() *imports.Imports {
	if gridColumnsImport == nil {
		gridColumnsImport = NewDefaultImports()
		gridColumnsImport.SetImportTable(gridColumnsImportTables)
		gridColumnsImportTables = nil
	}
	return gridColumnsImport
}
