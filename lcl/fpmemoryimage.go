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

// IFPMemoryImage Parent: IFPCustomImage
type IFPMemoryImage interface {
	IFPCustomImage
}

// TFPMemoryImage Parent: TFPCustomImage
type TFPMemoryImage struct {
	TFPCustomImage
}

func NewFPMemoryImage(AWidth, AHeight int32) IFPMemoryImage {
	r1 := LCL().SysCallN(3010, uintptr(AWidth), uintptr(AHeight))
	return AsFPMemoryImage(r1)
}

func FPMemoryImageClass() TClass {
	ret := LCL().SysCallN(3009)
	return TClass(ret)
}
