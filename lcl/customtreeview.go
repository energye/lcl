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

// ICustomTreeView Parent: ICustomControl
type ICustomTreeView interface {
	ICustomControl
	AlphaSort() bool                                                                                                      // function
	DefaultTreeViewSort(node1 ITreeNode, node2 ITreeNode) int32                                                           // function
	GetHitTestInfoAt(X int32, Y int32) types.THitTests                                                                    // function
	GetNodeAt(X int32, Y int32) ITreeNode                                                                                 // function
	GetNodeWithExpandSignAt(X int32, Y int32) ITreeNode                                                                   // function
	IsEditing() bool                                                                                                      // function
	GetFirstMultiSelected() ITreeNode                                                                                     // function
	GetLastMultiSelected() ITreeNode                                                                                      // function
	SelectionVisible() bool                                                                                               // function
	StoreCurrentSelection() IStringList                                                                                   // function
	ClearSelection(keepPrimary bool)                                                                                      // procedure
	ConsistencyCheck()                                                                                                    // procedure
	GetInsertMarkAt(X int32, Y int32, outAnInsertMarkNode *ITreeNode, outAnInsertMarkType *types.TTreeViewInsertMarkType) // procedure
	SetInsertMark(anInsertMarkNode ITreeNode, anInsertMarkType types.TTreeViewInsertMarkType)                             // procedure
	SetInsertMarkAt(X int32, Y int32)                                                                                     // procedure
	BeginUpdate()                                                                                                         // procedure
	EndUpdate()                                                                                                           // procedure
	FullCollapse()                                                                                                        // procedure
	FullExpand()                                                                                                          // procedure
	LoadFromFile(fileName string)                                                                                         // procedure
	LoadFromStream(stream IStream)                                                                                        // procedure
	SaveToFile(fileName string)                                                                                           // procedure
	SaveToStream(stream IStream)                                                                                          // procedure
	WriteDebugReport(prefix string, allNodes bool)                                                                        // procedure
	LockSelectionChangeEvent()                                                                                            // procedure
	UnlockSelectionChangeEvent()                                                                                          // procedure
	SelectWithTreeNodeShiftState(node ITreeNode, shiftState types.TShiftState)                                            // procedure
	SelectWithList(nodes IList)                                                                                           // procedure
	MakeSelectionVisible()                                                                                                // procedure
	ClearInvisibleSelection()                                                                                             // procedure
	ApplyStoredSelection(selection IStringList, freeList bool)                                                            // procedure
	MoveToNextNode(select_ bool)                                                                                          // procedure
	MoveToPrevNode(select_ bool)                                                                                          // procedure
	MovePageDown(select_ bool)                                                                                            // procedure
	MovePageUp(select_ bool)                                                                                              // procedure
	MoveLeft(select_ bool)                                                                                                // procedure
	MoveRight(select_ bool)                                                                                               // procedure
	MoveExpand(select_ bool)                                                                                              // procedure
	MoveCollapse(select_ bool)                                                                                            // procedure
	MoveHome(select_ bool)                                                                                                // procedure
	MoveEnd(select_ bool)                                                                                                 // procedure
	// AccessibilityOn
	//  This property is provided for the case when a tree view contains a huge amount of items,
	//  lets say 10.000+. In this case accessibility might slow the tree down, so turning
	//  it off might make things faster
	AccessibilityOn() bool                                 // property AccessibilityOn Getter
	SetAccessibilityOn(value bool)                         // property AccessibilityOn Setter
	BackgroundColor() types.TColor                         // property BackgroundColor Getter
	SetBackgroundColor(value types.TColor)                 // property BackgroundColor Setter
	BottomItem() ITreeNode                                 // property BottomItem Getter
	SetBottomItem(value ITreeNode)                         // property BottomItem Setter
	DefaultItemHeight() int32                              // property DefaultItemHeight Getter
	SetDefaultItemHeight(value int32)                      // property DefaultItemHeight Setter
	DropTarget() ITreeNode                                 // property DropTarget Getter
	SetDropTarget(value ITreeNode)                         // property DropTarget Setter
	ExpandSignColor() types.TColor                         // property ExpandSignColor Getter
	SetExpandSignColor(value types.TColor)                 // property ExpandSignColor Setter
	ExpandSignSize() int32                                 // property ExpandSignSize Getter
	SetExpandSignSize(value int32)                         // property ExpandSignSize Setter
	ExpandSignWidth() int32                                // property ExpandSignWidth Getter
	SetExpandSignWidth(value int32)                        // property ExpandSignWidth Setter
	ExpandSignType() types.TTreeViewExpandSignType         // property ExpandSignType Getter
	SetExpandSignType(value types.TTreeViewExpandSignType) // property ExpandSignType Setter
	Images() ICustomImageList                              // property Images Getter
	SetImages(value ICustomImageList)                      // property Images Setter
	ImagesWidth() int32                                    // property ImagesWidth Getter
	SetImagesWidth(value int32)                            // property ImagesWidth Setter
	InsertMarkNode() ITreeNode                             // property InsertMarkNode Getter
	SetInsertMarkNode(value ITreeNode)                     // property InsertMarkNode Setter
	InsertMarkType() types.TTreeViewInsertMarkType         // property InsertMarkType Getter
	SetInsertMarkType(value types.TTreeViewInsertMarkType) // property InsertMarkType Setter
	Items() ITreeNodes                                     // property Items Getter
	SetItems(value ITreeNodes)                             // property Items Setter
	KeepCollapsedNodes() bool                              // property KeepCollapsedNodes Getter
	SetKeepCollapsedNodes(value bool)                      // property KeepCollapsedNodes Setter
	MultiSelectStyle() types.TMultiSelectStyle             // property MultiSelectStyle Getter
	SetMultiSelectStyle(value types.TMultiSelectStyle)     // property MultiSelectStyle Setter
	Options() types.TTreeViewOptions                       // property Options Getter
	SetOptions(value types.TTreeViewOptions)               // property Options Setter
	ScrollBars() types.TScrollStyle                        // property ScrollBars Getter
	SetScrollBars(value types.TScrollStyle)                // property ScrollBars Setter
	Selected() ITreeNode                                   // property Selected Getter
	SetSelected(value ITreeNode)                           // property Selected Setter
	SelectionColor() types.TColor                          // property SelectionColor Getter
	SetSelectionColor(value types.TColor)                  // property SelectionColor Setter
	SelectionCount() uint32                                // property SelectionCount Getter
	SelectionFontColor() types.TColor                      // property SelectionFontColor Getter
	SetSelectionFontColor(value types.TColor)              // property SelectionFontColor Setter
	SelectionFontColorUsed() bool                          // property SelectionFontColorUsed Getter
	SetSelectionFontColorUsed(value bool)                  // property SelectionFontColorUsed Setter
	Selections(index int32) ITreeNode                      // property Selections Getter
	SeparatorColor() types.TColor                          // property SeparatorColor Getter
	SetSeparatorColor(value types.TColor)                  // property SeparatorColor Setter
	StateImages() ICustomImageList                         // property StateImages Getter
	SetStateImages(value ICustomImageList)                 // property StateImages Setter
	StateImagesWidth() int32                               // property StateImagesWidth Getter
	SetStateImagesWidth(value int32)                       // property StateImagesWidth Setter
	TopItem() ITreeNode                                    // property TopItem Getter
	SetTopItem(value ITreeNode)                            // property TopItem Setter
	TreeLineColor() types.TColor                           // property TreeLineColor Getter
	SetTreeLineColor(value types.TColor)                   // property TreeLineColor Setter
	TreeLinePenStyle() types.TPenStyle                     // property TreeLinePenStyle Getter
	SetTreeLinePenStyle(value types.TPenStyle)             // property TreeLinePenStyle Setter
}

type TCustomTreeView struct {
	TCustomControl
}

func (m *TCustomTreeView) AlphaSort() bool {
	if !m.IsValid() {
		return false
	}
	r := customTreeViewAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTreeView) DefaultTreeViewSort(node1 ITreeNode, node2 ITreeNode) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(node1), base.GetObjectUintptr(node2))
	return int32(r)
}

func (m *TCustomTreeView) GetHitTestInfoAt(X int32, Y int32) types.THitTests {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(3, m.Instance(), uintptr(X), uintptr(Y))
	return types.THitTests(r)
}

func (m *TCustomTreeView) GetNodeAt(X int32, Y int32) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(4, m.Instance(), uintptr(X), uintptr(Y))
	return AsTreeNode(r)
}

func (m *TCustomTreeView) GetNodeWithExpandSignAt(X int32, Y int32) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(5, m.Instance(), uintptr(X), uintptr(Y))
	return AsTreeNode(r)
}

func (m *TCustomTreeView) IsEditing() bool {
	if !m.IsValid() {
		return false
	}
	r := customTreeViewAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTreeView) GetFirstMultiSelected() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(7, m.Instance())
	return AsTreeNode(r)
}

func (m *TCustomTreeView) GetLastMultiSelected() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(8, m.Instance())
	return AsTreeNode(r)
}

func (m *TCustomTreeView) SelectionVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := customTreeViewAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTreeView) StoreCurrentSelection() IStringList {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(10, m.Instance())
	return AsStringList(r)
}

func (m *TCustomTreeView) ClearSelection(keepPrimary bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(11, m.Instance(), api.PasBool(keepPrimary))
}

func (m *TCustomTreeView) ConsistencyCheck() {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(12, m.Instance())
}

func (m *TCustomTreeView) GetInsertMarkAt(X int32, Y int32, outAnInsertMarkNode *ITreeNode, outAnInsertMarkType *types.TTreeViewInsertMarkType) {
	if !m.IsValid() {
		return
	}
	var anInsertMarkNodePtr uintptr
	var anInsertMarkTypePtr uintptr
	customTreeViewAPI().SysCallN(13, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(&anInsertMarkNodePtr)), uintptr(base.UnsafePointer(&anInsertMarkTypePtr)))
	*outAnInsertMarkNode = AsTreeNode(anInsertMarkNodePtr)
	*outAnInsertMarkType = types.TTreeViewInsertMarkType(anInsertMarkTypePtr)
}

func (m *TCustomTreeView) SetInsertMark(anInsertMarkNode ITreeNode, anInsertMarkType types.TTreeViewInsertMarkType) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(anInsertMarkNode), uintptr(anInsertMarkType))
}

func (m *TCustomTreeView) SetInsertMarkAt(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(15, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TCustomTreeView) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(16, m.Instance())
}

func (m *TCustomTreeView) EndUpdate() {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(17, m.Instance())
}

func (m *TCustomTreeView) FullCollapse() {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(18, m.Instance())
}

func (m *TCustomTreeView) FullExpand() {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(19, m.Instance())
}

func (m *TCustomTreeView) LoadFromFile(fileName string) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(20, m.Instance(), api.PasStr(fileName))
}

func (m *TCustomTreeView) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCustomTreeView) SaveToFile(fileName string) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(22, m.Instance(), api.PasStr(fileName))
}

func (m *TCustomTreeView) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(23, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCustomTreeView) WriteDebugReport(prefix string, allNodes bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(24, m.Instance(), api.PasStr(prefix), api.PasBool(allNodes))
}

func (m *TCustomTreeView) LockSelectionChangeEvent() {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(25, m.Instance())
}

func (m *TCustomTreeView) UnlockSelectionChangeEvent() {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(26, m.Instance())
}

func (m *TCustomTreeView) SelectWithTreeNodeShiftState(node ITreeNode, shiftState types.TShiftState) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(27, m.Instance(), base.GetObjectUintptr(node), uintptr(shiftState))
}

func (m *TCustomTreeView) SelectWithList(nodes IList) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(28, m.Instance(), base.GetObjectUintptr(nodes))
}

func (m *TCustomTreeView) MakeSelectionVisible() {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(29, m.Instance())
}

func (m *TCustomTreeView) ClearInvisibleSelection() {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(30, m.Instance())
}

func (m *TCustomTreeView) ApplyStoredSelection(selection IStringList, freeList bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(31, m.Instance(), base.GetObjectUintptr(selection), api.PasBool(freeList))
}

func (m *TCustomTreeView) MoveToNextNode(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(32, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) MoveToPrevNode(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(33, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) MovePageDown(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(34, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) MovePageUp(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(35, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) MoveLeft(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(36, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) MoveRight(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(37, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) MoveExpand(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(38, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) MoveCollapse(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(39, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) MoveHome(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(40, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) MoveEnd(select_ bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(41, m.Instance(), api.PasBool(select_))
}

func (m *TCustomTreeView) AccessibilityOn() bool {
	if !m.IsValid() {
		return false
	}
	r := customTreeViewAPI().SysCallN(42, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTreeView) SetAccessibilityOn(value bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(42, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTreeView) BackgroundColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(43, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomTreeView) SetBackgroundColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(43, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) BottomItem() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(44, 0, m.Instance())
	return AsTreeNode(r)
}

func (m *TCustomTreeView) SetBottomItem(value ITreeNode) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(44, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTreeView) DefaultItemHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(45, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTreeView) SetDefaultItemHeight(value int32) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(45, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) DropTarget() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(46, 0, m.Instance())
	return AsTreeNode(r)
}

func (m *TCustomTreeView) SetDropTarget(value ITreeNode) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(46, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTreeView) ExpandSignColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(47, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomTreeView) SetExpandSignColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(47, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) ExpandSignSize() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(48, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTreeView) SetExpandSignSize(value int32) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(48, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) ExpandSignWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(49, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTreeView) SetExpandSignWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(49, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) ExpandSignType() types.TTreeViewExpandSignType {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(50, 0, m.Instance())
	return types.TTreeViewExpandSignType(r)
}

func (m *TCustomTreeView) SetExpandSignType(value types.TTreeViewExpandSignType) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(50, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(51, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomTreeView) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(51, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTreeView) ImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(52, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTreeView) SetImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(52, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) InsertMarkNode() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(53, 0, m.Instance())
	return AsTreeNode(r)
}

func (m *TCustomTreeView) SetInsertMarkNode(value ITreeNode) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(53, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTreeView) InsertMarkType() types.TTreeViewInsertMarkType {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(54, 0, m.Instance())
	return types.TTreeViewInsertMarkType(r)
}

func (m *TCustomTreeView) SetInsertMarkType(value types.TTreeViewInsertMarkType) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(54, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) Items() ITreeNodes {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(55, 0, m.Instance())
	return AsTreeNodes(r)
}

func (m *TCustomTreeView) SetItems(value ITreeNodes) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(55, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTreeView) KeepCollapsedNodes() bool {
	if !m.IsValid() {
		return false
	}
	r := customTreeViewAPI().SysCallN(56, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTreeView) SetKeepCollapsedNodes(value bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(56, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTreeView) MultiSelectStyle() types.TMultiSelectStyle {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(57, 0, m.Instance())
	return types.TMultiSelectStyle(r)
}

func (m *TCustomTreeView) SetMultiSelectStyle(value types.TMultiSelectStyle) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(57, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) Options() types.TTreeViewOptions {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(58, 0, m.Instance())
	return types.TTreeViewOptions(r)
}

func (m *TCustomTreeView) SetOptions(value types.TTreeViewOptions) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(58, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) ScrollBars() types.TScrollStyle {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(59, 0, m.Instance())
	return types.TScrollStyle(r)
}

func (m *TCustomTreeView) SetScrollBars(value types.TScrollStyle) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(59, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) Selected() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(60, 0, m.Instance())
	return AsTreeNode(r)
}

func (m *TCustomTreeView) SetSelected(value ITreeNode) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(60, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTreeView) SelectionColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(61, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomTreeView) SetSelectionColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(61, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) SelectionCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(62, m.Instance())
	return uint32(r)
}

func (m *TCustomTreeView) SelectionFontColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(63, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomTreeView) SetSelectionFontColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(63, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) SelectionFontColorUsed() bool {
	if !m.IsValid() {
		return false
	}
	r := customTreeViewAPI().SysCallN(64, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTreeView) SetSelectionFontColorUsed(value bool) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(64, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTreeView) Selections(index int32) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(65, m.Instance(), uintptr(index))
	return AsTreeNode(r)
}

func (m *TCustomTreeView) SeparatorColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(66, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomTreeView) SetSeparatorColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(66, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) StateImages() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(67, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomTreeView) SetStateImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(67, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTreeView) StateImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(68, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTreeView) SetStateImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(68, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) TopItem() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := customTreeViewAPI().SysCallN(69, 0, m.Instance())
	return AsTreeNode(r)
}

func (m *TCustomTreeView) SetTopItem(value ITreeNode) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(69, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTreeView) TreeLineColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(70, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomTreeView) SetTreeLineColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(70, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTreeView) TreeLinePenStyle() types.TPenStyle {
	if !m.IsValid() {
		return 0
	}
	r := customTreeViewAPI().SysCallN(71, 0, m.Instance())
	return types.TPenStyle(r)
}

func (m *TCustomTreeView) SetTreeLinePenStyle(value types.TPenStyle) {
	if !m.IsValid() {
		return
	}
	customTreeViewAPI().SysCallN(71, 1, m.Instance(), uintptr(value))
}

// NewCustomTreeView class constructor
func NewCustomTreeView(anOwner IComponent) ICustomTreeView {
	r := customTreeViewAPI().SysCallN(0, base.GetObjectUintptr(anOwner))
	return AsCustomTreeView(r)
}

func TCustomTreeViewClass() types.TClass {
	r := customTreeViewAPI().SysCallN(72)
	return types.TClass(r)
}

var (
	customTreeViewOnce   base.Once
	customTreeViewImport *imports.Imports = nil
)

func customTreeViewAPI() *imports.Imports {
	customTreeViewOnce.Do(func() {
		customTreeViewImport = api.NewDefaultImports()
		customTreeViewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomTreeView_Create", 0), // constructor NewCustomTreeView
			/* 1 */ imports.NewTable("TCustomTreeView_AlphaSort", 0), // function AlphaSort
			/* 2 */ imports.NewTable("TCustomTreeView_DefaultTreeViewSort", 0), // function DefaultTreeViewSort
			/* 3 */ imports.NewTable("TCustomTreeView_GetHitTestInfoAt", 0), // function GetHitTestInfoAt
			/* 4 */ imports.NewTable("TCustomTreeView_GetNodeAt", 0), // function GetNodeAt
			/* 5 */ imports.NewTable("TCustomTreeView_GetNodeWithExpandSignAt", 0), // function GetNodeWithExpandSignAt
			/* 6 */ imports.NewTable("TCustomTreeView_IsEditing", 0), // function IsEditing
			/* 7 */ imports.NewTable("TCustomTreeView_GetFirstMultiSelected", 0), // function GetFirstMultiSelected
			/* 8 */ imports.NewTable("TCustomTreeView_GetLastMultiSelected", 0), // function GetLastMultiSelected
			/* 9 */ imports.NewTable("TCustomTreeView_SelectionVisible", 0), // function SelectionVisible
			/* 10 */ imports.NewTable("TCustomTreeView_StoreCurrentSelection", 0), // function StoreCurrentSelection
			/* 11 */ imports.NewTable("TCustomTreeView_ClearSelection", 0), // procedure ClearSelection
			/* 12 */ imports.NewTable("TCustomTreeView_ConsistencyCheck", 0), // procedure ConsistencyCheck
			/* 13 */ imports.NewTable("TCustomTreeView_GetInsertMarkAt", 0), // procedure GetInsertMarkAt
			/* 14 */ imports.NewTable("TCustomTreeView_SetInsertMark", 0), // procedure SetInsertMark
			/* 15 */ imports.NewTable("TCustomTreeView_SetInsertMarkAt", 0), // procedure SetInsertMarkAt
			/* 16 */ imports.NewTable("TCustomTreeView_BeginUpdate", 0), // procedure BeginUpdate
			/* 17 */ imports.NewTable("TCustomTreeView_EndUpdate", 0), // procedure EndUpdate
			/* 18 */ imports.NewTable("TCustomTreeView_FullCollapse", 0), // procedure FullCollapse
			/* 19 */ imports.NewTable("TCustomTreeView_FullExpand", 0), // procedure FullExpand
			/* 20 */ imports.NewTable("TCustomTreeView_LoadFromFile", 0), // procedure LoadFromFile
			/* 21 */ imports.NewTable("TCustomTreeView_LoadFromStream", 0), // procedure LoadFromStream
			/* 22 */ imports.NewTable("TCustomTreeView_SaveToFile", 0), // procedure SaveToFile
			/* 23 */ imports.NewTable("TCustomTreeView_SaveToStream", 0), // procedure SaveToStream
			/* 24 */ imports.NewTable("TCustomTreeView_WriteDebugReport", 0), // procedure WriteDebugReport
			/* 25 */ imports.NewTable("TCustomTreeView_LockSelectionChangeEvent", 0), // procedure LockSelectionChangeEvent
			/* 26 */ imports.NewTable("TCustomTreeView_UnlockSelectionChangeEvent", 0), // procedure UnlockSelectionChangeEvent
			/* 27 */ imports.NewTable("TCustomTreeView_SelectWithTreeNodeShiftState", 0), // procedure SelectWithTreeNodeShiftState
			/* 28 */ imports.NewTable("TCustomTreeView_SelectWithList", 0), // procedure SelectWithList
			/* 29 */ imports.NewTable("TCustomTreeView_MakeSelectionVisible", 0), // procedure MakeSelectionVisible
			/* 30 */ imports.NewTable("TCustomTreeView_ClearInvisibleSelection", 0), // procedure ClearInvisibleSelection
			/* 31 */ imports.NewTable("TCustomTreeView_ApplyStoredSelection", 0), // procedure ApplyStoredSelection
			/* 32 */ imports.NewTable("TCustomTreeView_MoveToNextNode", 0), // procedure MoveToNextNode
			/* 33 */ imports.NewTable("TCustomTreeView_MoveToPrevNode", 0), // procedure MoveToPrevNode
			/* 34 */ imports.NewTable("TCustomTreeView_MovePageDown", 0), // procedure MovePageDown
			/* 35 */ imports.NewTable("TCustomTreeView_MovePageUp", 0), // procedure MovePageUp
			/* 36 */ imports.NewTable("TCustomTreeView_MoveLeft", 0), // procedure MoveLeft
			/* 37 */ imports.NewTable("TCustomTreeView_MoveRight", 0), // procedure MoveRight
			/* 38 */ imports.NewTable("TCustomTreeView_MoveExpand", 0), // procedure MoveExpand
			/* 39 */ imports.NewTable("TCustomTreeView_MoveCollapse", 0), // procedure MoveCollapse
			/* 40 */ imports.NewTable("TCustomTreeView_MoveHome", 0), // procedure MoveHome
			/* 41 */ imports.NewTable("TCustomTreeView_MoveEnd", 0), // procedure MoveEnd
			/* 42 */ imports.NewTable("TCustomTreeView_AccessibilityOn", 0), // property AccessibilityOn
			/* 43 */ imports.NewTable("TCustomTreeView_BackgroundColor", 0), // property BackgroundColor
			/* 44 */ imports.NewTable("TCustomTreeView_BottomItem", 0), // property BottomItem
			/* 45 */ imports.NewTable("TCustomTreeView_DefaultItemHeight", 0), // property DefaultItemHeight
			/* 46 */ imports.NewTable("TCustomTreeView_DropTarget", 0), // property DropTarget
			/* 47 */ imports.NewTable("TCustomTreeView_ExpandSignColor", 0), // property ExpandSignColor
			/* 48 */ imports.NewTable("TCustomTreeView_ExpandSignSize", 0), // property ExpandSignSize
			/* 49 */ imports.NewTable("TCustomTreeView_ExpandSignWidth", 0), // property ExpandSignWidth
			/* 50 */ imports.NewTable("TCustomTreeView_ExpandSignType", 0), // property ExpandSignType
			/* 51 */ imports.NewTable("TCustomTreeView_Images", 0), // property Images
			/* 52 */ imports.NewTable("TCustomTreeView_ImagesWidth", 0), // property ImagesWidth
			/* 53 */ imports.NewTable("TCustomTreeView_InsertMarkNode", 0), // property InsertMarkNode
			/* 54 */ imports.NewTable("TCustomTreeView_InsertMarkType", 0), // property InsertMarkType
			/* 55 */ imports.NewTable("TCustomTreeView_Items", 0), // property Items
			/* 56 */ imports.NewTable("TCustomTreeView_KeepCollapsedNodes", 0), // property KeepCollapsedNodes
			/* 57 */ imports.NewTable("TCustomTreeView_MultiSelectStyle", 0), // property MultiSelectStyle
			/* 58 */ imports.NewTable("TCustomTreeView_Options", 0), // property Options
			/* 59 */ imports.NewTable("TCustomTreeView_ScrollBars", 0), // property ScrollBars
			/* 60 */ imports.NewTable("TCustomTreeView_Selected", 0), // property Selected
			/* 61 */ imports.NewTable("TCustomTreeView_SelectionColor", 0), // property SelectionColor
			/* 62 */ imports.NewTable("TCustomTreeView_SelectionCount", 0), // property SelectionCount
			/* 63 */ imports.NewTable("TCustomTreeView_SelectionFontColor", 0), // property SelectionFontColor
			/* 64 */ imports.NewTable("TCustomTreeView_SelectionFontColorUsed", 0), // property SelectionFontColorUsed
			/* 65 */ imports.NewTable("TCustomTreeView_Selections", 0), // property Selections
			/* 66 */ imports.NewTable("TCustomTreeView_SeparatorColor", 0), // property SeparatorColor
			/* 67 */ imports.NewTable("TCustomTreeView_StateImages", 0), // property StateImages
			/* 68 */ imports.NewTable("TCustomTreeView_StateImagesWidth", 0), // property StateImagesWidth
			/* 69 */ imports.NewTable("TCustomTreeView_TopItem", 0), // property TopItem
			/* 70 */ imports.NewTable("TCustomTreeView_TreeLineColor", 0), // property TreeLineColor
			/* 71 */ imports.NewTable("TCustomTreeView_TreeLinePenStyle", 0), // property TreeLinePenStyle
			/* 72 */ imports.NewTable("TCustomTreeView_TClass", 0), // function TCustomTreeViewClass
		}
	})
	return customTreeViewImport
}
