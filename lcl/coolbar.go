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

// ICoolBar Parent: ICustomCoolBar
type ICoolBar interface {
	ICustomCoolBar
	DragCursor() types.TCursor                     // property DragCursor Getter
	SetDragCursor(value types.TCursor)             // property DragCursor Setter
	DragKind() types.TDragKind                     // property DragKind Getter
	SetDragKind(value types.TDragKind)             // property DragKind Setter
	DragMode() types.TDragMode                     // property DragMode Getter
	SetDragMode(value types.TDragMode)             // property DragMode Setter
	ParentColor() bool                             // property ParentColor Getter
	SetParentColor(value bool)                     // property ParentColor Setter
	ParentFont() bool                              // property ParentFont Getter
	SetParentFont(value bool)                      // property ParentFont Setter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)         // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

type TCoolBar struct {
	TCustomCoolBar
}

func (m *TCoolBar) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := coolBarAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCoolBar) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	coolBarAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCoolBar) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := coolBarAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TCoolBar) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	coolBarAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCoolBar) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := coolBarAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TCoolBar) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	coolBarAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCoolBar) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBarAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBar) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	coolBarAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBar) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBarAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBar) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	coolBarAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBar) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBarAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBar) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	coolBarAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBar) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 11, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 12, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 13, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 17, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 19, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 22, coolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCoolBar) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 23, coolBarAPI(), api.MakeEventDataPtr(cb))
}

// NewCoolBar class constructor
func NewCoolBar(owner IComponent) ICoolBar {
	r := coolBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCoolBar(r)
}

func TCoolBarClass() types.TClass {
	r := coolBarAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	coolBarOnce   base.Once
	coolBarImport *imports.Imports = nil
)

func coolBarAPI() *imports.Imports {
	coolBarOnce.Do(func() {
		coolBarImport = api.NewDefaultImports()
		coolBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoolBar_Create", 0), // constructor NewCoolBar
			/* 1 */ imports.NewTable("TCoolBar_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TCoolBar_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TCoolBar_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TCoolBar_ParentColor", 0), // property ParentColor
			/* 5 */ imports.NewTable("TCoolBar_ParentFont", 0), // property ParentFont
			/* 6 */ imports.NewTable("TCoolBar_ParentShowHint", 0), // property ParentShowHint
			/* 7 */ imports.NewTable("TCoolBar_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TCoolBar_OnDblClick", 0), // event OnDblClick
			/* 9 */ imports.NewTable("TCoolBar_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TCoolBar_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TCoolBar_OnEndDock", 0), // event OnEndDock
			/* 12 */ imports.NewTable("TCoolBar_OnEndDrag", 0), // event OnEndDrag
			/* 13 */ imports.NewTable("TCoolBar_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 14 */ imports.NewTable("TCoolBar_OnMouseDown", 0), // event OnMouseDown
			/* 15 */ imports.NewTable("TCoolBar_OnMouseEnter", 0), // event OnMouseEnter
			/* 16 */ imports.NewTable("TCoolBar_OnMouseLeave", 0), // event OnMouseLeave
			/* 17 */ imports.NewTable("TCoolBar_OnMouseMove", 0), // event OnMouseMove
			/* 18 */ imports.NewTable("TCoolBar_OnMouseUp", 0), // event OnMouseUp
			/* 19 */ imports.NewTable("TCoolBar_OnMouseWheel", 0), // event OnMouseWheel
			/* 20 */ imports.NewTable("TCoolBar_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 21 */ imports.NewTable("TCoolBar_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 22 */ imports.NewTable("TCoolBar_OnStartDock", 0), // event OnStartDock
			/* 23 */ imports.NewTable("TCoolBar_OnStartDrag", 0), // event OnStartDrag
			/* 24 */ imports.NewTable("TCoolBar_TClass", 0), // function TCoolBarClass
		}
	})
	return coolBarImport
}
