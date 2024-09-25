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

// IFPCustomBrush Parent: IFPCanvasHelper
type IFPCustomBrush interface {
	IFPCanvasHelper
	Style() TFPBrushStyle           // property
	SetStyle(AValue TFPBrushStyle)  // property
	Image() IFPCustomImage          // property
	SetImage(AValue IFPCustomImage) // property
	CopyBrush() IFPCustomBrush      // function
}

// TFPCustomBrush Parent: TFPCanvasHelper
type TFPCustomBrush struct {
	TFPCanvasHelper
}

func NewFPCustomBrush() IFPCustomBrush {
	r1 := LCL().SysCallN(2862)
	return AsFPCustomBrush(r1)
}

func (m *TFPCustomBrush) Style() TFPBrushStyle {
	r1 := LCL().SysCallN(2864, 0, m.Instance(), 0)
	return TFPBrushStyle(r1)
}

func (m *TFPCustomBrush) SetStyle(AValue TFPBrushStyle) {
	LCL().SysCallN(2864, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomBrush) Image() IFPCustomImage {
	r1 := LCL().SysCallN(2863, 0, m.Instance(), 0)
	return AsFPCustomImage(r1)
}

func (m *TFPCustomBrush) SetImage(AValue IFPCustomImage) {
	LCL().SysCallN(2863, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomBrush) CopyBrush() IFPCustomBrush {
	r1 := LCL().SysCallN(2861, m.Instance())
	return AsFPCustomBrush(r1)
}

func FPCustomBrushClass() TClass {
	ret := LCL().SysCallN(2860)
	return TClass(ret)
}
