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

// ILazIntfImage Parent: IFPCustomImage
type ILazIntfImage interface {
	IFPCustomImage
	CheckDescription(description IRawImageDescriptionWrap, exceptionOnError bool) bool              // function
	GetDataLineStart(Y int32) uintptr                                                               // function
	HasTransparency() bool                                                                          // function
	HasMask() bool                                                                                  // function
	BeginUpdate()                                                                                   // procedure
	EndUpdate()                                                                                     // procedure
	LoadFromDevice(dC types.HDC)                                                                    // procedure
	LoadFromBitmap(bitmap types.HBitmap, maskBitmap types.HBitmap, width int32, height int32)       // procedure
	CreateBitmaps(outBitmap *types.HBitmap, outMask *types.HBitmap, skipMask bool)                  // procedure
	SetRawImage(rawImage IRawImageWrap, dataOwner bool)                                             // procedure
	GetRawImage(outRawImage *IRawImageWrap, transferOwnership bool)                                 // procedure
	FillPixels(color TFPColor)                                                                      // procedure
	CopyPixels(source IFPCustomImage, xDst int32, yDst int32, alphaMask bool, alphaTreshold uint16) // procedure
	AlphaBlend(source ILazIntfImage, sourceAlpha ILazIntfImage, destX int32, destY int32)           // procedure
	AlphaFromMask(keepAlpha bool)                                                                   // procedure
	Mask(color TFPColor, keepOldMask bool)                                                          // procedure
	GetXYDataPosition(X int32, Y int32, outPosition *TRawImagePosition)                             // procedure
	GetXYMaskPosition(X int32, Y int32, outPosition *TRawImagePosition)                             // procedure
	CreateData()                                                                                    // procedure
	SetDataDescriptionKeepData(description IRawImageDescriptionWrap)                                // procedure
	PixelData() types.PByte                                                                         // property PixelData Getter
	MaskData() types.PByte                                                                          // property MaskData Getter
	DataDescription() IRawImageDescriptionWrap                                                      // property DataDescription Getter
	SetDataDescription(value IRawImageDescriptionWrap)                                              // property DataDescription Setter
	TColors(X int32, Y int32) types.TGraphicsColor                                                  // property TColors Getter
	SetTColors(X int32, Y int32, value types.TGraphicsColor)                                        // property TColors Setter
	Masked(X int32, Y int32) bool                                                                   // property Masked Getter
	SetMasked(X int32, Y int32, value bool)                                                         // property Masked Setter
}

type TLazIntfImage struct {
	TFPCustomImage
}

func (m *TLazIntfImage) CheckDescription(description IRawImageDescriptionWrap, exceptionOnError bool) bool {
	if !m.IsValid() {
		return false
	}
	r := lazIntfImageAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(description), api.PasBool(exceptionOnError))
	return api.GoBool(r)
}

func (m *TLazIntfImage) GetDataLineStart(Y int32) uintptr {
	if !m.IsValid() {
		return 0
	}
	r := lazIntfImageAPI().SysCallN(4, m.Instance(), uintptr(Y))
	return uintptr(r)
}

func (m *TLazIntfImage) HasTransparency() bool {
	if !m.IsValid() {
		return false
	}
	r := lazIntfImageAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TLazIntfImage) HasMask() bool {
	if !m.IsValid() {
		return false
	}
	r := lazIntfImageAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TLazIntfImage) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(7, m.Instance())
}

func (m *TLazIntfImage) EndUpdate() {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(8, m.Instance())
}

func (m *TLazIntfImage) LoadFromDevice(dC types.HDC) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(9, m.Instance(), uintptr(dC))
}

func (m *TLazIntfImage) LoadFromBitmap(bitmap types.HBitmap, maskBitmap types.HBitmap, width int32, height int32) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(10, m.Instance(), uintptr(bitmap), uintptr(maskBitmap), uintptr(width), uintptr(height))
}

func (m *TLazIntfImage) CreateBitmaps(outBitmap *types.HBitmap, outMask *types.HBitmap, skipMask bool) {
	if !m.IsValid() {
		return
	}
	var bitmapPtr uintptr
	var maskPtr uintptr
	lazIntfImageAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(&bitmapPtr)), uintptr(base.UnsafePointer(&maskPtr)), api.PasBool(skipMask))
	*outBitmap = types.HBitmap(bitmapPtr)
	*outMask = types.HBitmap(maskPtr)
}

func (m *TLazIntfImage) SetRawImage(rawImage IRawImageWrap, dataOwner bool) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(rawImage), api.PasBool(dataOwner))
}

func (m *TLazIntfImage) GetRawImage(outRawImage *IRawImageWrap, transferOwnership bool) {
	if !m.IsValid() {
		return
	}
	var rawImagePtr uintptr
	lazIntfImageAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&rawImagePtr)), api.PasBool(transferOwnership))
	*outRawImage = AsRawImageWrap(rawImagePtr)
}

func (m *TLazIntfImage) FillPixels(color TFPColor) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&color)))
}

func (m *TLazIntfImage) CopyPixels(source IFPCustomImage, xDst int32, yDst int32, alphaMask bool, alphaTreshold uint16) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(source), uintptr(xDst), uintptr(yDst), api.PasBool(alphaMask), uintptr(alphaTreshold))
}

func (m *TLazIntfImage) AlphaBlend(source ILazIntfImage, sourceAlpha ILazIntfImage, destX int32, destY int32) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(source), base.GetObjectUintptr(sourceAlpha), uintptr(destX), uintptr(destY))
}

func (m *TLazIntfImage) AlphaFromMask(keepAlpha bool) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(17, m.Instance(), api.PasBool(keepAlpha))
}

func (m *TLazIntfImage) Mask(color TFPColor, keepOldMask bool) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&color)), api.PasBool(keepOldMask))
}

func (m *TLazIntfImage) GetXYDataPosition(X int32, Y int32, outPosition *TRawImagePosition) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(19, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(outPosition)))
}

func (m *TLazIntfImage) GetXYMaskPosition(X int32, Y int32, outPosition *TRawImagePosition) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(20, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(outPosition)))
}

func (m *TLazIntfImage) CreateData() {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(21, m.Instance())
}

func (m *TLazIntfImage) SetDataDescriptionKeepData(description IRawImageDescriptionWrap) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(22, m.Instance(), base.GetObjectUintptr(description))
}

func (m *TLazIntfImage) PixelData() types.PByte {
	if !m.IsValid() {
		return 0
	}
	r := lazIntfImageAPI().SysCallN(23, m.Instance())
	return types.PByte(r)
}

func (m *TLazIntfImage) MaskData() types.PByte {
	if !m.IsValid() {
		return 0
	}
	r := lazIntfImageAPI().SysCallN(24, m.Instance())
	return types.PByte(r)
}

func (m *TLazIntfImage) DataDescription() IRawImageDescriptionWrap {
	if !m.IsValid() {
		return nil
	}
	r := lazIntfImageAPI().SysCallN(25, 0, m.Instance())
	return AsRawImageDescriptionWrap(r)
}

func (m *TLazIntfImage) SetDataDescription(value IRawImageDescriptionWrap) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(25, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazIntfImage) TColors(X int32, Y int32) types.TGraphicsColor {
	if !m.IsValid() {
		return 0
	}
	r := lazIntfImageAPI().SysCallN(26, 0, m.Instance(), uintptr(X), uintptr(Y))
	return types.TGraphicsColor(r)
}

func (m *TLazIntfImage) SetTColors(X int32, Y int32, value types.TGraphicsColor) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(26, 1, m.Instance(), uintptr(X), uintptr(Y), uintptr(value))
}

func (m *TLazIntfImage) Masked(X int32, Y int32) bool {
	if !m.IsValid() {
		return false
	}
	r := lazIntfImageAPI().SysCallN(27, 0, m.Instance(), uintptr(X), uintptr(Y))
	return api.GoBool(r)
}

func (m *TLazIntfImage) SetMasked(X int32, Y int32, value bool) {
	if !m.IsValid() {
		return
	}
	lazIntfImageAPI().SysCallN(27, 1, m.Instance(), uintptr(X), uintptr(Y), api.PasBool(value))
}

// NewLazIntfImage class constructor
func NewLazIntfImage(width int32, height int32) ILazIntfImage {
	r := lazIntfImageAPI().SysCallN(0, uintptr(width), uintptr(height))
	return AsLazIntfImage(r)
}

// NewLazIntfImageWithIntX2RawImageQueryFlags class constructor
func NewLazIntfImageWithIntX2RawImageQueryFlags(width int32, height int32, flags types.TRawImageQueryFlags) ILazIntfImage {
	r := lazIntfImageAPI().SysCallN(1, uintptr(width), uintptr(height), uintptr(flags))
	return AsLazIntfImage(r)
}

// NewLazIntfImageCompatible class constructor
func NewLazIntfImageCompatible(intfImg ILazIntfImage, width int32, height int32) ILazIntfImage {
	r := lazIntfImageAPI().SysCallN(2, base.GetObjectUintptr(intfImg), uintptr(width), uintptr(height))
	return AsLazIntfImage(r)
}

var (
	lazIntfImageOnce   base.Once
	lazIntfImageImport *imports.Imports = nil
)

func lazIntfImageAPI() *imports.Imports {
	lazIntfImageOnce.Do(func() {
		lazIntfImageImport = api.NewDefaultImports()
		lazIntfImageImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazIntfImage_Create", 0), // constructor NewLazIntfImage
			/* 1 */ imports.NewTable("TLazIntfImage_CreateWithIntX2RawImageQueryFlags", 0), // constructor NewLazIntfImageWithIntX2RawImageQueryFlags
			/* 2 */ imports.NewTable("TLazIntfImage_CreateCompatible", 0), // constructor NewLazIntfImageCompatible
			/* 3 */ imports.NewTable("TLazIntfImage_CheckDescription", 0), // function CheckDescription
			/* 4 */ imports.NewTable("TLazIntfImage_GetDataLineStart", 0), // function GetDataLineStart
			/* 5 */ imports.NewTable("TLazIntfImage_HasTransparency", 0), // function HasTransparency
			/* 6 */ imports.NewTable("TLazIntfImage_HasMask", 0), // function HasMask
			/* 7 */ imports.NewTable("TLazIntfImage_BeginUpdate", 0), // procedure BeginUpdate
			/* 8 */ imports.NewTable("TLazIntfImage_EndUpdate", 0), // procedure EndUpdate
			/* 9 */ imports.NewTable("TLazIntfImage_LoadFromDevice", 0), // procedure LoadFromDevice
			/* 10 */ imports.NewTable("TLazIntfImage_LoadFromBitmap", 0), // procedure LoadFromBitmap
			/* 11 */ imports.NewTable("TLazIntfImage_CreateBitmaps", 0), // procedure CreateBitmaps
			/* 12 */ imports.NewTable("TLazIntfImage_SetRawImage", 0), // procedure SetRawImage
			/* 13 */ imports.NewTable("TLazIntfImage_GetRawImage", 0), // procedure GetRawImage
			/* 14 */ imports.NewTable("TLazIntfImage_FillPixels", 0), // procedure FillPixels
			/* 15 */ imports.NewTable("TLazIntfImage_CopyPixels", 0), // procedure CopyPixels
			/* 16 */ imports.NewTable("TLazIntfImage_AlphaBlend", 0), // procedure AlphaBlend
			/* 17 */ imports.NewTable("TLazIntfImage_AlphaFromMask", 0), // procedure AlphaFromMask
			/* 18 */ imports.NewTable("TLazIntfImage_Mask", 0), // procedure Mask
			/* 19 */ imports.NewTable("TLazIntfImage_GetXYDataPosition", 0), // procedure GetXYDataPosition
			/* 20 */ imports.NewTable("TLazIntfImage_GetXYMaskPosition", 0), // procedure GetXYMaskPosition
			/* 21 */ imports.NewTable("TLazIntfImage_CreateData", 0), // procedure CreateData
			/* 22 */ imports.NewTable("TLazIntfImage_SetDataDescriptionKeepData", 0), // procedure SetDataDescriptionKeepData
			/* 23 */ imports.NewTable("TLazIntfImage_PixelData", 0), // property PixelData
			/* 24 */ imports.NewTable("TLazIntfImage_MaskData", 0), // property MaskData
			/* 25 */ imports.NewTable("TLazIntfImage_DataDescription", 0), // property DataDescription
			/* 26 */ imports.NewTable("TLazIntfImage_TColors", 0), // property TColors
			/* 27 */ imports.NewTable("TLazIntfImage_Masked", 0), // property Masked
		}
	})
	return lazIntfImageImport
}
