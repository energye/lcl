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

// IRasterImage Parent: IGraphic
type IRasterImage interface {
	IGraphic
	BitmapHandleAllocated() bool                                                      // function
	MaskHandleAllocated() bool                                                        // function
	PaletteAllocated() bool                                                           // function
	ReleaseBitmapHandle() types.HBitmap                                               // function
	ReleaseMaskHandle() types.HBitmap                                                 // function
	ReleasePalette() types.HPALETTE                                                   // function
	CreateIntfImage() ILazIntfImage                                                   // function
	HandleAllocated() bool                                                            // function
	BeginUpdate(canvasOnly bool)                                                      // procedure
	EndUpdate(streamIsValid bool)                                                     // procedure
	FreeImage()                                                                       // procedure
	LoadFromBitmapHandles(bitmap types.HBitmap, mask types.HBitmap, rect types.TRect) // procedure
	LoadFromDevice(dC types.HDC)                                                      // procedure
	LoadFromStreamWithStreamCardinal(stream IStream, size uint32)                     // procedure
	LoadFromRawImage(iMage IRawImageWrap, dataOwner bool)                             // procedure
	LoadFromIntfImage(intfImage ILazIntfImage)                                        // procedure
	GetSize(outWidth *int32, outHeight *int32)                                        // procedure
	Mask(transparentColor types.TColor)                                               // procedure
	SetHandles(bitmap types.HBitmap, mask types.HBitmap)                              // procedure
	Canvas() ICanvas                                                                  // property Canvas Getter
	BitmapHandle() types.HBitmap                                                      // property BitmapHandle Getter
	SetBitmapHandle(value types.HBitmap)                                              // property BitmapHandle Setter
	Masked() bool                                                                     // property Masked Getter
	SetMasked(value bool)                                                             // property Masked Setter
	MaskHandle() types.HBitmap                                                        // property MaskHandle Getter
	SetMaskHandle(value types.HBitmap)                                                // property MaskHandle Setter
	PixelFormat() types.TPixelFormat                                                  // property PixelFormat Getter
	SetPixelFormat(value types.TPixelFormat)                                          // property PixelFormat Setter
	RawImage() IRawImageWrap                                                          // property RawImage Getter
	ScanLine(row int32) uintptr                                                       // property ScanLine Getter
	TransparentColor() types.TColor                                                   // property TransparentColor Getter
	SetTransparentColor(value types.TColor)                                           // property TransparentColor Setter
	TransparentMode() types.TTransparentMode                                          // property TransparentMode Getter
	SetTransparentMode(value types.TTransparentMode)                                  // property TransparentMode Setter
}

type TRasterImage struct {
	TGraphic
}

func (m *TRasterImage) BitmapHandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := rasterImageAPI().SysCallN(0, m.Instance())
	return api.GoBool(r)
}

func (m *TRasterImage) MaskHandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := rasterImageAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TRasterImage) PaletteAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := rasterImageAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TRasterImage) ReleaseBitmapHandle() types.HBitmap {
	if !m.IsValid() {
		return 0
	}
	r := rasterImageAPI().SysCallN(3, m.Instance())
	return types.HBitmap(r)
}

func (m *TRasterImage) ReleaseMaskHandle() types.HBitmap {
	if !m.IsValid() {
		return 0
	}
	r := rasterImageAPI().SysCallN(4, m.Instance())
	return types.HBitmap(r)
}

func (m *TRasterImage) ReleasePalette() types.HPALETTE {
	if !m.IsValid() {
		return 0
	}
	r := rasterImageAPI().SysCallN(5, m.Instance())
	return types.HPALETTE(r)
}

func (m *TRasterImage) CreateIntfImage() ILazIntfImage {
	if !m.IsValid() {
		return nil
	}
	r := rasterImageAPI().SysCallN(6, m.Instance())
	return AsLazIntfImage(r)
}

func (m *TRasterImage) HandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := rasterImageAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TRasterImage) BeginUpdate(canvasOnly bool) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(8, m.Instance(), api.PasBool(canvasOnly))
}

func (m *TRasterImage) EndUpdate(streamIsValid bool) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(9, m.Instance(), api.PasBool(streamIsValid))
}

func (m *TRasterImage) FreeImage() {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(10, m.Instance())
}

func (m *TRasterImage) LoadFromBitmapHandles(bitmap types.HBitmap, mask types.HBitmap, rect types.TRect) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(11, m.Instance(), uintptr(bitmap), uintptr(mask), uintptr(base.UnsafePointer(&rect)))
}

func (m *TRasterImage) LoadFromDevice(dC types.HDC) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(12, m.Instance(), uintptr(dC))
}

func (m *TRasterImage) LoadFromStreamWithStreamCardinal(stream IStream, size uint32) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(stream), uintptr(size))
}

func (m *TRasterImage) LoadFromRawImage(iMage IRawImageWrap, dataOwner bool) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(iMage), api.PasBool(dataOwner))
}

func (m *TRasterImage) LoadFromIntfImage(intfImage ILazIntfImage) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(intfImage))
}

func (m *TRasterImage) GetSize(outWidth *int32, outHeight *int32) {
	if !m.IsValid() {
		return
	}
	var widthPtr uintptr
	var heightPtr uintptr
	rasterImageAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&widthPtr)), uintptr(base.UnsafePointer(&heightPtr)))
	*outWidth = int32(widthPtr)
	*outHeight = int32(heightPtr)
}

func (m *TRasterImage) Mask(transparentColor types.TColor) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(17, m.Instance(), uintptr(transparentColor))
}

func (m *TRasterImage) SetHandles(bitmap types.HBitmap, mask types.HBitmap) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(18, m.Instance(), uintptr(bitmap), uintptr(mask))
}

func (m *TRasterImage) Canvas() ICanvas {
	if !m.IsValid() {
		return nil
	}
	r := rasterImageAPI().SysCallN(19, m.Instance())
	return AsCanvas(r)
}

func (m *TRasterImage) BitmapHandle() types.HBitmap {
	if !m.IsValid() {
		return 0
	}
	r := rasterImageAPI().SysCallN(20, 0, m.Instance())
	return types.HBitmap(r)
}

func (m *TRasterImage) SetBitmapHandle(value types.HBitmap) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TRasterImage) Masked() bool {
	if !m.IsValid() {
		return false
	}
	r := rasterImageAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TRasterImage) SetMasked(value bool) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TRasterImage) MaskHandle() types.HBitmap {
	if !m.IsValid() {
		return 0
	}
	r := rasterImageAPI().SysCallN(22, 0, m.Instance())
	return types.HBitmap(r)
}

func (m *TRasterImage) SetMaskHandle(value types.HBitmap) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TRasterImage) PixelFormat() types.TPixelFormat {
	if !m.IsValid() {
		return 0
	}
	r := rasterImageAPI().SysCallN(23, 0, m.Instance())
	return types.TPixelFormat(r)
}

func (m *TRasterImage) SetPixelFormat(value types.TPixelFormat) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TRasterImage) RawImage() IRawImageWrap {
	if !m.IsValid() {
		return nil
	}
	r := rasterImageAPI().SysCallN(24, m.Instance())
	return AsRawImageWrap(r)
}

func (m *TRasterImage) ScanLine(row int32) uintptr {
	if !m.IsValid() {
		return 0
	}
	r := rasterImageAPI().SysCallN(25, m.Instance(), uintptr(row))
	return uintptr(r)
}

func (m *TRasterImage) TransparentColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := rasterImageAPI().SysCallN(26, 0, m.Instance())
	return types.TColor(r)
}

func (m *TRasterImage) SetTransparentColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TRasterImage) TransparentMode() types.TTransparentMode {
	if !m.IsValid() {
		return 0
	}
	r := rasterImageAPI().SysCallN(27, 0, m.Instance())
	return types.TTransparentMode(r)
}

func (m *TRasterImage) SetTransparentMode(value types.TTransparentMode) {
	if !m.IsValid() {
		return
	}
	rasterImageAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

var (
	rasterImageOnce   base.Once
	rasterImageImport *imports.Imports = nil
)

func rasterImageAPI() *imports.Imports {
	rasterImageOnce.Do(func() {
		rasterImageImport = api.NewDefaultImports()
		rasterImageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRasterImage_BitmapHandleAllocated", 0), // function BitmapHandleAllocated
			/* 1 */ imports.NewTable("TRasterImage_MaskHandleAllocated", 0), // function MaskHandleAllocated
			/* 2 */ imports.NewTable("TRasterImage_PaletteAllocated", 0), // function PaletteAllocated
			/* 3 */ imports.NewTable("TRasterImage_ReleaseBitmapHandle", 0), // function ReleaseBitmapHandle
			/* 4 */ imports.NewTable("TRasterImage_ReleaseMaskHandle", 0), // function ReleaseMaskHandle
			/* 5 */ imports.NewTable("TRasterImage_ReleasePalette", 0), // function ReleasePalette
			/* 6 */ imports.NewTable("TRasterImage_CreateIntfImage", 0), // function CreateIntfImage
			/* 7 */ imports.NewTable("TRasterImage_HandleAllocated", 0), // function HandleAllocated
			/* 8 */ imports.NewTable("TRasterImage_BeginUpdate", 0), // procedure BeginUpdate
			/* 9 */ imports.NewTable("TRasterImage_EndUpdate", 0), // procedure EndUpdate
			/* 10 */ imports.NewTable("TRasterImage_FreeImage", 0), // procedure FreeImage
			/* 11 */ imports.NewTable("TRasterImage_LoadFromBitmapHandles", 0), // procedure LoadFromBitmapHandles
			/* 12 */ imports.NewTable("TRasterImage_LoadFromDevice", 0), // procedure LoadFromDevice
			/* 13 */ imports.NewTable("TRasterImage_LoadFromStreamWithStreamCardinal", 0), // procedure LoadFromStreamWithStreamCardinal
			/* 14 */ imports.NewTable("TRasterImage_LoadFromRawImage", 0), // procedure LoadFromRawImage
			/* 15 */ imports.NewTable("TRasterImage_LoadFromIntfImage", 0), // procedure LoadFromIntfImage
			/* 16 */ imports.NewTable("TRasterImage_GetSize", 0), // procedure GetSize
			/* 17 */ imports.NewTable("TRasterImage_Mask", 0), // procedure Mask
			/* 18 */ imports.NewTable("TRasterImage_SetHandles", 0), // procedure SetHandles
			/* 19 */ imports.NewTable("TRasterImage_Canvas", 0), // property Canvas
			/* 20 */ imports.NewTable("TRasterImage_BitmapHandle", 0), // property BitmapHandle
			/* 21 */ imports.NewTable("TRasterImage_Masked", 0), // property Masked
			/* 22 */ imports.NewTable("TRasterImage_MaskHandle", 0), // property MaskHandle
			/* 23 */ imports.NewTable("TRasterImage_PixelFormat", 0), // property PixelFormat
			/* 24 */ imports.NewTable("TRasterImage_RawImage", 0), // property RawImage
			/* 25 */ imports.NewTable("TRasterImage_ScanLine", 0), // property ScanLine
			/* 26 */ imports.NewTable("TRasterImage_TransparentColor", 0), // property TransparentColor
			/* 27 */ imports.NewTable("TRasterImage_TransparentMode", 0), // property TransparentMode
		}
	})
	return rasterImageImport
}
