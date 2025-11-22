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

// IPreviewFileDialog Parent: IOpenDialog
type IPreviewFileDialog interface {
	IOpenDialog
	PreviewFileControl() IPreviewFileControl // property PreviewFileControl Getter
}

type TPreviewFileDialog struct {
	TOpenDialog
}

func (m *TPreviewFileDialog) PreviewFileControl() IPreviewFileControl {
	if !m.IsValid() {
		return nil
	}
	r := previewFileDialogAPI().SysCallN(1, m.Instance())
	return AsPreviewFileControl(r)
}

// NewPreviewFileDialog class constructor
func NewPreviewFileDialog(theOwner IComponent) IPreviewFileDialog {
	r := previewFileDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsPreviewFileDialog(r)
}

func TPreviewFileDialogClass() types.TClass {
	r := previewFileDialogAPI().SysCallN(2)
	return types.TClass(r)
}

var (
	previewFileDialogOnce   base.Once
	previewFileDialogImport *imports.Imports = nil
)

func previewFileDialogAPI() *imports.Imports {
	previewFileDialogOnce.Do(func() {
		previewFileDialogImport = api.NewDefaultImports()
		previewFileDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPreviewFileDialog_Create", 0), // constructor NewPreviewFileDialog
			/* 1 */ imports.NewTable("TPreviewFileDialog_PreviewFileControl", 0), // property PreviewFileControl
			/* 2 */ imports.NewTable("TPreviewFileDialog_TClass", 0), // function TPreviewFileDialogClass
		}
	})
	return previewFileDialogImport
}
