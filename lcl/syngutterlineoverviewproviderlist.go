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

// ISynGutterLineOverviewProviderList Parent: ISynObjectList
type ISynGutterLineOverviewProviderList interface {
	ISynObjectList
	OwnerToSynGutterLineOverview() ISynGutterLineOverview // property Owner Getter
	Providers(index int32) ISynGutterLineOverviewProvider // property Providers Getter
}

type TSynGutterLineOverviewProviderList struct {
	TSynObjectList
}

func (m *TSynGutterLineOverviewProviderList) OwnerToSynGutterLineOverview() ISynGutterLineOverview {
	if !m.IsValid() {
		return nil
	}
	r := synGutterLineOverviewProviderListAPI().SysCallN(1, m.Instance())
	return AsSynGutterLineOverview(r)
}

func (m *TSynGutterLineOverviewProviderList) Providers(index int32) ISynGutterLineOverviewProvider {
	if !m.IsValid() {
		return nil
	}
	r := synGutterLineOverviewProviderListAPI().SysCallN(2, m.Instance(), uintptr(index))
	return AsSynGutterLineOverviewProvider(r)
}

// NewSynGutterLineOverviewProviderList class constructor
func NewSynGutterLineOverviewProviderList(owner IComponent) ISynGutterLineOverviewProviderList {
	r := synGutterLineOverviewProviderListAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynGutterLineOverviewProviderList(r)
}

func TSynGutterLineOverviewProviderListClass() types.TClass {
	r := synGutterLineOverviewProviderListAPI().SysCallN(3)
	return types.TClass(r)
}

var (
	synGutterLineOverviewProviderListOnce   base.Once
	synGutterLineOverviewProviderListImport *imports.Imports = nil
)

func synGutterLineOverviewProviderListAPI() *imports.Imports {
	synGutterLineOverviewProviderListOnce.Do(func() {
		synGutterLineOverviewProviderListImport = api.NewDefaultImports()
		synGutterLineOverviewProviderListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterLineOverviewProviderList_Create", 0), // constructor NewSynGutterLineOverviewProviderList
			/* 1 */ imports.NewTable("TSynGutterLineOverviewProviderList_OwnerToSynGutterLineOverview", 0), // property OwnerToSynGutterLineOverview
			/* 2 */ imports.NewTable("TSynGutterLineOverviewProviderList_Providers", 0), // property Providers
			/* 3 */ imports.NewTable("TSynGutterLineOverviewProviderList_TClass", 0), // function TSynGutterLineOverviewProviderListClass
		}
	})
	return synGutterLineOverviewProviderListImport
}
