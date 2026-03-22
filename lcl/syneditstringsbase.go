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

// ISynEditStringsBase Parent: IStrings
type ISynEditStringsBase interface {
	IStrings
	GetPCharWithInt(lineIndex int32) uintptr                  // function
	GetPCharWithIntX2(lineIndex int32, outLen *int32) uintptr // function
	SendHighlightChanged(index int32, count int32)            // procedure
	Ranges(index uintptr) ISynManagedStorageMem               // property Ranges Getter
	SetRanges(index uintptr, value ISynManagedStorageMem)     // property Ranges Setter
}

type TSynEditStringsBase struct {
	TStrings
}

func (m *TSynEditStringsBase) GetPCharWithInt(lineIndex int32) uintptr {
	if !m.IsValid() {
		return 0
	}
	r := synEditStringsBaseAPI().SysCallN(0, m.Instance(), uintptr(lineIndex))
	return uintptr(r)
}

func (m *TSynEditStringsBase) GetPCharWithIntX2(lineIndex int32, outLen *int32) uintptr {
	if !m.IsValid() {
		return 0
	}
	var lenPtr uintptr
	r := synEditStringsBaseAPI().SysCallN(1, m.Instance(), uintptr(lineIndex), uintptr(base.UnsafePointer(&lenPtr)))
	*outLen = int32(lenPtr)
	return uintptr(r)
}

func (m *TSynEditStringsBase) SendHighlightChanged(index int32, count int32) {
	if !m.IsValid() {
		return
	}
	synEditStringsBaseAPI().SysCallN(2, m.Instance(), uintptr(index), uintptr(count))
}

func (m *TSynEditStringsBase) Ranges(index uintptr) ISynManagedStorageMem {
	if !m.IsValid() {
		return nil
	}
	r := synEditStringsBaseAPI().SysCallN(3, 0, m.Instance(), uintptr(index))
	return AsSynManagedStorageMem(r)
}

func (m *TSynEditStringsBase) SetRanges(index uintptr, value ISynManagedStorageMem) {
	if !m.IsValid() {
		return
	}
	synEditStringsBaseAPI().SysCallN(3, 1, m.Instance(), uintptr(index), base.GetObjectUintptr(value))
}

var (
	synEditStringsBaseOnce   base.Once
	synEditStringsBaseImport *imports.Imports = nil
)

func synEditStringsBaseAPI() *imports.Imports {
	synEditStringsBaseOnce.Do(func() {
		synEditStringsBaseImport = api.NewDefaultImports()
		synEditStringsBaseImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynEditStringsBase_GetPCharWithInt", 0), // function GetPCharWithInt
			/* 1 */ imports.NewTable("TSynEditStringsBase_GetPCharWithIntX2", 0), // function GetPCharWithIntX2
			/* 2 */ imports.NewTable("TSynEditStringsBase_SendHighlightChanged", 0), // procedure SendHighlightChanged
			/* 3 */ imports.NewTable("TSynEditStringsBase_Ranges", 0), // property Ranges
		}
	})
	return synEditStringsBaseImport
}
