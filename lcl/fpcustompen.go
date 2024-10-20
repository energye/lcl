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
	r1 := fPCustomPenImportAPI().SysCallN(2)
	return AsFPCustomPen(r1)
}

func (m *TFPCustomPen) Style() TFPPenStyle {
	r1 := fPCustomPenImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TFPPenStyle(r1)
}

func (m *TFPCustomPen) SetStyle(AValue TFPPenStyle) {
	fPCustomPenImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) Width() int32 {
	r1 := fPCustomPenImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomPen) SetWidth(AValue int32) {
	fPCustomPenImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) Mode() TFPPenMode {
	r1 := fPCustomPenImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TFPPenMode(r1)
}

func (m *TFPCustomPen) SetMode(AValue TFPPenMode) {
	fPCustomPenImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) Pattern() uint32 {
	r1 := fPCustomPenImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TFPCustomPen) SetPattern(AValue uint32) {
	fPCustomPenImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) EndCap() TFPPenEndCap {
	r1 := fPCustomPenImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TFPPenEndCap(r1)
}

func (m *TFPCustomPen) SetEndCap(AValue TFPPenEndCap) {
	fPCustomPenImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) JoinStyle() TFPPenJoinStyle {
	r1 := fPCustomPenImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TFPPenJoinStyle(r1)
}

func (m *TFPCustomPen) SetJoinStyle(AValue TFPPenJoinStyle) {
	fPCustomPenImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomPen) CopyPen() IFPCustomPen {
	r1 := fPCustomPenImportAPI().SysCallN(1, m.Instance())
	return AsFPCustomPen(r1)
}

func FPCustomPenClass() TClass {
	ret := fPCustomPenImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	fPCustomPenImport       *imports.Imports = nil
	fPCustomPenImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCustomPen_Class", 0),
		/*1*/ imports.NewTable("FPCustomPen_CopyPen", 0),
		/*2*/ imports.NewTable("FPCustomPen_Create", 0),
		/*3*/ imports.NewTable("FPCustomPen_EndCap", 0),
		/*4*/ imports.NewTable("FPCustomPen_JoinStyle", 0),
		/*5*/ imports.NewTable("FPCustomPen_Mode", 0),
		/*6*/ imports.NewTable("FPCustomPen_Pattern", 0),
		/*7*/ imports.NewTable("FPCustomPen_Style", 0),
		/*8*/ imports.NewTable("FPCustomPen_Width", 0),
	}
)

func fPCustomPenImportAPI() *imports.Imports {
	if fPCustomPenImport == nil {
		fPCustomPenImport = NewDefaultImports()
		fPCustomPenImport.SetImportTable(fPCustomPenImportTables)
		fPCustomPenImportTables = nil
	}
	return fPCustomPenImport
}
