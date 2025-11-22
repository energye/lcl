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

// IStringWrap Parent: IObject
type IStringWrap interface {
	IObject
	Value() string         // property Value Getter
	SetValue(value string) // property Value Setter
	Len() int32            // property Len Getter
}

type TStringWrap struct {
	TObject
}

func (m *TStringWrap) Value() string {
	if !m.IsValid() {
		return ""
	}
	r := stringWrapAPI().SysCallN(1, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TStringWrap) SetValue(value string) {
	if !m.IsValid() {
		return
	}
	stringWrapAPI().SysCallN(1, 1, m.Instance(), api.PasStr(value))
}

func (m *TStringWrap) Len() int32 {
	if !m.IsValid() {
		return 0
	}
	r := stringWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

// NewStringWrap class constructor
func NewStringWrap() IStringWrap {
	r := stringWrapAPI().SysCallN(0)
	return AsStringWrap(r)
}

var (
	stringWrapOnce   base.Once
	stringWrapImport *imports.Imports = nil
)

func stringWrapAPI() *imports.Imports {
	stringWrapOnce.Do(func() {
		stringWrapImport = api.NewDefaultImports()
		stringWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TStringWrap_Create", 0), // constructor NewStringWrap
			/* 1 */ imports.NewTable("TStringWrap_Value", 0), // property Value
			/* 2 */ imports.NewTable("TStringWrap_Len", 0), // property Len
		}
	})
	return stringWrapImport
}
