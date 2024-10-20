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

// ITaskDialog Parent: ICustomTaskDialog
type ITaskDialog interface {
	ICustomTaskDialog
}

// TTaskDialog Parent: TCustomTaskDialog
type TTaskDialog struct {
	TCustomTaskDialog
}

func NewTaskDialog(AOwner IComponent) ITaskDialog {
	r1 := askDialogImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsTaskDialog(r1)
}

func TaskDialogClass() TClass {
	ret := askDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	askDialogImport       *imports.Imports = nil
	askDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TaskDialog_Class", 0),
		/*1*/ imports.NewTable("TaskDialog_Create", 0),
	}
)

func askDialogImportAPI() *imports.Imports {
	if askDialogImport == nil {
		askDialogImport = NewDefaultImports()
		askDialogImport.SetImportTable(askDialogImportTables)
		askDialogImportTables = nil
	}
	return askDialogImport
}
