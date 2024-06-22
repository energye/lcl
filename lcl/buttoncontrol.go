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

// IButtonControl Parent: IWinControl
type IButtonControl interface {
	IWinControl
}

// TButtonControl Parent: TWinControl
type TButtonControl struct {
	TWinControl
}

func NewButtonControl(TheOwner IComponent) IButtonControl {
	r1 := LCL().SysCallN(475, GetObjectUintptr(TheOwner))
	return AsButtonControl(r1)
}

func ButtonControlClass() TClass {
	ret := LCL().SysCallN(474)
	return TClass(ret)
}
