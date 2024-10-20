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

// IGraphic Is Abstract Class Parent: IPersistent
type IGraphic interface {
	IPersistent
	LoadFromBytes(data []byte)
	Empty() bool                                                                       // property
	Height() int32                                                                     // property
	SetHeight(AValue int32)                                                            // property
	Modified() bool                                                                    // property
	SetModified(AValue bool)                                                           // property
	MimeType() string                                                                  // property
	Palette() HPALETTE                                                                 // property
	SetPalette(AValue HPALETTE)                                                        // property
	PaletteModified() bool                                                             // property
	SetPaletteModified(AValue bool)                                                    // property
	Transparent() bool                                                                 // property
	SetTransparent(AValue bool)                                                        // property
	Width() int32                                                                      // property
	SetWidth(AValue int32)                                                             // property
	LazarusResourceTypeValid(AResourceType string) bool                                // function
	GetResourceType() TResourceType                                                    // function
	Clear()                                                                            // procedure
	LoadFromFile(Filename string)                                                      // procedure
	LoadFromStream(Stream IStream)                                                     // procedure Is Abstract
	LoadFromMimeStream(AStream IStream, AMimeType string)                              // procedure
	LoadFromLazarusResource(ResName string)                                            // procedure
	LoadFromResourceName(Instance THandle, ResName string)                             // procedure
	LoadFromResourceID(Instance THandle, ResID uint32)                                 // procedure
	LoadFromClipboardFormat(FormatID TClipboardFormat)                                 // procedure
	LoadFromClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) // procedure
	SaveToFile(Filename string)                                                        // procedure
	SaveToStream(Stream IStream)                                                       // procedure Is Abstract
	SaveToClipboardFormat(FormatID TClipboardFormat)                                   // procedure
	SaveToClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat)   // procedure
	GetSupportedSourceMimeTypes(List IStrings)                                         // procedure
	SetOnChange(fn TNotifyEvent)                                                       // property event
	SetOnProgress(fn TProgressEvent)                                                   // property event
}

// TGraphic Is Abstract Class Parent: TPersistent
type TGraphic struct {
	TPersistent
	changePtr   uintptr
	progressPtr uintptr
}

func (m *TGraphic) Empty() bool {
	r1 := graphicImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TGraphic) Height() int32 {
	r1 := graphicImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGraphic) SetHeight(AValue int32) {
	graphicImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TGraphic) Modified() bool {
	r1 := graphicImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGraphic) SetModified(AValue bool) {
	graphicImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGraphic) MimeType() string {
	r1 := graphicImportAPI().SysCallN(15, m.Instance())
	return GoStr(r1)
}

func (m *TGraphic) Palette() HPALETTE {
	r1 := graphicImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return HPALETTE(r1)
}

func (m *TGraphic) SetPalette(AValue HPALETTE) {
	graphicImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TGraphic) PaletteModified() bool {
	r1 := graphicImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGraphic) SetPaletteModified(AValue bool) {
	graphicImportAPI().SysCallN(18, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGraphic) Transparent() bool {
	r1 := graphicImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGraphic) SetTransparent(AValue bool) {
	graphicImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGraphic) Width() int32 {
	r1 := graphicImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TGraphic) SetWidth(AValue int32) {
	graphicImportAPI().SysCallN(26, 1, m.Instance(), uintptr(AValue))
}

func (m *TGraphic) LazarusResourceTypeValid(AResourceType string) bool {
	r1 := graphicImportAPI().SysCallN(6, m.Instance(), PascalStr(AResourceType))
	return GoBool(r1)
}

func (m *TGraphic) GetResourceType() TResourceType {
	r1 := graphicImportAPI().SysCallN(3, m.Instance())
	return TResourceType(r1)
}

func GraphicClass() TClass {
	ret := graphicImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TGraphic) Clear() {
	graphicImportAPI().SysCallN(1, m.Instance())
}

func (m *TGraphic) LoadFromFile(Filename string) {
	graphicImportAPI().SysCallN(9, m.Instance(), PascalStr(Filename))
}

func (m *TGraphic) LoadFromStream(Stream IStream) {
	graphicImportAPI().SysCallN(14, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TGraphic) LoadFromMimeStream(AStream IStream, AMimeType string) {
	graphicImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(AStream), PascalStr(AMimeType))
}

func (m *TGraphic) LoadFromLazarusResource(ResName string) {
	graphicImportAPI().SysCallN(10, m.Instance(), PascalStr(ResName))
}

func (m *TGraphic) LoadFromResourceName(Instance THandle, ResName string) {
	graphicImportAPI().SysCallN(13, m.Instance(), uintptr(Instance), PascalStr(ResName))
}

func (m *TGraphic) LoadFromResourceID(Instance THandle, ResID uint32) {
	graphicImportAPI().SysCallN(12, m.Instance(), uintptr(Instance), uintptr(ResID))
}

func (m *TGraphic) LoadFromClipboardFormat(FormatID TClipboardFormat) {
	graphicImportAPI().SysCallN(7, m.Instance(), uintptr(FormatID))
}

func (m *TGraphic) LoadFromClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) {
	graphicImportAPI().SysCallN(8, m.Instance(), uintptr(ClipboardType), uintptr(FormatID))
}

func (m *TGraphic) SaveToFile(Filename string) {
	graphicImportAPI().SysCallN(21, m.Instance(), PascalStr(Filename))
}

func (m *TGraphic) SaveToStream(Stream IStream) {
	graphicImportAPI().SysCallN(22, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TGraphic) SaveToClipboardFormat(FormatID TClipboardFormat) {
	graphicImportAPI().SysCallN(19, m.Instance(), uintptr(FormatID))
}

func (m *TGraphic) SaveToClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) {
	graphicImportAPI().SysCallN(20, m.Instance(), uintptr(ClipboardType), uintptr(FormatID))
}

func (m *TGraphic) GetSupportedSourceMimeTypes(List IStrings) {
	graphicImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(List))
}

func (m *TGraphic) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	graphicImportAPI().SysCallN(23, m.Instance(), m.changePtr)
}

func (m *TGraphic) SetOnProgress(fn TProgressEvent) {
	if m.progressPtr != 0 {
		RemoveEventElement(m.progressPtr)
	}
	m.progressPtr = MakeEventDataPtr(fn)
	graphicImportAPI().SysCallN(24, m.Instance(), m.progressPtr)
}

var (
	graphicImport       *imports.Imports = nil
	graphicImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Graphic_Class", 0),
		/*1*/ imports.NewTable("Graphic_Clear", 0),
		/*2*/ imports.NewTable("Graphic_Empty", 0),
		/*3*/ imports.NewTable("Graphic_GetResourceType", 0),
		/*4*/ imports.NewTable("Graphic_GetSupportedSourceMimeTypes", 0),
		/*5*/ imports.NewTable("Graphic_Height", 0),
		/*6*/ imports.NewTable("Graphic_LazarusResourceTypeValid", 0),
		/*7*/ imports.NewTable("Graphic_LoadFromClipboardFormat", 0),
		/*8*/ imports.NewTable("Graphic_LoadFromClipboardFormatID", 0),
		/*9*/ imports.NewTable("Graphic_LoadFromFile", 0),
		/*10*/ imports.NewTable("Graphic_LoadFromLazarusResource", 0),
		/*11*/ imports.NewTable("Graphic_LoadFromMimeStream", 0),
		/*12*/ imports.NewTable("Graphic_LoadFromResourceID", 0),
		/*13*/ imports.NewTable("Graphic_LoadFromResourceName", 0),
		/*14*/ imports.NewTable("Graphic_LoadFromStream", 0),
		/*15*/ imports.NewTable("Graphic_MimeType", 0),
		/*16*/ imports.NewTable("Graphic_Modified", 0),
		/*17*/ imports.NewTable("Graphic_Palette", 0),
		/*18*/ imports.NewTable("Graphic_PaletteModified", 0),
		/*19*/ imports.NewTable("Graphic_SaveToClipboardFormat", 0),
		/*20*/ imports.NewTable("Graphic_SaveToClipboardFormatID", 0),
		/*21*/ imports.NewTable("Graphic_SaveToFile", 0),
		/*22*/ imports.NewTable("Graphic_SaveToStream", 0),
		/*23*/ imports.NewTable("Graphic_SetOnChange", 0),
		/*24*/ imports.NewTable("Graphic_SetOnProgress", 0),
		/*25*/ imports.NewTable("Graphic_Transparent", 0),
		/*26*/ imports.NewTable("Graphic_Width", 0),
	}
)

func graphicImportAPI() *imports.Imports {
	if graphicImport == nil {
		graphicImport = NewDefaultImports()
		graphicImport.SetImportTable(graphicImportTables)
		graphicImportTables = nil
	}
	return graphicImport
}
