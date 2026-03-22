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

// IByteArrayWrap Parent: IObject
type IByteArrayWrap interface {
	IObject
	ByteArrayData() types.PByteArray  // function
	Size() int32                      // function
	GetValue(index int32) byte        // function
	SetValue(index int32, value byte) // procedure
}

type TByteArrayWrap struct {
	TObject
}

func (m *TByteArrayWrap) ByteArrayData() types.PByteArray {
	if !m.IsValid() {
		return 0
	}
	r := byteArrayWrapAPI().SysCallN(1, m.Instance())
	return types.PByteArray(r)
}

func (m *TByteArrayWrap) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := byteArrayWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TByteArrayWrap) GetValue(index int32) byte {
	if !m.IsValid() {
		return 0
	}
	r := byteArrayWrapAPI().SysCallN(3, m.Instance(), uintptr(index))
	return byte(r)
}

func (m *TByteArrayWrap) SetValue(index int32, value byte) {
	if !m.IsValid() {
		return
	}
	byteArrayWrapAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(value))
}

// NewByteArrayWrapWithInt class constructor
func NewByteArrayWrapWithInt(size int32) IByteArrayWrap {
	r := byteArrayWrapAPI().SysCallN(0, uintptr(size))
	return AsByteArrayWrap(r)
}

var (
	byteArrayWrapOnce   base.Once
	byteArrayWrapImport *imports.Imports = nil
)

func byteArrayWrapAPI() *imports.Imports {
	byteArrayWrapOnce.Do(func() {
		byteArrayWrapImport = api.NewDefaultImports()
		byteArrayWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TByteArrayWrap_CreateWithInt", 0), // constructor NewByteArrayWrapWithInt
			/* 1 */ imports.NewTable("TByteArrayWrap_ByteArrayData", 0), // function ByteArrayData
			/* 2 */ imports.NewTable("TByteArrayWrap_Size", 0), // function Size
			/* 3 */ imports.NewTable("TByteArrayWrap_GetValue", 0), // function GetValue
			/* 4 */ imports.NewTable("TByteArrayWrap_SetValue", 0), // procedure SetValue
		}
	})
	return byteArrayWrapImport
}
