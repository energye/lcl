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

// IHeaderSections Parent: ICollection
type IHeaderSections interface {
	ICollection
	ItemsForHeaderSection(Index int32) IHeaderSection            // property
	SetItemsForHeaderSection(Index int32, AValue IHeaderSection) // property
	AddForHeaderSection() IHeaderSection                         // function
	AddItem(Item IHeaderSection, Index int32) IHeaderSection     // function
	InsertForHeaderSection(Index int32) IHeaderSection           // function
}

// THeaderSections Parent: TCollection
type THeaderSections struct {
	TCollection
}

func NewHeaderSections(HeaderControl ICustomHeaderControl) IHeaderSections {
	r1 := headerSectionsImportAPI().SysCallN(3, GetObjectUintptr(HeaderControl))
	return AsHeaderSections(r1)
}

func (m *THeaderSections) ItemsForHeaderSection(Index int32) IHeaderSection {
	r1 := headerSectionsImportAPI().SysCallN(5, 0, m.Instance(), uintptr(Index))
	return AsHeaderSection(r1)
}

func (m *THeaderSections) SetItemsForHeaderSection(Index int32, AValue IHeaderSection) {
	headerSectionsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *THeaderSections) AddForHeaderSection() IHeaderSection {
	r1 := headerSectionsImportAPI().SysCallN(0, m.Instance())
	return AsHeaderSection(r1)
}

func (m *THeaderSections) AddItem(Item IHeaderSection, Index int32) IHeaderSection {
	r1 := headerSectionsImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(Item), uintptr(Index))
	return AsHeaderSection(r1)
}

func (m *THeaderSections) InsertForHeaderSection(Index int32) IHeaderSection {
	r1 := headerSectionsImportAPI().SysCallN(4, m.Instance(), uintptr(Index))
	return AsHeaderSection(r1)
}

func HeaderSectionsClass() TClass {
	ret := headerSectionsImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	headerSectionsImport       *imports.Imports = nil
	headerSectionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("HeaderSections_AddForHeaderSection", 0),
		/*1*/ imports.NewTable("HeaderSections_AddItem", 0),
		/*2*/ imports.NewTable("HeaderSections_Class", 0),
		/*3*/ imports.NewTable("HeaderSections_Create", 0),
		/*4*/ imports.NewTable("HeaderSections_InsertForHeaderSection", 0),
		/*5*/ imports.NewTable("HeaderSections_ItemsForHeaderSection", 0),
	}
)

func headerSectionsImportAPI() *imports.Imports {
	if headerSectionsImport == nil {
		headerSectionsImport = NewDefaultImports()
		headerSectionsImport.SetImportTable(headerSectionsImportTables)
		headerSectionsImportTables = nil
	}
	return headerSectionsImport
}
