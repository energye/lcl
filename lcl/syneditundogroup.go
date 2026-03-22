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

// ISynEditUndoGroup Parent: IObject
type ISynEditUndoGroup interface {
	IObject
	Pop() ISynEditUndoItem                       // function
	Assign(anUndoGroup ISynEditUndoGroup)        // procedure
	Add(anItem ISynEditUndoItem)                 // procedure
	Clear(onlyFreeItems bool)                    // procedure
	Insert(index int32, anItem ISynEditUndoItem) // procedure
	Delete(index int32)                          // procedure
	Count() int32                                // property Count Getter
	Items(index int32) ISynEditUndoItem          // property Items Getter
	Reason() types.TSynEditorCommand             // property Reason Getter
	SetReason(value types.TSynEditorCommand)     // property Reason Setter
}

type TSynEditUndoGroup struct {
	TObject
}

func (m *TSynEditUndoGroup) Pop() ISynEditUndoItem {
	if !m.IsValid() {
		return nil
	}
	r := synEditUndoGroupAPI().SysCallN(1, m.Instance())
	return AsSynEditUndoItem(r)
}

func (m *TSynEditUndoGroup) Assign(anUndoGroup ISynEditUndoGroup) {
	if !m.IsValid() {
		return
	}
	synEditUndoGroupAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(anUndoGroup))
}

func (m *TSynEditUndoGroup) Add(anItem ISynEditUndoItem) {
	if !m.IsValid() {
		return
	}
	synEditUndoGroupAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(anItem))
}

func (m *TSynEditUndoGroup) Clear(onlyFreeItems bool) {
	if !m.IsValid() {
		return
	}
	synEditUndoGroupAPI().SysCallN(4, m.Instance(), api.PasBool(onlyFreeItems))
}

func (m *TSynEditUndoGroup) Insert(index int32, anItem ISynEditUndoItem) {
	if !m.IsValid() {
		return
	}
	synEditUndoGroupAPI().SysCallN(5, m.Instance(), uintptr(index), base.GetObjectUintptr(anItem))
}

func (m *TSynEditUndoGroup) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	synEditUndoGroupAPI().SysCallN(6, m.Instance(), uintptr(index))
}

func (m *TSynEditUndoGroup) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditUndoGroupAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TSynEditUndoGroup) Items(index int32) ISynEditUndoItem {
	if !m.IsValid() {
		return nil
	}
	r := synEditUndoGroupAPI().SysCallN(8, m.Instance(), uintptr(index))
	return AsSynEditUndoItem(r)
}

func (m *TSynEditUndoGroup) Reason() types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	r := synEditUndoGroupAPI().SysCallN(9, 0, m.Instance())
	return types.TSynEditorCommand(r)
}

func (m *TSynEditUndoGroup) SetReason(value types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	synEditUndoGroupAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

// NewSynEditUndoGroup class constructor
func NewSynEditUndoGroup() ISynEditUndoGroup {
	r := synEditUndoGroupAPI().SysCallN(0)
	return AsSynEditUndoGroup(r)
}

var (
	synEditUndoGroupOnce   base.Once
	synEditUndoGroupImport *imports.Imports = nil
)

func synEditUndoGroupAPI() *imports.Imports {
	synEditUndoGroupOnce.Do(func() {
		synEditUndoGroupImport = api.NewDefaultImports()
		synEditUndoGroupImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditUndoGroup_Create", 0), // constructor NewSynEditUndoGroup
			/* 1 */ imports.NewTable("TSynEditUndoGroup_Pop", 0), // function Pop
			/* 2 */ imports.NewTable("TSynEditUndoGroup_Assign", 0), // procedure Assign
			/* 3 */ imports.NewTable("TSynEditUndoGroup_Add", 0), // procedure Add
			/* 4 */ imports.NewTable("TSynEditUndoGroup_Clear", 0), // procedure Clear
			/* 5 */ imports.NewTable("TSynEditUndoGroup_Insert", 0), // procedure Insert
			/* 6 */ imports.NewTable("TSynEditUndoGroup_Delete", 0), // procedure Delete
			/* 7 */ imports.NewTable("TSynEditUndoGroup_Count", 0), // property Count
			/* 8 */ imports.NewTable("TSynEditUndoGroup_Items", 0), // property Items
			/* 9 */ imports.NewTable("TSynEditUndoGroup_Reason", 0), // property Reason
		}
	})
	return synEditUndoGroupImport
}
