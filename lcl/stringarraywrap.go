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

// IStringArrayWrap Parent: IObject
type IStringArrayWrap interface {
	IObject
	StringArrayData() types.PStringArray // function
	Size() int32                         // function
	GetValue(index int32) string         // function
	SetValue(index int32, value string)  // procedure
}

type TStringArrayWrap struct {
	TObject
}

func (m *TStringArrayWrap) StringArrayData() types.PStringArray {
	if !m.IsValid() {
		return 0
	}
	r := stringArrayWrapAPI().SysCallN(1, m.Instance())
	return types.PStringArray(r)
}

func (m *TStringArrayWrap) Size() int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringArrayWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TStringArrayWrap) GetValue(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := stringArrayWrapAPI().SysCallN(3, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TStringArrayWrap) SetValue(index int32, value string) {
	if !m.IsValid() {
		return
	}
	stringArrayWrapAPI().SysCallN(4, m.Instance(), uintptr(index), api.PasStr(value))
}

// NewStringArrayWrapWithInt class constructor
func NewStringArrayWrapWithInt(size int32) IStringArrayWrap {
	r := stringArrayWrapAPI().SysCallN(0, uintptr(size))
	return AsStringArrayWrap(r)
}

var (
	stringArrayWrapOnce   base.Once
	stringArrayWrapImport *imports.Imports = nil
)

func stringArrayWrapAPI() *imports.Imports {
	stringArrayWrapOnce.Do(func() {
		stringArrayWrapImport = api.NewDefaultImports()
		stringArrayWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStringArrayWrap_CreateWithInt", 0), // constructor NewStringArrayWrapWithInt
			/* 1 */ imports.NewTable("TStringArrayWrap_StringArrayData", 0), // function StringArrayData
			/* 2 */ imports.NewTable("TStringArrayWrap_Size", 0), // function Size
			/* 3 */ imports.NewTable("TStringArrayWrap_GetValue", 0), // function GetValue
			/* 4 */ imports.NewTable("TStringArrayWrap_SetValue", 0), // procedure SetValue
		}
	})
	return stringArrayWrapImport
}
