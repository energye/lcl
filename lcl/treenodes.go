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
	r1 := reeNodesImportAPI().SysCallN(15, GetObjectUintptr(AnOwner))
	return AsTreeNodes(r1)
}

func (m *TTreeNodes) Count() int32 {
	r1 := reeNodesImportAPI().SysCallN(14, m.Instance())
	return int32(r1)
}

func (m *TTreeNodes) Item(Index int32) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(38, m.Instance(), uintptr(Index))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) KeepCollapsedNodes() bool {
	r1 := reeNodesImportAPI().SysCallN(39, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNodes) SetKeepCollapsedNodes(AValue bool) {
	reeNodesImportAPI().SysCallN(39, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNodes) Owner() ICustomTreeView {
	r1 := reeNodesImportAPI().SysCallN(41, m.Instance())
	return AsCustomTreeView(r1)
}

func (m *TTreeNodes) SelectionCount() uint32 {
	r1 := reeNodesImportAPI().SysCallN(43, m.Instance())
	return uint32(r1)
}

func (m *TTreeNodes) TopLvlCount() int32 {
	r1 := reeNodesImportAPI().SysCallN(45, m.Instance())
	return int32(r1)
}

func (m *TTreeNodes) TopLvlItems(Index int32) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(46, 0, m.Instance(), uintptr(Index))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) SetTopLvlItems(Index int32, AValue ITreeNode) {
	reeNodesImportAPI().SysCallN(46, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TTreeNodes) Add(SiblingNode ITreeNode, S string) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(SiblingNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddChild(ParentNode ITreeNode, S string) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(ParentNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddChildFirst(ParentNode ITreeNode, S string) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(ParentNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddChildObject(ParentNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(3, m.Instance(), GetObjectUintptr(ParentNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddChildObjectFirst(ParentNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(ParentNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddFirst(SiblingNode ITreeNode, S string) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(SiblingNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddNode(Node ITreeNode, Relative ITreeNode, S string, Ptr uintptr, Method TNodeAttachMode) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(Node), GetObjectUintptr(Relative), PascalStr(S), uintptr(Ptr), uintptr(Method))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddObject(SiblingNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(SiblingNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) AddObjectFirst(SiblingNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(SiblingNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) FindNodeWithData(NodeData uintptr) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(18, m.Instance(), uintptr(NodeData))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) FindNodeWithText(NodeText string) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(19, m.Instance(), PascalStr(NodeText))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) FindNodeWithTextPath(TextPath string) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(20, m.Instance(), PascalStr(TextPath))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) FindTopLvlNode(NodeText string) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(21, m.Instance(), PascalStr(NodeText))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetEnumerator() ITreeNodesEnumerator {
	r1 := reeNodesImportAPI().SysCallN(23, m.Instance())
	return AsTreeNodesEnumerator(r1)
}

func (m *TTreeNodes) GetFirstNode() ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(24, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetFirstVisibleNode() ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(26, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetFirstVisibleEnabledNode() ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(25, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastExpandedSubNode() ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(27, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastNode() ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(28, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastSubNode() ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(29, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastVisibleNode() ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(31, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetLastVisibleEnabledNode() ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(30, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodes) GetSelections(AIndex int32) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(32, m.Instance(), uintptr(AIndex))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) Insert(NextNode ITreeNode, S string) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(33, m.Instance(), GetObjectUintptr(NextNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) InsertBehind(PrevNode ITreeNode, S string) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(34, m.Instance(), GetObjectUintptr(PrevNode), PascalStr(S))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) InsertObject(NextNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(35, m.Instance(), GetObjectUintptr(NextNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) InsertObjectBehind(PrevNode ITreeNode, S string, Data uintptr) ITreeNode {
	r1 := reeNodesImportAPI().SysCallN(36, m.Instance(), GetObjectUintptr(PrevNode), PascalStr(S), uintptr(Data))
	return AsTreeNode(r1)
}

func (m *TTreeNodes) IsMultiSelection() bool {
	r1 := reeNodesImportAPI().SysCallN(37, m.Instance())
	return GoBool(r1)
}

func TreeNodesClass() TClass {
	ret := reeNodesImportAPI().SysCallN(10)
	return TClass(ret)
}

func (m *TTreeNodes) BeginUpdate() {
	reeNodesImportAPI().SysCallN(9, m.Instance())
}

func (m *TTreeNodes) Clear() {
	reeNodesImportAPI().SysCallN(11, m.Instance())
}

func (m *TTreeNodes) ClearMultiSelection(ClearSelected bool) {
	reeNodesImportAPI().SysCallN(12, m.Instance(), PascalBool(ClearSelected))
}

func (m *TTreeNodes) ConsistencyCheck() {
	reeNodesImportAPI().SysCallN(13, m.Instance())
}

func (m *TTreeNodes) Delete(Node ITreeNode) {
	reeNodesImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(Node))
}

func (m *TTreeNodes) EndUpdate() {
	reeNodesImportAPI().SysCallN(17, m.Instance())
}

func (m *TTreeNodes) FreeAllNodeData() {
	reeNodesImportAPI().SysCallN(22, m.Instance())
}

func (m *TTreeNodes) SelectionsChanged(ANode ITreeNode, AIsSelected bool) {
	reeNodesImportAPI().SysCallN(44, m.Instance(), GetObjectUintptr(ANode), PascalBool(AIsSelected))
}

func (m *TTreeNodes) SelectOnlyThis(Node ITreeNode) {
	reeNodesImportAPI().SysCallN(42, m.Instance(), GetObjectUintptr(Node))
}

func (m *TTreeNodes) MultiSelect(Node ITreeNode, ClearWholeSelection bool) {
	reeNodesImportAPI().SysCallN(40, m.Instance(), GetObjectUintptr(Node), PascalBool(ClearWholeSelection))
}

func (m *TTreeNodes) WriteDebugReport(Prefix string, AllNodes bool) {
	reeNodesImportAPI().SysCallN(47, m.Instance(), PascalStr(Prefix), PascalBool(AllNodes))
}

var (
	reeNodesImport       *imports.Imports = nil
	reeNodesImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TreeNodes_Add", 0),
		/*1*/ imports.NewTable("TreeNodes_AddChild", 0),
		/*2*/ imports.NewTable("TreeNodes_AddChildFirst", 0),
		/*3*/ imports.NewTable("TreeNodes_AddChildObject", 0),
		/*4*/ imports.NewTable("TreeNodes_AddChildObjectFirst", 0),
		/*5*/ imports.NewTable("TreeNodes_AddFirst", 0),
		/*6*/ imports.NewTable("TreeNodes_AddNode", 0),
		/*7*/ imports.NewTable("TreeNodes_AddObject", 0),
		/*8*/ imports.NewTable("TreeNodes_AddObjectFirst", 0),
		/*9*/ imports.NewTable("TreeNodes_BeginUpdate", 0),
		/*10*/ imports.NewTable("TreeNodes_Class", 0),
		/*11*/ imports.NewTable("TreeNodes_Clear", 0),
		/*12*/ imports.NewTable("TreeNodes_ClearMultiSelection", 0),
		/*13*/ imports.NewTable("TreeNodes_ConsistencyCheck", 0),
		/*14*/ imports.NewTable("TreeNodes_Count", 0),
		/*15*/ imports.NewTable("TreeNodes_Create", 0),
		/*16*/ imports.NewTable("TreeNodes_Delete", 0),
		/*17*/ imports.NewTable("TreeNodes_EndUpdate", 0),
		/*18*/ imports.NewTable("TreeNodes_FindNodeWithData", 0),
		/*19*/ imports.NewTable("TreeNodes_FindNodeWithText", 0),
		/*20*/ imports.NewTable("TreeNodes_FindNodeWithTextPath", 0),
		/*21*/ imports.NewTable("TreeNodes_FindTopLvlNode", 0),
		/*22*/ imports.NewTable("TreeNodes_FreeAllNodeData", 0),
		/*23*/ imports.NewTable("TreeNodes_GetEnumerator", 0),
		/*24*/ imports.NewTable("TreeNodes_GetFirstNode", 0),
		/*25*/ imports.NewTable("TreeNodes_GetFirstVisibleEnabledNode", 0),
		/*26*/ imports.NewTable("TreeNodes_GetFirstVisibleNode", 0),
		/*27*/ imports.NewTable("TreeNodes_GetLastExpandedSubNode", 0),
		/*28*/ imports.NewTable("TreeNodes_GetLastNode", 0),
		/*29*/ imports.NewTable("TreeNodes_GetLastSubNode", 0),
		/*30*/ imports.NewTable("TreeNodes_GetLastVisibleEnabledNode", 0),
		/*31*/ imports.NewTable("TreeNodes_GetLastVisibleNode", 0),
		/*32*/ imports.NewTable("TreeNodes_GetSelections", 0),
		/*33*/ imports.NewTable("TreeNodes_Insert", 0),
		/*34*/ imports.NewTable("TreeNodes_InsertBehind", 0),
		/*35*/ imports.NewTable("TreeNodes_InsertObject", 0),
		/*36*/ imports.NewTable("TreeNodes_InsertObjectBehind", 0),
		/*37*/ imports.NewTable("TreeNodes_IsMultiSelection", 0),
		/*38*/ imports.NewTable("TreeNodes_Item", 0),
		/*39*/ imports.NewTable("TreeNodes_KeepCollapsedNodes", 0),
		/*40*/ imports.NewTable("TreeNodes_MultiSelect", 0),
		/*41*/ imports.NewTable("TreeNodes_Owner", 0),
		/*42*/ imports.NewTable("TreeNodes_SelectOnlyThis", 0),
		/*43*/ imports.NewTable("TreeNodes_SelectionCount", 0),
		/*44*/ imports.NewTable("TreeNodes_SelectionsChanged", 0),
		/*45*/ imports.NewTable("TreeNodes_TopLvlCount", 0),
		/*46*/ imports.NewTable("TreeNodes_TopLvlItems", 0),
		/*47*/ imports.NewTable("TreeNodes_WriteDebugReport", 0),
	}
)

func reeNodesImportAPI() *imports.Imports {
	if reeNodesImport == nil {
		reeNodesImport = NewDefaultImports()
		reeNodesImport.SetImportTable(reeNodesImportTables)
		reeNodesImportTables = nil
	}
	return reeNodesImport
}
