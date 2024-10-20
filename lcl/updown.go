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

// IUpDown Parent: ICustomUpDown
type IUpDown interface {
	ICustomUpDown
	AlignButton() TUDAlignButton                    // property
	SetAlignButton(AValue TUDAlignButton)           // property
	ArrowKeys() bool                                // property
	SetArrowKeys(AValue bool)                       // property
	Associate() IWinControl                         // property
	SetAssociate(AValue IWinControl)                // property
	Increment() int32                               // property
	SetIncrement(AValue int32)                      // property
	Max() SmallInt                                  // property
	SetMax(AValue SmallInt)                         // property
	Min() SmallInt                                  // property
	SetMin(AValue SmallInt)                         // property
	MinRepeatInterval() Byte                        // property
	SetMinRepeatInterval(AValue Byte)               // property
	Orientation() TUDOrientation                    // property
	SetOrientation(AValue TUDOrientation)           // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	Position() SmallInt                             // property
	SetPosition(AValue SmallInt)                    // property
	Thousands() bool                                // property
	SetThousands(AValue bool)                       // property
	Flat() bool                                     // property
	SetFlat(AValue bool)                            // property
	Wrap() bool                                     // property
	SetWrap(AValue bool)                            // property
	SetOnChanging(fn TUDChangingEvent)              // property event
	SetOnChangingEx(fn TUDChangingEventEx)          // property event
	SetOnClickForUDClickEvent(fn TUDClickEvent)     // property event
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

// TUpDown Parent: TCustomUpDown
type TUpDown struct {
	TCustomUpDown
	changingPtr             uintptr
	changingExPtr           uintptr
	clickForUDClickEventPtr uintptr
	contextPopupPtr         uintptr
	mouseDownPtr            uintptr
	mouseEnterPtr           uintptr
	mouseLeavePtr           uintptr
	mouseMovePtr            uintptr
	mouseUpPtr              uintptr
	mouseWheelPtr           uintptr
	mouseWheelDownPtr       uintptr
	mouseWheelUpPtr         uintptr
	mouseWheelHorzPtr       uintptr
	mouseWheelLeftPtr       uintptr
	mouseWheelRightPtr      uintptr
}

func NewUpDown(AOwner IComponent) IUpDown {
	r1 := upDownImportAPI().SysCallN(4, GetObjectUintptr(AOwner))
	return AsUpDown(r1)
}

func (m *TUpDown) AlignButton() TUDAlignButton {
	r1 := upDownImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TUDAlignButton(r1)
}

func (m *TUpDown) SetAlignButton(AValue TUDAlignButton) {
	upDownImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) ArrowKeys() bool {
	r1 := upDownImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetArrowKeys(AValue bool) {
	upDownImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) Associate() IWinControl {
	r1 := upDownImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TUpDown) SetAssociate(AValue IWinControl) {
	upDownImportAPI().SysCallN(2, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TUpDown) Increment() int32 {
	r1 := upDownImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TUpDown) SetIncrement(AValue int32) {
	upDownImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) Max() SmallInt {
	r1 := upDownImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TUpDown) SetMax(AValue SmallInt) {
	upDownImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) Min() SmallInt {
	r1 := upDownImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TUpDown) SetMin(AValue SmallInt) {
	upDownImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) MinRepeatInterval() Byte {
	r1 := upDownImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TUpDown) SetMinRepeatInterval(AValue Byte) {
	upDownImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) Orientation() TUDOrientation {
	r1 := upDownImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TUDOrientation(r1)
}

func (m *TUpDown) SetOrientation(AValue TUDOrientation) {
	upDownImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) ParentColor() bool {
	r1 := upDownImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetParentColor(AValue bool) {
	upDownImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) ParentShowHint() bool {
	r1 := upDownImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetParentShowHint(AValue bool) {
	upDownImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) Position() SmallInt {
	r1 := upDownImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TUpDown) SetPosition(AValue SmallInt) {
	upDownImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) Thousands() bool {
	r1 := upDownImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetThousands(AValue bool) {
	upDownImportAPI().SysCallN(29, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) Flat() bool {
	r1 := upDownImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetFlat(AValue bool) {
	upDownImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) Wrap() bool {
	r1 := upDownImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetWrap(AValue bool) {
	upDownImportAPI().SysCallN(30, 1, m.Instance(), PascalBool(AValue))
}

func UpDownClass() TClass {
	ret := upDownImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TUpDown) SetOnChanging(fn TUDChangingEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(14, m.Instance(), m.changingPtr)
}

func (m *TUpDown) SetOnChangingEx(fn TUDChangingEventEx) {
	if m.changingExPtr != 0 {
		RemoveEventElement(m.changingExPtr)
	}
	m.changingExPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(15, m.Instance(), m.changingExPtr)
}

func (m *TUpDown) SetOnClickForUDClickEvent(fn TUDClickEvent) {
	if m.clickForUDClickEventPtr != 0 {
		RemoveEventElement(m.clickForUDClickEventPtr)
	}
	m.clickForUDClickEventPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(16, m.Instance(), m.clickForUDClickEventPtr)
}

func (m *TUpDown) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(17, m.Instance(), m.contextPopupPtr)
}

func (m *TUpDown) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(18, m.Instance(), m.mouseDownPtr)
}

func (m *TUpDown) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(19, m.Instance(), m.mouseEnterPtr)
}

func (m *TUpDown) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(20, m.Instance(), m.mouseLeavePtr)
}

func (m *TUpDown) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(21, m.Instance(), m.mouseMovePtr)
}

func (m *TUpDown) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(22, m.Instance(), m.mouseUpPtr)
}

func (m *TUpDown) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(23, m.Instance(), m.mouseWheelPtr)
}

func (m *TUpDown) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(24, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TUpDown) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(28, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TUpDown) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(25, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TUpDown) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(26, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TUpDown) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	upDownImportAPI().SysCallN(27, m.Instance(), m.mouseWheelRightPtr)
}

var (
	upDownImport       *imports.Imports = nil
	upDownImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("UpDown_AlignButton", 0),
		/*1*/ imports.NewTable("UpDown_ArrowKeys", 0),
		/*2*/ imports.NewTable("UpDown_Associate", 0),
		/*3*/ imports.NewTable("UpDown_Class", 0),
		/*4*/ imports.NewTable("UpDown_Create", 0),
		/*5*/ imports.NewTable("UpDown_Flat", 0),
		/*6*/ imports.NewTable("UpDown_Increment", 0),
		/*7*/ imports.NewTable("UpDown_Max", 0),
		/*8*/ imports.NewTable("UpDown_Min", 0),
		/*9*/ imports.NewTable("UpDown_MinRepeatInterval", 0),
		/*10*/ imports.NewTable("UpDown_Orientation", 0),
		/*11*/ imports.NewTable("UpDown_ParentColor", 0),
		/*12*/ imports.NewTable("UpDown_ParentShowHint", 0),
		/*13*/ imports.NewTable("UpDown_Position", 0),
		/*14*/ imports.NewTable("UpDown_SetOnChanging", 0),
		/*15*/ imports.NewTable("UpDown_SetOnChangingEx", 0),
		/*16*/ imports.NewTable("UpDown_SetOnClickForUDClickEvent", 0),
		/*17*/ imports.NewTable("UpDown_SetOnContextPopup", 0),
		/*18*/ imports.NewTable("UpDown_SetOnMouseDown", 0),
		/*19*/ imports.NewTable("UpDown_SetOnMouseEnter", 0),
		/*20*/ imports.NewTable("UpDown_SetOnMouseLeave", 0),
		/*21*/ imports.NewTable("UpDown_SetOnMouseMove", 0),
		/*22*/ imports.NewTable("UpDown_SetOnMouseUp", 0),
		/*23*/ imports.NewTable("UpDown_SetOnMouseWheel", 0),
		/*24*/ imports.NewTable("UpDown_SetOnMouseWheelDown", 0),
		/*25*/ imports.NewTable("UpDown_SetOnMouseWheelHorz", 0),
		/*26*/ imports.NewTable("UpDown_SetOnMouseWheelLeft", 0),
		/*27*/ imports.NewTable("UpDown_SetOnMouseWheelRight", 0),
		/*28*/ imports.NewTable("UpDown_SetOnMouseWheelUp", 0),
		/*29*/ imports.NewTable("UpDown_Thousands", 0),
		/*30*/ imports.NewTable("UpDown_Wrap", 0),
	}
)

func upDownImportAPI() *imports.Imports {
	if upDownImport == nil {
		upDownImport = NewDefaultImports()
		upDownImport.SetImportTable(upDownImportTables)
		upDownImportTables = nil
	}
	return upDownImport
}
