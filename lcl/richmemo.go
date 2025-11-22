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

// IRichMemo Parent: ICustomRichMemo
type IRichMemo interface {
	ICustomRichMemo
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
	Rtf() string                                   // property Rtf Getter
	SetRtf(value string)                           // property Rtf Setter
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

type TRichMemo struct {
	TCustomRichMemo
}

func (m *TRichMemo) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := richMemoAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TRichMemo) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	richMemoAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TRichMemo) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := richMemoAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TRichMemo) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	richMemoAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TRichMemo) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := richMemoAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TRichMemo) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	richMemoAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TRichMemo) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := richMemoAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRichMemo) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	richMemoAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TRichMemo) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := richMemoAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRichMemo) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	richMemoAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TRichMemo) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := richMemoAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRichMemo) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	richMemoAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TRichMemo) Rtf() string {
	if !m.IsValid() {
		return ""
	}
	r := richMemoAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TRichMemo) SetRtf(value string) {
	if !m.IsValid() {
		return
	}
	richMemoAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TRichMemo) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 8, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 10, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 11, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 13, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 14, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 17, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 18, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 19, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 20, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 21, richMemoAPI(), api.MakeEventDataPtr(cb))
}

func (m *TRichMemo) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 22, richMemoAPI(), api.MakeEventDataPtr(cb))
}

// NewRichMemo class constructor
func NewRichMemo(owner IComponent) IRichMemo {
	r := richMemoAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsRichMemo(r)
}

func TRichMemoClass() types.TClass {
	r := richMemoAPI().SysCallN(23)
	return types.TClass(r)
}

var (
	richMemoOnce   base.Once
	richMemoImport *imports.Imports = nil
)

func richMemoAPI() *imports.Imports {
	richMemoOnce.Do(func() {
		richMemoImport = api.NewDefaultImports()
		richMemoImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRichMemo_Create", 0), // constructor NewRichMemo
			/* 1 */ imports.NewTable("TRichMemo_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TRichMemo_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TRichMemo_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TRichMemo_ParentColor", 0), // property ParentColor
			/* 5 */ imports.NewTable("TRichMemo_ParentFont", 0), // property ParentFont
			/* 6 */ imports.NewTable("TRichMemo_ParentShowHint", 0), // property ParentShowHint
			/* 7 */ imports.NewTable("TRichMemo_Rtf", 0), // property Rtf
			/* 8 */ imports.NewTable("TRichMemo_OnContextPopup", 0), // event OnContextPopup
			/* 9 */ imports.NewTable("TRichMemo_OnDblClick", 0), // event OnDblClick
			/* 10 */ imports.NewTable("TRichMemo_OnDragDrop", 0), // event OnDragDrop
			/* 11 */ imports.NewTable("TRichMemo_OnDragOver", 0), // event OnDragOver
			/* 12 */ imports.NewTable("TRichMemo_OnEditingDone", 0), // event OnEditingDone
			/* 13 */ imports.NewTable("TRichMemo_OnEndDrag", 0), // event OnEndDrag
			/* 14 */ imports.NewTable("TRichMemo_OnMouseDown", 0), // event OnMouseDown
			/* 15 */ imports.NewTable("TRichMemo_OnMouseEnter", 0), // event OnMouseEnter
			/* 16 */ imports.NewTable("TRichMemo_OnMouseLeave", 0), // event OnMouseLeave
			/* 17 */ imports.NewTable("TRichMemo_OnMouseMove", 0), // event OnMouseMove
			/* 18 */ imports.NewTable("TRichMemo_OnMouseUp", 0), // event OnMouseUp
			/* 19 */ imports.NewTable("TRichMemo_OnMouseWheel", 0), // event OnMouseWheel
			/* 20 */ imports.NewTable("TRichMemo_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 21 */ imports.NewTable("TRichMemo_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 22 */ imports.NewTable("TRichMemo_OnStartDrag", 0), // event OnStartDrag
			/* 23 */ imports.NewTable("TRichMemo_TClass", 0), // function TRichMemoClass
		}
	})
	return richMemoImport
}
