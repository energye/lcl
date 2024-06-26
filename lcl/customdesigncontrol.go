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

// ICustomDesignControl Parent: IScrollingWinControl
type ICustomDesignControl interface {
	IScrollingWinControl
	DesignTimePPI() int32          // property
	SetDesignTimePPI(AValue int32) // property
	PixelsPerInch() int32          // property
	SetPixelsPerInch(AValue int32) // property
	Scaled() bool                  // property
	SetScaled(AValue bool)         // property
}

// TCustomDesignControl Parent: TScrollingWinControl
type TCustomDesignControl struct {
	TScrollingWinControl
}

func NewCustomDesignControl(TheOwner IComponent) ICustomDesignControl {
	r1 := LCL().SysCallN(1504, GetObjectUintptr(TheOwner))
	return AsCustomDesignControl(r1)
}

func (m *TCustomDesignControl) DesignTimePPI() int32 {
	r1 := LCL().SysCallN(1505, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDesignControl) SetDesignTimePPI(AValue int32) {
	LCL().SysCallN(1505, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDesignControl) PixelsPerInch() int32 {
	r1 := LCL().SysCallN(1506, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDesignControl) SetPixelsPerInch(AValue int32) {
	LCL().SysCallN(1506, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDesignControl) Scaled() bool {
	r1 := LCL().SysCallN(1507, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDesignControl) SetScaled(AValue bool) {
	LCL().SysCallN(1507, 1, m.Instance(), PascalBool(AValue))
}

func CustomDesignControlClass() TClass {
	ret := LCL().SysCallN(1503)
	return TClass(ret)
}
