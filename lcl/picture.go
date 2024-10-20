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

// IPicture Parent: IPersistent
type IPicture interface {
	IPersistent
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	Bitmap() IBitmap                                                                   // property
	SetBitmap(AValue IBitmap)                                                          // property
	Icon() IIcon                                                                       // property
	SetIcon(AValue IIcon)                                                              // property
	Jpeg() IJPEGImage                                                                  // property
	SetJpeg(AValue IJPEGImage)                                                         // property
	Pixmap() IPixmap                                                                   // property
	SetPixmap(AValue IPixmap)                                                          // property
	PNG() IPortableNetworkGraphic                                                      // property
	SetPNG(AValue IPortableNetworkGraphic)                                             // property
	PNM() IPortableAnyMapGraphic                                                       // property
	SetPNM(AValue IPortableAnyMapGraphic)                                              // property
	Graphic() IGraphic                                                                 // property
	SetGraphic(AValue IGraphic)                                                        // property
	Height() int32                                                                     // property
	Width() int32                                                                      // property
	Clear()                                                                            // procedure
	LoadFromClipboardFormat(FormatID TClipboardFormat)                                 // procedure
	LoadFromClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) // procedure
	LoadFromFile(Filename string)                                                      // procedure
	LoadFromResourceName(Instance THandle, ResName string)                             // procedure
	LoadFromResourceName1(Instance THandle, ResName string, AClass TGraphicClass)      // procedure
	LoadFromLazarusResource(AName string)                                              // procedure
	LoadFromStream(Stream IStream)                                                     // procedure
	LoadFromStreamWithFileExt(Stream IStream, FileExt string)                          // procedure
	SaveToClipboardFormat(FormatID TClipboardFormat)                                   // procedure
	SaveToFile(Filename string, FileExt string)                                        // procedure
	SaveToStream(Stream IStream)                                                       // procedure
	SaveToStreamWithFileExt(Stream IStream, FileExt string)                            // procedure
	SetOnChange(fn TNotifyEvent)                                                       // property event
	SetOnProgress(fn TProgressEvent)                                                   // property event
}

// TPicture Parent: TPersistent
type TPicture struct {
	TPersistent
	changePtr   uintptr
	progressPtr uintptr
}

func NewPicture() IPicture {
	r1 := pictureImportAPI().SysCallN(3)
	return AsPicture(r1)
}

func (m *TPicture) Bitmap() IBitmap {
	r1 := pictureImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TPicture) SetBitmap(AValue IBitmap) {
	pictureImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Icon() IIcon {
	r1 := pictureImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return AsIcon(r1)
}

func (m *TPicture) SetIcon(AValue IIcon) {
	pictureImportAPI().SysCallN(6, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Jpeg() IJPEGImage {
	r1 := pictureImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return AsJPEGImage(r1)
}

func (m *TPicture) SetJpeg(AValue IJPEGImage) {
	pictureImportAPI().SysCallN(7, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Pixmap() IPixmap {
	r1 := pictureImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return AsPixmap(r1)
}

func (m *TPicture) SetPixmap(AValue IPixmap) {
	pictureImportAPI().SysCallN(18, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) PNG() IPortableNetworkGraphic {
	r1 := pictureImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return AsPortableNetworkGraphic(r1)
}

func (m *TPicture) SetPNG(AValue IPortableNetworkGraphic) {
	pictureImportAPI().SysCallN(16, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) PNM() IPortableAnyMapGraphic {
	r1 := pictureImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return AsPortableAnyMapGraphic(r1)
}

func (m *TPicture) SetPNM(AValue IPortableAnyMapGraphic) {
	pictureImportAPI().SysCallN(17, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Graphic() IGraphic {
	r1 := pictureImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return AsGraphic(r1)
}

func (m *TPicture) SetGraphic(AValue IGraphic) {
	pictureImportAPI().SysCallN(4, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TPicture) Height() int32 {
	r1 := pictureImportAPI().SysCallN(5, m.Instance())
	return int32(r1)
}

func (m *TPicture) Width() int32 {
	r1 := pictureImportAPI().SysCallN(25, m.Instance())
	return int32(r1)
}

func PictureClass() TClass {
	ret := pictureImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TPicture) Clear() {
	pictureImportAPI().SysCallN(2, m.Instance())
}

func (m *TPicture) LoadFromClipboardFormat(FormatID TClipboardFormat) {
	pictureImportAPI().SysCallN(8, m.Instance(), uintptr(FormatID))
}

func (m *TPicture) LoadFromClipboardFormatID(ClipboardType TClipboardType, FormatID TClipboardFormat) {
	pictureImportAPI().SysCallN(9, m.Instance(), uintptr(ClipboardType), uintptr(FormatID))
}

func (m *TPicture) LoadFromFile(Filename string) {
	pictureImportAPI().SysCallN(10, m.Instance(), PascalStr(Filename))
}

func (m *TPicture) LoadFromResourceName(Instance THandle, ResName string) {
	pictureImportAPI().SysCallN(12, m.Instance(), uintptr(Instance), PascalStr(ResName))
}

func (m *TPicture) LoadFromResourceName1(Instance THandle, ResName string, AClass TGraphicClass) {
	pictureImportAPI().SysCallN(13, m.Instance(), uintptr(Instance), PascalStr(ResName), uintptr(AClass))
}

func (m *TPicture) LoadFromLazarusResource(AName string) {
	pictureImportAPI().SysCallN(11, m.Instance(), PascalStr(AName))
}

func (m *TPicture) LoadFromStream(Stream IStream) {
	pictureImportAPI().SysCallN(14, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TPicture) LoadFromStreamWithFileExt(Stream IStream, FileExt string) {
	pictureImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(Stream), PascalStr(FileExt))
}

func (m *TPicture) SaveToClipboardFormat(FormatID TClipboardFormat) {
	pictureImportAPI().SysCallN(19, m.Instance(), uintptr(FormatID))
}

func (m *TPicture) SaveToFile(Filename string, FileExt string) {
	pictureImportAPI().SysCallN(20, m.Instance(), PascalStr(Filename), PascalStr(FileExt))
}

func (m *TPicture) SaveToStream(Stream IStream) {
	pictureImportAPI().SysCallN(21, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TPicture) SaveToStreamWithFileExt(Stream IStream, FileExt string) {
	pictureImportAPI().SysCallN(22, m.Instance(), GetObjectUintptr(Stream), PascalStr(FileExt))
}

func (m *TPicture) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	pictureImportAPI().SysCallN(23, m.Instance(), m.changePtr)
}

func (m *TPicture) SetOnProgress(fn TProgressEvent) {
	if m.progressPtr != 0 {
		RemoveEventElement(m.progressPtr)
	}
	m.progressPtr = MakeEventDataPtr(fn)
	pictureImportAPI().SysCallN(24, m.Instance(), m.progressPtr)
}

var (
	pictureImport       *imports.Imports = nil
	pictureImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Picture_Bitmap", 0),
		/*1*/ imports.NewTable("Picture_Class", 0),
		/*2*/ imports.NewTable("Picture_Clear", 0),
		/*3*/ imports.NewTable("Picture_Create", 0),
		/*4*/ imports.NewTable("Picture_Graphic", 0),
		/*5*/ imports.NewTable("Picture_Height", 0),
		/*6*/ imports.NewTable("Picture_Icon", 0),
		/*7*/ imports.NewTable("Picture_Jpeg", 0),
		/*8*/ imports.NewTable("Picture_LoadFromClipboardFormat", 0),
		/*9*/ imports.NewTable("Picture_LoadFromClipboardFormatID", 0),
		/*10*/ imports.NewTable("Picture_LoadFromFile", 0),
		/*11*/ imports.NewTable("Picture_LoadFromLazarusResource", 0),
		/*12*/ imports.NewTable("Picture_LoadFromResourceName", 0),
		/*13*/ imports.NewTable("Picture_LoadFromResourceName1", 0),
		/*14*/ imports.NewTable("Picture_LoadFromStream", 0),
		/*15*/ imports.NewTable("Picture_LoadFromStreamWithFileExt", 0),
		/*16*/ imports.NewTable("Picture_PNG", 0),
		/*17*/ imports.NewTable("Picture_PNM", 0),
		/*18*/ imports.NewTable("Picture_Pixmap", 0),
		/*19*/ imports.NewTable("Picture_SaveToClipboardFormat", 0),
		/*20*/ imports.NewTable("Picture_SaveToFile", 0),
		/*21*/ imports.NewTable("Picture_SaveToStream", 0),
		/*22*/ imports.NewTable("Picture_SaveToStreamWithFileExt", 0),
		/*23*/ imports.NewTable("Picture_SetOnChange", 0),
		/*24*/ imports.NewTable("Picture_SetOnProgress", 0),
		/*25*/ imports.NewTable("Picture_Width", 0),
	}
)

func pictureImportAPI() *imports.Imports {
	if pictureImport == nil {
		pictureImport = NewDefaultImports()
		pictureImport.SetImportTable(pictureImportTables)
		pictureImportTables = nil
	}
	return pictureImport
}
