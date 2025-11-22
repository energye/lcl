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
)

// ITaskDialogRadioButtonItem Parent: ITaskDialogBaseButtonItem
type ITaskDialogRadioButtonItem interface {
	ITaskDialogBaseButtonItem
}

type TTaskDialogRadioButtonItem struct {
	TTaskDialogBaseButtonItem
}

// NewTaskDialogRadioButtonItem class constructor
func NewTaskDialogRadioButtonItem(collection ICollection) ITaskDialogRadioButtonItem {
	r := taskDialogRadioButtonItemAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsTaskDialogRadioButtonItem(r)
}

var (
	taskDialogRadioButtonItemOnce   base.Once
	taskDialogRadioButtonItemImport *imports.Imports = nil
)

func taskDialogRadioButtonItemAPI() *imports.Imports {
	taskDialogRadioButtonItemOnce.Do(func() {
		taskDialogRadioButtonItemImport = api.NewDefaultImports()
		taskDialogRadioButtonItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTaskDialogRadioButtonItem_Create", 0), // constructor NewTaskDialogRadioButtonItem
		}
	})
	return taskDialogRadioButtonItemImport
}
