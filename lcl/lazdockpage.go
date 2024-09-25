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
	r1 := LCL().SysCallN(3555, GetObjectUintptr(TheOwner))
	return AsLazDockPage(r1)
}

func (m *TLazDockPage) DockZone() IDockZone {
	r1 := LCL().SysCallN(3556, m.Instance())
	return AsDockZone(r1)
}

func (m *TLazDockPage) PageControl() ILazDockPages {
	r1 := LCL().SysCallN(3557, m.Instance())
	return AsLazDockPages(r1)
}

func LazDockPageClass() TClass {
	ret := LCL().SysCallN(3554)
	return TClass(ret)
}
