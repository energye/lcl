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

// ILazSynDisplayView Parent: IObject
type ILazSynDisplayView interface {
	IObject
	GetNextHighlighterToken(outTokenInfo *TLazSynDisplayTokenInfo) bool                                                       // function
	GetLinesCount() int32                                                                                                     // function
	GetDrawDividerInfo() TSynDividerDrawConfigSetting                                                                         // function
	TextToViewIndex(textIndex types.TLineIdx) TLineRange                                                                      // function
	ViewToTextIndex(viewIndex types.TLineIdx) types.TLineIdx                                                                  // function
	ViewToTextIndexEx(viewIndex types.TLineIdx, outViewRange *TLineRange) types.TLineIdx                                      // function
	InitHighlighterTokens(highlighter ISynCustomHighlighter)                                                                  // procedure
	SetHighlighterTokensLine(line types.TLineIdx, outRealLine *types.TLineIdx, outStartBytePos *int32, outLineByteLen *int32) // procedure
	FinishHighlighterTokens()                                                                                                 // procedure
	NextView() ILazSynDisplayView                                                                                             // property NextView Getter
	SetNextView(value ILazSynDisplayView)                                                                                     // property NextView Setter
}

type TLazSynDisplayView struct {
	TObject
}

func (m *TLazSynDisplayView) GetNextHighlighterToken(outTokenInfo *TLazSynDisplayTokenInfo) bool {
	if !m.IsValid() {
		return false
	}
	tokenInfoPtr := &tLazSynDisplayTokenInfo{}
	r := lazSynDisplayViewAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(tokenInfoPtr)))
	*outTokenInfo = tokenInfoPtr.ToGo()
	return api.GoBool(r)
}

func (m *TLazSynDisplayView) GetLinesCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynDisplayViewAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TLazSynDisplayView) GetDrawDividerInfo() (result TSynDividerDrawConfigSetting) {
	if !m.IsValid() {
		return
	}
	lazSynDisplayViewAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynDisplayView) TextToViewIndex(textIndex types.TLineIdx) (result TLineRange) {
	if !m.IsValid() {
		return
	}
	lazSynDisplayViewAPI().SysCallN(4, m.Instance(), uintptr(textIndex), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynDisplayView) ViewToTextIndex(viewIndex types.TLineIdx) types.TLineIdx {
	if !m.IsValid() {
		return 0
	}
	r := lazSynDisplayViewAPI().SysCallN(5, m.Instance(), uintptr(viewIndex))
	return types.TLineIdx(r)
}

func (m *TLazSynDisplayView) ViewToTextIndexEx(viewIndex types.TLineIdx, outViewRange *TLineRange) types.TLineIdx {
	if !m.IsValid() {
		return 0
	}
	r := lazSynDisplayViewAPI().SysCallN(6, m.Instance(), uintptr(viewIndex), uintptr(base.UnsafePointer(outViewRange)))
	return types.TLineIdx(r)
}

func (m *TLazSynDisplayView) InitHighlighterTokens(highlighter ISynCustomHighlighter) {
	if !m.IsValid() {
		return
	}
	lazSynDisplayViewAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(highlighter))
}

func (m *TLazSynDisplayView) SetHighlighterTokensLine(line types.TLineIdx, outRealLine *types.TLineIdx, outStartBytePos *int32, outLineByteLen *int32) {
	if !m.IsValid() {
		return
	}
	var realLinePtr uintptr
	var startBytePosPtr uintptr
	var lineByteLenPtr uintptr
	lazSynDisplayViewAPI().SysCallN(8, m.Instance(), uintptr(line), uintptr(base.UnsafePointer(&realLinePtr)), uintptr(base.UnsafePointer(&startBytePosPtr)), uintptr(base.UnsafePointer(&lineByteLenPtr)))
	*outRealLine = types.TLineIdx(realLinePtr)
	*outStartBytePos = int32(startBytePosPtr)
	*outLineByteLen = int32(lineByteLenPtr)
}

func (m *TLazSynDisplayView) FinishHighlighterTokens() {
	if !m.IsValid() {
		return
	}
	lazSynDisplayViewAPI().SysCallN(9, m.Instance())
}

func (m *TLazSynDisplayView) NextView() ILazSynDisplayView {
	if !m.IsValid() {
		return nil
	}
	r := lazSynDisplayViewAPI().SysCallN(10, 0, m.Instance())
	return AsLazSynDisplayView(r)
}

func (m *TLazSynDisplayView) SetNextView(value ILazSynDisplayView) {
	if !m.IsValid() {
		return
	}
	lazSynDisplayViewAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewLazSynDisplayView class constructor
func NewLazSynDisplayView() ILazSynDisplayView {
	r := lazSynDisplayViewAPI().SysCallN(0)
	return AsLazSynDisplayView(r)
}

var (
	lazSynDisplayViewOnce   base.Once
	lazSynDisplayViewImport *imports.Imports = nil
)

func lazSynDisplayViewAPI() *imports.Imports {
	lazSynDisplayViewOnce.Do(func() {
		lazSynDisplayViewImport = api.NewDefaultImports()
		lazSynDisplayViewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazSynDisplayView_Create", 0), // constructor NewLazSynDisplayView
			/* 1 */ imports.NewTable("TLazSynDisplayView_GetNextHighlighterToken", 0), // function GetNextHighlighterToken
			/* 2 */ imports.NewTable("TLazSynDisplayView_GetLinesCount", 0), // function GetLinesCount
			/* 3 */ imports.NewTable("TLazSynDisplayView_GetDrawDividerInfo", 0), // function GetDrawDividerInfo
			/* 4 */ imports.NewTable("TLazSynDisplayView_TextToViewIndex", 0), // function TextToViewIndex
			/* 5 */ imports.NewTable("TLazSynDisplayView_ViewToTextIndex", 0), // function ViewToTextIndex
			/* 6 */ imports.NewTable("TLazSynDisplayView_ViewToTextIndexEx", 0), // function ViewToTextIndexEx
			/* 7 */ imports.NewTable("TLazSynDisplayView_InitHighlighterTokens", 0), // procedure InitHighlighterTokens
			/* 8 */ imports.NewTable("TLazSynDisplayView_SetHighlighterTokensLine", 0), // procedure SetHighlighterTokensLine
			/* 9 */ imports.NewTable("TLazSynDisplayView_FinishHighlighterTokens", 0), // procedure FinishHighlighterTokens
			/* 10 */ imports.NewTable("TLazSynDisplayView_NextView", 0), // property NextView
		}
	})
	return lazSynDisplayViewImport
}
