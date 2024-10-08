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

// ICustomPrinterSetupDialog Parent: ICommonDialog
type ICustomPrinterSetupDialog interface {
	ICommonDialog
}

// TCustomPrinterSetupDialog Parent: TCommonDialog
type TCustomPrinterSetupDialog struct {
	TCommonDialog
}

func NewCustomPrinterSetupDialog(TheOwner IComponent) ICustomPrinterSetupDialog {
	r1 := LCL().SysCallN(2167, GetObjectUintptr(TheOwner))
	return AsCustomPrinterSetupDialog(r1)
}

func CustomPrinterSetupDialogClass() TClass {
	ret := LCL().SysCallN(2166)
	return TClass(ret)
}
