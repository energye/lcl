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

// ICustomSpinEdit Parent: ICustomFloatSpinEdit
type ICustomSpinEdit interface {
	ICustomFloatSpinEdit
	ValueForInteger() int32              // property
	SetValueForInteger(AValue int32)     // property
	MinValueForInteger() int32           // property
	SetMinValueForInteger(AValue int32)  // property
	MaxValueForInteger() int32           // property
	SetMaxValueForInteger(AValue int32)  // property
	IncrementForInteger() int32          // property
	SetIncrementForInteger(AValue int32) // property
}

// TCustomSpinEdit Parent: TCustomFloatSpinEdit
type TCustomSpinEdit struct {
	TCustomFloatSpinEdit
}

func NewCustomSpinEdit(TheOwner IComponent) ICustomSpinEdit {
	r1 := customSpinEditImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsCustomSpinEdit(r1)
}

func (m *TCustomSpinEdit) ValueForInteger() int32 {
	r1 := customSpinEditImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetValueForInteger(AValue int32) {
	customSpinEditImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpinEdit) MinValueForInteger() int32 {
	r1 := customSpinEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetMinValueForInteger(AValue int32) {
	customSpinEditImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpinEdit) MaxValueForInteger() int32 {
	r1 := customSpinEditImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetMaxValueForInteger(AValue int32) {
	customSpinEditImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpinEdit) IncrementForInteger() int32 {
	r1 := customSpinEditImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetIncrementForInteger(AValue int32) {
	customSpinEditImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func CustomSpinEditClass() TClass {
	ret := customSpinEditImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customSpinEditImport       *imports.Imports = nil
	customSpinEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomSpinEdit_Class", 0),
		/*1*/ imports.NewTable("CustomSpinEdit_Create", 0),
		/*2*/ imports.NewTable("CustomSpinEdit_IncrementForInteger", 0),
		/*3*/ imports.NewTable("CustomSpinEdit_MaxValueForInteger", 0),
		/*4*/ imports.NewTable("CustomSpinEdit_MinValueForInteger", 0),
		/*5*/ imports.NewTable("CustomSpinEdit_ValueForInteger", 0),
	}
)

func customSpinEditImportAPI() *imports.Imports {
	if customSpinEditImport == nil {
		customSpinEditImport = NewDefaultImports()
		customSpinEditImport.SetImportTable(customSpinEditImportTables)
		customSpinEditImportTables = nil
	}
	return customSpinEditImport
}
