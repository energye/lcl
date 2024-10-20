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

// ICustomStringTreeOptions Parent: ICustomVirtualTreeOptions
type ICustomStringTreeOptions interface {
	ICustomVirtualTreeOptions
}

// TCustomStringTreeOptions Parent: TCustomVirtualTreeOptions
type TCustomStringTreeOptions struct {
	TCustomVirtualTreeOptions
}

func NewCustomStringTreeOptions(AOwner IBaseVirtualTree) ICustomStringTreeOptions {
	r1 := customStringTreeOptionsImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomStringTreeOptions(r1)
}

func CustomStringTreeOptionsClass() TClass {
	ret := customStringTreeOptionsImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customStringTreeOptionsImport       *imports.Imports = nil
	customStringTreeOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomStringTreeOptions_Class", 0),
		/*1*/ imports.NewTable("CustomStringTreeOptions_Create", 0),
	}
)

func customStringTreeOptionsImportAPI() *imports.Imports {
	if customStringTreeOptionsImport == nil {
		customStringTreeOptionsImport = NewDefaultImports()
		customStringTreeOptionsImport.SetImportTable(customStringTreeOptionsImportTables)
		customStringTreeOptionsImportTables = nil
	}
	return customStringTreeOptionsImport
}
