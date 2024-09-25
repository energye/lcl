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
	. "github.com/energye/lcl/types"
)

// IXButton Parent: IGraphicControl
type IXButton interface {
	IGraphicControl
	Caption() string                   // property
	SetCaption(AValue string)          // property
	ShowCaption() bool                 // property
	SetShowCaption(AValue bool)        // property
	BackColor() TColor                 // property
	SetBackColor(AValue TColor)        // property
	HoverColor() TColor                // property
	SetHoverColor(AValue TColor)       // property
	DownColor() TColor                 // property
	SetDownColor(AValue TColor)        // property
	BorderWidth() int32                // property
	SetBorderWidth(AValue int32)       // property
	BorderColor() TColor               // property
	SetBorderColor(AValue TColor)      // property
	Picture() IPicture                 // property
	SetPicture(AValue IPicture)        // property
	DrawMode() TDrawImageMode          // property
	SetDrawMode(AValue TDrawImageMode) // property
	NormalFontColor() TColor           // property
	SetNormalFontColor(AValue TColor)  // property
	DownFontColor() TColor             // property
	SetDownFontColor(AValue TColor)    // property
	HoverFontColor() TColor            // property
	SetHoverFontColor(AValue TColor)   // property
	ParentFont() bool                  // property
	SetParentFont(AValue bool)         // property
	ParentShowHint() bool              // property
	SetParentShowHint(AValue bool)     // property
	Paint()                            // procedure
	Resize()                           // procedure
	SetOnDblClick(fn TNotifyEvent)     // property event
	SetOnMouseDown(fn TMouseEvent)     // property event
	SetOnMouseEnter(fn TNotifyEvent)   // property event
	SetOnMouseLeave(fn TNotifyEvent)   // property event
	SetOnMouseMove(fn TMouseMoveEvent) // property event
	SetOnMouseUp(fn TMouseEvent)       // property event
}

// TXButton Parent: TGraphicControl
type TXButton struct {
	TGraphicControl
	dblClickPtr   uintptr
	mouseDownPtr  uintptr
	mouseEnterPtr uintptr
	mouseLeavePtr uintptr
	mouseMovePtr  uintptr
	mouseUpPtr    uintptr
}

func NewXButton(AOwner IComponent) IXButton {
	r1 := LCL().SysCallN(6152, GetObjectUintptr(AOwner))
	return AsXButton(r1)
}

func (m *TXButton) Caption() string {
	r1 := LCL().SysCallN(6150, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TXButton) SetCaption(AValue string) {
	LCL().SysCallN(6150, 1, m.Instance(), PascalStr(AValue))
}

func (m *TXButton) ShowCaption() bool {
	r1 := LCL().SysCallN(6170, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetShowCaption(AValue bool) {
	LCL().SysCallN(6170, 1, m.Instance(), PascalBool(AValue))
}

func (m *TXButton) BackColor() TColor {
	r1 := LCL().SysCallN(6147, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetBackColor(AValue TColor) {
	LCL().SysCallN(6147, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) HoverColor() TColor {
	r1 := LCL().SysCallN(6156, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetHoverColor(AValue TColor) {
	LCL().SysCallN(6156, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) DownColor() TColor {
	r1 := LCL().SysCallN(6153, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetDownColor(AValue TColor) {
	LCL().SysCallN(6153, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) BorderWidth() int32 {
	r1 := LCL().SysCallN(6149, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TXButton) SetBorderWidth(AValue int32) {
	LCL().SysCallN(6149, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) BorderColor() TColor {
	r1 := LCL().SysCallN(6148, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetBorderColor(AValue TColor) {
	LCL().SysCallN(6148, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) Picture() IPicture {
	r1 := LCL().SysCallN(6162, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TXButton) SetPicture(AValue IPicture) {
	LCL().SysCallN(6162, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TXButton) DrawMode() TDrawImageMode {
	r1 := LCL().SysCallN(6155, 0, m.Instance(), 0)
	return TDrawImageMode(r1)
}

func (m *TXButton) SetDrawMode(AValue TDrawImageMode) {
	LCL().SysCallN(6155, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) NormalFontColor() TColor {
	r1 := LCL().SysCallN(6158, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetNormalFontColor(AValue TColor) {
	LCL().SysCallN(6158, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) DownFontColor() TColor {
	r1 := LCL().SysCallN(6154, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetDownFontColor(AValue TColor) {
	LCL().SysCallN(6154, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) HoverFontColor() TColor {
	r1 := LCL().SysCallN(6157, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetHoverFontColor(AValue TColor) {
	LCL().SysCallN(6157, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) ParentFont() bool {
	r1 := LCL().SysCallN(6160, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetParentFont(AValue bool) {
	LCL().SysCallN(6160, 1, m.Instance(), PascalBool(AValue))
}

func (m *TXButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(6161, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(6161, 1, m.Instance(), PascalBool(AValue))
}

func XButtonClass() TClass {
	ret := LCL().SysCallN(6151)
	return TClass(ret)
}

func (m *TXButton) Paint() {
	LCL().SysCallN(6159, m.Instance())
}

func (m *TXButton) Resize() {
	LCL().SysCallN(6163, m.Instance())
}

func (m *TXButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6164, m.Instance(), m.dblClickPtr)
}

func (m *TXButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6165, m.Instance(), m.mouseDownPtr)
}

func (m *TXButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6166, m.Instance(), m.mouseEnterPtr)
}

func (m *TXButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6167, m.Instance(), m.mouseLeavePtr)
}

func (m *TXButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6168, m.Instance(), m.mouseMovePtr)
}

func (m *TXButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6169, m.Instance(), m.mouseUpPtr)
}
