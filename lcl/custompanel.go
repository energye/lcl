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

// ICustomPanel Parent: ICustomControl
type ICustomPanel interface {
	ICustomControl
	Alignment() TAlignment            // property
	SetAlignment(AValue TAlignment)   // property
	BevelColor() TColor               // property
	SetBevelColor(AValue TColor)      // property
	BevelInner() TPanelBevel          // property
	SetBevelInner(AValue TPanelBevel) // property
	BevelOuter() TPanelBevel          // property
	SetBevelOuter(AValue TPanelBevel) // property
	BevelWidth() TBevelWidth          // property
	SetBevelWidth(AValue TBevelWidth) // property
	FullRepaint() bool                // property
	SetFullRepaint(AValue bool)       // property
	ParentBackground() bool           // property
	SetParentBackground(AValue bool)  // property
	ParentColor() bool                // property
	SetParentColor(AValue bool)       // property
}

// TCustomPanel Parent: TCustomControl
type TCustomPanel struct {
	TCustomControl
}

func NewCustomPanel(TheOwner IComponent) ICustomPanel {
	r1 := LCL().SysCallN(2122, GetObjectUintptr(TheOwner))
	return AsCustomPanel(r1)
}

func (m *TCustomPanel) Alignment() TAlignment {
	r1 := LCL().SysCallN(2116, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomPanel) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(2116, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) BevelColor() TColor {
	r1 := LCL().SysCallN(2117, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomPanel) SetBevelColor(AValue TColor) {
	LCL().SysCallN(2117, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) BevelInner() TPanelBevel {
	r1 := LCL().SysCallN(2118, 0, m.Instance(), 0)
	return TPanelBevel(r1)
}

func (m *TCustomPanel) SetBevelInner(AValue TPanelBevel) {
	LCL().SysCallN(2118, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) BevelOuter() TPanelBevel {
	r1 := LCL().SysCallN(2119, 0, m.Instance(), 0)
	return TPanelBevel(r1)
}

func (m *TCustomPanel) SetBevelOuter(AValue TPanelBevel) {
	LCL().SysCallN(2119, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) BevelWidth() TBevelWidth {
	r1 := LCL().SysCallN(2120, 0, m.Instance(), 0)
	return TBevelWidth(r1)
}

func (m *TCustomPanel) SetBevelWidth(AValue TBevelWidth) {
	LCL().SysCallN(2120, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPanel) FullRepaint() bool {
	r1 := LCL().SysCallN(2123, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPanel) SetFullRepaint(AValue bool) {
	LCL().SysCallN(2123, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPanel) ParentBackground() bool {
	r1 := LCL().SysCallN(2124, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPanel) SetParentBackground(AValue bool) {
	LCL().SysCallN(2124, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPanel) ParentColor() bool {
	r1 := LCL().SysCallN(2125, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPanel) SetParentColor(AValue bool) {
	LCL().SysCallN(2125, 1, m.Instance(), PascalBool(AValue))
}

func CustomPanelClass() TClass {
	ret := LCL().SysCallN(2121)
	return TClass(ret)
}
