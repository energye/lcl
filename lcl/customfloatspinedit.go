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

// ICustomFloatSpinEdit Parent: ICustomEdit
type ICustomFloatSpinEdit interface {
	ICustomEdit
	DecimalPlaces() int32                                  // property
	SetDecimalPlaces(AValue int32)                         // property
	EditorEnabled() bool                                   // property
	SetEditorEnabled(AValue bool)                          // property
	Increment() (resultDouble float64)                     // property
	SetIncrement(AValue float64)                           // property
	MinValue() (resultDouble float64)                      // property
	SetMinValue(AValue float64)                            // property
	MaxValue() (resultDouble float64)                      // property
	SetMaxValue(AValue float64)                            // property
	Value() (resultDouble float64)                         // property
	SetValue(AValue float64)                               // property
	ValueEmpty() bool                                      // property
	SetValueEmpty(AValue bool)                             // property
	GetLimitedValue(AValue float64) (resultDouble float64) // function
	ValueToStr(AValue float64) string                      // function
	StrToValue(S string) (resultDouble float64)            // function
}

// TCustomFloatSpinEdit Parent: TCustomEdit
type TCustomFloatSpinEdit struct {
	TCustomEdit
}

func NewCustomFloatSpinEdit(TheOwner IComponent) ICustomFloatSpinEdit {
	r1 := customFloatSpinEditImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsCustomFloatSpinEdit(r1)
}

func (m *TCustomFloatSpinEdit) DecimalPlaces() int32 {
	r1 := customFloatSpinEditImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomFloatSpinEdit) SetDecimalPlaces(AValue int32) {
	customFloatSpinEditImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomFloatSpinEdit) EditorEnabled() bool {
	r1 := customFloatSpinEditImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomFloatSpinEdit) SetEditorEnabled(AValue bool) {
	customFloatSpinEditImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomFloatSpinEdit) Increment() (resultDouble float64) {
	customFloatSpinEditImportAPI().SysCallN(5, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) SetIncrement(AValue float64) {
	customFloatSpinEditImportAPI().SysCallN(5, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCustomFloatSpinEdit) MinValue() (resultDouble float64) {
	customFloatSpinEditImportAPI().SysCallN(7, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) SetMinValue(AValue float64) {
	customFloatSpinEditImportAPI().SysCallN(7, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCustomFloatSpinEdit) MaxValue() (resultDouble float64) {
	customFloatSpinEditImportAPI().SysCallN(6, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) SetMaxValue(AValue float64) {
	customFloatSpinEditImportAPI().SysCallN(6, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCustomFloatSpinEdit) Value() (resultDouble float64) {
	customFloatSpinEditImportAPI().SysCallN(9, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) SetValue(AValue float64) {
	customFloatSpinEditImportAPI().SysCallN(9, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TCustomFloatSpinEdit) ValueEmpty() bool {
	r1 := customFloatSpinEditImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomFloatSpinEdit) SetValueEmpty(AValue bool) {
	customFloatSpinEditImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomFloatSpinEdit) GetLimitedValue(AValue float64) (resultDouble float64) {
	customFloatSpinEditImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TCustomFloatSpinEdit) ValueToStr(AValue float64) string {
	r1 := customFloatSpinEditImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&AValue)))
	return GoStr(r1)
}

func (m *TCustomFloatSpinEdit) StrToValue(S string) (resultDouble float64) {
	customFloatSpinEditImportAPI().SysCallN(8, m.Instance(), PascalStr(S), uintptr(unsafePointer(&resultDouble)))
	return
}

func CustomFloatSpinEditClass() TClass {
	ret := customFloatSpinEditImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customFloatSpinEditImport       *imports.Imports = nil
	customFloatSpinEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomFloatSpinEdit_Class", 0),
		/*1*/ imports.NewTable("CustomFloatSpinEdit_Create", 0),
		/*2*/ imports.NewTable("CustomFloatSpinEdit_DecimalPlaces", 0),
		/*3*/ imports.NewTable("CustomFloatSpinEdit_EditorEnabled", 0),
		/*4*/ imports.NewTable("CustomFloatSpinEdit_GetLimitedValue", 0),
		/*5*/ imports.NewTable("CustomFloatSpinEdit_Increment", 0),
		/*6*/ imports.NewTable("CustomFloatSpinEdit_MaxValue", 0),
		/*7*/ imports.NewTable("CustomFloatSpinEdit_MinValue", 0),
		/*8*/ imports.NewTable("CustomFloatSpinEdit_StrToValue", 0),
		/*9*/ imports.NewTable("CustomFloatSpinEdit_Value", 0),
		/*10*/ imports.NewTable("CustomFloatSpinEdit_ValueEmpty", 0),
		/*11*/ imports.NewTable("CustomFloatSpinEdit_ValueToStr", 0),
	}
)

func customFloatSpinEditImportAPI() *imports.Imports {
	if customFloatSpinEditImport == nil {
		customFloatSpinEditImport = NewDefaultImports()
		customFloatSpinEditImport.SetImportTable(customFloatSpinEditImportTables)
		customFloatSpinEditImportTables = nil
	}
	return customFloatSpinEditImport
}
