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

// IPrinterSetupDialog Parent: ICustomPrinterSetupDialog
type IPrinterSetupDialog interface {
	ICustomPrinterSetupDialog
	AttachTo() ICustomForm                   // property AttachTo Getter
	SetAttachTo(value ICustomForm)           // property AttachTo Setter
	SetOnDialogResult(fn TDialogResultEvent) // property event
}

type TPrinterSetupDialog struct {
	TCustomPrinterSetupDialog
}

func (m *TPrinterSetupDialog) AttachTo() ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := printerSetupDialogAPI().SysCallN(1, 0, m.Instance())
	return AsCustomForm(r)
}

func (m *TPrinterSetupDialog) SetAttachTo(value ICustomForm) {
	if !m.IsValid() {
		return
	}
	printerSetupDialogAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPrinterSetupDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDialogResultEvent(fn)
	base.SetEvent(m, 2, printerSetupDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewPrinterSetupDialog class constructor
func NewPrinterSetupDialog(theOwner IComponent) IPrinterSetupDialog {
	r := printerSetupDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsPrinterSetupDialog(r)
}

func TPrinterSetupDialogClass() types.TClass {
	r := printerSetupDialogAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	printerSetupDialogOnce   base.Once
	printerSetupDialogImport *imports.Imports = nil
)

func printerSetupDialogAPI() *imports.Imports {
	printerSetupDialogOnce.Do(func() {
		printerSetupDialogImport = api.NewDefaultImports()
		printerSetupDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPrinterSetupDialog_Create", 0), // constructor NewPrinterSetupDialog
			/* 1 */ imports.NewTable("TPrinterSetupDialog_AttachTo", 0), // property AttachTo
			/* 2 */ imports.NewTable("TPrinterSetupDialog_OnDialogResult", 0), // event OnDialogResult
			/* 3 */ imports.NewTable("TPrinterSetupDialog_TClass", 0), // function TPrinterSetupDialogClass
		}
	})
	return printerSetupDialogImport
}
