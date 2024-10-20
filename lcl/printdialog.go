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

// IPrintDialog Parent: ICustomPrintDialog
type IPrintDialog interface {
	ICustomPrintDialog
	AttachTo() ICustomForm                   // property
	SetAttachTo(AValue ICustomForm)          // property
	SetOnDialogResult(fn TDialogResultEvent) // property event
}

// TPrintDialog Parent: TCustomPrintDialog
type TPrintDialog struct {
	TCustomPrintDialog
	dialogResultPtr uintptr
}

func NewPrintDialog(TheOwner IComponent) IPrintDialog {
	r1 := printDialogImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsPrintDialog(r1)
}

func (m *TPrintDialog) AttachTo() ICustomForm {
	r1 := printDialogImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsCustomForm(r1)
}

func (m *TPrintDialog) SetAttachTo(AValue ICustomForm) {
	printDialogImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func PrintDialogClass() TClass {
	ret := printDialogImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TPrintDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if m.dialogResultPtr != 0 {
		RemoveEventElement(m.dialogResultPtr)
	}
	m.dialogResultPtr = MakeEventDataPtr(fn)
	printDialogImportAPI().SysCallN(3, m.Instance(), m.dialogResultPtr)
}

var (
	printDialogImport       *imports.Imports = nil
	printDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PrintDialog_AttachTo", 0),
		/*1*/ imports.NewTable("PrintDialog_Class", 0),
		/*2*/ imports.NewTable("PrintDialog_Create", 0),
		/*3*/ imports.NewTable("PrintDialog_SetOnDialogResult", 0),
	}
)

func printDialogImportAPI() *imports.Imports {
	if printDialogImport == nil {
		printDialogImport = NewDefaultImports()
		printDialogImport.SetImportTable(printDialogImportTables)
		printDialogImportTables = nil
	}
	return printDialogImport
}
