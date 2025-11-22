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

// IStaticText Parent: ICustomStaticText
type IStaticText interface {
	ICustomStaticText
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
	DragKind() types.TDragKind                      // property DragKind Getter
	SetDragKind(value types.TDragKind)              // property DragKind Setter
	DragMode() types.TDragMode                      // property DragMode Getter
	SetDragMode(value types.TDragMode)              // property DragMode Setter
	ParentFont() bool                               // property ParentFont Getter
	SetParentFont(value bool)                       // property ParentFont Setter
	ParentColor() bool                              // property ParentColor Getter
	SetParentColor(value bool)                      // property ParentColor Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

type TStaticText struct {
	TCustomStaticText
}

func (m *TStaticText) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := staticTextAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TStaticText) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	staticTextAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TStaticText) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := staticTextAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TStaticText) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	staticTextAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TStaticText) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := staticTextAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TStaticText) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	staticTextAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TStaticText) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := staticTextAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStaticText) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	staticTextAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TStaticText) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := staticTextAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStaticText) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	staticTextAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TStaticText) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := staticTextAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TStaticText) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	staticTextAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TStaticText) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 11, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 12, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 15, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 16, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 17, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 18, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 19, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 20, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 22, staticTextAPI(), api.MakeEventDataPtr(cb))
}

func (m *TStaticText) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 23, staticTextAPI(), api.MakeEventDataPtr(cb))
}

// NewStaticText class constructor
func NewStaticText(owner IComponent) IStaticText {
	r := staticTextAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsStaticText(r)
}

func TStaticTextClass() types.TClass {
	r := staticTextAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	staticTextOnce   base.Once
	staticTextImport *imports.Imports = nil
)

func staticTextAPI() *imports.Imports {
	staticTextOnce.Do(func() {
		staticTextImport = api.NewDefaultImports()
		staticTextImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStaticText_Create", 0), // constructor NewStaticText
			/* 1 */ imports.NewTable("TStaticText_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TStaticText_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TStaticText_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TStaticText_ParentFont", 0), // property ParentFont
			/* 5 */ imports.NewTable("TStaticText_ParentColor", 0), // property ParentColor
			/* 6 */ imports.NewTable("TStaticText_ParentShowHint", 0), // property ParentShowHint
			/* 7 */ imports.NewTable("TStaticText_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TStaticText_OnDblClick", 0), // event OnDblClick
			/* 9 */ imports.NewTable("TStaticText_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TStaticText_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TStaticText_OnEndDrag", 0), // event OnEndDrag
			/* 12 */ imports.NewTable("TStaticText_OnMouseDown", 0), // event OnMouseDown
			/* 13 */ imports.NewTable("TStaticText_OnMouseEnter", 0), // event OnMouseEnter
			/* 14 */ imports.NewTable("TStaticText_OnMouseLeave", 0), // event OnMouseLeave
			/* 15 */ imports.NewTable("TStaticText_OnMouseMove", 0), // event OnMouseMove
			/* 16 */ imports.NewTable("TStaticText_OnMouseUp", 0), // event OnMouseUp
			/* 17 */ imports.NewTable("TStaticText_OnMouseWheel", 0), // event OnMouseWheel
			/* 18 */ imports.NewTable("TStaticText_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 19 */ imports.NewTable("TStaticText_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 20 */ imports.NewTable("TStaticText_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 21 */ imports.NewTable("TStaticText_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 22 */ imports.NewTable("TStaticText_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 23 */ imports.NewTable("TStaticText_OnStartDrag", 0), // event OnStartDrag
			/* 24 */ imports.NewTable("TStaticText_TClass", 0), // function TStaticTextClass
		}
	})
	return staticTextImport
}
