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

// ICustomStaticText Parent: IWinControl
type ICustomStaticText interface {
	IWinControl
	Alignment() TAlignment                    // property
	SetAlignment(AValue TAlignment)           // property
	BorderStyle() TStaticBorderStyle          // property
	SetBorderStyle(AValue TStaticBorderStyle) // property
	FocusControl() IWinControl                // property
	SetFocusControl(AValue IWinControl)       // property
	ShowAccelChar() bool                      // property
	SetShowAccelChar(AValue bool)             // property
	Transparent() bool                        // property
	SetTransparent(AValue bool)               // property
}

// TCustomStaticText Parent: TWinControl
type TCustomStaticText struct {
	TWinControl
}

func NewCustomStaticText(AOwner IComponent) ICustomStaticText {
	r1 := LCL().SysCallN(2296, GetObjectUintptr(AOwner))
	return AsCustomStaticText(r1)
}

func (m *TCustomStaticText) Alignment() TAlignment {
	r1 := LCL().SysCallN(2293, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomStaticText) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(2293, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomStaticText) BorderStyle() TStaticBorderStyle {
	r1 := LCL().SysCallN(2294, 0, m.Instance(), 0)
	return TStaticBorderStyle(r1)
}

func (m *TCustomStaticText) SetBorderStyle(AValue TStaticBorderStyle) {
	LCL().SysCallN(2294, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomStaticText) FocusControl() IWinControl {
	r1 := LCL().SysCallN(2297, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TCustomStaticText) SetFocusControl(AValue IWinControl) {
	LCL().SysCallN(2297, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomStaticText) ShowAccelChar() bool {
	r1 := LCL().SysCallN(2298, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStaticText) SetShowAccelChar(AValue bool) {
	LCL().SysCallN(2298, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomStaticText) Transparent() bool {
	r1 := LCL().SysCallN(2299, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomStaticText) SetTransparent(AValue bool) {
	LCL().SysCallN(2299, 1, m.Instance(), PascalBool(AValue))
}

func CustomStaticTextClass() TClass {
	ret := LCL().SysCallN(2295)
	return TClass(ret)
}
