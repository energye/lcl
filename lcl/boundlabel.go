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

// IBoundLabel Parent: ICustomLabel
type IBoundLabel interface {
	ICustomLabel
	FocusControl() IWinControl                     // property FocusControl Getter
	SetFocusControl(value IWinControl)             // property FocusControl Setter
	Alignment() types.TAlignment                   // property Alignment Getter
	SetAlignment(value types.TAlignment)           // property Alignment Setter
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
	ShowAccelChar() bool                           // property ShowAccelChar Getter
	SetShowAccelChar(value bool)                   // property ShowAccelChar Setter
	Layout() types.TTextLayout                     // property Layout Getter
	SetLayout(value types.TTextLayout)             // property Layout Setter
	WordWrap() bool                                // property WordWrap Getter
	SetWordWrap(value bool)                        // property WordWrap Setter
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
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

type TBoundLabel struct {
	TCustomLabel
}

func (m *TBoundLabel) FocusControl() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := boundLabelAPI().SysCallN(1, 0, m.Instance())
	return AsWinControl(r)
}

func (m *TBoundLabel) SetFocusControl(value IWinControl) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TBoundLabel) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := boundLabelAPI().SysCallN(2, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TBoundLabel) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TBoundLabel) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := boundLabelAPI().SysCallN(3, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TBoundLabel) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TBoundLabel) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := boundLabelAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TBoundLabel) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TBoundLabel) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := boundLabelAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBoundLabel) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TBoundLabel) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := boundLabelAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBoundLabel) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TBoundLabel) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := boundLabelAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBoundLabel) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TBoundLabel) ShowAccelChar() bool {
	if !m.IsValid() {
		return false
	}
	r := boundLabelAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBoundLabel) SetShowAccelChar(value bool) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TBoundLabel) Layout() types.TTextLayout {
	if !m.IsValid() {
		return 0
	}
	r := boundLabelAPI().SysCallN(9, 0, m.Instance())
	return types.TTextLayout(r)
}

func (m *TBoundLabel) SetLayout(value types.TTextLayout) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TBoundLabel) WordWrap() bool {
	if !m.IsValid() {
		return false
	}
	r := boundLabelAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBoundLabel) SetWordWrap(value bool) {
	if !m.IsValid() {
		return
	}
	boundLabelAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TBoundLabel) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 12, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 13, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 14, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 15, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 17, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 18, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 19, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 20, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 22, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TBoundLabel) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 23, boundLabelAPI(), api.MakeEventDataPtr(cb))
}

// NewBoundLabel class constructor
func NewBoundLabel(theOwner IComponent) IBoundLabel {
	r := boundLabelAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsBoundLabel(r)
}

func TBoundLabelClass() types.TClass {
	r := boundLabelAPI().SysCallN(24)
	return types.TClass(r)
}

var (
	boundLabelOnce   base.Once
	boundLabelImport *imports.Imports = nil
)

func boundLabelAPI() *imports.Imports {
	boundLabelOnce.Do(func() {
		boundLabelImport = api.NewDefaultImports()
		boundLabelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBoundLabel_Create", 0), // constructor NewBoundLabel
			/* 1 */ imports.NewTable("TBoundLabel_FocusControl", 0), // property FocusControl
			/* 2 */ imports.NewTable("TBoundLabel_Alignment", 0), // property Alignment
			/* 3 */ imports.NewTable("TBoundLabel_DragCursor", 0), // property DragCursor
			/* 4 */ imports.NewTable("TBoundLabel_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TBoundLabel_ParentColor", 0), // property ParentColor
			/* 6 */ imports.NewTable("TBoundLabel_ParentFont", 0), // property ParentFont
			/* 7 */ imports.NewTable("TBoundLabel_ParentShowHint", 0), // property ParentShowHint
			/* 8 */ imports.NewTable("TBoundLabel_ShowAccelChar", 0), // property ShowAccelChar
			/* 9 */ imports.NewTable("TBoundLabel_Layout", 0), // property Layout
			/* 10 */ imports.NewTable("TBoundLabel_WordWrap", 0), // property WordWrap
			/* 11 */ imports.NewTable("TBoundLabel_OnDblClick", 0), // event OnDblClick
			/* 12 */ imports.NewTable("TBoundLabel_OnDragDrop", 0), // event OnDragDrop
			/* 13 */ imports.NewTable("TBoundLabel_OnDragOver", 0), // event OnDragOver
			/* 14 */ imports.NewTable("TBoundLabel_OnEndDrag", 0), // event OnEndDrag
			/* 15 */ imports.NewTable("TBoundLabel_OnMouseDown", 0), // event OnMouseDown
			/* 16 */ imports.NewTable("TBoundLabel_OnMouseEnter", 0), // event OnMouseEnter
			/* 17 */ imports.NewTable("TBoundLabel_OnMouseLeave", 0), // event OnMouseLeave
			/* 18 */ imports.NewTable("TBoundLabel_OnMouseMove", 0), // event OnMouseMove
			/* 19 */ imports.NewTable("TBoundLabel_OnMouseUp", 0), // event OnMouseUp
			/* 20 */ imports.NewTable("TBoundLabel_OnMouseWheel", 0), // event OnMouseWheel
			/* 21 */ imports.NewTable("TBoundLabel_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 22 */ imports.NewTable("TBoundLabel_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 23 */ imports.NewTable("TBoundLabel_OnStartDrag", 0), // event OnStartDrag
			/* 24 */ imports.NewTable("TBoundLabel_TClass", 0), // function TBoundLabelClass
		}
	})
	return boundLabelImport
}
