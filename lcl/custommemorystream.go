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

// ICustomMemoryStream Parent: IStream
type ICustomMemoryStream interface {
	IStream
	SaveToStream(stream IStream) // procedure
	SaveToFile(fileName string)  // procedure
	Memory() uintptr             // property Memory Getter
}

type TCustomMemoryStream struct {
	TStream
}

func (m *TCustomMemoryStream) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	customMemoryStreamAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCustomMemoryStream) SaveToFile(fileName string) {
	if !m.IsValid() {
		return
	}
	customMemoryStreamAPI().SysCallN(2, m.Instance(), api.PasStr(fileName))
}

func (m *TCustomMemoryStream) Memory() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := customMemoryStreamAPI().SysCallN(3, m.Instance())
	return uintptr(r)
}

// NewCustomMemoryStream class constructor
func NewCustomMemoryStream() ICustomMemoryStream {
	r := customMemoryStreamAPI().SysCallN(0)
	return AsCustomMemoryStream(r)
}

var (
	customMemoryStreamOnce   base.Once
	customMemoryStreamImport *imports.Imports = nil
)

func customMemoryStreamAPI() *imports.Imports {
	customMemoryStreamOnce.Do(func() {
		customMemoryStreamImport = api.NewDefaultImports()
		customMemoryStreamImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomMemoryStream_Create", 0), // constructor NewCustomMemoryStream
			/* 1 */ imports.NewTable("TCustomMemoryStream_SaveToStream", 0), // procedure SaveToStream
			/* 2 */ imports.NewTable("TCustomMemoryStream_SaveToFile", 0), // procedure SaveToFile
			/* 3 */ imports.NewTable("TCustomMemoryStream_Memory", 0), // property Memory
		}
	})
	return customMemoryStreamImport
}
