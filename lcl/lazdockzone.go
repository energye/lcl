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

// ILazDockZone Parent: IDockZone
type ILazDockZone interface {
	IDockZone
	Splitter() ILazDockSplitter          // property
	SetSplitter(AValue ILazDockSplitter) // property
	Pages() ILazDockPages                // property
	SetPages(AValue ILazDockPages)       // property
	Page() ILazDockPage                  // property
	SetPage(AValue ILazDockPage)         // property
	GetCaption() string                  // function
	GetParentControl() IWinControl       // function
	FreeSubComponents()                  // procedure
}

// TLazDockZone Parent: TDockZone
type TLazDockZone struct {
	TDockZone
}

func NewLazDockZone(TheTree IDockTree, TheChildControl IControl) ILazDockZone {
	r1 := lazDockZoneImportAPI().SysCallN(1, GetObjectUintptr(TheTree), GetObjectUintptr(TheChildControl))
	return AsLazDockZone(r1)
}

func (m *TLazDockZone) Splitter() ILazDockSplitter {
	r1 := lazDockZoneImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return AsLazDockSplitter(r1)
}

func (m *TLazDockZone) SetSplitter(AValue ILazDockSplitter) {
	lazDockZoneImportAPI().SysCallN(7, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockZone) Pages() ILazDockPages {
	r1 := lazDockZoneImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return AsLazDockPages(r1)
}

func (m *TLazDockZone) SetPages(AValue ILazDockPages) {
	lazDockZoneImportAPI().SysCallN(6, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockZone) Page() ILazDockPage {
	r1 := lazDockZoneImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsLazDockPage(r1)
}

func (m *TLazDockZone) SetPage(AValue ILazDockPage) {
	lazDockZoneImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockZone) GetCaption() string {
	r1 := lazDockZoneImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TLazDockZone) GetParentControl() IWinControl {
	r1 := lazDockZoneImportAPI().SysCallN(4, m.Instance())
	return AsWinControl(r1)
}

func LazDockZoneClass() TClass {
	ret := lazDockZoneImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TLazDockZone) FreeSubComponents() {
	lazDockZoneImportAPI().SysCallN(2, m.Instance())
}

var (
	lazDockZoneImport       *imports.Imports = nil
	lazDockZoneImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazDockZone_Class", 0),
		/*1*/ imports.NewTable("LazDockZone_Create", 0),
		/*2*/ imports.NewTable("LazDockZone_FreeSubComponents", 0),
		/*3*/ imports.NewTable("LazDockZone_GetCaption", 0),
		/*4*/ imports.NewTable("LazDockZone_GetParentControl", 0),
		/*5*/ imports.NewTable("LazDockZone_Page", 0),
		/*6*/ imports.NewTable("LazDockZone_Pages", 0),
		/*7*/ imports.NewTable("LazDockZone_Splitter", 0),
	}
)

func lazDockZoneImportAPI() *imports.Imports {
	if lazDockZoneImport == nil {
		lazDockZoneImport = NewDefaultImports()
		lazDockZoneImport.SetImportTable(lazDockZoneImportTables)
		lazDockZoneImportTables = nil
	}
	return lazDockZoneImport
}
