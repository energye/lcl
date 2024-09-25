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

// IFPCustomRegion Is Abstract Class Parent: IObject
type IFPCustomRegion interface {
	IObject
	GetBoundingRect() (resultRect TRect) // function Is Abstract
	IsPointInRegion(AX, AY int32) bool   // function Is Abstract
}

// TFPCustomRegion Is Abstract Class Parent: TObject
type TFPCustomRegion struct {
	TObject
}

func (m *TFPCustomRegion) GetBoundingRect() (resultRect TRect) {
	LCL().SysCallN(2978, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TFPCustomRegion) IsPointInRegion(AX, AY int32) bool {
	r1 := LCL().SysCallN(2979, m.Instance(), uintptr(AX), uintptr(AY))
	return GoBool(r1)
}

func FPCustomRegionClass() TClass {
	ret := LCL().SysCallN(2977)
	return TClass(ret)
}
