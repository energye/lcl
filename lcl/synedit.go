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

// ISynEdit Parent: ICustomSynEdit
type ISynEdit interface {
	ICustomSynEdit
	ParentColor() bool                             // property ParentColor Getter
	SetParentColor(value bool)                     // property ParentColor Setter
	ParentFont() bool                              // property ParentFont Getter
	SetParentFont(value bool)                      // property ParentFont Setter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnTripleClick(fn TNotifyEvent)              // property event
	SetOnQuadClick(fn TNotifyEvent)                // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnClickLink(fn TMouseEvent)                 // property event
	SetOnMouseLink(fn TSynMouseLinkEvent)          // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

type TSynEdit struct {
	TCustomSynEdit
}

func (m *TSynEdit) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEdit) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	synEditAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	synEditAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	synEditAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynEdit) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 4, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 5, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnTripleClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 6, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnQuadClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 8, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 9, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 10, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 11, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 12, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 13, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnClickLink(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 15, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnMouseLink(fn TSynMouseLinkEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynMouseLinkEvent(fn)
	base.SetEvent(m, 16, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 19, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 22, synEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynEdit) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 23, synEditAPI(), api.MakeEventDataPtr(cb))
}

// NewSynEdit class constructor
func NewSynEdit(owner IComponent) ISynEdit {
	r := synEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynEdit(r)
}

func TSynEditClass() types.TClass {
	r := synEditAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	synEditOnce   base.Once
	synEditImport *imports.Imports = nil
)

func synEditAPI() *imports.Imports {
	synEditOnce.Do(func() {
		synEditImport = api.NewDefaultImports()
		synEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEdit_Create", 0), // constructor NewSynEdit
			/* 1 */ imports.NewTable("TSynEdit_ParentColor", 0), // property ParentColor
			/* 2 */ imports.NewTable("TSynEdit_ParentFont", 0), // property ParentFont
			/* 3 */ imports.NewTable("TSynEdit_ParentShowHint", 0), // property ParentShowHint
			/* 4 */ imports.NewTable("TSynEdit_OnContextPopup", 0), // event OnContextPopup
			/* 5 */ imports.NewTable("TSynEdit_OnDblClick", 0), // event OnDblClick
			/* 6 */ imports.NewTable("TSynEdit_OnTripleClick", 0), // event OnTripleClick
			/* 7 */ imports.NewTable("TSynEdit_OnQuadClick", 0), // event OnQuadClick
			/* 8 */ imports.NewTable("TSynEdit_OnDragDrop", 0), // event OnDragDrop
			/* 9 */ imports.NewTable("TSynEdit_OnDragOver", 0), // event OnDragOver
			/* 10 */ imports.NewTable("TSynEdit_OnEndDock", 0), // event OnEndDock
			/* 11 */ imports.NewTable("TSynEdit_OnEndDrag", 0), // event OnEndDrag
			/* 12 */ imports.NewTable("TSynEdit_OnMouseDown", 0), // event OnMouseDown
			/* 13 */ imports.NewTable("TSynEdit_OnMouseMove", 0), // event OnMouseMove
			/* 14 */ imports.NewTable("TSynEdit_OnMouseUp", 0), // event OnMouseUp
			/* 15 */ imports.NewTable("TSynEdit_OnClickLink", 0), // event OnClickLink
			/* 16 */ imports.NewTable("TSynEdit_OnMouseLink", 0), // event OnMouseLink
			/* 17 */ imports.NewTable("TSynEdit_OnMouseEnter", 0), // event OnMouseEnter
			/* 18 */ imports.NewTable("TSynEdit_OnMouseLeave", 0), // event OnMouseLeave
			/* 19 */ imports.NewTable("TSynEdit_OnMouseWheel", 0), // event OnMouseWheel
			/* 20 */ imports.NewTable("TSynEdit_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 21 */ imports.NewTable("TSynEdit_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 22 */ imports.NewTable("TSynEdit_OnStartDock", 0), // event OnStartDock
			/* 23 */ imports.NewTable("TSynEdit_OnStartDrag", 0), // event OnStartDrag
			/* 24 */ imports.NewTable("TSynEdit_TClass", 0), // function TSynEditClass
		}
	})
	return synEditImport
}
