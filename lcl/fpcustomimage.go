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

// IFPCustomImage Is Abstract Class Parent: IPersistent
type IFPCustomImage interface {
	IPersistent
	Height() int32                                               // property
	SetHeight(AValue int32)                                      // property
	Width() int32                                                // property
	SetWidth(AValue int32)                                       // property
	Colors(x, y int32) (resultFPColor TFPColor)                  // property
	SetColors(x, y int32, AValue *TFPColor)                      // property
	UsePalette() bool                                            // property
	SetUsePalette(AValue bool)                                   // property
	Palette() IFPPalette                                         // property
	Pixels(x, y int32) int32                                     // property
	SetPixels(x, y int32, AValue int32)                          // property
	Extra(key string) string                                     // property
	SetExtra(key string, AValue string)                          // property
	ExtraValue(index int32) string                               // property
	SetExtraValue(index int32, AValue string)                    // property
	ExtraKey(index int32) string                                 // property
	SetExtraKey(index int32, AValue string)                      // property
	LoadFromFile(filename string) bool                           // function
	SaveToFile(filename string) bool                             // function
	ExtraCount() int32                                           // function
	LoadFromStream(Str IStream, Handler IFPCustomImageReader)    // procedure
	LoadFromStream1(Str IStream)                                 // procedure
	LoadFromFile1(filename string, Handler IFPCustomImageReader) // procedure
	SaveToStream(Str IStream, Handler IFPCustomImageWriter)      // procedure
	SaveToFile1(filename string, Handler IFPCustomImageWriter)   // procedure
	SetSize(AWidth, AHeight int32)                               // procedure
	RemoveExtra(key string)                                      // procedure
	SetOnProgress(fn TFPImgProgressEvent)                        // property event
}

// TFPCustomImage Is Abstract Class Parent: TPersistent
type TFPCustomImage struct {
	TPersistent
	progressPtr uintptr
}

func (m *TFPCustomImage) Height() int32 {
	r1 := fPCustomImageImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomImage) SetHeight(AValue int32) {
	fPCustomImageImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomImage) Width() int32 {
	r1 := fPCustomImageImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomImage) SetWidth(AValue int32) {
	fPCustomImageImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomImage) Colors(x, y int32) (resultFPColor TFPColor) {
	r1 := fPCustomImageImportAPI().SysCallN(1, 0, m.Instance(), uintptr(x), uintptr(y))
	return *(*TFPColor)(getPointer(r1))
}

func (m *TFPCustomImage) SetColors(x, y int32, AValue *TFPColor) {
	fPCustomImageImportAPI().SysCallN(1, 1, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(AValue)))
}

func (m *TFPCustomImage) UsePalette() bool {
	r1 := fPCustomImageImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomImage) SetUsePalette(AValue bool) {
	fPCustomImageImportAPI().SysCallN(19, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomImage) Palette() IFPPalette {
	r1 := fPCustomImageImportAPI().SysCallN(11, m.Instance())
	return AsFPPalette(r1)
}

func (m *TFPCustomImage) Pixels(x, y int32) int32 {
	r1 := fPCustomImageImportAPI().SysCallN(12, 0, m.Instance(), uintptr(x), uintptr(y))
	return int32(r1)
}

func (m *TFPCustomImage) SetPixels(x, y int32, AValue int32) {
	fPCustomImageImportAPI().SysCallN(12, 1, m.Instance(), uintptr(x), uintptr(y), uintptr(AValue))
}

func (m *TFPCustomImage) Extra(key string) string {
	r1 := fPCustomImageImportAPI().SysCallN(2, 0, m.Instance(), PascalStr(key))
	return GoStr(r1)
}

func (m *TFPCustomImage) SetExtra(key string, AValue string) {
	fPCustomImageImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(key), PascalStr(AValue))
}

func (m *TFPCustomImage) ExtraValue(index int32) string {
	r1 := fPCustomImageImportAPI().SysCallN(5, 0, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TFPCustomImage) SetExtraValue(index int32, AValue string) {
	fPCustomImageImportAPI().SysCallN(5, 1, m.Instance(), uintptr(index), PascalStr(AValue))
}

func (m *TFPCustomImage) ExtraKey(index int32) string {
	r1 := fPCustomImageImportAPI().SysCallN(4, 0, m.Instance(), uintptr(index))
	return GoStr(r1)
}

func (m *TFPCustomImage) SetExtraKey(index int32, AValue string) {
	fPCustomImageImportAPI().SysCallN(4, 1, m.Instance(), uintptr(index), PascalStr(AValue))
}

func (m *TFPCustomImage) LoadFromFile(filename string) bool {
	r1 := fPCustomImageImportAPI().SysCallN(7, m.Instance(), PascalStr(filename))
	return GoBool(r1)
}

func (m *TFPCustomImage) SaveToFile(filename string) bool {
	r1 := fPCustomImageImportAPI().SysCallN(14, m.Instance(), PascalStr(filename))
	return GoBool(r1)
}

func (m *TFPCustomImage) ExtraCount() int32 {
	r1 := fPCustomImageImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

func FPCustomImageClass() TClass {
	ret := fPCustomImageImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TFPCustomImage) LoadFromStream(Str IStream, Handler IFPCustomImageReader) {
	fPCustomImageImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(Str), GetObjectUintptr(Handler))
}

func (m *TFPCustomImage) LoadFromStream1(Str IStream) {
	fPCustomImageImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(Str))
}

func (m *TFPCustomImage) LoadFromFile1(filename string, Handler IFPCustomImageReader) {
	fPCustomImageImportAPI().SysCallN(8, m.Instance(), PascalStr(filename), GetObjectUintptr(Handler))
}

func (m *TFPCustomImage) SaveToStream(Str IStream, Handler IFPCustomImageWriter) {
	fPCustomImageImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(Str), GetObjectUintptr(Handler))
}

func (m *TFPCustomImage) SaveToFile1(filename string, Handler IFPCustomImageWriter) {
	fPCustomImageImportAPI().SysCallN(15, m.Instance(), PascalStr(filename), GetObjectUintptr(Handler))
}

func (m *TFPCustomImage) SetSize(AWidth, AHeight int32) {
	fPCustomImageImportAPI().SysCallN(18, m.Instance(), uintptr(AWidth), uintptr(AHeight))
}

func (m *TFPCustomImage) RemoveExtra(key string) {
	fPCustomImageImportAPI().SysCallN(13, m.Instance(), PascalStr(key))
}

func (m *TFPCustomImage) SetOnProgress(fn TFPImgProgressEvent) {
	if m.progressPtr != 0 {
		RemoveEventElement(m.progressPtr)
	}
	m.progressPtr = MakeEventDataPtr(fn)
	fPCustomImageImportAPI().SysCallN(17, m.Instance(), m.progressPtr)
}

var (
	fPCustomImageImport       *imports.Imports = nil
	fPCustomImageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCustomImage_Class", 0),
		/*1*/ imports.NewTable("FPCustomImage_Colors", 0),
		/*2*/ imports.NewTable("FPCustomImage_Extra", 0),
		/*3*/ imports.NewTable("FPCustomImage_ExtraCount", 0),
		/*4*/ imports.NewTable("FPCustomImage_ExtraKey", 0),
		/*5*/ imports.NewTable("FPCustomImage_ExtraValue", 0),
		/*6*/ imports.NewTable("FPCustomImage_Height", 0),
		/*7*/ imports.NewTable("FPCustomImage_LoadFromFile", 0),
		/*8*/ imports.NewTable("FPCustomImage_LoadFromFile1", 0),
		/*9*/ imports.NewTable("FPCustomImage_LoadFromStream", 0),
		/*10*/ imports.NewTable("FPCustomImage_LoadFromStream1", 0),
		/*11*/ imports.NewTable("FPCustomImage_Palette", 0),
		/*12*/ imports.NewTable("FPCustomImage_Pixels", 0),
		/*13*/ imports.NewTable("FPCustomImage_RemoveExtra", 0),
		/*14*/ imports.NewTable("FPCustomImage_SaveToFile", 0),
		/*15*/ imports.NewTable("FPCustomImage_SaveToFile1", 0),
		/*16*/ imports.NewTable("FPCustomImage_SaveToStream", 0),
		/*17*/ imports.NewTable("FPCustomImage_SetOnProgress", 0),
		/*18*/ imports.NewTable("FPCustomImage_SetSize", 0),
		/*19*/ imports.NewTable("FPCustomImage_UsePalette", 0),
		/*20*/ imports.NewTable("FPCustomImage_Width", 0),
	}
)

func fPCustomImageImportAPI() *imports.Imports {
	if fPCustomImageImport == nil {
		fPCustomImageImport = NewDefaultImports()
		fPCustomImageImport.SetImportTable(fPCustomImageImportTables)
		fPCustomImageImportTables = nil
	}
	return fPCustomImageImport
}
