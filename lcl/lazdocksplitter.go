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

// ILazDockSplitter Parent: ICustomSplitter
type ILazDockSplitter interface {
	ICustomSplitter
}

type TLazDockSplitter struct {
	TCustomSplitter
}

// NewLazDockSplitter class constructor
func NewLazDockSplitter(owner IComponent) ILazDockSplitter {
	r := lazDockSplitterAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsLazDockSplitter(r)
}

func TLazDockSplitterClass() types.TClass {
	r := lazDockSplitterAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	lazDockSplitterOnce   base.Once
	lazDockSplitterImport *imports.Imports = nil
)

func lazDockSplitterAPI() *imports.Imports {
	lazDockSplitterOnce.Do(func() {
		lazDockSplitterImport = api.NewDefaultImports()
		lazDockSplitterImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazDockSplitter_Create", 0), // constructor NewLazDockSplitter
			/* 1 */ imports.NewTable("TLazDockSplitter_TClass", 0), // function TLazDockSplitterClass
		}
	})
	return lazDockSplitterImport
}
