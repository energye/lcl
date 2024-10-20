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

// ICustomCheckBox Parent: IButtonControl
type ICustomCheckBox interface {
	IButtonControl
	Alignment() TLeftRight          // property
	SetAlignment(AValue TLeftRight) // property
	AllowGrayed() bool              // property
	SetAllowGrayed(AValue bool)     // property
	State() TCheckBoxState          // property
	SetState(AValue TCheckBoxState) // property
	ShortCut() TShortCut            // property
	ShortCutKey2() TShortCut        // property
	SetOnChange(fn TNotifyEvent)    // property event
}

// TCustomCheckBox Parent: TButtonControl
type TCustomCheckBox struct {
	TButtonControl
	changePtr uintptr
}

func NewCustomCheckBox(TheOwner IComponent) ICustomCheckBox {
	r1 := customCheckBoxImportAPI().SysCallN(3, GetObjectUintptr(TheOwner))
	return AsCustomCheckBox(r1)
}

func (m *TCustomCheckBox) Alignment() TLeftRight {
	r1 := customCheckBoxImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TLeftRight(r1)
}

func (m *TCustomCheckBox) SetAlignment(AValue TLeftRight) {
	customCheckBoxImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckBox) AllowGrayed() bool {
	r1 := customCheckBoxImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCheckBox) SetAllowGrayed(AValue bool) {
	customCheckBoxImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCheckBox) State() TCheckBoxState {
	r1 := customCheckBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TCheckBoxState(r1)
}

func (m *TCustomCheckBox) SetState(AValue TCheckBoxState) {
	customCheckBoxImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCheckBox) ShortCut() TShortCut {
	r1 := customCheckBoxImportAPI().SysCallN(5, m.Instance())
	return TShortCut(r1)
}

func (m *TCustomCheckBox) ShortCutKey2() TShortCut {
	r1 := customCheckBoxImportAPI().SysCallN(6, m.Instance())
	return TShortCut(r1)
}

func CustomCheckBoxClass() TClass {
	ret := customCheckBoxImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCustomCheckBox) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	customCheckBoxImportAPI().SysCallN(4, m.Instance(), m.changePtr)
}

var (
	customCheckBoxImport       *imports.Imports = nil
	customCheckBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomCheckBox_Alignment", 0),
		/*1*/ imports.NewTable("CustomCheckBox_AllowGrayed", 0),
		/*2*/ imports.NewTable("CustomCheckBox_Class", 0),
		/*3*/ imports.NewTable("CustomCheckBox_Create", 0),
		/*4*/ imports.NewTable("CustomCheckBox_SetOnChange", 0),
		/*5*/ imports.NewTable("CustomCheckBox_ShortCut", 0),
		/*6*/ imports.NewTable("CustomCheckBox_ShortCutKey2", 0),
		/*7*/ imports.NewTable("CustomCheckBox_State", 0),
	}
)

func customCheckBoxImportAPI() *imports.Imports {
	if customCheckBoxImport == nil {
		customCheckBoxImport = NewDefaultImports()
		customCheckBoxImport.SetImportTable(customCheckBoxImportTables)
		customCheckBoxImportTables = nil
	}
	return customCheckBoxImport
}
