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
	r1 := customTrayIconImportAPI().SysCallN(8, GetObjectUintptr(TheOwner))
	return AsCustomTrayIcon(r1)
}

func (m *TCustomTrayIcon) Animate() bool {
	r1 := customTrayIconImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrayIcon) SetAnimate(AValue bool) {
	customTrayIconImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrayIcon) AnimateInterval() uint32 {
	r1 := customTrayIconImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomTrayIcon) SetAnimateInterval(AValue uint32) {
	customTrayIconImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrayIcon) BalloonFlags() TBalloonFlags {
	r1 := customTrayIconImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TBalloonFlags(r1)
}

func (m *TCustomTrayIcon) SetBalloonFlags(AValue TBalloonFlags) {
	customTrayIconImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrayIcon) BalloonHint() string {
	r1 := customTrayIconImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTrayIcon) SetBalloonHint(AValue string) {
	customTrayIconImportAPI().SysCallN(3, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTrayIcon) BalloonTimeout() int32 {
	r1 := customTrayIconImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTrayIcon) SetBalloonTimeout(AValue int32) {
	customTrayIconImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTrayIcon) BalloonTitle() string {
	r1 := customTrayIconImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTrayIcon) SetBalloonTitle(AValue string) {
	customTrayIconImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTrayIcon) Canvas() ICanvas {
	r1 := customTrayIconImportAPI().SysCallN(6, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomTrayIcon) PopUpMenu() IPopupMenu {
	r1 := customTrayIconImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return AsPopupMenu(r1)
}

func (m *TCustomTrayIcon) SetPopUpMenu(AValue IPopupMenu) {
	customTrayIconImportAPI().SysCallN(15, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTrayIcon) Icon() IIcon {
	r1 := customTrayIconImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return AsIcon(r1)
}

func (m *TCustomTrayIcon) SetIcon(AValue IIcon) {
	customTrayIconImportAPI().SysCallN(12, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTrayIcon) Icons() ICustomImageList {
	r1 := customTrayIconImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomTrayIcon) SetIcons(AValue ICustomImageList) {
	customTrayIconImportAPI().SysCallN(13, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTrayIcon) Hint() string {
	r1 := customTrayIconImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTrayIcon) SetHint(AValue string) {
	customTrayIconImportAPI().SysCallN(11, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTrayIcon) ShowIcon() bool {
	r1 := customTrayIconImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrayIcon) SetShowIcon(AValue bool) {
	customTrayIconImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrayIcon) Visible() bool {
	r1 := customTrayIconImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTrayIcon) SetVisible(AValue bool) {
	customTrayIconImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTrayIcon) Hide() bool {
	r1 := customTrayIconImportAPI().SysCallN(10, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTrayIcon) Show() bool {
	r1 := customTrayIconImportAPI().SysCallN(22, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTrayIcon) GetPosition() (resultPoint TPoint) {
	customTrayIconImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func CustomTrayIconClass() TClass {
	ret := customTrayIconImportAPI().SysCallN(7)
	return TClass(ret)
}

func (m *TCustomTrayIcon) InternalUpdate() {
	customTrayIconImportAPI().SysCallN(14, m.Instance())
}

func (m *TCustomTrayIcon) ShowBalloonHint() {
	customTrayIconImportAPI().SysCallN(23, m.Instance())
}

func (m *TCustomTrayIcon) SetOnClick(fn TNotifyEvent) {
	if m.clickPtr != 0 {
		RemoveEventElement(m.clickPtr)
	}
	m.clickPtr = MakeEventDataPtr(fn)
	customTrayIconImportAPI().SysCallN(16, m.Instance(), m.clickPtr)
}

func (m *TCustomTrayIcon) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	customTrayIconImportAPI().SysCallN(17, m.Instance(), m.dblClickPtr)
}

func (m *TCustomTrayIcon) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	customTrayIconImportAPI().SysCallN(18, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomTrayIcon) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	customTrayIconImportAPI().SysCallN(20, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomTrayIcon) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	customTrayIconImportAPI().SysCallN(19, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomTrayIcon) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	customTrayIconImportAPI().SysCallN(21, m.Instance(), m.paintPtr)
}

var (
	customTrayIconImport       *imports.Imports = nil
	customTrayIconImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomTrayIcon_Animate", 0),
		/*1*/ imports.NewTable("CustomTrayIcon_AnimateInterval", 0),
		/*2*/ imports.NewTable("CustomTrayIcon_BalloonFlags", 0),
		/*3*/ imports.NewTable("CustomTrayIcon_BalloonHint", 0),
		/*4*/ imports.NewTable("CustomTrayIcon_BalloonTimeout", 0),
		/*5*/ imports.NewTable("CustomTrayIcon_BalloonTitle", 0),
		/*6*/ imports.NewTable("CustomTrayIcon_Canvas", 0),
		/*7*/ imports.NewTable("CustomTrayIcon_Class", 0),
		/*8*/ imports.NewTable("CustomTrayIcon_Create", 0),
		/*9*/ imports.NewTable("CustomTrayIcon_GetPosition", 0),
		/*10*/ imports.NewTable("CustomTrayIcon_Hide", 0),
		/*11*/ imports.NewTable("CustomTrayIcon_Hint", 0),
		/*12*/ imports.NewTable("CustomTrayIcon_Icon", 0),
		/*13*/ imports.NewTable("CustomTrayIcon_Icons", 0),
		/*14*/ imports.NewTable("CustomTrayIcon_InternalUpdate", 0),
		/*15*/ imports.NewTable("CustomTrayIcon_PopUpMenu", 0),
		/*16*/ imports.NewTable("CustomTrayIcon_SetOnClick", 0),
		/*17*/ imports.NewTable("CustomTrayIcon_SetOnDblClick", 0),
		/*18*/ imports.NewTable("CustomTrayIcon_SetOnMouseDown", 0),
		/*19*/ imports.NewTable("CustomTrayIcon_SetOnMouseMove", 0),
		/*20*/ imports.NewTable("CustomTrayIcon_SetOnMouseUp", 0),
		/*21*/ imports.NewTable("CustomTrayIcon_SetOnPaint", 0),
		/*22*/ imports.NewTable("CustomTrayIcon_Show", 0),
		/*23*/ imports.NewTable("CustomTrayIcon_ShowBalloonHint", 0),
		/*24*/ imports.NewTable("CustomTrayIcon_ShowIcon", 0),
		/*25*/ imports.NewTable("CustomTrayIcon_Visible", 0),
	}
)

func customTrayIconImportAPI() *imports.Imports {
	if customTrayIconImport == nil {
		customTrayIconImport = NewDefaultImports()
		customTrayIconImport.SetImportTable(customTrayIconImportTables)
		customTrayIconImportTables = nil
	}
	return customTrayIconImport
}
