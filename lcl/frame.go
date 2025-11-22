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

// IFrame Parent: ICustomFrame
type IFrame interface {
	ICustomFrame
	AutoScroll() bool                                  // property AutoScroll Getter
	SetAutoScroll(value bool)                          // property AutoScroll Setter
	DragCursor() types.TCursor                         // property DragCursor Getter
	SetDragCursor(value types.TCursor)                 // property DragCursor Setter
	DragKind() types.TDragKind                         // property DragKind Getter
	SetDragKind(value types.TDragKind)                 // property DragKind Setter
	DragMode() types.TDragMode                         // property DragMode Getter
	SetDragMode(value types.TDragMode)                 // property DragMode Setter
	LCLVersion() string                                // property LCLVersion Getter
	SetLCLVersion(value string)                        // property LCLVersion Setter
	ParentColor() bool                                 // property ParentColor Getter
	SetParentColor(value bool)                         // property ParentColor Setter
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
	SetOnMouseWheelHorz(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)    // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

type TFrame struct {
	TCustomFrame
}

func (m *TFrame) AutoScroll() bool {
	if !m.IsValid() {
		return false
	}
	r := frameAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFrame) SetAutoScroll(value bool) {
	if !m.IsValid() {
		return
	}
	frameAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TFrame) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := frameAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TFrame) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	frameAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TFrame) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := frameAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TFrame) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	frameAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TFrame) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := frameAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TFrame) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	frameAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TFrame) LCLVersion() string {
	if !m.IsValid() {
		return ""
	}
	r := frameAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFrame) SetLCLVersion(value string) {
	if !m.IsValid() {
		return
	}
	frameAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TFrame) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := frameAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFrame) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	frameAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TFrame) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := frameAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFrame) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	frameAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TFrame) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := frameAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFrame) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	frameAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TFrame) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTConstrainedResizeEvent(fn)
	base.SetEvent(m, 9, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 10, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 12, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 13, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 14, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 15, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetSiteInfoEvent(fn)
	base.SetEvent(m, 16, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 17, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 20, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 21, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 22, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 23, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 24, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 25, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 26, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 27, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 28, frameAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFrame) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 29, frameAPI(), api.MakeEventDataPtr(cb))
}

// NewFrame class constructor
func NewFrame(theOwner IComponent) IFrame {
	r := frameAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsFrame(r)
}

func TFrameClass() types.TClass {
	r := frameAPI().SysCallN(30)
	return types.TClass(r)
}

var (
	frameOnce   base.Once
	frameImport *imports.Imports = nil
)

func frameAPI() *imports.Imports {
	frameOnce.Do(func() {
		frameImport = api.NewDefaultImports()
		frameImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFrame_Create", 0), // constructor NewFrame
			/* 1 */ imports.NewTable("TFrame_AutoScroll", 0), // property AutoScroll
			/* 2 */ imports.NewTable("TFrame_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TFrame_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TFrame_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TFrame_LCLVersion", 0), // property LCLVersion
			/* 6 */ imports.NewTable("TFrame_ParentColor", 0), // property ParentColor
			/* 7 */ imports.NewTable("TFrame_ParentFont", 0), // property ParentFont
			/* 8 */ imports.NewTable("TFrame_ParentShowHint", 0), // property ParentShowHint
			/* 9 */ imports.NewTable("TFrame_OnConstrainedResize", 0), // event OnConstrainedResize
			/* 10 */ imports.NewTable("TFrame_OnContextPopup", 0), // event OnContextPopup
			/* 11 */ imports.NewTable("TFrame_OnDblClick", 0), // event OnDblClick
			/* 12 */ imports.NewTable("TFrame_OnDragDrop", 0), // event OnDragDrop
			/* 13 */ imports.NewTable("TFrame_OnDragOver", 0), // event OnDragOver
			/* 14 */ imports.NewTable("TFrame_OnEndDock", 0), // event OnEndDock
			/* 15 */ imports.NewTable("TFrame_OnEndDrag", 0), // event OnEndDrag
			/* 16 */ imports.NewTable("TFrame_OnGetSiteInfo", 0), // event OnGetSiteInfo
			/* 17 */ imports.NewTable("TFrame_OnMouseDown", 0), // event OnMouseDown
			/* 18 */ imports.NewTable("TFrame_OnMouseEnter", 0), // event OnMouseEnter
			/* 19 */ imports.NewTable("TFrame_OnMouseLeave", 0), // event OnMouseLeave
			/* 20 */ imports.NewTable("TFrame_OnMouseMove", 0), // event OnMouseMove
			/* 21 */ imports.NewTable("TFrame_OnMouseUp", 0), // event OnMouseUp
			/* 22 */ imports.NewTable("TFrame_OnMouseWheel", 0), // event OnMouseWheel
			/* 23 */ imports.NewTable("TFrame_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 24 */ imports.NewTable("TFrame_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 25 */ imports.NewTable("TFrame_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 26 */ imports.NewTable("TFrame_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 27 */ imports.NewTable("TFrame_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 28 */ imports.NewTable("TFrame_OnStartDock", 0), // event OnStartDock
			/* 29 */ imports.NewTable("TFrame_OnStartDrag", 0), // event OnStartDrag
			/* 30 */ imports.NewTable("TFrame_TClass", 0), // function TFrameClass
		}
	})
	return frameImport
}
