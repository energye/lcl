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

// IToolBarEnumerator Parent: IObject
type IToolBarEnumerator interface {
	IObject
	Current() IToolButton // property
	MoveNext() bool       // function
}

// TToolBarEnumerator Parent: TObject
type TToolBarEnumerator struct {
	TObject
}

func NewToolBarEnumerator(AToolBar IToolBar) IToolBarEnumerator {
	r1 := oolBarEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(AToolBar))
	return AsToolBarEnumerator(r1)
}

func (m *TToolBarEnumerator) Current() IToolButton {
	r1 := oolBarEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsToolButton(r1)
}

func (m *TToolBarEnumerator) MoveNext() bool {
	r1 := oolBarEnumeratorImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func ToolBarEnumeratorClass() TClass {
	ret := oolBarEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	oolBarEnumeratorImport       *imports.Imports = nil
	oolBarEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ToolBarEnumerator_Class", 0),
		/*1*/ imports.NewTable("ToolBarEnumerator_Create", 0),
		/*2*/ imports.NewTable("ToolBarEnumerator_Current", 0),
		/*3*/ imports.NewTable("ToolBarEnumerator_MoveNext", 0),
	}
)

func oolBarEnumeratorImportAPI() *imports.Imports {
	if oolBarEnumeratorImport == nil {
		oolBarEnumeratorImport = NewDefaultImports()
		oolBarEnumeratorImport.SetImportTable(oolBarEnumeratorImportTables)
		oolBarEnumeratorImportTables = nil
	}
	return oolBarEnumeratorImport
}
