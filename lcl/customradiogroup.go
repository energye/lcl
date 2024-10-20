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

// ICustomRadioGroup Parent: ICustomGroupBox
type ICustomRadioGroup interface {
	ICustomGroupBox
	AutoFill() bool                        // property
	SetAutoFill(AValue bool)               // property
	ItemIndex() int32                      // property
	SetItemIndex(AValue int32)             // property
	Items() IStrings                       // property
	SetItems(AValue IStrings)              // property
	Columns() int32                        // property
	SetColumns(AValue int32)               // property
	ColumnLayout() TColumnLayout           // property
	SetColumnLayout(AValue TColumnLayout)  // property
	CanModify() bool                       // function
	Rows() int32                           // function
	SetOnItemEnter(fn TNotifyEvent)        // property event
	SetOnItemExit(fn TNotifyEvent)         // property event
	SetOnSelectionChanged(fn TNotifyEvent) // property event
}

// TCustomRadioGroup Parent: TCustomGroupBox
type TCustomRadioGroup struct {
	TCustomGroupBox
	itemEnterPtr        uintptr
	itemExitPtr         uintptr
	selectionChangedPtr uintptr
}

func NewCustomRadioGroup(TheOwner IComponent) ICustomRadioGroup {
	r1 := customRadioGroupImportAPI().SysCallN(5, GetObjectUintptr(TheOwner))
	return AsCustomRadioGroup(r1)
}

func (m *TCustomRadioGroup) AutoFill() bool {
	r1 := customRadioGroupImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomRadioGroup) SetAutoFill(AValue bool) {
	customRadioGroupImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomRadioGroup) ItemIndex() int32 {
	r1 := customRadioGroupImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomRadioGroup) SetItemIndex(AValue int32) {
	customRadioGroupImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomRadioGroup) Items() IStrings {
	r1 := customRadioGroupImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomRadioGroup) SetItems(AValue IStrings) {
	customRadioGroupImportAPI().SysCallN(7, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomRadioGroup) Columns() int32 {
	r1 := customRadioGroupImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomRadioGroup) SetColumns(AValue int32) {
	customRadioGroupImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomRadioGroup) ColumnLayout() TColumnLayout {
	r1 := customRadioGroupImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TColumnLayout(r1)
}

func (m *TCustomRadioGroup) SetColumnLayout(AValue TColumnLayout) {
	customRadioGroupImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomRadioGroup) CanModify() bool {
	r1 := customRadioGroupImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCustomRadioGroup) Rows() int32 {
	r1 := customRadioGroupImportAPI().SysCallN(8, m.Instance())
	return int32(r1)
}

func CustomRadioGroupClass() TClass {
	ret := customRadioGroupImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCustomRadioGroup) SetOnItemEnter(fn TNotifyEvent) {
	if m.itemEnterPtr != 0 {
		RemoveEventElement(m.itemEnterPtr)
	}
	m.itemEnterPtr = MakeEventDataPtr(fn)
	customRadioGroupImportAPI().SysCallN(9, m.Instance(), m.itemEnterPtr)
}

func (m *TCustomRadioGroup) SetOnItemExit(fn TNotifyEvent) {
	if m.itemExitPtr != 0 {
		RemoveEventElement(m.itemExitPtr)
	}
	m.itemExitPtr = MakeEventDataPtr(fn)
	customRadioGroupImportAPI().SysCallN(10, m.Instance(), m.itemExitPtr)
}

func (m *TCustomRadioGroup) SetOnSelectionChanged(fn TNotifyEvent) {
	if m.selectionChangedPtr != 0 {
		RemoveEventElement(m.selectionChangedPtr)
	}
	m.selectionChangedPtr = MakeEventDataPtr(fn)
	customRadioGroupImportAPI().SysCallN(11, m.Instance(), m.selectionChangedPtr)
}

var (
	customRadioGroupImport       *imports.Imports = nil
	customRadioGroupImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomRadioGroup_AutoFill", 0),
		/*1*/ imports.NewTable("CustomRadioGroup_CanModify", 0),
		/*2*/ imports.NewTable("CustomRadioGroup_Class", 0),
		/*3*/ imports.NewTable("CustomRadioGroup_ColumnLayout", 0),
		/*4*/ imports.NewTable("CustomRadioGroup_Columns", 0),
		/*5*/ imports.NewTable("CustomRadioGroup_Create", 0),
		/*6*/ imports.NewTable("CustomRadioGroup_ItemIndex", 0),
		/*7*/ imports.NewTable("CustomRadioGroup_Items", 0),
		/*8*/ imports.NewTable("CustomRadioGroup_Rows", 0),
		/*9*/ imports.NewTable("CustomRadioGroup_SetOnItemEnter", 0),
		/*10*/ imports.NewTable("CustomRadioGroup_SetOnItemExit", 0),
		/*11*/ imports.NewTable("CustomRadioGroup_SetOnSelectionChanged", 0),
	}
)

func customRadioGroupImportAPI() *imports.Imports {
	if customRadioGroupImport == nil {
		customRadioGroupImport = NewDefaultImports()
		customRadioGroupImport.SetImportTable(customRadioGroupImportTables)
		customRadioGroupImportTables = nil
	}
	return customRadioGroupImport
}
