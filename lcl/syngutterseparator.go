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

// ISynGutterSeparator Parent: ISynGutterPartBase
type ISynGutterSeparator interface {
	ISynGutterPartBase
	LineWidth() int32          // property LineWidth Getter
	SetLineWidth(value int32)  // property LineWidth Setter
	LineOffset() int32         // property LineOffset Getter
	SetLineOffset(value int32) // property LineOffset Setter
	LineOnRight() bool         // property LineOnRight Getter
	SetLineOnRight(value bool) // property LineOnRight Setter
}

type TSynGutterSeparator struct {
	TSynGutterPartBase
}

func (m *TSynGutterSeparator) LineWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterSeparatorAPI().SysCallN(1, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterSeparator) SetLineWidth(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterSeparatorAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterSeparator) LineOffset() int32 {
	if !m.IsValid() {
		return 0
	}
	r := synGutterSeparatorAPI().SysCallN(2, 0, m.Instance())
	return int32(r)
}

func (m *TSynGutterSeparator) SetLineOffset(value int32) {
	if !m.IsValid() {
		return
	}
	synGutterSeparatorAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TSynGutterSeparator) LineOnRight() bool {
	if !m.IsValid() {
		return false
	}
	r := synGutterSeparatorAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TSynGutterSeparator) SetLineOnRight(value bool) {
	if !m.IsValid() {
		return
	}
	synGutterSeparatorAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

// NewSynGutterSeparator class constructor
func NewSynGutterSeparator(owner IComponent) ISynGutterSeparator {
	r := synGutterSeparatorAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynGutterSeparator(r)
}

func TSynGutterSeparatorClass() types.TClass {
	r := synGutterSeparatorAPI().SysCallN(4)
	return types.TClass(r)
}

var (
	synGutterSeparatorOnce   base.Once
	synGutterSeparatorImport *imports.Imports = nil
)

func synGutterSeparatorAPI() *imports.Imports {
	synGutterSeparatorOnce.Do(func() {
		synGutterSeparatorImport = api.NewDefaultImports()
		synGutterSeparatorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterSeparator_Create", 0), // constructor NewSynGutterSeparator
			/* 1 */ imports.NewTable("TSynGutterSeparator_LineWidth", 0), // property LineWidth
			/* 2 */ imports.NewTable("TSynGutterSeparator_LineOffset", 0), // property LineOffset
			/* 3 */ imports.NewTable("TSynGutterSeparator_LineOnRight", 0), // property LineOnRight
			/* 4 */ imports.NewTable("TSynGutterSeparator_TClass", 0), // function TSynGutterSeparatorClass
		}
	})
	return synGutterSeparatorImport
}
