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

// ISynScrollOnEditOptions Parent: IPersistent
type ISynScrollOnEditOptions interface {
	IPersistent
	SetDefaults()                             // procedure
	KeepBorderDistance() int32                // property KeepBorderDistance Getter
	SetKeepBorderDistance(value int32)        // property KeepBorderDistance Setter
	KeepBorderDistancePercent() int32         // property KeepBorderDistancePercent Getter
	SetKeepBorderDistancePercent(value int32) // property KeepBorderDistancePercent Setter
	ScrollExtraColumns() int32                // property ScrollExtraColumns Getter
	SetScrollExtraColumns(value int32)        // property ScrollExtraColumns Setter
	ScrollExtraPercent() int32                // property ScrollExtraPercent Getter
	SetScrollExtraPercent(value int32)        // property ScrollExtraPercent Setter
	ScrollExtraMax() int32                    // property ScrollExtraMax Getter
	SetScrollExtraMax(value int32)            // property ScrollExtraMax Setter
	SetOnChange(fn TNotifyEvent)              // property event
}

type TSynScrollOnEditOptions struct {
	TPersistent
}

func (m *TSynScrollOnEditOptions) SetDefaults() {
	if !m.IsValid() {
		return
	}
	synScrollOnEditOptionsAPI().SysCallN(1, m.Instance())
}

func (m *TSynScrollOnEditOptions) KeepBorderDistance() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synScrollOnEditOptionsAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TSynScrollOnEditOptions) SetKeepBorderDistance(value int32) {
	if !m.IsValid() {
		return
	}
	synScrollOnEditOptionsAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TSynScrollOnEditOptions) KeepBorderDistancePercent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synScrollOnEditOptionsAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TSynScrollOnEditOptions) SetKeepBorderDistancePercent(value int32) {
	if !m.IsValid() {
		return
	}
	synScrollOnEditOptionsAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TSynScrollOnEditOptions) ScrollExtraColumns() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synScrollOnEditOptionsAPI().SysCallN(4, 0, m.Instance())
	return int32(r)
}

func (m *TSynScrollOnEditOptions) SetScrollExtraColumns(value int32) {
	if !m.IsValid() {
		return
	}
	synScrollOnEditOptionsAPI().SysCallN(4, 1, m.Instance(), uintptr(value))
}

func (m *TSynScrollOnEditOptions) ScrollExtraPercent() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synScrollOnEditOptionsAPI().SysCallN(5, 0, m.Instance())
	return int32(r)
}

func (m *TSynScrollOnEditOptions) SetScrollExtraPercent(value int32) {
	if !m.IsValid() {
		return
	}
	synScrollOnEditOptionsAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TSynScrollOnEditOptions) ScrollExtraMax() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synScrollOnEditOptionsAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TSynScrollOnEditOptions) SetScrollExtraMax(value int32) {
	if !m.IsValid() {
		return
	}
	synScrollOnEditOptionsAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TSynScrollOnEditOptions) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 7, synScrollOnEditOptionsAPI(), api.MakeEventDataPtr(cb))
}

// NewSynScrollOnEditOptions class constructor
func NewSynScrollOnEditOptions() ISynScrollOnEditOptions {
	r := synScrollOnEditOptionsAPI().SysCallN(0)
	return AsSynScrollOnEditOptions(r)
}

var (
	synScrollOnEditOptionsOnce   base.Once
	synScrollOnEditOptionsImport *imports.Imports = nil
)

func synScrollOnEditOptionsAPI() *imports.Imports {
	synScrollOnEditOptionsOnce.Do(func() {
		synScrollOnEditOptionsImport = api.NewDefaultImports()
		synScrollOnEditOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynScrollOnEditOptions_Create", 0), // constructor NewSynScrollOnEditOptions
			/* 1 */ imports.NewTable("TSynScrollOnEditOptions_SetDefaults", 0), // procedure SetDefaults
			/* 2 */ imports.NewTable("TSynScrollOnEditOptions_KeepBorderDistance", 0), // property KeepBorderDistance
			/* 3 */ imports.NewTable("TSynScrollOnEditOptions_KeepBorderDistancePercent", 0), // property KeepBorderDistancePercent
			/* 4 */ imports.NewTable("TSynScrollOnEditOptions_ScrollExtraColumns", 0), // property ScrollExtraColumns
			/* 5 */ imports.NewTable("TSynScrollOnEditOptions_ScrollExtraPercent", 0), // property ScrollExtraPercent
			/* 6 */ imports.NewTable("TSynScrollOnEditOptions_ScrollExtraMax", 0), // property ScrollExtraMax
			/* 7 */ imports.NewTable("TSynScrollOnEditOptions_OnChange", 0), // event OnChange
		}
	})
	return synScrollOnEditOptionsImport
}
