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

// ITabSheet Parent: ICustomPage
type ITabSheet interface {
	ICustomPage
	PageControl() IPageControl                     // property PageControl Getter
	SetPageControl(value IPageControl)             // property PageControl Setter
	TabIndex() int32                               // property TabIndex Getter
	ParentFont() bool                              // property ParentFont Getter
	SetParentFont(value bool)                      // property ParentFont Setter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent)       // property event
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

type TTabSheet struct {
	TCustomPage
}

func (m *TTabSheet) PageControl() IPageControl {
	if !m.IsValid() {
		return nil
	}
	r := tabSheetAPI().SysCallN(1, 0, m.Instance())
	return AsPageControl(r)
}

func (m *TTabSheet) SetPageControl(value IPageControl) {
	if !m.IsValid() {
		return
	}
	tabSheetAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TTabSheet) TabIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := tabSheetAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TTabSheet) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := tabSheetAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTabSheet) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	tabSheetAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TTabSheet) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := tabSheetAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTabSheet) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	tabSheetAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TTabSheet) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 5, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 6, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 7, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 8, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 9, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 11, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 12, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 13, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 14, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 15, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 16, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTabSheet) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 17, tabSheetAPI(), api.MakeEventDataPtr(cb))
}

// NewTabSheet class constructor
func NewTabSheet(theOwner IComponent) ITabSheet {
	r := tabSheetAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsTabSheet(r)
}

func TTabSheetClass() types.TClass {
	r := tabSheetAPI().SysCallN(18)
	return types.TClass(r)
}

var (
	tabSheetOnce   base.Once
	tabSheetImport *imports.Imports = nil
)

func tabSheetAPI() *imports.Imports {
	tabSheetOnce.Do(func() {
		tabSheetImport = api.NewDefaultImports()
		tabSheetImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTabSheet_Create", 0), // constructor NewTabSheet
			/* 1 */ imports.NewTable("TTabSheet_PageControl", 0), // property PageControl
			/* 2 */ imports.NewTable("TTabSheet_TabIndex", 0), // property TabIndex
			/* 3 */ imports.NewTable("TTabSheet_ParentFont", 0), // property ParentFont
			/* 4 */ imports.NewTable("TTabSheet_ParentShowHint", 0), // property ParentShowHint
			/* 5 */ imports.NewTable("TTabSheet_OnContextPopup", 0), // event OnContextPopup
			/* 6 */ imports.NewTable("TTabSheet_OnDragDrop", 0), // event OnDragDrop
			/* 7 */ imports.NewTable("TTabSheet_OnDragOver", 0), // event OnDragOver
			/* 8 */ imports.NewTable("TTabSheet_OnEndDrag", 0), // event OnEndDrag
			/* 9 */ imports.NewTable("TTabSheet_OnMouseDown", 0), // event OnMouseDown
			/* 10 */ imports.NewTable("TTabSheet_OnMouseEnter", 0), // event OnMouseEnter
			/* 11 */ imports.NewTable("TTabSheet_OnMouseLeave", 0), // event OnMouseLeave
			/* 12 */ imports.NewTable("TTabSheet_OnMouseMove", 0), // event OnMouseMove
			/* 13 */ imports.NewTable("TTabSheet_OnMouseUp", 0), // event OnMouseUp
			/* 14 */ imports.NewTable("TTabSheet_OnMouseWheel", 0), // event OnMouseWheel
			/* 15 */ imports.NewTable("TTabSheet_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 16 */ imports.NewTable("TTabSheet_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 17 */ imports.NewTable("TTabSheet_OnStartDrag", 0), // event OnStartDrag
			/* 18 */ imports.NewTable("TTabSheet_TClass", 0), // function TTabSheetClass
		}
	})
	return tabSheetImport
}
