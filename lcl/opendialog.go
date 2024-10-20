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

// IOpenDialog Parent: IFileDialog
type IOpenDialog interface {
	IFileDialog
	Options() TOpenOptions                          // property
	SetOptions(AValue TOpenOptions)                 // property
	DoFolderChange()                                // procedure
	DoSelectionChange()                             // procedure
	IntfSetOption(AOption TOpenOption, AValue bool) // procedure
	SetOnFolderChange(fn TNotifyEvent)              // property event
	SetOnSelectionChange(fn TNotifyEvent)           // property event
}

// TOpenDialog Parent: TFileDialog
type TOpenDialog struct {
	TFileDialog
	folderChangePtr    uintptr
	selectionChangePtr uintptr
}

func NewOpenDialog(TheOwner IComponent) IOpenDialog {
	r1 := openDialogImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsOpenDialog(r1)
}

func (m *TOpenDialog) Options() TOpenOptions {
	r1 := openDialogImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TOpenOptions(r1)
}

func (m *TOpenDialog) SetOptions(AValue TOpenOptions) {
	openDialogImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func OpenDialogClass() TClass {
	ret := openDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TOpenDialog) DoFolderChange() {
	openDialogImportAPI().SysCallN(2, m.Instance())
}

func (m *TOpenDialog) DoSelectionChange() {
	openDialogImportAPI().SysCallN(3, m.Instance())
}

func (m *TOpenDialog) IntfSetOption(AOption TOpenOption, AValue bool) {
	openDialogImportAPI().SysCallN(4, m.Instance(), uintptr(AOption), PascalBool(AValue))
}

func (m *TOpenDialog) SetOnFolderChange(fn TNotifyEvent) {
	if m.folderChangePtr != 0 {
		RemoveEventElement(m.folderChangePtr)
	}
	m.folderChangePtr = MakeEventDataPtr(fn)
	openDialogImportAPI().SysCallN(6, m.Instance(), m.folderChangePtr)
}

func (m *TOpenDialog) SetOnSelectionChange(fn TNotifyEvent) {
	if m.selectionChangePtr != 0 {
		RemoveEventElement(m.selectionChangePtr)
	}
	m.selectionChangePtr = MakeEventDataPtr(fn)
	openDialogImportAPI().SysCallN(7, m.Instance(), m.selectionChangePtr)
}

var (
	openDialogImport       *imports.Imports = nil
	openDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("OpenDialog_Class", 0),
		/*1*/ imports.NewTable("OpenDialog_Create", 0),
		/*2*/ imports.NewTable("OpenDialog_DoFolderChange", 0),
		/*3*/ imports.NewTable("OpenDialog_DoSelectionChange", 0),
		/*4*/ imports.NewTable("OpenDialog_IntfSetOption", 0),
		/*5*/ imports.NewTable("OpenDialog_Options", 0),
		/*6*/ imports.NewTable("OpenDialog_SetOnFolderChange", 0),
		/*7*/ imports.NewTable("OpenDialog_SetOnSelectionChange", 0),
	}
)

func openDialogImportAPI() *imports.Imports {
	if openDialogImport == nil {
		openDialogImport = NewDefaultImports()
		openDialogImport.SetImportTable(openDialogImportTables)
		openDialogImportTables = nil
	}
	return openDialogImport
}
