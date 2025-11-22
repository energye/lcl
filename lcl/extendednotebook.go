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

// IExtendedNotebook Parent: IPageControl
type IExtendedNotebook interface {
	IPageControl
	BeginDragTab(tabIndex int32, immediate bool, threshold int32) // procedure
	DraggingTabIndex() int32                                      // property DraggingTabIndex Getter
	TabDragMode() types.TDragMode                                 // property TabDragMode Getter
	SetTabDragMode(value types.TDragMode)                         // property TabDragMode Setter
	TabDragAcceptMode() types.TDragMode                           // property TabDragAcceptMode Getter
	SetTabDragAcceptMode(value types.TDragMode)                   // property TabDragAcceptMode Setter
	SetOnTabDragOver(fn TDragOverEvent)                           // property event
	SetOnTabDragOverEx(fn TNotebookTabDragOverEvent)              // property event
	SetOnTabDragDrop(fn TDragDropEvent)                           // property event
	SetOnTabDragDropEx(fn TNotebookTabDragDropEvent)              // property event
	SetOnTabEndDrag(fn TEndDragEvent)                             // property event
	SetOnTabStartDrag(fn TStartDragEvent)                         // property event
}

type TExtendedNotebook struct {
	TPageControl
}

func (m *TExtendedNotebook) BeginDragTab(tabIndex int32, immediate bool, threshold int32) {
	if !m.IsValid() {
		return
	}
	extendedNotebookAPI().SysCallN(1, m.Instance(), uintptr(tabIndex), api.PasBool(immediate), uintptr(threshold))
}

func (m *TExtendedNotebook) DraggingTabIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := extendedNotebookAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TExtendedNotebook) TabDragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := extendedNotebookAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TExtendedNotebook) SetTabDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	extendedNotebookAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TExtendedNotebook) TabDragAcceptMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := extendedNotebookAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TExtendedNotebook) SetTabDragAcceptMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	extendedNotebookAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TExtendedNotebook) SetOnTabDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 5, extendedNotebookAPI(), api.MakeEventDataPtr(cb))
}

func (m *TExtendedNotebook) SetOnTabDragOverEx(fn TNotebookTabDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotebookTabDragOverEvent(fn)
	base.SetEvent(m, 6, extendedNotebookAPI(), api.MakeEventDataPtr(cb))
}

func (m *TExtendedNotebook) SetOnTabDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 7, extendedNotebookAPI(), api.MakeEventDataPtr(cb))
}

func (m *TExtendedNotebook) SetOnTabDragDropEx(fn TNotebookTabDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotebookTabDragDropEvent(fn)
	base.SetEvent(m, 8, extendedNotebookAPI(), api.MakeEventDataPtr(cb))
}

func (m *TExtendedNotebook) SetOnTabEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 9, extendedNotebookAPI(), api.MakeEventDataPtr(cb))
}

func (m *TExtendedNotebook) SetOnTabStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 10, extendedNotebookAPI(), api.MakeEventDataPtr(cb))
}

// NewExtendedNotebook class constructor
func NewExtendedNotebook(theOwner IComponent) IExtendedNotebook {
	r := extendedNotebookAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsExtendedNotebook(r)
}

func TExtendedNotebookClass() types.TClass {
	r := extendedNotebookAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	extendedNotebookOnce   base.Once
	extendedNotebookImport *imports.Imports = nil
)

func extendedNotebookAPI() *imports.Imports {
	extendedNotebookOnce.Do(func() {
		extendedNotebookImport = api.NewDefaultImports()
		extendedNotebookImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TExtendedNotebook_Create", 0), // constructor NewExtendedNotebook
			/* 1 */ imports.NewTable("TExtendedNotebook_BeginDragTab", 0), // procedure BeginDragTab
			/* 2 */ imports.NewTable("TExtendedNotebook_DraggingTabIndex", 0), // property DraggingTabIndex
			/* 3 */ imports.NewTable("TExtendedNotebook_TabDragMode", 0), // property TabDragMode
			/* 4 */ imports.NewTable("TExtendedNotebook_TabDragAcceptMode", 0), // property TabDragAcceptMode
			/* 5 */ imports.NewTable("TExtendedNotebook_OnTabDragOver", 0), // event OnTabDragOver
			/* 6 */ imports.NewTable("TExtendedNotebook_OnTabDragOverEx", 0), // event OnTabDragOverEx
			/* 7 */ imports.NewTable("TExtendedNotebook_OnTabDragDrop", 0), // event OnTabDragDrop
			/* 8 */ imports.NewTable("TExtendedNotebook_OnTabDragDropEx", 0), // event OnTabDragDropEx
			/* 9 */ imports.NewTable("TExtendedNotebook_OnTabEndDrag", 0), // event OnTabEndDrag
			/* 10 */ imports.NewTable("TExtendedNotebook_OnTabStartDrag", 0), // event OnTabStartDrag
			/* 11 */ imports.NewTable("TExtendedNotebook_TClass", 0), // function TExtendedNotebookClass
		}
	})
	return extendedNotebookImport
}
