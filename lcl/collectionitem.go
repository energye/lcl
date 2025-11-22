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

// ICollectionItem Parent: IPersistent
type ICollectionItem interface {
	IPersistent
	Collection() ICollection         // property Collection Getter
	SetCollection(value ICollection) // property Collection Setter
	ID() int32                       // property ID Getter
	Index() int32                    // property Index Getter
	SetIndex(value int32)            // property Index Setter
	DisplayName() string             // property DisplayName Getter
	SetDisplayName(value string)     // property DisplayName Setter
}

type TCollectionItem struct {
	TPersistent
}

func (m *TCollectionItem) Collection() ICollection {
	if !m.IsValid() {
		return nil
	}
	r := collectionItemAPI().SysCallN(1, 0, m.Instance())
	return AsCollection(r)
}

func (m *TCollectionItem) SetCollection(value ICollection) {
	if !m.IsValid() {
		return
	}
	collectionItemAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCollectionItem) ID() int32 {
	if !m.IsValid() {
		return 0
	}
	r := collectionItemAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TCollectionItem) Index() int32 {
	if !m.IsValid() {
		return 0
	}
	r := collectionItemAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TCollectionItem) SetIndex(value int32) {
	if !m.IsValid() {
		return
	}
	collectionItemAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCollectionItem) DisplayName() string {
	if !m.IsValid() {
		return ""
	}
	r := collectionItemAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCollectionItem) SetDisplayName(value string) {
	if !m.IsValid() {
		return
	}
	collectionItemAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

// NewCollectionItem class constructor
func NewCollectionItem(collection ICollection) ICollectionItem {
	r := collectionItemAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsCollectionItem(r)
}

var (
	collectionItemOnce   base.Once
	collectionItemImport *imports.Imports = nil
)

func collectionItemAPI() *imports.Imports {
	collectionItemOnce.Do(func() {
		collectionItemImport = api.NewDefaultImports()
		collectionItemImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCollectionItem_Create", 0), // constructor NewCollectionItem
			/* 1 */ imports.NewTable("TCollectionItem_Collection", 0), // property Collection
			/* 2 */ imports.NewTable("TCollectionItem_ID", 0), // property ID
			/* 3 */ imports.NewTable("TCollectionItem_Index", 0), // property Index
			/* 4 */ imports.NewTable("TCollectionItem_DisplayName", 0), // property DisplayName
		}
	})
	return collectionItemImport
}
