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

// IEditCopy Parent: IEditAction
type IEditCopy interface {
	IEditAction
}

type TEditCopy struct {
	TEditAction
}

// NewEditCopy class constructor
func NewEditCopy(owner IComponent) IEditCopy {
	r := editCopyAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEditCopy(r)
}

func TEditCopyClass() types.TClass {
	r := editCopyAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	editCopyOnce   base.Once
	editCopyImport *imports.Imports = nil
)

func editCopyAPI() *imports.Imports {
	editCopyOnce.Do(func() {
		editCopyImport = api.NewDefaultImports()
		editCopyImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEditCopy_Create", 0), // constructor NewEditCopy
			/* 1 */ imports.NewTable("TEditCopy_TClass", 0), // function TEditCopyClass
		}
	})
	return editCopyImport
}
