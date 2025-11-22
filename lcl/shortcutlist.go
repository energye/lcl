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

// IShortCutList Parent: IStringList
type IShortCutList interface {
	IStringList
	IndexOfShortCut(shortcut types.TShortCut) int32 // function
	ShortCuts(index int32) types.TShortCut          // property ShortCuts Getter
}

type TShortCutList struct {
	TStringList
}

func (m *TShortCutList) IndexOfShortCut(shortcut types.TShortCut) int32 {
	if !m.IsValid() {
		return 0
	}
	r := shortCutListAPI().SysCallN(1, m.Instance(), uintptr(shortcut))
	return int32(r)
}

func (m *TShortCutList) ShortCuts(index int32) types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := shortCutListAPI().SysCallN(2, m.Instance(), uintptr(index))
	return types.TShortCut(r)
}

// NewShortCutList class constructor
func NewShortCutList() IShortCutList {
	r := shortCutListAPI().SysCallN(0)
	return AsShortCutList(r)
}

var (
	shortCutListOnce   base.Once
	shortCutListImport *imports.Imports = nil
)

func shortCutListAPI() *imports.Imports {
	shortCutListOnce.Do(func() {
		shortCutListImport = api.NewDefaultImports()
		shortCutListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TShortCutList_Create", 0), // constructor NewShortCutList
			/* 1 */ imports.NewTable("TShortCutList_IndexOfShortCut", 0), // function IndexOfShortCut
			/* 2 */ imports.NewTable("TShortCutList_ShortCuts", 0), // property ShortCuts
		}
	})
	return shortCutListImport
}
