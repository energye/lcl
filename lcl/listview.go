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

// IListView Parent: ICustomListView
type IListView interface {
	ICustomListView
	AllocBy() int32                                                      // property AllocBy Getter
	SetAllocBy(value int32)                                              // property AllocBy Setter
	AutoSort() bool                                                      // property AutoSort Getter
	SetAutoSort(value bool)                                              // property AutoSort Setter
	AutoSortIndicator() bool                                             // property AutoSortIndicator Getter
	SetAutoSortIndicator(value bool)                                     // property AutoSortIndicator Setter
	AutoWidthLastColumn() bool                                           // property AutoWidthLastColumn Getter
	SetAutoWidthLastColumn(value bool)                                   // property AutoWidthLastColumn Setter
	Columns() IListColumns                                               // property Columns Getter
	SetColumns(value IListColumns)                                       // property Columns Setter
	ColumnClick() bool                                                   // property ColumnClick Getter
	SetColumnClick(value bool)                                           // property ColumnClick Setter
	DragCursor() types.TCursor                                           // property DragCursor Getter
	SetDragCursor(value types.TCursor)                                   // property DragCursor Setter
	DragKind() types.TDragKind                                           // property DragKind Getter
	SetDragKind(value types.TDragKind)                                   // property DragKind Setter
	DragMode() types.TDragMode                                           // property DragMode Getter
	SetDragMode(value types.TDragMode)                                   // property DragMode Setter
	HideSelection() bool                                                 // property HideSelection Getter
	SetHideSelection(value bool)                                         // property HideSelection Setter
	LargeImages() ICustomImageList                                       // property LargeImages Getter
	SetLargeImages(value ICustomImageList)                               // property LargeImages Setter
	LargeImagesWidth() int32                                             // property LargeImagesWidth Getter
	SetLargeImagesWidth(value int32)                                     // property LargeImagesWidth Setter
	OwnerDraw() bool                                                     // property OwnerDraw Getter
	SetOwnerDraw(value bool)                                             // property OwnerDraw Setter
	ParentColor() bool                                                   // property ParentColor Getter
	SetParentColor(value bool)                                           // property ParentColor Setter
	ParentFont() bool                                                    // property ParentFont Getter
	SetParentFont(value bool)                                            // property ParentFont Setter
	ParentShowHint() bool                                                // property ParentShowHint Getter
	SetParentShowHint(value bool)                                        // property ParentShowHint Setter
	ScrollBars() types.TScrollStyle                                      // property ScrollBars Getter
	SetScrollBars(value types.TScrollStyle)                              // property ScrollBars Setter
	ShowColumnHeaders() bool                                             // property ShowColumnHeaders Getter
	SetShowColumnHeaders(value bool)                                     // property ShowColumnHeaders Setter
	SmallImages() ICustomImageList                                       // property SmallImages Getter
	SetSmallImages(value ICustomImageList)                               // property SmallImages Setter
	SmallImagesWidth() int32                                             // property SmallImagesWidth Getter
	SetSmallImagesWidth(value int32)                                     // property SmallImagesWidth Setter
	SortColumn() int32                                                   // property SortColumn Getter
	SetSortColumn(value int32)                                           // property SortColumn Setter
	SortDirection() types.TSortDirection                                 // property SortDirection Getter
	SetSortDirection(value types.TSortDirection)                         // property SortDirection Setter
	SortType() types.TSortType                                           // property SortType Getter
	SetSortType(value types.TSortType)                                   // property SortType Setter
	StateImages() ICustomImageList                                       // property StateImages Getter
	SetStateImages(value ICustomImageList)                               // property StateImages Setter
	StateImagesWidth() int32                                             // property StateImagesWidth Getter
	SetStateImagesWidth(value int32)                                     // property StateImagesWidth Setter
	ToolTips() bool                                                      // property ToolTips Getter
	SetToolTips(value bool)                                              // property ToolTips Setter
	ViewStyle() types.TViewStyle                                         // property ViewStyle Getter
	SetViewStyle(value types.TViewStyle)                                 // property ViewStyle Setter
	SetOnAdvancedCustomDraw(fn TLVAdvancedCustomDrawEvent)               // property event
	SetOnAdvancedCustomDrawItem(fn TLVAdvancedCustomDrawItemEvent)       // property event
	SetOnAdvancedCustomDrawSubItem(fn TLVAdvancedCustomDrawSubItemEvent) // property event
	SetOnChange(fn TLVChangeEvent)                                       // property event
	SetOnChanging(fn TLVChangingEvent)                                   // property event
	SetOnColumnClick(fn TLVColumnClickEvent)                             // property event
	SetOnCompare(fn TLVCompareEvent)                                     // property event
	SetOnContextPopup(fn TContextPopupEvent)                             // property event
	SetOnCreateItemClass(fn TLVCreateItemClassEvent)                     // property event
	SetOnCustomDraw(fn TLVCustomDrawEvent)                               // property event
	SetOnCustomDrawItem(fn TLVCustomDrawItemEvent)                       // property event
	SetOnCustomDrawSubItem(fn TLVCustomDrawSubItemEvent)                 // property event
	SetOnData(fn TLVDataEvent)                                           // property event
	SetOnDataFind(fn TLVDataFindEvent)                                   // property event
	SetOnDataHint(fn TLVDataHintEvent)                                   // property event
	SetOnDataStateChange(fn TLVDataStateChangeEvent)                     // property event
	SetOnDblClick(fn TNotifyEvent)                                       // property event
	SetOnDeletion(fn TLVDeletedEvent)                                    // property event
	SetOnDragDrop(fn TDragDropEvent)                                     // property event
	SetOnDragOver(fn TDragOverEvent)                                     // property event
	SetOnDrawItem(fn TLVDrawItemEvent)                                   // property event
	SetOnEdited(fn TLVEditedEvent)                                       // property event
	SetOnEditing(fn TLVEditingEvent)                                     // property event
	SetOnEndDock(fn TEndDragEvent)                                       // property event
	SetOnEndDrag(fn TEndDragEvent)                                       // property event
	SetOnInsert(fn TLVInsertEvent)                                       // property event
	SetOnItemChecked(fn TLVCheckedItemEvent)                             // property event
	SetOnMouseDown(fn TMouseEvent)                                       // property event
	SetOnMouseEnter(fn TNotifyEvent)                                     // property event
	SetOnMouseLeave(fn TNotifyEvent)                                     // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                   // property event
	SetOnMouseUp(fn TMouseEvent)                                         // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                 // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                       // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                         // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)                             // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)                       // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)                      // property event
	SetOnSelectItem(fn TLVSelectItemEvent)                               // property event
	SetOnStartDock(fn TStartDockEvent)                                   // property event
	SetOnStartDrag(fn TStartDragEvent)                                   // property event
}

type TListView struct {
	TCustomListView
}

func (m *TListView) AllocBy() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TListView) SetAllocBy(value int32) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TListView) AutoSort() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetAutoSort(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) AutoSortIndicator() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetAutoSortIndicator(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) AutoWidthLastColumn() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetAutoWidthLastColumn(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) Columns() IListColumns {
	if !m.IsValid() {
		return nil
	}
	r := listViewAPI().SysCallN(5, 0, m.Instance())
	return AsListColumns(r)
}

func (m *TListView) SetColumns(value IListColumns) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TListView) ColumnClick() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetColumnClick(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(7, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TListView) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TListView) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(8, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TListView) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TListView) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(9, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TListView) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TListView) HideSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetHideSelection(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) LargeImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := listViewAPI().SysCallN(11, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TListView) SetLargeImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TListView) LargeImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TListView) SetLargeImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TListView) OwnerDraw() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetOwnerDraw(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) ScrollBars() types.TScrollStyle {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(17, 0, m.Instance())
	return types.TScrollStyle(r)
}

func (m *TListView) SetScrollBars(value types.TScrollStyle) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TListView) ShowColumnHeaders() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetShowColumnHeaders(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) SmallImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := listViewAPI().SysCallN(19, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TListView) SetSmallImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(19, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TListView) SmallImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TListView) SetSmallImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TListView) SortColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(21, 0, m.Instance())
	return int32(r)
}

func (m *TListView) SetSortColumn(value int32) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TListView) SortDirection() types.TSortDirection {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(22, 0, m.Instance())
	return types.TSortDirection(r)
}

func (m *TListView) SetSortDirection(value types.TSortDirection) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TListView) SortType() types.TSortType {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(23, 0, m.Instance())
	return types.TSortType(r)
}

func (m *TListView) SetSortType(value types.TSortType) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TListView) StateImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := listViewAPI().SysCallN(24, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TListView) SetStateImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(24, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TListView) StateImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(25, 0, m.Instance())
	return int32(r)
}

func (m *TListView) SetStateImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TListView) ToolTips() bool {
	if !m.IsValid() {
		return false
	}
	r := listViewAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TListView) SetToolTips(value bool) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TListView) ViewStyle() types.TViewStyle {
	if !m.IsValid() {
		return 0
	}
	r := listViewAPI().SysCallN(27, 0, m.Instance())
	return types.TViewStyle(r)
}

func (m *TListView) SetViewStyle(value types.TViewStyle) {
	if !m.IsValid() {
		return
	}
	listViewAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TListView) SetOnAdvancedCustomDraw(fn TLVAdvancedCustomDrawEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVAdvancedCustomDrawEvent(fn)
	base.SetEvent(m, 28, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnAdvancedCustomDrawItem(fn TLVAdvancedCustomDrawItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVAdvancedCustomDrawItemEvent(fn)
	base.SetEvent(m, 29, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnAdvancedCustomDrawSubItem(fn TLVAdvancedCustomDrawSubItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVAdvancedCustomDrawSubItemEvent(fn)
	base.SetEvent(m, 30, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnChange(fn TLVChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVChangeEvent(fn)
	base.SetEvent(m, 31, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnChanging(fn TLVChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVChangingEvent(fn)
	base.SetEvent(m, 32, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnColumnClick(fn TLVColumnClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVColumnClickEvent(fn)
	base.SetEvent(m, 33, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnCompare(fn TLVCompareEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVCompareEvent(fn)
	base.SetEvent(m, 34, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 35, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnCreateItemClass(fn TLVCreateItemClassEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVCreateItemClassEvent(fn)
	base.SetEvent(m, 36, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnCustomDraw(fn TLVCustomDrawEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVCustomDrawEvent(fn)
	base.SetEvent(m, 37, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnCustomDrawItem(fn TLVCustomDrawItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVCustomDrawItemEvent(fn)
	base.SetEvent(m, 38, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnCustomDrawSubItem(fn TLVCustomDrawSubItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVCustomDrawSubItemEvent(fn)
	base.SetEvent(m, 39, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnData(fn TLVDataEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVDataEvent(fn)
	base.SetEvent(m, 40, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnDataFind(fn TLVDataFindEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVDataFindEvent(fn)
	base.SetEvent(m, 41, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnDataHint(fn TLVDataHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVDataHintEvent(fn)
	base.SetEvent(m, 42, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnDataStateChange(fn TLVDataStateChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVDataStateChangeEvent(fn)
	base.SetEvent(m, 43, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 44, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnDeletion(fn TLVDeletedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVDeletedEvent(fn)
	base.SetEvent(m, 45, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 46, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 47, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnDrawItem(fn TLVDrawItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVDrawItemEvent(fn)
	base.SetEvent(m, 48, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnEdited(fn TLVEditedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVEditedEvent(fn)
	base.SetEvent(m, 49, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnEditing(fn TLVEditingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVEditingEvent(fn)
	base.SetEvent(m, 50, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 51, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 52, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnInsert(fn TLVInsertEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVInsertEvent(fn)
	base.SetEvent(m, 53, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnItemChecked(fn TLVCheckedItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVCheckedItemEvent(fn)
	base.SetEvent(m, 54, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 55, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 56, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 57, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 58, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 59, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 60, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 61, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 62, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 63, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 64, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 65, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnSelectItem(fn TLVSelectItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTLVSelectItemEvent(fn)
	base.SetEvent(m, 66, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 67, listViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TListView) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 68, listViewAPI(), api.MakeEventDataPtr(cb))
}

// NewListView class constructor
func NewListView(owner IComponent) IListView {
	r := listViewAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsListView(r)
}

func TListViewClass() types.TClass {
	r := listViewAPI().SysCallN(69)
	return types.TClass(r)
}

var (
	listViewOnce   base.Once
	listViewImport *imports.Imports = nil
)

func listViewAPI() *imports.Imports {
	listViewOnce.Do(func() {
		listViewImport = api.NewDefaultImports()
		listViewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TListView_Create", 0), // constructor NewListView
			/* 1 */ imports.NewTable("TListView_AllocBy", 0), // property AllocBy
			/* 2 */ imports.NewTable("TListView_AutoSort", 0), // property AutoSort
			/* 3 */ imports.NewTable("TListView_AutoSortIndicator", 0), // property AutoSortIndicator
			/* 4 */ imports.NewTable("TListView_AutoWidthLastColumn", 0), // property AutoWidthLastColumn
			/* 5 */ imports.NewTable("TListView_Columns", 0), // property Columns
			/* 6 */ imports.NewTable("TListView_ColumnClick", 0), // property ColumnClick
			/* 7 */ imports.NewTable("TListView_DragCursor", 0), // property DragCursor
			/* 8 */ imports.NewTable("TListView_DragKind", 0), // property DragKind
			/* 9 */ imports.NewTable("TListView_DragMode", 0), // property DragMode
			/* 10 */ imports.NewTable("TListView_HideSelection", 0), // property HideSelection
			/* 11 */ imports.NewTable("TListView_LargeImages", 0), // property LargeImages
			/* 12 */ imports.NewTable("TListView_LargeImagesWidth", 0), // property LargeImagesWidth
			/* 13 */ imports.NewTable("TListView_OwnerDraw", 0), // property OwnerDraw
			/* 14 */ imports.NewTable("TListView_ParentColor", 0), // property ParentColor
			/* 15 */ imports.NewTable("TListView_ParentFont", 0), // property ParentFont
			/* 16 */ imports.NewTable("TListView_ParentShowHint", 0), // property ParentShowHint
			/* 17 */ imports.NewTable("TListView_ScrollBars", 0), // property ScrollBars
			/* 18 */ imports.NewTable("TListView_ShowColumnHeaders", 0), // property ShowColumnHeaders
			/* 19 */ imports.NewTable("TListView_SmallImages", 0), // property SmallImages
			/* 20 */ imports.NewTable("TListView_SmallImagesWidth", 0), // property SmallImagesWidth
			/* 21 */ imports.NewTable("TListView_SortColumn", 0), // property SortColumn
			/* 22 */ imports.NewTable("TListView_SortDirection", 0), // property SortDirection
			/* 23 */ imports.NewTable("TListView_SortType", 0), // property SortType
			/* 24 */ imports.NewTable("TListView_StateImages", 0), // property StateImages
			/* 25 */ imports.NewTable("TListView_StateImagesWidth", 0), // property StateImagesWidth
			/* 26 */ imports.NewTable("TListView_ToolTips", 0), // property ToolTips
			/* 27 */ imports.NewTable("TListView_ViewStyle", 0), // property ViewStyle
			/* 28 */ imports.NewTable("TListView_OnAdvancedCustomDraw", 0), // event OnAdvancedCustomDraw
			/* 29 */ imports.NewTable("TListView_OnAdvancedCustomDrawItem", 0), // event OnAdvancedCustomDrawItem
			/* 30 */ imports.NewTable("TListView_OnAdvancedCustomDrawSubItem", 0), // event OnAdvancedCustomDrawSubItem
			/* 31 */ imports.NewTable("TListView_OnChange", 0), // event OnChange
			/* 32 */ imports.NewTable("TListView_OnChanging", 0), // event OnChanging
			/* 33 */ imports.NewTable("TListView_OnColumnClick", 0), // event OnColumnClick
			/* 34 */ imports.NewTable("TListView_OnCompare", 0), // event OnCompare
			/* 35 */ imports.NewTable("TListView_OnContextPopup", 0), // event OnContextPopup
			/* 36 */ imports.NewTable("TListView_OnCreateItemClass", 0), // event OnCreateItemClass
			/* 37 */ imports.NewTable("TListView_OnCustomDraw", 0), // event OnCustomDraw
			/* 38 */ imports.NewTable("TListView_OnCustomDrawItem", 0), // event OnCustomDrawItem
			/* 39 */ imports.NewTable("TListView_OnCustomDrawSubItem", 0), // event OnCustomDrawSubItem
			/* 40 */ imports.NewTable("TListView_OnData", 0), // event OnData
			/* 41 */ imports.NewTable("TListView_OnDataFind", 0), // event OnDataFind
			/* 42 */ imports.NewTable("TListView_OnDataHint", 0), // event OnDataHint
			/* 43 */ imports.NewTable("TListView_OnDataStateChange", 0), // event OnDataStateChange
			/* 44 */ imports.NewTable("TListView_OnDblClick", 0), // event OnDblClick
			/* 45 */ imports.NewTable("TListView_OnDeletion", 0), // event OnDeletion
			/* 46 */ imports.NewTable("TListView_OnDragDrop", 0), // event OnDragDrop
			/* 47 */ imports.NewTable("TListView_OnDragOver", 0), // event OnDragOver
			/* 48 */ imports.NewTable("TListView_OnDrawItem", 0), // event OnDrawItem
			/* 49 */ imports.NewTable("TListView_OnEdited", 0), // event OnEdited
			/* 50 */ imports.NewTable("TListView_OnEditing", 0), // event OnEditing
			/* 51 */ imports.NewTable("TListView_OnEndDock", 0), // event OnEndDock
			/* 52 */ imports.NewTable("TListView_OnEndDrag", 0), // event OnEndDrag
			/* 53 */ imports.NewTable("TListView_OnInsert", 0), // event OnInsert
			/* 54 */ imports.NewTable("TListView_OnItemChecked", 0), // event OnItemChecked
			/* 55 */ imports.NewTable("TListView_OnMouseDown", 0), // event OnMouseDown
			/* 56 */ imports.NewTable("TListView_OnMouseEnter", 0), // event OnMouseEnter
			/* 57 */ imports.NewTable("TListView_OnMouseLeave", 0), // event OnMouseLeave
			/* 58 */ imports.NewTable("TListView_OnMouseMove", 0), // event OnMouseMove
			/* 59 */ imports.NewTable("TListView_OnMouseUp", 0), // event OnMouseUp
			/* 60 */ imports.NewTable("TListView_OnMouseWheel", 0), // event OnMouseWheel
			/* 61 */ imports.NewTable("TListView_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 62 */ imports.NewTable("TListView_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 63 */ imports.NewTable("TListView_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 64 */ imports.NewTable("TListView_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 65 */ imports.NewTable("TListView_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 66 */ imports.NewTable("TListView_OnSelectItem", 0), // event OnSelectItem
			/* 67 */ imports.NewTable("TListView_OnStartDock", 0), // event OnStartDock
			/* 68 */ imports.NewTable("TListView_OnStartDrag", 0), // event OnStartDrag
			/* 69 */ imports.NewTable("TListView_TClass", 0), // function TListViewClass
		}
	})
	return listViewImport
}
