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

// IGroupBox Parent: ICustomGroupBox
type IGroupBox interface {
	ICustomGroupBox
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

type TGroupBox struct {
	TCustomGroupBox
}

func (m *TGroupBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := groupBoxAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TGroupBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	groupBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TGroupBox) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := groupBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TGroupBox) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	groupBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TGroupBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := groupBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TGroupBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	groupBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TGroupBox) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := groupBoxAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGroupBox) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	groupBoxAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TGroupBox) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := groupBoxAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGroupBox) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	groupBoxAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TGroupBox) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := groupBoxAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGroupBox) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	groupBoxAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TGroupBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 11, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 12, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 13, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 17, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 19, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 22, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGroupBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 23, groupBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewGroupBox class constructor
func NewGroupBox(owner IComponent) IGroupBox {
	r := groupBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsGroupBox(r)
}

func TGroupBoxClass() types.TClass {
	r := groupBoxAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	groupBoxOnce   base.Once
	groupBoxImport *imports.Imports = nil
)

func groupBoxAPI() *imports.Imports {
	groupBoxOnce.Do(func() {
		groupBoxImport = api.NewDefaultImports()
		groupBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGroupBox_Create", 0), // constructor NewGroupBox
			/* 1 */ imports.NewTable("TGroupBox_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TGroupBox_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TGroupBox_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TGroupBox_ParentColor", 0), // property ParentColor
			/* 5 */ imports.NewTable("TGroupBox_ParentFont", 0), // property ParentFont
			/* 6 */ imports.NewTable("TGroupBox_ParentShowHint", 0), // property ParentShowHint
			/* 7 */ imports.NewTable("TGroupBox_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TGroupBox_OnDblClick", 0), // event OnDblClick
			/* 9 */ imports.NewTable("TGroupBox_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TGroupBox_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TGroupBox_OnEndDock", 0), // event OnEndDock
			/* 12 */ imports.NewTable("TGroupBox_OnEndDrag", 0), // event OnEndDrag
			/* 13 */ imports.NewTable("TGroupBox_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 14 */ imports.NewTable("TGroupBox_OnMouseDown", 0), // event OnMouseDown
			/* 15 */ imports.NewTable("TGroupBox_OnMouseEnter", 0), // event OnMouseEnter
			/* 16 */ imports.NewTable("TGroupBox_OnMouseLeave", 0), // event OnMouseLeave
			/* 17 */ imports.NewTable("TGroupBox_OnMouseMove", 0), // event OnMouseMove
			/* 18 */ imports.NewTable("TGroupBox_OnMouseUp", 0), // event OnMouseUp
			/* 19 */ imports.NewTable("TGroupBox_OnMouseWheel", 0), // event OnMouseWheel
			/* 20 */ imports.NewTable("TGroupBox_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 21 */ imports.NewTable("TGroupBox_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 22 */ imports.NewTable("TGroupBox_OnStartDock", 0), // event OnStartDock
			/* 23 */ imports.NewTable("TGroupBox_OnStartDrag", 0), // event OnStartDrag
			/* 24 */ imports.NewTable("TGroupBox_TClass", 0), // function TGroupBoxClass
		}
	})
	return groupBoxImport
}
