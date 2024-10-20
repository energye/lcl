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

// IAVLTreeNode Parent: IObject
type IAVLTreeNode interface {
	IObject
	Successor() IAVLTreeNode        // function
	Precessor() IAVLTreeNode        // function
	TreeDepth() int32               // function
	GetCount() SizeInt              // function
	Clear()                         // procedure
	ConsistencyCheck(Tree IAVLTree) // procedure
}

// TAVLTreeNode Parent: TObject
type TAVLTreeNode struct {
	TObject
}

func NewAVLTreeNode() IAVLTreeNode {
	r1 := aVLTreeNodeImportAPI().SysCallN(3)
	return AsAVLTreeNode(r1)
}

func (m *TAVLTreeNode) Successor() IAVLTreeNode {
	r1 := aVLTreeNodeImportAPI().SysCallN(6, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTreeNode) Precessor() IAVLTreeNode {
	r1 := aVLTreeNodeImportAPI().SysCallN(5, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTreeNode) TreeDepth() int32 {
	r1 := aVLTreeNodeImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TAVLTreeNode) GetCount() SizeInt {
	r1 := aVLTreeNodeImportAPI().SysCallN(4, m.Instance())
	return SizeInt(r1)
}

func AVLTreeNodeClass() TClass {
	ret := aVLTreeNodeImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TAVLTreeNode) Clear() {
	aVLTreeNodeImportAPI().SysCallN(1, m.Instance())
}

func (m *TAVLTreeNode) ConsistencyCheck(Tree IAVLTree) {
	aVLTreeNodeImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(Tree))
}

var (
	aVLTreeNodeImport       *imports.Imports = nil
	aVLTreeNodeImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("AVLTreeNode_Class", 0),
		/*1*/ imports.NewTable("AVLTreeNode_Clear", 0),
		/*2*/ imports.NewTable("AVLTreeNode_ConsistencyCheck", 0),
		/*3*/ imports.NewTable("AVLTreeNode_Create", 0),
		/*4*/ imports.NewTable("AVLTreeNode_GetCount", 0),
		/*5*/ imports.NewTable("AVLTreeNode_Precessor", 0),
		/*6*/ imports.NewTable("AVLTreeNode_Successor", 0),
		/*7*/ imports.NewTable("AVLTreeNode_TreeDepth", 0),
	}
)

func aVLTreeNodeImportAPI() *imports.Imports {
	if aVLTreeNodeImport == nil {
		aVLTreeNodeImport = NewDefaultImports()
		aVLTreeNodeImport.SetImportTable(aVLTreeNodeImportTables)
		aVLTreeNodeImportTables = nil
	}
	return aVLTreeNodeImport
}
