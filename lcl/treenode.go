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

// ITreeNode Parent: IPersistent
type ITreeNode interface {
	IPersistent
	AlphaSort() bool                                               // function
	Bottom() int32                                                 // function
	BottomExpanded() int32                                         // function
	DefaultTreeViewSort(node1 ITreeNode, node2 ITreeNode) int32    // function
	DisplayExpandSignLeft() int32                                  // function
	DisplayExpandSignRect() types.TRect                            // function
	DisplayExpandSignRight() int32                                 // function
	DisplayIconLeft() int32                                        // function
	DisplayRect(textOnly bool) types.TRect                         // function
	DisplayStateIconLeft() int32                                   // function
	DisplayTextLeft() int32                                        // function
	DisplayTextRight() int32                                       // function
	EditText() bool                                                // function
	FindNode(nodeText string) ITreeNode                            // function
	GetFirstChild() ITreeNode                                      // function
	GetFirstSibling() ITreeNode                                    // function
	GetFirstVisibleChild(enabledOnly bool) ITreeNode               // function
	GetHandle() types.TLCLHandle                                   // function
	GetLastChild() ITreeNode                                       // function
	GetLastSibling() ITreeNode                                     // function
	GetLastSubChild() ITreeNode                                    // function
	GetLastVisibleChild(enabledOnly bool) ITreeNode                // function
	GetNext() ITreeNode                                            // function
	GetNextChild(value ITreeNode) ITreeNode                        // function
	GetNextExpanded(enabledOnly bool) ITreeNode                    // function
	GetNextMultiSelected() ITreeNode                               // function
	GetNextSibling() ITreeNode                                     // function
	GetNextSkipChildren() ITreeNode                                // function
	GetNextVisible(enabledOnly bool) ITreeNode                     // function
	GetNextVisibleSibling(enabledOnly bool) ITreeNode              // function
	GetParentNodeOfAbsoluteLevel(theAbsoluteLevel int32) ITreeNode // function
	GetPrev() ITreeNode                                            // function
	GetPrevChild(value ITreeNode) ITreeNode                        // function
	GetPrevExpanded(enabledOnly bool) ITreeNode                    // function
	GetPrevMultiSelected() ITreeNode                               // function
	GetPrevSibling() ITreeNode                                     // function
	GetPrevVisible(enabledOnly bool) ITreeNode                     // function
	GetPrevVisibleSibling(enabledOnly bool) ITreeNode              // function
	GetTextPath() string                                           // function
	HasAsParent(value ITreeNode) bool                              // function
	IndexOf(value ITreeNode) int32                                 // function
	IndexOfText(nodeText string) int32                             // function
	Collapse(recurse bool)                                         // procedure
	ConsistencyCheck()                                             // procedure
	Delete()                                                       // procedure
	DeleteChildren()                                               // procedure
	EndEdit(cancel bool)                                           // procedure
	Expand(recurse bool)                                           // procedure
	ExpandParents()                                                // procedure
	FreeAllNodeData()                                              // procedure
	MakeVisible()                                                  // procedure
	MoveTo(destination ITreeNode, mode types.TNodeAttachMode)      // procedure
	MultiSelectGroup()                                             // procedure
	Update()                                                       // procedure
	WriteDebugReport(prefix string, recurse bool)                  // procedure
	AbsoluteIndex() int32                                          // property AbsoluteIndex Getter
	Count() int32                                                  // property Count Getter
	Cut() bool                                                     // property Cut Getter
	SetCut(value bool)                                             // property Cut Setter
	Data() uintptr                                                 // property Data Getter
	SetData(value uintptr)                                         // property Data Setter
	Deleting() bool                                                // property Deleting Getter
	DropTarget() bool                                              // property DropTarget Getter
	SetDropTarget(value bool)                                      // property DropTarget Setter
	Expanded() bool                                                // property Expanded Getter
	SetExpanded(value bool)                                        // property Expanded Setter
	Focused() bool                                                 // property Focused Getter
	SetFocused(value bool)                                         // property Focused Setter
	Handle() types.TLCLHandle                                      // property Handle Getter
	HasChildren() bool                                             // property HasChildren Getter
	SetHasChildren(value bool)                                     // property HasChildren Setter
	Height() int32                                                 // property Height Getter
	SetHeight(value int32)                                         // property Height Setter
	ImageIndex() int32                                             // property ImageIndex Getter
	SetImageIndex(value int32)                                     // property ImageIndex Setter
	Index() int32                                                  // property Index Getter
	SetIndex(value int32)                                          // property Index Setter
	IsFullHeightVisible() bool                                     // property IsFullHeightVisible Getter
	IsVisible() bool                                               // property IsVisible Getter
	Items(itemIndex int32) ITreeNode                               // property Items Getter
	SetItems(itemIndex int32, value ITreeNode)                     // property Items Setter
	Level() int32                                                  // property Level Getter
	MultiSelected() bool                                           // property MultiSelected Getter
	SetMultiSelected(value bool)                                   // property MultiSelected Setter
	NodeEffect() types.TGraphicsDrawEffect                         // property NodeEffect Getter
	SetNodeEffect(value types.TGraphicsDrawEffect)                 // property NodeEffect Setter
	OverlayIndex() int32                                           // property OverlayIndex Getter
	SetOverlayIndex(value int32)                                   // property OverlayIndex Setter
	Owner() ITreeNodes                                             // property Owner Getter
	Parent() ITreeNode                                             // property Parent Getter
	Selected() bool                                                // property Selected Getter
	SetSelected(value bool)                                        // property Selected Setter
	SelectedIndex() int32                                          // property SelectedIndex Getter
	SetSelectedIndex(value int32)                                  // property SelectedIndex Setter
	StateIndex() int32                                             // property StateIndex Getter
	SetStateIndex(value int32)                                     // property StateIndex Setter
	States() types.TNodeStates                                     // property States Getter
	SubTreeCount() int32                                           // property SubTreeCount Getter
	Text() string                                                  // property Text Getter
	SetText(value string)                                          // property Text Setter
	Top() int32                                                    // property Top Getter
	TreeNodes() ITreeNodes                                         // property TreeNodes Getter
	TreeView() ICustomTreeView                                     // property TreeView Getter
	Visible() bool                                                 // property Visible Getter
	SetVisible(value bool)                                         // property Visible Setter
	Enabled() bool                                                 // property Enabled Getter
	SetEnabled(value bool)                                         // property Enabled Setter
}

type TTreeNode struct {
	TPersistent
}

func (m *TTreeNode) AlphaSort() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) Bottom() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TTreeNode) BottomExpanded() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TTreeNode) DefaultTreeViewSort(node1 ITreeNode, node2 ITreeNode) int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(node1), base.GetObjectUintptr(node2))
	return int32(r)
}

func (m *TTreeNode) DisplayExpandSignLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TTreeNode) DisplayExpandSignRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TTreeNode) DisplayExpandSignRight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TTreeNode) DisplayIconLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TTreeNode) DisplayRect(textOnly bool) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(9, m.Instance(), api.PasBool(textOnly), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TTreeNode) DisplayStateIconLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(10, m.Instance())
	return int32(r)
}

func (m *TTreeNode) DisplayTextLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(11, m.Instance())
	return int32(r)
}

func (m *TTreeNode) DisplayTextRight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(12, m.Instance())
	return int32(r)
}

func (m *TTreeNode) EditText() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(13, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) FindNode(nodeText string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(14, m.Instance(), api.PasStr(nodeText))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetFirstChild() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(15, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetFirstSibling() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(16, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetFirstVisibleChild(enabledOnly bool) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(17, m.Instance(), api.PasBool(enabledOnly))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetHandle() types.TLCLHandle {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(18, m.Instance())
	return types.TLCLHandle(r)
}

func (m *TTreeNode) GetLastChild() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(19, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetLastSibling() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(20, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetLastSubChild() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(21, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetLastVisibleChild(enabledOnly bool) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(22, m.Instance(), api.PasBool(enabledOnly))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetNext() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(23, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetNextChild(value ITreeNode) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(24, m.Instance(), base.GetObjectUintptr(value))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetNextExpanded(enabledOnly bool) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(25, m.Instance(), api.PasBool(enabledOnly))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetNextMultiSelected() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(26, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetNextSibling() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(27, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetNextSkipChildren() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(28, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetNextVisible(enabledOnly bool) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(29, m.Instance(), api.PasBool(enabledOnly))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetNextVisibleSibling(enabledOnly bool) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(30, m.Instance(), api.PasBool(enabledOnly))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetParentNodeOfAbsoluteLevel(theAbsoluteLevel int32) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(31, m.Instance(), uintptr(theAbsoluteLevel))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetPrev() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(32, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetPrevChild(value ITreeNode) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(33, m.Instance(), base.GetObjectUintptr(value))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetPrevExpanded(enabledOnly bool) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(34, m.Instance(), api.PasBool(enabledOnly))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetPrevMultiSelected() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(35, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetPrevSibling() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(36, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) GetPrevVisible(enabledOnly bool) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(37, m.Instance(), api.PasBool(enabledOnly))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetPrevVisibleSibling(enabledOnly bool) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(38, m.Instance(), api.PasBool(enabledOnly))
	return AsTreeNode(r)
}

func (m *TTreeNode) GetTextPath() string {
	if !m.IsValid() {
		return ""
	}
	r := treeNodeAPI().SysCallN(39, m.Instance())
	return api.GoStr(r)
}

func (m *TTreeNode) HasAsParent(value ITreeNode) bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(40, m.Instance(), base.GetObjectUintptr(value))
	return api.GoBool(r)
}

func (m *TTreeNode) IndexOf(value ITreeNode) int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(41, m.Instance(), base.GetObjectUintptr(value))
	return int32(r)
}

func (m *TTreeNode) IndexOfText(nodeText string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(42, m.Instance(), api.PasStr(nodeText))
	return int32(r)
}

func (m *TTreeNode) Collapse(recurse bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(43, m.Instance(), api.PasBool(recurse))
}

func (m *TTreeNode) ConsistencyCheck() {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(44, m.Instance())
}

func (m *TTreeNode) Delete() {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(45, m.Instance())
}

func (m *TTreeNode) DeleteChildren() {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(46, m.Instance())
}

func (m *TTreeNode) EndEdit(cancel bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(47, m.Instance(), api.PasBool(cancel))
}

func (m *TTreeNode) Expand(recurse bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(48, m.Instance(), api.PasBool(recurse))
}

func (m *TTreeNode) ExpandParents() {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(49, m.Instance())
}

func (m *TTreeNode) FreeAllNodeData() {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(50, m.Instance())
}

func (m *TTreeNode) MakeVisible() {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(51, m.Instance())
}

func (m *TTreeNode) MoveTo(destination ITreeNode, mode types.TNodeAttachMode) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(52, m.Instance(), base.GetObjectUintptr(destination), uintptr(mode))
}

func (m *TTreeNode) MultiSelectGroup() {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(53, m.Instance())
}

func (m *TTreeNode) Update() {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(54, m.Instance())
}

func (m *TTreeNode) WriteDebugReport(prefix string, recurse bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(55, m.Instance(), api.PasStr(prefix), api.PasBool(recurse))
}

func (m *TTreeNode) AbsoluteIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(56, m.Instance())
	return int32(r)
}

func (m *TTreeNode) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(57, m.Instance())
	return int32(r)
}

func (m *TTreeNode) Cut() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(58, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) SetCut(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(58, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeNode) Data() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(59, 0, m.Instance())
	return uintptr(r)
}

func (m *TTreeNode) SetData(value uintptr) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(59, 1, m.Instance(), uintptr(value))
}

func (m *TTreeNode) Deleting() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(60, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) DropTarget() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(61, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) SetDropTarget(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(61, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeNode) Expanded() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(62, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) SetExpanded(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(62, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeNode) Focused() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(63, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) SetFocused(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(63, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeNode) Handle() types.TLCLHandle {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(64, m.Instance())
	return types.TLCLHandle(r)
}

func (m *TTreeNode) HasChildren() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(65, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) SetHasChildren(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(65, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeNode) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(66, 0, m.Instance())
	return int32(r)
}

func (m *TTreeNode) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(66, 1, m.Instance(), uintptr(value))
}

func (m *TTreeNode) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(67, 0, m.Instance())
	return int32(r)
}

func (m *TTreeNode) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(67, 1, m.Instance(), uintptr(value))
}

func (m *TTreeNode) Index() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(68, 0, m.Instance())
	return int32(r)
}

func (m *TTreeNode) SetIndex(value int32) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(68, 1, m.Instance(), uintptr(value))
}

func (m *TTreeNode) IsFullHeightVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(69, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) IsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(70, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) Items(itemIndex int32) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(71, 0, m.Instance(), uintptr(itemIndex))
	return AsTreeNode(r)
}

func (m *TTreeNode) SetItems(itemIndex int32, value ITreeNode) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(71, 1, m.Instance(), uintptr(itemIndex), base.GetObjectUintptr(value))
}

func (m *TTreeNode) Level() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(72, m.Instance())
	return int32(r)
}

func (m *TTreeNode) MultiSelected() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(73, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) SetMultiSelected(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(73, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeNode) NodeEffect() types.TGraphicsDrawEffect {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(74, 0, m.Instance())
	return types.TGraphicsDrawEffect(r)
}

func (m *TTreeNode) SetNodeEffect(value types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(74, 1, m.Instance(), uintptr(value))
}

func (m *TTreeNode) OverlayIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(75, 0, m.Instance())
	return int32(r)
}

func (m *TTreeNode) SetOverlayIndex(value int32) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(75, 1, m.Instance(), uintptr(value))
}

func (m *TTreeNode) Owner() ITreeNodes {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(76, m.Instance())
	return AsTreeNodes(r)
}

func (m *TTreeNode) Parent() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(77, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNode) Selected() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(78, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) SetSelected(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(78, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeNode) SelectedIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(79, 0, m.Instance())
	return int32(r)
}

func (m *TTreeNode) SetSelectedIndex(value int32) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(79, 1, m.Instance(), uintptr(value))
}

func (m *TTreeNode) StateIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(80, 0, m.Instance())
	return int32(r)
}

func (m *TTreeNode) SetStateIndex(value int32) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(80, 1, m.Instance(), uintptr(value))
}

func (m *TTreeNode) States() types.TNodeStates {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(81, m.Instance())
	return types.TNodeStates(r)
}

func (m *TTreeNode) SubTreeCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(82, m.Instance())
	return int32(r)
}

func (m *TTreeNode) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := treeNodeAPI().SysCallN(83, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTreeNode) SetText(value string) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(83, 1, m.Instance(), api.PasStr(value))
}

func (m *TTreeNode) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodeAPI().SysCallN(84, m.Instance())
	return int32(r)
}

func (m *TTreeNode) TreeNodes() ITreeNodes {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(85, m.Instance())
	return AsTreeNodes(r)
}

func (m *TTreeNode) TreeView() ICustomTreeView {
	if !m.IsValid() {
		return nil
	}
	r := treeNodeAPI().SysCallN(86, m.Instance())
	return AsCustomTreeView(r)
}

func (m *TTreeNode) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(87, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(87, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeNode) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodeAPI().SysCallN(88, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNode) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodeAPI().SysCallN(88, 1, m.Instance(), api.PasBool(value))
}

// NewTreeNode class constructor
func NewTreeNode(anOwner ITreeNodes) ITreeNode {
	r := treeNodeAPI().SysCallN(0, base.GetObjectUintptr(anOwner))
	return AsTreeNode(r)
}

var (
	treeNodeOnce   base.Once
	treeNodeImport *imports.Imports = nil
)

func treeNodeAPI() *imports.Imports {
	treeNodeOnce.Do(func() {
		treeNodeImport = api.NewDefaultImports()
		treeNodeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTreeNode_Create", 0), // constructor NewTreeNode
			/* 1 */ imports.NewTable("TTreeNode_AlphaSort", 0), // function AlphaSort
			/* 2 */ imports.NewTable("TTreeNode_Bottom", 0), // function Bottom
			/* 3 */ imports.NewTable("TTreeNode_BottomExpanded", 0), // function BottomExpanded
			/* 4 */ imports.NewTable("TTreeNode_DefaultTreeViewSort", 0), // function DefaultTreeViewSort
			/* 5 */ imports.NewTable("TTreeNode_DisplayExpandSignLeft", 0), // function DisplayExpandSignLeft
			/* 6 */ imports.NewTable("TTreeNode_DisplayExpandSignRect", 0), // function DisplayExpandSignRect
			/* 7 */ imports.NewTable("TTreeNode_DisplayExpandSignRight", 0), // function DisplayExpandSignRight
			/* 8 */ imports.NewTable("TTreeNode_DisplayIconLeft", 0), // function DisplayIconLeft
			/* 9 */ imports.NewTable("TTreeNode_DisplayRect", 0), // function DisplayRect
			/* 10 */ imports.NewTable("TTreeNode_DisplayStateIconLeft", 0), // function DisplayStateIconLeft
			/* 11 */ imports.NewTable("TTreeNode_DisplayTextLeft", 0), // function DisplayTextLeft
			/* 12 */ imports.NewTable("TTreeNode_DisplayTextRight", 0), // function DisplayTextRight
			/* 13 */ imports.NewTable("TTreeNode_EditText", 0), // function EditText
			/* 14 */ imports.NewTable("TTreeNode_FindNode", 0), // function FindNode
			/* 15 */ imports.NewTable("TTreeNode_GetFirstChild", 0), // function GetFirstChild
			/* 16 */ imports.NewTable("TTreeNode_GetFirstSibling", 0), // function GetFirstSibling
			/* 17 */ imports.NewTable("TTreeNode_GetFirstVisibleChild", 0), // function GetFirstVisibleChild
			/* 18 */ imports.NewTable("TTreeNode_GetHandle", 0), // function GetHandle
			/* 19 */ imports.NewTable("TTreeNode_GetLastChild", 0), // function GetLastChild
			/* 20 */ imports.NewTable("TTreeNode_GetLastSibling", 0), // function GetLastSibling
			/* 21 */ imports.NewTable("TTreeNode_GetLastSubChild", 0), // function GetLastSubChild
			/* 22 */ imports.NewTable("TTreeNode_GetLastVisibleChild", 0), // function GetLastVisibleChild
			/* 23 */ imports.NewTable("TTreeNode_GetNext", 0), // function GetNext
			/* 24 */ imports.NewTable("TTreeNode_GetNextChild", 0), // function GetNextChild
			/* 25 */ imports.NewTable("TTreeNode_GetNextExpanded", 0), // function GetNextExpanded
			/* 26 */ imports.NewTable("TTreeNode_GetNextMultiSelected", 0), // function GetNextMultiSelected
			/* 27 */ imports.NewTable("TTreeNode_GetNextSibling", 0), // function GetNextSibling
			/* 28 */ imports.NewTable("TTreeNode_GetNextSkipChildren", 0), // function GetNextSkipChildren
			/* 29 */ imports.NewTable("TTreeNode_GetNextVisible", 0), // function GetNextVisible
			/* 30 */ imports.NewTable("TTreeNode_GetNextVisibleSibling", 0), // function GetNextVisibleSibling
			/* 31 */ imports.NewTable("TTreeNode_GetParentNodeOfAbsoluteLevel", 0), // function GetParentNodeOfAbsoluteLevel
			/* 32 */ imports.NewTable("TTreeNode_GetPrev", 0), // function GetPrev
			/* 33 */ imports.NewTable("TTreeNode_GetPrevChild", 0), // function GetPrevChild
			/* 34 */ imports.NewTable("TTreeNode_GetPrevExpanded", 0), // function GetPrevExpanded
			/* 35 */ imports.NewTable("TTreeNode_GetPrevMultiSelected", 0), // function GetPrevMultiSelected
			/* 36 */ imports.NewTable("TTreeNode_GetPrevSibling", 0), // function GetPrevSibling
			/* 37 */ imports.NewTable("TTreeNode_GetPrevVisible", 0), // function GetPrevVisible
			/* 38 */ imports.NewTable("TTreeNode_GetPrevVisibleSibling", 0), // function GetPrevVisibleSibling
			/* 39 */ imports.NewTable("TTreeNode_GetTextPath", 0), // function GetTextPath
			/* 40 */ imports.NewTable("TTreeNode_HasAsParent", 0), // function HasAsParent
			/* 41 */ imports.NewTable("TTreeNode_IndexOf", 0), // function IndexOf
			/* 42 */ imports.NewTable("TTreeNode_IndexOfText", 0), // function IndexOfText
			/* 43 */ imports.NewTable("TTreeNode_Collapse", 0), // procedure Collapse
			/* 44 */ imports.NewTable("TTreeNode_ConsistencyCheck", 0), // procedure ConsistencyCheck
			/* 45 */ imports.NewTable("TTreeNode_Delete", 0), // procedure Delete
			/* 46 */ imports.NewTable("TTreeNode_DeleteChildren", 0), // procedure DeleteChildren
			/* 47 */ imports.NewTable("TTreeNode_EndEdit", 0), // procedure EndEdit
			/* 48 */ imports.NewTable("TTreeNode_Expand", 0), // procedure Expand
			/* 49 */ imports.NewTable("TTreeNode_ExpandParents", 0), // procedure ExpandParents
			/* 50 */ imports.NewTable("TTreeNode_FreeAllNodeData", 0), // procedure FreeAllNodeData
			/* 51 */ imports.NewTable("TTreeNode_MakeVisible", 0), // procedure MakeVisible
			/* 52 */ imports.NewTable("TTreeNode_MoveTo", 0), // procedure MoveTo
			/* 53 */ imports.NewTable("TTreeNode_MultiSelectGroup", 0), // procedure MultiSelectGroup
			/* 54 */ imports.NewTable("TTreeNode_Update", 0), // procedure Update
			/* 55 */ imports.NewTable("TTreeNode_WriteDebugReport", 0), // procedure WriteDebugReport
			/* 56 */ imports.NewTable("TTreeNode_AbsoluteIndex", 0), // property AbsoluteIndex
			/* 57 */ imports.NewTable("TTreeNode_Count", 0), // property Count
			/* 58 */ imports.NewTable("TTreeNode_Cut", 0), // property Cut
			/* 59 */ imports.NewTable("TTreeNode_Data", 0), // property Data
			/* 60 */ imports.NewTable("TTreeNode_Deleting", 0), // property Deleting
			/* 61 */ imports.NewTable("TTreeNode_DropTarget", 0), // property DropTarget
			/* 62 */ imports.NewTable("TTreeNode_Expanded", 0), // property Expanded
			/* 63 */ imports.NewTable("TTreeNode_Focused", 0), // property Focused
			/* 64 */ imports.NewTable("TTreeNode_Handle", 0), // property Handle
			/* 65 */ imports.NewTable("TTreeNode_HasChildren", 0), // property HasChildren
			/* 66 */ imports.NewTable("TTreeNode_Height", 0), // property Height
			/* 67 */ imports.NewTable("TTreeNode_ImageIndex", 0), // property ImageIndex
			/* 68 */ imports.NewTable("TTreeNode_Index", 0), // property Index
			/* 69 */ imports.NewTable("TTreeNode_IsFullHeightVisible", 0), // property IsFullHeightVisible
			/* 70 */ imports.NewTable("TTreeNode_IsVisible", 0), // property IsVisible
			/* 71 */ imports.NewTable("TTreeNode_Items", 0), // property Items
			/* 72 */ imports.NewTable("TTreeNode_Level", 0), // property Level
			/* 73 */ imports.NewTable("TTreeNode_MultiSelected", 0), // property MultiSelected
			/* 74 */ imports.NewTable("TTreeNode_NodeEffect", 0), // property NodeEffect
			/* 75 */ imports.NewTable("TTreeNode_OverlayIndex", 0), // property OverlayIndex
			/* 76 */ imports.NewTable("TTreeNode_Owner", 0), // property Owner
			/* 77 */ imports.NewTable("TTreeNode_Parent", 0), // property Parent
			/* 78 */ imports.NewTable("TTreeNode_Selected", 0), // property Selected
			/* 79 */ imports.NewTable("TTreeNode_SelectedIndex", 0), // property SelectedIndex
			/* 80 */ imports.NewTable("TTreeNode_StateIndex", 0), // property StateIndex
			/* 81 */ imports.NewTable("TTreeNode_States", 0), // property States
			/* 82 */ imports.NewTable("TTreeNode_SubTreeCount", 0), // property SubTreeCount
			/* 83 */ imports.NewTable("TTreeNode_Text", 0), // property Text
			/* 84 */ imports.NewTable("TTreeNode_Top", 0), // property Top
			/* 85 */ imports.NewTable("TTreeNode_TreeNodes", 0), // property TreeNodes
			/* 86 */ imports.NewTable("TTreeNode_TreeView", 0), // property TreeView
			/* 87 */ imports.NewTable("TTreeNode_Visible", 0), // property Visible
			/* 88 */ imports.NewTable("TTreeNode_Enabled", 0), // property Enabled
		}
	})
	return treeNodeImport
}
