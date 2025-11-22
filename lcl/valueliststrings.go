//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
)

// IValueListStrings Parent: IStringList
type IValueListStrings interface {
	IStringList
}

type TValueListStrings struct {
	TStringList
}

// NewValueListStrings class constructor
func NewValueListStrings(owner IValueListEditor) IValueListStrings {
	r := valueListStringsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsValueListStrings(r)
}

var (
	valueListStringsOnce   base.Once
	valueListStringsImport *imports.Imports = nil
)

func valueListStringsAPI() *imports.Imports {
	valueListStringsOnce.Do(func() {
		valueListStringsImport = api.NewDefaultImports()
		valueListStringsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TValueListStrings_Create", 0), // constructor NewValueListStrings
		}
	})
	return valueListStringsImport
}
