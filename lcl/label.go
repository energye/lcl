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

// ILabel Parent: ICustomLabel
type ILabel interface {
	ICustomLabel
	Alignment() types.TAlignment                    // property Alignment Getter
	SetAlignment(value types.TAlignment)            // property Alignment Setter
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
	DragKind() types.TDragKind                      // property DragKind Getter
	SetDragKind(value types.TDragKind)              // property DragKind Setter
	DragMode() types.TDragMode                      // property DragMode Getter
	SetDragMode(value types.TDragMode)              // property DragMode Setter
	FocusControl() IWinControl                      // property FocusControl Getter
	SetFocusControl(value IWinControl)              // property FocusControl Setter
	Layout() types.TTextLayout                      // property Layout Getter
	SetLayout(value types.TTextLayout)              // property Layout Setter
	OptimalFill() bool                              // property OptimalFill Getter
	SetOptimalFill(value bool)                      // property OptimalFill Setter
	ParentColor() bool                              // property ParentColor Getter
	SetParentColor(value bool)                      // property ParentColor Setter
	ParentFont() bool                               // property ParentFont Getter
	SetParentFont(value bool)                       // property ParentFont Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	ShowAccelChar() bool                            // property ShowAccelChar Getter
	SetShowAccelChar(value bool)                    // property ShowAccelChar Setter
	Transparent() bool                              // property Transparent Getter
	SetTransparent(value bool)                      // property Transparent Setter
	WordWrap() bool                                 // property WordWrap Getter
	SetWordWrap(value bool)                         // property WordWrap Setter
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

type TLabel struct {
	TCustomLabel
}

func (m *TLabel) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := labelAPI().SysCallN(1, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TLabel) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TLabel) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := labelAPI().SysCallN(2, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TLabel) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TLabel) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := labelAPI().SysCallN(3, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TLabel) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TLabel) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := labelAPI().SysCallN(4, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TLabel) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TLabel) FocusControl() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := labelAPI().SysCallN(5, 0, m.Instance())
	return AsWinControl(r)
}

func (m *TLabel) SetFocusControl(value IWinControl) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLabel) Layout() types.TTextLayout {
	if !m.IsValid() {
		return 0
	}
	r := labelAPI().SysCallN(6, 0, m.Instance())
	return types.TTextLayout(r)
}

func (m *TLabel) SetLayout(value types.TTextLayout) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TLabel) OptimalFill() bool {
	if !m.IsValid() {
		return false
	}
	r := labelAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabel) SetOptimalFill(value bool) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabel) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := labelAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabel) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabel) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := labelAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabel) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabel) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := labelAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabel) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabel) ShowAccelChar() bool {
	if !m.IsValid() {
		return false
	}
	r := labelAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabel) SetShowAccelChar(value bool) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabel) Transparent() bool {
	if !m.IsValid() {
		return false
	}
	r := labelAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabel) SetTransparent(value bool) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabel) WordWrap() bool {
	if !m.IsValid() {
		return false
	}
	r := labelAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLabel) SetWordWrap(value bool) {
	if !m.IsValid() {
		return
	}
	labelAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TLabel) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 14, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 16, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 17, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 18, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 19, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 21, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 22, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 23, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 24, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 25, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 26, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 27, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 28, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 29, labelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLabel) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 30, labelAPI(), api.MakeEventDataPtr(cb))
}

// NewLabel class constructor
func NewLabel(theOwner IComponent) ILabel {
	r := labelAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsLabel(r)
}

func TLabelClass() types.TClass {
	r := labelAPI().SysCallN(31)
	return types.TClass(r)
}

var (
	labelOnce   base.Once
	labelImport *imports.Imports = nil
)

func labelAPI() *imports.Imports {
	labelOnce.Do(func() {
		labelImport = api.NewDefaultImports()
		labelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLabel_Create", 0), // constructor NewLabel
			/* 1 */ imports.NewTable("TLabel_Alignment", 0), // property Alignment
			/* 2 */ imports.NewTable("TLabel_DragCursor", 0), // property DragCursor
			/* 3 */ imports.NewTable("TLabel_DragKind", 0), // property DragKind
			/* 4 */ imports.NewTable("TLabel_DragMode", 0), // property DragMode
			/* 5 */ imports.NewTable("TLabel_FocusControl", 0), // property FocusControl
			/* 6 */ imports.NewTable("TLabel_Layout", 0), // property Layout
			/* 7 */ imports.NewTable("TLabel_OptimalFill", 0), // property OptimalFill
			/* 8 */ imports.NewTable("TLabel_ParentColor", 0), // property ParentColor
			/* 9 */ imports.NewTable("TLabel_ParentFont", 0), // property ParentFont
			/* 10 */ imports.NewTable("TLabel_ParentShowHint", 0), // property ParentShowHint
			/* 11 */ imports.NewTable("TLabel_ShowAccelChar", 0), // property ShowAccelChar
			/* 12 */ imports.NewTable("TLabel_Transparent", 0), // property Transparent
			/* 13 */ imports.NewTable("TLabel_WordWrap", 0), // property WordWrap
			/* 14 */ imports.NewTable("TLabel_OnContextPopup", 0), // event OnContextPopup
			/* 15 */ imports.NewTable("TLabel_OnDblClick", 0), // event OnDblClick
			/* 16 */ imports.NewTable("TLabel_OnDragDrop", 0), // event OnDragDrop
			/* 17 */ imports.NewTable("TLabel_OnDragOver", 0), // event OnDragOver
			/* 18 */ imports.NewTable("TLabel_OnEndDrag", 0), // event OnEndDrag
			/* 19 */ imports.NewTable("TLabel_OnMouseDown", 0), // event OnMouseDown
			/* 20 */ imports.NewTable("TLabel_OnMouseEnter", 0), // event OnMouseEnter
			/* 21 */ imports.NewTable("TLabel_OnMouseLeave", 0), // event OnMouseLeave
			/* 22 */ imports.NewTable("TLabel_OnMouseMove", 0), // event OnMouseMove
			/* 23 */ imports.NewTable("TLabel_OnMouseUp", 0), // event OnMouseUp
			/* 24 */ imports.NewTable("TLabel_OnMouseWheel", 0), // event OnMouseWheel
			/* 25 */ imports.NewTable("TLabel_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 26 */ imports.NewTable("TLabel_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 27 */ imports.NewTable("TLabel_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 28 */ imports.NewTable("TLabel_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 29 */ imports.NewTable("TLabel_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 30 */ imports.NewTable("TLabel_OnStartDrag", 0), // event OnStartDrag
			/* 31 */ imports.NewTable("TLabel_TClass", 0), // function TLabelClass
		}
	})
	return labelImport
}
