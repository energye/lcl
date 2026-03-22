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

// ISynSizedDifferentialAVLTree Parent: IObject
type ISynSizedDifferentialAVLTree interface {
	IObject
	FirstToSynSizedDifferentialAVLNode() ISynSizedDifferentialAVLNode                                                                                       // function
	LastToSynSizedDifferentialAVLNode() ISynSizedDifferentialAVLNode                                                                                        // function
	FirstWithIntX2(outStartPosition *int32, outSizesBeforeSum *int32) ISynSizedDifferentialAVLNode                                                          // function
	LastWithIntX2(outStartPosition *int32, outSizesBeforeSum *int32) ISynSizedDifferentialAVLNode                                                           // function
	FindNodeAtLeftSize(leftSum int32, outStartPosition *int32, outSizesBeforeSum *int32) ISynSizedDifferentialAVLNode                                       // function
	FindNodeAtPosition(position int32, mode types.TSynSizedDiffAVLFindMode, outStartPosition *int32, outSizesBeforeSum *int32) ISynSizedDifferentialAVLNode // function
	Clear()                                                                                                                                                 // procedure
	AdjustForLinesInserted(startLine int32, lineCount int32)                                                                                                // procedure
	AdjustForLinesDeleted(startLine int32, lineCount int32)                                                                                                 // procedure
}

type TSynSizedDifferentialAVLTree struct {
	TObject
}

func (m *TSynSizedDifferentialAVLTree) FirstToSynSizedDifferentialAVLNode() ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	r := synSizedDifferentialAVLTreeAPI().SysCallN(1, m.Instance())
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLTree) LastToSynSizedDifferentialAVLNode() ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	r := synSizedDifferentialAVLTreeAPI().SysCallN(2, m.Instance())
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLTree) FirstWithIntX2(outStartPosition *int32, outSizesBeforeSum *int32) ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	var startPositionPtr uintptr
	var sizesBeforeSumPtr uintptr
	r := synSizedDifferentialAVLTreeAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&startPositionPtr)), uintptr(base.UnsafePointer(&sizesBeforeSumPtr)))
	*outStartPosition = int32(startPositionPtr)
	*outSizesBeforeSum = int32(sizesBeforeSumPtr)
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLTree) LastWithIntX2(outStartPosition *int32, outSizesBeforeSum *int32) ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	var startPositionPtr uintptr
	var sizesBeforeSumPtr uintptr
	r := synSizedDifferentialAVLTreeAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&startPositionPtr)), uintptr(base.UnsafePointer(&sizesBeforeSumPtr)))
	*outStartPosition = int32(startPositionPtr)
	*outSizesBeforeSum = int32(sizesBeforeSumPtr)
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLTree) FindNodeAtLeftSize(leftSum int32, outStartPosition *int32, outSizesBeforeSum *int32) ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	var startPositionPtr uintptr
	var sizesBeforeSumPtr uintptr
	r := synSizedDifferentialAVLTreeAPI().SysCallN(5, m.Instance(), uintptr(leftSum), uintptr(base.UnsafePointer(&startPositionPtr)), uintptr(base.UnsafePointer(&sizesBeforeSumPtr)))
	*outStartPosition = int32(startPositionPtr)
	*outSizesBeforeSum = int32(sizesBeforeSumPtr)
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLTree) FindNodeAtPosition(position int32, mode types.TSynSizedDiffAVLFindMode, outStartPosition *int32, outSizesBeforeSum *int32) ISynSizedDifferentialAVLNode {
	if !m.IsValid() {
		return nil
	}
	var startPositionPtr uintptr
	var sizesBeforeSumPtr uintptr
	r := synSizedDifferentialAVLTreeAPI().SysCallN(6, m.Instance(), uintptr(position), uintptr(mode), uintptr(base.UnsafePointer(&startPositionPtr)), uintptr(base.UnsafePointer(&sizesBeforeSumPtr)))
	*outStartPosition = int32(startPositionPtr)
	*outSizesBeforeSum = int32(sizesBeforeSumPtr)
	return AsSynSizedDifferentialAVLNode(r)
}

func (m *TSynSizedDifferentialAVLTree) Clear() {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLTreeAPI().SysCallN(7, m.Instance())
}

func (m *TSynSizedDifferentialAVLTree) AdjustForLinesInserted(startLine int32, lineCount int32) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLTreeAPI().SysCallN(8, m.Instance(), uintptr(startLine), uintptr(lineCount))
}

func (m *TSynSizedDifferentialAVLTree) AdjustForLinesDeleted(startLine int32, lineCount int32) {
	if !m.IsValid() {
		return
	}
	synSizedDifferentialAVLTreeAPI().SysCallN(9, m.Instance(), uintptr(startLine), uintptr(lineCount))
}

// NewSynSizedDifferentialAVLTree class constructor
func NewSynSizedDifferentialAVLTree() ISynSizedDifferentialAVLTree {
	r := synSizedDifferentialAVLTreeAPI().SysCallN(0)
	return AsSynSizedDifferentialAVLTree(r)
}

var (
	synSizedDifferentialAVLTreeOnce   base.Once
	synSizedDifferentialAVLTreeImport *imports.Imports = nil
)

func synSizedDifferentialAVLTreeAPI() *imports.Imports {
	synSizedDifferentialAVLTreeOnce.Do(func() {
		synSizedDifferentialAVLTreeImport = api.NewDefaultImports()
		synSizedDifferentialAVLTreeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynSizedDifferentialAVLTree_Create", 0), // constructor NewSynSizedDifferentialAVLTree
			/* 1 */ imports.NewTable("TSynSizedDifferentialAVLTree_FirstToSynSizedDifferentialAVLNode", 0), // function FirstToSynSizedDifferentialAVLNode
			/* 2 */ imports.NewTable("TSynSizedDifferentialAVLTree_LastToSynSizedDifferentialAVLNode", 0), // function LastToSynSizedDifferentialAVLNode
			/* 3 */ imports.NewTable("TSynSizedDifferentialAVLTree_FirstWithIntX2", 0), // function FirstWithIntX2
			/* 4 */ imports.NewTable("TSynSizedDifferentialAVLTree_LastWithIntX2", 0), // function LastWithIntX2
			/* 5 */ imports.NewTable("TSynSizedDifferentialAVLTree_FindNodeAtLeftSize", 0), // function FindNodeAtLeftSize
			/* 6 */ imports.NewTable("TSynSizedDifferentialAVLTree_FindNodeAtPosition", 0), // function FindNodeAtPosition
			/* 7 */ imports.NewTable("TSynSizedDifferentialAVLTree_Clear", 0), // procedure Clear
			/* 8 */ imports.NewTable("TSynSizedDifferentialAVLTree_AdjustForLinesInserted", 0), // procedure AdjustForLinesInserted
			/* 9 */ imports.NewTable("TSynSizedDifferentialAVLTree_AdjustForLinesDeleted", 0), // procedure AdjustForLinesDeleted
		}
	})
	return synSizedDifferentialAVLTreeImport
}
