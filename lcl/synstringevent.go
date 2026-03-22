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

// ISynStringEvent Parent: ISynMacroEvent
type ISynStringEvent interface {
	ISynMacroEvent
	Value() string         // property Value Getter
	SetValue(value string) // property Value Setter
}

type TSynStringEvent struct {
	TSynMacroEvent
}

func (m *TSynStringEvent) Value() string {
	if !m.IsValid() {
		return ""
	}
	r := synStringEventAPI().SysCallN(1, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TSynStringEvent) SetValue(value string) {
	if !m.IsValid() {
		return
	}
	synStringEventAPI().SysCallN(1, 1, m.Instance(), api.PasStr(value))
}

// NewSynStringEvent class constructor
func NewSynStringEvent() ISynStringEvent {
	r := synStringEventAPI().SysCallN(0)
	return AsSynStringEvent(r)
}

var (
	synStringEventOnce   base.Once
	synStringEventImport *imports.Imports = nil
)

func synStringEventAPI() *imports.Imports {
	synStringEventOnce.Do(func() {
		synStringEventImport = api.NewDefaultImports()
		synStringEventImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynStringEvent_Create", 0), // constructor NewSynStringEvent
			/* 1 */ imports.NewTable("TSynStringEvent_Value", 0), // property Value
		}
	})
	return synStringEventImport
}
