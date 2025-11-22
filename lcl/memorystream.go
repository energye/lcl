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

// IMemoryStream Parent: ICustomMemoryStream
type IMemoryStream interface {
	ICustomMemoryStream
	Clear()                         // procedure
	LoadFromStream(stream IStream)  // procedure
	LoadFromFile(fileName string)   // procedure
	SetSizeWithInt64(newSize int64) // procedure
}

type TMemoryStream struct {
	TCustomMemoryStream
}

func (m *TMemoryStream) Clear() {
	if !m.IsValid() {
		return
	}
	memoryStreamAPI().SysCallN(1, m.Instance())
}

func (m *TMemoryStream) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	memoryStreamAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TMemoryStream) LoadFromFile(fileName string) {
	if !m.IsValid() {
		return
	}
	memoryStreamAPI().SysCallN(3, m.Instance(), api.PasStr(fileName))
}

func (m *TMemoryStream) SetSizeWithInt64(newSize int64) {
	if !m.IsValid() {
		return
	}
	memoryStreamAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&newSize)))
}

// NewMemoryStream class constructor
func NewMemoryStream() IMemoryStream {
	r := memoryStreamAPI().SysCallN(0)
	return AsMemoryStream(r)
}

var (
	memoryStreamOnce   base.Once
	memoryStreamImport *imports.Imports = nil
)

func memoryStreamAPI() *imports.Imports {
	memoryStreamOnce.Do(func() {
		memoryStreamImport = api.NewDefaultImports()
		memoryStreamImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TMemoryStream_Create", 0), // constructor NewMemoryStream
			/* 1 */ imports.NewTable("TMemoryStream_Clear", 0), // procedure Clear
			/* 2 */ imports.NewTable("TMemoryStream_LoadFromStream", 0), // procedure LoadFromStream
			/* 3 */ imports.NewTable("TMemoryStream_LoadFromFile", 0), // procedure LoadFromFile
			/* 4 */ imports.NewTable("TMemoryStream_SetSizeWithInt64", 0), // procedure SetSizeWithInt64
		}
	})
	return memoryStreamImport
}
