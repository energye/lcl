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

// IToolButton Parent: IGraphicControl
type IToolButton interface {
	IGraphicControl
	CheckMenuDropdown() bool                                                                           // function
	PointInArrow(X int32, Y int32) bool                                                                // function
	Click()                                                                                            // procedure
	ArrowClick()                                                                                       // procedure
	GetCurrentIcon(imageList *ICustomImageList, theIndex *int32, theEffect *types.TGraphicsDrawEffect) // procedure
	Index() int32                                                                                      // property Index Getter
	AllowAllUp() bool                                                                                  // property AllowAllUp Getter
	SetAllowAllUp(value bool)                                                                          // property AllowAllUp Setter
	Down() bool                                                                                        // property Down Getter
	SetDown(value bool)                                                                                // property Down Setter
	DragCursor() types.TCursor                                                                         // property DragCursor Getter
	SetDragCursor(value types.TCursor)                                                                 // property DragCursor Setter
	DragKind() types.TDragKind                                                                         // property DragKind Getter
	SetDragKind(value types.TDragKind)                                                                 // property DragKind Setter
	DragMode() types.TDragMode                                                                         // property DragMode Getter
	SetDragMode(value types.TDragMode)                                                                 // property DragMode Setter
	DropdownMenu() IPopupMenu                                                                          // property DropdownMenu Getter
	SetDropdownMenu(value IPopupMenu)                                                                  // property DropdownMenu Setter
	Grouped() bool                                                                                     // property Grouped Getter
	SetGrouped(value bool)                                                                             // property Grouped Setter
	ImageIndex() int32                                                                                 // property ImageIndex Getter
	SetImageIndex(value int32)                                                                         // property ImageIndex Setter
	Indeterminate() bool                                                                               // property Indeterminate Getter
	SetIndeterminate(value bool)                                                                       // property Indeterminate Setter
	Marked() bool                                                                                      // property Marked Getter
	SetMarked(value bool)                                                                              // property Marked Setter
	MenuItem() IMenuItem                                                                               // property MenuItem Getter
	SetMenuItem(value IMenuItem)                                                                       // property MenuItem Setter
	ParentShowHint() bool                                                                              // property ParentShowHint Getter
	SetParentShowHint(value bool)                                                                      // property ParentShowHint Setter
	ShowCaption() bool                                                                                 // property ShowCaption Getter
	SetShowCaption(value bool)                                                                         // property ShowCaption Setter
	Style() types.TToolButtonStyle                                                                     // property Style Getter
	SetStyle(value types.TToolButtonStyle)                                                             // property Style Setter
	Wrap() bool                                                                                        // property Wrap Getter
	SetWrap(value bool)                                                                                // property Wrap Setter
	SetOnArrowClick(fn TNotifyEvent)                                                                   // property event
	SetOnContextPopup(fn TContextPopupEvent)                                                           // property event
	SetOnDragDrop(fn TDragDropEvent)                                                                   // property event
	SetOnDragOver(fn TDragOverEvent)                                                                   // property event
	SetOnEndDock(fn TEndDragEvent)                                                                     // property event
	SetOnEndDrag(fn TEndDragEvent)                                                                     // property event
	SetOnMouseDown(fn TMouseEvent)                                                                     // property event
	SetOnMouseEnter(fn TNotifyEvent)                                                                   // property event
	SetOnMouseLeave(fn TNotifyEvent)                                                                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                                                 // property event
	SetOnMouseUp(fn TMouseEvent)                                                                       // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                                               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                                                     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                                                       // property event
	SetOnStartDock(fn TStartDockEvent)                                                                 // property event
	SetOnStartDrag(fn TStartDragEvent)                                                                 // property event
}

type TToolButton struct {
	TGraphicControl
}

func (m *TToolButton) CheckMenuDropdown() bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TToolButton) PointInArrow(X int32, Y int32) bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(2, m.Instance(), uintptr(X), uintptr(Y))
	return api.GoBool(r)
}

func (m *TToolButton) Click() {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(3, m.Instance())
}

func (m *TToolButton) ArrowClick() {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(4, m.Instance())
}

func (m *TToolButton) GetCurrentIcon(imageList *ICustomImageList, theIndex *int32, theEffect *types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	imageListPtr := base.GetObjectUintptr(*imageList)
	theIndexPtr := uintptr(*theIndex)
	theEffectPtr := uintptr(*theEffect)
	toolButtonAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&imageListPtr)), uintptr(base.UnsafePointer(&theIndexPtr)), uintptr(base.UnsafePointer(&theEffectPtr)))
	*imageList = AsCustomImageList(imageListPtr)
	*theIndex = int32(theIndexPtr)
	*theEffect = types.TGraphicsDrawEffect(theEffectPtr)
}

func (m *TToolButton) Index() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolButtonAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TToolButton) AllowAllUp() bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolButton) SetAllowAllUp(value bool) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolButton) Down() bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolButton) SetDown(value bool) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolButton) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := toolButtonAPI().SysCallN(9, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TToolButton) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TToolButton) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := toolButtonAPI().SysCallN(10, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TToolButton) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TToolButton) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := toolButtonAPI().SysCallN(11, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TToolButton) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TToolButton) DropdownMenu() IPopupMenu {
	if !m.IsValid() {
		return nil
	}
	r := toolButtonAPI().SysCallN(12, 0, m.Instance())
	return AsPopupMenu(r)
}

func (m *TToolButton) SetDropdownMenu(value IPopupMenu) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(12, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TToolButton) Grouped() bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolButton) SetGrouped(value bool) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolButton) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := toolButtonAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TToolButton) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TToolButton) Indeterminate() bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolButton) SetIndeterminate(value bool) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolButton) Marked() bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolButton) SetMarked(value bool) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolButton) MenuItem() IMenuItem {
	if !m.IsValid() {
		return nil
	}
	r := toolButtonAPI().SysCallN(17, 0, m.Instance())
	return AsMenuItem(r)
}

func (m *TToolButton) SetMenuItem(value IMenuItem) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(17, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TToolButton) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolButton) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolButton) ShowCaption() bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolButton) SetShowCaption(value bool) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolButton) Style() types.TToolButtonStyle {
	if !m.IsValid() {
		return 0
	}
	r := toolButtonAPI().SysCallN(20, 0, m.Instance())
	return types.TToolButtonStyle(r)
}

func (m *TToolButton) SetStyle(value types.TToolButtonStyle) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TToolButton) Wrap() bool {
	if !m.IsValid() {
		return false
	}
	r := toolButtonAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TToolButton) SetWrap(value bool) {
	if !m.IsValid() {
		return
	}
	toolButtonAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TToolButton) SetOnArrowClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 22, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 23, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 24, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 25, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 26, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 27, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 28, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 29, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 30, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 31, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 32, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 33, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 34, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 35, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 36, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

func (m *TToolButton) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 37, toolButtonAPI(), api.MakeEventDataPtr(cb))
}

// NewToolButton class constructor
func NewToolButton(theOwner IComponent) IToolButton {
	r := toolButtonAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsToolButton(r)
}

func TToolButtonClass() types.TClass {
	r := toolButtonAPI().SysCallN(38)
	return types.TClass(r)
}

var (
	toolButtonOnce   base.Once
	toolButtonImport *imports.Imports = nil
)

func toolButtonAPI() *imports.Imports {
	toolButtonOnce.Do(func() {
		toolButtonImport = api.NewDefaultImports()
		toolButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TToolButton_Create", 0), // constructor NewToolButton
			/* 1 */ imports.NewTable("TToolButton_CheckMenuDropdown", 0), // function CheckMenuDropdown
			/* 2 */ imports.NewTable("TToolButton_PointInArrow", 0), // function PointInArrow
			/* 3 */ imports.NewTable("TToolButton_Click", 0), // procedure Click
			/* 4 */ imports.NewTable("TToolButton_ArrowClick", 0), // procedure ArrowClick
			/* 5 */ imports.NewTable("TToolButton_GetCurrentIcon", 0), // procedure GetCurrentIcon
			/* 6 */ imports.NewTable("TToolButton_Index", 0), // property Index
			/* 7 */ imports.NewTable("TToolButton_AllowAllUp", 0), // property AllowAllUp
			/* 8 */ imports.NewTable("TToolButton_Down", 0), // property Down
			/* 9 */ imports.NewTable("TToolButton_DragCursor", 0), // property DragCursor
			/* 10 */ imports.NewTable("TToolButton_DragKind", 0), // property DragKind
			/* 11 */ imports.NewTable("TToolButton_DragMode", 0), // property DragMode
			/* 12 */ imports.NewTable("TToolButton_DropdownMenu", 0), // property DropdownMenu
			/* 13 */ imports.NewTable("TToolButton_Grouped", 0), // property Grouped
			/* 14 */ imports.NewTable("TToolButton_ImageIndex", 0), // property ImageIndex
			/* 15 */ imports.NewTable("TToolButton_Indeterminate", 0), // property Indeterminate
			/* 16 */ imports.NewTable("TToolButton_Marked", 0), // property Marked
			/* 17 */ imports.NewTable("TToolButton_MenuItem", 0), // property MenuItem
			/* 18 */ imports.NewTable("TToolButton_ParentShowHint", 0), // property ParentShowHint
			/* 19 */ imports.NewTable("TToolButton_ShowCaption", 0), // property ShowCaption
			/* 20 */ imports.NewTable("TToolButton_Style", 0), // property Style
			/* 21 */ imports.NewTable("TToolButton_Wrap", 0), // property Wrap
			/* 22 */ imports.NewTable("TToolButton_OnArrowClick", 0), // event OnArrowClick
			/* 23 */ imports.NewTable("TToolButton_OnContextPopup", 0), // event OnContextPopup
			/* 24 */ imports.NewTable("TToolButton_OnDragDrop", 0), // event OnDragDrop
			/* 25 */ imports.NewTable("TToolButton_OnDragOver", 0), // event OnDragOver
			/* 26 */ imports.NewTable("TToolButton_OnEndDock", 0), // event OnEndDock
			/* 27 */ imports.NewTable("TToolButton_OnEndDrag", 0), // event OnEndDrag
			/* 28 */ imports.NewTable("TToolButton_OnMouseDown", 0), // event OnMouseDown
			/* 29 */ imports.NewTable("TToolButton_OnMouseEnter", 0), // event OnMouseEnter
			/* 30 */ imports.NewTable("TToolButton_OnMouseLeave", 0), // event OnMouseLeave
			/* 31 */ imports.NewTable("TToolButton_OnMouseMove", 0), // event OnMouseMove
			/* 32 */ imports.NewTable("TToolButton_OnMouseUp", 0), // event OnMouseUp
			/* 33 */ imports.NewTable("TToolButton_OnMouseWheel", 0), // event OnMouseWheel
			/* 34 */ imports.NewTable("TToolButton_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 35 */ imports.NewTable("TToolButton_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 36 */ imports.NewTable("TToolButton_OnStartDock", 0), // event OnStartDock
			/* 37 */ imports.NewTable("TToolButton_OnStartDrag", 0), // event OnStartDrag
			/* 38 */ imports.NewTable("TToolButton_TClass", 0), // function TToolButtonClass
		}
	})
	return toolButtonImport
}
