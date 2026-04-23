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

// IEditPaste Parent: IEditAction
type IEditPaste interface {
	IEditAction
}

type TEditPaste struct {
	TEditAction
}

// NewEditPaste class constructor
func NewEditPaste(owner IComponent) IEditPaste {
	r := editPasteAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEditPaste(r)
}

func TEditPasteClass() types.TClass {
	r := editPasteAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	editPasteOnce   base.Once
	editPasteImport *imports.Imports = nil
)

func editPasteAPI() *imports.Imports {
	editPasteOnce.Do(func() {
		editPasteImport = api.NewDefaultImports()
		editPasteImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEditPaste_Create", 0), // constructor NewEditPaste
			/* 1 */ imports.NewTable("TEditPaste_TClass", 0), // function TEditPasteClass
		}
	})
	return editPasteImport
}
