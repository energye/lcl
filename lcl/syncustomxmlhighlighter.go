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

// ISynCustomXmlHighlighter Parent: ISynCustomFoldHighlighter
type ISynCustomXmlHighlighter interface {
	ISynCustomFoldHighlighter
}

type TSynCustomXmlHighlighter struct {
	TSynCustomFoldHighlighter
}

// NewSynCustomXmlHighlighter class constructor
func NewSynCustomXmlHighlighter(owner IComponent) ISynCustomXmlHighlighter {
	r := synCustomXmlHighlighterAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsSynCustomXmlHighlighter(r)
}

func TSynCustomXmlHighlighterClass() types.TClass {
	r := synCustomXmlHighlighterAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	synCustomXmlHighlighterOnce   base.Once
	synCustomXmlHighlighterImport *imports.Imports = nil
)

func synCustomXmlHighlighterAPI() *imports.Imports {
	synCustomXmlHighlighterOnce.Do(func() {
		synCustomXmlHighlighterImport = api.NewDefaultImports()
		synCustomXmlHighlighterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynCustomXmlHighlighter_Create", 0), // constructor NewSynCustomXmlHighlighter
			/* 1 */ imports.NewTable("TSynCustomXmlHighlighter_TClass", 0), // function TSynCustomXmlHighlighterClass
		}
	})
	return synCustomXmlHighlighterImport
}
