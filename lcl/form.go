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

// IForm Parent: ICustomForm
type IForm interface {
	ICustomForm
	// Cascade
	//  mdi related routine
	Cascade() // procedure
	// Next
	//  mdi related routine
	Next() // procedure
	// Previous
	//  mdi related routine
	Previous() // procedure
	// Tile
	//  mdi related routine
	Tile() // procedure
	// ArrangeIcons
	//  mdi related routine
	ArrangeIcons() // procedure
	// ClientHandle
	//  MDI implementation
	//  returns handle of MDIForm client handle (container for mdi children this
	//  is not Handle of form itself !)
	ClientHandle() types.HWND                          // property ClientHandle Getter
	DragKind() types.TDragKind                         // property DragKind Getter
	SetDragKind(value types.TDragKind)                 // property DragKind Setter
	DragMode() types.TDragMode                         // property DragMode Getter
	SetDragMode(value types.TDragMode)                 // property DragMode Setter
	SessionProperties() string                         // property SessionProperties Getter
	SetSessionProperties(value string)                 // property SessionProperties Setter
	LCLVersion() string                                // property LCLVersion Getter
	SetLCLVersion(value string)                        // property LCLVersion Setter
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
	SetOnContextPopup(fn TContextPopupEvent)           // property event
	SetOnDblClick(fn TNotifyEvent)                     // property event
	SetOnDragDrop(fn TDragDropEvent)                   // property event
	SetOnDragOver(fn TDragOverEvent)                   // property event
	SetOnEndDock(fn TEndDragEvent)                     // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)             // property event
	SetOnMouseDown(fn TMouseEvent)                     // property event
	SetOnMouseEnter(fn TNotifyEvent)                   // property event
	SetOnMouseLeave(fn TNotifyEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                       // property event
	SetOnMouseWheel(fn TMouseWheelEvent)               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)    // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
}

type TForm struct {
	TCustomForm
}

func (m *TForm) Cascade() {
	if !m.IsValid() {
		return
	}
	formAPI().SysCallN(1, m.Instance())
}

func (m *TForm) Next() {
	if !m.IsValid() {
		return
	}
	formAPI().SysCallN(2, m.Instance())
}

func (m *TForm) Previous() {
	if !m.IsValid() {
		return
	}
	formAPI().SysCallN(3, m.Instance())
}

func (m *TForm) Tile() {
	if !m.IsValid() {
		return
	}
	formAPI().SysCallN(4, m.Instance())
}

func (m *TForm) ArrangeIcons() {
	if !m.IsValid() {
		return
	}
	formAPI().SysCallN(5, m.Instance())
}

func (m *TForm) ClientHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	r := formAPI().SysCallN(6, m.Instance())
	return types.HWND(r)
}

func (m *TForm) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := formAPI().SysCallN(7, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TForm) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	formAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TForm) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := formAPI().SysCallN(8, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TForm) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	formAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TForm) SessionProperties() string {
	if !m.IsValid() {
		return ""
	}
	r := formAPI().SysCallN(9, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TForm) SetSessionProperties(value string) {
	if !m.IsValid() {
		return
	}
	formAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *TForm) LCLVersion() string {
	if !m.IsValid() {
		return ""
	}
	r := formAPI().SysCallN(10, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TForm) SetLCLVersion(value string) {
	if !m.IsValid() {
		return
	}
	formAPI().SysCallN(10, 1, m.Instance(), api.PasStr(value))
}

func (m *TForm) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTConstrainedResizeEvent(fn)
	base.SetEvent(m, 11, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 12, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 14, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 15, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 16, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 17, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 21, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 22, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 23, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 24, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 25, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 26, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 27, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 28, formAPI(), api.MakeEventDataPtr(cb))
}

func (m *TForm) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 29, formAPI(), api.MakeEventDataPtr(cb))
}

// NewForm class constructor
func NewForm(theOwner IComponent) IForm {
	r := formAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsForm(r)
}

func TFormClass() types.TClass {
	r := formAPI().SysCallN(30)
	return types.TClass(r)
}

var (
	formOnce   base.Once
	formImport *imports.Imports = nil
)

func formAPI() *imports.Imports {
	formOnce.Do(func() {
		formImport = api.NewDefaultImports()
		formImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TForm_Create", 0), // constructor NewForm
			/* 1 */ imports.NewTable("TForm_Cascade", 0), // procedure Cascade
			/* 2 */ imports.NewTable("TForm_Next", 0), // procedure Next
			/* 3 */ imports.NewTable("TForm_Previous", 0), // procedure Previous
			/* 4 */ imports.NewTable("TForm_Tile", 0), // procedure Tile
			/* 5 */ imports.NewTable("TForm_ArrangeIcons", 0), // procedure ArrangeIcons
			/* 6 */ imports.NewTable("TForm_ClientHandle", 0), // property ClientHandle
			/* 7 */ imports.NewTable("TForm_DragKind", 0), // property DragKind
			/* 8 */ imports.NewTable("TForm_DragMode", 0), // property DragMode
			/* 9 */ imports.NewTable("TForm_SessionProperties", 0), // property SessionProperties
			/* 10 */ imports.NewTable("TForm_LCLVersion", 0), // property LCLVersion
			/* 11 */ imports.NewTable("TForm_OnConstrainedResize", 0), // event OnConstrainedResize
			/* 12 */ imports.NewTable("TForm_OnContextPopup", 0), // event OnContextPopup
			/* 13 */ imports.NewTable("TForm_OnDblClick", 0), // event OnDblClick
			/* 14 */ imports.NewTable("TForm_OnDragDrop", 0), // event OnDragDrop
			/* 15 */ imports.NewTable("TForm_OnDragOver", 0), // event OnDragOver
			/* 16 */ imports.NewTable("TForm_OnEndDock", 0), // event OnEndDock
			/* 17 */ imports.NewTable("TForm_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 18 */ imports.NewTable("TForm_OnMouseDown", 0), // event OnMouseDown
			/* 19 */ imports.NewTable("TForm_OnMouseEnter", 0), // event OnMouseEnter
			/* 20 */ imports.NewTable("TForm_OnMouseLeave", 0), // event OnMouseLeave
			/* 21 */ imports.NewTable("TForm_OnMouseMove", 0), // event OnMouseMove
			/* 22 */ imports.NewTable("TForm_OnMouseUp", 0), // event OnMouseUp
			/* 23 */ imports.NewTable("TForm_OnMouseWheel", 0), // event OnMouseWheel
			/* 24 */ imports.NewTable("TForm_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 25 */ imports.NewTable("TForm_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 26 */ imports.NewTable("TForm_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 27 */ imports.NewTable("TForm_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 28 */ imports.NewTable("TForm_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 29 */ imports.NewTable("TForm_OnStartDock", 0), // event OnStartDock
			/* 30 */ imports.NewTable("TForm_TClass", 0), // function TFormClass
		}
	})
	return formImport
}
