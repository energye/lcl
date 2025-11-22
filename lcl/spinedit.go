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

// ISpinEdit Parent: ICustomSpinEdit
type ISpinEdit interface {
	ICustomSpinEdit
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

type TSpinEdit struct {
	TCustomSpinEdit
}

func (m *TSpinEdit) AutoSelected() bool {
	if !m.IsValid() {
		return false
	}
	r := spinEditAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSpinEdit) SetAutoSelected(value bool) {
	if !m.IsValid() {
		return
	}
	spinEditAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TSpinEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := spinEditAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSpinEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	spinEditAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TSpinEdit) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := spinEditAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSpinEdit) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	spinEditAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TSpinEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := spinEditAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSpinEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	spinEditAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TSpinEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := spinEditAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSpinEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	spinEditAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TSpinEdit) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 7, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 10, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 11, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 12, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 13, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 14, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 15, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, spinEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSpinEdit) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 17, spinEditAPI(), api.MakeEventDataPtr(cb))
}

// NewSpinEdit class constructor
func NewSpinEdit(theOwner IComponent) ISpinEdit {
	r := spinEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsSpinEdit(r)
}

func TSpinEditClass() types.TClass {
	r := spinEditAPI().SysCallN(18)
	return types.TClass(r)
}

var (
	spinEditOnce   base.Once
	spinEditImport *imports.Imports = nil
)

func spinEditAPI() *imports.Imports {
	spinEditOnce.Do(func() {
		spinEditImport = api.NewDefaultImports()
		spinEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSpinEdit_Create", 0), // constructor NewSpinEdit
			/* 1 */ imports.NewTable("TSpinEdit_AutoSelected", 0), // property AutoSelected
			/* 2 */ imports.NewTable("TSpinEdit_AutoSelect", 0), // property AutoSelect
			/* 3 */ imports.NewTable("TSpinEdit_ParentColor", 0), // property ParentColor
			/* 4 */ imports.NewTable("TSpinEdit_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TSpinEdit_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TSpinEdit_OnEditingDone", 0), // event OnEditingDone
			/* 7 */ imports.NewTable("TSpinEdit_OnMouseDown", 0), // event OnMouseDown
			/* 8 */ imports.NewTable("TSpinEdit_OnMouseEnter", 0), // event OnMouseEnter
			/* 9 */ imports.NewTable("TSpinEdit_OnMouseLeave", 0), // event OnMouseLeave
			/* 10 */ imports.NewTable("TSpinEdit_OnMouseMove", 0), // event OnMouseMove
			/* 11 */ imports.NewTable("TSpinEdit_OnMouseUp", 0), // event OnMouseUp
			/* 12 */ imports.NewTable("TSpinEdit_OnMouseWheel", 0), // event OnMouseWheel
			/* 13 */ imports.NewTable("TSpinEdit_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 14 */ imports.NewTable("TSpinEdit_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 15 */ imports.NewTable("TSpinEdit_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 16 */ imports.NewTable("TSpinEdit_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 17 */ imports.NewTable("TSpinEdit_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 18 */ imports.NewTable("TSpinEdit_TClass", 0), // function TSpinEditClass
		}
	})
	return spinEditImport
}
