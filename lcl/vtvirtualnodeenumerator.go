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

// IVTVirtualNodeEnumerator Parent: IObject
type IVTVirtualNodeEnumerator interface {
	IObject
	Current() IVirtualNode // property
	MoveNext() bool        // function
}

// TVTVirtualNodeEnumerator Parent: TObject
type TVTVirtualNodeEnumerator struct {
	TObject
}

func NewVTVirtualNodeEnumerator() IVTVirtualNodeEnumerator {
	r1 := vTVirtualNodeEnumeratorImportAPI().SysCallN(1)
	return AsVTVirtualNodeEnumerator(r1)
}

func (m *TVTVirtualNodeEnumerator) Current() IVirtualNode {
	r1 := vTVirtualNodeEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsVirtualNode(r1)
}

func (m *TVTVirtualNodeEnumerator) MoveNext() bool {
	r1 := vTVirtualNodeEnumeratorImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func VTVirtualNodeEnumeratorClass() TClass {
	ret := vTVirtualNodeEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	vTVirtualNodeEnumeratorImport       *imports.Imports = nil
	vTVirtualNodeEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("VTVirtualNodeEnumerator_Class", 0),
		/*1*/ imports.NewTable("VTVirtualNodeEnumerator_Create", 0),
		/*2*/ imports.NewTable("VTVirtualNodeEnumerator_Current", 0),
		/*3*/ imports.NewTable("VTVirtualNodeEnumerator_MoveNext", 0),
	}
)

func vTVirtualNodeEnumeratorImportAPI() *imports.Imports {
	if vTVirtualNodeEnumeratorImport == nil {
		vTVirtualNodeEnumeratorImport = NewDefaultImports()
		vTVirtualNodeEnumeratorImport.SetImportTable(vTVirtualNodeEnumeratorImportTables)
		vTVirtualNodeEnumeratorImportTables = nil
	}
	return vTVirtualNodeEnumeratorImport
}
