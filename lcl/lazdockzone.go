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
	r1 := LCL().SysCallN(3572, GetObjectUintptr(TheTree), GetObjectUintptr(TheChildControl))
	return AsLazDockZone(r1)
}

func (m *TLazDockZone) Splitter() ILazDockSplitter {
	r1 := LCL().SysCallN(3578, 0, m.Instance(), 0)
	return AsLazDockSplitter(r1)
}

func (m *TLazDockZone) SetSplitter(AValue ILazDockSplitter) {
	LCL().SysCallN(3578, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockZone) Pages() ILazDockPages {
	r1 := LCL().SysCallN(3577, 0, m.Instance(), 0)
	return AsLazDockPages(r1)
}

func (m *TLazDockZone) SetPages(AValue ILazDockPages) {
	LCL().SysCallN(3577, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockZone) Page() ILazDockPage {
	r1 := LCL().SysCallN(3576, 0, m.Instance(), 0)
	return AsLazDockPage(r1)
}

func (m *TLazDockZone) SetPage(AValue ILazDockPage) {
	LCL().SysCallN(3576, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockZone) GetCaption() string {
	r1 := LCL().SysCallN(3574, m.Instance())
	return GoStr(r1)
}

func (m *TLazDockZone) GetParentControl() IWinControl {
	r1 := LCL().SysCallN(3575, m.Instance())
	return AsWinControl(r1)
}

func LazDockZoneClass() TClass {
	ret := LCL().SysCallN(3571)
	return TClass(ret)
}

func (m *TLazDockZone) FreeSubComponents() {
	LCL().SysCallN(3573, m.Instance())
}
