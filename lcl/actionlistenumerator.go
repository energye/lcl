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

// IActionListEnumerator Parent: IObject
type IActionListEnumerator interface {
	IObject
	Current() IContainedAction // property
	MoveNext() bool            // function
}

// TActionListEnumerator Parent: TObject
type TActionListEnumerator struct {
	TObject
}

func NewActionListEnumerator(AList ICustomActionList) IActionListEnumerator {
	r1 := actionListEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(AList))
	return AsActionListEnumerator(r1)
}

func (m *TActionListEnumerator) Current() IContainedAction {
	r1 := actionListEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsContainedAction(r1)
}

func (m *TActionListEnumerator) MoveNext() bool {
	r1 := actionListEnumeratorImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func ActionListEnumeratorClass() TClass {
	ret := actionListEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	actionListEnumeratorImport       *imports.Imports = nil
	actionListEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ActionListEnumerator_Class", 0),
		/*1*/ imports.NewTable("ActionListEnumerator_Create", 0),
		/*2*/ imports.NewTable("ActionListEnumerator_Current", 0),
		/*3*/ imports.NewTable("ActionListEnumerator_MoveNext", 0),
	}
)

func actionListEnumeratorImportAPI() *imports.Imports {
	if actionListEnumeratorImport == nil {
		actionListEnumeratorImport = NewDefaultImports()
		actionListEnumeratorImport.SetImportTable(actionListEnumeratorImportTables)
		actionListEnumeratorImportTables = nil
	}
	return actionListEnumeratorImport
}
