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

// ITaskDialogButtonItem Parent: ITaskDialogBaseButtonItem
type ITaskDialogButtonItem interface {
	ITaskDialogBaseButtonItem
}

type TTaskDialogButtonItem struct {
	TTaskDialogBaseButtonItem
}

// NewTaskDialogButtonItem class constructor
func NewTaskDialogButtonItem(collection ICollection) ITaskDialogButtonItem {
	r := taskDialogButtonItemAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsTaskDialogButtonItem(r)
}

var (
	taskDialogButtonItemOnce   base.Once
	taskDialogButtonItemImport *imports.Imports = nil
)

func taskDialogButtonItemAPI() *imports.Imports {
	taskDialogButtonItemOnce.Do(func() {
		taskDialogButtonItemImport = api.NewDefaultImports()
		taskDialogButtonItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTaskDialogButtonItem_Create", 0), // constructor NewTaskDialogButtonItem
		}
	})
	return taskDialogButtonItemImport
}
