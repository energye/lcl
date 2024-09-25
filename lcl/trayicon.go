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

// ITrayIcon Parent: ICustomTrayIcon
type ITrayIcon interface {
	ICustomTrayIcon
}

// TTrayIcon Parent: TCustomTrayIcon
type TTrayIcon struct {
	TCustomTrayIcon
}

func NewTrayIcon(TheOwner IComponent) ITrayIcon {
	r1 := LCL().SysCallN(5614, GetObjectUintptr(TheOwner))
	return AsTrayIcon(r1)
}

func TrayIconClass() TClass {
	ret := LCL().SysCallN(5613)
	return TClass(ret)
}
