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
	r1 := lazIntfImageImportAPI().SysCallN(5, uintptr(AWidth), uintptr(AHeight))
	return AsLazIntfImage(r1)
}

func NewLazIntfImage1(AWidth, AHeight int32, AFlags TRawImageQueryFlags) ILazIntfImage {
	r1 := lazIntfImageImportAPI().SysCallN(6, uintptr(AWidth), uintptr(AHeight), uintptr(AFlags))
	return AsLazIntfImage(r1)
}

func NewLazIntfImage2(ARawImage *TRawImage, ADataOwner bool) ILazIntfImage {
	r1 := lazIntfImageImportAPI().SysCallN(7, uintptr(unsafePointer(ARawImage)), PascalBool(ADataOwner))
	return AsLazIntfImage(r1)
}

func NewLazIntfImageCompatible(IntfImg ILazIntfImage, AWidth, AHeight int32) ILazIntfImage {
	r1 := lazIntfImageImportAPI().SysCallN(9, GetObjectUintptr(IntfImg), uintptr(AWidth), uintptr(AHeight))
	return AsLazIntfImage(r1)
}

func (m *TLazIntfImage) PixelData() PByte {
	r1 := lazIntfImageImportAPI().SysCallN(24, m.Instance())
	return PByte(r1)
}

func (m *TLazIntfImage) MaskData() PByte {
	r1 := lazIntfImageImportAPI().SysCallN(22, m.Instance())
	return PByte(r1)
}

func (m *TLazIntfImage) TColors(x, y int32) TGraphicsColor {
	r1 := lazIntfImageImportAPI().SysCallN(26, 0, m.Instance(), uintptr(x), uintptr(y))
	return TGraphicsColor(r1)
}

func (m *TLazIntfImage) SetTColors(x, y int32, AValue TGraphicsColor) {
	lazIntfImageImportAPI().SysCallN(26, 1, m.Instance(), uintptr(x), uintptr(y), uintptr(AValue))
}

func (m *TLazIntfImage) Masked(x, y int32) bool {
	r1 := lazIntfImageImportAPI().SysCallN(23, 0, m.Instance(), uintptr(x), uintptr(y))
	return GoBool(r1)
}

func (m *TLazIntfImage) SetMasked(x, y int32, AValue bool) {
	lazIntfImageImportAPI().SysCallN(23, 1, m.Instance(), uintptr(x), uintptr(y), PascalBool(AValue))
}

func (m *TLazIntfImage) GetDataLineStart(y int32) uintptr {
	r1 := lazIntfImageImportAPI().SysCallN(13, m.Instance(), uintptr(y))
	return uintptr(r1)
}

func (m *TLazIntfImage) HasTransparency() bool {
	r1 := lazIntfImageImportAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TLazIntfImage) HasMask() bool {
	r1 := lazIntfImageImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

func LazIntfImageClass() TClass {
	ret := lazIntfImageImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TLazIntfImage) BeginUpdate() {
	lazIntfImageImportAPI().SysCallN(2, m.Instance())
}

func (m *TLazIntfImage) EndUpdate() {
	lazIntfImageImportAPI().SysCallN(11, m.Instance())
}

func (m *TLazIntfImage) LoadFromDevice(DC HDC) {
	lazIntfImageImportAPI().SysCallN(20, m.Instance(), uintptr(DC))
}

func (m *TLazIntfImage) LoadFromBitmap(ABitmap, AMaskBitmap HBITMAP, AWidth int32, AHeight int32) {
	lazIntfImageImportAPI().SysCallN(19, m.Instance(), uintptr(ABitmap), uintptr(AMaskBitmap), uintptr(AWidth), uintptr(AHeight))
}

func (m *TLazIntfImage) CreateBitmaps(OutBitmap, OutMask *HBITMAP, ASkipMask bool) {
	var result0 uintptr
	var result1 uintptr
	lazIntfImageImportAPI().SysCallN(8, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)), PascalBool(ASkipMask))
	*OutBitmap = HBITMAP(result0)
	*OutMask = HBITMAP(result1)
}

func (m *TLazIntfImage) SetRawImage(ARawImage *TRawImage, ADataOwner bool) {
	lazIntfImageImportAPI().SysCallN(25, m.Instance(), uintptr(unsafePointer(ARawImage)), PascalBool(ADataOwner))
}

func (m *TLazIntfImage) GetRawImage(OutRawImage *TRawImage, ATransferOwnership bool) {
	var result0 uintptr
	lazIntfImageImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(&result0)), PascalBool(ATransferOwnership))
	*OutRawImage = *(*TRawImage)(getPointer(result0))
}

func (m *TLazIntfImage) FillPixels(Color *TFPColor) {
	lazIntfImageImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(Color)))
}

func (m *TLazIntfImage) CopyPixels(ASource IFPCustomImage, XDst int32, YDst int32, AlphaMask bool, AlphaTreshold Word) {
	lazIntfImageImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(ASource), uintptr(XDst), uintptr(YDst), PascalBool(AlphaMask), uintptr(AlphaTreshold))
}

func (m *TLazIntfImage) AlphaBlend(ASource, ASourceAlpha ILazIntfImage, ADestX, ADestY int32) {
	lazIntfImageImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(ASource), GetObjectUintptr(ASourceAlpha), uintptr(ADestX), uintptr(ADestY))
}

func (m *TLazIntfImage) AlphaFromMask(AKeepAlpha bool) {
	lazIntfImageImportAPI().SysCallN(1, m.Instance(), PascalBool(AKeepAlpha))
}

func (m *TLazIntfImage) Mask(AColor *TFPColor, AKeepOldMask bool) {
	lazIntfImageImportAPI().SysCallN(21, m.Instance(), uintptr(unsafePointer(AColor)), PascalBool(AKeepOldMask))
}

func (m *TLazIntfImage) GetXYDataPosition(x, y int32, OutPosition *TRawImagePosition) {
	var result1 uintptr
	lazIntfImageImportAPI().SysCallN(15, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(&result1)))
	*OutPosition = *(*TRawImagePosition)(getPointer(result1))
}

func (m *TLazIntfImage) GetXYMaskPosition(x, y int32, OutPosition *TRawImagePosition) {
	var result1 uintptr
	lazIntfImageImportAPI().SysCallN(16, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(&result1)))
	*OutPosition = *(*TRawImagePosition)(getPointer(result1))
}

func (m *TLazIntfImage) CreateData() {
	lazIntfImageImportAPI().SysCallN(10, m.Instance())
}

var (
	lazIntfImageImport       *imports.Imports = nil
	lazIntfImageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LazIntfImage_AlphaBlend", 0),
		/*1*/ imports.NewTable("LazIntfImage_AlphaFromMask", 0),
		/*2*/ imports.NewTable("LazIntfImage_BeginUpdate", 0),
		/*3*/ imports.NewTable("LazIntfImage_Class", 0),
		/*4*/ imports.NewTable("LazIntfImage_CopyPixels", 0),
		/*5*/ imports.NewTable("LazIntfImage_Create", 0),
		/*6*/ imports.NewTable("LazIntfImage_Create1", 0),
		/*7*/ imports.NewTable("LazIntfImage_Create2", 0),
		/*8*/ imports.NewTable("LazIntfImage_CreateBitmaps", 0),
		/*9*/ imports.NewTable("LazIntfImage_CreateCompatible", 0),
		/*10*/ imports.NewTable("LazIntfImage_CreateData", 0),
		/*11*/ imports.NewTable("LazIntfImage_EndUpdate", 0),
		/*12*/ imports.NewTable("LazIntfImage_FillPixels", 0),
		/*13*/ imports.NewTable("LazIntfImage_GetDataLineStart", 0),
		/*14*/ imports.NewTable("LazIntfImage_GetRawImage", 0),
		/*15*/ imports.NewTable("LazIntfImage_GetXYDataPosition", 0),
		/*16*/ imports.NewTable("LazIntfImage_GetXYMaskPosition", 0),
		/*17*/ imports.NewTable("LazIntfImage_HasMask", 0),
		/*18*/ imports.NewTable("LazIntfImage_HasTransparency", 0),
		/*19*/ imports.NewTable("LazIntfImage_LoadFromBitmap", 0),
		/*20*/ imports.NewTable("LazIntfImage_LoadFromDevice", 0),
		/*21*/ imports.NewTable("LazIntfImage_Mask", 0),
		/*22*/ imports.NewTable("LazIntfImage_MaskData", 0),
		/*23*/ imports.NewTable("LazIntfImage_Masked", 0),
		/*24*/ imports.NewTable("LazIntfImage_PixelData", 0),
		/*25*/ imports.NewTable("LazIntfImage_SetRawImage", 0),
		/*26*/ imports.NewTable("LazIntfImage_TColors", 0),
	}
)

func lazIntfImageImportAPI() *imports.Imports {
	if lazIntfImageImport == nil {
		lazIntfImageImport = NewDefaultImports()
		lazIntfImageImport.SetImportTable(lazIntfImageImportTables)
		lazIntfImageImportTables = nil
	}
	return lazIntfImageImport
}
