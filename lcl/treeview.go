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

// ITreeView Parent: ICustomTreeView
type ITreeView interface {
	ICustomTreeView
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	AutoExpand() bool                                              // property
	SetAutoExpand(AValue bool)                                     // property
	DisabledFontColor() TColor                                     // property
	SetDisabledFontColor(AValue TColor)                            // property
	DragKind() TDragKind                                           // property
	SetDragKind(AValue TDragKind)                                  // property
	DragCursor() TCursor                                           // property
	SetDragCursor(AValue TCursor)                                  // property
	DragMode() TDragMode                                           // property
	SetDragMode(AValue TDragMode)                                  // property
	HideSelection() bool                                           // property
	SetHideSelection(AValue bool)                                  // property
	HotTrack() bool                                                // property
	SetHotTrack(AValue bool)                                       // property
	HotTrackColor() TColor                                         // property
	SetHotTrackColor(AValue TColor)                                // property
	Indent() int32                                                 // property
	SetIndent(AValue int32)                                        // property
	MultiSelect() bool                                             // property
	SetMultiSelect(AValue bool)                                    // property
	ParentColor() bool                                             // property
	SetParentColor(AValue bool)                                    // property
	ParentFont() bool                                              // property
	SetParentFont(AValue bool)                                     // property
	ParentShowHint() bool                                          // property
	SetParentShowHint(AValue bool)                                 // property
	ReadOnly() bool                                                // property
	SetReadOnly(AValue bool)                                       // property
	RightClickSelect() bool                                        // property
	SetRightClickSelect(AValue bool)                               // property
	RowSelect() bool                                               // property
	SetRowSelect(AValue bool)                                      // property
	ShowButtons() bool                                             // property
	SetShowButtons(AValue bool)                                    // property
	ShowLines() bool                                               // property
	SetShowLines(AValue bool)                                      // property
	ShowRoot() bool                                                // property
	SetShowRoot(AValue bool)                                       // property
	ShowSeparators() bool                                          // property
	SetShowSeparators(AValue bool)                                 // property
	SortType() TSortType                                           // property
	SetSortType(AValue TSortType)                                  // property
	ToolTips() bool                                                // property
	SetToolTips(AValue bool)                                       // property
	SetOnAddition(fn TTVExpandedEvent)                             // property event
	SetOnAdvancedCustomDraw(fn TTVAdvancedCustomDrawEvent)         // property event
	SetOnAdvancedCustomDrawItem(fn TTVAdvancedCustomDrawItemEvent) // property event
	SetOnChange(fn TTVChangedEvent)                                // property event
	SetOnChanging(fn TTVChangingEvent)                             // property event
	SetOnCollapsed(fn TTVExpandedEvent)                            // property event
	SetOnCollapsing(fn TTVCollapsingEvent)                         // property event
	SetOnCompare(fn TTVCompareEvent)                               // property event
	SetOnContextPopup(fn TContextPopupEvent)                       // property event
	SetOnCreateNodeClass(fn TTVCreateNodeClassEvent)               // property event
	SetOnCustomCreateItem(fn TTVCustomCreateNodeEvent)             // property event
	SetOnCustomDraw(fn TTVCustomDrawEvent)                         // property event
	SetOnCustomDrawItem(fn TTVCustomDrawItemEvent)                 // property event
	SetOnCustomDrawArrow(fn TTVCustomDrawArrowEvent)               // property event
	SetOnDblClick(fn TNotifyEvent)                                 // property event
	SetOnDeletion(fn TTVExpandedEvent)                             // property event
	SetOnDragDrop(fn TDragDropEvent)                               // property event
	SetOnDragOver(fn TDragOverEvent)                               // property event
	SetOnEdited(fn TTVEditedEvent)                                 // property event
	SetOnEditing(fn TTVEditingEvent)                               // property event
	SetOnEditingEnd(fn TTVEditingEndEvent)                         // property event
	SetOnEndDrag(fn TEndDragEvent)                                 // property event
	SetOnExpanded(fn TTVExpandedEvent)                             // property event
	SetOnExpanding(fn TTVExpandingEvent)                           // property event
	SetOnGetImageIndex(fn TTVExpandedEvent)                        // property event
	SetOnGetSelectedIndex(fn TTVExpandedEvent)                     // property event
	SetOnHasChildren(fn TTVHasChildrenEvent)                       // property event
	SetOnMouseDown(fn TMouseEvent)                                 // property event
	SetOnMouseEnter(fn TNotifyEvent)                               // property event
	SetOnMouseLeave(fn TNotifyEvent)                               // property event
	SetOnMouseMove(fn TMouseMoveEvent)                             // property event
	SetOnMouseUp(fn TMouseEvent)                                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                 // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                   // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)                       // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)                 // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)                // property event
	SetOnNodeChanged(fn TTVNodeChangedEvent)                       // property event
	SetOnSelectionChanged(fn TNotifyEvent)                         // property event
	SetOnStartDrag(fn TStartDragEvent)                             // property event
}

// TTreeView Parent: TCustomTreeView
type TTreeView struct {
	TCustomTreeView
	additionPtr               uintptr
	advancedCustomDrawPtr     uintptr
	advancedCustomDrawItemPtr uintptr
	changePtr                 uintptr
	changingPtr               uintptr
	collapsedPtr              uintptr
	collapsingPtr             uintptr
	comparePtr                uintptr
	contextPopupPtr           uintptr
	createNodeClassPtr        uintptr
	customCreateItemPtr       uintptr
	customDrawPtr             uintptr
	customDrawItemPtr         uintptr
	customDrawArrowPtr        uintptr
	dblClickPtr               uintptr
	deletionPtr               uintptr
	dragDropPtr               uintptr
	dragOverPtr               uintptr
	editedPtr                 uintptr
	editingPtr                uintptr
	editingEndPtr             uintptr
	endDragPtr                uintptr
	expandedPtr               uintptr
	expandingPtr              uintptr
	getImageIndexPtr          uintptr
	getSelectedIndexPtr       uintptr
	hasChildrenPtr            uintptr
	mouseDownPtr              uintptr
	mouseEnterPtr             uintptr
	mouseLeavePtr             uintptr
	mouseMovePtr              uintptr
	mouseUpPtr                uintptr
	mouseWheelPtr             uintptr
	mouseWheelDownPtr         uintptr
	mouseWheelUpPtr           uintptr
	mouseWheelHorzPtr         uintptr
	mouseWheelLeftPtr         uintptr
	mouseWheelRightPtr        uintptr
	nodeChangedPtr            uintptr
	selectionChangedPtr       uintptr
	startDragPtr              uintptr
}

func NewTreeView(AnOwner IComponent) ITreeView {
	r1 := reeViewImportAPI().SysCallN(2, GetObjectUintptr(AnOwner))
	return AsTreeView(r1)
}

func (m *TTreeView) AutoExpand() bool {
	r1 := reeViewImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetAutoExpand(AValue bool) {
	reeViewImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) DisabledFontColor() TColor {
	r1 := reeViewImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TTreeView) SetDisabledFontColor(AValue TColor) {
	reeViewImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) DragKind() TDragKind {
	r1 := reeViewImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TTreeView) SetDragKind(AValue TDragKind) {
	reeViewImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) DragCursor() TCursor {
	r1 := reeViewImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TTreeView) SetDragCursor(AValue TCursor) {
	reeViewImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) DragMode() TDragMode {
	r1 := reeViewImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TTreeView) SetDragMode(AValue TDragMode) {
	reeViewImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) HideSelection() bool {
	r1 := reeViewImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetHideSelection(AValue bool) {
	reeViewImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) HotTrack() bool {
	r1 := reeViewImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetHotTrack(AValue bool) {
	reeViewImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) HotTrackColor() TColor {
	r1 := reeViewImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TTreeView) SetHotTrackColor(AValue TColor) {
	reeViewImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) Indent() int32 {
	r1 := reeViewImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeView) SetIndent(AValue int32) {
	reeViewImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) MultiSelect() bool {
	r1 := reeViewImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetMultiSelect(AValue bool) {
	reeViewImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ParentColor() bool {
	r1 := reeViewImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetParentColor(AValue bool) {
	reeViewImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ParentFont() bool {
	r1 := reeViewImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetParentFont(AValue bool) {
	reeViewImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ParentShowHint() bool {
	r1 := reeViewImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetParentShowHint(AValue bool) {
	reeViewImportAPI().SysCallN(14, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ReadOnly() bool {
	r1 := reeViewImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetReadOnly(AValue bool) {
	reeViewImportAPI().SysCallN(15, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) RightClickSelect() bool {
	r1 := reeViewImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetRightClickSelect(AValue bool) {
	reeViewImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) RowSelect() bool {
	r1 := reeViewImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetRowSelect(AValue bool) {
	reeViewImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ShowButtons() bool {
	r1 := reeViewImportAPI().SysCallN(59, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetShowButtons(AValue bool) {
	reeViewImportAPI().SysCallN(59, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ShowLines() bool {
	r1 := reeViewImportAPI().SysCallN(60, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetShowLines(AValue bool) {
	reeViewImportAPI().SysCallN(60, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ShowRoot() bool {
	r1 := reeViewImportAPI().SysCallN(61, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetShowRoot(AValue bool) {
	reeViewImportAPI().SysCallN(61, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ShowSeparators() bool {
	r1 := reeViewImportAPI().SysCallN(62, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetShowSeparators(AValue bool) {
	reeViewImportAPI().SysCallN(62, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) SortType() TSortType {
	r1 := reeViewImportAPI().SysCallN(63, 0, m.Instance(), 0)
	return TSortType(r1)
}

func (m *TTreeView) SetSortType(AValue TSortType) {
	reeViewImportAPI().SysCallN(63, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) ToolTips() bool {
	r1 := reeViewImportAPI().SysCallN(64, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetToolTips(AValue bool) {
	reeViewImportAPI().SysCallN(64, 1, m.Instance(), PascalBool(AValue))
}

func TreeViewClass() TClass {
	ret := reeViewImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TTreeView) SetOnAddition(fn TTVExpandedEvent) {
	if m.additionPtr != 0 {
		RemoveEventElement(m.additionPtr)
	}
	m.additionPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(18, m.Instance(), m.additionPtr)
}

func (m *TTreeView) SetOnAdvancedCustomDraw(fn TTVAdvancedCustomDrawEvent) {
	if m.advancedCustomDrawPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawPtr)
	}
	m.advancedCustomDrawPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(19, m.Instance(), m.advancedCustomDrawPtr)
}

func (m *TTreeView) SetOnAdvancedCustomDrawItem(fn TTVAdvancedCustomDrawItemEvent) {
	if m.advancedCustomDrawItemPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawItemPtr)
	}
	m.advancedCustomDrawItemPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(20, m.Instance(), m.advancedCustomDrawItemPtr)
}

func (m *TTreeView) SetOnChange(fn TTVChangedEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(21, m.Instance(), m.changePtr)
}

func (m *TTreeView) SetOnChanging(fn TTVChangingEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(22, m.Instance(), m.changingPtr)
}

func (m *TTreeView) SetOnCollapsed(fn TTVExpandedEvent) {
	if m.collapsedPtr != 0 {
		RemoveEventElement(m.collapsedPtr)
	}
	m.collapsedPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(23, m.Instance(), m.collapsedPtr)
}

func (m *TTreeView) SetOnCollapsing(fn TTVCollapsingEvent) {
	if m.collapsingPtr != 0 {
		RemoveEventElement(m.collapsingPtr)
	}
	m.collapsingPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(24, m.Instance(), m.collapsingPtr)
}

func (m *TTreeView) SetOnCompare(fn TTVCompareEvent) {
	if m.comparePtr != 0 {
		RemoveEventElement(m.comparePtr)
	}
	m.comparePtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(25, m.Instance(), m.comparePtr)
}

func (m *TTreeView) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(26, m.Instance(), m.contextPopupPtr)
}

func (m *TTreeView) SetOnCreateNodeClass(fn TTVCreateNodeClassEvent) {
	if m.createNodeClassPtr != 0 {
		RemoveEventElement(m.createNodeClassPtr)
	}
	m.createNodeClassPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(27, m.Instance(), m.createNodeClassPtr)
}

func (m *TTreeView) SetOnCustomCreateItem(fn TTVCustomCreateNodeEvent) {
	if m.customCreateItemPtr != 0 {
		RemoveEventElement(m.customCreateItemPtr)
	}
	m.customCreateItemPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(28, m.Instance(), m.customCreateItemPtr)
}

func (m *TTreeView) SetOnCustomDraw(fn TTVCustomDrawEvent) {
	if m.customDrawPtr != 0 {
		RemoveEventElement(m.customDrawPtr)
	}
	m.customDrawPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(29, m.Instance(), m.customDrawPtr)
}

func (m *TTreeView) SetOnCustomDrawItem(fn TTVCustomDrawItemEvent) {
	if m.customDrawItemPtr != 0 {
		RemoveEventElement(m.customDrawItemPtr)
	}
	m.customDrawItemPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(31, m.Instance(), m.customDrawItemPtr)
}

func (m *TTreeView) SetOnCustomDrawArrow(fn TTVCustomDrawArrowEvent) {
	if m.customDrawArrowPtr != 0 {
		RemoveEventElement(m.customDrawArrowPtr)
	}
	m.customDrawArrowPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(30, m.Instance(), m.customDrawArrowPtr)
}

func (m *TTreeView) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(32, m.Instance(), m.dblClickPtr)
}

func (m *TTreeView) SetOnDeletion(fn TTVExpandedEvent) {
	if m.deletionPtr != 0 {
		RemoveEventElement(m.deletionPtr)
	}
	m.deletionPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(33, m.Instance(), m.deletionPtr)
}

func (m *TTreeView) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(34, m.Instance(), m.dragDropPtr)
}

func (m *TTreeView) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(35, m.Instance(), m.dragOverPtr)
}

func (m *TTreeView) SetOnEdited(fn TTVEditedEvent) {
	if m.editedPtr != 0 {
		RemoveEventElement(m.editedPtr)
	}
	m.editedPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(36, m.Instance(), m.editedPtr)
}

func (m *TTreeView) SetOnEditing(fn TTVEditingEvent) {
	if m.editingPtr != 0 {
		RemoveEventElement(m.editingPtr)
	}
	m.editingPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(37, m.Instance(), m.editingPtr)
}

func (m *TTreeView) SetOnEditingEnd(fn TTVEditingEndEvent) {
	if m.editingEndPtr != 0 {
		RemoveEventElement(m.editingEndPtr)
	}
	m.editingEndPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(38, m.Instance(), m.editingEndPtr)
}

func (m *TTreeView) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(39, m.Instance(), m.endDragPtr)
}

func (m *TTreeView) SetOnExpanded(fn TTVExpandedEvent) {
	if m.expandedPtr != 0 {
		RemoveEventElement(m.expandedPtr)
	}
	m.expandedPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(40, m.Instance(), m.expandedPtr)
}

func (m *TTreeView) SetOnExpanding(fn TTVExpandingEvent) {
	if m.expandingPtr != 0 {
		RemoveEventElement(m.expandingPtr)
	}
	m.expandingPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(41, m.Instance(), m.expandingPtr)
}

func (m *TTreeView) SetOnGetImageIndex(fn TTVExpandedEvent) {
	if m.getImageIndexPtr != 0 {
		RemoveEventElement(m.getImageIndexPtr)
	}
	m.getImageIndexPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(42, m.Instance(), m.getImageIndexPtr)
}

func (m *TTreeView) SetOnGetSelectedIndex(fn TTVExpandedEvent) {
	if m.getSelectedIndexPtr != 0 {
		RemoveEventElement(m.getSelectedIndexPtr)
	}
	m.getSelectedIndexPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(43, m.Instance(), m.getSelectedIndexPtr)
}

func (m *TTreeView) SetOnHasChildren(fn TTVHasChildrenEvent) {
	if m.hasChildrenPtr != 0 {
		RemoveEventElement(m.hasChildrenPtr)
	}
	m.hasChildrenPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(44, m.Instance(), m.hasChildrenPtr)
}

func (m *TTreeView) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(45, m.Instance(), m.mouseDownPtr)
}

func (m *TTreeView) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(46, m.Instance(), m.mouseEnterPtr)
}

func (m *TTreeView) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(47, m.Instance(), m.mouseLeavePtr)
}

func (m *TTreeView) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(48, m.Instance(), m.mouseMovePtr)
}

func (m *TTreeView) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(49, m.Instance(), m.mouseUpPtr)
}

func (m *TTreeView) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(50, m.Instance(), m.mouseWheelPtr)
}

func (m *TTreeView) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(51, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TTreeView) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(55, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TTreeView) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(52, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TTreeView) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(53, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TTreeView) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(54, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TTreeView) SetOnNodeChanged(fn TTVNodeChangedEvent) {
	if m.nodeChangedPtr != 0 {
		RemoveEventElement(m.nodeChangedPtr)
	}
	m.nodeChangedPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(56, m.Instance(), m.nodeChangedPtr)
}

func (m *TTreeView) SetOnSelectionChanged(fn TNotifyEvent) {
	if m.selectionChangedPtr != 0 {
		RemoveEventElement(m.selectionChangedPtr)
	}
	m.selectionChangedPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(57, m.Instance(), m.selectionChangedPtr)
}

func (m *TTreeView) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	reeViewImportAPI().SysCallN(58, m.Instance(), m.startDragPtr)
}

var (
	reeViewImport       *imports.Imports = nil
	reeViewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TreeView_AutoExpand", 0),
		/*1*/ imports.NewTable("TreeView_Class", 0),
		/*2*/ imports.NewTable("TreeView_Create", 0),
		/*3*/ imports.NewTable("TreeView_DisabledFontColor", 0),
		/*4*/ imports.NewTable("TreeView_DragCursor", 0),
		/*5*/ imports.NewTable("TreeView_DragKind", 0),
		/*6*/ imports.NewTable("TreeView_DragMode", 0),
		/*7*/ imports.NewTable("TreeView_HideSelection", 0),
		/*8*/ imports.NewTable("TreeView_HotTrack", 0),
		/*9*/ imports.NewTable("TreeView_HotTrackColor", 0),
		/*10*/ imports.NewTable("TreeView_Indent", 0),
		/*11*/ imports.NewTable("TreeView_MultiSelect", 0),
		/*12*/ imports.NewTable("TreeView_ParentColor", 0),
		/*13*/ imports.NewTable("TreeView_ParentFont", 0),
		/*14*/ imports.NewTable("TreeView_ParentShowHint", 0),
		/*15*/ imports.NewTable("TreeView_ReadOnly", 0),
		/*16*/ imports.NewTable("TreeView_RightClickSelect", 0),
		/*17*/ imports.NewTable("TreeView_RowSelect", 0),
		/*18*/ imports.NewTable("TreeView_SetOnAddition", 0),
		/*19*/ imports.NewTable("TreeView_SetOnAdvancedCustomDraw", 0),
		/*20*/ imports.NewTable("TreeView_SetOnAdvancedCustomDrawItem", 0),
		/*21*/ imports.NewTable("TreeView_SetOnChange", 0),
		/*22*/ imports.NewTable("TreeView_SetOnChanging", 0),
		/*23*/ imports.NewTable("TreeView_SetOnCollapsed", 0),
		/*24*/ imports.NewTable("TreeView_SetOnCollapsing", 0),
		/*25*/ imports.NewTable("TreeView_SetOnCompare", 0),
		/*26*/ imports.NewTable("TreeView_SetOnContextPopup", 0),
		/*27*/ imports.NewTable("TreeView_SetOnCreateNodeClass", 0),
		/*28*/ imports.NewTable("TreeView_SetOnCustomCreateItem", 0),
		/*29*/ imports.NewTable("TreeView_SetOnCustomDraw", 0),
		/*30*/ imports.NewTable("TreeView_SetOnCustomDrawArrow", 0),
		/*31*/ imports.NewTable("TreeView_SetOnCustomDrawItem", 0),
		/*32*/ imports.NewTable("TreeView_SetOnDblClick", 0),
		/*33*/ imports.NewTable("TreeView_SetOnDeletion", 0),
		/*34*/ imports.NewTable("TreeView_SetOnDragDrop", 0),
		/*35*/ imports.NewTable("TreeView_SetOnDragOver", 0),
		/*36*/ imports.NewTable("TreeView_SetOnEdited", 0),
		/*37*/ imports.NewTable("TreeView_SetOnEditing", 0),
		/*38*/ imports.NewTable("TreeView_SetOnEditingEnd", 0),
		/*39*/ imports.NewTable("TreeView_SetOnEndDrag", 0),
		/*40*/ imports.NewTable("TreeView_SetOnExpanded", 0),
		/*41*/ imports.NewTable("TreeView_SetOnExpanding", 0),
		/*42*/ imports.NewTable("TreeView_SetOnGetImageIndex", 0),
		/*43*/ imports.NewTable("TreeView_SetOnGetSelectedIndex", 0),
		/*44*/ imports.NewTable("TreeView_SetOnHasChildren", 0),
		/*45*/ imports.NewTable("TreeView_SetOnMouseDown", 0),
		/*46*/ imports.NewTable("TreeView_SetOnMouseEnter", 0),
		/*47*/ imports.NewTable("TreeView_SetOnMouseLeave", 0),
		/*48*/ imports.NewTable("TreeView_SetOnMouseMove", 0),
		/*49*/ imports.NewTable("TreeView_SetOnMouseUp", 0),
		/*50*/ imports.NewTable("TreeView_SetOnMouseWheel", 0),
		/*51*/ imports.NewTable("TreeView_SetOnMouseWheelDown", 0),
		/*52*/ imports.NewTable("TreeView_SetOnMouseWheelHorz", 0),
		/*53*/ imports.NewTable("TreeView_SetOnMouseWheelLeft", 0),
		/*54*/ imports.NewTable("TreeView_SetOnMouseWheelRight", 0),
		/*55*/ imports.NewTable("TreeView_SetOnMouseWheelUp", 0),
		/*56*/ imports.NewTable("TreeView_SetOnNodeChanged", 0),
		/*57*/ imports.NewTable("TreeView_SetOnSelectionChanged", 0),
		/*58*/ imports.NewTable("TreeView_SetOnStartDrag", 0),
		/*59*/ imports.NewTable("TreeView_ShowButtons", 0),
		/*60*/ imports.NewTable("TreeView_ShowLines", 0),
		/*61*/ imports.NewTable("TreeView_ShowRoot", 0),
		/*62*/ imports.NewTable("TreeView_ShowSeparators", 0),
		/*63*/ imports.NewTable("TreeView_SortType", 0),
		/*64*/ imports.NewTable("TreeView_ToolTips", 0),
	}
)

func reeViewImportAPI() *imports.Imports {
	if reeViewImport == nil {
		reeViewImport = NewDefaultImports()
		reeViewImport.SetImportTable(reeViewImportTables)
		reeViewImportTables = nil
	}
	return reeViewImport
}
