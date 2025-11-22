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

// IFindDialog Parent: ICommonDialog
type IFindDialog interface {
	ICommonDialog
	CloseDialog()                        // procedure
	Left() int32                         // property Left Getter
	SetLeft(value int32)                 // property Left Setter
	Position() types.TPoint              // property Position Getter
	SetPosition(value types.TPoint)      // property Position Setter
	Top() int32                          // property Top Getter
	SetTop(value int32)                  // property Top Setter
	FindText() string                    // property FindText Getter
	SetFindText(value string)            // property FindText Setter
	Options() types.TFindOptions         // property Options Getter
	SetOptions(value types.TFindOptions) // property Options Setter
	SetOnFind(fn TNotifyEvent)           // property event
	SetOnHelpClicked(fn TNotifyEvent)    // property event
}

type TFindDialog struct {
	TCommonDialog
}

func (m *TFindDialog) CloseDialog() {
	if !m.IsValid() {
		return
	}
	findDialogAPI().SysCallN(1, m.Instance())
}

func (m *TFindDialog) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := findDialogAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TFindDialog) SetLeft(value int32) {
	if !m.IsValid() {
		return
	}
	findDialogAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TFindDialog) Position() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	findDialogAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TFindDialog) SetPosition(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	findDialogAPI().SysCallN(3, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TFindDialog) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := findDialogAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TFindDialog) SetTop(value int32) {
	if !m.IsValid() {
		return
	}
	findDialogAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TFindDialog) FindText() string {
	if !m.IsValid() {
		return ""
	}
	r := findDialogAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFindDialog) SetFindText(value string) {
	if !m.IsValid() {
		return
	}
	findDialogAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TFindDialog) Options() types.TFindOptions {
	if !m.IsValid() {
		return 0
	}
	r := findDialogAPI().SysCallN(6, 0, m.Instance())
	return types.TFindOptions(r)
}

func (m *TFindDialog) SetOptions(value types.TFindOptions) {
	if !m.IsValid() {
		return
	}
	findDialogAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TFindDialog) SetOnFind(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, findDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFindDialog) SetOnHelpClicked(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, findDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewFindDialog class constructor
func NewFindDialog(owner IComponent) IFindDialog {
	r := findDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsFindDialog(r)
}

func TFindDialogClass() types.TClass {
	r := findDialogAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	findDialogOnce   base.Once
	findDialogImport *imports.Imports = nil
)

func findDialogAPI() *imports.Imports {
	findDialogOnce.Do(func() {
		findDialogImport = api.NewDefaultImports()
		findDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFindDialog_Create", 0), // constructor NewFindDialog
			/* 1 */ imports.NewTable("TFindDialog_CloseDialog", 0), // procedure CloseDialog
			/* 2 */ imports.NewTable("TFindDialog_Left", 0), // property Left
			/* 3 */ imports.NewTable("TFindDialog_Position", 0), // property Position
			/* 4 */ imports.NewTable("TFindDialog_Top", 0), // property Top
			/* 5 */ imports.NewTable("TFindDialog_FindText", 0), // property FindText
			/* 6 */ imports.NewTable("TFindDialog_Options", 0), // property Options
			/* 7 */ imports.NewTable("TFindDialog_OnFind", 0), // event OnFind
			/* 8 */ imports.NewTable("TFindDialog_OnHelpClicked", 0), // event OnHelpClicked
			/* 9 */ imports.NewTable("TFindDialog_TClass", 0), // function TFindDialogClass
		}
	})
	return findDialogImport
}
