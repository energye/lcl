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

// IScrollBarOptions Parent: IPersistent
// A class to manage scroll bar aspects.
type IScrollBarOptions interface {
	IPersistent
	AlwaysVisible() bool                              // property
	SetAlwaysVisible(AValue bool)                     // property
	HorizontalIncrement() TVTScrollIncrement          // property
	SetHorizontalIncrement(AValue TVTScrollIncrement) // property
	ScrollBars() TScrollStyle                         // property
	SetScrollBars(AValue TScrollStyle)                // property
	ScrollBarStyle() TVTScrollBarStyle                // property
	SetScrollBarStyle(AValue TVTScrollBarStyle)       // property
	VerticalIncrement() TVTScrollIncrement            // property
	SetVerticalIncrement(AValue TVTScrollIncrement)   // property
}

// TScrollBarOptions Parent: TPersistent
// A class to manage scroll bar aspects.
type TScrollBarOptions struct {
	TPersistent
}

func NewScrollBarOptions(AOwner IBaseVirtualTree) IScrollBarOptions {
	r1 := scrollBarOptionsImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsScrollBarOptions(r1)
}

func (m *TScrollBarOptions) AlwaysVisible() bool {
	r1 := scrollBarOptionsImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBarOptions) SetAlwaysVisible(AValue bool) {
	scrollBarOptionsImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBarOptions) HorizontalIncrement() TVTScrollIncrement {
	r1 := scrollBarOptionsImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TVTScrollIncrement(r1)
}

func (m *TScrollBarOptions) SetHorizontalIncrement(AValue TVTScrollIncrement) {
	scrollBarOptionsImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBarOptions) ScrollBars() TScrollStyle {
	r1 := scrollBarOptionsImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TScrollBarOptions) SetScrollBars(AValue TScrollStyle) {
	scrollBarOptionsImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBarOptions) ScrollBarStyle() TVTScrollBarStyle {
	r1 := scrollBarOptionsImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TVTScrollBarStyle(r1)
}

func (m *TScrollBarOptions) SetScrollBarStyle(AValue TVTScrollBarStyle) {
	scrollBarOptionsImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBarOptions) VerticalIncrement() TVTScrollIncrement {
	r1 := scrollBarOptionsImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TVTScrollIncrement(r1)
}

func (m *TScrollBarOptions) SetVerticalIncrement(AValue TVTScrollIncrement) {
	scrollBarOptionsImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func ScrollBarOptionsClass() TClass {
	ret := scrollBarOptionsImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	scrollBarOptionsImport       *imports.Imports = nil
	scrollBarOptionsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ScrollBarOptions_AlwaysVisible", 0),
		/*1*/ imports.NewTable("ScrollBarOptions_Class", 0),
		/*2*/ imports.NewTable("ScrollBarOptions_Create", 0),
		/*3*/ imports.NewTable("ScrollBarOptions_HorizontalIncrement", 0),
		/*4*/ imports.NewTable("ScrollBarOptions_ScrollBarStyle", 0),
		/*5*/ imports.NewTable("ScrollBarOptions_ScrollBars", 0),
		/*6*/ imports.NewTable("ScrollBarOptions_VerticalIncrement", 0),
	}
)

func scrollBarOptionsImportAPI() *imports.Imports {
	if scrollBarOptionsImport == nil {
		scrollBarOptionsImport = NewDefaultImports()
		scrollBarOptionsImport.SetImportTable(scrollBarOptionsImportTables)
		scrollBarOptionsImportTables = nil
	}
	return scrollBarOptionsImport
}
