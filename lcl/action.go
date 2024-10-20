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

// IAction Parent: ICustomAction
type IAction interface {
	ICustomAction
}

// TAction Parent: TCustomAction
type TAction struct {
	TCustomAction
}

func NewAction(AOwner IComponent) IAction {
	r1 := actionImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsAction(r1)
}

func ActionClass() TClass {
	ret := actionImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	actionImport       *imports.Imports = nil
	actionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Action_Class", 0),
		/*1*/ imports.NewTable("Action_Create", 0),
	}
)

func actionImportAPI() *imports.Imports {
	if actionImport == nil {
		actionImport = NewDefaultImports()
		actionImport.SetImportTable(actionImportTables)
		actionImportTables = nil
	}
	return actionImport
}
