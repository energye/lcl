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

// IGraphic Parent: IPersistent
type IGraphic interface {
	IPersistent
	LazarusResourceTypeValid(resourceType string) bool                                             // function
	GetResourceType() string                                                                       // function
	Clear()                                                                                        // procedure
	LoadFromFile(filename string)                                                                  // procedure
	LoadFromStreamWithStream(stream IStream)                                                       // procedure
	LoadFromMimeStream(stream IStream, mimeType string)                                            // procedure
	LoadFromLazarusResource(resName string)                                                        // procedure
	LoadFromResourceName(instance types.TLCLHandle, resName string)                                // procedure
	LoadFromResourceID(instance types.TLCLHandle, resID uintptr)                                   // procedure
	LoadFromClipboardFormat(formatID types.TClipboardFormat)                                       // procedure
	LoadFromClipboardFormatID(clipboardType types.TClipboardType, formatID types.TClipboardFormat) // procedure
	SaveToFile(filename string)                                                                    // procedure
	SaveToStream(stream IStream)                                                                   // procedure
	SaveToClipboardFormat(formatID types.TClipboardFormat)                                         // procedure
	SaveToClipboardFormatID(clipboardType types.TClipboardType, formatID types.TClipboardFormat)   // procedure
	GetSupportedSourceMimeTypes(list IStrings)                                                     // procedure
	Empty() bool                                                                                   // property Empty Getter
	Height() int32                                                                                 // property Height Getter
	SetHeight(value int32)                                                                         // property Height Setter
	Modified() bool                                                                                // property Modified Getter
	SetModified(value bool)                                                                        // property Modified Setter
	MimeType() string                                                                              // property MimeType Getter
	Palette() types.HPALETTE                                                                       // property Palette Getter
	SetPalette(value types.HPALETTE)                                                               // property Palette Setter
	PaletteModified() bool                                                                         // property PaletteModified Getter
	SetPaletteModified(value bool)                                                                 // property PaletteModified Setter
	Transparent() bool                                                                             // property Transparent Getter
	SetTransparent(value bool)                                                                     // property Transparent Setter
	Width() int32                                                                                  // property Width Getter
	SetWidth(value int32)                                                                          // property Width Setter
	SetOnChange(fn TNotifyEvent)                                                                   // property event
	SetOnProgress(fn TProgressEvent)                                                               // property event
}

type TGraphic struct {
	TPersistent
}

func (m *TGraphic) LazarusResourceTypeValid(resourceType string) bool {
	if !m.IsValid() {
		return false
	}
	r := graphicAPI().SysCallN(0, m.Instance(), api.PasStr(resourceType))
	return api.GoBool(r)
}

func (m *TGraphic) GetResourceType() string {
	if !m.IsValid() {
		return ""
	}
	r := graphicAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TGraphic) Clear() {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(4, m.Instance())
}

func (m *TGraphic) LoadFromFile(filename string) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(5, m.Instance(), api.PasStr(filename))
}

func (m *TGraphic) LoadFromStreamWithStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TGraphic) LoadFromMimeStream(stream IStream, mimeType string) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(stream), api.PasStr(mimeType))
}

func (m *TGraphic) LoadFromLazarusResource(resName string) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(8, m.Instance(), api.PasStr(resName))
}

func (m *TGraphic) LoadFromResourceName(instance types.TLCLHandle, resName string) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(9, m.Instance(), uintptr(instance), api.PasStr(resName))
}

func (m *TGraphic) LoadFromResourceID(instance types.TLCLHandle, resID uintptr) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(10, m.Instance(), uintptr(instance), uintptr(resID))
}

func (m *TGraphic) LoadFromClipboardFormat(formatID types.TClipboardFormat) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(11, m.Instance(), uintptr(formatID))
}

func (m *TGraphic) LoadFromClipboardFormatID(clipboardType types.TClipboardType, formatID types.TClipboardFormat) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(12, m.Instance(), uintptr(clipboardType), uintptr(formatID))
}

func (m *TGraphic) SaveToFile(filename string) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(13, m.Instance(), api.PasStr(filename))
}

func (m *TGraphic) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TGraphic) SaveToClipboardFormat(formatID types.TClipboardFormat) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(15, m.Instance(), uintptr(formatID))
}

func (m *TGraphic) SaveToClipboardFormatID(clipboardType types.TClipboardType, formatID types.TClipboardFormat) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(16, m.Instance(), uintptr(clipboardType), uintptr(formatID))
}

func (m *TGraphic) GetSupportedSourceMimeTypes(list IStrings) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(17, m.Instance(), base.GetObjectUintptr(list))
}

func (m *TGraphic) Empty() bool {
	if !m.IsValid() {
		return false
	}
	r := graphicAPI().SysCallN(18, m.Instance())
	return api.GoBool(r)
}

func (m *TGraphic) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := graphicAPI().SysCallN(19, 0, m.Instance())
	return int32(r)
}

func (m *TGraphic) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TGraphic) Modified() bool {
	if !m.IsValid() {
		return false
	}
	r := graphicAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGraphic) SetModified(value bool) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TGraphic) MimeType() string {
	if !m.IsValid() {
		return ""
	}
	r := graphicAPI().SysCallN(21, m.Instance())
	return api.GoStr(r)
}

func (m *TGraphic) Palette() types.HPALETTE {
	if !m.IsValid() {
		return 0
	}
	r := graphicAPI().SysCallN(22, 0, m.Instance())
	return types.HPALETTE(r)
}

func (m *TGraphic) SetPalette(value types.HPALETTE) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TGraphic) PaletteModified() bool {
	if !m.IsValid() {
		return false
	}
	r := graphicAPI().SysCallN(23, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGraphic) SetPaletteModified(value bool) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(23, 1, m.Instance(), api.PasBool(value))
}

func (m *TGraphic) Transparent() bool {
	if !m.IsValid() {
		return false
	}
	r := graphicAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TGraphic) SetTransparent(value bool) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TGraphic) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := graphicAPI().SysCallN(25, 0, m.Instance())
	return int32(r)
}

func (m *TGraphic) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	graphicAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TGraphic) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 26, graphicAPI(), api.MakeEventDataPtr(cb))
}

func (m *TGraphic) SetOnProgress(fn TProgressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTProgressEvent(fn)
	base.SetEvent(m, 27, graphicAPI(), api.MakeEventDataPtr(cb))
}

// Graphic  is static instance
var Graphic _GraphicClass

// _GraphicClass is class type defined by TGraphic
type _GraphicClass uintptr

func (_GraphicClass) GetFileExtensions() string {
	r := graphicAPI().SysCallN(2)
	return api.GoStr(r)
}

func (_GraphicClass) IsStreamFormatSupported(stream IStream) bool {
	r := graphicAPI().SysCallN(3, base.GetObjectUintptr(stream))
	return api.GoBool(r)
}

var (
	graphicOnce   base.Once
	graphicImport *imports.Imports = nil
)

func graphicAPI() *imports.Imports {
	graphicOnce.Do(func() {
		graphicImport = api.NewDefaultImports()
		graphicImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TGraphic_LazarusResourceTypeValid", 0), // function LazarusResourceTypeValid
			/* 1 */ imports.NewTable("TGraphic_GetResourceType", 0), // function GetResourceType
			/* 2 */ imports.NewTable("TGraphic_GetFileExtensions", 0), // static function GetFileExtensions
			/* 3 */ imports.NewTable("TGraphic_IsStreamFormatSupported", 0), // static function IsStreamFormatSupported
			/* 4 */ imports.NewTable("TGraphic_Clear", 0), // procedure Clear
			/* 5 */ imports.NewTable("TGraphic_LoadFromFile", 0), // procedure LoadFromFile
			/* 6 */ imports.NewTable("TGraphic_LoadFromStreamWithStream", 0), // procedure LoadFromStreamWithStream
			/* 7 */ imports.NewTable("TGraphic_LoadFromMimeStream", 0), // procedure LoadFromMimeStream
			/* 8 */ imports.NewTable("TGraphic_LoadFromLazarusResource", 0), // procedure LoadFromLazarusResource
			/* 9 */ imports.NewTable("TGraphic_LoadFromResourceName", 0), // procedure LoadFromResourceName
			/* 10 */ imports.NewTable("TGraphic_LoadFromResourceID", 0), // procedure LoadFromResourceID
			/* 11 */ imports.NewTable("TGraphic_LoadFromClipboardFormat", 0), // procedure LoadFromClipboardFormat
			/* 12 */ imports.NewTable("TGraphic_LoadFromClipboardFormatID", 0), // procedure LoadFromClipboardFormatID
			/* 13 */ imports.NewTable("TGraphic_SaveToFile", 0), // procedure SaveToFile
			/* 14 */ imports.NewTable("TGraphic_SaveToStream", 0), // procedure SaveToStream
			/* 15 */ imports.NewTable("TGraphic_SaveToClipboardFormat", 0), // procedure SaveToClipboardFormat
			/* 16 */ imports.NewTable("TGraphic_SaveToClipboardFormatID", 0), // procedure SaveToClipboardFormatID
			/* 17 */ imports.NewTable("TGraphic_GetSupportedSourceMimeTypes", 0), // procedure GetSupportedSourceMimeTypes
			/* 18 */ imports.NewTable("TGraphic_Empty", 0), // property Empty
			/* 19 */ imports.NewTable("TGraphic_Height", 0), // property Height
			/* 20 */ imports.NewTable("TGraphic_Modified", 0), // property Modified
			/* 21 */ imports.NewTable("TGraphic_MimeType", 0), // property MimeType
			/* 22 */ imports.NewTable("TGraphic_Palette", 0), // property Palette
			/* 23 */ imports.NewTable("TGraphic_PaletteModified", 0), // property PaletteModified
			/* 24 */ imports.NewTable("TGraphic_Transparent", 0), // property Transparent
			/* 25 */ imports.NewTable("TGraphic_Width", 0), // property Width
			/* 26 */ imports.NewTable("TGraphic_OnChange", 0), // event OnChange
			/* 27 */ imports.NewTable("TGraphic_OnProgress", 0), // event OnProgress
		}
	})
	return graphicImport
}
