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

// ICustomCheckListBox Parent: ICustomListBox
type ICustomCheckListBox interface {
	ICustomListBox
	CalculateStandardItemHeight() int32                                        // function
	Toggle(index int32)                                                        // procedure
	CheckAll(state types.TCheckBoxState, allowGrayed bool, allowDisabled bool) // procedure
	Exchange(index1 int32, index2 int32)                                       // procedure
	AllowGrayed() bool                                                         // property AllowGrayed Getter
	SetAllowGrayed(value bool)                                                 // property AllowGrayed Setter
	Checked(index int32) bool                                                  // property Checked Getter
	SetChecked(index int32, value bool)                                        // property Checked Setter
	Header(index int32) bool                                                   // property Header Getter
	SetHeader(index int32, value bool)                                         // property Header Setter
	HeaderBackgroundColor() types.TColor                                       // property HeaderBackgroundColor Getter
	SetHeaderBackgroundColor(value types.TColor)                               // property HeaderBackgroundColor Setter
	HeaderColor() types.TColor                                                 // property HeaderColor Getter
	SetHeaderColor(value types.TColor)                                         // property HeaderColor Setter
	ItemEnabled(index int32) bool                                              // property ItemEnabled Getter
	SetItemEnabled(index int32, value bool)                                    // property ItemEnabled Setter
	State(index int32) types.TCheckBoxState                                    // property State Getter
	SetState(index int32, value types.TCheckBoxState)                          // property State Setter
	SetOnClickCheck(fn TNotifyEvent)                                           // property event
}

type TCustomCheckListBox struct {
	TCustomListBox
}

func (m *TCustomCheckListBox) CalculateStandardItemHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customCheckListBoxAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TCustomCheckListBox) Toggle(index int32) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(2, m.Instance(), uintptr(index))
}

func (m *TCustomCheckListBox) CheckAll(state types.TCheckBoxState, allowGrayed bool, allowDisabled bool) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(3, m.Instance(), uintptr(state), api.PasBool(allowGrayed), api.PasBool(allowDisabled))
}

func (m *TCustomCheckListBox) Exchange(index1 int32, index2 int32) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(4, m.Instance(), uintptr(index1), uintptr(index2))
}

func (m *TCustomCheckListBox) AllowGrayed() bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckListBoxAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCheckListBox) SetAllowGrayed(value bool) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCheckListBox) Checked(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckListBoxAPI().SysCallN(6, 0, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomCheckListBox) SetChecked(index int32, value bool) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(6, 1, m.Instance(), uintptr(index), api.PasBool(value))
}

func (m *TCustomCheckListBox) Header(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckListBoxAPI().SysCallN(7, 0, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomCheckListBox) SetHeader(index int32, value bool) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(7, 1, m.Instance(), uintptr(index), api.PasBool(value))
}

func (m *TCustomCheckListBox) HeaderBackgroundColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customCheckListBoxAPI().SysCallN(8, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomCheckListBox) SetHeaderBackgroundColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCheckListBox) HeaderColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customCheckListBoxAPI().SysCallN(9, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomCheckListBox) SetHeaderColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCheckListBox) ItemEnabled(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckListBoxAPI().SysCallN(10, 0, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomCheckListBox) SetItemEnabled(index int32, value bool) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(10, 1, m.Instance(), uintptr(index), api.PasBool(value))
}

func (m *TCustomCheckListBox) State(index int32) types.TCheckBoxState {
	if !m.IsValid() {
		return 0
	}
	r := customCheckListBoxAPI().SysCallN(11, 0, m.Instance(), uintptr(index))
	return types.TCheckBoxState(r)
}

func (m *TCustomCheckListBox) SetState(index int32, value types.TCheckBoxState) {
	if !m.IsValid() {
		return
	}
	customCheckListBoxAPI().SysCallN(11, 1, m.Instance(), uintptr(index), uintptr(value))
}

func (m *TCustomCheckListBox) SetOnClickCheck(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, customCheckListBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomCheckListBox class constructor
func NewCustomCheckListBox(owner IComponent) ICustomCheckListBox {
	r := customCheckListBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomCheckListBox(r)
}

func TCustomCheckListBoxClass() types.TClass {
	r := customCheckListBoxAPI().SysCallN(13)
	return types.TClass(r)
}

var (
	customCheckListBoxOnce   base.Once
	customCheckListBoxImport *imports.Imports = nil
)

func customCheckListBoxAPI() *imports.Imports {
	customCheckListBoxOnce.Do(func() {
		customCheckListBoxImport = api.NewDefaultImports()
		customCheckListBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCheckListBox_Create", 0), // constructor NewCustomCheckListBox
			/* 1 */ imports.NewTable("TCustomCheckListBox_CalculateStandardItemHeight", 0), // function CalculateStandardItemHeight
			/* 2 */ imports.NewTable("TCustomCheckListBox_Toggle", 0), // procedure Toggle
			/* 3 */ imports.NewTable("TCustomCheckListBox_CheckAll", 0), // procedure CheckAll
			/* 4 */ imports.NewTable("TCustomCheckListBox_Exchange", 0), // procedure Exchange
			/* 5 */ imports.NewTable("TCustomCheckListBox_AllowGrayed", 0), // property AllowGrayed
			/* 6 */ imports.NewTable("TCustomCheckListBox_Checked", 0), // property Checked
			/* 7 */ imports.NewTable("TCustomCheckListBox_Header", 0), // property Header
			/* 8 */ imports.NewTable("TCustomCheckListBox_HeaderBackgroundColor", 0), // property HeaderBackgroundColor
			/* 9 */ imports.NewTable("TCustomCheckListBox_HeaderColor", 0), // property HeaderColor
			/* 10 */ imports.NewTable("TCustomCheckListBox_ItemEnabled", 0), // property ItemEnabled
			/* 11 */ imports.NewTable("TCustomCheckListBox_State", 0), // property State
			/* 12 */ imports.NewTable("TCustomCheckListBox_OnClickCheck", 0), // event OnClickCheck
			/* 13 */ imports.NewTable("TCustomCheckListBox_TClass", 0), // function TCustomCheckListBoxClass
		}
	})
	return customCheckListBoxImport
}
