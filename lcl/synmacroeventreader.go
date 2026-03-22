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

// ISynMacroEventReader Parent: IObject
type ISynMacroEventReader interface {
	IObject
	EventCommand() types.TSynEditorCommand          // function
	ParamCount() int32                              // function
	ParamType(index int32) types.TSynEventParamType // property ParamType Getter
	ParamAsString(index int32) string               // property ParamAsString Getter
	ParamAsInt(index int32) int32                   // property ParamAsInt Getter
}

type TSynMacroEventReader struct {
	TObject
}

func (m *TSynMacroEventReader) EventCommand() types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	r := synMacroEventReaderAPI().SysCallN(0, m.Instance())
	return types.TSynEditorCommand(r)
}

func (m *TSynMacroEventReader) ParamCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synMacroEventReaderAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TSynMacroEventReader) ParamType(index int32) types.TSynEventParamType {
	if !m.IsValid() {
		return 0
	}
	r := synMacroEventReaderAPI().SysCallN(2, m.Instance(), uintptr(index))
	return types.TSynEventParamType(r)
}

func (m *TSynMacroEventReader) ParamAsString(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := synMacroEventReaderAPI().SysCallN(3, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TSynMacroEventReader) ParamAsInt(index int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synMacroEventReaderAPI().SysCallN(4, m.Instance(), uintptr(index))
	return int32(r)
}

var (
	synMacroEventReaderOnce   base.Once
	synMacroEventReaderImport *imports.Imports = nil
)

func synMacroEventReaderAPI() *imports.Imports {
	synMacroEventReaderOnce.Do(func() {
		synMacroEventReaderImport = api.NewDefaultImports()
		synMacroEventReaderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynMacroEventReader_EventCommand", 0), // function EventCommand
			/* 1 */ imports.NewTable("TSynMacroEventReader_ParamCount", 0), // function ParamCount
			/* 2 */ imports.NewTable("TSynMacroEventReader_ParamType", 0), // property ParamType
			/* 3 */ imports.NewTable("TSynMacroEventReader_ParamAsString", 0), // property ParamAsString
			/* 4 */ imports.NewTable("TSynMacroEventReader_ParamAsInt", 0), // property ParamAsInt
		}
	})
	return synMacroEventReaderImport
}
