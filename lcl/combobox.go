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

// IComboBox Parent: ICustomComboBox
type IComboBox interface {
	ICustomComboBox
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
	SetOnDrawItem(fn TDrawItemEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnGetItems(fn TNotifyEvent)                 // property event
	SetOnMeasureItem(fn TMeasureItemEvent)         // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnSelect(fn TNotifyEvent)                   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

type TComboBox struct {
	TCustomComboBox
}

func (m *TComboBox) BorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxAPI().SysCallN(1, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TComboBox) SetBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TComboBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TComboBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TComboBox) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TComboBox) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TComboBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TComboBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TComboBox) ItemHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TComboBox) SetItemHeight(value int32) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TComboBox) ItemWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TComboBox) SetItemWidth(value int32) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TComboBox) MaxLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := comboBoxAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TComboBox) SetMaxLength(value int32) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TComboBox) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := comboBoxAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TComboBox) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TComboBox) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := comboBoxAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TComboBox) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TComboBox) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := comboBoxAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TComboBox) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TComboBox) Sorted() bool {
	if !m.IsValid() {
		return false
	}
	r := comboBoxAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TComboBox) SetSorted(value bool) {
	if !m.IsValid() {
		return
	}
	comboBoxAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TComboBox) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnCloseUp(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 13, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 14, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 16, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 17, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnDrawItem(fn TDrawItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDrawItemEvent(fn)
	base.SetEvent(m, 18, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 19, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnDropDown(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 21, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnGetItems(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 22, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnMeasureItem(fn TMeasureItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMeasureItemEvent(fn)
	base.SetEvent(m, 23, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 24, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 25, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 26, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 27, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 28, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 29, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 30, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 31, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnSelect(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 32, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TComboBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 33, comboBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewComboBox class constructor
func NewComboBox(theOwner IComponent) IComboBox {
	r := comboBoxAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsComboBox(r)
}

func TComboBoxClass() types.TClass {
	r := comboBoxAPI().SysCallN(34)
	return types.TClass(r)
}

var (
	comboBoxOnce   base.Once
	comboBoxImport *imports.Imports = nil
)

func comboBoxAPI() *imports.Imports {
	comboBoxOnce.Do(func() {
		comboBoxImport = api.NewDefaultImports()
		comboBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TComboBox_Create", 0), // constructor NewComboBox
			/* 1 */ imports.NewTable("TComboBox_BorderStyle", 0), // property BorderStyle
			/* 2 */ imports.NewTable("TComboBox_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TComboBox_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TComboBox_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TComboBox_ItemHeight", 0), // property ItemHeight
			/* 6 */ imports.NewTable("TComboBox_ItemWidth", 0), // property ItemWidth
			/* 7 */ imports.NewTable("TComboBox_MaxLength", 0), // property MaxLength
			/* 8 */ imports.NewTable("TComboBox_ParentColor", 0), // property ParentColor
			/* 9 */ imports.NewTable("TComboBox_ParentFont", 0), // property ParentFont
			/* 10 */ imports.NewTable("TComboBox_ParentShowHint", 0), // property ParentShowHint
			/* 11 */ imports.NewTable("TComboBox_Sorted", 0), // property Sorted
			/* 12 */ imports.NewTable("TComboBox_OnChange", 0), // event OnChange
			/* 13 */ imports.NewTable("TComboBox_OnCloseUp", 0), // event OnCloseUp
			/* 14 */ imports.NewTable("TComboBox_OnContextPopup", 0), // event OnContextPopup
			/* 15 */ imports.NewTable("TComboBox_OnDblClick", 0), // event OnDblClick
			/* 16 */ imports.NewTable("TComboBox_OnDragDrop", 0), // event OnDragDrop
			/* 17 */ imports.NewTable("TComboBox_OnDragOver", 0), // event OnDragOver
			/* 18 */ imports.NewTable("TComboBox_OnDrawItem", 0), // event OnDrawItem
			/* 19 */ imports.NewTable("TComboBox_OnEndDrag", 0), // event OnEndDrag
			/* 20 */ imports.NewTable("TComboBox_OnDropDown", 0), // event OnDropDown
			/* 21 */ imports.NewTable("TComboBox_OnEditingDone", 0), // event OnEditingDone
			/* 22 */ imports.NewTable("TComboBox_OnGetItems", 0), // event OnGetItems
			/* 23 */ imports.NewTable("TComboBox_OnMeasureItem", 0), // event OnMeasureItem
			/* 24 */ imports.NewTable("TComboBox_OnMouseDown", 0), // event OnMouseDown
			/* 25 */ imports.NewTable("TComboBox_OnMouseEnter", 0), // event OnMouseEnter
			/* 26 */ imports.NewTable("TComboBox_OnMouseLeave", 0), // event OnMouseLeave
			/* 27 */ imports.NewTable("TComboBox_OnMouseMove", 0), // event OnMouseMove
			/* 28 */ imports.NewTable("TComboBox_OnMouseUp", 0), // event OnMouseUp
			/* 29 */ imports.NewTable("TComboBox_OnMouseWheel", 0), // event OnMouseWheel
			/* 30 */ imports.NewTable("TComboBox_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 31 */ imports.NewTable("TComboBox_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 32 */ imports.NewTable("TComboBox_OnSelect", 0), // event OnSelect
			/* 33 */ imports.NewTable("TComboBox_OnStartDrag", 0), // event OnStartDrag
			/* 34 */ imports.NewTable("TComboBox_TClass", 0), // function TComboBoxClass
		}
	})
	return comboBoxImport
}
