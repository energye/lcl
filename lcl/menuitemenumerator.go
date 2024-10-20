//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IMenuItemEnumerator Parent: IObject
type IMenuItemEnumerator interface {
	IObject
	Current() IMenuItem // property
	MoveNext() bool     // function
}

// TMenuItemEnumerator Parent: TObject
type TMenuItemEnumerator struct {
	TObject
}

func NewMenuItemEnumerator(AMenuItem IMenuItem) IMenuItemEnumerator {
	r1 := menuItemEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(AMenuItem))
	return AsMenuItemEnumerator(r1)
}

func (m *TMenuItemEnumerator) Current() IMenuItem {
	r1 := menuItemEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItemEnumerator) MoveNext() bool {
	r1 := menuItemEnumeratorImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func MenuItemEnumeratorClass() TClass {
	ret := menuItemEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	menuItemEnumeratorImport       *imports.Imports = nil
	menuItemEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("MenuItemEnumerator_Class", 0),
		/*1*/ imports.NewTable("MenuItemEnumerator_Create", 0),
		/*2*/ imports.NewTable("MenuItemEnumerator_Current", 0),
		/*3*/ imports.NewTable("MenuItemEnumerator_MoveNext", 0),
	}
)

func menuItemEnumeratorImportAPI() *imports.Imports {
	if menuItemEnumeratorImport == nil {
		menuItemEnumeratorImport = NewDefaultImports()
		menuItemEnumeratorImport.SetImportTable(menuItemEnumeratorImportTables)
		menuItemEnumeratorImportTables = nil
	}
	return menuItemEnumeratorImport
}
