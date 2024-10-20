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

// IPrinterSetupDialog Parent: ICustomPrinterSetupDialog
type IPrinterSetupDialog interface {
	ICustomPrinterSetupDialog
	AttachTo() ICustomForm                   // property
	SetAttachTo(AValue ICustomForm)          // property
	SetOnDialogResult(fn TDialogResultEvent) // property event
}

// TPrinterSetupDialog Parent: TCustomPrinterSetupDialog
type TPrinterSetupDialog struct {
	TCustomPrinterSetupDialog
	dialogResultPtr uintptr
}

func NewPrinterSetupDialog(TheOwner IComponent) IPrinterSetupDialog {
	r1 := printerSetupDialogImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsPrinterSetupDialog(r1)
}

func (m *TPrinterSetupDialog) AttachTo() ICustomForm {
	r1 := printerSetupDialogImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsCustomForm(r1)
}

func (m *TPrinterSetupDialog) SetAttachTo(AValue ICustomForm) {
	printerSetupDialogImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func PrinterSetupDialogClass() TClass {
	ret := printerSetupDialogImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TPrinterSetupDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if m.dialogResultPtr != 0 {
		RemoveEventElement(m.dialogResultPtr)
	}
	m.dialogResultPtr = MakeEventDataPtr(fn)
	printerSetupDialogImportAPI().SysCallN(3, m.Instance(), m.dialogResultPtr)
}

var (
	printerSetupDialogImport       *imports.Imports = nil
	printerSetupDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PrinterSetupDialog_AttachTo", 0),
		/*1*/ imports.NewTable("PrinterSetupDialog_Class", 0),
		/*2*/ imports.NewTable("PrinterSetupDialog_Create", 0),
		/*3*/ imports.NewTable("PrinterSetupDialog_SetOnDialogResult", 0),
	}
)

func printerSetupDialogImportAPI() *imports.Imports {
	if printerSetupDialogImport == nil {
		printerSetupDialogImport = NewDefaultImports()
		printerSetupDialogImport.SetImportTable(printerSetupDialogImportTables)
		printerSetupDialogImportTables = nil
	}
	return printerSetupDialogImport
}
