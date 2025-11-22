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

// ITreeView Parent: ICustomTreeView
type ITreeView interface {
	ICustomTreeView
	PathDelimiter() string                                         // property PathDelimiter Getter
	SetPathDelimiter(value string)                                 // property PathDelimiter Setter
	ScrolledLeft() int32                                           // property ScrolledLeft Getter
	SetScrolledLeft(value int32)                                   // property ScrolledLeft Setter
	ScrolledTop() int32                                            // property ScrolledTop Getter
	SetScrolledTop(value int32)                                    // property ScrolledTop Setter
	AutoExpand() bool                                              // property AutoExpand Getter
	SetAutoExpand(value bool)                                      // property AutoExpand Setter
	DisabledFontColor() types.TColor                               // property DisabledFontColor Getter
	SetDisabledFontColor(value types.TColor)                       // property DisabledFontColor Setter
	DragKind() types.TDragKind                                     // property DragKind Getter
	SetDragKind(value types.TDragKind)                             // property DragKind Setter
	DragCursor() types.TCursor                                     // property DragCursor Getter
	SetDragCursor(value types.TCursor)                             // property DragCursor Setter
	DragMode() types.TDragMode                                     // property DragMode Getter
	SetDragMode(value types.TDragMode)                             // property DragMode Setter
	HideSelection() bool                                           // property HideSelection Getter
	SetHideSelection(value bool)                                   // property HideSelection Setter
	HotTrack() bool                                                // property HotTrack Getter
	SetHotTrack(value bool)                                        // property HotTrack Setter
	HotTrackColor() types.TColor                                   // property HotTrackColor Getter
	SetHotTrackColor(value types.TColor)                           // property HotTrackColor Setter
	Indent() int32                                                 // property Indent Getter
	SetIndent(value int32)                                         // property Indent Setter
	MultiSelect() bool                                             // property MultiSelect Getter
	SetMultiSelect(value bool)                                     // property MultiSelect Setter
	ParentColor() bool                                             // property ParentColor Getter
	SetParentColor(value bool)                                     // property ParentColor Setter
	ParentFont() bool                                              // property ParentFont Getter
	SetParentFont(value bool)                                      // property ParentFont Setter
	ParentShowHint() bool                                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                                  // property ParentShowHint Setter
	ReadOnly() bool                                                // property ReadOnly Getter
	SetReadOnly(value bool)                                        // property ReadOnly Setter
	RightClickSelect() bool                                        // property RightClickSelect Getter
	SetRightClickSelect(value bool)                                // property RightClickSelect Setter
	RowSelect() bool                                               // property RowSelect Getter
	SetRowSelect(value bool)                                       // property RowSelect Setter
	ShowButtons() bool                                             // property ShowButtons Getter
	SetShowButtons(value bool)                                     // property ShowButtons Setter
	ShowLines() bool                                               // property ShowLines Getter
	SetShowLines(value bool)                                       // property ShowLines Setter
	ShowRoot() bool                                                // property ShowRoot Getter
	SetShowRoot(value bool)                                        // property ShowRoot Setter
	ShowSeparators() bool                                          // property ShowSeparators Getter
	SetShowSeparators(value bool)                                  // property ShowSeparators Setter
	SortType() types.TSortType                                     // property SortType Getter
	SetSortType(value types.TSortType)                             // property SortType Setter
	ToolTips() bool                                                // property ToolTips Getter
	SetToolTips(value bool)                                        // property ToolTips Setter
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

type TTreeView struct {
	TCustomTreeView
}

func (m *TTreeView) PathDelimiter() string {
	if !m.IsValid() {
		return ""
	}
	r := treeViewAPI().SysCallN(1, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTreeView) SetPathDelimiter(value string) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(1, 1, m.Instance(), api.PasStr(value))
}

func (m *TTreeView) ScrolledLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeViewAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TTreeView) SetScrolledLeft(value int32) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TTreeView) ScrolledTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeViewAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TTreeView) SetScrolledTop(value int32) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TTreeView) AutoExpand() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetAutoExpand(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) DisabledFontColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := treeViewAPI().SysCallN(5, 0, m.Instance())
	return types.TColor(r)
}

func (m *TTreeView) SetDisabledFontColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TTreeView) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := treeViewAPI().SysCallN(6, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TTreeView) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TTreeView) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := treeViewAPI().SysCallN(7, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TTreeView) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TTreeView) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := treeViewAPI().SysCallN(8, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TTreeView) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TTreeView) HideSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetHideSelection(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) HotTrack() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetHotTrack(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) HotTrackColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := treeViewAPI().SysCallN(11, 0, m.Instance())
	return types.TColor(r)
}

func (m *TTreeView) SetHotTrackColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TTreeView) Indent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeViewAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TTreeView) SetIndent(value int32) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TTreeView) MultiSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetMultiSelect(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) RightClickSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetRightClickSelect(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) RowSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetRowSelect(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) ShowButtons() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetShowButtons(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) ShowLines() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetShowLines(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) ShowRoot() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(22, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetShowRoot(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(22, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) ShowSeparators() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(23, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetShowSeparators(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(23, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) SortType() types.TSortType {
	if !m.IsValid() {
		return 0
	}
	r := treeViewAPI().SysCallN(24, 0, m.Instance())
	return types.TSortType(r)
}

func (m *TTreeView) SetSortType(value types.TSortType) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TTreeView) ToolTips() bool {
	if !m.IsValid() {
		return false
	}
	r := treeViewAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeView) SetToolTips(value bool) {
	if !m.IsValid() {
		return
	}
	treeViewAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeView) SetOnAddition(fn TTVExpandedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVExpandedEvent(fn)
	base.SetEvent(m, 26, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnAdvancedCustomDraw(fn TTVAdvancedCustomDrawEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVAdvancedCustomDrawEvent(fn)
	base.SetEvent(m, 27, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnAdvancedCustomDrawItem(fn TTVAdvancedCustomDrawItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVAdvancedCustomDrawItemEvent(fn)
	base.SetEvent(m, 28, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnChange(fn TTVChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVChangedEvent(fn)
	base.SetEvent(m, 29, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnChanging(fn TTVChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVChangingEvent(fn)
	base.SetEvent(m, 30, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnCollapsed(fn TTVExpandedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVExpandedEvent(fn)
	base.SetEvent(m, 31, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnCollapsing(fn TTVCollapsingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVCollapsingEvent(fn)
	base.SetEvent(m, 32, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnCompare(fn TTVCompareEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVCompareEvent(fn)
	base.SetEvent(m, 33, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 34, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnCreateNodeClass(fn TTVCreateNodeClassEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVCreateNodeClassEvent(fn)
	base.SetEvent(m, 35, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnCustomCreateItem(fn TTVCustomCreateNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVCustomCreateNodeEvent(fn)
	base.SetEvent(m, 36, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnCustomDraw(fn TTVCustomDrawEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVCustomDrawEvent(fn)
	base.SetEvent(m, 37, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnCustomDrawItem(fn TTVCustomDrawItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVCustomDrawItemEvent(fn)
	base.SetEvent(m, 38, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnCustomDrawArrow(fn TTVCustomDrawArrowEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVCustomDrawArrowEvent(fn)
	base.SetEvent(m, 39, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 40, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnDeletion(fn TTVExpandedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVExpandedEvent(fn)
	base.SetEvent(m, 41, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnDragDrop(fn TDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragDropEvent(fn)
	base.SetEvent(m, 42, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnDragOver(fn TDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDragOverEvent(fn)
	base.SetEvent(m, 43, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnEdited(fn TTVEditedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVEditedEvent(fn)
	base.SetEvent(m, 44, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnEditing(fn TTVEditingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVEditingEvent(fn)
	base.SetEvent(m, 45, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnEditingEnd(fn TTVEditingEndEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVEditingEndEvent(fn)
	base.SetEvent(m, 46, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 47, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnExpanded(fn TTVExpandedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVExpandedEvent(fn)
	base.SetEvent(m, 48, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnExpanding(fn TTVExpandingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVExpandingEvent(fn)
	base.SetEvent(m, 49, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnGetImageIndex(fn TTVExpandedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVExpandedEvent(fn)
	base.SetEvent(m, 50, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnGetSelectedIndex(fn TTVExpandedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVExpandedEvent(fn)
	base.SetEvent(m, 51, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnHasChildren(fn TTVHasChildrenEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVHasChildrenEvent(fn)
	base.SetEvent(m, 52, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 53, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 54, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 55, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 56, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 57, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 58, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 59, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 60, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 61, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 62, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 63, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnNodeChanged(fn TTVNodeChangedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTVNodeChangedEvent(fn)
	base.SetEvent(m, 64, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnSelectionChanged(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 65, treeViewAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTreeView) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 66, treeViewAPI(), api.MakeEventDataPtr(cb))
}

// NewTreeView class constructor
func NewTreeView(anOwner IComponent) ITreeView {
	r := treeViewAPI().SysCallN(0, base.GetObjectUintptr(anOwner))
	return AsTreeView(r)
}

func TTreeViewClass() types.TClass {
	r := treeViewAPI().SysCallN(67)
	return types.TClass(r)
}

var (
	treeViewOnce   base.Once
	treeViewImport *imports.Imports = nil
)

func treeViewAPI() *imports.Imports {
	treeViewOnce.Do(func() {
		treeViewImport = api.NewDefaultImports()
		treeViewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTreeView_Create", 0), // constructor NewTreeView
			/* 1 */ imports.NewTable("TTreeView_PathDelimiter", 0), // property PathDelimiter
			/* 2 */ imports.NewTable("TTreeView_ScrolledLeft", 0), // property ScrolledLeft
			/* 3 */ imports.NewTable("TTreeView_ScrolledTop", 0), // property ScrolledTop
			/* 4 */ imports.NewTable("TTreeView_AutoExpand", 0), // property AutoExpand
			/* 5 */ imports.NewTable("TTreeView_DisabledFontColor", 0), // property DisabledFontColor
			/* 6 */ imports.NewTable("TTreeView_DragKind", 0), // property DragKind
			/* 7 */ imports.NewTable("TTreeView_DragCursor", 0), // property DragCursor
			/* 8 */ imports.NewTable("TTreeView_DragMode", 0), // property DragMode
			/* 9 */ imports.NewTable("TTreeView_HideSelection", 0), // property HideSelection
			/* 10 */ imports.NewTable("TTreeView_HotTrack", 0), // property HotTrack
			/* 11 */ imports.NewTable("TTreeView_HotTrackColor", 0), // property HotTrackColor
			/* 12 */ imports.NewTable("TTreeView_Indent", 0), // property Indent
			/* 13 */ imports.NewTable("TTreeView_MultiSelect", 0), // property MultiSelect
			/* 14 */ imports.NewTable("TTreeView_ParentColor", 0), // property ParentColor
			/* 15 */ imports.NewTable("TTreeView_ParentFont", 0), // property ParentFont
			/* 16 */ imports.NewTable("TTreeView_ParentShowHint", 0), // property ParentShowHint
			/* 17 */ imports.NewTable("TTreeView_ReadOnly", 0), // property ReadOnly
			/* 18 */ imports.NewTable("TTreeView_RightClickSelect", 0), // property RightClickSelect
			/* 19 */ imports.NewTable("TTreeView_RowSelect", 0), // property RowSelect
			/* 20 */ imports.NewTable("TTreeView_ShowButtons", 0), // property ShowButtons
			/* 21 */ imports.NewTable("TTreeView_ShowLines", 0), // property ShowLines
			/* 22 */ imports.NewTable("TTreeView_ShowRoot", 0), // property ShowRoot
			/* 23 */ imports.NewTable("TTreeView_ShowSeparators", 0), // property ShowSeparators
			/* 24 */ imports.NewTable("TTreeView_SortType", 0), // property SortType
			/* 25 */ imports.NewTable("TTreeView_ToolTips", 0), // property ToolTips
			/* 26 */ imports.NewTable("TTreeView_OnAddition", 0), // event OnAddition
			/* 27 */ imports.NewTable("TTreeView_OnAdvancedCustomDraw", 0), // event OnAdvancedCustomDraw
			/* 28 */ imports.NewTable("TTreeView_OnAdvancedCustomDrawItem", 0), // event OnAdvancedCustomDrawItem
			/* 29 */ imports.NewTable("TTreeView_OnChange", 0), // event OnChange
			/* 30 */ imports.NewTable("TTreeView_OnChanging", 0), // event OnChanging
			/* 31 */ imports.NewTable("TTreeView_OnCollapsed", 0), // event OnCollapsed
			/* 32 */ imports.NewTable("TTreeView_OnCollapsing", 0), // event OnCollapsing
			/* 33 */ imports.NewTable("TTreeView_OnCompare", 0), // event OnCompare
			/* 34 */ imports.NewTable("TTreeView_OnContextPopup", 0), // event OnContextPopup
			/* 35 */ imports.NewTable("TTreeView_OnCreateNodeClass", 0), // event OnCreateNodeClass
			/* 36 */ imports.NewTable("TTreeView_OnCustomCreateItem", 0), // event OnCustomCreateItem
			/* 37 */ imports.NewTable("TTreeView_OnCustomDraw", 0), // event OnCustomDraw
			/* 38 */ imports.NewTable("TTreeView_OnCustomDrawItem", 0), // event OnCustomDrawItem
			/* 39 */ imports.NewTable("TTreeView_OnCustomDrawArrow", 0), // event OnCustomDrawArrow
			/* 40 */ imports.NewTable("TTreeView_OnDblClick", 0), // event OnDblClick
			/* 41 */ imports.NewTable("TTreeView_OnDeletion", 0), // event OnDeletion
			/* 42 */ imports.NewTable("TTreeView_OnDragDrop", 0), // event OnDragDrop
			/* 43 */ imports.NewTable("TTreeView_OnDragOver", 0), // event OnDragOver
			/* 44 */ imports.NewTable("TTreeView_OnEdited", 0), // event OnEdited
			/* 45 */ imports.NewTable("TTreeView_OnEditing", 0), // event OnEditing
			/* 46 */ imports.NewTable("TTreeView_OnEditingEnd", 0), // event OnEditingEnd
			/* 47 */ imports.NewTable("TTreeView_OnEndDrag", 0), // event OnEndDrag
			/* 48 */ imports.NewTable("TTreeView_OnExpanded", 0), // event OnExpanded
			/* 49 */ imports.NewTable("TTreeView_OnExpanding", 0), // event OnExpanding
			/* 50 */ imports.NewTable("TTreeView_OnGetImageIndex", 0), // event OnGetImageIndex
			/* 51 */ imports.NewTable("TTreeView_OnGetSelectedIndex", 0), // event OnGetSelectedIndex
			/* 52 */ imports.NewTable("TTreeView_OnHasChildren", 0), // event OnHasChildren
			/* 53 */ imports.NewTable("TTreeView_OnMouseDown", 0), // event OnMouseDown
			/* 54 */ imports.NewTable("TTreeView_OnMouseEnter", 0), // event OnMouseEnter
			/* 55 */ imports.NewTable("TTreeView_OnMouseLeave", 0), // event OnMouseLeave
			/* 56 */ imports.NewTable("TTreeView_OnMouseMove", 0), // event OnMouseMove
			/* 57 */ imports.NewTable("TTreeView_OnMouseUp", 0), // event OnMouseUp
			/* 58 */ imports.NewTable("TTreeView_OnMouseWheel", 0), // event OnMouseWheel
			/* 59 */ imports.NewTable("TTreeView_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 60 */ imports.NewTable("TTreeView_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 61 */ imports.NewTable("TTreeView_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 62 */ imports.NewTable("TTreeView_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 63 */ imports.NewTable("TTreeView_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 64 */ imports.NewTable("TTreeView_OnNodeChanged", 0), // event OnNodeChanged
			/* 65 */ imports.NewTable("TTreeView_OnSelectionChanged", 0), // event OnSelectionChanged
			/* 66 */ imports.NewTable("TTreeView_OnStartDrag", 0), // event OnStartDrag
			/* 67 */ imports.NewTable("TTreeView_TClass", 0), // function TTreeViewClass
		}
	})
	return treeViewImport
}
