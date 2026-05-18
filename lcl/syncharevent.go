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

// ISynCharEvent Parent: ISynMacroEvent
type ISynCharEvent interface {
	ISynMacroEvent
	Key() string         // property Key Getter
	SetKey(value string) // property Key Setter
}

type TSynCharEvent struct {
	TSynMacroEvent
}

func (m *TSynCharEvent) Key() (result string) {
	if !m.IsValid() {
		return
	}
	strBuf := api.NewStringBuffer(0, 0)
	synCharEventAPI().SysCallN(1, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&strBuf.Data)), uintptr(base.UnsafePointer(&strBuf.Size)))
	defer strBuf.Release()
	result = strBuf.String()
	return
}

func (m *TSynCharEvent) SetKey(value string) {
	if !m.IsValid() {
		return
	}
	synCharEventAPI().SysCallN(1, 1, m.Instance(), api.PasStr(value))
}

// NewSynCharEvent class constructor
func NewSynCharEvent() ISynCharEvent {
	r := synCharEventAPI().SysCallN(0)
	return AsSynCharEvent(r)
}

var (
	synCharEventOnce   base.Once
	synCharEventImport *imports.Imports = nil
)

func synCharEventAPI() *imports.Imports {
	synCharEventOnce.Do(func() {
		synCharEventImport = api.NewDefaultImports()
		synCharEventImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynCharEvent_Create", 0), // constructor NewSynCharEvent
			/* 1 */ imports.NewTable("TSynCharEvent_Key", 0), // property Key
		}
	})
	return synCharEventImport
}
