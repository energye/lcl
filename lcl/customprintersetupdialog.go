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

// ICustomPrinterSetupDialog Parent: ICommonDialog
type ICustomPrinterSetupDialog interface {
	ICommonDialog
}

type TCustomPrinterSetupDialog struct {
	TCommonDialog
}

// NewCustomPrinterSetupDialog class constructor
func NewCustomPrinterSetupDialog(theOwner IComponent) ICustomPrinterSetupDialog {
	r := customPrinterSetupDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomPrinterSetupDialog(r)
}

func TCustomPrinterSetupDialogClass() types.TClass {
	r := customPrinterSetupDialogAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	customPrinterSetupDialogOnce   base.Once
	customPrinterSetupDialogImport *imports.Imports = nil
)

func customPrinterSetupDialogAPI() *imports.Imports {
	customPrinterSetupDialogOnce.Do(func() {
		customPrinterSetupDialogImport = api.NewDefaultImports()
		customPrinterSetupDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomPrinterSetupDialog_Create", 0), // constructor NewCustomPrinterSetupDialog
			/* 1 */ imports.NewTable("TCustomPrinterSetupDialog_TClass", 0), // function TCustomPrinterSetupDialogClass
		}
	})
	return customPrinterSetupDialogImport
}
