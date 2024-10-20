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

// ICustomFlowPanel Parent: ICustomPanel
type ICustomFlowPanel interface {
	ICustomPanel
	AutoWrap() bool                              // property
	SetAutoWrap(AValue bool)                     // property
	ControlList() IFlowPanelControlList          // property
	SetControlList(AValue IFlowPanelControlList) // property
	FlowStyle() TFlowStyle                       // property
	SetFlowStyle(AValue TFlowStyle)              // property
	FlowLayout() TTextLayout                     // property
	SetFlowLayout(AValue TTextLayout)            // property
}

// TCustomFlowPanel Parent: TCustomPanel
type TCustomFlowPanel struct {
	TCustomPanel
}

func NewCustomFlowPanel(AOwner IComponent) ICustomFlowPanel {
	r1 := customFlowPanelImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsCustomFlowPanel(r1)
}

func (m *TCustomFlowPanel) AutoWrap() bool {
	r1 := customFlowPanelImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomFlowPanel) SetAutoWrap(AValue bool) {
	customFlowPanelImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomFlowPanel) ControlList() IFlowPanelControlList {
	r1 := customFlowPanelImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return AsFlowPanelControlList(r1)
}

func (m *TCustomFlowPanel) SetControlList(AValue IFlowPanelControlList) {
	customFlowPanelImportAPI().SysCallN(2, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomFlowPanel) FlowStyle() TFlowStyle {
	r1 := customFlowPanelImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TFlowStyle(r1)
}

func (m *TCustomFlowPanel) SetFlowStyle(AValue TFlowStyle) {
	customFlowPanelImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomFlowPanel) FlowLayout() TTextLayout {
	r1 := customFlowPanelImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TCustomFlowPanel) SetFlowLayout(AValue TTextLayout) {
	customFlowPanelImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func CustomFlowPanelClass() TClass {
	ret := customFlowPanelImportAPI().SysCallN(1)
	return TClass(ret)
}

var (
	customFlowPanelImport       *imports.Imports = nil
	customFlowPanelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomFlowPanel_AutoWrap", 0),
		/*1*/ imports.NewTable("CustomFlowPanel_Class", 0),
		/*2*/ imports.NewTable("CustomFlowPanel_ControlList", 0),
		/*3*/ imports.NewTable("CustomFlowPanel_Create", 0),
		/*4*/ imports.NewTable("CustomFlowPanel_FlowLayout", 0),
		/*5*/ imports.NewTable("CustomFlowPanel_FlowStyle", 0),
	}
)

func customFlowPanelImportAPI() *imports.Imports {
	if customFlowPanelImport == nil {
		customFlowPanelImport = NewDefaultImports()
		customFlowPanelImport.SetImportTable(customFlowPanelImportTables)
		customFlowPanelImportTables = nil
	}
	return customFlowPanelImport
}
