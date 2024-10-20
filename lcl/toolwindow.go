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
	r1 := oolWindowImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsToolWindow(r1)
}

func (m *TToolWindow) EdgeBorders() TEdgeBorders {
	r1 := oolWindowImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TEdgeBorders(r1)
}

func (m *TToolWindow) SetEdgeBorders(AValue TEdgeBorders) {
	oolWindowImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolWindow) EdgeInner() TEdgeStyle {
	r1 := oolWindowImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TEdgeStyle(r1)
}

func (m *TToolWindow) SetEdgeInner(AValue TEdgeStyle) {
	oolWindowImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolWindow) EdgeOuter() TEdgeStyle {
	r1 := oolWindowImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TEdgeStyle(r1)
}

func (m *TToolWindow) SetEdgeOuter(AValue TEdgeStyle) {
	oolWindowImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func ToolWindowClass() TClass {
	ret := oolWindowImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TToolWindow) BeginUpdate() {
	oolWindowImportAPI().SysCallN(0, m.Instance())
}

func (m *TToolWindow) EndUpdate() {
	oolWindowImportAPI().SysCallN(6, m.Instance())
}

var (
	oolWindowImport       *imports.Imports = nil
	oolWindowImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ToolWindow_BeginUpdate", 0),
		/*1*/ imports.NewTable("ToolWindow_Class", 0),
		/*2*/ imports.NewTable("ToolWindow_Create", 0),
		/*3*/ imports.NewTable("ToolWindow_EdgeBorders", 0),
		/*4*/ imports.NewTable("ToolWindow_EdgeInner", 0),
		/*5*/ imports.NewTable("ToolWindow_EdgeOuter", 0),
		/*6*/ imports.NewTable("ToolWindow_EndUpdate", 0),
	}
)

func oolWindowImportAPI() *imports.Imports {
	if oolWindowImport == nil {
		oolWindowImport = NewDefaultImports()
		oolWindowImport.SetImportTable(oolWindowImportTables)
		oolWindowImportTables = nil
	}
	return oolWindowImport
}
