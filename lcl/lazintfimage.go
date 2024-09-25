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

// ILazIntfImage Parent: IFPCustomImage
type ILazIntfImage interface {
	IFPCustomImage
	DataDescription() *TRawImageDescription
	SetDataDescription(AValue *TRawImageDescription)
	CheckDescription(ADescription *TRawImageDescription, ExceptionOnError bool) bool
	SetDataDescriptionKeepData(ADescription *TRawImageDescription)
	PixelData() PByte                                                                              // property
	MaskData() PByte                                                                               // property
	TColors(x, y int32) TGraphicsColor                                                             // property
	SetTColors(x, y int32, AValue TGraphicsColor)                                                  // property
	Masked(x, y int32) bool                                                                        // property
	SetMasked(x, y int32, AValue bool)                                                             // property
	GetDataLineStart(y int32) uintptr                                                              // function
	HasTransparency() bool                                                                         // function
	HasMask() bool                                                                                 // function
	BeginUpdate()                                                                                  // procedure
	EndUpdate()                                                                                    // procedure
	LoadFromDevice(DC HDC)                                                                         // procedure
	LoadFromBitmap(ABitmap, AMaskBitmap HBITMAP, AWidth int32, AHeight int32)                      // procedure
	CreateBitmaps(OutBitmap, OutMask *HBITMAP, ASkipMask bool)                                     // procedure
	SetRawImage(ARawImage *TRawImage, ADataOwner bool)                                             // procedure
	GetRawImage(OutRawImage *TRawImage, ATransferOwnership bool)                                   // procedure
	FillPixels(Color *TFPColor)                                                                    // procedure
	CopyPixels(ASource IFPCustomImage, XDst int32, YDst int32, AlphaMask bool, AlphaTreshold Word) // procedure
	AlphaBlend(ASource, ASourceAlpha ILazIntfImage, ADestX, ADestY int32)                          // procedure
	AlphaFromMask(AKeepAlpha bool)                                                                 // procedure
	Mask(AColor *TFPColor, AKeepOldMask bool)                                                      // procedure
	GetXYDataPosition(x, y int32, OutPosition *TRawImagePosition)                                  // procedure
	GetXYMaskPosition(x, y int32, OutPosition *TRawImagePosition)                                  // procedure
	CreateData()                                                                                   // procedure
}

// TLazIntfImage Parent: TFPCustomImage
type TLazIntfImage struct {
	TFPCustomImage
}

func NewLazIntfImage(AWidth, AHeight int32) ILazIntfImage {
	r1 := LCL().SysCallN(3584, uintptr(AWidth), uintptr(AHeight))
	return AsLazIntfImage(r1)
}

func NewLazIntfImage1(AWidth, AHeight int32, AFlags TRawImageQueryFlags) ILazIntfImage {
	r1 := LCL().SysCallN(3585, uintptr(AWidth), uintptr(AHeight), uintptr(AFlags))
	return AsLazIntfImage(r1)
}

func NewLazIntfImage2(ARawImage *TRawImage, ADataOwner bool) ILazIntfImage {
	r1 := LCL().SysCallN(3586, uintptr(unsafePointer(ARawImage)), PascalBool(ADataOwner))
	return AsLazIntfImage(r1)
}

func NewLazIntfImageCompatible(IntfImg ILazIntfImage, AWidth, AHeight int32) ILazIntfImage {
	r1 := LCL().SysCallN(3588, GetObjectUintptr(IntfImg), uintptr(AWidth), uintptr(AHeight))
	return AsLazIntfImage(r1)
}

func (m *TLazIntfImage) PixelData() PByte {
	r1 := LCL().SysCallN(3603, m.Instance())
	return PByte(r1)
}

func (m *TLazIntfImage) MaskData() PByte {
	r1 := LCL().SysCallN(3601, m.Instance())
	return PByte(r1)
}

func (m *TLazIntfImage) TColors(x, y int32) TGraphicsColor {
	r1 := LCL().SysCallN(3605, 0, m.Instance(), uintptr(x), uintptr(y))
	return TGraphicsColor(r1)
}

func (m *TLazIntfImage) SetTColors(x, y int32, AValue TGraphicsColor) {
	LCL().SysCallN(3605, 1, m.Instance(), uintptr(x), uintptr(y), uintptr(AValue))
}

func (m *TLazIntfImage) Masked(x, y int32) bool {
	r1 := LCL().SysCallN(3602, 0, m.Instance(), uintptr(x), uintptr(y))
	return GoBool(r1)
}

func (m *TLazIntfImage) SetMasked(x, y int32, AValue bool) {
	LCL().SysCallN(3602, 1, m.Instance(), uintptr(x), uintptr(y), PascalBool(AValue))
}

func (m *TLazIntfImage) GetDataLineStart(y int32) uintptr {
	r1 := LCL().SysCallN(3592, m.Instance(), uintptr(y))
	return uintptr(r1)
}

func (m *TLazIntfImage) HasTransparency() bool {
	r1 := LCL().SysCallN(3597, m.Instance())
	return GoBool(r1)
}

func (m *TLazIntfImage) HasMask() bool {
	r1 := LCL().SysCallN(3596, m.Instance())
	return GoBool(r1)
}

func LazIntfImageClass() TClass {
	ret := LCL().SysCallN(3582)
	return TClass(ret)
}

func (m *TLazIntfImage) BeginUpdate() {
	LCL().SysCallN(3581, m.Instance())
}

func (m *TLazIntfImage) EndUpdate() {
	LCL().SysCallN(3590, m.Instance())
}

func (m *TLazIntfImage) LoadFromDevice(DC HDC) {
	LCL().SysCallN(3599, m.Instance(), uintptr(DC))
}

func (m *TLazIntfImage) LoadFromBitmap(ABitmap, AMaskBitmap HBITMAP, AWidth int32, AHeight int32) {
	LCL().SysCallN(3598, m.Instance(), uintptr(ABitmap), uintptr(AMaskBitmap), uintptr(AWidth), uintptr(AHeight))
}

func (m *TLazIntfImage) CreateBitmaps(OutBitmap, OutMask *HBITMAP, ASkipMask bool) {
	var result0 uintptr
	var result1 uintptr
	LCL().SysCallN(3587, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)), PascalBool(ASkipMask))
	*OutBitmap = HBITMAP(result0)
	*OutMask = HBITMAP(result1)
}

func (m *TLazIntfImage) SetRawImage(ARawImage *TRawImage, ADataOwner bool) {
	LCL().SysCallN(3604, m.Instance(), uintptr(unsafePointer(ARawImage)), PascalBool(ADataOwner))
}

func (m *TLazIntfImage) GetRawImage(OutRawImage *TRawImage, ATransferOwnership bool) {
	var result0 uintptr
	LCL().SysCallN(3593, m.Instance(), uintptr(unsafePointer(&result0)), PascalBool(ATransferOwnership))
	*OutRawImage = *(*TRawImage)(getPointer(result0))
}

func (m *TLazIntfImage) FillPixels(Color *TFPColor) {
	LCL().SysCallN(3591, m.Instance(), uintptr(unsafePointer(Color)))
}

func (m *TLazIntfImage) CopyPixels(ASource IFPCustomImage, XDst int32, YDst int32, AlphaMask bool, AlphaTreshold Word) {
	LCL().SysCallN(3583, m.Instance(), GetObjectUintptr(ASource), uintptr(XDst), uintptr(YDst), PascalBool(AlphaMask), uintptr(AlphaTreshold))
}

func (m *TLazIntfImage) AlphaBlend(ASource, ASourceAlpha ILazIntfImage, ADestX, ADestY int32) {
	LCL().SysCallN(3579, m.Instance(), GetObjectUintptr(ASource), GetObjectUintptr(ASourceAlpha), uintptr(ADestX), uintptr(ADestY))
}

func (m *TLazIntfImage) AlphaFromMask(AKeepAlpha bool) {
	LCL().SysCallN(3580, m.Instance(), PascalBool(AKeepAlpha))
}

func (m *TLazIntfImage) Mask(AColor *TFPColor, AKeepOldMask bool) {
	LCL().SysCallN(3600, m.Instance(), uintptr(unsafePointer(AColor)), PascalBool(AKeepOldMask))
}

func (m *TLazIntfImage) GetXYDataPosition(x, y int32, OutPosition *TRawImagePosition) {
	var result1 uintptr
	LCL().SysCallN(3594, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(&result1)))
	*OutPosition = *(*TRawImagePosition)(getPointer(result1))
}

func (m *TLazIntfImage) GetXYMaskPosition(x, y int32, OutPosition *TRawImagePosition) {
	var result1 uintptr
	LCL().SysCallN(3595, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(&result1)))
	*OutPosition = *(*TRawImagePosition)(getPointer(result1))
}

func (m *TLazIntfImage) CreateData() {
	LCL().SysCallN(3589, m.Instance())
}
