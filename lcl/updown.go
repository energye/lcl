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

// IUpDown Parent: ICustomUpDown
type IUpDown interface {
	ICustomUpDown
	AlignButton() types.TUDAlignButton              // property AlignButton Getter
	SetAlignButton(value types.TUDAlignButton)      // property AlignButton Setter
	ArrowKeys() bool                                // property ArrowKeys Getter
	SetArrowKeys(value bool)                        // property ArrowKeys Setter
	Associate() IWinControl                         // property Associate Getter
	SetAssociate(value IWinControl)                 // property Associate Setter
	Increment() int32                               // property Increment Getter
	SetIncrement(value int32)                       // property Increment Setter
	Max() types.SmallInt                            // property Max Getter
	SetMax(value types.SmallInt)                    // property Max Setter
	Min() types.SmallInt                            // property Min Getter
	SetMin(value types.SmallInt)                    // property Min Setter
	MinRepeatInterval() byte                        // property MinRepeatInterval Getter
	SetMinRepeatInterval(value byte)                // property MinRepeatInterval Setter
	Orientation() types.TUDOrientation              // property Orientation Getter
	SetOrientation(value types.TUDOrientation)      // property Orientation Setter
	ParentColor() bool                              // property ParentColor Getter
	SetParentColor(value bool)                      // property ParentColor Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	Position() types.SmallInt                       // property Position Getter
	SetPosition(value types.SmallInt)               // property Position Setter
	Thousands() bool                                // property Thousands Getter
	SetThousands(value bool)                        // property Thousands Setter
	Flat() bool                                     // property Flat Getter
	SetFlat(value bool)                             // property Flat Setter
	Wrap() bool                                     // property Wrap Getter
	SetWrap(value bool)                             // property Wrap Setter
	SetOnChanging(fn TUDChangingEvent)              // property event
	SetOnChangingEx(fn TUDChangingEventEx)          // property event
	SetOnUDClick(fn TUDClickEvent)                  // property event
	SetOnContextPopup(fn TContextPopupEvent)        // property event
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

type TUpDown struct {
	TCustomUpDown
}

func (m *TUpDown) AlignButton() types.TUDAlignButton {
	if !m.IsValid() {
		return 0
	}
	r := upDownAPI().SysCallN(1, 0, m.Instance())
	return types.TUDAlignButton(r)
}

func (m *TUpDown) SetAlignButton(value types.TUDAlignButton) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TUpDown) ArrowKeys() bool {
	if !m.IsValid() {
		return false
	}
	r := upDownAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TUpDown) SetArrowKeys(value bool) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TUpDown) Associate() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := upDownAPI().SysCallN(3, 0, m.Instance())
	return AsWinControl(r)
}

func (m *TUpDown) SetAssociate(value IWinControl) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TUpDown) Increment() int32 {
	if !m.IsValid() {
		return 0
	}
	r := upDownAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TUpDown) SetIncrement(value int32) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TUpDown) Max() types.SmallInt {
	if !m.IsValid() {
		return 0
	}
	r := upDownAPI().SysCallN(5, 0, m.Instance())
	return types.SmallInt(r)
}

func (m *TUpDown) SetMax(value types.SmallInt) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TUpDown) Min() types.SmallInt {
	if !m.IsValid() {
		return 0
	}
	r := upDownAPI().SysCallN(6, 0, m.Instance())
	return types.SmallInt(r)
}

func (m *TUpDown) SetMin(value types.SmallInt) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TUpDown) MinRepeatInterval() byte {
	if !m.IsValid() {
		return 0
	}
	r := upDownAPI().SysCallN(7, 0, m.Instance())
	return byte(r)
}

func (m *TUpDown) SetMinRepeatInterval(value byte) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TUpDown) Orientation() types.TUDOrientation {
	if !m.IsValid() {
		return 0
	}
	r := upDownAPI().SysCallN(8, 0, m.Instance())
	return types.TUDOrientation(r)
}

func (m *TUpDown) SetOrientation(value types.TUDOrientation) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TUpDown) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := upDownAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TUpDown) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TUpDown) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := upDownAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TUpDown) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TUpDown) Position() types.SmallInt {
	if !m.IsValid() {
		return 0
	}
	r := upDownAPI().SysCallN(11, 0, m.Instance())
	return types.SmallInt(r)
}

func (m *TUpDown) SetPosition(value types.SmallInt) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TUpDown) Thousands() bool {
	if !m.IsValid() {
		return false
	}
	r := upDownAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TUpDown) SetThousands(value bool) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TUpDown) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := upDownAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TUpDown) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TUpDown) Wrap() bool {
	if !m.IsValid() {
		return false
	}
	r := upDownAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TUpDown) SetWrap(value bool) {
	if !m.IsValid() {
		return
	}
	upDownAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TUpDown) SetOnChanging(fn TUDChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUDChangingEvent(fn)
	base.SetEvent(m, 15, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnChangingEx(fn TUDChangingEventEx) {
	if !m.IsValid() {
		return
	}
	cb := makeTUDChangingEventEx(fn)
	base.SetEvent(m, 16, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnUDClick(fn TUDClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUDClickEvent(fn)
	base.SetEvent(m, 17, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 18, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 19, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 21, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 22, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 23, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 24, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 25, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 26, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 27, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 28, upDownAPI(), api.MakeEventDataPtr(cb))
}

func (m *TUpDown) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 29, upDownAPI(), api.MakeEventDataPtr(cb))
}

// NewUpDown class constructor
func NewUpDown(owner IComponent) IUpDown {
	r := upDownAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsUpDown(r)
}

func TUpDownClass() types.TClass {
	r := upDownAPI().SysCallN(30)
	return types.TClass(r)
}

var (
	upDownOnce   base.Once
	upDownImport *imports.Imports = nil
)

func upDownAPI() *imports.Imports {
	upDownOnce.Do(func() {
		upDownImport = api.NewDefaultImports()
		upDownImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TUpDown_Create", 0), // constructor NewUpDown
			/* 1 */ imports.NewTable("TUpDown_AlignButton", 0), // property AlignButton
			/* 2 */ imports.NewTable("TUpDown_ArrowKeys", 0), // property ArrowKeys
			/* 3 */ imports.NewTable("TUpDown_Associate", 0), // property Associate
			/* 4 */ imports.NewTable("TUpDown_Increment", 0), // property Increment
			/* 5 */ imports.NewTable("TUpDown_Max", 0), // property Max
			/* 6 */ imports.NewTable("TUpDown_Min", 0), // property Min
			/* 7 */ imports.NewTable("TUpDown_MinRepeatInterval", 0), // property MinRepeatInterval
			/* 8 */ imports.NewTable("TUpDown_Orientation", 0), // property Orientation
			/* 9 */ imports.NewTable("TUpDown_ParentColor", 0), // property ParentColor
			/* 10 */ imports.NewTable("TUpDown_ParentShowHint", 0), // property ParentShowHint
			/* 11 */ imports.NewTable("TUpDown_Position", 0), // property Position
			/* 12 */ imports.NewTable("TUpDown_Thousands", 0), // property Thousands
			/* 13 */ imports.NewTable("TUpDown_Flat", 0), // property Flat
			/* 14 */ imports.NewTable("TUpDown_Wrap", 0), // property Wrap
			/* 15 */ imports.NewTable("TUpDown_OnChanging", 0), // event OnChanging
			/* 16 */ imports.NewTable("TUpDown_OnChangingEx", 0), // event OnChangingEx
			/* 17 */ imports.NewTable("TUpDown_OnClickToUDClickEvent", 0), // event OnUDClick
			/* 18 */ imports.NewTable("TUpDown_OnContextPopup", 0), // event OnContextPopup
			/* 19 */ imports.NewTable("TUpDown_OnMouseDown", 0), // event OnMouseDown
			/* 20 */ imports.NewTable("TUpDown_OnMouseEnter", 0), // event OnMouseEnter
			/* 21 */ imports.NewTable("TUpDown_OnMouseLeave", 0), // event OnMouseLeave
			/* 22 */ imports.NewTable("TUpDown_OnMouseMove", 0), // event OnMouseMove
			/* 23 */ imports.NewTable("TUpDown_OnMouseUp", 0), // event OnMouseUp
			/* 24 */ imports.NewTable("TUpDown_OnMouseWheel", 0), // event OnMouseWheel
			/* 25 */ imports.NewTable("TUpDown_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 26 */ imports.NewTable("TUpDown_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 27 */ imports.NewTable("TUpDown_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 28 */ imports.NewTable("TUpDown_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 29 */ imports.NewTable("TUpDown_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 30 */ imports.NewTable("TUpDown_TClass", 0), // function TUpDownClass
		}
	})
	return upDownImport
}
