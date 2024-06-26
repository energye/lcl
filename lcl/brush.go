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

// IBrush Parent: IFPCustomBrush
type IBrush interface {
	IFPCustomBrush
	Bitmap() ICustomBitmap          // property
	SetBitmap(AValue ICustomBitmap) // property
	Color() TColor                  // property
	SetColor(AValue TColor)         // property
	EqualsBrush(ABrush IBrush) bool // function
}

// TBrush Parent: TFPCustomBrush
type TBrush struct {
	TFPCustomBrush
}

func NewBrush() IBrush {
	r1 := LCL().SysCallN(472)
	return AsBrush(r1)
}

func (m *TBrush) Bitmap() ICustomBitmap {
	r1 := LCL().SysCallN(469, 0, m.Instance(), 0)
	return AsCustomBitmap(r1)
}

func (m *TBrush) SetBitmap(AValue ICustomBitmap) {
	LCL().SysCallN(469, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBrush) Color() TColor {
	r1 := LCL().SysCallN(471, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TBrush) SetColor(AValue TColor) {
	LCL().SysCallN(471, 1, m.Instance(), uintptr(AValue))
}

func (m *TBrush) EqualsBrush(ABrush IBrush) bool {
	r1 := LCL().SysCallN(473, m.Instance(), GetObjectUintptr(ABrush))
	return GoBool(r1)
}

func BrushClass() TClass {
	ret := LCL().SysCallN(470)
	return TClass(ret)
}
