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

// ITaskDialogButtons Parent: IOwnedCollection
type ITaskDialogButtons interface {
	IOwnedCollection
	DefaultButton() ITaskDialogBaseButtonItem                                          // property
	SetDefaultButton(AValue ITaskDialogBaseButtonItem)                                 // property
	ItemsForTaskDialogBaseButtonItem(Index int32) ITaskDialogBaseButtonItem            // property
	SetItemsForTaskDialogBaseButtonItem(Index int32, AValue ITaskDialogBaseButtonItem) // property
	AddForTaskDialogBaseButtonItem() ITaskDialogBaseButtonItem                         // function
	FindButton(AModalResult TModalResult) ITaskDialogBaseButtonItem                    // function
	GetEnumeratorForTaskDialogButtonsEnumerator() ITaskDialogButtonsEnumerator         // function
}

// TTaskDialogButtons Parent: TOwnedCollection
type TTaskDialogButtons struct {
	TOwnedCollection
}

func NewTaskDialogButtons(AOwner IPersistent, AItemClass TCollectionItemClass) ITaskDialogButtons {
	r1 := askDialogButtonsImportAPI().SysCallN(2, GetObjectUintptr(AOwner), uintptr(AItemClass))
	return AsTaskDialogButtons(r1)
}

func (m *TTaskDialogButtons) DefaultButton() ITaskDialogBaseButtonItem {
	r1 := askDialogButtonsImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtons) SetDefaultButton(AValue ITaskDialogBaseButtonItem) {
	askDialogButtonsImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TTaskDialogButtons) ItemsForTaskDialogBaseButtonItem(Index int32) ITaskDialogBaseButtonItem {
	r1 := askDialogButtonsImportAPI().SysCallN(6, 0, m.Instance(), uintptr(Index))
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtons) SetItemsForTaskDialogBaseButtonItem(Index int32, AValue ITaskDialogBaseButtonItem) {
	askDialogButtonsImportAPI().SysCallN(6, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TTaskDialogButtons) AddForTaskDialogBaseButtonItem() ITaskDialogBaseButtonItem {
	r1 := askDialogButtonsImportAPI().SysCallN(0, m.Instance())
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtons) FindButton(AModalResult TModalResult) ITaskDialogBaseButtonItem {
	r1 := askDialogButtonsImportAPI().SysCallN(4, m.Instance(), uintptr(AModalResult))
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtons) GetEnumeratorForTaskDialogButtonsEnumerator() ITaskDialogButtonsEnumerator {
	r1 := askDialogButtonsImportAPI().SysCallN(5, m.Instance())
	return AsTaskDialogButtonsEnumerator(r1)
}

func TaskDialogButtonsClass() TClass {
	ret := askDialogButtonsImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	askDialogButtonsImport       *imports.Imports = nil
	askDialogButtonsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TaskDialogButtons_AddForTaskDialogBaseButtonItem", 0),
		/*1*/ imports.NewTable("TaskDialogButtons_Class", 0),
		/*2*/ imports.NewTable("TaskDialogButtons_Create", 0),
		/*3*/ imports.NewTable("TaskDialogButtons_DefaultButton", 0),
		/*4*/ imports.NewTable("TaskDialogButtons_FindButton", 0),
		/*5*/ imports.NewTable("TaskDialogButtons_GetEnumeratorForTaskDialogButtonsEnumerator", 0),
		/*6*/ imports.NewTable("TaskDialogButtons_ItemsForTaskDialogBaseButtonItem", 0),
	}
)

func askDialogButtonsImportAPI() *imports.Imports {
	if askDialogButtonsImport == nil {
		askDialogButtonsImport = NewDefaultImports()
		askDialogButtonsImport.SetImportTable(askDialogButtonsImportTables)
		askDialogButtonsImportTables = nil
	}
	return askDialogButtonsImport
}
