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

// ITreeNodesEnumerator Parent: IObject
type ITreeNodesEnumerator interface {
	IObject
	Current() ITreeNode // property
	MoveNext() bool     // function
}

// TTreeNodesEnumerator Parent: TObject
type TTreeNodesEnumerator struct {
	TObject
}

func NewTreeNodesEnumerator(ANodes ITreeNodes) ITreeNodesEnumerator {
	r1 := reeNodesEnumeratorImportAPI().SysCallN(1, GetObjectUintptr(ANodes))
	return AsTreeNodesEnumerator(r1)
}

func (m *TTreeNodesEnumerator) Current() ITreeNode {
	r1 := reeNodesEnumeratorImportAPI().SysCallN(2, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodesEnumerator) MoveNext() bool {
	r1 := reeNodesEnumeratorImportAPI().SysCallN(3, m.Instance())
	return GoBool(r1)
}

func TreeNodesEnumeratorClass() TClass {
	ret := reeNodesEnumeratorImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	reeNodesEnumeratorImport       *imports.Imports = nil
	reeNodesEnumeratorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TreeNodesEnumerator_Class", 0),
		/*1*/ imports.NewTable("TreeNodesEnumerator_Create", 0),
		/*2*/ imports.NewTable("TreeNodesEnumerator_Current", 0),
		/*3*/ imports.NewTable("TreeNodesEnumerator_MoveNext", 0),
	}
)

func reeNodesEnumeratorImportAPI() *imports.Imports {
	if reeNodesEnumeratorImport == nil {
		reeNodesEnumeratorImport = NewDefaultImports()
		reeNodesEnumeratorImport.SetImportTable(reeNodesEnumeratorImportTables)
		reeNodesEnumeratorImportTables = nil
	}
	return reeNodesEnumeratorImport
}
