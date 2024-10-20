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

// IForm Parent: ICustomForm
type IForm interface {
	ICustomForm
	EnabledMaximize(v bool)
	EnabledMinimize(v bool)
	EnabledSystemMenu(v bool)
	ScreenCenter()
	WorkAreaCenter()
	InheritedWndProc(heMessage *TMessage)
	SetOnWndProc(fn TWndProcEvent)
	SetGoPtr(ptr uintptr)
	ScaleSelf()
	ScaleForPPI(newPPI int32)
	ScaleForCurrentDpi()
	ClientHandle() HWND                                // property
	DragKind() TDragKind                               // property
	SetDragKind(AValue TDragKind)                      // property
	DragMode() TDragMode                               // property
	SetDragMode(AValue TDragMode)                      // property
	SessionProperties() string                         // property
	SetSessionProperties(AValue string)                // property
	LCLVersion() string                                // property
	SetLCLVersion(AValue string)                       // property
	Cascade()                                          // procedure
	Next()                                             // procedure
	Previous()                                         // procedure
	Tile()                                             // procedure
	ArrangeIcons()                                     // procedure
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

// TForm Parent: TCustomForm
type TForm struct {
	TCustomForm
	constrainedResizePtr uintptr
	contextPopupPtr      uintptr
	dblClickPtr          uintptr
	dragDropPtr          uintptr
	dragOverPtr          uintptr
	endDockPtr           uintptr
	getSiteInfoPtr       uintptr
	mouseDownPtr         uintptr
	mouseEnterPtr        uintptr
	mouseLeavePtr        uintptr
	mouseMovePtr         uintptr
	mouseUpPtr           uintptr
	mouseWheelPtr        uintptr
	mouseWheelDownPtr    uintptr
	mouseWheelUpPtr      uintptr
	mouseWheelHorzPtr    uintptr
	mouseWheelLeftPtr    uintptr
	mouseWheelRightPtr   uintptr
	startDockPtr         uintptr
}

func NewForm(TheOwner IComponent) IForm {
	r1 := formImportAPI().SysCallN(4, GetObjectUintptr(TheOwner))
	return AsForm(r1)
}

func (m *TForm) ClientHandle() HWND {
	r1 := formImportAPI().SysCallN(3, m.Instance())
	return HWND(r1)
}

func (m *TForm) DragKind() TDragKind {
	r1 := formImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TForm) SetDragKind(AValue TDragKind) {
	formImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TForm) DragMode() TDragMode {
	r1 := formImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TForm) SetDragMode(AValue TDragMode) {
	formImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TForm) SessionProperties() string {
	r1 := formImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TForm) SetSessionProperties(AValue string) {
	formImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *TForm) LCLVersion() string {
	r1 := formImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TForm) SetLCLVersion(AValue string) {
	formImportAPI().SysCallN(7, 1, m.Instance(), PascalStr(AValue))
}

func FormClass() TClass {
	ret := formImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TForm) Cascade() {
	formImportAPI().SysCallN(1, m.Instance())
}

func (m *TForm) Next() {
	formImportAPI().SysCallN(8, m.Instance())
}

func (m *TForm) Previous() {
	formImportAPI().SysCallN(9, m.Instance())
}

func (m *TForm) Tile() {
	formImportAPI().SysCallN(30, m.Instance())
}

func (m *TForm) ArrangeIcons() {
	formImportAPI().SysCallN(0, m.Instance())
}

func (m *TForm) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(11, m.Instance(), m.constrainedResizePtr)
}

func (m *TForm) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(12, m.Instance(), m.contextPopupPtr)
}

func (m *TForm) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(13, m.Instance(), m.dblClickPtr)
}

func (m *TForm) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(14, m.Instance(), m.dragDropPtr)
}

func (m *TForm) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(15, m.Instance(), m.dragOverPtr)
}

func (m *TForm) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(16, m.Instance(), m.endDockPtr)
}

func (m *TForm) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(17, m.Instance(), m.getSiteInfoPtr)
}

func (m *TForm) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(18, m.Instance(), m.mouseDownPtr)
}

func (m *TForm) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(19, m.Instance(), m.mouseEnterPtr)
}

func (m *TForm) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(20, m.Instance(), m.mouseLeavePtr)
}

func (m *TForm) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(21, m.Instance(), m.mouseMovePtr)
}

func (m *TForm) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(22, m.Instance(), m.mouseUpPtr)
}

func (m *TForm) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(23, m.Instance(), m.mouseWheelPtr)
}

func (m *TForm) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(24, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TForm) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(28, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TForm) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(25, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TForm) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(26, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TForm) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(27, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TForm) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	formImportAPI().SysCallN(29, m.Instance(), m.startDockPtr)
}

var (
	formImport       *imports.Imports = nil
	formImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Form_ArrangeIcons", 0),
		/*1*/ imports.NewTable("Form_Cascade", 0),
		/*2*/ imports.NewTable("Form_Class", 0),
		/*3*/ imports.NewTable("Form_ClientHandle", 0),
		/*4*/ imports.NewTable("Form_Create", 0),
		/*5*/ imports.NewTable("Form_DragKind", 0),
		/*6*/ imports.NewTable("Form_DragMode", 0),
		/*7*/ imports.NewTable("Form_LCLVersion", 0),
		/*8*/ imports.NewTable("Form_Next", 0),
		/*9*/ imports.NewTable("Form_Previous", 0),
		/*10*/ imports.NewTable("Form_SessionProperties", 0),
		/*11*/ imports.NewTable("Form_SetOnConstrainedResize", 0),
		/*12*/ imports.NewTable("Form_SetOnContextPopup", 0),
		/*13*/ imports.NewTable("Form_SetOnDblClick", 0),
		/*14*/ imports.NewTable("Form_SetOnDragDrop", 0),
		/*15*/ imports.NewTable("Form_SetOnDragOver", 0),
		/*16*/ imports.NewTable("Form_SetOnEndDock", 0),
		/*17*/ imports.NewTable("Form_SetOnGetSiteInfo", 0),
		/*18*/ imports.NewTable("Form_SetOnMouseDown", 0),
		/*19*/ imports.NewTable("Form_SetOnMouseEnter", 0),
		/*20*/ imports.NewTable("Form_SetOnMouseLeave", 0),
		/*21*/ imports.NewTable("Form_SetOnMouseMove", 0),
		/*22*/ imports.NewTable("Form_SetOnMouseUp", 0),
		/*23*/ imports.NewTable("Form_SetOnMouseWheel", 0),
		/*24*/ imports.NewTable("Form_SetOnMouseWheelDown", 0),
		/*25*/ imports.NewTable("Form_SetOnMouseWheelHorz", 0),
		/*26*/ imports.NewTable("Form_SetOnMouseWheelLeft", 0),
		/*27*/ imports.NewTable("Form_SetOnMouseWheelRight", 0),
		/*28*/ imports.NewTable("Form_SetOnMouseWheelUp", 0),
		/*29*/ imports.NewTable("Form_SetOnStartDock", 0),
		/*30*/ imports.NewTable("Form_Tile", 0),
	}
)

func formImportAPI() *imports.Imports {
	if formImport == nil {
		formImport = NewDefaultImports()
		formImport.SetImportTable(formImportTables)
		formImportTables = nil
	}
	return formImport
}
