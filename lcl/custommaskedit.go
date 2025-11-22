//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// ICustomMaskEdit Parent: ICustomEdit
type ICustomMaskEdit interface {
	ICustomEdit
	Clear()                                                          // procedure
	ValidateEdit()                                                   // procedure
	EnableSets() bool                                                // property EnableSets Getter
	SetEnableSets(value bool)                                        // property EnableSets Setter
	ValidationErrorMode() types.TMaskEditValidationErrorMode         // property ValidationErrorMode Getter
	SetValidationErrorMode(value types.TMaskEditValidationErrorMode) // property ValidationErrorMode Setter
	SetOnValidationError(fn TNotifyEvent)                            // property event
}

type TCustomMaskEdit struct {
	TCustomEdit
}

func (m *TCustomMaskEdit) Clear() {
	if !m.IsValid() {
		return
	}
	customMaskEditAPI().SysCallN(1, m.Instance())
}

func (m *TCustomMaskEdit) ValidateEdit() {
	if !m.IsValid() {
		return
	}
	customMaskEditAPI().SysCallN(2, m.Instance())
}

func (m *TCustomMaskEdit) EnableSets() bool {
	if !m.IsValid() {
		return false
	}
	r := customMaskEditAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomMaskEdit) SetEnableSets(value bool) {
	if !m.IsValid() {
		return
	}
	customMaskEditAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomMaskEdit) ValidationErrorMode() types.TMaskEditValidationErrorMode {
	if !m.IsValid() {
		return 0
	}
	r := customMaskEditAPI().SysCallN(4, 0, m.Instance())
	return types.TMaskEditValidationErrorMode(r)
}

func (m *TCustomMaskEdit) SetValidationErrorMode(value types.TMaskEditValidationErrorMode) {
	if !m.IsValid() {
		return
	}
	customMaskEditAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomMaskEdit) SetOnValidationError(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 5, customMaskEditAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomMaskEdit class constructor
func NewCustomMaskEdit(theOwner IComponent) ICustomMaskEdit {
	r := customMaskEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomMaskEdit(r)
}

func TCustomMaskEditClass() types.TClass {
	r := customMaskEditAPI().SysCallN(6)
	return types.TClass(r)
}

var (
	customMaskEditOnce   base.Once
	customMaskEditImport *imports.Imports = nil
)

func customMaskEditAPI() *imports.Imports {
	customMaskEditOnce.Do(func() {
		customMaskEditImport = api.NewDefaultImports()
		customMaskEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomMaskEdit_Create", 0), // constructor NewCustomMaskEdit
			/* 1 */ imports.NewTable("TCustomMaskEdit_Clear", 0), // procedure Clear
			/* 2 */ imports.NewTable("TCustomMaskEdit_ValidateEdit", 0), // procedure ValidateEdit
			/* 3 */ imports.NewTable("TCustomMaskEdit_EnableSets", 0), // property EnableSets
			/* 4 */ imports.NewTable("TCustomMaskEdit_ValidationErrorMode", 0), // property ValidationErrorMode
			/* 5 */ imports.NewTable("TCustomMaskEdit_OnValidationError", 0), // event OnValidationError
			/* 6 */ imports.NewTable("TCustomMaskEdit_TClass", 0), // function TCustomMaskEditClass
		}
	})
	return customMaskEditImport
}
