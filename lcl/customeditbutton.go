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

// ICustomEditButton Parent: ICustomAbstractGroupedEdit
type ICustomEditButton interface {
	ICustomAbstractGroupedEdit
}

// TCustomEditButton Parent: TCustomAbstractGroupedEdit
type TCustomEditButton struct {
	TCustomAbstractGroupedEdit
}

func NewCustomEditButton(AOwner IComponent) ICustomEditButton {
	r1 := customEditButtonImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomEditButton(r1)
}

func CustomEditButtonClass() TClass {
	ret := customEditButtonImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customEditButtonImport       *imports.Imports = nil
	customEditButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomEditButton_Class", 0),
		/*1*/ imports.NewTable("CustomEditButton_Create", 0),
	}
)

func customEditButtonImportAPI() *imports.Imports {
	if customEditButtonImport == nil {
		customEditButtonImport = NewDefaultImports()
		customEditButtonImport.SetImportTable(customEditButtonImportTables)
		customEditButtonImportTables = nil
	}
	return customEditButtonImport
}
