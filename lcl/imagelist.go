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

// IImageList Parent: IDragImageList
type IImageList interface {
	IDragImageList
}

// TImageList Parent: TDragImageList
type TImageList struct {
	TDragImageList
}

func NewImageList(AOwner IComponent) IImageList {
	r1 := LCL().SysCallN(3379, GetObjectUintptr(AOwner))
	return AsImageList(r1)
}

func ImageListClass() TClass {
	ret := LCL().SysCallN(3378)
	return TClass(ret)
}
