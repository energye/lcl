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

// IControlBar Parent: ICustomControlBar
type IControlBar interface {
	ICustomControlBar
	DragCursor() types.TCursor                         // property DragCursor Getter
	SetDragCursor(value types.TCursor)                 // property DragCursor Setter
	DragKind() types.TDragKind                         // property DragKind Getter
	SetDragKind(value types.TDragKind)                 // property DragKind Setter
	DragMode() types.TDragMode                         // property DragMode Getter
	SetDragMode(value types.TDragMode)                 // property DragMode Setter
	ParentFont() bool                                  // property ParentFont Getter
	SetParentFont(value bool)                          // property ParentFont Setter
	ParentShowHint() bool                              // property ParentShowHint Getter
	SetParentShowHint(value bool)                      // property ParentShowHint Setter
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
	SetOnContextPopup(fn TContextPopupEvent)           // property event
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
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

type TControlBar struct {
	TCustomControlBar
}

func (m *TControlBar) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := controlBarAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TControlBar) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	controlBarAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TControlBar) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := controlBarAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TControlBar) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	controlBarAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TControlBar) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := controlBarAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TControlBar) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	controlBarAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TControlBar) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := controlBarAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControlBar) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	controlBarAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TControlBar) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := controlBarAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TControlBar) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	controlBarAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TControlBar) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTConstrainedResizeEvent(fn)
	base.SetEvent(m, 6, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 11, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 12, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 13, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 17, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 19, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 22, controlBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TControlBar) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 23, controlBarAPI(), api.MakeEventDataPtr(cb))
}

// NewControlBar class constructor
func NewControlBar(owner IComponent) IControlBar {
	r := controlBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsControlBar(r)
}

func TControlBarClass() types.TClass {
	r := controlBarAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	controlBarOnce   base.Once
	controlBarImport *imports.Imports = nil
)

func controlBarAPI() *imports.Imports {
	controlBarOnce.Do(func() {
		controlBarImport = api.NewDefaultImports()
		controlBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TControlBar_Create", 0), // constructor NewControlBar
			/* 1 */ imports.NewTable("TControlBar_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TControlBar_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TControlBar_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TControlBar_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TControlBar_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TControlBar_OnConstrainedResize", 0), // event OnConstrainedResize
			/* 7 */ imports.NewTable("TControlBar_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TControlBar_OnDblClick", 0), // event OnDblClick
			/* 9 */ imports.NewTable("TControlBar_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TControlBar_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TControlBar_OnEndDock", 0), // event OnEndDock
			/* 12 */ imports.NewTable("TControlBar_OnEndDrag", 0), // event OnEndDrag
			/* 13 */ imports.NewTable("TControlBar_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 14 */ imports.NewTable("TControlBar_OnMouseDown", 0), // event OnMouseDown
			/* 15 */ imports.NewTable("TControlBar_OnMouseEnter", 0), // event OnMouseEnter
			/* 16 */ imports.NewTable("TControlBar_OnMouseLeave", 0), // event OnMouseLeave
			/* 17 */ imports.NewTable("TControlBar_OnMouseMove", 0), // event OnMouseMove
			/* 18 */ imports.NewTable("TControlBar_OnMouseUp", 0), // event OnMouseUp
			/* 19 */ imports.NewTable("TControlBar_OnMouseWheel", 0), // event OnMouseWheel
			/* 20 */ imports.NewTable("TControlBar_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 21 */ imports.NewTable("TControlBar_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 22 */ imports.NewTable("TControlBar_OnStartDock", 0), // event OnStartDock
			/* 23 */ imports.NewTable("TControlBar_OnStartDrag", 0), // event OnStartDrag
			/* 24 */ imports.NewTable("TControlBar_TClass", 0), // function TControlBarClass
		}
	})
	return controlBarImport
}
