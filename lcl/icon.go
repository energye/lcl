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

// IIcon Parent: ICustomIcon
type IIcon interface {
	ICustomIcon
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Handle() HICON          // property
	SetHandle(AValue HICON) // property
	ReleaseHandle() HICON   // function
}

// TIcon Parent: TCustomIcon
type TIcon struct {
	TCustomIcon
}

func NewIcon() IIcon {
	r1 := conImportAPI().SysCallN(1)
	return AsIcon(r1)
}

func (m *TIcon) Handle() HICON {
	r1 := conImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return HICON(r1)
}

func (m *TIcon) SetHandle(AValue HICON) {
	conImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TIcon) ReleaseHandle() HICON {
	r1 := conImportAPI().SysCallN(3, m.Instance())
	return HICON(r1)
}

func IconClass() TClass {
	ret := conImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	conImport       *imports.Imports = nil
	conImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Icon_Class", 0),
		/*1*/ imports.NewTable("Icon_Create", 0),
		/*2*/ imports.NewTable("Icon_Handle", 0),
		/*3*/ imports.NewTable("Icon_ReleaseHandle", 0),
	}
)

func conImportAPI() *imports.Imports {
	if conImport == nil {
		conImport = NewDefaultImports()
		conImport.SetImportTable(conImportTables)
		conImportTables = nil
	}
	return conImport
}
