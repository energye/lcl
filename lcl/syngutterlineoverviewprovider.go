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

// ISynGutterLineOverviewProvider Parent: ISynObjectListItem
type ISynGutterLineOverviewProvider interface {
	ISynObjectListItem
	Height() int32               // property Height Getter
	SetHeight(value int32)       // property Height Setter
	Priority() int32             // property Priority Getter
	SetPriority(value int32)     // property Priority Setter
	Color() types.TColor         // property Color Getter
	SetColor(value types.TColor) // property Color Setter
}

type TSynGutterLineOverviewProvider struct {
	TSynObjectListItem
}

func (m *TSynGutterLineOverviewProvider) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLineOverviewProviderAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterLineOverviewProvider) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterLineOverviewProviderAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterLineOverviewProvider) Priority() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLineOverviewProviderAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterLineOverviewProvider) SetPriority(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterLineOverviewProviderAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterLineOverviewProvider) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := synGutterLineOverviewProviderAPI().SysCallN(3, 0, m.Instance())
	return types.TColor(r)
}

func (m *TSynGutterLineOverviewProvider) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	synGutterLineOverviewProviderAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

// NewSynGutterLineOverviewProvider class constructor
func NewSynGutterLineOverviewProvider(owner IComponent) ISynGutterLineOverviewProvider {
	r := synGutterLineOverviewProviderAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynGutterLineOverviewProvider(r)
}

func TSynGutterLineOverviewProviderClass() types.TClass {
	r := synGutterLineOverviewProviderAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	synGutterLineOverviewProviderOnce   base.Once
	synGutterLineOverviewProviderImport *imports.Imports = nil
)

func synGutterLineOverviewProviderAPI() *imports.Imports {
	synGutterLineOverviewProviderOnce.Do(func() {
		synGutterLineOverviewProviderImport = api.NewDefaultImports()
		synGutterLineOverviewProviderImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterLineOverviewProvider_Create", 0), // constructor NewSynGutterLineOverviewProvider
			/* 1 */ imports.NewTable("TSynGutterLineOverviewProvider_Height", 0), // property Height
			/* 2 */ imports.NewTable("TSynGutterLineOverviewProvider_Priority", 0), // property Priority
			/* 3 */ imports.NewTable("TSynGutterLineOverviewProvider_Color", 0), // property Color
			/* 4 */ imports.NewTable("TSynGutterLineOverviewProvider_TClass", 0), // function TSynGutterLineOverviewProviderClass
		}
	})
	return synGutterLineOverviewProviderImport
}
