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

// IIcon Parent: ICustomIcon
type IIcon interface {
	ICustomIcon
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Handle() HICON          // property
	SetHandle(AValue HICON) // property
	ReleaseHandle() HICON   // function
}

// TIcon Parent: TCustomIcon
type TIcon struct {
	TCustomIcon
}

func NewIcon() IIcon {
	r1 := LCL().SysCallN(3378)
	return AsIcon(r1)
}

func (m *TIcon) Handle() HICON {
	r1 := LCL().SysCallN(3379, 0, m.Instance(), 0)
	return HICON(r1)
}

func (m *TIcon) SetHandle(AValue HICON) {
	LCL().SysCallN(3379, 1, m.Instance(), uintptr(AValue))
}

func (m *TIcon) ReleaseHandle() HICON {
	r1 := LCL().SysCallN(3380, m.Instance())
	return HICON(r1)
}

func IconClass() TClass {
	ret := LCL().SysCallN(3377)
	return TClass(ret)
}
