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

// ICustomGrid Parent: ICustomControl
type ICustomGrid interface {
	ICustomControl
	CursorState() TGridCursorState                              // property
	SelectedRange(AIndex int32) (resultGridRect TGridRect)      // property
	SelectedRangeCount() int32                                  // property
	SortOrder() TSortOrder                                      // property
	SetSortOrder(AValue TSortOrder)                             // property
	SortColumn() int32                                          // property
	CellRect(ACol, ARow int32) (resultRect TRect)               // function
	CellToGridZone(aCol, aRow int32) TGridZone                  // function
	ClearCols() bool                                            // function
	ClearRows() bool                                            // function
	EditorByStyle(Style TColumnButtonStyle) IWinControl         // function
	HasMultiSelection() bool                                    // function
	IsCellVisible(aCol, aRow int32) bool                        // function
	IsFixedCellVisible(aCol, aRow int32) bool                   // function
	MouseCoord(X, Y int32) (resultGridCoord TGridCoord)         // function
	MouseToCell(Mouse *TPoint) (resultPoint TPoint)             // function
	MouseToLogcell(Mouse *TPoint) (resultPoint TPoint)          // function
	MouseToGridZone(X, Y int32) TGridZone                       // function
	AdjustInnerCellRect(ARect *TRect)                           // procedure
	AutoAdjustColumns()                                         // procedure
	BeginUpdate()                                               // procedure
	CheckPosition()                                             // procedure
	Clear()                                                     // procedure
	ClearSelections()                                           // procedure
	EditorKeyDown(Sender IObject, Key *Word, Shift TShiftState) // procedure
	EditorKeyPress(Sender IObject, Key *Char)                   // procedure
	EditorUTF8KeyPress(Sender IObject, UTF8Key *TUTF8Char)      // procedure
	EditorKeyUp(Sender IObject, key *Word, shift TShiftState)   // procedure
	EditorTextChanged(aCol, aRow int32, aText string)           // procedure
	EndUpdate(aRefresh bool)                                    // procedure
	HideSortArrow()                                             // procedure
	InvalidateCell(aCol, aRow int32)                            // procedure
	InvalidateCol(ACol int32)                                   // procedure
	InvalidateRange(aRange *TRect)                              // procedure
	InvalidateRow(ARow int32)                                   // procedure
	LoadFromFile(FileName string)                               // procedure
	LoadFromStream(AStream IStream)                             // procedure
	MouseToCell1(X, Y int32, OutCol, OutRow *int32)             // procedure
	SaveToFile(FileName string)                                 // procedure
	SaveToStream(AStream IStream)                               // procedure
}

// TCustomGrid Parent: TCustomControl
type TCustomGrid struct {
	TCustomControl
}

func NewCustomGrid(AOwner IComponent) ICustomGrid {
	r1 := customGridImportAPI().SysCallN(11, GetObjectUintptr(AOwner))
	return AsCustomGrid(r1)
}

func (m *TCustomGrid) CursorState() TGridCursorState {
	r1 := customGridImportAPI().SysCallN(12, m.Instance())
	return TGridCursorState(r1)
}

func (m *TCustomGrid) SelectedRange(AIndex int32) (resultGridRect TGridRect) {
	customGridImportAPI().SysCallN(37, m.Instance(), uintptr(AIndex), uintptr(unsafePointer(&resultGridRect)))
	return
}

func (m *TCustomGrid) SelectedRangeCount() int32 {
	r1 := customGridImportAPI().SysCallN(38, m.Instance())
	return int32(r1)
}

func (m *TCustomGrid) SortOrder() TSortOrder {
	r1 := customGridImportAPI().SysCallN(40, 0, m.Instance(), 0)
	return TSortOrder(r1)
}

func (m *TCustomGrid) SetSortOrder(AValue TSortOrder) {
	customGridImportAPI().SysCallN(40, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomGrid) SortColumn() int32 {
	r1 := customGridImportAPI().SysCallN(39, m.Instance())
	return int32(r1)
}

func (m *TCustomGrid) CellRect(ACol, ARow int32) (resultRect TRect) {
	customGridImportAPI().SysCallN(3, m.Instance(), uintptr(ACol), uintptr(ARow), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCustomGrid) CellToGridZone(aCol, aRow int32) TGridZone {
	r1 := customGridImportAPI().SysCallN(4, m.Instance(), uintptr(aCol), uintptr(aRow))
	return TGridZone(r1)
}

func (m *TCustomGrid) ClearCols() bool {
	r1 := customGridImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCustomGrid) ClearRows() bool {
	r1 := customGridImportAPI().SysCallN(9, m.Instance())
	return GoBool(r1)
}

func (m *TCustomGrid) EditorByStyle(Style TColumnButtonStyle) IWinControl {
	r1 := customGridImportAPI().SysCallN(13, m.Instance(), uintptr(Style))
	return AsWinControl(r1)
}

func (m *TCustomGrid) HasMultiSelection() bool {
	r1 := customGridImportAPI().SysCallN(20, m.Instance())
	return GoBool(r1)
}

func (m *TCustomGrid) IsCellVisible(aCol, aRow int32) bool {
	r1 := customGridImportAPI().SysCallN(26, m.Instance(), uintptr(aCol), uintptr(aRow))
	return GoBool(r1)
}

func (m *TCustomGrid) IsFixedCellVisible(aCol, aRow int32) bool {
	r1 := customGridImportAPI().SysCallN(27, m.Instance(), uintptr(aCol), uintptr(aRow))
	return GoBool(r1)
}

func (m *TCustomGrid) MouseCoord(X, Y int32) (resultGridCoord TGridCoord) {
	customGridImportAPI().SysCallN(30, m.Instance(), uintptr(X), uintptr(Y), uintptr(unsafePointer(&resultGridCoord)))
	return
}

func (m *TCustomGrid) MouseToCell(Mouse *TPoint) (resultPoint TPoint) {
	customGridImportAPI().SysCallN(31, m.Instance(), uintptr(unsafePointer(Mouse)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomGrid) MouseToLogcell(Mouse *TPoint) (resultPoint TPoint) {
	customGridImportAPI().SysCallN(34, m.Instance(), uintptr(unsafePointer(Mouse)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomGrid) MouseToGridZone(X, Y int32) TGridZone {
	r1 := customGridImportAPI().SysCallN(33, m.Instance(), uintptr(X), uintptr(Y))
	return TGridZone(r1)
}

func CustomGridClass() TClass {
	ret := customGridImportAPI().SysCallN(6)
	return TClass(ret)
}

func (m *TCustomGrid) AdjustInnerCellRect(ARect *TRect) {
	var result0 uintptr
	customGridImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&result0)))
	*ARect = *(*TRect)(getPointer(result0))
}

func (m *TCustomGrid) AutoAdjustColumns() {
	customGridImportAPI().SysCallN(1, m.Instance())
}

func (m *TCustomGrid) BeginUpdate() {
	customGridImportAPI().SysCallN(2, m.Instance())
}

func (m *TCustomGrid) CheckPosition() {
	customGridImportAPI().SysCallN(5, m.Instance())
}

func (m *TCustomGrid) Clear() {
	customGridImportAPI().SysCallN(7, m.Instance())
}

func (m *TCustomGrid) ClearSelections() {
	customGridImportAPI().SysCallN(10, m.Instance())
}

func (m *TCustomGrid) EditorKeyDown(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	customGridImportAPI().SysCallN(14, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TCustomGrid) EditorKeyPress(Sender IObject, Key *Char) {
	var result1 uintptr
	customGridImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)))
	*Key = Char(result1)
}

func (m *TCustomGrid) EditorUTF8KeyPress(Sender IObject, UTF8Key *TUTF8Char) {
	var result1 uintptr
	customGridImportAPI().SysCallN(18, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)))
	*UTF8Key = *(*TUTF8Char)(getPointer(result1))
}

func (m *TCustomGrid) EditorKeyUp(Sender IObject, key *Word, shift TShiftState) {
	var result1 uintptr
	customGridImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)), uintptr(shift))
	*key = Word(result1)
}

func (m *TCustomGrid) EditorTextChanged(aCol, aRow int32, aText string) {
	customGridImportAPI().SysCallN(17, m.Instance(), uintptr(aCol), uintptr(aRow), PascalStr(aText))
}

func (m *TCustomGrid) EndUpdate(aRefresh bool) {
	customGridImportAPI().SysCallN(19, m.Instance(), PascalBool(aRefresh))
}

func (m *TCustomGrid) HideSortArrow() {
	customGridImportAPI().SysCallN(21, m.Instance())
}

func (m *TCustomGrid) InvalidateCell(aCol, aRow int32) {
	customGridImportAPI().SysCallN(22, m.Instance(), uintptr(aCol), uintptr(aRow))
}

func (m *TCustomGrid) InvalidateCol(ACol int32) {
	customGridImportAPI().SysCallN(23, m.Instance(), uintptr(ACol))
}

func (m *TCustomGrid) InvalidateRange(aRange *TRect) {
	customGridImportAPI().SysCallN(24, m.Instance(), uintptr(unsafePointer(aRange)))
}

func (m *TCustomGrid) InvalidateRow(ARow int32) {
	customGridImportAPI().SysCallN(25, m.Instance(), uintptr(ARow))
}

func (m *TCustomGrid) LoadFromFile(FileName string) {
	customGridImportAPI().SysCallN(28, m.Instance(), PascalStr(FileName))
}

func (m *TCustomGrid) LoadFromStream(AStream IStream) {
	customGridImportAPI().SysCallN(29, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomGrid) MouseToCell1(X, Y int32, OutCol, OutRow *int32) {
	var result1 uintptr
	var result2 uintptr
	customGridImportAPI().SysCallN(32, m.Instance(), uintptr(X), uintptr(Y), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutCol = int32(result1)
	*OutRow = int32(result2)
}

func (m *TCustomGrid) SaveToFile(FileName string) {
	customGridImportAPI().SysCallN(35, m.Instance(), PascalStr(FileName))
}

func (m *TCustomGrid) SaveToStream(AStream IStream) {
	customGridImportAPI().SysCallN(36, m.Instance(), GetObjectUintptr(AStream))
}

var (
	customGridImport       *imports.Imports = nil
	customGridImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomGrid_AdjustInnerCellRect", 0),
		/*1*/ imports.NewTable("CustomGrid_AutoAdjustColumns", 0),
		/*2*/ imports.NewTable("CustomGrid_BeginUpdate", 0),
		/*3*/ imports.NewTable("CustomGrid_CellRect", 0),
		/*4*/ imports.NewTable("CustomGrid_CellToGridZone", 0),
		/*5*/ imports.NewTable("CustomGrid_CheckPosition", 0),
		/*6*/ imports.NewTable("CustomGrid_Class", 0),
		/*7*/ imports.NewTable("CustomGrid_Clear", 0),
		/*8*/ imports.NewTable("CustomGrid_ClearCols", 0),
		/*9*/ imports.NewTable("CustomGrid_ClearRows", 0),
		/*10*/ imports.NewTable("CustomGrid_ClearSelections", 0),
		/*11*/ imports.NewTable("CustomGrid_Create", 0),
		/*12*/ imports.NewTable("CustomGrid_CursorState", 0),
		/*13*/ imports.NewTable("CustomGrid_EditorByStyle", 0),
		/*14*/ imports.NewTable("CustomGrid_EditorKeyDown", 0),
		/*15*/ imports.NewTable("CustomGrid_EditorKeyPress", 0),
		/*16*/ imports.NewTable("CustomGrid_EditorKeyUp", 0),
		/*17*/ imports.NewTable("CustomGrid_EditorTextChanged", 0),
		/*18*/ imports.NewTable("CustomGrid_EditorUTF8KeyPress", 0),
		/*19*/ imports.NewTable("CustomGrid_EndUpdate", 0),
		/*20*/ imports.NewTable("CustomGrid_HasMultiSelection", 0),
		/*21*/ imports.NewTable("CustomGrid_HideSortArrow", 0),
		/*22*/ imports.NewTable("CustomGrid_InvalidateCell", 0),
		/*23*/ imports.NewTable("CustomGrid_InvalidateCol", 0),
		/*24*/ imports.NewTable("CustomGrid_InvalidateRange", 0),
		/*25*/ imports.NewTable("CustomGrid_InvalidateRow", 0),
		/*26*/ imports.NewTable("CustomGrid_IsCellVisible", 0),
		/*27*/ imports.NewTable("CustomGrid_IsFixedCellVisible", 0),
		/*28*/ imports.NewTable("CustomGrid_LoadFromFile", 0),
		/*29*/ imports.NewTable("CustomGrid_LoadFromStream", 0),
		/*30*/ imports.NewTable("CustomGrid_MouseCoord", 0),
		/*31*/ imports.NewTable("CustomGrid_MouseToCell", 0),
		/*32*/ imports.NewTable("CustomGrid_MouseToCell1", 0),
		/*33*/ imports.NewTable("CustomGrid_MouseToGridZone", 0),
		/*34*/ imports.NewTable("CustomGrid_MouseToLogcell", 0),
		/*35*/ imports.NewTable("CustomGrid_SaveToFile", 0),
		/*36*/ imports.NewTable("CustomGrid_SaveToStream", 0),
		/*37*/ imports.NewTable("CustomGrid_SelectedRange", 0),
		/*38*/ imports.NewTable("CustomGrid_SelectedRangeCount", 0),
		/*39*/ imports.NewTable("CustomGrid_SortColumn", 0),
		/*40*/ imports.NewTable("CustomGrid_SortOrder", 0),
	}
)

func customGridImportAPI() *imports.Imports {
	if customGridImport == nil {
		customGridImport = NewDefaultImports()
		customGridImport.SetImportTable(customGridImportTables)
		customGridImportTables = nil
	}
	return customGridImport
}
