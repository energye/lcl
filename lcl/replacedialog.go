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

// IReplaceDialog Parent: IFindDialog
type IReplaceDialog interface {
	IFindDialog
	ReplaceText() string          // property
	SetReplaceText(AValue string) // property
	SetOnReplace(fn TNotifyEvent) // property event
}

// TReplaceDialog Parent: TFindDialog
type TReplaceDialog struct {
	TFindDialog
	replacePtr uintptr
}

func NewReplaceDialog(AOwner IComponent) IReplaceDialog {
	r1 := replaceDialogImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsReplaceDialog(r1)
}

func (m *TReplaceDialog) ReplaceText() string {
	r1 := replaceDialogImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TReplaceDialog) SetReplaceText(AValue string) {
	replaceDialogImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(AValue))
}

func ReplaceDialogClass() TClass {
	ret := replaceDialogImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TReplaceDialog) SetOnReplace(fn TNotifyEvent) {
	if m.replacePtr != 0 {
		RemoveEventElement(m.replacePtr)
	}
	m.replacePtr = MakeEventDataPtr(fn)
	replaceDialogImportAPI().SysCallN(3, m.Instance(), m.replacePtr)
}

var (
	replaceDialogImport       *imports.Imports = nil
	replaceDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ReplaceDialog_Class", 0),
		/*1*/ imports.NewTable("ReplaceDialog_Create", 0),
		/*2*/ imports.NewTable("ReplaceDialog_ReplaceText", 0),
		/*3*/ imports.NewTable("ReplaceDialog_SetOnReplace", 0),
	}
)

func replaceDialogImportAPI() *imports.Imports {
	if replaceDialogImport == nil {
		replaceDialogImport = NewDefaultImports()
		replaceDialogImport.SetImportTable(replaceDialogImportTables)
		replaceDialogImportTables = nil
	}
	return replaceDialogImport
}
