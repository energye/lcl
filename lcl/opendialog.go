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

// IOpenDialog Parent: IFileDialog
type IOpenDialog interface {
	IFileDialog
	DoFolderChange()                                    // procedure
	DoSelectionChange()                                 // procedure
	IntfSetOption(option types.TOpenOption, value bool) // procedure
	Options() types.TOpenOptions                        // property Options Getter
	SetOptions(value types.TOpenOptions)                // property Options Setter
	OptionsEx() types.TOpenOptionsEx                    // property OptionsEx Getter
	SetOptionsEx(value types.TOpenOptionsEx)            // property OptionsEx Setter
	SetOnFolderChange(fn TNotifyEvent)                  // property event
	SetOnSelectionChange(fn TNotifyEvent)               // property event
}

type TOpenDialog struct {
	TFileDialog
}

func (m *TOpenDialog) DoFolderChange() {
	if !m.IsValid() {
		return
	}
	openDialogAPI().SysCallN(1, m.Instance())
}

func (m *TOpenDialog) DoSelectionChange() {
	if !m.IsValid() {
		return
	}
	openDialogAPI().SysCallN(2, m.Instance())
}

func (m *TOpenDialog) IntfSetOption(option types.TOpenOption, value bool) {
	if !m.IsValid() {
		return
	}
	openDialogAPI().SysCallN(3, m.Instance(), uintptr(option), api.PasBool(value))
}

func (m *TOpenDialog) Options() types.TOpenOptions {
	if !m.IsValid() {
		return 0
	}
	r := openDialogAPI().SysCallN(4, 0, m.Instance())
	return types.TOpenOptions(r)
}

func (m *TOpenDialog) SetOptions(value types.TOpenOptions) {
	if !m.IsValid() {
		return
	}
	openDialogAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TOpenDialog) OptionsEx() types.TOpenOptionsEx {
	if !m.IsValid() {
		return 0
	}
	r := openDialogAPI().SysCallN(5, 0, m.Instance())
	return types.TOpenOptionsEx(r)
}

func (m *TOpenDialog) SetOptionsEx(value types.TOpenOptionsEx) {
	if !m.IsValid() {
		return
	}
	openDialogAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TOpenDialog) SetOnFolderChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, openDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TOpenDialog) SetOnSelectionChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, openDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewOpenDialog class constructor
func NewOpenDialog(theOwner IComponent) IOpenDialog {
	r := openDialogAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsOpenDialog(r)
}

func TOpenDialogClass() types.TClass {
	r := openDialogAPI().SysCallN(8)
	return types.TClass(r)
}

var (
	openDialogOnce   base.Once
	openDialogImport *imports.Imports = nil
)

func openDialogAPI() *imports.Imports {
	openDialogOnce.Do(func() {
		openDialogImport = api.NewDefaultImports()
		openDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TOpenDialog_Create", 0), // constructor NewOpenDialog
			/* 1 */ imports.NewTable("TOpenDialog_DoFolderChange", 0), // procedure DoFolderChange
			/* 2 */ imports.NewTable("TOpenDialog_DoSelectionChange", 0), // procedure DoSelectionChange
			/* 3 */ imports.NewTable("TOpenDialog_IntfSetOption", 0), // procedure IntfSetOption
			/* 4 */ imports.NewTable("TOpenDialog_Options", 0), // property Options
			/* 5 */ imports.NewTable("TOpenDialog_OptionsEx", 0), // property OptionsEx
			/* 6 */ imports.NewTable("TOpenDialog_OnFolderChange", 0), // event OnFolderChange
			/* 7 */ imports.NewTable("TOpenDialog_OnSelectionChange", 0), // event OnSelectionChange
			/* 8 */ imports.NewTable("TOpenDialog_TClass", 0), // function TOpenDialogClass
		}
	})
	return openDialogImport
}
