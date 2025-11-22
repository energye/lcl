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

// IOpenPictureDialog Parent: IPreviewFileDialog
type IOpenPictureDialog interface {
	IPreviewFileDialog
	GetFilterExt() string  // function
	DefaultFilter() string // property DefaultFilter Getter
}

type TOpenPictureDialog struct {
	TPreviewFileDialog
}

func (m *TOpenPictureDialog) GetFilterExt() string {
	if !m.IsValid() {
		return ""
	}
	r := openPictureDialogAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TOpenPictureDialog) DefaultFilter() string {
	if !m.IsValid() {
		return ""
	}
	r := openPictureDialogAPI().SysCallN(2, m.Instance())
	return api.GoStr(r)
}

// NewOpenPictureDialog class constructor
func NewOpenPictureDialog(theOwner IComponent) IOpenPictureDialog {
	r := openPictureDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsOpenPictureDialog(r)
}

func TOpenPictureDialogClass() types.TClass {
	r := openPictureDialogAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	openPictureDialogOnce   base.Once
	openPictureDialogImport *imports.Imports = nil
)

func openPictureDialogAPI() *imports.Imports {
	openPictureDialogOnce.Do(func() {
		openPictureDialogImport = api.NewDefaultImports()
		openPictureDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TOpenPictureDialog_Create", 0), // constructor NewOpenPictureDialog
			/* 1 */ imports.NewTable("TOpenPictureDialog_GetFilterExt", 0), // function GetFilterExt
			/* 2 */ imports.NewTable("TOpenPictureDialog_DefaultFilter", 0), // property DefaultFilter
			/* 3 */ imports.NewTable("TOpenPictureDialog_TClass", 0), // function TOpenPictureDialogClass
		}
	})
	return openPictureDialogImport
}
