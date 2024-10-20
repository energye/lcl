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

// ILazVirtualStringTree Parent: ICustomVirtualStringTree
type ILazVirtualStringTree interface {
	ICustomVirtualStringTree
	RangeX() uint32                                                        // property
	LastDragEffect() uint32                                                // property
	Alignment() TAlignment                                                 // property
	SetAlignment(AValue TAlignment)                                        // property
	AnimationDuration() uint32                                             // property
	SetAnimationDuration(AValue uint32)                                    // property
	AutoExpandDelay() uint32                                               // property
	SetAutoExpandDelay(AValue uint32)                                      // property
	AutoScrollDelay() uint32                                               // property
	SetAutoScrollDelay(AValue uint32)                                      // property
	AutoScrollInterval() TAutoScrollInterval                               // property
	SetAutoScrollInterval(AValue TAutoScrollInterval)                      // property
	Background() IPicture                                                  // property
	SetBackground(AValue IPicture)                                         // property
	BackgroundOffsetX() int32                                              // property
	SetBackgroundOffsetX(AValue int32)                                     // property
	BackgroundOffsetY() int32                                              // property
	SetBackgroundOffsetY(AValue int32)                                     // property
	BottomSpace() uint32                                                   // property
	SetBottomSpace(AValue uint32)                                          // property
	ButtonFillMode() TVTButtonFillMode                                     // property
	SetButtonFillMode(AValue TVTButtonFillMode)                            // property
	ButtonStyle() TVTButtonStyle                                           // property
	SetButtonStyle(AValue TVTButtonStyle)                                  // property
	ChangeDelay() uint32                                                   // property
	SetChangeDelay(AValue uint32)                                          // property
	CheckImageKind() TCheckImageKind                                       // property
	SetCheckImageKind(AValue TCheckImageKind)                              // property
	ClipboardFormats() IClipboardFormats                                   // property
	SetClipboardFormats(AValue IClipboardFormats)                          // property
	Colors() IVTColors                                                     // property
	SetColors(AValue IVTColors)                                            // property
	CustomCheckImages() ICustomImageList                                   // property
	SetCustomCheckImages(AValue ICustomImageList)                          // property
	DefaultNodeHeight() uint32                                             // property
	SetDefaultNodeHeight(AValue uint32)                                    // property
	DefaultPasteMode() TVTNodeAttachMode                                   // property
	SetDefaultPasteMode(AValue TVTNodeAttachMode)                          // property
	DefaultText() string                                                   // property
	SetDefaultText(AValue string)                                          // property
	DragCursor() TCursor                                                   // property
	SetDragCursor(AValue TCursor)                                          // property
	DragHeight() int32                                                     // property
	SetDragHeight(AValue int32)                                            // property
	DragKind() TDragKind                                                   // property
	SetDragKind(AValue TDragKind)                                          // property
	DragImageKind() TVTDragImageKind                                       // property
	SetDragImageKind(AValue TVTDragImageKind)                              // property
	DragMode() TDragMode                                                   // property
	SetDragMode(AValue TDragMode)                                          // property
	DragOperations() TDragOperations                                       // property
	SetDragOperations(AValue TDragOperations)                              // property
	DragType() TVTDragType                                                 // property
	SetDragType(AValue TVTDragType)                                        // property
	DragWidth() int32                                                      // property
	SetDragWidth(AValue int32)                                             // property
	DrawSelectionMode() TVTDrawSelectionMode                               // property
	SetDrawSelectionMode(AValue TVTDrawSelectionMode)                      // property
	EditDelay() uint32                                                     // property
	SetEditDelay(AValue uint32)                                            // property
	Header() IVTHeader                                                     // property
	SetHeader(AValue IVTHeader)                                            // property
	HintMode() TVTHintMode                                                 // property
	SetHintMode(AValue TVTHintMode)                                        // property
	HotCursor() TCursor                                                    // property
	SetHotCursor(AValue TCursor)                                           // property
	Images() ICustomImageList                                              // property
	SetImages(AValue ICustomImageList)                                     // property
	IncrementalSearch() TVTIncrementalSearch                               // property
	SetIncrementalSearch(AValue TVTIncrementalSearch)                      // property
	IncrementalSearchDirection() TVTSearchDirection                        // property
	SetIncrementalSearchDirection(AValue TVTSearchDirection)               // property
	IncrementalSearchStart() TVTSearchStart                                // property
	SetIncrementalSearchStart(AValue TVTSearchStart)                       // property
	IncrementalSearchTimeout() uint32                                      // property
	SetIncrementalSearchTimeout(AValue uint32)                             // property
	Indent() uint32                                                        // property
	SetIndent(AValue uint32)                                               // property
	LineMode() TVTLineMode                                                 // property
	SetLineMode(AValue TVTLineMode)                                        // property
	LineStyle() TVTLineStyle                                               // property
	SetLineStyle(AValue TVTLineStyle)                                      // property
	Margin() int32                                                         // property
	SetMargin(AValue int32)                                                // property
	NodeAlignment() TVTNodeAlignment                                       // property
	SetNodeAlignment(AValue TVTNodeAlignment)                              // property
	NodeDataSize() int32                                                   // property
	SetNodeDataSize(AValue int32)                                          // property
	OperationCanceled() bool                                               // property
	ParentColor() bool                                                     // property
	SetParentColor(AValue bool)                                            // property
	ParentFont() bool                                                      // property
	SetParentFont(AValue bool)                                             // property
	ParentShowHint() bool                                                  // property
	SetParentShowHint(AValue bool)                                         // property
	RootNodeCount() uint32                                                 // property
	SetRootNodeCount(AValue uint32)                                        // property
	ScrollBarOptions() IScrollBarOptions                                   // property
	SetScrollBarOptions(AValue IScrollBarOptions)                          // property
	SelectionBlendFactor() Byte                                            // property
	SetSelectionBlendFactor(AValue Byte)                                   // property
	SelectionCurveRadius() uint32                                          // property
	SetSelectionCurveRadius(AValue uint32)                                 // property
	StateImages() ICustomImageList                                         // property
	SetStateImages(AValue ICustomImageList)                                // property
	TextMargin() int32                                                     // property
	SetTextMargin(AValue int32)                                            // property
	TreeOptions() IStringTreeOptions                                       // property
	SetTreeOptions(AValue IStringTreeOptions)                              // property
	WantTabs() bool                                                        // property
	SetWantTabs(AValue bool)                                               // property
	ImagesWidth() int32                                                    // property
	SetImagesWidth(AValue int32)                                           // property
	StateImagesWidth() int32                                               // property
	SetStateImagesWidth(AValue int32)                                      // property
	CustomCheckImagesWidth() int32                                         // property
	SetCustomCheckImagesWidth(AValue int32)                                // property
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
	SetOnGetHint(fn TVSTGetHintEvent)                                      // property event
	SetOnGetLineStyle(fn TVTGetLineStyleEvent)                             // property event
	SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent)                       // property event
	SetOnGetPopupMenu(fn TVTPopupEvent)                                    // property event
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
	SetOnShowScrollBar(fn TVTScrollBarShowEvent)                           // property event
	SetOnStartDock(fn TStartDockEvent)                                     // property event
	SetOnStartDrag(fn TStartDragEvent)                                     // property event
	SetOnStartOperation(fn TVTOperationEvent)                              // property event
	SetOnStateChange(fn TVTStateChangeEvent)                               // property event
	SetOnStructureChange(fn TVTStructureChangeEvent)                       // property event
	SetOnUpdating(fn TVTUpdatingEvent)                                     // property event
}

// TLazVirtualStringTree Parent: TCustomVirtualStringTree
type TLazVirtualStringTree struct {
	TCustomVirtualStringTree
	addToSelectionPtr             uintptr
	advancedHeaderDrawPtr         uintptr
	afterAutoFitColumnPtr         uintptr
	afterAutoFitColumnsPtr        uintptr
	afterCellPaintPtr             uintptr
	afterColumnExportPtr          uintptr
	afterColumnWidthTrackingPtr   uintptr
	afterGetMaxColumnWidthPtr     uintptr
	afterHeaderExportPtr          uintptr
	afterHeaderHeightTrackingPtr  uintptr
	afterItemErasePtr             uintptr
	afterItemPaintPtr             uintptr
	afterNodeExportPtr            uintptr
	afterPaintPtr                 uintptr
	afterTreeExportPtr            uintptr
	beforeAutoFitColumnPtr        uintptr
	beforeAutoFitColumnsPtr       uintptr
	beforeCellPaintPtr            uintptr
	beforeColumnExportPtr         uintptr
	beforeColumnWidthTrackingPtr  uintptr
	beforeDrawTreeLinePtr         uintptr
	beforeGetMaxColumnWidthPtr    uintptr
	beforeHeaderExportPtr         uintptr
	beforeHeaderHeightTrackingPtr uintptr
	beforeItemErasePtr            uintptr
	beforeItemPaintPtr            uintptr
	beforeNodeExportPtr           uintptr
	beforePaintPtr                uintptr
	beforeTreeExportPtr           uintptr
	canSplitterResizeColumnPtr    uintptr
	canSplitterResizeHeaderPtr    uintptr
	canSplitterResizeNodePtr      uintptr
	changePtr                     uintptr
	checkedPtr                    uintptr
	checkingPtr                   uintptr
	collapsedPtr                  uintptr
	collapsingPtr                 uintptr
	columnClickPtr                uintptr
	columnDblClickPtr             uintptr
	columnExportPtr               uintptr
	columnResizePtr               uintptr
	columnWidthDblClickResizePtr  uintptr
	columnWidthTrackingPtr        uintptr
	compareNodesPtr               uintptr
	contextPopupPtr               uintptr
	createDataObjectPtr           uintptr
	createDragManagerPtr          uintptr
	createEditorPtr               uintptr
	dblClickPtr                   uintptr
	dragAllowedPtr                uintptr
	dragOverPtr                   uintptr
	dragDropPtr                   uintptr
	drawHintPtr                   uintptr
	drawTextPtr                   uintptr
	editCancelledPtr              uintptr
	editedPtr                     uintptr
	editingPtr                    uintptr
	endDockPtr                    uintptr
	endDragPtr                    uintptr
	endOperationPtr               uintptr
	expandedPtr                   uintptr
	expandingPtr                  uintptr
	focusChangedPtr               uintptr
	focusChangingPtr              uintptr
	freeNodePtr                   uintptr
	getCellIsEmptyPtr             uintptr
	getCursorPtr                  uintptr
	getHeaderCursorPtr            uintptr
	getTextPtr                    uintptr
	paintTextPtr                  uintptr
	getHelpContextPtr             uintptr
	getHintKindPtr                uintptr
	getHintSizePtr                uintptr
	getImageIndexPtr              uintptr
	getImageIndexExPtr            uintptr
	getImageTextPtr               uintptr
	getHintPtr                    uintptr
	getLineStylePtr               uintptr
	getNodeDataSizePtr            uintptr
	getPopupMenuPtr               uintptr
	headerClickPtr                uintptr
	headerDblClickPtr             uintptr
	headerDraggedPtr              uintptr
	headerDraggedOutPtr           uintptr
	headerDraggingPtr             uintptr
	headerDrawPtr                 uintptr
	headerDrawQueryElementsPtr    uintptr
	headerHeightDblClickResizePtr uintptr
	headerHeightTrackingPtr       uintptr
	headerMouseDownPtr            uintptr
	headerMouseMovePtr            uintptr
	headerMouseUpPtr              uintptr
	hotChangePtr                  uintptr
	incrementalSearchPtr          uintptr
	initChildrenPtr               uintptr
	initNodePtr                   uintptr
	keyActionPtr                  uintptr
	loadNodePtr                   uintptr
	loadTreePtr                   uintptr
	measureItemPtr                uintptr
	measureTextWidthPtr           uintptr
	measureTextHeightPtr          uintptr
	mouseDownPtr                  uintptr
	mouseMovePtr                  uintptr
	mouseUpPtr                    uintptr
	mouseWheelPtr                 uintptr
	mouseEnterPtr                 uintptr
	mouseLeavePtr                 uintptr
	newTextPtr                    uintptr
	nodeClickPtr                  uintptr
	nodeCopiedPtr                 uintptr
	nodeCopyingPtr                uintptr
	nodeDblClickPtr               uintptr
	nodeExportPtr                 uintptr
	nodeHeightDblClickResizePtr   uintptr
	nodeHeightTrackingPtr         uintptr
	nodeMovedPtr                  uintptr
	nodeMovingPtr                 uintptr
	paintBackgroundPtr            uintptr
	removeFromSelectionPtr        uintptr
	resetNodePtr                  uintptr
	saveNodePtr                   uintptr
	saveTreePtr                   uintptr
	scrollPtr                     uintptr
	showScrollBarPtr              uintptr
	startDockPtr                  uintptr
	startDragPtr                  uintptr
	startOperationPtr             uintptr
	stateChangePtr                uintptr
	structureChangePtr            uintptr
	updatingPtr                   uintptr
}

func NewLazVirtualStringTree(AOwner IComponent) ILazVirtualStringTree {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(16, GetObjectUintptr(AOwner))
	return AsLazVirtualStringTree(r1)
}

func (m *TLazVirtualStringTree) RangeX() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(52, m.Instance())
	return uint32(r1)
}

func (m *TLazVirtualStringTree) LastDragEffect() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(42, m.Instance())
	return uint32(r1)
}

func (m *TLazVirtualStringTree) Alignment() TAlignment {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TLazVirtualStringTree) SetAlignment(AValue TAlignment) {
	lazVirtualStringTreeImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) AnimationDuration() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetAnimationDuration(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) AutoExpandDelay() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetAutoExpandDelay(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) AutoScrollDelay() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetAutoScrollDelay(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) AutoScrollInterval() TAutoScrollInterval {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TAutoScrollInterval(r1)
}

func (m *TLazVirtualStringTree) SetAutoScrollInterval(AValue TAutoScrollInterval) {
	lazVirtualStringTreeImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Background() IPicture {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TLazVirtualStringTree) SetBackground(AValue IPicture) {
	lazVirtualStringTreeImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) BackgroundOffsetX() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetBackgroundOffsetX(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) BackgroundOffsetY() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetBackgroundOffsetY(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) BottomSpace() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetBottomSpace(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ButtonFillMode() TVTButtonFillMode {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TVTButtonFillMode(r1)
}

func (m *TLazVirtualStringTree) SetButtonFillMode(AValue TVTButtonFillMode) {
	lazVirtualStringTreeImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ButtonStyle() TVTButtonStyle {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TVTButtonStyle(r1)
}

func (m *TLazVirtualStringTree) SetButtonStyle(AValue TVTButtonStyle) {
	lazVirtualStringTreeImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ChangeDelay() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetChangeDelay(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) CheckImageKind() TCheckImageKind {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TCheckImageKind(r1)
}

func (m *TLazVirtualStringTree) SetCheckImageKind(AValue TCheckImageKind) {
	lazVirtualStringTreeImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ClipboardFormats() IClipboardFormats {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return AsClipboardFormats(r1)
}

func (m *TLazVirtualStringTree) SetClipboardFormats(AValue IClipboardFormats) {
	lazVirtualStringTreeImportAPI().SysCallN(14, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) Colors() IVTColors {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return AsVTColors(r1)
}

func (m *TLazVirtualStringTree) SetColors(AValue IVTColors) {
	lazVirtualStringTreeImportAPI().SysCallN(15, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) CustomCheckImages() ICustomImageList {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualStringTree) SetCustomCheckImages(AValue ICustomImageList) {
	lazVirtualStringTreeImportAPI().SysCallN(17, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) DefaultNodeHeight() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetDefaultNodeHeight(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DefaultPasteMode() TVTNodeAttachMode {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return TVTNodeAttachMode(r1)
}

func (m *TLazVirtualStringTree) SetDefaultPasteMode(AValue TVTNodeAttachMode) {
	lazVirtualStringTreeImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DefaultText() string {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TLazVirtualStringTree) SetDefaultText(AValue string) {
	lazVirtualStringTreeImportAPI().SysCallN(21, 1, m.Instance(), PascalStr(AValue))
}

func (m *TLazVirtualStringTree) DragCursor() TCursor {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLazVirtualStringTree) SetDragCursor(AValue TCursor) {
	lazVirtualStringTreeImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragHeight() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetDragHeight(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragKind() TDragKind {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TLazVirtualStringTree) SetDragKind(AValue TDragKind) {
	lazVirtualStringTreeImportAPI().SysCallN(25, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragImageKind() TVTDragImageKind {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return TVTDragImageKind(r1)
}

func (m *TLazVirtualStringTree) SetDragImageKind(AValue TVTDragImageKind) {
	lazVirtualStringTreeImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragMode() TDragMode {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLazVirtualStringTree) SetDragMode(AValue TDragMode) {
	lazVirtualStringTreeImportAPI().SysCallN(26, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragOperations() TDragOperations {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return TDragOperations(r1)
}

func (m *TLazVirtualStringTree) SetDragOperations(AValue TDragOperations) {
	lazVirtualStringTreeImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragType() TVTDragType {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return TVTDragType(r1)
}

func (m *TLazVirtualStringTree) SetDragType(AValue TVTDragType) {
	lazVirtualStringTreeImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragWidth() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetDragWidth(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DrawSelectionMode() TVTDrawSelectionMode {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return TVTDrawSelectionMode(r1)
}

func (m *TLazVirtualStringTree) SetDrawSelectionMode(AValue TVTDrawSelectionMode) {
	lazVirtualStringTreeImportAPI().SysCallN(30, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) EditDelay() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetEditDelay(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(31, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Header() IVTHeader {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return AsVTHeader(r1)
}

func (m *TLazVirtualStringTree) SetHeader(AValue IVTHeader) {
	lazVirtualStringTreeImportAPI().SysCallN(32, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) HintMode() TVTHintMode {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return TVTHintMode(r1)
}

func (m *TLazVirtualStringTree) SetHintMode(AValue TVTHintMode) {
	lazVirtualStringTreeImportAPI().SysCallN(33, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) HotCursor() TCursor {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLazVirtualStringTree) SetHotCursor(AValue TCursor) {
	lazVirtualStringTreeImportAPI().SysCallN(34, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Images() ICustomImageList {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualStringTree) SetImages(AValue ICustomImageList) {
	lazVirtualStringTreeImportAPI().SysCallN(35, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) IncrementalSearch() TVTIncrementalSearch {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return TVTIncrementalSearch(r1)
}

func (m *TLazVirtualStringTree) SetIncrementalSearch(AValue TVTIncrementalSearch) {
	lazVirtualStringTreeImportAPI().SysCallN(37, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) IncrementalSearchDirection() TVTSearchDirection {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(38, 0, m.Instance(), 0)
	return TVTSearchDirection(r1)
}

func (m *TLazVirtualStringTree) SetIncrementalSearchDirection(AValue TVTSearchDirection) {
	lazVirtualStringTreeImportAPI().SysCallN(38, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) IncrementalSearchStart() TVTSearchStart {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(39, 0, m.Instance(), 0)
	return TVTSearchStart(r1)
}

func (m *TLazVirtualStringTree) SetIncrementalSearchStart(AValue TVTSearchStart) {
	lazVirtualStringTreeImportAPI().SysCallN(39, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) IncrementalSearchTimeout() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(40, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetIncrementalSearchTimeout(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(40, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Indent() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(41, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetIndent(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(41, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) LineMode() TVTLineMode {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(43, 0, m.Instance(), 0)
	return TVTLineMode(r1)
}

func (m *TLazVirtualStringTree) SetLineMode(AValue TVTLineMode) {
	lazVirtualStringTreeImportAPI().SysCallN(43, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) LineStyle() TVTLineStyle {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(44, 0, m.Instance(), 0)
	return TVTLineStyle(r1)
}

func (m *TLazVirtualStringTree) SetLineStyle(AValue TVTLineStyle) {
	lazVirtualStringTreeImportAPI().SysCallN(44, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Margin() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(45, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetMargin(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(45, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) NodeAlignment() TVTNodeAlignment {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(46, 0, m.Instance(), 0)
	return TVTNodeAlignment(r1)
}

func (m *TLazVirtualStringTree) SetNodeAlignment(AValue TVTNodeAlignment) {
	lazVirtualStringTreeImportAPI().SysCallN(46, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) NodeDataSize() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(47, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetNodeDataSize(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(47, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) OperationCanceled() bool {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(48, m.Instance())
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) ParentColor() bool {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(49, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) SetParentColor(AValue bool) {
	lazVirtualStringTreeImportAPI().SysCallN(49, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualStringTree) ParentFont() bool {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) SetParentFont(AValue bool) {
	lazVirtualStringTreeImportAPI().SysCallN(50, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualStringTree) ParentShowHint() bool {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(51, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) SetParentShowHint(AValue bool) {
	lazVirtualStringTreeImportAPI().SysCallN(51, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualStringTree) RootNodeCount() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(53, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetRootNodeCount(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(53, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ScrollBarOptions() IScrollBarOptions {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(54, 0, m.Instance(), 0)
	return AsScrollBarOptions(r1)
}

func (m *TLazVirtualStringTree) SetScrollBarOptions(AValue IScrollBarOptions) {
	lazVirtualStringTreeImportAPI().SysCallN(54, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) SelectionBlendFactor() Byte {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(55, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TLazVirtualStringTree) SetSelectionBlendFactor(AValue Byte) {
	lazVirtualStringTreeImportAPI().SysCallN(55, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) SelectionCurveRadius() uint32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(56, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetSelectionCurveRadius(AValue uint32) {
	lazVirtualStringTreeImportAPI().SysCallN(56, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) StateImages() ICustomImageList {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(188, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualStringTree) SetStateImages(AValue ICustomImageList) {
	lazVirtualStringTreeImportAPI().SysCallN(188, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) TextMargin() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(190, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetTextMargin(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(190, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) TreeOptions() IStringTreeOptions {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(191, 0, m.Instance(), 0)
	return AsStringTreeOptions(r1)
}

func (m *TLazVirtualStringTree) SetTreeOptions(AValue IStringTreeOptions) {
	lazVirtualStringTreeImportAPI().SysCallN(191, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) WantTabs() bool {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(192, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) SetWantTabs(AValue bool) {
	lazVirtualStringTreeImportAPI().SysCallN(192, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualStringTree) ImagesWidth() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetImagesWidth(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(36, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) StateImagesWidth() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(189, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetStateImagesWidth(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(189, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) CustomCheckImagesWidth() int32 {
	r1 := lazVirtualStringTreeImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetCustomCheckImagesWidth(AValue int32) {
	lazVirtualStringTreeImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func LazVirtualStringTreeClass() TClass {
	ret := lazVirtualStringTreeImportAPI().SysCallN(13)
	return TClass(ret)
}

func (m *TLazVirtualStringTree) SetOnAddToSelection(fn TVTAddToSelectionEvent) {
	if m.addToSelectionPtr != 0 {
		RemoveEventElement(m.addToSelectionPtr)
	}
	m.addToSelectionPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(57, m.Instance(), m.addToSelectionPtr)
}

func (m *TLazVirtualStringTree) SetOnAdvancedHeaderDraw(fn TVTAdvancedHeaderPaintEvent) {
	if m.advancedHeaderDrawPtr != 0 {
		RemoveEventElement(m.advancedHeaderDrawPtr)
	}
	m.advancedHeaderDrawPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(58, m.Instance(), m.advancedHeaderDrawPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterAutoFitColumn(fn TVTAfterAutoFitColumnEvent) {
	if m.afterAutoFitColumnPtr != 0 {
		RemoveEventElement(m.afterAutoFitColumnPtr)
	}
	m.afterAutoFitColumnPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(59, m.Instance(), m.afterAutoFitColumnPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterAutoFitColumns(fn TVTAfterAutoFitColumnsEvent) {
	if m.afterAutoFitColumnsPtr != 0 {
		RemoveEventElement(m.afterAutoFitColumnsPtr)
	}
	m.afterAutoFitColumnsPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(60, m.Instance(), m.afterAutoFitColumnsPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterCellPaint(fn TVTAfterCellPaintEvent) {
	if m.afterCellPaintPtr != 0 {
		RemoveEventElement(m.afterCellPaintPtr)
	}
	m.afterCellPaintPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(61, m.Instance(), m.afterCellPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterColumnExport(fn TVTColumnExportEvent) {
	if m.afterColumnExportPtr != 0 {
		RemoveEventElement(m.afterColumnExportPtr)
	}
	m.afterColumnExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(62, m.Instance(), m.afterColumnExportPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterColumnWidthTracking(fn TVTAfterColumnWidthTrackingEvent) {
	if m.afterColumnWidthTrackingPtr != 0 {
		RemoveEventElement(m.afterColumnWidthTrackingPtr)
	}
	m.afterColumnWidthTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(63, m.Instance(), m.afterColumnWidthTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterGetMaxColumnWidth(fn TVTAfterGetMaxColumnWidthEvent) {
	if m.afterGetMaxColumnWidthPtr != 0 {
		RemoveEventElement(m.afterGetMaxColumnWidthPtr)
	}
	m.afterGetMaxColumnWidthPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(64, m.Instance(), m.afterGetMaxColumnWidthPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterHeaderExport(fn TVTTreeExportEvent) {
	if m.afterHeaderExportPtr != 0 {
		RemoveEventElement(m.afterHeaderExportPtr)
	}
	m.afterHeaderExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(65, m.Instance(), m.afterHeaderExportPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterHeaderHeightTracking(fn TVTAfterHeaderHeightTrackingEvent) {
	if m.afterHeaderHeightTrackingPtr != 0 {
		RemoveEventElement(m.afterHeaderHeightTrackingPtr)
	}
	m.afterHeaderHeightTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(66, m.Instance(), m.afterHeaderHeightTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterItemErase(fn TVTAfterItemEraseEvent) {
	if m.afterItemErasePtr != 0 {
		RemoveEventElement(m.afterItemErasePtr)
	}
	m.afterItemErasePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(67, m.Instance(), m.afterItemErasePtr)
}

func (m *TLazVirtualStringTree) SetOnAfterItemPaint(fn TVTAfterItemPaintEvent) {
	if m.afterItemPaintPtr != 0 {
		RemoveEventElement(m.afterItemPaintPtr)
	}
	m.afterItemPaintPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(68, m.Instance(), m.afterItemPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterNodeExport(fn TVTNodeExportEvent) {
	if m.afterNodeExportPtr != 0 {
		RemoveEventElement(m.afterNodeExportPtr)
	}
	m.afterNodeExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(69, m.Instance(), m.afterNodeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterPaint(fn TVTPaintEvent) {
	if m.afterPaintPtr != 0 {
		RemoveEventElement(m.afterPaintPtr)
	}
	m.afterPaintPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(70, m.Instance(), m.afterPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterTreeExport(fn TVTTreeExportEvent) {
	if m.afterTreeExportPtr != 0 {
		RemoveEventElement(m.afterTreeExportPtr)
	}
	m.afterTreeExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(71, m.Instance(), m.afterTreeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeAutoFitColumn(fn TVTBeforeAutoFitColumnEvent) {
	if m.beforeAutoFitColumnPtr != 0 {
		RemoveEventElement(m.beforeAutoFitColumnPtr)
	}
	m.beforeAutoFitColumnPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(72, m.Instance(), m.beforeAutoFitColumnPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeAutoFitColumns(fn TVTBeforeAutoFitColumnsEvent) {
	if m.beforeAutoFitColumnsPtr != 0 {
		RemoveEventElement(m.beforeAutoFitColumnsPtr)
	}
	m.beforeAutoFitColumnsPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(73, m.Instance(), m.beforeAutoFitColumnsPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeCellPaint(fn TVTBeforeCellPaintEvent) {
	if m.beforeCellPaintPtr != 0 {
		RemoveEventElement(m.beforeCellPaintPtr)
	}
	m.beforeCellPaintPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(74, m.Instance(), m.beforeCellPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeColumnExport(fn TVTColumnExportEvent) {
	if m.beforeColumnExportPtr != 0 {
		RemoveEventElement(m.beforeColumnExportPtr)
	}
	m.beforeColumnExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(75, m.Instance(), m.beforeColumnExportPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeColumnWidthTracking(fn TVTBeforeColumnWidthTrackingEvent) {
	if m.beforeColumnWidthTrackingPtr != 0 {
		RemoveEventElement(m.beforeColumnWidthTrackingPtr)
	}
	m.beforeColumnWidthTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(76, m.Instance(), m.beforeColumnWidthTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeDrawTreeLine(fn TVTBeforeDrawLineImageEvent) {
	if m.beforeDrawTreeLinePtr != 0 {
		RemoveEventElement(m.beforeDrawTreeLinePtr)
	}
	m.beforeDrawTreeLinePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(77, m.Instance(), m.beforeDrawTreeLinePtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeGetMaxColumnWidth(fn TVTBeforeGetMaxColumnWidthEvent) {
	if m.beforeGetMaxColumnWidthPtr != 0 {
		RemoveEventElement(m.beforeGetMaxColumnWidthPtr)
	}
	m.beforeGetMaxColumnWidthPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(78, m.Instance(), m.beforeGetMaxColumnWidthPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeHeaderExport(fn TVTTreeExportEvent) {
	if m.beforeHeaderExportPtr != 0 {
		RemoveEventElement(m.beforeHeaderExportPtr)
	}
	m.beforeHeaderExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(79, m.Instance(), m.beforeHeaderExportPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeHeaderHeightTracking(fn TVTBeforeHeaderHeightTrackingEvent) {
	if m.beforeHeaderHeightTrackingPtr != 0 {
		RemoveEventElement(m.beforeHeaderHeightTrackingPtr)
	}
	m.beforeHeaderHeightTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(80, m.Instance(), m.beforeHeaderHeightTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeItemErase(fn TVTBeforeItemEraseEvent) {
	if m.beforeItemErasePtr != 0 {
		RemoveEventElement(m.beforeItemErasePtr)
	}
	m.beforeItemErasePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(81, m.Instance(), m.beforeItemErasePtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeItemPaint(fn TVTBeforeItemPaintEvent) {
	if m.beforeItemPaintPtr != 0 {
		RemoveEventElement(m.beforeItemPaintPtr)
	}
	m.beforeItemPaintPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(82, m.Instance(), m.beforeItemPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeNodeExport(fn TVTNodeExportEvent) {
	if m.beforeNodeExportPtr != 0 {
		RemoveEventElement(m.beforeNodeExportPtr)
	}
	m.beforeNodeExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(83, m.Instance(), m.beforeNodeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforePaint(fn TVTPaintEvent) {
	if m.beforePaintPtr != 0 {
		RemoveEventElement(m.beforePaintPtr)
	}
	m.beforePaintPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(84, m.Instance(), m.beforePaintPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeTreeExport(fn TVTTreeExportEvent) {
	if m.beforeTreeExportPtr != 0 {
		RemoveEventElement(m.beforeTreeExportPtr)
	}
	m.beforeTreeExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(85, m.Instance(), m.beforeTreeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnCanSplitterResizeColumn(fn TVTCanSplitterResizeColumnEvent) {
	if m.canSplitterResizeColumnPtr != 0 {
		RemoveEventElement(m.canSplitterResizeColumnPtr)
	}
	m.canSplitterResizeColumnPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(86, m.Instance(), m.canSplitterResizeColumnPtr)
}

func (m *TLazVirtualStringTree) SetOnCanSplitterResizeHeader(fn TVTCanSplitterResizeHeaderEvent) {
	if m.canSplitterResizeHeaderPtr != 0 {
		RemoveEventElement(m.canSplitterResizeHeaderPtr)
	}
	m.canSplitterResizeHeaderPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(87, m.Instance(), m.canSplitterResizeHeaderPtr)
}

func (m *TLazVirtualStringTree) SetOnCanSplitterResizeNode(fn TVTCanSplitterResizeNodeEvent) {
	if m.canSplitterResizeNodePtr != 0 {
		RemoveEventElement(m.canSplitterResizeNodePtr)
	}
	m.canSplitterResizeNodePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(88, m.Instance(), m.canSplitterResizeNodePtr)
}

func (m *TLazVirtualStringTree) SetOnChange(fn TVTChangeEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(89, m.Instance(), m.changePtr)
}

func (m *TLazVirtualStringTree) SetOnChecked(fn TVTChangeEvent) {
	if m.checkedPtr != 0 {
		RemoveEventElement(m.checkedPtr)
	}
	m.checkedPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(90, m.Instance(), m.checkedPtr)
}

func (m *TLazVirtualStringTree) SetOnChecking(fn TVTCheckChangingEvent) {
	if m.checkingPtr != 0 {
		RemoveEventElement(m.checkingPtr)
	}
	m.checkingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(91, m.Instance(), m.checkingPtr)
}

func (m *TLazVirtualStringTree) SetOnCollapsed(fn TVTChangeEvent) {
	if m.collapsedPtr != 0 {
		RemoveEventElement(m.collapsedPtr)
	}
	m.collapsedPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(92, m.Instance(), m.collapsedPtr)
}

func (m *TLazVirtualStringTree) SetOnCollapsing(fn TVTChangingEvent) {
	if m.collapsingPtr != 0 {
		RemoveEventElement(m.collapsingPtr)
	}
	m.collapsingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(93, m.Instance(), m.collapsingPtr)
}

func (m *TLazVirtualStringTree) SetOnColumnClick(fn TVTColumnClickEvent) {
	if m.columnClickPtr != 0 {
		RemoveEventElement(m.columnClickPtr)
	}
	m.columnClickPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(94, m.Instance(), m.columnClickPtr)
}

func (m *TLazVirtualStringTree) SetOnColumnDblClick(fn TVTColumnDblClickEvent) {
	if m.columnDblClickPtr != 0 {
		RemoveEventElement(m.columnDblClickPtr)
	}
	m.columnDblClickPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(95, m.Instance(), m.columnDblClickPtr)
}

func (m *TLazVirtualStringTree) SetOnColumnExport(fn TVTColumnExportEvent) {
	if m.columnExportPtr != 0 {
		RemoveEventElement(m.columnExportPtr)
	}
	m.columnExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(96, m.Instance(), m.columnExportPtr)
}

func (m *TLazVirtualStringTree) SetOnColumnResize(fn TVTHeaderNotifyEvent) {
	if m.columnResizePtr != 0 {
		RemoveEventElement(m.columnResizePtr)
	}
	m.columnResizePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(97, m.Instance(), m.columnResizePtr)
}

func (m *TLazVirtualStringTree) SetOnColumnWidthDblClickResize(fn TVTColumnWidthDblClickResizeEvent) {
	if m.columnWidthDblClickResizePtr != 0 {
		RemoveEventElement(m.columnWidthDblClickResizePtr)
	}
	m.columnWidthDblClickResizePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(98, m.Instance(), m.columnWidthDblClickResizePtr)
}

func (m *TLazVirtualStringTree) SetOnColumnWidthTracking(fn TVTColumnWidthTrackingEvent) {
	if m.columnWidthTrackingPtr != 0 {
		RemoveEventElement(m.columnWidthTrackingPtr)
	}
	m.columnWidthTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(99, m.Instance(), m.columnWidthTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnCompareNodes(fn TVTCompareEvent) {
	if m.compareNodesPtr != 0 {
		RemoveEventElement(m.compareNodesPtr)
	}
	m.compareNodesPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(100, m.Instance(), m.compareNodesPtr)
}

func (m *TLazVirtualStringTree) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(101, m.Instance(), m.contextPopupPtr)
}

func (m *TLazVirtualStringTree) SetOnCreateDataObject(fn TVTCreateDataObjectEvent) {
	if m.createDataObjectPtr != 0 {
		RemoveEventElement(m.createDataObjectPtr)
	}
	m.createDataObjectPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(102, m.Instance(), m.createDataObjectPtr)
}

func (m *TLazVirtualStringTree) SetOnCreateDragManager(fn TVTCreateDragManagerEvent) {
	if m.createDragManagerPtr != 0 {
		RemoveEventElement(m.createDragManagerPtr)
	}
	m.createDragManagerPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(103, m.Instance(), m.createDragManagerPtr)
}

func (m *TLazVirtualStringTree) SetOnCreateEditor(fn TVTCreateEditorEvent) {
	if m.createEditorPtr != 0 {
		RemoveEventElement(m.createEditorPtr)
	}
	m.createEditorPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(104, m.Instance(), m.createEditorPtr)
}

func (m *TLazVirtualStringTree) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(105, m.Instance(), m.dblClickPtr)
}

func (m *TLazVirtualStringTree) SetOnDragAllowed(fn TVTDragAllowedEvent) {
	if m.dragAllowedPtr != 0 {
		RemoveEventElement(m.dragAllowedPtr)
	}
	m.dragAllowedPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(106, m.Instance(), m.dragAllowedPtr)
}

func (m *TLazVirtualStringTree) SetOnDragOver(fn TVTDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(108, m.Instance(), m.dragOverPtr)
}

func (m *TLazVirtualStringTree) SetOnDragDrop(fn TVTDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(107, m.Instance(), m.dragDropPtr)
}

func (m *TLazVirtualStringTree) SetOnDrawHint(fn TVTDrawHintEvent) {
	if m.drawHintPtr != 0 {
		RemoveEventElement(m.drawHintPtr)
	}
	m.drawHintPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(109, m.Instance(), m.drawHintPtr)
}

func (m *TLazVirtualStringTree) SetOnDrawText(fn TVTDrawTextEvent) {
	if m.drawTextPtr != 0 {
		RemoveEventElement(m.drawTextPtr)
	}
	m.drawTextPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(110, m.Instance(), m.drawTextPtr)
}

func (m *TLazVirtualStringTree) SetOnEditCancelled(fn TVTEditCancelEvent) {
	if m.editCancelledPtr != 0 {
		RemoveEventElement(m.editCancelledPtr)
	}
	m.editCancelledPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(111, m.Instance(), m.editCancelledPtr)
}

func (m *TLazVirtualStringTree) SetOnEdited(fn TVTEditChangeEvent) {
	if m.editedPtr != 0 {
		RemoveEventElement(m.editedPtr)
	}
	m.editedPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(112, m.Instance(), m.editedPtr)
}

func (m *TLazVirtualStringTree) SetOnEditing(fn TVTEditChangingEvent) {
	if m.editingPtr != 0 {
		RemoveEventElement(m.editingPtr)
	}
	m.editingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(113, m.Instance(), m.editingPtr)
}

func (m *TLazVirtualStringTree) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(114, m.Instance(), m.endDockPtr)
}

func (m *TLazVirtualStringTree) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(115, m.Instance(), m.endDragPtr)
}

func (m *TLazVirtualStringTree) SetOnEndOperation(fn TVTOperationEvent) {
	if m.endOperationPtr != 0 {
		RemoveEventElement(m.endOperationPtr)
	}
	m.endOperationPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(116, m.Instance(), m.endOperationPtr)
}

func (m *TLazVirtualStringTree) SetOnExpanded(fn TVTChangeEvent) {
	if m.expandedPtr != 0 {
		RemoveEventElement(m.expandedPtr)
	}
	m.expandedPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(117, m.Instance(), m.expandedPtr)
}

func (m *TLazVirtualStringTree) SetOnExpanding(fn TVTChangingEvent) {
	if m.expandingPtr != 0 {
		RemoveEventElement(m.expandingPtr)
	}
	m.expandingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(118, m.Instance(), m.expandingPtr)
}

func (m *TLazVirtualStringTree) SetOnFocusChanged(fn TVTFocusChangeEvent) {
	if m.focusChangedPtr != 0 {
		RemoveEventElement(m.focusChangedPtr)
	}
	m.focusChangedPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(119, m.Instance(), m.focusChangedPtr)
}

func (m *TLazVirtualStringTree) SetOnFocusChanging(fn TVTFocusChangingEvent) {
	if m.focusChangingPtr != 0 {
		RemoveEventElement(m.focusChangingPtr)
	}
	m.focusChangingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(120, m.Instance(), m.focusChangingPtr)
}

func (m *TLazVirtualStringTree) SetOnFreeNode(fn TVTFreeNodeEvent) {
	if m.freeNodePtr != 0 {
		RemoveEventElement(m.freeNodePtr)
	}
	m.freeNodePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(121, m.Instance(), m.freeNodePtr)
}

func (m *TLazVirtualStringTree) SetOnGetCellIsEmpty(fn TVTGetCellIsEmptyEvent) {
	if m.getCellIsEmptyPtr != 0 {
		RemoveEventElement(m.getCellIsEmptyPtr)
	}
	m.getCellIsEmptyPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(122, m.Instance(), m.getCellIsEmptyPtr)
}

func (m *TLazVirtualStringTree) SetOnGetCursor(fn TVTGetCursorEvent) {
	if m.getCursorPtr != 0 {
		RemoveEventElement(m.getCursorPtr)
	}
	m.getCursorPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(123, m.Instance(), m.getCursorPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHeaderCursor(fn TVTGetHeaderCursorEvent) {
	if m.getHeaderCursorPtr != 0 {
		RemoveEventElement(m.getHeaderCursorPtr)
	}
	m.getHeaderCursorPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(124, m.Instance(), m.getHeaderCursorPtr)
}

func (m *TLazVirtualStringTree) SetOnGetText(fn TVSTGetTextEvent) {
	if m.getTextPtr != 0 {
		RemoveEventElement(m.getTextPtr)
	}
	m.getTextPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(135, m.Instance(), m.getTextPtr)
}

func (m *TLazVirtualStringTree) SetOnPaintText(fn TVTPaintText) {
	if m.paintTextPtr != 0 {
		RemoveEventElement(m.paintTextPtr)
	}
	m.paintTextPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(175, m.Instance(), m.paintTextPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHelpContext(fn TVTHelpContextEvent) {
	if m.getHelpContextPtr != 0 {
		RemoveEventElement(m.getHelpContextPtr)
	}
	m.getHelpContextPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(125, m.Instance(), m.getHelpContextPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHintKind(fn TVTHintKindEvent) {
	if m.getHintKindPtr != 0 {
		RemoveEventElement(m.getHintKindPtr)
	}
	m.getHintKindPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(127, m.Instance(), m.getHintKindPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHintSize(fn TVTGetHintSizeEvent) {
	if m.getHintSizePtr != 0 {
		RemoveEventElement(m.getHintSizePtr)
	}
	m.getHintSizePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(128, m.Instance(), m.getHintSizePtr)
}

func (m *TLazVirtualStringTree) SetOnGetImageIndex(fn TVTGetImageEvent) {
	if m.getImageIndexPtr != 0 {
		RemoveEventElement(m.getImageIndexPtr)
	}
	m.getImageIndexPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(129, m.Instance(), m.getImageIndexPtr)
}

func (m *TLazVirtualStringTree) SetOnGetImageIndexEx(fn TVTGetImageExEvent) {
	if m.getImageIndexExPtr != 0 {
		RemoveEventElement(m.getImageIndexExPtr)
	}
	m.getImageIndexExPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(130, m.Instance(), m.getImageIndexExPtr)
}

func (m *TLazVirtualStringTree) SetOnGetImageText(fn TVTGetImageTextEvent) {
	if m.getImageTextPtr != 0 {
		RemoveEventElement(m.getImageTextPtr)
	}
	m.getImageTextPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(131, m.Instance(), m.getImageTextPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHint(fn TVSTGetHintEvent) {
	if m.getHintPtr != 0 {
		RemoveEventElement(m.getHintPtr)
	}
	m.getHintPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(126, m.Instance(), m.getHintPtr)
}

func (m *TLazVirtualStringTree) SetOnGetLineStyle(fn TVTGetLineStyleEvent) {
	if m.getLineStylePtr != 0 {
		RemoveEventElement(m.getLineStylePtr)
	}
	m.getLineStylePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(132, m.Instance(), m.getLineStylePtr)
}

func (m *TLazVirtualStringTree) SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent) {
	if m.getNodeDataSizePtr != 0 {
		RemoveEventElement(m.getNodeDataSizePtr)
	}
	m.getNodeDataSizePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(133, m.Instance(), m.getNodeDataSizePtr)
}

func (m *TLazVirtualStringTree) SetOnGetPopupMenu(fn TVTPopupEvent) {
	if m.getPopupMenuPtr != 0 {
		RemoveEventElement(m.getPopupMenuPtr)
	}
	m.getPopupMenuPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(134, m.Instance(), m.getPopupMenuPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderClick(fn TVTHeaderClickEvent) {
	if m.headerClickPtr != 0 {
		RemoveEventElement(m.headerClickPtr)
	}
	m.headerClickPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(136, m.Instance(), m.headerClickPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDblClick(fn TVTHeaderClickEvent) {
	if m.headerDblClickPtr != 0 {
		RemoveEventElement(m.headerDblClickPtr)
	}
	m.headerDblClickPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(137, m.Instance(), m.headerDblClickPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDragged(fn TVTHeaderDraggedEvent) {
	if m.headerDraggedPtr != 0 {
		RemoveEventElement(m.headerDraggedPtr)
	}
	m.headerDraggedPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(138, m.Instance(), m.headerDraggedPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent) {
	if m.headerDraggedOutPtr != 0 {
		RemoveEventElement(m.headerDraggedOutPtr)
	}
	m.headerDraggedOutPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(139, m.Instance(), m.headerDraggedOutPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDragging(fn TVTHeaderDraggingEvent) {
	if m.headerDraggingPtr != 0 {
		RemoveEventElement(m.headerDraggingPtr)
	}
	m.headerDraggingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(140, m.Instance(), m.headerDraggingPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDraw(fn TVTHeaderPaintEvent) {
	if m.headerDrawPtr != 0 {
		RemoveEventElement(m.headerDrawPtr)
	}
	m.headerDrawPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(141, m.Instance(), m.headerDrawPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent) {
	if m.headerDrawQueryElementsPtr != 0 {
		RemoveEventElement(m.headerDrawQueryElementsPtr)
	}
	m.headerDrawQueryElementsPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(142, m.Instance(), m.headerDrawQueryElementsPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) {
	if m.headerHeightDblClickResizePtr != 0 {
		RemoveEventElement(m.headerHeightDblClickResizePtr)
	}
	m.headerHeightDblClickResizePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(143, m.Instance(), m.headerHeightDblClickResizePtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent) {
	if m.headerHeightTrackingPtr != 0 {
		RemoveEventElement(m.headerHeightTrackingPtr)
	}
	m.headerHeightTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(144, m.Instance(), m.headerHeightTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderMouseDown(fn TVTHeaderMouseEvent) {
	if m.headerMouseDownPtr != 0 {
		RemoveEventElement(m.headerMouseDownPtr)
	}
	m.headerMouseDownPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(145, m.Instance(), m.headerMouseDownPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderMouseMove(fn TVTHeaderMouseMoveEvent) {
	if m.headerMouseMovePtr != 0 {
		RemoveEventElement(m.headerMouseMovePtr)
	}
	m.headerMouseMovePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(146, m.Instance(), m.headerMouseMovePtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderMouseUp(fn TVTHeaderMouseEvent) {
	if m.headerMouseUpPtr != 0 {
		RemoveEventElement(m.headerMouseUpPtr)
	}
	m.headerMouseUpPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(147, m.Instance(), m.headerMouseUpPtr)
}

func (m *TLazVirtualStringTree) SetOnHotChange(fn TVTHotNodeChangeEvent) {
	if m.hotChangePtr != 0 {
		RemoveEventElement(m.hotChangePtr)
	}
	m.hotChangePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(148, m.Instance(), m.hotChangePtr)
}

func (m *TLazVirtualStringTree) SetOnIncrementalSearch(fn TVTIncrementalSearchEvent) {
	if m.incrementalSearchPtr != 0 {
		RemoveEventElement(m.incrementalSearchPtr)
	}
	m.incrementalSearchPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(149, m.Instance(), m.incrementalSearchPtr)
}

func (m *TLazVirtualStringTree) SetOnInitChildren(fn TVTInitChildrenEvent) {
	if m.initChildrenPtr != 0 {
		RemoveEventElement(m.initChildrenPtr)
	}
	m.initChildrenPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(150, m.Instance(), m.initChildrenPtr)
}

func (m *TLazVirtualStringTree) SetOnInitNode(fn TVTInitNodeEvent) {
	if m.initNodePtr != 0 {
		RemoveEventElement(m.initNodePtr)
	}
	m.initNodePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(151, m.Instance(), m.initNodePtr)
}

func (m *TLazVirtualStringTree) SetOnKeyAction(fn TVTKeyActionEvent) {
	if m.keyActionPtr != 0 {
		RemoveEventElement(m.keyActionPtr)
	}
	m.keyActionPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(152, m.Instance(), m.keyActionPtr)
}

func (m *TLazVirtualStringTree) SetOnLoadNode(fn TVTSaveNodeEvent) {
	if m.loadNodePtr != 0 {
		RemoveEventElement(m.loadNodePtr)
	}
	m.loadNodePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(153, m.Instance(), m.loadNodePtr)
}

func (m *TLazVirtualStringTree) SetOnLoadTree(fn TVTSaveTreeEvent) {
	if m.loadTreePtr != 0 {
		RemoveEventElement(m.loadTreePtr)
	}
	m.loadTreePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(154, m.Instance(), m.loadTreePtr)
}

func (m *TLazVirtualStringTree) SetOnMeasureItem(fn TVTMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(155, m.Instance(), m.measureItemPtr)
}

func (m *TLazVirtualStringTree) SetOnMeasureTextWidth(fn TVTMeasureTextEvent) {
	if m.measureTextWidthPtr != 0 {
		RemoveEventElement(m.measureTextWidthPtr)
	}
	m.measureTextWidthPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(157, m.Instance(), m.measureTextWidthPtr)
}

func (m *TLazVirtualStringTree) SetOnMeasureTextHeight(fn TVTMeasureTextEvent) {
	if m.measureTextHeightPtr != 0 {
		RemoveEventElement(m.measureTextHeightPtr)
	}
	m.measureTextHeightPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(156, m.Instance(), m.measureTextHeightPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(158, m.Instance(), m.mouseDownPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(161, m.Instance(), m.mouseMovePtr)
}

func (m *TLazVirtualStringTree) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(162, m.Instance(), m.mouseUpPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(163, m.Instance(), m.mouseWheelPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(159, m.Instance(), m.mouseEnterPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(160, m.Instance(), m.mouseLeavePtr)
}

func (m *TLazVirtualStringTree) SetOnNewText(fn TVSTNewTextEvent) {
	if m.newTextPtr != 0 {
		RemoveEventElement(m.newTextPtr)
	}
	m.newTextPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(164, m.Instance(), m.newTextPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeClick(fn TVTNodeClickEvent) {
	if m.nodeClickPtr != 0 {
		RemoveEventElement(m.nodeClickPtr)
	}
	m.nodeClickPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(165, m.Instance(), m.nodeClickPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeCopied(fn TVTNodeCopiedEvent) {
	if m.nodeCopiedPtr != 0 {
		RemoveEventElement(m.nodeCopiedPtr)
	}
	m.nodeCopiedPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(166, m.Instance(), m.nodeCopiedPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeCopying(fn TVTNodeCopyingEvent) {
	if m.nodeCopyingPtr != 0 {
		RemoveEventElement(m.nodeCopyingPtr)
	}
	m.nodeCopyingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(167, m.Instance(), m.nodeCopyingPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeDblClick(fn TVTNodeClickEvent) {
	if m.nodeDblClickPtr != 0 {
		RemoveEventElement(m.nodeDblClickPtr)
	}
	m.nodeDblClickPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(168, m.Instance(), m.nodeDblClickPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeExport(fn TVTNodeExportEvent) {
	if m.nodeExportPtr != 0 {
		RemoveEventElement(m.nodeExportPtr)
	}
	m.nodeExportPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(169, m.Instance(), m.nodeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent) {
	if m.nodeHeightDblClickResizePtr != 0 {
		RemoveEventElement(m.nodeHeightDblClickResizePtr)
	}
	m.nodeHeightDblClickResizePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(170, m.Instance(), m.nodeHeightDblClickResizePtr)
}

func (m *TLazVirtualStringTree) SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent) {
	if m.nodeHeightTrackingPtr != 0 {
		RemoveEventElement(m.nodeHeightTrackingPtr)
	}
	m.nodeHeightTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(171, m.Instance(), m.nodeHeightTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeMoved(fn TVTNodeMovedEvent) {
	if m.nodeMovedPtr != 0 {
		RemoveEventElement(m.nodeMovedPtr)
	}
	m.nodeMovedPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(172, m.Instance(), m.nodeMovedPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeMoving(fn TVTNodeMovingEvent) {
	if m.nodeMovingPtr != 0 {
		RemoveEventElement(m.nodeMovingPtr)
	}
	m.nodeMovingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(173, m.Instance(), m.nodeMovingPtr)
}

func (m *TLazVirtualStringTree) SetOnPaintBackground(fn TVTBackgroundPaintEvent) {
	if m.paintBackgroundPtr != 0 {
		RemoveEventElement(m.paintBackgroundPtr)
	}
	m.paintBackgroundPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(174, m.Instance(), m.paintBackgroundPtr)
}

func (m *TLazVirtualStringTree) SetOnRemoveFromSelection(fn TVTRemoveFromSelectionEvent) {
	if m.removeFromSelectionPtr != 0 {
		RemoveEventElement(m.removeFromSelectionPtr)
	}
	m.removeFromSelectionPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(176, m.Instance(), m.removeFromSelectionPtr)
}

func (m *TLazVirtualStringTree) SetOnResetNode(fn TVTChangeEvent) {
	if m.resetNodePtr != 0 {
		RemoveEventElement(m.resetNodePtr)
	}
	m.resetNodePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(177, m.Instance(), m.resetNodePtr)
}

func (m *TLazVirtualStringTree) SetOnSaveNode(fn TVTSaveNodeEvent) {
	if m.saveNodePtr != 0 {
		RemoveEventElement(m.saveNodePtr)
	}
	m.saveNodePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(178, m.Instance(), m.saveNodePtr)
}

func (m *TLazVirtualStringTree) SetOnSaveTree(fn TVTSaveTreeEvent) {
	if m.saveTreePtr != 0 {
		RemoveEventElement(m.saveTreePtr)
	}
	m.saveTreePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(179, m.Instance(), m.saveTreePtr)
}

func (m *TLazVirtualStringTree) SetOnScroll(fn TVTScrollEvent) {
	if m.scrollPtr != 0 {
		RemoveEventElement(m.scrollPtr)
	}
	m.scrollPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(180, m.Instance(), m.scrollPtr)
}

func (m *TLazVirtualStringTree) SetOnShowScrollBar(fn TVTScrollBarShowEvent) {
	if m.showScrollBarPtr != 0 {
		RemoveEventElement(m.showScrollBarPtr)
	}
	m.showScrollBarPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(181, m.Instance(), m.showScrollBarPtr)
}

func (m *TLazVirtualStringTree) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(182, m.Instance(), m.startDockPtr)
}

func (m *TLazVirtualStringTree) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(183, m.Instance(), m.startDragPtr)
}

func (m *TLazVirtualStringTree) SetOnStartOperation(fn TVTOperationEvent) {
	if m.startOperationPtr != 0 {
		RemoveEventElement(m.startOperationPtr)
	}
	m.startOperationPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(184, m.Instance(), m.startOperationPtr)
}

func (m *TLazVirtualStringTree) SetOnStateChange(fn TVTStateChangeEvent) {
	if m.stateChangePtr != 0 {
		RemoveEventElement(m.stateChangePtr)
	}
	m.stateChangePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(185, m.Instance(), m.stateChangePtr)
}

func (m *TLazVirtualStringTree) SetOnStructureChange(fn TVTStructureChangeEvent) {
	if m.structureChangePtr != 0 {
		RemoveEventElement(m.structureChangePtr)
	}
	m.structureChangePtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(186, m.Instance(), m.structureChangePtr)
}

func (m *TLazVirtualStringTree) SetOnUpdating(fn TVTUpdatingEvent) {
	if m.updatingPtr != 0 {
		RemoveEventElement(m.updatingPtr)
	}
	m.updatingPtr = MakeEventDataPtr(fn)
	lazVirtualStringTreeImportAPI().SysCallN(187, m.Instance(), m.updatingPtr)
}

var (
	lazVirtualStringTreeImport       *imports.Imports = nil
	lazVirtualStringTreeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazVirtualStringTree_Alignment", 0),
		/*1*/ imports.NewTable("LazVirtualStringTree_AnimationDuration", 0),
		/*2*/ imports.NewTable("LazVirtualStringTree_AutoExpandDelay", 0),
		/*3*/ imports.NewTable("LazVirtualStringTree_AutoScrollDelay", 0),
		/*4*/ imports.NewTable("LazVirtualStringTree_AutoScrollInterval", 0),
		/*5*/ imports.NewTable("LazVirtualStringTree_Background", 0),
		/*6*/ imports.NewTable("LazVirtualStringTree_BackgroundOffsetX", 0),
		/*7*/ imports.NewTable("LazVirtualStringTree_BackgroundOffsetY", 0),
		/*8*/ imports.NewTable("LazVirtualStringTree_BottomSpace", 0),
		/*9*/ imports.NewTable("LazVirtualStringTree_ButtonFillMode", 0),
		/*10*/ imports.NewTable("LazVirtualStringTree_ButtonStyle", 0),
		/*11*/ imports.NewTable("LazVirtualStringTree_ChangeDelay", 0),
		/*12*/ imports.NewTable("LazVirtualStringTree_CheckImageKind", 0),
		/*13*/ imports.NewTable("LazVirtualStringTree_Class", 0),
		/*14*/ imports.NewTable("LazVirtualStringTree_ClipboardFormats", 0),
		/*15*/ imports.NewTable("LazVirtualStringTree_Colors", 0),
		/*16*/ imports.NewTable("LazVirtualStringTree_Create", 0),
		/*17*/ imports.NewTable("LazVirtualStringTree_CustomCheckImages", 0),
		/*18*/ imports.NewTable("LazVirtualStringTree_CustomCheckImagesWidth", 0),
		/*19*/ imports.NewTable("LazVirtualStringTree_DefaultNodeHeight", 0),
		/*20*/ imports.NewTable("LazVirtualStringTree_DefaultPasteMode", 0),
		/*21*/ imports.NewTable("LazVirtualStringTree_DefaultText", 0),
		/*22*/ imports.NewTable("LazVirtualStringTree_DragCursor", 0),
		/*23*/ imports.NewTable("LazVirtualStringTree_DragHeight", 0),
		/*24*/ imports.NewTable("LazVirtualStringTree_DragImageKind", 0),
		/*25*/ imports.NewTable("LazVirtualStringTree_DragKind", 0),
		/*26*/ imports.NewTable("LazVirtualStringTree_DragMode", 0),
		/*27*/ imports.NewTable("LazVirtualStringTree_DragOperations", 0),
		/*28*/ imports.NewTable("LazVirtualStringTree_DragType", 0),
		/*29*/ imports.NewTable("LazVirtualStringTree_DragWidth", 0),
		/*30*/ imports.NewTable("LazVirtualStringTree_DrawSelectionMode", 0),
		/*31*/ imports.NewTable("LazVirtualStringTree_EditDelay", 0),
		/*32*/ imports.NewTable("LazVirtualStringTree_Header", 0),
		/*33*/ imports.NewTable("LazVirtualStringTree_HintMode", 0),
		/*34*/ imports.NewTable("LazVirtualStringTree_HotCursor", 0),
		/*35*/ imports.NewTable("LazVirtualStringTree_Images", 0),
		/*36*/ imports.NewTable("LazVirtualStringTree_ImagesWidth", 0),
		/*37*/ imports.NewTable("LazVirtualStringTree_IncrementalSearch", 0),
		/*38*/ imports.NewTable("LazVirtualStringTree_IncrementalSearchDirection", 0),
		/*39*/ imports.NewTable("LazVirtualStringTree_IncrementalSearchStart", 0),
		/*40*/ imports.NewTable("LazVirtualStringTree_IncrementalSearchTimeout", 0),
		/*41*/ imports.NewTable("LazVirtualStringTree_Indent", 0),
		/*42*/ imports.NewTable("LazVirtualStringTree_LastDragEffect", 0),
		/*43*/ imports.NewTable("LazVirtualStringTree_LineMode", 0),
		/*44*/ imports.NewTable("LazVirtualStringTree_LineStyle", 0),
		/*45*/ imports.NewTable("LazVirtualStringTree_Margin", 0),
		/*46*/ imports.NewTable("LazVirtualStringTree_NodeAlignment", 0),
		/*47*/ imports.NewTable("LazVirtualStringTree_NodeDataSize", 0),
		/*48*/ imports.NewTable("LazVirtualStringTree_OperationCanceled", 0),
		/*49*/ imports.NewTable("LazVirtualStringTree_ParentColor", 0),
		/*50*/ imports.NewTable("LazVirtualStringTree_ParentFont", 0),
		/*51*/ imports.NewTable("LazVirtualStringTree_ParentShowHint", 0),
		/*52*/ imports.NewTable("LazVirtualStringTree_RangeX", 0),
		/*53*/ imports.NewTable("LazVirtualStringTree_RootNodeCount", 0),
		/*54*/ imports.NewTable("LazVirtualStringTree_ScrollBarOptions", 0),
		/*55*/ imports.NewTable("LazVirtualStringTree_SelectionBlendFactor", 0),
		/*56*/ imports.NewTable("LazVirtualStringTree_SelectionCurveRadius", 0),
		/*57*/ imports.NewTable("LazVirtualStringTree_SetOnAddToSelection", 0),
		/*58*/ imports.NewTable("LazVirtualStringTree_SetOnAdvancedHeaderDraw", 0),
		/*59*/ imports.NewTable("LazVirtualStringTree_SetOnAfterAutoFitColumn", 0),
		/*60*/ imports.NewTable("LazVirtualStringTree_SetOnAfterAutoFitColumns", 0),
		/*61*/ imports.NewTable("LazVirtualStringTree_SetOnAfterCellPaint", 0),
		/*62*/ imports.NewTable("LazVirtualStringTree_SetOnAfterColumnExport", 0),
		/*63*/ imports.NewTable("LazVirtualStringTree_SetOnAfterColumnWidthTracking", 0),
		/*64*/ imports.NewTable("LazVirtualStringTree_SetOnAfterGetMaxColumnWidth", 0),
		/*65*/ imports.NewTable("LazVirtualStringTree_SetOnAfterHeaderExport", 0),
		/*66*/ imports.NewTable("LazVirtualStringTree_SetOnAfterHeaderHeightTracking", 0),
		/*67*/ imports.NewTable("LazVirtualStringTree_SetOnAfterItemErase", 0),
		/*68*/ imports.NewTable("LazVirtualStringTree_SetOnAfterItemPaint", 0),
		/*69*/ imports.NewTable("LazVirtualStringTree_SetOnAfterNodeExport", 0),
		/*70*/ imports.NewTable("LazVirtualStringTree_SetOnAfterPaint", 0),
		/*71*/ imports.NewTable("LazVirtualStringTree_SetOnAfterTreeExport", 0),
		/*72*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeAutoFitColumn", 0),
		/*73*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeAutoFitColumns", 0),
		/*74*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeCellPaint", 0),
		/*75*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeColumnExport", 0),
		/*76*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeColumnWidthTracking", 0),
		/*77*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeDrawTreeLine", 0),
		/*78*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeGetMaxColumnWidth", 0),
		/*79*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeHeaderExport", 0),
		/*80*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeHeaderHeightTracking", 0),
		/*81*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeItemErase", 0),
		/*82*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeItemPaint", 0),
		/*83*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeNodeExport", 0),
		/*84*/ imports.NewTable("LazVirtualStringTree_SetOnBeforePaint", 0),
		/*85*/ imports.NewTable("LazVirtualStringTree_SetOnBeforeTreeExport", 0),
		/*86*/ imports.NewTable("LazVirtualStringTree_SetOnCanSplitterResizeColumn", 0),
		/*87*/ imports.NewTable("LazVirtualStringTree_SetOnCanSplitterResizeHeader", 0),
		/*88*/ imports.NewTable("LazVirtualStringTree_SetOnCanSplitterResizeNode", 0),
		/*89*/ imports.NewTable("LazVirtualStringTree_SetOnChange", 0),
		/*90*/ imports.NewTable("LazVirtualStringTree_SetOnChecked", 0),
		/*91*/ imports.NewTable("LazVirtualStringTree_SetOnChecking", 0),
		/*92*/ imports.NewTable("LazVirtualStringTree_SetOnCollapsed", 0),
		/*93*/ imports.NewTable("LazVirtualStringTree_SetOnCollapsing", 0),
		/*94*/ imports.NewTable("LazVirtualStringTree_SetOnColumnClick", 0),
		/*95*/ imports.NewTable("LazVirtualStringTree_SetOnColumnDblClick", 0),
		/*96*/ imports.NewTable("LazVirtualStringTree_SetOnColumnExport", 0),
		/*97*/ imports.NewTable("LazVirtualStringTree_SetOnColumnResize", 0),
		/*98*/ imports.NewTable("LazVirtualStringTree_SetOnColumnWidthDblClickResize", 0),
		/*99*/ imports.NewTable("LazVirtualStringTree_SetOnColumnWidthTracking", 0),
		/*100*/ imports.NewTable("LazVirtualStringTree_SetOnCompareNodes", 0),
		/*101*/ imports.NewTable("LazVirtualStringTree_SetOnContextPopup", 0),
		/*102*/ imports.NewTable("LazVirtualStringTree_SetOnCreateDataObject", 0),
		/*103*/ imports.NewTable("LazVirtualStringTree_SetOnCreateDragManager", 0),
		/*104*/ imports.NewTable("LazVirtualStringTree_SetOnCreateEditor", 0),
		/*105*/ imports.NewTable("LazVirtualStringTree_SetOnDblClick", 0),
		/*106*/ imports.NewTable("LazVirtualStringTree_SetOnDragAllowed", 0),
		/*107*/ imports.NewTable("LazVirtualStringTree_SetOnDragDrop", 0),
		/*108*/ imports.NewTable("LazVirtualStringTree_SetOnDragOver", 0),
		/*109*/ imports.NewTable("LazVirtualStringTree_SetOnDrawHint", 0),
		/*110*/ imports.NewTable("LazVirtualStringTree_SetOnDrawText", 0),
		/*111*/ imports.NewTable("LazVirtualStringTree_SetOnEditCancelled", 0),
		/*112*/ imports.NewTable("LazVirtualStringTree_SetOnEdited", 0),
		/*113*/ imports.NewTable("LazVirtualStringTree_SetOnEditing", 0),
		/*114*/ imports.NewTable("LazVirtualStringTree_SetOnEndDock", 0),
		/*115*/ imports.NewTable("LazVirtualStringTree_SetOnEndDrag", 0),
		/*116*/ imports.NewTable("LazVirtualStringTree_SetOnEndOperation", 0),
		/*117*/ imports.NewTable("LazVirtualStringTree_SetOnExpanded", 0),
		/*118*/ imports.NewTable("LazVirtualStringTree_SetOnExpanding", 0),
		/*119*/ imports.NewTable("LazVirtualStringTree_SetOnFocusChanged", 0),
		/*120*/ imports.NewTable("LazVirtualStringTree_SetOnFocusChanging", 0),
		/*121*/ imports.NewTable("LazVirtualStringTree_SetOnFreeNode", 0),
		/*122*/ imports.NewTable("LazVirtualStringTree_SetOnGetCellIsEmpty", 0),
		/*123*/ imports.NewTable("LazVirtualStringTree_SetOnGetCursor", 0),
		/*124*/ imports.NewTable("LazVirtualStringTree_SetOnGetHeaderCursor", 0),
		/*125*/ imports.NewTable("LazVirtualStringTree_SetOnGetHelpContext", 0),
		/*126*/ imports.NewTable("LazVirtualStringTree_SetOnGetHint", 0),
		/*127*/ imports.NewTable("LazVirtualStringTree_SetOnGetHintKind", 0),
		/*128*/ imports.NewTable("LazVirtualStringTree_SetOnGetHintSize", 0),
		/*129*/ imports.NewTable("LazVirtualStringTree_SetOnGetImageIndex", 0),
		/*130*/ imports.NewTable("LazVirtualStringTree_SetOnGetImageIndexEx", 0),
		/*131*/ imports.NewTable("LazVirtualStringTree_SetOnGetImageText", 0),
		/*132*/ imports.NewTable("LazVirtualStringTree_SetOnGetLineStyle", 0),
		/*133*/ imports.NewTable("LazVirtualStringTree_SetOnGetNodeDataSize", 0),
		/*134*/ imports.NewTable("LazVirtualStringTree_SetOnGetPopupMenu", 0),
		/*135*/ imports.NewTable("LazVirtualStringTree_SetOnGetText", 0),
		/*136*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderClick", 0),
		/*137*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderDblClick", 0),
		/*138*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderDragged", 0),
		/*139*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderDraggedOut", 0),
		/*140*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderDragging", 0),
		/*141*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderDraw", 0),
		/*142*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderDrawQueryElements", 0),
		/*143*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderHeightDblClickResize", 0),
		/*144*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderHeightTracking", 0),
		/*145*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderMouseDown", 0),
		/*146*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderMouseMove", 0),
		/*147*/ imports.NewTable("LazVirtualStringTree_SetOnHeaderMouseUp", 0),
		/*148*/ imports.NewTable("LazVirtualStringTree_SetOnHotChange", 0),
		/*149*/ imports.NewTable("LazVirtualStringTree_SetOnIncrementalSearch", 0),
		/*150*/ imports.NewTable("LazVirtualStringTree_SetOnInitChildren", 0),
		/*151*/ imports.NewTable("LazVirtualStringTree_SetOnInitNode", 0),
		/*152*/ imports.NewTable("LazVirtualStringTree_SetOnKeyAction", 0),
		/*153*/ imports.NewTable("LazVirtualStringTree_SetOnLoadNode", 0),
		/*154*/ imports.NewTable("LazVirtualStringTree_SetOnLoadTree", 0),
		/*155*/ imports.NewTable("LazVirtualStringTree_SetOnMeasureItem", 0),
		/*156*/ imports.NewTable("LazVirtualStringTree_SetOnMeasureTextHeight", 0),
		/*157*/ imports.NewTable("LazVirtualStringTree_SetOnMeasureTextWidth", 0),
		/*158*/ imports.NewTable("LazVirtualStringTree_SetOnMouseDown", 0),
		/*159*/ imports.NewTable("LazVirtualStringTree_SetOnMouseEnter", 0),
		/*160*/ imports.NewTable("LazVirtualStringTree_SetOnMouseLeave", 0),
		/*161*/ imports.NewTable("LazVirtualStringTree_SetOnMouseMove", 0),
		/*162*/ imports.NewTable("LazVirtualStringTree_SetOnMouseUp", 0),
		/*163*/ imports.NewTable("LazVirtualStringTree_SetOnMouseWheel", 0),
		/*164*/ imports.NewTable("LazVirtualStringTree_SetOnNewText", 0),
		/*165*/ imports.NewTable("LazVirtualStringTree_SetOnNodeClick", 0),
		/*166*/ imports.NewTable("LazVirtualStringTree_SetOnNodeCopied", 0),
		/*167*/ imports.NewTable("LazVirtualStringTree_SetOnNodeCopying", 0),
		/*168*/ imports.NewTable("LazVirtualStringTree_SetOnNodeDblClick", 0),
		/*169*/ imports.NewTable("LazVirtualStringTree_SetOnNodeExport", 0),
		/*170*/ imports.NewTable("LazVirtualStringTree_SetOnNodeHeightDblClickResize", 0),
		/*171*/ imports.NewTable("LazVirtualStringTree_SetOnNodeHeightTracking", 0),
		/*172*/ imports.NewTable("LazVirtualStringTree_SetOnNodeMoved", 0),
		/*173*/ imports.NewTable("LazVirtualStringTree_SetOnNodeMoving", 0),
		/*174*/ imports.NewTable("LazVirtualStringTree_SetOnPaintBackground", 0),
		/*175*/ imports.NewTable("LazVirtualStringTree_SetOnPaintText", 0),
		/*176*/ imports.NewTable("LazVirtualStringTree_SetOnRemoveFromSelection", 0),
		/*177*/ imports.NewTable("LazVirtualStringTree_SetOnResetNode", 0),
		/*178*/ imports.NewTable("LazVirtualStringTree_SetOnSaveNode", 0),
		/*179*/ imports.NewTable("LazVirtualStringTree_SetOnSaveTree", 0),
		/*180*/ imports.NewTable("LazVirtualStringTree_SetOnScroll", 0),
		/*181*/ imports.NewTable("LazVirtualStringTree_SetOnShowScrollBar", 0),
		/*182*/ imports.NewTable("LazVirtualStringTree_SetOnStartDock", 0),
		/*183*/ imports.NewTable("LazVirtualStringTree_SetOnStartDrag", 0),
		/*184*/ imports.NewTable("LazVirtualStringTree_SetOnStartOperation", 0),
		/*185*/ imports.NewTable("LazVirtualStringTree_SetOnStateChange", 0),
		/*186*/ imports.NewTable("LazVirtualStringTree_SetOnStructureChange", 0),
		/*187*/ imports.NewTable("LazVirtualStringTree_SetOnUpdating", 0),
		/*188*/ imports.NewTable("LazVirtualStringTree_StateImages", 0),
		/*189*/ imports.NewTable("LazVirtualStringTree_StateImagesWidth", 0),
		/*190*/ imports.NewTable("LazVirtualStringTree_TextMargin", 0),
		/*191*/ imports.NewTable("LazVirtualStringTree_TreeOptions", 0),
		/*192*/ imports.NewTable("LazVirtualStringTree_WantTabs", 0),
	}
)

func lazVirtualStringTreeImportAPI() *imports.Imports {
	if lazVirtualStringTreeImport == nil {
		lazVirtualStringTreeImport = NewDefaultImports()
		lazVirtualStringTreeImport.SetImportTable(lazVirtualStringTreeImportTables)
		lazVirtualStringTreeImportTables = nil
	}
	return lazVirtualStringTreeImport
}
