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

// ICustomVirtualTreeOptions Parent: IPersistent
type ICustomVirtualTreeOptions interface {
	IPersistent
	AssignTo(dest IPersistent) // procedure
	Owner() IBaseVirtualTree   // property Owner Getter
}

type TCustomVirtualTreeOptions struct {
	TPersistent
}

func (m *TCustomVirtualTreeOptions) AssignTo(dest IPersistent) {
	if !m.IsValid() {
		return
	}
	customVirtualTreeOptionsAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(dest))
}

func (m *TCustomVirtualTreeOptions) Owner() IBaseVirtualTree {
	if !m.IsValid() {
		return nil
	}
	r := customVirtualTreeOptionsAPI().SysCallN(2, m.Instance())
	return AsBaseVirtualTree(r)
}

// NewCustomVirtualTreeOptions class constructor
func NewCustomVirtualTreeOptions(owner IBaseVirtualTree) ICustomVirtualTreeOptions {
	r := customVirtualTreeOptionsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomVirtualTreeOptions(r)
}

var (
	customVirtualTreeOptionsOnce   base.Once
	customVirtualTreeOptionsImport *imports.Imports = nil
)

func customVirtualTreeOptionsAPI() *imports.Imports {
	customVirtualTreeOptionsOnce.Do(func() {
		customVirtualTreeOptionsImport = api.NewDefaultImports()
		customVirtualTreeOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomVirtualTreeOptions_Create", 0), // constructor NewCustomVirtualTreeOptions
			/* 1 */ imports.NewTable("TCustomVirtualTreeOptions_AssignTo", 0), // procedure AssignTo
			/* 2 */ imports.NewTable("TCustomVirtualTreeOptions_Owner", 0), // property Owner
		}
	})
	return customVirtualTreeOptionsImport
}
