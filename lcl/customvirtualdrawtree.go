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

// ICustomVirtualDrawTree Parent: IBaseVirtualTree
type ICustomVirtualDrawTree interface {
	IBaseVirtualTree
}

type TCustomVirtualDrawTree struct {
	TBaseVirtualTree
}

// NewCustomVirtualDrawTree class constructor
func NewCustomVirtualDrawTree(owner IComponent) ICustomVirtualDrawTree {
	r := customVirtualDrawTreeAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomVirtualDrawTree(r)
}

func TCustomVirtualDrawTreeClass() types.TClass {
	r := customVirtualDrawTreeAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	customVirtualDrawTreeOnce   base.Once
	customVirtualDrawTreeImport *imports.Imports = nil
)

func customVirtualDrawTreeAPI() *imports.Imports {
	customVirtualDrawTreeOnce.Do(func() {
		customVirtualDrawTreeImport = api.NewDefaultImports()
		customVirtualDrawTreeImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomVirtualDrawTree_Create", 0), // constructor NewCustomVirtualDrawTree
			/* 1 */ imports.NewTable("TCustomVirtualDrawTree_TClass", 0), // function TCustomVirtualDrawTreeClass
		}
	})
	return customVirtualDrawTreeImport
}
