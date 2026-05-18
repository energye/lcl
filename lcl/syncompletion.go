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

// ISynCompletion Parent: ISynBaseCompletion
type ISynCompletion interface {
	ISynBaseCompletion
	AddCharAtCursor(utf8Char string)                // procedure
	DeleteCharBeforeCursor()                        // procedure
	ShortCut() types.TShortCut                      // property ShortCut Getter
	SetShortCut(value types.TShortCut)              // property ShortCut Setter
	EndOfTokenChr() string                          // property EndOfTokenChr Getter
	SetEndOfTokenChr(value string)                  // property EndOfTokenChr Setter
	ExecCommandID() types.TSynEditorCommand         // property ExecCommandID Getter
	SetExecCommandID(value types.TSynEditorCommand) // property ExecCommandID Setter
	ToggleReplaceWhole() bool                       // property ToggleReplaceWhole Getter
	SetToggleReplaceWhole(value bool)               // property ToggleReplaceWhole Setter
	SetOnCodeCompletion(fn TCodeCompletionEvent)    // property event
}

type TSynCompletion struct {
	TSynBaseCompletion
}

func (m *TSynCompletion) AddCharAtCursor(utf8Char string) {
	if !m.IsValid() {
		return
	}
	synCompletionAPI().SysCallN(1, m.Instance(), api.PasStr(utf8Char))
}

func (m *TSynCompletion) DeleteCharBeforeCursor() {
	if !m.IsValid() {
		return
	}
	synCompletionAPI().SysCallN(2, m.Instance())
}

func (m *TSynCompletion) ShortCut() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := synCompletionAPI().SysCallN(3, 0, m.Instance())
	return types.TShortCut(r)
}

func (m *TSynCompletion) SetShortCut(value types.TShortCut) {
	if !m.IsValid() {
		return
	}
	synCompletionAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TSynCompletion) EndOfTokenChr() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synCompletionAPI().SysCallN(4, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynCompletion) SetEndOfTokenChr(value string) {
	if !m.IsValid() {
		return
	}
	synCompletionAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TSynCompletion) ExecCommandID() types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	r := synCompletionAPI().SysCallN(5, 0, m.Instance())
	return types.TSynEditorCommand(r)
}

func (m *TSynCompletion) SetExecCommandID(value types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	synCompletionAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TSynCompletion) ToggleReplaceWhole() bool {
	if !m.IsValid() {
		return false
	}
	r := synCompletionAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynCompletion) SetToggleReplaceWhole(value bool) {
	if !m.IsValid() {
		return
	}
	synCompletionAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynCompletion) SetOnCodeCompletion(fn TCodeCompletionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCodeCompletionEvent(fn)
	base.SetEvent(m, 7, synCompletionAPI(), api.MakeEventDataPtr(cb))
}

// NewSynCompletion class constructor
func NewSynCompletion(owner IComponent) ISynCompletion {
	r := synCompletionAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynCompletion(r)
}

func TSynCompletionClass() types.TClass {
	r := synCompletionAPI().SysCallN(8)
	return types.TClass(r)
}

var (
	synCompletionOnce   base.Once
	synCompletionImport *imports.Imports = nil
)

func synCompletionAPI() *imports.Imports {
	synCompletionOnce.Do(func() {
		synCompletionImport = api.NewDefaultImports()
		synCompletionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynCompletion_Create", 0), // constructor NewSynCompletion
			/* 1 */ imports.NewTable("TSynCompletion_AddCharAtCursor", 0), // procedure AddCharAtCursor
			/* 2 */ imports.NewTable("TSynCompletion_DeleteCharBeforeCursor", 0), // procedure DeleteCharBeforeCursor
			/* 3 */ imports.NewTable("TSynCompletion_ShortCut", 0), // property ShortCut
			/* 4 */ imports.NewTable("TSynCompletion_EndOfTokenChr", 0), // property EndOfTokenChr
			/* 5 */ imports.NewTable("TSynCompletion_ExecCommandID", 0), // property ExecCommandID
			/* 6 */ imports.NewTable("TSynCompletion_ToggleReplaceWhole", 0), // property ToggleReplaceWhole
			/* 7 */ imports.NewTable("TSynCompletion_OnCodeCompletion", 0), // event OnCodeCompletion
			/* 8 */ imports.NewTable("TSynCompletion_TClass", 0), // function TSynCompletionClass
		}
	})
	return synCompletionImport
}
