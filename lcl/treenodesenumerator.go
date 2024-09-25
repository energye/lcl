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
	r1 := LCL().SysCallN(5707, GetObjectUintptr(ANodes))
	return AsTreeNodesEnumerator(r1)
}

func (m *TTreeNodesEnumerator) Current() ITreeNode {
	r1 := LCL().SysCallN(5708, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNodesEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(5709, m.Instance())
	return GoBool(r1)
}

func TreeNodesEnumeratorClass() TClass {
	ret := LCL().SysCallN(5706)
	return TClass(ret)
}
