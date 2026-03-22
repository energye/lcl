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

// ISynCustomBeautifier Parent: IComponent
type ISynCustomBeautifier interface {
	IComponent
	GetCopy() ISynCustomBeautifier // function
	// GetDesiredIndentForLineWithSEBaseSEStringsSECaret
	//  GetDesiredIndentForLine: Returns the 1-based Physical x pos
	GetDesiredIndentForLineWithSEBaseSEStringsSECaret(editor ISynEditBase, lines ISynEditStrings, caret ISynEditCaret) int32                                                                  // function
	GetDesiredIndentForLineWithSEBaseSEStringsSECaretBoolStr(editor ISynEditBase, lines ISynEditStrings, caret ISynEditCaret, outReplaceIndent *bool, outDesiredIndent *string) int32         // function
	BeforeCommand(editor ISynEditBase, lines ISynEditStrings, caret ISynEditCaret, command *types.TSynEditorCommand, initialCmd types.TSynEditorCommand)                                      // procedure
	AfterCommand(editor ISynEditBase, lines ISynEditStrings, caret ISynEditCaret, command *types.TSynEditorCommand, initialCmd types.TSynEditorCommand, startLinePos int32, endLinePos int32) // procedure
	AutoIndent() bool                                                                                                                                                                         // property AutoIndent Getter
	SetAutoIndent(value bool)                                                                                                                                                                 // property AutoIndent Setter
	SetOnGetDesiredIndent(fn TSynBeautifierGetIndentEvent)                                                                                                                                    // property event
}

type TSynCustomBeautifier struct {
	TComponent
}

func (m *TSynCustomBeautifier) GetCopy() ISynCustomBeautifier {
	if !m.IsValid() {
		return nil
	}
	r := synCustomBeautifierAPI().SysCallN(0, m.Instance())
	return AsSynCustomBeautifier(r)
}

func (m *TSynCustomBeautifier) GetDesiredIndentForLineWithSEBaseSEStringsSECaret(editor ISynEditBase, lines ISynEditStrings, caret ISynEditCaret) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synCustomBeautifierAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(editor), base.GetObjectUintptr(lines), base.GetObjectUintptr(caret))
	return int32(r)
}

func (m *TSynCustomBeautifier) GetDesiredIndentForLineWithSEBaseSEStringsSECaretBoolStr(editor ISynEditBase, lines ISynEditStrings, caret ISynEditCaret, outReplaceIndent *bool, outDesiredIndent *string) int32 {
	if !m.IsValid() {
		return 0
	}
	var desiredIndentPtr uintptr
	r := synCustomBeautifierAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(editor), base.GetObjectUintptr(lines), base.GetObjectUintptr(caret), uintptr(base.UnsafePointer(outReplaceIndent)), uintptr(base.UnsafePointer(&desiredIndentPtr)))
	*outDesiredIndent = api.GoStr(desiredIndentPtr)
	return int32(r)
}

func (m *TSynCustomBeautifier) BeforeCommand(editor ISynEditBase, lines ISynEditStrings, caret ISynEditCaret, command *types.TSynEditorCommand, initialCmd types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	commandPtr := uintptr(*command)
	synCustomBeautifierAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(editor), base.GetObjectUintptr(lines), base.GetObjectUintptr(caret), uintptr(base.UnsafePointer(&commandPtr)), uintptr(initialCmd))
	*command = types.TSynEditorCommand(commandPtr)
}

func (m *TSynCustomBeautifier) AfterCommand(editor ISynEditBase, lines ISynEditStrings, caret ISynEditCaret, command *types.TSynEditorCommand, initialCmd types.TSynEditorCommand, startLinePos int32, endLinePos int32) {
	if !m.IsValid() {
		return
	}
	commandPtr := uintptr(*command)
	synCustomBeautifierAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(editor), base.GetObjectUintptr(lines), base.GetObjectUintptr(caret), uintptr(base.UnsafePointer(&commandPtr)), uintptr(initialCmd), uintptr(startLinePos), uintptr(endLinePos))
	*command = types.TSynEditorCommand(commandPtr)
}

func (m *TSynCustomBeautifier) AutoIndent() bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomBeautifierAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynCustomBeautifier) SetAutoIndent(value bool) {
	if !m.IsValid() {
		return
	}
	synCustomBeautifierAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynCustomBeautifier) SetOnGetDesiredIndent(fn TSynBeautifierGetIndentEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynBeautifierGetIndentEvent(fn)
	base.SetEvent(m, 6, synCustomBeautifierAPI(), api.MakeEventDataPtr(cb))
}

var (
	synCustomBeautifierOnce   base.Once
	synCustomBeautifierImport *imports.Imports = nil
)

func synCustomBeautifierAPI() *imports.Imports {
	synCustomBeautifierOnce.Do(func() {
		synCustomBeautifierImport = api.NewDefaultImports()
		synCustomBeautifierImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynCustomBeautifier_GetCopy", 0), // function GetCopy
			/* 1 */ imports.NewTable("TSynCustomBeautifier_GetDesiredIndentForLineWithSEBaseSEStringsSECaret", 0), // function GetDesiredIndentForLineWithSEBaseSEStringsSECaret
			/* 2 */ imports.NewTable("TSynCustomBeautifier_GetDesiredIndentForLineWithSEBaseSEStringsSECaretBoolStr", 0), // function GetDesiredIndentForLineWithSEBaseSEStringsSECaretBoolStr
			/* 3 */ imports.NewTable("TSynCustomBeautifier_BeforeCommand", 0), // procedure BeforeCommand
			/* 4 */ imports.NewTable("TSynCustomBeautifier_AfterCommand", 0), // procedure AfterCommand
			/* 5 */ imports.NewTable("TSynCustomBeautifier_AutoIndent", 0), // property AutoIndent
			/* 6 */ imports.NewTable("TSynCustomBeautifier_OnGetDesiredIndent", 0), // event OnGetDesiredIndent
		}
	})
	return synCustomBeautifierImport
}
