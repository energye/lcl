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

// ICustomTrackBar Parent: IWinControl
type ICustomTrackBar interface {
	IWinControl
	Frequency() int32                           // property
	SetFrequency(AValue int32)                  // property
	LineSize() int32                            // property
	SetLineSize(AValue int32)                   // property
	Max() int32                                 // property
	SetMax(AValue int32)                        // property
	Min() int32                                 // property
	SetMin(AValue int32)                        // property
	Orientation() TTrackBarOrientation          // property
	SetOrientation(AValue TTrackBarOrientation) // property
	PageSize() int32                            // property
	SetPageSize(AValue int32)                   // property
	Position() int32                            // property
	SetPosition(AValue int32)                   // property
	Reversed() bool                             // property
	SetReversed(AValue bool)                    // property
	ScalePos() TTrackBarScalePos                // property
	SetScalePos(AValue TTrackBarScalePos)       // property
	SelEnd() int32                              // property
	SetSelEnd(AValue int32)                     // property
	SelStart() int32                            // property
	SetSelStart(AValue int32)                   // property
	ShowSelRange() bool                         // property
	SetShowSelRange(AValue bool)                // property
	TickMarks() TTickMark                       // property
	SetTickMarks(AValue TTickMark)              // property
	TickStyle() TTickStyle                      // property
	SetTickStyle(AValue TTickStyle)             // property
	SetTick(Value int32)                        // procedure
	SetParams(APosition, AMin, AMax int32)      // procedure
	SetOnChange(fn TNotifyEvent)                // property event
}

// TCustomTrackBar Parent: TWinControl
type TCustomTrackBar struct {
	TWinControl
	changePtr uintptr
}

func NewCustomTrackBar(AOwner IComponent) ICustomTrackBar {
	r1 := customTrackBarImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomTrackBar(r1)
}

func (m *TCustomTrackBar) Frequency() int32 {
	r1 := customTrackBarImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetFrequency(AValue int32) {
	customTrackBarImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) LineSize() int32 {
	r1 := customTrackBarImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetLineSize(AValue int32) {
	customTrackBarImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Max() int32 {
	r1 := customTrackBarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetMax(AValue int32) {
	customTrackBarImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Min() int32 {
	r1 := customTrackBarImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetMin(AValue int32) {
	customTrackBarImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Orientation() TTrackBarOrientation {
	r1 := customTrackBarImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TTrackBarOrientation(r1)
}

func (m *TCustomTrackBar) SetOrientation(AValue TTrackBarOrientation) {
	customTrackBarImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) PageSize() int32 {
	r1 := customTrackBarImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetPageSize(AValue int32) {
	customTrackBarImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Position() int32 {
	r1 := customTrackBarImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetPosition(AValue int32) {
	customTrackBarImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Reversed() bool {
	r1 := customTrackBarImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrackBar) SetReversed(AValue bool) {
	customTrackBarImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrackBar) ScalePos() TTrackBarScalePos {
	r1 := customTrackBarImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TTrackBarScalePos(r1)
}

func (m *TCustomTrackBar) SetScalePos(AValue TTrackBarScalePos) {
	customTrackBarImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) SelEnd() int32 {
	r1 := customTrackBarImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetSelEnd(AValue int32) {
	customTrackBarImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) SelStart() int32 {
	r1 := customTrackBarImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetSelStart(AValue int32) {
	customTrackBarImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) ShowSelRange() bool {
	r1 := customTrackBarImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrackBar) SetShowSelRange(AValue bool) {
	customTrackBarImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrackBar) TickMarks() TTickMark {
	r1 := customTrackBarImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return TTickMark(r1)
}

func (m *TCustomTrackBar) SetTickMarks(AValue TTickMark) {
	customTrackBarImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) TickStyle() TTickStyle {
	r1 := customTrackBarImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TTickStyle(r1)
}

func (m *TCustomTrackBar) SetTickStyle(AValue TTickStyle) {
	customTrackBarImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func CustomTrackBarClass() TClass {
	ret := customTrackBarImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomTrackBar) SetTick(Value int32) {
	customTrackBarImportAPI().SysCallN(15, m.Instance(), uintptr(Value))
}

func (m *TCustomTrackBar) SetParams(APosition, AMin, AMax int32) {
	customTrackBarImportAPI().SysCallN(14, m.Instance(), uintptr(APosition), uintptr(AMin), uintptr(AMax))
}

func (m *TCustomTrackBar) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	customTrackBarImportAPI().SysCallN(13, m.Instance(), m.changePtr)
}

var (
	customTrackBarImport       *imports.Imports = nil
	customTrackBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomTrackBar_Class", 0),
		/*1*/ imports.NewTable("CustomTrackBar_Create", 0),
		/*2*/ imports.NewTable("CustomTrackBar_Frequency", 0),
		/*3*/ imports.NewTable("CustomTrackBar_LineSize", 0),
		/*4*/ imports.NewTable("CustomTrackBar_Max", 0),
		/*5*/ imports.NewTable("CustomTrackBar_Min", 0),
		/*6*/ imports.NewTable("CustomTrackBar_Orientation", 0),
		/*7*/ imports.NewTable("CustomTrackBar_PageSize", 0),
		/*8*/ imports.NewTable("CustomTrackBar_Position", 0),
		/*9*/ imports.NewTable("CustomTrackBar_Reversed", 0),
		/*10*/ imports.NewTable("CustomTrackBar_ScalePos", 0),
		/*11*/ imports.NewTable("CustomTrackBar_SelEnd", 0),
		/*12*/ imports.NewTable("CustomTrackBar_SelStart", 0),
		/*13*/ imports.NewTable("CustomTrackBar_SetOnChange", 0),
		/*14*/ imports.NewTable("CustomTrackBar_SetParams", 0),
		/*15*/ imports.NewTable("CustomTrackBar_SetTick", 0),
		/*16*/ imports.NewTable("CustomTrackBar_ShowSelRange", 0),
		/*17*/ imports.NewTable("CustomTrackBar_TickMarks", 0),
		/*18*/ imports.NewTable("CustomTrackBar_TickStyle", 0),
	}
)

func customTrackBarImportAPI() *imports.Imports {
	if customTrackBarImport == nil {
		customTrackBarImport = NewDefaultImports()
		customTrackBarImport.SetImportTable(customTrackBarImportTables)
		customTrackBarImportTables = nil
	}
	return customTrackBarImport
}
