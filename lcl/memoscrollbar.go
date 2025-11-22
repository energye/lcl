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

// IMemoScrollbar Parent: IControlScrollBar
type IMemoScrollbar interface {
	IControlScrollBar
}

type TMemoScrollbar struct {
	TControlScrollBar
}

// NewMemoScrollbar class constructor
func NewMemoScrollbar(control IWinControl, kind types.TScrollBarKind) IMemoScrollbar {
	r := memoScrollbarAPI().SysCallN(0, base.GetObjectUintptr(control), uintptr(kind))
	return AsMemoScrollbar(r)
}

var (
	memoScrollbarOnce   base.Once
	memoScrollbarImport *imports.Imports = nil
)

func memoScrollbarAPI() *imports.Imports {
	memoScrollbarOnce.Do(func() {
		memoScrollbarImport = api.NewDefaultImports()
		memoScrollbarImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMemoScrollbar_Create", 0), // constructor NewMemoScrollbar
		}
	})
	return memoScrollbarImport
}
