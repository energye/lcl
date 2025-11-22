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

// IExtCommonDialog Parent: ICommonDialog
type IExtCommonDialog interface {
	ICommonDialog
	Left() int32                             // property Left Getter
	SetLeft(value int32)                     // property Left Setter
	Top() int32                              // property Top Getter
	SetTop(value int32)                      // property Top Setter
	DialogPosition() types.TPosition         // property DialogPosition Getter
	SetDialogPosition(value types.TPosition) // property DialogPosition Setter
}

type TExtCommonDialog struct {
	TCommonDialog
}

func (m *TExtCommonDialog) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := extCommonDialogAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TExtCommonDialog) SetLeft(value int32) {
	if !m.IsValid() {
		return
	}
	extCommonDialogAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TExtCommonDialog) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := extCommonDialogAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TExtCommonDialog) SetTop(value int32) {
	if !m.IsValid() {
		return
	}
	extCommonDialogAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TExtCommonDialog) DialogPosition() types.TPosition {
	if !m.IsValid() {
		return 0
	}
	r := extCommonDialogAPI().SysCallN(3, 0, m.Instance())
	return types.TPosition(r)
}

func (m *TExtCommonDialog) SetDialogPosition(value types.TPosition) {
	if !m.IsValid() {
		return
	}
	extCommonDialogAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

// NewExtCommonDialog class constructor
func NewExtCommonDialog(owner IComponent) IExtCommonDialog {
	r := extCommonDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsExtCommonDialog(r)
}

func TExtCommonDialogClass() types.TClass {
	r := extCommonDialogAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	extCommonDialogOnce   base.Once
	extCommonDialogImport *imports.Imports = nil
)

func extCommonDialogAPI() *imports.Imports {
	extCommonDialogOnce.Do(func() {
		extCommonDialogImport = api.NewDefaultImports()
		extCommonDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TExtCommonDialog_Create", 0), // constructor NewExtCommonDialog
			/* 1 */ imports.NewTable("TExtCommonDialog_Left", 0), // property Left
			/* 2 */ imports.NewTable("TExtCommonDialog_Top", 0), // property Top
			/* 3 */ imports.NewTable("TExtCommonDialog_DialogPosition", 0), // property DialogPosition
			/* 4 */ imports.NewTable("TExtCommonDialog_TClass", 0), // function TExtCommonDialogClass
		}
	})
	return extCommonDialogImport
}
