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

// IWindowMagnetOptions Parent: IPersistent
type IWindowMagnetOptions interface {
	IPersistent
	SnapToMonitor() bool           // property
	SetSnapToMonitor(AValue bool)  // property
	SnapToForms() bool             // property
	SetSnapToForms(AValue bool)    // property
	SnapFormTarget() bool          // property
	SetSnapFormTarget(AValue bool) // property
	Distance() int32               // property
	SetDistance(AValue int32)      // property
	AssignTo(Dest IPersistent)     // procedure
}

// TWindowMagnetOptions Parent: TPersistent
type TWindowMagnetOptions struct {
	TPersistent
}

func NewWindowMagnetOptions() IWindowMagnetOptions {
	r1 := windowMagnetOptionsImportAPI().SysCallN(2)
	return AsWindowMagnetOptions(r1)
}

func (m *TWindowMagnetOptions) SnapToMonitor() bool {
	r1 := windowMagnetOptionsImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWindowMagnetOptions) SetSnapToMonitor(AValue bool) {
	windowMagnetOptionsImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWindowMagnetOptions) SnapToForms() bool {
	r1 := windowMagnetOptionsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWindowMagnetOptions) SetSnapToForms(AValue bool) {
	windowMagnetOptionsImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWindowMagnetOptions) SnapFormTarget() bool {
	r1 := windowMagnetOptionsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWindowMagnetOptions) SetSnapFormTarget(AValue bool) {
	windowMagnetOptionsImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWindowMagnetOptions) Distance() int32 {
	r1 := windowMagnetOptionsImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TWindowMagnetOptions) SetDistance(AValue int32) {
	windowMagnetOptionsImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func WindowMagnetOptionsClass() TClass {
	ret := windowMagnetOptionsImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TWindowMagnetOptions) AssignTo(Dest IPersistent) {
	windowMagnetOptionsImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(Dest))
}

var (
	windowMagnetOptionsImport       *imports.Imports = nil
	windowMagnetOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WindowMagnetOptions_AssignTo", 0),
		/*1*/ imports.NewTable("WindowMagnetOptions_Class", 0),
		/*2*/ imports.NewTable("WindowMagnetOptions_Create", 0),
		/*3*/ imports.NewTable("WindowMagnetOptions_Distance", 0),
		/*4*/ imports.NewTable("WindowMagnetOptions_SnapFormTarget", 0),
		/*5*/ imports.NewTable("WindowMagnetOptions_SnapToForms", 0),
		/*6*/ imports.NewTable("WindowMagnetOptions_SnapToMonitor", 0),
	}
)

func windowMagnetOptionsImportAPI() *imports.Imports {
	if windowMagnetOptionsImport == nil {
		windowMagnetOptionsImport = NewDefaultImports()
		windowMagnetOptionsImport.SetImportTable(windowMagnetOptionsImportTables)
		windowMagnetOptionsImportTables = nil
	}
	return windowMagnetOptionsImport
}
