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

// IFileDialog Parent: ICommonDialog
type IFileDialog interface {
	ICommonDialog
	Files() IStrings                          // property
	HistoryList() IStrings                    // property
	SetHistoryList(AValue IStrings)           // property
	DefaultExt() string                       // property
	SetDefaultExt(AValue string)              // property
	FileName() string                         // property
	SetFileName(AValue string)                // property
	Filter() string                           // property
	SetFilter(AValue string)                  // property
	FilterIndex() int32                       // property
	SetFilterIndex(AValue int32)              // property
	InitialDir() string                       // property
	SetInitialDir(AValue string)              // property
	DoTypeChange()                            // procedure
	IntfFileTypeChanged(NewFilterIndex int32) // procedure
	SetOnHelpClicked(fn TNotifyEvent)         // property event
	SetOnTypeChange(fn TNotifyEvent)          // property event
}

// TFileDialog Parent: TCommonDialog
type TFileDialog struct {
	TCommonDialog
	helpClickedPtr uintptr
	typeChangePtr  uintptr
}

func NewFileDialog(TheOwner IComponent) IFileDialog {
	r1 := fileDialogImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsFileDialog(r1)
}

func (m *TFileDialog) Files() IStrings {
	r1 := fileDialogImportAPI().SysCallN(5, m.Instance())
	return AsStrings(r1)
}

func (m *TFileDialog) HistoryList() IStrings {
	r1 := fileDialogImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TFileDialog) SetHistoryList(AValue IStrings) {
	fileDialogImportAPI().SysCallN(8, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFileDialog) DefaultExt() string {
	r1 := fileDialogImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetDefaultExt(AValue string) {
	fileDialogImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFileDialog) FileName() string {
	r1 := fileDialogImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetFileName(AValue string) {
	fileDialogImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFileDialog) Filter() string {
	r1 := fileDialogImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetFilter(AValue string) {
	fileDialogImportAPI().SysCallN(6, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFileDialog) FilterIndex() int32 {
	r1 := fileDialogImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFileDialog) SetFilterIndex(AValue int32) {
	fileDialogImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TFileDialog) InitialDir() string {
	r1 := fileDialogImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetInitialDir(AValue string) {
	fileDialogImportAPI().SysCallN(9, 1, m.Instance(), PascalStr(AValue))
}

func FileDialogClass() TClass {
	ret := fileDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TFileDialog) DoTypeChange() {
	fileDialogImportAPI().SysCallN(3, m.Instance())
}

func (m *TFileDialog) IntfFileTypeChanged(NewFilterIndex int32) {
	fileDialogImportAPI().SysCallN(10, m.Instance(), uintptr(NewFilterIndex))
}

func (m *TFileDialog) SetOnHelpClicked(fn TNotifyEvent) {
	if m.helpClickedPtr != 0 {
		RemoveEventElement(m.helpClickedPtr)
	}
	m.helpClickedPtr = MakeEventDataPtr(fn)
	fileDialogImportAPI().SysCallN(11, m.Instance(), m.helpClickedPtr)
}

func (m *TFileDialog) SetOnTypeChange(fn TNotifyEvent) {
	if m.typeChangePtr != 0 {
		RemoveEventElement(m.typeChangePtr)
	}
	m.typeChangePtr = MakeEventDataPtr(fn)
	fileDialogImportAPI().SysCallN(12, m.Instance(), m.typeChangePtr)
}

var (
	fileDialogImport       *imports.Imports = nil
	fileDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FileDialog_Class", 0),
		/*1*/ imports.NewTable("FileDialog_Create", 0),
		/*2*/ imports.NewTable("FileDialog_DefaultExt", 0),
		/*3*/ imports.NewTable("FileDialog_DoTypeChange", 0),
		/*4*/ imports.NewTable("FileDialog_FileName", 0),
		/*5*/ imports.NewTable("FileDialog_Files", 0),
		/*6*/ imports.NewTable("FileDialog_Filter", 0),
		/*7*/ imports.NewTable("FileDialog_FilterIndex", 0),
		/*8*/ imports.NewTable("FileDialog_HistoryList", 0),
		/*9*/ imports.NewTable("FileDialog_InitialDir", 0),
		/*10*/ imports.NewTable("FileDialog_IntfFileTypeChanged", 0),
		/*11*/ imports.NewTable("FileDialog_SetOnHelpClicked", 0),
		/*12*/ imports.NewTable("FileDialog_SetOnTypeChange", 0),
	}
)

func fileDialogImportAPI() *imports.Imports {
	if fileDialogImport == nil {
		fileDialogImport = NewDefaultImports()
		fileDialogImport.SetImportTable(fileDialogImportTables)
		fileDialogImportTables = nil
	}
	return fileDialogImport
}
