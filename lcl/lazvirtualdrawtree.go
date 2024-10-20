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

// ILazVirtualDrawTree Parent: ICustomVirtualDrawTree
type ILazVirtualDrawTree interface {
	ICustomVirtualDrawTree
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
	TreeOptions() IVirtualTreeOptions                                      // property
	SetTreeOptions(AValue IVirtualTreeOptions)                             // property
	WantTabs() bool                                                        // property
	SetWantTabs(AValue bool)                                               // property
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
	SetOnGetHintKind(fn TVTHintKindEvent)                                  // property event
	SetOnGetHintSize(fn TVTGetHintSizeEvent)                               // property event
	SetOnGetImageIndex(fn TVTGetImageEvent)                                // property event
	SetOnGetImageIndexEx(fn TVTGetImageExEvent)                            // property event
	SetOnGetLineStyle(fn TVTGetLineStyleEvent)                             // property event
	SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent)                       // property event
	SetOnGetNodeWidth(fn TVTGetNodeWidthEvent)                             // property event
	SetOnGetPopupMenu(fn TVTPopupEvent)                                    // property event
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

// TLazVirtualDrawTree Parent: TCustomVirtualDrawTree
type TLazVirtualDrawTree struct {
	TCustomVirtualDrawTree
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
	drawNodePtr                   uintptr
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
	getHelpContextPtr             uintptr
	getHintKindPtr                uintptr
	getHintSizePtr                uintptr
	getImageIndexPtr              uintptr
	getImageIndexExPtr            uintptr
	getLineStylePtr               uintptr
	getNodeDataSizePtr            uintptr
	getNodeWidthPtr               uintptr
	getPopupMenuPtr               uintptr
	headerClickPtr                uintptr
	headerDblClickPtr             uintptr
	headerDraggedPtr              uintptr
	headerDraggedOutPtr           uintptr
	headerDraggingPtr             uintptr
	headerDrawPtr                 uintptr
	headerDrawQueryElementsPtr    uintptr
	headerHeightTrackingPtr       uintptr
	headerHeightDblClickResizePtr uintptr
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
	mouseDownPtr                  uintptr
	mouseMovePtr                  uintptr
	mouseUpPtr                    uintptr
	mouseWheelPtr                 uintptr
	nodeClickPtr                  uintptr
	nodeCopiedPtr                 uintptr
	nodeCopyingPtr                uintptr
	nodeDblClickPtr               uintptr
	nodeExportPtr                 uintptr
	nodeHeightTrackingPtr         uintptr
	nodeHeightDblClickResizePtr   uintptr
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

func NewLazVirtualDrawTree(AOwner IComponent) ILazVirtualDrawTree {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(16, GetObjectUintptr(AOwner))
	return AsLazVirtualDrawTree(r1)
}

func (m *TLazVirtualDrawTree) LastDragEffect() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(39, m.Instance())
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) Alignment() TAlignment {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TLazVirtualDrawTree) SetAlignment(AValue TAlignment) {
	lazVirtualDrawTreeImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) AnimationDuration() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetAnimationDuration(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) AutoExpandDelay() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetAutoExpandDelay(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) AutoScrollDelay() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetAutoScrollDelay(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) AutoScrollInterval() TAutoScrollInterval {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TAutoScrollInterval(r1)
}

func (m *TLazVirtualDrawTree) SetAutoScrollInterval(AValue TAutoScrollInterval) {
	lazVirtualDrawTreeImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Background() IPicture {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TLazVirtualDrawTree) SetBackground(AValue IPicture) {
	lazVirtualDrawTreeImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) BackgroundOffsetX() int32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetBackgroundOffsetX(AValue int32) {
	lazVirtualDrawTreeImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) BackgroundOffsetY() int32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetBackgroundOffsetY(AValue int32) {
	lazVirtualDrawTreeImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) BottomSpace() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetBottomSpace(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ButtonFillMode() TVTButtonFillMode {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TVTButtonFillMode(r1)
}

func (m *TLazVirtualDrawTree) SetButtonFillMode(AValue TVTButtonFillMode) {
	lazVirtualDrawTreeImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ButtonStyle() TVTButtonStyle {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TVTButtonStyle(r1)
}

func (m *TLazVirtualDrawTree) SetButtonStyle(AValue TVTButtonStyle) {
	lazVirtualDrawTreeImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ChangeDelay() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetChangeDelay(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) CheckImageKind() TCheckImageKind {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TCheckImageKind(r1)
}

func (m *TLazVirtualDrawTree) SetCheckImageKind(AValue TCheckImageKind) {
	lazVirtualDrawTreeImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ClipboardFormats() IClipboardFormats {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return AsClipboardFormats(r1)
}

func (m *TLazVirtualDrawTree) SetClipboardFormats(AValue IClipboardFormats) {
	lazVirtualDrawTreeImportAPI().SysCallN(14, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) Colors() IVTColors {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return AsVTColors(r1)
}

func (m *TLazVirtualDrawTree) SetColors(AValue IVTColors) {
	lazVirtualDrawTreeImportAPI().SysCallN(15, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) CustomCheckImages() ICustomImageList {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualDrawTree) SetCustomCheckImages(AValue ICustomImageList) {
	lazVirtualDrawTreeImportAPI().SysCallN(17, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) DefaultNodeHeight() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetDefaultNodeHeight(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DefaultPasteMode() TVTNodeAttachMode {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TVTNodeAttachMode(r1)
}

func (m *TLazVirtualDrawTree) SetDefaultPasteMode(AValue TVTNodeAttachMode) {
	lazVirtualDrawTreeImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragCursor() TCursor {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLazVirtualDrawTree) SetDragCursor(AValue TCursor) {
	lazVirtualDrawTreeImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragHeight() int32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetDragHeight(AValue int32) {
	lazVirtualDrawTreeImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragKind() TDragKind {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TLazVirtualDrawTree) SetDragKind(AValue TDragKind) {
	lazVirtualDrawTreeImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragImageKind() TVTDragImageKind {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return TVTDragImageKind(r1)
}

func (m *TLazVirtualDrawTree) SetDragImageKind(AValue TVTDragImageKind) {
	lazVirtualDrawTreeImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragMode() TDragMode {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLazVirtualDrawTree) SetDragMode(AValue TDragMode) {
	lazVirtualDrawTreeImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragOperations() TDragOperations {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return TDragOperations(r1)
}

func (m *TLazVirtualDrawTree) SetDragOperations(AValue TDragOperations) {
	lazVirtualDrawTreeImportAPI().SysCallN(25, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragType() TVTDragType {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return TVTDragType(r1)
}

func (m *TLazVirtualDrawTree) SetDragType(AValue TVTDragType) {
	lazVirtualDrawTreeImportAPI().SysCallN(26, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragWidth() int32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetDragWidth(AValue int32) {
	lazVirtualDrawTreeImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DrawSelectionMode() TVTDrawSelectionMode {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return TVTDrawSelectionMode(r1)
}

func (m *TLazVirtualDrawTree) SetDrawSelectionMode(AValue TVTDrawSelectionMode) {
	lazVirtualDrawTreeImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) EditDelay() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetEditDelay(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Header() IVTHeader {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return AsVTHeader(r1)
}

func (m *TLazVirtualDrawTree) SetHeader(AValue IVTHeader) {
	lazVirtualDrawTreeImportAPI().SysCallN(30, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) HintMode() TVTHintMode {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return TVTHintMode(r1)
}

func (m *TLazVirtualDrawTree) SetHintMode(AValue TVTHintMode) {
	lazVirtualDrawTreeImportAPI().SysCallN(31, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) HotCursor() TCursor {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLazVirtualDrawTree) SetHotCursor(AValue TCursor) {
	lazVirtualDrawTreeImportAPI().SysCallN(32, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Images() ICustomImageList {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualDrawTree) SetImages(AValue ICustomImageList) {
	lazVirtualDrawTreeImportAPI().SysCallN(33, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) IncrementalSearch() TVTIncrementalSearch {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return TVTIncrementalSearch(r1)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearch(AValue TVTIncrementalSearch) {
	lazVirtualDrawTreeImportAPI().SysCallN(34, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) IncrementalSearchDirection() TVTSearchDirection {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return TVTSearchDirection(r1)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearchDirection(AValue TVTSearchDirection) {
	lazVirtualDrawTreeImportAPI().SysCallN(35, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) IncrementalSearchStart() TVTSearchStart {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return TVTSearchStart(r1)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearchStart(AValue TVTSearchStart) {
	lazVirtualDrawTreeImportAPI().SysCallN(36, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) IncrementalSearchTimeout() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearchTimeout(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(37, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Indent() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(38, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetIndent(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(38, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) LineMode() TVTLineMode {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(40, 0, m.Instance(), 0)
	return TVTLineMode(r1)
}

func (m *TLazVirtualDrawTree) SetLineMode(AValue TVTLineMode) {
	lazVirtualDrawTreeImportAPI().SysCallN(40, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) LineStyle() TVTLineStyle {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(41, 0, m.Instance(), 0)
	return TVTLineStyle(r1)
}

func (m *TLazVirtualDrawTree) SetLineStyle(AValue TVTLineStyle) {
	lazVirtualDrawTreeImportAPI().SysCallN(41, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Margin() int32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetMargin(AValue int32) {
	lazVirtualDrawTreeImportAPI().SysCallN(42, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) NodeAlignment() TVTNodeAlignment {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(43, 0, m.Instance(), 0)
	return TVTNodeAlignment(r1)
}

func (m *TLazVirtualDrawTree) SetNodeAlignment(AValue TVTNodeAlignment) {
	lazVirtualDrawTreeImportAPI().SysCallN(43, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) NodeDataSize() int32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(44, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetNodeDataSize(AValue int32) {
	lazVirtualDrawTreeImportAPI().SysCallN(44, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) OperationCanceled() bool {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(45, m.Instance())
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) ParentColor() bool {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(46, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) SetParentColor(AValue bool) {
	lazVirtualDrawTreeImportAPI().SysCallN(46, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualDrawTree) ParentFont() bool {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(47, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) SetParentFont(AValue bool) {
	lazVirtualDrawTreeImportAPI().SysCallN(47, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualDrawTree) ParentShowHint() bool {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(48, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) SetParentShowHint(AValue bool) {
	lazVirtualDrawTreeImportAPI().SysCallN(48, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualDrawTree) RootNodeCount() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(49, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetRootNodeCount(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(49, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ScrollBarOptions() IScrollBarOptions {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return AsScrollBarOptions(r1)
}

func (m *TLazVirtualDrawTree) SetScrollBarOptions(AValue IScrollBarOptions) {
	lazVirtualDrawTreeImportAPI().SysCallN(50, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) SelectionBlendFactor() Byte {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(51, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TLazVirtualDrawTree) SetSelectionBlendFactor(AValue Byte) {
	lazVirtualDrawTreeImportAPI().SysCallN(51, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) SelectionCurveRadius() uint32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(52, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetSelectionCurveRadius(AValue uint32) {
	lazVirtualDrawTreeImportAPI().SysCallN(52, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) StateImages() ICustomImageList {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(175, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualDrawTree) SetStateImages(AValue ICustomImageList) {
	lazVirtualDrawTreeImportAPI().SysCallN(175, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) TextMargin() int32 {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(176, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetTextMargin(AValue int32) {
	lazVirtualDrawTreeImportAPI().SysCallN(176, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) TreeOptions() IVirtualTreeOptions {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(177, 0, m.Instance(), 0)
	return AsVirtualTreeOptions(r1)
}

func (m *TLazVirtualDrawTree) SetTreeOptions(AValue IVirtualTreeOptions) {
	lazVirtualDrawTreeImportAPI().SysCallN(177, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) WantTabs() bool {
	r1 := lazVirtualDrawTreeImportAPI().SysCallN(178, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) SetWantTabs(AValue bool) {
	lazVirtualDrawTreeImportAPI().SysCallN(178, 1, m.Instance(), PascalBool(AValue))
}

func LazVirtualDrawTreeClass() TClass {
	ret := lazVirtualDrawTreeImportAPI().SysCallN(13)
	return TClass(ret)
}

func (m *TLazVirtualDrawTree) SetOnAddToSelection(fn TVTAddToSelectionEvent) {
	if m.addToSelectionPtr != 0 {
		RemoveEventElement(m.addToSelectionPtr)
	}
	m.addToSelectionPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(53, m.Instance(), m.addToSelectionPtr)
}

func (m *TLazVirtualDrawTree) SetOnAdvancedHeaderDraw(fn TVTAdvancedHeaderPaintEvent) {
	if m.advancedHeaderDrawPtr != 0 {
		RemoveEventElement(m.advancedHeaderDrawPtr)
	}
	m.advancedHeaderDrawPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(54, m.Instance(), m.advancedHeaderDrawPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterAutoFitColumn(fn TVTAfterAutoFitColumnEvent) {
	if m.afterAutoFitColumnPtr != 0 {
		RemoveEventElement(m.afterAutoFitColumnPtr)
	}
	m.afterAutoFitColumnPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(55, m.Instance(), m.afterAutoFitColumnPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterAutoFitColumns(fn TVTAfterAutoFitColumnsEvent) {
	if m.afterAutoFitColumnsPtr != 0 {
		RemoveEventElement(m.afterAutoFitColumnsPtr)
	}
	m.afterAutoFitColumnsPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(56, m.Instance(), m.afterAutoFitColumnsPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterCellPaint(fn TVTAfterCellPaintEvent) {
	if m.afterCellPaintPtr != 0 {
		RemoveEventElement(m.afterCellPaintPtr)
	}
	m.afterCellPaintPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(57, m.Instance(), m.afterCellPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterColumnExport(fn TVTColumnExportEvent) {
	if m.afterColumnExportPtr != 0 {
		RemoveEventElement(m.afterColumnExportPtr)
	}
	m.afterColumnExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(58, m.Instance(), m.afterColumnExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterColumnWidthTracking(fn TVTAfterColumnWidthTrackingEvent) {
	if m.afterColumnWidthTrackingPtr != 0 {
		RemoveEventElement(m.afterColumnWidthTrackingPtr)
	}
	m.afterColumnWidthTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(59, m.Instance(), m.afterColumnWidthTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterGetMaxColumnWidth(fn TVTAfterGetMaxColumnWidthEvent) {
	if m.afterGetMaxColumnWidthPtr != 0 {
		RemoveEventElement(m.afterGetMaxColumnWidthPtr)
	}
	m.afterGetMaxColumnWidthPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(60, m.Instance(), m.afterGetMaxColumnWidthPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterHeaderExport(fn TVTTreeExportEvent) {
	if m.afterHeaderExportPtr != 0 {
		RemoveEventElement(m.afterHeaderExportPtr)
	}
	m.afterHeaderExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(61, m.Instance(), m.afterHeaderExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterHeaderHeightTracking(fn TVTAfterHeaderHeightTrackingEvent) {
	if m.afterHeaderHeightTrackingPtr != 0 {
		RemoveEventElement(m.afterHeaderHeightTrackingPtr)
	}
	m.afterHeaderHeightTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(62, m.Instance(), m.afterHeaderHeightTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterItemErase(fn TVTAfterItemEraseEvent) {
	if m.afterItemErasePtr != 0 {
		RemoveEventElement(m.afterItemErasePtr)
	}
	m.afterItemErasePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(63, m.Instance(), m.afterItemErasePtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterItemPaint(fn TVTAfterItemPaintEvent) {
	if m.afterItemPaintPtr != 0 {
		RemoveEventElement(m.afterItemPaintPtr)
	}
	m.afterItemPaintPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(64, m.Instance(), m.afterItemPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterNodeExport(fn TVTNodeExportEvent) {
	if m.afterNodeExportPtr != 0 {
		RemoveEventElement(m.afterNodeExportPtr)
	}
	m.afterNodeExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(65, m.Instance(), m.afterNodeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterPaint(fn TVTPaintEvent) {
	if m.afterPaintPtr != 0 {
		RemoveEventElement(m.afterPaintPtr)
	}
	m.afterPaintPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(66, m.Instance(), m.afterPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterTreeExport(fn TVTTreeExportEvent) {
	if m.afterTreeExportPtr != 0 {
		RemoveEventElement(m.afterTreeExportPtr)
	}
	m.afterTreeExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(67, m.Instance(), m.afterTreeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeAutoFitColumn(fn TVTBeforeAutoFitColumnEvent) {
	if m.beforeAutoFitColumnPtr != 0 {
		RemoveEventElement(m.beforeAutoFitColumnPtr)
	}
	m.beforeAutoFitColumnPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(68, m.Instance(), m.beforeAutoFitColumnPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeAutoFitColumns(fn TVTBeforeAutoFitColumnsEvent) {
	if m.beforeAutoFitColumnsPtr != 0 {
		RemoveEventElement(m.beforeAutoFitColumnsPtr)
	}
	m.beforeAutoFitColumnsPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(69, m.Instance(), m.beforeAutoFitColumnsPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeCellPaint(fn TVTBeforeCellPaintEvent) {
	if m.beforeCellPaintPtr != 0 {
		RemoveEventElement(m.beforeCellPaintPtr)
	}
	m.beforeCellPaintPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(70, m.Instance(), m.beforeCellPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeColumnExport(fn TVTColumnExportEvent) {
	if m.beforeColumnExportPtr != 0 {
		RemoveEventElement(m.beforeColumnExportPtr)
	}
	m.beforeColumnExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(71, m.Instance(), m.beforeColumnExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeColumnWidthTracking(fn TVTBeforeColumnWidthTrackingEvent) {
	if m.beforeColumnWidthTrackingPtr != 0 {
		RemoveEventElement(m.beforeColumnWidthTrackingPtr)
	}
	m.beforeColumnWidthTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(72, m.Instance(), m.beforeColumnWidthTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeDrawTreeLine(fn TVTBeforeDrawLineImageEvent) {
	if m.beforeDrawTreeLinePtr != 0 {
		RemoveEventElement(m.beforeDrawTreeLinePtr)
	}
	m.beforeDrawTreeLinePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(73, m.Instance(), m.beforeDrawTreeLinePtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeGetMaxColumnWidth(fn TVTBeforeGetMaxColumnWidthEvent) {
	if m.beforeGetMaxColumnWidthPtr != 0 {
		RemoveEventElement(m.beforeGetMaxColumnWidthPtr)
	}
	m.beforeGetMaxColumnWidthPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(74, m.Instance(), m.beforeGetMaxColumnWidthPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeHeaderExport(fn TVTTreeExportEvent) {
	if m.beforeHeaderExportPtr != 0 {
		RemoveEventElement(m.beforeHeaderExportPtr)
	}
	m.beforeHeaderExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(75, m.Instance(), m.beforeHeaderExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeHeaderHeightTracking(fn TVTBeforeHeaderHeightTrackingEvent) {
	if m.beforeHeaderHeightTrackingPtr != 0 {
		RemoveEventElement(m.beforeHeaderHeightTrackingPtr)
	}
	m.beforeHeaderHeightTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(76, m.Instance(), m.beforeHeaderHeightTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeItemErase(fn TVTBeforeItemEraseEvent) {
	if m.beforeItemErasePtr != 0 {
		RemoveEventElement(m.beforeItemErasePtr)
	}
	m.beforeItemErasePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(77, m.Instance(), m.beforeItemErasePtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeItemPaint(fn TVTBeforeItemPaintEvent) {
	if m.beforeItemPaintPtr != 0 {
		RemoveEventElement(m.beforeItemPaintPtr)
	}
	m.beforeItemPaintPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(78, m.Instance(), m.beforeItemPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeNodeExport(fn TVTNodeExportEvent) {
	if m.beforeNodeExportPtr != 0 {
		RemoveEventElement(m.beforeNodeExportPtr)
	}
	m.beforeNodeExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(79, m.Instance(), m.beforeNodeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforePaint(fn TVTPaintEvent) {
	if m.beforePaintPtr != 0 {
		RemoveEventElement(m.beforePaintPtr)
	}
	m.beforePaintPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(80, m.Instance(), m.beforePaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeTreeExport(fn TVTTreeExportEvent) {
	if m.beforeTreeExportPtr != 0 {
		RemoveEventElement(m.beforeTreeExportPtr)
	}
	m.beforeTreeExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(81, m.Instance(), m.beforeTreeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnCanSplitterResizeColumn(fn TVTCanSplitterResizeColumnEvent) {
	if m.canSplitterResizeColumnPtr != 0 {
		RemoveEventElement(m.canSplitterResizeColumnPtr)
	}
	m.canSplitterResizeColumnPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(82, m.Instance(), m.canSplitterResizeColumnPtr)
}

func (m *TLazVirtualDrawTree) SetOnCanSplitterResizeHeader(fn TVTCanSplitterResizeHeaderEvent) {
	if m.canSplitterResizeHeaderPtr != 0 {
		RemoveEventElement(m.canSplitterResizeHeaderPtr)
	}
	m.canSplitterResizeHeaderPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(83, m.Instance(), m.canSplitterResizeHeaderPtr)
}

func (m *TLazVirtualDrawTree) SetOnCanSplitterResizeNode(fn TVTCanSplitterResizeNodeEvent) {
	if m.canSplitterResizeNodePtr != 0 {
		RemoveEventElement(m.canSplitterResizeNodePtr)
	}
	m.canSplitterResizeNodePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(84, m.Instance(), m.canSplitterResizeNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnChange(fn TVTChangeEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(85, m.Instance(), m.changePtr)
}

func (m *TLazVirtualDrawTree) SetOnChecked(fn TVTChangeEvent) {
	if m.checkedPtr != 0 {
		RemoveEventElement(m.checkedPtr)
	}
	m.checkedPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(86, m.Instance(), m.checkedPtr)
}

func (m *TLazVirtualDrawTree) SetOnChecking(fn TVTCheckChangingEvent) {
	if m.checkingPtr != 0 {
		RemoveEventElement(m.checkingPtr)
	}
	m.checkingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(87, m.Instance(), m.checkingPtr)
}

func (m *TLazVirtualDrawTree) SetOnCollapsed(fn TVTChangeEvent) {
	if m.collapsedPtr != 0 {
		RemoveEventElement(m.collapsedPtr)
	}
	m.collapsedPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(88, m.Instance(), m.collapsedPtr)
}

func (m *TLazVirtualDrawTree) SetOnCollapsing(fn TVTChangingEvent) {
	if m.collapsingPtr != 0 {
		RemoveEventElement(m.collapsingPtr)
	}
	m.collapsingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(89, m.Instance(), m.collapsingPtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnClick(fn TVTColumnClickEvent) {
	if m.columnClickPtr != 0 {
		RemoveEventElement(m.columnClickPtr)
	}
	m.columnClickPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(90, m.Instance(), m.columnClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnDblClick(fn TVTColumnDblClickEvent) {
	if m.columnDblClickPtr != 0 {
		RemoveEventElement(m.columnDblClickPtr)
	}
	m.columnDblClickPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(91, m.Instance(), m.columnDblClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnExport(fn TVTColumnExportEvent) {
	if m.columnExportPtr != 0 {
		RemoveEventElement(m.columnExportPtr)
	}
	m.columnExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(92, m.Instance(), m.columnExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnResize(fn TVTHeaderNotifyEvent) {
	if m.columnResizePtr != 0 {
		RemoveEventElement(m.columnResizePtr)
	}
	m.columnResizePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(93, m.Instance(), m.columnResizePtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnWidthDblClickResize(fn TVTColumnWidthDblClickResizeEvent) {
	if m.columnWidthDblClickResizePtr != 0 {
		RemoveEventElement(m.columnWidthDblClickResizePtr)
	}
	m.columnWidthDblClickResizePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(94, m.Instance(), m.columnWidthDblClickResizePtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnWidthTracking(fn TVTColumnWidthTrackingEvent) {
	if m.columnWidthTrackingPtr != 0 {
		RemoveEventElement(m.columnWidthTrackingPtr)
	}
	m.columnWidthTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(95, m.Instance(), m.columnWidthTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnCompareNodes(fn TVTCompareEvent) {
	if m.compareNodesPtr != 0 {
		RemoveEventElement(m.compareNodesPtr)
	}
	m.compareNodesPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(96, m.Instance(), m.compareNodesPtr)
}

func (m *TLazVirtualDrawTree) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(97, m.Instance(), m.contextPopupPtr)
}

func (m *TLazVirtualDrawTree) SetOnCreateDataObject(fn TVTCreateDataObjectEvent) {
	if m.createDataObjectPtr != 0 {
		RemoveEventElement(m.createDataObjectPtr)
	}
	m.createDataObjectPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(98, m.Instance(), m.createDataObjectPtr)
}

func (m *TLazVirtualDrawTree) SetOnCreateDragManager(fn TVTCreateDragManagerEvent) {
	if m.createDragManagerPtr != 0 {
		RemoveEventElement(m.createDragManagerPtr)
	}
	m.createDragManagerPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(99, m.Instance(), m.createDragManagerPtr)
}

func (m *TLazVirtualDrawTree) SetOnCreateEditor(fn TVTCreateEditorEvent) {
	if m.createEditorPtr != 0 {
		RemoveEventElement(m.createEditorPtr)
	}
	m.createEditorPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(100, m.Instance(), m.createEditorPtr)
}

func (m *TLazVirtualDrawTree) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(101, m.Instance(), m.dblClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnDragAllowed(fn TVTDragAllowedEvent) {
	if m.dragAllowedPtr != 0 {
		RemoveEventElement(m.dragAllowedPtr)
	}
	m.dragAllowedPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(102, m.Instance(), m.dragAllowedPtr)
}

func (m *TLazVirtualDrawTree) SetOnDragOver(fn TVTDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(104, m.Instance(), m.dragOverPtr)
}

func (m *TLazVirtualDrawTree) SetOnDragDrop(fn TVTDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(103, m.Instance(), m.dragDropPtr)
}

func (m *TLazVirtualDrawTree) SetOnDrawHint(fn TVTDrawHintEvent) {
	if m.drawHintPtr != 0 {
		RemoveEventElement(m.drawHintPtr)
	}
	m.drawHintPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(105, m.Instance(), m.drawHintPtr)
}

func (m *TLazVirtualDrawTree) SetOnDrawNode(fn TVTDrawNodeEvent) {
	if m.drawNodePtr != 0 {
		RemoveEventElement(m.drawNodePtr)
	}
	m.drawNodePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(106, m.Instance(), m.drawNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnEdited(fn TVTEditChangeEvent) {
	if m.editedPtr != 0 {
		RemoveEventElement(m.editedPtr)
	}
	m.editedPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(107, m.Instance(), m.editedPtr)
}

func (m *TLazVirtualDrawTree) SetOnEditing(fn TVTEditChangingEvent) {
	if m.editingPtr != 0 {
		RemoveEventElement(m.editingPtr)
	}
	m.editingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(108, m.Instance(), m.editingPtr)
}

func (m *TLazVirtualDrawTree) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(109, m.Instance(), m.endDockPtr)
}

func (m *TLazVirtualDrawTree) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(110, m.Instance(), m.endDragPtr)
}

func (m *TLazVirtualDrawTree) SetOnEndOperation(fn TVTOperationEvent) {
	if m.endOperationPtr != 0 {
		RemoveEventElement(m.endOperationPtr)
	}
	m.endOperationPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(111, m.Instance(), m.endOperationPtr)
}

func (m *TLazVirtualDrawTree) SetOnExpanded(fn TVTChangeEvent) {
	if m.expandedPtr != 0 {
		RemoveEventElement(m.expandedPtr)
	}
	m.expandedPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(112, m.Instance(), m.expandedPtr)
}

func (m *TLazVirtualDrawTree) SetOnExpanding(fn TVTChangingEvent) {
	if m.expandingPtr != 0 {
		RemoveEventElement(m.expandingPtr)
	}
	m.expandingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(113, m.Instance(), m.expandingPtr)
}

func (m *TLazVirtualDrawTree) SetOnFocusChanged(fn TVTFocusChangeEvent) {
	if m.focusChangedPtr != 0 {
		RemoveEventElement(m.focusChangedPtr)
	}
	m.focusChangedPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(114, m.Instance(), m.focusChangedPtr)
}

func (m *TLazVirtualDrawTree) SetOnFocusChanging(fn TVTFocusChangingEvent) {
	if m.focusChangingPtr != 0 {
		RemoveEventElement(m.focusChangingPtr)
	}
	m.focusChangingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(115, m.Instance(), m.focusChangingPtr)
}

func (m *TLazVirtualDrawTree) SetOnFreeNode(fn TVTFreeNodeEvent) {
	if m.freeNodePtr != 0 {
		RemoveEventElement(m.freeNodePtr)
	}
	m.freeNodePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(116, m.Instance(), m.freeNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnGetCellIsEmpty(fn TVTGetCellIsEmptyEvent) {
	if m.getCellIsEmptyPtr != 0 {
		RemoveEventElement(m.getCellIsEmptyPtr)
	}
	m.getCellIsEmptyPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(117, m.Instance(), m.getCellIsEmptyPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetCursor(fn TVTGetCursorEvent) {
	if m.getCursorPtr != 0 {
		RemoveEventElement(m.getCursorPtr)
	}
	m.getCursorPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(118, m.Instance(), m.getCursorPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetHeaderCursor(fn TVTGetHeaderCursorEvent) {
	if m.getHeaderCursorPtr != 0 {
		RemoveEventElement(m.getHeaderCursorPtr)
	}
	m.getHeaderCursorPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(119, m.Instance(), m.getHeaderCursorPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetHelpContext(fn TVTHelpContextEvent) {
	if m.getHelpContextPtr != 0 {
		RemoveEventElement(m.getHelpContextPtr)
	}
	m.getHelpContextPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(120, m.Instance(), m.getHelpContextPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetHintKind(fn TVTHintKindEvent) {
	if m.getHintKindPtr != 0 {
		RemoveEventElement(m.getHintKindPtr)
	}
	m.getHintKindPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(121, m.Instance(), m.getHintKindPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetHintSize(fn TVTGetHintSizeEvent) {
	if m.getHintSizePtr != 0 {
		RemoveEventElement(m.getHintSizePtr)
	}
	m.getHintSizePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(122, m.Instance(), m.getHintSizePtr)
}

func (m *TLazVirtualDrawTree) SetOnGetImageIndex(fn TVTGetImageEvent) {
	if m.getImageIndexPtr != 0 {
		RemoveEventElement(m.getImageIndexPtr)
	}
	m.getImageIndexPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(123, m.Instance(), m.getImageIndexPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetImageIndexEx(fn TVTGetImageExEvent) {
	if m.getImageIndexExPtr != 0 {
		RemoveEventElement(m.getImageIndexExPtr)
	}
	m.getImageIndexExPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(124, m.Instance(), m.getImageIndexExPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetLineStyle(fn TVTGetLineStyleEvent) {
	if m.getLineStylePtr != 0 {
		RemoveEventElement(m.getLineStylePtr)
	}
	m.getLineStylePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(125, m.Instance(), m.getLineStylePtr)
}

func (m *TLazVirtualDrawTree) SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent) {
	if m.getNodeDataSizePtr != 0 {
		RemoveEventElement(m.getNodeDataSizePtr)
	}
	m.getNodeDataSizePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(126, m.Instance(), m.getNodeDataSizePtr)
}

func (m *TLazVirtualDrawTree) SetOnGetNodeWidth(fn TVTGetNodeWidthEvent) {
	if m.getNodeWidthPtr != 0 {
		RemoveEventElement(m.getNodeWidthPtr)
	}
	m.getNodeWidthPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(127, m.Instance(), m.getNodeWidthPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetPopupMenu(fn TVTPopupEvent) {
	if m.getPopupMenuPtr != 0 {
		RemoveEventElement(m.getPopupMenuPtr)
	}
	m.getPopupMenuPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(128, m.Instance(), m.getPopupMenuPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderClick(fn TVTHeaderClickEvent) {
	if m.headerClickPtr != 0 {
		RemoveEventElement(m.headerClickPtr)
	}
	m.headerClickPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(129, m.Instance(), m.headerClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDblClick(fn TVTHeaderClickEvent) {
	if m.headerDblClickPtr != 0 {
		RemoveEventElement(m.headerDblClickPtr)
	}
	m.headerDblClickPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(130, m.Instance(), m.headerDblClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDragged(fn TVTHeaderDraggedEvent) {
	if m.headerDraggedPtr != 0 {
		RemoveEventElement(m.headerDraggedPtr)
	}
	m.headerDraggedPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(131, m.Instance(), m.headerDraggedPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent) {
	if m.headerDraggedOutPtr != 0 {
		RemoveEventElement(m.headerDraggedOutPtr)
	}
	m.headerDraggedOutPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(132, m.Instance(), m.headerDraggedOutPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDragging(fn TVTHeaderDraggingEvent) {
	if m.headerDraggingPtr != 0 {
		RemoveEventElement(m.headerDraggingPtr)
	}
	m.headerDraggingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(133, m.Instance(), m.headerDraggingPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDraw(fn TVTHeaderPaintEvent) {
	if m.headerDrawPtr != 0 {
		RemoveEventElement(m.headerDrawPtr)
	}
	m.headerDrawPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(134, m.Instance(), m.headerDrawPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent) {
	if m.headerDrawQueryElementsPtr != 0 {
		RemoveEventElement(m.headerDrawQueryElementsPtr)
	}
	m.headerDrawQueryElementsPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(135, m.Instance(), m.headerDrawQueryElementsPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent) {
	if m.headerHeightTrackingPtr != 0 {
		RemoveEventElement(m.headerHeightTrackingPtr)
	}
	m.headerHeightTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(137, m.Instance(), m.headerHeightTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) {
	if m.headerHeightDblClickResizePtr != 0 {
		RemoveEventElement(m.headerHeightDblClickResizePtr)
	}
	m.headerHeightDblClickResizePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(136, m.Instance(), m.headerHeightDblClickResizePtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderMouseDown(fn TVTHeaderMouseEvent) {
	if m.headerMouseDownPtr != 0 {
		RemoveEventElement(m.headerMouseDownPtr)
	}
	m.headerMouseDownPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(138, m.Instance(), m.headerMouseDownPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderMouseMove(fn TVTHeaderMouseMoveEvent) {
	if m.headerMouseMovePtr != 0 {
		RemoveEventElement(m.headerMouseMovePtr)
	}
	m.headerMouseMovePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(139, m.Instance(), m.headerMouseMovePtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderMouseUp(fn TVTHeaderMouseEvent) {
	if m.headerMouseUpPtr != 0 {
		RemoveEventElement(m.headerMouseUpPtr)
	}
	m.headerMouseUpPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(140, m.Instance(), m.headerMouseUpPtr)
}

func (m *TLazVirtualDrawTree) SetOnHotChange(fn TVTHotNodeChangeEvent) {
	if m.hotChangePtr != 0 {
		RemoveEventElement(m.hotChangePtr)
	}
	m.hotChangePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(141, m.Instance(), m.hotChangePtr)
}

func (m *TLazVirtualDrawTree) SetOnIncrementalSearch(fn TVTIncrementalSearchEvent) {
	if m.incrementalSearchPtr != 0 {
		RemoveEventElement(m.incrementalSearchPtr)
	}
	m.incrementalSearchPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(142, m.Instance(), m.incrementalSearchPtr)
}

func (m *TLazVirtualDrawTree) SetOnInitChildren(fn TVTInitChildrenEvent) {
	if m.initChildrenPtr != 0 {
		RemoveEventElement(m.initChildrenPtr)
	}
	m.initChildrenPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(143, m.Instance(), m.initChildrenPtr)
}

func (m *TLazVirtualDrawTree) SetOnInitNode(fn TVTInitNodeEvent) {
	if m.initNodePtr != 0 {
		RemoveEventElement(m.initNodePtr)
	}
	m.initNodePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(144, m.Instance(), m.initNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnKeyAction(fn TVTKeyActionEvent) {
	if m.keyActionPtr != 0 {
		RemoveEventElement(m.keyActionPtr)
	}
	m.keyActionPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(145, m.Instance(), m.keyActionPtr)
}

func (m *TLazVirtualDrawTree) SetOnLoadNode(fn TVTSaveNodeEvent) {
	if m.loadNodePtr != 0 {
		RemoveEventElement(m.loadNodePtr)
	}
	m.loadNodePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(146, m.Instance(), m.loadNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnLoadTree(fn TVTSaveTreeEvent) {
	if m.loadTreePtr != 0 {
		RemoveEventElement(m.loadTreePtr)
	}
	m.loadTreePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(147, m.Instance(), m.loadTreePtr)
}

func (m *TLazVirtualDrawTree) SetOnMeasureItem(fn TVTMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(148, m.Instance(), m.measureItemPtr)
}

func (m *TLazVirtualDrawTree) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(149, m.Instance(), m.mouseDownPtr)
}

func (m *TLazVirtualDrawTree) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(150, m.Instance(), m.mouseMovePtr)
}

func (m *TLazVirtualDrawTree) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(151, m.Instance(), m.mouseUpPtr)
}

func (m *TLazVirtualDrawTree) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(152, m.Instance(), m.mouseWheelPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeClick(fn TVTNodeClickEvent) {
	if m.nodeClickPtr != 0 {
		RemoveEventElement(m.nodeClickPtr)
	}
	m.nodeClickPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(153, m.Instance(), m.nodeClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeCopied(fn TVTNodeCopiedEvent) {
	if m.nodeCopiedPtr != 0 {
		RemoveEventElement(m.nodeCopiedPtr)
	}
	m.nodeCopiedPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(154, m.Instance(), m.nodeCopiedPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeCopying(fn TVTNodeCopyingEvent) {
	if m.nodeCopyingPtr != 0 {
		RemoveEventElement(m.nodeCopyingPtr)
	}
	m.nodeCopyingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(155, m.Instance(), m.nodeCopyingPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeDblClick(fn TVTNodeClickEvent) {
	if m.nodeDblClickPtr != 0 {
		RemoveEventElement(m.nodeDblClickPtr)
	}
	m.nodeDblClickPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(156, m.Instance(), m.nodeDblClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeExport(fn TVTNodeExportEvent) {
	if m.nodeExportPtr != 0 {
		RemoveEventElement(m.nodeExportPtr)
	}
	m.nodeExportPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(157, m.Instance(), m.nodeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent) {
	if m.nodeHeightTrackingPtr != 0 {
		RemoveEventElement(m.nodeHeightTrackingPtr)
	}
	m.nodeHeightTrackingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(159, m.Instance(), m.nodeHeightTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent) {
	if m.nodeHeightDblClickResizePtr != 0 {
		RemoveEventElement(m.nodeHeightDblClickResizePtr)
	}
	m.nodeHeightDblClickResizePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(158, m.Instance(), m.nodeHeightDblClickResizePtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeMoved(fn TVTNodeMovedEvent) {
	if m.nodeMovedPtr != 0 {
		RemoveEventElement(m.nodeMovedPtr)
	}
	m.nodeMovedPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(160, m.Instance(), m.nodeMovedPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeMoving(fn TVTNodeMovingEvent) {
	if m.nodeMovingPtr != 0 {
		RemoveEventElement(m.nodeMovingPtr)
	}
	m.nodeMovingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(161, m.Instance(), m.nodeMovingPtr)
}

func (m *TLazVirtualDrawTree) SetOnPaintBackground(fn TVTBackgroundPaintEvent) {
	if m.paintBackgroundPtr != 0 {
		RemoveEventElement(m.paintBackgroundPtr)
	}
	m.paintBackgroundPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(162, m.Instance(), m.paintBackgroundPtr)
}

func (m *TLazVirtualDrawTree) SetOnRemoveFromSelection(fn TVTRemoveFromSelectionEvent) {
	if m.removeFromSelectionPtr != 0 {
		RemoveEventElement(m.removeFromSelectionPtr)
	}
	m.removeFromSelectionPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(163, m.Instance(), m.removeFromSelectionPtr)
}

func (m *TLazVirtualDrawTree) SetOnResetNode(fn TVTChangeEvent) {
	if m.resetNodePtr != 0 {
		RemoveEventElement(m.resetNodePtr)
	}
	m.resetNodePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(164, m.Instance(), m.resetNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnSaveNode(fn TVTSaveNodeEvent) {
	if m.saveNodePtr != 0 {
		RemoveEventElement(m.saveNodePtr)
	}
	m.saveNodePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(165, m.Instance(), m.saveNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnSaveTree(fn TVTSaveTreeEvent) {
	if m.saveTreePtr != 0 {
		RemoveEventElement(m.saveTreePtr)
	}
	m.saveTreePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(166, m.Instance(), m.saveTreePtr)
}

func (m *TLazVirtualDrawTree) SetOnScroll(fn TVTScrollEvent) {
	if m.scrollPtr != 0 {
		RemoveEventElement(m.scrollPtr)
	}
	m.scrollPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(167, m.Instance(), m.scrollPtr)
}

func (m *TLazVirtualDrawTree) SetOnShowScrollBar(fn TVTScrollBarShowEvent) {
	if m.showScrollBarPtr != 0 {
		RemoveEventElement(m.showScrollBarPtr)
	}
	m.showScrollBarPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(168, m.Instance(), m.showScrollBarPtr)
}

func (m *TLazVirtualDrawTree) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(169, m.Instance(), m.startDockPtr)
}

func (m *TLazVirtualDrawTree) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(170, m.Instance(), m.startDragPtr)
}

func (m *TLazVirtualDrawTree) SetOnStartOperation(fn TVTOperationEvent) {
	if m.startOperationPtr != 0 {
		RemoveEventElement(m.startOperationPtr)
	}
	m.startOperationPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(171, m.Instance(), m.startOperationPtr)
}

func (m *TLazVirtualDrawTree) SetOnStateChange(fn TVTStateChangeEvent) {
	if m.stateChangePtr != 0 {
		RemoveEventElement(m.stateChangePtr)
	}
	m.stateChangePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(172, m.Instance(), m.stateChangePtr)
}

func (m *TLazVirtualDrawTree) SetOnStructureChange(fn TVTStructureChangeEvent) {
	if m.structureChangePtr != 0 {
		RemoveEventElement(m.structureChangePtr)
	}
	m.structureChangePtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(173, m.Instance(), m.structureChangePtr)
}

func (m *TLazVirtualDrawTree) SetOnUpdating(fn TVTUpdatingEvent) {
	if m.updatingPtr != 0 {
		RemoveEventElement(m.updatingPtr)
	}
	m.updatingPtr = MakeEventDataPtr(fn)
	lazVirtualDrawTreeImportAPI().SysCallN(174, m.Instance(), m.updatingPtr)
}

var (
	lazVirtualDrawTreeImport       *imports.Imports = nil
	lazVirtualDrawTreeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazVirtualDrawTree_Alignment", 0),
		/*1*/ imports.NewTable("LazVirtualDrawTree_AnimationDuration", 0),
		/*2*/ imports.NewTable("LazVirtualDrawTree_AutoExpandDelay", 0),
		/*3*/ imports.NewTable("LazVirtualDrawTree_AutoScrollDelay", 0),
		/*4*/ imports.NewTable("LazVirtualDrawTree_AutoScrollInterval", 0),
		/*5*/ imports.NewTable("LazVirtualDrawTree_Background", 0),
		/*6*/ imports.NewTable("LazVirtualDrawTree_BackgroundOffsetX", 0),
		/*7*/ imports.NewTable("LazVirtualDrawTree_BackgroundOffsetY", 0),
		/*8*/ imports.NewTable("LazVirtualDrawTree_BottomSpace", 0),
		/*9*/ imports.NewTable("LazVirtualDrawTree_ButtonFillMode", 0),
		/*10*/ imports.NewTable("LazVirtualDrawTree_ButtonStyle", 0),
		/*11*/ imports.NewTable("LazVirtualDrawTree_ChangeDelay", 0),
		/*12*/ imports.NewTable("LazVirtualDrawTree_CheckImageKind", 0),
		/*13*/ imports.NewTable("LazVirtualDrawTree_Class", 0),
		/*14*/ imports.NewTable("LazVirtualDrawTree_ClipboardFormats", 0),
		/*15*/ imports.NewTable("LazVirtualDrawTree_Colors", 0),
		/*16*/ imports.NewTable("LazVirtualDrawTree_Create", 0),
		/*17*/ imports.NewTable("LazVirtualDrawTree_CustomCheckImages", 0),
		/*18*/ imports.NewTable("LazVirtualDrawTree_DefaultNodeHeight", 0),
		/*19*/ imports.NewTable("LazVirtualDrawTree_DefaultPasteMode", 0),
		/*20*/ imports.NewTable("LazVirtualDrawTree_DragCursor", 0),
		/*21*/ imports.NewTable("LazVirtualDrawTree_DragHeight", 0),
		/*22*/ imports.NewTable("LazVirtualDrawTree_DragImageKind", 0),
		/*23*/ imports.NewTable("LazVirtualDrawTree_DragKind", 0),
		/*24*/ imports.NewTable("LazVirtualDrawTree_DragMode", 0),
		/*25*/ imports.NewTable("LazVirtualDrawTree_DragOperations", 0),
		/*26*/ imports.NewTable("LazVirtualDrawTree_DragType", 0),
		/*27*/ imports.NewTable("LazVirtualDrawTree_DragWidth", 0),
		/*28*/ imports.NewTable("LazVirtualDrawTree_DrawSelectionMode", 0),
		/*29*/ imports.NewTable("LazVirtualDrawTree_EditDelay", 0),
		/*30*/ imports.NewTable("LazVirtualDrawTree_Header", 0),
		/*31*/ imports.NewTable("LazVirtualDrawTree_HintMode", 0),
		/*32*/ imports.NewTable("LazVirtualDrawTree_HotCursor", 0),
		/*33*/ imports.NewTable("LazVirtualDrawTree_Images", 0),
		/*34*/ imports.NewTable("LazVirtualDrawTree_IncrementalSearch", 0),
		/*35*/ imports.NewTable("LazVirtualDrawTree_IncrementalSearchDirection", 0),
		/*36*/ imports.NewTable("LazVirtualDrawTree_IncrementalSearchStart", 0),
		/*37*/ imports.NewTable("LazVirtualDrawTree_IncrementalSearchTimeout", 0),
		/*38*/ imports.NewTable("LazVirtualDrawTree_Indent", 0),
		/*39*/ imports.NewTable("LazVirtualDrawTree_LastDragEffect", 0),
		/*40*/ imports.NewTable("LazVirtualDrawTree_LineMode", 0),
		/*41*/ imports.NewTable("LazVirtualDrawTree_LineStyle", 0),
		/*42*/ imports.NewTable("LazVirtualDrawTree_Margin", 0),
		/*43*/ imports.NewTable("LazVirtualDrawTree_NodeAlignment", 0),
		/*44*/ imports.NewTable("LazVirtualDrawTree_NodeDataSize", 0),
		/*45*/ imports.NewTable("LazVirtualDrawTree_OperationCanceled", 0),
		/*46*/ imports.NewTable("LazVirtualDrawTree_ParentColor", 0),
		/*47*/ imports.NewTable("LazVirtualDrawTree_ParentFont", 0),
		/*48*/ imports.NewTable("LazVirtualDrawTree_ParentShowHint", 0),
		/*49*/ imports.NewTable("LazVirtualDrawTree_RootNodeCount", 0),
		/*50*/ imports.NewTable("LazVirtualDrawTree_ScrollBarOptions", 0),
		/*51*/ imports.NewTable("LazVirtualDrawTree_SelectionBlendFactor", 0),
		/*52*/ imports.NewTable("LazVirtualDrawTree_SelectionCurveRadius", 0),
		/*53*/ imports.NewTable("LazVirtualDrawTree_SetOnAddToSelection", 0),
		/*54*/ imports.NewTable("LazVirtualDrawTree_SetOnAdvancedHeaderDraw", 0),
		/*55*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterAutoFitColumn", 0),
		/*56*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterAutoFitColumns", 0),
		/*57*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterCellPaint", 0),
		/*58*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterColumnExport", 0),
		/*59*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterColumnWidthTracking", 0),
		/*60*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterGetMaxColumnWidth", 0),
		/*61*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterHeaderExport", 0),
		/*62*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterHeaderHeightTracking", 0),
		/*63*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterItemErase", 0),
		/*64*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterItemPaint", 0),
		/*65*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterNodeExport", 0),
		/*66*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterPaint", 0),
		/*67*/ imports.NewTable("LazVirtualDrawTree_SetOnAfterTreeExport", 0),
		/*68*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeAutoFitColumn", 0),
		/*69*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeAutoFitColumns", 0),
		/*70*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeCellPaint", 0),
		/*71*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeColumnExport", 0),
		/*72*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeColumnWidthTracking", 0),
		/*73*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeDrawTreeLine", 0),
		/*74*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeGetMaxColumnWidth", 0),
		/*75*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeHeaderExport", 0),
		/*76*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeHeaderHeightTracking", 0),
		/*77*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeItemErase", 0),
		/*78*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeItemPaint", 0),
		/*79*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeNodeExport", 0),
		/*80*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforePaint", 0),
		/*81*/ imports.NewTable("LazVirtualDrawTree_SetOnBeforeTreeExport", 0),
		/*82*/ imports.NewTable("LazVirtualDrawTree_SetOnCanSplitterResizeColumn", 0),
		/*83*/ imports.NewTable("LazVirtualDrawTree_SetOnCanSplitterResizeHeader", 0),
		/*84*/ imports.NewTable("LazVirtualDrawTree_SetOnCanSplitterResizeNode", 0),
		/*85*/ imports.NewTable("LazVirtualDrawTree_SetOnChange", 0),
		/*86*/ imports.NewTable("LazVirtualDrawTree_SetOnChecked", 0),
		/*87*/ imports.NewTable("LazVirtualDrawTree_SetOnChecking", 0),
		/*88*/ imports.NewTable("LazVirtualDrawTree_SetOnCollapsed", 0),
		/*89*/ imports.NewTable("LazVirtualDrawTree_SetOnCollapsing", 0),
		/*90*/ imports.NewTable("LazVirtualDrawTree_SetOnColumnClick", 0),
		/*91*/ imports.NewTable("LazVirtualDrawTree_SetOnColumnDblClick", 0),
		/*92*/ imports.NewTable("LazVirtualDrawTree_SetOnColumnExport", 0),
		/*93*/ imports.NewTable("LazVirtualDrawTree_SetOnColumnResize", 0),
		/*94*/ imports.NewTable("LazVirtualDrawTree_SetOnColumnWidthDblClickResize", 0),
		/*95*/ imports.NewTable("LazVirtualDrawTree_SetOnColumnWidthTracking", 0),
		/*96*/ imports.NewTable("LazVirtualDrawTree_SetOnCompareNodes", 0),
		/*97*/ imports.NewTable("LazVirtualDrawTree_SetOnContextPopup", 0),
		/*98*/ imports.NewTable("LazVirtualDrawTree_SetOnCreateDataObject", 0),
		/*99*/ imports.NewTable("LazVirtualDrawTree_SetOnCreateDragManager", 0),
		/*100*/ imports.NewTable("LazVirtualDrawTree_SetOnCreateEditor", 0),
		/*101*/ imports.NewTable("LazVirtualDrawTree_SetOnDblClick", 0),
		/*102*/ imports.NewTable("LazVirtualDrawTree_SetOnDragAllowed", 0),
		/*103*/ imports.NewTable("LazVirtualDrawTree_SetOnDragDrop", 0),
		/*104*/ imports.NewTable("LazVirtualDrawTree_SetOnDragOver", 0),
		/*105*/ imports.NewTable("LazVirtualDrawTree_SetOnDrawHint", 0),
		/*106*/ imports.NewTable("LazVirtualDrawTree_SetOnDrawNode", 0),
		/*107*/ imports.NewTable("LazVirtualDrawTree_SetOnEdited", 0),
		/*108*/ imports.NewTable("LazVirtualDrawTree_SetOnEditing", 0),
		/*109*/ imports.NewTable("LazVirtualDrawTree_SetOnEndDock", 0),
		/*110*/ imports.NewTable("LazVirtualDrawTree_SetOnEndDrag", 0),
		/*111*/ imports.NewTable("LazVirtualDrawTree_SetOnEndOperation", 0),
		/*112*/ imports.NewTable("LazVirtualDrawTree_SetOnExpanded", 0),
		/*113*/ imports.NewTable("LazVirtualDrawTree_SetOnExpanding", 0),
		/*114*/ imports.NewTable("LazVirtualDrawTree_SetOnFocusChanged", 0),
		/*115*/ imports.NewTable("LazVirtualDrawTree_SetOnFocusChanging", 0),
		/*116*/ imports.NewTable("LazVirtualDrawTree_SetOnFreeNode", 0),
		/*117*/ imports.NewTable("LazVirtualDrawTree_SetOnGetCellIsEmpty", 0),
		/*118*/ imports.NewTable("LazVirtualDrawTree_SetOnGetCursor", 0),
		/*119*/ imports.NewTable("LazVirtualDrawTree_SetOnGetHeaderCursor", 0),
		/*120*/ imports.NewTable("LazVirtualDrawTree_SetOnGetHelpContext", 0),
		/*121*/ imports.NewTable("LazVirtualDrawTree_SetOnGetHintKind", 0),
		/*122*/ imports.NewTable("LazVirtualDrawTree_SetOnGetHintSize", 0),
		/*123*/ imports.NewTable("LazVirtualDrawTree_SetOnGetImageIndex", 0),
		/*124*/ imports.NewTable("LazVirtualDrawTree_SetOnGetImageIndexEx", 0),
		/*125*/ imports.NewTable("LazVirtualDrawTree_SetOnGetLineStyle", 0),
		/*126*/ imports.NewTable("LazVirtualDrawTree_SetOnGetNodeDataSize", 0),
		/*127*/ imports.NewTable("LazVirtualDrawTree_SetOnGetNodeWidth", 0),
		/*128*/ imports.NewTable("LazVirtualDrawTree_SetOnGetPopupMenu", 0),
		/*129*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderClick", 0),
		/*130*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderDblClick", 0),
		/*131*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderDragged", 0),
		/*132*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderDraggedOut", 0),
		/*133*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderDragging", 0),
		/*134*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderDraw", 0),
		/*135*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderDrawQueryElements", 0),
		/*136*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderHeightDblClickResize", 0),
		/*137*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderHeightTracking", 0),
		/*138*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderMouseDown", 0),
		/*139*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderMouseMove", 0),
		/*140*/ imports.NewTable("LazVirtualDrawTree_SetOnHeaderMouseUp", 0),
		/*141*/ imports.NewTable("LazVirtualDrawTree_SetOnHotChange", 0),
		/*142*/ imports.NewTable("LazVirtualDrawTree_SetOnIncrementalSearch", 0),
		/*143*/ imports.NewTable("LazVirtualDrawTree_SetOnInitChildren", 0),
		/*144*/ imports.NewTable("LazVirtualDrawTree_SetOnInitNode", 0),
		/*145*/ imports.NewTable("LazVirtualDrawTree_SetOnKeyAction", 0),
		/*146*/ imports.NewTable("LazVirtualDrawTree_SetOnLoadNode", 0),
		/*147*/ imports.NewTable("LazVirtualDrawTree_SetOnLoadTree", 0),
		/*148*/ imports.NewTable("LazVirtualDrawTree_SetOnMeasureItem", 0),
		/*149*/ imports.NewTable("LazVirtualDrawTree_SetOnMouseDown", 0),
		/*150*/ imports.NewTable("LazVirtualDrawTree_SetOnMouseMove", 0),
		/*151*/ imports.NewTable("LazVirtualDrawTree_SetOnMouseUp", 0),
		/*152*/ imports.NewTable("LazVirtualDrawTree_SetOnMouseWheel", 0),
		/*153*/ imports.NewTable("LazVirtualDrawTree_SetOnNodeClick", 0),
		/*154*/ imports.NewTable("LazVirtualDrawTree_SetOnNodeCopied", 0),
		/*155*/ imports.NewTable("LazVirtualDrawTree_SetOnNodeCopying", 0),
		/*156*/ imports.NewTable("LazVirtualDrawTree_SetOnNodeDblClick", 0),
		/*157*/ imports.NewTable("LazVirtualDrawTree_SetOnNodeExport", 0),
		/*158*/ imports.NewTable("LazVirtualDrawTree_SetOnNodeHeightDblClickResize", 0),
		/*159*/ imports.NewTable("LazVirtualDrawTree_SetOnNodeHeightTracking", 0),
		/*160*/ imports.NewTable("LazVirtualDrawTree_SetOnNodeMoved", 0),
		/*161*/ imports.NewTable("LazVirtualDrawTree_SetOnNodeMoving", 0),
		/*162*/ imports.NewTable("LazVirtualDrawTree_SetOnPaintBackground", 0),
		/*163*/ imports.NewTable("LazVirtualDrawTree_SetOnRemoveFromSelection", 0),
		/*164*/ imports.NewTable("LazVirtualDrawTree_SetOnResetNode", 0),
		/*165*/ imports.NewTable("LazVirtualDrawTree_SetOnSaveNode", 0),
		/*166*/ imports.NewTable("LazVirtualDrawTree_SetOnSaveTree", 0),
		/*167*/ imports.NewTable("LazVirtualDrawTree_SetOnScroll", 0),
		/*168*/ imports.NewTable("LazVirtualDrawTree_SetOnShowScrollBar", 0),
		/*169*/ imports.NewTable("LazVirtualDrawTree_SetOnStartDock", 0),
		/*170*/ imports.NewTable("LazVirtualDrawTree_SetOnStartDrag", 0),
		/*171*/ imports.NewTable("LazVirtualDrawTree_SetOnStartOperation", 0),
		/*172*/ imports.NewTable("LazVirtualDrawTree_SetOnStateChange", 0),
		/*173*/ imports.NewTable("LazVirtualDrawTree_SetOnStructureChange", 0),
		/*174*/ imports.NewTable("LazVirtualDrawTree_SetOnUpdating", 0),
		/*175*/ imports.NewTable("LazVirtualDrawTree_StateImages", 0),
		/*176*/ imports.NewTable("LazVirtualDrawTree_TextMargin", 0),
		/*177*/ imports.NewTable("LazVirtualDrawTree_TreeOptions", 0),
		/*178*/ imports.NewTable("LazVirtualDrawTree_WantTabs", 0),
	}
)

func lazVirtualDrawTreeImportAPI() *imports.Imports {
	if lazVirtualDrawTreeImport == nil {
		lazVirtualDrawTreeImport = NewDefaultImports()
		lazVirtualDrawTreeImport.SetImportTable(lazVirtualDrawTreeImportTables)
		lazVirtualDrawTreeImportTables = nil
	}
	return lazVirtualDrawTreeImport
}
