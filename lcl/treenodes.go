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
	. "github.com/energye/lcl/types"
)

// ITreeNodes Parent: IPersistent
type ITreeNodes interface {
	IPersistent
	Count() int32                                                                                        // property
	Item(Index int32) ITreeNode                                                                          // property
	KeepCollapsedNodes() bool                                                                            // property
	SetKeepCollapsedNodes(AValue bool)                                                                   // property
	Owner() ICustomTreeView                                                                              // property
	SelectionCount() uint32                                                                              // property
	TopLvlCount() int32                                                                                  // property
	TopLvlItems(Index int32) ITreeNode                                                                   // property
	SetTopLvlItems(Index int32, AValue ITreeNode)                                                        // property
	Add(SiblingNode ITreeNode, S string) ITreeNode                                                       // function
	AddChild(ParentNode ITreeNode, S string) ITreeNode                                                   // function
	AddChildFirst(ParentNode ITreeNode, S string) ITreeNode                                              // function
	AddChildObject(ParentNode ITreeNode, S string, Data uintptr) ITreeNode                               // function
	AddChildObjectFirst(ParentNode ITreeNode, S string, Data uintptr) ITreeNode                          // function
	AddFirst(SiblingNode ITreeNode, S string) ITreeNode                                                  // function
	AddNode(Node ITreeNode, Relative ITreeNode, S string, Ptr uintptr, Method TNodeAttachMode) ITreeNode // function
	AddObject(SiblingNode ITreeNode, S string, Data uintptr) ITreeNode                                   // function
	AddObjectFirst(SiblingNode ITreeNode, S string, Data uintptr) ITreeNode                              // function
	FindNodeWithData(NodeData uintptr) ITreeNode                                                         // function
	FindNodeWithText(NodeText string) ITreeNode                                                          // function
	FindNodeWithTextPath(TextPath string) ITreeNode                                                      // function
	FindTopLvlNode(NodeText string) ITreeNode                                                            // function
	GetEnumerator() ITreeNodesEnumerator                                                                 // function
	GetFirstNode() ITreeNode                                                                             // function
	GetFirstVisibleNode() ITreeNode                                                                      // function
	GetFirstVisibleEnabledNode() ITreeNode                                                               // function
	GetLastExpandedSubNode() ITreeNode                                                                   // function
	GetLastNode() ITreeNode                                                                              // function
	GetLastSubNode() ITreeNode                                                                           // function
	GetLastVisibleNode() ITreeNode                                                                       // function
	GetLastVisibleEnabledNode() ITreeNode                                                                // function
	GetSelections(AIndex int32) ITreeNode                                                                // function
	Insert(NextNode ITreeNode, S string) ITreeNode                                                       // function
	InsertBehind(PrevNode ITreeNode, S string) ITreeNode                                                 // function
	InsertObject(NextNode ITreeNode, S string, Data uintptr) ITreeNode                                   // function
	InsertObjectBehind(PrevNode ITreeNode, S string, Data uintptr) ITreeNode                             // function
	IsMultiSelection() bool                                                                              // function
	BeginUpdate()                                                                                        // procedure
	Clear()                                                                                              // procedure
	ClearMultiSelection(ClearSelected bool)                                                              // procedure
	ConsistencyCheck()                                                                                   // procedure
	Delete(Node ITreeNode)                                                                               // procedure
	EndUpdate()                                                                                          // procedure
	FreeAllNodeData()                                                                                    // procedure
	SelectionsChanged(ANode ITreeNode, AIsSelected bool)                                                 // procedure
	SelectOnlyThis(Node ITreeNode)                                                                       // procedure
	MultiSelect(Node ITreeNode, ClearWholeSelection bool)                                                // procedure
	WriteDebugReport(Prefix string, AllNodes bool)                                                       // procedure
}

// TTreeNodes Parent: TPersistent
type TTreeNodes struct {
	TPersistent
}

func NewTreeNodes(AnOwner ICustomTreeView) ITreeNodes {
	r1 := LCL().SysCallN(5725, GetObjectUintptr(AnOwner))
	return AsTreeNodes(r1)
}

func (m *TTreeNodes) Count() int32 {
	r1 := LCL().SysCallN(5724, m.Instance())
	return int32(r1)
}

func (m *TTreeNodes) Item(Index int32) ITreeNode {
	r1 := LCL().SysCallN(5748, m.Instance(), uintptr(Index))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) KeepCollapsedNodes() bool {
	r1 := LCL().SysCallN(5749, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNodes) SetKeepCollapsedNodes(AValue bool) {
	LCL().SysCallN(5749, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNodes) Owner() ICustomTreeView {
	r1 := LCL().SysCallN(5751, m.Instance())
	return AsCustomTreeView(r1)
}

func (m *TTreeNodes) SelectionCount() uint32 {
	r1 := LCL().SysCallN(5753, m.Instance())
	return uint32(r1)
}

func (m *TTreeNodes) TopLvlCount() int32 {
	r1 := LCL().SysCallN(5755, m.Instance())
	return int32(r1)
}

func (m *TTreeNodes) TopLvlItems(Index int32) ITreeNode {
	r1 := LCL().SysCallN(5756, 0, m.Instance(), uintptr(Index))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) SetTopLvlItems(Index int32, AValue ITreeNode) {
	LCL().SysCallN(5756, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TTreeNodes) Add(SiblingNode ITreeNode, S string) ITreeNode {
	r1 := LCL().SysCallN(5710, m.Instance(), GetObjectUintptr(SiblingNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddChild(ParentNode ITreeNode, S string) ITreeNode {
	r1 := LCL().SysCallN(5711, m.Instance(), GetObjectUintptr(ParentNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddChildFirst(ParentNode ITreeNode, S string) ITreeNode {
	r1 := LCL().SysCallN(5712, m.Instance(), GetObjectUintptr(ParentNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddChildObject(ParentNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := LCL().SysCallN(5713, m.Instance(), GetObjectUintptr(ParentNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddChildObjectFirst(ParentNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := LCL().SysCallN(5714, m.Instance(), GetObjectUintptr(ParentNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddFirst(SiblingNode ITreeNode, S string) ITreeNode {
	r1 := LCL().SysCallN(5715, m.Instance(), GetObjectUintptr(SiblingNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddNode(Node ITreeNode, Relative ITreeNode, S string, Ptr uintptr, Method TNodeAttachMode) ITreeNode {
	r1 := LCL().SysCallN(5716, m.Instance(), GetObjectUintptr(Node), GetObjectUintptr(Relative), PascalStr(S), uintptr(Ptr), uintptr(Method))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddObject(SiblingNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := LCL().SysCallN(5717, m.Instance(), GetObjectUintptr(SiblingNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddObjectFirst(SiblingNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := LCL().SysCallN(5718, m.Instance(), GetObjectUintptr(SiblingNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) FindNodeWithData(NodeData uintptr) ITreeNode {
	r1 := LCL().SysCallN(5728, m.Instance(), uintptr(NodeData))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) FindNodeWithText(NodeText string) ITreeNode {
	r1 := LCL().SysCallN(5729, m.Instance(), PascalStr(NodeText))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) FindNodeWithTextPath(TextPath string) ITreeNode {
	r1 := LCL().SysCallN(5730, m.Instance(), PascalStr(TextPath))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) FindTopLvlNode(NodeText string) ITreeNode {
	r1 := LCL().SysCallN(5731, m.Instance(), PascalStr(NodeText))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetEnumerator() ITreeNodesEnumerator {
	r1 := LCL().SysCallN(5733, m.Instance())
	return AsTreeNodesEnumerator(r1)
}

func (m *TTreeNodes) GetFirstNode() ITreeNode {
	r1 := LCL().SysCallN(5734, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetFirstVisibleNode() ITreeNode {
	r1 := LCL().SysCallN(5736, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetFirstVisibleEnabledNode() ITreeNode {
	r1 := LCL().SysCallN(5735, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastExpandedSubNode() ITreeNode {
	r1 := LCL().SysCallN(5737, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastNode() ITreeNode {
	r1 := LCL().SysCallN(5738, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastSubNode() ITreeNode {
	r1 := LCL().SysCallN(5739, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastVisibleNode() ITreeNode {
	r1 := LCL().SysCallN(5741, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastVisibleEnabledNode() ITreeNode {
	r1 := LCL().SysCallN(5740, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetSelections(AIndex int32) ITreeNode {
	r1 := LCL().SysCallN(5742, m.Instance(), uintptr(AIndex))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) Insert(NextNode ITreeNode, S string) ITreeNode {
	r1 := LCL().SysCallN(5743, m.Instance(), GetObjectUintptr(NextNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) InsertBehind(PrevNode ITreeNode, S string) ITreeNode {
	r1 := LCL().SysCallN(5744, m.Instance(), GetObjectUintptr(PrevNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) InsertObject(NextNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := LCL().SysCallN(5745, m.Instance(), GetObjectUintptr(NextNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) InsertObjectBehind(PrevNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := LCL().SysCallN(5746, m.Instance(), GetObjectUintptr(PrevNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) IsMultiSelection() bool {
	r1 := LCL().SysCallN(5747, m.Instance())
	return GoBool(r1)
}

func TreeNodesClass() TClass {
	ret := LCL().SysCallN(5720)
	return TClass(ret)
}

func (m *TTreeNodes) BeginUpdate() {
	LCL().SysCallN(5719, m.Instance())
}

func (m *TTreeNodes) Clear() {
	LCL().SysCallN(5721, m.Instance())
}

func (m *TTreeNodes) ClearMultiSelection(ClearSelected bool) {
	LCL().SysCallN(5722, m.Instance(), PascalBool(ClearSelected))
}

func (m *TTreeNodes) ConsistencyCheck() {
	LCL().SysCallN(5723, m.Instance())
}

func (m *TTreeNodes) Delete(Node ITreeNode) {
	LCL().SysCallN(5726, m.Instance(), GetObjectUintptr(Node))
}

func (m *TTreeNodes) EndUpdate() {
	LCL().SysCallN(5727, m.Instance())
}

func (m *TTreeNodes) FreeAllNodeData() {
	LCL().SysCallN(5732, m.Instance())
}

func (m *TTreeNodes) SelectionsChanged(ANode ITreeNode, AIsSelected bool) {
	LCL().SysCallN(5754, m.Instance(), GetObjectUintptr(ANode), PascalBool(AIsSelected))
}

func (m *TTreeNodes) SelectOnlyThis(Node ITreeNode) {
	LCL().SysCallN(5752, m.Instance(), GetObjectUintptr(Node))
}

func (m *TTreeNodes) MultiSelect(Node ITreeNode, ClearWholeSelection bool) {
	LCL().SysCallN(5750, m.Instance(), GetObjectUintptr(Node), PascalBool(ClearWholeSelection))
}

func (m *TTreeNodes) WriteDebugReport(Prefix string, AllNodes bool) {
	LCL().SysCallN(5757, m.Instance(), PascalStr(Prefix), PascalBool(AllNodes))
}
