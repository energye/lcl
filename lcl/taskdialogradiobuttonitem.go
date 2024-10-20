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

// ITaskDialogRadioButtonItem Parent: ITaskDialogBaseButtonItem
type ITaskDialogRadioButtonItem interface {
	ITaskDialogBaseButtonItem
}

// TTaskDialogRadioButtonItem Parent: TTaskDialogBaseButtonItem
type TTaskDialogRadioButtonItem struct {
	TTaskDialogBaseButtonItem
}

func NewTaskDialogRadioButtonItem(ACollection ICollection) ITaskDialogRadioButtonItem {
	r1 := askDialogRadioButtonItemImportAPI().SysCallN(1, GetObjectUintptr(ACollection))
	return AsTaskDialogRadioButtonItem(r1)
}

func TaskDialogRadioButtonItemClass() TClass {
	ret := askDialogRadioButtonItemImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	askDialogRadioButtonItemImport       *imports.Imports = nil
	askDialogRadioButtonItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TaskDialogRadioButtonItem_Class", 0),
		/*1*/ imports.NewTable("TaskDialogRadioButtonItem_Create", 0),
	}
)

func askDialogRadioButtonItemImportAPI() *imports.Imports {
	if askDialogRadioButtonItemImport == nil {
		askDialogRadioButtonItemImport = NewDefaultImports()
		askDialogRadioButtonItemImport.SetImportTable(askDialogRadioButtonItemImportTables)
		askDialogRadioButtonItemImportTables = nil
	}
	return askDialogRadioButtonItemImport
}
