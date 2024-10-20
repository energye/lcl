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

// IFPListEnumerator Parent: IObject
type IFPListEnumerator interface {
	IObject
	Current() uintptr    // property
	GetCurrent() uintptr // function
	MoveNext() bool      // function
}

// TFPListEnumerator Parent: TObject
type TFPListEnumerator struct {
	TObject
}

func NewFPListEnumerator(AList IFPList) IFPListEnumerator {
	r1 := fPListEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(AList))
	return AsFPListEnumerator(r1)
}

func (m *TFPListEnumerator) Current() uintptr {
	r1 := fPListEnumeratorImportAPI().SysCallN(2, m.Instance())
	return uintptr(r1)
}

func (m *TFPListEnumerator) GetCurrent() uintptr {
	r1 := fPListEnumeratorImportAPI().SysCallN(3, m.Instance())
	return uintptr(r1)
}

func (m *TFPListEnumerator) MoveNext() bool {
	r1 := fPListEnumeratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func FPListEnumeratorClass() TClass {
	ret := fPListEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	fPListEnumeratorImport       *imports.Imports = nil
	fPListEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPListEnumerator_Class", 0),
		/*1*/ imports.NewTable("FPListEnumerator_Create", 0),
		/*2*/ imports.NewTable("FPListEnumerator_Current", 0),
		/*3*/ imports.NewTable("FPListEnumerator_GetCurrent", 0),
		/*4*/ imports.NewTable("FPListEnumerator_MoveNext", 0),
	}
)

func fPListEnumeratorImportAPI() *imports.Imports {
	if fPListEnumeratorImport == nil {
		fPListEnumeratorImport = NewDefaultImports()
		fPListEnumeratorImport.SetImportTable(fPListEnumeratorImportTables)
		fPListEnumeratorImportTables = nil
	}
	return fPListEnumeratorImport
}
