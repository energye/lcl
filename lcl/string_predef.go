//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy go string => pascal string
// String reference, used to read string values

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
)

// TString
// Pointer reference
type TString struct {
	instance unsafePointer
}

// NewTString
//
//	Create TString pointer reference
func NewTString() *TString {
	var result uintptr
	stringImportAPI().SysCallN(0, uintptr(unsafePointer(&result)))
	return &TString{instance: unsafePointer(result)}
}

// AsTString
//
//	Convert TString
//	instance must string(TString) pointer
func AsTString(instance uintptr) *TString {
	if instance == 0 {
		return nil
	}
	return &TString{instance: unsafePointer(instance)}
}

// IsValid
//
//	return true if created
func (m *TString) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

// Instance
//
//	return string value pointer
func (m *TString) Instance() uintptr {
	return uintptr(m.instance)
}

// Free
//
//	Destroy this reference
func (m *TString) Free() {
	if !m.IsValid() {
		return
	}
	var instance = m.Instance()
	stringImportAPI().SysCallN(1, uintptr(unsafePointer(&instance)))
	m.instance = nil
}

// SetValue set new value
func (m *TString) SetValue(newValue string) {
	if !m.IsValid() {
		return
	}
	stringImportAPI().SysCallN(2, m.Instance(), PascalStr(newValue))
}

// Value
//
//	get string pointer, string length bytes copy
func (m *TString) Value() string {
	if !m.IsValid() {
		return ""
	}
	var result uintptr
	s := stringImportAPI().SysCallN(3, m.Instance(), uintptr(unsafePointer(&result)))
	if result != 0 && s > 0 {
		str := make([]byte, s)
		for i := 0; i < int(s); i++ {
			str[i] = *(*byte)(unsafePointer(result + uintptr(i)))
		}
		return string(str)
	}
	return ""
}

var (
	stringImport       *imports.Imports = nil
	stringImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TString_Create", 0),
		/*1*/ imports.NewTable("TString_Free", 0),
		/*2*/ imports.NewTable("TString_SetValue", 0),
		/*3*/ imports.NewTable("TString_GetValue", 0),
	}
)

func stringImportAPI() *imports.Imports {
	if stringImport == nil {
		stringImport = NewDefaultImports()
		stringImport.SetImportTable(stringImportTables)
		stringImportTables = nil
	}
	return stringImport
}
