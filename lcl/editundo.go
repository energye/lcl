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

// IEditUndo Parent: IEditAction
type IEditUndo interface {
	IEditAction
}

type TEditUndo struct {
	TEditAction
}

// NewEditUndo class constructor
func NewEditUndo(owner IComponent) IEditUndo {
	r := editUndoAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEditUndo(r)
}

func TEditUndoClass() types.TClass {
	r := editUndoAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	editUndoOnce   base.Once
	editUndoImport *imports.Imports = nil
)

func editUndoAPI() *imports.Imports {
	editUndoOnce.Do(func() {
		editUndoImport = api.NewDefaultImports()
		editUndoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEditUndo_Create", 0), // constructor NewEditUndo
			/* 1 */ imports.NewTable("TEditUndo_TClass", 0), // function TEditUndoClass
		}
	})
	return editUndoImport
}
