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

// IFPPalette Parent: IObject
type IFPPalette interface {
	IObject
	IndexOf(color TFPColor) int32         // function
	Add(value TFPColor) int32             // function
	Build(img IFPCustomImage)             // procedure
	Copy(palette IFPPalette)              // procedure
	Merge(pal IFPPalette)                 // procedure
	Clear()                               // procedure
	Color(index int32) TFPColor           // property Color Getter
	SetColor(index int32, value TFPColor) // property Color Setter
	Count() int32                         // property Count Getter
	SetCount(value int32)                 // property Count Setter
	Capacity() int32                      // property Capacity Getter
	SetCapacity(value int32)              // property Capacity Setter
}

type TFPPalette struct {
	TObject
}

func (m *TFPPalette) IndexOf(color TFPColor) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPPaletteAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&color)))
	return int32(r)
}

func (m *TFPPalette) Add(value TFPColor) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPPaletteAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&value)))
	return int32(r)
}

func (m *TFPPalette) Build(img IFPCustomImage) {
	if !m.IsValid() {
		return
	}
	fPPaletteAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(img))
}

func (m *TFPPalette) Copy(palette IFPPalette) {
	if !m.IsValid() {
		return
	}
	fPPaletteAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(palette))
}

func (m *TFPPalette) Merge(pal IFPPalette) {
	if !m.IsValid() {
		return
	}
	fPPaletteAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(pal))
}

func (m *TFPPalette) Clear() {
	if !m.IsValid() {
		return
	}
	fPPaletteAPI().SysCallN(6, m.Instance())
}

func (m *TFPPalette) Color(index int32) (result TFPColor) {
	if !m.IsValid() {
		return
	}
	fPPaletteAPI().SysCallN(7, 0, m.Instance(), uintptr(index), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TFPPalette) SetColor(index int32, value TFPColor) {
	if !m.IsValid() {
		return
	}
	fPPaletteAPI().SysCallN(7, 1, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&value)))
}

func (m *TFPPalette) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPPaletteAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TFPPalette) SetCount(value int32) {
	if !m.IsValid() {
		return
	}
	fPPaletteAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TFPPalette) Capacity() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPPaletteAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TFPPalette) SetCapacity(value int32) {
	if !m.IsValid() {
		return
	}
	fPPaletteAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

// NewFPPalette class constructor
func NewFPPalette(count int32) IFPPalette {
	r := fPPaletteAPI().SysCallN(0, uintptr(count))
	return AsFPPalette(r)
}

var (
	fPPaletteOnce   base.Once
	fPPaletteImport *imports.Imports = nil
)

func fPPaletteAPI() *imports.Imports {
	fPPaletteOnce.Do(func() {
		fPPaletteImport = api.NewDefaultImports()
		fPPaletteImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPPalette_Create", 0), // constructor NewFPPalette
			/* 1 */ imports.NewTable("TFPPalette_IndexOf", 0), // function IndexOf
			/* 2 */ imports.NewTable("TFPPalette_Add", 0), // function Add
			/* 3 */ imports.NewTable("TFPPalette_Build", 0), // procedure Build
			/* 4 */ imports.NewTable("TFPPalette_Copy", 0), // procedure Copy
			/* 5 */ imports.NewTable("TFPPalette_Merge", 0), // procedure Merge
			/* 6 */ imports.NewTable("TFPPalette_Clear", 0), // procedure Clear
			/* 7 */ imports.NewTable("TFPPalette_Color", 0), // property Color
			/* 8 */ imports.NewTable("TFPPalette_Count", 0), // property Count
			/* 9 */ imports.NewTable("TFPPalette_Capacity", 0), // property Capacity
		}
	})
	return fPPaletteImport
}
