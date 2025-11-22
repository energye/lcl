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

// IScrollBox Parent: IScrollingWinControl
type IScrollBox interface {
	IScrollingWinControl
	AutoScroll() bool                                  // property AutoScroll Getter
	SetAutoScroll(value bool)                          // property AutoScroll Setter
	DragCursor() types.TCursor                         // property DragCursor Getter
	SetDragCursor(value types.TCursor)                 // property DragCursor Setter
	DragKind() types.TDragKind                         // property DragKind Getter
	SetDragKind(value types.TDragKind)                 // property DragKind Setter
	DragMode() types.TDragMode                         // property DragMode Getter
	SetDragMode(value types.TDragMode)                 // property DragMode Setter
	ParentBackground() bool                            // property ParentBackground Getter
	SetParentBackground(value bool)                    // property ParentBackground Setter
	ParentColor() bool                                 // property ParentColor Getter
	SetParentColor(value bool)                         // property ParentColor Setter
	ParentFont() bool                                  // property ParentFont Getter
	SetParentFont(value bool)                          // property ParentFont Setter
	ParentShowHint() bool                              // property ParentShowHint Getter
	SetParentShowHint(value bool)                      // property ParentShowHint Setter
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
	SetOnDblClick(fn TNotifyEvent)                     // property event
	SetOnDragDrop(fn TDragDropEvent)                   // property event
	SetOnDragOver(fn TDragOverEvent)                   // property event
	SetOnEndDock(fn TEndDragEvent)                     // property event
	SetOnEndDrag(fn TEndDragEvent)                     // property event
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
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

type TScrollBox struct {
	TScrollingWinControl
}

func (m *TScrollBox) AutoScroll() bool {
	if !m.IsValid() {
		return false
	}
	r := scrollBoxAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TScrollBox) SetAutoScroll(value bool) {
	if !m.IsValid() {
		return
	}
	scrollBoxAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TScrollBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := scrollBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TScrollBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	scrollBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TScrollBox) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := scrollBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TScrollBox) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	scrollBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TScrollBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := scrollBoxAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TScrollBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	scrollBoxAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TScrollBox) ParentBackground() bool {
	if !m.IsValid() {
		return false
	}
	r := scrollBoxAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TScrollBox) SetParentBackground(value bool) {
	if !m.IsValid() {
		return
	}
	scrollBoxAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TScrollBox) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := scrollBoxAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TScrollBox) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	scrollBoxAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TScrollBox) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := scrollBoxAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TScrollBox) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	scrollBoxAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TScrollBox) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := scrollBoxAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TScrollBox) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	scrollBoxAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TScrollBox) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTConstrainedResizeEvent(fn)
	base.SetEvent(m, 9, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 11, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 12, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 13, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 14, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 15, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 16, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 19, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 20, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 21, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 22, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 23, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 24, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 25, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 26, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 27, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 28, scrollBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewScrollBox class constructor
func NewScrollBox(owner IComponent) IScrollBox {
	r := scrollBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsScrollBox(r)
}

func TScrollBoxClass() types.TClass {
	r := scrollBoxAPI().SysCallN(29)
	return types.TClass(r)
}

var (
	scrollBoxOnce   base.Once
	scrollBoxImport *imports.Imports = nil
)

func scrollBoxAPI() *imports.Imports {
	scrollBoxOnce.Do(func() {
		scrollBoxImport = api.NewDefaultImports()
		scrollBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TScrollBox_Create", 0), // constructor NewScrollBox
			/* 1 */ imports.NewTable("TScrollBox_AutoScroll", 0), // property AutoScroll
			/* 2 */ imports.NewTable("TScrollBox_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TScrollBox_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TScrollBox_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TScrollBox_ParentBackground", 0), // property ParentBackground
			/* 6 */ imports.NewTable("TScrollBox_ParentColor", 0), // property ParentColor
			/* 7 */ imports.NewTable("TScrollBox_ParentFont", 0), // property ParentFont
			/* 8 */ imports.NewTable("TScrollBox_ParentShowHint", 0), // property ParentShowHint
			/* 9 */ imports.NewTable("TScrollBox_OnConstrainedResize", 0), // event OnConstrainedResize
			/* 10 */ imports.NewTable("TScrollBox_OnDblClick", 0), // event OnDblClick
			/* 11 */ imports.NewTable("TScrollBox_OnDragDrop", 0), // event OnDragDrop
			/* 12 */ imports.NewTable("TScrollBox_OnDragOver", 0), // event OnDragOver
			/* 13 */ imports.NewTable("TScrollBox_OnEndDock", 0), // event OnEndDock
			/* 14 */ imports.NewTable("TScrollBox_OnEndDrag", 0), // event OnEndDrag
			/* 15 */ imports.NewTable("TScrollBox_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 16 */ imports.NewTable("TScrollBox_OnMouseDown", 0), // event OnMouseDown
			/* 17 */ imports.NewTable("TScrollBox_OnMouseEnter", 0), // event OnMouseEnter
			/* 18 */ imports.NewTable("TScrollBox_OnMouseLeave", 0), // event OnMouseLeave
			/* 19 */ imports.NewTable("TScrollBox_OnMouseMove", 0), // event OnMouseMove
			/* 20 */ imports.NewTable("TScrollBox_OnMouseUp", 0), // event OnMouseUp
			/* 21 */ imports.NewTable("TScrollBox_OnMouseWheel", 0), // event OnMouseWheel
			/* 22 */ imports.NewTable("TScrollBox_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 23 */ imports.NewTable("TScrollBox_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 24 */ imports.NewTable("TScrollBox_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 25 */ imports.NewTable("TScrollBox_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 26 */ imports.NewTable("TScrollBox_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 27 */ imports.NewTable("TScrollBox_OnStartDock", 0), // event OnStartDock
			/* 28 */ imports.NewTable("TScrollBox_OnStartDrag", 0), // event OnStartDrag
			/* 29 */ imports.NewTable("TScrollBox_TClass", 0), // function TScrollBoxClass
		}
	})
	return scrollBoxImport
}
