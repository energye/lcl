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

// IHeaderSections Parent: ICollection
type IHeaderSections interface {
	ICollection
	AddToHeaderSection() IHeaderSection                               // function
	AddItem(item IHeaderSection, index int32) IHeaderSection          // function
	InsertWithInt(index int32) IHeaderSection                         // function
	DeleteWithInt(index int32)                                        // procedure
	ItemsWithIntToHeaderSection(index int32) IHeaderSection           // property Items Getter
	SetItemsWithIntToHeaderSection(index int32, value IHeaderSection) // property Items Setter
}

type THeaderSections struct {
	TCollection
}

func (m *THeaderSections) AddToHeaderSection() IHeaderSection {
	if !m.IsValid() {
		return nil
	}
	r := headerSectionsAPI().SysCallN(1, m.Instance())
	return AsHeaderSection(r)
}

func (m *THeaderSections) AddItem(item IHeaderSection, index int32) IHeaderSection {
	if !m.IsValid() {
		return nil
	}
	r := headerSectionsAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(item), uintptr(index))
	return AsHeaderSection(r)
}

func (m *THeaderSections) InsertWithInt(index int32) IHeaderSection {
	if !m.IsValid() {
		return nil
	}
	r := headerSectionsAPI().SysCallN(3, m.Instance(), uintptr(index))
	return AsHeaderSection(r)
}

func (m *THeaderSections) DeleteWithInt(index int32) {
	if !m.IsValid() {
		return
	}
	headerSectionsAPI().SysCallN(4, m.Instance(), uintptr(index))
}

func (m *THeaderSections) ItemsWithIntToHeaderSection(index int32) IHeaderSection {
	if !m.IsValid() {
		return nil
	}
	r := headerSectionsAPI().SysCallN(5, 0, m.Instance(), uintptr(index))
	return AsHeaderSection(r)
}

func (m *THeaderSections) SetItemsWithIntToHeaderSection(index int32, value IHeaderSection) {
	if !m.IsValid() {
		return
	}
	headerSectionsAPI().SysCallN(5, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

// NewHeaderSections class constructor
func NewHeaderSections(headerControl ICustomHeaderControl) IHeaderSections {
	r := headerSectionsAPI().SysCallN(0, base.GetObjectUintptr(headerControl))
	return AsHeaderSections(r)
}

var (
	headerSectionsOnce   base.Once
	headerSectionsImport *imports.Imports = nil
)

func headerSectionsAPI() *imports.Imports {
	headerSectionsOnce.Do(func() {
		headerSectionsImport = api.NewDefaultImports()
		headerSectionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("THeaderSections_Create", 0), // constructor NewHeaderSections
			/* 1 */ imports.NewTable("THeaderSections_AddToHeaderSection", 0), // function AddToHeaderSection
			/* 2 */ imports.NewTable("THeaderSections_AddItem", 0), // function AddItem
			/* 3 */ imports.NewTable("THeaderSections_InsertWithInt", 0), // function InsertWithInt
			/* 4 */ imports.NewTable("THeaderSections_DeleteWithInt", 0), // procedure DeleteWithInt
			/* 5 */ imports.NewTable("THeaderSections_ItemsWithIntToHeaderSection", 0), // property ItemsWithIntToHeaderSection
		}
	})
	return headerSectionsImport
}
