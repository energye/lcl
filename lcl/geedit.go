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

// IGEEdit Parent: ICustomMaskEdit
type IGEEdit interface {
	ICustomMaskEdit
}

type TGEEdit struct {
	TCustomMaskEdit
}

// NewGEEdit class constructor
func NewGEEdit(theOwner IComponent) IGEEdit {
	r := gEEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsGEEdit(r)
}

func TGEEditClass() types.TClass {
	r := gEEditAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	gEEditOnce   base.Once
	gEEditImport *imports.Imports = nil
)

func gEEditAPI() *imports.Imports {
	gEEditOnce.Do(func() {
		gEEditImport = api.NewDefaultImports()
		gEEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGEEdit_Create", 0), // constructor NewGEEdit
			/* 1 */ imports.NewTable("TGEEdit_TClass", 0), // function TGEEditClass
		}
	})
	return gEEditImport
}
