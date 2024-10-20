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

// IPen Parent: IFPCustomPen
type IPen interface {
	IFPCustomPen
	Color() TColor                                    // property
	SetColor(AValue TColor)                           // property
	Cosmetic() bool                                   // property
	SetCosmetic(AValue bool)                          // property
	EndCapForPenEndCap() TPenEndCap                   // property
	SetEndCapForPenEndCap(AValue TPenEndCap)          // property
	JoinStyleForPenJoinStyle() TPenJoinStyle          // property
	SetJoinStyleForPenJoinStyle(AValue TPenJoinStyle) // property
	GetPatternForUintptr() uintptr                    // function
	SetPatternForPointer(APattern uintptr)            // procedure
}

// TPen Parent: TFPCustomPen
type TPen struct {
	TFPCustomPen
}

func NewPen() IPen {
	r1 := penImportAPI().SysCallN(3)
	return AsPen(r1)
}

func (m *TPen) Color() TColor {
	r1 := penImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TPen) SetColor(AValue TColor) {
	penImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TPen) Cosmetic() bool {
	r1 := penImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPen) SetCosmetic(AValue bool) {
	penImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPen) EndCapForPenEndCap() TPenEndCap {
	r1 := penImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TPenEndCap(r1)
}

func (m *TPen) SetEndCapForPenEndCap(AValue TPenEndCap) {
	penImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TPen) JoinStyleForPenJoinStyle() TPenJoinStyle {
	r1 := penImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TPenJoinStyle(r1)
}

func (m *TPen) SetJoinStyleForPenJoinStyle(AValue TPenJoinStyle) {
	penImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TPen) GetPatternForUintptr() uintptr {
	r1 := penImportAPI().SysCallN(5, m.Instance())
	return uintptr(r1)
}

func PenClass() TClass {
	ret := penImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TPen) SetPatternForPointer(APattern uintptr) {
	penImportAPI().SysCallN(7, m.Instance(), uintptr(APattern))
}

var (
	penImport       *imports.Imports = nil
	penImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Pen_Class", 0),
		/*1*/ imports.NewTable("Pen_Color", 0),
		/*2*/ imports.NewTable("Pen_Cosmetic", 0),
		/*3*/ imports.NewTable("Pen_Create", 0),
		/*4*/ imports.NewTable("Pen_EndCapForPenEndCap", 0),
		/*5*/ imports.NewTable("Pen_GetPatternForUintptr", 0),
		/*6*/ imports.NewTable("Pen_JoinStyleForPenJoinStyle", 0),
		/*7*/ imports.NewTable("Pen_SetPatternForPointer", 0),
	}
)

func penImportAPI() *imports.Imports {
	if penImport == nil {
		penImport = NewDefaultImports()
		penImport.SetImportTable(penImportTables)
		penImportTables = nil
	}
	return penImport
}
