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

// IFloatSpinEdit Parent: ICustomFloatSpinEdit
type IFloatSpinEdit interface {
	ICustomFloatSpinEdit
	AutoSelected() bool                             // property AutoSelected Getter
	SetAutoSelected(value bool)                     // property AutoSelected Setter
	AutoSelect() bool                               // property AutoSelect Getter
	SetAutoSelect(value bool)                       // property AutoSelect Setter
	ParentColor() bool                              // property ParentColor Getter
	SetParentColor(value bool)                      // property ParentColor Setter
	ParentFont() bool                               // property ParentFont Getter
	SetParentFont(value bool)                       // property ParentFont Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	SetOnEditingDone(fn TNotifyEvent)               // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
}

type TFloatSpinEdit struct {
	TCustomFloatSpinEdit
}

func (m *TFloatSpinEdit) AutoSelected() bool {
	if !m.IsValid() {
		return false
	}
	r := floatSpinEditAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFloatSpinEdit) SetAutoSelected(value bool) {
	if !m.IsValid() {
		return
	}
	floatSpinEditAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TFloatSpinEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := floatSpinEditAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFloatSpinEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	floatSpinEditAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TFloatSpinEdit) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := floatSpinEditAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFloatSpinEdit) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	floatSpinEditAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TFloatSpinEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := floatSpinEditAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFloatSpinEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	floatSpinEditAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TFloatSpinEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := floatSpinEditAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFloatSpinEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	floatSpinEditAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TFloatSpinEdit) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 7, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 10, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 11, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 12, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 13, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 14, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 15, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFloatSpinEdit) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, floatSpinEditAPI(), api.MakeEventDataPtr(cb))
}

// NewFloatSpinEdit class constructor
func NewFloatSpinEdit(theOwner IComponent) IFloatSpinEdit {
	r := floatSpinEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsFloatSpinEdit(r)
}

func TFloatSpinEditClass() types.TClass {
	r := floatSpinEditAPI().SysCallN(18)
	return types.TClass(r)
}

var (
	floatSpinEditOnce   base.Once
	floatSpinEditImport *imports.Imports = nil
)

func floatSpinEditAPI() *imports.Imports {
	floatSpinEditOnce.Do(func() {
		floatSpinEditImport = api.NewDefaultImports()
		floatSpinEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFloatSpinEdit_Create", 0), // constructor NewFloatSpinEdit
			/* 1 */ imports.NewTable("TFloatSpinEdit_AutoSelected", 0), // property AutoSelected
			/* 2 */ imports.NewTable("TFloatSpinEdit_AutoSelect", 0), // property AutoSelect
			/* 3 */ imports.NewTable("TFloatSpinEdit_ParentColor", 0), // property ParentColor
			/* 4 */ imports.NewTable("TFloatSpinEdit_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TFloatSpinEdit_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TFloatSpinEdit_OnEditingDone", 0), // event OnEditingDone
			/* 7 */ imports.NewTable("TFloatSpinEdit_OnMouseDown", 0), // event OnMouseDown
			/* 8 */ imports.NewTable("TFloatSpinEdit_OnMouseEnter", 0), // event OnMouseEnter
			/* 9 */ imports.NewTable("TFloatSpinEdit_OnMouseLeave", 0), // event OnMouseLeave
			/* 10 */ imports.NewTable("TFloatSpinEdit_OnMouseMove", 0), // event OnMouseMove
			/* 11 */ imports.NewTable("TFloatSpinEdit_OnMouseUp", 0), // event OnMouseUp
			/* 12 */ imports.NewTable("TFloatSpinEdit_OnMouseWheel", 0), // event OnMouseWheel
			/* 13 */ imports.NewTable("TFloatSpinEdit_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 14 */ imports.NewTable("TFloatSpinEdit_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 15 */ imports.NewTable("TFloatSpinEdit_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 16 */ imports.NewTable("TFloatSpinEdit_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 17 */ imports.NewTable("TFloatSpinEdit_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 18 */ imports.NewTable("TFloatSpinEdit_TClass", 0), // function TFloatSpinEditClass
		}
	})
	return floatSpinEditImport
}
