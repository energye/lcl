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

// ICustomCheckCombo Parent: ICustomComboBox
type ICustomCheckCombo interface {
	ICustomComboBox
	AddItemWithStringCheckBoxStateBool(item string, state types.TCheckBoxState, enabled bool) // procedure
	AssignItems(items IStrings)                                                               // procedure
	DeleteItem(index int32)                                                                   // procedure
	CheckAll(state types.TCheckBoxState, allowGrayed bool, allowDisabled bool)                // procedure
	Toggle(index int32)                                                                       // procedure
	AllowGrayed() bool                                                                        // property AllowGrayed Getter
	SetAllowGrayed(value bool)                                                                // property AllowGrayed Setter
	Count() int32                                                                             // property Count Getter
	Checked(index int32) bool                                                                 // property Checked Getter
	SetChecked(index int32, value bool)                                                       // property Checked Setter
	ItemEnabled(index int32) bool                                                             // property ItemEnabled Getter
	SetItemEnabled(index int32, value bool)                                                   // property ItemEnabled Setter
	Objects(index int32) IObject                                                              // property Objects Getter
	SetObjects(index int32, value IObject)                                                    // property Objects Setter
	State(index int32) types.TCheckBoxState                                                   // property State Getter
	SetState(index int32, value types.TCheckBoxState)                                         // property State Setter
	SetOnItemChange(fn TCheckItemChange)                                                      // property event
}

type TCustomCheckCombo struct {
	TCustomComboBox
}

func (m *TCustomCheckCombo) AddItemWithStringCheckBoxStateBool(item string, state types.TCheckBoxState, enabled bool) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(1, m.Instance(), api.PasStr(item), uintptr(state), api.PasBool(enabled))
}

func (m *TCustomCheckCombo) AssignItems(items IStrings) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(items))
}

func (m *TCustomCheckCombo) DeleteItem(index int32) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(3, m.Instance(), uintptr(index))
}

func (m *TCustomCheckCombo) CheckAll(state types.TCheckBoxState, allowGrayed bool, allowDisabled bool) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(4, m.Instance(), uintptr(state), api.PasBool(allowGrayed), api.PasBool(allowDisabled))
}

func (m *TCustomCheckCombo) Toggle(index int32) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(5, m.Instance(), uintptr(index))
}

func (m *TCustomCheckCombo) AllowGrayed() bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckComboAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCheckCombo) SetAllowGrayed(value bool) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCheckCombo) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customCheckComboAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TCustomCheckCombo) Checked(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckComboAPI().SysCallN(8, 0, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomCheckCombo) SetChecked(index int32, value bool) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(8, 1, m.Instance(), uintptr(index), api.PasBool(value))
}

func (m *TCustomCheckCombo) ItemEnabled(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckComboAPI().SysCallN(9, 0, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomCheckCombo) SetItemEnabled(index int32, value bool) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(9, 1, m.Instance(), uintptr(index), api.PasBool(value))
}

func (m *TCustomCheckCombo) Objects(index int32) IObject {
	if !m.IsValid() {
		return nil
	}
	r := customCheckComboAPI().SysCallN(10, 0, m.Instance(), uintptr(index))
	return AsObject(r)
}

func (m *TCustomCheckCombo) SetObjects(index int32, value IObject) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(10, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TCustomCheckCombo) State(index int32) types.TCheckBoxState {
	if !m.IsValid() {
		return 0
	}
	r := customCheckComboAPI().SysCallN(11, 0, m.Instance(), uintptr(index))
	return types.TCheckBoxState(r)
}

func (m *TCustomCheckCombo) SetState(index int32, value types.TCheckBoxState) {
	if !m.IsValid() {
		return
	}
	customCheckComboAPI().SysCallN(11, 1, m.Instance(), uintptr(index), uintptr(value))
}

func (m *TCustomCheckCombo) SetOnItemChange(fn TCheckItemChange) {
	if !m.IsValid() {
		return
	}
	cb := makeTCheckItemChange(fn)
	base.SetEvent(m, 12, customCheckComboAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomCheckCombo class constructor
func NewCustomCheckCombo(owner IComponent) ICustomCheckCombo {
	r := customCheckComboAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomCheckCombo(r)
}

func TCustomCheckComboClass() types.TClass {
	r := customCheckComboAPI().SysCallN(13)
	return types.TClass(r)
}

var (
	customCheckComboOnce   base.Once
	customCheckComboImport *imports.Imports = nil
)

func customCheckComboAPI() *imports.Imports {
	customCheckComboOnce.Do(func() {
		customCheckComboImport = api.NewDefaultImports()
		customCheckComboImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCheckCombo_Create", 0), // constructor NewCustomCheckCombo
			/* 1 */ imports.NewTable("TCustomCheckCombo_AddItemWithStringCheckBoxStateBool", 0), // procedure AddItemWithStringCheckBoxStateBool
			/* 2 */ imports.NewTable("TCustomCheckCombo_AssignItems", 0), // procedure AssignItems
			/* 3 */ imports.NewTable("TCustomCheckCombo_DeleteItem", 0), // procedure DeleteItem
			/* 4 */ imports.NewTable("TCustomCheckCombo_CheckAll", 0), // procedure CheckAll
			/* 5 */ imports.NewTable("TCustomCheckCombo_Toggle", 0), // procedure Toggle
			/* 6 */ imports.NewTable("TCustomCheckCombo_AllowGrayed", 0), // property AllowGrayed
			/* 7 */ imports.NewTable("TCustomCheckCombo_Count", 0), // property Count
			/* 8 */ imports.NewTable("TCustomCheckCombo_Checked", 0), // property Checked
			/* 9 */ imports.NewTable("TCustomCheckCombo_ItemEnabled", 0), // property ItemEnabled
			/* 10 */ imports.NewTable("TCustomCheckCombo_Objects", 0), // property Objects
			/* 11 */ imports.NewTable("TCustomCheckCombo_State", 0), // property State
			/* 12 */ imports.NewTable("TCustomCheckCombo_OnItemChange", 0), // event OnItemChange
			/* 13 */ imports.NewTable("TCustomCheckCombo_TClass", 0), // function TCustomCheckComboClass
		}
	})
	return customCheckComboImport
}
