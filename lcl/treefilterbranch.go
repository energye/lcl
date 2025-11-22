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
)

// ITreeFilterBranch Parent: IObject
type ITreeFilterBranch interface {
	IObject
	GetData(index int32) IObject                                    // function
	AddNodeData(nodeText string, data IObject, fullFilename string) // procedure
	DeleteData(node ITreeNode)                                      // procedure
	FreeNodeData(node ITreeNode)                                    // procedure
	ClearNodeData()                                                 // procedure
	InvalidateBranch()                                              // procedure
	Move(curIndex int32, newIndex int32)                            // procedure
	Items() IStringList                                             // property Items Getter
}

type TTreeFilterBranch struct {
	TObject
}

func (m *TTreeFilterBranch) GetData(index int32) IObject {
	if !m.IsValid() {
		return nil
	}
	r := treeFilterBranchAPI().SysCallN(1, m.Instance(), uintptr(index))
	return AsObject(r)
}

func (m *TTreeFilterBranch) AddNodeData(nodeText string, data IObject, fullFilename string) {
	if !m.IsValid() {
		return
	}
	treeFilterBranchAPI().SysCallN(2, m.Instance(), api.PasStr(nodeText), base.GetObjectUintptr(data), api.PasStr(fullFilename))
}

func (m *TTreeFilterBranch) DeleteData(node ITreeNode) {
	if !m.IsValid() {
		return
	}
	treeFilterBranchAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(node))
}

func (m *TTreeFilterBranch) FreeNodeData(node ITreeNode) {
	if !m.IsValid() {
		return
	}
	treeFilterBranchAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(node))
}

func (m *TTreeFilterBranch) ClearNodeData() {
	if !m.IsValid() {
		return
	}
	treeFilterBranchAPI().SysCallN(5, m.Instance())
}

func (m *TTreeFilterBranch) InvalidateBranch() {
	if !m.IsValid() {
		return
	}
	treeFilterBranchAPI().SysCallN(6, m.Instance())
}

func (m *TTreeFilterBranch) Move(curIndex int32, newIndex int32) {
	if !m.IsValid() {
		return
	}
	treeFilterBranchAPI().SysCallN(7, m.Instance(), uintptr(curIndex), uintptr(newIndex))
}

func (m *TTreeFilterBranch) Items() IStringList {
	if !m.IsValid() {
		return nil
	}
	r := treeFilterBranchAPI().SysCallN(8, m.Instance())
	return AsStringList(r)
}

// NewTreeFilterBranch class constructor
func NewTreeFilterBranch(owner ITreeFilterEdit, rootNode ITreeNode) ITreeFilterBranch {
	r := treeFilterBranchAPI().SysCallN(0, base.GetObjectUintptr(owner), base.GetObjectUintptr(rootNode))
	return AsTreeFilterBranch(r)
}

var (
	treeFilterBranchOnce   base.Once
	treeFilterBranchImport *imports.Imports = nil
)

func treeFilterBranchAPI() *imports.Imports {
	treeFilterBranchOnce.Do(func() {
		treeFilterBranchImport = api.NewDefaultImports()
		treeFilterBranchImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTreeFilterBranch_Create", 0), // constructor NewTreeFilterBranch
			/* 1 */ imports.NewTable("TTreeFilterBranch_GetData", 0), // function GetData
			/* 2 */ imports.NewTable("TTreeFilterBranch_AddNodeData", 0), // procedure AddNodeData
			/* 3 */ imports.NewTable("TTreeFilterBranch_DeleteData", 0), // procedure DeleteData
			/* 4 */ imports.NewTable("TTreeFilterBranch_FreeNodeData", 0), // procedure FreeNodeData
			/* 5 */ imports.NewTable("TTreeFilterBranch_ClearNodeData", 0), // procedure ClearNodeData
			/* 6 */ imports.NewTable("TTreeFilterBranch_InvalidateBranch", 0), // procedure InvalidateBranch
			/* 7 */ imports.NewTable("TTreeFilterBranch_Move", 0), // procedure Move
			/* 8 */ imports.NewTable("TTreeFilterBranch_Items", 0), // property Items
		}
	})
	return treeFilterBranchImport
}
