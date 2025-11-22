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

// ILabeledEdit Parent: ICustomLabeledEdit
type ILabeledEdit interface {
	ICustomLabeledEdit
	AutoSelect() bool                              // property AutoSelect Getter
	SetAutoSelect(value bool)                      // property AutoSelect Setter
	DragCursor() types.TCursor                     // property DragCursor Getter
	SetDragCursor(value types.TCursor)             // property DragCursor Setter
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
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

type TLabeledEdit struct {
	TCustomLabeledEdit
}

func (m *TLabeledEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := labeledEditAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabeledEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	labeledEditAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabeledEdit) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := labeledEditAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TLabeledEdit) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	labeledEditAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TLabeledEdit) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := labeledEditAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TLabeledEdit) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	labeledEditAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TLabeledEdit) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := labeledEditAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabeledEdit) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	labeledEditAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabeledEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := labeledEditAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabeledEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	labeledEditAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabeledEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := labeledEditAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabeledEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	labeledEditAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabeledEdit) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 7, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 9, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 10, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 12, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 13, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 16, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 17, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 18, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 19, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabeledEdit) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 21, labeledEditAPI(), api.MakeEventDataPtr(cb))
}

// NewLabeledEdit class constructor
func NewLabeledEdit(theOwner IComponent) ILabeledEdit {
	r := labeledEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsLabeledEdit(r)
}

func TLabeledEditClass() types.TClass {
	r := labeledEditAPI().SysCallN(22)
	return types.TClass(r)
}

var (
	labeledEditOnce   base.Once
	labeledEditImport *imports.Imports = nil
)

func labeledEditAPI() *imports.Imports {
	labeledEditOnce.Do(func() {
		labeledEditImport = api.NewDefaultImports()
		labeledEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLabeledEdit_Create", 0), // constructor NewLabeledEdit
			/* 1 */ imports.NewTable("TLabeledEdit_AutoSelect", 0), // property AutoSelect
			/* 2 */ imports.NewTable("TLabeledEdit_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TLabeledEdit_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TLabeledEdit_ParentColor", 0), // property ParentColor
			/* 5 */ imports.NewTable("TLabeledEdit_ParentFont", 0), // property ParentFont
			/* 6 */ imports.NewTable("TLabeledEdit_ParentShowHint", 0), // property ParentShowHint
			/* 7 */ imports.NewTable("TLabeledEdit_OnContextPopup", 0), // event OnContextPopup
			/* 8 */ imports.NewTable("TLabeledEdit_OnDblClick", 0), // event OnDblClick
			/* 9 */ imports.NewTable("TLabeledEdit_OnDragDrop", 0), // event OnDragDrop
			/* 10 */ imports.NewTable("TLabeledEdit_OnDragOver", 0), // event OnDragOver
			/* 11 */ imports.NewTable("TLabeledEdit_OnEditingDone", 0), // event OnEditingDone
			/* 12 */ imports.NewTable("TLabeledEdit_OnEndDrag", 0), // event OnEndDrag
			/* 13 */ imports.NewTable("TLabeledEdit_OnMouseDown", 0), // event OnMouseDown
			/* 14 */ imports.NewTable("TLabeledEdit_OnMouseEnter", 0), // event OnMouseEnter
			/* 15 */ imports.NewTable("TLabeledEdit_OnMouseLeave", 0), // event OnMouseLeave
			/* 16 */ imports.NewTable("TLabeledEdit_OnMouseMove", 0), // event OnMouseMove
			/* 17 */ imports.NewTable("TLabeledEdit_OnMouseUp", 0), // event OnMouseUp
			/* 18 */ imports.NewTable("TLabeledEdit_OnMouseWheel", 0), // event OnMouseWheel
			/* 19 */ imports.NewTable("TLabeledEdit_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 20 */ imports.NewTable("TLabeledEdit_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 21 */ imports.NewTable("TLabeledEdit_OnStartDrag", 0), // event OnStartDrag
			/* 22 */ imports.NewTable("TLabeledEdit_TClass", 0), // function TLabeledEditClass
		}
	})
	return labeledEditImport
}
