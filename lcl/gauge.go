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

// IGauge Parent: IATGauge
type IGauge interface {
	IATGauge
	ParentFont() bool                          // property ParentFont Getter
	SetParentFont(value bool)                  // property ParentFont Setter
	ForeColor() types.TColor                   // property ForeColor Getter
	SetForeColor(value types.TColor)           // property ForeColor Setter
	BackColor() types.TColor                   // property BackColor Getter
	SetBackColor(value types.TColor)           // property BackColor Setter
	KindToGaugeKind() types.TGaugeKind         // property Kind Getter
	SetKindToGaugeKind(value types.TGaugeKind) // property Kind Setter
}

type TGauge struct {
	TATGauge
}

func (m *TGauge) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := gaugeAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGauge) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	gaugeAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TGauge) ForeColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := gaugeAPI().SysCallN(2, 0, m.Instance())
	return types.TColor(r)
}

func (m *TGauge) SetForeColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	gaugeAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TGauge) BackColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := gaugeAPI().SysCallN(3, 0, m.Instance())
	return types.TColor(r)
}

func (m *TGauge) SetBackColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	gaugeAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TGauge) KindToGaugeKind() types.TGaugeKind {
	if !m.IsValid() {
		return 0
	}
	r := gaugeAPI().SysCallN(4, 0, m.Instance())
	return types.TGaugeKind(r)
}

func (m *TGauge) SetKindToGaugeKind(value types.TGaugeKind) {
	if !m.IsValid() {
		return
	}
	gaugeAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

// NewGauge class constructor
func NewGauge(owner IComponent) IGauge {
	r := gaugeAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsGauge(r)
}

func TGaugeClass() types.TClass {
	r := gaugeAPI().SysCallN(5)
	return types.TClass(r)
}

var (
	gaugeOnce   base.Once
	gaugeImport *imports.Imports = nil
)

func gaugeAPI() *imports.Imports {
	gaugeOnce.Do(func() {
		gaugeImport = api.NewDefaultImports()
		gaugeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGauge_Create", 0), // constructor NewGauge
			/* 1 */ imports.NewTable("TGauge_ParentFont", 0), // property ParentFont
			/* 2 */ imports.NewTable("TGauge_ForeColor", 0), // property ForeColor
			/* 3 */ imports.NewTable("TGauge_BackColor", 0), // property BackColor
			/* 4 */ imports.NewTable("TGauge_KindToGaugeKind", 0), // property KindToGaugeKind
			/* 5 */ imports.NewTable("TGauge_TClass", 0), // function TGaugeClass
		}
	})
	return gaugeImport
}
