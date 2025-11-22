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

// ITreeNodes Parent: IPersistent
type ITreeNodes interface {
	IPersistent
	Add(siblingNode ITreeNode, S string) ITreeNode                                                             // function
	AddChild(parentNode ITreeNode, S string) ITreeNode                                                         // function
	AddChildFirst(parentNode ITreeNode, S string) ITreeNode                                                    // function
	AddChildObject(parentNode ITreeNode, S string, data uintptr) ITreeNode                                     // function
	AddChildObjectFirst(parentNode ITreeNode, S string, data uintptr) ITreeNode                                // function
	AddFirst(siblingNode ITreeNode, S string) ITreeNode                                                        // function
	AddNode(node ITreeNode, relative ITreeNode, S string, ptr uintptr, method types.TNodeAttachMode) ITreeNode // function
	AddObject(siblingNode ITreeNode, S string, data uintptr) ITreeNode                                         // function
	AddObjectFirst(siblingNode ITreeNode, S string, data uintptr) ITreeNode                                    // function
	FindNodeWithData(nodeData uintptr) ITreeNode                                                               // function
	FindNodeWithText(nodeText string) ITreeNode                                                                // function
	FindNodeWithTextPath(textPath string) ITreeNode                                                            // function
	FindTopLvlNode(nodeText string) ITreeNode                                                                  // function
	GetEnumerator() ITreeNodesEnumerator                                                                       // function
	GetFirstNode() ITreeNode                                                                                   // function
	GetFirstVisibleNode() ITreeNode                                                                            // function
	GetFirstVisibleEnabledNode() ITreeNode                                                                     // function
	GetLastExpandedSubNode() ITreeNode                                                                         // function
	GetLastNode() ITreeNode                                                                                    // function
	GetLastSubNode() ITreeNode                                                                                 // function
	GetLastVisibleNode() ITreeNode                                                                             // function
	GetLastVisibleEnabledNode() ITreeNode                                                                      // function
	GetSelections(index int32) ITreeNode                                                                       // function
	Insert(nextNode ITreeNode, S string) ITreeNode                                                             // function
	InsertBehind(prevNode ITreeNode, S string) ITreeNode                                                       // function
	InsertObject(nextNode ITreeNode, S string, data uintptr) ITreeNode                                         // function
	InsertObjectBehind(prevNode ITreeNode, S string, data uintptr) ITreeNode                                   // function
	IsMultiSelection() bool                                                                                    // function
	IsUpdating() bool                                                                                          // function
	BeginUpdate()                                                                                              // procedure
	Clear()                                                                                                    // procedure
	ClearMultiSelection(clearSelected bool)                                                                    // procedure
	ConsistencyCheck()                                                                                         // procedure
	Delete(node ITreeNode)                                                                                     // procedure
	EndUpdate()                                                                                                // procedure
	FreeAllNodeData()                                                                                          // procedure
	SelectionsChanged(node ITreeNode, isSelected bool)                                                         // procedure
	SelectOnlyThis(node ITreeNode)                                                                             // procedure
	MultiSelect(node ITreeNode, clearWholeSelection bool)                                                      // procedure
	WriteDebugReport(prefix string, allNodes bool)                                                             // procedure
	Count() int32                                                                                              // property Count Getter
	Item(index int32) ITreeNode                                                                                // property Item Getter
	KeepCollapsedNodes() bool                                                                                  // property KeepCollapsedNodes Getter
	SetKeepCollapsedNodes(value bool)                                                                          // property KeepCollapsedNodes Setter
	Owner() ICustomTreeView                                                                                    // property Owner Getter
	SelectionCount() uint32                                                                                    // property SelectionCount Getter
	TopLvlCount() int32                                                                                        // property TopLvlCount Getter
	TopLvlItems(index int32) ITreeNode                                                                         // property TopLvlItems Getter
	SetTopLvlItems(index int32, value ITreeNode)                                                               // property TopLvlItems Setter
}

type TTreeNodes struct {
	TPersistent
}

func (m *TTreeNodes) Add(siblingNode ITreeNode, S string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(siblingNode), api.PasStr(S))
	return AsTreeNode(r)
}

func (m *TTreeNodes) AddChild(parentNode ITreeNode, S string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(parentNode), api.PasStr(S))
	return AsTreeNode(r)
}

func (m *TTreeNodes) AddChildFirst(parentNode ITreeNode, S string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(parentNode), api.PasStr(S))
	return AsTreeNode(r)
}

func (m *TTreeNodes) AddChildObject(parentNode ITreeNode, S string, data uintptr) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(parentNode), api.PasStr(S), uintptr(data))
	return AsTreeNode(r)
}

func (m *TTreeNodes) AddChildObjectFirst(parentNode ITreeNode, S string, data uintptr) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(parentNode), api.PasStr(S), uintptr(data))
	return AsTreeNode(r)
}

func (m *TTreeNodes) AddFirst(siblingNode ITreeNode, S string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(siblingNode), api.PasStr(S))
	return AsTreeNode(r)
}

func (m *TTreeNodes) AddNode(node ITreeNode, relative ITreeNode, S string, ptr uintptr, method types.TNodeAttachMode) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(node), base.GetObjectUintptr(relative), api.PasStr(S), uintptr(ptr), uintptr(method))
	return AsTreeNode(r)
}

func (m *TTreeNodes) AddObject(siblingNode ITreeNode, S string, data uintptr) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(siblingNode), api.PasStr(S), uintptr(data))
	return AsTreeNode(r)
}

func (m *TTreeNodes) AddObjectFirst(siblingNode ITreeNode, S string, data uintptr) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(siblingNode), api.PasStr(S), uintptr(data))
	return AsTreeNode(r)
}

func (m *TTreeNodes) FindNodeWithData(nodeData uintptr) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(10, m.Instance(), uintptr(nodeData))
	return AsTreeNode(r)
}

func (m *TTreeNodes) FindNodeWithText(nodeText string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(11, m.Instance(), api.PasStr(nodeText))
	return AsTreeNode(r)
}

func (m *TTreeNodes) FindNodeWithTextPath(textPath string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(12, m.Instance(), api.PasStr(textPath))
	return AsTreeNode(r)
}

func (m *TTreeNodes) FindTopLvlNode(nodeText string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(13, m.Instance(), api.PasStr(nodeText))
	return AsTreeNode(r)
}

func (m *TTreeNodes) GetEnumerator() ITreeNodesEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(14, m.Instance())
	return AsTreeNodesEnumerator(r)
}

func (m *TTreeNodes) GetFirstNode() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(15, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNodes) GetFirstVisibleNode() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(16, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNodes) GetFirstVisibleEnabledNode() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(17, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNodes) GetLastExpandedSubNode() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(18, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNodes) GetLastNode() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(19, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNodes) GetLastSubNode() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(20, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNodes) GetLastVisibleNode() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(21, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNodes) GetLastVisibleEnabledNode() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(22, m.Instance())
	return AsTreeNode(r)
}

func (m *TTreeNodes) GetSelections(index int32) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(23, m.Instance(), uintptr(index))
	return AsTreeNode(r)
}

func (m *TTreeNodes) Insert(nextNode ITreeNode, S string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(24, m.Instance(), base.GetObjectUintptr(nextNode), api.PasStr(S))
	return AsTreeNode(r)
}

func (m *TTreeNodes) InsertBehind(prevNode ITreeNode, S string) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(25, m.Instance(), base.GetObjectUintptr(prevNode), api.PasStr(S))
	return AsTreeNode(r)
}

func (m *TTreeNodes) InsertObject(nextNode ITreeNode, S string, data uintptr) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(26, m.Instance(), base.GetObjectUintptr(nextNode), api.PasStr(S), uintptr(data))
	return AsTreeNode(r)
}

func (m *TTreeNodes) InsertObjectBehind(prevNode ITreeNode, S string, data uintptr) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(27, m.Instance(), base.GetObjectUintptr(prevNode), api.PasStr(S), uintptr(data))
	return AsTreeNode(r)
}

func (m *TTreeNodes) IsMultiSelection() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodesAPI().SysCallN(28, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNodes) IsUpdating() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodesAPI().SysCallN(29, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNodes) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(30, m.Instance())
}

func (m *TTreeNodes) Clear() {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(31, m.Instance())
}

func (m *TTreeNodes) ClearMultiSelection(clearSelected bool) {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(32, m.Instance(), api.PasBool(clearSelected))
}

func (m *TTreeNodes) ConsistencyCheck() {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(33, m.Instance())
}

func (m *TTreeNodes) Delete(node ITreeNode) {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(34, m.Instance(), base.GetObjectUintptr(node))
}

func (m *TTreeNodes) EndUpdate() {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(35, m.Instance())
}

func (m *TTreeNodes) FreeAllNodeData() {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(36, m.Instance())
}

func (m *TTreeNodes) SelectionsChanged(node ITreeNode, isSelected bool) {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(37, m.Instance(), base.GetObjectUintptr(node), api.PasBool(isSelected))
}

func (m *TTreeNodes) SelectOnlyThis(node ITreeNode) {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(38, m.Instance(), base.GetObjectUintptr(node))
}

func (m *TTreeNodes) MultiSelect(node ITreeNode, clearWholeSelection bool) {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(39, m.Instance(), base.GetObjectUintptr(node), api.PasBool(clearWholeSelection))
}

func (m *TTreeNodes) WriteDebugReport(prefix string, allNodes bool) {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(40, m.Instance(), api.PasStr(prefix), api.PasBool(allNodes))
}

func (m *TTreeNodes) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodesAPI().SysCallN(41, m.Instance())
	return int32(r)
}

func (m *TTreeNodes) Item(index int32) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(42, m.Instance(), uintptr(index))
	return AsTreeNode(r)
}

func (m *TTreeNodes) KeepCollapsedNodes() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodesAPI().SysCallN(43, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNodes) SetKeepCollapsedNodes(value bool) {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(43, 1, m.Instance(), api.PasBool(value))
}

func (m *TTreeNodes) Owner() ICustomTreeView {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(44, m.Instance())
	return AsCustomTreeView(r)
}

func (m *TTreeNodes) SelectionCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodesAPI().SysCallN(45, m.Instance())
	return uint32(r)
}

func (m *TTreeNodes) TopLvlCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := treeNodesAPI().SysCallN(46, m.Instance())
	return int32(r)
}

func (m *TTreeNodes) TopLvlItems(index int32) ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesAPI().SysCallN(47, 0, m.Instance(), uintptr(index))
	return AsTreeNode(r)
}

func (m *TTreeNodes) SetTopLvlItems(index int32, value ITreeNode) {
	if !m.IsValid() {
		return
	}
	treeNodesAPI().SysCallN(47, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

// NewTreeNodes class constructor
func NewTreeNodes(anOwner ICustomTreeView) ITreeNodes {
	r := treeNodesAPI().SysCallN(0, base.GetObjectUintptr(anOwner))
	return AsTreeNodes(r)
}

var (
	treeNodesOnce   base.Once
	treeNodesImport *imports.Imports = nil
)

func treeNodesAPI() *imports.Imports {
	treeNodesOnce.Do(func() {
		treeNodesImport = api.NewDefaultImports()
		treeNodesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTreeNodes_Create", 0), // constructor NewTreeNodes
			/* 1 */ imports.NewTable("TTreeNodes_Add", 0), // function Add
			/* 2 */ imports.NewTable("TTreeNodes_AddChild", 0), // function AddChild
			/* 3 */ imports.NewTable("TTreeNodes_AddChildFirst", 0), // function AddChildFirst
			/* 4 */ imports.NewTable("TTreeNodes_AddChildObject", 0), // function AddChildObject
			/* 5 */ imports.NewTable("TTreeNodes_AddChildObjectFirst", 0), // function AddChildObjectFirst
			/* 6 */ imports.NewTable("TTreeNodes_AddFirst", 0), // function AddFirst
			/* 7 */ imports.NewTable("TTreeNodes_AddNode", 0), // function AddNode
			/* 8 */ imports.NewTable("TTreeNodes_AddObject", 0), // function AddObject
			/* 9 */ imports.NewTable("TTreeNodes_AddObjectFirst", 0), // function AddObjectFirst
			/* 10 */ imports.NewTable("TTreeNodes_FindNodeWithData", 0), // function FindNodeWithData
			/* 11 */ imports.NewTable("TTreeNodes_FindNodeWithText", 0), // function FindNodeWithText
			/* 12 */ imports.NewTable("TTreeNodes_FindNodeWithTextPath", 0), // function FindNodeWithTextPath
			/* 13 */ imports.NewTable("TTreeNodes_FindTopLvlNode", 0), // function FindTopLvlNode
			/* 14 */ imports.NewTable("TTreeNodes_GetEnumerator", 0), // function GetEnumerator
			/* 15 */ imports.NewTable("TTreeNodes_GetFirstNode", 0), // function GetFirstNode
			/* 16 */ imports.NewTable("TTreeNodes_GetFirstVisibleNode", 0), // function GetFirstVisibleNode
			/* 17 */ imports.NewTable("TTreeNodes_GetFirstVisibleEnabledNode", 0), // function GetFirstVisibleEnabledNode
			/* 18 */ imports.NewTable("TTreeNodes_GetLastExpandedSubNode", 0), // function GetLastExpandedSubNode
			/* 19 */ imports.NewTable("TTreeNodes_GetLastNode", 0), // function GetLastNode
			/* 20 */ imports.NewTable("TTreeNodes_GetLastSubNode", 0), // function GetLastSubNode
			/* 21 */ imports.NewTable("TTreeNodes_GetLastVisibleNode", 0), // function GetLastVisibleNode
			/* 22 */ imports.NewTable("TTreeNodes_GetLastVisibleEnabledNode", 0), // function GetLastVisibleEnabledNode
			/* 23 */ imports.NewTable("TTreeNodes_GetSelections", 0), // function GetSelections
			/* 24 */ imports.NewTable("TTreeNodes_Insert", 0), // function Insert
			/* 25 */ imports.NewTable("TTreeNodes_InsertBehind", 0), // function InsertBehind
			/* 26 */ imports.NewTable("TTreeNodes_InsertObject", 0), // function InsertObject
			/* 27 */ imports.NewTable("TTreeNodes_InsertObjectBehind", 0), // function InsertObjectBehind
			/* 28 */ imports.NewTable("TTreeNodes_IsMultiSelection", 0), // function IsMultiSelection
			/* 29 */ imports.NewTable("TTreeNodes_IsUpdating", 0), // function IsUpdating
			/* 30 */ imports.NewTable("TTreeNodes_BeginUpdate", 0), // procedure BeginUpdate
			/* 31 */ imports.NewTable("TTreeNodes_Clear", 0), // procedure Clear
			/* 32 */ imports.NewTable("TTreeNodes_ClearMultiSelection", 0), // procedure ClearMultiSelection
			/* 33 */ imports.NewTable("TTreeNodes_ConsistencyCheck", 0), // procedure ConsistencyCheck
			/* 34 */ imports.NewTable("TTreeNodes_Delete", 0), // procedure Delete
			/* 35 */ imports.NewTable("TTreeNodes_EndUpdate", 0), // procedure EndUpdate
			/* 36 */ imports.NewTable("TTreeNodes_FreeAllNodeData", 0), // procedure FreeAllNodeData
			/* 37 */ imports.NewTable("TTreeNodes_SelectionsChanged", 0), // procedure SelectionsChanged
			/* 38 */ imports.NewTable("TTreeNodes_SelectOnlyThis", 0), // procedure SelectOnlyThis
			/* 39 */ imports.NewTable("TTreeNodes_MultiSelect", 0), // procedure MultiSelect
			/* 40 */ imports.NewTable("TTreeNodes_WriteDebugReport", 0), // procedure WriteDebugReport
			/* 41 */ imports.NewTable("TTreeNodes_Count", 0), // property Count
			/* 42 */ imports.NewTable("TTreeNodes_Item", 0), // property Item
			/* 43 */ imports.NewTable("TTreeNodes_KeepCollapsedNodes", 0), // property KeepCollapsedNodes
			/* 44 */ imports.NewTable("TTreeNodes_Owner", 0), // property Owner
			/* 45 */ imports.NewTable("TTreeNodes_SelectionCount", 0), // property SelectionCount
			/* 46 */ imports.NewTable("TTreeNodes_TopLvlCount", 0), // property TopLvlCount
			/* 47 */ imports.NewTable("TTreeNodes_TopLvlItems", 0), // property TopLvlItems
		}
	})
	return treeNodesImport
}
