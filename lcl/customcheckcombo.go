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
	r1 := customCheckComboImportAPI().SysCallN(7, GetObjectUintptr(AOwner))
	return AsCustomCheckCombo(r1)
}

func (m *TCustomCheckCombo) AllowGrayed() bool {
	r1 := customCheckComboImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCheckCombo) SetAllowGrayed(AValue bool) {
	customCheckComboImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCheckCombo) Count() int32 {
	r1 := customCheckComboImportAPI().SysCallN(6, m.Instance())
	return int32(r1)
}

func (m *TCustomCheckCombo) Checked(AIndex int32) bool {
	r1 := customCheckComboImportAPI().SysCallN(4, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckCombo) SetChecked(AIndex int32, AValue bool) {
	customCheckComboImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckCombo) ItemEnabled(AIndex int32) bool {
	r1 := customCheckComboImportAPI().SysCallN(9, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckCombo) SetItemEnabled(AIndex int32, AValue bool) {
	customCheckComboImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckCombo) Objects(AIndex int32) IObject {
	r1 := customCheckComboImportAPI().SysCallN(10, 0, m.Instance(), uintptr(AIndex))
	return AsObject(r1)
}

func (m *TCustomCheckCombo) SetObjects(AIndex int32, AValue IObject) {
	customCheckComboImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AIndex), GetObjectUintptr(AValue))
}

func (m *TCustomCheckCombo) State(AIndex int32) TCheckBoxState {
	r1 := customCheckComboImportAPI().SysCallN(12, 0, m.Instance(), uintptr(AIndex))
	return TCheckBoxState(r1)
}

func (m *TCustomCheckCombo) SetState(AIndex int32, AValue TCheckBoxState) {
	customCheckComboImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AIndex), uintptr(AValue))
}

func CustomCheckComboClass() TClass {
	ret := customCheckComboImportAPI().SysCallN(5)
	return TClass(ret)
}

func (m *TCustomCheckCombo) AddItemForPChar(AItem string, AState TCheckBoxState, AEnabled bool) {
	customCheckComboImportAPI().SysCallN(0, m.Instance(), PascalStr(AItem), uintptr(AState), PascalBool(AEnabled))
}

func (m *TCustomCheckCombo) AssignItems(AItems IStrings) {
	customCheckComboImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(AItems))
}

func (m *TCustomCheckCombo) DeleteItem(AIndex int32) {
	customCheckComboImportAPI().SysCallN(8, m.Instance(), uintptr(AIndex))
}

func (m *TCustomCheckCombo) CheckAll(AState TCheckBoxState, AAllowGrayed bool, AAllowDisabled bool) {
	customCheckComboImportAPI().SysCallN(3, m.Instance(), uintptr(AState), PascalBool(AAllowGrayed), PascalBool(AAllowDisabled))
}

func (m *TCustomCheckCombo) Toggle(AIndex int32) {
	customCheckComboImportAPI().SysCallN(13, m.Instance(), uintptr(AIndex))
}

func (m *TCustomCheckCombo) SetOnItemChange(fn TCheckItemChange) {
	if m.itemChangePtr != 0 {
		RemoveEventElement(m.itemChangePtr)
	}
	m.itemChangePtr = MakeEventDataPtr(fn)
	customCheckComboImportAPI().SysCallN(11, m.Instance(), m.itemChangePtr)
}

var (
	customCheckComboImport       *imports.Imports = nil
	customCheckComboImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCheckCombo_AddItemForPChar", 0),
		/*1*/ imports.NewTable("CustomCheckCombo_AllowGrayed", 0),
		/*2*/ imports.NewTable("CustomCheckCombo_AssignItems", 0),
		/*3*/ imports.NewTable("CustomCheckCombo_CheckAll", 0),
		/*4*/ imports.NewTable("CustomCheckCombo_Checked", 0),
		/*5*/ imports.NewTable("CustomCheckCombo_Class", 0),
		/*6*/ imports.NewTable("CustomCheckCombo_Count", 0),
		/*7*/ imports.NewTable("CustomCheckCombo_Create", 0),
		/*8*/ imports.NewTable("CustomCheckCombo_DeleteItem", 0),
		/*9*/ imports.NewTable("CustomCheckCombo_ItemEnabled", 0),
		/*10*/ imports.NewTable("CustomCheckCombo_Objects", 0),
		/*11*/ imports.NewTable("CustomCheckCombo_SetOnItemChange", 0),
		/*12*/ imports.NewTable("CustomCheckCombo_State", 0),
		/*13*/ imports.NewTable("CustomCheckCombo_Toggle", 0),
	}
)

func customCheckComboImportAPI() *imports.Imports {
	if customCheckComboImport == nil {
		customCheckComboImport = NewDefaultImports()
		customCheckComboImport.SetImportTable(customCheckComboImportTables)
		customCheckComboImportTables = nil
	}
	return customCheckComboImport
}
