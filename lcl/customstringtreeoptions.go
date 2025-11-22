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

// ICustomStringTreeOptions Parent: ICustomVirtualTreeOptions
type ICustomStringTreeOptions interface {
	ICustomVirtualTreeOptions
}

type TCustomStringTreeOptions struct {
	TCustomVirtualTreeOptions
}

// NewCustomStringTreeOptions class constructor
func NewCustomStringTreeOptions(owner IBaseVirtualTree) ICustomStringTreeOptions {
	r := customStringTreeOptionsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomStringTreeOptions(r)
}

var (
	customStringTreeOptionsOnce   base.Once
	customStringTreeOptionsImport *imports.Imports = nil
)

func customStringTreeOptionsAPI() *imports.Imports {
	customStringTreeOptionsOnce.Do(func() {
		customStringTreeOptionsImport = api.NewDefaultImports()
		customStringTreeOptionsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomStringTreeOptions_Create", 0), // constructor NewCustomStringTreeOptions
		}
	})
	return customStringTreeOptionsImport
}
