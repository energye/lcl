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

// ISynPluginSyncronizedEditBase Parent: ILazSynEditPlugin
type ISynPluginSyncronizedEditBase interface {
	ILazSynEditPlugin
	Clear()                                       // procedure
	Enabled() bool                                // property Enabled Getter
	SetEnabled(value bool)                        // property Enabled Setter
	Active() bool                                 // property Active Getter
	SetActive(value bool)                         // property Active Setter
	PreActive() bool                              // property PreActive Getter
	SetPreActive(value bool)                      // property PreActive Setter
	MarkupInfo() ISynSelectedColor                // property MarkupInfo Getter
	SetMarkupInfo(value ISynSelectedColor)        // property MarkupInfo Setter
	MarkupInfoCurrent() ISynSelectedColor         // property MarkupInfoCurrent Getter
	SetMarkupInfoCurrent(value ISynSelectedColor) // property MarkupInfoCurrent Setter
	MarkupInfoSync() ISynSelectedColor            // property MarkupInfoSync Getter
	SetMarkupInfoSync(value ISynSelectedColor)    // property MarkupInfoSync Setter
	MarkupInfoArea() ISynSelectedColor            // property MarkupInfoArea Getter
	SetMarkupInfoArea(value ISynSelectedColor)    // property MarkupInfoArea Setter
	SetOnActivate(fn TNotifyEvent)                // property event
	SetOnDeactivate(fn TNotifyEvent)              // property event
}

type TSynPluginSyncronizedEditBase struct {
	TLazSynEditPlugin
}

func (m *TSynPluginSyncronizedEditBase) Clear() {
	if !m.IsValid() {
		return
	}
	synPluginSyncronizedEditBaseAPI().SysCallN(1, m.Instance())
}

func (m *TSynPluginSyncronizedEditBase) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := synPluginSyncronizedEditBaseAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynPluginSyncronizedEditBase) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	synPluginSyncronizedEditBaseAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynPluginSyncronizedEditBase) Active() bool {
	if !m.IsValid() {
		return false
	}
	r := synPluginSyncronizedEditBaseAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynPluginSyncronizedEditBase) SetActive(value bool) {
	if !m.IsValid() {
		return
	}
	synPluginSyncronizedEditBaseAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynPluginSyncronizedEditBase) PreActive() bool {
	if !m.IsValid() {
		return false
	}
	r := synPluginSyncronizedEditBaseAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynPluginSyncronizedEditBase) SetPreActive(value bool) {
	if !m.IsValid() {
		return
	}
	synPluginSyncronizedEditBaseAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TSynPluginSyncronizedEditBase) MarkupInfo() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synPluginSyncronizedEditBaseAPI().SysCallN(5, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TSynPluginSyncronizedEditBase) SetMarkupInfo(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	synPluginSyncronizedEditBaseAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPluginSyncronizedEditBase) MarkupInfoCurrent() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synPluginSyncronizedEditBaseAPI().SysCallN(6, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TSynPluginSyncronizedEditBase) SetMarkupInfoCurrent(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	synPluginSyncronizedEditBaseAPI().SysCallN(6, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPluginSyncronizedEditBase) MarkupInfoSync() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synPluginSyncronizedEditBaseAPI().SysCallN(7, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TSynPluginSyncronizedEditBase) SetMarkupInfoSync(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	synPluginSyncronizedEditBaseAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPluginSyncronizedEditBase) MarkupInfoArea() ISynSelectedColor {
	if !m.IsValid() {
		return nil
	}
	r := synPluginSyncronizedEditBaseAPI().SysCallN(8, 0, m.Instance())
	return AsSynSelectedColor(r)
}

func (m *TSynPluginSyncronizedEditBase) SetMarkupInfoArea(value ISynSelectedColor) {
	if !m.IsValid() {
		return
	}
	synPluginSyncronizedEditBaseAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TSynPluginSyncronizedEditBase) SetOnActivate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 9, synPluginSyncronizedEditBaseAPI(), api.MakeEventDataPtr(cb))
}

func (m *TSynPluginSyncronizedEditBase) SetOnDeactivate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 10, synPluginSyncronizedEditBaseAPI(), api.MakeEventDataPtr(cb))
}

// NewSynPluginSyncronizedEditBase class constructor
func NewSynPluginSyncronizedEditBase(owner IComponent) ISynPluginSyncronizedEditBase {
	r := synPluginSyncronizedEditBaseAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynPluginSyncronizedEditBase(r)
}

func TSynPluginSyncronizedEditBaseClass() types.TClass {
	r := synPluginSyncronizedEditBaseAPI().SysCallN(11)
	return types.TClass(r)
}

var (
	synPluginSyncronizedEditBaseOnce   base.Once
	synPluginSyncronizedEditBaseImport *imports.Imports = nil
)

func synPluginSyncronizedEditBaseAPI() *imports.Imports {
	synPluginSyncronizedEditBaseOnce.Do(func() {
		synPluginSyncronizedEditBaseImport = api.NewDefaultImports()
		synPluginSyncronizedEditBaseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynPluginSyncronizedEditBase_Create", 0), // constructor NewSynPluginSyncronizedEditBase
			/* 1 */ imports.NewTable("TSynPluginSyncronizedEditBase_Clear", 0), // procedure Clear
			/* 2 */ imports.NewTable("TSynPluginSyncronizedEditBase_Enabled", 0), // property Enabled
			/* 3 */ imports.NewTable("TSynPluginSyncronizedEditBase_Active", 0), // property Active
			/* 4 */ imports.NewTable("TSynPluginSyncronizedEditBase_PreActive", 0), // property PreActive
			/* 5 */ imports.NewTable("TSynPluginSyncronizedEditBase_MarkupInfo", 0), // property MarkupInfo
			/* 6 */ imports.NewTable("TSynPluginSyncronizedEditBase_MarkupInfoCurrent", 0), // property MarkupInfoCurrent
			/* 7 */ imports.NewTable("TSynPluginSyncronizedEditBase_MarkupInfoSync", 0), // property MarkupInfoSync
			/* 8 */ imports.NewTable("TSynPluginSyncronizedEditBase_MarkupInfoArea", 0), // property MarkupInfoArea
			/* 9 */ imports.NewTable("TSynPluginSyncronizedEditBase_OnActivate", 0), // event OnActivate
			/* 10 */ imports.NewTable("TSynPluginSyncronizedEditBase_OnDeactivate", 0), // event OnDeactivate
			/* 11 */ imports.NewTable("TSynPluginSyncronizedEditBase_TClass", 0), // function TSynPluginSyncronizedEditBaseClass
		}
	})
	return synPluginSyncronizedEditBaseImport
}
