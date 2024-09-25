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
	r1 := LCL().SysCallN(3282, GetObjectUintptr(AGrid), uintptr(aItemClass))
	return AsGridColumns(r1)
}

func (m *TGridColumns) Grid() ICustomGrid {
	r1 := LCL().SysCallN(3284, m.Instance())
	return AsCustomGrid(r1)
}

func (m *TGridColumns) ItemsForGridColumn(Index int32) IGridColumn {
	r1 := LCL().SysCallN(3288, 0, m.Instance(), uintptr(Index))
	return AsGridColumn(r1)
}

func (m *TGridColumns) SetItemsForGridColumn(Index int32, AValue IGridColumn) {
	LCL().SysCallN(3288, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TGridColumns) VisibleCount() int32 {
	r1 := LCL().SysCallN(3290, m.Instance())
	return int32(r1)
}

func (m *TGridColumns) Enabled() bool {
	r1 := LCL().SysCallN(3283, m.Instance())
	return GoBool(r1)
}

func (m *TGridColumns) AddForGridColumn() IGridColumn {
	r1 := LCL().SysCallN(3279, m.Instance())
	return AsGridColumn(r1)
}

func (m *TGridColumns) ColumnByTitle(aTitle string) IGridColumn {
	r1 := LCL().SysCallN(3281, m.Instance(), PascalStr(aTitle))
	return AsGridColumn(r1)
}

func (m *TGridColumns) RealIndex(Index int32) int32 {
	r1 := LCL().SysCallN(3289, m.Instance(), uintptr(Index))
	return int32(r1)
}

func (m *TGridColumns) IndexOf(Column IGridColumn) int32 {
	r1 := LCL().SysCallN(3286, m.Instance(), GetObjectUintptr(Column))
	return int32(r1)
}

func (m *TGridColumns) IsDefault() bool {
	r1 := LCL().SysCallN(3287, m.Instance())
	return GoBool(r1)
}

func (m *TGridColumns) HasIndex(Index int32) bool {
	r1 := LCL().SysCallN(3285, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TGridColumns) VisibleIndex(Index int32) int32 {
	r1 := LCL().SysCallN(3291, m.Instance(), uintptr(Index))
	return int32(r1)
}

func GridColumnsClass() TClass {
	ret := LCL().SysCallN(3280)
	return TClass(ret)
}
