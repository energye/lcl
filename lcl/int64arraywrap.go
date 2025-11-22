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

// IInt64ArrayWrap Parent: IObject
type IInt64ArrayWrap interface {
	IObject
	Int64ArrayData() types.PInt64Array // function
	Size() int32                       // function
	GetValue(index int32) int64        // function
	SetValue(index int32, value int64) // procedure
}

type TInt64ArrayWrap struct {
	TObject
}

func (m *TInt64ArrayWrap) Int64ArrayData() types.PInt64Array {
	if !m.IsValid() {
		return 0
	}
	r := int64ArrayWrapAPI().SysCallN(1, m.Instance())
	return types.PInt64Array(r)
}

func (m *TInt64ArrayWrap) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := int64ArrayWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TInt64ArrayWrap) GetValue(index int32) (result int64) {
	if !m.IsValid() {
		return
	}
	int64ArrayWrapAPI().SysCallN(3, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TInt64ArrayWrap) SetValue(index int32, value int64) {
	if !m.IsValid() {
		return
	}
	int64ArrayWrapAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&value)))
}

// NewInt64ArrayWrapWithInt class constructor
func NewInt64ArrayWrapWithInt(size int32) IInt64ArrayWrap {
	r := int64ArrayWrapAPI().SysCallN(0, uintptr(size))
	return AsInt64ArrayWrap(r)
}

var (
	int64ArrayWrapOnce   base.Once
	int64ArrayWrapImport *imports.Imports = nil
)

func int64ArrayWrapAPI() *imports.Imports {
	int64ArrayWrapOnce.Do(func() {
		int64ArrayWrapImport = api.NewDefaultImports()
		int64ArrayWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TInt64ArrayWrap_CreateWithInt", 0), // constructor NewInt64ArrayWrapWithInt
			/* 1 */ imports.NewTable("TInt64ArrayWrap_Int64ArrayData", 0), // function Int64ArrayData
			/* 2 */ imports.NewTable("TInt64ArrayWrap_Size", 0), // function Size
			/* 3 */ imports.NewTable("TInt64ArrayWrap_GetValue", 0), // function GetValue
			/* 4 */ imports.NewTable("TInt64ArrayWrap_SetValue", 0), // procedure SetValue
		}
	})
	return int64ArrayWrapImport
}
