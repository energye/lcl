//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IListView Parent: ICustomListView
type IListView interface {
	ICustomListView
	DeleteSelected()
	AllocBy() int32                                                      // property
	SetAllocBy(AValue int32)                                             // property
	AutoSort() bool                                                      // property
	SetAutoSort(AValue bool)                                             // property
	AutoSortIndicator() bool                                             // property
	SetAutoSortIndicator(AValue bool)                                    // property
	AutoWidthLastColumn() bool                                           // property
	SetAutoWidthLastColumn(AValue bool)                                  // property
	Columns() IListColumns                                               // property
	SetColumns(AValue IListColumns)                                      // property
	ColumnClick() bool                                                   // property
	SetColumnClick(AValue bool)                                          // property
	DragCursor() TCursor                                                 // property
	SetDragCursor(AValue TCursor)                                        // property
	DragKind() TDragKind                                                 // property
	SetDragKind(AValue TDragKind)                                        // property
	DragMode() TDragMode                                                 // property
	SetDragMode(AValue TDragMode)                                        // property
	HideSelection() bool                                                 // property
	SetHideSelection(AValue bool)                                        // property
	LargeImages() ICustomImageList                                       // property
	SetLargeImages(AValue ICustomImageList)                              // property
	LargeImagesWidth() int32                                             // property
	SetLargeImagesWidth(AValue int32)                                    // property
	OwnerDraw() bool                                                     // property
	SetOwnerDraw(AValue bool)                                            // property
	ParentColor() bool                                                   // property
	SetParentColor(AValue bool)                                          // property
	ParentFont() bool                                                    // property
	SetParentFont(AValue bool)                                           // property
	ParentShowHint() bool                                                // property
	SetParentShowHint(AValue bool)                                       // property
	ScrollBars() TScrollStyle                                            // property
	SetScrollBars(AValue TScrollStyle)                                   // property
	ShowColumnHeaders() bool                                             // property
	SetShowColumnHeaders(AValue bool)                                    // property
	SmallImages() ICustomImageList                                       // property
	SetSmallImages(AValue ICustomImageList)                              // property
	SmallImagesWidth() int32                                             // property
	SetSmallImagesWidth(AValue int32)                                    // property
	SortColumn() int32                                                   // property
	SetSortColumn(AValue int32)                                          // property
	SortDirection() TSortDirection                                       // property
	SetSortDirection(AValue TSortDirection)                              // property
	SortType() TSortType                                                 // property
	SetSortType(AValue TSortType)                                        // property
	StateImages() ICustomImageList                                       // property
	SetStateImages(AValue ICustomImageList)                              // property
	StateImagesWidth() int32                                             // property
	SetStateImagesWidth(AValue int32)                                    // property
	ToolTips() bool                                                      // property
	SetToolTips(AValue bool)                                             // property
	ViewStyle() TViewStyle                                               // property
	SetViewStyle(AValue TViewStyle)                                      // property
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

// TListView Parent: TCustomListView
type TListView struct {
	TCustomListView
	advancedCustomDrawPtr        uintptr
	advancedCustomDrawItemPtr    uintptr
	advancedCustomDrawSubItemPtr uintptr
	changePtr                    uintptr
	changingPtr                  uintptr
	columnClickPtr               uintptr
	comparePtr                   uintptr
	contextPopupPtr              uintptr
	createItemClassPtr           uintptr
	customDrawPtr                uintptr
	customDrawItemPtr            uintptr
	customDrawSubItemPtr         uintptr
	dataPtr                      uintptr
	dataFindPtr                  uintptr
	dataHintPtr                  uintptr
	dataStateChangePtr           uintptr
	dblClickPtr                  uintptr
	deletionPtr                  uintptr
	dragDropPtr                  uintptr
	dragOverPtr                  uintptr
	drawItemPtr                  uintptr
	editedPtr                    uintptr
	editingPtr                   uintptr
	endDockPtr                   uintptr
	endDragPtr                   uintptr
	insertPtr                    uintptr
	itemCheckedPtr               uintptr
	mouseDownPtr                 uintptr
	mouseEnterPtr                uintptr
	mouseLeavePtr                uintptr
	mouseMovePtr                 uintptr
	mouseUpPtr                   uintptr
	mouseWheelPtr                uintptr
	mouseWheelDownPtr            uintptr
	mouseWheelUpPtr              uintptr
	mouseWheelHorzPtr            uintptr
	mouseWheelLeftPtr            uintptr
	mouseWheelRightPtr           uintptr
	selectItemPtr                uintptr
	startDockPtr                 uintptr
	startDragPtr                 uintptr
}

func NewListView(AOwner IComponent) IListView {
	r1 := listViewImportAPI().SysCallN(7, GetObjectUintptr(AOwner))
	return AsListView(r1)
}

func (m *TListView) AllocBy() int32 {
	r1 := listViewImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetAllocBy(AValue int32) {
	listViewImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) AutoSort() bool {
	r1 := listViewImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetAutoSort(AValue bool) {
	listViewImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) AutoSortIndicator() bool {
	r1 := listViewImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetAutoSortIndicator(AValue bool) {
	listViewImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) AutoWidthLastColumn() bool {
	r1 := listViewImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetAutoWidthLastColumn(AValue bool) {
	listViewImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) Columns() IListColumns {
	r1 := listViewImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return AsListColumns(r1)
}

func (m *TListView) SetColumns(AValue IListColumns) {
	listViewImportAPI().SysCallN(6, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListView) ColumnClick() bool {
	r1 := listViewImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetColumnClick(AValue bool) {
	listViewImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) DragCursor() TCursor {
	r1 := listViewImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TListView) SetDragCursor(AValue TCursor) {
	listViewImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) DragKind() TDragKind {
	r1 := listViewImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TListView) SetDragKind(AValue TDragKind) {
	listViewImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) DragMode() TDragMode {
	r1 := listViewImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TListView) SetDragMode(AValue TDragMode) {
	listViewImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) HideSelection() bool {
	r1 := listViewImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetHideSelection(AValue bool) {
	listViewImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) LargeImages() ICustomImageList {
	r1 := listViewImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TListView) SetLargeImages(AValue ICustomImageList) {
	listViewImportAPI().SysCallN(12, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListView) LargeImagesWidth() int32 {
	r1 := listViewImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetLargeImagesWidth(AValue int32) {
	listViewImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) OwnerDraw() bool {
	r1 := listViewImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetOwnerDraw(AValue bool) {
	listViewImportAPI().SysCallN(14, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ParentColor() bool {
	r1 := listViewImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetParentColor(AValue bool) {
	listViewImportAPI().SysCallN(15, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ParentFont() bool {
	r1 := listViewImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetParentFont(AValue bool) {
	listViewImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ParentShowHint() bool {
	r1 := listViewImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetParentShowHint(AValue bool) {
	listViewImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ScrollBars() TScrollStyle {
	r1 := listViewImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TListView) SetScrollBars(AValue TScrollStyle) {
	listViewImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) ShowColumnHeaders() bool {
	r1 := listViewImportAPI().SysCallN(60, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetShowColumnHeaders(AValue bool) {
	listViewImportAPI().SysCallN(60, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) SmallImages() ICustomImageList {
	r1 := listViewImportAPI().SysCallN(61, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TListView) SetSmallImages(AValue ICustomImageList) {
	listViewImportAPI().SysCallN(61, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListView) SmallImagesWidth() int32 {
	r1 := listViewImportAPI().SysCallN(62, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetSmallImagesWidth(AValue int32) {
	listViewImportAPI().SysCallN(62, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) SortColumn() int32 {
	r1 := listViewImportAPI().SysCallN(63, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetSortColumn(AValue int32) {
	listViewImportAPI().SysCallN(63, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) SortDirection() TSortDirection {
	r1 := listViewImportAPI().SysCallN(64, 0, m.Instance(), 0)
	return TSortDirection(r1)
}

func (m *TListView) SetSortDirection(AValue TSortDirection) {
	listViewImportAPI().SysCallN(64, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) SortType() TSortType {
	r1 := listViewImportAPI().SysCallN(65, 0, m.Instance(), 0)
	return TSortType(r1)
}

func (m *TListView) SetSortType(AValue TSortType) {
	listViewImportAPI().SysCallN(65, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) StateImages() ICustomImageList {
	r1 := listViewImportAPI().SysCallN(66, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TListView) SetStateImages(AValue ICustomImageList) {
	listViewImportAPI().SysCallN(66, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListView) StateImagesWidth() int32 {
	r1 := listViewImportAPI().SysCallN(67, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetStateImagesWidth(AValue int32) {
	listViewImportAPI().SysCallN(67, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) ToolTips() bool {
	r1 := listViewImportAPI().SysCallN(68, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetToolTips(AValue bool) {
	listViewImportAPI().SysCallN(68, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ViewStyle() TViewStyle {
	r1 := listViewImportAPI().SysCallN(69, 0, m.Instance(), 0)
	return TViewStyle(r1)
}

func (m *TListView) SetViewStyle(AValue TViewStyle) {
	listViewImportAPI().SysCallN(69, 1, m.Instance(), uintptr(AValue))
}

func ListViewClass() TClass {
	ret := listViewImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TListView) SetOnAdvancedCustomDraw(fn TLVAdvancedCustomDrawEvent) {
	if m.advancedCustomDrawPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawPtr)
	}
	m.advancedCustomDrawPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(19, m.Instance(), m.advancedCustomDrawPtr)
}

func (m *TListView) SetOnAdvancedCustomDrawItem(fn TLVAdvancedCustomDrawItemEvent) {
	if m.advancedCustomDrawItemPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawItemPtr)
	}
	m.advancedCustomDrawItemPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(20, m.Instance(), m.advancedCustomDrawItemPtr)
}

func (m *TListView) SetOnAdvancedCustomDrawSubItem(fn TLVAdvancedCustomDrawSubItemEvent) {
	if m.advancedCustomDrawSubItemPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawSubItemPtr)
	}
	m.advancedCustomDrawSubItemPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(21, m.Instance(), m.advancedCustomDrawSubItemPtr)
}

func (m *TListView) SetOnChange(fn TLVChangeEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(22, m.Instance(), m.changePtr)
}

func (m *TListView) SetOnChanging(fn TLVChangingEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(23, m.Instance(), m.changingPtr)
}

func (m *TListView) SetOnColumnClick(fn TLVColumnClickEvent) {
	if m.columnClickPtr != 0 {
		RemoveEventElement(m.columnClickPtr)
	}
	m.columnClickPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(24, m.Instance(), m.columnClickPtr)
}

func (m *TListView) SetOnCompare(fn TLVCompareEvent) {
	if m.comparePtr != 0 {
		RemoveEventElement(m.comparePtr)
	}
	m.comparePtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(25, m.Instance(), m.comparePtr)
}

func (m *TListView) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(26, m.Instance(), m.contextPopupPtr)
}

func (m *TListView) SetOnCreateItemClass(fn TLVCreateItemClassEvent) {
	if m.createItemClassPtr != 0 {
		RemoveEventElement(m.createItemClassPtr)
	}
	m.createItemClassPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(27, m.Instance(), m.createItemClassPtr)
}

func (m *TListView) SetOnCustomDraw(fn TLVCustomDrawEvent) {
	if m.customDrawPtr != 0 {
		RemoveEventElement(m.customDrawPtr)
	}
	m.customDrawPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(28, m.Instance(), m.customDrawPtr)
}

func (m *TListView) SetOnCustomDrawItem(fn TLVCustomDrawItemEvent) {
	if m.customDrawItemPtr != 0 {
		RemoveEventElement(m.customDrawItemPtr)
	}
	m.customDrawItemPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(29, m.Instance(), m.customDrawItemPtr)
}

func (m *TListView) SetOnCustomDrawSubItem(fn TLVCustomDrawSubItemEvent) {
	if m.customDrawSubItemPtr != 0 {
		RemoveEventElement(m.customDrawSubItemPtr)
	}
	m.customDrawSubItemPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(30, m.Instance(), m.customDrawSubItemPtr)
}

func (m *TListView) SetOnData(fn TLVDataEvent) {
	if m.dataPtr != 0 {
		RemoveEventElement(m.dataPtr)
	}
	m.dataPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(31, m.Instance(), m.dataPtr)
}

func (m *TListView) SetOnDataFind(fn TLVDataFindEvent) {
	if m.dataFindPtr != 0 {
		RemoveEventElement(m.dataFindPtr)
	}
	m.dataFindPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(32, m.Instance(), m.dataFindPtr)
}

func (m *TListView) SetOnDataHint(fn TLVDataHintEvent) {
	if m.dataHintPtr != 0 {
		RemoveEventElement(m.dataHintPtr)
	}
	m.dataHintPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(33, m.Instance(), m.dataHintPtr)
}

func (m *TListView) SetOnDataStateChange(fn TLVDataStateChangeEvent) {
	if m.dataStateChangePtr != 0 {
		RemoveEventElement(m.dataStateChangePtr)
	}
	m.dataStateChangePtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(34, m.Instance(), m.dataStateChangePtr)
}

func (m *TListView) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(35, m.Instance(), m.dblClickPtr)
}

func (m *TListView) SetOnDeletion(fn TLVDeletedEvent) {
	if m.deletionPtr != 0 {
		RemoveEventElement(m.deletionPtr)
	}
	m.deletionPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(36, m.Instance(), m.deletionPtr)
}

func (m *TListView) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(37, m.Instance(), m.dragDropPtr)
}

func (m *TListView) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(38, m.Instance(), m.dragOverPtr)
}

func (m *TListView) SetOnDrawItem(fn TLVDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(39, m.Instance(), m.drawItemPtr)
}

func (m *TListView) SetOnEdited(fn TLVEditedEvent) {
	if m.editedPtr != 0 {
		RemoveEventElement(m.editedPtr)
	}
	m.editedPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(40, m.Instance(), m.editedPtr)
}

func (m *TListView) SetOnEditing(fn TLVEditingEvent) {
	if m.editingPtr != 0 {
		RemoveEventElement(m.editingPtr)
	}
	m.editingPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(41, m.Instance(), m.editingPtr)
}

func (m *TListView) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(42, m.Instance(), m.endDockPtr)
}

func (m *TListView) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(43, m.Instance(), m.endDragPtr)
}

func (m *TListView) SetOnInsert(fn TLVInsertEvent) {
	if m.insertPtr != 0 {
		RemoveEventElement(m.insertPtr)
	}
	m.insertPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(44, m.Instance(), m.insertPtr)
}

func (m *TListView) SetOnItemChecked(fn TLVCheckedItemEvent) {
	if m.itemCheckedPtr != 0 {
		RemoveEventElement(m.itemCheckedPtr)
	}
	m.itemCheckedPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(45, m.Instance(), m.itemCheckedPtr)
}

func (m *TListView) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(46, m.Instance(), m.mouseDownPtr)
}

func (m *TListView) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(47, m.Instance(), m.mouseEnterPtr)
}

func (m *TListView) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(48, m.Instance(), m.mouseLeavePtr)
}

func (m *TListView) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(49, m.Instance(), m.mouseMovePtr)
}

func (m *TListView) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(50, m.Instance(), m.mouseUpPtr)
}

func (m *TListView) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(51, m.Instance(), m.mouseWheelPtr)
}

func (m *TListView) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(52, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TListView) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(56, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TListView) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(53, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TListView) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(54, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TListView) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(55, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TListView) SetOnSelectItem(fn TLVSelectItemEvent) {
	if m.selectItemPtr != 0 {
		RemoveEventElement(m.selectItemPtr)
	}
	m.selectItemPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(57, m.Instance(), m.selectItemPtr)
}

func (m *TListView) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(58, m.Instance(), m.startDockPtr)
}

func (m *TListView) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	listViewImportAPI().SysCallN(59, m.Instance(), m.startDragPtr)
}

var (
	listViewImport       *imports.Imports = nil
	listViewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListView_AllocBy", 0),
		/*1*/ imports.NewTable("ListView_AutoSort", 0),
		/*2*/ imports.NewTable("ListView_AutoSortIndicator", 0),
		/*3*/ imports.NewTable("ListView_AutoWidthLastColumn", 0),
		/*4*/ imports.NewTable("ListView_Class", 0),
		/*5*/ imports.NewTable("ListView_ColumnClick", 0),
		/*6*/ imports.NewTable("ListView_Columns", 0),
		/*7*/ imports.NewTable("ListView_Create", 0),
		/*8*/ imports.NewTable("ListView_DragCursor", 0),
		/*9*/ imports.NewTable("ListView_DragKind", 0),
		/*10*/ imports.NewTable("ListView_DragMode", 0),
		/*11*/ imports.NewTable("ListView_HideSelection", 0),
		/*12*/ imports.NewTable("ListView_LargeImages", 0),
		/*13*/ imports.NewTable("ListView_LargeImagesWidth", 0),
		/*14*/ imports.NewTable("ListView_OwnerDraw", 0),
		/*15*/ imports.NewTable("ListView_ParentColor", 0),
		/*16*/ imports.NewTable("ListView_ParentFont", 0),
		/*17*/ imports.NewTable("ListView_ParentShowHint", 0),
		/*18*/ imports.NewTable("ListView_ScrollBars", 0),
		/*19*/ imports.NewTable("ListView_SetOnAdvancedCustomDraw", 0),
		/*20*/ imports.NewTable("ListView_SetOnAdvancedCustomDrawItem", 0),
		/*21*/ imports.NewTable("ListView_SetOnAdvancedCustomDrawSubItem", 0),
		/*22*/ imports.NewTable("ListView_SetOnChange", 0),
		/*23*/ imports.NewTable("ListView_SetOnChanging", 0),
		/*24*/ imports.NewTable("ListView_SetOnColumnClick", 0),
		/*25*/ imports.NewTable("ListView_SetOnCompare", 0),
		/*26*/ imports.NewTable("ListView_SetOnContextPopup", 0),
		/*27*/ imports.NewTable("ListView_SetOnCreateItemClass", 0),
		/*28*/ imports.NewTable("ListView_SetOnCustomDraw", 0),
		/*29*/ imports.NewTable("ListView_SetOnCustomDrawItem", 0),
		/*30*/ imports.NewTable("ListView_SetOnCustomDrawSubItem", 0),
		/*31*/ imports.NewTable("ListView_SetOnData", 0),
		/*32*/ imports.NewTable("ListView_SetOnDataFind", 0),
		/*33*/ imports.NewTable("ListView_SetOnDataHint", 0),
		/*34*/ imports.NewTable("ListView_SetOnDataStateChange", 0),
		/*35*/ imports.NewTable("ListView_SetOnDblClick", 0),
		/*36*/ imports.NewTable("ListView_SetOnDeletion", 0),
		/*37*/ imports.NewTable("ListView_SetOnDragDrop", 0),
		/*38*/ imports.NewTable("ListView_SetOnDragOver", 0),
		/*39*/ imports.NewTable("ListView_SetOnDrawItem", 0),
		/*40*/ imports.NewTable("ListView_SetOnEdited", 0),
		/*41*/ imports.NewTable("ListView_SetOnEditing", 0),
		/*42*/ imports.NewTable("ListView_SetOnEndDock", 0),
		/*43*/ imports.NewTable("ListView_SetOnEndDrag", 0),
		/*44*/ imports.NewTable("ListView_SetOnInsert", 0),
		/*45*/ imports.NewTable("ListView_SetOnItemChecked", 0),
		/*46*/ imports.NewTable("ListView_SetOnMouseDown", 0),
		/*47*/ imports.NewTable("ListView_SetOnMouseEnter", 0),
		/*48*/ imports.NewTable("ListView_SetOnMouseLeave", 0),
		/*49*/ imports.NewTable("ListView_SetOnMouseMove", 0),
		/*50*/ imports.NewTable("ListView_SetOnMouseUp", 0),
		/*51*/ imports.NewTable("ListView_SetOnMouseWheel", 0),
		/*52*/ imports.NewTable("ListView_SetOnMouseWheelDown", 0),
		/*53*/ imports.NewTable("ListView_SetOnMouseWheelHorz", 0),
		/*54*/ imports.NewTable("ListView_SetOnMouseWheelLeft", 0),
		/*55*/ imports.NewTable("ListView_SetOnMouseWheelRight", 0),
		/*56*/ imports.NewTable("ListView_SetOnMouseWheelUp", 0),
		/*57*/ imports.NewTable("ListView_SetOnSelectItem", 0),
		/*58*/ imports.NewTable("ListView_SetOnStartDock", 0),
		/*59*/ imports.NewTable("ListView_SetOnStartDrag", 0),
		/*60*/ imports.NewTable("ListView_ShowColumnHeaders", 0),
		/*61*/ imports.NewTable("ListView_SmallImages", 0),
		/*62*/ imports.NewTable("ListView_SmallImagesWidth", 0),
		/*63*/ imports.NewTable("ListView_SortColumn", 0),
		/*64*/ imports.NewTable("ListView_SortDirection", 0),
		/*65*/ imports.NewTable("ListView_SortType", 0),
		/*66*/ imports.NewTable("ListView_StateImages", 0),
		/*67*/ imports.NewTable("ListView_StateImagesWidth", 0),
		/*68*/ imports.NewTable("ListView_ToolTips", 0),
		/*69*/ imports.NewTable("ListView_ViewStyle", 0),
	}
)

func listViewImportAPI() *imports.Imports {
	if listViewImport == nil {
		listViewImport = NewDefaultImports()
		listViewImport.SetImportTable(listViewImportTables)
		listViewImportTables = nil
	}
	return listViewImport
}
