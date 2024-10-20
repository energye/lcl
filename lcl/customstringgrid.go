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
	r1 := customStringGridImportAPI().SysCallN(10, GetObjectUintptr(AOwner))
	return AsCustomStringGrid(r1)
}

func (m *TCustomStringGrid) Cells(ACol, ARow int32) string {
	r1 := customStringGridImportAPI().SysCallN(2, 0, m.Instance(), uintptr(ACol), uintptr(ARow))
	return GoStr(r1)
}

func (m *TCustomStringGrid) SetCells(ACol, ARow int32, AValue string) {
	customStringGridImportAPI().SysCallN(2, 1, m.Instance(), uintptr(ACol), uintptr(ARow), PascalStr(AValue))
}

func (m *TCustomStringGrid) Cols(index int32) IStrings {
	r1 := customStringGridImportAPI().SysCallN(8, 0, m.Instance(), uintptr(index))
	return AsStrings(r1)
}

func (m *TCustomStringGrid) SetCols(index int32, AValue IStrings) {
	customStringGridImportAPI().SysCallN(8, 1, m.Instance(), uintptr(index), GetObjectUintptr(AValue))
}

func (m *TCustomStringGrid) DefaultTextStyle() (resultTextStyle TTextStyle) {
	customStringGridImportAPI().SysCallN(11, 0, m.Instance(), uintptr(unsafePointer(&resultTextStyle)), uintptr(unsafePointer(&resultTextStyle)))
	return
}

func (m *TCustomStringGrid) SetDefaultTextStyle(AValue *TTextStyle) {
	customStringGridImportAPI().SysCallN(11, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomStringGrid) ExtendedSelect() bool {
	r1 := customStringGridImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStringGrid) SetExtendedSelect(AValue bool) {
	customStringGridImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomStringGrid) Objects(ACol, ARow int32) IObject {
	r1 := customStringGridImportAPI().SysCallN(15, 0, m.Instance(), uintptr(ACol), uintptr(ARow))
	return AsObject(r1)
}

func (m *TCustomStringGrid) SetObjects(ACol, ARow int32, AValue IObject) {
	customStringGridImportAPI().SysCallN(15, 1, m.Instance(), uintptr(ACol), uintptr(ARow), GetObjectUintptr(AValue))
}

func (m *TCustomStringGrid) Rows(index int32) IStrings {
	r1 := customStringGridImportAPI().SysCallN(16, 0, m.Instance(), uintptr(index))
	return AsStrings(r1)
}

func (m *TCustomStringGrid) SetRows(index int32, AValue IStrings) {
	customStringGridImportAPI().SysCallN(16, 1, m.Instance(), uintptr(index), GetObjectUintptr(AValue))
}

func (m *TCustomStringGrid) ValidateOnSetSelection() bool {
	r1 := customStringGridImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStringGrid) SetValidateOnSetSelection(AValue bool) {
	customStringGridImportAPI().SysCallN(19, 1, m.Instance(), PascalBool(AValue))
}

func CustomStringGridClass() TClass {
	ret := customStringGridImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TCustomStringGrid) AutoSizeColumn(aCol int32) {
	customStringGridImportAPI().SysCallN(0, m.Instance(), uintptr(aCol))
}

func (m *TCustomStringGrid) AutoSizeColumns() {
	customStringGridImportAPI().SysCallN(1, m.Instance())
}

func (m *TCustomStringGrid) Clean() {
	customStringGridImportAPI().SysCallN(4, m.Instance())
}

func (m *TCustomStringGrid) Clean1(CleanOptions TGridZoneSet) {
	customStringGridImportAPI().SysCallN(5, m.Instance(), uintptr(CleanOptions))
}

func (m *TCustomStringGrid) Clean2(aRect *TRect, CleanOptions TGridZoneSet) {
	customStringGridImportAPI().SysCallN(6, m.Instance(), uintptr(unsafePointer(aRect)), uintptr(CleanOptions))
}

func (m *TCustomStringGrid) Clean3(StartCol, StartRow, EndCol, EndRow int32, CleanOptions TGridZoneSet) {
	customStringGridImportAPI().SysCallN(7, m.Instance(), uintptr(StartCol), uintptr(StartRow), uintptr(EndCol), uintptr(EndRow), uintptr(CleanOptions))
}

func (m *TCustomStringGrid) CopyToClipboard(AUseSelection bool) {
	customStringGridImportAPI().SysCallN(9, m.Instance(), PascalBool(AUseSelection))
}

func (m *TCustomStringGrid) LoadFromCSVStream(AStream IStream, ADelimiter Char, UseTitles bool, FromLine int32, SkipEmptyLines bool) {
	customStringGridImportAPI().SysCallN(14, m.Instance(), GetObjectUintptr(AStream), uintptr(ADelimiter), PascalBool(UseTitles), uintptr(FromLine), PascalBool(SkipEmptyLines))
}

func (m *TCustomStringGrid) LoadFromCSVFile(AFilename string, ADelimiter Char, UseTitles bool, FromLine int32, SkipEmptyLines bool) {
	customStringGridImportAPI().SysCallN(13, m.Instance(), PascalStr(AFilename), uintptr(ADelimiter), PascalBool(UseTitles), uintptr(FromLine), PascalBool(SkipEmptyLines))
}

func (m *TCustomStringGrid) SaveToCSVStream(AStream IStream, ADelimiter Char, WriteTitles bool, VisibleColumnsOnly bool) {
	customStringGridImportAPI().SysCallN(18, m.Instance(), GetObjectUintptr(AStream), uintptr(ADelimiter), PascalBool(WriteTitles), PascalBool(VisibleColumnsOnly))
}

func (m *TCustomStringGrid) SaveToCSVFile(AFileName string, ADelimiter Char, WriteTitles bool, VisibleColumnsOnly bool) {
	customStringGridImportAPI().SysCallN(17, m.Instance(), PascalStr(AFileName), uintptr(ADelimiter), PascalBool(WriteTitles), PascalBool(VisibleColumnsOnly))
}

var (
	customStringGridImport       *imports.Imports = nil
	customStringGridImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomStringGrid_AutoSizeColumn", 0),
		/*1*/ imports.NewTable("CustomStringGrid_AutoSizeColumns", 0),
		/*2*/ imports.NewTable("CustomStringGrid_Cells", 0),
		/*3*/ imports.NewTable("CustomStringGrid_Class", 0),
		/*4*/ imports.NewTable("CustomStringGrid_Clean", 0),
		/*5*/ imports.NewTable("CustomStringGrid_Clean1", 0),
		/*6*/ imports.NewTable("CustomStringGrid_Clean2", 0),
		/*7*/ imports.NewTable("CustomStringGrid_Clean3", 0),
		/*8*/ imports.NewTable("CustomStringGrid_Cols", 0),
		/*9*/ imports.NewTable("CustomStringGrid_CopyToClipboard", 0),
		/*10*/ imports.NewTable("CustomStringGrid_Create", 0),
		/*11*/ imports.NewTable("CustomStringGrid_DefaultTextStyle", 0),
		/*12*/ imports.NewTable("CustomStringGrid_ExtendedSelect", 0),
		/*13*/ imports.NewTable("CustomStringGrid_LoadFromCSVFile", 0),
		/*14*/ imports.NewTable("CustomStringGrid_LoadFromCSVStream", 0),
		/*15*/ imports.NewTable("CustomStringGrid_Objects", 0),
		/*16*/ imports.NewTable("CustomStringGrid_Rows", 0),
		/*17*/ imports.NewTable("CustomStringGrid_SaveToCSVFile", 0),
		/*18*/ imports.NewTable("CustomStringGrid_SaveToCSVStream", 0),
		/*19*/ imports.NewTable("CustomStringGrid_ValidateOnSetSelection", 0),
	}
)

func customStringGridImportAPI() *imports.Imports {
	if customStringGridImport == nil {
		customStringGridImport = NewDefaultImports()
		customStringGridImport.SetImportTable(customStringGridImportTables)
		customStringGridImportTables = nil
	}
	return customStringGridImport
}
