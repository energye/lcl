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

// IVTFixedAreaConstraints Parent: IPersistent
type IVTFixedAreaConstraints interface {
	IPersistent
	MaxHeightPercent() TVTConstraintPercent          // property
	SetMaxHeightPercent(AValue TVTConstraintPercent) // property
	MaxWidthPercent() TVTConstraintPercent           // property
	SetMaxWidthPercent(AValue TVTConstraintPercent)  // property
	MinHeightPercent() TVTConstraintPercent          // property
	SetMinHeightPercent(AValue TVTConstraintPercent) // property
	MinWidthPercent() TVTConstraintPercent           // property
	SetMinWidthPercent(AValue TVTConstraintPercent)  // property
	SetOnChange(fn TNotifyEvent)                     // property event
}

// TVTFixedAreaConstraints Parent: TPersistent
type TVTFixedAreaConstraints struct {
	TPersistent
	changePtr uintptr
}

func NewVTFixedAreaConstraints(AOwner IVTHeader) IVTFixedAreaConstraints {
	r1 := vTFixedAreaConstraintsImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsVTFixedAreaConstraints(r1)
}

func (m *TVTFixedAreaConstraints) MaxHeightPercent() TVTConstraintPercent {
	r1 := vTFixedAreaConstraintsImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMaxHeightPercent(AValue TVTConstraintPercent) {
	vTFixedAreaConstraintsImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTFixedAreaConstraints) MaxWidthPercent() TVTConstraintPercent {
	r1 := vTFixedAreaConstraintsImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMaxWidthPercent(AValue TVTConstraintPercent) {
	vTFixedAreaConstraintsImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTFixedAreaConstraints) MinHeightPercent() TVTConstraintPercent {
	r1 := vTFixedAreaConstraintsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMinHeightPercent(AValue TVTConstraintPercent) {
	vTFixedAreaConstraintsImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TVTFixedAreaConstraints) MinWidthPercent() TVTConstraintPercent {
	r1 := vTFixedAreaConstraintsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TVTConstraintPercent(r1)
}

func (m *TVTFixedAreaConstraints) SetMinWidthPercent(AValue TVTConstraintPercent) {
	vTFixedAreaConstraintsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func VTFixedAreaConstraintsClass() TClass {
	ret := vTFixedAreaConstraintsImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TVTFixedAreaConstraints) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	vTFixedAreaConstraintsImportAPI().SysCallN(6, m.Instance(), m.changePtr)
}

var (
	vTFixedAreaConstraintsImport       *imports.Imports = nil
	vTFixedAreaConstraintsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("VTFixedAreaConstraints_Class", 0),
		/*1*/ imports.NewTable("VTFixedAreaConstraints_Create", 0),
		/*2*/ imports.NewTable("VTFixedAreaConstraints_MaxHeightPercent", 0),
		/*3*/ imports.NewTable("VTFixedAreaConstraints_MaxWidthPercent", 0),
		/*4*/ imports.NewTable("VTFixedAreaConstraints_MinHeightPercent", 0),
		/*5*/ imports.NewTable("VTFixedAreaConstraints_MinWidthPercent", 0),
		/*6*/ imports.NewTable("VTFixedAreaConstraints_SetOnChange", 0),
	}
)

func vTFixedAreaConstraintsImportAPI() *imports.Imports {
	if vTFixedAreaConstraintsImport == nil {
		vTFixedAreaConstraintsImport = NewDefaultImports()
		vTFixedAreaConstraintsImport.SetImportTable(vTFixedAreaConstraintsImportTables)
		vTFixedAreaConstraintsImportTables = nil
	}
	return vTFixedAreaConstraintsImport
}
