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

// IPageControl Parent: ICustomTabControl
type IPageControl interface {
	ICustomTabControl
	FindNextPage(curPage ITabSheet, goForward bool, checkTabVisible bool) ITabSheet // function
	IndexOfTabAtWithIntX2(X int32, Y int32) int32                                   // function
	IndexOfTabAtWithPoint(P types.TPoint) int32                                     // function
	IndexOfPageAtWithIntX2(X int32, Y int32) int32                                  // function
	IndexOfPageAtWithPoint(P types.TPoint) int32                                    // function
	AddTabSheet() ITabSheet                                                         // function
	Clear()                                                                         // procedure
	SelectNextPageWithBool(goForward bool)                                          // procedure
	SelectNextPageWithBoolX2(goForward bool, checkTabVisible bool)                  // procedure
	ActivePageIndex() int32                                                         // property ActivePageIndex Getter
	SetActivePageIndex(value int32)                                                 // property ActivePageIndex Setter
	PagesWithIntToTabSheet(index int32) ITabSheet                                   // property Pages Getter
	ActivePage() ITabSheet                                                          // property ActivePage Getter
	SetActivePage(value ITabSheet)                                                  // property ActivePage Setter
	DragCursor() types.TCursor                                                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)                                              // property DragCursor Setter
	DragKind() types.TDragKind                                                      // property DragKind Getter
	SetDragKind(value types.TDragKind)                                              // property DragKind Setter
	DragMode() types.TDragMode                                                      // property DragMode Getter
	SetDragMode(value types.TDragMode)                                              // property DragMode Setter
	ParentFont() bool                                                               // property ParentFont Getter
	SetParentFont(value bool)                                                       // property ParentFont Setter
	ParentShowHint() bool                                                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                                                   // property ParentShowHint Setter
	TabIndex() int32                                                                // property TabIndex Getter
	SetTabIndex(value int32)                                                        // property TabIndex Setter
	SetOnGetDockCaption(fn TGetDockCaptionEvent)                                    // property event
	SetOnChange(fn TNotifyEvent)                                                    // property event
	SetOnContextPopup(fn TContextPopupEvent)                                        // property event
	SetOnDragDrop(fn TDragDropEvent)                                                // property event
	SetOnDragOver(fn TDragOverEvent)                                                // property event
	SetOnEndDock(fn TEndDragEvent)                                                  // property event
	SetOnEndDrag(fn TEndDragEvent)                                                  // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)                                          // property event
	SetOnMouseDown(fn TMouseEvent)                                                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                                                // property event
	SetOnMouseLeave(fn TNotifyEvent)                                                // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                              // property event
	SetOnMouseUp(fn TMouseEvent)                                                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                                  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                                    // property event
	SetOnStartDock(fn TStartDockEvent)                                              // property event
	SetOnStartDrag(fn TStartDragEvent)                                              // property event
}

type TPageControl struct {
	TCustomTabControl
}

func (m *TPageControl) FindNextPage(curPage ITabSheet, goForward bool, checkTabVisible bool) ITabSheet {
	if !m.IsValid() {
		return nil
	}
	r := pageControlAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(curPage), api.PasBool(goForward), api.PasBool(checkTabVisible))
	return AsTabSheet(r)
}

func (m *TPageControl) IndexOfTabAtWithIntX2(X int32, Y int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageControlAPI().SysCallN(2, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r)
}

func (m *TPageControl) IndexOfTabAtWithPoint(P types.TPoint) int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageControlAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&P)))
	return int32(r)
}

func (m *TPageControl) IndexOfPageAtWithIntX2(X int32, Y int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageControlAPI().SysCallN(4, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r)
}

func (m *TPageControl) IndexOfPageAtWithPoint(P types.TPoint) int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageControlAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&P)))
	return int32(r)
}

func (m *TPageControl) AddTabSheet() ITabSheet {
	if !m.IsValid() {
		return nil
	}
	r := pageControlAPI().SysCallN(6, m.Instance())
	return AsTabSheet(r)
}

func (m *TPageControl) Clear() {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(7, m.Instance())
}

func (m *TPageControl) SelectNextPageWithBool(goForward bool) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(8, m.Instance(), api.PasBool(goForward))
}

func (m *TPageControl) SelectNextPageWithBoolX2(goForward bool, checkTabVisible bool) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(9, m.Instance(), api.PasBool(goForward), api.PasBool(checkTabVisible))
}

func (m *TPageControl) ActivePageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageControlAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TPageControl) SetActivePageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TPageControl) PagesWithIntToTabSheet(index int32) ITabSheet {
	if !m.IsValid() {
		return nil
	}
	r := pageControlAPI().SysCallN(11, m.Instance(), uintptr(index))
	return AsTabSheet(r)
}

func (m *TPageControl) ActivePage() ITabSheet {
	if !m.IsValid() {
		return nil
	}
	r := pageControlAPI().SysCallN(12, 0, m.Instance())
	return AsTabSheet(r)
}

func (m *TPageControl) SetActivePage(value ITabSheet) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPageControl) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := pageControlAPI().SysCallN(13, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TPageControl) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TPageControl) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := pageControlAPI().SysCallN(14, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TPageControl) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TPageControl) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := pageControlAPI().SysCallN(15, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TPageControl) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TPageControl) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := pageControlAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPageControl) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TPageControl) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := pageControlAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPageControl) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TPageControl) TabIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pageControlAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TPageControl) SetTabIndex(value int32) {
	if !m.IsValid() {
		return
	}
	pageControlAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TPageControl) SetOnGetDockCaption(fn TGetDockCaptionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetDockCaptionEvent(fn)
	base.SetEvent(m, 19, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 21, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 22, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 23, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 24, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 25, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 26, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 27, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 28, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 29, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 30, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 31, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 32, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 33, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 34, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 35, pageControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPageControl) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 36, pageControlAPI(), api.MakeEventDataPtr(cb))
}

// NewPageControl class constructor
func NewPageControl(theOwner IComponent) IPageControl {
	r := pageControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsPageControl(r)
}

func TPageControlClass() types.TClass {
	r := pageControlAPI().SysCallN(37)
	return types.TClass(r)
}

var (
	pageControlOnce   base.Once
	pageControlImport *imports.Imports = nil
)

func pageControlAPI() *imports.Imports {
	pageControlOnce.Do(func() {
		pageControlImport = api.NewDefaultImports()
		pageControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPageControl_Create", 0), // constructor NewPageControl
			/* 1 */ imports.NewTable("TPageControl_FindNextPage", 0), // function FindNextPage
			/* 2 */ imports.NewTable("TPageControl_IndexOfTabAtWithIntX2", 0), // function IndexOfTabAtWithIntX2
			/* 3 */ imports.NewTable("TPageControl_IndexOfTabAtWithPoint", 0), // function IndexOfTabAtWithPoint
			/* 4 */ imports.NewTable("TPageControl_IndexOfPageAtWithIntX2", 0), // function IndexOfPageAtWithIntX2
			/* 5 */ imports.NewTable("TPageControl_IndexOfPageAtWithPoint", 0), // function IndexOfPageAtWithPoint
			/* 6 */ imports.NewTable("TPageControl_AddTabSheet", 0), // function AddTabSheet
			/* 7 */ imports.NewTable("TPageControl_Clear", 0), // procedure Clear
			/* 8 */ imports.NewTable("TPageControl_SelectNextPageWithBool", 0), // procedure SelectNextPageWithBool
			/* 9 */ imports.NewTable("TPageControl_SelectNextPageWithBoolX2", 0), // procedure SelectNextPageWithBoolX2
			/* 10 */ imports.NewTable("TPageControl_ActivePageIndex", 0), // property ActivePageIndex
			/* 11 */ imports.NewTable("TPageControl_PagesWithIntToTabSheet", 0), // property PagesWithIntToTabSheet
			/* 12 */ imports.NewTable("TPageControl_ActivePage", 0), // property ActivePage
			/* 13 */ imports.NewTable("TPageControl_DragCursor", 0), // property DragCursor
			/* 14 */ imports.NewTable("TPageControl_DragKind", 0), // property DragKind
			/* 15 */ imports.NewTable("TPageControl_DragMode", 0), // property DragMode
			/* 16 */ imports.NewTable("TPageControl_ParentFont", 0), // property ParentFont
			/* 17 */ imports.NewTable("TPageControl_ParentShowHint", 0), // property ParentShowHint
			/* 18 */ imports.NewTable("TPageControl_TabIndex", 0), // property TabIndex
			/* 19 */ imports.NewTable("TPageControl_OnGetDockCaption", 0), // event OnGetDockCaption
			/* 20 */ imports.NewTable("TPageControl_OnChange", 0), // event OnChange
			/* 21 */ imports.NewTable("TPageControl_OnContextPopup", 0), // event OnContextPopup
			/* 22 */ imports.NewTable("TPageControl_OnDragDrop", 0), // event OnDragDrop
			/* 23 */ imports.NewTable("TPageControl_OnDragOver", 0), // event OnDragOver
			/* 24 */ imports.NewTable("TPageControl_OnEndDock", 0), // event OnEndDock
			/* 25 */ imports.NewTable("TPageControl_OnEndDrag", 0), // event OnEndDrag
			/* 26 */ imports.NewTable("TPageControl_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 27 */ imports.NewTable("TPageControl_OnMouseDown", 0), // event OnMouseDown
			/* 28 */ imports.NewTable("TPageControl_OnMouseEnter", 0), // event OnMouseEnter
			/* 29 */ imports.NewTable("TPageControl_OnMouseLeave", 0), // event OnMouseLeave
			/* 30 */ imports.NewTable("TPageControl_OnMouseMove", 0), // event OnMouseMove
			/* 31 */ imports.NewTable("TPageControl_OnMouseUp", 0), // event OnMouseUp
			/* 32 */ imports.NewTable("TPageControl_OnMouseWheel", 0), // event OnMouseWheel
			/* 33 */ imports.NewTable("TPageControl_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 34 */ imports.NewTable("TPageControl_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 35 */ imports.NewTable("TPageControl_OnStartDock", 0), // event OnStartDock
			/* 36 */ imports.NewTable("TPageControl_OnStartDrag", 0), // event OnStartDrag
			/* 37 */ imports.NewTable("TPageControl_TClass", 0), // function TPageControlClass
		}
	})
	return pageControlImport
}
