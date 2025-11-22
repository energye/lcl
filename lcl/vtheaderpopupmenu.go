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

// IVTHeaderPopupMenu Parent: IPopupMenu
type IVTHeaderPopupMenu interface {
	IPopupMenu
	Options() types.TVTHeaderPopupOptions                // property Options Getter
	SetOptions(value types.TVTHeaderPopupOptions)        // property Options Setter
	SetOnAddHeaderPopupItem(fn TAddHeaderPopupItemEvent) // property event
	SetOnColumnChange(fn TColumnChangeEvent)             // property event
}

type TVTHeaderPopupMenu struct {
	TPopupMenu
}

func (m *TVTHeaderPopupMenu) Options() types.TVTHeaderPopupOptions {
	if !m.IsValid() {
		return 0
	}
	r := vTHeaderPopupMenuAPI().SysCallN(1, 0, m.Instance())
	return types.TVTHeaderPopupOptions(r)
}

func (m *TVTHeaderPopupMenu) SetOptions(value types.TVTHeaderPopupOptions) {
	if !m.IsValid() {
		return
	}
	vTHeaderPopupMenuAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TVTHeaderPopupMenu) SetOnAddHeaderPopupItem(fn TAddHeaderPopupItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTAddHeaderPopupItemEvent(fn)
	base.SetEvent(m, 2, vTHeaderPopupMenuAPI(), api.MakeEventDataPtr(cb))
}

func (m *TVTHeaderPopupMenu) SetOnColumnChange(fn TColumnChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTColumnChangeEvent(fn)
	base.SetEvent(m, 3, vTHeaderPopupMenuAPI(), api.MakeEventDataPtr(cb))
}

// NewVTHeaderPopupMenu class constructor
func NewVTHeaderPopupMenu(owner IComponent) IVTHeaderPopupMenu {
	r := vTHeaderPopupMenuAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsVTHeaderPopupMenu(r)
}

func TVTHeaderPopupMenuClass() types.TClass {
	r := vTHeaderPopupMenuAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	vTHeaderPopupMenuOnce   base.Once
	vTHeaderPopupMenuImport *imports.Imports = nil
)

func vTHeaderPopupMenuAPI() *imports.Imports {
	vTHeaderPopupMenuOnce.Do(func() {
		vTHeaderPopupMenuImport = api.NewDefaultImports()
		vTHeaderPopupMenuImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVTHeaderPopupMenu_Create", 0), // constructor NewVTHeaderPopupMenu
			/* 1 */ imports.NewTable("TVTHeaderPopupMenu_Options", 0), // property Options
			/* 2 */ imports.NewTable("TVTHeaderPopupMenu_OnAddHeaderPopupItem", 0), // event OnAddHeaderPopupItem
			/* 3 */ imports.NewTable("TVTHeaderPopupMenu_OnColumnChange", 0), // event OnColumnChange
			/* 4 */ imports.NewTable("TVTHeaderPopupMenu_TClass", 0), // function TVTHeaderPopupMenuClass
		}
	})
	return vTHeaderPopupMenuImport
}
