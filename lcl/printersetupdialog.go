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
	r1 := LCL().SysCallN(4648, GetObjectUintptr(TheOwner))
	return AsPrinterSetupDialog(r1)
}

func (m *TPrinterSetupDialog) AttachTo() ICustomForm {
	r1 := LCL().SysCallN(4646, 0, m.Instance(), 0)
	return AsCustomForm(r1)
}

func (m *TPrinterSetupDialog) SetAttachTo(AValue ICustomForm) {
	LCL().SysCallN(4646, 1, m.Instance(), GetObjectUintptr(AValue))
}

func PrinterSetupDialogClass() TClass {
	ret := LCL().SysCallN(4647)
	return TClass(ret)
}

func (m *TPrinterSetupDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if m.dialogResultPtr != 0 {
		RemoveEventElement(m.dialogResultPtr)
	}
	m.dialogResultPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4649, m.Instance(), m.dialogResultPtr)
}
