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

// ISynCustomFoldConfig Parent: IPersistent
type ISynCustomFoldConfig interface {
	IPersistent
	AssignWithSynCustomFoldConfig(src ISynCustomFoldConfig)  // procedure
	IsEssential() bool                                       // property IsEssential Getter
	SupportedModes() types.TSynCustomFoldConfigModes         // property SupportedModes Getter
	SetSupportedModes(value types.TSynCustomFoldConfigModes) // property SupportedModes Setter
	// FoldActions
	//  Actions representing the modes
	FoldActions() types.TSynFoldActions             // property FoldActions Getter
	Enabled() bool                                  // property Enabled Getter
	SetEnabled(value bool)                          // property Enabled Setter
	Modes() types.TSynCustomFoldConfigModes         // property Modes Getter
	SetModes(value types.TSynCustomFoldConfigModes) // property Modes Setter
	SetOnChange(fn TNotifyEvent)                    // property event
}

type TSynCustomFoldConfig struct {
	TPersistent
}

func (m *TSynCustomFoldConfig) AssignWithSynCustomFoldConfig(src ISynCustomFoldConfig) {
	if !m.IsValid() {
		return
	}
	synCustomFoldConfigAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(src))
}

func (m *TSynCustomFoldConfig) IsEssential() bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomFoldConfigAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TSynCustomFoldConfig) SupportedModes() types.TSynCustomFoldConfigModes {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldConfigAPI().SysCallN(4, 0, m.Instance())
	return types.TSynCustomFoldConfigModes(r)
}

func (m *TSynCustomFoldConfig) SetSupportedModes(value types.TSynCustomFoldConfigModes) {
	if !m.IsValid() {
		return
	}
	synCustomFoldConfigAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TSynCustomFoldConfig) FoldActions() types.TSynFoldActions {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldConfigAPI().SysCallN(5, m.Instance())
	return types.TSynFoldActions(r)
}

func (m *TSynCustomFoldConfig) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := synCustomFoldConfigAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynCustomFoldConfig) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	synCustomFoldConfigAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynCustomFoldConfig) Modes() types.TSynCustomFoldConfigModes {
	if !m.IsValid() {
		return 0
	}
	r := synCustomFoldConfigAPI().SysCallN(7, 0, m.Instance())
	return types.TSynCustomFoldConfigModes(r)
}

func (m *TSynCustomFoldConfig) SetModes(value types.TSynCustomFoldConfigModes) {
	if !m.IsValid() {
		return
	}
	synCustomFoldConfigAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TSynCustomFoldConfig) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 8, synCustomFoldConfigAPI(), api.MakeEventDataPtr(cb))
}

// NewSynCustomFoldConfig class constructor
func NewSynCustomFoldConfig() ISynCustomFoldConfig {
	r := synCustomFoldConfigAPI().SysCallN(0)
	return AsSynCustomFoldConfig(r)
}

// NewSynCustomFoldConfigWithSynCustomFoldConfigModesBool class constructor
func NewSynCustomFoldConfigWithSynCustomFoldConfigModesBool(supportedModes types.TSynCustomFoldConfigModes, anIsEssential bool) ISynCustomFoldConfig {
	r := synCustomFoldConfigAPI().SysCallN(1, uintptr(supportedModes), api.PasBool(anIsEssential))
	return AsSynCustomFoldConfig(r)
}

var (
	synCustomFoldConfigOnce   base.Once
	synCustomFoldConfigImport *imports.Imports = nil
)

func synCustomFoldConfigAPI() *imports.Imports {
	synCustomFoldConfigOnce.Do(func() {
		synCustomFoldConfigImport = api.NewDefaultImports()
		synCustomFoldConfigImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynCustomFoldConfig_Create", 0), // constructor NewSynCustomFoldConfig
			/* 1 */ imports.NewTable("TSynCustomFoldConfig_CreateWithSynCustomFoldConfigModesBool", 0), // constructor NewSynCustomFoldConfigWithSynCustomFoldConfigModesBool
			/* 2 */ imports.NewTable("TSynCustomFoldConfig_AssignWithSynCustomFoldConfig", 0), // procedure AssignWithSynCustomFoldConfig
			/* 3 */ imports.NewTable("TSynCustomFoldConfig_IsEssential", 0), // property IsEssential
			/* 4 */ imports.NewTable("TSynCustomFoldConfig_SupportedModes", 0), // property SupportedModes
			/* 5 */ imports.NewTable("TSynCustomFoldConfig_FoldActions", 0), // property FoldActions
			/* 6 */ imports.NewTable("TSynCustomFoldConfig_Enabled", 0), // property Enabled
			/* 7 */ imports.NewTable("TSynCustomFoldConfig_Modes", 0), // property Modes
			/* 8 */ imports.NewTable("TSynCustomFoldConfig_OnChange", 0), // event OnChange
		}
	})
	return synCustomFoldConfigImport
}
