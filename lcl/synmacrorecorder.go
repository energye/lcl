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

// ISynMacroRecorder Parent: ICustomSynMacroRecorder
type ISynMacroRecorder interface {
	ICustomSynMacroRecorder
	RecordCommandID() types.TSynEditorCommand           // property RecordCommandID Getter
	SetRecordCommandID(value types.TSynEditorCommand)   // property RecordCommandID Setter
	PlaybackCommandID() types.TSynEditorCommand         // property PlaybackCommandID Getter
	SetPlaybackCommandID(value types.TSynEditorCommand) // property PlaybackCommandID Setter
}

type TSynMacroRecorder struct {
	TCustomSynMacroRecorder
}

func (m *TSynMacroRecorder) RecordCommandID() types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	r := synMacroRecorderAPI().SysCallN(1, 0, m.Instance())
	return types.TSynEditorCommand(r)
}

func (m *TSynMacroRecorder) SetRecordCommandID(value types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	synMacroRecorderAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TSynMacroRecorder) PlaybackCommandID() types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	r := synMacroRecorderAPI().SysCallN(2, 0, m.Instance())
	return types.TSynEditorCommand(r)
}

func (m *TSynMacroRecorder) SetPlaybackCommandID(value types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	synMacroRecorderAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

// NewSynMacroRecorder class constructor
func NewSynMacroRecorder(owner IComponent) ISynMacroRecorder {
	r := synMacroRecorderAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynMacroRecorder(r)
}

func TSynMacroRecorderClass() types.TClass {
	r := synMacroRecorderAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	synMacroRecorderOnce   base.Once
	synMacroRecorderImport *imports.Imports = nil
)

func synMacroRecorderAPI() *imports.Imports {
	synMacroRecorderOnce.Do(func() {
		synMacroRecorderImport = api.NewDefaultImports()
		synMacroRecorderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynMacroRecorder_Create", 0), // constructor NewSynMacroRecorder
			/* 1 */ imports.NewTable("TSynMacroRecorder_RecordCommandID", 0), // property RecordCommandID
			/* 2 */ imports.NewTable("TSynMacroRecorder_PlaybackCommandID", 0), // property PlaybackCommandID
			/* 3 */ imports.NewTable("TSynMacroRecorder_TClass", 0), // function TSynMacroRecorderClass
		}
	})
	return synMacroRecorderImport
}
