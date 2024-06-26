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

// ICustomFrame Parent: ICustomDesignControl
type ICustomFrame interface {
	ICustomDesignControl
	ParentBackground() bool          // property
	SetParentBackground(AValue bool) // property
}

// TCustomFrame Parent: TCustomDesignControl
type TCustomFrame struct {
	TCustomDesignControl
}

func NewCustomFrame(AOwner IComponent) ICustomFrame {
	r1 := LCL().SysCallN(1735, GetObjectUintptr(AOwner))
	return AsCustomFrame(r1)
}

func (m *TCustomFrame) ParentBackground() bool {
	r1 := LCL().SysCallN(1736, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomFrame) SetParentBackground(AValue bool) {
	LCL().SysCallN(1736, 1, m.Instance(), PascalBool(AValue))
}

func CustomFrameClass() TClass {
	ret := LCL().SysCallN(1734)
	return TClass(ret)
}
