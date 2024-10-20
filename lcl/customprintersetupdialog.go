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

// ICustomPrinterSetupDialog Parent: ICommonDialog
type ICustomPrinterSetupDialog interface {
	ICommonDialog
}

// TCustomPrinterSetupDialog Parent: TCommonDialog
type TCustomPrinterSetupDialog struct {
	TCommonDialog
}

func NewCustomPrinterSetupDialog(TheOwner IComponent) ICustomPrinterSetupDialog {
	r1 := customPrinterSetupDialogImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsCustomPrinterSetupDialog(r1)
}

func CustomPrinterSetupDialogClass() TClass {
	ret := customPrinterSetupDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customPrinterSetupDialogImport       *imports.Imports = nil
	customPrinterSetupDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomPrinterSetupDialog_Class", 0),
		/*1*/ imports.NewTable("CustomPrinterSetupDialog_Create", 0),
	}
)

func customPrinterSetupDialogImportAPI() *imports.Imports {
	if customPrinterSetupDialogImport == nil {
		customPrinterSetupDialogImport = NewDefaultImports()
		customPrinterSetupDialogImport.SetImportTable(customPrinterSetupDialogImportTables)
		customPrinterSetupDialogImportTables = nil
	}
	return customPrinterSetupDialogImport
}
