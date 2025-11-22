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

// IEbEdit Parent: IGEEdit
type IEbEdit interface {
	IGEEdit
}

type TEbEdit struct {
	TGEEdit
}

// NewEbEdit class constructor
func NewEbEdit(theOwner IComponent) IEbEdit {
	r := ebEditAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsEbEdit(r)
}

func TEbEditClass() types.TClass {
	r := ebEditAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	ebEditOnce   base.Once
	ebEditImport *imports.Imports = nil
)

func ebEditAPI() *imports.Imports {
	ebEditOnce.Do(func() {
		ebEditImport = api.NewDefaultImports()
		ebEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEbEdit_Create", 0), // constructor NewEbEdit
			/* 1 */ imports.NewTable("TEbEdit_TClass", 0), // function TEbEditClass
		}
	})
	return ebEditImport
}
