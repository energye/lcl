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

// ICustomMaskEdit Parent: ICustomEdit
type ICustomMaskEdit interface {
	ICustomEdit
	EnableSets() bool                                           // property
	SetEnableSets(AValue bool)                                  // property
	ValidationErrorMode() TMaskEditValidationErrorMode          // property
	SetValidationErrorMode(AValue TMaskEditValidationErrorMode) // property
	ValidateEdit()                                              // procedure
	SetOnValidationError(fn TNotifyEvent)                       // property event
}

// TCustomMaskEdit Parent: TCustomEdit
type TCustomMaskEdit struct {
	TCustomEdit
	validationErrorPtr uintptr
}

func NewCustomMaskEdit(TheOwner IComponent) ICustomMaskEdit {
	r1 := customMaskEditImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsCustomMaskEdit(r1)
}

func (m *TCustomMaskEdit) EnableSets() bool {
	r1 := customMaskEditImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomMaskEdit) SetEnableSets(AValue bool) {
	customMaskEditImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomMaskEdit) ValidationErrorMode() TMaskEditValidationErrorMode {
	r1 := customMaskEditImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TMaskEditValidationErrorMode(r1)
}

func (m *TCustomMaskEdit) SetValidationErrorMode(AValue TMaskEditValidationErrorMode) {
	customMaskEditImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func CustomMaskEditClass() TClass {
	ret := customMaskEditImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomMaskEdit) ValidateEdit() {
	customMaskEditImportAPI().SysCallN(4, m.Instance())
}

func (m *TCustomMaskEdit) SetOnValidationError(fn TNotifyEvent) {
	if m.validationErrorPtr != 0 {
		RemoveEventElement(m.validationErrorPtr)
	}
	m.validationErrorPtr = MakeEventDataPtr(fn)
	customMaskEditImportAPI().SysCallN(3, m.Instance(), m.validationErrorPtr)
}

var (
	customMaskEditImport       *imports.Imports = nil
	customMaskEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomMaskEdit_Class", 0),
		/*1*/ imports.NewTable("CustomMaskEdit_Create", 0),
		/*2*/ imports.NewTable("CustomMaskEdit_EnableSets", 0),
		/*3*/ imports.NewTable("CustomMaskEdit_SetOnValidationError", 0),
		/*4*/ imports.NewTable("CustomMaskEdit_ValidateEdit", 0),
		/*5*/ imports.NewTable("CustomMaskEdit_ValidationErrorMode", 0),
	}
)

func customMaskEditImportAPI() *imports.Imports {
	if customMaskEditImport == nil {
		customMaskEditImport = NewDefaultImports()
		customMaskEditImport.SetImportTable(customMaskEditImportTables)
		customMaskEditImportTables = nil
	}
	return customMaskEditImport
}
