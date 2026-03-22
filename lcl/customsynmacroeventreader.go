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

// ICustomSynMacroEventReader Parent: ISynMacroEventReader
type ICustomSynMacroEventReader interface {
	ISynMacroEventReader
	SetOnGetParamAsInt(fn TSynMacroReaderGetParamAsIntEvent)       // property event
	SetOnGetParamAsString(fn TSynMacroReaderGetParamAsStringEvent) // property event
	SetOnGetParamType(fn TSynMacroReaderGetParamTypeEvent)         // property event
	SetOnEventCommand(fn TSynMacroReaderEventCommandEvent)         // property event
	SetOnParamCount(fn TSynMacroReaderParamCountEvent)             // property event
}

type TCustomSynMacroEventReader struct {
	TSynMacroEventReader
}

func (m *TCustomSynMacroEventReader) SetOnGetParamAsInt(fn TSynMacroReaderGetParamAsIntEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynMacroReaderGetParamAsIntEvent(fn)
	base.SetEvent(m, 1, customSynMacroEventReaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynMacroEventReader) SetOnGetParamAsString(fn TSynMacroReaderGetParamAsStringEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynMacroReaderGetParamAsStringEvent(fn)
	base.SetEvent(m, 2, customSynMacroEventReaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynMacroEventReader) SetOnGetParamType(fn TSynMacroReaderGetParamTypeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynMacroReaderGetParamTypeEvent(fn)
	base.SetEvent(m, 3, customSynMacroEventReaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynMacroEventReader) SetOnEventCommand(fn TSynMacroReaderEventCommandEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynMacroReaderEventCommandEvent(fn)
	base.SetEvent(m, 4, customSynMacroEventReaderAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomSynMacroEventReader) SetOnParamCount(fn TSynMacroReaderParamCountEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSynMacroReaderParamCountEvent(fn)
	base.SetEvent(m, 5, customSynMacroEventReaderAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomSynMacroEventReader class constructor
func NewCustomSynMacroEventReader() ICustomSynMacroEventReader {
	r := customSynMacroEventReaderAPI().SysCallN(0)
	return AsCustomSynMacroEventReader(r)
}

var (
	customSynMacroEventReaderOnce   base.Once
	customSynMacroEventReaderImport *imports.Imports = nil
)

func customSynMacroEventReaderAPI() *imports.Imports {
	customSynMacroEventReaderOnce.Do(func() {
		customSynMacroEventReaderImport = api.NewDefaultImports()
		customSynMacroEventReaderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomSynMacroEventReader_Create", 0), // constructor NewCustomSynMacroEventReader
			/* 1 */ imports.NewTable("TCustomSynMacroEventReader_OnGetParamAsInt", 0), // event OnGetParamAsInt
			/* 2 */ imports.NewTable("TCustomSynMacroEventReader_OnGetParamAsString", 0), // event OnGetParamAsString
			/* 3 */ imports.NewTable("TCustomSynMacroEventReader_OnGetParamType", 0), // event OnGetParamType
			/* 4 */ imports.NewTable("TCustomSynMacroEventReader_OnEventCommand", 0), // event OnEventCommand
			/* 5 */ imports.NewTable("TCustomSynMacroEventReader_OnParamCount", 0), // event OnParamCount
		}
	})
	return customSynMacroEventReaderImport
}
