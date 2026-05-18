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

// ISynAutoComplete Parent: ILazSynMultiEditPlugin
type ISynAutoComplete interface {
	ILazSynMultiEditPlugin
	EditorsCount() int32                            // function
	GetTokenList() string                           // function
	GetTokenValue(token string) string              // function
	Execute(token string, editor ICustomSynEdit)    // procedure
	AutoCompleteList() IStrings                     // property AutoCompleteList Getter
	SetAutoCompleteList(value IStrings)             // property AutoCompleteList Setter
	EndOfTokenChr() string                          // property EndOfTokenChr Getter
	SetEndOfTokenChr(value string)                  // property EndOfTokenChr Setter
	ShortCut() types.TShortCut                      // property ShortCut Getter
	SetShortCut(value types.TShortCut)              // property ShortCut Setter
	ExecCommandID() types.TSynEditorCommand         // property ExecCommandID Getter
	SetExecCommandID(value types.TSynEditorCommand) // property ExecCommandID Setter
}

type TSynAutoComplete struct {
	TLazSynMultiEditPlugin
}

func (m *TSynAutoComplete) EditorsCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synAutoCompleteAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TSynAutoComplete) GetTokenList() (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	synAutoCompleteAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynAutoComplete) GetTokenValue(token string) (result string) {
	if !m.IsValid() {
		return ""
	}
	strBuf := api.NewStringBuffer(0, 0)
	synAutoCompleteAPI().SysCallN(3, m.Instance(), api.PasStr(token), uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynAutoComplete) Execute(token string, editor ICustomSynEdit) {
	if !m.IsValid() {
		return
	}
	synAutoCompleteAPI().SysCallN(4, m.Instance(), api.PasStr(token), base.GetObjectUintptr(editor))
}

func (m *TSynAutoComplete) AutoCompleteList() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := synAutoCompleteAPI().SysCallN(5, 0, m.Instance())
	return AsStrings(r)
}

func (m *TSynAutoComplete) SetAutoCompleteList(value IStrings) {
	if !m.IsValid() {
		return
	}
	synAutoCompleteAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynAutoComplete) EndOfTokenChr() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synAutoCompleteAPI().SysCallN(6, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynAutoComplete) SetEndOfTokenChr(value string) {
	if !m.IsValid() {
		return
	}
	synAutoCompleteAPI().SysCallN(6, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynAutoComplete) ShortCut() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := synAutoCompleteAPI().SysCallN(7, 0, m.Instance())
	return types.TShortCut(r)
}

func (m *TSynAutoComplete) SetShortCut(value types.TShortCut) {
	if !m.IsValid() {
		return
	}
	synAutoCompleteAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TSynAutoComplete) ExecCommandID() types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	r := synAutoCompleteAPI().SysCallN(8, 0, m.Instance())
	return types.TSynEditorCommand(r)
}

func (m *TSynAutoComplete) SetExecCommandID(value types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	synAutoCompleteAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

// NewSynAutoComplete class constructor
func NewSynAutoComplete(owner IComponent) ISynAutoComplete {
	r := synAutoCompleteAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynAutoComplete(r)
}

func TSynAutoCompleteClass() types.TClass {
	r := synAutoCompleteAPI().SysCallN(9)
	return types.TClass(r)
}

var (
	synAutoCompleteOnce   base.Once
	synAutoCompleteImport *imports.Imports = nil
)

func synAutoCompleteAPI() *imports.Imports {
	synAutoCompleteOnce.Do(func() {
		synAutoCompleteImport = api.NewDefaultImports()
		synAutoCompleteImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynAutoComplete_Create", 0), // constructor NewSynAutoComplete
			/* 1 */ imports.NewTable("TSynAutoComplete_EditorsCount", 0), // function EditorsCount
			/* 2 */ imports.NewTable("TSynAutoComplete_GetTokenList", 0), // function GetTokenList
			/* 3 */ imports.NewTable("TSynAutoComplete_GetTokenValue", 0), // function GetTokenValue
			/* 4 */ imports.NewTable("TSynAutoComplete_Execute", 0), // procedure Execute
			/* 5 */ imports.NewTable("TSynAutoComplete_AutoCompleteList", 0), // property AutoCompleteList
			/* 6 */ imports.NewTable("TSynAutoComplete_EndOfTokenChr", 0), // property EndOfTokenChr
			/* 7 */ imports.NewTable("TSynAutoComplete_ShortCut", 0), // property ShortCut
			/* 8 */ imports.NewTable("TSynAutoComplete_ExecCommandID", 0), // property ExecCommandID
			/* 9 */ imports.NewTable("TSynAutoComplete_TClass", 0), // function TSynAutoCompleteClass
		}
	})
	return synAutoCompleteImport
}
