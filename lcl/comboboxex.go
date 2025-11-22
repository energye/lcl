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

// IComboBoxEx Parent: ICustomComboBoxEx
type IComboBoxEx interface {
	ICustomComboBoxEx
	// BorderStyle
	//  properties which are not supported by all descendents
	BorderStyle() types.TBorderStyle               // property BorderStyle Getter
	SetBorderStyle(value types.TBorderStyle)       // property BorderStyle Setter
	DragCursor() types.TCursor                     // property DragCursor Getter
	SetDragCursor(value types.TCursor)             // property DragCursor Setter
	DragKind() types.TDragKind                     // property DragKind Getter
	SetDragKind(value types.TDragKind)             // property DragKind Setter
	DragMode() types.TDragMode                     // property DragMode Getter
	SetDragMode(value types.TDragMode)             // property DragMode Setter
	ItemHeight() int32                             // property ItemHeight Getter
	SetItemHeight(value int32)                     // property ItemHeight Setter
	ItemWidth() int32                              // property ItemWidth Getter
	SetItemWidth(value int32)                      // property ItemWidth Setter
	MaxLength() int32                              // property MaxLength Getter
	SetMaxLength(value int32)                      // property MaxLength Setter
	ParentColor() bool                             // property ParentColor Getter
	SetParentColor(value bool)                     // property ParentColor Setter
	ParentFont() bool                              // property ParentFont Getter
	SetParentFont(value bool)                      // property ParentFont Setter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnCloseUp(fn TNotifyEvent)                  // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnGetItems(fn TNotifyEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnSelect(fn TNotifyEvent)                   // property event
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

type TComboBoxEx struct {
	TCustomComboBoxEx
}

func (m *TComboBoxEx) BorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxExAPI().SysCallN(1, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TComboBoxEx) SetBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TComboBoxEx) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxExAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TComboBoxEx) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TComboBoxEx) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxExAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TComboBoxEx) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TComboBoxEx) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxExAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TComboBoxEx) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TComboBoxEx) ItemHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxExAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TComboBoxEx) SetItemHeight(value int32) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TComboBoxEx) ItemWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxExAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TComboBoxEx) SetItemWidth(value int32) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TComboBoxEx) MaxLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxExAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TComboBoxEx) SetMaxLength(value int32) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TComboBoxEx) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := comboBoxExAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TComboBoxEx) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TComboBoxEx) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := comboBoxExAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TComboBoxEx) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TComboBoxEx) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := comboBoxExAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TComboBoxEx) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	comboBoxExAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TComboBoxEx) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnCloseUp(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 13, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 14, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 15, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 16, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnDropDown(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 19, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 20, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnGetItems(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 21, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 22, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 23, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 24, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 25, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 26, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 27, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 28, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 29, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnSelect(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 30, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 31, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBoxEx) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 32, comboBoxExAPI(), api.MakeEventDataPtr(cb))
}

// NewComboBoxEx class constructor
func NewComboBoxEx(theOwner IComponent) IComboBoxEx {
	r := comboBoxExAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsComboBoxEx(r)
}

func TComboBoxExClass() types.TClass {
	r := comboBoxExAPI().SysCallN(33)
	return types.TClass(r)
}

var (
	comboBoxExOnce   base.Once
	comboBoxExImport *imports.Imports = nil
)

func comboBoxExAPI() *imports.Imports {
	comboBoxExOnce.Do(func() {
		comboBoxExImport = api.NewDefaultImports()
		comboBoxExImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TComboBoxEx_Create", 0), // constructor NewComboBoxEx
			/* 1 */ imports.NewTable("TComboBoxEx_BorderStyle", 0), // property BorderStyle
			/* 2 */ imports.NewTable("TComboBoxEx_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TComboBoxEx_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TComboBoxEx_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TComboBoxEx_ItemHeight", 0), // property ItemHeight
			/* 6 */ imports.NewTable("TComboBoxEx_ItemWidth", 0), // property ItemWidth
			/* 7 */ imports.NewTable("TComboBoxEx_MaxLength", 0), // property MaxLength
			/* 8 */ imports.NewTable("TComboBoxEx_ParentColor", 0), // property ParentColor
			/* 9 */ imports.NewTable("TComboBoxEx_ParentFont", 0), // property ParentFont
			/* 10 */ imports.NewTable("TComboBoxEx_ParentShowHint", 0), // property ParentShowHint
			/* 11 */ imports.NewTable("TComboBoxEx_OnChange", 0), // event OnChange
			/* 12 */ imports.NewTable("TComboBoxEx_OnCloseUp", 0), // event OnCloseUp
			/* 13 */ imports.NewTable("TComboBoxEx_OnContextPopup", 0), // event OnContextPopup
			/* 14 */ imports.NewTable("TComboBoxEx_OnDblClick", 0), // event OnDblClick
			/* 15 */ imports.NewTable("TComboBoxEx_OnDragDrop", 0), // event OnDragDrop
			/* 16 */ imports.NewTable("TComboBoxEx_OnDragOver", 0), // event OnDragOver
			/* 17 */ imports.NewTable("TComboBoxEx_OnDropDown", 0), // event OnDropDown
			/* 18 */ imports.NewTable("TComboBoxEx_OnEditingDone", 0), // event OnEditingDone
			/* 19 */ imports.NewTable("TComboBoxEx_OnEndDock", 0), // event OnEndDock
			/* 20 */ imports.NewTable("TComboBoxEx_OnEndDrag", 0), // event OnEndDrag
			/* 21 */ imports.NewTable("TComboBoxEx_OnGetItems", 0), // event OnGetItems
			/* 22 */ imports.NewTable("TComboBoxEx_OnMouseDown", 0), // event OnMouseDown
			/* 23 */ imports.NewTable("TComboBoxEx_OnMouseEnter", 0), // event OnMouseEnter
			/* 24 */ imports.NewTable("TComboBoxEx_OnMouseLeave", 0), // event OnMouseLeave
			/* 25 */ imports.NewTable("TComboBoxEx_OnMouseMove", 0), // event OnMouseMove
			/* 26 */ imports.NewTable("TComboBoxEx_OnMouseUp", 0), // event OnMouseUp
			/* 27 */ imports.NewTable("TComboBoxEx_OnMouseWheel", 0), // event OnMouseWheel
			/* 28 */ imports.NewTable("TComboBoxEx_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 29 */ imports.NewTable("TComboBoxEx_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 30 */ imports.NewTable("TComboBoxEx_OnSelect", 0), // event OnSelect
			/* 31 */ imports.NewTable("TComboBoxEx_OnStartDock", 0), // event OnStartDock
			/* 32 */ imports.NewTable("TComboBoxEx_OnStartDrag", 0), // event OnStartDrag
			/* 33 */ imports.NewTable("TComboBoxEx_TClass", 0), // function TComboBoxExClass
		}
	})
	return comboBoxExImport
}
