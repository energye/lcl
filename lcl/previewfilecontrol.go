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

// IPreviewFileControl Parent: IWinControl
type IPreviewFileControl interface {
	IWinControl
	PreviewFileDialog() IPreviewFileDialog          // property
	SetPreviewFileDialog(AValue IPreviewFileDialog) // property
}

// TPreviewFileControl Parent: TWinControl
type TPreviewFileControl struct {
	TWinControl
}

func NewPreviewFileControl(TheOwner IComponent) IPreviewFileControl {
	r1 := previewFileControlImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsPreviewFileControl(r1)
}

func (m *TPreviewFileControl) PreviewFileDialog() IPreviewFileDialog {
	r1 := previewFileControlImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return AsPreviewFileDialog(r1)
}

func (m *TPreviewFileControl) SetPreviewFileDialog(AValue IPreviewFileDialog) {
	previewFileControlImportAPI().SysCallN(2, 1, m.Instance(), GetObjectUintptr(AValue))
}

func PreviewFileControlClass() TClass {
	ret := previewFileControlImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	previewFileControlImport       *imports.Imports = nil
	previewFileControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PreviewFileControl_Class", 0),
		/*1*/ imports.NewTable("PreviewFileControl_Create", 0),
		/*2*/ imports.NewTable("PreviewFileControl_PreviewFileDialog", 0),
	}
)

func previewFileControlImportAPI() *imports.Imports {
	if previewFileControlImport == nil {
		previewFileControlImport = NewDefaultImports()
		previewFileControlImport.SetImportTable(previewFileControlImportTables)
		previewFileControlImportTables = nil
	}
	return previewFileControlImport
}
