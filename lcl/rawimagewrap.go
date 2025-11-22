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

// IRawImageWrap Parent: IObject
type IRawImageWrap interface {
	IObject
	RawImageData() types.PRawImage                                                                                // function
	GetLineStart(line uint32) types.PByte                                                                         // function
	ReadBits(position TRawImagePosition, prec byte, shift byte) uint16                                            // function
	IsMasked(testPixels bool) bool                                                                                // function
	IsTransparent(testPixels bool) bool                                                                           // function
	IsEqual(image IRawImageWrap) bool                                                                             // function
	Init()                                                                                                        // procedure
	CreateData(zeroMem bool)                                                                                      // procedure
	FreeData()                                                                                                    // procedure
	ReleaseData()                                                                                                 // procedure
	ExtractRect(rect types.TRect, outDst *IRawImageWrap)                                                          // procedure
	PerformEffect(drawEffect types.TGraphicsDrawEffect, createNewData bool, freeOldData bool)                     // procedure
	ReadChannels(position TRawImagePosition, outRed *uint16, outGreen *uint16, outBlue *uint16, outAlpha *uint16) // procedure
	ReadMask(position TRawImagePosition, outMask *bool)                                                           // procedure
	WriteBits(position TRawImagePosition, prec byte, shift byte, bits uint16)                                     // procedure
	WriteChannels(position TRawImagePosition, red uint16, green uint16, blue uint16, alpha uint16)                // procedure
	WriteMask(position TRawImagePosition, mask bool)                                                              // procedure
	Description() IRawImageDescriptionWrap                                                                        // property Description Getter
	SetDescription(value IRawImageDescriptionWrap)                                                                // property Description Setter
	Data() types.PByte                                                                                            // property Data Getter
	SetData(value types.PByte)                                                                                    // property Data Setter
	DataSize() uintptr                                                                                            // property DataSize Getter
	SetDataSize(value uintptr)                                                                                    // property DataSize Setter
	Mask() types.PByte                                                                                            // property Mask Getter
	SetMask(value types.PByte)                                                                                    // property Mask Setter
	MaskSize() uintptr                                                                                            // property MaskSize Getter
	SetMaskSize(value uintptr)                                                                                    // property MaskSize Setter
	Palette() types.PByte                                                                                         // property Palette Getter
	SetPalette(value types.PByte)                                                                                 // property Palette Setter
	PaletteSize() uintptr                                                                                         // property PaletteSize Getter
	SetPaletteSize(value uintptr)                                                                                 // property PaletteSize Setter
}

type TRawImageWrap struct {
	TObject
}

func (m *TRawImageWrap) RawImageData() types.PRawImage {
	if !m.IsValid() {
		return 0
	}
	r := rawImageWrapAPI().SysCallN(1, m.Instance())
	return types.PRawImage(r)
}

func (m *TRawImageWrap) GetLineStart(line uint32) types.PByte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageWrapAPI().SysCallN(2, m.Instance(), uintptr(line))
	return types.PByte(r)
}

func (m *TRawImageWrap) ReadBits(position TRawImagePosition, prec byte, shift byte) uint16 {
	if !m.IsValid() {
		return 0
	}
	r := rawImageWrapAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&position)), uintptr(prec), uintptr(shift))
	return uint16(r)
}

func (m *TRawImageWrap) IsMasked(testPixels bool) bool {
	if !m.IsValid() {
		return false
	}
	r := rawImageWrapAPI().SysCallN(4, m.Instance(), api.PasBool(testPixels))
	return api.GoBool(r)
}

func (m *TRawImageWrap) IsTransparent(testPixels bool) bool {
	if !m.IsValid() {
		return false
	}
	r := rawImageWrapAPI().SysCallN(5, m.Instance(), api.PasBool(testPixels))
	return api.GoBool(r)
}

func (m *TRawImageWrap) IsEqual(image IRawImageWrap) bool {
	if !m.IsValid() {
		return false
	}
	r := rawImageWrapAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(image))
	return api.GoBool(r)
}

func (m *TRawImageWrap) Init() {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(8, m.Instance())
}

func (m *TRawImageWrap) CreateData(zeroMem bool) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(9, m.Instance(), api.PasBool(zeroMem))
}

func (m *TRawImageWrap) FreeData() {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(10, m.Instance())
}

func (m *TRawImageWrap) ReleaseData() {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(11, m.Instance())
}

func (m *TRawImageWrap) ExtractRect(rect types.TRect, outDst *IRawImageWrap) {
	if !m.IsValid() {
		return
	}
	var dstPtr uintptr
	rawImageWrapAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&rect)), uintptr(base.UnsafePointer(&dstPtr)))
	*outDst = AsRawImageWrap(dstPtr)
}

func (m *TRawImageWrap) PerformEffect(drawEffect types.TGraphicsDrawEffect, createNewData bool, freeOldData bool) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(13, m.Instance(), uintptr(drawEffect), api.PasBool(createNewData), api.PasBool(freeOldData))
}

func (m *TRawImageWrap) ReadChannels(position TRawImagePosition, outRed *uint16, outGreen *uint16, outBlue *uint16, outAlpha *uint16) {
	if !m.IsValid() {
		return
	}
	var redPtr uintptr
	var greenPtr uintptr
	var bluePtr uintptr
	var alphaPtr uintptr
	rawImageWrapAPI().SysCallN(14, m.Instance(), uintptr(base.UnsafePointer(&position)), uintptr(base.UnsafePointer(&redPtr)), uintptr(base.UnsafePointer(&greenPtr)), uintptr(base.UnsafePointer(&bluePtr)), uintptr(base.UnsafePointer(&alphaPtr)))
	*outRed = uint16(redPtr)
	*outGreen = uint16(greenPtr)
	*outBlue = uint16(bluePtr)
	*outAlpha = uint16(alphaPtr)
}

func (m *TRawImageWrap) ReadMask(position TRawImagePosition, outMask *bool) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&position)), uintptr(base.UnsafePointer(outMask)))
}

func (m *TRawImageWrap) WriteBits(position TRawImagePosition, prec byte, shift byte, bits uint16) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(16, m.Instance(), uintptr(base.UnsafePointer(&position)), uintptr(prec), uintptr(shift), uintptr(bits))
}

func (m *TRawImageWrap) WriteChannels(position TRawImagePosition, red uint16, green uint16, blue uint16, alpha uint16) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&position)), uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha))
}

func (m *TRawImageWrap) WriteMask(position TRawImagePosition, mask bool) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(18, m.Instance(), uintptr(base.UnsafePointer(&position)), api.PasBool(mask))
}

func (m *TRawImageWrap) Description() IRawImageDescriptionWrap {
	if !m.IsValid() {
		return nil
	}
	r := rawImageWrapAPI().SysCallN(19, 0, m.Instance())
	return AsRawImageDescriptionWrap(r)
}

func (m *TRawImageWrap) SetDescription(value IRawImageDescriptionWrap) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(19, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TRawImageWrap) Data() types.PByte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageWrapAPI().SysCallN(20, 0, m.Instance())
	return types.PByte(r)
}

func (m *TRawImageWrap) SetData(value types.PByte) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageWrap) DataSize() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := rawImageWrapAPI().SysCallN(21, 0, m.Instance())
	return uintptr(r)
}

func (m *TRawImageWrap) SetDataSize(value uintptr) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageWrap) Mask() types.PByte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageWrapAPI().SysCallN(22, 0, m.Instance())
	return types.PByte(r)
}

func (m *TRawImageWrap) SetMask(value types.PByte) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageWrap) MaskSize() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := rawImageWrapAPI().SysCallN(23, 0, m.Instance())
	return uintptr(r)
}

func (m *TRawImageWrap) SetMaskSize(value uintptr) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageWrap) Palette() types.PByte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageWrapAPI().SysCallN(24, 0, m.Instance())
	return types.PByte(r)
}

func (m *TRawImageWrap) SetPalette(value types.PByte) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageWrap) PaletteSize() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := rawImageWrapAPI().SysCallN(25, 0, m.Instance())
	return uintptr(r)
}

func (m *TRawImageWrap) SetPaletteSize(value uintptr) {
	if !m.IsValid() {
		return
	}
	rawImageWrapAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

// RawImageWrap  is static instance
var RawImageWrap _RawImageWrapClass

// _RawImageWrapClass is class type defined by TRawImageWrap
type _RawImageWrapClass uintptr

func (_RawImageWrapClass) UnWrap(data types.PRawImage) IRawImageWrap {
	r := rawImageWrapAPI().SysCallN(7, uintptr(data))
	return AsRawImageWrap(r)
}

// NewRawImageWrap class constructor
func NewRawImageWrap() IRawImageWrap {
	r := rawImageWrapAPI().SysCallN(0)
	return AsRawImageWrap(r)
}

var (
	rawImageWrapOnce   base.Once
	rawImageWrapImport *imports.Imports = nil
)

func rawImageWrapAPI() *imports.Imports {
	rawImageWrapOnce.Do(func() {
		rawImageWrapImport = api.NewDefaultImports()
		rawImageWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRawImageWrap_Create", 0), // constructor NewRawImageWrap
			/* 1 */ imports.NewTable("TRawImageWrap_RawImageData", 0), // function RawImageData
			/* 2 */ imports.NewTable("TRawImageWrap_GetLineStart", 0), // function GetLineStart
			/* 3 */ imports.NewTable("TRawImageWrap_ReadBits", 0), // function ReadBits
			/* 4 */ imports.NewTable("TRawImageWrap_IsMasked", 0), // function IsMasked
			/* 5 */ imports.NewTable("TRawImageWrap_IsTransparent", 0), // function IsTransparent
			/* 6 */ imports.NewTable("TRawImageWrap_IsEqual", 0), // function IsEqual
			/* 7 */ imports.NewTable("TRawImageWrap_UnWrap", 0), // static function UnWrap
			/* 8 */ imports.NewTable("TRawImageWrap_Init", 0), // procedure Init
			/* 9 */ imports.NewTable("TRawImageWrap_CreateData", 0), // procedure CreateData
			/* 10 */ imports.NewTable("TRawImageWrap_FreeData", 0), // procedure FreeData
			/* 11 */ imports.NewTable("TRawImageWrap_ReleaseData", 0), // procedure ReleaseData
			/* 12 */ imports.NewTable("TRawImageWrap_ExtractRect", 0), // procedure ExtractRect
			/* 13 */ imports.NewTable("TRawImageWrap_PerformEffect", 0), // procedure PerformEffect
			/* 14 */ imports.NewTable("TRawImageWrap_ReadChannels", 0), // procedure ReadChannels
			/* 15 */ imports.NewTable("TRawImageWrap_ReadMask", 0), // procedure ReadMask
			/* 16 */ imports.NewTable("TRawImageWrap_WriteBits", 0), // procedure WriteBits
			/* 17 */ imports.NewTable("TRawImageWrap_WriteChannels", 0), // procedure WriteChannels
			/* 18 */ imports.NewTable("TRawImageWrap_WriteMask", 0), // procedure WriteMask
			/* 19 */ imports.NewTable("TRawImageWrap_Description", 0), // property Description
			/* 20 */ imports.NewTable("TRawImageWrap_Data", 0), // property Data
			/* 21 */ imports.NewTable("TRawImageWrap_DataSize", 0), // property DataSize
			/* 22 */ imports.NewTable("TRawImageWrap_Mask", 0), // property Mask
			/* 23 */ imports.NewTable("TRawImageWrap_MaskSize", 0), // property MaskSize
			/* 24 */ imports.NewTable("TRawImageWrap_Palette", 0), // property Palette
			/* 25 */ imports.NewTable("TRawImageWrap_PaletteSize", 0), // property PaletteSize
		}
	})
	return rawImageWrapImport
}
