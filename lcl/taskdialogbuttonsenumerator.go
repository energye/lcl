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

// ITaskDialogButtonsEnumerator Parent: IObject
type ITaskDialogButtonsEnumerator interface {
	IObject
	Current() ITaskDialogBaseButtonItem    // property
	GetCurrent() ITaskDialogBaseButtonItem // function
	MoveNext() bool                        // function
}

// TTaskDialogButtonsEnumerator Parent: TObject
type TTaskDialogButtonsEnumerator struct {
	TObject
}

func NewTaskDialogButtonsEnumerator(ACollection ITaskDialogButtons) ITaskDialogButtonsEnumerator {
	r1 := LCL().SysCallN(5395, GetObjectUintptr(ACollection))
	return AsTaskDialogButtonsEnumerator(r1)
}

func (m *TTaskDialogButtonsEnumerator) Current() ITaskDialogBaseButtonItem {
	r1 := LCL().SysCallN(5396, m.Instance())
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtonsEnumerator) GetCurrent() ITaskDialogBaseButtonItem {
	r1 := LCL().SysCallN(5397, m.Instance())
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtonsEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(5398, m.Instance())
	return GoBool(r1)
}

func TaskDialogButtonsEnumeratorClass() TClass {
	ret := LCL().SysCallN(5394)
	return TClass(ret)
}
