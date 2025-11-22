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

// ISavePictureDialog Parent: IOpenPictureDialog
type ISavePictureDialog interface {
	IOpenPictureDialog
}

type TSavePictureDialog struct {
	TOpenPictureDialog
}

// NewSavePictureDialog class constructor
func NewSavePictureDialog(theOwner IComponent) ISavePictureDialog {
	r := savePictureDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsSavePictureDialog(r)
}

func TSavePictureDialogClass() types.TClass {
	r := savePictureDialogAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	savePictureDialogOnce   base.Once
	savePictureDialogImport *imports.Imports = nil
)

func savePictureDialogAPI() *imports.Imports {
	savePictureDialogOnce.Do(func() {
		savePictureDialogImport = api.NewDefaultImports()
		savePictureDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSavePictureDialog_Create", 0), // constructor NewSavePictureDialog
			/* 1 */ imports.NewTable("TSavePictureDialog_TClass", 0), // function TSavePictureDialogClass
		}
	})
	return savePictureDialogImport
}
