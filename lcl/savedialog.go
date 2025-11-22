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

// ISaveDialog Parent: IOpenDialog
type ISaveDialog interface {
	IOpenDialog
}

type TSaveDialog struct {
	TOpenDialog
}

// NewSaveDialog class constructor
func NewSaveDialog(owner IComponent) ISaveDialog {
	r := saveDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSaveDialog(r)
}

func TSaveDialogClass() types.TClass {
	r := saveDialogAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	saveDialogOnce   base.Once
	saveDialogImport *imports.Imports = nil
)

func saveDialogAPI() *imports.Imports {
	saveDialogOnce.Do(func() {
		saveDialogImport = api.NewDefaultImports()
		saveDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSaveDialog_Create", 0), // constructor NewSaveDialog
			/* 1 */ imports.NewTable("TSaveDialog_TClass", 0), // function TSaveDialogClass
		}
	})
	return saveDialogImport
}
