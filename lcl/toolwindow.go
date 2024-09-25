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

// IToolWindow Parent: ICustomControl
type IToolWindow interface {
	ICustomControl
	EdgeBorders() TEdgeBorders          // property
	SetEdgeBorders(AValue TEdgeBorders) // property
	EdgeInner() TEdgeStyle              // property
	SetEdgeInner(AValue TEdgeStyle)     // property
	EdgeOuter() TEdgeStyle              // property
	SetEdgeOuter(AValue TEdgeStyle)     // property
	BeginUpdate()                       // procedure
	EndUpdate()                         // procedure
}

// TToolWindow Parent: TCustomControl
type TToolWindow struct {
	TCustomControl
}

func NewToolWindow(TheOwner IComponent) IToolWindow {
	r1 := LCL().SysCallN(5585, GetObjectUintptr(TheOwner))
	return AsToolWindow(r1)
}

func (m *TToolWindow) EdgeBorders() TEdgeBorders {
	r1 := LCL().SysCallN(5586, 0, m.Instance(), 0)
	return TEdgeBorders(r1)
}

func (m *TToolWindow) SetEdgeBorders(AValue TEdgeBorders) {
	LCL().SysCallN(5586, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolWindow) EdgeInner() TEdgeStyle {
	r1 := LCL().SysCallN(5587, 0, m.Instance(), 0)
	return TEdgeStyle(r1)
}

func (m *TToolWindow) SetEdgeInner(AValue TEdgeStyle) {
	LCL().SysCallN(5587, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolWindow) EdgeOuter() TEdgeStyle {
	r1 := LCL().SysCallN(5588, 0, m.Instance(), 0)
	return TEdgeStyle(r1)
}

func (m *TToolWindow) SetEdgeOuter(AValue TEdgeStyle) {
	LCL().SysCallN(5588, 1, m.Instance(), uintptr(AValue))
}

func ToolWindowClass() TClass {
	ret := LCL().SysCallN(5584)
	return TClass(ret)
}

func (m *TToolWindow) BeginUpdate() {
	LCL().SysCallN(5583, m.Instance())
}

func (m *TToolWindow) EndUpdate() {
	LCL().SysCallN(5589, m.Instance())
}
