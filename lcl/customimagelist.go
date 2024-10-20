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

// ICustomImageList Parent: ILCLComponent
type ICustomImageList interface {
	ILCLComponent
	HasOverlays() bool                                                                                                                                          // property
	AllocBy() int32                                                                                                                                             // property
	SetAllocBy(AValue int32)                                                                                                                                    // property
	BlendColor() TColor                                                                                                                                         // property
	SetBlendColor(AValue TColor)                                                                                                                                // property
	BkColor() TColor                                                                                                                                            // property
	SetBkColor(AValue TColor)                                                                                                                                   // property
	Count() int32                                                                                                                                               // property
	DrawingStyle() TDrawingStyle                                                                                                                                // property
	SetDrawingStyle(AValue TDrawingStyle)                                                                                                                       // property
	Height() int32                                                                                                                                              // property
	SetHeight(AValue int32)                                                                                                                                     // property
	HeightForPPI(AImageWidth, APPI int32) int32                                                                                                                 // property
	HeightForWidth(AWidth int32) int32                                                                                                                          // property
	Width() int32                                                                                                                                               // property
	SetWidth(AValue int32)                                                                                                                                      // property
	WidthForPPI(AImageWidth, APPI int32) int32                                                                                                                  // property
	SizeForPPI(AImageWidth, APPI int32) (resultSize TSize)                                                                                                      // property
	Masked() bool                                                                                                                                               // property
	SetMasked(AValue bool)                                                                                                                                      // property
	Resolution(AImageWidth int32) ICustomImageListResolution                                                                                                    // property
	ResolutionByIndex(AIndex int32) ICustomImageListResolution                                                                                                  // property
	ResolutionForPPI(AImageWidth, APPI int32, ACanvasScaleFactor float64) (resultScaledImageListResolution TScaledImageListResolution)                          // property
	ResolutionCount() int32                                                                                                                                     // property
	Scaled() bool                                                                                                                                               // property
	SetScaled(AValue bool)                                                                                                                                      // property
	ShareImages() bool                                                                                                                                          // property
	SetShareImages(AValue bool)                                                                                                                                 // property
	ImageType() TImageType                                                                                                                                      // property
	SetImageType(AValue TImageType)                                                                                                                             // property
	Add(Image, Mask ICustomBitmap) int32                                                                                                                        // function
	AddSliced(Image ICustomBitmap, AHorizontalCount, AVerticalCount int32) int32                                                                                // function
	AddSlice(Image ICustomBitmap, AImageRect *TRect) int32                                                                                                      // function
	AddSliceCentered(Image ICustomBitmap) int32                                                                                                                 // function
	AddIcon(Image ICustomIcon) int32                                                                                                                            // function
	AddMasked(Image IBitmap, MaskColor TColor) int32                                                                                                            // function
	AddLazarusResource(ResourceName string, MaskColor TColor) int32                                                                                             // function
	AddResourceName(Instance THandle, ResourceName string, MaskColor TColor) int32                                                                              // function
	GetHotSpot() (resultPoint TPoint)                                                                                                                           // function
	FindResolution(AImageWidth int32, OutResolution *ICustomImageListResolution) bool                                                                           // function
	Resolutions() ICustomImageListResolutionEnumerator                                                                                                          // function
	ResolutionsDesc() ICustomImageListResolutionEnumerator                                                                                                      // function
	AssignTo(Dest IPersistent)                                                                                                                                  // procedure
	WriteData(AStream IStream)                                                                                                                                  // procedure
	ReadData(AStream IStream)                                                                                                                                   // procedure
	WriteAdvData(AStream IStream)                                                                                                                               // procedure
	ReadAdvData(AStream IStream)                                                                                                                                // procedure
	BeginUpdate()                                                                                                                                               // procedure
	EndUpdate()                                                                                                                                                 // procedure
	AddImages(AValue ICustomImageList)                                                                                                                          // procedure
	Change()                                                                                                                                                    // procedure
	Clear()                                                                                                                                                     // procedure
	Delete(AIndex int32)                                                                                                                                        // procedure
	Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool)                                                                                                  // procedure
	Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect)                                                                               // procedure
	Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool)                                             // procedure
	Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect)                           // procedure
	DrawForPPI(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, AEnabled bool)                               // procedure
	DrawForPPI1(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, ADrawEffect TGraphicsDrawEffect)            // procedure
	DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool)                                                                        // procedure
	DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect)                                                     // procedure
	DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) // procedure
	GetBitmap(Index int32, Image ICustomBitmap)                                                                                                                 // procedure
	GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect)                                                                                   // procedure
	GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect)                                                                                             // procedure
	GetFullRawImage(OutImage *TRawImage)                                                                                                                        // procedure
	GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect)                                                                                              // procedure
	GetIcon1(Index int32, Image IIcon)                                                                                                                          // procedure
	GetRawImage(Index int32, OutImage *TRawImage)                                                                                                               // procedure
	Insert(AIndex int32, AImage, AMask ICustomBitmap)                                                                                                           // procedure
	InsertIcon(AIndex int32, AIcon ICustomIcon)                                                                                                                 // procedure
	InsertMasked(Index int32, AImage ICustomBitmap, MaskColor TColor)                                                                                           // procedure
	Move(ACurIndex, ANewIndex int32)                                                                                                                            // procedure
	Overlay(AIndex int32, Overlay TOverlay)                                                                                                                     // procedure
	Replace(AIndex int32, AImage, AMask ICustomBitmap, AllResolutions bool)                                                                                     // procedure
	ReplaceSlice(AIndex int32, Image ICustomBitmap, AImageRect *TRect, AllResolutions bool)                                                                     // procedure
	ReplaceSliceCentered(AIndex, AImageWidth int32, Image ICustomBitmap, AllResolutions bool)                                                                   // procedure
	ReplaceIcon(AIndex int32, AIcon ICustomIcon)                                                                                                                // procedure
	ReplaceMasked(Index int32, NewImage ICustomBitmap, MaskColor TColor, AllResolutions bool)                                                                   // procedure
	RegisterChanges(Value IChangeLink)                                                                                                                          // procedure
	StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool)                                                                                        // procedure
	UnRegisterChanges(Value IChangeLink)                                                                                                                        // procedure
	DeleteResolution(AWidth int32)                                                                                                                              // procedure
	SetOnChange(fn TNotifyEvent)                                                                                                                                // property event
	SetOnGetWidthForPPI(fn TCustomImageListGetWidthForPPI)                                                                                                      // property event
}

// TCustomImageList Parent: TLCLComponent
type TCustomImageList struct {
	TLCLComponent
	changePtr         uintptr
	getWidthForPPIPtr uintptr
}

func NewCustomImageList(AOwner IComponent) ICustomImageList {
	r1 := customImageListImportAPI().SysCallN(18, GetObjectUintptr(AOwner))
	return AsCustomImageList(r1)
}

func NewCustomImageListSize(AWidth, AHeight int32) ICustomImageList {
	r1 := customImageListImportAPI().SysCallN(19, uintptr(AWidth), uintptr(AHeight))
	return AsCustomImageList(r1)
}

func (m *TCustomImageList) HasOverlays() bool {
	r1 := customImageListImportAPI().SysCallN(42, m.Instance())
	return GoBool(r1)
}

func (m *TCustomImageList) AllocBy() int32 {
	r1 := customImageListImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImageList) SetAllocBy(AValue int32) {
	customImageListImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) BlendColor() TColor {
	r1 := customImageListImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomImageList) SetBlendColor(AValue TColor) {
	customImageListImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) BkColor() TColor {
	r1 := customImageListImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomImageList) SetBkColor(AValue TColor) {
	customImageListImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) Count() int32 {
	r1 := customImageListImportAPI().SysCallN(17, m.Instance())
	return int32(r1)
}

func (m *TCustomImageList) DrawingStyle() TDrawingStyle {
	r1 := customImageListImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return TDrawingStyle(r1)
}

func (m *TCustomImageList) SetDrawingStyle(AValue TDrawingStyle) {
	customImageListImportAPI().SysCallN(31, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) Height() int32 {
	r1 := customImageListImportAPI().SysCallN(43, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImageList) SetHeight(AValue int32) {
	customImageListImportAPI().SysCallN(43, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) HeightForPPI(AImageWidth, APPI int32) int32 {
	r1 := customImageListImportAPI().SysCallN(44, m.Instance(), uintptr(AImageWidth), uintptr(APPI))
	return int32(r1)
}

func (m *TCustomImageList) HeightForWidth(AWidth int32) int32 {
	r1 := customImageListImportAPI().SysCallN(45, m.Instance(), uintptr(AWidth))
	return int32(r1)
}

func (m *TCustomImageList) Width() int32 {
	r1 := customImageListImportAPI().SysCallN(74, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImageList) SetWidth(AValue int32) {
	customImageListImportAPI().SysCallN(74, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) WidthForPPI(AImageWidth, APPI int32) int32 {
	r1 := customImageListImportAPI().SysCallN(75, m.Instance(), uintptr(AImageWidth), uintptr(APPI))
	return int32(r1)
}

func (m *TCustomImageList) SizeForPPI(AImageWidth, APPI int32) (resultSize TSize) {
	customImageListImportAPI().SysCallN(71, m.Instance(), uintptr(AImageWidth), uintptr(APPI), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TCustomImageList) Masked() bool {
	r1 := customImageListImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageList) SetMasked(AValue bool) {
	customImageListImportAPI().SysCallN(50, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageList) Resolution(AImageWidth int32) ICustomImageListResolution {
	r1 := customImageListImportAPI().SysCallN(61, m.Instance(), uintptr(AImageWidth))
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageList) ResolutionByIndex(AIndex int32) ICustomImageListResolution {
	r1 := customImageListImportAPI().SysCallN(62, m.Instance(), uintptr(AIndex))
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageList) ResolutionForPPI(AImageWidth, APPI int32, ACanvasScaleFactor float64) (resultScaledImageListResolution TScaledImageListResolution) {
	r1 := customImageListImportAPI().SysCallN(64, m.Instance(), uintptr(AImageWidth), uintptr(APPI), uintptr(unsafePointer(&ACanvasScaleFactor)))
	return *(*TScaledImageListResolution)(getPointer(r1))
}

func (m *TCustomImageList) ResolutionCount() int32 {
	r1 := customImageListImportAPI().SysCallN(63, m.Instance())
	return int32(r1)
}

func (m *TCustomImageList) Scaled() bool {
	r1 := customImageListImportAPI().SysCallN(67, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageList) SetScaled(AValue bool) {
	customImageListImportAPI().SysCallN(67, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageList) ShareImages() bool {
	r1 := customImageListImportAPI().SysCallN(70, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageList) SetShareImages(AValue bool) {
	customImageListImportAPI().SysCallN(70, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageList) ImageType() TImageType {
	r1 := customImageListImportAPI().SysCallN(46, 0, m.Instance(), 0)
	return TImageType(r1)
}

func (m *TCustomImageList) SetImageType(AValue TImageType) {
	customImageListImportAPI().SysCallN(46, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImageList) Add(Image, Mask ICustomBitmap) int32 {
	r1 := customImageListImportAPI().SysCallN(0, m.Instance(), GetObjectUintptr(Image), GetObjectUintptr(Mask))
	return int32(r1)
}

func (m *TCustomImageList) AddSliced(Image ICustomBitmap, AHorizontalCount, AVerticalCount int32) int32 {
	r1 := customImageListImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(Image), uintptr(AHorizontalCount), uintptr(AVerticalCount))
	return int32(r1)
}

func (m *TCustomImageList) AddSlice(Image ICustomBitmap, AImageRect *TRect) int32 {
	r1 := customImageListImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(Image), uintptr(unsafePointer(AImageRect)))
	return int32(r1)
}

func (m *TCustomImageList) AddSliceCentered(Image ICustomBitmap) int32 {
	r1 := customImageListImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(Image))
	return int32(r1)
}

func (m *TCustomImageList) AddIcon(Image ICustomIcon) int32 {
	r1 := customImageListImportAPI().SysCallN(1, m.Instance(), GetObjectUintptr(Image))
	return int32(r1)
}

func (m *TCustomImageList) AddMasked(Image IBitmap, MaskColor TColor) int32 {
	r1 := customImageListImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(Image), uintptr(MaskColor))
	return int32(r1)
}

func (m *TCustomImageList) AddLazarusResource(ResourceName string, MaskColor TColor) int32 {
	r1 := customImageListImportAPI().SysCallN(3, m.Instance(), PascalStr(ResourceName), uintptr(MaskColor))
	return int32(r1)
}

func (m *TCustomImageList) AddResourceName(Instance THandle, ResourceName string, MaskColor TColor) int32 {
	r1 := customImageListImportAPI().SysCallN(5, m.Instance(), uintptr(Instance), PascalStr(ResourceName), uintptr(MaskColor))
	return int32(r1)
}

func (m *TCustomImageList) GetHotSpot() (resultPoint TPoint) {
	customImageListImportAPI().SysCallN(38, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomImageList) FindResolution(AImageWidth int32, OutResolution *ICustomImageListResolution) bool {
	var result1 uintptr
	r1 := customImageListImportAPI().SysCallN(33, m.Instance(), uintptr(AImageWidth), uintptr(unsafePointer(&result1)))
	*OutResolution = AsCustomImageListResolution(result1)
	return GoBool(r1)
}

func (m *TCustomImageList) Resolutions() ICustomImageListResolutionEnumerator {
	r1 := customImageListImportAPI().SysCallN(65, m.Instance())
	return AsCustomImageListResolutionEnumerator(r1)
}

func (m *TCustomImageList) ResolutionsDesc() ICustomImageListResolutionEnumerator {
	r1 := customImageListImportAPI().SysCallN(66, m.Instance())
	return AsCustomImageListResolutionEnumerator(r1)
}

func CustomImageListClass() TClass {
	ret := customImageListImportAPI().SysCallN(15)
	return TClass(ret)
}

func (m *TCustomImageList) AssignTo(Dest IPersistent) {
	customImageListImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(Dest))
}

func (m *TCustomImageList) WriteData(AStream IStream) {
	customImageListImportAPI().SysCallN(77, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) ReadData(AStream IStream) {
	customImageListImportAPI().SysCallN(54, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) WriteAdvData(AStream IStream) {
	customImageListImportAPI().SysCallN(76, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) ReadAdvData(AStream IStream) {
	customImageListImportAPI().SysCallN(53, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomImageList) BeginUpdate() {
	customImageListImportAPI().SysCallN(11, m.Instance())
}

func (m *TCustomImageList) EndUpdate() {
	customImageListImportAPI().SysCallN(32, m.Instance())
}

func (m *TCustomImageList) AddImages(AValue ICustomImageList) {
	customImageListImportAPI().SysCallN(2, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomImageList) Change() {
	customImageListImportAPI().SysCallN(14, m.Instance())
}

func (m *TCustomImageList) Clear() {
	customImageListImportAPI().SysCallN(16, m.Instance())
}

func (m *TCustomImageList) Delete(AIndex int32) {
	customImageListImportAPI().SysCallN(20, m.Instance(), uintptr(AIndex))
}

func (m *TCustomImageList) Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool) {
	customImageListImportAPI().SysCallN(22, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), PascalBool(AEnabled))
}

func (m *TCustomImageList) Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect) {
	customImageListImportAPI().SysCallN(23, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawEffect))
}

func (m *TCustomImageList) Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool) {
	customImageListImportAPI().SysCallN(24, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), PascalBool(AEnabled))
}

func (m *TCustomImageList) Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	customImageListImportAPI().SysCallN(25, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

func (m *TCustomImageList) DrawForPPI(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, AEnabled bool) {
	customImageListImportAPI().SysCallN(26, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AImageWidthAt96PPI), uintptr(ATargetPPI), uintptr(unsafePointer(&ACanvasFactor)), PascalBool(AEnabled))
}

func (m *TCustomImageList) DrawForPPI1(ACanvas ICanvas, AX, AY, AIndex int32, AImageWidthAt96PPI, ATargetPPI int32, ACanvasFactor float64, ADrawEffect TGraphicsDrawEffect) {
	customImageListImportAPI().SysCallN(27, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AImageWidthAt96PPI), uintptr(ATargetPPI), uintptr(unsafePointer(&ACanvasFactor)), uintptr(ADrawEffect))
}

func (m *TCustomImageList) DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool) {
	customImageListImportAPI().SysCallN(28, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), PascalBool(AEnabled))
}

func (m *TCustomImageList) DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect) {
	customImageListImportAPI().SysCallN(29, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawEffect))
}

func (m *TCustomImageList) DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	customImageListImportAPI().SysCallN(30, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

func (m *TCustomImageList) GetBitmap(Index int32, Image ICustomBitmap) {
	customImageListImportAPI().SysCallN(34, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageList) GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	customImageListImportAPI().SysCallN(35, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageList) GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	customImageListImportAPI().SysCallN(36, m.Instance(), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageList) GetFullRawImage(OutImage *TRawImage) {
	var result0 uintptr
	customImageListImportAPI().SysCallN(37, m.Instance(), uintptr(unsafePointer(&result0)))
	*OutImage = *(*TRawImage)(getPointer(result0))
}

func (m *TCustomImageList) GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect) {
	customImageListImportAPI().SysCallN(39, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageList) GetIcon1(Index int32, Image IIcon) {
	customImageListImportAPI().SysCallN(40, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageList) GetRawImage(Index int32, OutImage *TRawImage) {
	var result1 uintptr
	customImageListImportAPI().SysCallN(41, m.Instance(), uintptr(Index), uintptr(unsafePointer(&result1)))
	*OutImage = *(*TRawImage)(getPointer(result1))
}

func (m *TCustomImageList) Insert(AIndex int32, AImage, AMask ICustomBitmap) {
	customImageListImportAPI().SysCallN(47, m.Instance(), uintptr(AIndex), GetObjectUintptr(AImage), GetObjectUintptr(AMask))
}

func (m *TCustomImageList) InsertIcon(AIndex int32, AIcon ICustomIcon) {
	customImageListImportAPI().SysCallN(48, m.Instance(), uintptr(AIndex), GetObjectUintptr(AIcon))
}

func (m *TCustomImageList) InsertMasked(Index int32, AImage ICustomBitmap, MaskColor TColor) {
	customImageListImportAPI().SysCallN(49, m.Instance(), uintptr(Index), GetObjectUintptr(AImage), uintptr(MaskColor))
}

func (m *TCustomImageList) Move(ACurIndex, ANewIndex int32) {
	customImageListImportAPI().SysCallN(51, m.Instance(), uintptr(ACurIndex), uintptr(ANewIndex))
}

func (m *TCustomImageList) Overlay(AIndex int32, Overlay TOverlay) {
	customImageListImportAPI().SysCallN(52, m.Instance(), uintptr(AIndex), uintptr(Overlay))
}

func (m *TCustomImageList) Replace(AIndex int32, AImage, AMask ICustomBitmap, AllResolutions bool) {
	customImageListImportAPI().SysCallN(56, m.Instance(), uintptr(AIndex), GetObjectUintptr(AImage), GetObjectUintptr(AMask), PascalBool(AllResolutions))
}

func (m *TCustomImageList) ReplaceSlice(AIndex int32, Image ICustomBitmap, AImageRect *TRect, AllResolutions bool) {
	customImageListImportAPI().SysCallN(59, m.Instance(), uintptr(AIndex), GetObjectUintptr(Image), uintptr(unsafePointer(AImageRect)), PascalBool(AllResolutions))
}

func (m *TCustomImageList) ReplaceSliceCentered(AIndex, AImageWidth int32, Image ICustomBitmap, AllResolutions bool) {
	customImageListImportAPI().SysCallN(60, m.Instance(), uintptr(AIndex), uintptr(AImageWidth), GetObjectUintptr(Image), PascalBool(AllResolutions))
}

func (m *TCustomImageList) ReplaceIcon(AIndex int32, AIcon ICustomIcon) {
	customImageListImportAPI().SysCallN(57, m.Instance(), uintptr(AIndex), GetObjectUintptr(AIcon))
}

func (m *TCustomImageList) ReplaceMasked(Index int32, NewImage ICustomBitmap, MaskColor TColor, AllResolutions bool) {
	customImageListImportAPI().SysCallN(58, m.Instance(), uintptr(Index), GetObjectUintptr(NewImage), uintptr(MaskColor), PascalBool(AllResolutions))
}

func (m *TCustomImageList) RegisterChanges(Value IChangeLink) {
	customImageListImportAPI().SysCallN(55, m.Instance(), GetObjectUintptr(Value))
}

func (m *TCustomImageList) StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool) {
	customImageListImportAPI().SysCallN(72, m.Instance(), GetObjectUintptr(Canvas), uintptr(Index), uintptr(unsafePointer(ARect)), PascalBool(Enabled))
}

func (m *TCustomImageList) UnRegisterChanges(Value IChangeLink) {
	customImageListImportAPI().SysCallN(73, m.Instance(), GetObjectUintptr(Value))
}

func (m *TCustomImageList) DeleteResolution(AWidth int32) {
	customImageListImportAPI().SysCallN(21, m.Instance(), uintptr(AWidth))
}

func (m *TCustomImageList) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	customImageListImportAPI().SysCallN(68, m.Instance(), m.changePtr)
}

func (m *TCustomImageList) SetOnGetWidthForPPI(fn TCustomImageListGetWidthForPPI) {
	if m.getWidthForPPIPtr != 0 {
		RemoveEventElement(m.getWidthForPPIPtr)
	}
	m.getWidthForPPIPtr = MakeEventDataPtr(fn)
	customImageListImportAPI().SysCallN(69, m.Instance(), m.getWidthForPPIPtr)
}

var (
	customImageListImport       *imports.Imports = nil
	customImageListImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomImageList_Add", 0),
		/*1*/ imports.NewTable("CustomImageList_AddIcon", 0),
		/*2*/ imports.NewTable("CustomImageList_AddImages", 0),
		/*3*/ imports.NewTable("CustomImageList_AddLazarusResource", 0),
		/*4*/ imports.NewTable("CustomImageList_AddMasked", 0),
		/*5*/ imports.NewTable("CustomImageList_AddResourceName", 0),
		/*6*/ imports.NewTable("CustomImageList_AddSlice", 0),
		/*7*/ imports.NewTable("CustomImageList_AddSliceCentered", 0),
		/*8*/ imports.NewTable("CustomImageList_AddSliced", 0),
		/*9*/ imports.NewTable("CustomImageList_AllocBy", 0),
		/*10*/ imports.NewTable("CustomImageList_AssignTo", 0),
		/*11*/ imports.NewTable("CustomImageList_BeginUpdate", 0),
		/*12*/ imports.NewTable("CustomImageList_BkColor", 0),
		/*13*/ imports.NewTable("CustomImageList_BlendColor", 0),
		/*14*/ imports.NewTable("CustomImageList_Change", 0),
		/*15*/ imports.NewTable("CustomImageList_Class", 0),
		/*16*/ imports.NewTable("CustomImageList_Clear", 0),
		/*17*/ imports.NewTable("CustomImageList_Count", 0),
		/*18*/ imports.NewTable("CustomImageList_Create", 0),
		/*19*/ imports.NewTable("CustomImageList_CreateSize", 0),
		/*20*/ imports.NewTable("CustomImageList_Delete", 0),
		/*21*/ imports.NewTable("CustomImageList_DeleteResolution", 0),
		/*22*/ imports.NewTable("CustomImageList_Draw", 0),
		/*23*/ imports.NewTable("CustomImageList_Draw1", 0),
		/*24*/ imports.NewTable("CustomImageList_Draw2", 0),
		/*25*/ imports.NewTable("CustomImageList_Draw3", 0),
		/*26*/ imports.NewTable("CustomImageList_DrawForPPI", 0),
		/*27*/ imports.NewTable("CustomImageList_DrawForPPI1", 0),
		/*28*/ imports.NewTable("CustomImageList_DrawOverlay", 0),
		/*29*/ imports.NewTable("CustomImageList_DrawOverlay1", 0),
		/*30*/ imports.NewTable("CustomImageList_DrawOverlay2", 0),
		/*31*/ imports.NewTable("CustomImageList_DrawingStyle", 0),
		/*32*/ imports.NewTable("CustomImageList_EndUpdate", 0),
		/*33*/ imports.NewTable("CustomImageList_FindResolution", 0),
		/*34*/ imports.NewTable("CustomImageList_GetBitmap", 0),
		/*35*/ imports.NewTable("CustomImageList_GetBitmap1", 0),
		/*36*/ imports.NewTable("CustomImageList_GetFullBitmap", 0),
		/*37*/ imports.NewTable("CustomImageList_GetFullRawImage", 0),
		/*38*/ imports.NewTable("CustomImageList_GetHotSpot", 0),
		/*39*/ imports.NewTable("CustomImageList_GetIcon", 0),
		/*40*/ imports.NewTable("CustomImageList_GetIcon1", 0),
		/*41*/ imports.NewTable("CustomImageList_GetRawImage", 0),
		/*42*/ imports.NewTable("CustomImageList_HasOverlays", 0),
		/*43*/ imports.NewTable("CustomImageList_Height", 0),
		/*44*/ imports.NewTable("CustomImageList_HeightForPPI", 0),
		/*45*/ imports.NewTable("CustomImageList_HeightForWidth", 0),
		/*46*/ imports.NewTable("CustomImageList_ImageType", 0),
		/*47*/ imports.NewTable("CustomImageList_Insert", 0),
		/*48*/ imports.NewTable("CustomImageList_InsertIcon", 0),
		/*49*/ imports.NewTable("CustomImageList_InsertMasked", 0),
		/*50*/ imports.NewTable("CustomImageList_Masked", 0),
		/*51*/ imports.NewTable("CustomImageList_Move", 0),
		/*52*/ imports.NewTable("CustomImageList_Overlay", 0),
		/*53*/ imports.NewTable("CustomImageList_ReadAdvData", 0),
		/*54*/ imports.NewTable("CustomImageList_ReadData", 0),
		/*55*/ imports.NewTable("CustomImageList_RegisterChanges", 0),
		/*56*/ imports.NewTable("CustomImageList_Replace", 0),
		/*57*/ imports.NewTable("CustomImageList_ReplaceIcon", 0),
		/*58*/ imports.NewTable("CustomImageList_ReplaceMasked", 0),
		/*59*/ imports.NewTable("CustomImageList_ReplaceSlice", 0),
		/*60*/ imports.NewTable("CustomImageList_ReplaceSliceCentered", 0),
		/*61*/ imports.NewTable("CustomImageList_Resolution", 0),
		/*62*/ imports.NewTable("CustomImageList_ResolutionByIndex", 0),
		/*63*/ imports.NewTable("CustomImageList_ResolutionCount", 0),
		/*64*/ imports.NewTable("CustomImageList_ResolutionForPPI", 0),
		/*65*/ imports.NewTable("CustomImageList_Resolutions", 0),
		/*66*/ imports.NewTable("CustomImageList_ResolutionsDesc", 0),
		/*67*/ imports.NewTable("CustomImageList_Scaled", 0),
		/*68*/ imports.NewTable("CustomImageList_SetOnChange", 0),
		/*69*/ imports.NewTable("CustomImageList_SetOnGetWidthForPPI", 0),
		/*70*/ imports.NewTable("CustomImageList_ShareImages", 0),
		/*71*/ imports.NewTable("CustomImageList_SizeForPPI", 0),
		/*72*/ imports.NewTable("CustomImageList_StretchDraw", 0),
		/*73*/ imports.NewTable("CustomImageList_UnRegisterChanges", 0),
		/*74*/ imports.NewTable("CustomImageList_Width", 0),
		/*75*/ imports.NewTable("CustomImageList_WidthForPPI", 0),
		/*76*/ imports.NewTable("CustomImageList_WriteAdvData", 0),
		/*77*/ imports.NewTable("CustomImageList_WriteData", 0),
	}
)

func customImageListImportAPI() *imports.Imports {
	if customImageListImport == nil {
		customImageListImport = NewDefaultImports()
		customImageListImport.SetImportTable(customImageListImportTables)
		customImageListImportTables = nil
	}
	return customImageListImport
}
