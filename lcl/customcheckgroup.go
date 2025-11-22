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

// ICustomCheckGroup Parent: ICustomGroupBox
type ICustomCheckGroup interface {
	ICustomGroupBox
	Rows() int32                               // function
	AutoFill() bool                            // property AutoFill Getter
	SetAutoFill(value bool)                    // property AutoFill Setter
	Buttons(index int32) ICheckBox             // property Buttons Getter
	Items() IStrings                           // property Items Getter
	SetItems(value IStrings)                   // property Items Setter
	Checked(index int32) bool                  // property Checked Getter
	SetChecked(index int32, value bool)        // property Checked Setter
	CheckEnabled(index int32) bool             // property CheckEnabled Getter
	SetCheckEnabled(index int32, value bool)   // property CheckEnabled Setter
	Columns() int32                            // property Columns Getter
	SetColumns(value int32)                    // property Columns Setter
	ColumnLayout() types.TColumnLayout         // property ColumnLayout Getter
	SetColumnLayout(value types.TColumnLayout) // property ColumnLayout Setter
	SetOnItemClick(fn TCheckGroupClicked)      // property event
}

type TCustomCheckGroup struct {
	TCustomGroupBox
}

func (m *TCustomCheckGroup) Rows() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customCheckGroupAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TCustomCheckGroup) AutoFill() bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckGroupAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCheckGroup) SetAutoFill(value bool) {
	if !m.IsValid() {
		return
	}
	customCheckGroupAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCheckGroup) Buttons(index int32) ICheckBox {
	if !m.IsValid() {
		return nil
	}
	r := customCheckGroupAPI().SysCallN(3, m.Instance(), uintptr(index))
	return AsCheckBox(r)
}

func (m *TCustomCheckGroup) Items() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := customCheckGroupAPI().SysCallN(4, 0, m.Instance())
	return AsStrings(r)
}

func (m *TCustomCheckGroup) SetItems(value IStrings) {
	if !m.IsValid() {
		return
	}
	customCheckGroupAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomCheckGroup) Checked(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckGroupAPI().SysCallN(5, 0, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomCheckGroup) SetChecked(index int32, value bool) {
	if !m.IsValid() {
		return
	}
	customCheckGroupAPI().SysCallN(5, 1, m.Instance(), uintptr(index), api.PasBool(value))
}

func (m *TCustomCheckGroup) CheckEnabled(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckGroupAPI().SysCallN(6, 0, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomCheckGroup) SetCheckEnabled(index int32, value bool) {
	if !m.IsValid() {
		return
	}
	customCheckGroupAPI().SysCallN(6, 1, m.Instance(), uintptr(index), api.PasBool(value))
}

func (m *TCustomCheckGroup) Columns() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customCheckGroupAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TCustomCheckGroup) SetColumns(value int32) {
	if !m.IsValid() {
		return
	}
	customCheckGroupAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCheckGroup) ColumnLayout() types.TColumnLayout {
	if !m.IsValid() {
		return 0
	}
	r := customCheckGroupAPI().SysCallN(8, 0, m.Instance())
	return types.TColumnLayout(r)
}

func (m *TCustomCheckGroup) SetColumnLayout(value types.TColumnLayout) {
	if !m.IsValid() {
		return
	}
	customCheckGroupAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCheckGroup) SetOnItemClick(fn TCheckGroupClicked) {
	if !m.IsValid() {
		return
	}
	cb := makeTCheckGroupClicked(fn)
	base.SetEvent(m, 9, customCheckGroupAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomCheckGroup class constructor
func NewCustomCheckGroup(theOwner IComponent) ICustomCheckGroup {
	r := customCheckGroupAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomCheckGroup(r)
}

func TCustomCheckGroupClass() types.TClass {
	r := customCheckGroupAPI().SysCallN(10)
	return types.TClass(r)
}

var (
	customCheckGroupOnce   base.Once
	customCheckGroupImport *imports.Imports = nil
)

func customCheckGroupAPI() *imports.Imports {
	customCheckGroupOnce.Do(func() {
		customCheckGroupImport = api.NewDefaultImports()
		customCheckGroupImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCheckGroup_Create", 0), // constructor NewCustomCheckGroup
			/* 1 */ imports.NewTable("TCustomCheckGroup_Rows", 0), // function Rows
			/* 2 */ imports.NewTable("TCustomCheckGroup_AutoFill", 0), // property AutoFill
			/* 3 */ imports.NewTable("TCustomCheckGroup_Buttons", 0), // property Buttons
			/* 4 */ imports.NewTable("TCustomCheckGroup_Items", 0), // property Items
			/* 5 */ imports.NewTable("TCustomCheckGroup_Checked", 0), // property Checked
			/* 6 */ imports.NewTable("TCustomCheckGroup_CheckEnabled", 0), // property CheckEnabled
			/* 7 */ imports.NewTable("TCustomCheckGroup_Columns", 0), // property Columns
			/* 8 */ imports.NewTable("TCustomCheckGroup_ColumnLayout", 0), // property ColumnLayout
			/* 9 */ imports.NewTable("TCustomCheckGroup_OnItemClick", 0), // event OnItemClick
			/* 10 */ imports.NewTable("TCustomCheckGroup_TClass", 0), // function TCustomCheckGroupClass
		}
	})
	return customCheckGroupImport
}
