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
	r1 := LCL().SysCallN(6019, GetObjectUintptr(AOwner))
	return AsVirtualTreeColumns(r1)
}

func (m *TVirtualTreeColumns) ClickIndex() TColumnIndex {
	r1 := LCL().SysCallN(6016, m.Instance())
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) DefaultWidth() int32 {
	r1 := LCL().SysCallN(6020, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumns) SetDefaultWidth(AValue int32) {
	LCL().SysCallN(6020, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumns) ItemsForVirtualTreeColumn(Index TColumnIndex) IVirtualTreeColumn {
	r1 := LCL().SysCallN(6033, 0, m.Instance(), uintptr(Index))
	return AsVirtualTreeColumn(r1)
}

func (m *TVirtualTreeColumns) SetItemsForVirtualTreeColumn(Index TColumnIndex, AValue IVirtualTreeColumn) {
	LCL().SysCallN(6033, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TVirtualTreeColumns) Header() IVTHeader {
	r1 := LCL().SysCallN(6031, m.Instance())
	return AsVTHeader(r1)
}

func (m *TVirtualTreeColumns) TrackIndex() TColumnIndex {
	r1 := LCL().SysCallN(6039, m.Instance())
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) AddForVirtualTreeColumn() IVirtualTreeColumn {
	r1 := LCL().SysCallN(6013, m.Instance())
	return AsVirtualTreeColumn(r1)
}

func (m *TVirtualTreeColumns) ColumnFromPosition(P *TPoint, Relative bool) TColumnIndex {
	r1 := LCL().SysCallN(6017, m.Instance(), uintptr(unsafePointer(P)), PascalBool(Relative))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) ColumnFromPosition1(PositionIndex TColumnPosition) TColumnIndex {
	r1 := LCL().SysCallN(6018, m.Instance(), uintptr(PositionIndex))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetFirstVisibleColumn(ConsiderAllowFocus bool) TColumnIndex {
	r1 := LCL().SysCallN(6023, m.Instance(), PascalBool(ConsiderAllowFocus))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetLastVisibleColumn(ConsiderAllowFocus bool) TColumnIndex {
	r1 := LCL().SysCallN(6024, m.Instance(), PascalBool(ConsiderAllowFocus))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetFirstColumn() TColumnIndex {
	r1 := LCL().SysCallN(6022, m.Instance())
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetNextColumn(Column TColumnIndex) TColumnIndex {
	r1 := LCL().SysCallN(6025, m.Instance(), uintptr(Column))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetNextVisibleColumn(Column TColumnIndex, ConsiderAllowFocus bool) TColumnIndex {
	r1 := LCL().SysCallN(6026, m.Instance(), uintptr(Column), PascalBool(ConsiderAllowFocus))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetPreviousColumn(Column TColumnIndex) TColumnIndex {
	r1 := LCL().SysCallN(6027, m.Instance(), uintptr(Column))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetPreviousVisibleColumn(Column TColumnIndex, ConsiderAllowFocus bool) TColumnIndex {
	r1 := LCL().SysCallN(6028, m.Instance(), uintptr(Column), PascalBool(ConsiderAllowFocus))
	return TColumnIndex(r1)
}

func (m *TVirtualTreeColumns) GetScrollWidth() int32 {
	r1 := LCL().SysCallN(6029, m.Instance())
	return int32(r1)
}

func (m *TVirtualTreeColumns) GetVisibleFixedWidth() int32 {
	r1 := LCL().SysCallN(6030, m.Instance())
	return int32(r1)
}

func (m *TVirtualTreeColumns) IsValidColumn(Column TColumnIndex) bool {
	r1 := LCL().SysCallN(6032, m.Instance(), uintptr(Column))
	return GoBool(r1)
}

func (m *TVirtualTreeColumns) TotalWidth() int32 {
	r1 := LCL().SysCallN(6038, m.Instance())
	return int32(r1)
}

func VirtualTreeColumnsClass() TClass {
	ret := LCL().SysCallN(6015)
	return TClass(ret)
}

func (m *TVirtualTreeColumns) AnimatedResize(Column TColumnIndex, NewWidth int32) {
	LCL().SysCallN(6014, m.Instance(), uintptr(Column), uintptr(NewWidth))
}

func (m *TVirtualTreeColumns) GetColumnBounds(Column TColumnIndex, OutLeft, OutRight *int32) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(6021, m.Instance(), uintptr(Column), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutLeft = int32(result1)
	*OutRight = int32(result2)
}

func (m *TVirtualTreeColumns) LoadFromStream(Stream IStream, Version int32) {
	LCL().SysCallN(6034, m.Instance(), GetObjectUintptr(Stream), uintptr(Version))
}

func (m *TVirtualTreeColumns) PaintHeader(DC HDC, R *TRect, HOffset int32) {
	LCL().SysCallN(6035, m.Instance(), uintptr(DC), uintptr(unsafePointer(R)), uintptr(HOffset))
}

func (m *TVirtualTreeColumns) PaintHeader1(TargetCanvas ICanvas, R *TRect, Target *TPoint, RTLOffset int32) {
	LCL().SysCallN(6036, m.Instance(), GetObjectUintptr(TargetCanvas), uintptr(unsafePointer(R)), uintptr(unsafePointer(Target)), uintptr(RTLOffset))
}

func (m *TVirtualTreeColumns) SaveToStream(Stream IStream) {
	LCL().SysCallN(6037, m.Instance(), GetObjectUintptr(Stream))
}
