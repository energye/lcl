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

// IShortCutList Parent: IStringList
type IShortCutList interface {
	IStringList
	ShortCuts(Index int32) TShortCut          // property
	IndexOfShortCut(Shortcut TShortCut) int32 // function
}

// TShortCutList Parent: TStringList
type TShortCutList struct {
	TStringList
}

func NewShortCutList() IShortCutList {
	r1 := shortCutListImportAPI().SysCallN(1)
	return AsShortCutList(r1)
}

func (m *TShortCutList) ShortCuts(Index int32) TShortCut {
	r1 := shortCutListImportAPI().SysCallN(3, m.Instance(), uintptr(Index))
	return TShortCut(r1)
}

func (m *TShortCutList) IndexOfShortCut(Shortcut TShortCut) int32 {
	r1 := shortCutListImportAPI().SysCallN(2, m.Instance(), uintptr(Shortcut))
	return int32(r1)
}

func ShortCutListClass() TClass {
	ret := shortCutListImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	shortCutListImport       *imports.Imports = nil
	shortCutListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ShortCutList_Class", 0),
		/*1*/ imports.NewTable("ShortCutList_Create", 0),
		/*2*/ imports.NewTable("ShortCutList_IndexOfShortCut", 0),
		/*3*/ imports.NewTable("ShortCutList_ShortCuts", 0),
	}
)

func shortCutListImportAPI() *imports.Imports {
	if shortCutListImport == nil {
		shortCutListImport = NewDefaultImports()
		shortCutListImport.SetImportTable(shortCutListImportTables)
		shortCutListImportTables = nil
	}
	return shortCutListImport
}
