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

// ICustomGrid Parent: ICustomControl
type ICustomGrid interface {
	ICustomControl
	CellRect(col int32, row int32) types.TRect                // function
	CellToGridZone(col int32, row int32) types.TGridZone      // function
	ClearCols() bool                                          // function
	ClearRows() bool                                          // function
	EditorByStyle(style types.TColumnButtonStyle) IWinControl // function
	HasMultiSelection() bool                                  // function
	IsCellVisible(col int32, row int32) bool                  // function
	IsFixedCellVisible(col int32, row int32) bool             // function
	MouseCoord(X int32, Y int32) types.TGridCoord             // function
	MouseToCellWithPoint(mouse types.TPoint) types.TPoint     // function
	MouseToLogcell(mouse types.TPoint) types.TPoint           // function
	MouseToGridZone(X int32, Y int32) types.TGridZone         // function
	// AdjustInnerCellRect
	//  Exposed procs
	AdjustInnerCellRect(rect *types.TRect)                               // procedure
	AutoAdjustColumns()                                                  // procedure
	BeginUpdate()                                                        // procedure
	CheckPosition()                                                      // procedure
	Clear()                                                              // procedure
	ClearSelections()                                                    // procedure
	EditorKeyDown(sender IObject, key *uint16, shift types.TShiftState)  // procedure
	EditorKeyPress(sender IObject, key *uint16)                          // procedure
	EditorUTF8KeyPress(sender IObject, uTF8Key *string)                  // procedure
	EditorKeyUp(sender IObject, key *uint16, shift types.TShiftState)    // procedure
	EditorTextChanged(col int32, row int32, text string)                 // procedure
	EndUpdate(refresh bool)                                              // procedure
	HideSortArrow()                                                      // procedure
	InvalidateCell(col int32, row int32)                                 // procedure
	InvalidateCol(col int32)                                             // procedure
	InvalidateRange(range_ types.TRect)                                  // procedure
	InvalidateRow(row int32)                                             // procedure
	LoadFromFile(fileName string)                                        // procedure
	LoadFromStream(stream IStream)                                       // procedure
	MouseToCellWithIntX4(X int32, Y int32, outCol *int32, outRow *int32) // procedure
	SaveToFile(fileName string)                                          // procedure
	SaveToStream(stream IStream)                                         // procedure
	CursorState() types.TGridCursorState                                 // property CursorState Getter
	SelectedRange(index int32) types.TGridRect                           // property SelectedRange Getter
	SelectedRangeCount() int32                                           // property SelectedRangeCount Getter
	SortOrder() types.TSortOrder                                         // property SortOrder Getter
	SetSortOrder(value types.TSortOrder)                                 // property SortOrder Setter
	SortColumn() int32                                                   // property SortColumn Getter
}

type TCustomGrid struct {
	TCustomControl
}

func (m *TCustomGrid) CellRect(col int32, row int32) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(1, m.Instance(), uintptr(col), uintptr(row), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomGrid) CellToGridZone(col int32, row int32) types.TGridZone {
	if !m.IsValid() {
		return 0
	}
	r := customGridAPI().SysCallN(2, m.Instance(), uintptr(col), uintptr(row))
	return types.TGridZone(r)
}

func (m *TCustomGrid) ClearCols() bool {
	if !m.IsValid() {
		return false
	}
	r := customGridAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGrid) ClearRows() bool {
	if !m.IsValid() {
		return false
	}
	r := customGridAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGrid) EditorByStyle(style types.TColumnButtonStyle) IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := customGridAPI().SysCallN(5, m.Instance(), uintptr(style))
	return AsWinControl(r)
}

func (m *TCustomGrid) HasMultiSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := customGridAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomGrid) IsCellVisible(col int32, row int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customGridAPI().SysCallN(7, m.Instance(), uintptr(col), uintptr(row))
	return api.GoBool(r)
}

func (m *TCustomGrid) IsFixedCellVisible(col int32, row int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customGridAPI().SysCallN(8, m.Instance(), uintptr(col), uintptr(row))
	return api.GoBool(r)
}

func (m *TCustomGrid) MouseCoord(X int32, Y int32) (result types.TGridCoord) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(9, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomGrid) MouseToCellWithPoint(mouse types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&mouse)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomGrid) MouseToLogcell(mouse types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&mouse)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomGrid) MouseToGridZone(X int32, Y int32) types.TGridZone {
	if !m.IsValid() {
		return 0
	}
	r := customGridAPI().SysCallN(12, m.Instance(), uintptr(X), uintptr(Y))
	return types.TGridZone(r)
}

func (m *TCustomGrid) AdjustInnerCellRect(rect *types.TRect) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(rect)))
}

func (m *TCustomGrid) AutoAdjustColumns() {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(14, m.Instance())
}

func (m *TCustomGrid) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(15, m.Instance())
}

func (m *TCustomGrid) CheckPosition() {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(16, m.Instance())
}

func (m *TCustomGrid) Clear() {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(17, m.Instance())
}

func (m *TCustomGrid) ClearSelections() {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(18, m.Instance())
}

func (m *TCustomGrid) EditorKeyDown(sender IObject, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	customGridAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(sender), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TCustomGrid) EditorKeyPress(sender IObject, key *uint16) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	customGridAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(sender), uintptr(base.UnsafePointer(&keyPtr)))
	*key = uint16(keyPtr)
}

func (m *TCustomGrid) EditorUTF8KeyPress(sender IObject, uTF8Key *string) {
	if !m.IsValid() {
		return
	}
	uTF8KeyPtr := api.PasStr(*uTF8Key)
	customGridAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(sender), uintptr(base.UnsafePointer(&uTF8KeyPtr)))
	*uTF8Key = api.GoStr(uTF8KeyPtr)
}

func (m *TCustomGrid) EditorKeyUp(sender IObject, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	customGridAPI().SysCallN(22, m.Instance(), base.GetObjectUintptr(sender), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TCustomGrid) EditorTextChanged(col int32, row int32, text string) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(23, m.Instance(), uintptr(col), uintptr(row), api.PasStr(text))
}

func (m *TCustomGrid) EndUpdate(refresh bool) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(24, m.Instance(), api.PasBool(refresh))
}

func (m *TCustomGrid) HideSortArrow() {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(25, m.Instance())
}

func (m *TCustomGrid) InvalidateCell(col int32, row int32) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(26, m.Instance(), uintptr(col), uintptr(row))
}

func (m *TCustomGrid) InvalidateCol(col int32) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(27, m.Instance(), uintptr(col))
}

func (m *TCustomGrid) InvalidateRange(range_ types.TRect) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(28, m.Instance(), uintptr(base.UnsafePointer(&range_)))
}

func (m *TCustomGrid) InvalidateRow(row int32) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(29, m.Instance(), uintptr(row))
}

func (m *TCustomGrid) LoadFromFile(fileName string) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(30, m.Instance(), api.PasStr(fileName))
}

func (m *TCustomGrid) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(31, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCustomGrid) MouseToCellWithIntX4(X int32, Y int32, outCol *int32, outRow *int32) {
	if !m.IsValid() {
		return
	}
	var colPtr uintptr
	var rowPtr uintptr
	customGridAPI().SysCallN(32, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(&colPtr)), uintptr(base.UnsafePointer(&rowPtr)))
	*outCol = int32(colPtr)
	*outRow = int32(rowPtr)
}

func (m *TCustomGrid) SaveToFile(fileName string) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(33, m.Instance(), api.PasStr(fileName))
}

func (m *TCustomGrid) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(34, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCustomGrid) CursorState() types.TGridCursorState {
	if !m.IsValid() {
		return 0
	}
	r := customGridAPI().SysCallN(35, m.Instance())
	return types.TGridCursorState(r)
}

func (m *TCustomGrid) SelectedRange(index int32) (result types.TGridRect) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(36, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomGrid) SelectedRangeCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customGridAPI().SysCallN(37, m.Instance())
	return int32(r)
}

func (m *TCustomGrid) SortOrder() types.TSortOrder {
	if !m.IsValid() {
		return 0
	}
	r := customGridAPI().SysCallN(38, 0, m.Instance())
	return types.TSortOrder(r)
}

func (m *TCustomGrid) SetSortOrder(value types.TSortOrder) {
	if !m.IsValid() {
		return
	}
	customGridAPI().SysCallN(38, 1, m.Instance(), uintptr(value))
}

func (m *TCustomGrid) SortColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customGridAPI().SysCallN(39, m.Instance())
	return int32(r)
}

// NewCustomGrid class constructor
func NewCustomGrid(owner IComponent) ICustomGrid {
	r := customGridAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomGrid(r)
}

func TCustomGridClass() types.TClass {
	r := customGridAPI().SysCallN(40)
	return types.TClass(r)
}

var (
	customGridOnce   base.Once
	customGridImport *imports.Imports = nil
)

func customGridAPI() *imports.Imports {
	customGridOnce.Do(func() {
		customGridImport = api.NewDefaultImports()
		customGridImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomGrid_Create", 0), // constructor NewCustomGrid
			/* 1 */ imports.NewTable("TCustomGrid_CellRect", 0), // function CellRect
			/* 2 */ imports.NewTable("TCustomGrid_CellToGridZone", 0), // function CellToGridZone
			/* 3 */ imports.NewTable("TCustomGrid_ClearCols", 0), // function ClearCols
			/* 4 */ imports.NewTable("TCustomGrid_ClearRows", 0), // function ClearRows
			/* 5 */ imports.NewTable("TCustomGrid_EditorByStyle", 0), // function EditorByStyle
			/* 6 */ imports.NewTable("TCustomGrid_HasMultiSelection", 0), // function HasMultiSelection
			/* 7 */ imports.NewTable("TCustomGrid_IsCellVisible", 0), // function IsCellVisible
			/* 8 */ imports.NewTable("TCustomGrid_IsFixedCellVisible", 0), // function IsFixedCellVisible
			/* 9 */ imports.NewTable("TCustomGrid_MouseCoord", 0), // function MouseCoord
			/* 10 */ imports.NewTable("TCustomGrid_MouseToCellWithPoint", 0), // function MouseToCellWithPoint
			/* 11 */ imports.NewTable("TCustomGrid_MouseToLogcell", 0), // function MouseToLogcell
			/* 12 */ imports.NewTable("TCustomGrid_MouseToGridZone", 0), // function MouseToGridZone
			/* 13 */ imports.NewTable("TCustomGrid_AdjustInnerCellRect", 0), // procedure AdjustInnerCellRect
			/* 14 */ imports.NewTable("TCustomGrid_AutoAdjustColumns", 0), // procedure AutoAdjustColumns
			/* 15 */ imports.NewTable("TCustomGrid_BeginUpdate", 0), // procedure BeginUpdate
			/* 16 */ imports.NewTable("TCustomGrid_CheckPosition", 0), // procedure CheckPosition
			/* 17 */ imports.NewTable("TCustomGrid_Clear", 0), // procedure Clear
			/* 18 */ imports.NewTable("TCustomGrid_ClearSelections", 0), // procedure ClearSelections
			/* 19 */ imports.NewTable("TCustomGrid_EditorKeyDown", 0), // procedure EditorKeyDown
			/* 20 */ imports.NewTable("TCustomGrid_EditorKeyPress", 0), // procedure EditorKeyPress
			/* 21 */ imports.NewTable("TCustomGrid_EditorUTF8KeyPress", 0), // procedure EditorUTF8KeyPress
			/* 22 */ imports.NewTable("TCustomGrid_EditorKeyUp", 0), // procedure EditorKeyUp
			/* 23 */ imports.NewTable("TCustomGrid_EditorTextChanged", 0), // procedure EditorTextChanged
			/* 24 */ imports.NewTable("TCustomGrid_EndUpdate", 0), // procedure EndUpdate
			/* 25 */ imports.NewTable("TCustomGrid_HideSortArrow", 0), // procedure HideSortArrow
			/* 26 */ imports.NewTable("TCustomGrid_InvalidateCell", 0), // procedure InvalidateCell
			/* 27 */ imports.NewTable("TCustomGrid_InvalidateCol", 0), // procedure InvalidateCol
			/* 28 */ imports.NewTable("TCustomGrid_InvalidateRange", 0), // procedure InvalidateRange
			/* 29 */ imports.NewTable("TCustomGrid_InvalidateRow", 0), // procedure InvalidateRow
			/* 30 */ imports.NewTable("TCustomGrid_LoadFromFile", 0), // procedure LoadFromFile
			/* 31 */ imports.NewTable("TCustomGrid_LoadFromStream", 0), // procedure LoadFromStream
			/* 32 */ imports.NewTable("TCustomGrid_MouseToCellWithIntX4", 0), // procedure MouseToCellWithIntX4
			/* 33 */ imports.NewTable("TCustomGrid_SaveToFile", 0), // procedure SaveToFile
			/* 34 */ imports.NewTable("TCustomGrid_SaveToStream", 0), // procedure SaveToStream
			/* 35 */ imports.NewTable("TCustomGrid_CursorState", 0), // property CursorState
			/* 36 */ imports.NewTable("TCustomGrid_SelectedRange", 0), // property SelectedRange
			/* 37 */ imports.NewTable("TCustomGrid_SelectedRangeCount", 0), // property SelectedRangeCount
			/* 38 */ imports.NewTable("TCustomGrid_SortOrder", 0), // property SortOrder
			/* 39 */ imports.NewTable("TCustomGrid_SortColumn", 0), // property SortColumn
			/* 40 */ imports.NewTable("TCustomGrid_TClass", 0), // function TCustomGridClass
		}
	})
	return customGridImport
}
