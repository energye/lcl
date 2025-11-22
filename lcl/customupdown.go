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

// ICustomUpDown Parent: ICustomControl
type ICustomUpDown interface {
	ICustomControl
}

type TCustomUpDown struct {
	TCustomControl
}

// NewCustomUpDown class constructor
func NewCustomUpDown(owner IComponent) ICustomUpDown {
	r := customUpDownAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomUpDown(r)
}

func TCustomUpDownClass() types.TClass {
	r := customUpDownAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	customUpDownOnce   base.Once
	customUpDownImport *imports.Imports = nil
)

func customUpDownAPI() *imports.Imports {
	customUpDownOnce.Do(func() {
		customUpDownImport = api.NewDefaultImports()
		customUpDownImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomUpDown_Create", 0), // constructor NewCustomUpDown
			/* 1 */ imports.NewTable("TCustomUpDown_TClass", 0), // function TCustomUpDownClass
		}
	})
	return customUpDownImport
}
