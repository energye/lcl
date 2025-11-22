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

// ICustomCheckBox Parent: IButtonControl
type ICustomCheckBox interface {
	IButtonControl
	Alignment() types.TLeftRight         // property Alignment Getter
	SetAlignment(value types.TLeftRight) // property Alignment Setter
	AllowGrayed() bool                   // property AllowGrayed Getter
	SetAllowGrayed(value bool)           // property AllowGrayed Setter
	State() types.TCheckBoxState         // property State Getter
	SetState(value types.TCheckBoxState) // property State Setter
	ShortCut() types.TShortCut           // property ShortCut Getter
	ShortCutKey2() types.TShortCut       // property ShortCutKey2 Getter
	SetOnChange(fn TNotifyEvent)         // property event
}

type TCustomCheckBox struct {
	TButtonControl
}

func (m *TCustomCheckBox) Alignment() types.TLeftRight {
	if !m.IsValid() {
		return 0
	}
	r := customCheckBoxAPI().SysCallN(1, 0, m.Instance())
	return types.TLeftRight(r)
}

func (m *TCustomCheckBox) SetAlignment(value types.TLeftRight) {
	if !m.IsValid() {
		return
	}
	customCheckBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCheckBox) AllowGrayed() bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckBoxAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCheckBox) SetAllowGrayed(value bool) {
	if !m.IsValid() {
		return
	}
	customCheckBoxAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCheckBox) State() types.TCheckBoxState {
	if !m.IsValid() {
		return 0
	}
	r := customCheckBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TCheckBoxState(r)
}

func (m *TCustomCheckBox) SetState(value types.TCheckBoxState) {
	if !m.IsValid() {
		return
	}
	customCheckBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCheckBox) ShortCut() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := customCheckBoxAPI().SysCallN(4, m.Instance())
	return types.TShortCut(r)
}

func (m *TCustomCheckBox) ShortCutKey2() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := customCheckBoxAPI().SysCallN(5, m.Instance())
	return types.TShortCut(r)
}

func (m *TCustomCheckBox) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, customCheckBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomCheckBox class constructor
func NewCustomCheckBox(theOwner IComponent) ICustomCheckBox {
	r := customCheckBoxAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomCheckBox(r)
}

func TCustomCheckBoxClass() types.TClass {
	r := customCheckBoxAPI().SysCallN(7)
	return types.TClass(r)
}

var (
	customCheckBoxOnce   base.Once
	customCheckBoxImport *imports.Imports = nil
)

func customCheckBoxAPI() *imports.Imports {
	customCheckBoxOnce.Do(func() {
		customCheckBoxImport = api.NewDefaultImports()
		customCheckBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCheckBox_Create", 0), // constructor NewCustomCheckBox
			/* 1 */ imports.NewTable("TCustomCheckBox_Alignment", 0), // property Alignment
			/* 2 */ imports.NewTable("TCustomCheckBox_AllowGrayed", 0), // property AllowGrayed
			/* 3 */ imports.NewTable("TCustomCheckBox_State", 0), // property State
			/* 4 */ imports.NewTable("TCustomCheckBox_ShortCut", 0), // property ShortCut
			/* 5 */ imports.NewTable("TCustomCheckBox_ShortCutKey2", 0), // property ShortCutKey2
			/* 6 */ imports.NewTable("TCustomCheckBox_OnChange", 0), // event OnChange
			/* 7 */ imports.NewTable("TCustomCheckBox_TClass", 0), // function TCustomCheckBoxClass
		}
	})
	return customCheckBoxImport
}
