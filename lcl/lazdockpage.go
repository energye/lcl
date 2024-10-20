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

// ILazDockPage Parent: ICustomPage
type ILazDockPage interface {
	ICustomPage
	DockZone() IDockZone        // property
	PageControl() ILazDockPages // property
}

// TLazDockPage Parent: TCustomPage
type TLazDockPage struct {
	TCustomPage
}

func NewLazDockPage(TheOwner IComponent) ILazDockPage {
	r1 := lazDockPageImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsLazDockPage(r1)
}

func (m *TLazDockPage) DockZone() IDockZone {
	r1 := lazDockPageImportAPI().SysCallN(2, m.Instance())
	return AsDockZone(r1)
}

func (m *TLazDockPage) PageControl() ILazDockPages {
	r1 := lazDockPageImportAPI().SysCallN(3, m.Instance())
	return AsLazDockPages(r1)
}

func LazDockPageClass() TClass {
	ret := lazDockPageImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	lazDockPageImport       *imports.Imports = nil
	lazDockPageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazDockPage_Class", 0),
		/*1*/ imports.NewTable("LazDockPage_Create", 0),
		/*2*/ imports.NewTable("LazDockPage_DockZone", 0),
		/*3*/ imports.NewTable("LazDockPage_PageControl", 0),
	}
)

func lazDockPageImportAPI() *imports.Imports {
	if lazDockPageImport == nil {
		lazDockPageImport = NewDefaultImports()
		lazDockPageImport.SetImportTable(lazDockPageImportTables)
		lazDockPageImportTables = nil
	}
	return lazDockPageImport
}
