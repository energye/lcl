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

// ISynIgnoredEvent Parent: ISynMacroEvent
type ISynIgnoredEvent interface {
	ISynMacroEvent
}

type TSynIgnoredEvent struct {
	TSynMacroEvent
}

// NewSynIgnoredEvent class constructor
func NewSynIgnoredEvent() ISynIgnoredEvent {
	r := synIgnoredEventAPI().SysCallN(0)
	return AsSynIgnoredEvent(r)
}

var (
	synIgnoredEventOnce   base.Once
	synIgnoredEventImport *imports.Imports = nil
)

func synIgnoredEventAPI() *imports.Imports {
	synIgnoredEventOnce.Do(func() {
		synIgnoredEventImport = api.NewDefaultImports()
		synIgnoredEventImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynIgnoredEvent_Create", 0), // constructor NewSynIgnoredEvent
		}
	})
	return synIgnoredEventImport
}
