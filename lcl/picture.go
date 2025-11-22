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

// IPicture Parent: IPersistent
type IPicture interface {
	IPersistent
	Clear() // procedure
	// LoadFromClipboardFormat
	//  load methods
	LoadFromClipboardFormat(formatID types.TClipboardFormat)                                                       // procedure
	LoadFromClipboardFormatID(clipboardType types.TClipboardType, formatID types.TClipboardFormat)                 // procedure
	LoadFromFile(filename string)                                                                                  // procedure
	LoadFromResourceNameWithLCLHandleString(instance types.TLCLHandle, resName string)                             // procedure
	LoadFromResourceNameWithLCLHandleStringGraphicClass(instance types.TLCLHandle, resName string, class IGraphic) // procedure
	LoadFromLazarusResource(name string)                                                                           // procedure
	LoadFromStream(stream IStream)                                                                                 // procedure
	LoadFromStreamWithFileExt(stream IStream, fileExt string)                                                      // procedure
	// SaveToClipboardFormat
	//  save methods
	SaveToClipboardFormat(formatID types.TClipboardFormat)  // procedure
	SaveToFile(filename string, fileExt string)             // procedure
	SaveToStream(stream IStream)                            // procedure
	SaveToStreamWithFileExt(stream IStream, fileExt string) // procedure
	Bitmap() IBitmap                                        // property Bitmap Getter
	SetBitmap(value IBitmap)                                // property Bitmap Setter
	Icon() IIcon                                            // property Icon Getter
	SetIcon(value IIcon)                                    // property Icon Setter
	Jpeg() IJPEGImage                                       // property Jpeg Getter
	SetJpeg(value IJPEGImage)                               // property Jpeg Setter
	Pixmap() IPixmap                                        // property Pixmap Getter
	SetPixmap(value IPixmap)                                // property Pixmap Setter
	PNG() IPortableNetworkGraphic                           // property PNG Getter
	SetPNG(value IPortableNetworkGraphic)                   // property PNG Setter
	PNM() IPortableAnyMapGraphic                            // property PNM Getter
	SetPNM(value IPortableAnyMapGraphic)                    // property PNM Setter
	Graphic() IGraphic                                      // property Graphic Getter
	SetGraphic(value IGraphic)                              // property Graphic Setter
	// Height
	//  property PictureAdapter: IChangeNotifier read FNotify write FNotify;
	Height() int32                   // property Height Getter
	Width() int32                    // property Width Getter
	SetOnChange(fn TNotifyEvent)     // property event
	SetOnProgress(fn TProgressEvent) // property event
}

type TPicture struct {
	TPersistent
}

func (m *TPicture) Clear() {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(3, m.Instance())
}

func (m *TPicture) LoadFromClipboardFormat(formatID types.TClipboardFormat) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(4, m.Instance(), uintptr(formatID))
}

func (m *TPicture) LoadFromClipboardFormatID(clipboardType types.TClipboardType, formatID types.TClipboardFormat) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(5, m.Instance(), uintptr(clipboardType), uintptr(formatID))
}

func (m *TPicture) LoadFromFile(filename string) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(6, m.Instance(), api.PasStr(filename))
}

func (m *TPicture) LoadFromResourceNameWithLCLHandleString(instance types.TLCLHandle, resName string) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(7, m.Instance(), uintptr(instance), api.PasStr(resName))
}

func (m *TPicture) LoadFromResourceNameWithLCLHandleStringGraphicClass(instance types.TLCLHandle, resName string, class IGraphic) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(8, m.Instance(), uintptr(instance), api.PasStr(resName), base.GetObjectUintptr(class))
}

func (m *TPicture) LoadFromLazarusResource(name string) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(9, m.Instance(), api.PasStr(name))
}

func (m *TPicture) LoadFromStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TPicture) LoadFromStreamWithFileExt(stream IStream, fileExt string) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(stream), api.PasStr(fileExt))
}

func (m *TPicture) SaveToClipboardFormat(formatID types.TClipboardFormat) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(12, m.Instance(), uintptr(formatID))
}

func (m *TPicture) SaveToFile(filename string, fileExt string) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(13, m.Instance(), api.PasStr(filename), api.PasStr(fileExt))
}

func (m *TPicture) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TPicture) SaveToStreamWithFileExt(stream IStream, fileExt string) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(stream), api.PasStr(fileExt))
}

func (m *TPicture) Bitmap() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := pictureAPI().SysCallN(19, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TPicture) SetBitmap(value IBitmap) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(19, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPicture) Icon() IIcon {
	if !m.IsValid() {
		return nil
	}
	r := pictureAPI().SysCallN(20, 0, m.Instance())
	return AsIcon(r)
}

func (m *TPicture) SetIcon(value IIcon) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(20, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPicture) Jpeg() IJPEGImage {
	if !m.IsValid() {
		return nil
	}
	r := pictureAPI().SysCallN(21, 0, m.Instance())
	return AsJPEGImage(r)
}

func (m *TPicture) SetJpeg(value IJPEGImage) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(21, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPicture) Pixmap() IPixmap {
	if !m.IsValid() {
		return nil
	}
	r := pictureAPI().SysCallN(22, 0, m.Instance())
	return AsPixmap(r)
}

func (m *TPicture) SetPixmap(value IPixmap) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPicture) PNG() IPortableNetworkGraphic {
	if !m.IsValid() {
		return nil
	}
	r := pictureAPI().SysCallN(23, 0, m.Instance())
	return AsPortableNetworkGraphic(r)
}

func (m *TPicture) SetPNG(value IPortableNetworkGraphic) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(23, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPicture) PNM() IPortableAnyMapGraphic {
	if !m.IsValid() {
		return nil
	}
	r := pictureAPI().SysCallN(24, 0, m.Instance())
	return AsPortableAnyMapGraphic(r)
}

func (m *TPicture) SetPNM(value IPortableAnyMapGraphic) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(24, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPicture) Graphic() IGraphic {
	if !m.IsValid() {
		return nil
	}
	r := pictureAPI().SysCallN(25, 0, m.Instance())
	return AsGraphic(r)
}

func (m *TPicture) SetGraphic(value IGraphic) {
	if !m.IsValid() {
		return
	}
	pictureAPI().SysCallN(25, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TPicture) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pictureAPI().SysCallN(26, m.Instance())
	return int32(r)
}

func (m *TPicture) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := pictureAPI().SysCallN(27, m.Instance())
	return int32(r)
}

func (m *TPicture) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 28, pictureAPI(), api.MakeEventDataPtr(cb))
}

func (m *TPicture) SetOnProgress(fn TProgressEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTProgressEvent(fn)
	base.SetEvent(m, 29, pictureAPI(), api.MakeEventDataPtr(cb))
}

// Picture  is static instance
var Picture _PictureClass

// _PictureClass is class type defined by TPicture
type _PictureClass uintptr

func (_PictureClass) SupportsClipboardFormat(formatID types.TClipboardFormat) bool {
	r := pictureAPI().SysCallN(1, uintptr(formatID))
	return api.GoBool(r)
}

func (_PictureClass) FindGraphicClassWithFileExt(ext string, exceptionOnNotFound bool) IGraphic {
	r := pictureAPI().SysCallN(2, api.PasStr(ext), api.PasBool(exceptionOnNotFound))
	return AsGraphic(r)
}

func (_PictureClass) RegisterFileFormat(anExtension string, description string, graphicClass IGraphic) {
	pictureAPI().SysCallN(16, api.PasStr(anExtension), api.PasStr(description), base.GetObjectUintptr(graphicClass))
}

func (_PictureClass) RegisterClipboardFormat(formatID types.TClipboardFormat, graphicClass IGraphic) {
	pictureAPI().SysCallN(17, uintptr(formatID), base.GetObjectUintptr(graphicClass))
}

func (_PictureClass) UnregisterGraphicClass(class IGraphic) {
	pictureAPI().SysCallN(18, base.GetObjectUintptr(class))
}

// NewPicture class constructor
func NewPicture() IPicture {
	r := pictureAPI().SysCallN(0)
	return AsPicture(r)
}

var (
	pictureOnce   base.Once
	pictureImport *imports.Imports = nil
)

func pictureAPI() *imports.Imports {
	pictureOnce.Do(func() {
		pictureImport = api.NewDefaultImports()
		pictureImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TPicture_Create", 0), // constructor NewPicture
			/* 1 */ imports.NewTable("TPicture_SupportsClipboardFormat", 0), // static function SupportsClipboardFormat
			/* 2 */ imports.NewTable("TPicture_FindGraphicClassWithFileExt", 0), // static function FindGraphicClassWithFileExt
			/* 3 */ imports.NewTable("TPicture_Clear", 0), // procedure Clear
			/* 4 */ imports.NewTable("TPicture_LoadFromClipboardFormat", 0), // procedure LoadFromClipboardFormat
			/* 5 */ imports.NewTable("TPicture_LoadFromClipboardFormatID", 0), // procedure LoadFromClipboardFormatID
			/* 6 */ imports.NewTable("TPicture_LoadFromFile", 0), // procedure LoadFromFile
			/* 7 */ imports.NewTable("TPicture_LoadFromResourceNameWithLCLHandleString", 0), // procedure LoadFromResourceNameWithLCLHandleString
			/* 8 */ imports.NewTable("TPicture_LoadFromResourceNameWithLCLHandleStringGraphicClass", 0), // procedure LoadFromResourceNameWithLCLHandleStringGraphicClass
			/* 9 */ imports.NewTable("TPicture_LoadFromLazarusResource", 0), // procedure LoadFromLazarusResource
			/* 10 */ imports.NewTable("TPicture_LoadFromStream", 0), // procedure LoadFromStream
			/* 11 */ imports.NewTable("TPicture_LoadFromStreamWithFileExt", 0), // procedure LoadFromStreamWithFileExt
			/* 12 */ imports.NewTable("TPicture_SaveToClipboardFormat", 0), // procedure SaveToClipboardFormat
			/* 13 */ imports.NewTable("TPicture_SaveToFile", 0), // procedure SaveToFile
			/* 14 */ imports.NewTable("TPicture_SaveToStream", 0), // procedure SaveToStream
			/* 15 */ imports.NewTable("TPicture_SaveToStreamWithFileExt", 0), // procedure SaveToStreamWithFileExt
			/* 16 */ imports.NewTable("TPicture_RegisterFileFormat", 0), // static procedure RegisterFileFormat
			/* 17 */ imports.NewTable("TPicture_RegisterClipboardFormat", 0), // static procedure RegisterClipboardFormat
			/* 18 */ imports.NewTable("TPicture_UnregisterGraphicClass", 0), // static procedure UnregisterGraphicClass
			/* 19 */ imports.NewTable("TPicture_Bitmap", 0), // property Bitmap
			/* 20 */ imports.NewTable("TPicture_Icon", 0), // property Icon
			/* 21 */ imports.NewTable("TPicture_Jpeg", 0), // property Jpeg
			/* 22 */ imports.NewTable("TPicture_Pixmap", 0), // property Pixmap
			/* 23 */ imports.NewTable("TPicture_PNG", 0), // property PNG
			/* 24 */ imports.NewTable("TPicture_PNM", 0), // property PNM
			/* 25 */ imports.NewTable("TPicture_Graphic", 0), // property Graphic
			/* 26 */ imports.NewTable("TPicture_Height", 0), // property Height
			/* 27 */ imports.NewTable("TPicture_Width", 0), // property Width
			/* 28 */ imports.NewTable("TPicture_OnChange", 0), // event OnChange
			/* 29 */ imports.NewTable("TPicture_OnProgress", 0), // event OnProgress
		}
	})
	return pictureImport
}
