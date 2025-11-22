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

// ICustomEditButton Parent: ICustomAbstractGroupedEdit
type ICustomEditButton interface {
	ICustomAbstractGroupedEdit
}

type TCustomEditButton struct {
	TCustomAbstractGroupedEdit
}

// NewCustomEditButton class constructor
func NewCustomEditButton(owner IComponent) ICustomEditButton {
	r := customEditButtonAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomEditButton(r)
}

func TCustomEditButtonClass() types.TClass {
	r := customEditButtonAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	customEditButtonOnce   base.Once
	customEditButtonImport *imports.Imports = nil
)

func customEditButtonAPI() *imports.Imports {
	customEditButtonOnce.Do(func() {
		customEditButtonImport = api.NewDefaultImports()
		customEditButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomEditButton_Create", 0), // constructor NewCustomEditButton
			/* 1 */ imports.NewTable("TCustomEditButton_TClass", 0), // function TCustomEditButtonClass
		}
	})
	return customEditButtonImport
}
