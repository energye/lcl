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

// IEtoBuffer Parent: IObject
type IEtoBuffer interface {
	IObject
	Eto() types.PInteger    // function
	Len() int32             // function
	Clear()                 // procedure
	SetMinLength(len int32) // procedure
}

type TEtoBuffer struct {
	TObject
}

func (m *TEtoBuffer) Eto() types.PInteger {
	if !m.IsValid() {
		return 0
	}
	r := etoBufferAPI().SysCallN(0, m.Instance())
	return types.PInteger(r)
}

func (m *TEtoBuffer) Len() int32 {
	if !m.IsValid() {
		return 0
	}
	r := etoBufferAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TEtoBuffer) Clear() {
	if !m.IsValid() {
		return
	}
	etoBufferAPI().SysCallN(2, m.Instance())
}

func (m *TEtoBuffer) SetMinLength(len int32) {
	if !m.IsValid() {
		return
	}
	etoBufferAPI().SysCallN(3, m.Instance(), uintptr(len))
}

var (
	etoBufferOnce   base.Once
	etoBufferImport *imports.Imports = nil
)

func etoBufferAPI() *imports.Imports {
	etoBufferOnce.Do(func() {
		etoBufferImport = api.NewDefaultImports()
		etoBufferImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEtoBuffer_Eto", 0), // function Eto
			/* 1 */ imports.NewTable("TEtoBuffer_Len", 0), // function Len
			/* 2 */ imports.NewTable("TEtoBuffer_Clear", 0), // procedure Clear
			/* 3 */ imports.NewTable("TEtoBuffer_SetMinLength", 0), // procedure SetMinLength
		}
	})
	return etoBufferImport
}
