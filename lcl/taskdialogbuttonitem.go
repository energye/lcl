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

// ITaskDialogButtonItem Parent: ITaskDialogBaseButtonItem
type ITaskDialogButtonItem interface {
	ITaskDialogBaseButtonItem
}

// TTaskDialogButtonItem Parent: TTaskDialogBaseButtonItem
type TTaskDialogButtonItem struct {
	TTaskDialogBaseButtonItem
}

func NewTaskDialogButtonItem(ACollection ICollection) ITaskDialogButtonItem {
	r1 := LCL().SysCallN(5393, GetObjectUintptr(ACollection))
	return AsTaskDialogButtonItem(r1)
}

func TaskDialogButtonItemClass() TClass {
	ret := LCL().SysCallN(5392)
	return TClass(ret)
}
