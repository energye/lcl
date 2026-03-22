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

// ILazSynFoldNodeInfoList Parent: IObject
type ILazSynFoldNodeInfoList interface {
	IObject
	CountAll() int32 // function
	Count() int32    // function
	// CountEx
	//  all items / filtered on the fly
	CountEx(anActionFilter types.TSynFoldActions, groupFilter int32) int32                            // function
	NodeInfoEx(index int32, anActionFilter types.TSynFoldActions, groupFilter int32) TSynFoldNodeInfo // function
	// Add
	//  used by HighLighters to add data
	Add(anInfo TSynFoldNodeInfo) // procedure
	Delete(anIndex int32)        // procedure
	// ClearFilter
	//  filtered items
	ClearFilter()                                // procedure
	ItemPointer(anIndex int32) TSynFoldNodeInfo  // property ItemPointer Getter
	LastItemPointer() TSynFoldNodeInfo           // property LastItemPointer Getter
	Item(index int32) TSynFoldNodeInfo           // property Item Getter
	ActionFilter() types.TSynFoldActions         // property ActionFilter Getter
	SetActionFilter(value types.TSynFoldActions) // property ActionFilter Setter
	GroupFilter() int32                          // property GroupFilter Getter
	SetGroupFilter(value int32)                  // property GroupFilter Setter
	// Line
	//  Only allowed to be set, if highlighter has CurrentLines (and is scanned)
	Line() types.TLineIdx         // property Line Getter
	SetLine(value types.TLineIdx) // property Line Setter
}

type TLazSynFoldNodeInfoList struct {
	TObject
}

func (m *TLazSynFoldNodeInfoList) CountAll() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynFoldNodeInfoListAPI().SysCallN(0, m.Instance())
	return int32(r)
}

func (m *TLazSynFoldNodeInfoList) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynFoldNodeInfoListAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TLazSynFoldNodeInfoList) CountEx(anActionFilter types.TSynFoldActions, groupFilter int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynFoldNodeInfoListAPI().SysCallN(2, m.Instance(), uintptr(anActionFilter), uintptr(groupFilter))
	return int32(r)
}

func (m *TLazSynFoldNodeInfoList) NodeInfoEx(index int32, anActionFilter types.TSynFoldActions, groupFilter int32) (result TSynFoldNodeInfo) {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(3, m.Instance(), uintptr(index), uintptr(anActionFilter), uintptr(groupFilter), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynFoldNodeInfoList) Add(anInfo TSynFoldNodeInfo) {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&anInfo)))
}

func (m *TLazSynFoldNodeInfoList) Delete(anIndex int32) {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(5, m.Instance(), uintptr(anIndex))
}

func (m *TLazSynFoldNodeInfoList) ClearFilter() {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(6, m.Instance())
}

func (m *TLazSynFoldNodeInfoList) ItemPointer(anIndex int32) (result TSynFoldNodeInfo) {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(7, m.Instance(), uintptr(anIndex), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynFoldNodeInfoList) LastItemPointer() (result TSynFoldNodeInfo) {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(8, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynFoldNodeInfoList) Item(index int32) (result TSynFoldNodeInfo) {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(9, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynFoldNodeInfoList) ActionFilter() types.TSynFoldActions {
	if !m.IsValid() {
		return 0
	}
	r := lazSynFoldNodeInfoListAPI().SysCallN(10, 0, m.Instance())
	return types.TSynFoldActions(r)
}

func (m *TLazSynFoldNodeInfoList) SetActionFilter(value types.TSynFoldActions) {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynFoldNodeInfoList) GroupFilter() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynFoldNodeInfoListAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynFoldNodeInfoList) SetGroupFilter(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynFoldNodeInfoList) Line() types.TLineIdx {
	if !m.IsValid() {
		return 0
	}
	r := lazSynFoldNodeInfoListAPI().SysCallN(12, 0, m.Instance())
	return types.TLineIdx(r)
}

func (m *TLazSynFoldNodeInfoList) SetLine(value types.TLineIdx) {
	if !m.IsValid() {
		return
	}
	lazSynFoldNodeInfoListAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

var (
	lazSynFoldNodeInfoListOnce   base.Once
	lazSynFoldNodeInfoListImport *imports.Imports = nil
)

func lazSynFoldNodeInfoListAPI() *imports.Imports {
	lazSynFoldNodeInfoListOnce.Do(func() {
		lazSynFoldNodeInfoListImport = api.NewDefaultImports()
		lazSynFoldNodeInfoListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazSynFoldNodeInfoList_CountAll", 0), // function CountAll
			/* 1 */ imports.NewTable("TLazSynFoldNodeInfoList_Count", 0), // function Count
			/* 2 */ imports.NewTable("TLazSynFoldNodeInfoList_CountEx", 0), // function CountEx
			/* 3 */ imports.NewTable("TLazSynFoldNodeInfoList_NodeInfoEx", 0), // function NodeInfoEx
			/* 4 */ imports.NewTable("TLazSynFoldNodeInfoList_Add", 0), // procedure Add
			/* 5 */ imports.NewTable("TLazSynFoldNodeInfoList_Delete", 0), // procedure Delete
			/* 6 */ imports.NewTable("TLazSynFoldNodeInfoList_ClearFilter", 0), // procedure ClearFilter
			/* 7 */ imports.NewTable("TLazSynFoldNodeInfoList_ItemPointer", 0), // property ItemPointer
			/* 8 */ imports.NewTable("TLazSynFoldNodeInfoList_LastItemPointer", 0), // property LastItemPointer
			/* 9 */ imports.NewTable("TLazSynFoldNodeInfoList_Item", 0), // property Item
			/* 10 */ imports.NewTable("TLazSynFoldNodeInfoList_ActionFilter", 0), // property ActionFilter
			/* 11 */ imports.NewTable("TLazSynFoldNodeInfoList_GroupFilter", 0), // property GroupFilter
			/* 12 */ imports.NewTable("TLazSynFoldNodeInfoList_Line", 0), // property Line
		}
	})
	return lazSynFoldNodeInfoListImport
}
