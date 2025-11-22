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

// IFlowPanelControlList Parent: IOwnedCollection
type IFlowPanelControlList interface {
	IOwnedCollection
	IndexOf(control IControl) int32 // function
	// AllowAdd
	//  These methods are used by the Object Inspector only
	AllowAdd() bool                                                         // function
	AllowDelete() bool                                                      // function
	ItemsWithIntToFlowPanelControl(index int32) IFlowPanelControl           // property Items Getter
	SetItemsWithIntToFlowPanelControl(index int32, value IFlowPanelControl) // property Items Setter
	AsIntfObjInspInterface() uintptr
}

type TFlowPanelControlList struct {
	TOwnedCollection
}

func (m *TFlowPanelControlList) IndexOf(control IControl) int32 {
	if !m.IsValid() {
		return 0
	}
	r := flowPanelControlListAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(control))
	return int32(r)
}

func (m *TFlowPanelControlList) AllowAdd() bool {
	if !m.IsValid() {
		return false
	}
	r := flowPanelControlListAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TFlowPanelControlList) AllowDelete() bool {
	if !m.IsValid() {
		return false
	}
	r := flowPanelControlListAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TFlowPanelControlList) ItemsWithIntToFlowPanelControl(index int32) IFlowPanelControl {
	if !m.IsValid() {
		return nil
	}
	r := flowPanelControlListAPI().SysCallN(4, 0, m.Instance(), uintptr(index))
	return AsFlowPanelControl(r)
}

func (m *TFlowPanelControlList) SetItemsWithIntToFlowPanelControl(index int32, value IFlowPanelControl) {
	if !m.IsValid() {
		return
	}
	flowPanelControlListAPI().SysCallN(4, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TFlowPanelControlList) AsIntfObjInspInterface() uintptr {
	return m.GetIntfPointer(0)
}

// NewFlowPanelControlList class constructor
func NewFlowPanelControlList(owner IPersistent) IFlowPanelControlList {
	var objInspInterfacePtr uintptr // IObjInspInterface
	r := flowPanelControlListAPI().SysCallN(0, base.GetObjectUintptr(owner), uintptr(base.UnsafePointer(&objInspInterfacePtr)))
	ret := AsFlowPanelControlList(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, objInspInterfacePtr)
	}
	return ret
}

var (
	flowPanelControlListOnce   base.Once
	flowPanelControlListImport *imports.Imports = nil
)

func flowPanelControlListAPI() *imports.Imports {
	flowPanelControlListOnce.Do(func() {
		flowPanelControlListImport = api.NewDefaultImports()
		flowPanelControlListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFlowPanelControlList_Create", 0), // constructor NewFlowPanelControlList
			/* 1 */ imports.NewTable("TFlowPanelControlList_IndexOf", 0), // function IndexOf
			/* 2 */ imports.NewTable("TFlowPanelControlList_AllowAdd", 0), // function AllowAdd
			/* 3 */ imports.NewTable("TFlowPanelControlList_AllowDelete", 0), // function AllowDelete
			/* 4 */ imports.NewTable("TFlowPanelControlList_ItemsWithIntToFlowPanelControl", 0), // property ItemsWithIntToFlowPanelControl
		}
	})
	return flowPanelControlListImport
}
