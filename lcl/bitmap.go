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

// IBitmap Parent: IFPImageBitmap
type IBitmap interface {
	IFPImageBitmap
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
}

// TBitmap Parent: TFPImageBitmap
type TBitmap struct {
	TFPImageBitmap
}

func NewBitmap() IBitmap {
	r1 := LCL().SysCallN(444)
	return AsBitmap(r1)
}

func BitmapClass() TClass {
	ret := LCL().SysCallN(443)
	return TClass(ret)
}
