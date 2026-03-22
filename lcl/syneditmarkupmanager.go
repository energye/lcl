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

// ISynEditMarkupManager Parent: ISynEditMarkup
type ISynEditMarkupManager interface {
	ISynEditMarkup
	Count() int32                                                 // function
	DoVisibleChanged(visible bool)                                // procedure
	AddMarkUp(markUp ISynEditMarkup, asFirst bool)                // procedure
	RemoveMarkUp(markUp ISynEditMarkup)                           // procedure
	Markup(index int32) ISynEditMarkup                            // property Markup Getter
	MarkupByClass(index types.TSynEditMarkupClass) ISynEditMarkup // property MarkupByClass Getter
}

type TSynEditMarkupManager struct {
	TSynEditMarkup
}

func (m *TSynEditMarkupManager) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synEditMarkupManagerAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TSynEditMarkupManager) DoVisibleChanged(visible bool) {
	if !m.IsValid() {
		return
	}
	synEditMarkupManagerAPI().SysCallN(2, m.Instance(), api.PasBool(visible))
}

func (m *TSynEditMarkupManager) AddMarkUp(markUp ISynEditMarkup, asFirst bool) {
	if !m.IsValid() {
		return
	}
	synEditMarkupManagerAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(markUp), api.PasBool(asFirst))
}

func (m *TSynEditMarkupManager) RemoveMarkUp(markUp ISynEditMarkup) {
	if !m.IsValid() {
		return
	}
	synEditMarkupManagerAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(markUp))
}

func (m *TSynEditMarkupManager) Markup(index int32) ISynEditMarkup {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkupManagerAPI().SysCallN(5, m.Instance(), uintptr(index))
	return AsSynEditMarkup(r)
}

func (m *TSynEditMarkupManager) MarkupByClass(index types.TSynEditMarkupClass) ISynEditMarkup {
	if !m.IsValid() {
		return nil
	}
	r := synEditMarkupManagerAPI().SysCallN(6, m.Instance(), uintptr(index))
	return AsSynEditMarkup(r)
}

// NewSynEditMarkupManager class constructor
func NewSynEditMarkupManager(synEdit ISynEditBase) ISynEditMarkupManager {
	r := synEditMarkupManagerAPI().SysCallN(0, base.GetObjectUintptr(synEdit))
	return AsSynEditMarkupManager(r)
}

var (
	synEditMarkupManagerOnce   base.Once
	synEditMarkupManagerImport *imports.Imports = nil
)

func synEditMarkupManagerAPI() *imports.Imports {
	synEditMarkupManagerOnce.Do(func() {
		synEditMarkupManagerImport = api.NewDefaultImports()
		synEditMarkupManagerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditMarkupManager_Create", 0), // constructor NewSynEditMarkupManager
			/* 1 */ imports.NewTable("TSynEditMarkupManager_Count", 0), // function Count
			/* 2 */ imports.NewTable("TSynEditMarkupManager_DoVisibleChanged", 0), // procedure DoVisibleChanged
			/* 3 */ imports.NewTable("TSynEditMarkupManager_AddMarkUp", 0), // procedure AddMarkUp
			/* 4 */ imports.NewTable("TSynEditMarkupManager_RemoveMarkUp", 0), // procedure RemoveMarkUp
			/* 5 */ imports.NewTable("TSynEditMarkupManager_Markup", 0), // property Markup
			/* 6 */ imports.NewTable("TSynEditMarkupManager_MarkupByClass", 0), // property MarkupByClass
		}
	})
	return synEditMarkupManagerImport
}
