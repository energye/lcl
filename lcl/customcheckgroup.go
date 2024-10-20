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

// ICustomCheckGroup Parent: ICustomGroupBox
type ICustomCheckGroup interface {
	ICustomGroupBox
	AutoFill() bool                           // property
	SetAutoFill(AValue bool)                  // property
	Items() IStrings                          // property
	SetItems(AValue IStrings)                 // property
	Checked(Index int32) bool                 // property
	SetChecked(Index int32, AValue bool)      // property
	CheckEnabled(Index int32) bool            // property
	SetCheckEnabled(Index int32, AValue bool) // property
	Columns() int32                           // property
	SetColumns(AValue int32)                  // property
	ColumnLayout() TColumnLayout              // property
	SetColumnLayout(AValue TColumnLayout)     // property
	Rows() int32                              // function
}

// TCustomCheckGroup Parent: TCustomGroupBox
type TCustomCheckGroup struct {
	TCustomGroupBox
}

func NewCustomCheckGroup(TheOwner IComponent) ICustomCheckGroup {
	r1 := customCheckGroupImportAPI().SysCallN(6, GetObjectUintptr(TheOwner))
	return AsCustomCheckGroup(r1)
}

func (m *TCustomCheckGroup) AutoFill() bool {
	r1 := customCheckGroupImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCheckGroup) SetAutoFill(AValue bool) {
	customCheckGroupImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCheckGroup) Items() IStrings {
	r1 := customCheckGroupImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomCheckGroup) SetItems(AValue IStrings) {
	customCheckGroupImportAPI().SysCallN(7, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomCheckGroup) Checked(Index int32) bool {
	r1 := customCheckGroupImportAPI().SysCallN(2, 0, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomCheckGroup) SetChecked(Index int32, AValue bool) {
	customCheckGroupImportAPI().SysCallN(2, 1, m.Instance(), uintptr(Index), PascalBool(AValue))
}

func (m *TCustomCheckGroup) CheckEnabled(Index int32) bool {
	r1 := customCheckGroupImportAPI().SysCallN(1, 0, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomCheckGroup) SetCheckEnabled(Index int32, AValue bool) {
	customCheckGroupImportAPI().SysCallN(1, 1, m.Instance(), uintptr(Index), PascalBool(AValue))
}

func (m *TCustomCheckGroup) Columns() int32 {
	r1 := customCheckGroupImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCheckGroup) SetColumns(AValue int32) {
	customCheckGroupImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckGroup) ColumnLayout() TColumnLayout {
	r1 := customCheckGroupImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TColumnLayout(r1)
}

func (m *TCustomCheckGroup) SetColumnLayout(AValue TColumnLayout) {
	customCheckGroupImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckGroup) Rows() int32 {
	r1 := customCheckGroupImportAPI().SysCallN(8, m.Instance())
	return int32(r1)
}

func CustomCheckGroupClass() TClass {
	ret := customCheckGroupImportAPI().SysCallN(3)
	return TClass(ret)
}

var (
	customCheckGroupImport       *imports.Imports = nil
	customCheckGroupImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCheckGroup_AutoFill", 0),
		/*1*/ imports.NewTable("CustomCheckGroup_CheckEnabled", 0),
		/*2*/ imports.NewTable("CustomCheckGroup_Checked", 0),
		/*3*/ imports.NewTable("CustomCheckGroup_Class", 0),
		/*4*/ imports.NewTable("CustomCheckGroup_ColumnLayout", 0),
		/*5*/ imports.NewTable("CustomCheckGroup_Columns", 0),
		/*6*/ imports.NewTable("CustomCheckGroup_Create", 0),
		/*7*/ imports.NewTable("CustomCheckGroup_Items", 0),
		/*8*/ imports.NewTable("CustomCheckGroup_Rows", 0),
	}
)

func customCheckGroupImportAPI() *imports.Imports {
	if customCheckGroupImport == nil {
		customCheckGroupImport = NewDefaultImports()
		customCheckGroupImport.SetImportTable(customCheckGroupImportTables)
		customCheckGroupImportTables = nil
	}
	return customCheckGroupImport
}
