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

// ICustomStringGrid Parent: ICustomDrawGrid
type ICustomStringGrid interface {
	ICustomDrawGrid
	Cells(ACol, ARow int32) string                                                                           // property
	SetCells(ACol, ARow int32, AValue string)                                                                // property
	Cols(index int32) IStrings                                                                               // property
	SetCols(index int32, AValue IStrings)                                                                    // property
	DefaultTextStyle() (resultTextStyle TTextStyle)                                                          // property
	SetDefaultTextStyle(AValue *TTextStyle)                                                                  // property
	ExtendedSelect() bool                                                                                    // property
	SetExtendedSelect(AValue bool)                                                                           // property
	Objects(ACol, ARow int32) IObject                                                                        // property
	SetObjects(ACol, ARow int32, AValue IObject)                                                             // property
	Rows(index int32) IStrings                                                                               // property
	SetRows(index int32, AValue IStrings)                                                                    // property
	ValidateOnSetSelection() bool                                                                            // property
	SetValidateOnSetSelection(AValue bool)                                                                   // property
	AutoSizeColumn(aCol int32)                                                                               // procedure
	AutoSizeColumns()                                                                                        // procedure
	Clean()                                                                                                  // procedure
	Clean1(CleanOptions TGridZoneSet)                                                                        // procedure
	Clean2(aRect *TRect, CleanOptions TGridZoneSet)                                                          // procedure
	Clean3(StartCol, StartRow, EndCol, EndRow int32, CleanOptions TGridZoneSet)                              // procedure
	CopyToClipboard(AUseSelection bool)                                                                      // procedure
	LoadFromCSVStream(AStream IStream, ADelimiter Char, UseTitles bool, FromLine int32, SkipEmptyLines bool) // procedure
	LoadFromCSVFile(AFilename string, ADelimiter Char, UseTitles bool, FromLine int32, SkipEmptyLines bool)  // procedure
	SaveToCSVStream(AStream IStream, ADelimiter Char, WriteTitles bool, VisibleColumnsOnly bool)             // procedure
	SaveToCSVFile(AFileName string, ADelimiter Char, WriteTitles bool, VisibleColumnsOnly bool)              // procedure
}

// TCustomStringGrid Parent: TCustomDrawGrid
type TCustomStringGrid struct {
	TCustomDrawGrid
}

func NewCustomStringGrid(AOwner IComponent) ICustomStringGrid {
	r1 := LCL().SysCallN(2281, GetObjectUintptr(AOwner))
	return AsCustomStringGrid(r1)
}

func (m *TCustomStringGrid) Cells(ACol, ARow int32) string {
	r1 := LCL().SysCallN(2273, 0, m.Instance(), uintptr(ACol), uintptr(ARow))
	return GoStr(r1)
}

func (m *TCustomStringGrid) SetCells(ACol, ARow int32, AValue string) {
	LCL().SysCallN(2273, 1, m.Instance(), uintptr(ACol), uintptr(ARow), PascalStr(AValue))
}

func (m *TCustomStringGrid) Cols(index int32) IStrings {
	r1 := LCL().SysCallN(2279, 0, m.Instance(), uintptr(index))
	return AsStrings(r1)
}

func (m *TCustomStringGrid) SetCols(index int32, AValue IStrings) {
	LCL().SysCallN(2279, 1, m.Instance(), uintptr(index), GetObjectUintptr(AValue))
}

func (m *TCustomStringGrid) DefaultTextStyle() (resultTextStyle TTextStyle) {
	LCL().SysCallN(2282, 0, m.Instance(), uintptr(unsafePointer(&resultTextStyle)), uintptr(unsafePointer(&resultTextStyle)))
	return
}

func (m *TCustomStringGrid) SetDefaultTextStyle(AValue *TTextStyle) {
	LCL().SysCallN(2282, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomStringGrid) ExtendedSelect() bool {
	r1 := LCL().SysCallN(2283, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStringGrid) SetExtendedSelect(AValue bool) {
	LCL().SysCallN(2283, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomStringGrid) Objects(ACol, ARow int32) IObject {
	r1 := LCL().SysCallN(2286, 0, m.Instance(), uintptr(ACol), uintptr(ARow))
	return AsObject(r1)
}

func (m *TCustomStringGrid) SetObjects(ACol, ARow int32, AValue IObject) {
	LCL().SysCallN(2286, 1, m.Instance(), uintptr(ACol), uintptr(ARow), GetObjectUintptr(AValue))
}

func (m *TCustomStringGrid) Rows(index int32) IStrings {
	r1 := LCL().SysCallN(2287, 0, m.Instance(), uintptr(index))
	return AsStrings(r1)
}

func (m *TCustomStringGrid) SetRows(index int32, AValue IStrings) {
	LCL().SysCallN(2287, 1, m.Instance(), uintptr(index), GetObjectUintptr(AValue))
}

func (m *TCustomStringGrid) ValidateOnSetSelection() bool {
	r1 := LCL().SysCallN(2290, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStringGrid) SetValidateOnSetSelection(AValue bool) {
	LCL().SysCallN(2290, 1, m.Instance(), PascalBool(AValue))
}

func CustomStringGridClass() TClass {
	ret := LCL().SysCallN(2274)
	return TClass(ret)
}

func (m *TCustomStringGrid) AutoSizeColumn(aCol int32) {
	LCL().SysCallN(2271, m.Instance(), uintptr(aCol))
}

func (m *TCustomStringGrid) AutoSizeColumns() {
	LCL().SysCallN(2272, m.Instance())
}

func (m *TCustomStringGrid) Clean() {
	LCL().SysCallN(2275, m.Instance())
}

func (m *TCustomStringGrid) Clean1(CleanOptions TGridZoneSet) {
	LCL().SysCallN(2276, m.Instance(), uintptr(CleanOptions))
}

func (m *TCustomStringGrid) Clean2(aRect *TRect, CleanOptions TGridZoneSet) {
	LCL().SysCallN(2277, m.Instance(), uintptr(unsafePointer(aRect)), uintptr(CleanOptions))
}

func (m *TCustomStringGrid) Clean3(StartCol, StartRow, EndCol, EndRow int32, CleanOptions TGridZoneSet) {
	LCL().SysCallN(2278, m.Instance(), uintptr(StartCol), uintptr(StartRow), uintptr(EndCol), uintptr(EndRow), uintptr(CleanOptions))
}

func (m *TCustomStringGrid) CopyToClipboard(AUseSelection bool) {
	LCL().SysCallN(2280, m.Instance(), PascalBool(AUseSelection))
}

func (m *TCustomStringGrid) LoadFromCSVStream(AStream IStream, ADelimiter Char, UseTitles bool, FromLine int32, SkipEmptyLines bool) {
	LCL().SysCallN(2285, m.Instance(), GetObjectUintptr(AStream), uintptr(ADelimiter), PascalBool(UseTitles), uintptr(FromLine), PascalBool(SkipEmptyLines))
}

func (m *TCustomStringGrid) LoadFromCSVFile(AFilename string, ADelimiter Char, UseTitles bool, FromLine int32, SkipEmptyLines bool) {
	LCL().SysCallN(2284, m.Instance(), PascalStr(AFilename), uintptr(ADelimiter), PascalBool(UseTitles), uintptr(FromLine), PascalBool(SkipEmptyLines))
}

func (m *TCustomStringGrid) SaveToCSVStream(AStream IStream, ADelimiter Char, WriteTitles bool, VisibleColumnsOnly bool) {
	LCL().SysCallN(2289, m.Instance(), GetObjectUintptr(AStream), uintptr(ADelimiter), PascalBool(WriteTitles), PascalBool(VisibleColumnsOnly))
}

func (m *TCustomStringGrid) SaveToCSVFile(AFileName string, ADelimiter Char, WriteTitles bool, VisibleColumnsOnly bool) {
	LCL().SysCallN(2288, m.Instance(), PascalStr(AFileName), uintptr(ADelimiter), PascalBool(WriteTitles), PascalBool(VisibleColumnsOnly))
}
