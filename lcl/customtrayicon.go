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

// ICustomTrayIcon Parent: ILCLComponent
type ICustomTrayIcon interface {
	ILCLComponent
	Animate() bool                        // property
	SetAnimate(AValue bool)               // property
	AnimateInterval() uint32              // property
	SetAnimateInterval(AValue uint32)     // property
	BalloonFlags() TBalloonFlags          // property
	SetBalloonFlags(AValue TBalloonFlags) // property
	BalloonHint() string                  // property
	SetBalloonHint(AValue string)         // property
	BalloonTimeout() int32                // property
	SetBalloonTimeout(AValue int32)       // property
	BalloonTitle() string                 // property
	SetBalloonTitle(AValue string)        // property
	Canvas() ICanvas                      // property
	PopUpMenu() IPopupMenu                // property
	SetPopUpMenu(AValue IPopupMenu)       // property
	Icon() IIcon                          // property
	SetIcon(AValue IIcon)                 // property
	Icons() ICustomImageList              // property
	SetIcons(AValue ICustomImageList)     // property
	Hint() string                         // property
	SetHint(AValue string)                // property
	ShowIcon() bool                       // property
	SetShowIcon(AValue bool)              // property
	Visible() bool                        // property
	SetVisible(AValue bool)               // property
	Hide() bool                           // function
	Show() bool                           // function
	GetPosition() (resultPoint TPoint)    // function
	InternalUpdate()                      // procedure
	ShowBalloonHint()                     // procedure
	SetOnClick(fn TNotifyEvent)           // property event
	SetOnDblClick(fn TNotifyEvent)        // property event
	SetOnMouseDown(fn TMouseEvent)        // property event
	SetOnMouseUp(fn TMouseEvent)          // property event
	SetOnMouseMove(fn TMouseMoveEvent)    // property event
	SetOnPaint(fn TNotifyEvent)           // property event
}

// TCustomTrayIcon Parent: TLCLComponent
type TCustomTrayIcon struct {
	TLCLComponent
	clickPtr     uintptr
	dblClickPtr  uintptr
	mouseDownPtr uintptr
	mouseUpPtr   uintptr
	mouseMovePtr uintptr
	paintPtr     uintptr
}

func NewCustomTrayIcon(TheOwner IComponent) ICustomTrayIcon {
	r1 := LCL().SysCallN(2413, GetObjectUintptr(TheOwner))
	return AsCustomTrayIcon(r1)
}

func (m *TCustomTrayIcon) Animate() bool {
	r1 := LCL().SysCallN(2405, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrayIcon) SetAnimate(AValue bool) {
	LCL().SysCallN(2405, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrayIcon) AnimateInterval() uint32 {
	r1 := LCL().SysCallN(2406, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomTrayIcon) SetAnimateInterval(AValue uint32) {
	LCL().SysCallN(2406, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrayIcon) BalloonFlags() TBalloonFlags {
	r1 := LCL().SysCallN(2407, 0, m.Instance(), 0)
	return TBalloonFlags(r1)
}

func (m *TCustomTrayIcon) SetBalloonFlags(AValue TBalloonFlags) {
	LCL().SysCallN(2407, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrayIcon) BalloonHint() string {
	r1 := LCL().SysCallN(2408, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTrayIcon) SetBalloonHint(AValue string) {
	LCL().SysCallN(2408, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTrayIcon) BalloonTimeout() int32 {
	r1 := LCL().SysCallN(2409, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrayIcon) SetBalloonTimeout(AValue int32) {
	LCL().SysCallN(2409, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrayIcon) BalloonTitle() string {
	r1 := LCL().SysCallN(2410, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTrayIcon) SetBalloonTitle(AValue string) {
	LCL().SysCallN(2410, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTrayIcon) Canvas() ICanvas {
	r1 := LCL().SysCallN(2411, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomTrayIcon) PopUpMenu() IPopupMenu {
	r1 := LCL().SysCallN(2420, 0, m.Instance(), 0)
	return AsPopupMenu(r1)
}

func (m *TCustomTrayIcon) SetPopUpMenu(AValue IPopupMenu) {
	LCL().SysCallN(2420, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTrayIcon) Icon() IIcon {
	r1 := LCL().SysCallN(2417, 0, m.Instance(), 0)
	return AsIcon(r1)
}

func (m *TCustomTrayIcon) SetIcon(AValue IIcon) {
	LCL().SysCallN(2417, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTrayIcon) Icons() ICustomImageList {
	r1 := LCL().SysCallN(2418, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomTrayIcon) SetIcons(AValue ICustomImageList) {
	LCL().SysCallN(2418, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTrayIcon) Hint() string {
	r1 := LCL().SysCallN(2416, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTrayIcon) SetHint(AValue string) {
	LCL().SysCallN(2416, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTrayIcon) ShowIcon() bool {
	r1 := LCL().SysCallN(2429, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrayIcon) SetShowIcon(AValue bool) {
	LCL().SysCallN(2429, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrayIcon) Visible() bool {
	r1 := LCL().SysCallN(2430, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrayIcon) SetVisible(AValue bool) {
	LCL().SysCallN(2430, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrayIcon) Hide() bool {
	r1 := LCL().SysCallN(2415, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTrayIcon) Show() bool {
	r1 := LCL().SysCallN(2427, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTrayIcon) GetPosition() (resultPoint TPoint) {
	LCL().SysCallN(2414, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func CustomTrayIconClass() TClass {
	ret := LCL().SysCallN(2412)
	return TClass(ret)
}

func (m *TCustomTrayIcon) InternalUpdate() {
	LCL().SysCallN(2419, m.Instance())
}

func (m *TCustomTrayIcon) ShowBalloonHint() {
	LCL().SysCallN(2428, m.Instance())
}

func (m *TCustomTrayIcon) SetOnClick(fn TNotifyEvent) {
	if m.clickPtr != 0 {
		RemoveEventElement(m.clickPtr)
	}
	m.clickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2421, m.Instance(), m.clickPtr)
}

func (m *TCustomTrayIcon) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2422, m.Instance(), m.dblClickPtr)
}

func (m *TCustomTrayIcon) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2423, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomTrayIcon) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2425, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomTrayIcon) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2424, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomTrayIcon) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2426, m.Instance(), m.paintPtr)
}
