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

// IPreviewFileControl Parent: IWinControl
type IPreviewFileControl interface {
	IWinControl
	PreviewFileDialog() IPreviewFileDialog         // property PreviewFileDialog Getter
	SetPreviewFileDialog(value IPreviewFileDialog) // property PreviewFileDialog Setter
}

type TPreviewFileControl struct {
	TWinControl
}

func (m *TPreviewFileControl) PreviewFileDialog() IPreviewFileDialog {
	if !m.IsValid() {
		return nil
	}
	r := previewFileControlAPI().SysCallN(1, 0, m.Instance())
	return AsPreviewFileDialog(r)
}

func (m *TPreviewFileControl) SetPreviewFileDialog(value IPreviewFileDialog) {
	if !m.IsValid() {
		return
	}
	previewFileControlAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewPreviewFileControl class constructor
func NewPreviewFileControl(theOwner IComponent) IPreviewFileControl {
	r := previewFileControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsPreviewFileControl(r)
}

func TPreviewFileControlClass() types.TClass {
	r := previewFileControlAPI().SysCallN(2)
	return types.TClass(r)
}

var (
	previewFileControlOnce   base.Once
	previewFileControlImport *imports.Imports = nil
)

func previewFileControlAPI() *imports.Imports {
	previewFileControlOnce.Do(func() {
		previewFileControlImport = api.NewDefaultImports()
		previewFileControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPreviewFileControl_Create", 0), // constructor NewPreviewFileControl
			/* 1 */ imports.NewTable("TPreviewFileControl_PreviewFileDialog", 0), // property PreviewFileDialog
			/* 2 */ imports.NewTable("TPreviewFileControl_TClass", 0), // function TPreviewFileControlClass
		}
	})
	return previewFileControlImport
}
