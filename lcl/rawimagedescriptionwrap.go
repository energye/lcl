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

// IRawImageDescriptionWrap Parent: IObject
type IRawImageDescriptionWrap interface {
	IObject
	RawImageDescriptionData() types.PRawImageDescription // function
	GetDescriptionFromMask() IRawImageDescriptionWrap    // function
	GetDescriptionFromAlpha() IRawImageDescriptionWrap   // function
	BytesPerLine() uintptr                               // function
	BitsPerLine() uintptr                                // function
	MaskBytesPerLine() uintptr                           // function
	MaskBitsPerLine() uintptr                            // function
	AsString() string                                    // function
	IsEqual(desc IRawImageDescriptionWrap) bool          // function
	// Init
	//  don't use a constructor here, it will break compatibility with a record
	Init() // procedure
	// InitBPP1
	//  1-bit mono format
	InitBPP1(width int32, height int32) // procedure
	// InitBPP16R5G6B5
	//  16-bits formats
	InitBPP16R5G6B5(width int32, height int32) // procedure
	// InitBPP24R8G8B8BIOTTB
	//  Formats in RGB order
	InitBPP24R8G8B8BIOTTB(width int32, height int32)           // procedure
	InitBPP24R8G8B8BIOTTBUpsideDown(width int32, height int32) // procedure
	InitBPP32A8R8G8B8BIOTTB(width int32, height int32)         // procedure
	InitBPP32R8G8B8A8BIOTTB(width int32, height int32)         // procedure
	// InitBPP24B8G8R8BIOTTB
	//  Formats in Windows pixels order: BGR
	InitBPP24B8G8R8BIOTTB(width int32, height int32)     // procedure
	InitBPP24B8G8R8M1BIOTTB(width int32, height int32)   // procedure
	InitBPP32B8G8R8BIOTTB(width int32, height int32)     // procedure
	InitBPP32B8G8R8M1BIOTTB(width int32, height int32)   // procedure
	InitBPP32B8G8R8A8BIOTTB(width int32, height int32)   // procedure
	InitBPP32B8G8R8A8M1BIOTTB(width int32, height int32) // procedure
	// GetRGBIndices
	//  returns indices of channels in four-element array
	GetRGBIndices(outRidx *byte, outGidx *byte, outBidx *byte, outAidx *byte) // procedure
	Format() types.TRawImageColorFormat                                       // property Format Getter
	SetFormat(value types.TRawImageColorFormat)                               // property Format Setter
	Width() uint32                                                            // property Width Getter
	SetWidth(value uint32)                                                    // property Width Setter
	Height() uint32                                                           // property Height Getter
	SetHeight(value uint32)                                                   // property Height Setter
	// Depth
	//  used bits per pixel
	Depth() byte                                 // property Depth Getter
	SetDepth(value byte)                         // property Depth Setter
	BitOrder() types.TRawImageBitOrder           // property BitOrder Getter
	SetBitOrder(value types.TRawImageBitOrder)   // property BitOrder Setter
	ByteOrder() types.TRawImageByteOrder         // property ByteOrder Getter
	SetByteOrder(value types.TRawImageByteOrder) // property ByteOrder Setter
	LineOrder() types.TRawImageLineOrder         // property LineOrder Getter
	SetLineOrder(value types.TRawImageLineOrder) // property LineOrder Setter
	LineEnd() types.TRawImageLineEnd             // property LineEnd Getter
	SetLineEnd(value types.TRawImageLineEnd)     // property LineEnd Setter
	// BitsPerPixel
	//  bits per pixel. can be greater than Depth.
	BitsPerPixel() byte         // property BitsPerPixel Getter
	SetBitsPerPixel(value byte) // property BitsPerPixel Setter
	// RedPrec
	//  red or gray precision. bits for red
	RedPrec() byte         // property RedPrec Getter
	SetRedPrec(value byte) // property RedPrec Setter
	// RedShift
	//  bitshift. Direction: from least to most significant
	RedShift() byte           // property RedShift Getter
	SetRedShift(value byte)   // property RedShift Setter
	GreenPrec() byte          // property GreenPrec Getter
	SetGreenPrec(value byte)  // property GreenPrec Setter
	GreenShift() byte         // property GreenShift Getter
	SetGreenShift(value byte) // property GreenShift Setter
	BluePrec() byte           // property BluePrec Getter
	SetBluePrec(value byte)   // property BluePrec Setter
	BlueShift() byte          // property BlueShift Getter
	SetBlueShift(value byte)  // property BlueShift Setter
	AlphaPrec() byte          // property AlphaPrec Getter
	SetAlphaPrec(value byte)  // property AlphaPrec Setter
	AlphaShift() byte         // property AlphaShift Getter
	SetAlphaShift(value byte) // property AlphaShift Setter
	// MaskBitsPerPixel
	//  The next values are only valid, if there is a mask (MaskBitsPerPixel > 0)
	//  Masks are always separate with a depth of 1 bpp. One pixel can occupy
	//  one byte at most
	//  a value of 1 means that pixel is masked
	//  a value of 0 means the pixel value is shown
	//  bits per mask pixel, usually 1, 0 when no mask
	MaskBitsPerPixel() byte         // property MaskBitsPerPixel Getter
	SetMaskBitsPerPixel(value byte) // property MaskBitsPerPixel Setter
	// MaskShift
	//  the shift (=position) of the mask bit
	MaskShift() byte                               // property MaskShift Getter
	SetMaskShift(value byte)                       // property MaskShift Setter
	MaskLineEnd() types.TRawImageLineEnd           // property MaskLineEnd Getter
	SetMaskLineEnd(value types.TRawImageLineEnd)   // property MaskLineEnd Setter
	MaskBitOrder() types.TRawImageBitOrder         // property MaskBitOrder Getter
	SetMaskBitOrder(value types.TRawImageBitOrder) // property MaskBitOrder Setter
	// PaletteColorCount
	//  The next values are only valid, if there is a palette (PaletteColorCount > 0)
	//  entries in color palette. 0 when no palette.
	PaletteColorCount() uint16         // property PaletteColorCount Getter
	SetPaletteColorCount(value uint16) // property PaletteColorCount Setter
	// PaletteBitsPerIndex
	//  bits per palette index, this can be larger than the colors used
	PaletteBitsPerIndex() byte         // property PaletteBitsPerIndex Getter
	SetPaletteBitsPerIndex(value byte) // property PaletteBitsPerIndex Setter
	// PaletteShift
	//  bitshift. Direction: from least to most significant
	PaletteShift() byte                                 // property PaletteShift Getter
	SetPaletteShift(value byte)                         // property PaletteShift Setter
	PaletteLineEnd() types.TRawImageLineEnd             // property PaletteLineEnd Getter
	SetPaletteLineEnd(value types.TRawImageLineEnd)     // property PaletteLineEnd Setter
	PaletteBitOrder() types.TRawImageBitOrder           // property PaletteBitOrder Getter
	SetPaletteBitOrder(value types.TRawImageBitOrder)   // property PaletteBitOrder Setter
	PaletteByteOrder() types.TRawImageByteOrder         // property PaletteByteOrder Getter
	SetPaletteByteOrder(value types.TRawImageByteOrder) // property PaletteByteOrder Setter
}

type TRawImageDescriptionWrap struct {
	TObject
}

func (m *TRawImageDescriptionWrap) RawImageDescriptionData() types.PRawImageDescription {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(1, m.Instance())
	return types.PRawImageDescription(r)
}

func (m *TRawImageDescriptionWrap) GetDescriptionFromMask() IRawImageDescriptionWrap {
	if !m.IsValid() {
		return nil
	}
	r := rawImageDescriptionWrapAPI().SysCallN(2, m.Instance())
	return AsRawImageDescriptionWrap(r)
}

func (m *TRawImageDescriptionWrap) GetDescriptionFromAlpha() IRawImageDescriptionWrap {
	if !m.IsValid() {
		return nil
	}
	r := rawImageDescriptionWrapAPI().SysCallN(3, m.Instance())
	return AsRawImageDescriptionWrap(r)
}

func (m *TRawImageDescriptionWrap) BytesPerLine() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(4, m.Instance())
	return uintptr(r)
}

func (m *TRawImageDescriptionWrap) BitsPerLine() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(5, m.Instance())
	return uintptr(r)
}

func (m *TRawImageDescriptionWrap) MaskBytesPerLine() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(6, m.Instance())
	return uintptr(r)
}

func (m *TRawImageDescriptionWrap) MaskBitsPerLine() uintptr {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(7, m.Instance())
	return uintptr(r)
}

func (m *TRawImageDescriptionWrap) AsString() string {
	if !m.IsValid() {
		return ""
	}
	r := rawImageDescriptionWrapAPI().SysCallN(8, m.Instance())
	return api.GoStr(r)
}

func (m *TRawImageDescriptionWrap) IsEqual(desc IRawImageDescriptionWrap) bool {
	if !m.IsValid() {
		return false
	}
	r := rawImageDescriptionWrapAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(desc))
	return api.GoBool(r)
}

func (m *TRawImageDescriptionWrap) Init() {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(11, m.Instance())
}

func (m *TRawImageDescriptionWrap) InitBPP1(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(12, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP16R5G6B5(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(13, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP24R8G8B8BIOTTB(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(14, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP24R8G8B8BIOTTBUpsideDown(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(15, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP32A8R8G8B8BIOTTB(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(16, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP32R8G8B8A8BIOTTB(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(17, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP24B8G8R8BIOTTB(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(18, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP24B8G8R8M1BIOTTB(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(19, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP32B8G8R8BIOTTB(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(20, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP32B8G8R8M1BIOTTB(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(21, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP32B8G8R8A8BIOTTB(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(22, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) InitBPP32B8G8R8A8M1BIOTTB(width int32, height int32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(23, m.Instance(), uintptr(width), uintptr(height))
}

func (m *TRawImageDescriptionWrap) GetRGBIndices(outRidx *byte, outGidx *byte, outBidx *byte, outAidx *byte) {
	if !m.IsValid() {
		return
	}
	var ridxPtr uintptr
	var gidxPtr uintptr
	var bidxPtr uintptr
	var aidxPtr uintptr
	rawImageDescriptionWrapAPI().SysCallN(24, m.Instance(), uintptr(base.UnsafePointer(&ridxPtr)), uintptr(base.UnsafePointer(&gidxPtr)), uintptr(base.UnsafePointer(&bidxPtr)), uintptr(base.UnsafePointer(&aidxPtr)))
	*outRidx = byte(ridxPtr)
	*outGidx = byte(gidxPtr)
	*outBidx = byte(bidxPtr)
	*outAidx = byte(aidxPtr)
}

func (m *TRawImageDescriptionWrap) Format() types.TRawImageColorFormat {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(25, 0, m.Instance())
	return types.TRawImageColorFormat(r)
}

func (m *TRawImageDescriptionWrap) SetFormat(value types.TRawImageColorFormat) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) Width() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(26, 0, m.Instance())
	return uint32(r)
}

func (m *TRawImageDescriptionWrap) SetWidth(value uint32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) Height() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(27, 0, m.Instance())
	return uint32(r)
}

func (m *TRawImageDescriptionWrap) SetHeight(value uint32) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) Depth() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(28, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetDepth(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) BitOrder() types.TRawImageBitOrder {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(29, 0, m.Instance())
	return types.TRawImageBitOrder(r)
}

func (m *TRawImageDescriptionWrap) SetBitOrder(value types.TRawImageBitOrder) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) ByteOrder() types.TRawImageByteOrder {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(30, 0, m.Instance())
	return types.TRawImageByteOrder(r)
}

func (m *TRawImageDescriptionWrap) SetByteOrder(value types.TRawImageByteOrder) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) LineOrder() types.TRawImageLineOrder {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(31, 0, m.Instance())
	return types.TRawImageLineOrder(r)
}

func (m *TRawImageDescriptionWrap) SetLineOrder(value types.TRawImageLineOrder) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) LineEnd() types.TRawImageLineEnd {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(32, 0, m.Instance())
	return types.TRawImageLineEnd(r)
}

func (m *TRawImageDescriptionWrap) SetLineEnd(value types.TRawImageLineEnd) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(32, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) BitsPerPixel() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(33, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetBitsPerPixel(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(33, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) RedPrec() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(34, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetRedPrec(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(34, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) RedShift() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(35, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetRedShift(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(35, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) GreenPrec() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(36, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetGreenPrec(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) GreenShift() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(37, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetGreenShift(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(37, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) BluePrec() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(38, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetBluePrec(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(38, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) BlueShift() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(39, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetBlueShift(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(39, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) AlphaPrec() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(40, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetAlphaPrec(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(40, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) AlphaShift() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(41, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetAlphaShift(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(41, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) MaskBitsPerPixel() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(42, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetMaskBitsPerPixel(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(42, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) MaskShift() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(43, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetMaskShift(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(43, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) MaskLineEnd() types.TRawImageLineEnd {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(44, 0, m.Instance())
	return types.TRawImageLineEnd(r)
}

func (m *TRawImageDescriptionWrap) SetMaskLineEnd(value types.TRawImageLineEnd) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(44, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) MaskBitOrder() types.TRawImageBitOrder {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(45, 0, m.Instance())
	return types.TRawImageBitOrder(r)
}

func (m *TRawImageDescriptionWrap) SetMaskBitOrder(value types.TRawImageBitOrder) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(45, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) PaletteColorCount() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(46, 0, m.Instance())
	return uint16(r)
}

func (m *TRawImageDescriptionWrap) SetPaletteColorCount(value uint16) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(46, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) PaletteBitsPerIndex() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(47, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetPaletteBitsPerIndex(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(47, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) PaletteShift() byte {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(48, 0, m.Instance())
	return byte(r)
}

func (m *TRawImageDescriptionWrap) SetPaletteShift(value byte) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(48, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) PaletteLineEnd() types.TRawImageLineEnd {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(49, 0, m.Instance())
	return types.TRawImageLineEnd(r)
}

func (m *TRawImageDescriptionWrap) SetPaletteLineEnd(value types.TRawImageLineEnd) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(49, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) PaletteBitOrder() types.TRawImageBitOrder {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(50, 0, m.Instance())
	return types.TRawImageBitOrder(r)
}

func (m *TRawImageDescriptionWrap) SetPaletteBitOrder(value types.TRawImageBitOrder) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(50, 1, m.Instance(), uintptr(value))
}

func (m *TRawImageDescriptionWrap) PaletteByteOrder() types.TRawImageByteOrder {
	if !m.IsValid() {
		return 0
	}
	r := rawImageDescriptionWrapAPI().SysCallN(51, 0, m.Instance())
	return types.TRawImageByteOrder(r)
}

func (m *TRawImageDescriptionWrap) SetPaletteByteOrder(value types.TRawImageByteOrder) {
	if !m.IsValid() {
		return
	}
	rawImageDescriptionWrapAPI().SysCallN(51, 1, m.Instance(), uintptr(value))
}

// RawImageDescriptionWrap  is static instance
var RawImageDescriptionWrap _RawImageDescriptionWrapClass

// _RawImageDescriptionWrapClass is class type defined by TRawImageDescriptionWrap
type _RawImageDescriptionWrapClass uintptr

func (_RawImageDescriptionWrapClass) UnWrap(data types.PRawImageDescription) IRawImageDescriptionWrap {
	r := rawImageDescriptionWrapAPI().SysCallN(10, uintptr(data))
	return AsRawImageDescriptionWrap(r)
}

// NewRawImageDescriptionWrap class constructor
func NewRawImageDescriptionWrap() IRawImageDescriptionWrap {
	r := rawImageDescriptionWrapAPI().SysCallN(0)
	return AsRawImageDescriptionWrap(r)
}

var (
	rawImageDescriptionWrapOnce   base.Once
	rawImageDescriptionWrapImport *imports.Imports = nil
)

func rawImageDescriptionWrapAPI() *imports.Imports {
	rawImageDescriptionWrapOnce.Do(func() {
		rawImageDescriptionWrapImport = api.NewDefaultImports()
		rawImageDescriptionWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TRawImageDescriptionWrap_Create", 0), // constructor NewRawImageDescriptionWrap
			/* 1 */ imports.NewTable("TRawImageDescriptionWrap_RawImageDescriptionData", 0), // function RawImageDescriptionData
			/* 2 */ imports.NewTable("TRawImageDescriptionWrap_GetDescriptionFromMask", 0), // function GetDescriptionFromMask
			/* 3 */ imports.NewTable("TRawImageDescriptionWrap_GetDescriptionFromAlpha", 0), // function GetDescriptionFromAlpha
			/* 4 */ imports.NewTable("TRawImageDescriptionWrap_BytesPerLine", 0), // function BytesPerLine
			/* 5 */ imports.NewTable("TRawImageDescriptionWrap_BitsPerLine", 0), // function BitsPerLine
			/* 6 */ imports.NewTable("TRawImageDescriptionWrap_MaskBytesPerLine", 0), // function MaskBytesPerLine
			/* 7 */ imports.NewTable("TRawImageDescriptionWrap_MaskBitsPerLine", 0), // function MaskBitsPerLine
			/* 8 */ imports.NewTable("TRawImageDescriptionWrap_AsString", 0), // function AsString
			/* 9 */ imports.NewTable("TRawImageDescriptionWrap_IsEqual", 0), // function IsEqual
			/* 10 */ imports.NewTable("TRawImageDescriptionWrap_UnWrap", 0), // static function UnWrap
			/* 11 */ imports.NewTable("TRawImageDescriptionWrap_Init", 0), // procedure Init
			/* 12 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP1", 0), // procedure InitBPP1
			/* 13 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP16_R5G6B5", 0), // procedure InitBPP16R5G6B5
			/* 14 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP24_R8G8B8_BIO_TTB", 0), // procedure InitBPP24R8G8B8BIOTTB
			/* 15 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP24_R8G8B8_BIO_TTB_UpsideDown", 0), // procedure InitBPP24R8G8B8BIOTTBUpsideDown
			/* 16 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP32_A8R8G8B8_BIO_TTB", 0), // procedure InitBPP32A8R8G8B8BIOTTB
			/* 17 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP32_R8G8B8A8_BIO_TTB", 0), // procedure InitBPP32R8G8B8A8BIOTTB
			/* 18 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP24_B8G8R8_BIO_TTB", 0), // procedure InitBPP24B8G8R8BIOTTB
			/* 19 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP24_B8G8R8_M1_BIO_TTB", 0), // procedure InitBPP24B8G8R8M1BIOTTB
			/* 20 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP32_B8G8R8_BIO_TTB", 0), // procedure InitBPP32B8G8R8BIOTTB
			/* 21 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP32_B8G8R8_M1_BIO_TTB", 0), // procedure InitBPP32B8G8R8M1BIOTTB
			/* 22 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP32_B8G8R8A8_BIO_TTB", 0), // procedure InitBPP32B8G8R8A8BIOTTB
			/* 23 */ imports.NewTable("TRawImageDescriptionWrap_Init_BPP32_B8G8R8A8_M1_BIO_TTB", 0), // procedure InitBPP32B8G8R8A8M1BIOTTB
			/* 24 */ imports.NewTable("TRawImageDescriptionWrap_GetRGBIndices", 0), // procedure GetRGBIndices
			/* 25 */ imports.NewTable("TRawImageDescriptionWrap_Format", 0), // property Format
			/* 26 */ imports.NewTable("TRawImageDescriptionWrap_Width", 0), // property Width
			/* 27 */ imports.NewTable("TRawImageDescriptionWrap_Height", 0), // property Height
			/* 28 */ imports.NewTable("TRawImageDescriptionWrap_Depth", 0), // property Depth
			/* 29 */ imports.NewTable("TRawImageDescriptionWrap_BitOrder", 0), // property BitOrder
			/* 30 */ imports.NewTable("TRawImageDescriptionWrap_ByteOrder", 0), // property ByteOrder
			/* 31 */ imports.NewTable("TRawImageDescriptionWrap_LineOrder", 0), // property LineOrder
			/* 32 */ imports.NewTable("TRawImageDescriptionWrap_LineEnd", 0), // property LineEnd
			/* 33 */ imports.NewTable("TRawImageDescriptionWrap_BitsPerPixel", 0), // property BitsPerPixel
			/* 34 */ imports.NewTable("TRawImageDescriptionWrap_RedPrec", 0), // property RedPrec
			/* 35 */ imports.NewTable("TRawImageDescriptionWrap_RedShift", 0), // property RedShift
			/* 36 */ imports.NewTable("TRawImageDescriptionWrap_GreenPrec", 0), // property GreenPrec
			/* 37 */ imports.NewTable("TRawImageDescriptionWrap_GreenShift", 0), // property GreenShift
			/* 38 */ imports.NewTable("TRawImageDescriptionWrap_BluePrec", 0), // property BluePrec
			/* 39 */ imports.NewTable("TRawImageDescriptionWrap_BlueShift", 0), // property BlueShift
			/* 40 */ imports.NewTable("TRawImageDescriptionWrap_AlphaPrec", 0), // property AlphaPrec
			/* 41 */ imports.NewTable("TRawImageDescriptionWrap_AlphaShift", 0), // property AlphaShift
			/* 42 */ imports.NewTable("TRawImageDescriptionWrap_MaskBitsPerPixel", 0), // property MaskBitsPerPixel
			/* 43 */ imports.NewTable("TRawImageDescriptionWrap_MaskShift", 0), // property MaskShift
			/* 44 */ imports.NewTable("TRawImageDescriptionWrap_MaskLineEnd", 0), // property MaskLineEnd
			/* 45 */ imports.NewTable("TRawImageDescriptionWrap_MaskBitOrder", 0), // property MaskBitOrder
			/* 46 */ imports.NewTable("TRawImageDescriptionWrap_PaletteColorCount", 0), // property PaletteColorCount
			/* 47 */ imports.NewTable("TRawImageDescriptionWrap_PaletteBitsPerIndex", 0), // property PaletteBitsPerIndex
			/* 48 */ imports.NewTable("TRawImageDescriptionWrap_PaletteShift", 0), // property PaletteShift
			/* 49 */ imports.NewTable("TRawImageDescriptionWrap_PaletteLineEnd", 0), // property PaletteLineEnd
			/* 50 */ imports.NewTable("TRawImageDescriptionWrap_PaletteBitOrder", 0), // property PaletteBitOrder
			/* 51 */ imports.NewTable("TRawImageDescriptionWrap_PaletteByteOrder", 0), // property PaletteByteOrder
		}
	})
	return rawImageDescriptionWrapImport
}
