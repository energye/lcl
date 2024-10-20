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

// ICustomVirtualTreeOptions Parent: IPersistent
type ICustomVirtualTreeOptions interface {
	IPersistent
	Owner() IBaseVirtualTree   // property
	AssignTo(Dest IPersistent) // procedure
}

// TCustomVirtualTreeOptions Parent: TPersistent
type TCustomVirtualTreeOptions struct {
	TPersistent
}

func NewCustomVirtualTreeOptions(AOwner IBaseVirtualTree) ICustomVirtualTreeOptions {
	r1 := customVirtualTreeOptionsImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsCustomVirtualTreeOptions(r1)
}

func (m *TCustomVirtualTreeOptions) Owner() IBaseVirtualTree {
	r1 := customVirtualTreeOptionsImportAPI().SysCallN(3, m.Instance())
	return AsBaseVirtualTree(r1)
}

func CustomVirtualTreeOptionsClass() TClass {
	ret := customVirtualTreeOptionsImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCustomVirtualTreeOptions) AssignTo(Dest IPersistent) {
	customVirtualTreeOptionsImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(Dest))
}

var (
	customVirtualTreeOptionsImport       *imports.Imports = nil
	customVirtualTreeOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomVirtualTreeOptions_AssignTo", 0),
		/*1*/ imports.NewTable("CustomVirtualTreeOptions_Class", 0),
		/*2*/ imports.NewTable("CustomVirtualTreeOptions_Create", 0),
		/*3*/ imports.NewTable("CustomVirtualTreeOptions_Owner", 0),
	}
)

func customVirtualTreeOptionsImportAPI() *imports.Imports {
	if customVirtualTreeOptionsImport == nil {
		customVirtualTreeOptionsImport = NewDefaultImports()
		customVirtualTreeOptionsImport.SetImportTable(customVirtualTreeOptionsImportTables)
		customVirtualTreeOptionsImportTables = nil
	}
	return customVirtualTreeOptionsImport
}
