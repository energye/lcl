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

// IValueListStrings Parent: IStringList
type IValueListStrings interface {
	IStringList
}

// TValueListStrings Parent: TStringList
type TValueListStrings struct {
	TStringList
}

func NewValueListStrings(AOwner IValueListEditor) IValueListStrings {
	r1 := valueListStringsImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsValueListStrings(r1)
}

func ValueListStringsClass() TClass {
	ret := valueListStringsImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	valueListStringsImport       *imports.Imports = nil
	valueListStringsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ValueListStrings_Class", 0),
		/*1*/ imports.NewTable("ValueListStrings_Create", 0),
	}
)

func valueListStringsImportAPI() *imports.Imports {
	if valueListStringsImport == nil {
		valueListStringsImport = NewDefaultImports()
		valueListStringsImport.SetImportTable(valueListStringsImportTables)
		valueListStringsImportTables = nil
	}
	return valueListStringsImport
}
