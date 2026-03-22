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

// ISynMacroEvent Parent: IObject
type ISynMacroEvent interface {
	IObject
	Initialize(cmd types.TSynEditorCommand, char string, data uintptr) // procedure
	Assign(anEvent ISynMacroEvent)                                     // procedure
	// LoadFromStream
	//  the CommandID must not be read inside LoadFromStream/SaveToStream. It's read by the
	//  MacroRecorder component to decide which MacroEvent class to instanciate
	LoadFromStream(stream IStream)              // procedure
	SaveToStream(stream IStream)                // procedure
	LoadFromReader(reader ISynMacroEventReader) // procedure
	SaveToWriter(writer ISynMacroEventWriter)   // procedure
	Playback(editor ICustomSynEdit)             // procedure
	AsString() string                           // property AsString Getter
	RepeatCount() int32                         // property RepeatCount Getter
	SetRepeatCount(value int32)                 // property RepeatCount Setter
}

type TSynMacroEvent struct {
	TObject
}

func (m *TSynMacroEvent) Initialize(cmd types.TSynEditorCommand, char string, data uintptr) {
	if !m.IsValid() {
		return
	}
	synMacroEventAPI().SysCallN(0, m.Instance(), uintptr(cmd), api.PasStr(char), uintptr(data))
}

func (m *TSynMacroEvent) Assign(anEvent ISynMacroEvent) {
	if !m.IsValid() {
		return
	}
	synMacroEventAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(anEvent))
}

func (m *TSynMacroEvent) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	synMacroEventAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TSynMacroEvent) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	synMacroEventAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TSynMacroEvent) LoadFromReader(reader ISynMacroEventReader) {
	if !m.IsValid() {
		return
	}
	synMacroEventAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(reader))
}

func (m *TSynMacroEvent) SaveToWriter(writer ISynMacroEventWriter) {
	if !m.IsValid() {
		return
	}
	synMacroEventAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(writer))
}

func (m *TSynMacroEvent) Playback(editor ICustomSynEdit) {
	if !m.IsValid() {
		return
	}
	synMacroEventAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(editor))
}

func (m *TSynMacroEvent) AsString() string {
	if !m.IsValid() {
		return ""
	}
	r := synMacroEventAPI().SysCallN(7, m.Instance())
	return api.GoStr(r)
}

func (m *TSynMacroEvent) RepeatCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synMacroEventAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TSynMacroEvent) SetRepeatCount(value int32) {
	if !m.IsValid() {
		return
	}
	synMacroEventAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

var (
	synMacroEventOnce   base.Once
	synMacroEventImport *imports.Imports = nil
)

func synMacroEventAPI() *imports.Imports {
	synMacroEventOnce.Do(func() {
		synMacroEventImport = api.NewDefaultImports()
		synMacroEventImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynMacroEvent_Initialize", 0), // procedure Initialize
			/* 1 */ imports.NewTable("TSynMacroEvent_Assign", 0), // procedure Assign
			/* 2 */ imports.NewTable("TSynMacroEvent_LoadFromStream", 0), // procedure LoadFromStream
			/* 3 */ imports.NewTable("TSynMacroEvent_SaveToStream", 0), // procedure SaveToStream
			/* 4 */ imports.NewTable("TSynMacroEvent_LoadFromReader", 0), // procedure LoadFromReader
			/* 5 */ imports.NewTable("TSynMacroEvent_SaveToWriter", 0), // procedure SaveToWriter
			/* 6 */ imports.NewTable("TSynMacroEvent_Playback", 0), // procedure Playback
			/* 7 */ imports.NewTable("TSynMacroEvent_AsString", 0), // property AsString
			/* 8 */ imports.NewTable("TSynMacroEvent_RepeatCount", 0), // property RepeatCount
		}
	})
	return synMacroEventImport
}
