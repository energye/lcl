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

// ICustomRadioGroup Parent: ICustomGroupBox
type ICustomRadioGroup interface {
	ICustomGroupBox
	CanModify() bool                           // function
	Rows() int32                               // function
	AutoFill() bool                            // property AutoFill Getter
	SetAutoFill(value bool)                    // property AutoFill Setter
	Buttons(index int32) IRadioButton          // property Buttons Getter
	ItemIndex() int32                          // property ItemIndex Getter
	SetItemIndex(value int32)                  // property ItemIndex Setter
	Items() IStrings                           // property Items Getter
	SetItems(value IStrings)                   // property Items Setter
	Columns() int32                            // property Columns Getter
	SetColumns(value int32)                    // property Columns Setter
	ColumnLayout() types.TColumnLayout         // property ColumnLayout Getter
	SetColumnLayout(value types.TColumnLayout) // property ColumnLayout Setter
	SetOnItemEnter(fn TNotifyEvent)            // property event
	SetOnItemExit(fn TNotifyEvent)             // property event
	SetOnSelectionChanged(fn TNotifyEvent)     // property event
}

type TCustomRadioGroup struct {
	TCustomGroupBox
}

func (m *TCustomRadioGroup) CanModify() bool {
	if !m.IsValid() {
		return false
	}
	r := customRadioGroupAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomRadioGroup) Rows() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customRadioGroupAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TCustomRadioGroup) AutoFill() bool {
	if !m.IsValid() {
		return false
	}
	r := customRadioGroupAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomRadioGroup) SetAutoFill(value bool) {
	if !m.IsValid() {
		return
	}
	customRadioGroupAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomRadioGroup) Buttons(index int32) IRadioButton {
	if !m.IsValid() {
		return nil
	}
	r := customRadioGroupAPI().SysCallN(4, m.Instance(), uintptr(index))
	return AsRadioButton(r)
}

func (m *TCustomRadioGroup) ItemIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customRadioGroupAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TCustomRadioGroup) SetItemIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customRadioGroupAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomRadioGroup) Items() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := customRadioGroupAPI().SysCallN(6, 0, m.Instance())
	return AsStrings(r)
}

func (m *TCustomRadioGroup) SetItems(value IStrings) {
	if !m.IsValid() {
		return
	}
	customRadioGroupAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomRadioGroup) Columns() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customRadioGroupAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TCustomRadioGroup) SetColumns(value int32) {
	if !m.IsValid() {
		return
	}
	customRadioGroupAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomRadioGroup) ColumnLayout() types.TColumnLayout {
	if !m.IsValid() {
		return 0
	}
	r := customRadioGroupAPI().SysCallN(8, 0, m.Instance())
	return types.TColumnLayout(r)
}

func (m *TCustomRadioGroup) SetColumnLayout(value types.TColumnLayout) {
	if !m.IsValid() {
		return
	}
	customRadioGroupAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomRadioGroup) SetOnItemEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, customRadioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomRadioGroup) SetOnItemExit(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, customRadioGroupAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomRadioGroup) SetOnSelectionChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, customRadioGroupAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomRadioGroup class constructor
func NewCustomRadioGroup(theOwner IComponent) ICustomRadioGroup {
	r := customRadioGroupAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomRadioGroup(r)
}

func TCustomRadioGroupClass() types.TClass {
	r := customRadioGroupAPI().SysCallN(12)
	return types.TClass(r)
}

var (
	customRadioGroupOnce   base.Once
	customRadioGroupImport *imports.Imports = nil
)

func customRadioGroupAPI() *imports.Imports {
	customRadioGroupOnce.Do(func() {
		customRadioGroupImport = api.NewDefaultImports()
		customRadioGroupImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomRadioGroup_Create", 0), // constructor NewCustomRadioGroup
			/* 1 */ imports.NewTable("TCustomRadioGroup_CanModify", 0), // function CanModify
			/* 2 */ imports.NewTable("TCustomRadioGroup_Rows", 0), // function Rows
			/* 3 */ imports.NewTable("TCustomRadioGroup_AutoFill", 0), // property AutoFill
			/* 4 */ imports.NewTable("TCustomRadioGroup_Buttons", 0), // property Buttons
			/* 5 */ imports.NewTable("TCustomRadioGroup_ItemIndex", 0), // property ItemIndex
			/* 6 */ imports.NewTable("TCustomRadioGroup_Items", 0), // property Items
			/* 7 */ imports.NewTable("TCustomRadioGroup_Columns", 0), // property Columns
			/* 8 */ imports.NewTable("TCustomRadioGroup_ColumnLayout", 0), // property ColumnLayout
			/* 9 */ imports.NewTable("TCustomRadioGroup_OnItemEnter", 0), // event OnItemEnter
			/* 10 */ imports.NewTable("TCustomRadioGroup_OnItemExit", 0), // event OnItemExit
			/* 11 */ imports.NewTable("TCustomRadioGroup_OnSelectionChanged", 0), // event OnSelectionChanged
			/* 12 */ imports.NewTable("TCustomRadioGroup_TClass", 0), // function TCustomRadioGroupClass
		}
	})
	return customRadioGroupImport
}
