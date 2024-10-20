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

// ICustomCheckListBox Parent: ICustomListBox
type ICustomCheckListBox interface {
	ICustomListBox
	AllowGrayed() bool                                                      // property
	SetAllowGrayed(AValue bool)                                             // property
	Checked(AIndex int32) bool                                              // property
	SetChecked(AIndex int32, AValue bool)                                   // property
	Header(AIndex int32) bool                                               // property
	SetHeader(AIndex int32, AValue bool)                                    // property
	HeaderBackgroundColor() TColor                                          // property
	SetHeaderBackgroundColor(AValue TColor)                                 // property
	HeaderColor() TColor                                                    // property
	SetHeaderColor(AValue TColor)                                           // property
	ItemEnabled(AIndex int32) bool                                          // property
	SetItemEnabled(AIndex int32, AValue bool)                               // property
	State(AIndex int32) TCheckBoxState                                      // property
	SetState(AIndex int32, AValue TCheckBoxState)                           // property
	CalculateStandardItemHeight() int32                                     // function
	Toggle(AIndex int32)                                                    // procedure
	CheckAll(AState TCheckBoxState, aAllowGrayed bool, aAllowDisabled bool) // procedure
	Exchange(AIndex1, AIndex2 int32)                                        // procedure
	SetOnClickCheck(fn TNotifyEvent)                                        // property event
}

// TCustomCheckListBox Parent: TCustomListBox
type TCustomCheckListBox struct {
	TCustomListBox
	clickCheckPtr uintptr
}

func NewCustomCheckListBox(AOwner IComponent) ICustomCheckListBox {
	r1 := customCheckListBoxImportAPI().SysCallN(5, GetObjectUintptr(AOwner))
	return AsCustomCheckListBox(r1)
}

func (m *TCustomCheckListBox) AllowGrayed() bool {
	r1 := customCheckListBoxImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCheckListBox) SetAllowGrayed(AValue bool) {
	customCheckListBoxImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCheckListBox) Checked(AIndex int32) bool {
	r1 := customCheckListBoxImportAPI().SysCallN(3, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckListBox) SetChecked(AIndex int32, AValue bool) {
	customCheckListBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckListBox) Header(AIndex int32) bool {
	r1 := customCheckListBoxImportAPI().SysCallN(7, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckListBox) SetHeader(AIndex int32, AValue bool) {
	customCheckListBoxImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckListBox) HeaderBackgroundColor() TColor {
	r1 := customCheckListBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomCheckListBox) SetHeaderBackgroundColor(AValue TColor) {
	customCheckListBoxImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckListBox) HeaderColor() TColor {
	r1 := customCheckListBoxImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomCheckListBox) SetHeaderColor(AValue TColor) {
	customCheckListBoxImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckListBox) ItemEnabled(AIndex int32) bool {
	r1 := customCheckListBoxImportAPI().SysCallN(10, 0, m.Instance(), uintptr(AIndex))
	return GoBool(r1)
}

func (m *TCustomCheckListBox) SetItemEnabled(AIndex int32, AValue bool) {
	customCheckListBoxImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AIndex), PascalBool(AValue))
}

func (m *TCustomCheckListBox) State(AIndex int32) TCheckBoxState {
	r1 := customCheckListBoxImportAPI().SysCallN(12, 0, m.Instance(), uintptr(AIndex))
	return TCheckBoxState(r1)
}

func (m *TCustomCheckListBox) SetState(AIndex int32, AValue TCheckBoxState) {
	customCheckListBoxImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AIndex), uintptr(AValue))
}

func (m *TCustomCheckListBox) CalculateStandardItemHeight() int32 {
	r1 := customCheckListBoxImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func CustomCheckListBoxClass() TClass {
	ret := customCheckListBoxImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TCustomCheckListBox) Toggle(AIndex int32) {
	customCheckListBoxImportAPI().SysCallN(13, m.Instance(), uintptr(AIndex))
}

func (m *TCustomCheckListBox) CheckAll(AState TCheckBoxState, aAllowGrayed bool, aAllowDisabled bool) {
	customCheckListBoxImportAPI().SysCallN(2, m.Instance(), uintptr(AState), PascalBool(aAllowGrayed), PascalBool(aAllowDisabled))
}

func (m *TCustomCheckListBox) Exchange(AIndex1, AIndex2 int32) {
	customCheckListBoxImportAPI().SysCallN(6, m.Instance(), uintptr(AIndex1), uintptr(AIndex2))
}

func (m *TCustomCheckListBox) SetOnClickCheck(fn TNotifyEvent) {
	if m.clickCheckPtr != 0 {
		RemoveEventElement(m.clickCheckPtr)
	}
	m.clickCheckPtr = MakeEventDataPtr(fn)
	customCheckListBoxImportAPI().SysCallN(11, m.Instance(), m.clickCheckPtr)
}

var (
	customCheckListBoxImport       *imports.Imports = nil
	customCheckListBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCheckListBox_AllowGrayed", 0),
		/*1*/ imports.NewTable("CustomCheckListBox_CalculateStandardItemHeight", 0),
		/*2*/ imports.NewTable("CustomCheckListBox_CheckAll", 0),
		/*3*/ imports.NewTable("CustomCheckListBox_Checked", 0),
		/*4*/ imports.NewTable("CustomCheckListBox_Class", 0),
		/*5*/ imports.NewTable("CustomCheckListBox_Create", 0),
		/*6*/ imports.NewTable("CustomCheckListBox_Exchange", 0),
		/*7*/ imports.NewTable("CustomCheckListBox_Header", 0),
		/*8*/ imports.NewTable("CustomCheckListBox_HeaderBackgroundColor", 0),
		/*9*/ imports.NewTable("CustomCheckListBox_HeaderColor", 0),
		/*10*/ imports.NewTable("CustomCheckListBox_ItemEnabled", 0),
		/*11*/ imports.NewTable("CustomCheckListBox_SetOnClickCheck", 0),
		/*12*/ imports.NewTable("CustomCheckListBox_State", 0),
		/*13*/ imports.NewTable("CustomCheckListBox_Toggle", 0),
	}
)

func customCheckListBoxImportAPI() *imports.Imports {
	if customCheckListBoxImport == nil {
		customCheckListBoxImport = NewDefaultImports()
		customCheckListBoxImport.SetImportTable(customCheckListBoxImportTables)
		customCheckListBoxImportTables = nil
	}
	return customCheckListBoxImport
}
