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

// ITaskDialogButtons Parent: IOwnedCollection
type ITaskDialogButtons interface {
	IOwnedCollection
	AddToTaskDialogBaseButtonItem() ITaskDialogBaseButtonItem                               // function
	FindButton(modalResult types.TModalResult) ITaskDialogBaseButtonItem                    // function
	GetEnumeratorToTaskDialogButtonsEnumerator() ITaskDialogButtonsEnumerator               // function
	DefaultButton() ITaskDialogBaseButtonItem                                               // property DefaultButton Getter
	SetDefaultButton(value ITaskDialogBaseButtonItem)                                       // property DefaultButton Setter
	ItemsWithIntToTaskDialogBaseButtonItem(index int32) ITaskDialogBaseButtonItem           // property Items Getter
	SetItemsWithIntToTaskDialogBaseButtonItem(index int32, value ITaskDialogBaseButtonItem) // property Items Setter
}

type TTaskDialogButtons struct {
	TOwnedCollection
}

func (m *TTaskDialogButtons) AddToTaskDialogBaseButtonItem() ITaskDialogBaseButtonItem {
	if !m.IsValid() {
		return nil
	}
	r := taskDialogButtonsAPI().SysCallN(1, m.Instance())
	return AsTaskDialogBaseButtonItem(r)
}

func (m *TTaskDialogButtons) FindButton(modalResult types.TModalResult) ITaskDialogBaseButtonItem {
	if !m.IsValid() {
		return nil
	}
	r := taskDialogButtonsAPI().SysCallN(2, m.Instance(), uintptr(modalResult))
	return AsTaskDialogBaseButtonItem(r)
}

func (m *TTaskDialogButtons) GetEnumeratorToTaskDialogButtonsEnumerator() ITaskDialogButtonsEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := taskDialogButtonsAPI().SysCallN(3, m.Instance())
	return AsTaskDialogButtonsEnumerator(r)
}

func (m *TTaskDialogButtons) DefaultButton() ITaskDialogBaseButtonItem {
	if !m.IsValid() {
		return nil
	}
	r := taskDialogButtonsAPI().SysCallN(4, 0, m.Instance())
	return AsTaskDialogBaseButtonItem(r)
}

func (m *TTaskDialogButtons) SetDefaultButton(value ITaskDialogBaseButtonItem) {
	if !m.IsValid() {
		return
	}
	taskDialogButtonsAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TTaskDialogButtons) ItemsWithIntToTaskDialogBaseButtonItem(index int32) ITaskDialogBaseButtonItem {
	if !m.IsValid() {
		return nil
	}
	r := taskDialogButtonsAPI().SysCallN(5, 0, m.Instance(), uintptr(index))
	return AsTaskDialogBaseButtonItem(r)
}

func (m *TTaskDialogButtons) SetItemsWithIntToTaskDialogBaseButtonItem(index int32, value ITaskDialogBaseButtonItem) {
	if !m.IsValid() {
		return
	}
	taskDialogButtonsAPI().SysCallN(5, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

// NewTaskDialogButtons class constructor
func NewTaskDialogButtons(owner IPersistent, itemClass types.TCollectionItemClass) ITaskDialogButtons {
	r := taskDialogButtonsAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(itemClass))
	return AsTaskDialogButtons(r)
}

var (
	taskDialogButtonsOnce   base.Once
	taskDialogButtonsImport *imports.Imports = nil
)

func taskDialogButtonsAPI() *imports.Imports {
	taskDialogButtonsOnce.Do(func() {
		taskDialogButtonsImport = api.NewDefaultImports()
		taskDialogButtonsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTaskDialogButtons_Create", 0), // constructor NewTaskDialogButtons
			/* 1 */ imports.NewTable("TTaskDialogButtons_AddToTaskDialogBaseButtonItem", 0), // function AddToTaskDialogBaseButtonItem
			/* 2 */ imports.NewTable("TTaskDialogButtons_FindButton", 0), // function FindButton
			/* 3 */ imports.NewTable("TTaskDialogButtons_GetEnumeratorToTaskDialogButtonsEnumerator", 0), // function GetEnumeratorToTaskDialogButtonsEnumerator
			/* 4 */ imports.NewTable("TTaskDialogButtons_DefaultButton", 0), // property DefaultButton
			/* 5 */ imports.NewTable("TTaskDialogButtons_ItemsWithIntToTaskDialogBaseButtonItem", 0), // property ItemsWithIntToTaskDialogBaseButtonItem
		}
	})
	return taskDialogButtonsImport
}
