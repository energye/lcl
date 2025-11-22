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

// IScrollBarOptions Parent: IPersistent
type IScrollBarOptions interface {
	IPersistent
	AlwaysVisible() bool                                   // property AlwaysVisible Getter
	SetAlwaysVisible(value bool)                           // property AlwaysVisible Setter
	HorizontalIncrement() types.TVTScrollIncrement         // property HorizontalIncrement Getter
	SetHorizontalIncrement(value types.TVTScrollIncrement) // property HorizontalIncrement Setter
	ScrollBars() types.TScrollStyle                        // property ScrollBars Getter
	SetScrollBars(value types.TScrollStyle)                // property ScrollBars Setter
	ScrollBarStyle() types.TVTScrollBarStyle               // property ScrollBarStyle Getter
	SetScrollBarStyle(value types.TVTScrollBarStyle)       // property ScrollBarStyle Setter
	VerticalIncrement() types.TVTScrollIncrement           // property VerticalIncrement Getter
	SetVerticalIncrement(value types.TVTScrollIncrement)   // property VerticalIncrement Setter
}

type TScrollBarOptions struct {
	TPersistent
}

func (m *TScrollBarOptions) AlwaysVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := scrollBarOptionsAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TScrollBarOptions) SetAlwaysVisible(value bool) {
	if !m.IsValid() {
		return
	}
	scrollBarOptionsAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TScrollBarOptions) HorizontalIncrement() types.TVTScrollIncrement {
	if !m.IsValid() {
		return 0
	}
	r := scrollBarOptionsAPI().SysCallN(2, 0, m.Instance())
	return types.TVTScrollIncrement(r)
}

func (m *TScrollBarOptions) SetHorizontalIncrement(value types.TVTScrollIncrement) {
	if !m.IsValid() {
		return
	}
	scrollBarOptionsAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TScrollBarOptions) ScrollBars() types.TScrollStyle {
	if !m.IsValid() {
		return 0
	}
	r := scrollBarOptionsAPI().SysCallN(3, 0, m.Instance())
	return types.TScrollStyle(r)
}

func (m *TScrollBarOptions) SetScrollBars(value types.TScrollStyle) {
	if !m.IsValid() {
		return
	}
	scrollBarOptionsAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TScrollBarOptions) ScrollBarStyle() types.TVTScrollBarStyle {
	if !m.IsValid() {
		return 0
	}
	r := scrollBarOptionsAPI().SysCallN(4, 0, m.Instance())
	return types.TVTScrollBarStyle(r)
}

func (m *TScrollBarOptions) SetScrollBarStyle(value types.TVTScrollBarStyle) {
	if !m.IsValid() {
		return
	}
	scrollBarOptionsAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TScrollBarOptions) VerticalIncrement() types.TVTScrollIncrement {
	if !m.IsValid() {
		return 0
	}
	r := scrollBarOptionsAPI().SysCallN(5, 0, m.Instance())
	return types.TVTScrollIncrement(r)
}

func (m *TScrollBarOptions) SetVerticalIncrement(value types.TVTScrollIncrement) {
	if !m.IsValid() {
		return
	}
	scrollBarOptionsAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

// NewScrollBarOptions class constructor
func NewScrollBarOptions(owner IBaseVirtualTree) IScrollBarOptions {
	r := scrollBarOptionsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsScrollBarOptions(r)
}

var (
	scrollBarOptionsOnce   base.Once
	scrollBarOptionsImport *imports.Imports = nil
)

func scrollBarOptionsAPI() *imports.Imports {
	scrollBarOptionsOnce.Do(func() {
		scrollBarOptionsImport = api.NewDefaultImports()
		scrollBarOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TScrollBarOptions_Create", 0), // constructor NewScrollBarOptions
			/* 1 */ imports.NewTable("TScrollBarOptions_AlwaysVisible", 0), // property AlwaysVisible
			/* 2 */ imports.NewTable("TScrollBarOptions_HorizontalIncrement", 0), // property HorizontalIncrement
			/* 3 */ imports.NewTable("TScrollBarOptions_ScrollBars", 0), // property ScrollBars
			/* 4 */ imports.NewTable("TScrollBarOptions_ScrollBarStyle", 0), // property ScrollBarStyle
			/* 5 */ imports.NewTable("TScrollBarOptions_VerticalIncrement", 0), // property VerticalIncrement
		}
	})
	return scrollBarOptionsImport
}
