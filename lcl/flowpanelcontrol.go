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

// IFlowPanelControl Parent: ICollectionItem
type IFlowPanelControl interface {
	ICollectionItem
	// AllowAdd
	//  These methods are used by the Object Inspector only
	AllowAdd() bool                      // function
	AllowDelete() bool                   // function
	Control() IControl                   // property Control Getter
	SetControl(value IControl)           // property Control Setter
	WrapAfter() types.TWrapAfter         // property WrapAfter Getter
	SetWrapAfter(value types.TWrapAfter) // property WrapAfter Setter
	AsIntfObjInspInterface() uintptr
}

type TFlowPanelControl struct {
	TCollectionItem
}

func (m *TFlowPanelControl) AllowAdd() bool {
	if !m.IsValid() {
		return false
	}
	r := flowPanelControlAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TFlowPanelControl) AllowDelete() bool {
	if !m.IsValid() {
		return false
	}
	r := flowPanelControlAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TFlowPanelControl) Control() IControl {
	if !m.IsValid() {
		return nil
	}
	r := flowPanelControlAPI().SysCallN(3, 0, m.Instance())
	return AsControl(r)
}

func (m *TFlowPanelControl) SetControl(value IControl) {
	if !m.IsValid() {
		return
	}
	flowPanelControlAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFlowPanelControl) WrapAfter() types.TWrapAfter {
	if !m.IsValid() {
		return 0
	}
	r := flowPanelControlAPI().SysCallN(4, 0, m.Instance())
	return types.TWrapAfter(r)
}

func (m *TFlowPanelControl) SetWrapAfter(value types.TWrapAfter) {
	if !m.IsValid() {
		return
	}
	flowPanelControlAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TFlowPanelControl) AsIntfObjInspInterface() uintptr {
	return m.GetIntfPointer(0)
}

// NewFlowPanelControl class constructor
func NewFlowPanelControl(collection ICollection) IFlowPanelControl {
	var objInspInterfacePtr uintptr // IObjInspInterface
	r := flowPanelControlAPI().SysCallN(0, base.GetObjectUintptr(collection), uintptr(base.UnsafePointer(&objInspInterfacePtr)))
	ret := AsFlowPanelControl(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, objInspInterfacePtr)
	}
	return ret
}

var (
	flowPanelControlOnce   base.Once
	flowPanelControlImport *imports.Imports = nil
)

func flowPanelControlAPI() *imports.Imports {
	flowPanelControlOnce.Do(func() {
		flowPanelControlImport = api.NewDefaultImports()
		flowPanelControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFlowPanelControl_Create", 0), // constructor NewFlowPanelControl
			/* 1 */ imports.NewTable("TFlowPanelControl_AllowAdd", 0), // function AllowAdd
			/* 2 */ imports.NewTable("TFlowPanelControl_AllowDelete", 0), // function AllowDelete
			/* 3 */ imports.NewTable("TFlowPanelControl_Control", 0), // property Control
			/* 4 */ imports.NewTable("TFlowPanelControl_WrapAfter", 0), // property WrapAfter
		}
	})
	return flowPanelControlImport
}
