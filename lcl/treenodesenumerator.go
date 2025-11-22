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

// ITreeNodesEnumerator Parent: IObject
type ITreeNodesEnumerator interface {
	IObject
	MoveNext() bool     // function
	Current() ITreeNode // property Current Getter
}

type TTreeNodesEnumerator struct {
	TObject
}

func (m *TTreeNodesEnumerator) MoveNext() bool {
	if !m.IsValid() {
		return false
	}
	r := treeNodesEnumeratorAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TTreeNodesEnumerator) Current() ITreeNode {
	if !m.IsValid() {
		return nil
	}
	r := treeNodesEnumeratorAPI().SysCallN(2, m.Instance())
	return AsTreeNode(r)
}

// NewTreeNodesEnumerator class constructor
func NewTreeNodesEnumerator(nodes ITreeNodes) ITreeNodesEnumerator {
	r := treeNodesEnumeratorAPI().SysCallN(0, base.GetObjectUintptr(nodes))
	return AsTreeNodesEnumerator(r)
}

var (
	treeNodesEnumeratorOnce   base.Once
	treeNodesEnumeratorImport *imports.Imports = nil
)

func treeNodesEnumeratorAPI() *imports.Imports {
	treeNodesEnumeratorOnce.Do(func() {
		treeNodesEnumeratorImport = api.NewDefaultImports()
		treeNodesEnumeratorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTreeNodesEnumerator_Create", 0), // constructor NewTreeNodesEnumerator
			/* 1 */ imports.NewTable("TTreeNodesEnumerator_MoveNext", 0), // function MoveNext
			/* 2 */ imports.NewTable("TTreeNodesEnumerator_Current", 0), // property Current
		}
	})
	return treeNodesEnumeratorImport
}
