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

// ILazSynEditPlugin Parent: ISynEditFriend
type ILazSynEditPlugin interface {
	ISynEditFriend
	Editor() ICustomSynEdit         // property Editor Getter
	SetEditor(value ICustomSynEdit) // property Editor Setter
}

type TLazSynEditPlugin struct {
	TSynEditFriend
}

func (m *TLazSynEditPlugin) Editor() ICustomSynEdit {
	if !m.IsValid() {
		return nil
	}
	r := lazSynEditPluginAPI().SysCallN(1, 0, m.Instance())
	return AsCustomSynEdit(r)
}

func (m *TLazSynEditPlugin) SetEditor(value ICustomSynEdit) {
	if !m.IsValid() {
		return
	}
	lazSynEditPluginAPI().SysCallN(1, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewLazSynEditPlugin class constructor
func NewLazSynEditPlugin(owner IComponent) ILazSynEditPlugin {
	r := lazSynEditPluginAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLazSynEditPlugin(r)
}

func TLazSynEditPluginClass() types.TClass {
	r := lazSynEditPluginAPI().SysCallN(2)
	return types.TClass(r)
}

var (
	lazSynEditPluginOnce   base.Once
	lazSynEditPluginImport *imports.Imports = nil
)

func lazSynEditPluginAPI() *imports.Imports {
	lazSynEditPluginOnce.Do(func() {
		lazSynEditPluginImport = api.NewDefaultImports()
		lazSynEditPluginImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazSynEditPlugin_Create", 0), // constructor NewLazSynEditPlugin
			/* 1 */ imports.NewTable("TLazSynEditPlugin_Editor", 0), // property Editor
			/* 2 */ imports.NewTable("TLazSynEditPlugin_TClass", 0), // function TLazSynEditPluginClass
		}
	})
	return lazSynEditPluginImport
}
