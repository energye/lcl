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

// ISynGutterMarks Parent: ISynGutterPartBase
type ISynGutterMarks interface {
	ISynGutterPartBase
}

type TSynGutterMarks struct {
	TSynGutterPartBase
}

// NewSynGutterMarks class constructor
func NewSynGutterMarks(anOwner IComponent) ISynGutterMarks {
	r := synGutterMarksAPI().SysCallN(0, base.GetObjectUintptr(anOwner))
	return AsSynGutterMarks(r)
}

func TSynGutterMarksClass() types.TClass {
	r := synGutterMarksAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	synGutterMarksOnce   base.Once
	synGutterMarksImport *imports.Imports = nil
)

func synGutterMarksAPI() *imports.Imports {
	synGutterMarksOnce.Do(func() {
		synGutterMarksImport = api.NewDefaultImports()
		synGutterMarksImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynGutterMarks_Create", 0), // constructor NewSynGutterMarks
			/* 1 */ imports.NewTable("TSynGutterMarks_TClass", 0), // function TSynGutterMarksClass
		}
	})
	return synGutterMarksImport
}
