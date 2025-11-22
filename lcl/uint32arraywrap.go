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

// IUint32ArrayWrap Parent: IObject
type IUint32ArrayWrap interface {
	IObject
	Uint32ArrayData() types.PUint32Array // function
	Size() int32                         // function
	GetValue(index int32) uint32         // function
	SetValue(index int32, value uint32)  // procedure
}

type TUint32ArrayWrap struct {
	TObject
}

func (m *TUint32ArrayWrap) Uint32ArrayData() types.PUint32Array {
	if !m.IsValid() {
		return 0
	}
	r := uint32ArrayWrapAPI().SysCallN(2, m.Instance())
	return types.PUint32Array(r)
}

func (m *TUint32ArrayWrap) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := uint32ArrayWrapAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TUint32ArrayWrap) GetValue(index int32) uint32 {
	if !m.IsValid() {
		return 0
	}
	r := uint32ArrayWrapAPI().SysCallN(4, m.Instance(), uintptr(index))
	return uint32(r)
}

func (m *TUint32ArrayWrap) SetValue(index int32, value uint32) {
	if !m.IsValid() {
		return
	}
	uint32ArrayWrapAPI().SysCallN(5, m.Instance(), uintptr(index), uintptr(value))
}

// NewUint32ArrayWrap class constructor
func NewUint32ArrayWrap(data IUint32ArrayWrap) IUint32ArrayWrap {
	r := uint32ArrayWrapAPI().SysCallN(0, base.GetObjectUintptr(data))
	return AsUint32ArrayWrap(r)
}

// NewUint32ArrayWrapWithInt class constructor
func NewUint32ArrayWrapWithInt(size int32) IUint32ArrayWrap {
	r := uint32ArrayWrapAPI().SysCallN(1, uintptr(size))
	return AsUint32ArrayWrap(r)
}

var (
	uint32ArrayWrapOnce   base.Once
	uint32ArrayWrapImport *imports.Imports = nil
)

func uint32ArrayWrapAPI() *imports.Imports {
	uint32ArrayWrapOnce.Do(func() {
		uint32ArrayWrapImport = api.NewDefaultImports()
		uint32ArrayWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TUint32ArrayWrap_Create", 0), // constructor NewUint32ArrayWrap
			/* 1 */ imports.NewTable("TUint32ArrayWrap_CreateWithInt", 0), // constructor NewUint32ArrayWrapWithInt
			/* 2 */ imports.NewTable("TUint32ArrayWrap_Uint32ArrayData", 0), // function Uint32ArrayData
			/* 3 */ imports.NewTable("TUint32ArrayWrap_Size", 0), // function Size
			/* 4 */ imports.NewTable("TUint32ArrayWrap_GetValue", 0), // function GetValue
			/* 5 */ imports.NewTable("TUint32ArrayWrap_SetValue", 0), // procedure SetValue
		}
	})
	return uint32ArrayWrapImport
}
