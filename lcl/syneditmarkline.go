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

// ISynEditMarkLine Parent: ISynSizedDifferentialAVLNode
type ISynEditMarkLine interface {
	ISynSizedDifferentialAVLNode
	Add(item ISynEditMark) int32                                                               // function
	Remove(item ISynEditMark) int32                                                            // function
	Count() int32                                                                              // function
	VisibleCount() int32                                                                       // function
	IndexOf(mark ISynEditMark) int32                                                           // function
	Sort(primaryOrder types.TSynEditMarkSortOrder, secondaryOrder types.TSynEditMarkSortOrder) // procedure
	Delete(index int32)                                                                        // procedure
	Clear(freeMarks bool)                                                                      // procedure
	Items(index int32) ISynEditMark                                                            // property Items Getter
	SetItems(index int32, value ISynEditMark)                                                  // property Items Setter
	LineNum() int32                                                                            // property LineNum Getter
}

type TSynEditMarkLine struct {
	TSynSizedDifferentialAVLNode
}

func (m *TSynEditMarkLine) Add(item ISynEditMark) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(item))
	return int32(r)
}

func (m *TSynEditMarkLine) Remove(item ISynEditMark) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(item))
	return int32(r)
}

func (m *TSynEditMarkLine) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TSynEditMarkLine) VisibleCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TSynEditMarkLine) IndexOf(mark ISynEditMark) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(mark))
	return int32(r)
}

func (m *TSynEditMarkLine) Sort(primaryOrder types.TSynEditMarkSortOrder, secondaryOrder types.TSynEditMarkSortOrder) {
	if !m.IsValid() {
		return
	}
	synEditMarkLineAPI().SysCallN(6, m.Instance(), uintptr(primaryOrder), uintptr(secondaryOrder))
}

func (m *TSynEditMarkLine) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	synEditMarkLineAPI().SysCallN(7, m.Instance(), uintptr(index))
}

func (m *TSynEditMarkLine) Clear(freeMarks bool) {
	if !m.IsValid() {
		return
	}
	synEditMarkLineAPI().SysCallN(8, m.Instance(), api.PasBool(freeMarks))
}

func (m *TSynEditMarkLine) Items(index int32) ISynEditMark {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkLineAPI().SysCallN(9, 0, m.Instance(), uintptr(index))
	return AsSynEditMark(r)
}

func (m *TSynEditMarkLine) SetItems(index int32, value ISynEditMark) {
	if !m.IsValid() {
		return
	}
	synEditMarkLineAPI().SysCallN(9, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TSynEditMarkLine) LineNum() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkLineAPI().SysCallN(10, m.Instance())
	return int32(r)
}

// NewSynEditMarkLine class constructor
func NewSynEditMarkLine(lineNum int32, owner ISynEditMarkLineList) ISynEditMarkLine {
	r := synEditMarkLineAPI().SysCallN(0, uintptr(lineNum), base.GetObjectUintptr(owner))
	return AsSynEditMarkLine(r)
}

var (
	synEditMarkLineOnce   base.Once
	synEditMarkLineImport *imports.Imports = nil
)

func synEditMarkLineAPI() *imports.Imports {
	synEditMarkLineOnce.Do(func() {
		synEditMarkLineImport = api.NewDefaultImports()
		synEditMarkLineImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditMarkLine_Create", 0), // constructor NewSynEditMarkLine
			/* 1 */ imports.NewTable("TSynEditMarkLine_Add", 0), // function Add
			/* 2 */ imports.NewTable("TSynEditMarkLine_Remove", 0), // function Remove
			/* 3 */ imports.NewTable("TSynEditMarkLine_Count", 0), // function Count
			/* 4 */ imports.NewTable("TSynEditMarkLine_VisibleCount", 0), // function VisibleCount
			/* 5 */ imports.NewTable("TSynEditMarkLine_IndexOf", 0), // function IndexOf
			/* 6 */ imports.NewTable("TSynEditMarkLine_Sort", 0), // procedure Sort
			/* 7 */ imports.NewTable("TSynEditMarkLine_Delete", 0), // procedure Delete
			/* 8 */ imports.NewTable("TSynEditMarkLine_Clear", 0), // procedure Clear
			/* 9 */ imports.NewTable("TSynEditMarkLine_Items", 0), // property Items
			/* 10 */ imports.NewTable("TSynEditMarkLine_LineNum", 0), // property LineNum
		}
	})
	return synEditMarkLineImport
}
