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

// IHelpAction Parent: IAction
type IHelpAction interface {
	IAction
}

type THelpAction struct {
	TAction
}

// NewHelpAction class constructor
func NewHelpAction(theOwner IComponent) IHelpAction {
	r := helpActionAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsHelpAction(r)
}

func THelpActionClass() types.TClass {
	r := helpActionAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	helpActionOnce   base.Once
	helpActionImport *imports.Imports = nil
)

func helpActionAPI() *imports.Imports {
	helpActionOnce.Do(func() {
		helpActionImport = api.NewDefaultImports()
		helpActionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("THelpAction_Create", 0), // constructor NewHelpAction
			/* 1 */ imports.NewTable("THelpAction_TClass", 0), // function THelpActionClass
		}
	})
	return helpActionImport
}
