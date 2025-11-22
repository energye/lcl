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

// IClipboardFormats Parent: IStringList
type IClipboardFormats interface {
	IStringList
	Owner() IBaseVirtualTree // property Owner Getter
}

type TClipboardFormats struct {
	TStringList
}

func (m *TClipboardFormats) Owner() IBaseVirtualTree {
	if !m.IsValid() {
		return nil
	}
	r := clipboardFormatsAPI().SysCallN(1, m.Instance())
	return AsBaseVirtualTree(r)
}

// NewClipboardFormats class constructor
func NewClipboardFormats(owner IBaseVirtualTree) IClipboardFormats {
	r := clipboardFormatsAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsClipboardFormats(r)
}

var (
	clipboardFormatsOnce   base.Once
	clipboardFormatsImport *imports.Imports = nil
)

func clipboardFormatsAPI() *imports.Imports {
	clipboardFormatsOnce.Do(func() {
		clipboardFormatsImport = api.NewDefaultImports()
		clipboardFormatsImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TClipboardFormats_Create", 0), // constructor NewClipboardFormats
			/* 1 */ imports.NewTable("TClipboardFormats_Owner", 0), // property Owner
		}
	})
	return clipboardFormatsImport
}
