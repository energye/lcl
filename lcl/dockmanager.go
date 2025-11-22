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

// IDockManager Parent: IPersistent
type IDockManager interface {
	IPersistent
	GetDockEdge(dockObject IDragDockObject) bool                                                                             // function
	AutoFreeByControl() bool                                                                                                 // function
	IsEnabledControl(control IControl) bool                                                                                  // function
	CanBeDoubleDocked() bool                                                                                                 // function
	BeginUpdate()                                                                                                            // procedure
	EndUpdate()                                                                                                              // procedure
	GetControlBounds(control IControl, outControlBounds *types.TRect)                                                        // procedure
	InsertControlWithDragDockObject(dockObject IDragDockObject)                                                              // procedure
	InsertControlWithControlX2Align(control IControl, insertAt types.TAlign, dropCtl IControl)                               // procedure
	LoadFromStream(stream IStream)                                                                                           // procedure
	PaintSite(dC types.HDC)                                                                                                  // procedure
	MessageHandler(sender IControl, message *types.TLMessage)                                                                // procedure
	PositionDockRectWithDragDockObject(dockObject IDragDockObject)                                                           // procedure
	PositionDockRectWithControlX2AlignRect(client IControl, dropCtl IControl, dropAlign types.TAlign, dockRect *types.TRect) // procedure
	RemoveControl(control IControl)                                                                                          // procedure
	ResetBounds(force bool)                                                                                                  // procedure
	SaveToStream(stream IStream)                                                                                             // procedure
	SetReplacingControl(control IControl)                                                                                    // procedure
}

type TDockManager struct {
	TPersistent
}

func (m *TDockManager) GetDockEdge(dockObject IDragDockObject) bool {
	if !m.IsValid() {
		return false
	}
	r := dockManagerAPI().SysCallN(0, m.Instance(), base.GetObjectUintptr(dockObject))
	return api.GoBool(r)
}

func (m *TDockManager) AutoFreeByControl() bool {
	if !m.IsValid() {
		return false
	}
	r := dockManagerAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TDockManager) IsEnabledControl(control IControl) bool {
	if !m.IsValid() {
		return false
	}
	r := dockManagerAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(control))
	return api.GoBool(r)
}

func (m *TDockManager) CanBeDoubleDocked() bool {
	if !m.IsValid() {
		return false
	}
	r := dockManagerAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TDockManager) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(4, m.Instance())
}

func (m *TDockManager) EndUpdate() {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(5, m.Instance())
}

func (m *TDockManager) GetControlBounds(control IControl, outControlBounds *types.TRect) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(control), uintptr(base.UnsafePointer(outControlBounds)))
}

func (m *TDockManager) InsertControlWithDragDockObject(dockObject IDragDockObject) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(dockObject))
}

func (m *TDockManager) InsertControlWithControlX2Align(control IControl, insertAt types.TAlign, dropCtl IControl) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(control), uintptr(insertAt), base.GetObjectUintptr(dropCtl))
}

func (m *TDockManager) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TDockManager) PaintSite(dC types.HDC) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(10, m.Instance(), uintptr(dC))
}

func (m *TDockManager) MessageHandler(sender IControl, message *types.TLMessage) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(sender), uintptr(base.UnsafePointer(message)))
}

func (m *TDockManager) PositionDockRectWithDragDockObject(dockObject IDragDockObject) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(dockObject))
}

func (m *TDockManager) PositionDockRectWithControlX2AlignRect(client IControl, dropCtl IControl, dropAlign types.TAlign, dockRect *types.TRect) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(client), base.GetObjectUintptr(dropCtl), uintptr(dropAlign), uintptr(base.UnsafePointer(dockRect)))
}

func (m *TDockManager) RemoveControl(control IControl) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(control))
}

func (m *TDockManager) ResetBounds(force bool) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(15, m.Instance(), api.PasBool(force))
}

func (m *TDockManager) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TDockManager) SetReplacingControl(control IControl) {
	if !m.IsValid() {
		return
	}
	dockManagerAPI().SysCallN(17, m.Instance(), base.GetObjectUintptr(control))
}

var (
	dockManagerOnce   base.Once
	dockManagerImport *imports.Imports = nil
)

func dockManagerAPI() *imports.Imports {
	dockManagerOnce.Do(func() {
		dockManagerImport = api.NewDefaultImports()
		dockManagerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDockManager_GetDockEdge", 0), // function GetDockEdge
			/* 1 */ imports.NewTable("TDockManager_AutoFreeByControl", 0), // function AutoFreeByControl
			/* 2 */ imports.NewTable("TDockManager_IsEnabledControl", 0), // function IsEnabledControl
			/* 3 */ imports.NewTable("TDockManager_CanBeDoubleDocked", 0), // function CanBeDoubleDocked
			/* 4 */ imports.NewTable("TDockManager_BeginUpdate", 0), // procedure BeginUpdate
			/* 5 */ imports.NewTable("TDockManager_EndUpdate", 0), // procedure EndUpdate
			/* 6 */ imports.NewTable("TDockManager_GetControlBounds", 0), // procedure GetControlBounds
			/* 7 */ imports.NewTable("TDockManager_InsertControlWithDragDockObject", 0), // procedure InsertControlWithDragDockObject
			/* 8 */ imports.NewTable("TDockManager_InsertControlWithControlX2Align", 0), // procedure InsertControlWithControlX2Align
			/* 9 */ imports.NewTable("TDockManager_LoadFromStream", 0), // procedure LoadFromStream
			/* 10 */ imports.NewTable("TDockManager_PaintSite", 0), // procedure PaintSite
			/* 11 */ imports.NewTable("TDockManager_MessageHandler", 0), // procedure MessageHandler
			/* 12 */ imports.NewTable("TDockManager_PositionDockRectWithDragDockObject", 0), // procedure PositionDockRectWithDragDockObject
			/* 13 */ imports.NewTable("TDockManager_PositionDockRectWithControlX2AlignRect", 0), // procedure PositionDockRectWithControlX2AlignRect
			/* 14 */ imports.NewTable("TDockManager_RemoveControl", 0), // procedure RemoveControl
			/* 15 */ imports.NewTable("TDockManager_ResetBounds", 0), // procedure ResetBounds
			/* 16 */ imports.NewTable("TDockManager_SaveToStream", 0), // procedure SaveToStream
			/* 17 */ imports.NewTable("TDockManager_SetReplacingControl", 0), // procedure SetReplacingControl
		}
	})
	return dockManagerImport
}
