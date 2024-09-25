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

// IMergedMenuItems Parent: IObject
type IMergedMenuItems interface {
	IObject
	VisibleCount() int32                  // property
	VisibleItems(Index int32) IMenuItem   // property
	InvisibleCount() int32                // property
	InvisibleItems(Index int32) IMenuItem // property
}

// TMergedMenuItems Parent: TObject
type TMergedMenuItems struct {
	TObject
}

func NewMergedMenuItems(aParent IMenuItem) IMergedMenuItems {
	r1 := LCL().SysCallN(4362, GetObjectUintptr(aParent))
	return AsMergedMenuItems(r1)
}

func (m *TMergedMenuItems) VisibleCount() int32 {
	r1 := LCL().SysCallN(4365, m.Instance())
	return int32(r1)
}

func (m *TMergedMenuItems) VisibleItems(Index int32) IMenuItem {
	r1 := LCL().SysCallN(4366, m.Instance(), uintptr(Index))
	return AsMenuItem(r1)
}

func (m *TMergedMenuItems) InvisibleCount() int32 {
	r1 := LCL().SysCallN(4363, m.Instance())
	return int32(r1)
}

func (m *TMergedMenuItems) InvisibleItems(Index int32) IMenuItem {
	r1 := LCL().SysCallN(4364, m.Instance(), uintptr(Index))
	return AsMenuItem(r1)
}

func MergedMenuItemsClass() TClass {
	ret := LCL().SysCallN(4361)
	return TClass(ret)
}
