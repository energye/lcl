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

// IPointArrayWrap Parent: IObject
type IPointArrayWrap interface {
	IObject
	PointArrayData() types.PPointArray        // function
	Size() int32                              // function
	GetValue(index int32) types.TPoint        // function
	SetValue(index int32, value types.TPoint) // procedure
}

type TPointArrayWrap struct {
	TObject
}

func (m *TPointArrayWrap) PointArrayData() types.PPointArray {
	if !m.IsValid() {
		return 0
	}
	r := pointArrayWrapAPI().SysCallN(1, m.Instance())
	return types.PPointArray(r)
}

func (m *TPointArrayWrap) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pointArrayWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TPointArrayWrap) GetValue(index int32) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	pointArrayWrapAPI().SysCallN(3, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TPointArrayWrap) SetValue(index int32, value types.TPoint) {
	if !m.IsValid() {
		return
	}
	pointArrayWrapAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&value)))
}

// NewPointArrayWrapWithInt class constructor
func NewPointArrayWrapWithInt(size int32) IPointArrayWrap {
	r := pointArrayWrapAPI().SysCallN(0, uintptr(size))
	return AsPointArrayWrap(r)
}

var (
	pointArrayWrapOnce   base.Once
	pointArrayWrapImport *imports.Imports = nil
)

func pointArrayWrapAPI() *imports.Imports {
	pointArrayWrapOnce.Do(func() {
		pointArrayWrapImport = api.NewDefaultImports()
		pointArrayWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPointArrayWrap_CreateWithInt", 0), // constructor NewPointArrayWrapWithInt
			/* 1 */ imports.NewTable("TPointArrayWrap_PointArrayData", 0), // function PointArrayData
			/* 2 */ imports.NewTable("TPointArrayWrap_Size", 0), // function Size
			/* 3 */ imports.NewTable("TPointArrayWrap_GetValue", 0), // function GetValue
			/* 4 */ imports.NewTable("TPointArrayWrap_SetValue", 0), // procedure SetValue
		}
	})
	return pointArrayWrapImport
}
