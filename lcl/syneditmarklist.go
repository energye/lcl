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
)

// ISynEditMarkList Parent: IObject
type ISynEditMarkList interface {
	IObject
	Remove(item ISynEditMark) int32           // function
	IndexOf(item ISynEditMark) int32          // function
	Count() int32                             // function
	Add(item ISynEditMark)                    // procedure
	Delete(index int32)                       // procedure
	ClearLine(line int32)                     // procedure
	Items(index int32) ISynEditMark           // property Items Getter
	SetItems(index int32, value ISynEditMark) // property Items Setter
	Line(lineNum int32) ISynEditMarkLine      // property Line Getter
	SetOnChange(fn TNotifyEvent)              // property event
}

type TSynEditMarkList struct {
	TObject
}

func (m *TSynEditMarkList) Remove(item ISynEditMark) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkListAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(item))
	return int32(r)
}

func (m *TSynEditMarkList) IndexOf(item ISynEditMark) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkListAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(item))
	return int32(r)
}

func (m *TSynEditMarkList) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkListAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TSynEditMarkList) Add(item ISynEditMark) {
	if !m.IsValid() {
		return
	}
	synEditMarkListAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(item))
}

func (m *TSynEditMarkList) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkListAPI().SysCallN(5, m.Instance(), uintptr(index))
}

func (m *TSynEditMarkList) ClearLine(line int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkListAPI().SysCallN(6, m.Instance(), uintptr(line))
}

func (m *TSynEditMarkList) Items(index int32) ISynEditMark {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkListAPI().SysCallN(7, 0, m.Instance(), uintptr(index))
	return AsSynEditMark(r)
}

func (m *TSynEditMarkList) SetItems(index int32, value ISynEditMark) {
	if !m.IsValid() {
		return
	}
	synEditMarkListAPI().SysCallN(7, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TSynEditMarkList) Line(lineNum int32) ISynEditMarkLine {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkListAPI().SysCallN(8, m.Instance(), uintptr(lineNum))
	return AsSynEditMarkLine(r)
}

func (m *TSynEditMarkList) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, synEditMarkListAPI(), api.MakeEventDataPtr(cb))
}

// NewSynEditMarkList class constructor
func NewSynEditMarkList(owner ISynEditBase, lines ISynEditStringsLinked) ISynEditMarkList {
	r := synEditMarkListAPI().SysCallN(0, base.GetObjectUintptr(owner), base.GetObjectUintptr(lines))
	return AsSynEditMarkList(r)
}

var (
	synEditMarkListOnce   base.Once
	synEditMarkListImport *imports.Imports = nil
)

func synEditMarkListAPI() *imports.Imports {
	synEditMarkListOnce.Do(func() {
		synEditMarkListImport = api.NewDefaultImports()
		synEditMarkListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditMarkList_Create", 0), // constructor NewSynEditMarkList
			/* 1 */ imports.NewTable("TSynEditMarkList_Remove", 0), // function Remove
			/* 2 */ imports.NewTable("TSynEditMarkList_IndexOf", 0), // function IndexOf
			/* 3 */ imports.NewTable("TSynEditMarkList_Count", 0), // function Count
			/* 4 */ imports.NewTable("TSynEditMarkList_Add", 0), // procedure Add
			/* 5 */ imports.NewTable("TSynEditMarkList_Delete", 0), // procedure Delete
			/* 6 */ imports.NewTable("TSynEditMarkList_ClearLine", 0), // procedure ClearLine
			/* 7 */ imports.NewTable("TSynEditMarkList_Items", 0), // property Items
			/* 8 */ imports.NewTable("TSynEditMarkList_Line", 0), // property Line
			/* 9 */ imports.NewTable("TSynEditMarkList_OnChange", 0), // event OnChange
		}
	})
	return synEditMarkListImport
}
