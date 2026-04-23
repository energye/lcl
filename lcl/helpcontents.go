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

// IHelpContents Parent: IHelpAction
type IHelpContents interface {
	IHelpAction
}

type THelpContents struct {
	THelpAction
}

// NewHelpContents class constructor
func NewHelpContents(theOwner IComponent) IHelpContents {
	r := helpContentsAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsHelpContents(r)
}

func THelpContentsClass() types.TClass {
	r := helpContentsAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	helpContentsOnce   base.Once
	helpContentsImport *imports.Imports = nil
)

func helpContentsAPI() *imports.Imports {
	helpContentsOnce.Do(func() {
		helpContentsImport = api.NewDefaultImports()
		helpContentsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("THelpContents_Create", 0), // constructor NewHelpContents
			/* 1 */ imports.NewTable("THelpContents_TClass", 0), // function THelpContentsClass
		}
	})
	return helpContentsImport
}
