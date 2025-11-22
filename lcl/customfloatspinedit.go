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

// ICustomFloatSpinEdit Parent: ICustomEdit
type ICustomFloatSpinEdit interface {
	ICustomEdit
	GetLimitedValue(value float64) float64 // function
	ValueToStr(value float64) string       // function
	StrToValue(S string) float64           // function
	DecimalPlaces() int32                  // property DecimalPlaces Getter
	SetDecimalPlaces(value int32)          // property DecimalPlaces Setter
	EditorEnabled() bool                   // property EditorEnabled Getter
	SetEditorEnabled(value bool)           // property EditorEnabled Setter
	Increment() float64                    // property Increment Getter
	SetIncrement(value float64)            // property Increment Setter
	MinValue() float64                     // property MinValue Getter
	SetMinValue(value float64)             // property MinValue Setter
	MaxValue() float64                     // property MaxValue Getter
	SetMaxValue(value float64)             // property MaxValue Setter
	Value() float64                        // property Value Getter
	SetValue(value float64)                // property Value Setter
	ValueEmpty() bool                      // property ValueEmpty Getter
	SetValueEmpty(value bool)              // property ValueEmpty Setter
}

type TCustomFloatSpinEdit struct {
	TCustomEdit
}

func (m *TCustomFloatSpinEdit) GetLimitedValue(value float64) (result float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&value)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomFloatSpinEdit) ValueToStr(value float64) string {
	if !m.IsValid() {
		return ""
	}
	r := customFloatSpinEditAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&value)))
	return api.GoStr(r)
}

func (m *TCustomFloatSpinEdit) StrToValue(S string) (result float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(3, m.Instance(), api.PasStr(S), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomFloatSpinEdit) DecimalPlaces() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customFloatSpinEditAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TCustomFloatSpinEdit) SetDecimalPlaces(value int32) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCustomFloatSpinEdit) EditorEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := customFloatSpinEditAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomFloatSpinEdit) SetEditorEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomFloatSpinEdit) Increment() (result float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(6, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomFloatSpinEdit) SetIncrement(value float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(6, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomFloatSpinEdit) MinValue() (result float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomFloatSpinEdit) SetMinValue(value float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(7, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomFloatSpinEdit) MaxValue() (result float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomFloatSpinEdit) SetMaxValue(value float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomFloatSpinEdit) Value() (result float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(9, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomFloatSpinEdit) SetValue(value float64) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(9, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomFloatSpinEdit) ValueEmpty() bool {
	if !m.IsValid() {
		return false
	}
	r := customFloatSpinEditAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomFloatSpinEdit) SetValueEmpty(value bool) {
	if !m.IsValid() {
		return
	}
	customFloatSpinEditAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

// NewCustomFloatSpinEdit class constructor
func NewCustomFloatSpinEdit(theOwner IComponent) ICustomFloatSpinEdit {
	r := customFloatSpinEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomFloatSpinEdit(r)
}

func TCustomFloatSpinEditClass() types.TClass {
	r := customFloatSpinEditAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	customFloatSpinEditOnce   base.Once
	customFloatSpinEditImport *imports.Imports = nil
)

func customFloatSpinEditAPI() *imports.Imports {
	customFloatSpinEditOnce.Do(func() {
		customFloatSpinEditImport = api.NewDefaultImports()
		customFloatSpinEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomFloatSpinEdit_Create", 0), // constructor NewCustomFloatSpinEdit
			/* 1 */ imports.NewTable("TCustomFloatSpinEdit_GetLimitedValue", 0), // function GetLimitedValue
			/* 2 */ imports.NewTable("TCustomFloatSpinEdit_ValueToStr", 0), // function ValueToStr
			/* 3 */ imports.NewTable("TCustomFloatSpinEdit_StrToValue", 0), // function StrToValue
			/* 4 */ imports.NewTable("TCustomFloatSpinEdit_DecimalPlaces", 0), // property DecimalPlaces
			/* 5 */ imports.NewTable("TCustomFloatSpinEdit_EditorEnabled", 0), // property EditorEnabled
			/* 6 */ imports.NewTable("TCustomFloatSpinEdit_Increment", 0), // property Increment
			/* 7 */ imports.NewTable("TCustomFloatSpinEdit_MinValue", 0), // property MinValue
			/* 8 */ imports.NewTable("TCustomFloatSpinEdit_MaxValue", 0), // property MaxValue
			/* 9 */ imports.NewTable("TCustomFloatSpinEdit_Value", 0), // property Value
			/* 10 */ imports.NewTable("TCustomFloatSpinEdit_ValueEmpty", 0), // property ValueEmpty
			/* 11 */ imports.NewTable("TCustomFloatSpinEdit_TClass", 0), // function TCustomFloatSpinEditClass
		}
	})
	return customFloatSpinEditImport
}
