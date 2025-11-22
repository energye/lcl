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
)

// IMenuItemEnumerator Parent: IObject
type IMenuItemEnumerator interface {
	IObject
	MoveNext() bool     // function
	Current() IMenuItem // property Current Getter
}

type TMenuItemEnumerator struct {
	TObject
}

func (m *TMenuItemEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := menuItemEnumeratorAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TMenuItemEnumerator) Current() IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := menuItemEnumeratorAPI().SysCallN(2, m.Instance())
	return AsMenuItem(r)
}

// NewMenuItemEnumerator class constructor
func NewMenuItemEnumerator(menuItem IMenuItem) IMenuItemEnumerator {
	r := menuItemEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(menuItem))
	return AsMenuItemEnumerator(r)
}

var (
	menuItemEnumeratorOnce   base.Once
	menuItemEnumeratorImport *imports.Imports = nil
)

func menuItemEnumeratorAPI() *imports.Imports {
	menuItemEnumeratorOnce.Do(func() {
		menuItemEnumeratorImport = api.NewDefaultImports()
		menuItemEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMenuItemEnumerator_Create", 0), // constructor NewMenuItemEnumerator
			/* 1 */ imports.NewTable("TMenuItemEnumerator_MoveNext", 0), // function MoveNext
			/* 2 */ imports.NewTable("TMenuItemEnumerator_Current", 0), // property Current
		}
	})
	return menuItemEnumeratorImport
}
