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

// ILinkLabel Parent: ICustomLabel
type ILinkLabel interface {
	ICustomLabel
	Alignment() types.TAlignment             // property Alignment Getter
	SetAlignment(value types.TAlignment)     // property Alignment Setter
	DragCursor() types.TCursor               // property DragCursor Getter
	SetDragCursor(value types.TCursor)       // property DragCursor Setter
	DragKind() types.TDragKind               // property DragKind Getter
	SetDragKind(value types.TDragKind)       // property DragKind Setter
	DragMode() types.TDragMode               // property DragMode Getter
	SetDragMode(value types.TDragMode)       // property DragMode Setter
	ParentColor() bool                       // property ParentColor Getter
	SetParentColor(value bool)               // property ParentColor Setter
	ParentFont() bool                        // property ParentFont Getter
	SetParentFont(value bool)                // property ParentFont Setter
	ParentShowHint() bool                    // property ParentShowHint Getter
	SetParentShowHint(value bool)            // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent) // property event
	SetOnDblClick(fn TNotifyEvent)           // property event
	SetOnDragDrop(fn TDragDropEvent)         // property event
	SetOnDragOver(fn TDragOverEvent)         // property event
	SetOnEndDock(fn TEndDragEvent)           // property event
	SetOnEndDrag(fn TEndDragEvent)           // property event
	SetOnMouseDown(fn TMouseEvent)           // property event
	SetOnMouseEnter(fn TNotifyEvent)         // property event
	SetOnMouseLeave(fn TNotifyEvent)         // property event
	SetOnMouseMove(fn TMouseMoveEvent)       // property event
	SetOnMouseUp(fn TMouseEvent)             // property event
	SetOnStartDock(fn TStartDockEvent)       // property event
	SetOnStartDrag(fn TStartDragEvent)       // property event
	SetOnLinkClick(fn TSysLinkEvent)         // property event
}

type TLinkLabel struct {
	TCustomLabel
}

func (m *TLinkLabel) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := linkLabelAPI().SysCallN(1, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TLinkLabel) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	linkLabelAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TLinkLabel) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := linkLabelAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TLinkLabel) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	linkLabelAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TLinkLabel) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := linkLabelAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TLinkLabel) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	linkLabelAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TLinkLabel) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := linkLabelAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TLinkLabel) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	linkLabelAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TLinkLabel) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := linkLabelAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLinkLabel) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	linkLabelAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TLinkLabel) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := linkLabelAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLinkLabel) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	linkLabelAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TLinkLabel) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := linkLabelAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLinkLabel) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	linkLabelAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TLinkLabel) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 8, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 10, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 11, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 12, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 13, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 17, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 19, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 20, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLinkLabel) SetOnLinkClick(fn TSysLinkEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSysLinkEvent(fn)
	base.SetEvent(m, 21, linkLabelAPI(), api.MakeEventDataPtr(cb))
}

// NewLinkLabel class constructor
func NewLinkLabel(owner IComponent) ILinkLabel {
	r := linkLabelAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLinkLabel(r)
}

func TLinkLabelClass() types.TClass {
	r := linkLabelAPI().SysCallN(22)
	return types.TClass(r)
}

var (
	linkLabelOnce   base.Once
	linkLabelImport *imports.Imports = nil
)

func linkLabelAPI() *imports.Imports {
	linkLabelOnce.Do(func() {
		linkLabelImport = api.NewDefaultImports()
		linkLabelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLinkLabel_Create", 0), // constructor NewLinkLabel
			/* 1 */ imports.NewTable("TLinkLabel_Alignment", 0), // property Alignment
			/* 2 */ imports.NewTable("TLinkLabel_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TLinkLabel_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TLinkLabel_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TLinkLabel_ParentColor", 0), // property ParentColor
			/* 6 */ imports.NewTable("TLinkLabel_ParentFont", 0), // property ParentFont
			/* 7 */ imports.NewTable("TLinkLabel_ParentShowHint", 0), // property ParentShowHint
			/* 8 */ imports.NewTable("TLinkLabel_OnContextPopup", 0), // event OnContextPopup
			/* 9 */ imports.NewTable("TLinkLabel_OnDblClick", 0), // event OnDblClick
			/* 10 */ imports.NewTable("TLinkLabel_OnDragDrop", 0), // event OnDragDrop
			/* 11 */ imports.NewTable("TLinkLabel_OnDragOver", 0), // event OnDragOver
			/* 12 */ imports.NewTable("TLinkLabel_OnEndDock", 0), // event OnEndDock
			/* 13 */ imports.NewTable("TLinkLabel_OnEndDrag", 0), // event OnEndDrag
			/* 14 */ imports.NewTable("TLinkLabel_OnMouseDown", 0), // event OnMouseDown
			/* 15 */ imports.NewTable("TLinkLabel_OnMouseEnter", 0), // event OnMouseEnter
			/* 16 */ imports.NewTable("TLinkLabel_OnMouseLeave", 0), // event OnMouseLeave
			/* 17 */ imports.NewTable("TLinkLabel_OnMouseMove", 0), // event OnMouseMove
			/* 18 */ imports.NewTable("TLinkLabel_OnMouseUp", 0), // event OnMouseUp
			/* 19 */ imports.NewTable("TLinkLabel_OnStartDock", 0), // event OnStartDock
			/* 20 */ imports.NewTable("TLinkLabel_OnStartDrag", 0), // event OnStartDrag
			/* 21 */ imports.NewTable("TLinkLabel_OnLinkClick", 0), // event OnLinkClick
			/* 22 */ imports.NewTable("TLinkLabel_TClass", 0), // function TLinkLabelClass
		}
	})
	return linkLabelImport
}
