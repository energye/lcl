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

// IMaskEdit Parent: ICustomMaskEdit
type IMaskEdit interface {
	ICustomMaskEdit
	IsMasked() bool                                // property IsMasked Getter
	EditText() string                              // property EditText Getter
	SetEditText(value string)                      // property EditText Setter
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
	EditMask() string                              // property EditMask Getter
	SetEditMask(value string)                      // property EditMask Setter
	SpaceChar() uint16                             // property SpaceChar Getter
	SetSpaceChar(value uint16)                     // property SpaceChar Setter
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
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

type TMaskEdit struct {
	TCustomMaskEdit
}

func (m *TMaskEdit) IsMasked() bool {
	if !m.IsValid() {
		return false
	}
	r := maskEditAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TMaskEdit) EditText() string {
	if !m.IsValid() {
		return ""
	}
	r := maskEditAPI().SysCallN(2, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TMaskEdit) SetEditText(value string) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(2, 1, m.Instance(), api.PasStr(value))
}

func (m *TMaskEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := maskEditAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMaskEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TMaskEdit) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := maskEditAPI().SysCallN(4, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TMaskEdit) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TMaskEdit) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := maskEditAPI().SysCallN(5, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TMaskEdit) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TMaskEdit) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := maskEditAPI().SysCallN(6, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TMaskEdit) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TMaskEdit) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := maskEditAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMaskEdit) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TMaskEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := maskEditAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMaskEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TMaskEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := maskEditAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TMaskEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TMaskEdit) EditMask() string {
	if !m.IsValid() {
		return ""
	}
	r := maskEditAPI().SysCallN(10, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TMaskEdit) SetEditMask(value string) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(10, 1, m.Instance(), api.PasStr(value))
}

func (m *TMaskEdit) SpaceChar() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := maskEditAPI().SysCallN(11, 0, m.Instance())
	return uint16(r)
}

func (m *TMaskEdit) SetSpaceChar(value uint16) {
	if !m.IsValid() {
		return
	}
	maskEditAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TMaskEdit) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 13, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 14, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 16, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 17, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 21, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 22, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 23, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 24, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 25, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 26, maskEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TMaskEdit) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 27, maskEditAPI(), api.MakeEventDataPtr(cb))
}

// NewMaskEdit class constructor
func NewMaskEdit(theOwner IComponent) IMaskEdit {
	r := maskEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsMaskEdit(r)
}

func TMaskEditClass() types.TClass {
	r := maskEditAPI().SysCallN(28)
	return types.TClass(r)
}

var (
	maskEditOnce   base.Once
	maskEditImport *imports.Imports = nil
)

func maskEditAPI() *imports.Imports {
	maskEditOnce.Do(func() {
		maskEditImport = api.NewDefaultImports()
		maskEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMaskEdit_Create", 0), // constructor NewMaskEdit
			/* 1 */ imports.NewTable("TMaskEdit_IsMasked", 0), // property IsMasked
			/* 2 */ imports.NewTable("TMaskEdit_EditText", 0), // property EditText
			/* 3 */ imports.NewTable("TMaskEdit_AutoSelect", 0), // property AutoSelect
			/* 4 */ imports.NewTable("TMaskEdit_DragCursor", 0), // property DragCursor
			/* 5 */ imports.NewTable("TMaskEdit_DragKind", 0), // property DragKind
			/* 6 */ imports.NewTable("TMaskEdit_DragMode", 0), // property DragMode
			/* 7 */ imports.NewTable("TMaskEdit_ParentColor", 0), // property ParentColor
			/* 8 */ imports.NewTable("TMaskEdit_ParentFont", 0), // property ParentFont
			/* 9 */ imports.NewTable("TMaskEdit_ParentShowHint", 0), // property ParentShowHint
			/* 10 */ imports.NewTable("TMaskEdit_EditMask", 0), // property EditMask
			/* 11 */ imports.NewTable("TMaskEdit_SpaceChar", 0), // property SpaceChar
			/* 12 */ imports.NewTable("TMaskEdit_OnDblClick", 0), // event OnDblClick
			/* 13 */ imports.NewTable("TMaskEdit_OnDragDrop", 0), // event OnDragDrop
			/* 14 */ imports.NewTable("TMaskEdit_OnDragOver", 0), // event OnDragOver
			/* 15 */ imports.NewTable("TMaskEdit_OnEditingDone", 0), // event OnEditingDone
			/* 16 */ imports.NewTable("TMaskEdit_OnEndDock", 0), // event OnEndDock
			/* 17 */ imports.NewTable("TMaskEdit_OnEndDrag", 0), // event OnEndDrag
			/* 18 */ imports.NewTable("TMaskEdit_OnMouseDown", 0), // event OnMouseDown
			/* 19 */ imports.NewTable("TMaskEdit_OnMouseEnter", 0), // event OnMouseEnter
			/* 20 */ imports.NewTable("TMaskEdit_OnMouseLeave", 0), // event OnMouseLeave
			/* 21 */ imports.NewTable("TMaskEdit_OnMouseMove", 0), // event OnMouseMove
			/* 22 */ imports.NewTable("TMaskEdit_OnMouseUp", 0), // event OnMouseUp
			/* 23 */ imports.NewTable("TMaskEdit_OnMouseWheel", 0), // event OnMouseWheel
			/* 24 */ imports.NewTable("TMaskEdit_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 25 */ imports.NewTable("TMaskEdit_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 26 */ imports.NewTable("TMaskEdit_OnStartDock", 0), // event OnStartDock
			/* 27 */ imports.NewTable("TMaskEdit_OnStartDrag", 0), // event OnStartDrag
			/* 28 */ imports.NewTable("TMaskEdit_TClass", 0), // function TMaskEditClass
		}
	})
	return maskEditImport
}
