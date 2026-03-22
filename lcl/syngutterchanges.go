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

// ISynGutterChanges Parent: ISynGutterPartBase
type ISynGutterChanges interface {
	ISynGutterPartBase
	ModifiedColor() types.TColor         // property ModifiedColor Getter
	SetModifiedColor(value types.TColor) // property ModifiedColor Setter
	SavedColor() types.TColor            // property SavedColor Getter
	SetSavedColor(value types.TColor)    // property SavedColor Setter
}

type TSynGutterChanges struct {
	TSynGutterPartBase
}

func (m *TSynGutterChanges) ModifiedColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synGutterChangesAPI().SysCallN(1, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynGutterChanges) SetModifiedColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synGutterChangesAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterChanges) SavedColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synGutterChangesAPI().SysCallN(2, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynGutterChanges) SetSavedColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synGutterChangesAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

// NewSynGutterChanges class constructor
func NewSynGutterChanges(owner IComponent) ISynGutterChanges {
	r := synGutterChangesAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynGutterChanges(r)
}

func TSynGutterChangesClass() types.TClass {
	r := synGutterChangesAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	synGutterChangesOnce   base.Once
	synGutterChangesImport *imports.Imports = nil
)

func synGutterChangesAPI() *imports.Imports {
	synGutterChangesOnce.Do(func() {
		synGutterChangesImport = api.NewDefaultImports()
		synGutterChangesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterChanges_Create", 0), // constructor NewSynGutterChanges
			/* 1 */ imports.NewTable("TSynGutterChanges_ModifiedColor", 0), // property ModifiedColor
			/* 2 */ imports.NewTable("TSynGutterChanges_SavedColor", 0), // property SavedColor
			/* 3 */ imports.NewTable("TSynGutterChanges_TClass", 0), // function TSynGutterChangesClass
		}
	})
	return synGutterChangesImport
}
