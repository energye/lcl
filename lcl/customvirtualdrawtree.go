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

// ICustomVirtualDrawTree Parent: IBaseVirtualTree
// Tree descendant to let an application draw its stuff itself.
type ICustomVirtualDrawTree interface {
	IBaseVirtualTree
}

// TCustomVirtualDrawTree Parent: TBaseVirtualTree
// Tree descendant to let an application draw its stuff itself.
type TCustomVirtualDrawTree struct {
	TBaseVirtualTree
}

func NewCustomVirtualDrawTree(AOwner IComponent) ICustomVirtualDrawTree {
	r1 := customVirtualDrawTreeImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomVirtualDrawTree(r1)
}

func CustomVirtualDrawTreeClass() TClass {
	ret := customVirtualDrawTreeImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customVirtualDrawTreeImport       *imports.Imports = nil
	customVirtualDrawTreeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomVirtualDrawTree_Class", 0),
		/*1*/ imports.NewTable("CustomVirtualDrawTree_Create", 0),
	}
)

func customVirtualDrawTreeImportAPI() *imports.Imports {
	if customVirtualDrawTreeImport == nil {
		customVirtualDrawTreeImport = NewDefaultImports()
		customVirtualDrawTreeImport.SetImportTable(customVirtualDrawTreeImportTables)
		customVirtualDrawTreeImportTables = nil
	}
	return customVirtualDrawTreeImport
}
