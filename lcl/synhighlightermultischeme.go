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

// ISynHighlighterMultiScheme Parent: ICollectionItem
type ISynHighlighterMultiScheme interface {
	ICollectionItem
	FindStartPosInLine(searchPos int32) int32          // function
	FindEndPosInLine(searchPos int32) int32            // function
	ClearLinesSet()                                    // procedure
	LastMatchLen() int32                               // property LastMatchLen Getter
	NeedHLScan() bool                                  // property NeedHLScan Getter
	VirtualLines() ISynHLightMultiVirtualLines         // property VirtualLines Getter
	SetVirtualLines(value ISynHLightMultiVirtualLines) // property VirtualLines Setter
	CaseSensitive() bool                               // property CaseSensitive Getter
	SetCaseSensitive(value bool)                       // property CaseSensitive Setter
	StartExpr() string                                 // property StartExpr Getter
	SetStartExpr(value string)                         // property StartExpr Setter
	EndExpr() string                                   // property EndExpr Getter
	SetEndExpr(value string)                           // property EndExpr Setter
	Highlighter() ISynCustomHighlighter                // property Highlighter Getter
	SetHighlighter(value ISynCustomHighlighter)        // property Highlighter Setter
	MarkerAttri() ISynHighlighterAttributes            // property MarkerAttri Getter
	SetMarkerAttri(value ISynHighlighterAttributes)    // property MarkerAttri Setter
	SchemeName() string                                // property SchemeName Getter
	SetSchemeName(value string)                        // property SchemeName Setter
	SetOnCheckStartMarker(fn TOnCheckMarker)           // property event
	SetOnCheckEndMarker(fn TOnCheckMarker)             // property event
}

type TSynHighlighterMultiScheme struct {
	TCollectionItem
}

func (m *TSynHighlighterMultiScheme) FindStartPosInLine(searchPos int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterMultiSchemeAPI().SysCallN(1, m.Instance(), uintptr(searchPos))
	return int32(r)
}

func (m *TSynHighlighterMultiScheme) FindEndPosInLine(searchPos int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterMultiSchemeAPI().SysCallN(2, m.Instance(), uintptr(searchPos))
	return int32(r)
}

func (m *TSynHighlighterMultiScheme) ClearLinesSet() {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeAPI().SysCallN(3, m.Instance())
}

func (m *TSynHighlighterMultiScheme) LastMatchLen() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synHighlighterMultiSchemeAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TSynHighlighterMultiScheme) NeedHLScan() bool {
	if !m.IsValid() {
		return false
	}
	r := synHighlighterMultiSchemeAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TSynHighlighterMultiScheme) VirtualLines() ISynHLightMultiVirtualLines {
	if !m.IsValid() {
		return nil
	}
	r := synHighlighterMultiSchemeAPI().SysCallN(6, 0, m.Instance())
	return AsSynHLightMultiVirtualLines(r)
}

func (m *TSynHighlighterMultiScheme) SetVirtualLines(value ISynHLightMultiVirtualLines) {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHighlighterMultiScheme) CaseSensitive() bool {
	if !m.IsValid() {
		return false
	}
	r := synHighlighterMultiSchemeAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynHighlighterMultiScheme) SetCaseSensitive(value bool) {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynHighlighterMultiScheme) StartExpr() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synHighlighterMultiSchemeAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynHighlighterMultiScheme) SetStartExpr(value string) {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeAPI().SysCallN(8, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynHighlighterMultiScheme) EndExpr() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synHighlighterMultiSchemeAPI().SysCallN(9, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynHighlighterMultiScheme) SetEndExpr(value string) {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynHighlighterMultiScheme) Highlighter() ISynCustomHighlighter {
	if !m.IsValid() {
		return nil
	}
	r := synHighlighterMultiSchemeAPI().SysCallN(10, 0, m.Instance())
	return AsSynCustomHighlighter(r)
}

func (m *TSynHighlighterMultiScheme) SetHighlighter(value ISynCustomHighlighter) {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHighlighterMultiScheme) MarkerAttri() ISynHighlighterAttributes {
	if !m.IsValid() {
		return nil
	}
	r := synHighlighterMultiSchemeAPI().SysCallN(11, 0, m.Instance())
	return AsSynHighlighterAttributes(r)
}

func (m *TSynHighlighterMultiScheme) SetMarkerAttri(value ISynHighlighterAttributes) {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynHighlighterMultiScheme) SchemeName() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synHighlighterMultiSchemeAPI().SysCallN(12, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynHighlighterMultiScheme) SetSchemeName(value string) {
	if !m.IsValid() {
		return
	}
	synHighlighterMultiSchemeAPI().SysCallN(12, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynHighlighterMultiScheme) SetOnCheckStartMarker(fn TOnCheckMarker) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCheckMarker(fn)
	base.SetEvent(m, 13, synHighlighterMultiSchemeAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynHighlighterMultiScheme) SetOnCheckEndMarker(fn TOnCheckMarker) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnCheckMarker(fn)
	base.SetEvent(m, 14, synHighlighterMultiSchemeAPI(), api.MakeEventDataPtr(cb))
}

// NewSynHighlighterMultiScheme class constructor
func NewSynHighlighterMultiScheme(theCollection ICollection) ISynHighlighterMultiScheme {
	r := synHighlighterMultiSchemeAPI().SysCallN(0, base.GetObjectUintptr(theCollection))
	return AsSynHighlighterMultiScheme(r)
}

var (
	synHighlighterMultiSchemeOnce   base.Once
	synHighlighterMultiSchemeImport *imports.Imports = nil
)

func synHighlighterMultiSchemeAPI() *imports.Imports {
	synHighlighterMultiSchemeOnce.Do(func() {
		synHighlighterMultiSchemeImport = api.NewDefaultImports()
		synHighlighterMultiSchemeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynHighlighterMultiScheme_Create", 0), // constructor NewSynHighlighterMultiScheme
			/* 1 */ imports.NewTable("TSynHighlighterMultiScheme_FindStartPosInLine", 0), // function FindStartPosInLine
			/* 2 */ imports.NewTable("TSynHighlighterMultiScheme_FindEndPosInLine", 0), // function FindEndPosInLine
			/* 3 */ imports.NewTable("TSynHighlighterMultiScheme_ClearLinesSet", 0), // procedure ClearLinesSet
			/* 4 */ imports.NewTable("TSynHighlighterMultiScheme_LastMatchLen", 0), // property LastMatchLen
			/* 5 */ imports.NewTable("TSynHighlighterMultiScheme_NeedHLScan", 0), // property NeedHLScan
			/* 6 */ imports.NewTable("TSynHighlighterMultiScheme_VirtualLines", 0), // property VirtualLines
			/* 7 */ imports.NewTable("TSynHighlighterMultiScheme_CaseSensitive", 0), // property CaseSensitive
			/* 8 */ imports.NewTable("TSynHighlighterMultiScheme_StartExpr", 0), // property StartExpr
			/* 9 */ imports.NewTable("TSynHighlighterMultiScheme_EndExpr", 0), // property EndExpr
			/* 10 */ imports.NewTable("TSynHighlighterMultiScheme_Highlighter", 0), // property Highlighter
			/* 11 */ imports.NewTable("TSynHighlighterMultiScheme_MarkerAttri", 0), // property MarkerAttri
			/* 12 */ imports.NewTable("TSynHighlighterMultiScheme_SchemeName", 0), // property SchemeName
			/* 13 */ imports.NewTable("TSynHighlighterMultiScheme_OnCheckStartMarker", 0), // event OnCheckStartMarker
			/* 14 */ imports.NewTable("TSynHighlighterMultiScheme_OnCheckEndMarker", 0), // event OnCheckEndMarker
		}
	})
	return synHighlighterMultiSchemeImport
}
