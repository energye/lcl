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

// IDockTree Parent: IDockManager
type IDockTree interface {
	IDockManager
	DockZoneClass() TDockZoneClass                  // property
	DockSite() IWinControl                          // property
	SetDockSite(AValue IWinControl)                 // property
	RootZone() IDockZone                            // property
	AdjustDockRect(AControl IControl, ARect *TRect) // procedure
	DumpLayout(FileName string)                     // procedure
}

// TDockTree Parent: TDockManager
type TDockTree struct {
	TDockManager
}

func NewDockTree(TheDockSite IWinControl) IDockTree {
	r1 := dockTreeImportAPI().SysCallN(2, GetObjectUintptr(TheDockSite))
	return AsDockTree(r1)
}

func (m *TDockTree) DockZoneClass() TDockZoneClass {
	r1 := dockTreeImportAPI().SysCallN(4, m.Instance())
	return TDockZoneClass(r1)
}

func (m *TDockTree) DockSite() IWinControl {
	r1 := dockTreeImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TDockTree) SetDockSite(AValue IWinControl) {
	dockTreeImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDockTree) RootZone() IDockZone {
	r1 := dockTreeImportAPI().SysCallN(6, m.Instance())
	return AsDockZone(r1)
}

func DockTreeClass() TClass {
	ret := dockTreeImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TDockTree) AdjustDockRect(AControl IControl, ARect *TRect) {
	var result1 uintptr
	dockTreeImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafePointer(&result1)))
	*ARect = *(*TRect)(getPointer(result1))
}

func (m *TDockTree) DumpLayout(FileName string) {
	dockTreeImportAPI().SysCallN(5, m.Instance(), PascalStr(FileName))
}

var (
	dockTreeImport       *imports.Imports = nil
	dockTreeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DockTree_AdjustDockRect", 0),
		/*1*/ imports.NewTable("DockTree_Class", 0),
		/*2*/ imports.NewTable("DockTree_Create", 0),
		/*3*/ imports.NewTable("DockTree_DockSite", 0),
		/*4*/ imports.NewTable("DockTree_DockZoneClass", 0),
		/*5*/ imports.NewTable("DockTree_DumpLayout", 0),
		/*6*/ imports.NewTable("DockTree_RootZone", 0),
	}
)

func dockTreeImportAPI() *imports.Imports {
	if dockTreeImport == nil {
		dockTreeImport = NewDefaultImports()
		dockTreeImport.SetImportTable(dockTreeImportTables)
		dockTreeImportTables = nil
	}
	return dockTreeImport
}
