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

// IListControlItem Parent: ICollectionItem
type IListControlItem interface {
	ICollectionItem
	Data() types.TCustomData         // property Data Getter
	SetData(value types.TCustomData) // property Data Setter
	Caption() string                 // property Caption Getter
	SetCaption(value string)         // property Caption Setter
	ImageIndex() int32               // property ImageIndex Getter
	SetImageIndex(value int32)       // property ImageIndex Setter
}

type TListControlItem struct {
	TCollectionItem
}

func (m *TListControlItem) Data() types.TCustomData {
	if !m.IsValid() {
		return 0
	}
	r := listControlItemAPI().SysCallN(1, 0, m.Instance())
	return types.TCustomData(r)
}

func (m *TListControlItem) SetData(value types.TCustomData) {
	if !m.IsValid() {
		return
	}
	listControlItemAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TListControlItem) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := listControlItemAPI().SysCallN(2, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TListControlItem) SetCaption(value string) {
	if !m.IsValid() {
		return
	}
	listControlItemAPI().SysCallN(2, 1, m.Instance(), api.PasStr(value))
}

func (m *TListControlItem) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listControlItemAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TListControlItem) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	listControlItemAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

// NewListControlItem class constructor
func NewListControlItem(collection ICollection) IListControlItem {
	r := listControlItemAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsListControlItem(r)
}

var (
	listControlItemOnce   base.Once
	listControlItemImport *imports.Imports = nil
)

func listControlItemAPI() *imports.Imports {
	listControlItemOnce.Do(func() {
		listControlItemImport = api.NewDefaultImports()
		listControlItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListControlItem_Create", 0), // constructor NewListControlItem
			/* 1 */ imports.NewTable("TListControlItem_Data", 0), // property Data
			/* 2 */ imports.NewTable("TListControlItem_Caption", 0), // property Caption
			/* 3 */ imports.NewTable("TListControlItem_ImageIndex", 0), // property ImageIndex
		}
	})
	return listControlItemImport
}
