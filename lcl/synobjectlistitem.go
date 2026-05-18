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

// ISynObjectListItem Parent: ISynEditFriend
type ISynObjectListItem interface {
	ISynEditFriend
	Index() int32         // property Index Getter
	SetIndex(value int32) // property Index Setter
	DisplayName() string  // property DisplayName Getter
}

type TSynObjectListItem struct {
	TSynEditFriend
}

func (m *TSynObjectListItem) Index() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synObjectListItemAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TSynObjectListItem) SetIndex(value int32) {
	if !m.IsValid() {
		return
	}
	synObjectListItemAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TSynObjectListItem) DisplayName() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synObjectListItemAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

// NewSynObjectListItem class constructor
func NewSynObjectListItem(owner IComponent) ISynObjectListItem {
	r := synObjectListItemAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynObjectListItem(r)
}

func TSynObjectListItemClass() types.TClass {
	r := synObjectListItemAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	synObjectListItemOnce   base.Once
	synObjectListItemImport *imports.Imports = nil
)

func synObjectListItemAPI() *imports.Imports {
	synObjectListItemOnce.Do(func() {
		synObjectListItemImport = api.NewDefaultImports()
		synObjectListItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynObjectListItem_Create", 0), // constructor NewSynObjectListItem
			/* 1 */ imports.NewTable("TSynObjectListItem_Index", 0), // property Index
			/* 2 */ imports.NewTable("TSynObjectListItem_DisplayName", 0), // property DisplayName
			/* 3 */ imports.NewTable("TSynObjectListItem_TClass", 0), // function TSynObjectListItemClass
		}
	})
	return synObjectListItemImport
}
