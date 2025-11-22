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

// IButtonControl Parent: IWinControl
type IButtonControl interface {
	IWinControl
}

type TButtonControl struct {
	TWinControl
}

// NewButtonControl class constructor
func NewButtonControl(theOwner IComponent) IButtonControl {
	r := buttonControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsButtonControl(r)
}

func TButtonControlClass() types.TClass {
	r := buttonControlAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	buttonControlOnce   base.Once
	buttonControlImport *imports.Imports = nil
)

func buttonControlAPI() *imports.Imports {
	buttonControlOnce.Do(func() {
		buttonControlImport = api.NewDefaultImports()
		buttonControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TButtonControl_Create", 0), // constructor NewButtonControl
			/* 1 */ imports.NewTable("TButtonControl_TClass", 0), // function TButtonControlClass
		}
	})
	return buttonControlImport
}
