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
	r1 := gaugeImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsGauge(r1)
}

func (m *TGauge) ParentFont() bool {
	r1 := gaugeImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGauge) SetParentFont(AValue bool) {
	gaugeImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGauge) ForeColor() TColor {
	r1 := gaugeImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGauge) SetForeColor(AValue TColor) {
	gaugeImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TGauge) BackColor() TColor {
	r1 := gaugeImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TGauge) SetBackColor(AValue TColor) {
	gaugeImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TGauge) KindForGaugeKind() TGaugeKind {
	r1 := gaugeImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TGaugeKind(r1)
}

func (m *TGauge) SetKindForGaugeKind(AValue TGaugeKind) {
	gaugeImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func GaugeClass() TClass {
	ret := gaugeImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	gaugeImport       *imports.Imports = nil
	gaugeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Gauge_BackColor", 0),
		/*1*/ imports.NewTable("Gauge_Class", 0),
		/*2*/ imports.NewTable("Gauge_Create", 0),
		/*3*/ imports.NewTable("Gauge_ForeColor", 0),
		/*4*/ imports.NewTable("Gauge_KindForGaugeKind", 0),
		/*5*/ imports.NewTable("Gauge_ParentFont", 0),
	}
)

func gaugeImportAPI() *imports.Imports {
	if gaugeImport == nil {
		gaugeImport = NewDefaultImports()
		gaugeImport.SetImportTable(gaugeImportTables)
		gaugeImportTables = nil
	}
	return gaugeImport
}
