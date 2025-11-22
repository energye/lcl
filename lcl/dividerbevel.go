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

// IDividerBevel Parent: IGraphicControl
type IDividerBevel interface {
	IGraphicControl
	BevelStyle() types.TBevelStyle                   // property BevelStyle Getter
	SetBevelStyle(value types.TBevelStyle)           // property BevelStyle Setter
	BevelWidth() int32                               // property BevelWidth Getter
	SetBevelWidth(value int32)                       // property BevelWidth Setter
	CaptionSpacing() int32                           // property CaptionSpacing Getter
	SetCaptionSpacing(value int32)                   // property CaptionSpacing Setter
	DragCursor() types.TCursor                       // property DragCursor Getter
	SetDragCursor(value types.TCursor)               // property DragCursor Setter
	DragKind() types.TDragKind                       // property DragKind Getter
	SetDragKind(value types.TDragKind)               // property DragKind Setter
	DragMode() types.TDragMode                       // property DragMode Getter
	SetDragMode(value types.TDragMode)               // property DragMode Setter
	LeftIndent() int32                               // property LeftIndent Getter
	SetLeftIndent(value int32)                       // property LeftIndent Setter
	Orientation() types.TTrackBarOrientation         // property Orientation Getter
	SetOrientation(value types.TTrackBarOrientation) // property Orientation Setter
	ParentColor() bool                               // property ParentColor Getter
	SetParentColor(value bool)                       // property ParentColor Setter
	ParentFont() bool                                // property ParentFont Getter
	SetParentFont(value bool)                        // property ParentFont Setter
	ParentShowHint() bool                            // property ParentShowHint Getter
	SetParentShowHint(value bool)                    // property ParentShowHint Setter
	Style() types.TGrabStyle                         // property Style Getter
	SetStyle(value types.TGrabStyle)                 // property Style Setter
	Transparent() bool                               // property Transparent Getter
	SetTransparent(value bool)                       // property Transparent Setter
	SetOnContextPopup(fn TContextPopupEvent)         // property event
	SetOnDblClick(fn TNotifyEvent)                   // property event
	SetOnDragDrop(fn TDragDropEvent)                 // property event
	SetOnDragOver(fn TDragOverEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                   // property event
	SetOnMouseDown(fn TMouseEvent)                   // property event
	SetOnMouseEnter(fn TNotifyEvent)                 // property event
	SetOnMouseLeave(fn TNotifyEvent)                 // property event
	SetOnMouseMove(fn TMouseMoveEvent)               // property event
	SetOnMouseUp(fn TMouseEvent)                     // property event
	SetOnStartDrag(fn TStartDragEvent)               // property event
}

type TDividerBevel struct {
	TGraphicControl
}

func (m *TDividerBevel) BevelStyle() types.TBevelStyle {
	if !m.IsValid() {
		return 0
	}
	r := dividerBevelAPI().SysCallN(1, 0, m.Instance())
	return types.TBevelStyle(r)
}

func (m *TDividerBevel) SetBevelStyle(value types.TBevelStyle) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TDividerBevel) BevelWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dividerBevelAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TDividerBevel) SetBevelWidth(value int32) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TDividerBevel) CaptionSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dividerBevelAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TDividerBevel) SetCaptionSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TDividerBevel) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := dividerBevelAPI().SysCallN(4, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TDividerBevel) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TDividerBevel) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := dividerBevelAPI().SysCallN(5, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TDividerBevel) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TDividerBevel) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := dividerBevelAPI().SysCallN(6, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TDividerBevel) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TDividerBevel) LeftIndent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dividerBevelAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TDividerBevel) SetLeftIndent(value int32) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TDividerBevel) Orientation() types.TTrackBarOrientation {
	if !m.IsValid() {
		return 0
	}
	r := dividerBevelAPI().SysCallN(8, 0, m.Instance())
	return types.TTrackBarOrientation(r)
}

func (m *TDividerBevel) SetOrientation(value types.TTrackBarOrientation) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TDividerBevel) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := dividerBevelAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDividerBevel) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TDividerBevel) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := dividerBevelAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDividerBevel) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TDividerBevel) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := dividerBevelAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDividerBevel) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TDividerBevel) Style() types.TGrabStyle {
	if !m.IsValid() {
		return 0
	}
	r := dividerBevelAPI().SysCallN(12, 0, m.Instance())
	return types.TGrabStyle(r)
}

func (m *TDividerBevel) SetStyle(value types.TGrabStyle) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TDividerBevel) Transparent() bool {
	if !m.IsValid() {
		return false
	}
	r := dividerBevelAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDividerBevel) SetTransparent(value bool) {
	if !m.IsValid() {
		return
	}
	dividerBevelAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TDividerBevel) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 14, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 16, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 17, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 18, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 19, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 21, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 22, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 23, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDividerBevel) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 24, dividerBevelAPI(), api.MakeEventDataPtr(cb))
}

// NewDividerBevel class constructor
func NewDividerBevel(owner IComponent) IDividerBevel {
	r := dividerBevelAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsDividerBevel(r)
}

func TDividerBevelClass() types.TClass {
	r := dividerBevelAPI().SysCallN(25)
	return types.TClass(r)
}

var (
	dividerBevelOnce   base.Once
	dividerBevelImport *imports.Imports = nil
)

func dividerBevelAPI() *imports.Imports {
	dividerBevelOnce.Do(func() {
		dividerBevelImport = api.NewDefaultImports()
		dividerBevelImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDividerBevel_Create", 0), // constructor NewDividerBevel
			/* 1 */ imports.NewTable("TDividerBevel_BevelStyle", 0), // property BevelStyle
			/* 2 */ imports.NewTable("TDividerBevel_BevelWidth", 0), // property BevelWidth
			/* 3 */ imports.NewTable("TDividerBevel_CaptionSpacing", 0), // property CaptionSpacing
			/* 4 */ imports.NewTable("TDividerBevel_DragCursor", 0), // property DragCursor
			/* 5 */ imports.NewTable("TDividerBevel_DragKind", 0), // property DragKind
			/* 6 */ imports.NewTable("TDividerBevel_DragMode", 0), // property DragMode
			/* 7 */ imports.NewTable("TDividerBevel_LeftIndent", 0), // property LeftIndent
			/* 8 */ imports.NewTable("TDividerBevel_Orientation", 0), // property Orientation
			/* 9 */ imports.NewTable("TDividerBevel_ParentColor", 0), // property ParentColor
			/* 10 */ imports.NewTable("TDividerBevel_ParentFont", 0), // property ParentFont
			/* 11 */ imports.NewTable("TDividerBevel_ParentShowHint", 0), // property ParentShowHint
			/* 12 */ imports.NewTable("TDividerBevel_Style", 0), // property Style
			/* 13 */ imports.NewTable("TDividerBevel_Transparent", 0), // property Transparent
			/* 14 */ imports.NewTable("TDividerBevel_OnContextPopup", 0), // event OnContextPopup
			/* 15 */ imports.NewTable("TDividerBevel_OnDblClick", 0), // event OnDblClick
			/* 16 */ imports.NewTable("TDividerBevel_OnDragDrop", 0), // event OnDragDrop
			/* 17 */ imports.NewTable("TDividerBevel_OnDragOver", 0), // event OnDragOver
			/* 18 */ imports.NewTable("TDividerBevel_OnEndDrag", 0), // event OnEndDrag
			/* 19 */ imports.NewTable("TDividerBevel_OnMouseDown", 0), // event OnMouseDown
			/* 20 */ imports.NewTable("TDividerBevel_OnMouseEnter", 0), // event OnMouseEnter
			/* 21 */ imports.NewTable("TDividerBevel_OnMouseLeave", 0), // event OnMouseLeave
			/* 22 */ imports.NewTable("TDividerBevel_OnMouseMove", 0), // event OnMouseMove
			/* 23 */ imports.NewTable("TDividerBevel_OnMouseUp", 0), // event OnMouseUp
			/* 24 */ imports.NewTable("TDividerBevel_OnStartDrag", 0), // event OnStartDrag
			/* 25 */ imports.NewTable("TDividerBevel_TClass", 0), // function TDividerBevelClass
		}
	})
	return dividerBevelImport
}
