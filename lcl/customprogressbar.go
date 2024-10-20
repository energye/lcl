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

// ICustomProgressBar Parent: IWinControl
type ICustomProgressBar interface {
	IWinControl
	Max() int32                                    // property
	SetMax(AValue int32)                           // property
	Min() int32                                    // property
	SetMin(AValue int32)                           // property
	Orientation() TProgressBarOrientation          // property
	SetOrientation(AValue TProgressBarOrientation) // property
	Position() int32                               // property
	SetPosition(AValue int32)                      // property
	Smooth() bool                                  // property
	SetSmooth(AValue bool)                         // property
	Step() int32                                   // property
	SetStep(AValue int32)                          // property
	Style() TProgressBarStyle                      // property
	SetStyle(AValue TProgressBarStyle)             // property
	BarShowText() bool                             // property
	SetBarShowText(AValue bool)                    // property
	StepIt()                                       // procedure
	StepBy(Delta int32)                            // procedure
}

// TCustomProgressBar Parent: TWinControl
type TCustomProgressBar struct {
	TWinControl
}

func NewCustomProgressBar(AOwner IComponent) ICustomProgressBar {
	r1 := customProgressBarImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsCustomProgressBar(r1)
}

func (m *TCustomProgressBar) Max() int32 {
	r1 := customProgressBarImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomProgressBar) SetMax(AValue int32) {
	customProgressBarImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Min() int32 {
	r1 := customProgressBarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomProgressBar) SetMin(AValue int32) {
	customProgressBarImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Orientation() TProgressBarOrientation {
	r1 := customProgressBarImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TProgressBarOrientation(r1)
}

func (m *TCustomProgressBar) SetOrientation(AValue TProgressBarOrientation) {
	customProgressBarImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Position() int32 {
	r1 := customProgressBarImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomProgressBar) SetPosition(AValue int32) {
	customProgressBarImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Smooth() bool {
	r1 := customProgressBarImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomProgressBar) SetSmooth(AValue bool) {
	customProgressBarImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomProgressBar) Step() int32 {
	r1 := customProgressBarImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomProgressBar) SetStep(AValue int32) {
	customProgressBarImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) Style() TProgressBarStyle {
	r1 := customProgressBarImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TProgressBarStyle(r1)
}

func (m *TCustomProgressBar) SetStyle(AValue TProgressBarStyle) {
	customProgressBarImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomProgressBar) BarShowText() bool {
	r1 := customProgressBarImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomProgressBar) SetBarShowText(AValue bool) {
	customProgressBarImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func CustomProgressBarClass() TClass {
	ret := customProgressBarImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCustomProgressBar) StepIt() {
	customProgressBarImportAPI().SysCallN(10, m.Instance())
}

func (m *TCustomProgressBar) StepBy(Delta int32) {
	customProgressBarImportAPI().SysCallN(9, m.Instance(), uintptr(Delta))
}

var (
	customProgressBarImport       *imports.Imports = nil
	customProgressBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomProgressBar_BarShowText", 0),
		/*1*/ imports.NewTable("CustomProgressBar_Class", 0),
		/*2*/ imports.NewTable("CustomProgressBar_Create", 0),
		/*3*/ imports.NewTable("CustomProgressBar_Max", 0),
		/*4*/ imports.NewTable("CustomProgressBar_Min", 0),
		/*5*/ imports.NewTable("CustomProgressBar_Orientation", 0),
		/*6*/ imports.NewTable("CustomProgressBar_Position", 0),
		/*7*/ imports.NewTable("CustomProgressBar_Smooth", 0),
		/*8*/ imports.NewTable("CustomProgressBar_Step", 0),
		/*9*/ imports.NewTable("CustomProgressBar_StepBy", 0),
		/*10*/ imports.NewTable("CustomProgressBar_StepIt", 0),
		/*11*/ imports.NewTable("CustomProgressBar_Style", 0),
	}
)

func customProgressBarImportAPI() *imports.Imports {
	if customProgressBarImport == nil {
		customProgressBarImport = NewDefaultImports()
		customProgressBarImport.SetImportTable(customProgressBarImportTables)
		customProgressBarImportTables = nil
	}
	return customProgressBarImport
}
