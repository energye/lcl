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

// ICustomFlowPanel Parent: ICustomPanel
type ICustomFlowPanel interface {
	ICustomPanel
	GetControlIndexWithControl(control IControl) int32           // function
	SetControlIndexWithControlInt(control IControl, index int32) // procedure
	AutoWrap() bool                                              // property AutoWrap Getter
	SetAutoWrap(value bool)                                      // property AutoWrap Setter
	ControlList() IFlowPanelControlList                          // property ControlList Getter
	SetControlList(value IFlowPanelControlList)                  // property ControlList Setter
	FlowStyle() types.TFlowStyle                                 // property FlowStyle Getter
	SetFlowStyle(value types.TFlowStyle)                         // property FlowStyle Setter
	FlowLayout() types.TTextLayout                               // property FlowLayout Getter
	SetFlowLayout(value types.TTextLayout)                       // property FlowLayout Setter
}

type TCustomFlowPanel struct {
	TCustomPanel
}

func (m *TCustomFlowPanel) GetControlIndexWithControl(control IControl) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customFlowPanelAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(control))
	return int32(r)
}

func (m *TCustomFlowPanel) SetControlIndexWithControlInt(control IControl, index int32) {
	if !m.IsValid() {
		return
	}
	customFlowPanelAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(control), uintptr(index))
}

func (m *TCustomFlowPanel) AutoWrap() bool {
	if !m.IsValid() {
		return false
	}
	r := customFlowPanelAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomFlowPanel) SetAutoWrap(value bool) {
	if !m.IsValid() {
		return
	}
	customFlowPanelAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomFlowPanel) ControlList() IFlowPanelControlList {
	if !m.IsValid() {
		return nil
	}
	r := customFlowPanelAPI().SysCallN(4, 0, m.Instance())
	return AsFlowPanelControlList(r)
}

func (m *TCustomFlowPanel) SetControlList(value IFlowPanelControlList) {
	if !m.IsValid() {
		return
	}
	customFlowPanelAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomFlowPanel) FlowStyle() types.TFlowStyle {
	if !m.IsValid() {
		return 0
	}
	r := customFlowPanelAPI().SysCallN(5, 0, m.Instance())
	return types.TFlowStyle(r)
}

func (m *TCustomFlowPanel) SetFlowStyle(value types.TFlowStyle) {
	if !m.IsValid() {
		return
	}
	customFlowPanelAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomFlowPanel) FlowLayout() types.TTextLayout {
	if !m.IsValid() {
		return 0
	}
	r := customFlowPanelAPI().SysCallN(6, 0, m.Instance())
	return types.TTextLayout(r)
}

func (m *TCustomFlowPanel) SetFlowLayout(value types.TTextLayout) {
	if !m.IsValid() {
		return
	}
	customFlowPanelAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

// NewCustomFlowPanel class constructor
func NewCustomFlowPanel(owner IComponent) ICustomFlowPanel {
	r := customFlowPanelAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomFlowPanel(r)
}

func TCustomFlowPanelClass() types.TClass {
	r := customFlowPanelAPI().SysCallN(7)
	return types.TClass(r)
}

var (
	customFlowPanelOnce   base.Once
	customFlowPanelImport *imports.Imports = nil
)

func customFlowPanelAPI() *imports.Imports {
	customFlowPanelOnce.Do(func() {
		customFlowPanelImport = api.NewDefaultImports()
		customFlowPanelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomFlowPanel_Create", 0), // constructor NewCustomFlowPanel
			/* 1 */ imports.NewTable("TCustomFlowPanel_GetControlIndexWithControl", 0), // function GetControlIndexWithControl
			/* 2 */ imports.NewTable("TCustomFlowPanel_SetControlIndexWithControlInt", 0), // procedure SetControlIndexWithControlInt
			/* 3 */ imports.NewTable("TCustomFlowPanel_AutoWrap", 0), // property AutoWrap
			/* 4 */ imports.NewTable("TCustomFlowPanel_ControlList", 0), // property ControlList
			/* 5 */ imports.NewTable("TCustomFlowPanel_FlowStyle", 0), // property FlowStyle
			/* 6 */ imports.NewTable("TCustomFlowPanel_FlowLayout", 0), // property FlowLayout
			/* 7 */ imports.NewTable("TCustomFlowPanel_TClass", 0), // function TCustomFlowPanelClass
		}
	})
	return customFlowPanelImport
}
