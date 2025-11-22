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

// IToolWindow Parent: ICustomControl
type IToolWindow interface {
	ICustomControl
	BeginUpdate()                            // procedure
	EndUpdate()                              // procedure
	EdgeBorders() types.TEdgeBorders         // property EdgeBorders Getter
	SetEdgeBorders(value types.TEdgeBorders) // property EdgeBorders Setter
	EdgeInner() types.TEdgeStyle             // property EdgeInner Getter
	SetEdgeInner(value types.TEdgeStyle)     // property EdgeInner Setter
	EdgeOuter() types.TEdgeStyle             // property EdgeOuter Getter
	SetEdgeOuter(value types.TEdgeStyle)     // property EdgeOuter Setter
}

type TToolWindow struct {
	TCustomControl
}

func (m *TToolWindow) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	toolWindowAPI().SysCallN(1, m.Instance())
}

func (m *TToolWindow) EndUpdate() {
	if !m.IsValid() {
		return
	}
	toolWindowAPI().SysCallN(2, m.Instance())
}

func (m *TToolWindow) EdgeBorders() types.TEdgeBorders {
	if !m.IsValid() {
		return 0
	}
	r := toolWindowAPI().SysCallN(3, 0, m.Instance())
	return types.TEdgeBorders(r)
}

func (m *TToolWindow) SetEdgeBorders(value types.TEdgeBorders) {
	if !m.IsValid() {
		return
	}
	toolWindowAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TToolWindow) EdgeInner() types.TEdgeStyle {
	if !m.IsValid() {
		return 0
	}
	r := toolWindowAPI().SysCallN(4, 0, m.Instance())
	return types.TEdgeStyle(r)
}

func (m *TToolWindow) SetEdgeInner(value types.TEdgeStyle) {
	if !m.IsValid() {
		return
	}
	toolWindowAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TToolWindow) EdgeOuter() types.TEdgeStyle {
	if !m.IsValid() {
		return 0
	}
	r := toolWindowAPI().SysCallN(5, 0, m.Instance())
	return types.TEdgeStyle(r)
}

func (m *TToolWindow) SetEdgeOuter(value types.TEdgeStyle) {
	if !m.IsValid() {
		return
	}
	toolWindowAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

// NewToolWindow class constructor
func NewToolWindow(theOwner IComponent) IToolWindow {
	r := toolWindowAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsToolWindow(r)
}

func TToolWindowClass() types.TClass {
	r := toolWindowAPI().SysCallN(6)
	return types.TClass(r)
}

var (
	toolWindowOnce   base.Once
	toolWindowImport *imports.Imports = nil
)

func toolWindowAPI() *imports.Imports {
	toolWindowOnce.Do(func() {
		toolWindowImport = api.NewDefaultImports()
		toolWindowImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TToolWindow_Create", 0), // constructor NewToolWindow
			/* 1 */ imports.NewTable("TToolWindow_BeginUpdate", 0), // procedure BeginUpdate
			/* 2 */ imports.NewTable("TToolWindow_EndUpdate", 0), // procedure EndUpdate
			/* 3 */ imports.NewTable("TToolWindow_EdgeBorders", 0), // property EdgeBorders
			/* 4 */ imports.NewTable("TToolWindow_EdgeInner", 0), // property EdgeInner
			/* 5 */ imports.NewTable("TToolWindow_EdgeOuter", 0), // property EdgeOuter
			/* 6 */ imports.NewTable("TToolWindow_TClass", 0), // function TToolWindowClass
		}
	})
	return toolWindowImport
}
