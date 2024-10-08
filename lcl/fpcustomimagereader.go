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

// IFPCustomImageReader Is Abstract Class Parent: IFPCustomImageHandler
type IFPCustomImageReader interface {
	IFPCustomImageHandler
	DefaultImageClass() TFPCustomImageClass                   // property
	SetDefaultImageClass(AValue TFPCustomImageClass)          // property
	ImageRead(Str IStream, Img IFPCustomImage) IFPCustomImage // function
	CheckContents(Str IStream) bool                           // function
}

// TFPCustomImageReader Is Abstract Class Parent: TFPCustomImageHandler
type TFPCustomImageReader struct {
	TFPCustomImageHandler
}

func (m *TFPCustomImageReader) DefaultImageClass() TFPCustomImageClass {
	r1 := LCL().SysCallN(2943, 0, m.Instance(), 0)
	return TFPCustomImageClass(r1)
}

func (m *TFPCustomImageReader) SetDefaultImageClass(AValue TFPCustomImageClass) {
	LCL().SysCallN(2943, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomImageReader) ImageRead(Str IStream, Img IFPCustomImage) IFPCustomImage {
	r1 := LCL().SysCallN(2944, m.Instance(), GetObjectUintptr(Str), GetObjectUintptr(Img))
	return AsFPCustomImage(r1)
}

func (m *TFPCustomImageReader) CheckContents(Str IStream) bool {
	r1 := LCL().SysCallN(2941, m.Instance(), GetObjectUintptr(Str))
	return GoBool(r1)
}

func FPCustomImageReaderClass() TClass {
	ret := LCL().SysCallN(2942)
	return TClass(ret)
}
