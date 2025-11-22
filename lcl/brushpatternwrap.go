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

// IBrushPatternWrap Parent: IObject
type IBrushPatternWrap interface {
	IObject
	BrushPatternData() types.PBrushPattern // function
	Size() int32                           // function
	GetValue(index int32) uint32           // function
	SetValue(index int32, value uint32)    // procedure
}

type TBrushPatternWrap struct {
	TObject
}

func (m *TBrushPatternWrap) BrushPatternData() types.PBrushPattern {
	if !m.IsValid() {
		return 0
	}
	r := brushPatternWrapAPI().SysCallN(1, m.Instance())
	return types.PBrushPattern(r)
}

func (m *TBrushPatternWrap) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := brushPatternWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TBrushPatternWrap) GetValue(index int32) uint32 {
	if !m.IsValid() {
		return 0
	}
	r := brushPatternWrapAPI().SysCallN(3, m.Instance(), uintptr(index))
	return uint32(r)
}

func (m *TBrushPatternWrap) SetValue(index int32, value uint32) {
	if !m.IsValid() {
		return
	}
	brushPatternWrapAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(value))
}

// NewBrushPatternWrapToBrushPatternWrap class constructor
func NewBrushPatternWrapToBrushPatternWrap() IBrushPatternWrap {
	r := brushPatternWrapAPI().SysCallN(0)
	return AsBrushPatternWrap(r)
}

var (
	brushPatternWrapOnce   base.Once
	brushPatternWrapImport *imports.Imports = nil
)

func brushPatternWrapAPI() *imports.Imports {
	brushPatternWrapOnce.Do(func() {
		brushPatternWrapImport = api.NewDefaultImports()
		brushPatternWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TBrushPatternWrap_CreateToBrushPatternWrap", 0), // constructor NewBrushPatternWrapToBrushPatternWrap
			/* 1 */ imports.NewTable("TBrushPatternWrap_BrushPatternData", 0), // function BrushPatternData
			/* 2 */ imports.NewTable("TBrushPatternWrap_Size", 0), // function Size
			/* 3 */ imports.NewTable("TBrushPatternWrap_GetValue", 0), // function GetValue
			/* 4 */ imports.NewTable("TBrushPatternWrap_SetValue", 0), // procedure SetValue
		}
	})
	return brushPatternWrapImport
}
