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

// ICustomStringGrid Parent: ICustomDrawGrid
type ICustomStringGrid interface {
	ICustomDrawGrid
	AutoSizeColumn(col int32)                                                                                              // procedure
	AutoSizeColumns()                                                                                                      // procedure
	Clean()                                                                                                                // procedure
	CleanWithGridZoneSet(cleanOptions types.TGridZoneSet)                                                                  // procedure
	CleanWithRectGridZoneSet(rect types.TRect, cleanOptions types.TGridZoneSet)                                            // procedure
	CleanWithIntX4GridZoneSet(startCol int32, startRow int32, endCol int32, endRow int32, cleanOptions types.TGridZoneSet) // procedure
	CopyToClipboard(useSelection bool)                                                                                     // procedure
	LoadFromCSVStream(stream IStream, delimiter uint16, useTitles bool, fromLine int32, skipEmptyLines bool)               // procedure
	LoadFromCSVFile(filename string, delimiter uint16, useTitles bool, fromLine int32, skipEmptyLines bool)                // procedure
	SaveToCSVStream(stream IStream, delimiter uint16, writeTitles bool, visibleColumnsOnly bool)                           // procedure
	SaveToCSVFile(fileName string, delimiter uint16, writeTitles bool, visibleColumnsOnly bool)                            // procedure
	Cells(col int32, row int32) string                                                                                     // property Cells Getter
	SetCells(col int32, row int32, value string)                                                                           // property Cells Setter
	Cols(index int32) IStrings                                                                                             // property Cols Getter
	SetCols(index int32, value IStrings)                                                                                   // property Cols Setter
	DefaultTextStyle() TTextStyle                                                                                          // property DefaultTextStyle Getter
	SetDefaultTextStyle(value TTextStyle)                                                                                  // property DefaultTextStyle Setter
	ExtendedSelect() bool                                                                                                  // property ExtendedSelect Getter
	SetExtendedSelect(value bool)                                                                                          // property ExtendedSelect Setter
	Objects(col int32, row int32) IObject                                                                                  // property Objects Getter
	SetObjects(col int32, row int32, value IObject)                                                                        // property Objects Setter
	Rows(index int32) IStrings                                                                                             // property Rows Getter
	SetRows(index int32, value IStrings)                                                                                   // property Rows Setter
	ValidateOnSetSelection() bool                                                                                          // property ValidateOnSetSelection Getter
	SetValidateOnSetSelection(value bool)                                                                                  // property ValidateOnSetSelection Setter
}

type TCustomStringGrid struct {
	TCustomDrawGrid
}

func (m *TCustomStringGrid) AutoSizeColumn(col int32) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(1, m.Instance(), uintptr(col))
}

func (m *TCustomStringGrid) AutoSizeColumns() {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(2, m.Instance())
}

func (m *TCustomStringGrid) Clean() {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(3, m.Instance())
}

func (m *TCustomStringGrid) CleanWithGridZoneSet(cleanOptions types.TGridZoneSet) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(4, m.Instance(), uintptr(cleanOptions))
}

func (m *TCustomStringGrid) CleanWithRectGridZoneSet(rect types.TRect, cleanOptions types.TGridZoneSet) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&rect)), uintptr(cleanOptions))
}

func (m *TCustomStringGrid) CleanWithIntX4GridZoneSet(startCol int32, startRow int32, endCol int32, endRow int32, cleanOptions types.TGridZoneSet) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(6, m.Instance(), uintptr(startCol), uintptr(startRow), uintptr(endCol), uintptr(endRow), uintptr(cleanOptions))
}

func (m *TCustomStringGrid) CopyToClipboard(useSelection bool) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(7, m.Instance(), api.PasBool(useSelection))
}

func (m *TCustomStringGrid) LoadFromCSVStream(stream IStream, delimiter uint16, useTitles bool, fromLine int32, skipEmptyLines bool) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(stream), uintptr(delimiter), api.PasBool(useTitles), uintptr(fromLine), api.PasBool(skipEmptyLines))
}

func (m *TCustomStringGrid) LoadFromCSVFile(filename string, delimiter uint16, useTitles bool, fromLine int32, skipEmptyLines bool) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(9, m.Instance(), api.PasStr(filename), uintptr(delimiter), api.PasBool(useTitles), uintptr(fromLine), api.PasBool(skipEmptyLines))
}

func (m *TCustomStringGrid) SaveToCSVStream(stream IStream, delimiter uint16, writeTitles bool, visibleColumnsOnly bool) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(stream), uintptr(delimiter), api.PasBool(writeTitles), api.PasBool(visibleColumnsOnly))
}

func (m *TCustomStringGrid) SaveToCSVFile(fileName string, delimiter uint16, writeTitles bool, visibleColumnsOnly bool) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(11, m.Instance(), api.PasStr(fileName), uintptr(delimiter), api.PasBool(writeTitles), api.PasBool(visibleColumnsOnly))
}

func (m *TCustomStringGrid) Cells(col int32, row int32) string {
	if !m.IsValid() {
		return ""
	}
	r := customStringGridAPI().SysCallN(12, 0, m.Instance(), uintptr(col), uintptr(row))
	return api.GoStr(r)
}

func (m *TCustomStringGrid) SetCells(col int32, row int32, value string) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(12, 1, m.Instance(), uintptr(col), uintptr(row), api.PasStr(value))
}

func (m *TCustomStringGrid) Cols(index int32) IStrings {
	if !m.IsValid() {
		return nil
	}
	r := customStringGridAPI().SysCallN(13, 0, m.Instance(), uintptr(index))
	return AsStrings(r)
}

func (m *TCustomStringGrid) SetCols(index int32, value IStrings) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(13, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TCustomStringGrid) DefaultTextStyle() (result TTextStyle) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(14, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomStringGrid) SetDefaultTextStyle(value TTextStyle) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(14, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomStringGrid) ExtendedSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := customStringGridAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomStringGrid) SetExtendedSelect(value bool) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomStringGrid) Objects(col int32, row int32) IObject {
	if !m.IsValid() {
		return nil
	}
	r := customStringGridAPI().SysCallN(16, 0, m.Instance(), uintptr(col), uintptr(row))
	return AsObject(r)
}

func (m *TCustomStringGrid) SetObjects(col int32, row int32, value IObject) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(16, 1, m.Instance(), uintptr(col), uintptr(row), base.GetObjectUintptr(value))
}

func (m *TCustomStringGrid) Rows(index int32) IStrings {
	if !m.IsValid() {
		return nil
	}
	r := customStringGridAPI().SysCallN(17, 0, m.Instance(), uintptr(index))
	return AsStrings(r)
}

func (m *TCustomStringGrid) SetRows(index int32, value IStrings) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(17, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TCustomStringGrid) ValidateOnSetSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := customStringGridAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomStringGrid) SetValidateOnSetSelection(value bool) {
	if !m.IsValid() {
		return
	}
	customStringGridAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

// NewCustomStringGrid class constructor
func NewCustomStringGrid(owner IComponent) ICustomStringGrid {
	r := customStringGridAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomStringGrid(r)
}

func TCustomStringGridClass() types.TClass {
	r := customStringGridAPI().SysCallN(19)
	return types.TClass(r)
}

var (
	customStringGridOnce   base.Once
	customStringGridImport *imports.Imports = nil
)

func customStringGridAPI() *imports.Imports {
	customStringGridOnce.Do(func() {
		customStringGridImport = api.NewDefaultImports()
		customStringGridImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomStringGrid_Create", 0), // constructor NewCustomStringGrid
			/* 1 */ imports.NewTable("TCustomStringGrid_AutoSizeColumn", 0), // procedure AutoSizeColumn
			/* 2 */ imports.NewTable("TCustomStringGrid_AutoSizeColumns", 0), // procedure AutoSizeColumns
			/* 3 */ imports.NewTable("TCustomStringGrid_Clean", 0), // procedure Clean
			/* 4 */ imports.NewTable("TCustomStringGrid_CleanWithGridZoneSet", 0), // procedure CleanWithGridZoneSet
			/* 5 */ imports.NewTable("TCustomStringGrid_CleanWithRectGridZoneSet", 0), // procedure CleanWithRectGridZoneSet
			/* 6 */ imports.NewTable("TCustomStringGrid_CleanWithIntX4GridZoneSet", 0), // procedure CleanWithIntX4GridZoneSet
			/* 7 */ imports.NewTable("TCustomStringGrid_CopyToClipboard", 0), // procedure CopyToClipboard
			/* 8 */ imports.NewTable("TCustomStringGrid_LoadFromCSVStream", 0), // procedure LoadFromCSVStream
			/* 9 */ imports.NewTable("TCustomStringGrid_LoadFromCSVFile", 0), // procedure LoadFromCSVFile
			/* 10 */ imports.NewTable("TCustomStringGrid_SaveToCSVStream", 0), // procedure SaveToCSVStream
			/* 11 */ imports.NewTable("TCustomStringGrid_SaveToCSVFile", 0), // procedure SaveToCSVFile
			/* 12 */ imports.NewTable("TCustomStringGrid_Cells", 0), // property Cells
			/* 13 */ imports.NewTable("TCustomStringGrid_Cols", 0), // property Cols
			/* 14 */ imports.NewTable("TCustomStringGrid_DefaultTextStyle", 0), // property DefaultTextStyle
			/* 15 */ imports.NewTable("TCustomStringGrid_ExtendedSelect", 0), // property ExtendedSelect
			/* 16 */ imports.NewTable("TCustomStringGrid_Objects", 0), // property Objects
			/* 17 */ imports.NewTable("TCustomStringGrid_Rows", 0), // property Rows
			/* 18 */ imports.NewTable("TCustomStringGrid_ValidateOnSetSelection", 0), // property ValidateOnSetSelection
			/* 19 */ imports.NewTable("TCustomStringGrid_TClass", 0), // function TCustomStringGridClass
		}
	})
	return customStringGridImport
}
