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

// ITimer Parent: ICustomTimer
type ITimer interface {
	ICustomTimer
}

// TTimer Parent: TCustomTimer
type TTimer struct {
	TCustomTimer
}

func NewTimer(AOwner IComponent) ITimer {
	r1 := LCL().SysCallN(5432, GetObjectUintptr(AOwner))
	return AsTimer(r1)
}

func TimerClass() TClass {
	ret := LCL().SysCallN(5431)
	return TClass(ret)
}
