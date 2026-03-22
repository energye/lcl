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

// ISynEditStringsLinked Parent: ISynEditStrings
type ISynEditStringsLinked interface {
	ISynEditStrings
	RemoveHandlers(owner IObject) // procedure
	// NextLines
	//  function GetPhysicalCharWidths(Line: PChar; LineLen, Index: Integer): TPhysicalCharWidths; override;
	NextLines() ISynEditStrings         // property NextLines Getter
	SetNextLines(value ISynEditStrings) // property NextLines Setter
}

type TSynEditStringsLinked struct {
	TSynEditStrings
}

func (m *TSynEditStringsLinked) RemoveHandlers(owner IObject) {
	if !m.IsValid() {
		return
	}
	synEditStringsLinkedAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(owner))
}

func (m *TSynEditStringsLinked) NextLines() ISynEditStrings {
	if !m.IsValid() {
		return nil
	}
	r := synEditStringsLinkedAPI().SysCallN(2, 0, m.Instance())
	return AsSynEditStrings(r)
}

func (m *TSynEditStringsLinked) SetNextLines(value ISynEditStrings) {
	if !m.IsValid() {
		return
	}
	synEditStringsLinkedAPI().SysCallN(2, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewSynEditStringsLinked class constructor
func NewSynEditStringsLinked() ISynEditStringsLinked {
	r := synEditStringsLinkedAPI().SysCallN(0)
	return AsSynEditStringsLinked(r)
}

var (
	synEditStringsLinkedOnce   base.Once
	synEditStringsLinkedImport *imports.Imports = nil
)

func synEditStringsLinkedAPI() *imports.Imports {
	synEditStringsLinkedOnce.Do(func() {
		synEditStringsLinkedImport = api.NewDefaultImports()
		synEditStringsLinkedImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditStringsLinked_Create", 0), // constructor NewSynEditStringsLinked
			/* 1 */ imports.NewTable("TSynEditStringsLinked_RemoveHandlers", 0), // procedure RemoveHandlers
			/* 2 */ imports.NewTable("TSynEditStringsLinked_NextLines", 0), // property NextLines
		}
	})
	return synEditStringsLinkedImport
}
