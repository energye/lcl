//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IFlowPanelControlList Parent: IOwnedCollection
type IFlowPanelControlList interface {
	IOwnedCollection
	ItemsForFlowPanelControl(Index int32) IFlowPanelControl            // property
	SetItemsForFlowPanelControl(Index int32, AValue IFlowPanelControl) // property
	IndexOf(AControl IControl) int32                                   // function
	AllowAdd() bool                                                    // function
	AllowDelete() bool                                                 // function
}

// TFlowPanelControlList Parent: TOwnedCollection
type TFlowPanelControlList struct {
	TOwnedCollection
}

func NewFlowPanelControlList(AOwner IPersistent) IFlowPanelControlList {
	r1 := flowPanelControlListImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsFlowPanelControlList(r1)
}

func (m *TFlowPanelControlList) ItemsForFlowPanelControl(Index int32) IFlowPanelControl {
	r1 := flowPanelControlListImportAPI().SysCallN(5, 0, m.Instance(), uintptr(Index))
	return AsFlowPanelControl(r1)
}

func (m *TFlowPanelControlList) SetItemsForFlowPanelControl(Index int32, AValue IFlowPanelControl) {
	flowPanelControlListImportAPI().SysCallN(5, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TFlowPanelControlList) IndexOf(AControl IControl) int32 {
	r1 := flowPanelControlListImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(AControl))
	return int32(r1)
}

func (m *TFlowPanelControlList) AllowAdd() bool {
	r1 := flowPanelControlListImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TFlowPanelControlList) AllowDelete() bool {
	r1 := flowPanelControlListImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func FlowPanelControlListClass() TClass {
	ret := flowPanelControlListImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	flowPanelControlListImport       *imports.Imports = nil
	flowPanelControlListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FlowPanelControlList_AllowAdd", 0),
		/*1*/ imports.NewTable("FlowPanelControlList_AllowDelete", 0),
		/*2*/ imports.NewTable("FlowPanelControlList_Class", 0),
		/*3*/ imports.NewTable("FlowPanelControlList_Create", 0),
		/*4*/ imports.NewTable("FlowPanelControlList_IndexOf", 0),
		/*5*/ imports.NewTable("FlowPanelControlList_ItemsForFlowPanelControl", 0),
	}
)

func flowPanelControlListImportAPI() *imports.Imports {
	if flowPanelControlListImport == nil {
		flowPanelControlListImport = NewDefaultImports()
		flowPanelControlListImport.SetImportTable(flowPanelControlListImportTables)
		flowPanelControlListImportTables = nil
	}
	return flowPanelControlListImport
}
