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
)

// IComboExItem Parent: IListControlItem
type IComboExItem interface {
	IListControlItem
	Indent() int32                     // property Indent Getter
	SetIndent(value int32)             // property Indent Setter
	OverlayImageIndex() int32          // property OverlayImageIndex Getter
	SetOverlayImageIndex(value int32)  // property OverlayImageIndex Setter
	SelectedImageIndex() int32         // property SelectedImageIndex Getter
	SetSelectedImageIndex(value int32) // property SelectedImageIndex Setter
}

type TComboExItem struct {
	TListControlItem
}

func (m *TComboExItem) Indent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := comboExItemAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TComboExItem) SetIndent(value int32) {
	if !m.IsValid() {
		return
	}
	comboExItemAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TComboExItem) OverlayImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := comboExItemAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TComboExItem) SetOverlayImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	comboExItemAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TComboExItem) SelectedImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := comboExItemAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TComboExItem) SetSelectedImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	comboExItemAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

// NewComboExItem class constructor
func NewComboExItem(collection ICollection) IComboExItem {
	r := comboExItemAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsComboExItem(r)
}

var (
	comboExItemOnce   base.Once
	comboExItemImport *imports.Imports = nil
)

func comboExItemAPI() *imports.Imports {
	comboExItemOnce.Do(func() {
		comboExItemImport = api.NewDefaultImports()
		comboExItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TComboExItem_Create", 0), // constructor NewComboExItem
			/* 1 */ imports.NewTable("TComboExItem_Indent", 0), // property Indent
			/* 2 */ imports.NewTable("TComboExItem_OverlayImageIndex", 0), // property OverlayImageIndex
			/* 3 */ imports.NewTable("TComboExItem_SelectedImageIndex", 0), // property SelectedImageIndex
		}
	})
	return comboExItemImport
}
