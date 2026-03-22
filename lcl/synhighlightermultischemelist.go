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
)

// ISynHighlighterMultiSchemeList Parent: ICollection
type ISynHighlighterMultiSchemeList interface {
	ICollection
	IndexOf(anItem ISynHighlighterMultiScheme) int32                                          // function
	ItemsWithIntToSynHighlighterMultiScheme(index int32) ISynHighlighterMultiScheme           // property Items Getter
	SetItemsWithIntToSynHighlighterMultiScheme(index int32, value ISynHighlighterMultiScheme) // property Items Setter
	ConvertedCurrentLine() string                                                             // property ConvertedCurrentLine Getter
	CurrentLine() string                                                                      // property CurrentLine Getter
	SetCurrentLine(value string)                                                              // property CurrentLine Setter
	OwnerToSynMultiSyn() ISynMultiSyn                                                         // property Owner Getter
}

type TSynHighlighterMultiSchemeList struct {
	TCollection
}

func (m *TSynHighlighterMultiSchemeList) IndexOf(anItem ISynHighlighterMultiScheme) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterMultiSchemeListAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(anItem))
	return int32(r)
}

func (m *TSynHighlighterMultiSchemeList) ItemsWithIntToSynHighlighterMultiScheme(index int32) ISynHighlighterMultiScheme {
	if !m.IsValid() {
		return nil
	}
	r := synHighlighterMultiSchemeListAPI().SysCallN(2, 0, m.Instance(), uintptr(index))
	return AsSynHighlighterMultiScheme(r)
}

func (m *TSynHighlighterMultiSchemeList) SetItemsWithIntToSynHighlighterMultiScheme(index int32, value ISynHighlighterMultiScheme) {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeListAPI().SysCallN(2, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TSynHighlighterMultiSchemeList) ConvertedCurrentLine() string {
	if !m.IsValid() {
		return ""
	}
	r := synHighlighterMultiSchemeListAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TSynHighlighterMultiSchemeList) CurrentLine() string {
	if !m.IsValid() {
		return ""
	}
	r := synHighlighterMultiSchemeListAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TSynHighlighterMultiSchemeList) SetCurrentLine(value string) {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeListAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynHighlighterMultiSchemeList) OwnerToSynMultiSyn() ISynMultiSyn {
	if !m.IsValid() {
		return nil
	}
	r := synHighlighterMultiSchemeListAPI().SysCallN(5, m.Instance())
	return AsSynMultiSyn(r)
}

// NewSynHighlighterMultiSchemeList class constructor
func NewSynHighlighterMultiSchemeList(owner ISynMultiSyn) ISynHighlighterMultiSchemeList {
	r := synHighlighterMultiSchemeListAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynHighlighterMultiSchemeList(r)
}

var (
	synHighlighterMultiSchemeListOnce   base.Once
	synHighlighterMultiSchemeListImport *imports.Imports = nil
)

func synHighlighterMultiSchemeListAPI() *imports.Imports {
	synHighlighterMultiSchemeListOnce.Do(func() {
		synHighlighterMultiSchemeListImport = api.NewDefaultImports()
		synHighlighterMultiSchemeListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynHighlighterMultiSchemeList_Create", 0), // constructor NewSynHighlighterMultiSchemeList
			/* 1 */ imports.NewTable("TSynHighlighterMultiSchemeList_IndexOf", 0), // function IndexOf
			/* 2 */ imports.NewTable("TSynHighlighterMultiSchemeList_ItemsWithIntToSynHighlighterMultiScheme", 0), // property ItemsWithIntToSynHighlighterMultiScheme
			/* 3 */ imports.NewTable("TSynHighlighterMultiSchemeList_ConvertedCurrentLine", 0), // property ConvertedCurrentLine
			/* 4 */ imports.NewTable("TSynHighlighterMultiSchemeList_CurrentLine", 0), // property CurrentLine
			/* 5 */ imports.NewTable("TSynHighlighterMultiSchemeList_OwnerToSynMultiSyn", 0), // property OwnerToSynMultiSyn
		}
	})
	return synHighlighterMultiSchemeListImport
}
