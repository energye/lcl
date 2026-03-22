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

// ISynPluginCustomSyncroEdit Parent: ISynPluginSyncronizedEditBase
type ISynPluginCustomSyncroEdit interface {
	ISynPluginSyncronizedEditBase
	// IncExternalEditLock
	//  destructor Destroy; override;
	IncExternalEditLock() // procedure
	DecExternalEditLock() // procedure
}

type TSynPluginCustomSyncroEdit struct {
	TSynPluginSyncronizedEditBase
}

func (m *TSynPluginCustomSyncroEdit) IncExternalEditLock() {
	if !m.IsValid() {
		return
	}
	synPluginCustomSyncroEditAPI().SysCallN(1, m.Instance())
}

func (m *TSynPluginCustomSyncroEdit) DecExternalEditLock() {
	if !m.IsValid() {
		return
	}
	synPluginCustomSyncroEditAPI().SysCallN(2, m.Instance())
}

// NewSynPluginCustomSyncroEdit class constructor
func NewSynPluginCustomSyncroEdit(owner IComponent) ISynPluginCustomSyncroEdit {
	r := synPluginCustomSyncroEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynPluginCustomSyncroEdit(r)
}

func TSynPluginCustomSyncroEditClass() types.TClass {
	r := synPluginCustomSyncroEditAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	synPluginCustomSyncroEditOnce   base.Once
	synPluginCustomSyncroEditImport *imports.Imports = nil
)

func synPluginCustomSyncroEditAPI() *imports.Imports {
	synPluginCustomSyncroEditOnce.Do(func() {
		synPluginCustomSyncroEditImport = api.NewDefaultImports()
		synPluginCustomSyncroEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynPluginCustomSyncroEdit_Create", 0), // constructor NewSynPluginCustomSyncroEdit
			/* 1 */ imports.NewTable("TSynPluginCustomSyncroEdit_IncExternalEditLock", 0), // procedure IncExternalEditLock
			/* 2 */ imports.NewTable("TSynPluginCustomSyncroEdit_DecExternalEditLock", 0), // procedure DecExternalEditLock
			/* 3 */ imports.NewTable("TSynPluginCustomSyncroEdit_TClass", 0), // function TSynPluginCustomSyncroEditClass
		}
	})
	return synPluginCustomSyncroEditImport
}
