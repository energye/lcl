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

// ISaveDialog Parent: IOpenDialog
type ISaveDialog interface {
	IOpenDialog
}

// TSaveDialog Parent: TOpenDialog
type TSaveDialog struct {
	TOpenDialog
}

func NewSaveDialog(AOwner IComponent) ISaveDialog {
	r1 := saveDialogImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsSaveDialog(r1)
}

func SaveDialogClass() TClass {
	ret := saveDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	saveDialogImport       *imports.Imports = nil
	saveDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("SaveDialog_Class", 0),
		/*1*/ imports.NewTable("SaveDialog_Create", 0),
	}
)

func saveDialogImportAPI() *imports.Imports {
	if saveDialogImport == nil {
		saveDialogImport = NewDefaultImports()
		saveDialogImport.SetImportTable(saveDialogImportTables)
		saveDialogImportTables = nil
	}
	return saveDialogImport
}
