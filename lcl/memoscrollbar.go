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

// IMemoScrollBar Parent: IControlScrollBar
type IMemoScrollBar interface {
	IControlScrollBar
}

// TMemoScrollBar Parent: TControlScrollBar
type TMemoScrollBar struct {
	TControlScrollBar
}

func NewMemoScrollBar(AControl IWinControl, AKind TScrollBarKind) IMemoScrollBar {
	r1 := memoScrollBarImportAPI().SysCallN(1, GetObjectUintptr(AControl), uintptr(AKind))
	return AsMemoScrollBar(r1)
}

func MemoScrollBarClass() TClass {
	ret := memoScrollBarImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	memoScrollBarImport       *imports.Imports = nil
	memoScrollBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("MemoScrollBar_Class", 0),
		/*1*/ imports.NewTable("MemoScrollBar_Create", 0),
	}
)

func memoScrollBarImportAPI() *imports.Imports {
	if memoScrollBarImport == nil {
		memoScrollBarImport = NewDefaultImports()
		memoScrollBarImport.SetImportTable(memoScrollBarImportTables)
		memoScrollBarImportTables = nil
	}
	return memoScrollBarImport
}
