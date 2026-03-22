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

// ISynHighlighterRangeList Parent: ISynManagedStorageMem
type ISynHighlighterRangeList interface {
	ISynManagedStorageMem
	ClearReScanNeeded()                  // procedure
	AdjustReScanStart(newStart int32)    // procedure
	InvalidateAll()                      // procedure
	IncRefCount()                        // procedure
	DecRefCount()                        // procedure
	Range(index int32) uintptr           // property Range Getter
	SetRange(index int32, value uintptr) // property Range Setter
	RefCount() int32                     // property RefCount Getter
	NeedsReScanStartIndex() int32        // property NeedsReScanStartIndex Getter
	NeedsReScanEndIndex() int32          // property NeedsReScanEndIndex Getter
	NeedsReScanRealStartIndex() int32    // property NeedsReScanRealStartIndex Getter
}

type TSynHighlighterRangeList struct {
	TSynManagedStorageMem
}

func (m *TSynHighlighterRangeList) ClearReScanNeeded() {
	if !m.IsValid() {
		return
	}
	synHighlighterRangeListAPI().SysCallN(1, m.Instance())
}

func (m *TSynHighlighterRangeList) AdjustReScanStart(newStart int32) {
	if !m.IsValid() {
		return
	}
	synHighlighterRangeListAPI().SysCallN(2, m.Instance(), uintptr(newStart))
}

func (m *TSynHighlighterRangeList) InvalidateAll() {
	if !m.IsValid() {
		return
	}
	synHighlighterRangeListAPI().SysCallN(3, m.Instance())
}

func (m *TSynHighlighterRangeList) IncRefCount() {
	if !m.IsValid() {
		return
	}
	synHighlighterRangeListAPI().SysCallN(4, m.Instance())
}

func (m *TSynHighlighterRangeList) DecRefCount() {
	if !m.IsValid() {
		return
	}
	synHighlighterRangeListAPI().SysCallN(5, m.Instance())
}

func (m *TSynHighlighterRangeList) Range(index int32) uintptr {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterRangeListAPI().SysCallN(6, 0, m.Instance(), uintptr(index))
	return uintptr(r)
}

func (m *TSynHighlighterRangeList) SetRange(index int32, value uintptr) {
	if !m.IsValid() {
		return
	}
	synHighlighterRangeListAPI().SysCallN(6, 1, m.Instance(), uintptr(index), uintptr(value))
}

func (m *TSynHighlighterRangeList) RefCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterRangeListAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TSynHighlighterRangeList) NeedsReScanStartIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterRangeListAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TSynHighlighterRangeList) NeedsReScanEndIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterRangeListAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TSynHighlighterRangeList) NeedsReScanRealStartIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterRangeListAPI().SysCallN(10, m.Instance())
	return int32(r)
}

// NewSynHighlighterRangeList class constructor
func NewSynHighlighterRangeList() ISynHighlighterRangeList {
	r := synHighlighterRangeListAPI().SysCallN(0)
	return AsSynHighlighterRangeList(r)
}

var (
	synHighlighterRangeListOnce   base.Once
	synHighlighterRangeListImport *imports.Imports = nil
)

func synHighlighterRangeListAPI() *imports.Imports {
	synHighlighterRangeListOnce.Do(func() {
		synHighlighterRangeListImport = api.NewDefaultImports()
		synHighlighterRangeListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynHighlighterRangeList_Create", 0), // constructor NewSynHighlighterRangeList
			/* 1 */ imports.NewTable("TSynHighlighterRangeList_ClearReScanNeeded", 0), // procedure ClearReScanNeeded
			/* 2 */ imports.NewTable("TSynHighlighterRangeList_AdjustReScanStart", 0), // procedure AdjustReScanStart
			/* 3 */ imports.NewTable("TSynHighlighterRangeList_InvalidateAll", 0), // procedure InvalidateAll
			/* 4 */ imports.NewTable("TSynHighlighterRangeList_IncRefCount", 0), // procedure IncRefCount
			/* 5 */ imports.NewTable("TSynHighlighterRangeList_DecRefCount", 0), // procedure DecRefCount
			/* 6 */ imports.NewTable("TSynHighlighterRangeList_Range", 0), // property Range
			/* 7 */ imports.NewTable("TSynHighlighterRangeList_RefCount", 0), // property RefCount
			/* 8 */ imports.NewTable("TSynHighlighterRangeList_NeedsReScanStartIndex", 0), // property NeedsReScanStartIndex
			/* 9 */ imports.NewTable("TSynHighlighterRangeList_NeedsReScanEndIndex", 0), // property NeedsReScanEndIndex
			/* 10 */ imports.NewTable("TSynHighlighterRangeList_NeedsReScanRealStartIndex", 0), // property NeedsReScanRealStartIndex
		}
	})
	return synHighlighterRangeListImport
}
