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

// IVirtualTreeColumns Parent: ICollection
type IVirtualTreeColumns interface {
	ICollection
	GetVisibleColumns() *TColumnsArray
	ClickIndex() TColumnIndex                                                           // property
	DefaultWidth() int32                                                                // property
	SetDefaultWidth(AValue int32)                                                       // property
	ItemsForVirtualTreeColumn(Index TColumnIndex) IVirtualTreeColumn                    // property
	SetItemsForVirtualTreeColumn(Index TColumnIndex, AValue IVirtualTreeColumn)         // property
	Header() IVTHeader                                                                  // property
	TrackIndex() TColumnIndex                                                           // property
	AddForVirtualTreeColumn() IVirtualTreeColumn                                        // function
	ColumnFromPosition(P *TPoint, Relative bool) TColumnIndex                           // function
	ColumnFromPosition1(PositionIndex TColumnPosition) TColumnIndex                     // function
	GetFirstVisibleColumn(ConsiderAllowFocus bool) TColumnIndex                         // function
	GetLastVisibleColumn(ConsiderAllowFocus bool) TColumnIndex                          // function
	GetFirstColumn() TColumnIndex                                                       // function
	GetNextColumn(Column TColumnIndex) TColumnIndex                                     // function
	GetNextVisibleColumn(Column TColumnIndex, ConsiderAllowFocus bool) TColumnIndex     // function
	GetPreviousColumn(Column TColumnIndex) TColumnIndex                                 // function
	GetPreviousVisibleColumn(Column TColumnIndex, ConsiderAllowFocus bool) TColumnIndex // function
	GetScrollWidth() int32                                                              // function
	GetVisibleFixedWidth() int32                                                        // function
	IsValidColumn(Column TColumnIndex) bool                                             // function
	TotalWidth() int32                                                                  // function
	AnimatedResize(Column TColumnIndex, NewWidth int32)                                 // procedure
	GetColumnBounds(Column TColumnIndex, OutLeft, OutRight *int32)                      // procedure
	LoadFromStream(Stream IStream, Version int32)                                       // procedure
	PaintHeader(DC HDC, R *TRect, HOffset int32)                                        // procedure
	PaintHeader1(TargetCanvas ICanvas, R *TRect, Target *TPoint, RTLOffset int32)       // procedure
	SaveToStream(Stream IStream)                                                        // procedure
}

// TVirtualTreeColumns Parent: TCollection
type TVirtualTreeColumns struct {
	TCollection
}

func NewVirtualTreeColumns(AOwner IVTHeader) IVirtualTreeColumns {
	r1 := virtualTreeColumnsImportAPI().SysCallN(6, GetObjectUintptr(AOwner))
	return AsVirtualTreeColumns(r1)
}

func (m *TVirtualTreeColumns) ClickIndex() TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(3, m.Instance())
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) DefaultWidth() int32 {
	r1 := virtualTreeColumnsImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumns) SetDefaultWidth(AValue int32) {
	virtualTreeColumnsImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumns) ItemsForVirtualTreeColumn(Index TColumnIndex) IVirtualTreeColumn {
	r1 := virtualTreeColumnsImportAPI().SysCallN(20, 0, m.Instance(), uintptr(Index))
	return AsVirtualTreeColumn(r1)
}

func (m *TVirtualTreeColumns) SetItemsForVirtualTreeColumn(Index TColumnIndex, AValue IVirtualTreeColumn) {
	virtualTreeColumnsImportAPI().SysCallN(20, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TVirtualTreeColumns) Header() IVTHeader {
	r1 := virtualTreeColumnsImportAPI().SysCallN(18, m.Instance())
	return AsVTHeader(r1)
}

func (m *TVirtualTreeColumns) TrackIndex() TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(26, m.Instance())
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) AddForVirtualTreeColumn() IVirtualTreeColumn {
	r1 := virtualTreeColumnsImportAPI().SysCallN(0, m.Instance())
	return AsVirtualTreeColumn(r1)
}

func (m *TVirtualTreeColumns) ColumnFromPosition(P *TPoint, Relative bool) TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(P)), PascalBool(Relative))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) ColumnFromPosition1(PositionIndex TColumnPosition) TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(5, m.Instance(), uintptr(PositionIndex))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetFirstVisibleColumn(ConsiderAllowFocus bool) TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(10, m.Instance(), PascalBool(ConsiderAllowFocus))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetLastVisibleColumn(ConsiderAllowFocus bool) TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(11, m.Instance(), PascalBool(ConsiderAllowFocus))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetFirstColumn() TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(9, m.Instance())
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetNextColumn(Column TColumnIndex) TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(12, m.Instance(), uintptr(Column))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetNextVisibleColumn(Column TColumnIndex, ConsiderAllowFocus bool) TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(13, m.Instance(), uintptr(Column), PascalBool(ConsiderAllowFocus))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetPreviousColumn(Column TColumnIndex) TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(14, m.Instance(), uintptr(Column))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetPreviousVisibleColumn(Column TColumnIndex, ConsiderAllowFocus bool) TColumnIndex {
	r1 := virtualTreeColumnsImportAPI().SysCallN(15, m.Instance(), uintptr(Column), PascalBool(ConsiderAllowFocus))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetScrollWidth() int32 {
	r1 := virtualTreeColumnsImportAPI().SysCallN(16, m.Instance())
	return int32(r1)
}

func (m *TVirtualTreeColumns) GetVisibleFixedWidth() int32 {
	r1 := virtualTreeColumnsImportAPI().SysCallN(17, m.Instance())
	return int32(r1)
}

func (m *TVirtualTreeColumns) IsValidColumn(Column TColumnIndex) bool {
	r1 := virtualTreeColumnsImportAPI().SysCallN(19, m.Instance(), uintptr(Column))
	return GoBool(r1)
}

func (m *TVirtualTreeColumns) TotalWidth() int32 {
	r1 := virtualTreeColumnsImportAPI().SysCallN(25, m.Instance())
	return int32(r1)
}

func VirtualTreeColumnsClass() TClass {
	ret := virtualTreeColumnsImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TVirtualTreeColumns) AnimatedResize(Column TColumnIndex, NewWidth int32) {
	virtualTreeColumnsImportAPI().SysCallN(1, m.Instance(), uintptr(Column), uintptr(NewWidth))
}

func (m *TVirtualTreeColumns) GetColumnBounds(Column TColumnIndex, OutLeft, OutRight *int32) {
	var result1 uintptr
	var result2 uintptr
	virtualTreeColumnsImportAPI().SysCallN(8, m.Instance(), uintptr(Column), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutLeft = int32(result1)
	*OutRight = int32(result2)
}

func (m *TVirtualTreeColumns) LoadFromStream(Stream IStream, Version int32) {
	virtualTreeColumnsImportAPI().SysCallN(21, m.Instance(), GetObjectUintptr(Stream), uintptr(Version))
}

func (m *TVirtualTreeColumns) PaintHeader(DC HDC, R *TRect, HOffset int32) {
	virtualTreeColumnsImportAPI().SysCallN(22, m.Instance(), uintptr(DC), uintptr(unsafePointer(R)), uintptr(HOffset))
}

func (m *TVirtualTreeColumns) PaintHeader1(TargetCanvas ICanvas, R *TRect, Target *TPoint, RTLOffset int32) {
	virtualTreeColumnsImportAPI().SysCallN(23, m.Instance(), GetObjectUintptr(TargetCanvas), uintptr(unsafePointer(R)), uintptr(unsafePointer(Target)), uintptr(RTLOffset))
}

func (m *TVirtualTreeColumns) SaveToStream(Stream IStream) {
	virtualTreeColumnsImportAPI().SysCallN(24, m.Instance(), GetObjectUintptr(Stream))
}

var (
	virtualTreeColumnsImport       *imports.Imports = nil
	virtualTreeColumnsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("VirtualTreeColumns_AddForVirtualTreeColumn", 0),
		/*1*/ imports.NewTable("VirtualTreeColumns_AnimatedResize", 0),
		/*2*/ imports.NewTable("VirtualTreeColumns_Class", 0),
		/*3*/ imports.NewTable("VirtualTreeColumns_ClickIndex", 0),
		/*4*/ imports.NewTable("VirtualTreeColumns_ColumnFromPosition", 0),
		/*5*/ imports.NewTable("VirtualTreeColumns_ColumnFromPosition1", 0),
		/*6*/ imports.NewTable("VirtualTreeColumns_Create", 0),
		/*7*/ imports.NewTable("VirtualTreeColumns_DefaultWidth", 0),
		/*8*/ imports.NewTable("VirtualTreeColumns_GetColumnBounds", 0),
		/*9*/ imports.NewTable("VirtualTreeColumns_GetFirstColumn", 0),
		/*10*/ imports.NewTable("VirtualTreeColumns_GetFirstVisibleColumn", 0),
		/*11*/ imports.NewTable("VirtualTreeColumns_GetLastVisibleColumn", 0),
		/*12*/ imports.NewTable("VirtualTreeColumns_GetNextColumn", 0),
		/*13*/ imports.NewTable("VirtualTreeColumns_GetNextVisibleColumn", 0),
		/*14*/ imports.NewTable("VirtualTreeColumns_GetPreviousColumn", 0),
		/*15*/ imports.NewTable("VirtualTreeColumns_GetPreviousVisibleColumn", 0),
		/*16*/ imports.NewTable("VirtualTreeColumns_GetScrollWidth", 0),
		/*17*/ imports.NewTable("VirtualTreeColumns_GetVisibleFixedWidth", 0),
		/*18*/ imports.NewTable("VirtualTreeColumns_Header", 0),
		/*19*/ imports.NewTable("VirtualTreeColumns_IsValidColumn", 0),
		/*20*/ imports.NewTable("VirtualTreeColumns_ItemsForVirtualTreeColumn", 0),
		/*21*/ imports.NewTable("VirtualTreeColumns_LoadFromStream", 0),
		/*22*/ imports.NewTable("VirtualTreeColumns_PaintHeader", 0),
		/*23*/ imports.NewTable("VirtualTreeColumns_PaintHeader1", 0),
		/*24*/ imports.NewTable("VirtualTreeColumns_SaveToStream", 0),
		/*25*/ imports.NewTable("VirtualTreeColumns_TotalWidth", 0),
		/*26*/ imports.NewTable("VirtualTreeColumns_TrackIndex", 0),
	}
)

func virtualTreeColumnsImportAPI() *imports.Imports {
	if virtualTreeColumnsImport == nil {
		virtualTreeColumnsImport = NewDefaultImports()
		virtualTreeColumnsImport.SetImportTable(virtualTreeColumnsImportTables)
		virtualTreeColumnsImportTables = nil
	}
	return virtualTreeColumnsImport
}
