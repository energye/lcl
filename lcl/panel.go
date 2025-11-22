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

// IPanel Parent: ICustomPanel
type IPanel interface {
	ICustomPanel
	DragCursor() types.TCursor                           // property DragCursor Getter
	SetDragCursor(value types.TCursor)                   // property DragCursor Setter
	DragKind() types.TDragKind                           // property DragKind Getter
	SetDragKind(value types.TDragKind)                   // property DragKind Setter
	DragMode() types.TDragMode                           // property DragMode Getter
	SetDragMode(value types.TDragMode)                   // property DragMode Setter
	ParentFont() bool                                    // property ParentFont Getter
	SetParentFont(value bool)                            // property ParentFont Setter
	ParentShowHint() bool                                // property ParentShowHint Getter
	SetParentShowHint(value bool)                        // property ParentShowHint Setter
	ShowAccelChar() bool                                 // property ShowAccelChar Getter
	SetShowAccelChar(value bool)                         // property ShowAccelChar Setter
	VerticalAlignment() types.TVerticalAlignment         // property VerticalAlignment Getter
	SetVerticalAlignment(value types.TVerticalAlignment) // property VerticalAlignment Setter
	Wordwrap() bool                                      // property Wordwrap Getter
	SetWordwrap(value bool)                              // property Wordwrap Setter
	SetOnContextPopup(fn TContextPopupEvent)             // property event
	SetOnDblClick(fn TNotifyEvent)                       // property event
	SetOnDragDrop(fn TDragDropEvent)                     // property event
	SetOnDragOver(fn TDragOverEvent)                     // property event
	SetOnEndDock(fn TEndDragEvent)                       // property event
	SetOnEndDrag(fn TEndDragEvent)                       // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)               // property event
	SetOnGetDockCaption(fn TGetDockCaptionEvent)         // property event
	SetOnMouseDown(fn TMouseEvent)                       // property event
	SetOnMouseEnter(fn TNotifyEvent)                     // property event
	SetOnMouseLeave(fn TNotifyEvent)                     // property event
	SetOnMouseMove(fn TMouseMoveEvent)                   // property event
	SetOnMouseUp(fn TMouseEvent)                         // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                 // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)         // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)             // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)      // property event
	SetOnStartDock(fn TStartDockEvent)                   // property event
	SetOnStartDrag(fn TStartDragEvent)                   // property event
}

type TPanel struct {
	TCustomPanel
}

func (m *TPanel) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := panelAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TPanel) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	panelAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TPanel) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := panelAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TPanel) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	panelAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TPanel) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := panelAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TPanel) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	panelAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TPanel) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := panelAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPanel) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	panelAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TPanel) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := panelAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPanel) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	panelAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TPanel) ShowAccelChar() bool {
	if !m.IsValid() {
		return false
	}
	r := panelAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPanel) SetShowAccelChar(value bool) {
	if !m.IsValid() {
		return
	}
	panelAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TPanel) VerticalAlignment() types.TVerticalAlignment {
	if !m.IsValid() {
		return 0
	}
	r := panelAPI().SysCallN(7, 0, m.Instance())
	return types.TVerticalAlignment(r)
}

func (m *TPanel) SetVerticalAlignment(value types.TVerticalAlignment) {
	if !m.IsValid() {
		return
	}
	panelAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TPanel) Wordwrap() bool {
	if !m.IsValid() {
		return false
	}
	r := panelAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TPanel) SetWordwrap(value bool) {
	if !m.IsValid() {
		return
	}
	panelAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TPanel) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 9, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 11, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 12, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 13, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 14, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 15, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnGetDockCaption(fn TGetDockCaptionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetDockCaptionEvent(fn)
	base.SetEvent(m, 16, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 17, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 20, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 21, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 22, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 23, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 24, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 25, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 26, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 27, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 28, panelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPanel) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 29, panelAPI(), api.MakeEventDataPtr(cb))
}

// NewPanel class constructor
func NewPanel(theOwner IComponent) IPanel {
	r := panelAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsPanel(r)
}

func TPanelClass() types.TClass {
	r := panelAPI().SysCallN(30)
	return types.TClass(r)
}

var (
	panelOnce   base.Once
	panelImport *imports.Imports = nil
)

func panelAPI() *imports.Imports {
	panelOnce.Do(func() {
		panelImport = api.NewDefaultImports()
		panelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPanel_Create", 0), // constructor NewPanel
			/* 1 */ imports.NewTable("TPanel_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TPanel_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TPanel_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TPanel_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TPanel_ParentShowHint", 0), // property ParentShowHint
			/* 6 */ imports.NewTable("TPanel_ShowAccelChar", 0), // property ShowAccelChar
			/* 7 */ imports.NewTable("TPanel_VerticalAlignment", 0), // property VerticalAlignment
			/* 8 */ imports.NewTable("TPanel_Wordwrap", 0), // property Wordwrap
			/* 9 */ imports.NewTable("TPanel_OnContextPopup", 0), // event OnContextPopup
			/* 10 */ imports.NewTable("TPanel_OnDblClick", 0), // event OnDblClick
			/* 11 */ imports.NewTable("TPanel_OnDragDrop", 0), // event OnDragDrop
			/* 12 */ imports.NewTable("TPanel_OnDragOver", 0), // event OnDragOver
			/* 13 */ imports.NewTable("TPanel_OnEndDock", 0), // event OnEndDock
			/* 14 */ imports.NewTable("TPanel_OnEndDrag", 0), // event OnEndDrag
			/* 15 */ imports.NewTable("TPanel_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 16 */ imports.NewTable("TPanel_OnGetDockCaption", 0), // event OnGetDockCaption
			/* 17 */ imports.NewTable("TPanel_OnMouseDown", 0), // event OnMouseDown
			/* 18 */ imports.NewTable("TPanel_OnMouseEnter", 0), // event OnMouseEnter
			/* 19 */ imports.NewTable("TPanel_OnMouseLeave", 0), // event OnMouseLeave
			/* 20 */ imports.NewTable("TPanel_OnMouseMove", 0), // event OnMouseMove
			/* 21 */ imports.NewTable("TPanel_OnMouseUp", 0), // event OnMouseUp
			/* 22 */ imports.NewTable("TPanel_OnMouseWheel", 0), // event OnMouseWheel
			/* 23 */ imports.NewTable("TPanel_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 24 */ imports.NewTable("TPanel_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 25 */ imports.NewTable("TPanel_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 26 */ imports.NewTable("TPanel_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 27 */ imports.NewTable("TPanel_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 28 */ imports.NewTable("TPanel_OnStartDock", 0), // event OnStartDock
			/* 29 */ imports.NewTable("TPanel_OnStartDrag", 0), // event OnStartDrag
			/* 30 */ imports.NewTable("TPanel_TClass", 0), // function TPanelClass
		}
	})
	return panelImport
}
