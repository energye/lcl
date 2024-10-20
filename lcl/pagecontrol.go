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

// IPageControl Parent: ICustomTabControl
type IPageControl interface {
	ICustomTabControl
	ActivePageIndex() int32                                                    // property
	SetActivePageIndex(AValue int32)                                           // property
	PagesForTabSheet(Index int32) ITabSheet                                    // property
	ActivePage() ITabSheet                                                     // property
	SetActivePage(AValue ITabSheet)                                            // property
	DragCursor() TCursor                                                       // property
	SetDragCursor(AValue TCursor)                                              // property
	DragKind() TDragKind                                                       // property
	SetDragKind(AValue TDragKind)                                              // property
	DragMode() TDragMode                                                       // property
	SetDragMode(AValue TDragMode)                                              // property
	ParentFont() bool                                                          // property
	SetParentFont(AValue bool)                                                 // property
	ParentShowHint() bool                                                      // property
	SetParentShowHint(AValue bool)                                             // property
	TabIndex() int32                                                           // property
	SetTabIndex(AValue int32)                                                  // property
	FindNextPage(CurPage ITabSheet, GoForward, CheckTabVisible bool) ITabSheet // function
	IndexOfTabAt(X, Y int32) int32                                             // function
	IndexOfTabAt1(P *TPoint) int32                                             // function
	IndexOfPageAt(X, Y int32) int32                                            // function
	IndexOfPageAt1(P *TPoint) int32                                            // function
	AddTabSheet() ITabSheet                                                    // function
	Clear()                                                                    // procedure
	SelectNextPage(GoForward bool)                                             // procedure
	SelectNextPage1(GoForward bool, CheckTabVisible bool)                      // procedure
	SetOnGetDockCaption(fn TGetDockCaptionEvent)                               // property event
	SetOnChange(fn TNotifyEvent)                                               // property event
	SetOnContextPopup(fn TContextPopupEvent)                                   // property event
	SetOnDragDrop(fn TDragDropEvent)                                           // property event
	SetOnDragOver(fn TDragOverEvent)                                           // property event
	SetOnEndDock(fn TEndDragEvent)                                             // property event
	SetOnEndDrag(fn TEndDragEvent)                                             // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)                                     // property event
	SetOnMouseDown(fn TMouseEvent)                                             // property event
	SetOnMouseEnter(fn TNotifyEvent)                                           // property event
	SetOnMouseLeave(fn TNotifyEvent)                                           // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                         // property event
	SetOnMouseUp(fn TMouseEvent)                                               // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                       // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                             // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                               // property event
	SetOnStartDock(fn TStartDockEvent)                                         // property event
	SetOnStartDrag(fn TStartDragEvent)                                         // property event
}

// TPageControl Parent: TCustomTabControl
type TPageControl struct {
	TCustomTabControl
	getDockCaptionPtr uintptr
	changePtr         uintptr
	contextPopupPtr   uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDockPtr        uintptr
	endDragPtr        uintptr
	getSiteInfoPtr    uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewPageControl(TheOwner IComponent) IPageControl {
	r1 := pageControlImportAPI().SysCallN(5, GetObjectUintptr(TheOwner))
	return AsPageControl(r1)
}

func (m *TPageControl) ActivePageIndex() int32 {
	r1 := pageControlImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageControl) SetActivePageIndex(AValue int32) {
	pageControlImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) PagesForTabSheet(Index int32) ITabSheet {
	r1 := pageControlImportAPI().SysCallN(14, m.Instance(), uintptr(Index))
	return AsTabSheet(r1)
}

func (m *TPageControl) ActivePage() ITabSheet {
	r1 := pageControlImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsTabSheet(r1)
}

func (m *TPageControl) SetActivePage(AValue ITabSheet) {
	pageControlImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPageControl) DragCursor() TCursor {
	r1 := pageControlImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TPageControl) SetDragCursor(AValue TCursor) {
	pageControlImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) DragKind() TDragKind {
	r1 := pageControlImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TPageControl) SetDragKind(AValue TDragKind) {
	pageControlImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) DragMode() TDragMode {
	r1 := pageControlImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TPageControl) SetDragMode(AValue TDragMode) {
	pageControlImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) ParentFont() bool {
	r1 := pageControlImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPageControl) SetParentFont(AValue bool) {
	pageControlImportAPI().SysCallN(15, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPageControl) ParentShowHint() bool {
	r1 := pageControlImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPageControl) SetParentShowHint(AValue bool) {
	pageControlImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPageControl) TabIndex() int32 {
	r1 := pageControlImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TPageControl) SetTabIndex(AValue int32) {
	pageControlImportAPI().SysCallN(37, 1, m.Instance(), uintptr(AValue))
}

func (m *TPageControl) FindNextPage(CurPage ITabSheet, GoForward, CheckTabVisible bool) ITabSheet {
	r1 := pageControlImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(CurPage), PascalBool(GoForward), PascalBool(CheckTabVisible))
	return AsTabSheet(r1)
}

func (m *TPageControl) IndexOfTabAt(X, Y int32) int32 {
	r1 := pageControlImportAPI().SysCallN(12, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TPageControl) IndexOfTabAt1(P *TPoint) int32 {
	r1 := pageControlImportAPI().SysCallN(13, m.Instance(), uintptr(unsafePointer(P)))
	return int32(r1)
}

func (m *TPageControl) IndexOfPageAt(X, Y int32) int32 {
	r1 := pageControlImportAPI().SysCallN(10, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TPageControl) IndexOfPageAt1(P *TPoint) int32 {
	r1 := pageControlImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(P)))
	return int32(r1)
}

func (m *TPageControl) AddTabSheet() ITabSheet {
	r1 := pageControlImportAPI().SysCallN(2, m.Instance())
	return AsTabSheet(r1)
}

func PageControlClass() TClass {
	ret := pageControlImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TPageControl) Clear() {
	pageControlImportAPI().SysCallN(4, m.Instance())
}

func (m *TPageControl) SelectNextPage(GoForward bool) {
	pageControlImportAPI().SysCallN(17, m.Instance(), PascalBool(GoForward))
}

func (m *TPageControl) SelectNextPage1(GoForward bool, CheckTabVisible bool) {
	pageControlImportAPI().SysCallN(18, m.Instance(), PascalBool(GoForward), PascalBool(CheckTabVisible))
}

func (m *TPageControl) SetOnGetDockCaption(fn TGetDockCaptionEvent) {
	if m.getDockCaptionPtr != 0 {
		RemoveEventElement(m.getDockCaptionPtr)
	}
	m.getDockCaptionPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(25, m.Instance(), m.getDockCaptionPtr)
}

func (m *TPageControl) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(19, m.Instance(), m.changePtr)
}

func (m *TPageControl) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(20, m.Instance(), m.contextPopupPtr)
}

func (m *TPageControl) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(21, m.Instance(), m.dragDropPtr)
}

func (m *TPageControl) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(22, m.Instance(), m.dragOverPtr)
}

func (m *TPageControl) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(23, m.Instance(), m.endDockPtr)
}

func (m *TPageControl) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(24, m.Instance(), m.endDragPtr)
}

func (m *TPageControl) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(26, m.Instance(), m.getSiteInfoPtr)
}

func (m *TPageControl) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(27, m.Instance(), m.mouseDownPtr)
}

func (m *TPageControl) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(28, m.Instance(), m.mouseEnterPtr)
}

func (m *TPageControl) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(29, m.Instance(), m.mouseLeavePtr)
}

func (m *TPageControl) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(30, m.Instance(), m.mouseMovePtr)
}

func (m *TPageControl) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(31, m.Instance(), m.mouseUpPtr)
}

func (m *TPageControl) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(32, m.Instance(), m.mouseWheelPtr)
}

func (m *TPageControl) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(33, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TPageControl) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(34, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TPageControl) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(35, m.Instance(), m.startDockPtr)
}

func (m *TPageControl) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	pageControlImportAPI().SysCallN(36, m.Instance(), m.startDragPtr)
}

var (
	pageControlImport       *imports.Imports = nil
	pageControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PageControl_ActivePage", 0),
		/*1*/ imports.NewTable("PageControl_ActivePageIndex", 0),
		/*2*/ imports.NewTable("PageControl_AddTabSheet", 0),
		/*3*/ imports.NewTable("PageControl_Class", 0),
		/*4*/ imports.NewTable("PageControl_Clear", 0),
		/*5*/ imports.NewTable("PageControl_Create", 0),
		/*6*/ imports.NewTable("PageControl_DragCursor", 0),
		/*7*/ imports.NewTable("PageControl_DragKind", 0),
		/*8*/ imports.NewTable("PageControl_DragMode", 0),
		/*9*/ imports.NewTable("PageControl_FindNextPage", 0),
		/*10*/ imports.NewTable("PageControl_IndexOfPageAt", 0),
		/*11*/ imports.NewTable("PageControl_IndexOfPageAt1", 0),
		/*12*/ imports.NewTable("PageControl_IndexOfTabAt", 0),
		/*13*/ imports.NewTable("PageControl_IndexOfTabAt1", 0),
		/*14*/ imports.NewTable("PageControl_PagesForTabSheet", 0),
		/*15*/ imports.NewTable("PageControl_ParentFont", 0),
		/*16*/ imports.NewTable("PageControl_ParentShowHint", 0),
		/*17*/ imports.NewTable("PageControl_SelectNextPage", 0),
		/*18*/ imports.NewTable("PageControl_SelectNextPage1", 0),
		/*19*/ imports.NewTable("PageControl_SetOnChange", 0),
		/*20*/ imports.NewTable("PageControl_SetOnContextPopup", 0),
		/*21*/ imports.NewTable("PageControl_SetOnDragDrop", 0),
		/*22*/ imports.NewTable("PageControl_SetOnDragOver", 0),
		/*23*/ imports.NewTable("PageControl_SetOnEndDock", 0),
		/*24*/ imports.NewTable("PageControl_SetOnEndDrag", 0),
		/*25*/ imports.NewTable("PageControl_SetOnGetDockCaption", 0),
		/*26*/ imports.NewTable("PageControl_SetOnGetSiteInfo", 0),
		/*27*/ imports.NewTable("PageControl_SetOnMouseDown", 0),
		/*28*/ imports.NewTable("PageControl_SetOnMouseEnter", 0),
		/*29*/ imports.NewTable("PageControl_SetOnMouseLeave", 0),
		/*30*/ imports.NewTable("PageControl_SetOnMouseMove", 0),
		/*31*/ imports.NewTable("PageControl_SetOnMouseUp", 0),
		/*32*/ imports.NewTable("PageControl_SetOnMouseWheel", 0),
		/*33*/ imports.NewTable("PageControl_SetOnMouseWheelDown", 0),
		/*34*/ imports.NewTable("PageControl_SetOnMouseWheelUp", 0),
		/*35*/ imports.NewTable("PageControl_SetOnStartDock", 0),
		/*36*/ imports.NewTable("PageControl_SetOnStartDrag", 0),
		/*37*/ imports.NewTable("PageControl_TabIndex", 0),
	}
)

func pageControlImportAPI() *imports.Imports {
	if pageControlImport == nil {
		pageControlImport = NewDefaultImports()
		pageControlImport.SetImportTable(pageControlImportTables)
		pageControlImportTables = nil
	}
	return pageControlImport
}
