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
	"github.com/energye/lcl/types"
)

// IEditSelectAll Parent: IEditAction
type IEditSelectAll interface {
	IEditAction
}

type TEditSelectAll struct {
	TEditAction
}

// NewEditSelectAll class constructor
func NewEditSelectAll(owner IComponent) IEditSelectAll {
	r := editSelectAllAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEditSelectAll(r)
}

func TEditSelectAllClass() types.TClass {
	r := editSelectAllAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	editSelectAllOnce   base.Once
	editSelectAllImport *imports.Imports = nil
)

func editSelectAllAPI() *imports.Imports {
	editSelectAllOnce.Do(func() {
		editSelectAllImport = api.NewDefaultImports()
		editSelectAllImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEditSelectAll_Create", 0), // constructor NewEditSelectAll
			/* 1 */ imports.NewTable("TEditSelectAll_TClass", 0), // function TEditSelectAllClass
		}
	})
	return editSelectAllImport
}
