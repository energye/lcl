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

// IFlowPanelControl Parent: ICollectionItem
type IFlowPanelControl interface {
	ICollectionItem
	Control() IControl              // property
	SetControl(AValue IControl)     // property
	WrapAfter() TWrapAfter          // property
	SetWrapAfter(AValue TWrapAfter) // property
	AllowAdd() bool                 // function
	AllowDelete() bool              // function
}

// TFlowPanelControl Parent: TCollectionItem
type TFlowPanelControl struct {
	TCollectionItem
}

func NewFlowPanelControl(ACollection ICollection) IFlowPanelControl {
	r1 := flowPanelControlImportAPI().SysCallN(4, GetObjectUintptr(ACollection))
	return AsFlowPanelControl(r1)
}

func (m *TFlowPanelControl) Control() IControl {
	r1 := flowPanelControlImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TFlowPanelControl) SetControl(AValue IControl) {
	flowPanelControlImportAPI().SysCallN(3, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFlowPanelControl) WrapAfter() TWrapAfter {
	r1 := flowPanelControlImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TWrapAfter(r1)
}

func (m *TFlowPanelControl) SetWrapAfter(AValue TWrapAfter) {
	flowPanelControlImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TFlowPanelControl) AllowAdd() bool {
	r1 := flowPanelControlImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TFlowPanelControl) AllowDelete() bool {
	r1 := flowPanelControlImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func FlowPanelControlClass() TClass {
	ret := flowPanelControlImportAPI().SysCallN(2)
	return TClass(ret)
}

var (
	flowPanelControlImport       *imports.Imports = nil
	flowPanelControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FlowPanelControl_AllowAdd", 0),
		/*1*/ imports.NewTable("FlowPanelControl_AllowDelete", 0),
		/*2*/ imports.NewTable("FlowPanelControl_Class", 0),
		/*3*/ imports.NewTable("FlowPanelControl_Control", 0),
		/*4*/ imports.NewTable("FlowPanelControl_Create", 0),
		/*5*/ imports.NewTable("FlowPanelControl_WrapAfter", 0),
	}
)

func flowPanelControlImportAPI() *imports.Imports {
	if flowPanelControlImport == nil {
		flowPanelControlImport = NewDefaultImports()
		flowPanelControlImport.SetImportTable(flowPanelControlImportTables)
		flowPanelControlImportTables = nil
	}
	return flowPanelControlImport
}
