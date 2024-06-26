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
	r1 := LCL().SysCallN(2087, GetObjectUintptr(TheOwner))
	return AsCustomMaskEdit(r1)
}

func (m *TCustomMaskEdit) EnableSets() bool {
	r1 := LCL().SysCallN(2088, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomMaskEdit) SetEnableSets(AValue bool) {
	LCL().SysCallN(2088, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomMaskEdit) ValidationErrorMode() TMaskEditValidationErrorMode {
	r1 := LCL().SysCallN(2091, 0, m.Instance(), 0)
	return TMaskEditValidationErrorMode(r1)
}

func (m *TCustomMaskEdit) SetValidationErrorMode(AValue TMaskEditValidationErrorMode) {
	LCL().SysCallN(2091, 1, m.Instance(), uintptr(AValue))
}

func CustomMaskEditClass() TClass {
	ret := LCL().SysCallN(2086)
	return TClass(ret)
}

func (m *TCustomMaskEdit) ValidateEdit() {
	LCL().SysCallN(2090, m.Instance())
}

func (m *TCustomMaskEdit) SetOnValidationError(fn TNotifyEvent) {
	if m.validationErrorPtr != 0 {
		RemoveEventElement(m.validationErrorPtr)
	}
	m.validationErrorPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2089, m.Instance(), m.validationErrorPtr)
}
