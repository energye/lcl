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

// ILCLComponent Parent: IComponent
type ILCLComponent interface {
	IComponent
}

type TLCLComponent struct {
	TComponent
}

// NewLCLComponent class constructor
func NewLCLComponent(owner IComponent) ILCLComponent {
	r := lCLComponentAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLCLComponent(r)
}

func TLCLComponentClass() types.TClass {
	r := lCLComponentAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	lCLComponentOnce   base.Once
	lCLComponentImport *imports.Imports = nil
)

func lCLComponentAPI() *imports.Imports {
	lCLComponentOnce.Do(func() {
		lCLComponentImport = api.NewDefaultImports()
		lCLComponentImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLCLComponent_Create", 0), // constructor NewLCLComponent
			/* 1 */ imports.NewTable("TLCLComponent_TClass", 0), // function TLCLComponentClass
		}
	})
	return lCLComponentImport
}
