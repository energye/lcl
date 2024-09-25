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

// IFPCustomPen Parent: IFPCanvasHelper
type IFPCustomPen interface {
	IFPCanvasHelper
	Style() TFPPenStyle                  // property
	SetStyle(AValue TFPPenStyle)         // property
	Width() int32                        // property
	SetWidth(AValue int32)               // property
	Mode() TFPPenMode                    // property
	SetMode(AValue TFPPenMode)           // property
	Pattern() uint32                     // property
	SetPattern(AValue uint32)            // property
	EndCap() TFPPenEndCap                // property
	SetEndCap(AValue TFPPenEndCap)       // property
	JoinStyle() TFPPenJoinStyle          // property
	SetJoinStyle(AValue TFPPenJoinStyle) // property
	CopyPen() IFPCustomPen               // function
}

// TFPCustomPen Parent: TFPCanvasHelper
type TFPCustomPen struct {
	TFPCanvasHelper
}

func NewFPCustomPen() IFPCustomPen {
	r1 := LCL().SysCallN(2970)
	return AsFPCustomPen(r1)
}

func (m *TFPCustomPen) Style() TFPPenStyle {
	r1 := LCL().SysCallN(2975, 0, m.Instance(), 0)
	return TFPPenStyle(r1)
}

func (m *TFPCustomPen) SetStyle(AValue TFPPenStyle) {
	LCL().SysCallN(2975, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) Width() int32 {
	r1 := LCL().SysCallN(2976, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomPen) SetWidth(AValue int32) {
	LCL().SysCallN(2976, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) Mode() TFPPenMode {
	r1 := LCL().SysCallN(2973, 0, m.Instance(), 0)
	return TFPPenMode(r1)
}

func (m *TFPCustomPen) SetMode(AValue TFPPenMode) {
	LCL().SysCallN(2973, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) Pattern() uint32 {
	r1 := LCL().SysCallN(2974, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TFPCustomPen) SetPattern(AValue uint32) {
	LCL().SysCallN(2974, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) EndCap() TFPPenEndCap {
	r1 := LCL().SysCallN(2971, 0, m.Instance(), 0)
	return TFPPenEndCap(r1)
}

func (m *TFPCustomPen) SetEndCap(AValue TFPPenEndCap) {
	LCL().SysCallN(2971, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) JoinStyle() TFPPenJoinStyle {
	r1 := LCL().SysCallN(2972, 0, m.Instance(), 0)
	return TFPPenJoinStyle(r1)
}

func (m *TFPCustomPen) SetJoinStyle(AValue TFPPenJoinStyle) {
	LCL().SysCallN(2972, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) CopyPen() IFPCustomPen {
	r1 := LCL().SysCallN(2969, m.Instance())
	return AsFPCustomPen(r1)
}

func FPCustomPenClass() TClass {
	ret := LCL().SysCallN(2968)
	return TClass(ret)
}
