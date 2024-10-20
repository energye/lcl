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

// IScrollingWinControl Parent: ICustomControl
type IScrollingWinControl interface {
	ICustomControl
	HorzScrollBar() IControlScrollBar          // property
	SetHorzScrollBar(AValue IControlScrollBar) // property
	VertScrollBar() IControlScrollBar          // property
	SetVertScrollBar(AValue IControlScrollBar) // property
	UpdateScrollbars()                         // procedure
	ScrollInView(AControl IControl)            // procedure
}

// TScrollingWinControl Parent: TCustomControl
type TScrollingWinControl struct {
	TCustomControl
}

func NewScrollingWinControl(TheOwner IComponent) IScrollingWinControl {
	r1 := scrollingWinControlImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsScrollingWinControl(r1)
}

func (m *TScrollingWinControl) HorzScrollBar() IControlScrollBar {
	r1 := scrollingWinControlImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return AsControlScrollBar(r1)
}

func (m *TScrollingWinControl) SetHorzScrollBar(AValue IControlScrollBar) {
	scrollingWinControlImportAPI().SysCallN(2, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScrollingWinControl) VertScrollBar() IControlScrollBar {
	r1 := scrollingWinControlImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsControlScrollBar(r1)
}

func (m *TScrollingWinControl) SetVertScrollBar(AValue IControlScrollBar) {
	scrollingWinControlImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func ScrollingWinControlClass() TClass {
	ret := scrollingWinControlImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TScrollingWinControl) UpdateScrollbars() {
	scrollingWinControlImportAPI().SysCallN(4, m.Instance())
}

func (m *TScrollingWinControl) ScrollInView(AControl IControl) {
	scrollingWinControlImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(AControl))
}

var (
	scrollingWinControlImport       *imports.Imports = nil
	scrollingWinControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ScrollingWinControl_Class", 0),
		/*1*/ imports.NewTable("ScrollingWinControl_Create", 0),
		/*2*/ imports.NewTable("ScrollingWinControl_HorzScrollBar", 0),
		/*3*/ imports.NewTable("ScrollingWinControl_ScrollInView", 0),
		/*4*/ imports.NewTable("ScrollingWinControl_UpdateScrollbars", 0),
		/*5*/ imports.NewTable("ScrollingWinControl_VertScrollBar", 0),
	}
)

func scrollingWinControlImportAPI() *imports.Imports {
	if scrollingWinControlImport == nil {
		scrollingWinControlImport = NewDefaultImports()
		scrollingWinControlImport.SetImportTable(scrollingWinControlImportTables)
		scrollingWinControlImportTables = nil
	}
	return scrollingWinControlImport
}
