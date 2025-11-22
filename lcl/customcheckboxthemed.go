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

// ICustomCheckBoxThemed Parent: ICustomControl
type ICustomCheckBoxThemed interface {
	ICustomControl
	Alignment() types.TLeftRight         // property Alignment Getter
	SetAlignment(value types.TLeftRight) // property Alignment Setter
	AllowGrayed() bool                   // property AllowGrayed Getter
	SetAllowGrayed(value bool)           // property AllowGrayed Setter
	Checked() bool                       // property Checked Getter
	SetChecked(value bool)               // property Checked Setter
	State() types.TCheckBoxState         // property State Getter
	SetState(value types.TCheckBoxState) // property State Setter
	SetOnChange(fn TNotifyEvent)         // property event
}

type TCustomCheckBoxThemed struct {
	TCustomControl
}

func (m *TCustomCheckBoxThemed) Alignment() types.TLeftRight {
	if !m.IsValid() {
		return 0
	}
	r := customCheckBoxThemedAPI().SysCallN(3, 0, m.Instance())
	return types.TLeftRight(r)
}

func (m *TCustomCheckBoxThemed) SetAlignment(value types.TLeftRight) {
	if !m.IsValid() {
		return
	}
	customCheckBoxThemedAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCheckBoxThemed) AllowGrayed() bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckBoxThemedAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCheckBoxThemed) SetAllowGrayed(value bool) {
	if !m.IsValid() {
		return
	}
	customCheckBoxThemedAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCheckBoxThemed) Checked() bool {
	if !m.IsValid() {
		return false
	}
	r := customCheckBoxThemedAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomCheckBoxThemed) SetChecked(value bool) {
	if !m.IsValid() {
		return
	}
	customCheckBoxThemedAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomCheckBoxThemed) State() types.TCheckBoxState {
	if !m.IsValid() {
		return 0
	}
	r := customCheckBoxThemedAPI().SysCallN(6, 0, m.Instance())
	return types.TCheckBoxState(r)
}

func (m *TCustomCheckBoxThemed) SetState(value types.TCheckBoxState) {
	if !m.IsValid() {
		return
	}
	customCheckBoxThemedAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCustomCheckBoxThemed) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, customCheckBoxThemedAPI(), api.MakeEventDataPtr(cb))
}

// CustomCheckBoxThemed  is static instance
var CustomCheckBoxThemed _CustomCheckBoxThemedClass

// _CustomCheckBoxThemedClass is class type defined by TCustomCheckBoxThemed
type _CustomCheckBoxThemedClass uintptr

func (_CustomCheckBoxThemedClass) GetCheckBoxSize(pixelsPerInch int32) (result types.TSize) {
	customCheckBoxThemedAPI().SysCallN(1, uintptr(pixelsPerInch), uintptr(base.UnsafePointer(&result)))
	return
}

func (_CustomCheckBoxThemedClass) PaintSelf(canvas ICanvas, caption string, rect types.TRect, state types.TCheckBoxState, rightToLeft bool, hovered bool, pressed bool, focused bool, alignment types.TLeftRight, enabled bool) {
	customCheckBoxThemedAPI().SysCallN(2, base.GetObjectUintptr(canvas), api.PasStr(caption), uintptr(base.UnsafePointer(&rect)), uintptr(state), api.PasBool(rightToLeft), api.PasBool(hovered), api.PasBool(pressed), api.PasBool(focused), uintptr(alignment), api.PasBool(enabled))
}

// NewCustomCheckBoxThemed class constructor
func NewCustomCheckBoxThemed(owner IComponent) ICustomCheckBoxThemed {
	r := customCheckBoxThemedAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomCheckBoxThemed(r)
}

func TCustomCheckBoxThemedClass() types.TClass {
	r := customCheckBoxThemedAPI().SysCallN(8)
	return types.TClass(r)
}

var (
	customCheckBoxThemedOnce   base.Once
	customCheckBoxThemedImport *imports.Imports = nil
)

func customCheckBoxThemedAPI() *imports.Imports {
	customCheckBoxThemedOnce.Do(func() {
		customCheckBoxThemedImport = api.NewDefaultImports()
		customCheckBoxThemedImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomCheckBoxThemed_Create", 0), // constructor NewCustomCheckBoxThemed
			/* 1 */ imports.NewTable("TCustomCheckBoxThemed_GetCheckBoxSize", 0), // static function GetCheckBoxSize
			/* 2 */ imports.NewTable("TCustomCheckBoxThemed_PaintSelf", 0), // static procedure PaintSelf
			/* 3 */ imports.NewTable("TCustomCheckBoxThemed_Alignment", 0), // property Alignment
			/* 4 */ imports.NewTable("TCustomCheckBoxThemed_AllowGrayed", 0), // property AllowGrayed
			/* 5 */ imports.NewTable("TCustomCheckBoxThemed_Checked", 0), // property Checked
			/* 6 */ imports.NewTable("TCustomCheckBoxThemed_State", 0), // property State
			/* 7 */ imports.NewTable("TCustomCheckBoxThemed_OnChange", 0), // event OnChange
			/* 8 */ imports.NewTable("TCustomCheckBoxThemed_TClass", 0), // function TCustomCheckBoxThemedClass
		}
	})
	return customCheckBoxThemedImport
}
