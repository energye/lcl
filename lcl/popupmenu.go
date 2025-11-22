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

// IPopupMenu Parent: IMenu
type IPopupMenu interface {
	IMenu
	PopUp()                                   // procedure
	PopUpWithIntX2(X int32, Y int32)          // procedure
	Close()                                   // procedure
	PopupComponent() IComponent               // property PopupComponent Getter
	SetPopupComponent(value IComponent)       // property PopupComponent Setter
	PopupPoint() types.TPoint                 // property PopupPoint Getter
	SetPopupPoint(value types.TPoint)         // property PopupPoint Setter
	Alignment() types.TPopupAlignment         // property Alignment Getter
	SetAlignment(value types.TPopupAlignment) // property Alignment Setter
	AutoPopup() bool                          // property AutoPopup Getter
	SetAutoPopup(value bool)                  // property AutoPopup Setter
	HelpContext() types.THelpContext          // property HelpContext Getter
	SetHelpContext(value types.THelpContext)  // property HelpContext Setter
	TrackButton() types.TTrackButton          // property TrackButton Getter
	SetTrackButton(value types.TTrackButton)  // property TrackButton Setter
	SetOnPopup(fn TNotifyEvent)               // property event
	SetOnClose(fn TNotifyEvent)               // property event
}

type TPopupMenu struct {
	TMenu
}

func (m *TPopupMenu) PopUp() {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(1, m.Instance())
}

func (m *TPopupMenu) PopUpWithIntX2(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(2, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TPopupMenu) Close() {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(3, m.Instance())
}

func (m *TPopupMenu) PopupComponent() IComponent {
	if !m.IsValid() {
		return nil
	}
	r := popupMenuAPI().SysCallN(4, 0, m.Instance())
	return AsComponent(r)
}

func (m *TPopupMenu) SetPopupComponent(value IComponent) {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPopupMenu) PopupPoint() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(5, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPopupMenu) SetPopupPoint(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(5, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TPopupMenu) Alignment() types.TPopupAlignment {
	if !m.IsValid() {
		return 0
	}
	r := popupMenuAPI().SysCallN(6, 0, m.Instance())
	return types.TPopupAlignment(r)
}

func (m *TPopupMenu) SetAlignment(value types.TPopupAlignment) {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TPopupMenu) AutoPopup() bool {
	if !m.IsValid() {
		return false
	}
	r := popupMenuAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPopupMenu) SetAutoPopup(value bool) {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TPopupMenu) HelpContext() types.THelpContext {
	if !m.IsValid() {
		return 0
	}
	r := popupMenuAPI().SysCallN(8, 0, m.Instance())
	return types.THelpContext(r)
}

func (m *TPopupMenu) SetHelpContext(value types.THelpContext) {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TPopupMenu) TrackButton() types.TTrackButton {
	if !m.IsValid() {
		return 0
	}
	r := popupMenuAPI().SysCallN(9, 0, m.Instance())
	return types.TTrackButton(r)
}

func (m *TPopupMenu) SetTrackButton(value types.TTrackButton) {
	if !m.IsValid() {
		return
	}
	popupMenuAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TPopupMenu) SetOnPopup(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, popupMenuAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPopupMenu) SetOnClose(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, popupMenuAPI(), api.MakeEventDataPtr(cb))
}

// NewPopupMenu class constructor
func NewPopupMenu(owner IComponent) IPopupMenu {
	r := popupMenuAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsPopupMenu(r)
}

func TPopupMenuClass() types.TClass {
	r := popupMenuAPI().SysCallN(12)
	return types.TClass(r)
}

var (
	popupMenuOnce   base.Once
	popupMenuImport *imports.Imports = nil
)

func popupMenuAPI() *imports.Imports {
	popupMenuOnce.Do(func() {
		popupMenuImport = api.NewDefaultImports()
		popupMenuImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPopupMenu_Create", 0), // constructor NewPopupMenu
			/* 1 */ imports.NewTable("TPopupMenu_PopUp", 0), // procedure PopUp
			/* 2 */ imports.NewTable("TPopupMenu_PopUpWithIntX2", 0), // procedure PopUpWithIntX2
			/* 3 */ imports.NewTable("TPopupMenu_Close", 0), // procedure Close
			/* 4 */ imports.NewTable("TPopupMenu_PopupComponent", 0), // property PopupComponent
			/* 5 */ imports.NewTable("TPopupMenu_PopupPoint", 0), // property PopupPoint
			/* 6 */ imports.NewTable("TPopupMenu_Alignment", 0), // property Alignment
			/* 7 */ imports.NewTable("TPopupMenu_AutoPopup", 0), // property AutoPopup
			/* 8 */ imports.NewTable("TPopupMenu_HelpContext", 0), // property HelpContext
			/* 9 */ imports.NewTable("TPopupMenu_TrackButton", 0), // property TrackButton
			/* 10 */ imports.NewTable("TPopupMenu_OnPopup", 0), // event OnPopup
			/* 11 */ imports.NewTable("TPopupMenu_OnClose", 0), // event OnClose
			/* 12 */ imports.NewTable("TPopupMenu_TClass", 0), // function TPopupMenuClass
		}
	})
	return popupMenuImport
}
