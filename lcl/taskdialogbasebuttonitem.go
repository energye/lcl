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

// ITaskDialogBaseButtonItem Parent: ICollectionItem
type ITaskDialogBaseButtonItem interface {
	ICollectionItem
	ModalResult() types.TModalResult         // property ModalResult Getter
	SetModalResult(value types.TModalResult) // property ModalResult Setter
	CommandLinkHint() string                 // property CommandLinkHint Getter
	SetCommandLinkHint(value string)         // property CommandLinkHint Setter
	Caption() string                         // property Caption Getter
	SetCaption(value string)                 // property Caption Setter
	Default() bool                           // property Default Getter
	SetDefault(value bool)                   // property Default Setter
}

type TTaskDialogBaseButtonItem struct {
	TCollectionItem
}

func (m *TTaskDialogBaseButtonItem) ModalResult() types.TModalResult {
	if !m.IsValid() {
		return 0
	}
	r := taskDialogBaseButtonItemAPI().SysCallN(1, 0, m.Instance())
	return types.TModalResult(r)
}

func (m *TTaskDialogBaseButtonItem) SetModalResult(value types.TModalResult) {
	if !m.IsValid() {
		return
	}
	taskDialogBaseButtonItemAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TTaskDialogBaseButtonItem) CommandLinkHint() string {
	if !m.IsValid() {
		return ""
	}
	r := taskDialogBaseButtonItemAPI().SysCallN(2, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTaskDialogBaseButtonItem) SetCommandLinkHint(value string) {
	if !m.IsValid() {
		return
	}
	taskDialogBaseButtonItemAPI().SysCallN(2, 1, m.Instance(), api.PasStr(value))
}

func (m *TTaskDialogBaseButtonItem) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := taskDialogBaseButtonItemAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTaskDialogBaseButtonItem) SetCaption(value string) {
	if !m.IsValid() {
		return
	}
	taskDialogBaseButtonItemAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TTaskDialogBaseButtonItem) Default() bool {
	if !m.IsValid() {
		return false
	}
	r := taskDialogBaseButtonItemAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTaskDialogBaseButtonItem) SetDefault(value bool) {
	if !m.IsValid() {
		return
	}
	taskDialogBaseButtonItemAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

// NewTaskDialogBaseButtonItem class constructor
func NewTaskDialogBaseButtonItem(collection ICollection) ITaskDialogBaseButtonItem {
	r := taskDialogBaseButtonItemAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsTaskDialogBaseButtonItem(r)
}

var (
	taskDialogBaseButtonItemOnce   base.Once
	taskDialogBaseButtonItemImport *imports.Imports = nil
)

func taskDialogBaseButtonItemAPI() *imports.Imports {
	taskDialogBaseButtonItemOnce.Do(func() {
		taskDialogBaseButtonItemImport = api.NewDefaultImports()
		taskDialogBaseButtonItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTaskDialogBaseButtonItem_Create", 0), // constructor NewTaskDialogBaseButtonItem
			/* 1 */ imports.NewTable("TTaskDialogBaseButtonItem_ModalResult", 0), // property ModalResult
			/* 2 */ imports.NewTable("TTaskDialogBaseButtonItem_CommandLinkHint", 0), // property CommandLinkHint
			/* 3 */ imports.NewTable("TTaskDialogBaseButtonItem_Caption", 0), // property Caption
			/* 4 */ imports.NewTable("TTaskDialogBaseButtonItem_Default", 0), // property Default
		}
	})
	return taskDialogBaseButtonItemImport
}
