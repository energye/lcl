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

// ICustomTrackBar Parent: IWinControl
type ICustomTrackBar interface {
	IWinControl
	SetTick(value int32)                             // procedure
	SetParams(position int32, min int32, max int32)  // procedure
	Frequency() int32                                // property Frequency Getter
	SetFrequency(value int32)                        // property Frequency Setter
	LineSize() int32                                 // property LineSize Getter
	SetLineSize(value int32)                         // property LineSize Setter
	Max() int32                                      // property Max Getter
	SetMax(value int32)                              // property Max Setter
	Min() int32                                      // property Min Getter
	SetMin(value int32)                              // property Min Setter
	Orientation() types.TTrackBarOrientation         // property Orientation Getter
	SetOrientation(value types.TTrackBarOrientation) // property Orientation Setter
	PageSize() int32                                 // property PageSize Getter
	SetPageSize(value int32)                         // property PageSize Setter
	Position() int32                                 // property Position Getter
	SetPosition(value int32)                         // property Position Setter
	Reversed() bool                                  // property Reversed Getter
	SetReversed(value bool)                          // property Reversed Setter
	ScalePos() types.TTrackBarScalePos               // property ScalePos Getter
	SetScalePos(value types.TTrackBarScalePos)       // property ScalePos Setter
	SelEnd() int32                                   // property SelEnd Getter
	SetSelEnd(value int32)                           // property SelEnd Setter
	SelStart() int32                                 // property SelStart Getter
	SetSelStart(value int32)                         // property SelStart Setter
	ShowSelRange() bool                              // property ShowSelRange Getter
	SetShowSelRange(value bool)                      // property ShowSelRange Setter
	TickMarks() types.TTickMark                      // property TickMarks Getter
	SetTickMarks(value types.TTickMark)              // property TickMarks Setter
	TickStyle() types.TTickStyle                     // property TickStyle Getter
	SetTickStyle(value types.TTickStyle)             // property TickStyle Setter
	SetOnChange(fn TNotifyEvent)                     // property event
}

type TCustomTrackBar struct {
	TWinControl
}

func (m *TCustomTrackBar) SetTick(value int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) SetParams(position int32, min int32, max int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(2, m.Instance(), uintptr(position), uintptr(min), uintptr(max))
}

func (m *TCustomTrackBar) Frequency() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTrackBar) SetFrequency(value int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) LineSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTrackBar) SetLineSize(value int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) Max() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTrackBar) SetMax(value int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) Min() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTrackBar) SetMin(value int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) Orientation() types.TTrackBarOrientation {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(7, 0, m.Instance())
	return types.TTrackBarOrientation(r)
}

func (m *TCustomTrackBar) SetOrientation(value types.TTrackBarOrientation) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) PageSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTrackBar) SetPageSize(value int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) Position() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTrackBar) SetPosition(value int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) Reversed() bool {
	if !m.IsValid() {
		return false
	}
	r := customTrackBarAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTrackBar) SetReversed(value bool) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTrackBar) ScalePos() types.TTrackBarScalePos {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(11, 0, m.Instance())
	return types.TTrackBarScalePos(r)
}

func (m *TCustomTrackBar) SetScalePos(value types.TTrackBarScalePos) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) SelEnd() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTrackBar) SetSelEnd(value int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) SelStart() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTrackBar) SetSelStart(value int32) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) ShowSelRange() bool {
	if !m.IsValid() {
		return false
	}
	r := customTrackBarAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTrackBar) SetShowSelRange(value bool) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTrackBar) TickMarks() types.TTickMark {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(15, 0, m.Instance())
	return types.TTickMark(r)
}

func (m *TCustomTrackBar) SetTickMarks(value types.TTickMark) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) TickStyle() types.TTickStyle {
	if !m.IsValid() {
		return 0
	}
	r := customTrackBarAPI().SysCallN(16, 0, m.Instance())
	return types.TTickStyle(r)
}

func (m *TCustomTrackBar) SetTickStyle(value types.TTickStyle) {
	if !m.IsValid() {
		return
	}
	customTrackBarAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrackBar) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, customTrackBarAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomTrackBar class constructor
func NewCustomTrackBar(owner IComponent) ICustomTrackBar {
	r := customTrackBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomTrackBar(r)
}

func TCustomTrackBarClass() types.TClass {
	r := customTrackBarAPI().SysCallN(18)
	return types.TClass(r)
}

var (
	customTrackBarOnce   base.Once
	customTrackBarImport *imports.Imports = nil
)

func customTrackBarAPI() *imports.Imports {
	customTrackBarOnce.Do(func() {
		customTrackBarImport = api.NewDefaultImports()
		customTrackBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomTrackBar_Create", 0), // constructor NewCustomTrackBar
			/* 1 */ imports.NewTable("TCustomTrackBar_SetTick", 0), // procedure SetTick
			/* 2 */ imports.NewTable("TCustomTrackBar_SetParams", 0), // procedure SetParams
			/* 3 */ imports.NewTable("TCustomTrackBar_Frequency", 0), // property Frequency
			/* 4 */ imports.NewTable("TCustomTrackBar_LineSize", 0), // property LineSize
			/* 5 */ imports.NewTable("TCustomTrackBar_Max", 0), // property Max
			/* 6 */ imports.NewTable("TCustomTrackBar_Min", 0), // property Min
			/* 7 */ imports.NewTable("TCustomTrackBar_Orientation", 0), // property Orientation
			/* 8 */ imports.NewTable("TCustomTrackBar_PageSize", 0), // property PageSize
			/* 9 */ imports.NewTable("TCustomTrackBar_Position", 0), // property Position
			/* 10 */ imports.NewTable("TCustomTrackBar_Reversed", 0), // property Reversed
			/* 11 */ imports.NewTable("TCustomTrackBar_ScalePos", 0), // property ScalePos
			/* 12 */ imports.NewTable("TCustomTrackBar_SelEnd", 0), // property SelEnd
			/* 13 */ imports.NewTable("TCustomTrackBar_SelStart", 0), // property SelStart
			/* 14 */ imports.NewTable("TCustomTrackBar_ShowSelRange", 0), // property ShowSelRange
			/* 15 */ imports.NewTable("TCustomTrackBar_TickMarks", 0), // property TickMarks
			/* 16 */ imports.NewTable("TCustomTrackBar_TickStyle", 0), // property TickStyle
			/* 17 */ imports.NewTable("TCustomTrackBar_OnChange", 0), // event OnChange
			/* 18 */ imports.NewTable("TCustomTrackBar_TClass", 0), // function TCustomTrackBarClass
		}
	})
	return customTrackBarImport
}
