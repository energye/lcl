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

// ITaskDialogButtonsEnumerator Parent: IObject
type ITaskDialogButtonsEnumerator interface {
	IObject
	Current() ITaskDialogBaseButtonItem    // property
	GetCurrent() ITaskDialogBaseButtonItem // function
	MoveNext() bool                        // function
}

// TTaskDialogButtonsEnumerator Parent: TObject
type TTaskDialogButtonsEnumerator struct {
	TObject
}

func NewTaskDialogButtonsEnumerator(ACollection ITaskDialogButtons) ITaskDialogButtonsEnumerator {
	r1 := askDialogButtonsEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(ACollection))
	return AsTaskDialogButtonsEnumerator(r1)
}

func (m *TTaskDialogButtonsEnumerator) Current() ITaskDialogBaseButtonItem {
	r1 := askDialogButtonsEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtonsEnumerator) GetCurrent() ITaskDialogBaseButtonItem {
	r1 := askDialogButtonsEnumeratorImportAPI().SysCallN(3, m.Instance())
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtonsEnumerator) MoveNext() bool {
	r1 := askDialogButtonsEnumeratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func TaskDialogButtonsEnumeratorClass() TClass {
	ret := askDialogButtonsEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	askDialogButtonsEnumeratorImport       *imports.Imports = nil
	askDialogButtonsEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TaskDialogButtonsEnumerator_Class", 0),
		/*1*/ imports.NewTable("TaskDialogButtonsEnumerator_Create", 0),
		/*2*/ imports.NewTable("TaskDialogButtonsEnumerator_Current", 0),
		/*3*/ imports.NewTable("TaskDialogButtonsEnumerator_GetCurrent", 0),
		/*4*/ imports.NewTable("TaskDialogButtonsEnumerator_MoveNext", 0),
	}
)

func askDialogButtonsEnumeratorImportAPI() *imports.Imports {
	if askDialogButtonsEnumeratorImport == nil {
		askDialogButtonsEnumeratorImport = NewDefaultImports()
		askDialogButtonsEnumeratorImport.SetImportTable(askDialogButtonsEnumeratorImportTables)
		askDialogButtonsEnumeratorImportTables = nil
	}
	return askDialogButtonsEnumeratorImport
}
