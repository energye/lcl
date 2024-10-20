//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IFPPalette Parent: IObject
type IFPPalette interface {
	IObject
	Color(Index int32) (resultFPColor TFPColor) // property
	SetColor(Index int32, AValue *TFPColor)     // property
	Count() int32                               // property
	SetCount(AValue int32)                      // property
	Capacity() int32                            // property
	SetCapacity(AValue int32)                   // property
	IndexOf(AColor *TFPColor) int32             // function
	Add(Value *TFPColor) int32                  // function
	Build(Img IFPCustomImage)                   // procedure
	Copy(APalette IFPPalette)                   // procedure
	Merge(pal IFPPalette)                       // procedure
	Clear()                                     // procedure
}

// TFPPalette Parent: TObject
type TFPPalette struct {
	TObject
}

func NewFPPalette(ACount int32) IFPPalette {
	r1 := fPPaletteImportAPI().SysCallN(8, uintptr(ACount))
	return AsFPPalette(r1)
}

func (m *TFPPalette) Color(Index int32) (resultFPColor TFPColor) {
	r1 := fPPaletteImportAPI().SysCallN(5, 0, m.Instance(), uintptr(Index))
	return *(*TFPColor)(getPointer(r1))
}

func (m *TFPPalette) SetColor(Index int32, AValue *TFPColor) {
	fPPaletteImportAPI().SysCallN(5, 1, m.Instance(), uintptr(Index), uintptr(unsafePointer(AValue)))
}

func (m *TFPPalette) Count() int32 {
	r1 := fPPaletteImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPPalette) SetCount(AValue int32) {
	fPPaletteImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPPalette) Capacity() int32 {
	r1 := fPPaletteImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPPalette) SetCapacity(AValue int32) {
	fPPaletteImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPPalette) IndexOf(AColor *TFPColor) int32 {
	r1 := fPPaletteImportAPI().SysCallN(9, m.Instance(), uintptr(unsafePointer(AColor)))
	return int32(r1)
}

func (m *TFPPalette) Add(Value *TFPColor) int32 {
	r1 := fPPaletteImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(Value)))
	return int32(r1)
}

func FPPaletteClass() TClass {
	ret := fPPaletteImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TFPPalette) Build(Img IFPCustomImage) {
	fPPaletteImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(Img))
}

func (m *TFPPalette) Copy(APalette IFPPalette) {
	fPPaletteImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(APalette))
}

func (m *TFPPalette) Merge(pal IFPPalette) {
	fPPaletteImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(pal))
}

func (m *TFPPalette) Clear() {
	fPPaletteImportAPI().SysCallN(4, m.Instance())
}

var (
	fPPaletteImport       *imports.Imports = nil
	fPPaletteImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPPalette_Add", 0),
		/*1*/ imports.NewTable("FPPalette_Build", 0),
		/*2*/ imports.NewTable("FPPalette_Capacity", 0),
		/*3*/ imports.NewTable("FPPalette_Class", 0),
		/*4*/ imports.NewTable("FPPalette_Clear", 0),
		/*5*/ imports.NewTable("FPPalette_Color", 0),
		/*6*/ imports.NewTable("FPPalette_Copy", 0),
		/*7*/ imports.NewTable("FPPalette_Count", 0),
		/*8*/ imports.NewTable("FPPalette_Create", 0),
		/*9*/ imports.NewTable("FPPalette_IndexOf", 0),
		/*10*/ imports.NewTable("FPPalette_Merge", 0),
	}
)

func fPPaletteImportAPI() *imports.Imports {
	if fPPaletteImport == nil {
		fPPaletteImport = NewDefaultImports()
		fPPaletteImport.SetImportTable(fPPaletteImportTables)
		fPPaletteImportTables = nil
	}
	return fPPaletteImport
}
