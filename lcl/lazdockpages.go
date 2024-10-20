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

// ILazDockPages Parent: ICustomTabControl
type ILazDockPages interface {
	ICustomTabControl
	PageForLazDockPage(Index int32) ILazDockPage // property
	ActivePageComponent() ILazDockPage           // property
	SetActivePageComponent(AValue ILazDockPage)  // property
}

// TLazDockPages Parent: TCustomTabControl
type TLazDockPages struct {
	TCustomTabControl
}

func NewLazDockPages(TheOwner IComponent) ILazDockPages {
	r1 := lazDockPagesImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsLazDockPages(r1)
}

func (m *TLazDockPages) PageForLazDockPage(Index int32) ILazDockPage {
	r1 := lazDockPagesImportAPI().SysCallN(3, m.Instance(), uintptr(Index))
	return AsLazDockPage(r1)
}

func (m *TLazDockPages) ActivePageComponent() ILazDockPage {
	r1 := lazDockPagesImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsLazDockPage(r1)
}

func (m *TLazDockPages) SetActivePageComponent(AValue ILazDockPage) {
	lazDockPagesImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func LazDockPagesClass() TClass {
	ret := lazDockPagesImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	lazDockPagesImport       *imports.Imports = nil
	lazDockPagesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazDockPages_ActivePageComponent", 0),
		/*1*/ imports.NewTable("LazDockPages_Class", 0),
		/*2*/ imports.NewTable("LazDockPages_Create", 0),
		/*3*/ imports.NewTable("LazDockPages_PageForLazDockPage", 0),
	}
)

func lazDockPagesImportAPI() *imports.Imports {
	if lazDockPagesImport == nil {
		lazDockPagesImport = NewDefaultImports()
		lazDockPagesImport.SetImportTable(lazDockPagesImportTables)
		lazDockPagesImportTables = nil
	}
	return lazDockPagesImport
}
