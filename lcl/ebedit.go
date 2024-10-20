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

// IEbEdit Parent: IGEEdit
type IEbEdit interface {
	IGEEdit
}

// TEbEdit Parent: TGEEdit
type TEbEdit struct {
	TGEEdit
}

func NewEbEdit(TheOwner IComponent) IEbEdit {
	r1 := ebEditImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsEbEdit(r1)
}

func EbEditClass() TClass {
	ret := ebEditImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	ebEditImport       *imports.Imports = nil
	ebEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("EbEdit_Class", 0),
		/*1*/ imports.NewTable("EbEdit_Create", 0),
	}
)

func ebEditImportAPI() *imports.Imports {
	if ebEditImport == nil {
		ebEditImport = NewDefaultImports()
		ebEditImport.SetImportTable(ebEditImportTables)
		ebEditImportTables = nil
	}
	return ebEditImport
}
