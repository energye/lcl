//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// ITimer Parent: ICustomTimer
type ITimer interface {
	ICustomTimer
}

type TTimer struct {
	TCustomTimer
}

// NewTimer class constructor
func NewTimer(owner IComponent) ITimer {
	r := timerAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsTimer(r)
}

func TTimerClass() types.TClass {
	r := timerAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	timerOnce   base.Once
	timerImport *imports.Imports = nil
)

func timerAPI() *imports.Imports {
	timerOnce.Do(func() {
		timerImport = api.NewDefaultImports()
		timerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTimer_Create", 0), // constructor NewTimer
			/* 1 */ imports.NewTable("TTimer_TClass", 0), // function TTimerClass
		}
	})
	return timerImport
}
