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

// IListFilterEdit Parent: ICustomControlFilterEdit
type IListFilterEdit interface {
	ICustomControlFilterEdit
	RemoveItem(item string)                     // procedure
	ItemWasClicked(item string, isChecked bool) // procedure
	SimpleSelection() bool                      // property SimpleSelection Getter
	SetSimpleSelection(value bool)              // property SimpleSelection Setter
	SelectionList() IStringList                 // property SelectionList Getter
	Items() IStringList                         // property Items Getter
	FilteredListbox() ICustomListBox            // property FilteredListbox Getter
	SetFilteredListbox(value ICustomListBox)    // property FilteredListbox Setter
}

type TListFilterEdit struct {
	TCustomControlFilterEdit
}

func (m *TListFilterEdit) RemoveItem(item string) {
	if !m.IsValid() {
		return
	}
	listFilterEditAPI().SysCallN(1, m.Instance(), api.PasStr(item))
}

func (m *TListFilterEdit) ItemWasClicked(item string, isChecked bool) {
	if !m.IsValid() {
		return
	}
	listFilterEditAPI().SysCallN(2, m.Instance(), api.PasStr(item), api.PasBool(isChecked))
}

func (m *TListFilterEdit) SimpleSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := listFilterEditAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListFilterEdit) SetSimpleSelection(value bool) {
	if !m.IsValid() {
		return
	}
	listFilterEditAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TListFilterEdit) SelectionList() IStringList {
	if !m.IsValid() {
		return nil
	}
	r := listFilterEditAPI().SysCallN(4, m.Instance())
	return AsStringList(r)
}

func (m *TListFilterEdit) Items() IStringList {
	if !m.IsValid() {
		return nil
	}
	r := listFilterEditAPI().SysCallN(5, m.Instance())
	return AsStringList(r)
}

func (m *TListFilterEdit) FilteredListbox() ICustomListBox {
	if !m.IsValid() {
		return nil
	}
	r := listFilterEditAPI().SysCallN(6, 0, m.Instance())
	return AsCustomListBox(r)
}

func (m *TListFilterEdit) SetFilteredListbox(value ICustomListBox) {
	if !m.IsValid() {
		return
	}
	listFilterEditAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewListFilterEdit class constructor
func NewListFilterEdit(owner IComponent) IListFilterEdit {
	r := listFilterEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsListFilterEdit(r)
}

func TListFilterEditClass() types.TClass {
	r := listFilterEditAPI().SysCallN(7)
	return types.TClass(r)
}

var (
	listFilterEditOnce   base.Once
	listFilterEditImport *imports.Imports = nil
)

func listFilterEditAPI() *imports.Imports {
	listFilterEditOnce.Do(func() {
		listFilterEditImport = api.NewDefaultImports()
		listFilterEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListFilterEdit_Create", 0), // constructor NewListFilterEdit
			/* 1 */ imports.NewTable("TListFilterEdit_RemoveItem", 0), // procedure RemoveItem
			/* 2 */ imports.NewTable("TListFilterEdit_ItemWasClicked", 0), // procedure ItemWasClicked
			/* 3 */ imports.NewTable("TListFilterEdit_SimpleSelection", 0), // property SimpleSelection
			/* 4 */ imports.NewTable("TListFilterEdit_SelectionList", 0), // property SelectionList
			/* 5 */ imports.NewTable("TListFilterEdit_Items", 0), // property Items
			/* 6 */ imports.NewTable("TListFilterEdit_FilteredListbox", 0), // property FilteredListbox
			/* 7 */ imports.NewTable("TListFilterEdit_TClass", 0), // function TListFilterEditClass
		}
	})
	return listFilterEditImport
}
