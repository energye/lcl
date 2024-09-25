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

// IFPCustomImageWriter Is Abstract Class Parent: IFPCustomImageHandler
type IFPCustomImageWriter interface {
	IFPCustomImageHandler
	ImageWrite(Str IStream, Img IFPCustomImage) // procedure
}

// TFPCustomImageWriter Is Abstract Class Parent: TFPCustomImageHandler
type TFPCustomImageWriter struct {
	TFPCustomImageHandler
}

func FPCustomImageWriterClass() TClass {
	ret := LCL().SysCallN(2945)
	return TClass(ret)
}

func (m *TFPCustomImageWriter) ImageWrite(Str IStream, Img IFPCustomImage) {
	LCL().SysCallN(2946, m.Instance(), GetObjectUintptr(Str), GetObjectUintptr(Img))
}
