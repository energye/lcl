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

// ISynMacroEventWriter Parent: IObject
type ISynMacroEventWriter interface {
	IObject
	WriteEventCommand(cmd types.TSynEditorCommand) // procedure
	WriteEventParamWithStr(param string)           // procedure
	WriteEventParamWithInt(param int32)            // procedure
}

type TSynMacroEventWriter struct {
	TObject
}

func (m *TSynMacroEventWriter) WriteEventCommand(cmd types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	synMacroEventWriterAPI().SysCallN(0, m.Instance(), uintptr(cmd))
}

func (m *TSynMacroEventWriter) WriteEventParamWithStr(param string) {
	if !m.IsValid() {
		return
	}
	synMacroEventWriterAPI().SysCallN(1, m.Instance(), api.PasStr(param))
}

func (m *TSynMacroEventWriter) WriteEventParamWithInt(param int32) {
	if !m.IsValid() {
		return
	}
	synMacroEventWriterAPI().SysCallN(2, m.Instance(), uintptr(param))
}

var (
	synMacroEventWriterOnce   base.Once
	synMacroEventWriterImport *imports.Imports = nil
)

func synMacroEventWriterAPI() *imports.Imports {
	synMacroEventWriterOnce.Do(func() {
		synMacroEventWriterImport = api.NewDefaultImports()
		synMacroEventWriterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynMacroEventWriter_WriteEventCommand", 0), // procedure WriteEventCommand
			/* 1 */ imports.NewTable("TSynMacroEventWriter_WriteEventParamWithStr", 0), // procedure WriteEventParamWithStr
			/* 2 */ imports.NewTable("TSynMacroEventWriter_WriteEventParamWithInt", 0), // procedure WriteEventParamWithInt
		}
	})
	return synMacroEventWriterImport
}
