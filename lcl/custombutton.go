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

// ICustomButton Parent: IButtonControl
type ICustomButton interface {
	IButtonControl
	Active() bool                       // property
	Default() bool                      // property
	SetDefault(AValue bool)             // property
	ModalResult() TModalResult          // property
	SetModalResult(AValue TModalResult) // property
	ShortCut() TShortCut                // property
	ShortCutKey2() TShortCut            // property
	Cancel() bool                       // property
	SetCancel(AValue bool)              // property
	Click()                             // procedure
}

// TCustomButton Parent: TButtonControl
type TCustomButton struct {
	TButtonControl
}

func NewCustomButton(TheOwner IComponent) ICustomButton {
	r1 := customButtonImportAPI().SysCallN(4, GetObjectUintptr(TheOwner))
	return AsCustomButton(r1)
}

func (m *TCustomButton) Active() bool {
	r1 := customButtonImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCustomButton) Default() bool {
	r1 := customButtonImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomButton) SetDefault(AValue bool) {
	customButtonImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomButton) ModalResult() TModalResult {
	r1 := customButtonImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TCustomButton) SetModalResult(AValue TModalResult) {
	customButtonImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomButton) ShortCut() TShortCut {
	r1 := customButtonImportAPI().SysCallN(7, m.Instance())
	return TShortCut(r1)
}

func (m *TCustomButton) ShortCutKey2() TShortCut {
	r1 := customButtonImportAPI().SysCallN(8, m.Instance())
	return TShortCut(r1)
}

func (m *TCustomButton) Cancel() bool {
	r1 := customButtonImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomButton) SetCancel(AValue bool) {
	customButtonImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func CustomButtonClass() TClass {
	ret := customButtonImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCustomButton) Click() {
	customButtonImportAPI().SysCallN(3, m.Instance())
}

var (
	customButtonImport       *imports.Imports = nil
	customButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomButton_Active", 0),
		/*1*/ imports.NewTable("CustomButton_Cancel", 0),
		/*2*/ imports.NewTable("CustomButton_Class", 0),
		/*3*/ imports.NewTable("CustomButton_Click", 0),
		/*4*/ imports.NewTable("CustomButton_Create", 0),
		/*5*/ imports.NewTable("CustomButton_Default", 0),
		/*6*/ imports.NewTable("CustomButton_ModalResult", 0),
		/*7*/ imports.NewTable("CustomButton_ShortCut", 0),
		/*8*/ imports.NewTable("CustomButton_ShortCutKey2", 0),
	}
)

func customButtonImportAPI() *imports.Imports {
	if customButtonImport == nil {
		customButtonImport = NewDefaultImports()
		customButtonImport.SetImportTable(customButtonImportTables)
		customButtonImportTables = nil
	}
	return customButtonImport
}
