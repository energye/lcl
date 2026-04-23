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

// IEditCut Parent: IEditAction
type IEditCut interface {
	IEditAction
}

type TEditCut struct {
	TEditAction
}

// NewEditCut class constructor
func NewEditCut(owner IComponent) IEditCut {
	r := editCutAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEditCut(r)
}

func TEditCutClass() types.TClass {
	r := editCutAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	editCutOnce   base.Once
	editCutImport *imports.Imports = nil
)

func editCutAPI() *imports.Imports {
	editCutOnce.Do(func() {
		editCutImport = api.NewDefaultImports()
		editCutImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEditCut_Create", 0), // constructor NewEditCut
			/* 1 */ imports.NewTable("TEditCut_TClass", 0), // function TEditCutClass
		}
	})
	return editCutImport
}
