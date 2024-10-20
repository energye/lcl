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
	"github.com/energye/lcl/api/imports"
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
	r1 := imerImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsTimer(r1)
}

func TimerClass() TClass {
	ret := imerImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	imerImport       *imports.Imports = nil
	imerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Timer_Class", 0),
		/*1*/ imports.NewTable("Timer_Create", 0),
	}
)

func imerImportAPI() *imports.Imports {
	if imerImport == nil {
		imerImport = NewDefaultImports()
		imerImport.SetImportTable(imerImportTables)
		imerImportTables = nil
	}
	return imerImport
}
