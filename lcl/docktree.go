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

// IDockTree Parent: IDockManager
type IDockTree interface {
	IDockManager
	AdjustDockRect(control IControl, rect *types.TRect) // procedure
	DumpLayout(fileName string)                         // procedure
	DockZoneClass() types.TDockZoneClass                // property DockZoneClass Getter
	DockSite() IWinControl                              // property DockSite Getter
	SetDockSite(value IWinControl)                      // property DockSite Setter
	RootZone() IDockZone                                // property RootZone Getter
}

type TDockTree struct {
	TDockManager
}

func (m *TDockTree) AdjustDockRect(control IControl, rect *types.TRect) {
	if !m.IsValid() {
		return
	}
	dockTreeAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(control), uintptr(base.UnsafePointer(rect)))
}

func (m *TDockTree) DumpLayout(fileName string) {
	if !m.IsValid() {
		return
	}
	dockTreeAPI().SysCallN(2, m.Instance(), api.PasStr(fileName))
}

func (m *TDockTree) DockZoneClass() types.TDockZoneClass {
	if !m.IsValid() {
		return 0
	}
	r := dockTreeAPI().SysCallN(3, m.Instance())
	return types.TDockZoneClass(r)
}

func (m *TDockTree) DockSite() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := dockTreeAPI().SysCallN(4, 0, m.Instance())
	return AsWinControl(r)
}

func (m *TDockTree) SetDockSite(value IWinControl) {
	if !m.IsValid() {
		return
	}
	dockTreeAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDockTree) RootZone() IDockZone {
	if !m.IsValid() {
		return nil
	}
	r := dockTreeAPI().SysCallN(5, m.Instance())
	return AsDockZone(r)
}

// NewDockTree class constructor
func NewDockTree(theDockSite IWinControl) IDockTree {
	r := dockTreeAPI().SysCallN(0, base.GetObjectUintptr(theDockSite))
	return AsDockTree(r)
}

var (
	dockTreeOnce   base.Once
	dockTreeImport *imports.Imports = nil
)

func dockTreeAPI() *imports.Imports {
	dockTreeOnce.Do(func() {
		dockTreeImport = api.NewDefaultImports()
		dockTreeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDockTree_Create", 0), // constructor NewDockTree
			/* 1 */ imports.NewTable("TDockTree_AdjustDockRect", 0), // procedure AdjustDockRect
			/* 2 */ imports.NewTable("TDockTree_DumpLayout", 0), // procedure DumpLayout
			/* 3 */ imports.NewTable("TDockTree_DockZoneClass", 0), // property DockZoneClass
			/* 4 */ imports.NewTable("TDockTree_DockSite", 0), // property DockSite
			/* 5 */ imports.NewTable("TDockTree_RootZone", 0), // property RootZone
		}
	})
	return dockTreeImport
}
