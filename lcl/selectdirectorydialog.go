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

// ISelectDirectoryDialog Parent: IOpenDialog
type ISelectDirectoryDialog interface {
	IOpenDialog
}

// TSelectDirectoryDialog Parent: TOpenDialog
type TSelectDirectoryDialog struct {
	TOpenDialog
}

func NewSelectDirectoryDialog(AOwner IComponent) ISelectDirectoryDialog {
	r1 := selectDirectoryDialogImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsSelectDirectoryDialog(r1)
}

func SelectDirectoryDialogClass() TClass {
	ret := selectDirectoryDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	selectDirectoryDialogImport       *imports.Imports = nil
	selectDirectoryDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("SelectDirectoryDialog_Class", 0),
		/*1*/ imports.NewTable("SelectDirectoryDialog_Create", 0),
	}
)

func selectDirectoryDialogImportAPI() *imports.Imports {
	if selectDirectoryDialogImport == nil {
		selectDirectoryDialogImport = NewDefaultImports()
		selectDirectoryDialogImport.SetImportTable(selectDirectoryDialogImportTables)
		selectDirectoryDialogImportTables = nil
	}
	return selectDirectoryDialogImport
}
