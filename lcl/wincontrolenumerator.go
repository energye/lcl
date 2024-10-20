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

// IWinControlEnumerator Parent: IObject
type IWinControlEnumerator interface {
	IObject
	Current() IControl                    // property
	GetEnumerator() IWinControlEnumerator // function
	MoveNext() bool                       // function
}

// TWinControlEnumerator Parent: TObject
type TWinControlEnumerator struct {
	TObject
}

func NewWinControlEnumerator(Parent IWinControl, aLowToHigh bool) IWinControlEnumerator {
	r1 := winControlEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(Parent), PascalBool(aLowToHigh))
	return AsWinControlEnumerator(r1)
}

func (m *TWinControlEnumerator) Current() IControl {
	r1 := winControlEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsControl(r1)
}

func (m *TWinControlEnumerator) GetEnumerator() IWinControlEnumerator {
	r1 := winControlEnumeratorImportAPI().SysCallN(3, m.Instance())
	return AsWinControlEnumerator(r1)
}

func (m *TWinControlEnumerator) MoveNext() bool {
	r1 := winControlEnumeratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func WinControlEnumeratorClass() TClass {
	ret := winControlEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	winControlEnumeratorImport       *imports.Imports = nil
	winControlEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("WinControlEnumerator_Class", 0),
		/*1*/ imports.NewTable("WinControlEnumerator_Create", 0),
		/*2*/ imports.NewTable("WinControlEnumerator_Current", 0),
		/*3*/ imports.NewTable("WinControlEnumerator_GetEnumerator", 0),
		/*4*/ imports.NewTable("WinControlEnumerator_MoveNext", 0),
	}
)

func winControlEnumeratorImportAPI() *imports.Imports {
	if winControlEnumeratorImport == nil {
		winControlEnumeratorImport = NewDefaultImports()
		winControlEnumeratorImport.SetImportTable(winControlEnumeratorImportTables)
		winControlEnumeratorImportTables = nil
	}
	return winControlEnumeratorImport
}
