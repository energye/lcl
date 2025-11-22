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
)

// ITaskDialogButtonsEnumerator Parent: IObject
type ITaskDialogButtonsEnumerator interface {
	IObject
	GetCurrent() ITaskDialogBaseButtonItem // function
	MoveNext() bool                        // function
	Current() ITaskDialogBaseButtonItem    // property Current Getter
}

type TTaskDialogButtonsEnumerator struct {
	TObject
}

func (m *TTaskDialogButtonsEnumerator) GetCurrent() ITaskDialogBaseButtonItem {
	if !m.IsValid() {
		return nil
	}
	r := taskDialogButtonsEnumeratorAPI().SysCallN(1, m.Instance())
	return AsTaskDialogBaseButtonItem(r)
}

func (m *TTaskDialogButtonsEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := taskDialogButtonsEnumeratorAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TTaskDialogButtonsEnumerator) Current() ITaskDialogBaseButtonItem {
	if !m.IsValid() {
		return nil
	}
	r := taskDialogButtonsEnumeratorAPI().SysCallN(3, m.Instance())
	return AsTaskDialogBaseButtonItem(r)
}

// NewTaskDialogButtonsEnumerator class constructor
func NewTaskDialogButtonsEnumerator(collection ITaskDialogButtons) ITaskDialogButtonsEnumerator {
	r := taskDialogButtonsEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsTaskDialogButtonsEnumerator(r)
}

var (
	taskDialogButtonsEnumeratorOnce   base.Once
	taskDialogButtonsEnumeratorImport *imports.Imports = nil
)

func taskDialogButtonsEnumeratorAPI() *imports.Imports {
	taskDialogButtonsEnumeratorOnce.Do(func() {
		taskDialogButtonsEnumeratorImport = api.NewDefaultImports()
		taskDialogButtonsEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTaskDialogButtonsEnumerator_Create", 0), // constructor NewTaskDialogButtonsEnumerator
			/* 1 */ imports.NewTable("TTaskDialogButtonsEnumerator_GetCurrent", 0), // function GetCurrent
			/* 2 */ imports.NewTable("TTaskDialogButtonsEnumerator_MoveNext", 0), // function MoveNext
			/* 3 */ imports.NewTable("TTaskDialogButtonsEnumerator_Current", 0), // property Current
		}
	})
	return taskDialogButtonsEnumeratorImport
}
