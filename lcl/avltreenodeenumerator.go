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

// IAVLTreeNodeEnumerator Parent: IObject
type IAVLTreeNodeEnumerator interface {
	IObject
	Current() IAVLTreeNode                 // property
	LowToHigh() bool                       // property
	GetEnumerator() IAVLTreeNodeEnumerator // function
	MoveNext() bool                        // function
}

// TAVLTreeNodeEnumerator Parent: TObject
type TAVLTreeNodeEnumerator struct {
	TObject
}

func NewAVLTreeNodeEnumerator(Tree IAVLTree, aLowToHigh bool) IAVLTreeNodeEnumerator {
	r1 := aVLTreeNodeEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(Tree), PascalBool(aLowToHigh))
	return AsAVLTreeNodeEnumerator(r1)
}

func (m *TAVLTreeNodeEnumerator) Current() IAVLTreeNode {
	r1 := aVLTreeNodeEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTreeNodeEnumerator) LowToHigh() bool {
	r1 := aVLTreeNodeEnumeratorImportAPI().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TAVLTreeNodeEnumerator) GetEnumerator() IAVLTreeNodeEnumerator {
	r1 := aVLTreeNodeEnumeratorImportAPI().SysCallN(3, m.Instance())
	return AsAVLTreeNodeEnumerator(r1)
}

func (m *TAVLTreeNodeEnumerator) MoveNext() bool {
	r1 := aVLTreeNodeEnumeratorImportAPI().SysCallN(5, m.Instance())
	return GoBool(r1)
}

func AVLTreeNodeEnumeratorClass() TClass {
	ret := aVLTreeNodeEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	aVLTreeNodeEnumeratorImport       *imports.Imports = nil
	aVLTreeNodeEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("AVLTreeNodeEnumerator_Class", 0),
		/*1*/ imports.NewTable("AVLTreeNodeEnumerator_Create", 0),
		/*2*/ imports.NewTable("AVLTreeNodeEnumerator_Current", 0),
		/*3*/ imports.NewTable("AVLTreeNodeEnumerator_GetEnumerator", 0),
		/*4*/ imports.NewTable("AVLTreeNodeEnumerator_LowToHigh", 0),
		/*5*/ imports.NewTable("AVLTreeNodeEnumerator_MoveNext", 0),
	}
)

func aVLTreeNodeEnumeratorImportAPI() *imports.Imports {
	if aVLTreeNodeEnumeratorImport == nil {
		aVLTreeNodeEnumeratorImport = NewDefaultImports()
		aVLTreeNodeEnumeratorImport.SetImportTable(aVLTreeNodeEnumeratorImportTables)
		aVLTreeNodeEnumeratorImportTables = nil
	}
	return aVLTreeNodeEnumeratorImport
}
