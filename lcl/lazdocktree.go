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
	r1 := LCL().SysCallN(3567, GetObjectUintptr(TheDockSite))
	return AsLazDockTree(r1)
}

func (m *TLazDockTree) AutoFreeDockSite() bool {
	r1 := LCL().SysCallN(3564, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazDockTree) SetAutoFreeDockSite(AValue bool) {
	LCL().SysCallN(3564, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazDockTree) FindBorderControl(Zone ILazDockZone, Side TAnchorKind) IControl {
	r1 := LCL().SysCallN(3568, m.Instance(), GetObjectUintptr(Zone), uintptr(Side))
	return AsControl(r1)
}

func (m *TLazDockTree) GetAnchorControl(Zone ILazDockZone, Side TAnchorKind, OutSide bool) IControl {
	r1 := LCL().SysCallN(3570, m.Instance(), GetObjectUintptr(Zone), uintptr(Side), PascalBool(OutSide))
	return AsControl(r1)
}

func LazDockTreeClass() TClass {
	ret := LCL().SysCallN(3566)
	return TClass(ret)
}

func (m *TLazDockTree) BuildDockLayout(Zone ILazDockZone) {
	LCL().SysCallN(3565, m.Instance(), GetObjectUintptr(Zone))
}

func (m *TLazDockTree) FindBorderControls(Zone ILazDockZone, Side TAnchorKind, List *IFPList) {
	var result2 uintptr
	LCL().SysCallN(3569, m.Instance(), GetObjectUintptr(Zone), uintptr(Side), uintptr(unsafePointer(&result2)))
	*List = AsFPList(result2)
}
