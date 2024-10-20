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

// ITaskDialogBaseButtonItem Parent: ICollectionItem
type ITaskDialogBaseButtonItem interface {
	ICollectionItem
	ModalResult() TModalResult          // property
	SetModalResult(AValue TModalResult) // property
	Caption() string                    // property
	SetCaption(AValue string)           // property
	Default() bool                      // property
	SetDefault(AValue bool)             // property
}

// TTaskDialogBaseButtonItem Parent: TCollectionItem
type TTaskDialogBaseButtonItem struct {
	TCollectionItem
}

func NewTaskDialogBaseButtonItem(ACollection ICollection) ITaskDialogBaseButtonItem {
	r1 := askDialogBaseButtonItemImportAPI().SysCallN(2, GetObjectUintptr(ACollection))
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogBaseButtonItem) ModalResult() TModalResult {
	r1 := askDialogBaseButtonItemImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TTaskDialogBaseButtonItem) SetModalResult(AValue TModalResult) {
	askDialogBaseButtonItemImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TTaskDialogBaseButtonItem) Caption() string {
	r1 := askDialogBaseButtonItemImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TTaskDialogBaseButtonItem) SetCaption(AValue string) {
	askDialogBaseButtonItemImportAPI().SysCallN(0, 1, m.Instance(), PascalStr(AValue))
}

func (m *TTaskDialogBaseButtonItem) Default() bool {
	r1 := askDialogBaseButtonItemImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTaskDialogBaseButtonItem) SetDefault(AValue bool) {
	askDialogBaseButtonItemImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func TaskDialogBaseButtonItemClass() TClass {
	ret := askDialogBaseButtonItemImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	askDialogBaseButtonItemImport       *imports.Imports = nil
	askDialogBaseButtonItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TaskDialogBaseButtonItem_Caption", 0),
		/*1*/ imports.NewTable("TaskDialogBaseButtonItem_Class", 0),
		/*2*/ imports.NewTable("TaskDialogBaseButtonItem_Create", 0),
		/*3*/ imports.NewTable("TaskDialogBaseButtonItem_Default", 0),
		/*4*/ imports.NewTable("TaskDialogBaseButtonItem_ModalResult", 0),
	}
)

func askDialogBaseButtonItemImportAPI() *imports.Imports {
	if askDialogBaseButtonItemImport == nil {
		askDialogBaseButtonItemImport = NewDefaultImports()
		askDialogBaseButtonItemImport.SetImportTable(askDialogBaseButtonItemImportTables)
		askDialogBaseButtonItemImportTables = nil
	}
	return askDialogBaseButtonItemImport
}
