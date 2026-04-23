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

// IHelpOnHelp Parent: IHelpAction
type IHelpOnHelp interface {
	IHelpAction
}

type THelpOnHelp struct {
	THelpAction
}

// NewHelpOnHelp class constructor
func NewHelpOnHelp(theOwner IComponent) IHelpOnHelp {
	r := helpOnHelpAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsHelpOnHelp(r)
}

func THelpOnHelpClass() types.TClass {
	r := helpOnHelpAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	helpOnHelpOnce   base.Once
	helpOnHelpImport *imports.Imports = nil
)

func helpOnHelpAPI() *imports.Imports {
	helpOnHelpOnce.Do(func() {
		helpOnHelpImport = api.NewDefaultImports()
		helpOnHelpImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("THelpOnHelp_Create", 0), // constructor NewHelpOnHelp
			/* 1 */ imports.NewTable("THelpOnHelp_TClass", 0), // function THelpOnHelpClass
		}
	})
	return helpOnHelpImport
}
