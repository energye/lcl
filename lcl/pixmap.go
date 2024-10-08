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

// IPixmap Parent: IFPImageBitmap
type IPixmap interface {
	IFPImageBitmap
}

// TPixmap Parent: TFPImageBitmap
type TPixmap struct {
	TFPImageBitmap
}

func NewPixmap() IPixmap {
	r1 := LCL().SysCallN(4597)
	return AsPixmap(r1)
}

func PixmapClass() TClass {
	ret := LCL().SysCallN(4596)
	return TClass(ret)
}
