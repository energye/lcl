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

// IStatusBar Parent: IWinControl
type IStatusBar interface {
	IWinControl
	GetPanelIndexAt(X int32, Y int32) int32                         // function
	SizeGripEnabled() bool                                          // function
	UpdatingStatusBar() bool                                        // function
	InvalidatePanel(panelIndex int32, panelParts types.TPanelParts) // procedure
	BeginUpdate()                                                   // procedure
	EndUpdate()                                                     // procedure
	Canvas() ICanvas                                                // property Canvas Getter
	AutoHint() bool                                                 // property AutoHint Getter
	SetAutoHint(value bool)                                         // property AutoHint Setter
	DragCursor() types.TCursor                                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)                              // property DragCursor Setter
	DragKind() types.TDragKind                                      // property DragKind Getter
	SetDragKind(value types.TDragKind)                              // property DragKind Setter
	DragMode() types.TDragMode                                      // property DragMode Getter
	SetDragMode(value types.TDragMode)                              // property DragMode Setter
	Panels() IStatusPanels                                          // property Panels Getter
	SetPanels(value IStatusPanels)                                  // property Panels Setter
	ParentColor() bool                                              // property ParentColor Getter
	SetParentColor(value bool)                                      // property ParentColor Setter
	ParentFont() bool                                               // property ParentFont Getter
	SetParentFont(value bool)                                       // property ParentFont Setter
	ParentShowHint() bool                                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                                   // property ParentShowHint Setter
	SimpleText() string                                             // property SimpleText Getter
	SetSimpleText(value string)                                     // property SimpleText Setter
	SimplePanel() bool                                              // property SimplePanel Getter
	SetSimplePanel(value bool)                                      // property SimplePanel Setter
	SizeGrip() bool                                                 // property SizeGrip Getter
	SetSizeGrip(value bool)                                         // property SizeGrip Setter
	UseSystemFont() bool                                            // property UseSystemFont Getter
	SetUseSystemFont(value bool)                                    // property UseSystemFont Setter
	SetOnContextPopup(fn TContextPopupEvent)                        // property event
	SetOnCreatePanelClass(fn TSBCreatePanelClassEvent)              // property event
	SetOnDblClick(fn TNotifyEvent)                                  // property event
	SetOnDragDrop(fn TDragDropEvent)                                // property event
	SetOnDragOver(fn TDragOverEvent)                                // property event
	SetOnDrawPanel(fn TDrawPanelEvent)                              // property event
	SetOnEndDock(fn TEndDragEvent)                                  // property event
	SetOnEndDrag(fn TEndDragEvent)                                  // property event
	SetOnHint(fn TNotifyEvent)                                      // property event
	SetOnMouseDown(fn TMouseEvent)                                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                                // property event
	SetOnMouseLeave(fn TNotifyEvent)                                // property event
	SetOnMouseMove(fn TMouseMoveEvent)                              // property event
	SetOnMouseUp(fn TMouseEvent)                                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                    // property event
	SetOnStartDock(fn TStartDockEvent)                              // property event
	SetOnStartDrag(fn TStartDragEvent)                              // property event
}

type TStatusBar struct {
	TWinControl
}

func (m *TStatusBar) GetPanelIndexAt(X int32, Y int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := statusBarAPI().SysCallN(1, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r)
}

func (m *TStatusBar) SizeGripEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := statusBarAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusBar) UpdatingStatusBar() bool {
	if !m.IsValid() {
		return false
	}
	r := statusBarAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusBar) InvalidatePanel(panelIndex int32, panelParts types.TPanelParts) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(4, m.Instance(), uintptr(panelIndex), uintptr(panelParts))
}

func (m *TStatusBar) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(5, m.Instance())
}

func (m *TStatusBar) EndUpdate() {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(6, m.Instance())
}

func (m *TStatusBar) Canvas() ICanvas {
	if !m.IsValid() {
		return nil
	}
	r := statusBarAPI().SysCallN(7, m.Instance())
	return AsCanvas(r)
}

func (m *TStatusBar) AutoHint() bool {
	if !m.IsValid() {
		return false
	}
	r := statusBarAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusBar) SetAutoHint(value bool) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TStatusBar) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := statusBarAPI().SysCallN(9, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TStatusBar) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TStatusBar) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := statusBarAPI().SysCallN(10, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TStatusBar) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TStatusBar) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := statusBarAPI().SysCallN(11, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TStatusBar) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TStatusBar) Panels() IStatusPanels {
	if !m.IsValid() {
		return nil
	}
	r := statusBarAPI().SysCallN(12, 0, m.Instance())
	return AsStatusPanels(r)
}

func (m *TStatusBar) SetPanels(value IStatusPanels) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TStatusBar) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := statusBarAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusBar) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TStatusBar) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := statusBarAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusBar) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TStatusBar) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := statusBarAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusBar) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TStatusBar) SimpleText() string {
	if !m.IsValid() {
		return ""
	}
	r := statusBarAPI().SysCallN(16, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TStatusBar) SetSimpleText(value string) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(16, 1, m.Instance(), api.PasStr(value))
}

func (m *TStatusBar) SimplePanel() bool {
	if !m.IsValid() {
		return false
	}
	r := statusBarAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusBar) SetSimplePanel(value bool) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TStatusBar) SizeGrip() bool {
	if !m.IsValid() {
		return false
	}
	r := statusBarAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusBar) SetSizeGrip(value bool) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TStatusBar) UseSystemFont() bool {
	if !m.IsValid() {
		return false
	}
	r := statusBarAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStatusBar) SetUseSystemFont(value bool) {
	if !m.IsValid() {
		return
	}
	statusBarAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TStatusBar) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 20, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnCreatePanelClass(fn TSBCreatePanelClassEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSBCreatePanelClassEvent(fn)
	base.SetEvent(m, 21, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 22, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 23, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 24, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnDrawPanel(fn TDrawPanelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDrawPanelEvent(fn)
	base.SetEvent(m, 25, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 26, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 27, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnHint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 28, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 29, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 30, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 31, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 32, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 33, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 34, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 35, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 36, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 37, statusBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStatusBar) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 38, statusBarAPI(), api.MakeEventDataPtr(cb))
}

// NewStatusBar class constructor
func NewStatusBar(theOwner IComponent) IStatusBar {
	r := statusBarAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsStatusBar(r)
}

func TStatusBarClass() types.TClass {
	r := statusBarAPI().SysCallN(39)
	return types.TClass(r)
}

var (
	statusBarOnce   base.Once
	statusBarImport *imports.Imports = nil
)

func statusBarAPI() *imports.Imports {
	statusBarOnce.Do(func() {
		statusBarImport = api.NewDefaultImports()
		statusBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStatusBar_Create", 0), // constructor NewStatusBar
			/* 1 */ imports.NewTable("TStatusBar_GetPanelIndexAt", 0), // function GetPanelIndexAt
			/* 2 */ imports.NewTable("TStatusBar_SizeGripEnabled", 0), // function SizeGripEnabled
			/* 3 */ imports.NewTable("TStatusBar_UpdatingStatusBar", 0), // function UpdatingStatusBar
			/* 4 */ imports.NewTable("TStatusBar_InvalidatePanel", 0), // procedure InvalidatePanel
			/* 5 */ imports.NewTable("TStatusBar_BeginUpdate", 0), // procedure BeginUpdate
			/* 6 */ imports.NewTable("TStatusBar_EndUpdate", 0), // procedure EndUpdate
			/* 7 */ imports.NewTable("TStatusBar_Canvas", 0), // property Canvas
			/* 8 */ imports.NewTable("TStatusBar_AutoHint", 0), // property AutoHint
			/* 9 */ imports.NewTable("TStatusBar_DragCursor", 0), // property DragCursor
			/* 10 */ imports.NewTable("TStatusBar_DragKind", 0), // property DragKind
			/* 11 */ imports.NewTable("TStatusBar_DragMode", 0), // property DragMode
			/* 12 */ imports.NewTable("TStatusBar_Panels", 0), // property Panels
			/* 13 */ imports.NewTable("TStatusBar_ParentColor", 0), // property ParentColor
			/* 14 */ imports.NewTable("TStatusBar_ParentFont", 0), // property ParentFont
			/* 15 */ imports.NewTable("TStatusBar_ParentShowHint", 0), // property ParentShowHint
			/* 16 */ imports.NewTable("TStatusBar_SimpleText", 0), // property SimpleText
			/* 17 */ imports.NewTable("TStatusBar_SimplePanel", 0), // property SimplePanel
			/* 18 */ imports.NewTable("TStatusBar_SizeGrip", 0), // property SizeGrip
			/* 19 */ imports.NewTable("TStatusBar_UseSystemFont", 0), // property UseSystemFont
			/* 20 */ imports.NewTable("TStatusBar_OnContextPopup", 0), // event OnContextPopup
			/* 21 */ imports.NewTable("TStatusBar_OnCreatePanelClass", 0), // event OnCreatePanelClass
			/* 22 */ imports.NewTable("TStatusBar_OnDblClick", 0), // event OnDblClick
			/* 23 */ imports.NewTable("TStatusBar_OnDragDrop", 0), // event OnDragDrop
			/* 24 */ imports.NewTable("TStatusBar_OnDragOver", 0), // event OnDragOver
			/* 25 */ imports.NewTable("TStatusBar_OnDrawPanel", 0), // event OnDrawPanel
			/* 26 */ imports.NewTable("TStatusBar_OnEndDock", 0), // event OnEndDock
			/* 27 */ imports.NewTable("TStatusBar_OnEndDrag", 0), // event OnEndDrag
			/* 28 */ imports.NewTable("TStatusBar_OnHint", 0), // event OnHint
			/* 29 */ imports.NewTable("TStatusBar_OnMouseDown", 0), // event OnMouseDown
			/* 30 */ imports.NewTable("TStatusBar_OnMouseEnter", 0), // event OnMouseEnter
			/* 31 */ imports.NewTable("TStatusBar_OnMouseLeave", 0), // event OnMouseLeave
			/* 32 */ imports.NewTable("TStatusBar_OnMouseMove", 0), // event OnMouseMove
			/* 33 */ imports.NewTable("TStatusBar_OnMouseUp", 0), // event OnMouseUp
			/* 34 */ imports.NewTable("TStatusBar_OnMouseWheel", 0), // event OnMouseWheel
			/* 35 */ imports.NewTable("TStatusBar_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 36 */ imports.NewTable("TStatusBar_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 37 */ imports.NewTable("TStatusBar_OnStartDock", 0), // event OnStartDock
			/* 38 */ imports.NewTable("TStatusBar_OnStartDrag", 0), // event OnStartDrag
			/* 39 */ imports.NewTable("TStatusBar_TClass", 0), // function TStatusBarClass
		}
	})
	return statusBarImport
}
