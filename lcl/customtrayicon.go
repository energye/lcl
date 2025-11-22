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

// ICustomTrayIcon Parent: ILCLComponent
type ICustomTrayIcon interface {
	ILCLComponent
	Hide() bool                // function
	Show() bool                // function
	GetPosition() types.TPoint // function
	InternalUpdate()           // procedure
	ShowBalloonHint()          // procedure
	// Animate
	//  Properties
	Animate() bool                             // property Animate Getter
	SetAnimate(value bool)                     // property Animate Setter
	AnimateInterval() uint32                   // property AnimateInterval Getter
	SetAnimateInterval(value uint32)           // property AnimateInterval Setter
	BalloonFlags() types.TBalloonFlags         // property BalloonFlags Getter
	SetBalloonFlags(value types.TBalloonFlags) // property BalloonFlags Setter
	BalloonHint() string                       // property BalloonHint Getter
	SetBalloonHint(value string)               // property BalloonHint Setter
	BalloonTimeout() int32                     // property BalloonTimeout Getter
	SetBalloonTimeout(value int32)             // property BalloonTimeout Setter
	BalloonTitle() string                      // property BalloonTitle Getter
	SetBalloonTitle(value string)              // property BalloonTitle Setter
	Canvas() ICanvas                           // property Canvas Getter
	PopUpMenu() IPopupMenu                     // property PopUpMenu Getter
	SetPopUpMenu(value IPopupMenu)             // property PopUpMenu Setter
	Icon() IIcon                               // property Icon Getter
	SetIcon(value IIcon)                       // property Icon Setter
	Icons() ICustomImageList                   // property Icons Getter
	SetIcons(value ICustomImageList)           // property Icons Setter
	Hint() string                              // property Hint Getter
	SetHint(value string)                      // property Hint Setter
	ShowIcon() bool                            // property ShowIcon Getter
	SetShowIcon(value bool)                    // property ShowIcon Setter
	Visible() bool                             // property Visible Getter
	SetVisible(value bool)                     // property Visible Setter
	SetOnClick(fn TNotifyEvent)                // property event
	SetOnDblClick(fn TNotifyEvent)             // property event
	SetOnMouseDown(fn TMouseEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)         // property event
	SetOnPaint(fn TNotifyEvent)                // property event
}

type TCustomTrayIcon struct {
	TLCLComponent
}

func (m *TCustomTrayIcon) Hide() bool {
	if !m.IsValid() {
		return false
	}
	r := customTrayIconAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTrayIcon) Show() bool {
	if !m.IsValid() {
		return false
	}
	r := customTrayIconAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTrayIcon) GetPosition() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomTrayIcon) InternalUpdate() {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(4, m.Instance())
}

func (m *TCustomTrayIcon) ShowBalloonHint() {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(5, m.Instance())
}

func (m *TCustomTrayIcon) Animate() bool {
	if !m.IsValid() {
		return false
	}
	r := customTrayIconAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTrayIcon) SetAnimate(value bool) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTrayIcon) AnimateInterval() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrayIconAPI().SysCallN(7, 0, m.Instance())
	return uint32(r)
}

func (m *TCustomTrayIcon) SetAnimateInterval(value uint32) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrayIcon) BalloonFlags() types.TBalloonFlags {
	if !m.IsValid() {
		return 0
	}
	r := customTrayIconAPI().SysCallN(8, 0, m.Instance())
	return types.TBalloonFlags(r)
}

func (m *TCustomTrayIcon) SetBalloonFlags(value types.TBalloonFlags) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrayIcon) BalloonHint() string {
	if !m.IsValid() {
		return ""
	}
	r := customTrayIconAPI().SysCallN(9, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTrayIcon) SetBalloonHint(value string) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTrayIcon) BalloonTimeout() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTrayIconAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTrayIcon) SetBalloonTimeout(value int32) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTrayIcon) BalloonTitle() string {
	if !m.IsValid() {
		return ""
	}
	r := customTrayIconAPI().SysCallN(11, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTrayIcon) SetBalloonTitle(value string) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(11, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTrayIcon) Canvas() ICanvas {
	if !m.IsValid() {
		return nil
	}
	r := customTrayIconAPI().SysCallN(12, m.Instance())
	return AsCanvas(r)
}

func (m *TCustomTrayIcon) PopUpMenu() IPopupMenu {
	if !m.IsValid() {
		return nil
	}
	r := customTrayIconAPI().SysCallN(13, 0, m.Instance())
	return AsPopupMenu(r)
}

func (m *TCustomTrayIcon) SetPopUpMenu(value IPopupMenu) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTrayIcon) Icon() IIcon {
	if !m.IsValid() {
		return nil
	}
	r := customTrayIconAPI().SysCallN(14, 0, m.Instance())
	return AsIcon(r)
}

func (m *TCustomTrayIcon) SetIcon(value IIcon) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTrayIcon) Icons() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customTrayIconAPI().SysCallN(15, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomTrayIcon) SetIcons(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(15, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTrayIcon) Hint() string {
	if !m.IsValid() {
		return ""
	}
	r := customTrayIconAPI().SysCallN(16, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTrayIcon) SetHint(value string) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(16, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTrayIcon) ShowIcon() bool {
	if !m.IsValid() {
		return false
	}
	r := customTrayIconAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTrayIcon) SetShowIcon(value bool) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTrayIcon) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := customTrayIconAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTrayIcon) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	customTrayIconAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTrayIcon) SetOnClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, customTrayIconAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTrayIcon) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, customTrayIconAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTrayIcon) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 21, customTrayIconAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTrayIcon) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 22, customTrayIconAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTrayIcon) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 23, customTrayIconAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTrayIcon) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 24, customTrayIconAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomTrayIcon class constructor
func NewCustomTrayIcon(theOwner IComponent) ICustomTrayIcon {
	r := customTrayIconAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomTrayIcon(r)
}

func TCustomTrayIconClass() types.TClass {
	r := customTrayIconAPI().SysCallN(25)
	return types.TClass(r)
}

var (
	customTrayIconOnce   base.Once
	customTrayIconImport *imports.Imports = nil
)

func customTrayIconAPI() *imports.Imports {
	customTrayIconOnce.Do(func() {
		customTrayIconImport = api.NewDefaultImports()
		customTrayIconImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomTrayIcon_Create", 0), // constructor NewCustomTrayIcon
			/* 1 */ imports.NewTable("TCustomTrayIcon_Hide", 0), // function Hide
			/* 2 */ imports.NewTable("TCustomTrayIcon_Show", 0), // function Show
			/* 3 */ imports.NewTable("TCustomTrayIcon_GetPosition", 0), // function GetPosition
			/* 4 */ imports.NewTable("TCustomTrayIcon_InternalUpdate", 0), // procedure InternalUpdate
			/* 5 */ imports.NewTable("TCustomTrayIcon_ShowBalloonHint", 0), // procedure ShowBalloonHint
			/* 6 */ imports.NewTable("TCustomTrayIcon_Animate", 0), // property Animate
			/* 7 */ imports.NewTable("TCustomTrayIcon_AnimateInterval", 0), // property AnimateInterval
			/* 8 */ imports.NewTable("TCustomTrayIcon_BalloonFlags", 0), // property BalloonFlags
			/* 9 */ imports.NewTable("TCustomTrayIcon_BalloonHint", 0), // property BalloonHint
			/* 10 */ imports.NewTable("TCustomTrayIcon_BalloonTimeout", 0), // property BalloonTimeout
			/* 11 */ imports.NewTable("TCustomTrayIcon_BalloonTitle", 0), // property BalloonTitle
			/* 12 */ imports.NewTable("TCustomTrayIcon_Canvas", 0), // property Canvas
			/* 13 */ imports.NewTable("TCustomTrayIcon_PopUpMenu", 0), // property PopUpMenu
			/* 14 */ imports.NewTable("TCustomTrayIcon_Icon", 0), // property Icon
			/* 15 */ imports.NewTable("TCustomTrayIcon_Icons", 0), // property Icons
			/* 16 */ imports.NewTable("TCustomTrayIcon_Hint", 0), // property Hint
			/* 17 */ imports.NewTable("TCustomTrayIcon_ShowIcon", 0), // property ShowIcon
			/* 18 */ imports.NewTable("TCustomTrayIcon_Visible", 0), // property Visible
			/* 19 */ imports.NewTable("TCustomTrayIcon_OnClick", 0), // event OnClick
			/* 20 */ imports.NewTable("TCustomTrayIcon_OnDblClick", 0), // event OnDblClick
			/* 21 */ imports.NewTable("TCustomTrayIcon_OnMouseDown", 0), // event OnMouseDown
			/* 22 */ imports.NewTable("TCustomTrayIcon_OnMouseUp", 0), // event OnMouseUp
			/* 23 */ imports.NewTable("TCustomTrayIcon_OnMouseMove", 0), // event OnMouseMove
			/* 24 */ imports.NewTable("TCustomTrayIcon_OnPaint", 0), // event OnPaint
			/* 25 */ imports.NewTable("TCustomTrayIcon_TClass", 0), // function TCustomTrayIconClass
		}
	})
	return customTrayIconImport
}
