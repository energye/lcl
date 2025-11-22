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

// ILazDockTree Parent: IDockTree
type ILazDockTree interface {
	IDockTree
	FindBorderControl(zone ILazDockZone, side types.TAnchorKind) IControl              // function
	GetAnchorControl(zone ILazDockZone, side types.TAnchorKind, outSide bool) IControl // function
	BuildDockLayout(zone ILazDockZone)                                                 // procedure
	FindBorderControls(zone ILazDockZone, side types.TAnchorKind, list *IFPList)       // procedure
	AutoFreeDockSite() bool                                                            // property AutoFreeDockSite Getter
	SetAutoFreeDockSite(value bool)                                                    // property AutoFreeDockSite Setter
}

type TLazDockTree struct {
	TDockTree
}

func (m *TLazDockTree) FindBorderControl(zone ILazDockZone, side types.TAnchorKind) IControl {
	if !m.IsValid() {
		return nil
	}
	r := lazDockTreeAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(zone), uintptr(side))
	return AsControl(r)
}

func (m *TLazDockTree) GetAnchorControl(zone ILazDockZone, side types.TAnchorKind, outSide bool) IControl {
	if !m.IsValid() {
		return nil
	}
	r := lazDockTreeAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(zone), uintptr(side), api.PasBool(outSide))
	return AsControl(r)
}

func (m *TLazDockTree) BuildDockLayout(zone ILazDockZone) {
	if !m.IsValid() {
		return
	}
	lazDockTreeAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(zone))
}

func (m *TLazDockTree) FindBorderControls(zone ILazDockZone, side types.TAnchorKind, list *IFPList) {
	if !m.IsValid() {
		return
	}
	listPtr := base.GetObjectUintptr(*list)
	lazDockTreeAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(zone), uintptr(side), uintptr(base.UnsafePointer(&listPtr)))
	*list = AsFPList(listPtr)
}

func (m *TLazDockTree) AutoFreeDockSite() bool {
	if !m.IsValid() {
		return false
	}
	r := lazDockTreeAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazDockTree) SetAutoFreeDockSite(value bool) {
	if !m.IsValid() {
		return
	}
	lazDockTreeAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

// NewLazDockTree class constructor
func NewLazDockTree(theDockSite IWinControl) ILazDockTree {
	r := lazDockTreeAPI().SysCallN(0, base.GetObjectUintptr(theDockSite))
	return AsLazDockTree(r)
}

var (
	lazDockTreeOnce   base.Once
	lazDockTreeImport *imports.Imports = nil
)

func lazDockTreeAPI() *imports.Imports {
	lazDockTreeOnce.Do(func() {
		lazDockTreeImport = api.NewDefaultImports()
		lazDockTreeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazDockTree_Create", 0), // constructor NewLazDockTree
			/* 1 */ imports.NewTable("TLazDockTree_FindBorderControl", 0), // function FindBorderControl
			/* 2 */ imports.NewTable("TLazDockTree_GetAnchorControl", 0), // function GetAnchorControl
			/* 3 */ imports.NewTable("TLazDockTree_BuildDockLayout", 0), // procedure BuildDockLayout
			/* 4 */ imports.NewTable("TLazDockTree_FindBorderControls", 0), // procedure FindBorderControls
			/* 5 */ imports.NewTable("TLazDockTree_AutoFreeDockSite", 0), // property AutoFreeDockSite
		}
	})
	return lazDockTreeImport
}
