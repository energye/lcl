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

// IGauge Parent: IATGauge
type IGauge interface {
	IATGauge
	ParentFont() bool                      // property
	SetParentFont(AValue bool)             // property
	ForeColor() TColor                     // property
	SetForeColor(AValue TColor)            // property
	BackColor() TColor                     // property
	SetBackColor(AValue TColor)            // property
	KindForGaugeKind() TGaugeKind          // property
	SetKindForGaugeKind(AValue TGaugeKind) // property
}

// TGauge Parent: TATGauge
type TGauge struct {
	TATGauge
}

func NewGauge(AOwner IComponent) IGauge {
	r1 := LCL().SysCallN(3197, GetObjectUintptr(AOwner))
	return AsGauge(r1)
}

func (m *TGauge) ParentFont() bool {
	r1 := LCL().SysCallN(3200, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGauge) SetParentFont(AValue bool) {
	LCL().SysCallN(3200, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGauge) ForeColor() TColor {
	r1 := LCL().SysCallN(3198, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGauge) SetForeColor(AValue TColor) {
	LCL().SysCallN(3198, 1, m.Instance(), uintptr(AValue))
}

func (m *TGauge) BackColor() TColor {
	r1 := LCL().SysCallN(3195, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGauge) SetBackColor(AValue TColor) {
	LCL().SysCallN(3195, 1, m.Instance(), uintptr(AValue))
}

func (m *TGauge) KindForGaugeKind() TGaugeKind {
	r1 := LCL().SysCallN(3199, 0, m.Instance(), 0)
	return TGaugeKind(r1)
}

func (m *TGauge) SetKindForGaugeKind(AValue TGaugeKind) {
	LCL().SysCallN(3199, 1, m.Instance(), uintptr(AValue))
}

func GaugeClass() TClass {
	ret := LCL().SysCallN(3196)
	return TClass(ret)
}
