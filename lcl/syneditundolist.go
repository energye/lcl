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

// ISynEditUndoList Parent: IObject
type ISynEditUndoList interface {
	IObject
	GetLastChange() ISynEditUndoItem                     // function
	PopLastChange() ISynEditUndoItem                     // function
	PopItem() ISynEditUndoGroup                          // function
	PeekItem() ISynEditUndoGroup                         // function
	IsLocked() bool                                      // function
	RealCount() int32                                    // function
	IsTopMarkedAsUnmodified() bool                       // function
	UnModifiedMarkerExists() bool                        // function
	IsTopMarkedAsSaved() bool                            // function
	SavedMarkerExists() bool                             // function
	AddChange(change ISynEditUndoItem, forceLocked bool) // procedure
	// AppendToLastChange
	//  "LastChange: Either in current Group, or in last Group, if no current
	AppendToLastChange(change ISynEditUndoItem) // procedure
	BeginBlock()                                // procedure
	EndBlock()                                  // procedure
	Clear()                                     // procedure
	Lock()                                      // procedure
	Unlock()                                    // procedure
	// MarkTopAsUnmodified
	//  * Historically SynEdit has
	//  TSynEdit.MarkTextAsSaved; // Affects the "Changes Gutter" only
	//  TSynEdit.Modified := False;
	MarkTopAsUnmodified()                           // procedure
	MarkTopAsSaved()                                // procedure
	ForceGroupEnd()                                 // procedure
	CanUndo() bool                                  // property CanUndo Getter
	FullUndoImpossible() bool                       // property FullUndoImpossible Getter
	ItemCount() int32                               // property ItemCount Getter
	MaxUndoActions() int32                          // property MaxUndoActions Getter
	SetMaxUndoActions(value int32)                  // property MaxUndoActions Setter
	IsInsideRedo() bool                             // property IsInsideRedo Getter
	SetIsInsideRedo(value bool)                     // property IsInsideRedo Setter
	GroupUndo() bool                                // property GroupUndo Getter
	SetGroupUndo(value bool)                        // property GroupUndo Setter
	CurrentGroup() ISynEditUndoGroup                // property CurrentGroup Getter
	CurrentReason() types.TSynEditorCommand         // property CurrentReason Getter
	SetCurrentReason(value types.TSynEditorCommand) // property CurrentReason Setter
	SetOnAddedUndo(fn TNotifyEvent)                 // property event
	SetOnNeedCaretUndo(fn TSynGetCaretUndoProc)     // property event
}

type TSynEditUndoList struct {
	TObject
}

func (m *TSynEditUndoList) GetLastChange() ISynEditUndoItem {
	if !m.IsValid() {
		return nil
	}
	r := synEditUndoListAPI().SysCallN(1, m.Instance())
	return AsSynEditUndoItem(r)
}

func (m *TSynEditUndoList) PopLastChange() ISynEditUndoItem {
	if !m.IsValid() {
		return nil
	}
	r := synEditUndoListAPI().SysCallN(2, m.Instance())
	return AsSynEditUndoItem(r)
}

func (m *TSynEditUndoList) PopItem() ISynEditUndoGroup {
	if !m.IsValid() {
		return nil
	}
	r := synEditUndoListAPI().SysCallN(3, m.Instance())
	return AsSynEditUndoGroup(r)
}

func (m *TSynEditUndoList) PeekItem() ISynEditUndoGroup {
	if !m.IsValid() {
		return nil
	}
	r := synEditUndoListAPI().SysCallN(4, m.Instance())
	return AsSynEditUndoGroup(r)
}

func (m *TSynEditUndoList) IsLocked() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoListAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoList) RealCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditUndoListAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TSynEditUndoList) IsTopMarkedAsUnmodified() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoListAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoList) UnModifiedMarkerExists() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoListAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoList) IsTopMarkedAsSaved() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoListAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoList) SavedMarkerExists() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoListAPI().SysCallN(10, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoList) AddChange(change ISynEditUndoItem, forceLocked bool) {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(change), api.PasBool(forceLocked))
}

func (m *TSynEditUndoList) AppendToLastChange(change ISynEditUndoItem) {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(change))
}

func (m *TSynEditUndoList) BeginBlock() {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(13, m.Instance())
}

func (m *TSynEditUndoList) EndBlock() {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(14, m.Instance())
}

func (m *TSynEditUndoList) Clear() {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(15, m.Instance())
}

func (m *TSynEditUndoList) Lock() {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(16, m.Instance())
}

func (m *TSynEditUndoList) Unlock() {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(17, m.Instance())
}

func (m *TSynEditUndoList) MarkTopAsUnmodified() {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(18, m.Instance())
}

func (m *TSynEditUndoList) MarkTopAsSaved() {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(19, m.Instance())
}

func (m *TSynEditUndoList) ForceGroupEnd() {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(20, m.Instance())
}

func (m *TSynEditUndoList) CanUndo() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoListAPI().SysCallN(21, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoList) FullUndoImpossible() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoListAPI().SysCallN(22, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoList) ItemCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditUndoListAPI().SysCallN(23, m.Instance())
	return int32(r)
}

func (m *TSynEditUndoList) MaxUndoActions() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditUndoListAPI().SysCallN(24, 0, m.Instance())
	return int32(r)
}

func (m *TSynEditUndoList) SetMaxUndoActions(value int32) {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditUndoList) IsInsideRedo() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoListAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoList) SetIsInsideRedo(value bool) {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditUndoList) GroupUndo() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditUndoListAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEditUndoList) SetGroupUndo(value bool) {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEditUndoList) CurrentGroup() ISynEditUndoGroup {
	if !m.IsValid() {
		return nil
	}
	r := synEditUndoListAPI().SysCallN(27, m.Instance())
	return AsSynEditUndoGroup(r)
}

func (m *TSynEditUndoList) CurrentReason() types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	r := synEditUndoListAPI().SysCallN(28, 0, m.Instance())
	return types.TSynEditorCommand(r)
}

func (m *TSynEditUndoList) SetCurrentReason(value types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	synEditUndoListAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TSynEditUndoList) SetOnAddedUndo(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 29, synEditUndoListAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEditUndoList) SetOnNeedCaretUndo(fn TSynGetCaretUndoProc) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynGetCaretUndoProc(fn)
	base.SetEvent(m, 30, synEditUndoListAPI(), api.MakeEventDataPtr(cb))
}

// NewSynEditUndoList class constructor
func NewSynEditUndoList() ISynEditUndoList {
	r := synEditUndoListAPI().SysCallN(0)
	return AsSynEditUndoList(r)
}

var (
	synEditUndoListOnce   base.Once
	synEditUndoListImport *imports.Imports = nil
)

func synEditUndoListAPI() *imports.Imports {
	synEditUndoListOnce.Do(func() {
		synEditUndoListImport = api.NewDefaultImports()
		synEditUndoListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditUndoList_Create", 0), // constructor NewSynEditUndoList
			/* 1 */ imports.NewTable("TSynEditUndoList_GetLastChange", 0), // function GetLastChange
			/* 2 */ imports.NewTable("TSynEditUndoList_PopLastChange", 0), // function PopLastChange
			/* 3 */ imports.NewTable("TSynEditUndoList_PopItem", 0), // function PopItem
			/* 4 */ imports.NewTable("TSynEditUndoList_PeekItem", 0), // function PeekItem
			/* 5 */ imports.NewTable("TSynEditUndoList_IsLocked", 0), // function IsLocked
			/* 6 */ imports.NewTable("TSynEditUndoList_RealCount", 0), // function RealCount
			/* 7 */ imports.NewTable("TSynEditUndoList_IsTopMarkedAsUnmodified", 0), // function IsTopMarkedAsUnmodified
			/* 8 */ imports.NewTable("TSynEditUndoList_UnModifiedMarkerExists", 0), // function UnModifiedMarkerExists
			/* 9 */ imports.NewTable("TSynEditUndoList_IsTopMarkedAsSaved", 0), // function IsTopMarkedAsSaved
			/* 10 */ imports.NewTable("TSynEditUndoList_SavedMarkerExists", 0), // function SavedMarkerExists
			/* 11 */ imports.NewTable("TSynEditUndoList_AddChange", 0), // procedure AddChange
			/* 12 */ imports.NewTable("TSynEditUndoList_AppendToLastChange", 0), // procedure AppendToLastChange
			/* 13 */ imports.NewTable("TSynEditUndoList_BeginBlock", 0), // procedure BeginBlock
			/* 14 */ imports.NewTable("TSynEditUndoList_EndBlock", 0), // procedure EndBlock
			/* 15 */ imports.NewTable("TSynEditUndoList_Clear", 0), // procedure Clear
			/* 16 */ imports.NewTable("TSynEditUndoList_Lock", 0), // procedure Lock
			/* 17 */ imports.NewTable("TSynEditUndoList_Unlock", 0), // procedure Unlock
			/* 18 */ imports.NewTable("TSynEditUndoList_MarkTopAsUnmodified", 0), // procedure MarkTopAsUnmodified
			/* 19 */ imports.NewTable("TSynEditUndoList_MarkTopAsSaved", 0), // procedure MarkTopAsSaved
			/* 20 */ imports.NewTable("TSynEditUndoList_ForceGroupEnd", 0), // procedure ForceGroupEnd
			/* 21 */ imports.NewTable("TSynEditUndoList_CanUndo", 0), // property CanUndo
			/* 22 */ imports.NewTable("TSynEditUndoList_FullUndoImpossible", 0), // property FullUndoImpossible
			/* 23 */ imports.NewTable("TSynEditUndoList_ItemCount", 0), // property ItemCount
			/* 24 */ imports.NewTable("TSynEditUndoList_MaxUndoActions", 0), // property MaxUndoActions
			/* 25 */ imports.NewTable("TSynEditUndoList_IsInsideRedo", 0), // property IsInsideRedo
			/* 26 */ imports.NewTable("TSynEditUndoList_GroupUndo", 0), // property GroupUndo
			/* 27 */ imports.NewTable("TSynEditUndoList_CurrentGroup", 0), // property CurrentGroup
			/* 28 */ imports.NewTable("TSynEditUndoList_CurrentReason", 0), // property CurrentReason
			/* 29 */ imports.NewTable("TSynEditUndoList_OnAddedUndo", 0), // event OnAddedUndo
			/* 30 */ imports.NewTable("TSynEditUndoList_OnNeedCaretUndo", 0), // event OnNeedCaretUndo
		}
	})
	return synEditUndoListImport
}
