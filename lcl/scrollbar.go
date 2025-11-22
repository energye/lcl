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

// IScrollBar Parent: ICustomScrollBar
type IScrollBar interface {
	ICustomScrollBar
	DragCursor() types.TCursor               // property DragCursor Getter
	SetDragCursor(value types.TCursor)       // property DragCursor Setter
	DragKind() types.TDragKind               // property DragKind Getter
	SetDragKind(value types.TDragKind)       // property DragKind Setter
	DragMode() types.TDragMode               // property DragMode Getter
	SetDragMode(value types.TDragMode)       // property DragMode Setter
	ParentShowHint() bool                    // property ParentShowHint Getter
	SetParentShowHint(value bool)            // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent) // property event
	SetOnDragDrop(fn TDragDropEvent)         // property event
	SetOnDragOver(fn TDragOverEvent)         // property event
	SetOnEndDrag(fn TEndDragEvent)           // property event
	SetOnStartDrag(fn TStartDragEvent)       // property event
}

type TScrollBar struct {
	TCustomScrollBar
}

func (m *TScrollBar) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := scrollBarAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TScrollBar) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	scrollBarAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TScrollBar) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := scrollBarAPI().SysCallN(2, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TScrollBar) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	scrollBarAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TScrollBar) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := scrollBarAPI().SysCallN(3, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TScrollBar) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	scrollBarAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TScrollBar) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := scrollBarAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TScrollBar) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	scrollBarAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TScrollBar) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 5, scrollBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBar) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 6, scrollBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBar) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 7, scrollBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBar) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 8, scrollBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScrollBar) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 9, scrollBarAPI(), api.MakeEventDataPtr(cb))
}

// NewScrollBar class constructor
func NewScrollBar(owner IComponent) IScrollBar {
	r := scrollBarAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsScrollBar(r)
}

func TScrollBarClass() types.TClass {
	r := scrollBarAPI().SysCallN(10)
	return types.TClass(r)
}

var (
	scrollBarOnce   base.Once
	scrollBarImport *imports.Imports = nil
)

func scrollBarAPI() *imports.Imports {
	scrollBarOnce.Do(func() {
		scrollBarImport = api.NewDefaultImports()
		scrollBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TScrollBar_Create", 0), // constructor NewScrollBar
			/* 1 */ imports.NewTable("TScrollBar_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TScrollBar_DragKind", 0), // property DragKind
			/* 3 */ imports.NewTable("TScrollBar_DragMode", 0), // property DragMode
			/* 4 */ imports.NewTable("TScrollBar_ParentShowHint", 0), // property ParentShowHint
			/* 5 */ imports.NewTable("TScrollBar_OnContextPopup", 0), // event OnContextPopup
			/* 6 */ imports.NewTable("TScrollBar_OnDragDrop", 0), // event OnDragDrop
			/* 7 */ imports.NewTable("TScrollBar_OnDragOver", 0), // event OnDragOver
			/* 8 */ imports.NewTable("TScrollBar_OnEndDrag", 0), // event OnEndDrag
			/* 9 */ imports.NewTable("TScrollBar_OnStartDrag", 0), // event OnStartDrag
			/* 10 */ imports.NewTable("TScrollBar_TClass", 0), // function TScrollBarClass
		}
	})
	return scrollBarImport
}
