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

// IFindDialog Parent: ICommonDialog
type IFindDialog interface {
	ICommonDialog
	Left() int32                      // property
	SetLeft(AValue int32)             // property
	Position() (resultPoint TPoint)   // property
	SetPosition(AValue *TPoint)       // property
	Top() int32                       // property
	SetTop(AValue int32)              // property
	FindText() string                 // property
	SetFindText(AValue string)        // property
	Options() TFindOptions            // property
	SetOptions(AValue TFindOptions)   // property
	CloseDialog()                     // procedure
	SetOnFind(fn TNotifyEvent)        // property event
	SetOnHelpClicked(fn TNotifyEvent) // property event
}

// TFindDialog Parent: TCommonDialog
type TFindDialog struct {
	TCommonDialog
	findPtr        uintptr
	helpClickedPtr uintptr
}

func NewFindDialog(AOwner IComponent) IFindDialog {
	r1 := findDialogImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsFindDialog(r1)
}

func (m *TFindDialog) Left() int32 {
	r1 := findDialogImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFindDialog) SetLeft(AValue int32) {
	findDialogImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TFindDialog) Position() (resultPoint TPoint) {
	findDialogImportAPI().SysCallN(6, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TFindDialog) SetPosition(AValue *TPoint) {
	findDialogImportAPI().SysCallN(6, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TFindDialog) Top() int32 {
	r1 := findDialogImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFindDialog) SetTop(AValue int32) {
	findDialogImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TFindDialog) FindText() string {
	r1 := findDialogImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFindDialog) SetFindText(AValue string) {
	findDialogImportAPI().SysCallN(3, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFindDialog) Options() TFindOptions {
	r1 := findDialogImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TFindOptions(r1)
}

func (m *TFindDialog) SetOptions(AValue TFindOptions) {
	findDialogImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func FindDialogClass() TClass {
	ret := findDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TFindDialog) CloseDialog() {
	findDialogImportAPI().SysCallN(1, m.Instance())
}

func (m *TFindDialog) SetOnFind(fn TNotifyEvent) {
	if m.findPtr != 0 {
		RemoveEventElement(m.findPtr)
	}
	m.findPtr = MakeEventDataPtr(fn)
	findDialogImportAPI().SysCallN(7, m.Instance(), m.findPtr)
}

func (m *TFindDialog) SetOnHelpClicked(fn TNotifyEvent) {
	if m.helpClickedPtr != 0 {
		RemoveEventElement(m.helpClickedPtr)
	}
	m.helpClickedPtr = MakeEventDataPtr(fn)
	findDialogImportAPI().SysCallN(8, m.Instance(), m.helpClickedPtr)
}

var (
	findDialogImport       *imports.Imports = nil
	findDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FindDialog_Class", 0),
		/*1*/ imports.NewTable("FindDialog_CloseDialog", 0),
		/*2*/ imports.NewTable("FindDialog_Create", 0),
		/*3*/ imports.NewTable("FindDialog_FindText", 0),
		/*4*/ imports.NewTable("FindDialog_Left", 0),
		/*5*/ imports.NewTable("FindDialog_Options", 0),
		/*6*/ imports.NewTable("FindDialog_Position", 0),
		/*7*/ imports.NewTable("FindDialog_SetOnFind", 0),
		/*8*/ imports.NewTable("FindDialog_SetOnHelpClicked", 0),
		/*9*/ imports.NewTable("FindDialog_Top", 0),
	}
)

func findDialogImportAPI() *imports.Imports {
	if findDialogImport == nil {
		findDialogImport = NewDefaultImports()
		findDialogImport.SetImportTable(findDialogImportTables)
		findDialogImportTables = nil
	}
	return findDialogImport
}
