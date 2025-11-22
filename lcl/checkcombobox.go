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

// ICheckComboBox Parent: ICustomCheckCombo
type ICheckComboBox interface {
	ICustomCheckCombo
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
	Sorted() bool                                  // property Sorted Getter
	SetSorted(value bool)                          // property Sorted Setter
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnCloseUp(fn TNotifyEvent)                  // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnGetItems(fn TNotifyEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
	SetOnSelect(fn TNotifyEvent)                   // property event
}

type TCheckComboBox struct {
	TCustomCheckCombo
}

func (m *TCheckComboBox) BorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := checkComboBoxAPI().SysCallN(1, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TCheckComboBox) SetBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TCheckComboBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := checkComboBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TCheckComboBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TCheckComboBox) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := checkComboBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TCheckComboBox) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TCheckComboBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := checkComboBoxAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TCheckComboBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TCheckComboBox) ItemHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := checkComboBoxAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TCheckComboBox) SetItemHeight(value int32) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCheckComboBox) ItemWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := checkComboBoxAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TCheckComboBox) SetItemWidth(value int32) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TCheckComboBox) MaxLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := checkComboBoxAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TCheckComboBox) SetMaxLength(value int32) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCheckComboBox) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := checkComboBoxAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckComboBox) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckComboBox) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := checkComboBoxAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckComboBox) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckComboBox) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := checkComboBoxAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckComboBox) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckComboBox) Sorted() bool {
	if !m.IsValid() {
		return false
	}
	r := checkComboBoxAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCheckComboBox) SetSorted(value bool) {
	if !m.IsValid() {
		return
	}
	checkComboBoxAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TCheckComboBox) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnCloseUp(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 14, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 16, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 17, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 18, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnDropDown(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnGetItems(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 21, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 22, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 23, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 24, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 25, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 26, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 27, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 28, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 29, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 30, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCheckComboBox) SetOnSelect(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 31, checkComboBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewCheckComboBox class constructor
func NewCheckComboBox(owner IComponent) ICheckComboBox {
	r := checkComboBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCheckComboBox(r)
}

func TCheckComboBoxClass() types.TClass {
	r := checkComboBoxAPI().SysCallN(32)
	return types.TClass(r)
}

var (
	checkComboBoxOnce   base.Once
	checkComboBoxImport *imports.Imports = nil
)

func checkComboBoxAPI() *imports.Imports {
	checkComboBoxOnce.Do(func() {
		checkComboBoxImport = api.NewDefaultImports()
		checkComboBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCheckComboBox_Create", 0), // constructor NewCheckComboBox
			/* 1 */ imports.NewTable("TCheckComboBox_BorderStyle", 0), // property BorderStyle
			/* 2 */ imports.NewTable("TCheckComboBox_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TCheckComboBox_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TCheckComboBox_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TCheckComboBox_ItemHeight", 0), // property ItemHeight
			/* 6 */ imports.NewTable("TCheckComboBox_ItemWidth", 0), // property ItemWidth
			/* 7 */ imports.NewTable("TCheckComboBox_MaxLength", 0), // property MaxLength
			/* 8 */ imports.NewTable("TCheckComboBox_ParentColor", 0), // property ParentColor
			/* 9 */ imports.NewTable("TCheckComboBox_ParentFont", 0), // property ParentFont
			/* 10 */ imports.NewTable("TCheckComboBox_ParentShowHint", 0), // property ParentShowHint
			/* 11 */ imports.NewTable("TCheckComboBox_Sorted", 0), // property Sorted
			/* 12 */ imports.NewTable("TCheckComboBox_OnChange", 0), // event OnChange
			/* 13 */ imports.NewTable("TCheckComboBox_OnCloseUp", 0), // event OnCloseUp
			/* 14 */ imports.NewTable("TCheckComboBox_OnContextPopup", 0), // event OnContextPopup
			/* 15 */ imports.NewTable("TCheckComboBox_OnDblClick", 0), // event OnDblClick
			/* 16 */ imports.NewTable("TCheckComboBox_OnDragDrop", 0), // event OnDragDrop
			/* 17 */ imports.NewTable("TCheckComboBox_OnDragOver", 0), // event OnDragOver
			/* 18 */ imports.NewTable("TCheckComboBox_OnEndDrag", 0), // event OnEndDrag
			/* 19 */ imports.NewTable("TCheckComboBox_OnDropDown", 0), // event OnDropDown
			/* 20 */ imports.NewTable("TCheckComboBox_OnEditingDone", 0), // event OnEditingDone
			/* 21 */ imports.NewTable("TCheckComboBox_OnGetItems", 0), // event OnGetItems
			/* 22 */ imports.NewTable("TCheckComboBox_OnMouseDown", 0), // event OnMouseDown
			/* 23 */ imports.NewTable("TCheckComboBox_OnMouseEnter", 0), // event OnMouseEnter
			/* 24 */ imports.NewTable("TCheckComboBox_OnMouseLeave", 0), // event OnMouseLeave
			/* 25 */ imports.NewTable("TCheckComboBox_OnMouseMove", 0), // event OnMouseMove
			/* 26 */ imports.NewTable("TCheckComboBox_OnMouseUp", 0), // event OnMouseUp
			/* 27 */ imports.NewTable("TCheckComboBox_OnMouseWheel", 0), // event OnMouseWheel
			/* 28 */ imports.NewTable("TCheckComboBox_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 29 */ imports.NewTable("TCheckComboBox_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 30 */ imports.NewTable("TCheckComboBox_OnStartDrag", 0), // event OnStartDrag
			/* 31 */ imports.NewTable("TCheckComboBox_OnSelect", 0), // event OnSelect
			/* 32 */ imports.NewTable("TCheckComboBox_TClass", 0), // function TCheckComboBoxClass
		}
	})
	return checkComboBoxImport
}
