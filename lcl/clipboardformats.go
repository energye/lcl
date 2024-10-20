//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IClipboardFormats Parent: IStringList
type IClipboardFormats interface {
	IStringList
	Owner() IBaseVirtualTree // property
}

// TClipboardFormats Parent: TStringList
type TClipboardFormats struct {
	TStringList
}

func NewClipboardFormats(AOwner IBaseVirtualTree) IClipboardFormats {
	r1 := clipboardFormatsImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsClipboardFormats(r1)
}

func (m *TClipboardFormats) Owner() IBaseVirtualTree {
	r1 := clipboardFormatsImportAPI().SysCallN(2, m.Instance())
	return AsBaseVirtualTree(r1)
}

func ClipboardFormatsClass() TClass {
	ret := clipboardFormatsImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	clipboardFormatsImport       *imports.Imports = nil
	clipboardFormatsImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ClipboardFormats_Class", 0),
		/*1*/ imports.NewTable("ClipboardFormats_Create", 0),
		/*2*/ imports.NewTable("ClipboardFormats_Owner", 0),
	}
)

func clipboardFormatsImportAPI() *imports.Imports {
	if clipboardFormatsImport == nil {
		clipboardFormatsImport = NewDefaultImports()
		clipboardFormatsImport.SetImportTable(clipboardFormatsImportTables)
		clipboardFormatsImportTables = nil
	}
	return clipboardFormatsImport
}
