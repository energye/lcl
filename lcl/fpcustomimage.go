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

// IFPCustomImage Parent: IPersistent
type IFPCustomImage interface {
	IPersistent
	LoadFromFileWithString(filename string) bool // function
	SaveToFileWithString(filename string) bool   // function
	ExtraCount() int32                           // function
	// LoadFromStreamWithStreamFPCustomImageReader
	//  Saving and loading
	LoadFromStreamWithStreamFPCustomImageReader(str IStream, handler IFPCustomImageReader)   // procedure
	LoadFromStreamWithStream(str IStream)                                                    // procedure
	LoadFromFileWithStringFPCustomImageReader(filename string, handler IFPCustomImageReader) // procedure
	SaveToStream(str IStream, handler IFPCustomImageWriter)                                  // procedure
	SaveToFileWithStringFPCustomImageWriter(filename string, handler IFPCustomImageWriter)   // procedure
	// SetSize
	//  Size and data
	SetSize(width int32, height int32)          // procedure
	RemoveExtra(key string)                     // procedure
	Height() int32                              // property Height Getter
	SetHeight(value int32)                      // property Height Setter
	Width() int32                               // property Width Getter
	SetWidth(value int32)                       // property Width Setter
	Colors(X int32, Y int32) TFPColor           // property Colors Getter
	SetColors(X int32, Y int32, value TFPColor) // property Colors Setter
	// UsePalette
	//  Use of palette for colors
	UsePalette() bool                        // property UsePalette Getter
	SetUsePalette(value bool)                // property UsePalette Setter
	Palette() IFPPalette                     // property Palette Getter
	Pixels(X int32, Y int32) int32           // property Pixels Getter
	SetPixels(X int32, Y int32, value int32) // property Pixels Setter
	// Extra
	//  Info unrelated with the image representation
	Extra(key string) string                 // property Extra Getter
	SetExtra(key string, value string)       // property Extra Setter
	ExtraValue(index int32) string           // property ExtraValue Getter
	SetExtraValue(index int32, value string) // property ExtraValue Setter
	ExtraKey(index int32) string             // property ExtraKey Getter
	SetExtraKey(index int32, value string)   // property ExtraKey Setter
	SetOnProgress(fn TFPImgProgressEvent)    // property event
}

type TFPCustomImage struct {
	TPersistent
}

func (m *TFPCustomImage) LoadFromFileWithString(filename string) bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomImageAPI().SysCallN(0, m.Instance(), api.PasStr(filename))
	return api.GoBool(r)
}

func (m *TFPCustomImage) SaveToFileWithString(filename string) bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomImageAPI().SysCallN(1, m.Instance(), api.PasStr(filename))
	return api.GoBool(r)
}

func (m *TFPCustomImage) ExtraCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomImageAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TFPCustomImage) LoadFromStreamWithStreamFPCustomImageReader(str IStream, handler IFPCustomImageReader) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(str), base.GetObjectUintptr(handler))
}

func (m *TFPCustomImage) LoadFromStreamWithStream(str IStream) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(str))
}

func (m *TFPCustomImage) LoadFromFileWithStringFPCustomImageReader(filename string, handler IFPCustomImageReader) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(5, m.Instance(), api.PasStr(filename), base.GetObjectUintptr(handler))
}

func (m *TFPCustomImage) SaveToStream(str IStream, handler IFPCustomImageWriter) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(str), base.GetObjectUintptr(handler))
}

func (m *TFPCustomImage) SaveToFileWithStringFPCustomImageWriter(filename string, handler IFPCustomImageWriter) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(7, m.Instance(), api.PasStr(filename), base.GetObjectUintptr(handler))
}

func (m *TFPCustomImage) SetSize(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(8, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TFPCustomImage) RemoveExtra(key string) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(9, m.Instance(), api.PasStr(key))
}

func (m *TFPCustomImage) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomImageAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TFPCustomImage) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomImage) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomImageAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TFPCustomImage) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomImage) Colors(X int32, Y int32) (result TFPColor) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(12, 0, m.Instance(), uintptr(X), uintptr(Y), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TFPCustomImage) SetColors(X int32, Y int32, value TFPColor) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(12, 1, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(&value)))
}

func (m *TFPCustomImage) UsePalette() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomImageAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCustomImage) SetUsePalette(value bool) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TFPCustomImage) Palette() IFPPalette {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomImageAPI().SysCallN(14, m.Instance())
	return AsFPPalette(r)
}

func (m *TFPCustomImage) Pixels(X int32, Y int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomImageAPI().SysCallN(15, 0, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r)
}

func (m *TFPCustomImage) SetPixels(X int32, Y int32, value int32) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(15, 1, m.Instance(), uintptr(X), uintptr(Y), uintptr(value))
}

func (m *TFPCustomImage) Extra(key string) string {
	if !m.IsValid() {
		return ""
	}
	r := fPCustomImageAPI().SysCallN(16, 0, m.Instance(), api.PasStr(key))
	return api.GoStr(r)
}

func (m *TFPCustomImage) SetExtra(key string, value string) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(16, 1, m.Instance(), api.PasStr(key), api.PasStr(value))
}

func (m *TFPCustomImage) ExtraValue(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := fPCustomImageAPI().SysCallN(17, 0, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TFPCustomImage) SetExtraValue(index int32, value string) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(17, 1, m.Instance(), uintptr(index), api.PasStr(value))
}

func (m *TFPCustomImage) ExtraKey(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := fPCustomImageAPI().SysCallN(18, 0, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TFPCustomImage) SetExtraKey(index int32, value string) {
	if !m.IsValid() {
		return
	}
	fPCustomImageAPI().SysCallN(18, 1, m.Instance(), uintptr(index), api.PasStr(value))
}

func (m *TFPCustomImage) SetOnProgress(fn TFPImgProgressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTFPImgProgressEvent(fn)
	base.SetEvent(m, 19, fPCustomImageAPI(), api.MakeEventDataPtr(cb))
}

var (
	fPCustomImageOnce   base.Once
	fPCustomImageImport *imports.Imports = nil
)

func fPCustomImageAPI() *imports.Imports {
	fPCustomImageOnce.Do(func() {
		fPCustomImageImport = api.NewDefaultImports()
		fPCustomImageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomImage_LoadFromFileWithString", 0), // function LoadFromFileWithString
			/* 1 */ imports.NewTable("TFPCustomImage_SaveToFileWithString", 0), // function SaveToFileWithString
			/* 2 */ imports.NewTable("TFPCustomImage_ExtraCount", 0), // function ExtraCount
			/* 3 */ imports.NewTable("TFPCustomImage_LoadFromStreamWithStreamFPCustomImageReader", 0), // procedure LoadFromStreamWithStreamFPCustomImageReader
			/* 4 */ imports.NewTable("TFPCustomImage_LoadFromStreamWithStream", 0), // procedure LoadFromStreamWithStream
			/* 5 */ imports.NewTable("TFPCustomImage_LoadFromFileWithStringFPCustomImageReader", 0), // procedure LoadFromFileWithStringFPCustomImageReader
			/* 6 */ imports.NewTable("TFPCustomImage_SaveToStream", 0), // procedure SaveToStream
			/* 7 */ imports.NewTable("TFPCustomImage_SaveToFileWithStringFPCustomImageWriter", 0), // procedure SaveToFileWithStringFPCustomImageWriter
			/* 8 */ imports.NewTable("TFPCustomImage_SetSize", 0), // procedure SetSize
			/* 9 */ imports.NewTable("TFPCustomImage_RemoveExtra", 0), // procedure RemoveExtra
			/* 10 */ imports.NewTable("TFPCustomImage_Height", 0), // property Height
			/* 11 */ imports.NewTable("TFPCustomImage_Width", 0), // property Width
			/* 12 */ imports.NewTable("TFPCustomImage_Colors", 0), // property Colors
			/* 13 */ imports.NewTable("TFPCustomImage_UsePalette", 0), // property UsePalette
			/* 14 */ imports.NewTable("TFPCustomImage_Palette", 0), // property Palette
			/* 15 */ imports.NewTable("TFPCustomImage_Pixels", 0), // property Pixels
			/* 16 */ imports.NewTable("TFPCustomImage_Extra", 0), // property Extra
			/* 17 */ imports.NewTable("TFPCustomImage_ExtraValue", 0), // property ExtraValue
			/* 18 */ imports.NewTable("TFPCustomImage_ExtraKey", 0), // property ExtraKey
			/* 19 */ imports.NewTable("TFPCustomImage_OnProgress", 0), // event OnProgress
		}
	})
	return fPCustomImageImport
}
