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

// IReplaceDialog Parent: IFindDialog
type IReplaceDialog interface {
	IFindDialog
	ReplaceText() string          // property ReplaceText Getter
	SetReplaceText(value string)  // property ReplaceText Setter
	SetOnReplace(fn TNotifyEvent) // property event
}

type TReplaceDialog struct {
	TFindDialog
}

func (m *TReplaceDialog) ReplaceText() string {
	if !m.IsValid() {
		return ""
	}
	r := replaceDialogAPI().SysCallN(1, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TReplaceDialog) SetReplaceText(value string) {
	if !m.IsValid() {
		return
	}
	replaceDialogAPI().SysCallN(1, 1, m.Instance(), api.PasStr(value))
}

func (m *TReplaceDialog) SetOnReplace(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 2, replaceDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewReplaceDialog class constructor
func NewReplaceDialog(owner IComponent) IReplaceDialog {
	r := replaceDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsReplaceDialog(r)
}

func TReplaceDialogClass() types.TClass {
	r := replaceDialogAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	replaceDialogOnce   base.Once
	replaceDialogImport *imports.Imports = nil
)

func replaceDialogAPI() *imports.Imports {
	replaceDialogOnce.Do(func() {
		replaceDialogImport = api.NewDefaultImports()
		replaceDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TReplaceDialog_Create", 0), // constructor NewReplaceDialog
			/* 1 */ imports.NewTable("TReplaceDialog_ReplaceText", 0), // property ReplaceText
			/* 2 */ imports.NewTable("TReplaceDialog_OnReplace", 0), // event OnReplace
			/* 3 */ imports.NewTable("TReplaceDialog_TClass", 0), // function TReplaceDialogClass
		}
	})
	return replaceDialogImport
}
