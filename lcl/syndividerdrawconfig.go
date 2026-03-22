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

// ISynDividerDrawConfig Parent: IObject
type ISynDividerDrawConfig interface {
	IObject
	Assign(src ISynDividerDrawConfig) // procedure
	// TopSetting
	//  Do not use to set values, or you skip the change notification
	TopSetting() TSynDividerDrawConfigSetting  // property TopSetting Getter
	NestSetting() TSynDividerDrawConfigSetting // property NestSetting Getter
	MaxDrawDepth() int32                       // property MaxDrawDepth Getter
	SetMaxDrawDepth(value int32)               // property MaxDrawDepth Setter
	TopColor() types.TColor                    // property TopColor Getter
	SetTopColor(value types.TColor)            // property TopColor Setter
	NestColor() types.TColor                   // property NestColor Getter
	SetNestColor(value types.TColor)           // property NestColor Setter
	SetOnChange(fn TNotifyEvent)               // property event
}

type TSynDividerDrawConfig struct {
	TObject
}

func (m *TSynDividerDrawConfig) Assign(src ISynDividerDrawConfig) {
	if !m.IsValid() {
		return
	}
	synDividerDrawConfigAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(src))
}

func (m *TSynDividerDrawConfig) TopSetting() (result TSynDividerDrawConfigSetting) {
	if !m.IsValid() {
		return
	}
	synDividerDrawConfigAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynDividerDrawConfig) NestSetting() (result TSynDividerDrawConfigSetting) {
	if !m.IsValid() {
		return
	}
	synDividerDrawConfigAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynDividerDrawConfig) MaxDrawDepth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synDividerDrawConfigAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TSynDividerDrawConfig) SetMaxDrawDepth(value int32) {
	if !m.IsValid() {
		return
	}
	synDividerDrawConfigAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TSynDividerDrawConfig) TopColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synDividerDrawConfigAPI().SysCallN(5, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynDividerDrawConfig) SetTopColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synDividerDrawConfigAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TSynDividerDrawConfig) NestColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synDividerDrawConfigAPI().SysCallN(6, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynDividerDrawConfig) SetNestColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synDividerDrawConfigAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TSynDividerDrawConfig) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, synDividerDrawConfigAPI(), api.MakeEventDataPtr(cb))
}

// NewSynDividerDrawConfig class constructor
func NewSynDividerDrawConfig() ISynDividerDrawConfig {
	r := synDividerDrawConfigAPI().SysCallN(0)
	return AsSynDividerDrawConfig(r)
}

var (
	synDividerDrawConfigOnce   base.Once
	synDividerDrawConfigImport *imports.Imports = nil
)

func synDividerDrawConfigAPI() *imports.Imports {
	synDividerDrawConfigOnce.Do(func() {
		synDividerDrawConfigImport = api.NewDefaultImports()
		synDividerDrawConfigImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynDividerDrawConfig_Create", 0), // constructor NewSynDividerDrawConfig
			/* 1 */ imports.NewTable("TSynDividerDrawConfig_Assign", 0), // procedure Assign
			/* 2 */ imports.NewTable("TSynDividerDrawConfig_TopSetting", 0), // property TopSetting
			/* 3 */ imports.NewTable("TSynDividerDrawConfig_NestSetting", 0), // property NestSetting
			/* 4 */ imports.NewTable("TSynDividerDrawConfig_MaxDrawDepth", 0), // property MaxDrawDepth
			/* 5 */ imports.NewTable("TSynDividerDrawConfig_TopColor", 0), // property TopColor
			/* 6 */ imports.NewTable("TSynDividerDrawConfig_NestColor", 0), // property NestColor
			/* 7 */ imports.NewTable("TSynDividerDrawConfig_OnChange", 0), // event OnChange
		}
	})
	return synDividerDrawConfigImport
}
