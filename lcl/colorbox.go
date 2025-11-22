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

// IColorBox Parent: ICustomColorBox
type IColorBox interface {
	ICustomColorBox
	DragCursor() types.TCursor                     // property DragCursor Getter
	SetDragCursor(value types.TCursor)             // property DragCursor Setter
	DragMode() types.TDragMode                     // property DragMode Getter
	SetDragMode(value types.TDragMode)             // property DragMode Setter
	ItemHeight() int32                             // property ItemHeight Getter
	SetItemHeight(value int32)                     // property ItemHeight Setter
	ItemWidth() int32                              // property ItemWidth Getter
	SetItemWidth(value int32)                      // property ItemWidth Setter
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
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
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

type TColorBox struct {
	TCustomColorBox
}

func (m *TColorBox) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := colorBoxAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TColorBox) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	colorBoxAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TColorBox) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := colorBoxAPI().SysCallN(2, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TColorBox) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	colorBoxAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TColorBox) ItemHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := colorBoxAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TColorBox) SetItemHeight(value int32) {
	if !m.IsValid() {
		return
	}
	colorBoxAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TColorBox) ItemWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := colorBoxAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TColorBox) SetItemWidth(value int32) {
	if !m.IsValid() {
		return
	}
	colorBoxAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TColorBox) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := colorBoxAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TColorBox) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	colorBoxAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TColorBox) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := colorBoxAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TColorBox) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	colorBoxAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TColorBox) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := colorBoxAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TColorBox) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	colorBoxAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TColorBox) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnCloseUp(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 10, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 12, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 13, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 14, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnDropDown(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 17, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 18, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 20, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 21, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 22, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 23, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 24, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 25, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TColorBox) SetOnSelect(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 26, colorBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewColorBox class constructor
func NewColorBox(owner IComponent) IColorBox {
	r := colorBoxAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsColorBox(r)
}

func TColorBoxClass() types.TClass {
	r := colorBoxAPI().SysCallN(27)
	return types.TClass(r)
}

var (
	colorBoxOnce   base.Once
	colorBoxImport *imports.Imports = nil
)

func colorBoxAPI() *imports.Imports {
	colorBoxOnce.Do(func() {
		colorBoxImport = api.NewDefaultImports()
		colorBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TColorBox_Create", 0), // constructor NewColorBox
			/* 1 */ imports.NewTable("TColorBox_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TColorBox_DragMode", 0), // property DragMode
			/* 3 */ imports.NewTable("TColorBox_ItemHeight", 0), // property ItemHeight
			/* 4 */ imports.NewTable("TColorBox_ItemWidth", 0), // property ItemWidth
			/* 5 */ imports.NewTable("TColorBox_ParentColor", 0), // property ParentColor
			/* 6 */ imports.NewTable("TColorBox_ParentFont", 0), // property ParentFont
			/* 7 */ imports.NewTable("TColorBox_ParentShowHint", 0), // property ParentShowHint
			/* 8 */ imports.NewTable("TColorBox_OnChange", 0), // event OnChange
			/* 9 */ imports.NewTable("TColorBox_OnCloseUp", 0), // event OnCloseUp
			/* 10 */ imports.NewTable("TColorBox_OnContextPopup", 0), // event OnContextPopup
			/* 11 */ imports.NewTable("TColorBox_OnDblClick", 0), // event OnDblClick
			/* 12 */ imports.NewTable("TColorBox_OnDragDrop", 0), // event OnDragDrop
			/* 13 */ imports.NewTable("TColorBox_OnDragOver", 0), // event OnDragOver
			/* 14 */ imports.NewTable("TColorBox_OnEndDrag", 0), // event OnEndDrag
			/* 15 */ imports.NewTable("TColorBox_OnDropDown", 0), // event OnDropDown
			/* 16 */ imports.NewTable("TColorBox_OnEditingDone", 0), // event OnEditingDone
			/* 17 */ imports.NewTable("TColorBox_OnMouseDown", 0), // event OnMouseDown
			/* 18 */ imports.NewTable("TColorBox_OnMouseEnter", 0), // event OnMouseEnter
			/* 19 */ imports.NewTable("TColorBox_OnMouseLeave", 0), // event OnMouseLeave
			/* 20 */ imports.NewTable("TColorBox_OnMouseMove", 0), // event OnMouseMove
			/* 21 */ imports.NewTable("TColorBox_OnMouseUp", 0), // event OnMouseUp
			/* 22 */ imports.NewTable("TColorBox_OnMouseWheel", 0), // event OnMouseWheel
			/* 23 */ imports.NewTable("TColorBox_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 24 */ imports.NewTable("TColorBox_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 25 */ imports.NewTable("TColorBox_OnStartDrag", 0), // event OnStartDrag
			/* 26 */ imports.NewTable("TColorBox_OnSelect", 0), // event OnSelect
			/* 27 */ imports.NewTable("TColorBox_TClass", 0), // function TColorBoxClass
		}
	})
	return colorBoxImport
}
