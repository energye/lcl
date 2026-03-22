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

// ISynMultiSyn Parent: ISynCustomHighlighter
type ISynMultiSyn interface {
	ISynCustomHighlighter
	UpdateRangeInfoAtLine(index int32) bool            // function
	DefaultVirtualLines() ISynHLightMultiVirtualLines  // property DefaultVirtualLines Getter
	Schemes() ISynHighlighterMultiSchemeList           // property Schemes Getter
	SetSchemes(value ISynHighlighterMultiSchemeList)   // property Schemes Setter
	DefaultHighlighter() ISynCustomHighlighter         // property DefaultHighlighter Getter
	SetDefaultHighlighter(value ISynCustomHighlighter) // property DefaultHighlighter Setter
	DefaultLanguageName() string                       // property DefaultLanguageName Getter
	SetDefaultLanguageName(value string)               // property DefaultLanguageName Setter
}

type TSynMultiSyn struct {
	TSynCustomHighlighter
}

func (m *TSynMultiSyn) UpdateRangeInfoAtLine(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := synMultiSynAPI().SysCallN(1, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TSynMultiSyn) DefaultVirtualLines() ISynHLightMultiVirtualLines {
	if !m.IsValid() {
		return nil
	}
	r := synMultiSynAPI().SysCallN(2, m.Instance())
	return AsSynHLightMultiVirtualLines(r)
}

func (m *TSynMultiSyn) Schemes() ISynHighlighterMultiSchemeList {
	if !m.IsValid() {
		return nil
	}
	r := synMultiSynAPI().SysCallN(3, 0, m.Instance())
	return AsSynHighlighterMultiSchemeList(r)
}

func (m *TSynMultiSyn) SetSchemes(value ISynHighlighterMultiSchemeList) {
	if !m.IsValid() {
		return
	}
	synMultiSynAPI().SysCallN(3, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynMultiSyn) DefaultHighlighter() ISynCustomHighlighter {
	if !m.IsValid() {
		return nil
	}
	r := synMultiSynAPI().SysCallN(4, 0, m.Instance())
	return AsSynCustomHighlighter(r)
}

func (m *TSynMultiSyn) SetDefaultHighlighter(value ISynCustomHighlighter) {
	if !m.IsValid() {
		return
	}
	synMultiSynAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynMultiSyn) DefaultLanguageName() string {
	if !m.IsValid() {
		return ""
	}
	r := synMultiSynAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TSynMultiSyn) SetDefaultLanguageName(value string) {
	if !m.IsValid() {
		return
	}
	synMultiSynAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

// NewSynMultiSyn class constructor
func NewSynMultiSyn(owner IComponent) ISynMultiSyn {
	r := synMultiSynAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynMultiSyn(r)
}

func TSynMultiSynClass() types.TClass {
	r := synMultiSynAPI().SysCallN(6)
	return types.TClass(r)
}

var (
	synMultiSynOnce   base.Once
	synMultiSynImport *imports.Imports = nil
)

func synMultiSynAPI() *imports.Imports {
	synMultiSynOnce.Do(func() {
		synMultiSynImport = api.NewDefaultImports()
		synMultiSynImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynMultiSyn_Create", 0), // constructor NewSynMultiSyn
			/* 1 */ imports.NewTable("TSynMultiSyn_UpdateRangeInfoAtLine", 0), // function UpdateRangeInfoAtLine
			/* 2 */ imports.NewTable("TSynMultiSyn_DefaultVirtualLines", 0), // property DefaultVirtualLines
			/* 3 */ imports.NewTable("TSynMultiSyn_Schemes", 0), // property Schemes
			/* 4 */ imports.NewTable("TSynMultiSyn_DefaultHighlighter", 0), // property DefaultHighlighter
			/* 5 */ imports.NewTable("TSynMultiSyn_DefaultLanguageName", 0), // property DefaultLanguageName
			/* 6 */ imports.NewTable("TSynMultiSyn_TClass", 0), // function TSynMultiSynClass
		}
	})
	return synMultiSynImport
}
