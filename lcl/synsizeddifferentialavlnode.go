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

// ISynSizedDifferentialAVLNode Parent: IObject
type ISynSizedDifferentialAVLNode interface {
	IObject
	TreeDepth() int32                                                                                                                                               // function
	ReplaceChildWithSynSizedDifferentialAVLNodeX2(oldNode ISynSizedDifferentialAVLNode, node ISynSizedDifferentialAVLNode) types.TReplacedChildSite                 // function
	ReplaceChildWithSSDAVLNodeX2Int(oldNode ISynSizedDifferentialAVLNode, node ISynSizedDifferentialAVLNode, anAdjustChildPosOffset int32) types.TReplacedChildSite // function
	PrecessorToSynSizedDifferentialAVLNode() ISynSizedDifferentialAVLNode                                                                                           // function
	SuccessorToSynSizedDifferentialAVLNode() ISynSizedDifferentialAVLNode                                                                                           // function
	PrecessorWithIntX2(startPosition *int32, sizesBeforeSum *int32) ISynSizedDifferentialAVLNode                                                                    // function
	SuccessorWithIntX2(startPosition *int32, sizesBeforeSum *int32) ISynSizedDifferentialAVLNode                                                                    // function
	GetSizesBeforeSum() int32                                                                                                                                       // function
	GetPosition() int32                                                                                                                                             // function
	SetLeftChildWithSynSizedDifferentialAVLNode(node ISynSizedDifferentialAVLNode)                                                                                  // procedure
	SetLeftChildWithSynSizedDifferentialAVLNodeInt(node ISynSizedDifferentialAVLNode, anAdjustChildPosOffset int32)                                                 // procedure
	SetLeftChildWithSSDAVLNodeIntX2(node ISynSizedDifferentialAVLNode, anAdjustChildPosOffset int32, leftSizeSum int32)                                             // procedure
	SetRightChildWithSynSizedDifferentialAVLNode(node ISynSizedDifferentialAVLNode)                                                                                 // procedure
	SetRightChildWithSynSizedDifferentialAVLNodeInt(node ISynSizedDifferentialAVLNode, anAdjustChildPosOffset int32)                                                // procedure
	AdjustLeftCount(value int32)                                                                                                                                    // procedure
	AdjustParentLeftCount(value int32)                                                                                                                              // procedure
	AdjustPosition(value int32)                                                                                                                                     // procedure
}

type TSynSizedDifferentialAVLNode struct {
	TObject
}

func (m *TSynSizedDifferentialAVLNode) TreeDepth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synSizedDifferentialAVLNodeAPI().SysCallN(0, m.Instance())
	return int32(r)
}

func (m *TSynSizedDifferentialAVLNode) ReplaceChildWithSynSizedDifferentialAVLNodeX2(oldNode ISynSizedDifferentialAVLNode, node ISynSizedDifferentialAVLNode) types.TReplacedChildSite {
	if !m.IsValid() {
		return 0
	}
	r := synSizedDifferentialAVLNodeAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(oldNode), base.GetObjectUintptr(node))
	return types.TReplacedChildSite(r)
}

func (m *TSynSizedDifferentialAVLNode) ReplaceChildWithSSDAVLNodeX2Int(oldNode ISynSizedDifferentialAVLNode, node ISynSizedDifferentialAVLNode, anAdjustChildPosOffset int32) types.TReplacedChildSite {
	if !m.IsValid() {
		return 0
	}
	r := synSizedDifferentialAVLNodeAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(oldNode), base.GetObjectUintptr(node), uintptr(anAdjustChildPosOffset))
	return types.TReplacedChildSite(r)
}

func (m *TSynSizedDifferentialAVLNode) PrecessorToSynSizedDifferentialAVLNode() ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	r := synSizedDifferentialAVLNodeAPI().SysCallN(3, m.Instance())
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLNode) SuccessorToSynSizedDifferentialAVLNode() ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	r := synSizedDifferentialAVLNodeAPI().SysCallN(4, m.Instance())
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLNode) PrecessorWithIntX2(startPosition *int32, sizesBeforeSum *int32) ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	startPositionPtr := uintptr(*startPosition)
	sizesBeforeSumPtr := uintptr(*sizesBeforeSum)
	r := synSizedDifferentialAVLNodeAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&startPositionPtr)), uintptr(base.UnsafePointer(&sizesBeforeSumPtr)))
	*startPosition = int32(startPositionPtr)
	*sizesBeforeSum = int32(sizesBeforeSumPtr)
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLNode) SuccessorWithIntX2(startPosition *int32, sizesBeforeSum *int32) ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	startPositionPtr := uintptr(*startPosition)
	sizesBeforeSumPtr := uintptr(*sizesBeforeSum)
	r := synSizedDifferentialAVLNodeAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&startPositionPtr)), uintptr(base.UnsafePointer(&sizesBeforeSumPtr)))
	*startPosition = int32(startPositionPtr)
	*sizesBeforeSum = int32(sizesBeforeSumPtr)
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLNode) GetSizesBeforeSum() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synSizedDifferentialAVLNodeAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TSynSizedDifferentialAVLNode) GetPosition() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synSizedDifferentialAVLNodeAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TSynSizedDifferentialAVLNode) SetLeftChildWithSynSizedDifferentialAVLNode(node ISynSizedDifferentialAVLNode) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLNodeAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(node))
}

func (m *TSynSizedDifferentialAVLNode) SetLeftChildWithSynSizedDifferentialAVLNodeInt(node ISynSizedDifferentialAVLNode, anAdjustChildPosOffset int32) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLNodeAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(node), uintptr(anAdjustChildPosOffset))
}

func (m *TSynSizedDifferentialAVLNode) SetLeftChildWithSSDAVLNodeIntX2(node ISynSizedDifferentialAVLNode, anAdjustChildPosOffset int32, leftSizeSum int32) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLNodeAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(node), uintptr(anAdjustChildPosOffset), uintptr(leftSizeSum))
}

func (m *TSynSizedDifferentialAVLNode) SetRightChildWithSynSizedDifferentialAVLNode(node ISynSizedDifferentialAVLNode) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLNodeAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(node))
}

func (m *TSynSizedDifferentialAVLNode) SetRightChildWithSynSizedDifferentialAVLNodeInt(node ISynSizedDifferentialAVLNode, anAdjustChildPosOffset int32) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLNodeAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(node), uintptr(anAdjustChildPosOffset))
}

func (m *TSynSizedDifferentialAVLNode) AdjustLeftCount(value int32) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLNodeAPI().SysCallN(14, m.Instance(), uintptr(value))
}

func (m *TSynSizedDifferentialAVLNode) AdjustParentLeftCount(value int32) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLNodeAPI().SysCallN(15, m.Instance(), uintptr(value))
}

func (m *TSynSizedDifferentialAVLNode) AdjustPosition(value int32) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLNodeAPI().SysCallN(16, m.Instance(), uintptr(value))
}

var (
	synSizedDifferentialAVLNodeOnce   base.Once
	synSizedDifferentialAVLNodeImport *imports.Imports = nil
)

func synSizedDifferentialAVLNodeAPI() *imports.Imports {
	synSizedDifferentialAVLNodeOnce.Do(func() {
		synSizedDifferentialAVLNodeImport = api.NewDefaultImports()
		synSizedDifferentialAVLNodeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynSizedDifferentialAVLNode_TreeDepth", 0), // function TreeDepth
			/* 1 */ imports.NewTable("TSynSizedDifferentialAVLNode_ReplaceChildWithSynSizedDifferentialAVLNodeX2", 0), // function ReplaceChildWithSynSizedDifferentialAVLNodeX2
			/* 2 */ imports.NewTable("TSynSizedDifferentialAVLNode_ReplaceChildWithSSDAVLNodeX2Int", 0), // function ReplaceChildWithSSDAVLNodeX2Int
			/* 3 */ imports.NewTable("TSynSizedDifferentialAVLNode_PrecessorToSynSizedDifferentialAVLNode", 0), // function PrecessorToSynSizedDifferentialAVLNode
			/* 4 */ imports.NewTable("TSynSizedDifferentialAVLNode_SuccessorToSynSizedDifferentialAVLNode", 0), // function SuccessorToSynSizedDifferentialAVLNode
			/* 5 */ imports.NewTable("TSynSizedDifferentialAVLNode_PrecessorWithIntX2", 0), // function PrecessorWithIntX2
			/* 6 */ imports.NewTable("TSynSizedDifferentialAVLNode_SuccessorWithIntX2", 0), // function SuccessorWithIntX2
			/* 7 */ imports.NewTable("TSynSizedDifferentialAVLNode_GetSizesBeforeSum", 0), // function GetSizesBeforeSum
			/* 8 */ imports.NewTable("TSynSizedDifferentialAVLNode_GetPosition", 0), // function GetPosition
			/* 9 */ imports.NewTable("TSynSizedDifferentialAVLNode_SetLeftChildWithSynSizedDifferentialAVLNode", 0), // procedure SetLeftChildWithSynSizedDifferentialAVLNode
			/* 10 */ imports.NewTable("TSynSizedDifferentialAVLNode_SetLeftChildWithSynSizedDifferentialAVLNodeInt", 0), // procedure SetLeftChildWithSynSizedDifferentialAVLNodeInt
			/* 11 */ imports.NewTable("TSynSizedDifferentialAVLNode_SetLeftChildWithSSDAVLNodeIntX2", 0), // procedure SetLeftChildWithSSDAVLNodeIntX2
			/* 12 */ imports.NewTable("TSynSizedDifferentialAVLNode_SetRightChildWithSynSizedDifferentialAVLNode", 0), // procedure SetRightChildWithSynSizedDifferentialAVLNode
			/* 13 */ imports.NewTable("TSynSizedDifferentialAVLNode_SetRightChildWithSynSizedDifferentialAVLNodeInt", 0), // procedure SetRightChildWithSynSizedDifferentialAVLNodeInt
			/* 14 */ imports.NewTable("TSynSizedDifferentialAVLNode_AdjustLeftCount", 0), // procedure AdjustLeftCount
			/* 15 */ imports.NewTable("TSynSizedDifferentialAVLNode_AdjustParentLeftCount", 0), // procedure AdjustParentLeftCount
			/* 16 */ imports.NewTable("TSynSizedDifferentialAVLNode_AdjustPosition", 0), // procedure AdjustPosition
		}
	})
	return synSizedDifferentialAVLNodeImport
}
