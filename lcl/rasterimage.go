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

// IRasterImage Is Abstract Class Parent: IGraphic
type IRasterImage interface {
	IGraphic
	Canvas() ICanvas                                            // property
	BitmapHandle() HBITMAP                                      // property
	SetBitmapHandle(AValue HBITMAP)                             // property
	Masked() bool                                               // property
	SetMasked(AValue bool)                                      // property
	MaskHandle() HBITMAP                                        // property
	SetMaskHandle(AValue HBITMAP)                               // property
	PixelFormat() TPixelFormat                                  // property
	SetPixelFormat(AValue TPixelFormat)                         // property
	RawImage() (resultRawImage TRawImage)                       // property
	ScanLine(Row int32) uintptr                                 // property
	TransparentColor() TColor                                   // property
	SetTransparentColor(AValue TColor)                          // property
	TransparentMode() TTransparentMode                          // property
	SetTransparentMode(AValue TTransparentMode)                 // property
	BitmapHandleAllocated() bool                                // function Is Abstract
	MaskHandleAllocated() bool                                  // function Is Abstract
	PaletteAllocated() bool                                     // function Is Abstract
	ReleaseBitmapHandle() HBITMAP                               // function
	ReleaseMaskHandle() HBITMAP                                 // function
	ReleasePalette() HPALETTE                                   // function
	HandleAllocated() bool                                      // function
	BeginUpdate(ACanvasOnly bool)                               // procedure
	EndUpdate(AStreamIsValid bool)                              // procedure
	FreeImage()                                                 // procedure
	LoadFromBitmapHandles(ABitmap, AMask HBITMAP, ARect *TRect) // procedure
	LoadFromDevice(DC HDC)                                      // procedure
	LoadFromStreamForStream(AStream IStream, ASize uint32)      // procedure
	LoadFromRawImage(AIMage *TRawImage, ADataOwner bool)        // procedure
	GetSize(OutWidth, OutHeight *int32)                         // procedure
	Mask(ATransparentColor TColor)                              // procedure
	SetHandles(ABitmap, AMask HBITMAP)                          // procedure Is Abstract
}

// TRasterImage Is Abstract Class Parent: TGraphic
type TRasterImage struct {
	TGraphic
}

func (m *TRasterImage) Canvas() ICanvas {
	r1 := rasterImageImportAPI().SysCallN(3, m.Instance())
	return AsCanvas(r1)
}

func (m *TRasterImage) BitmapHandle() HBITMAP {
	r1 := rasterImageImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return HBITMAP(r1)
}

func (m *TRasterImage) SetBitmapHandle(AValue HBITMAP) {
	rasterImageImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) Masked() bool {
	r1 := rasterImageImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRasterImage) SetMasked(AValue bool) {
	rasterImageImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRasterImage) MaskHandle() HBITMAP {
	r1 := rasterImageImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return HBITMAP(r1)
}

func (m *TRasterImage) SetMaskHandle(AValue HBITMAP) {
	rasterImageImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) PixelFormat() TPixelFormat {
	r1 := rasterImageImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TPixelFormat(r1)
}

func (m *TRasterImage) SetPixelFormat(AValue TPixelFormat) {
	rasterImageImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) RawImage() (resultRawImage TRawImage) {
	r1 := rasterImageImportAPI().SysCallN(19, m.Instance())
	return *(*TRawImage)(getPointer(r1))
}

func (m *TRasterImage) ScanLine(Row int32) uintptr {
	r1 := rasterImageImportAPI().SysCallN(23, m.Instance(), uintptr(Row))
	return uintptr(r1)
}

func (m *TRasterImage) TransparentColor() TColor {
	r1 := rasterImageImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TRasterImage) SetTransparentColor(AValue TColor) {
	rasterImageImportAPI().SysCallN(25, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) TransparentMode() TTransparentMode {
	r1 := rasterImageImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return TTransparentMode(r1)
}

func (m *TRasterImage) SetTransparentMode(AValue TTransparentMode) {
	rasterImageImportAPI().SysCallN(26, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) BitmapHandleAllocated() bool {
	r1 := rasterImageImportAPI().SysCallN(2, m.Instance())
	return GoBool(r1)
}

func (m *TRasterImage) MaskHandleAllocated() bool {
	r1 := rasterImageImportAPI().SysCallN(15, m.Instance())
	return GoBool(r1)
}

func (m *TRasterImage) PaletteAllocated() bool {
	r1 := rasterImageImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

func (m *TRasterImage) ReleaseBitmapHandle() HBITMAP {
	r1 := rasterImageImportAPI().SysCallN(20, m.Instance())
	return HBITMAP(r1)
}

func (m *TRasterImage) ReleaseMaskHandle() HBITMAP {
	r1 := rasterImageImportAPI().SysCallN(21, m.Instance())
	return HBITMAP(r1)
}

func (m *TRasterImage) ReleasePalette() HPALETTE {
	r1 := rasterImageImportAPI().SysCallN(22, m.Instance())
	return HPALETTE(r1)
}

func (m *TRasterImage) HandleAllocated() bool {
	r1 := rasterImageImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func RasterImageClass() TClass {
	ret := rasterImageImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TRasterImage) BeginUpdate(ACanvasOnly bool) {
	rasterImageImportAPI().SysCallN(0, m.Instance(), PascalBool(ACanvasOnly))
}

func (m *TRasterImage) EndUpdate(AStreamIsValid bool) {
	rasterImageImportAPI().SysCallN(5, m.Instance(), PascalBool(AStreamIsValid))
}

func (m *TRasterImage) FreeImage() {
	rasterImageImportAPI().SysCallN(6, m.Instance())
}

func (m *TRasterImage) LoadFromBitmapHandles(ABitmap, AMask HBITMAP, ARect *TRect) {
	rasterImageImportAPI().SysCallN(9, m.Instance(), uintptr(ABitmap), uintptr(AMask), uintptr(unsafePointer(ARect)))
}

func (m *TRasterImage) LoadFromDevice(DC HDC) {
	rasterImageImportAPI().SysCallN(10, m.Instance(), uintptr(DC))
}

func (m *TRasterImage) LoadFromStreamForStream(AStream IStream, ASize uint32) {
	rasterImageImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(AStream), uintptr(ASize))
}

func (m *TRasterImage) LoadFromRawImage(AIMage *TRawImage, ADataOwner bool) {
	rasterImageImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(AIMage)), PascalBool(ADataOwner))
}

func (m *TRasterImage) GetSize(OutWidth, OutHeight *int32) {
	var result0 uintptr
	var result1 uintptr
	rasterImageImportAPI().SysCallN(7, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*OutWidth = int32(result0)
	*OutHeight = int32(result1)
}

func (m *TRasterImage) Mask(ATransparentColor TColor) {
	rasterImageImportAPI().SysCallN(13, m.Instance(), uintptr(ATransparentColor))
}

func (m *TRasterImage) SetHandles(ABitmap, AMask HBITMAP) {
	rasterImageImportAPI().SysCallN(24, m.Instance(), uintptr(ABitmap), uintptr(AMask))
}

var (
	rasterImageImport       *imports.Imports = nil
	rasterImageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("RasterImage_BeginUpdate", 0),
		/*1*/ imports.NewTable("RasterImage_BitmapHandle", 0),
		/*2*/ imports.NewTable("RasterImage_BitmapHandleAllocated", 0),
		/*3*/ imports.NewTable("RasterImage_Canvas", 0),
		/*4*/ imports.NewTable("RasterImage_Class", 0),
		/*5*/ imports.NewTable("RasterImage_EndUpdate", 0),
		/*6*/ imports.NewTable("RasterImage_FreeImage", 0),
		/*7*/ imports.NewTable("RasterImage_GetSize", 0),
		/*8*/ imports.NewTable("RasterImage_HandleAllocated", 0),
		/*9*/ imports.NewTable("RasterImage_LoadFromBitmapHandles", 0),
		/*10*/ imports.NewTable("RasterImage_LoadFromDevice", 0),
		/*11*/ imports.NewTable("RasterImage_LoadFromRawImage", 0),
		/*12*/ imports.NewTable("RasterImage_LoadFromStreamForStream", 0),
		/*13*/ imports.NewTable("RasterImage_Mask", 0),
		/*14*/ imports.NewTable("RasterImage_MaskHandle", 0),
		/*15*/ imports.NewTable("RasterImage_MaskHandleAllocated", 0),
		/*16*/ imports.NewTable("RasterImage_Masked", 0),
		/*17*/ imports.NewTable("RasterImage_PaletteAllocated", 0),
		/*18*/ imports.NewTable("RasterImage_PixelFormat", 0),
		/*19*/ imports.NewTable("RasterImage_RawImage", 0),
		/*20*/ imports.NewTable("RasterImage_ReleaseBitmapHandle", 0),
		/*21*/ imports.NewTable("RasterImage_ReleaseMaskHandle", 0),
		/*22*/ imports.NewTable("RasterImage_ReleasePalette", 0),
		/*23*/ imports.NewTable("RasterImage_ScanLine", 0),
		/*24*/ imports.NewTable("RasterImage_SetHandles", 0),
		/*25*/ imports.NewTable("RasterImage_TransparentColor", 0),
		/*26*/ imports.NewTable("RasterImage_TransparentMode", 0),
	}
)

func rasterImageImportAPI() *imports.Imports {
	if rasterImageImport == nil {
		rasterImageImport = NewDefaultImports()
		rasterImageImport.SetImportTable(rasterImageImportTables)
		rasterImageImportTables = nil
	}
	return rasterImageImport
}
