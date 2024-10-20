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

// ISizeConstraints Parent: IPersistent
type ISizeConstraints interface {
	IPersistent
	MaxInterfaceHeight() int32                            // property
	MaxInterfaceWidth() int32                             // property
	MinInterfaceHeight() int32                            // property
	MinInterfaceWidth() int32                             // property
	Control() IControl                                    // property
	Options() TSizeConstraintsOptions                     // property
	SetOptions(AValue TSizeConstraintsOptions)            // property
	MaxHeight() TConstraintSize                           // property
	SetMaxHeight(AValue TConstraintSize)                  // property
	MaxWidth() TConstraintSize                            // property
	SetMaxWidth(AValue TConstraintSize)                   // property
	MinHeight() TConstraintSize                           // property
	SetMinHeight(AValue TConstraintSize)                  // property
	MinWidth() TConstraintSize                            // property
	SetMinWidth(AValue TConstraintSize)                   // property
	EffectiveMinWidth() int32                             // function
	EffectiveMinHeight() int32                            // function
	EffectiveMaxWidth() int32                             // function
	EffectiveMaxHeight() int32                            // function
	MinMaxWidth(Width int32) int32                        // function
	MinMaxHeight(Height int32) int32                      // function
	UpdateInterfaceConstraints()                          // procedure
	SetInterfaceConstraints(MinW, MinH, MaxW, MaxH int32) // procedure
	AutoAdjustLayout(AXProportion, AYProportion float64)  // procedure
	SetOnChange(fn TNotifyEvent)                          // property event
}

// TSizeConstraints Parent: TPersistent
type TSizeConstraints struct {
	TPersistent
	changePtr uintptr
}

func NewSizeConstraints(AControl IControl) ISizeConstraints {
	r1 := sizeConstraintsImportAPI().SysCallN(3, GetObjectUintptr(AControl))
	return AsSizeConstraints(r1)
}

func (m *TSizeConstraints) MaxInterfaceHeight() int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MaxInterfaceWidth() int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MinInterfaceHeight() int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(13, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MinInterfaceWidth() int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(14, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) Control() IControl {
	r1 := sizeConstraintsImportAPI().SysCallN(2, m.Instance())
	return AsControl(r1)
}

func (m *TSizeConstraints) Options() TSizeConstraintsOptions {
	r1 := sizeConstraintsImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TSizeConstraintsOptions(r1)
}

func (m *TSizeConstraints) SetOptions(AValue TSizeConstraintsOptions) {
	sizeConstraintsImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MaxHeight() TConstraintSize {
	r1 := sizeConstraintsImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMaxHeight(AValue TConstraintSize) {
	sizeConstraintsImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MaxWidth() TConstraintSize {
	r1 := sizeConstraintsImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMaxWidth(AValue TConstraintSize) {
	sizeConstraintsImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MinHeight() TConstraintSize {
	r1 := sizeConstraintsImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMinHeight(AValue TConstraintSize) {
	sizeConstraintsImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) MinWidth() TConstraintSize {
	r1 := sizeConstraintsImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return TConstraintSize(r1)
}

func (m *TSizeConstraints) SetMinWidth(AValue TConstraintSize) {
	sizeConstraintsImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TSizeConstraints) EffectiveMinWidth() int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) EffectiveMinHeight() int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(6, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) EffectiveMaxWidth() int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) EffectiveMaxHeight() int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TSizeConstraints) MinMaxWidth(Width int32) int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(16, m.Instance(), uintptr(Width))
	return int32(r1)
}

func (m *TSizeConstraints) MinMaxHeight(Height int32) int32 {
	r1 := sizeConstraintsImportAPI().SysCallN(15, m.Instance(), uintptr(Height))
	return int32(r1)
}

func SizeConstraintsClass() TClass {
	ret := sizeConstraintsImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TSizeConstraints) UpdateInterfaceConstraints() {
	sizeConstraintsImportAPI().SysCallN(21, m.Instance())
}

func (m *TSizeConstraints) SetInterfaceConstraints(MinW, MinH, MaxW, MaxH int32) {
	sizeConstraintsImportAPI().SysCallN(19, m.Instance(), uintptr(MinW), uintptr(MinH), uintptr(MaxW), uintptr(MaxH))
}

func (m *TSizeConstraints) AutoAdjustLayout(AXProportion, AYProportion float64) {
	sizeConstraintsImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(&AXProportion)), uintptr(unsafePointer(&AYProportion)))
}

func (m *TSizeConstraints) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	sizeConstraintsImportAPI().SysCallN(20, m.Instance(), m.changePtr)
}

var (
	sizeConstraintsImport       *imports.Imports = nil
	sizeConstraintsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("SizeConstraints_AutoAdjustLayout", 0),
		/*1*/ imports.NewTable("SizeConstraints_Class", 0),
		/*2*/ imports.NewTable("SizeConstraints_Control", 0),
		/*3*/ imports.NewTable("SizeConstraints_Create", 0),
		/*4*/ imports.NewTable("SizeConstraints_EffectiveMaxHeight", 0),
		/*5*/ imports.NewTable("SizeConstraints_EffectiveMaxWidth", 0),
		/*6*/ imports.NewTable("SizeConstraints_EffectiveMinHeight", 0),
		/*7*/ imports.NewTable("SizeConstraints_EffectiveMinWidth", 0),
		/*8*/ imports.NewTable("SizeConstraints_MaxHeight", 0),
		/*9*/ imports.NewTable("SizeConstraints_MaxInterfaceHeight", 0),
		/*10*/ imports.NewTable("SizeConstraints_MaxInterfaceWidth", 0),
		/*11*/ imports.NewTable("SizeConstraints_MaxWidth", 0),
		/*12*/ imports.NewTable("SizeConstraints_MinHeight", 0),
		/*13*/ imports.NewTable("SizeConstraints_MinInterfaceHeight", 0),
		/*14*/ imports.NewTable("SizeConstraints_MinInterfaceWidth", 0),
		/*15*/ imports.NewTable("SizeConstraints_MinMaxHeight", 0),
		/*16*/ imports.NewTable("SizeConstraints_MinMaxWidth", 0),
		/*17*/ imports.NewTable("SizeConstraints_MinWidth", 0),
		/*18*/ imports.NewTable("SizeConstraints_Options", 0),
		/*19*/ imports.NewTable("SizeConstraints_SetInterfaceConstraints", 0),
		/*20*/ imports.NewTable("SizeConstraints_SetOnChange", 0),
		/*21*/ imports.NewTable("SizeConstraints_UpdateInterfaceConstraints", 0),
	}
)

func sizeConstraintsImportAPI() *imports.Imports {
	if sizeConstraintsImport == nil {
		sizeConstraintsImport = NewDefaultImports()
		sizeConstraintsImport.SetImportTable(sizeConstraintsImportTables)
		sizeConstraintsImportTables = nil
	}
	return sizeConstraintsImport
}
