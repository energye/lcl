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

// ISynPopupMenu Parent: IPopupMenu
type ISynPopupMenu interface {
	IPopupMenu
	DefaultPopupMenu() types.TSynDefaultPopupMenu         // property DefaultPopupMenu Getter
	SetDefaultPopupMenu(value types.TSynDefaultPopupMenu) // property DefaultPopupMenu Setter
}

type TSynPopupMenu struct {
	TPopupMenu
}

func (m *TSynPopupMenu) DefaultPopupMenu() types.TSynDefaultPopupMenu {
	if !m.IsValid() {
		return 0
	}
	r := synPopupMenuAPI().SysCallN(1, 0, m.Instance())
	return types.TSynDefaultPopupMenu(r)
}

func (m *TSynPopupMenu) SetDefaultPopupMenu(value types.TSynDefaultPopupMenu) {
	if !m.IsValid() {
		return
	}
	synPopupMenuAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

// NewSynPopupMenu class constructor
func NewSynPopupMenu(owner IComponent) ISynPopupMenu {
	r := synPopupMenuAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynPopupMenu(r)
}

func TSynPopupMenuClass() types.TClass {
	r := synPopupMenuAPI().SysCallN(2)
	return types.TClass(r)
}

var (
	synPopupMenuOnce   base.Once
	synPopupMenuImport *imports.Imports = nil
)

func synPopupMenuAPI() *imports.Imports {
	synPopupMenuOnce.Do(func() {
		synPopupMenuImport = api.NewDefaultImports()
		synPopupMenuImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynPopupMenu_Create", 0), // constructor NewSynPopupMenu
			/* 1 */ imports.NewTable("TSynPopupMenu_DefaultPopupMenu", 0), // property DefaultPopupMenu
			/* 2 */ imports.NewTable("TSynPopupMenu_TClass", 0), // function TSynPopupMenuClass
		}
	})
	return synPopupMenuImport
}
