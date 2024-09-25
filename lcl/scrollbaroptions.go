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
	r1 := LCL().SysCallN(4960, GetObjectUintptr(AOwner))
	return AsScrollBarOptions(r1)
}

func (m *TScrollBarOptions) AlwaysVisible() bool {
	r1 := LCL().SysCallN(4958, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBarOptions) SetAlwaysVisible(AValue bool) {
	LCL().SysCallN(4958, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBarOptions) HorizontalIncrement() TVTScrollIncrement {
	r1 := LCL().SysCallN(4961, 0, m.Instance(), 0)
	return TVTScrollIncrement(r1)
}

func (m *TScrollBarOptions) SetHorizontalIncrement(AValue TVTScrollIncrement) {
	LCL().SysCallN(4961, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBarOptions) ScrollBars() TScrollStyle {
	r1 := LCL().SysCallN(4963, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TScrollBarOptions) SetScrollBars(AValue TScrollStyle) {
	LCL().SysCallN(4963, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBarOptions) ScrollBarStyle() TVTScrollBarStyle {
	r1 := LCL().SysCallN(4962, 0, m.Instance(), 0)
	return TVTScrollBarStyle(r1)
}

func (m *TScrollBarOptions) SetScrollBarStyle(AValue TVTScrollBarStyle) {
	LCL().SysCallN(4962, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBarOptions) VerticalIncrement() TVTScrollIncrement {
	r1 := LCL().SysCallN(4964, 0, m.Instance(), 0)
	return TVTScrollIncrement(r1)
}

func (m *TScrollBarOptions) SetVerticalIncrement(AValue TVTScrollIncrement) {
	LCL().SysCallN(4964, 1, m.Instance(), uintptr(AValue))
}

func ScrollBarOptionsClass() TClass {
	ret := LCL().SysCallN(4959)
	return TClass(ret)
}
