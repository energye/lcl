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

// IFormatEtcArrayWrap Parent: IObject
type IFormatEtcArrayWrap interface {
	IObject
	FormatEtcArrayData() types.PFormatEtcArray    // function
	Size() int32                                  // function
	GetValue(index int32) types.PFormatEtc        // function
	SetValue(index int32, value types.PFormatEtc) // procedure
}

type TFormatEtcArrayWrap struct {
	TObject
}

func (m *TFormatEtcArrayWrap) FormatEtcArrayData() types.PFormatEtcArray {
	if !m.IsValid() {
		return 0
	}
	r := formatEtcArrayWrapAPI().SysCallN(1, m.Instance())
	return types.PFormatEtcArray(r)
}

func (m *TFormatEtcArrayWrap) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := formatEtcArrayWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TFormatEtcArrayWrap) GetValue(index int32) types.PFormatEtc {
	if !m.IsValid() {
		return 0
	}
	r := formatEtcArrayWrapAPI().SysCallN(3, m.Instance(), uintptr(index))
	return types.PFormatEtc(r)
}

func (m *TFormatEtcArrayWrap) SetValue(index int32, value types.PFormatEtc) {
	if !m.IsValid() {
		return
	}
	formatEtcArrayWrapAPI().SysCallN(4, m.Instance(), uintptr(index), uintptr(value))
}

// NewFormatEtcArrayWrap class constructor
func NewFormatEtcArrayWrap(size int32) IFormatEtcArrayWrap {
	r := formatEtcArrayWrapAPI().SysCallN(0, uintptr(size))
	return AsFormatEtcArrayWrap(r)
}

var (
	formatEtcArrayWrapOnce   base.Once
	formatEtcArrayWrapImport *imports.Imports = nil
)

func formatEtcArrayWrapAPI() *imports.Imports {
	formatEtcArrayWrapOnce.Do(func() {
		formatEtcArrayWrapImport = api.NewDefaultImports()
		formatEtcArrayWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFormatEtcArrayWrap_Create", 0), // constructor NewFormatEtcArrayWrap
			/* 1 */ imports.NewTable("TFormatEtcArrayWrap_FormatEtcArrayData", 0), // function FormatEtcArrayData
			/* 2 */ imports.NewTable("TFormatEtcArrayWrap_Size", 0), // function Size
			/* 3 */ imports.NewTable("TFormatEtcArrayWrap_GetValue", 0), // function GetValue
			/* 4 */ imports.NewTable("TFormatEtcArrayWrap_SetValue", 0), // procedure SetValue
		}
	})
	return formatEtcArrayWrapImport
}
