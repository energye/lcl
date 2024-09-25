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
	r1 := LCL().SysCallN(4756, m.Instance())
	return AsCanvas(r1)
}

func (m *TRasterImage) BitmapHandle() HBITMAP {
	r1 := LCL().SysCallN(4754, 0, m.Instance(), 0)
	return HBITMAP(r1)
}

func (m *TRasterImage) SetBitmapHandle(AValue HBITMAP) {
	LCL().SysCallN(4754, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) Masked() bool {
	r1 := LCL().SysCallN(4769, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRasterImage) SetMasked(AValue bool) {
	LCL().SysCallN(4769, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRasterImage) MaskHandle() HBITMAP {
	r1 := LCL().SysCallN(4767, 0, m.Instance(), 0)
	return HBITMAP(r1)
}

func (m *TRasterImage) SetMaskHandle(AValue HBITMAP) {
	LCL().SysCallN(4767, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) PixelFormat() TPixelFormat {
	r1 := LCL().SysCallN(4771, 0, m.Instance(), 0)
	return TPixelFormat(r1)
}

func (m *TRasterImage) SetPixelFormat(AValue TPixelFormat) {
	LCL().SysCallN(4771, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) RawImage() (resultRawImage TRawImage) {
	r1 := LCL().SysCallN(4772, m.Instance())
	return *(*TRawImage)(getPointer(r1))
}

func (m *TRasterImage) ScanLine(Row int32) uintptr {
	r1 := LCL().SysCallN(4776, m.Instance(), uintptr(Row))
	return uintptr(r1)
}

func (m *TRasterImage) TransparentColor() TColor {
	r1 := LCL().SysCallN(4778, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TRasterImage) SetTransparentColor(AValue TColor) {
	LCL().SysCallN(4778, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) TransparentMode() TTransparentMode {
	r1 := LCL().SysCallN(4779, 0, m.Instance(), 0)
	return TTransparentMode(r1)
}

func (m *TRasterImage) SetTransparentMode(AValue TTransparentMode) {
	LCL().SysCallN(4779, 1, m.Instance(), uintptr(AValue))
}

func (m *TRasterImage) BitmapHandleAllocated() bool {
	r1 := LCL().SysCallN(4755, m.Instance())
	return GoBool(r1)
}

func (m *TRasterImage) MaskHandleAllocated() bool {
	r1 := LCL().SysCallN(4768, m.Instance())
	return GoBool(r1)
}

func (m *TRasterImage) PaletteAllocated() bool {
	r1 := LCL().SysCallN(4770, m.Instance())
	return GoBool(r1)
}

func (m *TRasterImage) ReleaseBitmapHandle() HBITMAP {
	r1 := LCL().SysCallN(4773, m.Instance())
	return HBITMAP(r1)
}

func (m *TRasterImage) ReleaseMaskHandle() HBITMAP {
	r1 := LCL().SysCallN(4774, m.Instance())
	return HBITMAP(r1)
}

func (m *TRasterImage) ReleasePalette() HPALETTE {
	r1 := LCL().SysCallN(4775, m.Instance())
	return HPALETTE(r1)
}

func (m *TRasterImage) HandleAllocated() bool {
	r1 := LCL().SysCallN(4761, m.Instance())
	return GoBool(r1)
}

func RasterImageClass() TClass {
	ret := LCL().SysCallN(4757)
	return TClass(ret)
}

func (m *TRasterImage) BeginUpdate(ACanvasOnly bool) {
	LCL().SysCallN(4753, m.Instance(), PascalBool(ACanvasOnly))
}

func (m *TRasterImage) EndUpdate(AStreamIsValid bool) {
	LCL().SysCallN(4758, m.Instance(), PascalBool(AStreamIsValid))
}

func (m *TRasterImage) FreeImage() {
	LCL().SysCallN(4759, m.Instance())
}

func (m *TRasterImage) LoadFromBitmapHandles(ABitmap, AMask HBITMAP, ARect *TRect) {
	LCL().SysCallN(4762, m.Instance(), uintptr(ABitmap), uintptr(AMask), uintptr(unsafePointer(ARect)))
}

func (m *TRasterImage) LoadFromDevice(DC HDC) {
	LCL().SysCallN(4763, m.Instance(), uintptr(DC))
}

func (m *TRasterImage) LoadFromStreamForStream(AStream IStream, ASize uint32) {
	LCL().SysCallN(4765, m.Instance(), GetObjectUintptr(AStream), uintptr(ASize))
}

func (m *TRasterImage) LoadFromRawImage(AIMage *TRawImage, ADataOwner bool) {
	LCL().SysCallN(4764, m.Instance(), uintptr(unsafePointer(AIMage)), PascalBool(ADataOwner))
}

func (m *TRasterImage) GetSize(OutWidth, OutHeight *int32) {
	var result0 uintptr
	var result1 uintptr
	LCL().SysCallN(4760, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*OutWidth = int32(result0)
	*OutHeight = int32(result1)
}

func (m *TRasterImage) Mask(ATransparentColor TColor) {
	LCL().SysCallN(4766, m.Instance(), uintptr(ATransparentColor))
}

func (m *TRasterImage) SetHandles(ABitmap, AMask HBITMAP) {
	LCL().SysCallN(4777, m.Instance(), uintptr(ABitmap), uintptr(AMask))
}
