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

// ISynGutterPartListBase Parent: ISynObjectList
type ISynGutterPartListBase interface {
	ISynObjectList
	Part(index int32) ISynGutterPartBase                                         // property Part Getter
	SetPart(index int32, value ISynGutterPartBase)                               // property Part Setter
	ByClassCount(class types.TSynGutterPartBaseClass) int32                      // property ByClassCount Getter
	ByClass(class types.TSynGutterPartBaseClass, index int32) ISynGutterPartBase // property ByClass Getter
}

type TSynGutterPartListBase struct {
	TSynObjectList
}

func (m *TSynGutterPartListBase) Part(index int32) ISynGutterPartBase {
	if !m.IsValid() {
		return nil
	}
	r := synGutterPartListBaseAPI().SysCallN(2, 0, m.Instance(), uintptr(index))
	return AsSynGutterPartBase(r)
}

func (m *TSynGutterPartListBase) SetPart(index int32, value ISynGutterPartBase) {
	if !m.IsValid() {
		return
	}
	synGutterPartListBaseAPI().SysCallN(2, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

func (m *TSynGutterPartListBase) ByClassCount(class types.TSynGutterPartBaseClass) int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterPartListBaseAPI().SysCallN(3, m.Instance(), uintptr(class))
	return int32(r)
}

func (m *TSynGutterPartListBase) ByClass(class types.TSynGutterPartBaseClass, index int32) ISynGutterPartBase {
	if !m.IsValid() {
		return nil
	}
	r := synGutterPartListBaseAPI().SysCallN(4, m.Instance(), uintptr(class), uintptr(index))
	return AsSynGutterPartBase(r)
}

// NewSynGutterPartListBase class constructor
func NewSynGutterPartListBase(owner IComponent) ISynGutterPartListBase {
	r := synGutterPartListBaseAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynGutterPartListBase(r)
}

// NewSynGutterPartListBaseWithComponentSynGutterBase class constructor
func NewSynGutterPartListBaseWithComponentSynGutterBase(owner IComponent, gutter ISynGutterBase) ISynGutterPartListBase {
	r := synGutterPartListBaseAPI().SysCallN(1, base.GetObjectUintptr(owner), base.GetObjectUintptr(gutter))
	return AsSynGutterPartListBase(r)
}

func TSynGutterPartListBaseClass() types.TClass {
	r := synGutterPartListBaseAPI().SysCallN(5)
	return types.TClass(r)
}

var (
	synGutterPartListBaseOnce   base.Once
	synGutterPartListBaseImport *imports.Imports = nil
)

func synGutterPartListBaseAPI() *imports.Imports {
	synGutterPartListBaseOnce.Do(func() {
		synGutterPartListBaseImport = api.NewDefaultImports()
		synGutterPartListBaseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterPartListBase_Create", 0), // constructor NewSynGutterPartListBase
			/* 1 */ imports.NewTable("TSynGutterPartListBase_CreateWithComponentSynGutterBase", 0), // constructor NewSynGutterPartListBaseWithComponentSynGutterBase
			/* 2 */ imports.NewTable("TSynGutterPartListBase_Part", 0), // property Part
			/* 3 */ imports.NewTable("TSynGutterPartListBase_ByClassCount", 0), // property ByClassCount
			/* 4 */ imports.NewTable("TSynGutterPartListBase_ByClass", 0), // property ByClass
			/* 5 */ imports.NewTable("TSynGutterPartListBase_TClass", 0), // function TSynGutterPartListBaseClass
		}
	})
	return synGutterPartListBaseImport
}
