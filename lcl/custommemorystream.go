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

// ICustomMemoryStream Parent: IStream
type ICustomMemoryStream interface {
	IStream
	Memory() uintptr             // property
	SaveToStream(Stream IStream) // procedure
	SaveToFile(FileName string)  // procedure
}

// TCustomMemoryStream Parent: TStream
type TCustomMemoryStream struct {
	TStream
}

func NewCustomMemoryStream() ICustomMemoryStream {
	r1 := customMemoryStreamImportAPI().SysCallN(1)
	return AsCustomMemoryStream(r1)
}

func (m *TCustomMemoryStream) Memory() uintptr {
	r1 := customMemoryStreamImportAPI().SysCallN(2, m.Instance())
	return uintptr(r1)
}

func CustomMemoryStreamClass() TClass {
	ret := customMemoryStreamImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomMemoryStream) SaveToStream(Stream IStream) {
	customMemoryStreamImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TCustomMemoryStream) SaveToFile(FileName string) {
	customMemoryStreamImportAPI().SysCallN(3, m.Instance(), PascalStr(FileName))
}

var (
	customMemoryStreamImport       *imports.Imports = nil
	customMemoryStreamImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomMemoryStream_Class", 0),
		/*1*/ imports.NewTable("CustomMemoryStream_Create", 0),
		/*2*/ imports.NewTable("CustomMemoryStream_Memory", 0),
		/*3*/ imports.NewTable("CustomMemoryStream_SaveToFile", 0),
		/*4*/ imports.NewTable("CustomMemoryStream_SaveToStream", 0),
	}
)

func customMemoryStreamImportAPI() *imports.Imports {
	if customMemoryStreamImport == nil {
		customMemoryStreamImport = NewDefaultImports()
		customMemoryStreamImport.SetImportTable(customMemoryStreamImportTables)
		customMemoryStreamImportTables = nil
	}
	return customMemoryStreamImport
}
