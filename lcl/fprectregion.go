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

// IFPRectRegion Parent: IFPCustomRegion
type IFPRectRegion interface {
	IFPCustomRegion
}

// TFPRectRegion Parent: TFPCustomRegion
type TFPRectRegion struct {
	TFPCustomRegion
}

func NewFPRectRegion() IFPRectRegion {
	r1 := LCL().SysCallN(2994)
	return AsFPRectRegion(r1)
}

func FPRectRegionClass() TClass {
	ret := LCL().SysCallN(2993)
	return TClass(ret)
}
