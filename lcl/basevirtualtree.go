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

// IBaseVirtualTree Parent: ICustomControl
type IBaseVirtualTree interface {
	ICustomControl
	AbsoluteIndex(node types.PVirtualNode) uint32                                                                                                                                   // function
	AddChild(parent types.PVirtualNode, userData uintptr) types.PVirtualNode                                                                                                        // function
	CancelEditNode() bool                                                                                                                                                           // function
	CanEdit(node types.PVirtualNode, column int32) bool                                                                                                                             // function
	CopyToWithPVirtualNodeBaseVirtualTreeVTNodeAttachModeBool(source types.PVirtualNode, tree IBaseVirtualTree, mode types.TVTNodeAttachMode, childrenOnly bool) types.PVirtualNode // function
	CopyToWithPVirtualNodeX2VTNodeAttachModeBool(source types.PVirtualNode, target types.PVirtualNode, mode types.TVTNodeAttachMode, childrenOnly bool) types.PVirtualNode          // function
	DraggingToBool() bool                                                                                                                                                           // function
	EditNode(node types.PVirtualNode, column int32) bool                                                                                                                            // function
	EndEditNode() bool                                                                                                                                                              // function
	GetDisplayRect(node types.PVirtualNode, column int32, textOnly bool, unclipped bool, applyCellContentMargin bool) types.TRect                                                   // function
	GetEffectivelyFiltered(node types.PVirtualNode) bool                                                                                                                            // function
	GetEffectivelyVisible(node types.PVirtualNode) bool                                                                                                                             // function
	GetFirst(considerChildrenAbove bool) types.PVirtualNode                                                                                                                         // function
	GetFirstChecked(state types.TCheckState, considerChildrenAbove bool) types.PVirtualNode                                                                                         // function
	GetFirstChild(node types.PVirtualNode) types.PVirtualNode                                                                                                                       // function
	GetFirstChildNoInit(node types.PVirtualNode) types.PVirtualNode                                                                                                                 // function
	GetFirstCutCopy(considerChildrenAbove bool) types.PVirtualNode                                                                                                                  // function
	GetFirstInitialized(considerChildrenAbove bool) types.PVirtualNode                                                                                                              // function
	GetFirstLeaf() types.PVirtualNode                                                                                                                                               // function
	GetFirstLevel(nodeLevel uint32) types.PVirtualNode                                                                                                                              // function
	GetFirstNoInit(considerChildrenAbove bool) types.PVirtualNode                                                                                                                   // function
	GetFirstSelected(considerChildrenAbove bool) types.PVirtualNode                                                                                                                 // function
	GetFirstVisible(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) types.PVirtualNode                                                                   // function
	GetFirstVisibleChild(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode                                                                                          // function
	GetFirstVisibleChildNoInit(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode                                                                                    // function
	GetFirstVisibleNoInit(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) types.PVirtualNode                                                             // function
	GetLast(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                                 // function
	GetLastInitialized(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                      // function
	GetLastNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                           // function
	GetLastChild(node types.PVirtualNode) types.PVirtualNode                                                                                                                        // function
	GetLastChildNoInit(node types.PVirtualNode) types.PVirtualNode                                                                                                                  // function
	GetLastVisible(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) types.PVirtualNode                                                                    // function
	GetLastVisibleChild(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode                                                                                           // function
	GetLastVisibleChildNoInit(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode                                                                                     // function
	GetLastVisibleNoInit(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) types.PVirtualNode                                                              // function
	GetMaxColumnWidth(column int32, useSmartColumnWidth bool) int32                                                                                                                 // function
	GetNext(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                                 // function
	GetNextCheckedWithPVirtualNodeCheckStateBool(node types.PVirtualNode, state types.TCheckState, considerChildrenAbove bool) types.PVirtualNode                                   // function
	GetNextCheckedWithPVirtualNodeBool(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                      // function
	GetNextCutCopy(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                          // function
	GetNextInitialized(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                      // function
	GetNextLeaf(node types.PVirtualNode) types.PVirtualNode                                                                                                                         // function
	GetNextLevel(node types.PVirtualNode, nodeLevel uint32) types.PVirtualNode                                                                                                      // function
	GetNextNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                           // function
	GetNextSelected(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                         // function
	GetNextSibling(node types.PVirtualNode) types.PVirtualNode                                                                                                                      // function
	GetNextSiblingNoInit(node types.PVirtualNode) types.PVirtualNode                                                                                                                // function
	GetNextVisible(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                          // function
	GetNextVisibleNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                    // function
	GetNextVisibleSibling(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode                                                                                         // function
	GetNextVisibleSiblingNoInit(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode                                                                                   // function
	GetNodeAtWithPoint(P types.TPoint) types.PVirtualNode                                                                                                                           // function
	GetNodeAtWithIntX2(X int32, Y int32) types.PVirtualNode                                                                                                                         // function
	GetNodeAtWithIntX3Bool(X int32, Y int32, relative bool, nodeTop *int32) types.PVirtualNode                                                                                      // function
	GetNodeData(node types.PVirtualNode) uintptr                                                                                                                                    // function
	GetNodeLevel(node types.PVirtualNode) uint32                                                                                                                                    // function
	GetPrevious(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                             // function
	GetPreviousChecked(node types.PVirtualNode, state types.TCheckState, considerChildrenAbove bool) types.PVirtualNode                                                             // function
	GetPreviousCutCopy(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                      // function
	GetPreviousInitialized(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                  // function
	GetPreviousLeaf(node types.PVirtualNode) types.PVirtualNode                                                                                                                     // function
	GetPreviousLevel(node types.PVirtualNode, nodeLevel uint32) types.PVirtualNode                                                                                                  // function
	GetPreviousNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                       // function
	GetPreviousSelected(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                     // function
	GetPreviousSibling(node types.PVirtualNode) types.PVirtualNode                                                                                                                  // function
	GetPreviousSiblingNoInit(node types.PVirtualNode) types.PVirtualNode                                                                                                            // function
	GetPreviousVisible(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                      // function
	GetPreviousVisibleNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode                                                                                // function
	GetPreviousVisibleSibling(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode                                                                                     // function
	GetPreviousVisibleSiblingNoInit(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode                                                                               // function
	GetSortedCutCopySet(resolve bool) types.TNodeArray                                                                                                                              // function
	GetSortedSelection(resolve bool) types.TNodeArray                                                                                                                               // function
	GetTreeRect() types.TRect                                                                                                                                                       // function
	GetVisibleParent(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode                                                                                              // function
	HasAsParent(node types.PVirtualNode, potentialParent types.PVirtualNode) bool                                                                                                   // function
	InsertNode(node types.PVirtualNode, mode types.TVTNodeAttachMode, userData uintptr) types.PVirtualNode                                                                          // function
	InvalidateNode(node types.PVirtualNode) types.TRect                                                                                                                             // function
	IsEditing() bool                                                                                                                                                                // function
	IsMouseSelecting() bool                                                                                                                                                         // function
	IsEmpty() bool                                                                                                                                                                  // function
	PasteFromClipboard() bool                                                                                                                                                       // function
	ScrollIntoViewWithPVirtualNodeBoolX2(node types.PVirtualNode, center bool, horizontally bool) bool                                                                              // function
	ScrollIntoViewWithColumnIndexBool(column int32, center bool) bool                                                                                                               // function
	// Nodes
	//  Enumerations
	Nodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap                                                                                            // function
	CheckedNodes(state types.TCheckState, considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap                                                            // function
	ChildNodes(node types.PVirtualNode) IVTVirtualNodeEnumerationWrap                                                                                          // function
	CutCopyNodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap                                                                                     // function
	InitializedNodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap                                                                                 // function
	LeafNodes() IVTVirtualNodeEnumerationWrap                                                                                                                  // function
	LevelNodes(nodeLevel uint32) IVTVirtualNodeEnumerationWrap                                                                                                 // function
	NoInitNodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap                                                                                      // function
	SelectedNodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap                                                                                    // function
	VisibleNodes(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) IVTVirtualNodeEnumerationWrap                                      // function
	VisibleChildNodes(node types.PVirtualNode, includeFiltered bool) IVTVirtualNodeEnumerationWrap                                                             // function
	VisibleChildNoInitNodes(node types.PVirtualNode, includeFiltered bool) IVTVirtualNodeEnumerationWrap                                                       // function
	VisibleNoInitNodes(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) IVTVirtualNodeEnumerationWrap                                // function
	AddFromStream(stream IStream, targetNode types.PVirtualNode)                                                                                               // procedure
	BeginDragWithBoolInt(immediate bool, threshold int32)                                                                                                      // procedure
	BeginSynch()                                                                                                                                               // procedure
	BeginUpdate()                                                                                                                                              // procedure
	CancelCutOrCopy()                                                                                                                                          // procedure
	CancelOperation()                                                                                                                                          // procedure
	Clear()                                                                                                                                                    // procedure
	ClearChecked()                                                                                                                                             // procedure
	ClearSelection()                                                                                                                                           // procedure
	CopyToClipboard()                                                                                                                                          // procedure
	CutToClipboard()                                                                                                                                           // procedure
	DeleteChildren(node types.PVirtualNode, resetHasChildren bool)                                                                                             // procedure
	DeleteNode(node types.PVirtualNode, reindex bool)                                                                                                          // procedure
	DeleteSelectedNodes()                                                                                                                                      // procedure
	EndSynch()                                                                                                                                                 // procedure
	EndUpdate()                                                                                                                                                // procedure
	EnsureNodeSelected()                                                                                                                                       // procedure
	FinishCutOrCopy()                                                                                                                                          // procedure
	FlushClipboard()                                                                                                                                           // procedure
	FullCollapse(node types.PVirtualNode)                                                                                                                      // procedure
	FullExpand(node types.PVirtualNode)                                                                                                                        // procedure
	GetHitTestInfoAt(X int32, Y int32, relative bool, hitInfo *THitInfo)                                                                                       // procedure
	GetTextInfo(node types.PVirtualNode, column int32, font IFont, R *types.TRect, outText *string)                                                            // procedure
	InvalidateChildren(node types.PVirtualNode, recursive bool)                                                                                                // procedure
	InvalidateColumn(column int32)                                                                                                                             // procedure
	InvalidateToBottom(node types.PVirtualNode)                                                                                                                // procedure
	InvertSelection(visibleOnly bool)                                                                                                                          // procedure
	LoadFromFile(fileName string)                                                                                                                              // procedure
	LoadFromStream(stream IStream)                                                                                                                             // procedure
	MeasureItemHeight(canvas ICanvas, node types.PVirtualNode)                                                                                                 // procedure
	MoveToWithPVirtualNodeX2VTNodeAttachModeBool(source types.PVirtualNode, target types.PVirtualNode, mode types.TVTNodeAttachMode, childrenOnly bool)        // procedure
	MoveToWithPVirtualNodeBaseVirtualTreeVTNodeAttachModeBool(node types.PVirtualNode, tree IBaseVirtualTree, mode types.TVTNodeAttachMode, childrenOnly bool) // procedure
	PaintTree(targetCanvas ICanvas, window types.TRect, target types.TPoint, paintOptions types.TVTInternalPaintOptions, pixelFormat types.TPixelFormat)       // procedure
	RepaintNode(node types.PVirtualNode)                                                                                                                       // procedure
	ReinitChildren(node types.PVirtualNode, recursive bool)                                                                                                    // procedure
	ReinitNode(node types.PVirtualNode, recursive bool)                                                                                                        // procedure
	ResetNode(node types.PVirtualNode)                                                                                                                         // procedure
	SaveToFile(fileName string)                                                                                                                                // procedure
	SaveToStream(stream IStream, node types.PVirtualNode)                                                                                                      // procedure
	SelectAll(visibleOnly bool)                                                                                                                                // procedure
	Sort(node types.PVirtualNode, column int32, direction types.TSortDirection, doInit bool)                                                                   // procedure
	SortTree(column int32, direction types.TSortDirection, doInit bool)                                                                                        // procedure
	ToggleNode(node types.PVirtualNode)                                                                                                                        // procedure
	UpdateHorizontalRange()                                                                                                                                    // procedure
	UpdateHorizontalScrollBar(doRepaint bool)                                                                                                                  // procedure
	UpdateRanges()                                                                                                                                             // procedure
	UpdateScrollBars(doRepaint bool)                                                                                                                           // procedure
	UpdateVerticalRange()                                                                                                                                      // procedure
	UpdateVerticalScrollBar(doRepaint bool)                                                                                                                    // procedure
	// ValidateChildren
	//  lcl: reenable in case TControl implementation change to match Delphi
	//  function UseRightToLeftReading: Boolean;
	ValidateChildren(node types.PVirtualNode, recursive bool)        // procedure
	ValidateNode(node types.PVirtualNode, recursive bool)            // procedure
	BottomNode() types.PVirtualNode                                  // property BottomNode Getter
	SetBottomNode(value types.PVirtualNode)                          // property BottomNode Setter
	CheckedCount() int32                                             // property CheckedCount Getter
	CheckImages() ICustomImageList                                   // property CheckImages Getter
	CheckState(node types.PVirtualNode) types.TCheckState            // property CheckState Getter
	SetCheckState(node types.PVirtualNode, value types.TCheckState)  // property CheckState Setter
	CheckType(node types.PVirtualNode) types.TCheckType              // property CheckType Getter
	SetCheckType(node types.PVirtualNode, value types.TCheckType)    // property CheckType Setter
	ChildCount(node types.PVirtualNode) uint32                       // property ChildCount Getter
	SetChildCount(node types.PVirtualNode, value uint32)             // property ChildCount Setter
	ChildrenInitialized(node types.PVirtualNode) bool                // property ChildrenInitialized Getter
	CutCopyCount() int32                                             // property CutCopyCount Getter
	DragImage() IVTDragImage                                         // property DragImage Getter
	VTVDragManager() IVTDragManager                                  // property VTVDragManager Getter
	DropTargetNode() types.PVirtualNode                              // property DropTargetNode Getter
	SetDropTargetNode(value types.PVirtualNode)                      // property DropTargetNode Setter
	EditLink() IVTEditLink                                           // property EditLink Getter
	EmptyListMessage() string                                        // property EmptyListMessage Getter
	SetEmptyListMessage(value string)                                // property EmptyListMessage Setter
	Expanded(node types.PVirtualNode) bool                           // property Expanded Getter
	SetExpanded(node types.PVirtualNode, value bool)                 // property Expanded Setter
	FocusedColumn() int32                                            // property FocusedColumn Getter
	SetFocusedColumn(value int32)                                    // property FocusedColumn Setter
	FocusedNode() types.PVirtualNode                                 // property FocusedNode Getter
	SetFocusedNode(value types.PVirtualNode)                         // property FocusedNode Setter
	FullyVisible(node types.PVirtualNode) bool                       // property FullyVisible Getter
	SetFullyVisible(node types.PVirtualNode, value bool)             // property FullyVisible Setter
	HasChildren(node types.PVirtualNode) bool                        // property HasChildren Getter
	SetHasChildren(node types.PVirtualNode, value bool)              // property HasChildren Setter
	HotNode() types.PVirtualNode                                     // property HotNode Getter
	IsDisabled(node types.PVirtualNode) bool                         // property IsDisabled Getter
	SetIsDisabled(node types.PVirtualNode, value bool)               // property IsDisabled Setter
	IsEffectivelyFiltered(node types.PVirtualNode) bool              // property IsEffectivelyFiltered Getter
	IsEffectivelyVisible(node types.PVirtualNode) bool               // property IsEffectivelyVisible Getter
	IsFiltered(node types.PVirtualNode) bool                         // property IsFiltered Getter
	SetIsFiltered(node types.PVirtualNode, value bool)               // property IsFiltered Setter
	IsVisibleWithPVirtualNodeToBool(node types.PVirtualNode) bool    // property IsVisible Getter
	SetIsVisible(node types.PVirtualNode, value bool)                // property IsVisible Setter
	MultiLine(node types.PVirtualNode) bool                          // property MultiLine Getter
	SetMultiLine(node types.PVirtualNode, value bool)                // property MultiLine Setter
	NodeHeight(node types.PVirtualNode) uint32                       // property NodeHeight Getter
	SetNodeHeight(node types.PVirtualNode, value uint32)             // property NodeHeight Setter
	NodeParent(node types.PVirtualNode) types.PVirtualNode           // property NodeParent Getter
	SetNodeParent(node types.PVirtualNode, value types.PVirtualNode) // property NodeParent Setter
	OffsetX() int32                                                  // property OffsetX Getter
	SetOffsetX(value int32)                                          // property OffsetX Setter
	OffsetXY() types.TPoint                                          // property OffsetXY Getter
	SetOffsetXY(value types.TPoint)                                  // property OffsetXY Setter
	OffsetY() int32                                                  // property OffsetY Getter
	SetOffsetY(value int32)                                          // property OffsetY Setter
	OperationCount() uint32                                          // property OperationCount Getter
	RootNode() types.PVirtualNode                                    // property RootNode Getter
	SearchBuffer() string                                            // property SearchBuffer Getter
	Selected(node types.PVirtualNode) bool                           // property Selected Getter
	SetSelected(node types.PVirtualNode, value bool)                 // property Selected Setter
	SelectionLocked() bool                                           // property SelectionLocked Getter
	SetSelectionLocked(value bool)                                   // property SelectionLocked Setter
	TotalCount() uint32                                              // property TotalCount Getter
	TreeStates() types.TVirtualTreeStates                            // property TreeStates Getter
	SetTreeStates(value types.TVirtualTreeStates)                    // property TreeStates Setter
	SelectedCount() int32                                            // property SelectedCount Getter
	TopNode() types.PVirtualNode                                     // property TopNode Getter
	SetTopNode(value types.PVirtualNode)                             // property TopNode Setter
	VerticalAlignment(node types.PVirtualNode) byte                  // property VerticalAlignment Getter
	SetVerticalAlignment(node types.PVirtualNode, value byte)        // property VerticalAlignment Setter
	VisibleCount() uint32                                            // property VisibleCount Getter
	VisiblePath(node types.PVirtualNode) bool                        // property VisiblePath Getter
	SetVisiblePath(node types.PVirtualNode, value bool)              // property VisiblePath Setter
	UpdateCount() uint32                                             // property UpdateCount Getter
}

type TBaseVirtualTree struct {
	TCustomControl
}

func (m *TBaseVirtualTree) AbsoluteIndex(node types.PVirtualNode) uint32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(1, m.Instance(), uintptr(node))
	return uint32(r)
}

func (m *TBaseVirtualTree) AddChild(parent types.PVirtualNode, userData uintptr) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(2, m.Instance(), uintptr(parent), uintptr(userData))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) CancelEditNode() bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) CanEdit(node types.PVirtualNode, column int32) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(4, m.Instance(), uintptr(node), uintptr(column))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) CopyToWithPVirtualNodeBaseVirtualTreeVTNodeAttachModeBool(source types.PVirtualNode, tree IBaseVirtualTree, mode types.TVTNodeAttachMode, childrenOnly bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(5, m.Instance(), uintptr(source), base.GetObjectUintptr(tree), uintptr(mode), api.PasBool(childrenOnly))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) CopyToWithPVirtualNodeX2VTNodeAttachModeBool(source types.PVirtualNode, target types.PVirtualNode, mode types.TVTNodeAttachMode, childrenOnly bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(6, m.Instance(), uintptr(source), uintptr(target), uintptr(mode), api.PasBool(childrenOnly))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) DraggingToBool() bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) EditNode(node types.PVirtualNode, column int32) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(8, m.Instance(), uintptr(node), uintptr(column))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) EndEditNode() bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) GetDisplayRect(node types.PVirtualNode, column int32, textOnly bool, unclipped bool, applyCellContentMargin bool) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(10, m.Instance(), uintptr(node), uintptr(column), api.PasBool(textOnly), api.PasBool(unclipped), api.PasBool(applyCellContentMargin), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TBaseVirtualTree) GetEffectivelyFiltered(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(11, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) GetEffectivelyVisible(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(12, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) GetFirst(considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(13, m.Instance(), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstChecked(state types.TCheckState, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(14, m.Instance(), uintptr(state), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstChild(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(15, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstChildNoInit(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(16, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstCutCopy(considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(17, m.Instance(), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstInitialized(considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(18, m.Instance(), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstLeaf() types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(19, m.Instance())
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstLevel(nodeLevel uint32) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(20, m.Instance(), uintptr(nodeLevel))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstNoInit(considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(21, m.Instance(), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstSelected(considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(22, m.Instance(), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstVisible(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(23, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstVisibleChild(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(24, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstVisibleChildNoInit(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(25, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetFirstVisibleNoInit(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(26, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetLast(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(27, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetLastInitialized(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(28, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetLastNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(29, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetLastChild(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(30, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetLastChildNoInit(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(31, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetLastVisible(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(32, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetLastVisibleChild(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(33, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetLastVisibleChildNoInit(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(34, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetLastVisibleNoInit(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(35, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetMaxColumnWidth(column int32, useSmartColumnWidth bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(36, m.Instance(), uintptr(column), api.PasBool(useSmartColumnWidth))
	return int32(r)
}

func (m *TBaseVirtualTree) GetNext(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(37, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextCheckedWithPVirtualNodeCheckStateBool(node types.PVirtualNode, state types.TCheckState, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(38, m.Instance(), uintptr(node), uintptr(state), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextCheckedWithPVirtualNodeBool(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(39, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextCutCopy(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(40, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextInitialized(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(41, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextLeaf(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(42, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextLevel(node types.PVirtualNode, nodeLevel uint32) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(43, m.Instance(), uintptr(node), uintptr(nodeLevel))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(44, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextSelected(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(45, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextSibling(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(46, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextSiblingNoInit(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(47, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextVisible(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(48, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextVisibleNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(49, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextVisibleSibling(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(50, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNextVisibleSiblingNoInit(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(51, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNodeAtWithPoint(P types.TPoint) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(52, m.Instance(), uintptr(base.UnsafePointer(&P)))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNodeAtWithIntX2(X int32, Y int32) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(53, m.Instance(), uintptr(X), uintptr(Y))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNodeAtWithIntX3Bool(X int32, Y int32, relative bool, nodeTop *int32) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	nodeTopPtr := uintptr(*nodeTop)
	r := baseVirtualTreeAPI().SysCallN(54, m.Instance(), uintptr(X), uintptr(Y), api.PasBool(relative), uintptr(base.UnsafePointer(&nodeTopPtr)))
	*nodeTop = int32(nodeTopPtr)
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetNodeData(node types.PVirtualNode) uintptr {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(55, m.Instance(), uintptr(node))
	return uintptr(r)
}

func (m *TBaseVirtualTree) GetNodeLevel(node types.PVirtualNode) uint32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(56, m.Instance(), uintptr(node))
	return uint32(r)
}

func (m *TBaseVirtualTree) GetPrevious(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(57, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousChecked(node types.PVirtualNode, state types.TCheckState, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(58, m.Instance(), uintptr(node), uintptr(state), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousCutCopy(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(59, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousInitialized(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(60, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousLeaf(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(61, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousLevel(node types.PVirtualNode, nodeLevel uint32) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(62, m.Instance(), uintptr(node), uintptr(nodeLevel))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(63, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousSelected(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(64, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousSibling(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(65, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousSiblingNoInit(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(66, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousVisible(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(67, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousVisibleNoInit(node types.PVirtualNode, considerChildrenAbove bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(68, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousVisibleSibling(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(69, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetPreviousVisibleSiblingNoInit(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(70, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) GetSortedCutCopySet(resolve bool) types.TNodeArray {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(71, m.Instance(), api.PasBool(resolve))
	return types.TNodeArray(r)
}

func (m *TBaseVirtualTree) GetSortedSelection(resolve bool) types.TNodeArray {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(72, m.Instance(), api.PasBool(resolve))
	return types.TNodeArray(r)
}

func (m *TBaseVirtualTree) GetTreeRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(73, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TBaseVirtualTree) GetVisibleParent(node types.PVirtualNode, includeFiltered bool) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(74, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) HasAsParent(node types.PVirtualNode, potentialParent types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(75, m.Instance(), uintptr(node), uintptr(potentialParent))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) InsertNode(node types.PVirtualNode, mode types.TVTNodeAttachMode, userData uintptr) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(76, m.Instance(), uintptr(node), uintptr(mode), uintptr(userData))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) InvalidateNode(node types.PVirtualNode) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(77, m.Instance(), uintptr(node), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TBaseVirtualTree) IsEditing() bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(78, m.Instance())
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) IsMouseSelecting() bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(79, m.Instance())
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) IsEmpty() bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(80, m.Instance())
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) PasteFromClipboard() bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(81, m.Instance())
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) ScrollIntoViewWithPVirtualNodeBoolX2(node types.PVirtualNode, center bool, horizontally bool) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(82, m.Instance(), uintptr(node), api.PasBool(center), api.PasBool(horizontally))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) ScrollIntoViewWithColumnIndexBool(column int32, center bool) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(83, m.Instance(), uintptr(column), api.PasBool(center))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) Nodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(84, m.Instance(), api.PasBool(considerChildrenAbove))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) CheckedNodes(state types.TCheckState, considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(85, m.Instance(), uintptr(state), api.PasBool(considerChildrenAbove))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) ChildNodes(node types.PVirtualNode) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(86, m.Instance(), uintptr(node))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) CutCopyNodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(87, m.Instance(), api.PasBool(considerChildrenAbove))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) InitializedNodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(88, m.Instance(), api.PasBool(considerChildrenAbove))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) LeafNodes() IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(89, m.Instance())
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) LevelNodes(nodeLevel uint32) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(90, m.Instance(), uintptr(nodeLevel))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) NoInitNodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(91, m.Instance(), api.PasBool(considerChildrenAbove))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) SelectedNodes(considerChildrenAbove bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(92, m.Instance(), api.PasBool(considerChildrenAbove))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) VisibleNodes(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(93, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove), api.PasBool(includeFiltered))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) VisibleChildNodes(node types.PVirtualNode, includeFiltered bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(94, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) VisibleChildNoInitNodes(node types.PVirtualNode, includeFiltered bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(95, m.Instance(), uintptr(node), api.PasBool(includeFiltered))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) VisibleNoInitNodes(node types.PVirtualNode, considerChildrenAbove bool, includeFiltered bool) IVTVirtualNodeEnumerationWrap {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(96, m.Instance(), uintptr(node), api.PasBool(considerChildrenAbove), api.PasBool(includeFiltered))
	return AsVTVirtualNodeEnumerationWrap(r)
}

func (m *TBaseVirtualTree) AddFromStream(stream IStream, targetNode types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(97, m.Instance(), base.GetObjectUintptr(stream), uintptr(targetNode))
}

func (m *TBaseVirtualTree) BeginDragWithBoolInt(immediate bool, threshold int32) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(98, m.Instance(), api.PasBool(immediate), uintptr(threshold))
}

func (m *TBaseVirtualTree) BeginSynch() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(99, m.Instance())
}

func (m *TBaseVirtualTree) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(100, m.Instance())
}

func (m *TBaseVirtualTree) CancelCutOrCopy() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(101, m.Instance())
}

func (m *TBaseVirtualTree) CancelOperation() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(102, m.Instance())
}

func (m *TBaseVirtualTree) Clear() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(103, m.Instance())
}

func (m *TBaseVirtualTree) ClearChecked() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(104, m.Instance())
}

func (m *TBaseVirtualTree) ClearSelection() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(105, m.Instance())
}

func (m *TBaseVirtualTree) CopyToClipboard() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(106, m.Instance())
}

func (m *TBaseVirtualTree) CutToClipboard() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(107, m.Instance())
}

func (m *TBaseVirtualTree) DeleteChildren(node types.PVirtualNode, resetHasChildren bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(108, m.Instance(), uintptr(node), api.PasBool(resetHasChildren))
}

func (m *TBaseVirtualTree) DeleteNode(node types.PVirtualNode, reindex bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(109, m.Instance(), uintptr(node), api.PasBool(reindex))
}

func (m *TBaseVirtualTree) DeleteSelectedNodes() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(110, m.Instance())
}

func (m *TBaseVirtualTree) EndSynch() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(111, m.Instance())
}

func (m *TBaseVirtualTree) EndUpdate() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(112, m.Instance())
}

func (m *TBaseVirtualTree) EnsureNodeSelected() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(113, m.Instance())
}

func (m *TBaseVirtualTree) FinishCutOrCopy() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(114, m.Instance())
}

func (m *TBaseVirtualTree) FlushClipboard() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(115, m.Instance())
}

func (m *TBaseVirtualTree) FullCollapse(node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(116, m.Instance(), uintptr(node))
}

func (m *TBaseVirtualTree) FullExpand(node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(117, m.Instance(), uintptr(node))
}

func (m *TBaseVirtualTree) GetHitTestInfoAt(X int32, Y int32, relative bool, hitInfo *THitInfo) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(118, m.Instance(), uintptr(X), uintptr(Y), api.PasBool(relative), uintptr(base.UnsafePointer(hitInfo)))
}

func (m *TBaseVirtualTree) GetTextInfo(node types.PVirtualNode, column int32, font IFont, R *types.TRect, outText *string) {
	if !m.IsValid() {
		return
	}
	var textPtr uintptr
	baseVirtualTreeAPI().SysCallN(119, m.Instance(), uintptr(node), uintptr(column), base.GetObjectUintptr(font), uintptr(base.UnsafePointer(R)), uintptr(base.UnsafePointer(&textPtr)))
	*outText = api.GoStr(textPtr)
}

func (m *TBaseVirtualTree) InvalidateChildren(node types.PVirtualNode, recursive bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(120, m.Instance(), uintptr(node), api.PasBool(recursive))
}

func (m *TBaseVirtualTree) InvalidateColumn(column int32) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(121, m.Instance(), uintptr(column))
}

func (m *TBaseVirtualTree) InvalidateToBottom(node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(122, m.Instance(), uintptr(node))
}

func (m *TBaseVirtualTree) InvertSelection(visibleOnly bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(123, m.Instance(), api.PasBool(visibleOnly))
}

func (m *TBaseVirtualTree) LoadFromFile(fileName string) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(124, m.Instance(), api.PasStr(fileName))
}

func (m *TBaseVirtualTree) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(125, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TBaseVirtualTree) MeasureItemHeight(canvas ICanvas, node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(126, m.Instance(), base.GetObjectUintptr(canvas), uintptr(node))
}

func (m *TBaseVirtualTree) MoveToWithPVirtualNodeX2VTNodeAttachModeBool(source types.PVirtualNode, target types.PVirtualNode, mode types.TVTNodeAttachMode, childrenOnly bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(127, m.Instance(), uintptr(source), uintptr(target), uintptr(mode), api.PasBool(childrenOnly))
}

func (m *TBaseVirtualTree) MoveToWithPVirtualNodeBaseVirtualTreeVTNodeAttachModeBool(node types.PVirtualNode, tree IBaseVirtualTree, mode types.TVTNodeAttachMode, childrenOnly bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(128, m.Instance(), uintptr(node), base.GetObjectUintptr(tree), uintptr(mode), api.PasBool(childrenOnly))
}

func (m *TBaseVirtualTree) PaintTree(targetCanvas ICanvas, window types.TRect, target types.TPoint, paintOptions types.TVTInternalPaintOptions, pixelFormat types.TPixelFormat) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(129, m.Instance(), base.GetObjectUintptr(targetCanvas), uintptr(base.UnsafePointer(&window)), uintptr(base.UnsafePointer(&target)), uintptr(paintOptions), uintptr(pixelFormat))
}

func (m *TBaseVirtualTree) RepaintNode(node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(130, m.Instance(), uintptr(node))
}

func (m *TBaseVirtualTree) ReinitChildren(node types.PVirtualNode, recursive bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(131, m.Instance(), uintptr(node), api.PasBool(recursive))
}

func (m *TBaseVirtualTree) ReinitNode(node types.PVirtualNode, recursive bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(132, m.Instance(), uintptr(node), api.PasBool(recursive))
}

func (m *TBaseVirtualTree) ResetNode(node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(133, m.Instance(), uintptr(node))
}

func (m *TBaseVirtualTree) SaveToFile(fileName string) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(134, m.Instance(), api.PasStr(fileName))
}

func (m *TBaseVirtualTree) SaveToStream(stream IStream, node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(135, m.Instance(), base.GetObjectUintptr(stream), uintptr(node))
}

func (m *TBaseVirtualTree) SelectAll(visibleOnly bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(136, m.Instance(), api.PasBool(visibleOnly))
}

func (m *TBaseVirtualTree) Sort(node types.PVirtualNode, column int32, direction types.TSortDirection, doInit bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(137, m.Instance(), uintptr(node), uintptr(column), uintptr(direction), api.PasBool(doInit))
}

func (m *TBaseVirtualTree) SortTree(column int32, direction types.TSortDirection, doInit bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(138, m.Instance(), uintptr(column), uintptr(direction), api.PasBool(doInit))
}

func (m *TBaseVirtualTree) ToggleNode(node types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(139, m.Instance(), uintptr(node))
}

func (m *TBaseVirtualTree) UpdateHorizontalRange() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(140, m.Instance())
}

func (m *TBaseVirtualTree) UpdateHorizontalScrollBar(doRepaint bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(141, m.Instance(), api.PasBool(doRepaint))
}

func (m *TBaseVirtualTree) UpdateRanges() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(142, m.Instance())
}

func (m *TBaseVirtualTree) UpdateScrollBars(doRepaint bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(143, m.Instance(), api.PasBool(doRepaint))
}

func (m *TBaseVirtualTree) UpdateVerticalRange() {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(144, m.Instance())
}

func (m *TBaseVirtualTree) UpdateVerticalScrollBar(doRepaint bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(145, m.Instance(), api.PasBool(doRepaint))
}

func (m *TBaseVirtualTree) ValidateChildren(node types.PVirtualNode, recursive bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(146, m.Instance(), uintptr(node), api.PasBool(recursive))
}

func (m *TBaseVirtualTree) ValidateNode(node types.PVirtualNode, recursive bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(147, m.Instance(), uintptr(node), api.PasBool(recursive))
}

func (m *TBaseVirtualTree) BottomNode() types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(148, 0, m.Instance())
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) SetBottomNode(value types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(148, 1, m.Instance(), uintptr(value))
}

func (m *TBaseVirtualTree) CheckedCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(149, m.Instance())
	return int32(r)
}

func (m *TBaseVirtualTree) CheckImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(150, m.Instance())
	return AsCustomImageList(r)
}

func (m *TBaseVirtualTree) CheckState(node types.PVirtualNode) types.TCheckState {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(151, 0, m.Instance(), uintptr(node))
	return types.TCheckState(r)
}

func (m *TBaseVirtualTree) SetCheckState(node types.PVirtualNode, value types.TCheckState) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(151, 1, m.Instance(), uintptr(node), uintptr(value))
}

func (m *TBaseVirtualTree) CheckType(node types.PVirtualNode) types.TCheckType {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(152, 0, m.Instance(), uintptr(node))
	return types.TCheckType(r)
}

func (m *TBaseVirtualTree) SetCheckType(node types.PVirtualNode, value types.TCheckType) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(152, 1, m.Instance(), uintptr(node), uintptr(value))
}

func (m *TBaseVirtualTree) ChildCount(node types.PVirtualNode) uint32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(153, 0, m.Instance(), uintptr(node))
	return uint32(r)
}

func (m *TBaseVirtualTree) SetChildCount(node types.PVirtualNode, value uint32) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(153, 1, m.Instance(), uintptr(node), uintptr(value))
}

func (m *TBaseVirtualTree) ChildrenInitialized(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(154, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) CutCopyCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(155, m.Instance())
	return int32(r)
}

func (m *TBaseVirtualTree) DragImage() IVTDragImage {
	if !m.IsValid() {
		return nil
	}
	r := baseVirtualTreeAPI().SysCallN(156, m.Instance())
	return AsVTDragImage(r)
}

func (m *TBaseVirtualTree) VTVDragManager() (result IVTDragManager) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	baseVirtualTreeAPI().SysCallN(157, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsVTDragManager(resultPtr)
	return
}

func (m *TBaseVirtualTree) DropTargetNode() types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(158, 0, m.Instance())
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) SetDropTargetNode(value types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(158, 1, m.Instance(), uintptr(value))
}

func (m *TBaseVirtualTree) EditLink() (result IVTEditLink) {
	if !m.IsValid() {
		return
	}
	var resultPtr uintptr
	baseVirtualTreeAPI().SysCallN(159, m.Instance(), uintptr(base.UnsafePointer(&resultPtr)))
	result = AsVTEditLink(resultPtr)
	return
}

func (m *TBaseVirtualTree) EmptyListMessage() string {
	if !m.IsValid() {
		return ""
	}
	r := baseVirtualTreeAPI().SysCallN(160, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TBaseVirtualTree) SetEmptyListMessage(value string) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(160, 1, m.Instance(), api.PasStr(value))
}

func (m *TBaseVirtualTree) Expanded(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(161, 0, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetExpanded(node types.PVirtualNode, value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(161, 1, m.Instance(), uintptr(node), api.PasBool(value))
}

func (m *TBaseVirtualTree) FocusedColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(162, 0, m.Instance())
	return int32(r)
}

func (m *TBaseVirtualTree) SetFocusedColumn(value int32) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(162, 1, m.Instance(), uintptr(value))
}

func (m *TBaseVirtualTree) FocusedNode() types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(163, 0, m.Instance())
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) SetFocusedNode(value types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(163, 1, m.Instance(), uintptr(value))
}

func (m *TBaseVirtualTree) FullyVisible(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(164, 0, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetFullyVisible(node types.PVirtualNode, value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(164, 1, m.Instance(), uintptr(node), api.PasBool(value))
}

func (m *TBaseVirtualTree) HasChildren(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(165, 0, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetHasChildren(node types.PVirtualNode, value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(165, 1, m.Instance(), uintptr(node), api.PasBool(value))
}

func (m *TBaseVirtualTree) HotNode() types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(166, m.Instance())
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) IsDisabled(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(167, 0, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetIsDisabled(node types.PVirtualNode, value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(167, 1, m.Instance(), uintptr(node), api.PasBool(value))
}

func (m *TBaseVirtualTree) IsEffectivelyFiltered(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(168, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) IsEffectivelyVisible(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(169, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) IsFiltered(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(170, 0, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetIsFiltered(node types.PVirtualNode, value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(170, 1, m.Instance(), uintptr(node), api.PasBool(value))
}

func (m *TBaseVirtualTree) IsVisibleWithPVirtualNodeToBool(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(171, 0, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetIsVisible(node types.PVirtualNode, value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(171, 1, m.Instance(), uintptr(node), api.PasBool(value))
}

func (m *TBaseVirtualTree) MultiLine(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(172, 0, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetMultiLine(node types.PVirtualNode, value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(172, 1, m.Instance(), uintptr(node), api.PasBool(value))
}

func (m *TBaseVirtualTree) NodeHeight(node types.PVirtualNode) uint32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(173, 0, m.Instance(), uintptr(node))
	return uint32(r)
}

func (m *TBaseVirtualTree) SetNodeHeight(node types.PVirtualNode, value uint32) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(173, 1, m.Instance(), uintptr(node), uintptr(value))
}

func (m *TBaseVirtualTree) NodeParent(node types.PVirtualNode) types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(174, 0, m.Instance(), uintptr(node))
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) SetNodeParent(node types.PVirtualNode, value types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(174, 1, m.Instance(), uintptr(node), uintptr(value))
}

func (m *TBaseVirtualTree) OffsetX() int32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(175, 0, m.Instance())
	return int32(r)
}

func (m *TBaseVirtualTree) SetOffsetX(value int32) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(175, 1, m.Instance(), uintptr(value))
}

func (m *TBaseVirtualTree) OffsetXY() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(176, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TBaseVirtualTree) SetOffsetXY(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(176, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TBaseVirtualTree) OffsetY() int32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(177, 0, m.Instance())
	return int32(r)
}

func (m *TBaseVirtualTree) SetOffsetY(value int32) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(177, 1, m.Instance(), uintptr(value))
}

func (m *TBaseVirtualTree) OperationCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(178, m.Instance())
	return uint32(r)
}

func (m *TBaseVirtualTree) RootNode() types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(179, m.Instance())
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) SearchBuffer() string {
	if !m.IsValid() {
		return ""
	}
	r := baseVirtualTreeAPI().SysCallN(180, m.Instance())
	return api.GoStr(r)
}

func (m *TBaseVirtualTree) Selected(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(181, 0, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetSelected(node types.PVirtualNode, value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(181, 1, m.Instance(), uintptr(node), api.PasBool(value))
}

func (m *TBaseVirtualTree) SelectionLocked() bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(182, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetSelectionLocked(value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(182, 1, m.Instance(), api.PasBool(value))
}

func (m *TBaseVirtualTree) TotalCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(183, m.Instance())
	return uint32(r)
}

func (m *TBaseVirtualTree) TreeStates() types.TVirtualTreeStates {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(184, 0, m.Instance())
	return types.TVirtualTreeStates(r)
}

func (m *TBaseVirtualTree) SetTreeStates(value types.TVirtualTreeStates) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(184, 1, m.Instance(), uintptr(value))
}

func (m *TBaseVirtualTree) SelectedCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(185, m.Instance())
	return int32(r)
}

func (m *TBaseVirtualTree) TopNode() types.PVirtualNode {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(186, 0, m.Instance())
	return types.PVirtualNode(r)
}

func (m *TBaseVirtualTree) SetTopNode(value types.PVirtualNode) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(186, 1, m.Instance(), uintptr(value))
}

func (m *TBaseVirtualTree) VerticalAlignment(node types.PVirtualNode) byte {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(187, 0, m.Instance(), uintptr(node))
	return byte(r)
}

func (m *TBaseVirtualTree) SetVerticalAlignment(node types.PVirtualNode, value byte) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(187, 1, m.Instance(), uintptr(node), uintptr(value))
}

func (m *TBaseVirtualTree) VisibleCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(188, m.Instance())
	return uint32(r)
}

func (m *TBaseVirtualTree) VisiblePath(node types.PVirtualNode) bool {
	if !m.IsValid() {
		return false
	}
	r := baseVirtualTreeAPI().SysCallN(189, 0, m.Instance(), uintptr(node))
	return api.GoBool(r)
}

func (m *TBaseVirtualTree) SetVisiblePath(node types.PVirtualNode, value bool) {
	if !m.IsValid() {
		return
	}
	baseVirtualTreeAPI().SysCallN(189, 1, m.Instance(), uintptr(node), api.PasBool(value))
}

func (m *TBaseVirtualTree) UpdateCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := baseVirtualTreeAPI().SysCallN(190, m.Instance())
	return uint32(r)
}

// NewBaseVirtualTree class constructor
func NewBaseVirtualTree(owner IComponent) IBaseVirtualTree {
	r := baseVirtualTreeAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsBaseVirtualTree(r)
}

func TBaseVirtualTreeClass() types.TClass {
	r := baseVirtualTreeAPI().SysCallN(191)
	return types.TClass(r)
}

var (
	baseVirtualTreeOnce   base.Once
	baseVirtualTreeImport *imports.Imports = nil
)

func baseVirtualTreeAPI() *imports.Imports {
	baseVirtualTreeOnce.Do(func() {
		baseVirtualTreeImport = api.NewDefaultImports()
		baseVirtualTreeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBaseVirtualTree_Create", 0), // constructor NewBaseVirtualTree
			/* 1 */ imports.NewTable("TBaseVirtualTree_AbsoluteIndex", 0), // function AbsoluteIndex
			/* 2 */ imports.NewTable("TBaseVirtualTree_AddChild", 0), // function AddChild
			/* 3 */ imports.NewTable("TBaseVirtualTree_CancelEditNode", 0), // function CancelEditNode
			/* 4 */ imports.NewTable("TBaseVirtualTree_CanEdit", 0), // function CanEdit
			/* 5 */ imports.NewTable("TBaseVirtualTree_CopyToWithPVirtualNodeBaseVirtualTreeVTNodeAttachModeBool", 0), // function CopyToWithPVirtualNodeBaseVirtualTreeVTNodeAttachModeBool
			/* 6 */ imports.NewTable("TBaseVirtualTree_CopyToWithPVirtualNodeX2VTNodeAttachModeBool", 0), // function CopyToWithPVirtualNodeX2VTNodeAttachModeBool
			/* 7 */ imports.NewTable("TBaseVirtualTree_DraggingToBool", 0), // function DraggingToBool
			/* 8 */ imports.NewTable("TBaseVirtualTree_EditNode", 0), // function EditNode
			/* 9 */ imports.NewTable("TBaseVirtualTree_EndEditNode", 0), // function EndEditNode
			/* 10 */ imports.NewTable("TBaseVirtualTree_GetDisplayRect", 0), // function GetDisplayRect
			/* 11 */ imports.NewTable("TBaseVirtualTree_GetEffectivelyFiltered", 0), // function GetEffectivelyFiltered
			/* 12 */ imports.NewTable("TBaseVirtualTree_GetEffectivelyVisible", 0), // function GetEffectivelyVisible
			/* 13 */ imports.NewTable("TBaseVirtualTree_GetFirst", 0), // function GetFirst
			/* 14 */ imports.NewTable("TBaseVirtualTree_GetFirstChecked", 0), // function GetFirstChecked
			/* 15 */ imports.NewTable("TBaseVirtualTree_GetFirstChild", 0), // function GetFirstChild
			/* 16 */ imports.NewTable("TBaseVirtualTree_GetFirstChildNoInit", 0), // function GetFirstChildNoInit
			/* 17 */ imports.NewTable("TBaseVirtualTree_GetFirstCutCopy", 0), // function GetFirstCutCopy
			/* 18 */ imports.NewTable("TBaseVirtualTree_GetFirstInitialized", 0), // function GetFirstInitialized
			/* 19 */ imports.NewTable("TBaseVirtualTree_GetFirstLeaf", 0), // function GetFirstLeaf
			/* 20 */ imports.NewTable("TBaseVirtualTree_GetFirstLevel", 0), // function GetFirstLevel
			/* 21 */ imports.NewTable("TBaseVirtualTree_GetFirstNoInit", 0), // function GetFirstNoInit
			/* 22 */ imports.NewTable("TBaseVirtualTree_GetFirstSelected", 0), // function GetFirstSelected
			/* 23 */ imports.NewTable("TBaseVirtualTree_GetFirstVisible", 0), // function GetFirstVisible
			/* 24 */ imports.NewTable("TBaseVirtualTree_GetFirstVisibleChild", 0), // function GetFirstVisibleChild
			/* 25 */ imports.NewTable("TBaseVirtualTree_GetFirstVisibleChildNoInit", 0), // function GetFirstVisibleChildNoInit
			/* 26 */ imports.NewTable("TBaseVirtualTree_GetFirstVisibleNoInit", 0), // function GetFirstVisibleNoInit
			/* 27 */ imports.NewTable("TBaseVirtualTree_GetLast", 0), // function GetLast
			/* 28 */ imports.NewTable("TBaseVirtualTree_GetLastInitialized", 0), // function GetLastInitialized
			/* 29 */ imports.NewTable("TBaseVirtualTree_GetLastNoInit", 0), // function GetLastNoInit
			/* 30 */ imports.NewTable("TBaseVirtualTree_GetLastChild", 0), // function GetLastChild
			/* 31 */ imports.NewTable("TBaseVirtualTree_GetLastChildNoInit", 0), // function GetLastChildNoInit
			/* 32 */ imports.NewTable("TBaseVirtualTree_GetLastVisible", 0), // function GetLastVisible
			/* 33 */ imports.NewTable("TBaseVirtualTree_GetLastVisibleChild", 0), // function GetLastVisibleChild
			/* 34 */ imports.NewTable("TBaseVirtualTree_GetLastVisibleChildNoInit", 0), // function GetLastVisibleChildNoInit
			/* 35 */ imports.NewTable("TBaseVirtualTree_GetLastVisibleNoInit", 0), // function GetLastVisibleNoInit
			/* 36 */ imports.NewTable("TBaseVirtualTree_GetMaxColumnWidth", 0), // function GetMaxColumnWidth
			/* 37 */ imports.NewTable("TBaseVirtualTree_GetNext", 0), // function GetNext
			/* 38 */ imports.NewTable("TBaseVirtualTree_GetNextCheckedWithPVirtualNodeCheckStateBool", 0), // function GetNextCheckedWithPVirtualNodeCheckStateBool
			/* 39 */ imports.NewTable("TBaseVirtualTree_GetNextCheckedWithPVirtualNodeBool", 0), // function GetNextCheckedWithPVirtualNodeBool
			/* 40 */ imports.NewTable("TBaseVirtualTree_GetNextCutCopy", 0), // function GetNextCutCopy
			/* 41 */ imports.NewTable("TBaseVirtualTree_GetNextInitialized", 0), // function GetNextInitialized
			/* 42 */ imports.NewTable("TBaseVirtualTree_GetNextLeaf", 0), // function GetNextLeaf
			/* 43 */ imports.NewTable("TBaseVirtualTree_GetNextLevel", 0), // function GetNextLevel
			/* 44 */ imports.NewTable("TBaseVirtualTree_GetNextNoInit", 0), // function GetNextNoInit
			/* 45 */ imports.NewTable("TBaseVirtualTree_GetNextSelected", 0), // function GetNextSelected
			/* 46 */ imports.NewTable("TBaseVirtualTree_GetNextSibling", 0), // function GetNextSibling
			/* 47 */ imports.NewTable("TBaseVirtualTree_GetNextSiblingNoInit", 0), // function GetNextSiblingNoInit
			/* 48 */ imports.NewTable("TBaseVirtualTree_GetNextVisible", 0), // function GetNextVisible
			/* 49 */ imports.NewTable("TBaseVirtualTree_GetNextVisibleNoInit", 0), // function GetNextVisibleNoInit
			/* 50 */ imports.NewTable("TBaseVirtualTree_GetNextVisibleSibling", 0), // function GetNextVisibleSibling
			/* 51 */ imports.NewTable("TBaseVirtualTree_GetNextVisibleSiblingNoInit", 0), // function GetNextVisibleSiblingNoInit
			/* 52 */ imports.NewTable("TBaseVirtualTree_GetNodeAtWithPoint", 0), // function GetNodeAtWithPoint
			/* 53 */ imports.NewTable("TBaseVirtualTree_GetNodeAtWithIntX2", 0), // function GetNodeAtWithIntX2
			/* 54 */ imports.NewTable("TBaseVirtualTree_GetNodeAtWithIntX3Bool", 0), // function GetNodeAtWithIntX3Bool
			/* 55 */ imports.NewTable("TBaseVirtualTree_GetNodeData", 0), // function GetNodeData
			/* 56 */ imports.NewTable("TBaseVirtualTree_GetNodeLevel", 0), // function GetNodeLevel
			/* 57 */ imports.NewTable("TBaseVirtualTree_GetPrevious", 0), // function GetPrevious
			/* 58 */ imports.NewTable("TBaseVirtualTree_GetPreviousChecked", 0), // function GetPreviousChecked
			/* 59 */ imports.NewTable("TBaseVirtualTree_GetPreviousCutCopy", 0), // function GetPreviousCutCopy
			/* 60 */ imports.NewTable("TBaseVirtualTree_GetPreviousInitialized", 0), // function GetPreviousInitialized
			/* 61 */ imports.NewTable("TBaseVirtualTree_GetPreviousLeaf", 0), // function GetPreviousLeaf
			/* 62 */ imports.NewTable("TBaseVirtualTree_GetPreviousLevel", 0), // function GetPreviousLevel
			/* 63 */ imports.NewTable("TBaseVirtualTree_GetPreviousNoInit", 0), // function GetPreviousNoInit
			/* 64 */ imports.NewTable("TBaseVirtualTree_GetPreviousSelected", 0), // function GetPreviousSelected
			/* 65 */ imports.NewTable("TBaseVirtualTree_GetPreviousSibling", 0), // function GetPreviousSibling
			/* 66 */ imports.NewTable("TBaseVirtualTree_GetPreviousSiblingNoInit", 0), // function GetPreviousSiblingNoInit
			/* 67 */ imports.NewTable("TBaseVirtualTree_GetPreviousVisible", 0), // function GetPreviousVisible
			/* 68 */ imports.NewTable("TBaseVirtualTree_GetPreviousVisibleNoInit", 0), // function GetPreviousVisibleNoInit
			/* 69 */ imports.NewTable("TBaseVirtualTree_GetPreviousVisibleSibling", 0), // function GetPreviousVisibleSibling
			/* 70 */ imports.NewTable("TBaseVirtualTree_GetPreviousVisibleSiblingNoInit", 0), // function GetPreviousVisibleSiblingNoInit
			/* 71 */ imports.NewTable("TBaseVirtualTree_GetSortedCutCopySet", 0), // function GetSortedCutCopySet
			/* 72 */ imports.NewTable("TBaseVirtualTree_GetSortedSelection", 0), // function GetSortedSelection
			/* 73 */ imports.NewTable("TBaseVirtualTree_GetTreeRect", 0), // function GetTreeRect
			/* 74 */ imports.NewTable("TBaseVirtualTree_GetVisibleParent", 0), // function GetVisibleParent
			/* 75 */ imports.NewTable("TBaseVirtualTree_HasAsParent", 0), // function HasAsParent
			/* 76 */ imports.NewTable("TBaseVirtualTree_InsertNode", 0), // function InsertNode
			/* 77 */ imports.NewTable("TBaseVirtualTree_InvalidateNode", 0), // function InvalidateNode
			/* 78 */ imports.NewTable("TBaseVirtualTree_IsEditing", 0), // function IsEditing
			/* 79 */ imports.NewTable("TBaseVirtualTree_IsMouseSelecting", 0), // function IsMouseSelecting
			/* 80 */ imports.NewTable("TBaseVirtualTree_IsEmpty", 0), // function IsEmpty
			/* 81 */ imports.NewTable("TBaseVirtualTree_PasteFromClipboard", 0), // function PasteFromClipboard
			/* 82 */ imports.NewTable("TBaseVirtualTree_ScrollIntoViewWithPVirtualNodeBoolX2", 0), // function ScrollIntoViewWithPVirtualNodeBoolX2
			/* 83 */ imports.NewTable("TBaseVirtualTree_ScrollIntoViewWithColumnIndexBool", 0), // function ScrollIntoViewWithColumnIndexBool
			/* 84 */ imports.NewTable("TBaseVirtualTree_Nodes", 0), // function Nodes
			/* 85 */ imports.NewTable("TBaseVirtualTree_CheckedNodes", 0), // function CheckedNodes
			/* 86 */ imports.NewTable("TBaseVirtualTree_ChildNodes", 0), // function ChildNodes
			/* 87 */ imports.NewTable("TBaseVirtualTree_CutCopyNodes", 0), // function CutCopyNodes
			/* 88 */ imports.NewTable("TBaseVirtualTree_InitializedNodes", 0), // function InitializedNodes
			/* 89 */ imports.NewTable("TBaseVirtualTree_LeafNodes", 0), // function LeafNodes
			/* 90 */ imports.NewTable("TBaseVirtualTree_LevelNodes", 0), // function LevelNodes
			/* 91 */ imports.NewTable("TBaseVirtualTree_NoInitNodes", 0), // function NoInitNodes
			/* 92 */ imports.NewTable("TBaseVirtualTree_SelectedNodes", 0), // function SelectedNodes
			/* 93 */ imports.NewTable("TBaseVirtualTree_VisibleNodes", 0), // function VisibleNodes
			/* 94 */ imports.NewTable("TBaseVirtualTree_VisibleChildNodes", 0), // function VisibleChildNodes
			/* 95 */ imports.NewTable("TBaseVirtualTree_VisibleChildNoInitNodes", 0), // function VisibleChildNoInitNodes
			/* 96 */ imports.NewTable("TBaseVirtualTree_VisibleNoInitNodes", 0), // function VisibleNoInitNodes
			/* 97 */ imports.NewTable("TBaseVirtualTree_AddFromStream", 0), // procedure AddFromStream
			/* 98 */ imports.NewTable("TBaseVirtualTree_BeginDragWithBoolInt", 0), // procedure BeginDragWithBoolInt
			/* 99 */ imports.NewTable("TBaseVirtualTree_BeginSynch", 0), // procedure BeginSynch
			/* 100 */ imports.NewTable("TBaseVirtualTree_BeginUpdate", 0), // procedure BeginUpdate
			/* 101 */ imports.NewTable("TBaseVirtualTree_CancelCutOrCopy", 0), // procedure CancelCutOrCopy
			/* 102 */ imports.NewTable("TBaseVirtualTree_CancelOperation", 0), // procedure CancelOperation
			/* 103 */ imports.NewTable("TBaseVirtualTree_Clear", 0), // procedure Clear
			/* 104 */ imports.NewTable("TBaseVirtualTree_ClearChecked", 0), // procedure ClearChecked
			/* 105 */ imports.NewTable("TBaseVirtualTree_ClearSelection", 0), // procedure ClearSelection
			/* 106 */ imports.NewTable("TBaseVirtualTree_CopyToClipboard", 0), // procedure CopyToClipboard
			/* 107 */ imports.NewTable("TBaseVirtualTree_CutToClipboard", 0), // procedure CutToClipboard
			/* 108 */ imports.NewTable("TBaseVirtualTree_DeleteChildren", 0), // procedure DeleteChildren
			/* 109 */ imports.NewTable("TBaseVirtualTree_DeleteNode", 0), // procedure DeleteNode
			/* 110 */ imports.NewTable("TBaseVirtualTree_DeleteSelectedNodes", 0), // procedure DeleteSelectedNodes
			/* 111 */ imports.NewTable("TBaseVirtualTree_EndSynch", 0), // procedure EndSynch
			/* 112 */ imports.NewTable("TBaseVirtualTree_EndUpdate", 0), // procedure EndUpdate
			/* 113 */ imports.NewTable("TBaseVirtualTree_EnsureNodeSelected", 0), // procedure EnsureNodeSelected
			/* 114 */ imports.NewTable("TBaseVirtualTree_FinishCutOrCopy", 0), // procedure FinishCutOrCopy
			/* 115 */ imports.NewTable("TBaseVirtualTree_FlushClipboard", 0), // procedure FlushClipboard
			/* 116 */ imports.NewTable("TBaseVirtualTree_FullCollapse", 0), // procedure FullCollapse
			/* 117 */ imports.NewTable("TBaseVirtualTree_FullExpand", 0), // procedure FullExpand
			/* 118 */ imports.NewTable("TBaseVirtualTree_GetHitTestInfoAt", 0), // procedure GetHitTestInfoAt
			/* 119 */ imports.NewTable("TBaseVirtualTree_GetTextInfo", 0), // procedure GetTextInfo
			/* 120 */ imports.NewTable("TBaseVirtualTree_InvalidateChildren", 0), // procedure InvalidateChildren
			/* 121 */ imports.NewTable("TBaseVirtualTree_InvalidateColumn", 0), // procedure InvalidateColumn
			/* 122 */ imports.NewTable("TBaseVirtualTree_InvalidateToBottom", 0), // procedure InvalidateToBottom
			/* 123 */ imports.NewTable("TBaseVirtualTree_InvertSelection", 0), // procedure InvertSelection
			/* 124 */ imports.NewTable("TBaseVirtualTree_LoadFromFile", 0), // procedure LoadFromFile
			/* 125 */ imports.NewTable("TBaseVirtualTree_LoadFromStream", 0), // procedure LoadFromStream
			/* 126 */ imports.NewTable("TBaseVirtualTree_MeasureItemHeight", 0), // procedure MeasureItemHeight
			/* 127 */ imports.NewTable("TBaseVirtualTree_MoveToWithPVirtualNodeX2VTNodeAttachModeBool", 0), // procedure MoveToWithPVirtualNodeX2VTNodeAttachModeBool
			/* 128 */ imports.NewTable("TBaseVirtualTree_MoveToWithPVirtualNodeBaseVirtualTreeVTNodeAttachModeBool", 0), // procedure MoveToWithPVirtualNodeBaseVirtualTreeVTNodeAttachModeBool
			/* 129 */ imports.NewTable("TBaseVirtualTree_PaintTree", 0), // procedure PaintTree
			/* 130 */ imports.NewTable("TBaseVirtualTree_RepaintNode", 0), // procedure RepaintNode
			/* 131 */ imports.NewTable("TBaseVirtualTree_ReinitChildren", 0), // procedure ReinitChildren
			/* 132 */ imports.NewTable("TBaseVirtualTree_ReinitNode", 0), // procedure ReinitNode
			/* 133 */ imports.NewTable("TBaseVirtualTree_ResetNode", 0), // procedure ResetNode
			/* 134 */ imports.NewTable("TBaseVirtualTree_SaveToFile", 0), // procedure SaveToFile
			/* 135 */ imports.NewTable("TBaseVirtualTree_SaveToStream", 0), // procedure SaveToStream
			/* 136 */ imports.NewTable("TBaseVirtualTree_SelectAll", 0), // procedure SelectAll
			/* 137 */ imports.NewTable("TBaseVirtualTree_Sort", 0), // procedure Sort
			/* 138 */ imports.NewTable("TBaseVirtualTree_SortTree", 0), // procedure SortTree
			/* 139 */ imports.NewTable("TBaseVirtualTree_ToggleNode", 0), // procedure ToggleNode
			/* 140 */ imports.NewTable("TBaseVirtualTree_UpdateHorizontalRange", 0), // procedure UpdateHorizontalRange
			/* 141 */ imports.NewTable("TBaseVirtualTree_UpdateHorizontalScrollBar", 0), // procedure UpdateHorizontalScrollBar
			/* 142 */ imports.NewTable("TBaseVirtualTree_UpdateRanges", 0), // procedure UpdateRanges
			/* 143 */ imports.NewTable("TBaseVirtualTree_UpdateScrollBars", 0), // procedure UpdateScrollBars
			/* 144 */ imports.NewTable("TBaseVirtualTree_UpdateVerticalRange", 0), // procedure UpdateVerticalRange
			/* 145 */ imports.NewTable("TBaseVirtualTree_UpdateVerticalScrollBar", 0), // procedure UpdateVerticalScrollBar
			/* 146 */ imports.NewTable("TBaseVirtualTree_ValidateChildren", 0), // procedure ValidateChildren
			/* 147 */ imports.NewTable("TBaseVirtualTree_ValidateNode", 0), // procedure ValidateNode
			/* 148 */ imports.NewTable("TBaseVirtualTree_BottomNode", 0), // property BottomNode
			/* 149 */ imports.NewTable("TBaseVirtualTree_CheckedCount", 0), // property CheckedCount
			/* 150 */ imports.NewTable("TBaseVirtualTree_CheckImages", 0), // property CheckImages
			/* 151 */ imports.NewTable("TBaseVirtualTree_CheckState", 0), // property CheckState
			/* 152 */ imports.NewTable("TBaseVirtualTree_CheckType", 0), // property CheckType
			/* 153 */ imports.NewTable("TBaseVirtualTree_ChildCount", 0), // property ChildCount
			/* 154 */ imports.NewTable("TBaseVirtualTree_ChildrenInitialized", 0), // property ChildrenInitialized
			/* 155 */ imports.NewTable("TBaseVirtualTree_CutCopyCount", 0), // property CutCopyCount
			/* 156 */ imports.NewTable("TBaseVirtualTree_DragImage", 0), // property DragImage
			/* 157 */ imports.NewTable("TBaseVirtualTree_VTVDragManager", 0), // property VTVDragManager
			/* 158 */ imports.NewTable("TBaseVirtualTree_DropTargetNode", 0), // property DropTargetNode
			/* 159 */ imports.NewTable("TBaseVirtualTree_EditLink", 0), // property EditLink
			/* 160 */ imports.NewTable("TBaseVirtualTree_EmptyListMessage", 0), // property EmptyListMessage
			/* 161 */ imports.NewTable("TBaseVirtualTree_Expanded", 0), // property Expanded
			/* 162 */ imports.NewTable("TBaseVirtualTree_FocusedColumn", 0), // property FocusedColumn
			/* 163 */ imports.NewTable("TBaseVirtualTree_FocusedNode", 0), // property FocusedNode
			/* 164 */ imports.NewTable("TBaseVirtualTree_FullyVisible", 0), // property FullyVisible
			/* 165 */ imports.NewTable("TBaseVirtualTree_HasChildren", 0), // property HasChildren
			/* 166 */ imports.NewTable("TBaseVirtualTree_HotNode", 0), // property HotNode
			/* 167 */ imports.NewTable("TBaseVirtualTree_IsDisabled", 0), // property IsDisabled
			/* 168 */ imports.NewTable("TBaseVirtualTree_IsEffectivelyFiltered", 0), // property IsEffectivelyFiltered
			/* 169 */ imports.NewTable("TBaseVirtualTree_IsEffectivelyVisible", 0), // property IsEffectivelyVisible
			/* 170 */ imports.NewTable("TBaseVirtualTree_IsFiltered", 0), // property IsFiltered
			/* 171 */ imports.NewTable("TBaseVirtualTree_IsVisibleWithPVirtualNodeToBool", 0), // property IsVisibleWithPVirtualNodeToBool
			/* 172 */ imports.NewTable("TBaseVirtualTree_MultiLine", 0), // property MultiLine
			/* 173 */ imports.NewTable("TBaseVirtualTree_NodeHeight", 0), // property NodeHeight
			/* 174 */ imports.NewTable("TBaseVirtualTree_NodeParent", 0), // property NodeParent
			/* 175 */ imports.NewTable("TBaseVirtualTree_OffsetX", 0), // property OffsetX
			/* 176 */ imports.NewTable("TBaseVirtualTree_OffsetXY", 0), // property OffsetXY
			/* 177 */ imports.NewTable("TBaseVirtualTree_OffsetY", 0), // property OffsetY
			/* 178 */ imports.NewTable("TBaseVirtualTree_OperationCount", 0), // property OperationCount
			/* 179 */ imports.NewTable("TBaseVirtualTree_RootNode", 0), // property RootNode
			/* 180 */ imports.NewTable("TBaseVirtualTree_SearchBuffer", 0), // property SearchBuffer
			/* 181 */ imports.NewTable("TBaseVirtualTree_Selected", 0), // property Selected
			/* 182 */ imports.NewTable("TBaseVirtualTree_SelectionLocked", 0), // property SelectionLocked
			/* 183 */ imports.NewTable("TBaseVirtualTree_TotalCount", 0), // property TotalCount
			/* 184 */ imports.NewTable("TBaseVirtualTree_TreeStates", 0), // property TreeStates
			/* 185 */ imports.NewTable("TBaseVirtualTree_SelectedCount", 0), // property SelectedCount
			/* 186 */ imports.NewTable("TBaseVirtualTree_TopNode", 0), // property TopNode
			/* 187 */ imports.NewTable("TBaseVirtualTree_VerticalAlignment", 0), // property VerticalAlignment
			/* 188 */ imports.NewTable("TBaseVirtualTree_VisibleCount", 0), // property VisibleCount
			/* 189 */ imports.NewTable("TBaseVirtualTree_VisiblePath", 0), // property VisiblePath
			/* 190 */ imports.NewTable("TBaseVirtualTree_UpdateCount", 0), // property UpdateCount
			/* 191 */ imports.NewTable("TBaseVirtualTree_TClass", 0), // function TBaseVirtualTreeClass
		}
	})
	return baseVirtualTreeImport
}
