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
	r1 := LCL().SysCallN(2387, GetObjectUintptr(AOwner))
	return AsCustomTrackBar(r1)
}

func (m *TCustomTrackBar) Frequency() int32 {
	r1 := LCL().SysCallN(2388, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetFrequency(AValue int32) {
	LCL().SysCallN(2388, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) LineSize() int32 {
	r1 := LCL().SysCallN(2389, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetLineSize(AValue int32) {
	LCL().SysCallN(2389, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Max() int32 {
	r1 := LCL().SysCallN(2390, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetMax(AValue int32) {
	LCL().SysCallN(2390, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Min() int32 {
	r1 := LCL().SysCallN(2391, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetMin(AValue int32) {
	LCL().SysCallN(2391, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Orientation() TTrackBarOrientation {
	r1 := LCL().SysCallN(2392, 0, m.Instance(), 0)
	return TTrackBarOrientation(r1)
}

func (m *TCustomTrackBar) SetOrientation(AValue TTrackBarOrientation) {
	LCL().SysCallN(2392, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) PageSize() int32 {
	r1 := LCL().SysCallN(2393, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetPageSize(AValue int32) {
	LCL().SysCallN(2393, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Position() int32 {
	r1 := LCL().SysCallN(2394, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetPosition(AValue int32) {
	LCL().SysCallN(2394, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) Reversed() bool {
	r1 := LCL().SysCallN(2395, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrackBar) SetReversed(AValue bool) {
	LCL().SysCallN(2395, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrackBar) ScalePos() TTrackBarScalePos {
	r1 := LCL().SysCallN(2396, 0, m.Instance(), 0)
	return TTrackBarScalePos(r1)
}

func (m *TCustomTrackBar) SetScalePos(AValue TTrackBarScalePos) {
	LCL().SysCallN(2396, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) SelEnd() int32 {
	r1 := LCL().SysCallN(2397, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetSelEnd(AValue int32) {
	LCL().SysCallN(2397, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) SelStart() int32 {
	r1 := LCL().SysCallN(2398, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrackBar) SetSelStart(AValue int32) {
	LCL().SysCallN(2398, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) ShowSelRange() bool {
	r1 := LCL().SysCallN(2402, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrackBar) SetShowSelRange(AValue bool) {
	LCL().SysCallN(2402, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrackBar) TickMarks() TTickMark {
	r1 := LCL().SysCallN(2403, 0, m.Instance(), 0)
	return TTickMark(r1)
}

func (m *TCustomTrackBar) SetTickMarks(AValue TTickMark) {
	LCL().SysCallN(2403, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrackBar) TickStyle() TTickStyle {
	r1 := LCL().SysCallN(2404, 0, m.Instance(), 0)
	return TTickStyle(r1)
}

func (m *TCustomTrackBar) SetTickStyle(AValue TTickStyle) {
	LCL().SysCallN(2404, 1, m.Instance(), uintptr(AValue))
}

func CustomTrackBarClass() TClass {
	ret := LCL().SysCallN(2386)
	return TClass(ret)
}

func (m *TCustomTrackBar) SetTick(Value int32) {
	LCL().SysCallN(2401, m.Instance(), uintptr(Value))
}

func (m *TCustomTrackBar) SetParams(APosition, AMin, AMax int32) {
	LCL().SysCallN(2400, m.Instance(), uintptr(APosition), uintptr(AMin), uintptr(AMax))
}

func (m *TCustomTrackBar) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2399, m.Instance(), m.changePtr)
}
