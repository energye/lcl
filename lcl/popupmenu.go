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

// IPopupMenu Parent: IMenu
type IPopupMenu interface {
	IMenu
	PopupComponent() IComponent          // property
	SetPopupComponent(AValue IComponent) // property
	PopupPoint() (resultPoint TPoint)    // property
	Alignment() TPopupAlignment          // property
	SetAlignment(AValue TPopupAlignment) // property
	AutoPopup() bool                     // property
	SetAutoPopup(AValue bool)            // property
	HelpContext() THelpContext           // property
	SetHelpContext(AValue THelpContext)  // property
	TrackButton() TTrackButton           // property
	SetTrackButton(AValue TTrackButton)  // property
	PopUp()                              // procedure
	PopUp1(X, Y int32)                   // procedure
	Close()                              // procedure
	SetOnPopup(fn TNotifyEvent)          // property event
	SetOnClose(fn TNotifyEvent)          // property event
}

// TPopupMenu Parent: TMenu
type TPopupMenu struct {
	TMenu
	popupPtr uintptr
	closePtr uintptr
}

func NewPopupMenu(AOwner IComponent) IPopupMenu {
	r1 := popupMenuImportAPI().SysCallN(4, GetObjectUintptr(AOwner))
	return AsPopupMenu(r1)
}

func (m *TPopupMenu) PopupComponent() IComponent {
	r1 := popupMenuImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return AsComponent(r1)
}

func (m *TPopupMenu) SetPopupComponent(AValue IComponent) {
	popupMenuImportAPI().SysCallN(8, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPopupMenu) PopupPoint() (resultPoint TPoint) {
	popupMenuImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TPopupMenu) Alignment() TPopupAlignment {
	r1 := popupMenuImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TPopupAlignment(r1)
}

func (m *TPopupMenu) SetAlignment(AValue TPopupAlignment) {
	popupMenuImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TPopupMenu) AutoPopup() bool {
	r1 := popupMenuImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPopupMenu) SetAutoPopup(AValue bool) {
	popupMenuImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPopupMenu) HelpContext() THelpContext {
	r1 := popupMenuImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TPopupMenu) SetHelpContext(AValue THelpContext) {
	popupMenuImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TPopupMenu) TrackButton() TTrackButton {
	r1 := popupMenuImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TTrackButton(r1)
}

func (m *TPopupMenu) SetTrackButton(AValue TTrackButton) {
	popupMenuImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func PopupMenuClass() TClass {
	ret := popupMenuImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TPopupMenu) PopUp() {
	popupMenuImportAPI().SysCallN(6, m.Instance())
}

func (m *TPopupMenu) PopUp1(X, Y int32) {
	popupMenuImportAPI().SysCallN(7, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TPopupMenu) Close() {
	popupMenuImportAPI().SysCallN(3, m.Instance())
}

func (m *TPopupMenu) SetOnPopup(fn TNotifyEvent) {
	if m.popupPtr != 0 {
		RemoveEventElement(m.popupPtr)
	}
	m.popupPtr = MakeEventDataPtr(fn)
	popupMenuImportAPI().SysCallN(11, m.Instance(), m.popupPtr)
}

func (m *TPopupMenu) SetOnClose(fn TNotifyEvent) {
	if m.closePtr != 0 {
		RemoveEventElement(m.closePtr)
	}
	m.closePtr = MakeEventDataPtr(fn)
	popupMenuImportAPI().SysCallN(10, m.Instance(), m.closePtr)
}

var (
	popupMenuImport       *imports.Imports = nil
	popupMenuImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PopupMenu_Alignment", 0),
		/*1*/ imports.NewTable("PopupMenu_AutoPopup", 0),
		/*2*/ imports.NewTable("PopupMenu_Class", 0),
		/*3*/ imports.NewTable("PopupMenu_Close", 0),
		/*4*/ imports.NewTable("PopupMenu_Create", 0),
		/*5*/ imports.NewTable("PopupMenu_HelpContext", 0),
		/*6*/ imports.NewTable("PopupMenu_PopUp", 0),
		/*7*/ imports.NewTable("PopupMenu_PopUp1", 0),
		/*8*/ imports.NewTable("PopupMenu_PopupComponent", 0),
		/*9*/ imports.NewTable("PopupMenu_PopupPoint", 0),
		/*10*/ imports.NewTable("PopupMenu_SetOnClose", 0),
		/*11*/ imports.NewTable("PopupMenu_SetOnPopup", 0),
		/*12*/ imports.NewTable("PopupMenu_TrackButton", 0),
	}
)

func popupMenuImportAPI() *imports.Imports {
	if popupMenuImport == nil {
		popupMenuImport = NewDefaultImports()
		popupMenuImport.SetImportTable(popupMenuImportTables)
		popupMenuImportTables = nil
	}
	return popupMenuImport
}
