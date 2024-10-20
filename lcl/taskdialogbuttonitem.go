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

// ITaskDialogButtonItem Parent: ITaskDialogBaseButtonItem
type ITaskDialogButtonItem interface {
	ITaskDialogBaseButtonItem
}

// TTaskDialogButtonItem Parent: TTaskDialogBaseButtonItem
type TTaskDialogButtonItem struct {
	TTaskDialogBaseButtonItem
}

func NewTaskDialogButtonItem(ACollection ICollection) ITaskDialogButtonItem {
	r1 := askDialogButtonItemImportAPI().SysCallN(1, GetObjectUintptr(ACollection))
	return AsTaskDialogButtonItem(r1)
}

func TaskDialogButtonItemClass() TClass {
	ret := askDialogButtonItemImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	askDialogButtonItemImport       *imports.Imports = nil
	askDialogButtonItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TaskDialogButtonItem_Class", 0),
		/*1*/ imports.NewTable("TaskDialogButtonItem_Create", 0),
	}
)

func askDialogButtonItemImportAPI() *imports.Imports {
	if askDialogButtonItemImport == nil {
		askDialogButtonItemImport = NewDefaultImports()
		askDialogButtonItemImport.SetImportTable(askDialogButtonItemImportTables)
		askDialogButtonItemImportTables = nil
	}
	return askDialogButtonItemImport
}
