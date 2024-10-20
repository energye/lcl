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

// IComboExItem Parent: IListControlItem
type IComboExItem interface {
	IListControlItem
	Indent() int32                            // property
	SetIndent(AValue int32)                   // property
	OverlayImageIndex() TImageIndex           // property
	SetOverlayImageIndex(AValue TImageIndex)  // property
	SelectedImageIndex() TImageIndex          // property
	SetSelectedImageIndex(AValue TImageIndex) // property
}

// TComboExItem Parent: TListControlItem
type TComboExItem struct {
	TListControlItem
}

func NewComboExItem(ACollection ICollection) IComboExItem {
	r1 := comboExItemImportAPI().SysCallN(1, GetObjectUintptr(ACollection))
	return AsComboExItem(r1)
}

func (m *TComboExItem) Indent() int32 {
	r1 := comboExItemImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboExItem) SetIndent(AValue int32) {
	comboExItemImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboExItem) OverlayImageIndex() TImageIndex {
	r1 := comboExItemImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TComboExItem) SetOverlayImageIndex(AValue TImageIndex) {
	comboExItemImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboExItem) SelectedImageIndex() TImageIndex {
	r1 := comboExItemImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TComboExItem) SetSelectedImageIndex(AValue TImageIndex) {
	comboExItemImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func ComboExItemClass() TClass {
	ret := comboExItemImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	comboExItemImport       *imports.Imports = nil
	comboExItemImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ComboExItem_Class", 0),
		/*1*/ imports.NewTable("ComboExItem_Create", 0),
		/*2*/ imports.NewTable("ComboExItem_Indent", 0),
		/*3*/ imports.NewTable("ComboExItem_OverlayImageIndex", 0),
		/*4*/ imports.NewTable("ComboExItem_SelectedImageIndex", 0),
	}
)

func comboExItemImportAPI() *imports.Imports {
	if comboExItemImport == nil {
		comboExItemImport = NewDefaultImports()
		comboExItemImport.SetImportTable(comboExItemImportTables)
		comboExItemImportTables = nil
	}
	return comboExItemImport
}
