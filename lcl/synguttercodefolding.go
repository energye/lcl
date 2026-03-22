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

// ISynGutterCodeFolding Parent: ISynGutterPartBase
type ISynGutterCodeFolding interface {
	ISynGutterPartBase
	PaintWithCanvasRectIntX2(canvas ICanvas, clip types.TRect, firstLine int32, lastLine int32) // procedure
	MarkupInfoCurrentFold() ISynSelectedColor                                                   // property MarkupInfoCurrentFold Getter
	SetMarkupInfoCurrentFold(value ISynSelectedColor)                                           // property MarkupInfoCurrentFold Setter
	MouseActionsExpanded() ISynEditMouseActions                                                 // property MouseActionsExpanded Getter
	SetMouseActionsExpanded(value ISynEditMouseActions)                                         // property MouseActionsExpanded Setter
	MouseActionsCollapsed() ISynEditMouseActions                                                // property MouseActionsCollapsed Getter
	SetMouseActionsCollapsed(value ISynEditMouseActions)                                        // property MouseActionsCollapsed Setter
	ReversePopMenuOrder() bool                                                                  // property ReversePopMenuOrder Getter
	SetReversePopMenuOrder(value bool)                                                          // property ReversePopMenuOrder Setter
}

type TSynGutterCodeFolding struct {
	TSynGutterPartBase
}

func (m *TSynGutterCodeFolding) PaintWithCanvasRectIntX2(canvas ICanvas, clip types.TRect, firstLine int32, lastLine int32) {
	if !m.IsValid() {
		return
	}
	synGutterCodeFoldingAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(canvas), uintptr(base.UnsafePointer(&clip)), uintptr(firstLine), uintptr(lastLine))
}

func (m *TSynGutterCodeFolding) MarkupInfoCurrentFold() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synGutterCodeFoldingAPI().SysCallN(2, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TSynGutterCodeFolding) SetMarkupInfoCurrentFold(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	synGutterCodeFoldingAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynGutterCodeFolding) MouseActionsExpanded() ISynEditMouseActions {
	if !m.IsValid() {
		return nil
	}
	r := synGutterCodeFoldingAPI().SysCallN(3, 0, m.Instance())
	return AsSynEditMouseActions(r)
}

func (m *TSynGutterCodeFolding) SetMouseActionsExpanded(value ISynEditMouseActions) {
	if !m.IsValid() {
		return
	}
	synGutterCodeFoldingAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynGutterCodeFolding) MouseActionsCollapsed() ISynEditMouseActions {
	if !m.IsValid() {
		return nil
	}
	r := synGutterCodeFoldingAPI().SysCallN(4, 0, m.Instance())
	return AsSynEditMouseActions(r)
}

func (m *TSynGutterCodeFolding) SetMouseActionsCollapsed(value ISynEditMouseActions) {
	if !m.IsValid() {
		return
	}
	synGutterCodeFoldingAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynGutterCodeFolding) ReversePopMenuOrder() bool {
	if !m.IsValid() {
		return false
	}
	r := synGutterCodeFoldingAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynGutterCodeFolding) SetReversePopMenuOrder(value bool) {
	if !m.IsValid() {
		return
	}
	synGutterCodeFoldingAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

// NewSynGutterCodeFolding class constructor
func NewSynGutterCodeFolding(owner IComponent) ISynGutterCodeFolding {
	r := synGutterCodeFoldingAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynGutterCodeFolding(r)
}

func TSynGutterCodeFoldingClass() types.TClass {
	r := synGutterCodeFoldingAPI().SysCallN(6)
	return types.TClass(r)
}

var (
	synGutterCodeFoldingOnce   base.Once
	synGutterCodeFoldingImport *imports.Imports = nil
)

func synGutterCodeFoldingAPI() *imports.Imports {
	synGutterCodeFoldingOnce.Do(func() {
		synGutterCodeFoldingImport = api.NewDefaultImports()
		synGutterCodeFoldingImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterCodeFolding_Create", 0), // constructor NewSynGutterCodeFolding
			/* 1 */ imports.NewTable("TSynGutterCodeFolding_PaintWithCanvasRectIntX2", 0), // procedure PaintWithCanvasRectIntX2
			/* 2 */ imports.NewTable("TSynGutterCodeFolding_MarkupInfoCurrentFold", 0), // property MarkupInfoCurrentFold
			/* 3 */ imports.NewTable("TSynGutterCodeFolding_MouseActionsExpanded", 0), // property MouseActionsExpanded
			/* 4 */ imports.NewTable("TSynGutterCodeFolding_MouseActionsCollapsed", 0), // property MouseActionsCollapsed
			/* 5 */ imports.NewTable("TSynGutterCodeFolding_ReversePopMenuOrder", 0), // property ReversePopMenuOrder
			/* 6 */ imports.NewTable("TSynGutterCodeFolding_TClass", 0), // function TSynGutterCodeFoldingClass
		}
	})
	return synGutterCodeFoldingImport
}
