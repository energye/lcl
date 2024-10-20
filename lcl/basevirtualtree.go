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

// IBaseVirtualTree Parent: ICustomControl
// ----- TBaseVirtualTree
type IBaseVirtualTree interface {
	ICustomControl
	BottomNode() IVirtualNode                                                                                                             // property
	SetBottomNode(AValue IVirtualNode)                                                                                                    // property
	CheckedCount() int32                                                                                                                  // property
	CheckImages() ICustomImageList                                                                                                        // property
	CheckState(Node IVirtualNode) TCheckState                                                                                             // property
	SetCheckState(Node IVirtualNode, AValue TCheckState)                                                                                  // property
	CheckType(Node IVirtualNode) TCheckType                                                                                               // property
	SetCheckType(Node IVirtualNode, AValue TCheckType)                                                                                    // property
	ChildCount(Node IVirtualNode) uint32                                                                                                  // property
	SetChildCount(Node IVirtualNode, AValue uint32)                                                                                       // property
	ChildrenInitialized(Node IVirtualNode) bool                                                                                           // property
	CutCopyCount() int32                                                                                                                  // property
	DragImage() IVTDragImage                                                                                                              // property
	DropTargetNode() IVirtualNode                                                                                                         // property
	SetDropTargetNode(AValue IVirtualNode)                                                                                                // property
	EmptyListMessage() string                                                                                                             // property
	SetEmptyListMessage(AValue string)                                                                                                    // property
	Expanded(Node IVirtualNode) bool                                                                                                      // property
	SetExpanded(Node IVirtualNode, AValue bool)                                                                                           // property
	FocusedColumn() TColumnIndex                                                                                                          // property
	SetFocusedColumn(AValue TColumnIndex)                                                                                                 // property
	FocusedNode() IVirtualNode                                                                                                            // property
	SetFocusedNode(AValue IVirtualNode)                                                                                                   // property
	FullyVisible(Node IVirtualNode) bool                                                                                                  // property
	SetFullyVisible(Node IVirtualNode, AValue bool)                                                                                       // property
	HasChildren(Node IVirtualNode) bool                                                                                                   // property
	SetHasChildren(Node IVirtualNode, AValue bool)                                                                                        // property
	HotNode() IVirtualNode                                                                                                                // property
	IsDisabled(Node IVirtualNode) bool                                                                                                    // property
	SetIsDisabled(Node IVirtualNode, AValue bool)                                                                                         // property
	IsEffectivelyFiltered(Node IVirtualNode) bool                                                                                         // property
	IsEffectivelyVisible(Node IVirtualNode) bool                                                                                          // property
	IsFiltered(Node IVirtualNode) bool                                                                                                    // property
	SetIsFiltered(Node IVirtualNode, AValue bool)                                                                                         // property
	NodeIsVisible(Node IVirtualNode) bool                                                                                                 // property
	SetNodeIsVisible(Node IVirtualNode, AValue bool)                                                                                      // property
	MultiLine(Node IVirtualNode) bool                                                                                                     // property
	SetMultiLine(Node IVirtualNode, AValue bool)                                                                                          // property
	NodeHeight(Node IVirtualNode) uint32                                                                                                  // property
	SetNodeHeight(Node IVirtualNode, AValue uint32)                                                                                       // property
	NodeParent(Node IVirtualNode) IVirtualNode                                                                                            // property
	SetNodeParent(Node IVirtualNode, AValue IVirtualNode)                                                                                 // property
	OffsetX() int32                                                                                                                       // property
	SetOffsetX(AValue int32)                                                                                                              // property
	OffsetXY() (resultPoint TPoint)                                                                                                       // property
	SetOffsetXY(AValue *TPoint)                                                                                                           // property
	OffsetY() int32                                                                                                                       // property
	SetOffsetY(AValue int32)                                                                                                              // property
	OperationCount() uint32                                                                                                               // property
	RootNode() IVirtualNode                                                                                                               // property
	SearchBuffer() string                                                                                                                 // property
	Selected(Node IVirtualNode) bool                                                                                                      // property
	SetSelected(Node IVirtualNode, AValue bool)                                                                                           // property
	SelectionLocked() bool                                                                                                                // property
	SetSelectionLocked(AValue bool)                                                                                                       // property
	TotalCount() uint32                                                                                                                   // property
	TreeStates() TVirtualTreeStates                                                                                                       // property
	SetTreeStates(AValue TVirtualTreeStates)                                                                                              // property
	SelectedCount() int32                                                                                                                 // property
	TopNode() IVirtualNode                                                                                                                // property
	SetTopNode(AValue IVirtualNode)                                                                                                       // property
	VerticalAlignment(Node IVirtualNode) Byte                                                                                             // property
	SetVerticalAlignment(Node IVirtualNode, AValue Byte)                                                                                  // property
	VisibleCount() uint32                                                                                                                 // property
	VisiblePath(Node IVirtualNode) bool                                                                                                   // property
	SetVisiblePath(Node IVirtualNode, AValue bool)                                                                                        // property
	UpdateCount() uint32                                                                                                                  // property
	AbsoluteIndex(Node IVirtualNode) uint32                                                                                               // function
	AddChild(Parent IVirtualNode, UserData uintptr) IVirtualNode                                                                          // function
	CancelEditNode() bool                                                                                                                 // function
	CanEdit(Node IVirtualNode, Column TColumnIndex) bool                                                                                  // function
	CopyTo(Source IVirtualNode, Tree IBaseVirtualTree, Mode TVTNodeAttachMode, ChildrenOnly bool) IVirtualNode                            // function
	CopyTo1(Source, Target IVirtualNode, Mode TVTNodeAttachMode, ChildrenOnly bool) IVirtualNode                                          // function
	EditNode(Node IVirtualNode, Column TColumnIndex) bool                                                                                 // function
	EndEditNode() bool                                                                                                                    // function
	GetDisplayRect(Node IVirtualNode, Column TColumnIndex, TextOnly bool, Unclipped bool, ApplyCellContentMargin bool) (resultRect TRect) // function
	GetEffectivelyFiltered(Node IVirtualNode) bool                                                                                        // function
	GetEffectivelyVisible(Node IVirtualNode) bool                                                                                         // function
	GetFirst(ConsiderChildrenAbove bool) IVirtualNode                                                                                     // function
	GetFirstChecked(State TCheckState, ConsiderChildrenAbove bool) IVirtualNode                                                           // function
	GetFirstChild(Node IVirtualNode) IVirtualNode                                                                                         // function
	GetFirstChildNoInit(Node IVirtualNode) IVirtualNode                                                                                   // function
	GetFirstCutCopy(ConsiderChildrenAbove bool) IVirtualNode                                                                              // function
	GetFirstInitialized(ConsiderChildrenAbove bool) IVirtualNode                                                                          // function
	GetFirstLeaf() IVirtualNode                                                                                                           // function
	GetFirstLevel(NodeLevel uint32) IVirtualNode                                                                                          // function
	GetFirstNoInit(ConsiderChildrenAbove bool) IVirtualNode                                                                               // function
	GetFirstSelected(ConsiderChildrenAbove bool) IVirtualNode                                                                             // function
	GetFirstVisible(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode                                     // function
	GetFirstVisibleChild(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                            // function
	GetFirstVisibleChildNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                      // function
	GetFirstVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode                               // function
	GetLast(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                                   // function
	GetLastInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                        // function
	GetLastNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                             // function
	GetLastChild(Node IVirtualNode) IVirtualNode                                                                                          // function
	GetLastChildNoInit(Node IVirtualNode) IVirtualNode                                                                                    // function
	GetLastVisible(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode                                      // function
	GetLastVisibleChild(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                             // function
	GetLastVisibleChildNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                       // function
	GetLastVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode                                // function
	GetMaxColumnWidth(Column TColumnIndex, UseSmartColumnWidth bool) int32                                                                // function
	GetNext(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                                   // function
	GetNextChecked(Node IVirtualNode, State TCheckState, ConsiderChildrenAbove bool) IVirtualNode                                         // function
	GetNextChecked1(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                           // function
	GetNextCutCopy(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                            // function
	GetNextInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                        // function
	GetNextLeaf(Node IVirtualNode) IVirtualNode                                                                                           // function
	GetNextLevel(Node IVirtualNode, NodeLevel uint32) IVirtualNode                                                                        // function
	GetNextNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                             // function
	GetNextSelected(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                           // function
	GetNextSibling(Node IVirtualNode) IVirtualNode                                                                                        // function
	GetNextSiblingNoInit(Node IVirtualNode) IVirtualNode                                                                                  // function
	GetNextVisible(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                            // function
	GetNextVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                      // function
	GetNextVisibleSibling(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                           // function
	GetNextVisibleSiblingNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                     // function
	GetNodeAt(P *TPoint) IVirtualNode                                                                                                     // function
	GetNodeAt1(X, Y int32) IVirtualNode                                                                                                   // function
	GetNodeAt2(X, Y int32, Relative bool, NodeTop *int32) IVirtualNode                                                                    // function
	GetNodeData(Node IVirtualNode) uintptr                                                                                                // function
	GetNodeLevel(Node IVirtualNode) uint32                                                                                                // function
	GetPrevious(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                               // function
	GetPreviousChecked(Node IVirtualNode, State TCheckState, ConsiderChildrenAbove bool) IVirtualNode                                     // function
	GetPreviousCutCopy(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                        // function
	GetPreviousInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                    // function
	GetPreviousLeaf(Node IVirtualNode) IVirtualNode                                                                                       // function
	GetPreviousLevel(Node IVirtualNode, NodeLevel uint32) IVirtualNode                                                                    // function
	GetPreviousNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                         // function
	GetPreviousSelected(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                       // function
	GetPreviousSibling(Node IVirtualNode) IVirtualNode                                                                                    // function
	GetPreviousSiblingNoInit(Node IVirtualNode) IVirtualNode                                                                              // function
	GetPreviousVisible(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                        // function
	GetPreviousVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                  // function
	GetPreviousVisibleSibling(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                       // function
	GetPreviousVisibleSiblingNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                 // function
	GetSortedCutCopySet(Resolve bool) TNodeArray                                                                                          // function
	GetSortedSelection(Resolve bool) TNodeArray                                                                                           // function
	GetTreeRect() (resultRect TRect)                                                                                                      // function
	GetVisibleParent(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                                // function
	HasAsParent(Node, PotentialParent IVirtualNode) bool                                                                                  // function
	InsertNode(Node IVirtualNode, Mode TVTNodeAttachMode, UserData uintptr) IVirtualNode                                                  // function
	InvalidateNode(Node IVirtualNode) (resultRect TRect)                                                                                  // function
	IsEditing() bool                                                                                                                      // function
	IsMouseSelecting() bool                                                                                                               // function
	IsEmpty() bool                                                                                                                        // function
	PasteFromClipboard() bool                                                                                                             // function
	ScrollIntoView(Node IVirtualNode, Center bool, Horizontally bool) bool                                                                // function
	ScrollIntoView1(Column TColumnIndex, Center bool) bool                                                                                // function
	Nodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                           // function
	CheckedNodes(State TCheckState, ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                 // function
	ChildNodes(Node IVirtualNode) IVTVirtualNodeEnumeration                                                                               // function
	CutCopyNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                    // function
	InitializedNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                // function
	LeafNodes() IVTVirtualNodeEnumeration                                                                                                 // function
	LevelNodes(NodeLevel uint32) IVTVirtualNodeEnumeration                                                                                // function
	NoInitNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                     // function
	SelectedNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                   // function
	VisibleNodes(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVTVirtualNodeEnumeration                           // function
	VisibleChildNodes(Node IVirtualNode, IncludeFiltered bool) IVTVirtualNodeEnumeration                                                  // function
	VisibleChildNoInitNodes(Node IVirtualNode, IncludeFiltered bool) IVTVirtualNodeEnumeration                                            // function
	VisibleNoInitNodes(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVTVirtualNodeEnumeration                     // function
	AddFromStream(Stream IStream, TargetNode IVirtualNode)                                                                                // procedure
	BeginSynch()                                                                                                                          // procedure
	BeginUpdate()                                                                                                                         // procedure
	CancelCutOrCopy()                                                                                                                     // procedure
	CancelOperation()                                                                                                                     // procedure
	Clear()                                                                                                                               // procedure
	ClearChecked()                                                                                                                        // procedure
	ClearSelection()                                                                                                                      // procedure
	CopyToClipboard()                                                                                                                     // procedure
	CutToClipboard()                                                                                                                      // procedure
	DeleteChildren(Node IVirtualNode, ResetHasChildren bool)                                                                              // procedure
	DeleteNode(Node IVirtualNode, Reindex bool)                                                                                           // procedure
	DeleteSelectedNodes()                                                                                                                 // procedure
	EndSynch()                                                                                                                            // procedure
	EndUpdate()                                                                                                                           // procedure
	EnsureNodeSelected()                                                                                                                  // procedure
	FinishCutOrCopy()                                                                                                                     // procedure
	FlushClipboard()                                                                                                                      // procedure
	FullCollapse(Node IVirtualNode)                                                                                                       // procedure
	FullExpand(Node IVirtualNode)                                                                                                         // procedure
	GetTextInfo(Node IVirtualNode, Column TColumnIndex, AFont IFont, R *TRect, OutText *string)                                           // procedure
	InvalidateChildren(Node IVirtualNode, Recursive bool)                                                                                 // procedure
	InvalidateColumn(Column TColumnIndex)                                                                                                 // procedure
	InvalidateToBottom(Node IVirtualNode)                                                                                                 // procedure
	InvertSelection(VisibleOnly bool)                                                                                                     // procedure
	LoadFromFile(FileName string)                                                                                                         // procedure
	LoadFromStream(Stream IStream)                                                                                                        // procedure
	MeasureItemHeight(Canvas ICanvas, Node IVirtualNode)                                                                                  // procedure
	MoveTo(Source, Target IVirtualNode, Mode TVTNodeAttachMode, ChildrenOnly bool)                                                        // procedure
	MoveTo1(Node IVirtualNode, Tree IBaseVirtualTree, Mode TVTNodeAttachMode, ChildrenOnly bool)                                          // procedure
	PaintTree(TargetCanvas ICanvas, Window *TRect, Target *TPoint, PaintOptions TVTInternalPaintOptions, PixelFormat TPixelFormat)        // procedure
	RepaintNode(Node IVirtualNode)                                                                                                        // procedure
	ReinitChildren(Node IVirtualNode, Recursive bool)                                                                                     // procedure
	ReinitNode(Node IVirtualNode, Recursive bool)                                                                                         // procedure
	ResetNode(Node IVirtualNode)                                                                                                          // procedure
	SaveToFile(FileName string)                                                                                                           // procedure
	SaveToStream(Stream IStream, Node IVirtualNode)                                                                                       // procedure
	SelectAll(VisibleOnly bool)                                                                                                           // procedure
	Sort(Node IVirtualNode, Column TColumnIndex, Direction TSortDirection, DoInit bool)                                                   // procedure
	SortTree(Column TColumnIndex, Direction TSortDirection, DoInit bool)                                                                  // procedure
	ToggleNode(Node IVirtualNode)                                                                                                         // procedure
	UpdateHorizontalRange()                                                                                                               // procedure
	UpdateHorizontalScrollBar(DoRepaint bool)                                                                                             // procedure
	UpdateRanges()                                                                                                                        // procedure
	UpdateScrollBars(DoRepaint bool)                                                                                                      // procedure
	UpdateVerticalRange()                                                                                                                 // procedure
	UpdateVerticalScrollBar(DoRepaint bool)                                                                                               // procedure
	ValidateChildren(Node IVirtualNode, Recursive bool)                                                                                   // procedure
	ValidateNode(Node IVirtualNode, Recursive bool)                                                                                       // procedure
}

// TBaseVirtualTree Parent: TCustomControl
// ----- TBaseVirtualTree
type TBaseVirtualTree struct {
	TCustomControl
}

func NewBaseVirtualTree(AOwner IComponent) IBaseVirtualTree {
	r1 := baseVirtualTreeImportAPI().SysCallN(25, GetObjectUintptr(AOwner))
	return AsBaseVirtualTree(r1)
}

func (m *TBaseVirtualTree) BottomNode() IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetBottomNode(AValue IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) CheckedCount() int32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(13, m.Instance())
	return int32(r1)
}

func (m *TBaseVirtualTree) CheckImages() ICustomImageList {
	r1 := baseVirtualTreeImportAPI().SysCallN(10, m.Instance())
	return AsCustomImageList(r1)
}

func (m *TBaseVirtualTree) CheckState(Node IVirtualNode) TCheckState {
	r1 := baseVirtualTreeImportAPI().SysCallN(11, 0, m.Instance(), GetObjectUintptr(Node))
	return TCheckState(r1)
}

func (m *TBaseVirtualTree) SetCheckState(Node IVirtualNode, AValue TCheckState) {
	baseVirtualTreeImportAPI().SysCallN(11, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) CheckType(Node IVirtualNode) TCheckType {
	r1 := baseVirtualTreeImportAPI().SysCallN(12, 0, m.Instance(), GetObjectUintptr(Node))
	return TCheckType(r1)
}

func (m *TBaseVirtualTree) SetCheckType(Node IVirtualNode, AValue TCheckType) {
	baseVirtualTreeImportAPI().SysCallN(12, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) ChildCount(Node IVirtualNode) uint32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(15, 0, m.Instance(), GetObjectUintptr(Node))
	return uint32(r1)
}

func (m *TBaseVirtualTree) SetChildCount(Node IVirtualNode, AValue uint32) {
	baseVirtualTreeImportAPI().SysCallN(15, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) ChildrenInitialized(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) CutCopyCount() int32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(26, m.Instance())
	return int32(r1)
}

func (m *TBaseVirtualTree) DragImage() IVTDragImage {
	r1 := baseVirtualTreeImportAPI().SysCallN(32, m.Instance())
	return AsVTDragImage(r1)
}

func (m *TBaseVirtualTree) DropTargetNode() IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetDropTargetNode(AValue IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(33, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) EmptyListMessage() string {
	r1 := baseVirtualTreeImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TBaseVirtualTree) SetEmptyListMessage(AValue string) {
	baseVirtualTreeImportAPI().SysCallN(35, 1, m.Instance(), PascalStr(AValue))
}

func (m *TBaseVirtualTree) Expanded(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(40, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetExpanded(Node IVirtualNode, AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(40, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) FocusedColumn() TColumnIndex {
	r1 := baseVirtualTreeImportAPI().SysCallN(43, 0, m.Instance(), 0)
	return TColumnIndex(r1)
}

func (m *TBaseVirtualTree) SetFocusedColumn(AValue TColumnIndex) {
	baseVirtualTreeImportAPI().SysCallN(43, 1, m.Instance(), uintptr(AValue))
}

func (m *TBaseVirtualTree) FocusedNode() IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(44, 0, m.Instance(), 0)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetFocusedNode(AValue IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(44, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) FullyVisible(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(47, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetFullyVisible(Node IVirtualNode, AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(47, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) HasChildren(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(115, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetHasChildren(Node IVirtualNode, AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(115, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) HotNode() IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(116, m.Instance())
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) IsDisabled(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(124, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetIsDisabled(Node IVirtualNode, AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(124, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) IsEffectivelyFiltered(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(126, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) IsEffectivelyVisible(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(127, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) IsFiltered(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(129, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetIsFiltered(Node IVirtualNode, AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(129, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) NodeIsVisible(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(141, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetNodeIsVisible(Node IVirtualNode, AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(141, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) MultiLine(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(138, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetMultiLine(Node IVirtualNode, AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(138, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) NodeHeight(Node IVirtualNode) uint32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(140, 0, m.Instance(), GetObjectUintptr(Node))
	return uint32(r1)
}

func (m *TBaseVirtualTree) SetNodeHeight(Node IVirtualNode, AValue uint32) {
	baseVirtualTreeImportAPI().SysCallN(140, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) NodeParent(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(142, 0, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetNodeParent(Node IVirtualNode, AValue IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(142, 1, m.Instance(), GetObjectUintptr(Node), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) OffsetX() int32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(144, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TBaseVirtualTree) SetOffsetX(AValue int32) {
	baseVirtualTreeImportAPI().SysCallN(144, 1, m.Instance(), uintptr(AValue))
}

func (m *TBaseVirtualTree) OffsetXY() (resultPoint TPoint) {
	baseVirtualTreeImportAPI().SysCallN(145, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TBaseVirtualTree) SetOffsetXY(AValue *TPoint) {
	baseVirtualTreeImportAPI().SysCallN(145, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TBaseVirtualTree) OffsetY() int32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(146, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TBaseVirtualTree) SetOffsetY(AValue int32) {
	baseVirtualTreeImportAPI().SysCallN(146, 1, m.Instance(), uintptr(AValue))
}

func (m *TBaseVirtualTree) OperationCount() uint32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(147, m.Instance())
	return uint32(r1)
}

func (m *TBaseVirtualTree) RootNode() IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(154, m.Instance())
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SearchBuffer() string {
	r1 := baseVirtualTreeImportAPI().SysCallN(159, m.Instance())
	return GoStr(r1)
}

func (m *TBaseVirtualTree) Selected(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(161, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetSelected(Node IVirtualNode, AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(161, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) SelectionLocked() bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(164, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetSelectionLocked(AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(164, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBaseVirtualTree) TotalCount() uint32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(169, m.Instance())
	return uint32(r1)
}

func (m *TBaseVirtualTree) TreeStates() TVirtualTreeStates {
	r1 := baseVirtualTreeImportAPI().SysCallN(170, 0, m.Instance(), 0)
	return TVirtualTreeStates(r1)
}

func (m *TBaseVirtualTree) SetTreeStates(AValue TVirtualTreeStates) {
	baseVirtualTreeImportAPI().SysCallN(170, 1, m.Instance(), uintptr(AValue))
}

func (m *TBaseVirtualTree) SelectedCount() int32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(162, m.Instance())
	return int32(r1)
}

func (m *TBaseVirtualTree) TopNode() IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(168, 0, m.Instance(), 0)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetTopNode(AValue IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(168, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) VerticalAlignment(Node IVirtualNode) Byte {
	r1 := baseVirtualTreeImportAPI().SysCallN(180, 0, m.Instance(), GetObjectUintptr(Node))
	return Byte(r1)
}

func (m *TBaseVirtualTree) SetVerticalAlignment(Node IVirtualNode, AValue Byte) {
	baseVirtualTreeImportAPI().SysCallN(180, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) VisibleCount() uint32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(183, m.Instance())
	return uint32(r1)
}

func (m *TBaseVirtualTree) VisiblePath(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(186, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetVisiblePath(Node IVirtualNode, AValue bool) {
	baseVirtualTreeImportAPI().SysCallN(186, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) UpdateCount() uint32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(171, m.Instance())
	return uint32(r1)
}

func (m *TBaseVirtualTree) AbsoluteIndex(Node IVirtualNode) uint32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(Node))
	return uint32(r1)
}

func (m *TBaseVirtualTree) AddChild(Parent IVirtualNode, UserData uintptr) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(Parent), uintptr(UserData))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) CancelEditNode() bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) CanEdit(Node IVirtualNode, Column TColumnIndex) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(Node), uintptr(Column))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) CopyTo(Source IVirtualNode, Tree IBaseVirtualTree, Mode TVTNodeAttachMode, ChildrenOnly bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(22, m.Instance(), GetObjectUintptr(Source), GetObjectUintptr(Tree), uintptr(Mode), PascalBool(ChildrenOnly))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) CopyTo1(Source, Target IVirtualNode, Mode TVTNodeAttachMode, ChildrenOnly bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(23, m.Instance(), GetObjectUintptr(Source), GetObjectUintptr(Target), uintptr(Mode), PascalBool(ChildrenOnly))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) EditNode(Node IVirtualNode, Column TColumnIndex) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(34, m.Instance(), GetObjectUintptr(Node), uintptr(Column))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) EndEditNode() bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(36, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) GetDisplayRect(Node IVirtualNode, Column TColumnIndex, TextOnly bool, Unclipped bool, ApplyCellContentMargin bool) (resultRect TRect) {
	baseVirtualTreeImportAPI().SysCallN(48, m.Instance(), GetObjectUintptr(Node), uintptr(Column), PascalBool(TextOnly), PascalBool(Unclipped), PascalBool(ApplyCellContentMargin), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TBaseVirtualTree) GetEffectivelyFiltered(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(49, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) GetEffectivelyVisible(Node IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(50, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) GetFirst(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(51, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstChecked(State TCheckState, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(52, m.Instance(), uintptr(State), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstChild(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(53, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstChildNoInit(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(54, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstCutCopy(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(55, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstInitialized(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(56, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstLeaf() IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(57, m.Instance())
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstLevel(NodeLevel uint32) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(58, m.Instance(), uintptr(NodeLevel))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstNoInit(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(59, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstSelected(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(60, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstVisible(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(61, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstVisibleChild(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(62, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstVisibleChildNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(63, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(64, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLast(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(65, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(68, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(69, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastChild(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(66, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastChildNoInit(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(67, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastVisible(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(70, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastVisibleChild(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(71, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastVisibleChildNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(72, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(73, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetMaxColumnWidth(Column TColumnIndex, UseSmartColumnWidth bool) int32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(74, m.Instance(), uintptr(Column), PascalBool(UseSmartColumnWidth))
	return int32(r1)
}

func (m *TBaseVirtualTree) GetNext(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(75, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextChecked(Node IVirtualNode, State TCheckState, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(76, m.Instance(), GetObjectUintptr(Node), uintptr(State), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextChecked1(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(77, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextCutCopy(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(78, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(79, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextLeaf(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(80, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextLevel(Node IVirtualNode, NodeLevel uint32) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(81, m.Instance(), GetObjectUintptr(Node), uintptr(NodeLevel))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(82, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextSelected(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(83, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextSibling(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(84, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextSiblingNoInit(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(85, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextVisible(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(86, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(87, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextVisibleSibling(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(88, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextVisibleSiblingNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(89, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNodeAt(P *TPoint) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(90, m.Instance(), uintptr(unsafePointer(P)))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNodeAt1(X, Y int32) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(91, m.Instance(), uintptr(X), uintptr(Y))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNodeAt2(X, Y int32, Relative bool, NodeTop *int32) IVirtualNode {
	var result2 uintptr
	r1 := baseVirtualTreeImportAPI().SysCallN(92, m.Instance(), uintptr(X), uintptr(Y), PascalBool(Relative), uintptr(unsafePointer(&result2)))
	*NodeTop = int32(result2)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNodeData(Node IVirtualNode) uintptr {
	r1 := baseVirtualTreeImportAPI().SysCallN(93, m.Instance(), GetObjectUintptr(Node))
	return uintptr(r1)
}

func (m *TBaseVirtualTree) GetNodeLevel(Node IVirtualNode) uint32 {
	r1 := baseVirtualTreeImportAPI().SysCallN(94, m.Instance(), GetObjectUintptr(Node))
	return uint32(r1)
}

func (m *TBaseVirtualTree) GetPrevious(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(95, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousChecked(Node IVirtualNode, State TCheckState, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(96, m.Instance(), GetObjectUintptr(Node), uintptr(State), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousCutCopy(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(97, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(98, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousLeaf(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(99, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousLevel(Node IVirtualNode, NodeLevel uint32) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(100, m.Instance(), GetObjectUintptr(Node), uintptr(NodeLevel))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(101, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousSelected(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(102, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousSibling(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(103, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousSiblingNoInit(Node IVirtualNode) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(104, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousVisible(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(105, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(106, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousVisibleSibling(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(107, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousVisibleSiblingNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(108, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetSortedCutCopySet(Resolve bool) TNodeArray {
	r1 := baseVirtualTreeImportAPI().SysCallN(109, m.Instance(), PascalBool(Resolve))
	return TNodeArray(r1)
}

func (m *TBaseVirtualTree) GetSortedSelection(Resolve bool) TNodeArray {
	r1 := baseVirtualTreeImportAPI().SysCallN(110, m.Instance(), PascalBool(Resolve))
	return TNodeArray(r1)
}

func (m *TBaseVirtualTree) GetTreeRect() (resultRect TRect) {
	baseVirtualTreeImportAPI().SysCallN(112, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TBaseVirtualTree) GetVisibleParent(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(113, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) HasAsParent(Node, PotentialParent IVirtualNode) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(114, m.Instance(), GetObjectUintptr(Node), GetObjectUintptr(PotentialParent))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) InsertNode(Node IVirtualNode, Mode TVTNodeAttachMode, UserData uintptr) IVirtualNode {
	r1 := baseVirtualTreeImportAPI().SysCallN(118, m.Instance(), GetObjectUintptr(Node), uintptr(Mode), uintptr(UserData))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) InvalidateNode(Node IVirtualNode) (resultRect TRect) {
	baseVirtualTreeImportAPI().SysCallN(121, m.Instance(), GetObjectUintptr(Node), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TBaseVirtualTree) IsEditing() bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(125, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) IsMouseSelecting() bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(130, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) IsEmpty() bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(128, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) PasteFromClipboard() bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(149, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) ScrollIntoView(Node IVirtualNode, Center bool, Horizontally bool) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(157, m.Instance(), GetObjectUintptr(Node), PascalBool(Center), PascalBool(Horizontally))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) ScrollIntoView1(Column TColumnIndex, Center bool) bool {
	r1 := baseVirtualTreeImportAPI().SysCallN(158, m.Instance(), uintptr(Column), PascalBool(Center))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) Nodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(143, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) CheckedNodes(State TCheckState, ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(14, m.Instance(), uintptr(State), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) ChildNodes(Node IVirtualNode) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(Node))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) CutCopyNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(27, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) InitializedNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(117, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) LeafNodes() IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(131, m.Instance())
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) LevelNodes(NodeLevel uint32) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(132, m.Instance(), uintptr(NodeLevel))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) NoInitNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(139, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) SelectedNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(163, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) VisibleNodes(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(185, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) VisibleChildNodes(Node IVirtualNode, IncludeFiltered bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(182, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) VisibleChildNoInitNodes(Node IVirtualNode, IncludeFiltered bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(181, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) VisibleNoInitNodes(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVTVirtualNodeEnumeration {
	r1 := baseVirtualTreeImportAPI().SysCallN(184, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVTVirtualNodeEnumeration(r1)
}

func BaseVirtualTreeClass() TClass {
	ret := baseVirtualTreeImportAPI().SysCallN(18)
	return TClass(ret)
}

func (m *TBaseVirtualTree) AddFromStream(Stream IStream, TargetNode IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(Stream), GetObjectUintptr(TargetNode))
}

func (m *TBaseVirtualTree) BeginSynch() {
	baseVirtualTreeImportAPI().SysCallN(3, m.Instance())
}

func (m *TBaseVirtualTree) BeginUpdate() {
	baseVirtualTreeImportAPI().SysCallN(4, m.Instance())
}

func (m *TBaseVirtualTree) CancelCutOrCopy() {
	baseVirtualTreeImportAPI().SysCallN(7, m.Instance())
}

func (m *TBaseVirtualTree) CancelOperation() {
	baseVirtualTreeImportAPI().SysCallN(9, m.Instance())
}

func (m *TBaseVirtualTree) Clear() {
	baseVirtualTreeImportAPI().SysCallN(19, m.Instance())
}

func (m *TBaseVirtualTree) ClearChecked() {
	baseVirtualTreeImportAPI().SysCallN(20, m.Instance())
}

func (m *TBaseVirtualTree) ClearSelection() {
	baseVirtualTreeImportAPI().SysCallN(21, m.Instance())
}

func (m *TBaseVirtualTree) CopyToClipboard() {
	baseVirtualTreeImportAPI().SysCallN(24, m.Instance())
}

func (m *TBaseVirtualTree) CutToClipboard() {
	baseVirtualTreeImportAPI().SysCallN(28, m.Instance())
}

func (m *TBaseVirtualTree) DeleteChildren(Node IVirtualNode, ResetHasChildren bool) {
	baseVirtualTreeImportAPI().SysCallN(29, m.Instance(), GetObjectUintptr(Node), PascalBool(ResetHasChildren))
}

func (m *TBaseVirtualTree) DeleteNode(Node IVirtualNode, Reindex bool) {
	baseVirtualTreeImportAPI().SysCallN(30, m.Instance(), GetObjectUintptr(Node), PascalBool(Reindex))
}

func (m *TBaseVirtualTree) DeleteSelectedNodes() {
	baseVirtualTreeImportAPI().SysCallN(31, m.Instance())
}

func (m *TBaseVirtualTree) EndSynch() {
	baseVirtualTreeImportAPI().SysCallN(37, m.Instance())
}

func (m *TBaseVirtualTree) EndUpdate() {
	baseVirtualTreeImportAPI().SysCallN(38, m.Instance())
}

func (m *TBaseVirtualTree) EnsureNodeSelected() {
	baseVirtualTreeImportAPI().SysCallN(39, m.Instance())
}

func (m *TBaseVirtualTree) FinishCutOrCopy() {
	baseVirtualTreeImportAPI().SysCallN(41, m.Instance())
}

func (m *TBaseVirtualTree) FlushClipboard() {
	baseVirtualTreeImportAPI().SysCallN(42, m.Instance())
}

func (m *TBaseVirtualTree) FullCollapse(Node IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(45, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) FullExpand(Node IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(46, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) GetTextInfo(Node IVirtualNode, Column TColumnIndex, AFont IFont, R *TRect, OutText *string) {
	var result3 uintptr
	var result4 uintptr
	baseVirtualTreeImportAPI().SysCallN(111, m.Instance(), GetObjectUintptr(Node), uintptr(Column), GetObjectUintptr(AFont), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*R = *(*TRect)(getPointer(result3))
	*OutText = GoStr(result4)
}

func (m *TBaseVirtualTree) InvalidateChildren(Node IVirtualNode, Recursive bool) {
	baseVirtualTreeImportAPI().SysCallN(119, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}

func (m *TBaseVirtualTree) InvalidateColumn(Column TColumnIndex) {
	baseVirtualTreeImportAPI().SysCallN(120, m.Instance(), uintptr(Column))
}

func (m *TBaseVirtualTree) InvalidateToBottom(Node IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(122, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) InvertSelection(VisibleOnly bool) {
	baseVirtualTreeImportAPI().SysCallN(123, m.Instance(), PascalBool(VisibleOnly))
}

func (m *TBaseVirtualTree) LoadFromFile(FileName string) {
	baseVirtualTreeImportAPI().SysCallN(133, m.Instance(), PascalStr(FileName))
}

func (m *TBaseVirtualTree) LoadFromStream(Stream IStream) {
	baseVirtualTreeImportAPI().SysCallN(134, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TBaseVirtualTree) MeasureItemHeight(Canvas ICanvas, Node IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(135, m.Instance(), GetObjectUintptr(Canvas), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) MoveTo(Source, Target IVirtualNode, Mode TVTNodeAttachMode, ChildrenOnly bool) {
	baseVirtualTreeImportAPI().SysCallN(136, m.Instance(), GetObjectUintptr(Source), GetObjectUintptr(Target), uintptr(Mode), PascalBool(ChildrenOnly))
}

func (m *TBaseVirtualTree) MoveTo1(Node IVirtualNode, Tree IBaseVirtualTree, Mode TVTNodeAttachMode, ChildrenOnly bool) {
	baseVirtualTreeImportAPI().SysCallN(137, m.Instance(), GetObjectUintptr(Node), GetObjectUintptr(Tree), uintptr(Mode), PascalBool(ChildrenOnly))
}

func (m *TBaseVirtualTree) PaintTree(TargetCanvas ICanvas, Window *TRect, Target *TPoint, PaintOptions TVTInternalPaintOptions, PixelFormat TPixelFormat) {
	baseVirtualTreeImportAPI().SysCallN(148, m.Instance(), GetObjectUintptr(TargetCanvas), uintptr(unsafePointer(Window)), uintptr(unsafePointer(Target)), uintptr(PaintOptions), uintptr(PixelFormat))
}

func (m *TBaseVirtualTree) RepaintNode(Node IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(152, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) ReinitChildren(Node IVirtualNode, Recursive bool) {
	baseVirtualTreeImportAPI().SysCallN(150, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}

func (m *TBaseVirtualTree) ReinitNode(Node IVirtualNode, Recursive bool) {
	baseVirtualTreeImportAPI().SysCallN(151, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}

func (m *TBaseVirtualTree) ResetNode(Node IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(153, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) SaveToFile(FileName string) {
	baseVirtualTreeImportAPI().SysCallN(155, m.Instance(), PascalStr(FileName))
}

func (m *TBaseVirtualTree) SaveToStream(Stream IStream, Node IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(156, m.Instance(), GetObjectUintptr(Stream), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) SelectAll(VisibleOnly bool) {
	baseVirtualTreeImportAPI().SysCallN(160, m.Instance(), PascalBool(VisibleOnly))
}

func (m *TBaseVirtualTree) Sort(Node IVirtualNode, Column TColumnIndex, Direction TSortDirection, DoInit bool) {
	baseVirtualTreeImportAPI().SysCallN(165, m.Instance(), GetObjectUintptr(Node), uintptr(Column), uintptr(Direction), PascalBool(DoInit))
}

func (m *TBaseVirtualTree) SortTree(Column TColumnIndex, Direction TSortDirection, DoInit bool) {
	baseVirtualTreeImportAPI().SysCallN(166, m.Instance(), uintptr(Column), uintptr(Direction), PascalBool(DoInit))
}

func (m *TBaseVirtualTree) ToggleNode(Node IVirtualNode) {
	baseVirtualTreeImportAPI().SysCallN(167, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) UpdateHorizontalRange() {
	baseVirtualTreeImportAPI().SysCallN(172, m.Instance())
}

func (m *TBaseVirtualTree) UpdateHorizontalScrollBar(DoRepaint bool) {
	baseVirtualTreeImportAPI().SysCallN(173, m.Instance(), PascalBool(DoRepaint))
}

func (m *TBaseVirtualTree) UpdateRanges() {
	baseVirtualTreeImportAPI().SysCallN(174, m.Instance())
}

func (m *TBaseVirtualTree) UpdateScrollBars(DoRepaint bool) {
	baseVirtualTreeImportAPI().SysCallN(175, m.Instance(), PascalBool(DoRepaint))
}

func (m *TBaseVirtualTree) UpdateVerticalRange() {
	baseVirtualTreeImportAPI().SysCallN(176, m.Instance())
}

func (m *TBaseVirtualTree) UpdateVerticalScrollBar(DoRepaint bool) {
	baseVirtualTreeImportAPI().SysCallN(177, m.Instance(), PascalBool(DoRepaint))
}

func (m *TBaseVirtualTree) ValidateChildren(Node IVirtualNode, Recursive bool) {
	baseVirtualTreeImportAPI().SysCallN(178, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}

func (m *TBaseVirtualTree) ValidateNode(Node IVirtualNode, Recursive bool) {
	baseVirtualTreeImportAPI().SysCallN(179, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}

var (
	baseVirtualTreeImport       *imports.Imports = nil
	baseVirtualTreeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("BaseVirtualTree_AbsoluteIndex", 0),
		/*1*/ imports.NewTable("BaseVirtualTree_AddChild", 0),
		/*2*/ imports.NewTable("BaseVirtualTree_AddFromStream", 0),
		/*3*/ imports.NewTable("BaseVirtualTree_BeginSynch", 0),
		/*4*/ imports.NewTable("BaseVirtualTree_BeginUpdate", 0),
		/*5*/ imports.NewTable("BaseVirtualTree_BottomNode", 0),
		/*6*/ imports.NewTable("BaseVirtualTree_CanEdit", 0),
		/*7*/ imports.NewTable("BaseVirtualTree_CancelCutOrCopy", 0),
		/*8*/ imports.NewTable("BaseVirtualTree_CancelEditNode", 0),
		/*9*/ imports.NewTable("BaseVirtualTree_CancelOperation", 0),
		/*10*/ imports.NewTable("BaseVirtualTree_CheckImages", 0),
		/*11*/ imports.NewTable("BaseVirtualTree_CheckState", 0),
		/*12*/ imports.NewTable("BaseVirtualTree_CheckType", 0),
		/*13*/ imports.NewTable("BaseVirtualTree_CheckedCount", 0),
		/*14*/ imports.NewTable("BaseVirtualTree_CheckedNodes", 0),
		/*15*/ imports.NewTable("BaseVirtualTree_ChildCount", 0),
		/*16*/ imports.NewTable("BaseVirtualTree_ChildNodes", 0),
		/*17*/ imports.NewTable("BaseVirtualTree_ChildrenInitialized", 0),
		/*18*/ imports.NewTable("BaseVirtualTree_Class", 0),
		/*19*/ imports.NewTable("BaseVirtualTree_Clear", 0),
		/*20*/ imports.NewTable("BaseVirtualTree_ClearChecked", 0),
		/*21*/ imports.NewTable("BaseVirtualTree_ClearSelection", 0),
		/*22*/ imports.NewTable("BaseVirtualTree_CopyTo", 0),
		/*23*/ imports.NewTable("BaseVirtualTree_CopyTo1", 0),
		/*24*/ imports.NewTable("BaseVirtualTree_CopyToClipboard", 0),
		/*25*/ imports.NewTable("BaseVirtualTree_Create", 0),
		/*26*/ imports.NewTable("BaseVirtualTree_CutCopyCount", 0),
		/*27*/ imports.NewTable("BaseVirtualTree_CutCopyNodes", 0),
		/*28*/ imports.NewTable("BaseVirtualTree_CutToClipboard", 0),
		/*29*/ imports.NewTable("BaseVirtualTree_DeleteChildren", 0),
		/*30*/ imports.NewTable("BaseVirtualTree_DeleteNode", 0),
		/*31*/ imports.NewTable("BaseVirtualTree_DeleteSelectedNodes", 0),
		/*32*/ imports.NewTable("BaseVirtualTree_DragImage", 0),
		/*33*/ imports.NewTable("BaseVirtualTree_DropTargetNode", 0),
		/*34*/ imports.NewTable("BaseVirtualTree_EditNode", 0),
		/*35*/ imports.NewTable("BaseVirtualTree_EmptyListMessage", 0),
		/*36*/ imports.NewTable("BaseVirtualTree_EndEditNode", 0),
		/*37*/ imports.NewTable("BaseVirtualTree_EndSynch", 0),
		/*38*/ imports.NewTable("BaseVirtualTree_EndUpdate", 0),
		/*39*/ imports.NewTable("BaseVirtualTree_EnsureNodeSelected", 0),
		/*40*/ imports.NewTable("BaseVirtualTree_Expanded", 0),
		/*41*/ imports.NewTable("BaseVirtualTree_FinishCutOrCopy", 0),
		/*42*/ imports.NewTable("BaseVirtualTree_FlushClipboard", 0),
		/*43*/ imports.NewTable("BaseVirtualTree_FocusedColumn", 0),
		/*44*/ imports.NewTable("BaseVirtualTree_FocusedNode", 0),
		/*45*/ imports.NewTable("BaseVirtualTree_FullCollapse", 0),
		/*46*/ imports.NewTable("BaseVirtualTree_FullExpand", 0),
		/*47*/ imports.NewTable("BaseVirtualTree_FullyVisible", 0),
		/*48*/ imports.NewTable("BaseVirtualTree_GetDisplayRect", 0),
		/*49*/ imports.NewTable("BaseVirtualTree_GetEffectivelyFiltered", 0),
		/*50*/ imports.NewTable("BaseVirtualTree_GetEffectivelyVisible", 0),
		/*51*/ imports.NewTable("BaseVirtualTree_GetFirst", 0),
		/*52*/ imports.NewTable("BaseVirtualTree_GetFirstChecked", 0),
		/*53*/ imports.NewTable("BaseVirtualTree_GetFirstChild", 0),
		/*54*/ imports.NewTable("BaseVirtualTree_GetFirstChildNoInit", 0),
		/*55*/ imports.NewTable("BaseVirtualTree_GetFirstCutCopy", 0),
		/*56*/ imports.NewTable("BaseVirtualTree_GetFirstInitialized", 0),
		/*57*/ imports.NewTable("BaseVirtualTree_GetFirstLeaf", 0),
		/*58*/ imports.NewTable("BaseVirtualTree_GetFirstLevel", 0),
		/*59*/ imports.NewTable("BaseVirtualTree_GetFirstNoInit", 0),
		/*60*/ imports.NewTable("BaseVirtualTree_GetFirstSelected", 0),
		/*61*/ imports.NewTable("BaseVirtualTree_GetFirstVisible", 0),
		/*62*/ imports.NewTable("BaseVirtualTree_GetFirstVisibleChild", 0),
		/*63*/ imports.NewTable("BaseVirtualTree_GetFirstVisibleChildNoInit", 0),
		/*64*/ imports.NewTable("BaseVirtualTree_GetFirstVisibleNoInit", 0),
		/*65*/ imports.NewTable("BaseVirtualTree_GetLast", 0),
		/*66*/ imports.NewTable("BaseVirtualTree_GetLastChild", 0),
		/*67*/ imports.NewTable("BaseVirtualTree_GetLastChildNoInit", 0),
		/*68*/ imports.NewTable("BaseVirtualTree_GetLastInitialized", 0),
		/*69*/ imports.NewTable("BaseVirtualTree_GetLastNoInit", 0),
		/*70*/ imports.NewTable("BaseVirtualTree_GetLastVisible", 0),
		/*71*/ imports.NewTable("BaseVirtualTree_GetLastVisibleChild", 0),
		/*72*/ imports.NewTable("BaseVirtualTree_GetLastVisibleChildNoInit", 0),
		/*73*/ imports.NewTable("BaseVirtualTree_GetLastVisibleNoInit", 0),
		/*74*/ imports.NewTable("BaseVirtualTree_GetMaxColumnWidth", 0),
		/*75*/ imports.NewTable("BaseVirtualTree_GetNext", 0),
		/*76*/ imports.NewTable("BaseVirtualTree_GetNextChecked", 0),
		/*77*/ imports.NewTable("BaseVirtualTree_GetNextChecked1", 0),
		/*78*/ imports.NewTable("BaseVirtualTree_GetNextCutCopy", 0),
		/*79*/ imports.NewTable("BaseVirtualTree_GetNextInitialized", 0),
		/*80*/ imports.NewTable("BaseVirtualTree_GetNextLeaf", 0),
		/*81*/ imports.NewTable("BaseVirtualTree_GetNextLevel", 0),
		/*82*/ imports.NewTable("BaseVirtualTree_GetNextNoInit", 0),
		/*83*/ imports.NewTable("BaseVirtualTree_GetNextSelected", 0),
		/*84*/ imports.NewTable("BaseVirtualTree_GetNextSibling", 0),
		/*85*/ imports.NewTable("BaseVirtualTree_GetNextSiblingNoInit", 0),
		/*86*/ imports.NewTable("BaseVirtualTree_GetNextVisible", 0),
		/*87*/ imports.NewTable("BaseVirtualTree_GetNextVisibleNoInit", 0),
		/*88*/ imports.NewTable("BaseVirtualTree_GetNextVisibleSibling", 0),
		/*89*/ imports.NewTable("BaseVirtualTree_GetNextVisibleSiblingNoInit", 0),
		/*90*/ imports.NewTable("BaseVirtualTree_GetNodeAt", 0),
		/*91*/ imports.NewTable("BaseVirtualTree_GetNodeAt1", 0),
		/*92*/ imports.NewTable("BaseVirtualTree_GetNodeAt2", 0),
		/*93*/ imports.NewTable("BaseVirtualTree_GetNodeData", 0),
		/*94*/ imports.NewTable("BaseVirtualTree_GetNodeLevel", 0),
		/*95*/ imports.NewTable("BaseVirtualTree_GetPrevious", 0),
		/*96*/ imports.NewTable("BaseVirtualTree_GetPreviousChecked", 0),
		/*97*/ imports.NewTable("BaseVirtualTree_GetPreviousCutCopy", 0),
		/*98*/ imports.NewTable("BaseVirtualTree_GetPreviousInitialized", 0),
		/*99*/ imports.NewTable("BaseVirtualTree_GetPreviousLeaf", 0),
		/*100*/ imports.NewTable("BaseVirtualTree_GetPreviousLevel", 0),
		/*101*/ imports.NewTable("BaseVirtualTree_GetPreviousNoInit", 0),
		/*102*/ imports.NewTable("BaseVirtualTree_GetPreviousSelected", 0),
		/*103*/ imports.NewTable("BaseVirtualTree_GetPreviousSibling", 0),
		/*104*/ imports.NewTable("BaseVirtualTree_GetPreviousSiblingNoInit", 0),
		/*105*/ imports.NewTable("BaseVirtualTree_GetPreviousVisible", 0),
		/*106*/ imports.NewTable("BaseVirtualTree_GetPreviousVisibleNoInit", 0),
		/*107*/ imports.NewTable("BaseVirtualTree_GetPreviousVisibleSibling", 0),
		/*108*/ imports.NewTable("BaseVirtualTree_GetPreviousVisibleSiblingNoInit", 0),
		/*109*/ imports.NewTable("BaseVirtualTree_GetSortedCutCopySet", 0),
		/*110*/ imports.NewTable("BaseVirtualTree_GetSortedSelection", 0),
		/*111*/ imports.NewTable("BaseVirtualTree_GetTextInfo", 0),
		/*112*/ imports.NewTable("BaseVirtualTree_GetTreeRect", 0),
		/*113*/ imports.NewTable("BaseVirtualTree_GetVisibleParent", 0),
		/*114*/ imports.NewTable("BaseVirtualTree_HasAsParent", 0),
		/*115*/ imports.NewTable("BaseVirtualTree_HasChildren", 0),
		/*116*/ imports.NewTable("BaseVirtualTree_HotNode", 0),
		/*117*/ imports.NewTable("BaseVirtualTree_InitializedNodes", 0),
		/*118*/ imports.NewTable("BaseVirtualTree_InsertNode", 0),
		/*119*/ imports.NewTable("BaseVirtualTree_InvalidateChildren", 0),
		/*120*/ imports.NewTable("BaseVirtualTree_InvalidateColumn", 0),
		/*121*/ imports.NewTable("BaseVirtualTree_InvalidateNode", 0),
		/*122*/ imports.NewTable("BaseVirtualTree_InvalidateToBottom", 0),
		/*123*/ imports.NewTable("BaseVirtualTree_InvertSelection", 0),
		/*124*/ imports.NewTable("BaseVirtualTree_IsDisabled", 0),
		/*125*/ imports.NewTable("BaseVirtualTree_IsEditing", 0),
		/*126*/ imports.NewTable("BaseVirtualTree_IsEffectivelyFiltered", 0),
		/*127*/ imports.NewTable("BaseVirtualTree_IsEffectivelyVisible", 0),
		/*128*/ imports.NewTable("BaseVirtualTree_IsEmpty", 0),
		/*129*/ imports.NewTable("BaseVirtualTree_IsFiltered", 0),
		/*130*/ imports.NewTable("BaseVirtualTree_IsMouseSelecting", 0),
		/*131*/ imports.NewTable("BaseVirtualTree_LeafNodes", 0),
		/*132*/ imports.NewTable("BaseVirtualTree_LevelNodes", 0),
		/*133*/ imports.NewTable("BaseVirtualTree_LoadFromFile", 0),
		/*134*/ imports.NewTable("BaseVirtualTree_LoadFromStream", 0),
		/*135*/ imports.NewTable("BaseVirtualTree_MeasureItemHeight", 0),
		/*136*/ imports.NewTable("BaseVirtualTree_MoveTo", 0),
		/*137*/ imports.NewTable("BaseVirtualTree_MoveTo1", 0),
		/*138*/ imports.NewTable("BaseVirtualTree_MultiLine", 0),
		/*139*/ imports.NewTable("BaseVirtualTree_NoInitNodes", 0),
		/*140*/ imports.NewTable("BaseVirtualTree_NodeHeight", 0),
		/*141*/ imports.NewTable("BaseVirtualTree_NodeIsVisible", 0),
		/*142*/ imports.NewTable("BaseVirtualTree_NodeParent", 0),
		/*143*/ imports.NewTable("BaseVirtualTree_Nodes", 0),
		/*144*/ imports.NewTable("BaseVirtualTree_OffsetX", 0),
		/*145*/ imports.NewTable("BaseVirtualTree_OffsetXY", 0),
		/*146*/ imports.NewTable("BaseVirtualTree_OffsetY", 0),
		/*147*/ imports.NewTable("BaseVirtualTree_OperationCount", 0),
		/*148*/ imports.NewTable("BaseVirtualTree_PaintTree", 0),
		/*149*/ imports.NewTable("BaseVirtualTree_PasteFromClipboard", 0),
		/*150*/ imports.NewTable("BaseVirtualTree_ReinitChildren", 0),
		/*151*/ imports.NewTable("BaseVirtualTree_ReinitNode", 0),
		/*152*/ imports.NewTable("BaseVirtualTree_RepaintNode", 0),
		/*153*/ imports.NewTable("BaseVirtualTree_ResetNode", 0),
		/*154*/ imports.NewTable("BaseVirtualTree_RootNode", 0),
		/*155*/ imports.NewTable("BaseVirtualTree_SaveToFile", 0),
		/*156*/ imports.NewTable("BaseVirtualTree_SaveToStream", 0),
		/*157*/ imports.NewTable("BaseVirtualTree_ScrollIntoView", 0),
		/*158*/ imports.NewTable("BaseVirtualTree_ScrollIntoView1", 0),
		/*159*/ imports.NewTable("BaseVirtualTree_SearchBuffer", 0),
		/*160*/ imports.NewTable("BaseVirtualTree_SelectAll", 0),
		/*161*/ imports.NewTable("BaseVirtualTree_Selected", 0),
		/*162*/ imports.NewTable("BaseVirtualTree_SelectedCount", 0),
		/*163*/ imports.NewTable("BaseVirtualTree_SelectedNodes", 0),
		/*164*/ imports.NewTable("BaseVirtualTree_SelectionLocked", 0),
		/*165*/ imports.NewTable("BaseVirtualTree_Sort", 0),
		/*166*/ imports.NewTable("BaseVirtualTree_SortTree", 0),
		/*167*/ imports.NewTable("BaseVirtualTree_ToggleNode", 0),
		/*168*/ imports.NewTable("BaseVirtualTree_TopNode", 0),
		/*169*/ imports.NewTable("BaseVirtualTree_TotalCount", 0),
		/*170*/ imports.NewTable("BaseVirtualTree_TreeStates", 0),
		/*171*/ imports.NewTable("BaseVirtualTree_UpdateCount", 0),
		/*172*/ imports.NewTable("BaseVirtualTree_UpdateHorizontalRange", 0),
		/*173*/ imports.NewTable("BaseVirtualTree_UpdateHorizontalScrollBar", 0),
		/*174*/ imports.NewTable("BaseVirtualTree_UpdateRanges", 0),
		/*175*/ imports.NewTable("BaseVirtualTree_UpdateScrollBars", 0),
		/*176*/ imports.NewTable("BaseVirtualTree_UpdateVerticalRange", 0),
		/*177*/ imports.NewTable("BaseVirtualTree_UpdateVerticalScrollBar", 0),
		/*178*/ imports.NewTable("BaseVirtualTree_ValidateChildren", 0),
		/*179*/ imports.NewTable("BaseVirtualTree_ValidateNode", 0),
		/*180*/ imports.NewTable("BaseVirtualTree_VerticalAlignment", 0),
		/*181*/ imports.NewTable("BaseVirtualTree_VisibleChildNoInitNodes", 0),
		/*182*/ imports.NewTable("BaseVirtualTree_VisibleChildNodes", 0),
		/*183*/ imports.NewTable("BaseVirtualTree_VisibleCount", 0),
		/*184*/ imports.NewTable("BaseVirtualTree_VisibleNoInitNodes", 0),
		/*185*/ imports.NewTable("BaseVirtualTree_VisibleNodes", 0),
		/*186*/ imports.NewTable("BaseVirtualTree_VisiblePath", 0),
	}
)

func baseVirtualTreeImportAPI() *imports.Imports {
	if baseVirtualTreeImport == nil {
		baseVirtualTreeImport = NewDefaultImports()
		baseVirtualTreeImport.SetImportTable(baseVirtualTreeImportTables)
		baseVirtualTreeImportTables = nil
	}
	return baseVirtualTreeImport
}
