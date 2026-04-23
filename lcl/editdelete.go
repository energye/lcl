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

// IEditDelete Parent: IEditAction
type IEditDelete interface {
	IEditAction
}

type TEditDelete struct {
	TEditAction
}

// NewEditDelete class constructor
func NewEditDelete(owner IComponent) IEditDelete {
	r := editDeleteAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEditDelete(r)
}

func TEditDeleteClass() types.TClass {
	r := editDeleteAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	editDeleteOnce   base.Once
	editDeleteImport *imports.Imports = nil
)

func editDeleteAPI() *imports.Imports {
	editDeleteOnce.Do(func() {
		editDeleteImport = api.NewDefaultImports()
		editDeleteImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEditDelete_Create", 0), // constructor NewEditDelete
			/* 1 */ imports.NewTable("TEditDelete_TClass", 0), // function TEditDeleteClass
		}
	})
	return editDeleteImport
}
