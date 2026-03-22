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

// ISynEditPointBase Parent: IObject
type ISynEditPointBase interface {
	IObject
	Lock()                                // procedure
	Unlock()                              // procedure
	Lines() ISynEditStringsLinked         // property Lines Getter
	SetLines(value ISynEditStringsLinked) // property Lines Setter
	Locked() bool                         // property Locked Getter
}

type TSynEditPointBase struct {
	TObject
}

func (m *TSynEditPointBase) Lock() {
	if !m.IsValid() {
		return
	}
	synEditPointBaseAPI().SysCallN(2, m.Instance())
}

func (m *TSynEditPointBase) Unlock() {
	if !m.IsValid() {
		return
	}
	synEditPointBaseAPI().SysCallN(3, m.Instance())
}

func (m *TSynEditPointBase) Lines() ISynEditStringsLinked {
	if !m.IsValid() {
		return nil
	}
	r := synEditPointBaseAPI().SysCallN(4, 0, m.Instance())
	return AsSynEditStringsLinked(r)
}

func (m *TSynEditPointBase) SetLines(value ISynEditStringsLinked) {
	if !m.IsValid() {
		return
	}
	synEditPointBaseAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynEditPointBase) Locked() bool {
	if !m.IsValid() {
		return false
	}
	r := synEditPointBaseAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

// NewSynEditPointBase class constructor
func NewSynEditPointBase() ISynEditPointBase {
	r := synEditPointBaseAPI().SysCallN(0)
	return AsSynEditPointBase(r)
}

// NewSynEditPointBaseWithSynEditStringsLinked class constructor
func NewSynEditPointBaseWithSynEditStringsLinked(lines ISynEditStringsLinked) ISynEditPointBase {
	r := synEditPointBaseAPI().SysCallN(1, base.GetObjectUintptr(lines))
	return AsSynEditPointBase(r)
}

var (
	synEditPointBaseOnce   base.Once
	synEditPointBaseImport *imports.Imports = nil
)

func synEditPointBaseAPI() *imports.Imports {
	synEditPointBaseOnce.Do(func() {
		synEditPointBaseImport = api.NewDefaultImports()
		synEditPointBaseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditPointBase_Create", 0), // constructor NewSynEditPointBase
			/* 1 */ imports.NewTable("TSynEditPointBase_CreateWithSynEditStringsLinked", 0), // constructor NewSynEditPointBaseWithSynEditStringsLinked
			/* 2 */ imports.NewTable("TSynEditPointBase_Lock", 0), // procedure Lock
			/* 3 */ imports.NewTable("TSynEditPointBase_Unlock", 0), // procedure Unlock
			/* 4 */ imports.NewTable("TSynEditPointBase_Lines", 0), // property Lines
			/* 5 */ imports.NewTable("TSynEditPointBase_Locked", 0), // property Locked
		}
	})
	return synEditPointBaseImport
}
