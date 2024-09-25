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
	r1 := LCL().SysCallN(2272, GetObjectUintptr(TheOwner))
	return AsCustomSpinEdit(r1)
}

func (m *TCustomSpinEdit) ValueForInteger() int32 {
	r1 := LCL().SysCallN(2276, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetValueForInteger(AValue int32) {
	LCL().SysCallN(2276, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpinEdit) MinValueForInteger() int32 {
	r1 := LCL().SysCallN(2275, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetMinValueForInteger(AValue int32) {
	LCL().SysCallN(2275, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpinEdit) MaxValueForInteger() int32 {
	r1 := LCL().SysCallN(2274, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetMaxValueForInteger(AValue int32) {
	LCL().SysCallN(2274, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpinEdit) IncrementForInteger() int32 {
	r1 := LCL().SysCallN(2273, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpinEdit) SetIncrementForInteger(AValue int32) {
	LCL().SysCallN(2273, 1, m.Instance(), uintptr(AValue))
}

func CustomSpinEditClass() TClass {
	ret := LCL().SysCallN(2271)
	return TClass(ret)
}
