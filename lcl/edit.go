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

// IEdit Parent: ICustomEdit
type IEdit interface {
	ICustomEdit
	AutoSelected() bool                            // property AutoSelected Getter
	SetAutoSelected(value bool)                    // property AutoSelected Setter
	AutoSelect() bool                              // property AutoSelect Getter
	SetAutoSelect(value bool)                      // property AutoSelect Setter
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

type TEdit struct {
	TCustomEdit
}

func (m *TEdit) AutoSelected() bool {
	if !m.IsValid() {
		return false
	}
	r := editAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEdit) SetAutoSelected(value bool) {
	if !m.IsValid() {
		return
	}
	editAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := editAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	editAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TEdit) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := editAPI().SysCallN(3, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TEdit) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	editAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TEdit) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := editAPI().SysCallN(4, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TEdit) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	editAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TEdit) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := editAPI().SysCallN(5, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TEdit) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	editAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TEdit) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := editAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEdit) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	editAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := editAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	editAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := editAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	editAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TEdit) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 9, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 11, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 12, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 14, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 15, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 18, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 19, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 20, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 22, editAPI(), api.MakeEventDataPtr(cb))
}

func (m *TEdit) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 23, editAPI(), api.MakeEventDataPtr(cb))
}

// NewEdit class constructor
func NewEdit(owner IComponent) IEdit {
	r := editAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEdit(r)
}

func TEditClass() types.TClass {
	r := editAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	editOnce   base.Once
	editImport *imports.Imports = nil
)

func editAPI() *imports.Imports {
	editOnce.Do(func() {
		editImport = api.NewDefaultImports()
		editImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEdit_Create", 0), // constructor NewEdit
			/* 1 */ imports.NewTable("TEdit_AutoSelected", 0), // property AutoSelected
			/* 2 */ imports.NewTable("TEdit_AutoSelect", 0), // property AutoSelect
			/* 3 */ imports.NewTable("TEdit_DragCursor", 0), // property DragCursor
			/* 4 */ imports.NewTable("TEdit_DragKind", 0), // property DragKind
			/* 5 */ imports.NewTable("TEdit_DragMode", 0), // property DragMode
			/* 6 */ imports.NewTable("TEdit_ParentColor", 0), // property ParentColor
			/* 7 */ imports.NewTable("TEdit_ParentFont", 0), // property ParentFont
			/* 8 */ imports.NewTable("TEdit_ParentShowHint", 0), // property ParentShowHint
			/* 9 */ imports.NewTable("TEdit_OnContextPopup", 0), // event OnContextPopup
			/* 10 */ imports.NewTable("TEdit_OnDblClick", 0), // event OnDblClick
			/* 11 */ imports.NewTable("TEdit_OnDragDrop", 0), // event OnDragDrop
			/* 12 */ imports.NewTable("TEdit_OnDragOver", 0), // event OnDragOver
			/* 13 */ imports.NewTable("TEdit_OnEditingDone", 0), // event OnEditingDone
			/* 14 */ imports.NewTable("TEdit_OnEndDrag", 0), // event OnEndDrag
			/* 15 */ imports.NewTable("TEdit_OnMouseDown", 0), // event OnMouseDown
			/* 16 */ imports.NewTable("TEdit_OnMouseEnter", 0), // event OnMouseEnter
			/* 17 */ imports.NewTable("TEdit_OnMouseLeave", 0), // event OnMouseLeave
			/* 18 */ imports.NewTable("TEdit_OnMouseMove", 0), // event OnMouseMove
			/* 19 */ imports.NewTable("TEdit_OnMouseUp", 0), // event OnMouseUp
			/* 20 */ imports.NewTable("TEdit_OnMouseWheel", 0), // event OnMouseWheel
			/* 21 */ imports.NewTable("TEdit_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 22 */ imports.NewTable("TEdit_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 23 */ imports.NewTable("TEdit_OnStartDrag", 0), // event OnStartDrag
			/* 24 */ imports.NewTable("TEdit_TClass", 0), // function TEditClass
		}
	})
	return editImport
}
