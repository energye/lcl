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

// IPreviewFileDialog Parent: IOpenDialog
type IPreviewFileDialog interface {
	IOpenDialog
	PreviewFileControl() IPreviewFileControl // property
}

// TPreviewFileDialog Parent: TOpenDialog
type TPreviewFileDialog struct {
	TOpenDialog
}

func NewPreviewFileDialog(TheOwner IComponent) IPreviewFileDialog {
	r1 := previewFileDialogImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsPreviewFileDialog(r1)
}

func (m *TPreviewFileDialog) PreviewFileControl() IPreviewFileControl {
	r1 := previewFileDialogImportAPI().SysCallN(2, m.Instance())
	return AsPreviewFileControl(r1)
}

func PreviewFileDialogClass() TClass {
	ret := previewFileDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	previewFileDialogImport       *imports.Imports = nil
	previewFileDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PreviewFileDialog_Class", 0),
		/*1*/ imports.NewTable("PreviewFileDialog_Create", 0),
		/*2*/ imports.NewTable("PreviewFileDialog_PreviewFileControl", 0),
	}
)

func previewFileDialogImportAPI() *imports.Imports {
	if previewFileDialogImport == nil {
		previewFileDialogImport = NewDefaultImports()
		previewFileDialogImport.SetImportTable(previewFileDialogImportTables)
		previewFileDialogImportTables = nil
	}
	return previewFileDialogImport
}
