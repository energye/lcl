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

// IIconOptions Parent: IPersistent
type IIconOptions interface {
	IPersistent
	Arrangement() types.TIconArrangement         // property Arrangement Getter
	SetArrangement(value types.TIconArrangement) // property Arrangement Setter
	AutoArrange() bool                           // property AutoArrange Getter
	SetAutoArrange(value bool)                   // property AutoArrange Setter
	WrapText() bool                              // property WrapText Getter
	SetWrapText(value bool)                      // property WrapText Setter
}

type TIconOptions struct {
	TPersistent
}

func (m *TIconOptions) Arrangement() types.TIconArrangement {
	if !m.IsValid() {
		return 0
	}
	r := iconOptionsAPI().SysCallN(1, 0, m.Instance())
	return types.TIconArrangement(r)
}

func (m *TIconOptions) SetArrangement(value types.TIconArrangement) {
	if !m.IsValid() {
		return
	}
	iconOptionsAPI().SysCallN(1, 1, m.Instance(), uintptr(value))
}

func (m *TIconOptions) AutoArrange() bool {
	if !m.IsValid() {
		return false
	}
	r := iconOptionsAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TIconOptions) SetAutoArrange(value bool) {
	if !m.IsValid() {
		return
	}
	iconOptionsAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TIconOptions) WrapText() bool {
	if !m.IsValid() {
		return false
	}
	r := iconOptionsAPI().SysCallN(3, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TIconOptions) SetWrapText(value bool) {
	if !m.IsValid() {
		return
	}
	iconOptionsAPI().SysCallN(3, 1, m.Instance(), api.PasBool(value))
}

// NewIconOptions class constructor
func NewIconOptions(owner ICustomListView) IIconOptions {
	r := iconOptionsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsIconOptions(r)
}

var (
	iconOptionsOnce   base.Once
	iconOptionsImport *imports.Imports = nil
)

func iconOptionsAPI() *imports.Imports {
	iconOptionsOnce.Do(func() {
		iconOptionsImport = api.NewDefaultImports()
		iconOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TIconOptions_Create", 0), // constructor NewIconOptions
			/* 1 */ imports.NewTable("TIconOptions_Arrangement", 0), // property Arrangement
			/* 2 */ imports.NewTable("TIconOptions_AutoArrange", 0), // property AutoArrange
			/* 3 */ imports.NewTable("TIconOptions_WrapText", 0), // property WrapText
		}
	})
	return iconOptionsImport
}
