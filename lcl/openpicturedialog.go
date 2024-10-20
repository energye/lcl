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

// IOpenPictureDialog Parent: IPreviewFileDialog
type IOpenPictureDialog interface {
	IPreviewFileDialog
	DefaultFilter() string // property
	GetFilterExt() string  // function
}

// TOpenPictureDialog Parent: TPreviewFileDialog
type TOpenPictureDialog struct {
	TPreviewFileDialog
}

func NewOpenPictureDialog(TheOwner IComponent) IOpenPictureDialog {
	r1 := openPictureDialogImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsOpenPictureDialog(r1)
}

func (m *TOpenPictureDialog) DefaultFilter() string {
	r1 := openPictureDialogImportAPI().SysCallN(2, m.Instance())
	return GoStr(r1)
}

func (m *TOpenPictureDialog) GetFilterExt() string {
	r1 := openPictureDialogImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func OpenPictureDialogClass() TClass {
	ret := openPictureDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	openPictureDialogImport       *imports.Imports = nil
	openPictureDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("OpenPictureDialog_Class", 0),
		/*1*/ imports.NewTable("OpenPictureDialog_Create", 0),
		/*2*/ imports.NewTable("OpenPictureDialog_DefaultFilter", 0),
		/*3*/ imports.NewTable("OpenPictureDialog_GetFilterExt", 0),
	}
)

func openPictureDialogImportAPI() *imports.Imports {
	if openPictureDialogImport == nil {
		openPictureDialogImport = NewDefaultImports()
		openPictureDialogImport.SetImportTable(openPictureDialogImportTables)
		openPictureDialogImportTables = nil
	}
	return openPictureDialogImport
}
