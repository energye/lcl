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

// IWinControlEnumerator Parent: IObject
type IWinControlEnumerator interface {
	IObject
	Current() IControl                    // property
	GetEnumerator() IWinControlEnumerator // function
	MoveNext() bool                       // function
}

// TWinControlEnumerator Parent: TObject
type TWinControlEnumerator struct {
	TObject
}

func NewWinControlEnumerator(Parent IWinControl, aLowToHigh bool) IWinControlEnumerator {
	r1 := LCL().SysCallN(6049, GetObjectUintptr(Parent), PascalBool(aLowToHigh))
	return AsWinControlEnumerator(r1)
}

func (m *TWinControlEnumerator) Current() IControl {
	r1 := LCL().SysCallN(6050, m.Instance())
	return AsControl(r1)
}

func (m *TWinControlEnumerator) GetEnumerator() IWinControlEnumerator {
	r1 := LCL().SysCallN(6051, m.Instance())
	return AsWinControlEnumerator(r1)
}

func (m *TWinControlEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(6052, m.Instance())
	return GoBool(r1)
}

func WinControlEnumeratorClass() TClass {
	ret := LCL().SysCallN(6048)
	return TClass(ret)
}
