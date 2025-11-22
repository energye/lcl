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

// IVirtualTreeColumns Parent: ICollection
type IVirtualTreeColumns interface {
	ICollection
	AddToVirtualTreeColumn() IVirtualTreeColumn                                                                  // function
	ColumnFromPositionWithPointBool(P types.TPoint, relative bool) int32                                         // function
	ColumnFromPositionWithColumnPosition(positionIndex uint32) int32                                             // function
	GetFirstVisibleColumn(considerAllowFocus bool) int32                                                         // function
	GetLastVisibleColumn(considerAllowFocus bool) int32                                                          // function
	GetFirstColumn() int32                                                                                       // function
	GetNextColumn(column int32) int32                                                                            // function
	GetNextVisibleColumn(column int32, considerAllowFocus bool) int32                                            // function
	GetPreviousColumn(column int32) int32                                                                        // function
	GetPreviousVisibleColumn(column int32, considerAllowFocus bool) int32                                        // function
	GetScrollWidth() int32                                                                                       // function
	GetVisibleColumns() types.TColumnsArray                                                                      // function
	GetVisibleFixedWidth() int32                                                                                 // function
	IsValidColumn(column int32) bool                                                                             // function
	TotalWidth() int32                                                                                           // function
	AnimatedResize(column int32, newWidth int32)                                                                 // procedure
	Clear()                                                                                                      // procedure
	GetColumnBounds(column int32, outLeft *int32, outRight *int32)                                               // procedure
	LoadFromStream(stream IStream, version int32)                                                                // procedure
	PaintHeaderWithHDCRectInt(dC types.HDC, R types.TRect, hOffset int32)                                        // procedure
	PaintHeaderWithCanvasRectPointInt(targetCanvas ICanvas, R types.TRect, target types.TPoint, rTLOffset int32) // procedure
	SaveToStream(stream IStream)                                                                                 // procedure
	ClickIndex() int32                                                                                           // property ClickIndex Getter
	DefaultWidth() int32                                                                                         // property DefaultWidth Getter
	SetDefaultWidth(value int32)                                                                                 // property DefaultWidth Setter
	ItemsWithColumnIndexToVirtualTreeColumn(index int32) IVirtualTreeColumn                                      // property Items Getter
	SetItemsWithColumnIndexToVirtualTreeColumn(index int32, value IVirtualTreeColumn)                            // property Items Setter
	Header() IVTHeader                                                                                           // property Header Getter
	TrackIndex() int32                                                                                           // property TrackIndex Getter
}

type TVirtualTreeColumns struct {
	TCollection
}

func (m *TVirtualTreeColumns) AddToVirtualTreeColumn() IVirtualTreeColumn {
	if !m.IsValid() {
		return nil
	}
	r := virtualTreeColumnsAPI().SysCallN(1, m.Instance())
	return AsVirtualTreeColumn(r)
}

func (m *TVirtualTreeColumns) ColumnFromPositionWithPointBool(P types.TPoint, relative bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&P)), api.PasBool(relative))
	return int32(r)
}

func (m *TVirtualTreeColumns) ColumnFromPositionWithColumnPosition(positionIndex uint32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(3, m.Instance(), uintptr(positionIndex))
	return int32(r)
}

func (m *TVirtualTreeColumns) GetFirstVisibleColumn(considerAllowFocus bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(4, m.Instance(), api.PasBool(considerAllowFocus))
	return int32(r)
}

func (m *TVirtualTreeColumns) GetLastVisibleColumn(considerAllowFocus bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(5, m.Instance(), api.PasBool(considerAllowFocus))
	return int32(r)
}

func (m *TVirtualTreeColumns) GetFirstColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumns) GetNextColumn(column int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(7, m.Instance(), uintptr(column))
	return int32(r)
}

func (m *TVirtualTreeColumns) GetNextVisibleColumn(column int32, considerAllowFocus bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(8, m.Instance(), uintptr(column), api.PasBool(considerAllowFocus))
	return int32(r)
}

func (m *TVirtualTreeColumns) GetPreviousColumn(column int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(9, m.Instance(), uintptr(column))
	return int32(r)
}

func (m *TVirtualTreeColumns) GetPreviousVisibleColumn(column int32, considerAllowFocus bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(10, m.Instance(), uintptr(column), api.PasBool(considerAllowFocus))
	return int32(r)
}

func (m *TVirtualTreeColumns) GetScrollWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumns) GetVisibleColumns() types.TColumnsArray {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(12, m.Instance())
	return types.TColumnsArray(r)
}

func (m *TVirtualTreeColumns) GetVisibleFixedWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(13, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumns) IsValidColumn(column int32) bool {
	if !m.IsValid() {
		return false
	}
	r := virtualTreeColumnsAPI().SysCallN(14, m.Instance(), uintptr(column))
	return api.GoBool(r)
}

func (m *TVirtualTreeColumns) TotalWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(15, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumns) AnimatedResize(column int32, newWidth int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnsAPI().SysCallN(16, m.Instance(), uintptr(column), uintptr(newWidth))
}

func (m *TVirtualTreeColumns) Clear() {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnsAPI().SysCallN(17, m.Instance())
}

func (m *TVirtualTreeColumns) GetColumnBounds(column int32, outLeft *int32, outRight *int32) {
	if !m.IsValid() {
		return
	}
	var leftPtr uintptr
	var rightPtr uintptr
	virtualTreeColumnsAPI().SysCallN(18, m.Instance(), uintptr(column), uintptr(base.UnsafePointer(&leftPtr)), uintptr(base.UnsafePointer(&rightPtr)))
	*outLeft = int32(leftPtr)
	*outRight = int32(rightPtr)
}

func (m *TVirtualTreeColumns) LoadFromStream(stream IStream, version int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnsAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(stream), uintptr(version))
}

func (m *TVirtualTreeColumns) PaintHeaderWithHDCRectInt(dC types.HDC, R types.TRect, hOffset int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnsAPI().SysCallN(20, m.Instance(), uintptr(dC), uintptr(base.UnsafePointer(&R)), uintptr(hOffset))
}

func (m *TVirtualTreeColumns) PaintHeaderWithCanvasRectPointInt(targetCanvas ICanvas, R types.TRect, target types.TPoint, rTLOffset int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnsAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(targetCanvas), uintptr(base.UnsafePointer(&R)), uintptr(base.UnsafePointer(&target)), uintptr(rTLOffset))
}

func (m *TVirtualTreeColumns) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnsAPI().SysCallN(22, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TVirtualTreeColumns) ClickIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(23, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumns) DefaultWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(24, 0, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumns) SetDefaultWidth(value int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnsAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumns) ItemsWithColumnIndexToVirtualTreeColumn(index int32) IVirtualTreeColumn {
	if !m.IsValid() {
		return nil
	}
	r := virtualTreeColumnsAPI().SysCallN(25, 0, m.Instance(), uintptr(index))
	return AsVirtualTreeColumn(r)
}

func (m *TVirtualTreeColumns) SetItemsWithColumnIndexToVirtualTreeColumn(index int32, value IVirtualTreeColumn) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnsAPI().SysCallN(25, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TVirtualTreeColumns) Header() IVTHeader {
	if !m.IsValid() {
		return nil
	}
	r := virtualTreeColumnsAPI().SysCallN(26, m.Instance())
	return AsVTHeader(r)
}

func (m *TVirtualTreeColumns) TrackIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnsAPI().SysCallN(27, m.Instance())
	return int32(r)
}

// NewVirtualTreeColumns class constructor
func NewVirtualTreeColumns(owner IVTHeader) IVirtualTreeColumns {
	r := virtualTreeColumnsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsVirtualTreeColumns(r)
}

var (
	virtualTreeColumnsOnce   base.Once
	virtualTreeColumnsImport *imports.Imports = nil
)

func virtualTreeColumnsAPI() *imports.Imports {
	virtualTreeColumnsOnce.Do(func() {
		virtualTreeColumnsImport = api.NewDefaultImports()
		virtualTreeColumnsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVirtualTreeColumns_Create", 0), // constructor NewVirtualTreeColumns
			/* 1 */ imports.NewTable("TVirtualTreeColumns_AddToVirtualTreeColumn", 0), // function AddToVirtualTreeColumn
			/* 2 */ imports.NewTable("TVirtualTreeColumns_ColumnFromPositionWithPointBool", 0), // function ColumnFromPositionWithPointBool
			/* 3 */ imports.NewTable("TVirtualTreeColumns_ColumnFromPositionWithColumnPosition", 0), // function ColumnFromPositionWithColumnPosition
			/* 4 */ imports.NewTable("TVirtualTreeColumns_GetFirstVisibleColumn", 0), // function GetFirstVisibleColumn
			/* 5 */ imports.NewTable("TVirtualTreeColumns_GetLastVisibleColumn", 0), // function GetLastVisibleColumn
			/* 6 */ imports.NewTable("TVirtualTreeColumns_GetFirstColumn", 0), // function GetFirstColumn
			/* 7 */ imports.NewTable("TVirtualTreeColumns_GetNextColumn", 0), // function GetNextColumn
			/* 8 */ imports.NewTable("TVirtualTreeColumns_GetNextVisibleColumn", 0), // function GetNextVisibleColumn
			/* 9 */ imports.NewTable("TVirtualTreeColumns_GetPreviousColumn", 0), // function GetPreviousColumn
			/* 10 */ imports.NewTable("TVirtualTreeColumns_GetPreviousVisibleColumn", 0), // function GetPreviousVisibleColumn
			/* 11 */ imports.NewTable("TVirtualTreeColumns_GetScrollWidth", 0), // function GetScrollWidth
			/* 12 */ imports.NewTable("TVirtualTreeColumns_GetVisibleColumns", 0), // function GetVisibleColumns
			/* 13 */ imports.NewTable("TVirtualTreeColumns_GetVisibleFixedWidth", 0), // function GetVisibleFixedWidth
			/* 14 */ imports.NewTable("TVirtualTreeColumns_IsValidColumn", 0), // function IsValidColumn
			/* 15 */ imports.NewTable("TVirtualTreeColumns_TotalWidth", 0), // function TotalWidth
			/* 16 */ imports.NewTable("TVirtualTreeColumns_AnimatedResize", 0), // procedure AnimatedResize
			/* 17 */ imports.NewTable("TVirtualTreeColumns_Clear", 0), // procedure Clear
			/* 18 */ imports.NewTable("TVirtualTreeColumns_GetColumnBounds", 0), // procedure GetColumnBounds
			/* 19 */ imports.NewTable("TVirtualTreeColumns_LoadFromStream", 0), // procedure LoadFromStream
			/* 20 */ imports.NewTable("TVirtualTreeColumns_PaintHeaderWithHDCRectInt", 0), // procedure PaintHeaderWithHDCRectInt
			/* 21 */ imports.NewTable("TVirtualTreeColumns_PaintHeaderWithCanvasRectPointInt", 0), // procedure PaintHeaderWithCanvasRectPointInt
			/* 22 */ imports.NewTable("TVirtualTreeColumns_SaveToStream", 0), // procedure SaveToStream
			/* 23 */ imports.NewTable("TVirtualTreeColumns_ClickIndex", 0), // property ClickIndex
			/* 24 */ imports.NewTable("TVirtualTreeColumns_DefaultWidth", 0), // property DefaultWidth
			/* 25 */ imports.NewTable("TVirtualTreeColumns_ItemsWithColumnIndexToVirtualTreeColumn", 0), // property ItemsWithColumnIndexToVirtualTreeColumn
			/* 26 */ imports.NewTable("TVirtualTreeColumns_Header", 0), // property Header
			/* 27 */ imports.NewTable("TVirtualTreeColumns_TrackIndex", 0), // property TrackIndex
		}
	})
	return virtualTreeColumnsImport
}
