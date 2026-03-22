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

// ISynGutterLineOverview Parent: ISynGutterPartBase
type ISynGutterLineOverview interface {
	ISynGutterPartBase
	AddMark(mark ISynGutterLOvMark)                     // procedure
	Providers() ISynGutterLineOverviewProviderList      // property Providers Getter
	MarkHeight() int32                                  // property MarkHeight Getter
	SetMarkHeight(value int32)                          // property MarkHeight Setter
	MouseActionsForMarks() ISynEditMouseActions         // property MouseActionsForMarks Getter
	SetMouseActionsForMarks(value ISynEditMouseActions) // property MouseActionsForMarks Setter
}

type TSynGutterLineOverview struct {
	TSynGutterPartBase
}

func (m *TSynGutterLineOverview) AddMark(mark ISynGutterLOvMark) {
	if !m.IsValid() {
		return
	}
	synGutterLineOverviewAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(mark))
}

func (m *TSynGutterLineOverview) Providers() ISynGutterLineOverviewProviderList {
	if !m.IsValid() {
		return nil
	}
	r := synGutterLineOverviewAPI().SysCallN(2, m.Instance())
	return AsSynGutterLineOverviewProviderList(r)
}

func (m *TSynGutterLineOverview) MarkHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLineOverviewAPI().SysCallN(3, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterLineOverview) SetMarkHeight(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterLineOverviewAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterLineOverview) MouseActionsForMarks() ISynEditMouseActions {
	if !m.IsValid() {
		return nil
	}
	r := synGutterLineOverviewAPI().SysCallN(4, 0, m.Instance())
	return AsSynEditMouseActions(r)
}

func (m *TSynGutterLineOverview) SetMouseActionsForMarks(value ISynEditMouseActions) {
	if !m.IsValid() {
		return
	}
	synGutterLineOverviewAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

// NewSynGutterLineOverview class constructor
func NewSynGutterLineOverview(owner IComponent) ISynGutterLineOverview {
	r := synGutterLineOverviewAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynGutterLineOverview(r)
}

func TSynGutterLineOverviewClass() types.TClass {
	r := synGutterLineOverviewAPI().SysCallN(5)
	return types.TClass(r)
}

var (
	synGutterLineOverviewOnce   base.Once
	synGutterLineOverviewImport *imports.Imports = nil
)

func synGutterLineOverviewAPI() *imports.Imports {
	synGutterLineOverviewOnce.Do(func() {
		synGutterLineOverviewImport = api.NewDefaultImports()
		synGutterLineOverviewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterLineOverview_Create", 0), // constructor NewSynGutterLineOverview
			/* 1 */ imports.NewTable("TSynGutterLineOverview_AddMark", 0), // procedure AddMark
			/* 2 */ imports.NewTable("TSynGutterLineOverview_Providers", 0), // property Providers
			/* 3 */ imports.NewTable("TSynGutterLineOverview_MarkHeight", 0), // property MarkHeight
			/* 4 */ imports.NewTable("TSynGutterLineOverview_MouseActionsForMarks", 0), // property MouseActionsForMarks
			/* 5 */ imports.NewTable("TSynGutterLineOverview_TClass", 0), // function TSynGutterLineOverviewClass
		}
	})
	return synGutterLineOverviewImport
}
