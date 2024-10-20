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

// ISavePictureDialog Parent: IOpenPictureDialog
type ISavePictureDialog interface {
	IOpenPictureDialog
}

// TSavePictureDialog Parent: TOpenPictureDialog
type TSavePictureDialog struct {
	TOpenPictureDialog
}

func NewSavePictureDialog(TheOwner IComponent) ISavePictureDialog {
	r1 := savePictureDialogImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsSavePictureDialog(r1)
}

func SavePictureDialogClass() TClass {
	ret := savePictureDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	savePictureDialogImport       *imports.Imports = nil
	savePictureDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("SavePictureDialog_Class", 0),
		/*1*/ imports.NewTable("SavePictureDialog_Create", 0),
	}
)

func savePictureDialogImportAPI() *imports.Imports {
	if savePictureDialogImport == nil {
		savePictureDialogImport = NewDefaultImports()
		savePictureDialogImport.SetImportTable(savePictureDialogImportTables)
		savePictureDialogImportTables = nil
	}
	return savePictureDialogImport
}
