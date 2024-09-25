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

// IFPImageBitmap Parent: ICustomBitmap
type IFPImageBitmap interface {
	ICustomBitmap
}

// TFPImageBitmap Parent: TCustomBitmap
type TFPImageBitmap struct {
	TCustomBitmap
}

func NewFPImageBitmap() IFPImageBitmap {
	r1 := LCL().SysCallN(2981)
	return AsFPImageBitmap(r1)
}

func FPImageBitmapClass() TClass {
	ret := LCL().SysCallN(2980)
	return TClass(ret)
}
