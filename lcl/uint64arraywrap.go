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

// IUint64ArrayWrap Parent: IObject
type IUint64ArrayWrap interface {
	IObject
	Uint64ArrayData() types.PUint64Array // function
	Size() int32                         // function
	GetValue(index int32) uint64         // function
	SetValue(index int32, value uint64)  // procedure
}

type TUint64ArrayWrap struct {
	TObject
}

func (m *TUint64ArrayWrap) Uint64ArrayData() types.PUint64Array {
	if !m.IsValid() {
		return 0
	}
	r := uint64ArrayWrapAPI().SysCallN(1, m.Instance())
	return types.PUint64Array(r)
}

func (m *TUint64ArrayWrap) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := uint64ArrayWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TUint64ArrayWrap) GetValue(index int32) (result uint64) {
	if !m.IsValid() {
		return
	}
	uint64ArrayWrapAPI().SysCallN(3, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TUint64ArrayWrap) SetValue(index int32, value uint64) {
	if !m.IsValid() {
		return
	}
	uint64ArrayWrapAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&value)))
}

// NewUint64ArrayWrapWithInt class constructor
func NewUint64ArrayWrapWithInt(size int32) IUint64ArrayWrap {
	r := uint64ArrayWrapAPI().SysCallN(0, uintptr(size))
	return AsUint64ArrayWrap(r)
}

var (
	uint64ArrayWrapOnce   base.Once
	uint64ArrayWrapImport *imports.Imports = nil
)

func uint64ArrayWrapAPI() *imports.Imports {
	uint64ArrayWrapOnce.Do(func() {
		uint64ArrayWrapImport = api.NewDefaultImports()
		uint64ArrayWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TUint64ArrayWrap_CreateWithInt", 0), // constructor NewUint64ArrayWrapWithInt
			/* 1 */ imports.NewTable("TUint64ArrayWrap_Uint64ArrayData", 0), // function Uint64ArrayData
			/* 2 */ imports.NewTable("TUint64ArrayWrap_Size", 0), // function Size
			/* 3 */ imports.NewTable("TUint64ArrayWrap_GetValue", 0), // function GetValue
			/* 4 */ imports.NewTable("TUint64ArrayWrap_SetValue", 0), // procedure SetValue
		}
	})
	return uint64ArrayWrapImport
}
