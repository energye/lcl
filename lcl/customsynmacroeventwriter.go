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

// ICustomSynMacroEventWriter Parent: ISynMacroEventWriter
type ICustomSynMacroEventWriter interface {
	ISynMacroEventWriter
	SetOnWriteEventCommand(fn TSynMacroWriteCommandEvent) // property event
	SetOnWriteEventStr(fn TSynMacroWriteEventStrEvent)    // property event
	SetOnWriteEventInt(fn TSynMacroWriteEventIntEvent)    // property event
}

type TCustomSynMacroEventWriter struct {
	TSynMacroEventWriter
}

func (m *TCustomSynMacroEventWriter) SetOnWriteEventCommand(fn TSynMacroWriteCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynMacroWriteCommandEvent(fn)
	base.SetEvent(m, 1, customSynMacroEventWriterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynMacroEventWriter) SetOnWriteEventStr(fn TSynMacroWriteEventStrEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynMacroWriteEventStrEvent(fn)
	base.SetEvent(m, 2, customSynMacroEventWriterAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynMacroEventWriter) SetOnWriteEventInt(fn TSynMacroWriteEventIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynMacroWriteEventIntEvent(fn)
	base.SetEvent(m, 3, customSynMacroEventWriterAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomSynMacroEventWriter class constructor
func NewCustomSynMacroEventWriter() ICustomSynMacroEventWriter {
	r := customSynMacroEventWriterAPI().SysCallN(0)
	return AsCustomSynMacroEventWriter(r)
}

var (
	customSynMacroEventWriterOnce   base.Once
	customSynMacroEventWriterImport *imports.Imports = nil
)

func customSynMacroEventWriterAPI() *imports.Imports {
	customSynMacroEventWriterOnce.Do(func() {
		customSynMacroEventWriterImport = api.NewDefaultImports()
		customSynMacroEventWriterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomSynMacroEventWriter_Create", 0), // constructor NewCustomSynMacroEventWriter
			/* 1 */ imports.NewTable("TCustomSynMacroEventWriter_OnWriteEventCommand", 0), // event OnWriteEventCommand
			/* 2 */ imports.NewTable("TCustomSynMacroEventWriter_OnWriteEventStr", 0), // event OnWriteEventStr
			/* 3 */ imports.NewTable("TCustomSynMacroEventWriter_OnWriteEventInt", 0), // event OnWriteEventInt
		}
	})
	return customSynMacroEventWriterImport
}
