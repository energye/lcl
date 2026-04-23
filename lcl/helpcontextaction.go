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

// IHelpContextAction Parent: IHelpAction
type IHelpContextAction interface {
	IHelpAction
}

type THelpContextAction struct {
	THelpAction
}

// NewHelpContextAction class constructor
func NewHelpContextAction(theOwner IComponent) IHelpContextAction {
	r := helpContextActionAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsHelpContextAction(r)
}

func THelpContextActionClass() types.TClass {
	r := helpContextActionAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	helpContextActionOnce   base.Once
	helpContextActionImport *imports.Imports = nil
)

func helpContextActionAPI() *imports.Imports {
	helpContextActionOnce.Do(func() {
		helpContextActionImport = api.NewDefaultImports()
		helpContextActionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("THelpContextAction_Create", 0), // constructor NewHelpContextAction
			/* 1 */ imports.NewTable("THelpContextAction_TClass", 0), // function THelpContextActionClass
		}
	})
	return helpContextActionImport
}
