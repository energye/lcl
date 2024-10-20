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

// IMemoryStream Parent: ICustomMemoryStream
type IMemoryStream interface {
	ICustomMemoryStream
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Clear()                        // procedure
	LoadFromStream(Stream IStream) // procedure
	LoadFromFile(FileName string)  // procedure
}

// TMemoryStream Parent: TCustomMemoryStream
type TMemoryStream struct {
	TCustomMemoryStream
}

func NewMemoryStream() IMemoryStream {
	r1 := memoryStreamImportAPI().SysCallN(2)
	return AsMemoryStream(r1)
}

func MemoryStreamClass() TClass {
	ret := memoryStreamImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TMemoryStream) Clear() {
	memoryStreamImportAPI().SysCallN(1, m.Instance())
}

func (m *TMemoryStream) LoadFromStream(Stream IStream) {
	memoryStreamImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TMemoryStream) LoadFromFile(FileName string) {
	memoryStreamImportAPI().SysCallN(3, m.Instance(), PascalStr(FileName))
}

var (
	memoryStreamImport       *imports.Imports = nil
	memoryStreamImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("MemoryStream_Class", 0),
		/*1*/ imports.NewTable("MemoryStream_Clear", 0),
		/*2*/ imports.NewTable("MemoryStream_Create", 0),
		/*3*/ imports.NewTable("MemoryStream_LoadFromFile", 0),
		/*4*/ imports.NewTable("MemoryStream_LoadFromStream", 0),
	}
)

func memoryStreamImportAPI() *imports.Imports {
	if memoryStreamImport == nil {
		memoryStreamImport = NewDefaultImports()
		memoryStreamImport.SetImportTable(memoryStreamImportTables)
		memoryStreamImportTables = nil
	}
	return memoryStreamImport
}
