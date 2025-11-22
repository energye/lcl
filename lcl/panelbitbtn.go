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

// IPanelBitBtn Parent: ICustomBitBtn
type IPanelBitBtn interface {
	ICustomBitBtn
}

type TPanelBitBtn struct {
	TCustomBitBtn
}

// NewPanelBitBtn class constructor
func NewPanelBitBtn(owner IComponent) IPanelBitBtn {
	r := panelBitBtnAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsPanelBitBtn(r)
}

func TPanelBitBtnClass() types.TClass {
	r := panelBitBtnAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	panelBitBtnOnce   base.Once
	panelBitBtnImport *imports.Imports = nil
)

func panelBitBtnAPI() *imports.Imports {
	panelBitBtnOnce.Do(func() {
		panelBitBtnImport = api.NewDefaultImports()
		panelBitBtnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPanelBitBtn_Create", 0), // constructor NewPanelBitBtn
			/* 1 */ imports.NewTable("TPanelBitBtn_TClass", 0), // function TPanelBitBtnClass
		}
	})
	return panelBitBtnImport
}
