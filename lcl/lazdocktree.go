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

// ILazDockTree Parent: IDockTree
type ILazDockTree interface {
	IDockTree
	AutoFreeDockSite() bool                                                      // property
	SetAutoFreeDockSite(AValue bool)                                             // property
	FindBorderControl(Zone ILazDockZone, Side TAnchorKind) IControl              // function
	GetAnchorControl(Zone ILazDockZone, Side TAnchorKind, OutSide bool) IControl // function
	BuildDockLayout(Zone ILazDockZone)                                           // procedure
	FindBorderControls(Zone ILazDockZone, Side TAnchorKind, List *IFPList)       // procedure
}

// TLazDockTree Parent: TDockTree
type TLazDockTree struct {
	TDockTree
}

func NewLazDockTree(TheDockSite IWinControl) ILazDockTree {
	r1 := lazDockTreeImportAPI().SysCallN(3, GetObjectUintptr(TheDockSite))
	return AsLazDockTree(r1)
}

func (m *TLazDockTree) AutoFreeDockSite() bool {
	r1 := lazDockTreeImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazDockTree) SetAutoFreeDockSite(AValue bool) {
	lazDockTreeImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazDockTree) FindBorderControl(Zone ILazDockZone, Side TAnchorKind) IControl {
	r1 := lazDockTreeImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(Zone), uintptr(Side))
	return AsControl(r1)
}

func (m *TLazDockTree) GetAnchorControl(Zone ILazDockZone, Side TAnchorKind, OutSide bool) IControl {
	r1 := lazDockTreeImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(Zone), uintptr(Side), PascalBool(OutSide))
	return AsControl(r1)
}

func LazDockTreeClass() TClass {
	ret := lazDockTreeImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TLazDockTree) BuildDockLayout(Zone ILazDockZone) {
	lazDockTreeImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(Zone))
}

func (m *TLazDockTree) FindBorderControls(Zone ILazDockZone, Side TAnchorKind, List *IFPList) {
	var result2 uintptr
	lazDockTreeImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(Zone), uintptr(Side), uintptr(unsafePointer(&result2)))
	*List = AsFPList(result2)
}

var (
	lazDockTreeImport       *imports.Imports = nil
	lazDockTreeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazDockTree_AutoFreeDockSite", 0),
		/*1*/ imports.NewTable("LazDockTree_BuildDockLayout", 0),
		/*2*/ imports.NewTable("LazDockTree_Class", 0),
		/*3*/ imports.NewTable("LazDockTree_Create", 0),
		/*4*/ imports.NewTable("LazDockTree_FindBorderControl", 0),
		/*5*/ imports.NewTable("LazDockTree_FindBorderControls", 0),
		/*6*/ imports.NewTable("LazDockTree_GetAnchorControl", 0),
	}
)

func lazDockTreeImportAPI() *imports.Imports {
	if lazDockTreeImport == nil {
		lazDockTreeImport = NewDefaultImports()
		lazDockTreeImport.SetImportTable(lazDockTreeImportTables)
		lazDockTreeImportTables = nil
	}
	return lazDockTreeImport
}
