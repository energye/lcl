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

// ISynEditMarkLineList Parent: ISynSizedDifferentialAVLTree
type ISynEditMarkLineList interface {
	ISynSizedDifferentialAVLTree
	GetOrAddLine(lineNum int32) ISynEditMarkLine  // function
	GetLineOrNext(lineNum int32) ISynEditMarkLine // function
	AddMark(item ISynEditMark) int32              // function
	RemoveMark(item ISynEditMark) int32           // function
	IndexOfMark(mark ISynEditMark) int32          // function
	Remove(item ISynEditMarkLine, freeMarks bool) // procedure
	ClearWithBool(freeMarks bool)                 // procedure
	DeleteMark(index int32)                       // procedure
	Lines(lineNum int32) ISynEditMarkLine         // property Lines Getter
	MarkCount() int32                             // property MarkCount Getter
	Mark(index int32) ISynEditMark                // property Mark Getter
	SetMark(index int32, value ISynEditMark)      // property Mark Setter
}

type TSynEditMarkLineList struct {
	TSynSizedDifferentialAVLTree
}

func (m *TSynEditMarkLineList) GetOrAddLine(lineNum int32) ISynEditMarkLine {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkLineListAPI().SysCallN(1, m.Instance(), uintptr(lineNum))
	return AsSynEditMarkLine(r)
}

func (m *TSynEditMarkLineList) GetLineOrNext(lineNum int32) ISynEditMarkLine {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkLineListAPI().SysCallN(2, m.Instance(), uintptr(lineNum))
	return AsSynEditMarkLine(r)
}

func (m *TSynEditMarkLineList) AddMark(item ISynEditMark) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineListAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(item))
	return int32(r)
}

func (m *TSynEditMarkLineList) RemoveMark(item ISynEditMark) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineListAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(item))
	return int32(r)
}

func (m *TSynEditMarkLineList) IndexOfMark(mark ISynEditMark) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineListAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(mark))
	return int32(r)
}

func (m *TSynEditMarkLineList) Remove(item ISynEditMarkLine, freeMarks bool) {
	if !m.IsValid() {
		return
	}
	synEditMarkLineListAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(item), api.PasBool(freeMarks))
}

func (m *TSynEditMarkLineList) ClearWithBool(freeMarks bool) {
	if !m.IsValid() {
		return
	}
	synEditMarkLineListAPI().SysCallN(7, m.Instance(), api.PasBool(freeMarks))
}

func (m *TSynEditMarkLineList) DeleteMark(index int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkLineListAPI().SysCallN(8, m.Instance(), uintptr(index))
}

func (m *TSynEditMarkLineList) Lines(lineNum int32) ISynEditMarkLine {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkLineListAPI().SysCallN(9, m.Instance(), uintptr(lineNum))
	return AsSynEditMarkLine(r)
}

func (m *TSynEditMarkLineList) MarkCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineListAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TSynEditMarkLineList) Mark(index int32) ISynEditMark {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkLineListAPI().SysCallN(11, 0, m.Instance(), uintptr(index))
	return AsSynEditMark(r)
}

func (m *TSynEditMarkLineList) SetMark(index int32, value ISynEditMark) {
	if !m.IsValid() {
		return
	}
	synEditMarkLineListAPI().SysCallN(11, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

// NewSynEditMarkLineList class constructor
func NewSynEditMarkLineList(owner ISynEditMarkList) ISynEditMarkLineList {
	r := synEditMarkLineListAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynEditMarkLineList(r)
}

var (
	synEditMarkLineListOnce   base.Once
	synEditMarkLineListImport *imports.Imports = nil
)

func synEditMarkLineListAPI() *imports.Imports {
	synEditMarkLineListOnce.Do(func() {
		synEditMarkLineListImport = api.NewDefaultImports()
		synEditMarkLineListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditMarkLineList_Create", 0), // constructor NewSynEditMarkLineList
			/* 1 */ imports.NewTable("TSynEditMarkLineList_GetOrAddLine", 0), // function GetOrAddLine
			/* 2 */ imports.NewTable("TSynEditMarkLineList_GetLineOrNext", 0), // function GetLineOrNext
			/* 3 */ imports.NewTable("TSynEditMarkLineList_AddMark", 0), // function AddMark
			/* 4 */ imports.NewTable("TSynEditMarkLineList_RemoveMark", 0), // function RemoveMark
			/* 5 */ imports.NewTable("TSynEditMarkLineList_IndexOfMark", 0), // function IndexOfMark
			/* 6 */ imports.NewTable("TSynEditMarkLineList_Remove", 0), // procedure Remove
			/* 7 */ imports.NewTable("TSynEditMarkLineList_ClearWithBool", 0), // procedure ClearWithBool
			/* 8 */ imports.NewTable("TSynEditMarkLineList_DeleteMark", 0), // procedure DeleteMark
			/* 9 */ imports.NewTable("TSynEditMarkLineList_Lines", 0), // property Lines
			/* 10 */ imports.NewTable("TSynEditMarkLineList_MarkCount", 0), // property MarkCount
			/* 11 */ imports.NewTable("TSynEditMarkLineList_Mark", 0), // property Mark
		}
	})
	return synEditMarkLineListImport
}
