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

// ISynBasicEvent Parent: ISynMacroEvent
type ISynBasicEvent interface {
	ISynMacroEvent
	Command() types.TSynEditorCommand         // property Command Getter
	SetCommand(value types.TSynEditorCommand) // property Command Setter
}

type TSynBasicEvent struct {
	TSynMacroEvent
}

func (m *TSynBasicEvent) Command() types.TSynEditorCommand {
	if !m.IsValid() {
		return 0
	}
	r := synBasicEventAPI().SysCallN(1, 0, m.Instance())
	return types.TSynEditorCommand(r)
}

func (m *TSynBasicEvent) SetCommand(value types.TSynEditorCommand) {
	if !m.IsValid() {
		return
	}
	synBasicEventAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

// NewSynBasicEvent class constructor
func NewSynBasicEvent() ISynBasicEvent {
	r := synBasicEventAPI().SysCallN(0)
	return AsSynBasicEvent(r)
}

var (
	synBasicEventOnce   base.Once
	synBasicEventImport *imports.Imports = nil
)

func synBasicEventAPI() *imports.Imports {
	synBasicEventOnce.Do(func() {
		synBasicEventImport = api.NewDefaultImports()
		synBasicEventImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynBasicEvent_Create", 0), // constructor NewSynBasicEvent
			/* 1 */ imports.NewTable("TSynBasicEvent_Command", 0), // property Command
		}
	})
	return synBasicEventImport
}
