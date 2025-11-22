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

// IFlowPanel Parent: ICustomFlowPanel
type IFlowPanel interface {
	ICustomFlowPanel
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
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

type TFlowPanel struct {
	TCustomFlowPanel
}

func (m *TFlowPanel) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := flowPanelAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TFlowPanel) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	flowPanelAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TFlowPanel) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := flowPanelAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TFlowPanel) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	flowPanelAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TFlowPanel) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := flowPanelAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TFlowPanel) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	flowPanelAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TFlowPanel) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := flowPanelAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFlowPanel) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	flowPanelAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TFlowPanel) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := flowPanelAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFlowPanel) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	flowPanelAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TFlowPanel) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTConstrainedResizeEvent(fn)
	base.SetEvent(m, 6, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 11, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 12, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 13, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 17, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 19, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFlowPanel) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 20, flowPanelAPI(), api.MakeEventDataPtr(cb))
}

// NewFlowPanel class constructor
func NewFlowPanel(owner IComponent) IFlowPanel {
	r := flowPanelAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsFlowPanel(r)
}

func TFlowPanelClass() types.TClass {
	r := flowPanelAPI().SysCallN(21)
	return types.TClass(r)
}

var (
	flowPanelOnce   base.Once
	flowPanelImport *imports.Imports = nil
)

func flowPanelAPI() *imports.Imports {
	flowPanelOnce.Do(func() {
		flowPanelImport = api.NewDefaultImports()
		flowPanelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFlowPanel_Create", 0), // constructor NewFlowPanel
			/* 1 */ imports.NewTable("TFlowPanel_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TFlowPanel_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TFlowPanel_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TFlowPanel_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TFlowPanel_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TFlowPanel_OnConstrainedResize", 0), // event OnConstrainedResize
			/* 7 */ imports.NewTable("TFlowPanel_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TFlowPanel_OnDblClick", 0), // event OnDblClick
			/* 9 */ imports.NewTable("TFlowPanel_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TFlowPanel_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TFlowPanel_OnEndDock", 0), // event OnEndDock
			/* 12 */ imports.NewTable("TFlowPanel_OnEndDrag", 0), // event OnEndDrag
			/* 13 */ imports.NewTable("TFlowPanel_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 14 */ imports.NewTable("TFlowPanel_OnMouseDown", 0), // event OnMouseDown
			/* 15 */ imports.NewTable("TFlowPanel_OnMouseEnter", 0), // event OnMouseEnter
			/* 16 */ imports.NewTable("TFlowPanel_OnMouseLeave", 0), // event OnMouseLeave
			/* 17 */ imports.NewTable("TFlowPanel_OnMouseMove", 0), // event OnMouseMove
			/* 18 */ imports.NewTable("TFlowPanel_OnMouseUp", 0), // event OnMouseUp
			/* 19 */ imports.NewTable("TFlowPanel_OnStartDock", 0), // event OnStartDock
			/* 20 */ imports.NewTable("TFlowPanel_OnStartDrag", 0), // event OnStartDrag
			/* 21 */ imports.NewTable("TFlowPanel_TClass", 0), // function TFlowPanelClass
		}
	})
	return flowPanelImport
}
