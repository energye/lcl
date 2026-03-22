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

// ILazSynMultiEditPlugin Parent: ILazSynEditPlugin
type ILazSynMultiEditPlugin interface {
	ILazSynEditPlugin
	// AddEditor
	//  * Add/RemoveEditor versus Editor
	//  Unless stated otherwise Plugins inherting from TLazSynMultiEditPlugin can
	//  either use Add/RemoveEditor or Editor (single-editor-property).
	//  If Editors are added via AddEditor, then an Editor only set via "Editor:="
	//  may be lost/ignored.
	//  If using AddEditor, the "Editor" property may be used to set/read the
	//  current Editor (out of those in the list). This does however depend on the
	//  inherited class.
	AddEditor(value ICustomSynEdit) int32    // function
	RemoveEditor(value ICustomSynEdit) int32 // function
	Editors(index int32) ICustomSynEdit      // property Editors Getter
	EditorCount() int32                      // property EditorCount Getter
}

type TLazSynMultiEditPlugin struct {
	TLazSynEditPlugin
}

func (m *TLazSynMultiEditPlugin) AddEditor(value ICustomSynEdit) int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynMultiEditPluginAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(value))
	return int32(r)
}

func (m *TLazSynMultiEditPlugin) RemoveEditor(value ICustomSynEdit) int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynMultiEditPluginAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(value))
	return int32(r)
}

func (m *TLazSynMultiEditPlugin) Editors(index int32) ICustomSynEdit {
	if !m.IsValid() {
		return nil
	}
	r := lazSynMultiEditPluginAPI().SysCallN(3, m.Instance(), uintptr(index))
	return AsCustomSynEdit(r)
}

func (m *TLazSynMultiEditPlugin) EditorCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynMultiEditPluginAPI().SysCallN(4, m.Instance())
	return int32(r)
}

// NewLazSynMultiEditPlugin class constructor
func NewLazSynMultiEditPlugin(owner IComponent) ILazSynMultiEditPlugin {
	r := lazSynMultiEditPluginAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLazSynMultiEditPlugin(r)
}

func TLazSynMultiEditPluginClass() types.TClass {
	r := lazSynMultiEditPluginAPI().SysCallN(5)
	return types.TClass(r)
}

var (
	lazSynMultiEditPluginOnce   base.Once
	lazSynMultiEditPluginImport *imports.Imports = nil
)

func lazSynMultiEditPluginAPI() *imports.Imports {
	lazSynMultiEditPluginOnce.Do(func() {
		lazSynMultiEditPluginImport = api.NewDefaultImports()
		lazSynMultiEditPluginImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazSynMultiEditPlugin_Create", 0), // constructor NewLazSynMultiEditPlugin
			/* 1 */ imports.NewTable("TLazSynMultiEditPlugin_AddEditor", 0), // function AddEditor
			/* 2 */ imports.NewTable("TLazSynMultiEditPlugin_RemoveEditor", 0), // function RemoveEditor
			/* 3 */ imports.NewTable("TLazSynMultiEditPlugin_Editors", 0), // property Editors
			/* 4 */ imports.NewTable("TLazSynMultiEditPlugin_EditorCount", 0), // property EditorCount
			/* 5 */ imports.NewTable("TLazSynMultiEditPlugin_TClass", 0), // function TLazSynMultiEditPluginClass
		}
	})
	return lazSynMultiEditPluginImport
}
