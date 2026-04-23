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

// IHelpTopicSearch Parent: IHelpAction
type IHelpTopicSearch interface {
	IHelpAction
}

type THelpTopicSearch struct {
	THelpAction
}

// NewHelpTopicSearch class constructor
func NewHelpTopicSearch(theOwner IComponent) IHelpTopicSearch {
	r := helpTopicSearchAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsHelpTopicSearch(r)
}

func THelpTopicSearchClass() types.TClass {
	r := helpTopicSearchAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	helpTopicSearchOnce   base.Once
	helpTopicSearchImport *imports.Imports = nil
)

func helpTopicSearchAPI() *imports.Imports {
	helpTopicSearchOnce.Do(func() {
		helpTopicSearchImport = api.NewDefaultImports()
		helpTopicSearchImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("THelpTopicSearch_Create", 0), // constructor NewHelpTopicSearch
			/* 1 */ imports.NewTable("THelpTopicSearch_TClass", 0), // function THelpTopicSearchClass
		}
	})
	return helpTopicSearchImport
}
