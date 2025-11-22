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

// IScrollingWinControl Parent: ICustomControl
type IScrollingWinControl interface {
	ICustomControl
	UpdateScrollbars()                        // procedure
	ScrollInView(control IControl)            // procedure
	HorzScrollBar() IControlScrollBar         // property HorzScrollBar Getter
	SetHorzScrollBar(value IControlScrollBar) // property HorzScrollBar Setter
	VertScrollBar() IControlScrollBar         // property VertScrollBar Getter
	SetVertScrollBar(value IControlScrollBar) // property VertScrollBar Setter
}

type TScrollingWinControl struct {
	TCustomControl
}

func (m *TScrollingWinControl) UpdateScrollbars() {
	if !m.IsValid() {
		return
	}
	scrollingWinControlAPI().SysCallN(2, m.Instance())
}

func (m *TScrollingWinControl) ScrollInView(control IControl) {
	if !m.IsValid() {
		return
	}
	scrollingWinControlAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(control))
}

func (m *TScrollingWinControl) HorzScrollBar() IControlScrollBar {
	if !m.IsValid() {
		return nil
	}
	r := scrollingWinControlAPI().SysCallN(4, 0, m.Instance())
	return AsControlScrollBar(r)
}

func (m *TScrollingWinControl) SetHorzScrollBar(value IControlScrollBar) {
	if !m.IsValid() {
		return
	}
	scrollingWinControlAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TScrollingWinControl) VertScrollBar() IControlScrollBar {
	if !m.IsValid() {
		return nil
	}
	r := scrollingWinControlAPI().SysCallN(5, 0, m.Instance())
	return AsControlScrollBar(r)
}

func (m *TScrollingWinControl) SetVertScrollBar(value IControlScrollBar) {
	if !m.IsValid() {
		return
	}
	scrollingWinControlAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

// ScrollingWinControl  is static instance
var ScrollingWinControl _ScrollingWinControlClass

// _ScrollingWinControlClass is class type defined by TScrollingWinControl
type _ScrollingWinControlClass uintptr

func (_ScrollingWinControlClass) GetControlClassDefaultSize() (result types.TSize) {
	scrollingWinControlAPI().SysCallN(1, uintptr(base.UnsafePointer(&result)))
	return
}

// NewScrollingWinControl class constructor
func NewScrollingWinControl(theOwner IComponent) IScrollingWinControl {
	r := scrollingWinControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsScrollingWinControl(r)
}

func TScrollingWinControlClass() types.TClass {
	r := scrollingWinControlAPI().SysCallN(6)
	return types.TClass(r)
}

var (
	scrollingWinControlOnce   base.Once
	scrollingWinControlImport *imports.Imports = nil
)

func scrollingWinControlAPI() *imports.Imports {
	scrollingWinControlOnce.Do(func() {
		scrollingWinControlImport = api.NewDefaultImports()
		scrollingWinControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TScrollingWinControl_Create", 0), // constructor NewScrollingWinControl
			/* 1 */ imports.NewTable("TScrollingWinControl_GetControlClassDefaultSize", 0), // static function GetControlClassDefaultSize
			/* 2 */ imports.NewTable("TScrollingWinControl_UpdateScrollbars", 0), // procedure UpdateScrollbars
			/* 3 */ imports.NewTable("TScrollingWinControl_ScrollInView", 0), // procedure ScrollInView
			/* 4 */ imports.NewTable("TScrollingWinControl_HorzScrollBar", 0), // property HorzScrollBar
			/* 5 */ imports.NewTable("TScrollingWinControl_VertScrollBar", 0), // property VertScrollBar
			/* 6 */ imports.NewTable("TScrollingWinControl_TClass", 0), // function TScrollingWinControlClass
		}
	})
	return scrollingWinControlImport
}
