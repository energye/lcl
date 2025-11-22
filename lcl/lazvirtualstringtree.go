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

// ILazVirtualStringTree Parent: ICustomVirtualStringTree
type ILazVirtualStringTree interface {
	ICustomVirtualStringTree
	RangeX() uint32                                        // property RangeX Getter
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
	DefaultText() string                                                   // property DefaultText Getter
	SetDefaultText(value string)                                           // property DefaultText Setter
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
	TreeOptions() IStringTreeOptions                                       // property TreeOptions Getter
	SetTreeOptions(value IStringTreeOptions)                               // property TreeOptions Setter
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
	SetOnDrawText(fn TVTDrawTextEvent)                                     // property event
	SetOnEditCancelled(fn TVTEditCancelEvent)                              // property event
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
	SetOnGetText(fn TVSTGetTextEvent)                                      // property event
	SetOnPaintText(fn TVTPaintText)                                        // property event
	SetOnGetHelpContext(fn TVTHelpContextEvent)                            // property event
	SetOnGetHintKind(fn TVTHintKindEvent)                                  // property event
	SetOnGetHintSize(fn TVTGetHintSizeEvent)                               // property event
	SetOnGetImageIndex(fn TVTGetImageEvent)                                // property event
	SetOnGetImageIndexEx(fn TVTGetImageExEvent)                            // property event
	SetOnGetImageText(fn TVTGetImageTextEvent)                             // property event
	SetOnGetHint(fn TVTGetHintEvent)                                       // property event
	SetOnGetLineStyle(fn TVTGetLineStyleEvent)                             // property event
	SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent)                       // property event
	SetOnGetPopupMenu(fn TVTPopupEvent)                                    // property event
	SetOnGetUserClipboardFormats(fn TVTGetUserClipboardFormatsEvent)       // property event
	SetOnHeaderClick(fn TVTHeaderClickEvent)                               // property event
	SetOnHeaderDblClick(fn TVTHeaderClickEvent)                            // property event
	SetOnHeaderDragged(fn TVTHeaderDraggedEvent)                           // property event
	SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent)                     // property event
	SetOnHeaderDragging(fn TVTHeaderDraggingEvent)                         // property event
	SetOnHeaderDraw(fn TVTHeaderPaintEvent)                                // property event
	SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent)      // property event
	SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) // property event
	SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent)             // property event
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
	SetOnMeasureTextWidth(fn TVTMeasureTextEvent)                          // property event
	SetOnMeasureTextHeight(fn TVTMeasureTextEvent)                         // property event
	SetOnMouseDown(fn TMouseEvent)                                         // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                     // property event
	SetOnMouseUp(fn TMouseEvent)                                           // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                   // property event
	SetOnMouseEnter(fn TNotifyEvent)                                       // property event
	SetOnMouseLeave(fn TNotifyEvent)                                       // property event
	SetOnNewText(fn TVSTNewTextEvent)                                      // property event
	SetOnNodeClick(fn TVTNodeClickEvent)                                   // property event
	SetOnNodeCopied(fn TVTNodeCopiedEvent)                                 // property event
	SetOnNodeCopying(fn TVTNodeCopyingEvent)                               // property event
	SetOnNodeDblClick(fn TVTNodeClickEvent)                                // property event
	SetOnNodeExport(fn TVTNodeExportEvent)                                 // property event
	SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent)     // property event
	SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent)                 // property event
	SetOnNodeMoved(fn TVTNodeMovedEvent)                                   // property event
	SetOnNodeMoving(fn TVTNodeMovingEvent)                                 // property event
	SetOnPaintBackground(fn TVTBackgroundPaintEvent)                       // property event
	SetOnRemoveFromSelection(fn TVTRemoveFromSelectionEvent)               // property event
	SetOnResetNode(fn TVTChangeEvent)                                      // property event
	SetOnSaveNode(fn TVTSaveNodeEvent)                                     // property event
	SetOnSaveTree(fn TVTSaveTreeEvent)                                     // property event
	SetOnScroll(fn TVTScrollEvent)                                         // property event
	SetOnShortenString(fn TVSTShortenStringEvent)                          // property event
	SetOnShowScrollBar(fn TVTScrollBarShowEvent)                           // property event
	SetOnStartDock(fn TStartDockEvent)                                     // property event
	SetOnStartDrag(fn TStartDragEvent)                                     // property event
	SetOnStartOperation(fn TVTOperationEvent)                              // property event
	SetOnStateChange(fn TVTStateChangeEvent)                               // property event
	SetOnStructureChange(fn TVTStructureChangeEvent)                       // property event
	SetOnUpdating(fn TVTUpdatingEvent)                                     // property event
}

type TLazVirtualStringTree struct {
	TCustomVirtualStringTree
}

func (m *TLazVirtualStringTree) RangeX() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(1, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) LastDragEffect() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(2, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(3, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TLazVirtualStringTree) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) AnimationDuration() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(4, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetAnimationDuration(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) AutoExpandDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(5, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetAutoExpandDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) AutoScrollDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(6, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetAutoScrollDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) AutoScrollInterval() types.TAutoScrollInterval {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(7, 0, m.Instance())
	return types.TAutoScrollInterval(r)
}

func (m *TLazVirtualStringTree) SetAutoScrollInterval(value types.TAutoScrollInterval) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) Background() IPicture {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualStringTreeAPI().SysCallN(8, 0, m.Instance())
	return AsPicture(r)
}

func (m *TLazVirtualStringTree) SetBackground(value IPicture) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualStringTree) BackgroundOffsetX() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualStringTree) SetBackgroundOffsetX(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) BackgroundOffsetY() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualStringTree) SetBackgroundOffsetY(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) BottomSpace() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(11, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetBottomSpace(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) ButtonFillMode() types.TVTButtonFillMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(12, 0, m.Instance())
	return types.TVTButtonFillMode(r)
}

func (m *TLazVirtualStringTree) SetButtonFillMode(value types.TVTButtonFillMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) ButtonStyle() types.TVTButtonStyle {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(13, 0, m.Instance())
	return types.TVTButtonStyle(r)
}

func (m *TLazVirtualStringTree) SetButtonStyle(value types.TVTButtonStyle) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) ChangeDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(14, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetChangeDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) CheckImageKind() types.TCheckImageKind {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(15, 0, m.Instance())
	return types.TCheckImageKind(r)
}

func (m *TLazVirtualStringTree) SetCheckImageKind(value types.TCheckImageKind) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) ClipboardFormats() IClipboardFormats {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualStringTreeAPI().SysCallN(16, 0, m.Instance())
	return AsClipboardFormats(r)
}

func (m *TLazVirtualStringTree) SetClipboardFormats(value IClipboardFormats) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualStringTree) Colors() IVTColors {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualStringTreeAPI().SysCallN(17, 0, m.Instance())
	return AsVTColors(r)
}

func (m *TLazVirtualStringTree) SetColors(value IVTColors) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(17, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualStringTree) CustomCheckImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualStringTreeAPI().SysCallN(18, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TLazVirtualStringTree) SetCustomCheckImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(18, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualStringTree) DefaultNodeHeight() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(19, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetDefaultNodeHeight(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DefaultPasteMode() types.TVTNodeAttachMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(20, 0, m.Instance())
	return types.TVTNodeAttachMode(r)
}

func (m *TLazVirtualStringTree) SetDefaultPasteMode(value types.TVTNodeAttachMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DefaultText() string {
	if !m.IsValid() {
		return ""
	}
	r := lazVirtualStringTreeAPI().SysCallN(21, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TLazVirtualStringTree) SetDefaultText(value string) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(21, 1, m.Instance(), api.PasStr(value))
}

func (m *TLazVirtualStringTree) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(22, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TLazVirtualStringTree) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DragHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualStringTree) SetDragHeight(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(24, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TLazVirtualStringTree) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DragImageKind() types.TVTDragImageKind {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(25, 0, m.Instance())
	return types.TVTDragImageKind(r)
}

func (m *TLazVirtualStringTree) SetDragImageKind(value types.TVTDragImageKind) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(26, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TLazVirtualStringTree) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DragOperations() types.TDragOperations {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(27, 0, m.Instance())
	return types.TDragOperations(r)
}

func (m *TLazVirtualStringTree) SetDragOperations(value types.TDragOperations) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DragType() types.TVTDragType {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(28, 0, m.Instance())
	return types.TVTDragType(r)
}

func (m *TLazVirtualStringTree) SetDragType(value types.TVTDragType) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DragWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(29, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualStringTree) SetDragWidth(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) DrawSelectionMode() types.TVTDrawSelectionMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(30, 0, m.Instance())
	return types.TVTDrawSelectionMode(r)
}

func (m *TLazVirtualStringTree) SetDrawSelectionMode(value types.TVTDrawSelectionMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) EditDelay() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(31, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetEditDelay(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) Header() IVTHeader {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualStringTreeAPI().SysCallN(32, 0, m.Instance())
	return AsVTHeader(r)
}

func (m *TLazVirtualStringTree) SetHeader(value IVTHeader) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(32, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualStringTree) HintMode() types.TVTHintMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(33, 0, m.Instance())
	return types.TVTHintMode(r)
}

func (m *TLazVirtualStringTree) SetHintMode(value types.TVTHintMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(33, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) HotCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(34, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TLazVirtualStringTree) SetHotCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(34, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualStringTreeAPI().SysCallN(35, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TLazVirtualStringTree) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(35, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualStringTree) IncrementalSearch() types.TVTIncrementalSearch {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(36, 0, m.Instance())
	return types.TVTIncrementalSearch(r)
}

func (m *TLazVirtualStringTree) SetIncrementalSearch(value types.TVTIncrementalSearch) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) IncrementalSearchDirection() types.TVTSearchDirection {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(37, 0, m.Instance())
	return types.TVTSearchDirection(r)
}

func (m *TLazVirtualStringTree) SetIncrementalSearchDirection(value types.TVTSearchDirection) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(37, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) IncrementalSearchStart() types.TVTSearchStart {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(38, 0, m.Instance())
	return types.TVTSearchStart(r)
}

func (m *TLazVirtualStringTree) SetIncrementalSearchStart(value types.TVTSearchStart) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(38, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) IncrementalSearchTimeout() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(39, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetIncrementalSearchTimeout(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(39, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) Indent() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(40, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetIndent(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(40, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) LineMode() types.TVTLineMode {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(41, 0, m.Instance())
	return types.TVTLineMode(r)
}

func (m *TLazVirtualStringTree) SetLineMode(value types.TVTLineMode) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(41, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) LineStyle() types.TVTLineStyle {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(42, 0, m.Instance())
	return types.TVTLineStyle(r)
}

func (m *TLazVirtualStringTree) SetLineStyle(value types.TVTLineStyle) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(42, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) Margin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(43, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualStringTree) SetMargin(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(43, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) NodeAlignment() types.TVTNodeAlignment {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(44, 0, m.Instance())
	return types.TVTNodeAlignment(r)
}

func (m *TLazVirtualStringTree) SetNodeAlignment(value types.TVTNodeAlignment) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(44, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) NodeDataSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(45, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualStringTree) SetNodeDataSize(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(45, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) OperationCanceled() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualStringTreeAPI().SysCallN(46, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualStringTree) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualStringTreeAPI().SysCallN(47, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualStringTree) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(47, 1, m.Instance(), api.PasBool(value))
}

func (m *TLazVirtualStringTree) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualStringTreeAPI().SysCallN(48, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualStringTree) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(48, 1, m.Instance(), api.PasBool(value))
}

func (m *TLazVirtualStringTree) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualStringTreeAPI().SysCallN(49, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualStringTree) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(49, 1, m.Instance(), api.PasBool(value))
}

func (m *TLazVirtualStringTree) RootNodeCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(50, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetRootNodeCount(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(50, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) ScrollBarOptions() IScrollBarOptions {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualStringTreeAPI().SysCallN(51, 0, m.Instance())
	return AsScrollBarOptions(r)
}

func (m *TLazVirtualStringTree) SetScrollBarOptions(value IScrollBarOptions) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(51, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualStringTree) SelectionBlendFactor() byte {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(52, 0, m.Instance())
	return byte(r)
}

func (m *TLazVirtualStringTree) SetSelectionBlendFactor(value byte) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(52, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) SelectionCurveRadius() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(53, 0, m.Instance())
	return uint32(r)
}

func (m *TLazVirtualStringTree) SetSelectionCurveRadius(value uint32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(53, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) StateImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualStringTreeAPI().SysCallN(54, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TLazVirtualStringTree) SetStateImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(54, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualStringTree) TextMargin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazVirtualStringTreeAPI().SysCallN(55, 0, m.Instance())
	return int32(r)
}

func (m *TLazVirtualStringTree) SetTextMargin(value int32) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(55, 1, m.Instance(), uintptr(value))
}

func (m *TLazVirtualStringTree) TreeOptions() IStringTreeOptions {
	if !m.IsValid() {
		return nil
	}
	r := lazVirtualStringTreeAPI().SysCallN(56, 0, m.Instance())
	return AsStringTreeOptions(r)
}

func (m *TLazVirtualStringTree) SetTreeOptions(value IStringTreeOptions) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(56, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazVirtualStringTree) WantTabs() bool {
	if !m.IsValid() {
		return false
	}
	r := lazVirtualStringTreeAPI().SysCallN(57, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazVirtualStringTree) SetWantTabs(value bool) {
	if !m.IsValid() {
		return
	}
	lazVirtualStringTreeAPI().SysCallN(57, 1, m.Instance(), api.PasBool(value))
}

func (m *TLazVirtualStringTree) SetOnAddToSelection(fn TVTAddToSelectionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAddToSelectionEvent(fn)
	base.SetEvent(m, 58, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAdvancedHeaderDraw(fn TVTAdvancedHeaderPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAdvancedHeaderPaintEvent(fn)
	base.SetEvent(m, 59, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterAutoFitColumn(fn TVTAfterAutoFitColumnEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterAutoFitColumnEvent(fn)
	base.SetEvent(m, 60, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterAutoFitColumns(fn TVTAfterAutoFitColumnsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterAutoFitColumnsEvent(fn)
	base.SetEvent(m, 61, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterCellPaint(fn TVTAfterCellPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterCellPaintEvent(fn)
	base.SetEvent(m, 62, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterColumnExport(fn TVTColumnExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnExportEvent(fn)
	base.SetEvent(m, 63, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterColumnWidthTracking(fn TVTAfterColumnWidthTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterColumnWidthTrackingEvent(fn)
	base.SetEvent(m, 64, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterGetMaxColumnWidth(fn TVTAfterGetMaxColumnWidthEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterGetMaxColumnWidthEvent(fn)
	base.SetEvent(m, 65, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterHeaderExport(fn TVTTreeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTTreeExportEvent(fn)
	base.SetEvent(m, 66, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterHeaderHeightTracking(fn TVTAfterHeaderHeightTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterHeaderHeightTrackingEvent(fn)
	base.SetEvent(m, 67, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterItemErase(fn TVTAfterItemEraseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterItemEraseEvent(fn)
	base.SetEvent(m, 68, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterItemPaint(fn TVTAfterItemPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTAfterItemPaintEvent(fn)
	base.SetEvent(m, 69, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterNodeExport(fn TVTNodeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeExportEvent(fn)
	base.SetEvent(m, 70, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterPaint(fn TVTPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTPaintEvent(fn)
	base.SetEvent(m, 71, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnAfterTreeExport(fn TVTTreeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTTreeExportEvent(fn)
	base.SetEvent(m, 72, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeAutoFitColumn(fn TVTBeforeAutoFitColumnEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeAutoFitColumnEvent(fn)
	base.SetEvent(m, 73, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeAutoFitColumns(fn TVTBeforeAutoFitColumnsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeAutoFitColumnsEvent(fn)
	base.SetEvent(m, 74, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeCellPaint(fn TVTBeforeCellPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeCellPaintEvent(fn)
	base.SetEvent(m, 75, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeColumnExport(fn TVTColumnExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnExportEvent(fn)
	base.SetEvent(m, 76, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeColumnWidthTracking(fn TVTBeforeColumnWidthTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeColumnWidthTrackingEvent(fn)
	base.SetEvent(m, 77, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeDrawTreeLine(fn TVTBeforeDrawLineImageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeDrawLineImageEvent(fn)
	base.SetEvent(m, 78, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeGetMaxColumnWidth(fn TVTBeforeGetMaxColumnWidthEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeGetMaxColumnWidthEvent(fn)
	base.SetEvent(m, 79, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeHeaderExport(fn TVTTreeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTTreeExportEvent(fn)
	base.SetEvent(m, 80, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeHeaderHeightTracking(fn TVTBeforeHeaderHeightTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeHeaderHeightTrackingEvent(fn)
	base.SetEvent(m, 81, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeItemErase(fn TVTBeforeItemEraseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeItemEraseEvent(fn)
	base.SetEvent(m, 82, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeItemPaint(fn TVTBeforeItemPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBeforeItemPaintEvent(fn)
	base.SetEvent(m, 83, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeNodeExport(fn TVTNodeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeExportEvent(fn)
	base.SetEvent(m, 84, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforePaint(fn TVTPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTPaintEvent(fn)
	base.SetEvent(m, 85, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnBeforeTreeExport(fn TVTTreeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTTreeExportEvent(fn)
	base.SetEvent(m, 86, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnCanSplitterResizeColumn(fn TVTCanSplitterResizeColumnEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCanSplitterResizeColumnEvent(fn)
	base.SetEvent(m, 87, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnCanSplitterResizeHeader(fn TVTCanSplitterResizeHeaderEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCanSplitterResizeHeaderEvent(fn)
	base.SetEvent(m, 88, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnCanSplitterResizeNode(fn TVTCanSplitterResizeNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCanSplitterResizeNodeEvent(fn)
	base.SetEvent(m, 89, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnChange(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 90, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnChecked(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 91, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnChecking(fn TVTCheckChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCheckChangingEvent(fn)
	base.SetEvent(m, 92, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnCollapsed(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 93, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnCollapsing(fn TVTChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangingEvent(fn)
	base.SetEvent(m, 94, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnColumnClick(fn TVTColumnClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnClickEvent(fn)
	base.SetEvent(m, 95, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnColumnDblClick(fn TVTColumnDblClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnDblClickEvent(fn)
	base.SetEvent(m, 96, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnColumnExport(fn TVTColumnExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnExportEvent(fn)
	base.SetEvent(m, 97, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnColumnResize(fn TVTHeaderNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderNotifyEvent(fn)
	base.SetEvent(m, 98, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnColumnWidthDblClickResize(fn TVTColumnWidthDblClickResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnWidthDblClickResizeEvent(fn)
	base.SetEvent(m, 99, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnColumnWidthTracking(fn TVTColumnWidthTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTColumnWidthTrackingEvent(fn)
	base.SetEvent(m, 100, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnCompareNodes(fn TVTCompareEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCompareEvent(fn)
	base.SetEvent(m, 101, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 102, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnCreateDataObject(fn TVTCreateDataObjectEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCreateDataObjectEvent(fn)
	base.SetEvent(m, 103, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnCreateDragManager(fn TVTCreateDragManagerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCreateDragManagerEvent(fn)
	base.SetEvent(m, 104, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnCreateEditor(fn TVTCreateEditorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTCreateEditorEvent(fn)
	base.SetEvent(m, 105, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 106, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnDragAllowed(fn TVTDragAllowedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDragAllowedEvent(fn)
	base.SetEvent(m, 107, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnDragOver(fn TVTDragOverEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDragOverEvent(fn)
	base.SetEvent(m, 108, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnDragDrop(fn TVTDragDropEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDragDropEvent(fn)
	base.SetEvent(m, 109, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnDrawHint(fn TVTDrawHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDrawHintEvent(fn)
	base.SetEvent(m, 110, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnDrawText(fn TVTDrawTextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTDrawTextEvent(fn)
	base.SetEvent(m, 111, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnEditCancelled(fn TVTEditCancelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTEditCancelEvent(fn)
	base.SetEvent(m, 112, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnEdited(fn TVTEditChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTEditChangeEvent(fn)
	base.SetEvent(m, 113, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnEditing(fn TVTEditChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTEditChangingEvent(fn)
	base.SetEvent(m, 114, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnEndDock(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 115, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnEndDrag(fn TEndDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTEndDragEvent(fn)
	base.SetEvent(m, 116, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnEndOperation(fn TVTOperationEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTOperationEvent(fn)
	base.SetEvent(m, 117, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnExpanded(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 118, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnExpanding(fn TVTChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangingEvent(fn)
	base.SetEvent(m, 119, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnFocusChanged(fn TVTFocusChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTFocusChangeEvent(fn)
	base.SetEvent(m, 120, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnFocusChanging(fn TVTFocusChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTFocusChangingEvent(fn)
	base.SetEvent(m, 121, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnFreeNode(fn TVTFreeNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTFreeNodeEvent(fn)
	base.SetEvent(m, 122, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetCellIsEmpty(fn TVTGetCellIsEmptyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetCellIsEmptyEvent(fn)
	base.SetEvent(m, 123, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetCursor(fn TVTGetCursorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetCursorEvent(fn)
	base.SetEvent(m, 124, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetHeaderCursor(fn TVTGetHeaderCursorEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetHeaderCursorEvent(fn)
	base.SetEvent(m, 125, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetText(fn TVSTGetTextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVSTGetTextEvent(fn)
	base.SetEvent(m, 126, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnPaintText(fn TVTPaintText) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTPaintText(fn)
	base.SetEvent(m, 127, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetHelpContext(fn TVTHelpContextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHelpContextEvent(fn)
	base.SetEvent(m, 128, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetHintKind(fn TVTHintKindEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHintKindEvent(fn)
	base.SetEvent(m, 129, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetHintSize(fn TVTGetHintSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetHintSizeEvent(fn)
	base.SetEvent(m, 130, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetImageIndex(fn TVTGetImageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetImageEvent(fn)
	base.SetEvent(m, 131, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetImageIndexEx(fn TVTGetImageExEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetImageExEvent(fn)
	base.SetEvent(m, 132, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetImageText(fn TVTGetImageTextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetImageTextEvent(fn)
	base.SetEvent(m, 133, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetHint(fn TVTGetHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetHintEvent(fn)
	base.SetEvent(m, 134, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetLineStyle(fn TVTGetLineStyleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetLineStyleEvent(fn)
	base.SetEvent(m, 135, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetNodeDataSizeEvent(fn)
	base.SetEvent(m, 136, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetPopupMenu(fn TVTPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTPopupEvent(fn)
	base.SetEvent(m, 137, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnGetUserClipboardFormats(fn TVTGetUserClipboardFormatsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTGetUserClipboardFormatsEvent(fn)
	base.SetEvent(m, 138, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderClick(fn TVTHeaderClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderClickEvent(fn)
	base.SetEvent(m, 139, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderDblClick(fn TVTHeaderClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderClickEvent(fn)
	base.SetEvent(m, 140, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderDragged(fn TVTHeaderDraggedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderDraggedEvent(fn)
	base.SetEvent(m, 141, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderDraggedOutEvent(fn)
	base.SetEvent(m, 142, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderDragging(fn TVTHeaderDraggingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderDraggingEvent(fn)
	base.SetEvent(m, 143, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderDraw(fn TVTHeaderPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderPaintEvent(fn)
	base.SetEvent(m, 144, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderPaintQueryElementsEvent(fn)
	base.SetEvent(m, 145, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderHeightDblClickResizeEvent(fn)
	base.SetEvent(m, 146, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderHeightTrackingEvent(fn)
	base.SetEvent(m, 147, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderMouseDown(fn TVTHeaderMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderMouseEvent(fn)
	base.SetEvent(m, 148, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderMouseMove(fn TVTHeaderMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderMouseMoveEvent(fn)
	base.SetEvent(m, 149, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHeaderMouseUp(fn TVTHeaderMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHeaderMouseEvent(fn)
	base.SetEvent(m, 150, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnHotChange(fn TVTHotNodeChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTHotNodeChangeEvent(fn)
	base.SetEvent(m, 151, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnIncrementalSearch(fn TVTIncrementalSearchEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTIncrementalSearchEvent(fn)
	base.SetEvent(m, 152, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnInitChildren(fn TVTInitChildrenEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTInitChildrenEvent(fn)
	base.SetEvent(m, 153, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnInitNode(fn TVTInitNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTInitNodeEvent(fn)
	base.SetEvent(m, 154, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnKeyAction(fn TVTKeyActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTKeyActionEvent(fn)
	base.SetEvent(m, 155, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnLoadNode(fn TVTSaveNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTSaveNodeEvent(fn)
	base.SetEvent(m, 156, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnLoadTree(fn TVTSaveTreeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTSaveTreeEvent(fn)
	base.SetEvent(m, 157, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnMeasureItem(fn TVTMeasureItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTMeasureItemEvent(fn)
	base.SetEvent(m, 158, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnMeasureTextWidth(fn TVTMeasureTextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTMeasureTextEvent(fn)
	base.SetEvent(m, 159, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnMeasureTextHeight(fn TVTMeasureTextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTMeasureTextEvent(fn)
	base.SetEvent(m, 160, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 161, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 162, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 163, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 164, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 165, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 166, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNewText(fn TVSTNewTextEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVSTNewTextEvent(fn)
	base.SetEvent(m, 167, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNodeClick(fn TVTNodeClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeClickEvent(fn)
	base.SetEvent(m, 168, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNodeCopied(fn TVTNodeCopiedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeCopiedEvent(fn)
	base.SetEvent(m, 169, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNodeCopying(fn TVTNodeCopyingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeCopyingEvent(fn)
	base.SetEvent(m, 170, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNodeDblClick(fn TVTNodeClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeClickEvent(fn)
	base.SetEvent(m, 171, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNodeExport(fn TVTNodeExportEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeExportEvent(fn)
	base.SetEvent(m, 172, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeHeightDblClickResizeEvent(fn)
	base.SetEvent(m, 173, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeHeightTrackingEvent(fn)
	base.SetEvent(m, 174, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNodeMoved(fn TVTNodeMovedEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeMovedEvent(fn)
	base.SetEvent(m, 175, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnNodeMoving(fn TVTNodeMovingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTNodeMovingEvent(fn)
	base.SetEvent(m, 176, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnPaintBackground(fn TVTBackgroundPaintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTBackgroundPaintEvent(fn)
	base.SetEvent(m, 177, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnRemoveFromSelection(fn TVTRemoveFromSelectionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTRemoveFromSelectionEvent(fn)
	base.SetEvent(m, 178, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnResetNode(fn TVTChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTChangeEvent(fn)
	base.SetEvent(m, 179, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnSaveNode(fn TVTSaveNodeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTSaveNodeEvent(fn)
	base.SetEvent(m, 180, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnSaveTree(fn TVTSaveTreeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTSaveTreeEvent(fn)
	base.SetEvent(m, 181, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnScroll(fn TVTScrollEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTScrollEvent(fn)
	base.SetEvent(m, 182, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnShortenString(fn TVSTShortenStringEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVSTShortenStringEvent(fn)
	base.SetEvent(m, 183, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnShowScrollBar(fn TVTScrollBarShowEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTScrollBarShowEvent(fn)
	base.SetEvent(m, 184, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnStartDock(fn TStartDockEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDockEvent(fn)
	base.SetEvent(m, 185, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnStartDrag(fn TStartDragEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTStartDragEvent(fn)
	base.SetEvent(m, 186, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnStartOperation(fn TVTOperationEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTOperationEvent(fn)
	base.SetEvent(m, 187, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnStateChange(fn TVTStateChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTStateChangeEvent(fn)
	base.SetEvent(m, 188, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnStructureChange(fn TVTStructureChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTStructureChangeEvent(fn)
	base.SetEvent(m, 189, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TLazVirtualStringTree) SetOnUpdating(fn TVTUpdatingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTVTUpdatingEvent(fn)
	base.SetEvent(m, 190, lazVirtualStringTreeAPI(), api.MakeEventDataPtr(cb))
}

// NewLazVirtualStringTree class constructor
func NewLazVirtualStringTree(owner IComponent) ILazVirtualStringTree {
	r := lazVirtualStringTreeAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLazVirtualStringTree(r)
}

func TLazVirtualStringTreeClass() types.TClass {
	r := lazVirtualStringTreeAPI().SysCallN(191)
	return types.TClass(r)
}

var (
	lazVirtualStringTreeOnce   base.Once
	lazVirtualStringTreeImport *imports.Imports = nil
)

func lazVirtualStringTreeAPI() *imports.Imports {
	lazVirtualStringTreeOnce.Do(func() {
		lazVirtualStringTreeImport = api.NewDefaultImports()
		lazVirtualStringTreeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazVirtualStringTree_Create", 0), // constructor NewLazVirtualStringTree
			/* 1 */ imports.NewTable("TLazVirtualStringTree_RangeX", 0), // property RangeX
			/* 2 */ imports.NewTable("TLazVirtualStringTree_LastDragEffect", 0), // property LastDragEffect
			/* 3 */ imports.NewTable("TLazVirtualStringTree_Alignment", 0), // property Alignment
			/* 4 */ imports.NewTable("TLazVirtualStringTree_AnimationDuration", 0), // property AnimationDuration
			/* 5 */ imports.NewTable("TLazVirtualStringTree_AutoExpandDelay", 0), // property AutoExpandDelay
			/* 6 */ imports.NewTable("TLazVirtualStringTree_AutoScrollDelay", 0), // property AutoScrollDelay
			/* 7 */ imports.NewTable("TLazVirtualStringTree_AutoScrollInterval", 0), // property AutoScrollInterval
			/* 8 */ imports.NewTable("TLazVirtualStringTree_Background", 0), // property Background
			/* 9 */ imports.NewTable("TLazVirtualStringTree_BackgroundOffsetX", 0), // property BackgroundOffsetX
			/* 10 */ imports.NewTable("TLazVirtualStringTree_BackgroundOffsetY", 0), // property BackgroundOffsetY
			/* 11 */ imports.NewTable("TLazVirtualStringTree_BottomSpace", 0), // property BottomSpace
			/* 12 */ imports.NewTable("TLazVirtualStringTree_ButtonFillMode", 0), // property ButtonFillMode
			/* 13 */ imports.NewTable("TLazVirtualStringTree_ButtonStyle", 0), // property ButtonStyle
			/* 14 */ imports.NewTable("TLazVirtualStringTree_ChangeDelay", 0), // property ChangeDelay
			/* 15 */ imports.NewTable("TLazVirtualStringTree_CheckImageKind", 0), // property CheckImageKind
			/* 16 */ imports.NewTable("TLazVirtualStringTree_ClipboardFormats", 0), // property ClipboardFormats
			/* 17 */ imports.NewTable("TLazVirtualStringTree_Colors", 0), // property Colors
			/* 18 */ imports.NewTable("TLazVirtualStringTree_CustomCheckImages", 0), // property CustomCheckImages
			/* 19 */ imports.NewTable("TLazVirtualStringTree_DefaultNodeHeight", 0), // property DefaultNodeHeight
			/* 20 */ imports.NewTable("TLazVirtualStringTree_DefaultPasteMode", 0), // property DefaultPasteMode
			/* 21 */ imports.NewTable("TLazVirtualStringTree_DefaultText", 0), // property DefaultText
			/* 22 */ imports.NewTable("TLazVirtualStringTree_DragCursor", 0), // property DragCursor
			/* 23 */ imports.NewTable("TLazVirtualStringTree_DragHeight", 0), // property DragHeight
			/* 24 */ imports.NewTable("TLazVirtualStringTree_DragKind", 0), // property DragKind
			/* 25 */ imports.NewTable("TLazVirtualStringTree_DragImageKind", 0), // property DragImageKind
			/* 26 */ imports.NewTable("TLazVirtualStringTree_DragMode", 0), // property DragMode
			/* 27 */ imports.NewTable("TLazVirtualStringTree_DragOperations", 0), // property DragOperations
			/* 28 */ imports.NewTable("TLazVirtualStringTree_DragType", 0), // property DragType
			/* 29 */ imports.NewTable("TLazVirtualStringTree_DragWidth", 0), // property DragWidth
			/* 30 */ imports.NewTable("TLazVirtualStringTree_DrawSelectionMode", 0), // property DrawSelectionMode
			/* 31 */ imports.NewTable("TLazVirtualStringTree_EditDelay", 0), // property EditDelay
			/* 32 */ imports.NewTable("TLazVirtualStringTree_Header", 0), // property Header
			/* 33 */ imports.NewTable("TLazVirtualStringTree_HintMode", 0), // property HintMode
			/* 34 */ imports.NewTable("TLazVirtualStringTree_HotCursor", 0), // property HotCursor
			/* 35 */ imports.NewTable("TLazVirtualStringTree_Images", 0), // property Images
			/* 36 */ imports.NewTable("TLazVirtualStringTree_IncrementalSearch", 0), // property IncrementalSearch
			/* 37 */ imports.NewTable("TLazVirtualStringTree_IncrementalSearchDirection", 0), // property IncrementalSearchDirection
			/* 38 */ imports.NewTable("TLazVirtualStringTree_IncrementalSearchStart", 0), // property IncrementalSearchStart
			/* 39 */ imports.NewTable("TLazVirtualStringTree_IncrementalSearchTimeout", 0), // property IncrementalSearchTimeout
			/* 40 */ imports.NewTable("TLazVirtualStringTree_Indent", 0), // property Indent
			/* 41 */ imports.NewTable("TLazVirtualStringTree_LineMode", 0), // property LineMode
			/* 42 */ imports.NewTable("TLazVirtualStringTree_LineStyle", 0), // property LineStyle
			/* 43 */ imports.NewTable("TLazVirtualStringTree_Margin", 0), // property Margin
			/* 44 */ imports.NewTable("TLazVirtualStringTree_NodeAlignment", 0), // property NodeAlignment
			/* 45 */ imports.NewTable("TLazVirtualStringTree_NodeDataSize", 0), // property NodeDataSize
			/* 46 */ imports.NewTable("TLazVirtualStringTree_OperationCanceled", 0), // property OperationCanceled
			/* 47 */ imports.NewTable("TLazVirtualStringTree_ParentColor", 0), // property ParentColor
			/* 48 */ imports.NewTable("TLazVirtualStringTree_ParentFont", 0), // property ParentFont
			/* 49 */ imports.NewTable("TLazVirtualStringTree_ParentShowHint", 0), // property ParentShowHint
			/* 50 */ imports.NewTable("TLazVirtualStringTree_RootNodeCount", 0), // property RootNodeCount
			/* 51 */ imports.NewTable("TLazVirtualStringTree_ScrollBarOptions", 0), // property ScrollBarOptions
			/* 52 */ imports.NewTable("TLazVirtualStringTree_SelectionBlendFactor", 0), // property SelectionBlendFactor
			/* 53 */ imports.NewTable("TLazVirtualStringTree_SelectionCurveRadius", 0), // property SelectionCurveRadius
			/* 54 */ imports.NewTable("TLazVirtualStringTree_StateImages", 0), // property StateImages
			/* 55 */ imports.NewTable("TLazVirtualStringTree_TextMargin", 0), // property TextMargin
			/* 56 */ imports.NewTable("TLazVirtualStringTree_TreeOptions", 0), // property TreeOptions
			/* 57 */ imports.NewTable("TLazVirtualStringTree_WantTabs", 0), // property WantTabs
			/* 58 */ imports.NewTable("TLazVirtualStringTree_OnAddToSelection", 0), // event OnAddToSelection
			/* 59 */ imports.NewTable("TLazVirtualStringTree_OnAdvancedHeaderDraw", 0), // event OnAdvancedHeaderDraw
			/* 60 */ imports.NewTable("TLazVirtualStringTree_OnAfterAutoFitColumn", 0), // event OnAfterAutoFitColumn
			/* 61 */ imports.NewTable("TLazVirtualStringTree_OnAfterAutoFitColumns", 0), // event OnAfterAutoFitColumns
			/* 62 */ imports.NewTable("TLazVirtualStringTree_OnAfterCellPaint", 0), // event OnAfterCellPaint
			/* 63 */ imports.NewTable("TLazVirtualStringTree_OnAfterColumnExport", 0), // event OnAfterColumnExport
			/* 64 */ imports.NewTable("TLazVirtualStringTree_OnAfterColumnWidthTracking", 0), // event OnAfterColumnWidthTracking
			/* 65 */ imports.NewTable("TLazVirtualStringTree_OnAfterGetMaxColumnWidth", 0), // event OnAfterGetMaxColumnWidth
			/* 66 */ imports.NewTable("TLazVirtualStringTree_OnAfterHeaderExport", 0), // event OnAfterHeaderExport
			/* 67 */ imports.NewTable("TLazVirtualStringTree_OnAfterHeaderHeightTracking", 0), // event OnAfterHeaderHeightTracking
			/* 68 */ imports.NewTable("TLazVirtualStringTree_OnAfterItemErase", 0), // event OnAfterItemErase
			/* 69 */ imports.NewTable("TLazVirtualStringTree_OnAfterItemPaint", 0), // event OnAfterItemPaint
			/* 70 */ imports.NewTable("TLazVirtualStringTree_OnAfterNodeExport", 0), // event OnAfterNodeExport
			/* 71 */ imports.NewTable("TLazVirtualStringTree_OnAfterPaint", 0), // event OnAfterPaint
			/* 72 */ imports.NewTable("TLazVirtualStringTree_OnAfterTreeExport", 0), // event OnAfterTreeExport
			/* 73 */ imports.NewTable("TLazVirtualStringTree_OnBeforeAutoFitColumn", 0), // event OnBeforeAutoFitColumn
			/* 74 */ imports.NewTable("TLazVirtualStringTree_OnBeforeAutoFitColumns", 0), // event OnBeforeAutoFitColumns
			/* 75 */ imports.NewTable("TLazVirtualStringTree_OnBeforeCellPaint", 0), // event OnBeforeCellPaint
			/* 76 */ imports.NewTable("TLazVirtualStringTree_OnBeforeColumnExport", 0), // event OnBeforeColumnExport
			/* 77 */ imports.NewTable("TLazVirtualStringTree_OnBeforeColumnWidthTracking", 0), // event OnBeforeColumnWidthTracking
			/* 78 */ imports.NewTable("TLazVirtualStringTree_OnBeforeDrawTreeLine", 0), // event OnBeforeDrawTreeLine
			/* 79 */ imports.NewTable("TLazVirtualStringTree_OnBeforeGetMaxColumnWidth", 0), // event OnBeforeGetMaxColumnWidth
			/* 80 */ imports.NewTable("TLazVirtualStringTree_OnBeforeHeaderExport", 0), // event OnBeforeHeaderExport
			/* 81 */ imports.NewTable("TLazVirtualStringTree_OnBeforeHeaderHeightTracking", 0), // event OnBeforeHeaderHeightTracking
			/* 82 */ imports.NewTable("TLazVirtualStringTree_OnBeforeItemErase", 0), // event OnBeforeItemErase
			/* 83 */ imports.NewTable("TLazVirtualStringTree_OnBeforeItemPaint", 0), // event OnBeforeItemPaint
			/* 84 */ imports.NewTable("TLazVirtualStringTree_OnBeforeNodeExport", 0), // event OnBeforeNodeExport
			/* 85 */ imports.NewTable("TLazVirtualStringTree_OnBeforePaint", 0), // event OnBeforePaint
			/* 86 */ imports.NewTable("TLazVirtualStringTree_OnBeforeTreeExport", 0), // event OnBeforeTreeExport
			/* 87 */ imports.NewTable("TLazVirtualStringTree_OnCanSplitterResizeColumn", 0), // event OnCanSplitterResizeColumn
			/* 88 */ imports.NewTable("TLazVirtualStringTree_OnCanSplitterResizeHeader", 0), // event OnCanSplitterResizeHeader
			/* 89 */ imports.NewTable("TLazVirtualStringTree_OnCanSplitterResizeNode", 0), // event OnCanSplitterResizeNode
			/* 90 */ imports.NewTable("TLazVirtualStringTree_OnChange", 0), // event OnChange
			/* 91 */ imports.NewTable("TLazVirtualStringTree_OnChecked", 0), // event OnChecked
			/* 92 */ imports.NewTable("TLazVirtualStringTree_OnChecking", 0), // event OnChecking
			/* 93 */ imports.NewTable("TLazVirtualStringTree_OnCollapsed", 0), // event OnCollapsed
			/* 94 */ imports.NewTable("TLazVirtualStringTree_OnCollapsing", 0), // event OnCollapsing
			/* 95 */ imports.NewTable("TLazVirtualStringTree_OnColumnClick", 0), // event OnColumnClick
			/* 96 */ imports.NewTable("TLazVirtualStringTree_OnColumnDblClick", 0), // event OnColumnDblClick
			/* 97 */ imports.NewTable("TLazVirtualStringTree_OnColumnExport", 0), // event OnColumnExport
			/* 98 */ imports.NewTable("TLazVirtualStringTree_OnColumnResize", 0), // event OnColumnResize
			/* 99 */ imports.NewTable("TLazVirtualStringTree_OnColumnWidthDblClickResize", 0), // event OnColumnWidthDblClickResize
			/* 100 */ imports.NewTable("TLazVirtualStringTree_OnColumnWidthTracking", 0), // event OnColumnWidthTracking
			/* 101 */ imports.NewTable("TLazVirtualStringTree_OnCompareNodes", 0), // event OnCompareNodes
			/* 102 */ imports.NewTable("TLazVirtualStringTree_OnContextPopup", 0), // event OnContextPopup
			/* 103 */ imports.NewTable("TLazVirtualStringTree_OnCreateDataObject", 0), // event OnCreateDataObject
			/* 104 */ imports.NewTable("TLazVirtualStringTree_OnCreateDragManager", 0), // event OnCreateDragManager
			/* 105 */ imports.NewTable("TLazVirtualStringTree_OnCreateEditor", 0), // event OnCreateEditor
			/* 106 */ imports.NewTable("TLazVirtualStringTree_OnDblClick", 0), // event OnDblClick
			/* 107 */ imports.NewTable("TLazVirtualStringTree_OnDragAllowed", 0), // event OnDragAllowed
			/* 108 */ imports.NewTable("TLazVirtualStringTree_OnDragOver", 0), // event OnDragOver
			/* 109 */ imports.NewTable("TLazVirtualStringTree_OnDragDrop", 0), // event OnDragDrop
			/* 110 */ imports.NewTable("TLazVirtualStringTree_OnDrawHint", 0), // event OnDrawHint
			/* 111 */ imports.NewTable("TLazVirtualStringTree_OnDrawText", 0), // event OnDrawText
			/* 112 */ imports.NewTable("TLazVirtualStringTree_OnEditCancelled", 0), // event OnEditCancelled
			/* 113 */ imports.NewTable("TLazVirtualStringTree_OnEdited", 0), // event OnEdited
			/* 114 */ imports.NewTable("TLazVirtualStringTree_OnEditing", 0), // event OnEditing
			/* 115 */ imports.NewTable("TLazVirtualStringTree_OnEndDock", 0), // event OnEndDock
			/* 116 */ imports.NewTable("TLazVirtualStringTree_OnEndDrag", 0), // event OnEndDrag
			/* 117 */ imports.NewTable("TLazVirtualStringTree_OnEndOperation", 0), // event OnEndOperation
			/* 118 */ imports.NewTable("TLazVirtualStringTree_OnExpanded", 0), // event OnExpanded
			/* 119 */ imports.NewTable("TLazVirtualStringTree_OnExpanding", 0), // event OnExpanding
			/* 120 */ imports.NewTable("TLazVirtualStringTree_OnFocusChanged", 0), // event OnFocusChanged
			/* 121 */ imports.NewTable("TLazVirtualStringTree_OnFocusChanging", 0), // event OnFocusChanging
			/* 122 */ imports.NewTable("TLazVirtualStringTree_OnFreeNode", 0), // event OnFreeNode
			/* 123 */ imports.NewTable("TLazVirtualStringTree_OnGetCellIsEmpty", 0), // event OnGetCellIsEmpty
			/* 124 */ imports.NewTable("TLazVirtualStringTree_OnGetCursor", 0), // event OnGetCursor
			/* 125 */ imports.NewTable("TLazVirtualStringTree_OnGetHeaderCursor", 0), // event OnGetHeaderCursor
			/* 126 */ imports.NewTable("TLazVirtualStringTree_OnGetText", 0), // event OnGetText
			/* 127 */ imports.NewTable("TLazVirtualStringTree_OnPaintText", 0), // event OnPaintText
			/* 128 */ imports.NewTable("TLazVirtualStringTree_OnGetHelpContext", 0), // event OnGetHelpContext
			/* 129 */ imports.NewTable("TLazVirtualStringTree_OnGetHintKind", 0), // event OnGetHintKind
			/* 130 */ imports.NewTable("TLazVirtualStringTree_OnGetHintSize", 0), // event OnGetHintSize
			/* 131 */ imports.NewTable("TLazVirtualStringTree_OnGetImageIndex", 0), // event OnGetImageIndex
			/* 132 */ imports.NewTable("TLazVirtualStringTree_OnGetImageIndexEx", 0), // event OnGetImageIndexEx
			/* 133 */ imports.NewTable("TLazVirtualStringTree_OnGetImageText", 0), // event OnGetImageText
			/* 134 */ imports.NewTable("TLazVirtualStringTree_OnGetHint", 0), // event OnGetHint
			/* 135 */ imports.NewTable("TLazVirtualStringTree_OnGetLineStyle", 0), // event OnGetLineStyle
			/* 136 */ imports.NewTable("TLazVirtualStringTree_OnGetNodeDataSize", 0), // event OnGetNodeDataSize
			/* 137 */ imports.NewTable("TLazVirtualStringTree_OnGetPopupMenu", 0), // event OnGetPopupMenu
			/* 138 */ imports.NewTable("TLazVirtualStringTree_OnGetUserClipboardFormats", 0), // event OnGetUserClipboardFormats
			/* 139 */ imports.NewTable("TLazVirtualStringTree_OnHeaderClick", 0), // event OnHeaderClick
			/* 140 */ imports.NewTable("TLazVirtualStringTree_OnHeaderDblClick", 0), // event OnHeaderDblClick
			/* 141 */ imports.NewTable("TLazVirtualStringTree_OnHeaderDragged", 0), // event OnHeaderDragged
			/* 142 */ imports.NewTable("TLazVirtualStringTree_OnHeaderDraggedOut", 0), // event OnHeaderDraggedOut
			/* 143 */ imports.NewTable("TLazVirtualStringTree_OnHeaderDragging", 0), // event OnHeaderDragging
			/* 144 */ imports.NewTable("TLazVirtualStringTree_OnHeaderDraw", 0), // event OnHeaderDraw
			/* 145 */ imports.NewTable("TLazVirtualStringTree_OnHeaderDrawQueryElements", 0), // event OnHeaderDrawQueryElements
			/* 146 */ imports.NewTable("TLazVirtualStringTree_OnHeaderHeightDblClickResize", 0), // event OnHeaderHeightDblClickResize
			/* 147 */ imports.NewTable("TLazVirtualStringTree_OnHeaderHeightTracking", 0), // event OnHeaderHeightTracking
			/* 148 */ imports.NewTable("TLazVirtualStringTree_OnHeaderMouseDown", 0), // event OnHeaderMouseDown
			/* 149 */ imports.NewTable("TLazVirtualStringTree_OnHeaderMouseMove", 0), // event OnHeaderMouseMove
			/* 150 */ imports.NewTable("TLazVirtualStringTree_OnHeaderMouseUp", 0), // event OnHeaderMouseUp
			/* 151 */ imports.NewTable("TLazVirtualStringTree_OnHotChange", 0), // event OnHotChange
			/* 152 */ imports.NewTable("TLazVirtualStringTree_OnIncrementalSearch", 0), // event OnIncrementalSearch
			/* 153 */ imports.NewTable("TLazVirtualStringTree_OnInitChildren", 0), // event OnInitChildren
			/* 154 */ imports.NewTable("TLazVirtualStringTree_OnInitNode", 0), // event OnInitNode
			/* 155 */ imports.NewTable("TLazVirtualStringTree_OnKeyAction", 0), // event OnKeyAction
			/* 156 */ imports.NewTable("TLazVirtualStringTree_OnLoadNode", 0), // event OnLoadNode
			/* 157 */ imports.NewTable("TLazVirtualStringTree_OnLoadTree", 0), // event OnLoadTree
			/* 158 */ imports.NewTable("TLazVirtualStringTree_OnMeasureItem", 0), // event OnMeasureItem
			/* 159 */ imports.NewTable("TLazVirtualStringTree_OnMeasureTextWidth", 0), // event OnMeasureTextWidth
			/* 160 */ imports.NewTable("TLazVirtualStringTree_OnMeasureTextHeight", 0), // event OnMeasureTextHeight
			/* 161 */ imports.NewTable("TLazVirtualStringTree_OnMouseDown", 0), // event OnMouseDown
			/* 162 */ imports.NewTable("TLazVirtualStringTree_OnMouseMove", 0), // event OnMouseMove
			/* 163 */ imports.NewTable("TLazVirtualStringTree_OnMouseUp", 0), // event OnMouseUp
			/* 164 */ imports.NewTable("TLazVirtualStringTree_OnMouseWheel", 0), // event OnMouseWheel
			/* 165 */ imports.NewTable("TLazVirtualStringTree_OnMouseEnter", 0), // event OnMouseEnter
			/* 166 */ imports.NewTable("TLazVirtualStringTree_OnMouseLeave", 0), // event OnMouseLeave
			/* 167 */ imports.NewTable("TLazVirtualStringTree_OnNewText", 0), // event OnNewText
			/* 168 */ imports.NewTable("TLazVirtualStringTree_OnNodeClick", 0), // event OnNodeClick
			/* 169 */ imports.NewTable("TLazVirtualStringTree_OnNodeCopied", 0), // event OnNodeCopied
			/* 170 */ imports.NewTable("TLazVirtualStringTree_OnNodeCopying", 0), // event OnNodeCopying
			/* 171 */ imports.NewTable("TLazVirtualStringTree_OnNodeDblClick", 0), // event OnNodeDblClick
			/* 172 */ imports.NewTable("TLazVirtualStringTree_OnNodeExport", 0), // event OnNodeExport
			/* 173 */ imports.NewTable("TLazVirtualStringTree_OnNodeHeightDblClickResize", 0), // event OnNodeHeightDblClickResize
			/* 174 */ imports.NewTable("TLazVirtualStringTree_OnNodeHeightTracking", 0), // event OnNodeHeightTracking
			/* 175 */ imports.NewTable("TLazVirtualStringTree_OnNodeMoved", 0), // event OnNodeMoved
			/* 176 */ imports.NewTable("TLazVirtualStringTree_OnNodeMoving", 0), // event OnNodeMoving
			/* 177 */ imports.NewTable("TLazVirtualStringTree_OnPaintBackground", 0), // event OnPaintBackground
			/* 178 */ imports.NewTable("TLazVirtualStringTree_OnRemoveFromSelection", 0), // event OnRemoveFromSelection
			/* 179 */ imports.NewTable("TLazVirtualStringTree_OnResetNode", 0), // event OnResetNode
			/* 180 */ imports.NewTable("TLazVirtualStringTree_OnSaveNode", 0), // event OnSaveNode
			/* 181 */ imports.NewTable("TLazVirtualStringTree_OnSaveTree", 0), // event OnSaveTree
			/* 182 */ imports.NewTable("TLazVirtualStringTree_OnScroll", 0), // event OnScroll
			/* 183 */ imports.NewTable("TLazVirtualStringTree_OnShortenString", 0), // event OnShortenString
			/* 184 */ imports.NewTable("TLazVirtualStringTree_OnShowScrollBar", 0), // event OnShowScrollBar
			/* 185 */ imports.NewTable("TLazVirtualStringTree_OnStartDock", 0), // event OnStartDock
			/* 186 */ imports.NewTable("TLazVirtualStringTree_OnStartDrag", 0), // event OnStartDrag
			/* 187 */ imports.NewTable("TLazVirtualStringTree_OnStartOperation", 0), // event OnStartOperation
			/* 188 */ imports.NewTable("TLazVirtualStringTree_OnStateChange", 0), // event OnStateChange
			/* 189 */ imports.NewTable("TLazVirtualStringTree_OnStructureChange", 0), // event OnStructureChange
			/* 190 */ imports.NewTable("TLazVirtualStringTree_OnUpdating", 0), // event OnUpdating
			/* 191 */ imports.NewTable("TLazVirtualStringTree_TClass", 0), // function TLazVirtualStringTreeClass
		}
	})
	return lazVirtualStringTreeImport
}
