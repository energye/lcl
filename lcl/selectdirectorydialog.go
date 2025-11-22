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

// ISelectDirectoryDialog Parent: IOpenDialog
type ISelectDirectoryDialog interface {
	IOpenDialog
}

type TSelectDirectoryDialog struct {
	TOpenDialog
}

// NewSelectDirectoryDialog class constructor
func NewSelectDirectoryDialog(owner IComponent) ISelectDirectoryDialog {
	r := selectDirectoryDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSelectDirectoryDialog(r)
}

func TSelectDirectoryDialogClass() types.TClass {
	r := selectDirectoryDialogAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	selectDirectoryDialogOnce   base.Once
	selectDirectoryDialogImport *imports.Imports = nil
)

func selectDirectoryDialogAPI() *imports.Imports {
	selectDirectoryDialogOnce.Do(func() {
		selectDirectoryDialogImport = api.NewDefaultImports()
		selectDirectoryDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSelectDirectoryDialog_Create", 0), // constructor NewSelectDirectoryDialog
			/* 1 */ imports.NewTable("TSelectDirectoryDialog_TClass", 0), // function TSelectDirectoryDialogClass
		}
	})
	return selectDirectoryDialogImport
}
