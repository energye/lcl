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

// ICustomCheckCombo Parent: ICustomComboBox
type ICustomCheckCombo interface {
	ICustomComboBox
	AllowGrayed() bool                                                      // property
	SetAllowGrayed(AValue bool)                                             // property
	Count() int32                                                           // property
	Checked(AIndex int32) bool                                              // property
	SetChecked(AIndex int32, AValue bool)                                   // property
	ItemEnabled(AIndex int32) bool                                          // property
	SetItemEnabled(AIndex int32, AValue bool)                               // property
	Objects(AIndex int32) IObject                                           // property
	SetObjects(AIndex int32, AValue IObject)                                // property
	State(AIndex int32) TCheckBoxState                                      // property
	SetState(AIndex int32, AValue TCheckBoxState)                           // property
	AddItemForPChar(AItem string, AState TCheckBoxState, AEnabled bool)     // procedure
	AssignItems(AItems IStrings)                                            // procedure
	DeleteItem(AIndex int32)                                                // procedure
	CheckAll(AState TCheckBoxState, AAllowGrayed bool, AAllowDisabled bool) // procedure
	Toggle(AIndex int32)                                                    // procedure
	SetOnItemChange(fn TCheckItemChange)                                    // property event
}

// TCustomCheckCombo Parent: TCustomComboBox
type TCustomCheckCombo struct {
	TCustomComboBox
	itemChangePtr uintptr
}

func NewCustomCheckCombo(AOwner IComponent) ICustomCheckCombo {
	r1 := LCL().SysCallN(1371, GetObjectUintptr(AOwner))
	return AsCustomCheckCombo(r1)
}

func (m *TCustomCheckCombo) AllowGrayed() bool {
	r1 := LCL().SysCallN(1365, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCheckCombo) SetAllowGrayed(AValue bool) {
	LCL().SysCallN(1365, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCheckCombo) Count() int32 {
	r1 := LCL().SysCallN(1370, m.Instance())
	return int32(r1)
}

func (m *TCustomCheckCombo) Checked(AIndex int32) bool {
	r1 := LCL().SysCallN(1368, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckCombo) SetChecked(AIndex int32, AValue bool) {
	LCL().SysCallN(1368, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckCombo) ItemEnabled(AIndex int32) bool {
	r1 := LCL().SysCallN(1373, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckCombo) SetItemEnabled(AIndex int32, AValue bool) {
	LCL().SysCallN(1373, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckCombo) Objects(AIndex int32) IObject {
	r1 := LCL().SysCallN(1374, 0, m.Instance(), uintptr(AIndex))
	return AsObject(r1)
}

func (m *TCustomCheckCombo) SetObjects(AIndex int32, AValue IObject) {
	LCL().SysCallN(1374, 1, m.Instance(), uintptr(AIndex), GetObjectUintptr(AValue))
}

func (m *TCustomCheckCombo) State(AIndex int32) TCheckBoxState {
	r1 := LCL().SysCallN(1376, 0, m.Instance(), uintptr(AIndex))
	return TCheckBoxState(r1)
}

func (m *TCustomCheckCombo) SetState(AIndex int32, AValue TCheckBoxState) {
	LCL().SysCallN(1376, 1, m.Instance(), uintptr(AIndex), uintptr(AValue))
}

func CustomCheckComboClass() TClass {
	ret := LCL().SysCallN(1369)
	return TClass(ret)
}

func (m *TCustomCheckCombo) AddItemForPChar(AItem string, AState TCheckBoxState, AEnabled bool) {
	LCL().SysCallN(1364, m.Instance(), PascalStr(AItem), uintptr(AState), PascalBool(AEnabled))
}

func (m *TCustomCheckCombo) AssignItems(AItems IStrings) {
	LCL().SysCallN(1366, m.Instance(), GetObjectUintptr(AItems))
}

func (m *TCustomCheckCombo) DeleteItem(AIndex int32) {
	LCL().SysCallN(1372, m.Instance(), uintptr(AIndex))
}

func (m *TCustomCheckCombo) CheckAll(AState TCheckBoxState, AAllowGrayed bool, AAllowDisabled bool) {
	LCL().SysCallN(1367, m.Instance(), uintptr(AState), PascalBool(AAllowGrayed), PascalBool(AAllowDisabled))
}

func (m *TCustomCheckCombo) Toggle(AIndex int32) {
	LCL().SysCallN(1377, m.Instance(), uintptr(AIndex))
}

func (m *TCustomCheckCombo) SetOnItemChange(fn TCheckItemChange) {
	if m.itemChangePtr != 0 {
		RemoveEventElement(m.itemChangePtr)
	}
	m.itemChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1375, m.Instance(), m.itemChangePtr)
}
