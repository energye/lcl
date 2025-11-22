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

// IAction Parent: ICustomAction
type IAction interface {
	ICustomAction
}

type TAction struct {
	TCustomAction
}

// NewAction class constructor
func NewAction(owner IComponent) IAction {
	r := actionAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsAction(r)
}

func TActionClass() types.TClass {
	r := actionAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	actionOnce   base.Once
	actionImport *imports.Imports = nil
)

func actionAPI() *imports.Imports {
	actionOnce.Do(func() {
		actionImport = api.NewDefaultImports()
		actionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TAction_Create", 0), // constructor NewAction
			/* 1 */ imports.NewTable("TAction_TClass", 0), // function TActionClass
		}
	})
	return actionImport
}
