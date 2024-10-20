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

// ITreeNode Parent: IPersistent
type ITreeNode interface {
	IPersistent
	AbsoluteIndex() int32                                          // property
	Count() int32                                                  // property
	Cut() bool                                                     // property
	SetCut(AValue bool)                                            // property
	Data() uintptr                                                 // property
	SetData(AValue uintptr)                                        // property
	Deleting() bool                                                // property
	DropTarget() bool                                              // property
	SetDropTarget(AValue bool)                                     // property
	Expanded() bool                                                // property
	SetExpanded(AValue bool)                                       // property
	Focused() bool                                                 // property
	SetFocused(AValue bool)                                        // property
	Handle() THandle                                               // property
	HasChildren() bool                                             // property
	SetHasChildren(AValue bool)                                    // property
	Height() int32                                                 // property
	SetHeight(AValue int32)                                        // property
	ImageIndex() TImageIndex                                       // property
	SetImageIndex(AValue TImageIndex)                              // property
	Index() int32                                                  // property
	SetIndex(AValue int32)                                         // property
	IsFullHeightVisible() bool                                     // property
	IsVisible() bool                                               // property
	Items(ItemIndex int32) ITreeNode                               // property
	SetItems(ItemIndex int32, AValue ITreeNode)                    // property
	Level() int32                                                  // property
	MultiSelected() bool                                           // property
	SetMultiSelected(AValue bool)                                  // property
	NodeEffect() TGraphicsDrawEffect                               // property
	SetNodeEffect(AValue TGraphicsDrawEffect)                      // property
	OverlayIndex() int32                                           // property
	SetOverlayIndex(AValue int32)                                  // property
	Owner() ITreeNodes                                             // property
	Parent() ITreeNode                                             // property
	Selected() bool                                                // property
	SetSelected(AValue bool)                                       // property
	SelectedIndex() int32                                          // property
	SetSelectedIndex(AValue int32)                                 // property
	StateIndex() int32                                             // property
	SetStateIndex(AValue int32)                                    // property
	States() TNodeStates                                           // property
	SubTreeCount() int32                                           // property
	Text() string                                                  // property
	SetText(AValue string)                                         // property
	Top() int32                                                    // property
	TreeNodes() ITreeNodes                                         // property
	TreeView() ICustomTreeView                                     // property
	Visible() bool                                                 // property
	SetVisible(AValue bool)                                        // property
	Enabled() bool                                                 // property
	SetEnabled(AValue bool)                                        // property
	AlphaSort() bool                                               // function
	Bottom() int32                                                 // function
	BottomExpanded() int32                                         // function
	CustomSort(fn TTreeNodeCompare) bool                           // function
	DefaultTreeViewSort(Node1, Node2 ITreeNode) int32              // function
	DisplayExpandSignLeft() int32                                  // function
	DisplayExpandSignRect() (resultRect TRect)                     // function
	DisplayExpandSignRight() int32                                 // function
	DisplayIconLeft() int32                                        // function
	DisplayRect(TextOnly bool) (resultRect TRect)                  // function
	DisplayStateIconLeft() int32                                   // function
	DisplayTextLeft() int32                                        // function
	DisplayTextRight() int32                                       // function
	EditText() bool                                                // function
	FindNode(NodeText string) ITreeNode                            // function
	GetFirstChild() ITreeNode                                      // function
	GetFirstSibling() ITreeNode                                    // function
	GetFirstVisibleChild(aEnabledOnly bool) ITreeNode              // function
	GetHandle() THandle                                            // function
	GetLastChild() ITreeNode                                       // function
	GetLastSibling() ITreeNode                                     // function
	GetLastSubChild() ITreeNode                                    // function
	GetLastVisibleChild(aEnabledOnly bool) ITreeNode               // function
	GetNext() ITreeNode                                            // function
	GetNextChild(AValue ITreeNode) ITreeNode                       // function
	GetNextExpanded(aEnabledOnly bool) ITreeNode                   // function
	GetNextMultiSelected() ITreeNode                               // function
	GetNextSibling() ITreeNode                                     // function
	GetNextSkipChildren() ITreeNode                                // function
	GetNextVisible(aEnabledOnly bool) ITreeNode                    // function
	GetNextVisibleSibling(aEnabledOnly bool) ITreeNode             // function
	GetParentNodeOfAbsoluteLevel(TheAbsoluteLevel int32) ITreeNode // function
	GetPrev() ITreeNode                                            // function
	GetPrevChild(AValue ITreeNode) ITreeNode                       // function
	GetPrevExpanded(aEnabledOnly bool) ITreeNode                   // function
	GetPrevMultiSelected() ITreeNode                               // function
	GetPrevSibling() ITreeNode                                     // function
	GetPrevVisible(aEnabledOnly bool) ITreeNode                    // function
	GetPrevVisibleSibling(aEnabledOnly bool) ITreeNode             // function
	GetTextPath() string                                           // function
	HasAsParent(AValue ITreeNode) bool                             // function
	IndexOf(AValue ITreeNode) int32                                // function
	IndexOfText(NodeText string) int32                             // function
	Collapse(Recurse bool)                                         // procedure
	ConsistencyCheck()                                             // procedure
	Delete()                                                       // procedure
	DeleteChildren()                                               // procedure
	EndEdit(Cancel bool)                                           // procedure
	Expand(Recurse bool)                                           // procedure
	ExpandParents()                                                // procedure
	FreeAllNodeData()                                              // procedure
	MakeVisible()                                                  // procedure
	MoveTo(Destination ITreeNode, Mode TNodeAttachMode)            // procedure
	MultiSelectGroup()                                             // procedure
	Update()                                                       // procedure
	WriteDebugReport(Prefix string, Recurse bool)                  // procedure
}

// TTreeNode Parent: TPersistent
type TTreeNode struct {
	TPersistent
	customSortPtr uintptr
}

func NewTreeNode(AnOwner ITreeNodes) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(8, GetObjectUintptr(AnOwner))
	return AsTreeNode(r1)
}

func (m *TTreeNode) AbsoluteIndex() int32 {
	r1 := reeNodeImportAPI().SysCallN(0, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) Count() int32 {
	r1 := reeNodeImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) Cut() bool {
	r1 := reeNodeImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetCut(AValue bool) {
	reeNodeImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Data() uintptr {
	r1 := reeNodeImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TTreeNode) SetData(AValue uintptr) {
	reeNodeImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) Deleting() bool {
	r1 := reeNodeImportAPI().SysCallN(15, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) DropTarget() bool {
	r1 := reeNodeImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetDropTarget(AValue bool) {
	reeNodeImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Expanded() bool {
	r1 := reeNodeImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetExpanded(AValue bool) {
	reeNodeImportAPI().SysCallN(30, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Focused() bool {
	r1 := reeNodeImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetFocused(AValue bool) {
	reeNodeImportAPI().SysCallN(32, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Handle() THandle {
	r1 := reeNodeImportAPI().SysCallN(59, m.Instance())
	return THandle(r1)
}

func (m *TTreeNode) HasChildren() bool {
	r1 := reeNodeImportAPI().SysCallN(61, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetHasChildren(AValue bool) {
	reeNodeImportAPI().SysCallN(61, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Height() int32 {
	r1 := reeNodeImportAPI().SysCallN(62, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetHeight(AValue int32) {
	reeNodeImportAPI().SysCallN(62, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) ImageIndex() TImageIndex {
	r1 := reeNodeImportAPI().SysCallN(63, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TTreeNode) SetImageIndex(AValue TImageIndex) {
	reeNodeImportAPI().SysCallN(63, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) Index() int32 {
	r1 := reeNodeImportAPI().SysCallN(64, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetIndex(AValue int32) {
	reeNodeImportAPI().SysCallN(64, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) IsFullHeightVisible() bool {
	r1 := reeNodeImportAPI().SysCallN(67, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) IsVisible() bool {
	r1 := reeNodeImportAPI().SysCallN(68, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) Items(ItemIndex int32) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(69, 0, m.Instance(), uintptr(ItemIndex))
	return AsTreeNode(r1)
}

func (m *TTreeNode) SetItems(ItemIndex int32, AValue ITreeNode) {
	reeNodeImportAPI().SysCallN(69, 1, m.Instance(), uintptr(ItemIndex), GetObjectUintptr(AValue))
}

func (m *TTreeNode) Level() int32 {
	r1 := reeNodeImportAPI().SysCallN(70, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) MultiSelected() bool {
	r1 := reeNodeImportAPI().SysCallN(74, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetMultiSelected(AValue bool) {
	reeNodeImportAPI().SysCallN(74, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) NodeEffect() TGraphicsDrawEffect {
	r1 := reeNodeImportAPI().SysCallN(75, 0, m.Instance(), 0)
	return TGraphicsDrawEffect(r1)
}

func (m *TTreeNode) SetNodeEffect(AValue TGraphicsDrawEffect) {
	reeNodeImportAPI().SysCallN(75, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) OverlayIndex() int32 {
	r1 := reeNodeImportAPI().SysCallN(76, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetOverlayIndex(AValue int32) {
	reeNodeImportAPI().SysCallN(76, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) Owner() ITreeNodes {
	r1 := reeNodeImportAPI().SysCallN(77, m.Instance())
	return AsTreeNodes(r1)
}

func (m *TTreeNode) Parent() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(78, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) Selected() bool {
	r1 := reeNodeImportAPI().SysCallN(79, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetSelected(AValue bool) {
	reeNodeImportAPI().SysCallN(79, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) SelectedIndex() int32 {
	r1 := reeNodeImportAPI().SysCallN(80, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetSelectedIndex(AValue int32) {
	reeNodeImportAPI().SysCallN(80, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) StateIndex() int32 {
	r1 := reeNodeImportAPI().SysCallN(81, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetStateIndex(AValue int32) {
	reeNodeImportAPI().SysCallN(81, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) States() TNodeStates {
	r1 := reeNodeImportAPI().SysCallN(82, m.Instance())
	return TNodeStates(r1)
}

func (m *TTreeNode) SubTreeCount() int32 {
	r1 := reeNodeImportAPI().SysCallN(83, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) Text() string {
	r1 := reeNodeImportAPI().SysCallN(84, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TTreeNode) SetText(AValue string) {
	reeNodeImportAPI().SysCallN(84, 1, m.Instance(), PascalStr(AValue))
}

func (m *TTreeNode) Top() int32 {
	r1 := reeNodeImportAPI().SysCallN(85, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) TreeNodes() ITreeNodes {
	r1 := reeNodeImportAPI().SysCallN(86, m.Instance())
	return AsTreeNodes(r1)
}

func (m *TTreeNode) TreeView() ICustomTreeView {
	r1 := reeNodeImportAPI().SysCallN(87, m.Instance())
	return AsCustomTreeView(r1)
}

func (m *TTreeNode) Visible() bool {
	r1 := reeNodeImportAPI().SysCallN(89, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetVisible(AValue bool) {
	reeNodeImportAPI().SysCallN(89, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Enabled() bool {
	r1 := reeNodeImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetEnabled(AValue bool) {
	reeNodeImportAPI().SysCallN(26, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) AlphaSort() bool {
	r1 := reeNodeImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) Bottom() int32 {
	r1 := reeNodeImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) BottomExpanded() int32 {
	r1 := reeNodeImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) CustomSort(fn TTreeNodeCompare) bool {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	r1 := reeNodeImportAPI().SysCallN(9, m.Instance(), m.customSortPtr)
	return GoBool(r1)
}

func (m *TTreeNode) DefaultTreeViewSort(Node1, Node2 ITreeNode) int32 {
	r1 := reeNodeImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(Node1), GetObjectUintptr(Node2))
	return int32(r1)
}

func (m *TTreeNode) DisplayExpandSignLeft() int32 {
	r1 := reeNodeImportAPI().SysCallN(16, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayExpandSignRect() (resultRect TRect) {
	reeNodeImportAPI().SysCallN(17, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TTreeNode) DisplayExpandSignRight() int32 {
	r1 := reeNodeImportAPI().SysCallN(18, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayIconLeft() int32 {
	r1 := reeNodeImportAPI().SysCallN(19, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayRect(TextOnly bool) (resultRect TRect) {
	reeNodeImportAPI().SysCallN(20, m.Instance(), PascalBool(TextOnly), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TTreeNode) DisplayStateIconLeft() int32 {
	r1 := reeNodeImportAPI().SysCallN(21, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayTextLeft() int32 {
	r1 := reeNodeImportAPI().SysCallN(22, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayTextRight() int32 {
	r1 := reeNodeImportAPI().SysCallN(23, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) EditText() bool {
	r1 := reeNodeImportAPI().SysCallN(25, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) FindNode(NodeText string) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(31, m.Instance(), PascalStr(NodeText))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetFirstChild() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(34, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetFirstSibling() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(35, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetFirstVisibleChild(aEnabledOnly bool) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(36, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetHandle() THandle {
	r1 := reeNodeImportAPI().SysCallN(37, m.Instance())
	return THandle(r1)
}

func (m *TTreeNode) GetLastChild() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(38, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetLastSibling() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(39, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetLastSubChild() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(40, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetLastVisibleChild(aEnabledOnly bool) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(41, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNext() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(42, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextChild(AValue ITreeNode) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(43, m.Instance(), GetObjectUintptr(AValue))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextExpanded(aEnabledOnly bool) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(44, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextMultiSelected() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(45, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextSibling() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(46, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextSkipChildren() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(47, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextVisible(aEnabledOnly bool) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(48, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextVisibleSibling(aEnabledOnly bool) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(49, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetParentNodeOfAbsoluteLevel(TheAbsoluteLevel int32) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(50, m.Instance(), uintptr(TheAbsoluteLevel))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrev() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(51, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevChild(AValue ITreeNode) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(52, m.Instance(), GetObjectUintptr(AValue))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevExpanded(aEnabledOnly bool) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(53, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevMultiSelected() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(54, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevSibling() ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(55, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevVisible(aEnabledOnly bool) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(56, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevVisibleSibling(aEnabledOnly bool) ITreeNode {
	r1 := reeNodeImportAPI().SysCallN(57, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetTextPath() string {
	r1 := reeNodeImportAPI().SysCallN(58, m.Instance())
	return GoStr(r1)
}

func (m *TTreeNode) HasAsParent(AValue ITreeNode) bool {
	r1 := reeNodeImportAPI().SysCallN(60, m.Instance(), GetObjectUintptr(AValue))
	return GoBool(r1)
}

func (m *TTreeNode) IndexOf(AValue ITreeNode) int32 {
	r1 := reeNodeImportAPI().SysCallN(65, m.Instance(), GetObjectUintptr(AValue))
	return int32(r1)
}

func (m *TTreeNode) IndexOfText(NodeText string) int32 {
	r1 := reeNodeImportAPI().SysCallN(66, m.Instance(), PascalStr(NodeText))
	return int32(r1)
}

func TreeNodeClass() TClass {
	ret := reeNodeImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TTreeNode) Collapse(Recurse bool) {
	reeNodeImportAPI().SysCallN(5, m.Instance(), PascalBool(Recurse))
}

func (m *TTreeNode) ConsistencyCheck() {
	reeNodeImportAPI().SysCallN(6, m.Instance())
}

func (m *TTreeNode) Delete() {
	reeNodeImportAPI().SysCallN(13, m.Instance())
}

func (m *TTreeNode) DeleteChildren() {
	reeNodeImportAPI().SysCallN(14, m.Instance())
}

func (m *TTreeNode) EndEdit(Cancel bool) {
	reeNodeImportAPI().SysCallN(27, m.Instance(), PascalBool(Cancel))
}

func (m *TTreeNode) Expand(Recurse bool) {
	reeNodeImportAPI().SysCallN(28, m.Instance(), PascalBool(Recurse))
}

func (m *TTreeNode) ExpandParents() {
	reeNodeImportAPI().SysCallN(29, m.Instance())
}

func (m *TTreeNode) FreeAllNodeData() {
	reeNodeImportAPI().SysCallN(33, m.Instance())
}

func (m *TTreeNode) MakeVisible() {
	reeNodeImportAPI().SysCallN(71, m.Instance())
}

func (m *TTreeNode) MoveTo(Destination ITreeNode, Mode TNodeAttachMode) {
	reeNodeImportAPI().SysCallN(72, m.Instance(), GetObjectUintptr(Destination), uintptr(Mode))
}

func (m *TTreeNode) MultiSelectGroup() {
	reeNodeImportAPI().SysCallN(73, m.Instance())
}

func (m *TTreeNode) Update() {
	reeNodeImportAPI().SysCallN(88, m.Instance())
}

func (m *TTreeNode) WriteDebugReport(Prefix string, Recurse bool) {
	reeNodeImportAPI().SysCallN(90, m.Instance(), PascalStr(Prefix), PascalBool(Recurse))
}

var (
	reeNodeImport       *imports.Imports = nil
	reeNodeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TreeNode_AbsoluteIndex", 0),
		/*1*/ imports.NewTable("TreeNode_AlphaSort", 0),
		/*2*/ imports.NewTable("TreeNode_Bottom", 0),
		/*3*/ imports.NewTable("TreeNode_BottomExpanded", 0),
		/*4*/ imports.NewTable("TreeNode_Class", 0),
		/*5*/ imports.NewTable("TreeNode_Collapse", 0),
		/*6*/ imports.NewTable("TreeNode_ConsistencyCheck", 0),
		/*7*/ imports.NewTable("TreeNode_Count", 0),
		/*8*/ imports.NewTable("TreeNode_Create", 0),
		/*9*/ imports.NewTable("TreeNode_CustomSort", 0),
		/*10*/ imports.NewTable("TreeNode_Cut", 0),
		/*11*/ imports.NewTable("TreeNode_Data", 0),
		/*12*/ imports.NewTable("TreeNode_DefaultTreeViewSort", 0),
		/*13*/ imports.NewTable("TreeNode_Delete", 0),
		/*14*/ imports.NewTable("TreeNode_DeleteChildren", 0),
		/*15*/ imports.NewTable("TreeNode_Deleting", 0),
		/*16*/ imports.NewTable("TreeNode_DisplayExpandSignLeft", 0),
		/*17*/ imports.NewTable("TreeNode_DisplayExpandSignRect", 0),
		/*18*/ imports.NewTable("TreeNode_DisplayExpandSignRight", 0),
		/*19*/ imports.NewTable("TreeNode_DisplayIconLeft", 0),
		/*20*/ imports.NewTable("TreeNode_DisplayRect", 0),
		/*21*/ imports.NewTable("TreeNode_DisplayStateIconLeft", 0),
		/*22*/ imports.NewTable("TreeNode_DisplayTextLeft", 0),
		/*23*/ imports.NewTable("TreeNode_DisplayTextRight", 0),
		/*24*/ imports.NewTable("TreeNode_DropTarget", 0),
		/*25*/ imports.NewTable("TreeNode_EditText", 0),
		/*26*/ imports.NewTable("TreeNode_Enabled", 0),
		/*27*/ imports.NewTable("TreeNode_EndEdit", 0),
		/*28*/ imports.NewTable("TreeNode_Expand", 0),
		/*29*/ imports.NewTable("TreeNode_ExpandParents", 0),
		/*30*/ imports.NewTable("TreeNode_Expanded", 0),
		/*31*/ imports.NewTable("TreeNode_FindNode", 0),
		/*32*/ imports.NewTable("TreeNode_Focused", 0),
		/*33*/ imports.NewTable("TreeNode_FreeAllNodeData", 0),
		/*34*/ imports.NewTable("TreeNode_GetFirstChild", 0),
		/*35*/ imports.NewTable("TreeNode_GetFirstSibling", 0),
		/*36*/ imports.NewTable("TreeNode_GetFirstVisibleChild", 0),
		/*37*/ imports.NewTable("TreeNode_GetHandle", 0),
		/*38*/ imports.NewTable("TreeNode_GetLastChild", 0),
		/*39*/ imports.NewTable("TreeNode_GetLastSibling", 0),
		/*40*/ imports.NewTable("TreeNode_GetLastSubChild", 0),
		/*41*/ imports.NewTable("TreeNode_GetLastVisibleChild", 0),
		/*42*/ imports.NewTable("TreeNode_GetNext", 0),
		/*43*/ imports.NewTable("TreeNode_GetNextChild", 0),
		/*44*/ imports.NewTable("TreeNode_GetNextExpanded", 0),
		/*45*/ imports.NewTable("TreeNode_GetNextMultiSelected", 0),
		/*46*/ imports.NewTable("TreeNode_GetNextSibling", 0),
		/*47*/ imports.NewTable("TreeNode_GetNextSkipChildren", 0),
		/*48*/ imports.NewTable("TreeNode_GetNextVisible", 0),
		/*49*/ imports.NewTable("TreeNode_GetNextVisibleSibling", 0),
		/*50*/ imports.NewTable("TreeNode_GetParentNodeOfAbsoluteLevel", 0),
		/*51*/ imports.NewTable("TreeNode_GetPrev", 0),
		/*52*/ imports.NewTable("TreeNode_GetPrevChild", 0),
		/*53*/ imports.NewTable("TreeNode_GetPrevExpanded", 0),
		/*54*/ imports.NewTable("TreeNode_GetPrevMultiSelected", 0),
		/*55*/ imports.NewTable("TreeNode_GetPrevSibling", 0),
		/*56*/ imports.NewTable("TreeNode_GetPrevVisible", 0),
		/*57*/ imports.NewTable("TreeNode_GetPrevVisibleSibling", 0),
		/*58*/ imports.NewTable("TreeNode_GetTextPath", 0),
		/*59*/ imports.NewTable("TreeNode_Handle", 0),
		/*60*/ imports.NewTable("TreeNode_HasAsParent", 0),
		/*61*/ imports.NewTable("TreeNode_HasChildren", 0),
		/*62*/ imports.NewTable("TreeNode_Height", 0),
		/*63*/ imports.NewTable("TreeNode_ImageIndex", 0),
		/*64*/ imports.NewTable("TreeNode_Index", 0),
		/*65*/ imports.NewTable("TreeNode_IndexOf", 0),
		/*66*/ imports.NewTable("TreeNode_IndexOfText", 0),
		/*67*/ imports.NewTable("TreeNode_IsFullHeightVisible", 0),
		/*68*/ imports.NewTable("TreeNode_IsVisible", 0),
		/*69*/ imports.NewTable("TreeNode_Items", 0),
		/*70*/ imports.NewTable("TreeNode_Level", 0),
		/*71*/ imports.NewTable("TreeNode_MakeVisible", 0),
		/*72*/ imports.NewTable("TreeNode_MoveTo", 0),
		/*73*/ imports.NewTable("TreeNode_MultiSelectGroup", 0),
		/*74*/ imports.NewTable("TreeNode_MultiSelected", 0),
		/*75*/ imports.NewTable("TreeNode_NodeEffect", 0),
		/*76*/ imports.NewTable("TreeNode_OverlayIndex", 0),
		/*77*/ imports.NewTable("TreeNode_Owner", 0),
		/*78*/ imports.NewTable("TreeNode_Parent", 0),
		/*79*/ imports.NewTable("TreeNode_Selected", 0),
		/*80*/ imports.NewTable("TreeNode_SelectedIndex", 0),
		/*81*/ imports.NewTable("TreeNode_StateIndex", 0),
		/*82*/ imports.NewTable("TreeNode_States", 0),
		/*83*/ imports.NewTable("TreeNode_SubTreeCount", 0),
		/*84*/ imports.NewTable("TreeNode_Text", 0),
		/*85*/ imports.NewTable("TreeNode_Top", 0),
		/*86*/ imports.NewTable("TreeNode_TreeNodes", 0),
		/*87*/ imports.NewTable("TreeNode_TreeView", 0),
		/*88*/ imports.NewTable("TreeNode_Update", 0),
		/*89*/ imports.NewTable("TreeNode_Visible", 0),
		/*90*/ imports.NewTable("TreeNode_WriteDebugReport", 0),
	}
)

func reeNodeImportAPI() *imports.Imports {
	if reeNodeImport == nil {
		reeNodeImport = NewDefaultImports()
		reeNodeImport.SetImportTable(reeNodeImportTables)
		reeNodeImportTables = nil
	}
	return reeNodeImport
}
