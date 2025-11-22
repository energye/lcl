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

// ILazVirtualDrawTree Parent: ICustomVirtualDrawTree
type ILazVirtualDrawTree interface {
	ICustomVirtualDrawTree
	LastDragEffect() uint32                                // property LastDragEffect Getter
	Alignment() types.TAlignment                           // property Alignment Getter
	SetAlignment(value types.TAlignment)                   // property Alignment Setter
	AnimationDuration() uint32                             // property AnimationDuration Getter
	SetAnimationDuration(value uint32)                     // property AnimationDuration Setter
	AutoExpandDelay() uint32                               // property AutoExpandDelay Getter
	SetAutoExpandDelay(value uint32)                       // property AutoExpandDelay Setter
	AutoScrollDelay() uint32                               // property AutoScrollDelay Getter
	SetAutoScrollDelay(value uint32)                       // property AutoScrollDelay Setter
	AutoScrollInterval() types.TAutoScrollInterval         // property AutoScrollInterval Getter
	SetAutoScrollInterval(value types.TAutoScrollInterval) // property AutoScrollInterval Setter
	Background() IPicture                                  // property Background Getter
	SetBackground(value IPicture)                          // property Background Setter
	BackgroundOffsetX() int32                              // property BackgroundOffsetX Getter
	SetBackgroundOffsetX(value int32)                      // property BackgroundOffsetX Setter
	BackgroundOffsetY() int32                              // property BackgroundOffsetY Getter
	SetBackgroundOffsetY(value int32)                      // property BackgroundOffsetY Setter
	// BottomSpace
	//  property BorderStyle: TBorderStyle read FBorderStyle write SetBorderStyle default bsSingle;
	BottomSpace() uint32                                                   // property BottomSpace Getter
	SetBottomSpace(value uint32)                                           // property BottomSpace Setter
	ButtonFillMode() types.TVTButtonFillMode                               // property ButtonFillMode Getter
	SetButtonFillMode(value types.TVTButtonFillMode)                       // property ButtonFillMode Setter
	ButtonStyle() types.TVTButtonStyle                                     // property ButtonStyle Getter
	SetButtonStyle(value types.TVTButtonStyle)                             // property ButtonStyle Setter
	ChangeDelay() uint32                                                   // property ChangeDelay Getter
	SetChangeDelay(value uint32)                                           // property ChangeDelay Setter
	CheckImageKind() types.TCheckImageKind                                 // property CheckImageKind Getter
	SetCheckImageKind(value types.TCheckImageKind)                         // property CheckImageKind Setter
	ClipboardFormats() IClipboardFormats                                   // property ClipboardFormats Getter
	SetClipboardFormats(value IClipboardFormats)                           // property ClipboardFormats Setter
	Colors() IVTColors                                                     // property Colors Getter
	SetColors(value IVTColors)                                             // property Colors Setter
	CustomCheckImages() ICustomImageList                                   // property CustomCheckImages Getter
	SetCustomCheckImages(value ICustomImageList)                           // property CustomCheckImages Setter
	DefaultNodeHeight() uint32                                             // property DefaultNodeHeight Getter
	SetDefaultNodeHeight(value uint32)                                     // property DefaultNodeHeight Setter
	DefaultPasteMode() types.TVTNodeAttachMode                             // property DefaultPasteMode Getter
	SetDefaultPasteMode(value types.TVTNodeAttachMode)                     // property DefaultPasteMode Setter
	DragCursor() types.TCursor                                             // property DragCursor Getter
	SetDragCursor(value types.TCursor)                                     // property DragCursor Setter
	DragHeight() int32                                                     // property DragHeight Getter
	SetDragHeight(value int32)                                             // property DragHeight Setter
	DragKind() types.TDragKind                                             // property DragKind Getter
	SetDragKind(value types.TDragKind)                                     // property DragKind Setter
	DragImageKind() types.TVTDragImageKind                                 // property DragImageKind Getter
	SetDragImageKind(value types.TVTDragImageKind)                         // property DragImageKind Setter
	DragMode() types.TDragMode                                             // property DragMode Getter
	SetDragMode(value types.TDragMode)                                     // property DragMode Setter
	DragOperations() types.TDragOperations                                 // property DragOperations Getter
	SetDragOperations(value types.TDragOperations)                         // property DragOperations Setter
	DragType() types.TVTDragType                                           // property DragType Getter
	SetDragType(value types.TVTDragType)                                   // property DragType Setter
	DragWidth() int32                                                      // property DragWidth Getter
	SetDragWidth(value int32)                                              // property DragWidth Setter
	DrawSelectionMode() types.TVTDrawSelectionMode                         // property DrawSelectionMode Getter
	SetDrawSelectionMode(value types.TVTDrawSelectionMode)                 // property DrawSelectionMode Setter
	EditDelay() uint32                                                     // property EditDelay Getter
	SetEditDelay(value uint32)                                             // property EditDelay Setter
	Header() IVTHeader                                                     // property Header Getter
	SetHeader(value IVTHeader)                                             // property Header Setter
	HintMode() types.TVTHintMode                                           // property HintMode Getter
	SetHintMode(value types.TVTHintMode)                                   // property HintMode Setter
	HotCursor() types.TCursor                                              // property HotCursor Getter
	SetHotCursor(value types.TCursor)                                      // property HotCursor Setter
	Images() ICustomImageList                                              // property Images Getter
	SetImages(value ICustomImageList)                                      // property Images Setter
	IncrementalSearch() types.TVTIncrementalSearch                         // property IncrementalSearch Getter
	SetIncrementalSearch(value types.TVTIncrementalSearch)                 // property IncrementalSearch Setter
	IncrementalSearchDirection() types.TVTSearchDirection                  // property IncrementalSearchDirection Getter
	SetIncrementalSearchDirection(value types.TVTSearchDirection)          // property IncrementalSearchDirection Setter
	IncrementalSearchStart() types.TVTSearchStart                          // property IncrementalSearchStart Getter
	SetIncrementalSearchStart(value types.TVTSearchStart)                  // property IncrementalSearchStart Setter
	IncrementalSearchTimeout() uint32                                      // property IncrementalSearchTimeout Getter
	SetIncrementalSearchTimeout(value uint32)                              // property IncrementalSearchTimeout Setter
	Indent() uint32                                                        // property Indent Getter
	SetIndent(value uint32)                                                // property Indent Setter
	LineMode() types.TVTLineMode                                           // property LineMode Getter
	SetLineMode(value types.TVTLineMode)                                   // property LineMode Setter
	LineStyle() types.TVTLineStyle                                         // property LineStyle Getter
	SetLineStyle(value types.TVTLineStyle)                                 // property LineStyle Setter
	Margin() int32                                                         // property Margin Getter
	SetMargin(value int32)                                                 // property Margin Setter
	NodeAlignment() types.TVTNodeAlignment                                 // property NodeAlignment Getter
	SetNodeAlignment(value types.TVTNodeAlignment)                         // property NodeAlignment Setter
	NodeDataSize() int32                                                   // property NodeDataSize Getter
	SetNodeDataSize(value int32)                                           // property NodeDataSize Setter
	OperationCanceled() bool                                               // property OperationCanceled Getter
	ParentColor() bool                                                     // property ParentColor Getter
	SetParentColor(value bool)                                             // property ParentColor Setter
	ParentFont() bool                                                      // property ParentFont Getter
	SetParentFont(value bool)                                              // property ParentFont Setter
	ParentShowHint() bool                                                  // property ParentShowHint Getter
	SetParentShowHint(value bool)                                          // property ParentShowHint Setter
	RootNodeCount() uint32                                                 // property RootNodeCount Getter
	SetRootNodeCount(value uint32)                                         // property RootNodeCount Setter
	ScrollBarOptions() IScrollBarOptions                                   // property ScrollBarOptions Getter
	SetScrollBarOptions(value IScrollBarOptions)                           // property ScrollBarOptions Setter
	SelectionBlendFactor() byte                                            // property SelectionBlendFactor Getter
	SetSelectionBlendFactor(value byte)                                    // property SelectionBlendFactor Setter
	SelectionCurveRadius() uint32                                          // property SelectionCurveRadius Getter
	SetSelectionCurveRadius(value uint32)                                  // property SelectionCurveRadius Setter
	StateImages() ICustomImageList                                         // property StateImages Getter
	SetStateImages(value ICustomImageList)                                 // property StateImages Setter
	TextMargin() int32                                                     // property TextMargin Getter
	SetTextMargin(value int32)                                             // property TextMargin Setter
	TreeOptions() IVirtualTreeOptions                                      // property TreeOptions Getter
	SetTreeOptions(value IVirtualTreeOptions)                              // property TreeOptions Setter
	WantTabs() bool                                                        // property WantTabs Getter
	SetWantTabs(value bool)                                                // property WantTabs Setter
	SetOnAddToSelection(fn TVTAddToSelectionEvent)                         // property event
	SetOnAdvancedHeaderDraw(fn TVTAdvancedHeaderPaintEvent)                // property event
	SetOnAfterAutoFitColumn(fn TVTAfterAutoFitColumnEvent)                 // property event
	SetOnAfterAutoFitColumns(fn TVTAfterAutoFitColumnsEvent)               // property event
	SetOnAfterCellPaint(fn TVTAfterCellPaintEvent)                         // property event
	SetOnAfterColumnExport(fn TVTColumnExportEvent)                        // property event
	SetOnAfterColumnWidthTracking(fn TVTAfterColumnWidthTrackingEvent)     // property event
	SetOnAfterGetMaxColumnWidth(fn TVTAfterGetMaxColumnWidthEvent)         // property event
	SetOnAfterHeaderExport(fn TVTTreeExportEvent)                          // property event
	SetOnAfterHeaderHeightTracking(fn TVTAfterHeaderHeightTrackingEvent)   // property event
	SetOnAfterItemErase(fn TVTAfterItemEraseEvent)                         // property event
	SetOnAfterItemPaint(fn TVTAfterItemPaintEvent)                         // property event
	SetOnAfterNodeExport(fn TVTNodeExportEvent)                            // property event
	SetOnAfterPaint(fn TVTPaintEvent)                                      // property event
	SetOnAfterTreeExport(fn TVTTreeExportEvent)                            // property event
	SetOnBeforeAutoFitColumn(fn TVTBeforeAutoFitColumnEvent)               // property event
	SetOnBeforeAutoFitColumns(fn TVTBeforeAutoFitColumnsEvent)             // property event
	SetOnBeforeCellPaint(fn TVTBeforeCellPaintEvent)                       // property event
	SetOnBeforeColumnExport(fn TVTColumnExportEvent)                       // property event
	SetOnBeforeColumnWidthTracking(fn TVTBeforeColumnWidthTrackingEvent)   // property event
	SetOnBeforeDrawTreeLine(fn TVTBeforeDrawLineImageEvent)                // property event
	SetOnBeforeGetMaxColumnWidth(fn TVTBeforeGetMaxColumnWidthEvent)       // property event
	SetOnBeforeHeaderExport(fn TVTTreeExportEvent)                         // property event
	SetOnBeforeHeaderHeightTracking(fn TVTBeforeHeaderHeightTrackingEvent) // property event
	SetOnBeforeItemErase(fn TVTBeforeItemEraseEvent)                       // property event
	SetOnBeforeItemPaint(fn TVTBeforeItemPaintEvent)                       // property event
	SetOnBeforeNodeExport(fn TVTNodeExportEvent)                           // property event
	SetOnBeforePaint(fn TVTPaintEvent)                                     // property event
	SetOnBeforeTreeExport(fn TVTTreeExportEvent)                           // property event
	SetOnCanSplitterResizeColumn(fn TVTCanSplitterResizeColumnEvent)       // property event
	SetOnCanSplitterResizeHeader(fn TVTCanSplitterResizeHeaderEvent)       // property event
	SetOnCanSplitterResizeNode(fn TVTCanSplitterResizeNodeEvent)           // property event
	SetOnChange(fn TVTChangeEvent)                                         // property event
	SetOnChecked(fn TVTChangeEvent)                                        // property event
	SetOnChecking(fn TVTCheckChangingEvent)                                // property event
	SetOnCollapsed(fn TVTChangeEvent)                                      // property event
	SetOnCollapsing(fn TVTChangingEvent)                                   // property event
	SetOnColumnClick(fn TVTColumnClickEvent)                               // property event
	SetOnColumnDblClick(fn TVTColumnDblClickEvent)                         // property event
	SetOnColumnExport(fn TVTColumnExportEvent)                             // property event
	SetOnColumnResize(fn TVTHeaderNotifyEvent)                             // property event
	SetOnColumnWidthDblClickResize(fn TVTColumnWidthDblClickResizeEvent)   // property event
	SetOnColumnWidthTracking(fn TVTColumnWidthTrackingEvent)               // property event
	SetOnCompareNodes(fn TVTCompareEvent)                                  // property event
	SetOnContextPopup(fn TContextPopupEvent)                               // property event
	SetOnCreateDataObject(fn TVTCreateDataObjectEvent)                     // property event
	SetOnCreateDragManager(fn TVTCreateDragManagerEvent)                   // property event
	SetOnCreateEditor(fn TVTCreateEditorEvent)                             // property event
	SetOnDblClick(fn TNotifyEvent)                                         // property event
	SetOnDragAllowed(fn TVTDragAllowedEvent)                               // property event
	SetOnDragOver(fn TVTDragOverEvent)                                     // property event
	SetOnDragDrop(fn TVTDragDropEvent)                                     // property event
	SetOnDrawHint(fn TVTDrawHintEvent)                                     // property event
	SetOnDrawNode(fn TVTDrawNodeEvent)                                     // property event
	SetOnEdited(fn TVTEditChangeEvent)                                     // property event
	SetOnEditing(fn TVTEditChangingEvent)                                  // property event
	SetOnEndDock(fn TEndDragEvent)                                         // property event
	SetOnEndDrag(fn TEndDragEvent)                                         // property event
	SetOnEndOperation(fn TVTOperationEvent)                                // property event
	SetOnExpanded(fn TVTChangeEvent)                                       // property event
	SetOnExpanding(fn TVTChangingEvent)                                    // property event
	SetOnFocusChanged(fn TVTFocusChangeEvent)                              // property event
	SetOnFocusChanging(fn TVTFocusChangingEvent)                           // property event
	SetOnFreeNode(fn TVTFreeNodeEvent)                                     // property event
	SetOnGetCellIsEmpty(fn TVTGetCellIsEmptyEvent)                         // property event
	SetOnGetCursor(fn TVTGetCursorEvent)                                   // property event
	SetOnGetHeaderCursor(fn TVTGetHeaderCursorEvent)                       // property event
	SetOnGetHelpContext(fn TVTHelpContextEvent)                            // property event
	SetOnGetHint(fn TVTGetHintEvent)                                       // property event
	SetOnGetHintKind(fn TVTHintKindEvent)                                  // property event
	SetOnGetHintSize(fn TVTGetHintSizeEvent)                               // property event
	SetOnGetImageIndex(fn TVTGetImageEvent)                                // property event
	SetOnGetImageIndexEx(fn TVTGetImageExEvent)                            // property event
	SetOnGetLineStyle(fn TVTGetLineStyleEvent)                             // property event
	SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent)                       // property event
	SetOnGetNodeWidth(fn TVTGetNodeWidthEvent)                             // property event
	SetOnGetPopupMenu(fn TVTPopupEvent)                                    // property event
	SetOnGetUserClipboardFormats(fn TVTGetUserClipboardFormatsEvent)       // property event
	SetOnHeaderClick(fn TVTHeaderClickEvent)                               // property event
	SetOnHeaderDblClick(fn TVTHeaderClickEvent)                            // property event
	SetOnHeaderDragged(fn TVTHeaderDraggedEvent)                           // property event
	SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent)                     // property event
	SetOnHeaderDragging(fn TVTHeaderDraggingEvent)                         // property event
	SetOnHeaderDraw(fn TVTHeaderPaintEvent)                                // property event
	SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent)      // property event
	SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent)             // property event
	SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) // property event
	SetOnHeaderMouseDown(fn TVTHeaderMouseEvent)                           // property event
	SetOnHeaderMouseMove(fn TVTHeaderMouseMoveEvent)                       // property event
	SetOnHeaderMouseUp(fn TVTHeaderMouseEvent)                             // property event
	SetOnHotChange(fn TVTHotNodeChangeEvent)                               // property event
	SetOnIncrementalSearch(fn TVTIncrementalSearchEvent)                   // property event
	SetOnInitChildren(fn TVTInitChildrenEvent)                             // property event
	SetOnInitNode(fn TVTInitNodeEvent)                                     // property event
	SetOnKeyAction(fn TVTKeyActionEvent)                                   // property event
	SetOnLoadNode(fn TVTSaveNodeEvent)                                     // property event
	SetOnLoadTree(fn TVTSaveTreeEvent)                                     // property event
	SetOnMeasureItem(fn TVTMeasureItemEvent)                               // property event
	SetOnMouseDown(fn TMouseEvent)                                         // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                     // property event
	SetOnMouseUp(fn TMouseEvent)                                           // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                   // property event
	SetOnNodeClick(fn TVTNodeClickEvent)                                   // property event
	SetOnNodeCopied(fn TVTNodeCopiedEvent)                                 // property event
	SetOnNodeCopying(fn TVTNodeCopyingEvent)                               // property event
	SetOnNodeDblClick(fn TVTNodeClickEvent)                                // property event
	SetOnNodeExport(fn TVTNodeExportEvent)                                 // property event
	SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent)                 // property event
	SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent)     // property event
	SetOnNodeMoved(fn TVTNodeMovedEvent)                                   // property event
	SetOnNodeMoving(fn TVTNodeMovingEvent)                                 // property event
	SetOnPaintBackground(fn TVTBackgroundPaintEvent)                       // property event
	SetOnRemoveFromSelection(fn TVTRemoveFromSelectionEvent)               // property event
	SetOnResetNode(fn TVTChangeEvent)                                      // property event
	SetOnSaveNode(fn TVTSaveNodeEvent)                                     // property event
	SetOnSaveTree(fn TVTSaveTreeEvent)                                     // property event
	SetOnScroll(fn TVTScrollEvent)                                         // property event
	SetOnShowScrollBar(fn TVTScrollBarShowEvent)                           // property event
	SetOnStartDock(fn TStartDockEvent)                                     // property event
	SetOnStartDrag(fn TStartDragEvent)                                     // property event
	SetOnStartOperation(fn TVTOperationEvent)                              // property event
	SetOnStateChange(fn TVTStateChangeEvent)                               // property event
	SetOnStructureChange(fn TVTStructureChangeEvent)                       // property event
	SetOnUpdating(fn TVTUpdatingEvent)                                     // property event
}

type TLazVirtualDrawTree struct {
	TCustomVirtualDrawTree
}

func (m *TLazVirtualDrawTree) LastDragEffect() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(1, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(2, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TLazVirtualDrawTree) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) AnimationDuration() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(3, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetAnimationDuration(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) AutoExpandDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(4, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetAutoExpandDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) AutoScrollDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(5, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetAutoScrollDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) AutoScrollInterval() types.TAutoScrollInterval {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(6, 0, m.Instance())
	return types.TAutoScrollInterval(r)
}

func (m *TLazVirtualDrawTree) SetAutoScrollInterval(value types.TAutoScrollInterval) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) Background() IPicture {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualDrawTreeAPI().SysCallN(7, 0, m.Instance())
	return AsPicture(r)
}

func (m *TLazVirtualDrawTree) SetBackground(value IPicture) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualDrawTree) BackgroundOffsetX() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualDrawTree) SetBackgroundOffsetX(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) BackgroundOffsetY() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualDrawTree) SetBackgroundOffsetY(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) BottomSpace() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(10, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetBottomSpace(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) ButtonFillMode() types.TVTButtonFillMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(11, 0, m.Instance())
	return types.TVTButtonFillMode(r)
}

func (m *TLazVirtualDrawTree) SetButtonFillMode(value types.TVTButtonFillMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) ButtonStyle() types.TVTButtonStyle {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(12, 0, m.Instance())
	return types.TVTButtonStyle(r)
}

func (m *TLazVirtualDrawTree) SetButtonStyle(value types.TVTButtonStyle) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) ChangeDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(13, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetChangeDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) CheckImageKind() types.TCheckImageKind {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(14, 0, m.Instance())
	return types.TCheckImageKind(r)
}

func (m *TLazVirtualDrawTree) SetCheckImageKind(value types.TCheckImageKind) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) ClipboardFormats() IClipboardFormats {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualDrawTreeAPI().SysCallN(15, 0, m.Instance())
	return AsClipboardFormats(r)
}

func (m *TLazVirtualDrawTree) SetClipboardFormats(value IClipboardFormats) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(15, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualDrawTree) Colors() IVTColors {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualDrawTreeAPI().SysCallN(16, 0, m.Instance())
	return AsVTColors(r)
}

func (m *TLazVirtualDrawTree) SetColors(value IVTColors) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualDrawTree) CustomCheckImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualDrawTreeAPI().SysCallN(17, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TLazVirtualDrawTree) SetCustomCheckImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(17, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualDrawTree) DefaultNodeHeight() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(18, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetDefaultNodeHeight(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DefaultPasteMode() types.TVTNodeAttachMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(19, 0, m.Instance())
	return types.TVTNodeAttachMode(r)
}

func (m *TLazVirtualDrawTree) SetDefaultPasteMode(value types.TVTNodeAttachMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(20, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TLazVirtualDrawTree) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DragHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(21, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualDrawTree) SetDragHeight(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(22, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TLazVirtualDrawTree) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DragImageKind() types.TVTDragImageKind {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(23, 0, m.Instance())
	return types.TVTDragImageKind(r)
}

func (m *TLazVirtualDrawTree) SetDragImageKind(value types.TVTDragImageKind) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(24, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TLazVirtualDrawTree) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DragOperations() types.TDragOperations {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(25, 0, m.Instance())
	return types.TDragOperations(r)
}

func (m *TLazVirtualDrawTree) SetDragOperations(value types.TDragOperations) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DragType() types.TVTDragType {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(26, 0, m.Instance())
	return types.TVTDragType(r)
}

func (m *TLazVirtualDrawTree) SetDragType(value types.TVTDragType) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DragWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(27, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualDrawTree) SetDragWidth(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) DrawSelectionMode() types.TVTDrawSelectionMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(28, 0, m.Instance())
	return types.TVTDrawSelectionMode(r)
}

func (m *TLazVirtualDrawTree) SetDrawSelectionMode(value types.TVTDrawSelectionMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) EditDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(29, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetEditDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) Header() IVTHeader {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualDrawTreeAPI().SysCallN(30, 0, m.Instance())
	return AsVTHeader(r)
}

func (m *TLazVirtualDrawTree) SetHeader(value IVTHeader) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(30, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualDrawTree) HintMode() types.TVTHintMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(31, 0, m.Instance())
	return types.TVTHintMode(r)
}

func (m *TLazVirtualDrawTree) SetHintMode(value types.TVTHintMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) HotCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(32, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TLazVirtualDrawTree) SetHotCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(32, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualDrawTreeAPI().SysCallN(33, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TLazVirtualDrawTree) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(33, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualDrawTree) IncrementalSearch() types.TVTIncrementalSearch {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(34, 0, m.Instance())
	return types.TVTIncrementalSearch(r)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearch(value types.TVTIncrementalSearch) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(34, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) IncrementalSearchDirection() types.TVTSearchDirection {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(35, 0, m.Instance())
	return types.TVTSearchDirection(r)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearchDirection(value types.TVTSearchDirection) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(35, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) IncrementalSearchStart() types.TVTSearchStart {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(36, 0, m.Instance())
	return types.TVTSearchStart(r)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearchStart(value types.TVTSearchStart) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) IncrementalSearchTimeout() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(37, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearchTimeout(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(37, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) Indent() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(38, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetIndent(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(38, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) LineMode() types.TVTLineMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(39, 0, m.Instance())
	return types.TVTLineMode(r)
}

func (m *TLazVirtualDrawTree) SetLineMode(value types.TVTLineMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(39, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) LineStyle() types.TVTLineStyle {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(40, 0, m.Instance())
	return types.TVTLineStyle(r)
}

func (m *TLazVirtualDrawTree) SetLineStyle(value types.TVTLineStyle) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(40, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) Margin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(41, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualDrawTree) SetMargin(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(41, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) NodeAlignment() types.TVTNodeAlignment {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(42, 0, m.Instance())
	return types.TVTNodeAlignment(r)
}

func (m *TLazVirtualDrawTree) SetNodeAlignment(value types.TVTNodeAlignment) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(42, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) NodeDataSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(43, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualDrawTree) SetNodeDataSize(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(43, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) OperationCanceled() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualDrawTreeAPI().SysCallN(44, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualDrawTree) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualDrawTreeAPI().SysCallN(45, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualDrawTree) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(45, 1, m.Instance(), api.PasBool(value))
}

func (m *TLazVirtualDrawTree) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualDrawTreeAPI().SysCallN(46, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualDrawTree) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(46, 1, m.Instance(), api.PasBool(value))
}

func (m *TLazVirtualDrawTree) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualDrawTreeAPI().SysCallN(47, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualDrawTree) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(47, 1, m.Instance(), api.PasBool(value))
}

func (m *TLazVirtualDrawTree) RootNodeCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(48, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetRootNodeCount(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(48, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) ScrollBarOptions() IScrollBarOptions {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualDrawTreeAPI().SysCallN(49, 0, m.Instance())
	return AsScrollBarOptions(r)
}

func (m *TLazVirtualDrawTree) SetScrollBarOptions(value IScrollBarOptions) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(49, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualDrawTree) SelectionBlendFactor() byte {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(50, 0, m.Instance())
	return byte(r)
}

func (m *TLazVirtualDrawTree) SetSelectionBlendFactor(value byte) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(50, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) SelectionCurveRadius() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(51, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualDrawTree) SetSelectionCurveRadius(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(51, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) StateImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualDrawTreeAPI().SysCallN(52, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TLazVirtualDrawTree) SetStateImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(52, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualDrawTree) TextMargin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualDrawTreeAPI().SysCallN(53, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualDrawTree) SetTextMargin(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(53, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualDrawTree) TreeOptions() IVirtualTreeOptions {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualDrawTreeAPI().SysCallN(54, 0, m.Instance())
	return AsVirtualTreeOptions(r)
}

func (m *TLazVirtualDrawTree) SetTreeOptions(value IVirtualTreeOptions) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(54, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualDrawTree) WantTabs() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualDrawTreeAPI().SysCallN(55, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualDrawTree) SetWantTabs(value bool) {
	if !m.IsValid() {
		return
	}
	lazVirtualDrawTreeAPI().SysCallN(55, 1, m.Instance(), api.PasBool(value))
}

func (m *TLazVirtualDrawTree) SetOnAddToSelection(fn TVTAddToSelectionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAddToSelectionEvent(fn)
	base.SetEvent(m, 56, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAdvancedHeaderDraw(fn TVTAdvancedHeaderPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAdvancedHeaderPaintEvent(fn)
	base.SetEvent(m, 57, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterAutoFitColumn(fn TVTAfterAutoFitColumnEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterAutoFitColumnEvent(fn)
	base.SetEvent(m, 58, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterAutoFitColumns(fn TVTAfterAutoFitColumnsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterAutoFitColumnsEvent(fn)
	base.SetEvent(m, 59, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterCellPaint(fn TVTAfterCellPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterCellPaintEvent(fn)
	base.SetEvent(m, 60, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterColumnExport(fn TVTColumnExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnExportEvent(fn)
	base.SetEvent(m, 61, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterColumnWidthTracking(fn TVTAfterColumnWidthTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterColumnWidthTrackingEvent(fn)
	base.SetEvent(m, 62, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterGetMaxColumnWidth(fn TVTAfterGetMaxColumnWidthEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterGetMaxColumnWidthEvent(fn)
	base.SetEvent(m, 63, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterHeaderExport(fn TVTTreeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTTreeExportEvent(fn)
	base.SetEvent(m, 64, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterHeaderHeightTracking(fn TVTAfterHeaderHeightTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterHeaderHeightTrackingEvent(fn)
	base.SetEvent(m, 65, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterItemErase(fn TVTAfterItemEraseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterItemEraseEvent(fn)
	base.SetEvent(m, 66, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterItemPaint(fn TVTAfterItemPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterItemPaintEvent(fn)
	base.SetEvent(m, 67, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterNodeExport(fn TVTNodeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeExportEvent(fn)
	base.SetEvent(m, 68, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterPaint(fn TVTPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTPaintEvent(fn)
	base.SetEvent(m, 69, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnAfterTreeExport(fn TVTTreeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTTreeExportEvent(fn)
	base.SetEvent(m, 70, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeAutoFitColumn(fn TVTBeforeAutoFitColumnEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeAutoFitColumnEvent(fn)
	base.SetEvent(m, 71, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeAutoFitColumns(fn TVTBeforeAutoFitColumnsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeAutoFitColumnsEvent(fn)
	base.SetEvent(m, 72, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeCellPaint(fn TVTBeforeCellPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeCellPaintEvent(fn)
	base.SetEvent(m, 73, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeColumnExport(fn TVTColumnExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnExportEvent(fn)
	base.SetEvent(m, 74, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeColumnWidthTracking(fn TVTBeforeColumnWidthTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeColumnWidthTrackingEvent(fn)
	base.SetEvent(m, 75, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeDrawTreeLine(fn TVTBeforeDrawLineImageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeDrawLineImageEvent(fn)
	base.SetEvent(m, 76, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeGetMaxColumnWidth(fn TVTBeforeGetMaxColumnWidthEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeGetMaxColumnWidthEvent(fn)
	base.SetEvent(m, 77, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeHeaderExport(fn TVTTreeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTTreeExportEvent(fn)
	base.SetEvent(m, 78, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeHeaderHeightTracking(fn TVTBeforeHeaderHeightTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeHeaderHeightTrackingEvent(fn)
	base.SetEvent(m, 79, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeItemErase(fn TVTBeforeItemEraseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeItemEraseEvent(fn)
	base.SetEvent(m, 80, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeItemPaint(fn TVTBeforeItemPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeItemPaintEvent(fn)
	base.SetEvent(m, 81, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeNodeExport(fn TVTNodeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeExportEvent(fn)
	base.SetEvent(m, 82, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforePaint(fn TVTPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTPaintEvent(fn)
	base.SetEvent(m, 83, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnBeforeTreeExport(fn TVTTreeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTTreeExportEvent(fn)
	base.SetEvent(m, 84, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnCanSplitterResizeColumn(fn TVTCanSplitterResizeColumnEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCanSplitterResizeColumnEvent(fn)
	base.SetEvent(m, 85, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnCanSplitterResizeHeader(fn TVTCanSplitterResizeHeaderEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCanSplitterResizeHeaderEvent(fn)
	base.SetEvent(m, 86, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnCanSplitterResizeNode(fn TVTCanSplitterResizeNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCanSplitterResizeNodeEvent(fn)
	base.SetEvent(m, 87, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnChange(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 88, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnChecked(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 89, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnChecking(fn TVTCheckChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCheckChangingEvent(fn)
	base.SetEvent(m, 90, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnCollapsed(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 91, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnCollapsing(fn TVTChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangingEvent(fn)
	base.SetEvent(m, 92, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnColumnClick(fn TVTColumnClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnClickEvent(fn)
	base.SetEvent(m, 93, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnColumnDblClick(fn TVTColumnDblClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnDblClickEvent(fn)
	base.SetEvent(m, 94, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnColumnExport(fn TVTColumnExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnExportEvent(fn)
	base.SetEvent(m, 95, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnColumnResize(fn TVTHeaderNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderNotifyEvent(fn)
	base.SetEvent(m, 96, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnColumnWidthDblClickResize(fn TVTColumnWidthDblClickResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnWidthDblClickResizeEvent(fn)
	base.SetEvent(m, 97, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnColumnWidthTracking(fn TVTColumnWidthTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnWidthTrackingEvent(fn)
	base.SetEvent(m, 98, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnCompareNodes(fn TVTCompareEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCompareEvent(fn)
	base.SetEvent(m, 99, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 100, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnCreateDataObject(fn TVTCreateDataObjectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCreateDataObjectEvent(fn)
	base.SetEvent(m, 101, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnCreateDragManager(fn TVTCreateDragManagerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCreateDragManagerEvent(fn)
	base.SetEvent(m, 102, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnCreateEditor(fn TVTCreateEditorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCreateEditorEvent(fn)
	base.SetEvent(m, 103, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 104, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnDragAllowed(fn TVTDragAllowedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDragAllowedEvent(fn)
	base.SetEvent(m, 105, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnDragOver(fn TVTDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDragOverEvent(fn)
	base.SetEvent(m, 106, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnDragDrop(fn TVTDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDragDropEvent(fn)
	base.SetEvent(m, 107, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnDrawHint(fn TVTDrawHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDrawHintEvent(fn)
	base.SetEvent(m, 108, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnDrawNode(fn TVTDrawNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDrawNodeEvent(fn)
	base.SetEvent(m, 109, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnEdited(fn TVTEditChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTEditChangeEvent(fn)
	base.SetEvent(m, 110, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnEditing(fn TVTEditChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTEditChangingEvent(fn)
	base.SetEvent(m, 111, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 112, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 113, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnEndOperation(fn TVTOperationEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTOperationEvent(fn)
	base.SetEvent(m, 114, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnExpanded(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 115, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnExpanding(fn TVTChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangingEvent(fn)
	base.SetEvent(m, 116, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnFocusChanged(fn TVTFocusChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTFocusChangeEvent(fn)
	base.SetEvent(m, 117, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnFocusChanging(fn TVTFocusChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTFocusChangingEvent(fn)
	base.SetEvent(m, 118, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnFreeNode(fn TVTFreeNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTFreeNodeEvent(fn)
	base.SetEvent(m, 119, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetCellIsEmpty(fn TVTGetCellIsEmptyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetCellIsEmptyEvent(fn)
	base.SetEvent(m, 120, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetCursor(fn TVTGetCursorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetCursorEvent(fn)
	base.SetEvent(m, 121, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetHeaderCursor(fn TVTGetHeaderCursorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetHeaderCursorEvent(fn)
	base.SetEvent(m, 122, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetHelpContext(fn TVTHelpContextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHelpContextEvent(fn)
	base.SetEvent(m, 123, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetHint(fn TVTGetHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetHintEvent(fn)
	base.SetEvent(m, 124, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetHintKind(fn TVTHintKindEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHintKindEvent(fn)
	base.SetEvent(m, 125, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetHintSize(fn TVTGetHintSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetHintSizeEvent(fn)
	base.SetEvent(m, 126, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetImageIndex(fn TVTGetImageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetImageEvent(fn)
	base.SetEvent(m, 127, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetImageIndexEx(fn TVTGetImageExEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetImageExEvent(fn)
	base.SetEvent(m, 128, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetLineStyle(fn TVTGetLineStyleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetLineStyleEvent(fn)
	base.SetEvent(m, 129, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetNodeDataSizeEvent(fn)
	base.SetEvent(m, 130, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetNodeWidth(fn TVTGetNodeWidthEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetNodeWidthEvent(fn)
	base.SetEvent(m, 131, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetPopupMenu(fn TVTPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTPopupEvent(fn)
	base.SetEvent(m, 132, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnGetUserClipboardFormats(fn TVTGetUserClipboardFormatsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetUserClipboardFormatsEvent(fn)
	base.SetEvent(m, 133, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderClick(fn TVTHeaderClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderClickEvent(fn)
	base.SetEvent(m, 134, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderDblClick(fn TVTHeaderClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderClickEvent(fn)
	base.SetEvent(m, 135, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderDragged(fn TVTHeaderDraggedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderDraggedEvent(fn)
	base.SetEvent(m, 136, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderDraggedOutEvent(fn)
	base.SetEvent(m, 137, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderDragging(fn TVTHeaderDraggingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderDraggingEvent(fn)
	base.SetEvent(m, 138, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderDraw(fn TVTHeaderPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderPaintEvent(fn)
	base.SetEvent(m, 139, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderPaintQueryElementsEvent(fn)
	base.SetEvent(m, 140, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderHeightTrackingEvent(fn)
	base.SetEvent(m, 141, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderHeightDblClickResizeEvent(fn)
	base.SetEvent(m, 142, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderMouseDown(fn TVTHeaderMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderMouseEvent(fn)
	base.SetEvent(m, 143, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderMouseMove(fn TVTHeaderMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderMouseMoveEvent(fn)
	base.SetEvent(m, 144, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHeaderMouseUp(fn TVTHeaderMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderMouseEvent(fn)
	base.SetEvent(m, 145, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnHotChange(fn TVTHotNodeChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHotNodeChangeEvent(fn)
	base.SetEvent(m, 146, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnIncrementalSearch(fn TVTIncrementalSearchEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTIncrementalSearchEvent(fn)
	base.SetEvent(m, 147, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnInitChildren(fn TVTInitChildrenEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTInitChildrenEvent(fn)
	base.SetEvent(m, 148, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnInitNode(fn TVTInitNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTInitNodeEvent(fn)
	base.SetEvent(m, 149, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnKeyAction(fn TVTKeyActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTKeyActionEvent(fn)
	base.SetEvent(m, 150, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnLoadNode(fn TVTSaveNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTSaveNodeEvent(fn)
	base.SetEvent(m, 151, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnLoadTree(fn TVTSaveTreeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTSaveTreeEvent(fn)
	base.SetEvent(m, 152, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnMeasureItem(fn TVTMeasureItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTMeasureItemEvent(fn)
	base.SetEvent(m, 153, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 154, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 155, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 156, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 157, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnNodeClick(fn TVTNodeClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeClickEvent(fn)
	base.SetEvent(m, 158, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnNodeCopied(fn TVTNodeCopiedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeCopiedEvent(fn)
	base.SetEvent(m, 159, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnNodeCopying(fn TVTNodeCopyingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeCopyingEvent(fn)
	base.SetEvent(m, 160, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnNodeDblClick(fn TVTNodeClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeClickEvent(fn)
	base.SetEvent(m, 161, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnNodeExport(fn TVTNodeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeExportEvent(fn)
	base.SetEvent(m, 162, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeHeightTrackingEvent(fn)
	base.SetEvent(m, 163, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeHeightDblClickResizeEvent(fn)
	base.SetEvent(m, 164, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnNodeMoved(fn TVTNodeMovedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeMovedEvent(fn)
	base.SetEvent(m, 165, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnNodeMoving(fn TVTNodeMovingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeMovingEvent(fn)
	base.SetEvent(m, 166, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnPaintBackground(fn TVTBackgroundPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBackgroundPaintEvent(fn)
	base.SetEvent(m, 167, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnRemoveFromSelection(fn TVTRemoveFromSelectionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTRemoveFromSelectionEvent(fn)
	base.SetEvent(m, 168, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnResetNode(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 169, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnSaveNode(fn TVTSaveNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTSaveNodeEvent(fn)
	base.SetEvent(m, 170, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnSaveTree(fn TVTSaveTreeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTSaveTreeEvent(fn)
	base.SetEvent(m, 171, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnScroll(fn TVTScrollEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTScrollEvent(fn)
	base.SetEvent(m, 172, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnShowScrollBar(fn TVTScrollBarShowEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTScrollBarShowEvent(fn)
	base.SetEvent(m, 173, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 174, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 175, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnStartOperation(fn TVTOperationEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTOperationEvent(fn)
	base.SetEvent(m, 176, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnStateChange(fn TVTStateChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTStateChangeEvent(fn)
	base.SetEvent(m, 177, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnStructureChange(fn TVTStructureChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTStructureChangeEvent(fn)
	base.SetEvent(m, 178, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualDrawTree) SetOnUpdating(fn TVTUpdatingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTUpdatingEvent(fn)
	base.SetEvent(m, 179, lazVirtualDrawTreeAPI(), api.MakeEventDataPtr(cb))
}

// NewLazVirtualDrawTree class constructor
func NewLazVirtualDrawTree(owner IComponent) ILazVirtualDrawTree {
	r := lazVirtualDrawTreeAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLazVirtualDrawTree(r)
}

func TLazVirtualDrawTreeClass() types.TClass {
	r := lazVirtualDrawTreeAPI().SysCallN(180)
	return types.TClass(r)
}

var (
	lazVirtualDrawTreeOnce   base.Once
	lazVirtualDrawTreeImport *imports.Imports = nil
)

func lazVirtualDrawTreeAPI() *imports.Imports {
	lazVirtualDrawTreeOnce.Do(func() {
		lazVirtualDrawTreeImport = api.NewDefaultImports()
		lazVirtualDrawTreeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazVirtualDrawTree_Create", 0), // constructor NewLazVirtualDrawTree
			/* 1 */ imports.NewTable("TLazVirtualDrawTree_LastDragEffect", 0), // property LastDragEffect
			/* 2 */ imports.NewTable("TLazVirtualDrawTree_Alignment", 0), // property Alignment
			/* 3 */ imports.NewTable("TLazVirtualDrawTree_AnimationDuration", 0), // property AnimationDuration
			/* 4 */ imports.NewTable("TLazVirtualDrawTree_AutoExpandDelay", 0), // property AutoExpandDelay
			/* 5 */ imports.NewTable("TLazVirtualDrawTree_AutoScrollDelay", 0), // property AutoScrollDelay
			/* 6 */ imports.NewTable("TLazVirtualDrawTree_AutoScrollInterval", 0), // property AutoScrollInterval
			/* 7 */ imports.NewTable("TLazVirtualDrawTree_Background", 0), // property Background
			/* 8 */ imports.NewTable("TLazVirtualDrawTree_BackgroundOffsetX", 0), // property BackgroundOffsetX
			/* 9 */ imports.NewTable("TLazVirtualDrawTree_BackgroundOffsetY", 0), // property BackgroundOffsetY
			/* 10 */ imports.NewTable("TLazVirtualDrawTree_BottomSpace", 0), // property BottomSpace
			/* 11 */ imports.NewTable("TLazVirtualDrawTree_ButtonFillMode", 0), // property ButtonFillMode
			/* 12 */ imports.NewTable("TLazVirtualDrawTree_ButtonStyle", 0), // property ButtonStyle
			/* 13 */ imports.NewTable("TLazVirtualDrawTree_ChangeDelay", 0), // property ChangeDelay
			/* 14 */ imports.NewTable("TLazVirtualDrawTree_CheckImageKind", 0), // property CheckImageKind
			/* 15 */ imports.NewTable("TLazVirtualDrawTree_ClipboardFormats", 0), // property ClipboardFormats
			/* 16 */ imports.NewTable("TLazVirtualDrawTree_Colors", 0), // property Colors
			/* 17 */ imports.NewTable("TLazVirtualDrawTree_CustomCheckImages", 0), // property CustomCheckImages
			/* 18 */ imports.NewTable("TLazVirtualDrawTree_DefaultNodeHeight", 0), // property DefaultNodeHeight
			/* 19 */ imports.NewTable("TLazVirtualDrawTree_DefaultPasteMode", 0), // property DefaultPasteMode
			/* 20 */ imports.NewTable("TLazVirtualDrawTree_DragCursor", 0), // property DragCursor
			/* 21 */ imports.NewTable("TLazVirtualDrawTree_DragHeight", 0), // property DragHeight
			/* 22 */ imports.NewTable("TLazVirtualDrawTree_DragKind", 0), // property DragKind
			/* 23 */ imports.NewTable("TLazVirtualDrawTree_DragImageKind", 0), // property DragImageKind
			/* 24 */ imports.NewTable("TLazVirtualDrawTree_DragMode", 0), // property DragMode
			/* 25 */ imports.NewTable("TLazVirtualDrawTree_DragOperations", 0), // property DragOperations
			/* 26 */ imports.NewTable("TLazVirtualDrawTree_DragType", 0), // property DragType
			/* 27 */ imports.NewTable("TLazVirtualDrawTree_DragWidth", 0), // property DragWidth
			/* 28 */ imports.NewTable("TLazVirtualDrawTree_DrawSelectionMode", 0), // property DrawSelectionMode
			/* 29 */ imports.NewTable("TLazVirtualDrawTree_EditDelay", 0), // property EditDelay
			/* 30 */ imports.NewTable("TLazVirtualDrawTree_Header", 0), // property Header
			/* 31 */ imports.NewTable("TLazVirtualDrawTree_HintMode", 0), // property HintMode
			/* 32 */ imports.NewTable("TLazVirtualDrawTree_HotCursor", 0), // property HotCursor
			/* 33 */ imports.NewTable("TLazVirtualDrawTree_Images", 0), // property Images
			/* 34 */ imports.NewTable("TLazVirtualDrawTree_IncrementalSearch", 0), // property IncrementalSearch
			/* 35 */ imports.NewTable("TLazVirtualDrawTree_IncrementalSearchDirection", 0), // property IncrementalSearchDirection
			/* 36 */ imports.NewTable("TLazVirtualDrawTree_IncrementalSearchStart", 0), // property IncrementalSearchStart
			/* 37 */ imports.NewTable("TLazVirtualDrawTree_IncrementalSearchTimeout", 0), // property IncrementalSearchTimeout
			/* 38 */ imports.NewTable("TLazVirtualDrawTree_Indent", 0), // property Indent
			/* 39 */ imports.NewTable("TLazVirtualDrawTree_LineMode", 0), // property LineMode
			/* 40 */ imports.NewTable("TLazVirtualDrawTree_LineStyle", 0), // property LineStyle
			/* 41 */ imports.NewTable("TLazVirtualDrawTree_Margin", 0), // property Margin
			/* 42 */ imports.NewTable("TLazVirtualDrawTree_NodeAlignment", 0), // property NodeAlignment
			/* 43 */ imports.NewTable("TLazVirtualDrawTree_NodeDataSize", 0), // property NodeDataSize
			/* 44 */ imports.NewTable("TLazVirtualDrawTree_OperationCanceled", 0), // property OperationCanceled
			/* 45 */ imports.NewTable("TLazVirtualDrawTree_ParentColor", 0), // property ParentColor
			/* 46 */ imports.NewTable("TLazVirtualDrawTree_ParentFont", 0), // property ParentFont
			/* 47 */ imports.NewTable("TLazVirtualDrawTree_ParentShowHint", 0), // property ParentShowHint
			/* 48 */ imports.NewTable("TLazVirtualDrawTree_RootNodeCount", 0), // property RootNodeCount
			/* 49 */ imports.NewTable("TLazVirtualDrawTree_ScrollBarOptions", 0), // property ScrollBarOptions
			/* 50 */ imports.NewTable("TLazVirtualDrawTree_SelectionBlendFactor", 0), // property SelectionBlendFactor
			/* 51 */ imports.NewTable("TLazVirtualDrawTree_SelectionCurveRadius", 0), // property SelectionCurveRadius
			/* 52 */ imports.NewTable("TLazVirtualDrawTree_StateImages", 0), // property StateImages
			/* 53 */ imports.NewTable("TLazVirtualDrawTree_TextMargin", 0), // property TextMargin
			/* 54 */ imports.NewTable("TLazVirtualDrawTree_TreeOptions", 0), // property TreeOptions
			/* 55 */ imports.NewTable("TLazVirtualDrawTree_WantTabs", 0), // property WantTabs
			/* 56 */ imports.NewTable("TLazVirtualDrawTree_OnAddToSelection", 0), // event OnAddToSelection
			/* 57 */ imports.NewTable("TLazVirtualDrawTree_OnAdvancedHeaderDraw", 0), // event OnAdvancedHeaderDraw
			/* 58 */ imports.NewTable("TLazVirtualDrawTree_OnAfterAutoFitColumn", 0), // event OnAfterAutoFitColumn
			/* 59 */ imports.NewTable("TLazVirtualDrawTree_OnAfterAutoFitColumns", 0), // event OnAfterAutoFitColumns
			/* 60 */ imports.NewTable("TLazVirtualDrawTree_OnAfterCellPaint", 0), // event OnAfterCellPaint
			/* 61 */ imports.NewTable("TLazVirtualDrawTree_OnAfterColumnExport", 0), // event OnAfterColumnExport
			/* 62 */ imports.NewTable("TLazVirtualDrawTree_OnAfterColumnWidthTracking", 0), // event OnAfterColumnWidthTracking
			/* 63 */ imports.NewTable("TLazVirtualDrawTree_OnAfterGetMaxColumnWidth", 0), // event OnAfterGetMaxColumnWidth
			/* 64 */ imports.NewTable("TLazVirtualDrawTree_OnAfterHeaderExport", 0), // event OnAfterHeaderExport
			/* 65 */ imports.NewTable("TLazVirtualDrawTree_OnAfterHeaderHeightTracking", 0), // event OnAfterHeaderHeightTracking
			/* 66 */ imports.NewTable("TLazVirtualDrawTree_OnAfterItemErase", 0), // event OnAfterItemErase
			/* 67 */ imports.NewTable("TLazVirtualDrawTree_OnAfterItemPaint", 0), // event OnAfterItemPaint
			/* 68 */ imports.NewTable("TLazVirtualDrawTree_OnAfterNodeExport", 0), // event OnAfterNodeExport
			/* 69 */ imports.NewTable("TLazVirtualDrawTree_OnAfterPaint", 0), // event OnAfterPaint
			/* 70 */ imports.NewTable("TLazVirtualDrawTree_OnAfterTreeExport", 0), // event OnAfterTreeExport
			/* 71 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeAutoFitColumn", 0), // event OnBeforeAutoFitColumn
			/* 72 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeAutoFitColumns", 0), // event OnBeforeAutoFitColumns
			/* 73 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeCellPaint", 0), // event OnBeforeCellPaint
			/* 74 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeColumnExport", 0), // event OnBeforeColumnExport
			/* 75 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeColumnWidthTracking", 0), // event OnBeforeColumnWidthTracking
			/* 76 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeDrawTreeLine", 0), // event OnBeforeDrawTreeLine
			/* 77 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeGetMaxColumnWidth", 0), // event OnBeforeGetMaxColumnWidth
			/* 78 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeHeaderExport", 0), // event OnBeforeHeaderExport
			/* 79 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeHeaderHeightTracking", 0), // event OnBeforeHeaderHeightTracking
			/* 80 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeItemErase", 0), // event OnBeforeItemErase
			/* 81 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeItemPaint", 0), // event OnBeforeItemPaint
			/* 82 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeNodeExport", 0), // event OnBeforeNodeExport
			/* 83 */ imports.NewTable("TLazVirtualDrawTree_OnBeforePaint", 0), // event OnBeforePaint
			/* 84 */ imports.NewTable("TLazVirtualDrawTree_OnBeforeTreeExport", 0), // event OnBeforeTreeExport
			/* 85 */ imports.NewTable("TLazVirtualDrawTree_OnCanSplitterResizeColumn", 0), // event OnCanSplitterResizeColumn
			/* 86 */ imports.NewTable("TLazVirtualDrawTree_OnCanSplitterResizeHeader", 0), // event OnCanSplitterResizeHeader
			/* 87 */ imports.NewTable("TLazVirtualDrawTree_OnCanSplitterResizeNode", 0), // event OnCanSplitterResizeNode
			/* 88 */ imports.NewTable("TLazVirtualDrawTree_OnChange", 0), // event OnChange
			/* 89 */ imports.NewTable("TLazVirtualDrawTree_OnChecked", 0), // event OnChecked
			/* 90 */ imports.NewTable("TLazVirtualDrawTree_OnChecking", 0), // event OnChecking
			/* 91 */ imports.NewTable("TLazVirtualDrawTree_OnCollapsed", 0), // event OnCollapsed
			/* 92 */ imports.NewTable("TLazVirtualDrawTree_OnCollapsing", 0), // event OnCollapsing
			/* 93 */ imports.NewTable("TLazVirtualDrawTree_OnColumnClick", 0), // event OnColumnClick
			/* 94 */ imports.NewTable("TLazVirtualDrawTree_OnColumnDblClick", 0), // event OnColumnDblClick
			/* 95 */ imports.NewTable("TLazVirtualDrawTree_OnColumnExport", 0), // event OnColumnExport
			/* 96 */ imports.NewTable("TLazVirtualDrawTree_OnColumnResize", 0), // event OnColumnResize
			/* 97 */ imports.NewTable("TLazVirtualDrawTree_OnColumnWidthDblClickResize", 0), // event OnColumnWidthDblClickResize
			/* 98 */ imports.NewTable("TLazVirtualDrawTree_OnColumnWidthTracking", 0), // event OnColumnWidthTracking
			/* 99 */ imports.NewTable("TLazVirtualDrawTree_OnCompareNodes", 0), // event OnCompareNodes
			/* 100 */ imports.NewTable("TLazVirtualDrawTree_OnContextPopup", 0), // event OnContextPopup
			/* 101 */ imports.NewTable("TLazVirtualDrawTree_OnCreateDataObject", 0), // event OnCreateDataObject
			/* 102 */ imports.NewTable("TLazVirtualDrawTree_OnCreateDragManager", 0), // event OnCreateDragManager
			/* 103 */ imports.NewTable("TLazVirtualDrawTree_OnCreateEditor", 0), // event OnCreateEditor
			/* 104 */ imports.NewTable("TLazVirtualDrawTree_OnDblClick", 0), // event OnDblClick
			/* 105 */ imports.NewTable("TLazVirtualDrawTree_OnDragAllowed", 0), // event OnDragAllowed
			/* 106 */ imports.NewTable("TLazVirtualDrawTree_OnDragOver", 0), // event OnDragOver
			/* 107 */ imports.NewTable("TLazVirtualDrawTree_OnDragDrop", 0), // event OnDragDrop
			/* 108 */ imports.NewTable("TLazVirtualDrawTree_OnDrawHint", 0), // event OnDrawHint
			/* 109 */ imports.NewTable("TLazVirtualDrawTree_OnDrawNode", 0), // event OnDrawNode
			/* 110 */ imports.NewTable("TLazVirtualDrawTree_OnEdited", 0), // event OnEdited
			/* 111 */ imports.NewTable("TLazVirtualDrawTree_OnEditing", 0), // event OnEditing
			/* 112 */ imports.NewTable("TLazVirtualDrawTree_OnEndDock", 0), // event OnEndDock
			/* 113 */ imports.NewTable("TLazVirtualDrawTree_OnEndDrag", 0), // event OnEndDrag
			/* 114 */ imports.NewTable("TLazVirtualDrawTree_OnEndOperation", 0), // event OnEndOperation
			/* 115 */ imports.NewTable("TLazVirtualDrawTree_OnExpanded", 0), // event OnExpanded
			/* 116 */ imports.NewTable("TLazVirtualDrawTree_OnExpanding", 0), // event OnExpanding
			/* 117 */ imports.NewTable("TLazVirtualDrawTree_OnFocusChanged", 0), // event OnFocusChanged
			/* 118 */ imports.NewTable("TLazVirtualDrawTree_OnFocusChanging", 0), // event OnFocusChanging
			/* 119 */ imports.NewTable("TLazVirtualDrawTree_OnFreeNode", 0), // event OnFreeNode
			/* 120 */ imports.NewTable("TLazVirtualDrawTree_OnGetCellIsEmpty", 0), // event OnGetCellIsEmpty
			/* 121 */ imports.NewTable("TLazVirtualDrawTree_OnGetCursor", 0), // event OnGetCursor
			/* 122 */ imports.NewTable("TLazVirtualDrawTree_OnGetHeaderCursor", 0), // event OnGetHeaderCursor
			/* 123 */ imports.NewTable("TLazVirtualDrawTree_OnGetHelpContext", 0), // event OnGetHelpContext
			/* 124 */ imports.NewTable("TLazVirtualDrawTree_OnGetHint", 0), // event OnGetHint
			/* 125 */ imports.NewTable("TLazVirtualDrawTree_OnGetHintKind", 0), // event OnGetHintKind
			/* 126 */ imports.NewTable("TLazVirtualDrawTree_OnGetHintSize", 0), // event OnGetHintSize
			/* 127 */ imports.NewTable("TLazVirtualDrawTree_OnGetImageIndex", 0), // event OnGetImageIndex
			/* 128 */ imports.NewTable("TLazVirtualDrawTree_OnGetImageIndexEx", 0), // event OnGetImageIndexEx
			/* 129 */ imports.NewTable("TLazVirtualDrawTree_OnGetLineStyle", 0), // event OnGetLineStyle
			/* 130 */ imports.NewTable("TLazVirtualDrawTree_OnGetNodeDataSize", 0), // event OnGetNodeDataSize
			/* 131 */ imports.NewTable("TLazVirtualDrawTree_OnGetNodeWidth", 0), // event OnGetNodeWidth
			/* 132 */ imports.NewTable("TLazVirtualDrawTree_OnGetPopupMenu", 0), // event OnGetPopupMenu
			/* 133 */ imports.NewTable("TLazVirtualDrawTree_OnGetUserClipboardFormats", 0), // event OnGetUserClipboardFormats
			/* 134 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderClick", 0), // event OnHeaderClick
			/* 135 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderDblClick", 0), // event OnHeaderDblClick
			/* 136 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderDragged", 0), // event OnHeaderDragged
			/* 137 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderDraggedOut", 0), // event OnHeaderDraggedOut
			/* 138 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderDragging", 0), // event OnHeaderDragging
			/* 139 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderDraw", 0), // event OnHeaderDraw
			/* 140 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderDrawQueryElements", 0), // event OnHeaderDrawQueryElements
			/* 141 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderHeightTracking", 0), // event OnHeaderHeightTracking
			/* 142 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderHeightDblClickResize", 0), // event OnHeaderHeightDblClickResize
			/* 143 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderMouseDown", 0), // event OnHeaderMouseDown
			/* 144 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderMouseMove", 0), // event OnHeaderMouseMove
			/* 145 */ imports.NewTable("TLazVirtualDrawTree_OnHeaderMouseUp", 0), // event OnHeaderMouseUp
			/* 146 */ imports.NewTable("TLazVirtualDrawTree_OnHotChange", 0), // event OnHotChange
			/* 147 */ imports.NewTable("TLazVirtualDrawTree_OnIncrementalSearch", 0), // event OnIncrementalSearch
			/* 148 */ imports.NewTable("TLazVirtualDrawTree_OnInitChildren", 0), // event OnInitChildren
			/* 149 */ imports.NewTable("TLazVirtualDrawTree_OnInitNode", 0), // event OnInitNode
			/* 150 */ imports.NewTable("TLazVirtualDrawTree_OnKeyAction", 0), // event OnKeyAction
			/* 151 */ imports.NewTable("TLazVirtualDrawTree_OnLoadNode", 0), // event OnLoadNode
			/* 152 */ imports.NewTable("TLazVirtualDrawTree_OnLoadTree", 0), // event OnLoadTree
			/* 153 */ imports.NewTable("TLazVirtualDrawTree_OnMeasureItem", 0), // event OnMeasureItem
			/* 154 */ imports.NewTable("TLazVirtualDrawTree_OnMouseDown", 0), // event OnMouseDown
			/* 155 */ imports.NewTable("TLazVirtualDrawTree_OnMouseMove", 0), // event OnMouseMove
			/* 156 */ imports.NewTable("TLazVirtualDrawTree_OnMouseUp", 0), // event OnMouseUp
			/* 157 */ imports.NewTable("TLazVirtualDrawTree_OnMouseWheel", 0), // event OnMouseWheel
			/* 158 */ imports.NewTable("TLazVirtualDrawTree_OnNodeClick", 0), // event OnNodeClick
			/* 159 */ imports.NewTable("TLazVirtualDrawTree_OnNodeCopied", 0), // event OnNodeCopied
			/* 160 */ imports.NewTable("TLazVirtualDrawTree_OnNodeCopying", 0), // event OnNodeCopying
			/* 161 */ imports.NewTable("TLazVirtualDrawTree_OnNodeDblClick", 0), // event OnNodeDblClick
			/* 162 */ imports.NewTable("TLazVirtualDrawTree_OnNodeExport", 0), // event OnNodeExport
			/* 163 */ imports.NewTable("TLazVirtualDrawTree_OnNodeHeightTracking", 0), // event OnNodeHeightTracking
			/* 164 */ imports.NewTable("TLazVirtualDrawTree_OnNodeHeightDblClickResize", 0), // event OnNodeHeightDblClickResize
			/* 165 */ imports.NewTable("TLazVirtualDrawTree_OnNodeMoved", 0), // event OnNodeMoved
			/* 166 */ imports.NewTable("TLazVirtualDrawTree_OnNodeMoving", 0), // event OnNodeMoving
			/* 167 */ imports.NewTable("TLazVirtualDrawTree_OnPaintBackground", 0), // event OnPaintBackground
			/* 168 */ imports.NewTable("TLazVirtualDrawTree_OnRemoveFromSelection", 0), // event OnRemoveFromSelection
			/* 169 */ imports.NewTable("TLazVirtualDrawTree_OnResetNode", 0), // event OnResetNode
			/* 170 */ imports.NewTable("TLazVirtualDrawTree_OnSaveNode", 0), // event OnSaveNode
			/* 171 */ imports.NewTable("TLazVirtualDrawTree_OnSaveTree", 0), // event OnSaveTree
			/* 172 */ imports.NewTable("TLazVirtualDrawTree_OnScroll", 0), // event OnScroll
			/* 173 */ imports.NewTable("TLazVirtualDrawTree_OnShowScrollBar", 0), // event OnShowScrollBar
			/* 174 */ imports.NewTable("TLazVirtualDrawTree_OnStartDock", 0), // event OnStartDock
			/* 175 */ imports.NewTable("TLazVirtualDrawTree_OnStartDrag", 0), // event OnStartDrag
			/* 176 */ imports.NewTable("TLazVirtualDrawTree_OnStartOperation", 0), // event OnStartOperation
			/* 177 */ imports.NewTable("TLazVirtualDrawTree_OnStateChange", 0), // event OnStateChange
			/* 178 */ imports.NewTable("TLazVirtualDrawTree_OnStructureChange", 0), // event OnStructureChange
			/* 179 */ imports.NewTable("TLazVirtualDrawTree_OnUpdating", 0), // event OnUpdating
			/* 180 */ imports.NewTable("TLazVirtualDrawTree_TClass", 0), // function TLazVirtualDrawTreeClass
		}
	})
	return lazVirtualDrawTreeImport
}
