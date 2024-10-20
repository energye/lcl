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

// IGEEdit Parent: ICustomMaskEdit
type IGEEdit interface {
	ICustomMaskEdit
}

// TGEEdit Parent: TCustomMaskEdit
type TGEEdit struct {
	TCustomMaskEdit
}

func NewGEEdit(TheOwner IComponent) IGEEdit {
	r1 := gEEditImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsGEEdit(r1)
}

func GEEditClass() TClass {
	ret := gEEditImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	gEEditImport       *imports.Imports = nil
	gEEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("GEEdit_Class", 0),
		/*1*/ imports.NewTable("GEEdit_Create", 0),
	}
)

func gEEditImportAPI() *imports.Imports {
	if gEEditImport == nil {
		gEEditImport = NewDefaultImports()
		gEEditImport.SetImportTable(gEEditImportTables)
		gEEditImportTables = nil
	}
	return gEEditImport
}
