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

// ITaskDialog Parent: ICustomTaskDialog
type ITaskDialog interface {
	ICustomTaskDialog
}

type TTaskDialog struct {
	TCustomTaskDialog
}

// NewTaskDialog class constructor
func NewTaskDialog(owner IComponent) ITaskDialog {
	r := taskDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsTaskDialog(r)
}

func TTaskDialogClass() types.TClass {
	r := taskDialogAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	taskDialogOnce   base.Once
	taskDialogImport *imports.Imports = nil
)

func taskDialogAPI() *imports.Imports {
	taskDialogOnce.Do(func() {
		taskDialogImport = api.NewDefaultImports()
		taskDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTaskDialog_Create", 0), // constructor NewTaskDialog
			/* 1 */ imports.NewTable("TTaskDialog_TClass", 0), // function TTaskDialogClass
		}
	})
	return taskDialogImport
}
