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

// IPrintDialog Parent: ICustomPrintDialog
type IPrintDialog interface {
	ICustomPrintDialog
	AttachTo() ICustomForm                   // property AttachTo Getter
	SetAttachTo(value ICustomForm)           // property AttachTo Setter
	SetOnDialogResult(fn TDialogResultEvent) // property event
}

type TPrintDialog struct {
	TCustomPrintDialog
}

func (m *TPrintDialog) AttachTo() ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := printDialogAPI().SysCallN(1, 0, m.Instance())
	return AsCustomForm(r)
}

func (m *TPrintDialog) SetAttachTo(value ICustomForm) {
	if !m.IsValid() {
		return
	}
	printDialogAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPrintDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDialogResultEvent(fn)
	base.SetEvent(m, 2, printDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewPrintDialog class constructor
func NewPrintDialog(theOwner IComponent) IPrintDialog {
	r := printDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsPrintDialog(r)
}

func TPrintDialogClass() types.TClass {
	r := printDialogAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	printDialogOnce   base.Once
	printDialogImport *imports.Imports = nil
)

func printDialogAPI() *imports.Imports {
	printDialogOnce.Do(func() {
		printDialogImport = api.NewDefaultImports()
		printDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPrintDialog_Create", 0), // constructor NewPrintDialog
			/* 1 */ imports.NewTable("TPrintDialog_AttachTo", 0), // property AttachTo
			/* 2 */ imports.NewTable("TPrintDialog_OnDialogResult", 0), // event OnDialogResult
			/* 3 */ imports.NewTable("TPrintDialog_TClass", 0), // function TPrintDialogClass
		}
	})
	return printDialogImport
}
