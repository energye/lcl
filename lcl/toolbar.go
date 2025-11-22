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

// IToolBar Parent: IToolWindow
type IToolBar interface {
	IToolWindow
	GetEnumeratorToToolBarEnumerator() IToolBarEnumerator      // function
	SetButtonSize(newButtonWidth int32, newButtonHeight int32) // procedure
	ButtonCount() int32                                        // property ButtonCount Getter
	Buttons(index int32) IToolButton                           // property Buttons Getter
	ButtonList() IList                                         // property ButtonList Getter
	RowCount() int32                                           // property RowCount Getter
	ButtonDropWidth() int32                                    // property ButtonDropWidth Getter
	ButtonHeight() int32                                       // property ButtonHeight Getter
	SetButtonHeight(value int32)                               // property ButtonHeight Setter
	ButtonWidth() int32                                        // property ButtonWidth Getter
	SetButtonWidth(value int32)                                // property ButtonWidth Setter
	DisabledImages() ICustomImageList                          // property DisabledImages Getter
	SetDisabledImages(value ICustomImageList)                  // property DisabledImages Setter
	DragCursor() types.TCursor                                 // property DragCursor Getter
	SetDragCursor(value types.TCursor)                         // property DragCursor Setter
	DragKind() types.TDragKind                                 // property DragKind Getter
	SetDragKind(value types.TDragKind)                         // property DragKind Setter
	DragMode() types.TDragMode                                 // property DragMode Getter
	SetDragMode(value types.TDragMode)                         // property DragMode Setter
	DropDownWidth() int32                                      // property DropDownWidth Getter
	SetDropDownWidth(value int32)                              // property DropDownWidth Setter
	Flat() bool                                                // property Flat Getter
	SetFlat(value bool)                                        // property Flat Setter
	HotImages() ICustomImageList                               // property HotImages Getter
	SetHotImages(value ICustomImageList)                       // property HotImages Setter
	Images() ICustomImageList                                  // property Images Getter
	SetImages(value ICustomImageList)                          // property Images Setter
	ImagesWidth() int32                                        // property ImagesWidth Getter
	SetImagesWidth(value int32)                                // property ImagesWidth Setter
	Indent() int32                                             // property Indent Getter
	SetIndent(value int32)                                     // property Indent Setter
	List() bool                                                // property List Getter
	SetList(value bool)                                        // property List Setter
	ParentColor() bool                                         // property ParentColor Getter
	SetParentColor(value bool)                                 // property ParentColor Setter
	ParentFont() bool                                          // property ParentFont Getter
	SetParentFont(value bool)                                  // property ParentFont Setter
	ParentShowHint() bool                                      // property ParentShowHint Getter
	SetParentShowHint(value bool)                              // property ParentShowHint Setter
	ShowCaptions() bool                                        // property ShowCaptions Getter
	SetShowCaptions(value bool)                                // property ShowCaptions Setter
	Transparent() bool                                         // property Transparent Getter
	SetTransparent(value bool)                                 // property Transparent Setter
	Wrapable() bool                                            // property Wrapable Getter
	SetWrapable(value bool)                                    // property Wrapable Setter
	SetOnContextPopup(fn TContextPopupEvent)                   // property event
	SetOnDblClick(fn TNotifyEvent)                             // property event
	SetOnDragDrop(fn TDragDropEvent)                           // property event
	SetOnDragOver(fn TDragOverEvent)                           // property event
	SetOnPaintButton(fn TToolBarOnPaintButton)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                             // property event
	SetOnMouseDown(fn TMouseEvent)                             // property event
	SetOnMouseEnter(fn TNotifyEvent)                           // property event
	SetOnMouseLeave(fn TNotifyEvent)                           // property event
	SetOnMouseMove(fn TMouseMoveEvent)                         // property event
	SetOnMouseUp(fn TMouseEvent)                               // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                       // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)             // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)               // property event
	SetOnStartDrag(fn TStartDragEvent)                         // property event
}

type TToolBar struct {
	TToolWindow
}

func (m *TToolBar) GetEnumeratorToToolBarEnumerator() IToolBarEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := toolBarAPI().SysCallN(1, m.Instance())
	return AsToolBarEnumerator(r)
}

func (m *TToolBar) SetButtonSize(newButtonWidth int32, newButtonHeight int32) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(2, m.Instance(), uintptr(newButtonWidth), uintptr(newButtonHeight))
}

func (m *TToolBar) ButtonCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TToolBar) Buttons(index int32) IToolButton {
	if !m.IsValid() {
		return nil
	}
	r := toolBarAPI().SysCallN(4, m.Instance(), uintptr(index))
	return AsToolButton(r)
}

func (m *TToolBar) ButtonList() IList {
	if !m.IsValid() {
		return nil
	}
	r := toolBarAPI().SysCallN(5, m.Instance())
	return AsList(r)
}

func (m *TToolBar) RowCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TToolBar) ButtonDropWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TToolBar) ButtonHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TToolBar) SetButtonHeight(value int32) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TToolBar) ButtonWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TToolBar) SetButtonWidth(value int32) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TToolBar) DisabledImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := toolBarAPI().SysCallN(10, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TToolBar) SetDisabledImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TToolBar) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(11, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TToolBar) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TToolBar) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(12, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TToolBar) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TToolBar) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(13, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TToolBar) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TToolBar) DropDownWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TToolBar) SetDropDownWidth(value int32) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TToolBar) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := toolBarAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolBar) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolBar) HotImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := toolBarAPI().SysCallN(16, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TToolBar) SetHotImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TToolBar) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := toolBarAPI().SysCallN(17, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TToolBar) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(17, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TToolBar) ImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TToolBar) SetImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TToolBar) Indent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolBarAPI().SysCallN(19, 0, m.Instance())
	return int32(r)
}

func (m *TToolBar) SetIndent(value int32) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TToolBar) List() bool {
	if !m.IsValid() {
		return false
	}
	r := toolBarAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolBar) SetList(value bool) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolBar) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := toolBarAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolBar) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolBar) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := toolBarAPI().SysCallN(22, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolBar) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(22, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolBar) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := toolBarAPI().SysCallN(23, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolBar) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(23, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolBar) ShowCaptions() bool {
	if !m.IsValid() {
		return false
	}
	r := toolBarAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolBar) SetShowCaptions(value bool) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolBar) Transparent() bool {
	if !m.IsValid() {
		return false
	}
	r := toolBarAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolBar) SetTransparent(value bool) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolBar) Wrapable() bool {
	if !m.IsValid() {
		return false
	}
	r := toolBarAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolBar) SetWrapable(value bool) {
	if !m.IsValid() {
		return
	}
	toolBarAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolBar) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 27, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 28, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 29, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 30, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnPaintButton(fn TToolBarOnPaintButton) {
	if !m.IsValid() {
		return
	}
	cb := makeTToolBarOnPaintButton(fn)
	base.SetEvent(m, 31, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 32, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 33, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 34, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 35, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 36, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 37, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 38, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 39, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 40, toolBarAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolBar) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 41, toolBarAPI(), api.MakeEventDataPtr(cb))
}

// NewToolBar class constructor
func NewToolBar(theOwner IComponent) IToolBar {
	r := toolBarAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsToolBar(r)
}

func TToolBarClass() types.TClass {
	r := toolBarAPI().SysCallN(42)
	return types.TClass(r)
}

var (
	toolBarOnce   base.Once
	toolBarImport *imports.Imports = nil
)

func toolBarAPI() *imports.Imports {
	toolBarOnce.Do(func() {
		toolBarImport = api.NewDefaultImports()
		toolBarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TToolBar_Create", 0), // constructor NewToolBar
			/* 1 */ imports.NewTable("TToolBar_GetEnumeratorToToolBarEnumerator", 0), // function GetEnumeratorToToolBarEnumerator
			/* 2 */ imports.NewTable("TToolBar_SetButtonSize", 0), // procedure SetButtonSize
			/* 3 */ imports.NewTable("TToolBar_ButtonCount", 0), // property ButtonCount
			/* 4 */ imports.NewTable("TToolBar_Buttons", 0), // property Buttons
			/* 5 */ imports.NewTable("TToolBar_ButtonList", 0), // property ButtonList
			/* 6 */ imports.NewTable("TToolBar_RowCount", 0), // property RowCount
			/* 7 */ imports.NewTable("TToolBar_ButtonDropWidth", 0), // property ButtonDropWidth
			/* 8 */ imports.NewTable("TToolBar_ButtonHeight", 0), // property ButtonHeight
			/* 9 */ imports.NewTable("TToolBar_ButtonWidth", 0), // property ButtonWidth
			/* 10 */ imports.NewTable("TToolBar_DisabledImages", 0), // property DisabledImages
			/* 11 */ imports.NewTable("TToolBar_DragCursor", 0), // property DragCursor
			/* 12 */ imports.NewTable("TToolBar_DragKind", 0), // property DragKind
			/* 13 */ imports.NewTable("TToolBar_DragMode", 0), // property DragMode
			/* 14 */ imports.NewTable("TToolBar_DropDownWidth", 0), // property DropDownWidth
			/* 15 */ imports.NewTable("TToolBar_Flat", 0), // property Flat
			/* 16 */ imports.NewTable("TToolBar_HotImages", 0), // property HotImages
			/* 17 */ imports.NewTable("TToolBar_Images", 0), // property Images
			/* 18 */ imports.NewTable("TToolBar_ImagesWidth", 0), // property ImagesWidth
			/* 19 */ imports.NewTable("TToolBar_Indent", 0), // property Indent
			/* 20 */ imports.NewTable("TToolBar_List", 0), // property List
			/* 21 */ imports.NewTable("TToolBar_ParentColor", 0), // property ParentColor
			/* 22 */ imports.NewTable("TToolBar_ParentFont", 0), // property ParentFont
			/* 23 */ imports.NewTable("TToolBar_ParentShowHint", 0), // property ParentShowHint
			/* 24 */ imports.NewTable("TToolBar_ShowCaptions", 0), // property ShowCaptions
			/* 25 */ imports.NewTable("TToolBar_Transparent", 0), // property Transparent
			/* 26 */ imports.NewTable("TToolBar_Wrapable", 0), // property Wrapable
			/* 27 */ imports.NewTable("TToolBar_OnContextPopup", 0), // event OnContextPopup
			/* 28 */ imports.NewTable("TToolBar_OnDblClick", 0), // event OnDblClick
			/* 29 */ imports.NewTable("TToolBar_OnDragDrop", 0), // event OnDragDrop
			/* 30 */ imports.NewTable("TToolBar_OnDragOver", 0), // event OnDragOver
			/* 31 */ imports.NewTable("TToolBar_OnPaintButton", 0), // event OnPaintButton
			/* 32 */ imports.NewTable("TToolBar_OnEndDrag", 0), // event OnEndDrag
			/* 33 */ imports.NewTable("TToolBar_OnMouseDown", 0), // event OnMouseDown
			/* 34 */ imports.NewTable("TToolBar_OnMouseEnter", 0), // event OnMouseEnter
			/* 35 */ imports.NewTable("TToolBar_OnMouseLeave", 0), // event OnMouseLeave
			/* 36 */ imports.NewTable("TToolBar_OnMouseMove", 0), // event OnMouseMove
			/* 37 */ imports.NewTable("TToolBar_OnMouseUp", 0), // event OnMouseUp
			/* 38 */ imports.NewTable("TToolBar_OnMouseWheel", 0), // event OnMouseWheel
			/* 39 */ imports.NewTable("TToolBar_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 40 */ imports.NewTable("TToolBar_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 41 */ imports.NewTable("TToolBar_OnStartDrag", 0), // event OnStartDrag
			/* 42 */ imports.NewTable("TToolBar_TClass", 0), // function TToolBarClass
		}
	})
	return toolBarImport
}
