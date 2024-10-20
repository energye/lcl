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

// ILazDockSplitter Parent: ICustomSplitter
type ILazDockSplitter interface {
	ICustomSplitter
}

// TLazDockSplitter Parent: TCustomSplitter
type TLazDockSplitter struct {
	TCustomSplitter
}

func NewLazDockSplitter(AOwner IComponent) ILazDockSplitter {
	r1 := lazDockSplitterImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsLazDockSplitter(r1)
}

func LazDockSplitterClass() TClass {
	ret := lazDockSplitterImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	lazDockSplitterImport       *imports.Imports = nil
	lazDockSplitterImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazDockSplitter_Class", 0),
		/*1*/ imports.NewTable("LazDockSplitter_Create", 0),
	}
)

func lazDockSplitterImportAPI() *imports.Imports {
	if lazDockSplitterImport == nil {
		lazDockSplitterImport = NewDefaultImports()
		lazDockSplitterImport.SetImportTable(lazDockSplitterImportTables)
		lazDockSplitterImportTables = nil
	}
	return lazDockSplitterImport
}
