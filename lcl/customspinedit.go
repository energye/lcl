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

// ICustomSpinEdit Parent: ICustomFloatSpinEdit
type ICustomSpinEdit interface {
	ICustomFloatSpinEdit
	ValueToInt() int32             // property Value Getter
	SetValueToInt(value int32)     // property Value Setter
	MinValueToInt() int32          // property MinValue Getter
	SetMinValueToInt(value int32)  // property MinValue Setter
	MaxValueToInt() int32          // property MaxValue Getter
	SetMaxValueToInt(value int32)  // property MaxValue Setter
	IncrementToInt() int32         // property Increment Getter
	SetIncrementToInt(value int32) // property Increment Setter
}

type TCustomSpinEdit struct {
	TCustomFloatSpinEdit
}

func (m *TCustomSpinEdit) ValueToInt() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpinEditAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpinEdit) SetValueToInt(value int32) {
	if !m.IsValid() {
		return
	}
	customSpinEditAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpinEdit) MinValueToInt() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpinEditAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpinEdit) SetMinValueToInt(value int32) {
	if !m.IsValid() {
		return
	}
	customSpinEditAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpinEdit) MaxValueToInt() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpinEditAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpinEdit) SetMaxValueToInt(value int32) {
	if !m.IsValid() {
		return
	}
	customSpinEditAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpinEdit) IncrementToInt() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpinEditAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpinEdit) SetIncrementToInt(value int32) {
	if !m.IsValid() {
		return
	}
	customSpinEditAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

// NewCustomSpinEdit class constructor
func NewCustomSpinEdit(theOwner IComponent) ICustomSpinEdit {
	r := customSpinEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomSpinEdit(r)
}

func TCustomSpinEditClass() types.TClass {
	r := customSpinEditAPI().SysCallN(5)
	return types.TClass(r)
}

var (
	customSpinEditOnce   base.Once
	customSpinEditImport *imports.Imports = nil
)

func customSpinEditAPI() *imports.Imports {
	customSpinEditOnce.Do(func() {
		customSpinEditImport = api.NewDefaultImports()
		customSpinEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomSpinEdit_Create", 0), // constructor NewCustomSpinEdit
			/* 1 */ imports.NewTable("TCustomSpinEdit_ValueToInt", 0), // property ValueToInt
			/* 2 */ imports.NewTable("TCustomSpinEdit_MinValueToInt", 0), // property MinValueToInt
			/* 3 */ imports.NewTable("TCustomSpinEdit_MaxValueToInt", 0), // property MaxValueToInt
			/* 4 */ imports.NewTable("TCustomSpinEdit_IncrementToInt", 0), // property IncrementToInt
			/* 5 */ imports.NewTable("TCustomSpinEdit_TClass", 0), // function TCustomSpinEditClass
		}
	})
	return customSpinEditImport
}
