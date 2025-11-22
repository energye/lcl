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

// IImage Parent: ICustomImage
type IImage interface {
	ICustomImage
	DragCursor() types.TCursor                      // property DragCursor Getter
	SetDragCursor(value types.TCursor)              // property DragCursor Setter
	DragMode() types.TDragMode                      // property DragMode Getter
	SetDragMode(value types.TDragMode)              // property DragMode Setter
	ParentShowHint() bool                           // property ParentShowHint Getter
	SetParentShowHint(value bool)                   // property ParentShowHint Setter
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnPaint(fn TNotifyEvent)                     // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

type TImage struct {
	TCustomImage
}

func (m *TImage) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := imageAPI().SysCallN(1, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TImage) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	imageAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TImage) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := imageAPI().SysCallN(2, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TImage) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	imageAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TImage) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := imageAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TImage) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	imageAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TImage) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 4, imageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImage) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 5, imageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImage) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 6, imageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImage) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 7, imageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImage) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 8, imageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImage) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 9, imageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImage) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 10, imageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImage) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 11, imageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImage) SetOnPaint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 12, imageAPI(), api.MakeEventDataPtr(cb))
}

func (m *TImage) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 13, imageAPI(), api.MakeEventDataPtr(cb))
}

// NewImage class constructor
func NewImage(owner IComponent) IImage {
	r := imageAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsImage(r)
}

func TImageClass() types.TClass {
	r := imageAPI().SysCallN(14)
	return types.TClass(r)
}

var (
	imageOnce   base.Once
	imageImport *imports.Imports = nil
)

func imageAPI() *imports.Imports {
	imageOnce.Do(func() {
		imageImport = api.NewDefaultImports()
		imageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TImage_Create", 0), // constructor NewImage
			/* 1 */ imports.NewTable("TImage_DragCursor", 0), // property DragCursor
			/* 2 */ imports.NewTable("TImage_DragMode", 0), // property DragMode
			/* 3 */ imports.NewTable("TImage_ParentShowHint", 0), // property ParentShowHint
			/* 4 */ imports.NewTable("TImage_OnContextPopup", 0), // event OnContextPopup
			/* 5 */ imports.NewTable("TImage_OnDblClick", 0), // event OnDblClick
			/* 6 */ imports.NewTable("TImage_OnDragDrop", 0), // event OnDragDrop
			/* 7 */ imports.NewTable("TImage_OnDragOver", 0), // event OnDragOver
			/* 8 */ imports.NewTable("TImage_OnEndDrag", 0), // event OnEndDrag
			/* 9 */ imports.NewTable("TImage_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 10 */ imports.NewTable("TImage_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 11 */ imports.NewTable("TImage_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 12 */ imports.NewTable("TImage_OnPaint", 0), // event OnPaint
			/* 13 */ imports.NewTable("TImage_OnStartDrag", 0), // event OnStartDrag
			/* 14 */ imports.NewTable("TImage_TClass", 0), // function TImageClass
		}
	})
	return imageImport
}
