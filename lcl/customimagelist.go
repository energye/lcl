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

// ICustomImageList Parent: ILCLComponent
type ICustomImageList interface {
	ILCLComponent
	Add(image ICustomBitmap, mask ICustomBitmap) int32                                                                                                                                                                                                  // function
	AddSliced(image ICustomBitmap, horizontalCount int32, verticalCount int32) int32                                                                                                                                                                    // function
	AddSlice(image ICustomBitmap, imageRect types.TRect) int32                                                                                                                                                                                          // function
	AddSliceCentered(image ICustomBitmap) int32                                                                                                                                                                                                         // function
	AddIcon(image ICustomIcon) int32                                                                                                                                                                                                                    // function
	AddMasked(image IBitmap, maskColor types.TColor) int32                                                                                                                                                                                              // function
	AddLazarusResource(resourceName string, maskColor types.TColor) int32                                                                                                                                                                               // function
	AddResourceName(instance types.TLCLHandle, resourceName string, maskColor types.TColor) int32                                                                                                                                                       // function
	GetHotSpot() types.TPoint                                                                                                                                                                                                                           // function
	FindResolution(imageWidth int32, outResolution *ICustomImageListResolution) bool                                                                                                                                                                    // function
	Resolutions() ICustomImageListResolutionEnumerator                                                                                                                                                                                                  // function
	ResolutionsDesc() ICustomImageListResolutionEnumerator                                                                                                                                                                                              // function
	AssignTo(dest IPersistent)                                                                                                                                                                                                                          // procedure
	WriteData(stream IStream)                                                                                                                                                                                                                           // procedure
	ReadData(stream IStream)                                                                                                                                                                                                                            // procedure
	WriteAdvData(stream IStream)                                                                                                                                                                                                                        // procedure
	ReadAdvData(stream IStream)                                                                                                                                                                                                                         // procedure
	BeginUpdate()                                                                                                                                                                                                                                       // procedure
	EndUpdate()                                                                                                                                                                                                                                         // procedure
	AddImages(value ICustomImageList)                                                                                                                                                                                                                   // procedure
	Change()                                                                                                                                                                                                                                            // procedure
	Clear()                                                                                                                                                                                                                                             // procedure
	Delete(index int32)                                                                                                                                                                                                                                 // procedure
	DrawWithCanvasIntX3Bool(canvas ICanvas, X int32, Y int32, index int32, enabled bool)                                                                                                                                                                // procedure
	DrawWithCanvasIntX3GraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, drawEffect types.TGraphicsDrawEffect)                                                                                                                          // procedure
	DrawWithCanvasIntX3DrawingStyleImageTypeBool(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, enabled bool)                                                                             // procedure
	DrawWithCanvasIntX3DrawingStyleImageTypeGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect)                                       // procedure
	DrawForPPIWithCanvasIntX5DoubleBool(canvas ICanvas, X int32, Y int32, index int32, imageWidthAt96PPI int32, targetPPI int32, canvasFactor float64, enabled bool)                                                                                    // procedure
	DrawForPPIWithCanvasIntX5DoubleGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, imageWidthAt96PPI int32, targetPPI int32, canvasFactor float64, drawEffect types.TGraphicsDrawEffect)                                              // procedure
	DrawOverlayWithCanvasIntX3OverlayBool(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, enabled bool)                                                                                                                          // procedure
	DrawOverlayWithCanvasIntX3OverlayGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawEffect types.TGraphicsDrawEffect)                                                                                    // procedure
	DrawOverlayWithCanvasIntX3OverlayDrawingStyleImageTypeGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect) // procedure
	GetBitmapWithIntCustomBitmap(index int32, image ICustomBitmap)                                                                                                                                                                                      // procedure
	GetBitmapWithIntCustomBitmapGraphicsDrawEffect(index int32, image ICustomBitmap, effect types.TGraphicsDrawEffect)                                                                                                                                  // procedure
	GetFullBitmap(image ICustomBitmap, effect types.TGraphicsDrawEffect)                                                                                                                                                                                // procedure
	GetFullRawImage(outImage *IRawImageWrap)                                                                                                                                                                                                            // procedure
	GetIconWithIntIconGraphicsDrawEffect(index int32, image IIcon, effect types.TGraphicsDrawEffect)                                                                                                                                                    // procedure
	GetIconWithIntIcon(index int32, image IIcon)                                                                                                                                                                                                        // procedure
	GetRawImage(index int32, outImage *IRawImageWrap)                                                                                                                                                                                                   // procedure
	Insert(index int32, image ICustomBitmap, mask ICustomBitmap)                                                                                                                                                                                        // procedure
	InsertIcon(index int32, icon ICustomIcon)                                                                                                                                                                                                           // procedure
	InsertMasked(index int32, image ICustomBitmap, maskColor types.TColor)                                                                                                                                                                              // procedure
	Move(curIndex int32, newIndex int32)                                                                                                                                                                                                                // procedure
	Overlay(index int32, overlay types.TOverlay)                                                                                                                                                                                                        // procedure
	Replace(index int32, image ICustomBitmap, mask ICustomBitmap, allResolutions bool)                                                                                                                                                                  // procedure
	ReplaceSlice(index int32, image ICustomBitmap, imageRect types.TRect, allResolutions bool)                                                                                                                                                          // procedure
	ReplaceSliceCentered(index int32, imageWidth int32, image ICustomBitmap, allResolutions bool)                                                                                                                                                       // procedure
	ReplaceIcon(index int32, icon ICustomIcon)                                                                                                                                                                                                          // procedure
	ReplaceMasked(index int32, newImage ICustomBitmap, maskColor types.TColor, allResolutions bool)                                                                                                                                                     // procedure
	RegisterChanges(value IChangeLink)                                                                                                                                                                                                                  // procedure
	StretchDraw(canvas ICanvas, index int32, rect types.TRect, enabled bool)                                                                                                                                                                            // procedure
	UnRegisterChanges(value IChangeLink)                                                                                                                                                                                                                // procedure
	DeleteResolution(width int32)                                                                                                                                                                                                                       // procedure
	HasOverlays() bool                                                                                                                                                                                                                                  // property HasOverlays Getter
	AllocBy() int32                                                                                                                                                                                                                                     // property AllocBy Getter
	SetAllocBy(value int32)                                                                                                                                                                                                                             // property AllocBy Setter
	BlendColor() types.TColor                                                                                                                                                                                                                           // property BlendColor Getter
	SetBlendColor(value types.TColor)                                                                                                                                                                                                                   // property BlendColor Setter
	BkColor() types.TColor                                                                                                                                                                                                                              // property BkColor Getter
	SetBkColor(value types.TColor)                                                                                                                                                                                                                      // property BkColor Setter
	Count() int32                                                                                                                                                                                                                                       // property Count Getter
	DrawingStyle() types.TDrawingStyle                                                                                                                                                                                                                  // property DrawingStyle Getter
	SetDrawingStyle(value types.TDrawingStyle)                                                                                                                                                                                                          // property DrawingStyle Setter
	Handle() types.TLCLHandle                                                                                                                                                                                                                           // property Handle Getter
	Height() int32                                                                                                                                                                                                                                      // property Height Getter
	SetHeight(value int32)                                                                                                                                                                                                                              // property Height Setter
	HeightForPPI(imageWidth int32, pPI int32) int32                                                                                                                                                                                                     // property HeightForPPI Getter
	HeightForWidth(width int32) int32                                                                                                                                                                                                                   // property HeightForWidth Getter
	Width() int32                                                                                                                                                                                                                                       // property Width Getter
	SetWidth(value int32)                                                                                                                                                                                                                               // property Width Setter
	WidthForPPI(imageWidth int32, pPI int32) int32                                                                                                                                                                                                      // property WidthForPPI Getter
	SizeForPPI(imageWidth int32, pPI int32) types.TSize                                                                                                                                                                                                 // property SizeForPPI Getter
	Masked() bool                                                                                                                                                                                                                                       // property Masked Getter
	SetMasked(value bool)                                                                                                                                                                                                                               // property Masked Setter
	ResolutionWithIntToCustomImageListResolution(imageWidth int32) ICustomImageListResolution                                                                                                                                                           // property Resolution Getter
	ResolutionByIndex(index int32) ICustomImageListResolution                                                                                                                                                                                           // property ResolutionByIndex Getter
	ResolutionForPPI(imageWidth int32, pPI int32, canvasScaleFactor float64) IScaledImageListResolutionWrap                                                                                                                                             // property ResolutionForPPI Getter
	ResolutionCount() int32                                                                                                                                                                                                                             // property ResolutionCount Getter
	Scaled() bool                                                                                                                                                                                                                                       // property Scaled Getter
	SetScaled(value bool)                                                                                                                                                                                                                               // property Scaled Setter
	ShareImages() bool                                                                                                                                                                                                                                  // property ShareImages Getter
	SetShareImages(value bool)                                                                                                                                                                                                                          // property ShareImages Setter
	ImageType() types.TImageType                                                                                                                                                                                                                        // property ImageType Getter
	SetImageType(value types.TImageType)                                                                                                                                                                                                                // property ImageType Setter
	SetOnChange(fn TNotifyEvent)                                                                                                                                                                                                                        // property event
	SetOnGetWidthForPPI(fn TCustomImageListGetWidthForPPI)                                                                                                                                                                                              // property event
}

type TCustomImageList struct {
	TLCLComponent
}

func (m *TCustomImageList) Add(image ICustomBitmap, mask ICustomBitmap) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(image), base.GetObjectUintptr(mask))
	return int32(r)
}

func (m *TCustomImageList) AddSliced(image ICustomBitmap, horizontalCount int32, verticalCount int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(image), uintptr(horizontalCount), uintptr(verticalCount))
	return int32(r)
}

func (m *TCustomImageList) AddSlice(image ICustomBitmap, imageRect types.TRect) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(4, m.Instance(), base.GetObjectUintptr(image), uintptr(base.UnsafePointer(&imageRect)))
	return int32(r)
}

func (m *TCustomImageList) AddSliceCentered(image ICustomBitmap) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(image))
	return int32(r)
}

func (m *TCustomImageList) AddIcon(image ICustomIcon) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(image))
	return int32(r)
}

func (m *TCustomImageList) AddMasked(image IBitmap, maskColor types.TColor) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(image), uintptr(maskColor))
	return int32(r)
}

func (m *TCustomImageList) AddLazarusResource(resourceName string, maskColor types.TColor) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(8, m.Instance(), api.PasStr(resourceName), uintptr(maskColor))
	return int32(r)
}

func (m *TCustomImageList) AddResourceName(instance types.TLCLHandle, resourceName string, maskColor types.TColor) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(9, m.Instance(), uintptr(instance), api.PasStr(resourceName), uintptr(maskColor))
	return int32(r)
}

func (m *TCustomImageList) GetHotSpot() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(10, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomImageList) FindResolution(imageWidth int32, outResolution *ICustomImageListResolution) bool {
	if !m.IsValid() {
		return false
	}
	var resolutionPtr uintptr
	r := customImageListAPI().SysCallN(11, m.Instance(), uintptr(imageWidth), uintptr(base.UnsafePointer(&resolutionPtr)))
	*outResolution = AsCustomImageListResolution(resolutionPtr)
	return api.GoBool(r)
}

func (m *TCustomImageList) Resolutions() ICustomImageListResolutionEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := customImageListAPI().SysCallN(12, m.Instance())
	return AsCustomImageListResolutionEnumerator(r)
}

func (m *TCustomImageList) ResolutionsDesc() ICustomImageListResolutionEnumerator {
	if !m.IsValid() {
		return nil
	}
	r := customImageListAPI().SysCallN(13, m.Instance())
	return AsCustomImageListResolutionEnumerator(r)
}

func (m *TCustomImageList) AssignTo(dest IPersistent) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(17, m.Instance(), base.GetObjectUintptr(dest))
}

func (m *TCustomImageList) WriteData(stream IStream) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(18, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCustomImageList) ReadData(stream IStream) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCustomImageList) WriteAdvData(stream IStream) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCustomImageList) ReadAdvData(stream IStream) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TCustomImageList) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(22, m.Instance())
}

func (m *TCustomImageList) EndUpdate() {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(23, m.Instance())
}

func (m *TCustomImageList) AddImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(24, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomImageList) Change() {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(25, m.Instance())
}

func (m *TCustomImageList) Clear() {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(26, m.Instance())
}

func (m *TCustomImageList) Delete(index int32) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(27, m.Instance(), uintptr(index))
}

func (m *TCustomImageList) DrawWithCanvasIntX3Bool(canvas ICanvas, X int32, Y int32, index int32, enabled bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(28, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), api.PasBool(enabled))
}

func (m *TCustomImageList) DrawWithCanvasIntX3GraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(29, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(drawEffect))
}

func (m *TCustomImageList) DrawWithCanvasIntX3DrawingStyleImageTypeBool(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, enabled bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(30, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(drawingStyle), uintptr(imageType), api.PasBool(enabled))
}

func (m *TCustomImageList) DrawWithCanvasIntX3DrawingStyleImageTypeGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(31, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(drawingStyle), uintptr(imageType), uintptr(drawEffect))
}

func (m *TCustomImageList) DrawForPPIWithCanvasIntX5DoubleBool(canvas ICanvas, X int32, Y int32, index int32, imageWidthAt96PPI int32, targetPPI int32, canvasFactor float64, enabled bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(32, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(imageWidthAt96PPI), uintptr(targetPPI), uintptr(base.UnsafePointer(&canvasFactor)), api.PasBool(enabled))
}

func (m *TCustomImageList) DrawForPPIWithCanvasIntX5DoubleGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, imageWidthAt96PPI int32, targetPPI int32, canvasFactor float64, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(33, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(imageWidthAt96PPI), uintptr(targetPPI), uintptr(base.UnsafePointer(&canvasFactor)), uintptr(drawEffect))
}

func (m *TCustomImageList) DrawOverlayWithCanvasIntX3OverlayBool(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, enabled bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(34, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(overlay), api.PasBool(enabled))
}

func (m *TCustomImageList) DrawOverlayWithCanvasIntX3OverlayGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(35, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(overlay), uintptr(drawEffect))
}

func (m *TCustomImageList) DrawOverlayWithCanvasIntX3OverlayDrawingStyleImageTypeGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(36, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(overlay), uintptr(drawingStyle), uintptr(imageType), uintptr(drawEffect))
}

func (m *TCustomImageList) GetBitmapWithIntCustomBitmap(index int32, image ICustomBitmap) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(37, m.Instance(), uintptr(index), base.GetObjectUintptr(image))
}

func (m *TCustomImageList) GetBitmapWithIntCustomBitmapGraphicsDrawEffect(index int32, image ICustomBitmap, effect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(38, m.Instance(), uintptr(index), base.GetObjectUintptr(image), uintptr(effect))
}

func (m *TCustomImageList) GetFullBitmap(image ICustomBitmap, effect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(39, m.Instance(), base.GetObjectUintptr(image), uintptr(effect))
}

func (m *TCustomImageList) GetFullRawImage(outImage *IRawImageWrap) {
	if !m.IsValid() {
		return
	}
	var imagePtr uintptr
	customImageListAPI().SysCallN(40, m.Instance(), uintptr(base.UnsafePointer(&imagePtr)))
	*outImage = AsRawImageWrap(imagePtr)
}

func (m *TCustomImageList) GetIconWithIntIconGraphicsDrawEffect(index int32, image IIcon, effect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(41, m.Instance(), uintptr(index), base.GetObjectUintptr(image), uintptr(effect))
}

func (m *TCustomImageList) GetIconWithIntIcon(index int32, image IIcon) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(42, m.Instance(), uintptr(index), base.GetObjectUintptr(image))
}

func (m *TCustomImageList) GetRawImage(index int32, outImage *IRawImageWrap) {
	if !m.IsValid() {
		return
	}
	var imagePtr uintptr
	customImageListAPI().SysCallN(43, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&imagePtr)))
	*outImage = AsRawImageWrap(imagePtr)
}

func (m *TCustomImageList) Insert(index int32, image ICustomBitmap, mask ICustomBitmap) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(44, m.Instance(), uintptr(index), base.GetObjectUintptr(image), base.GetObjectUintptr(mask))
}

func (m *TCustomImageList) InsertIcon(index int32, icon ICustomIcon) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(45, m.Instance(), uintptr(index), base.GetObjectUintptr(icon))
}

func (m *TCustomImageList) InsertMasked(index int32, image ICustomBitmap, maskColor types.TColor) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(46, m.Instance(), uintptr(index), base.GetObjectUintptr(image), uintptr(maskColor))
}

func (m *TCustomImageList) Move(curIndex int32, newIndex int32) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(47, m.Instance(), uintptr(curIndex), uintptr(newIndex))
}

func (m *TCustomImageList) Overlay(index int32, overlay types.TOverlay) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(48, m.Instance(), uintptr(index), uintptr(overlay))
}

func (m *TCustomImageList) Replace(index int32, image ICustomBitmap, mask ICustomBitmap, allResolutions bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(49, m.Instance(), uintptr(index), base.GetObjectUintptr(image), base.GetObjectUintptr(mask), api.PasBool(allResolutions))
}

func (m *TCustomImageList) ReplaceSlice(index int32, image ICustomBitmap, imageRect types.TRect, allResolutions bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(50, m.Instance(), uintptr(index), base.GetObjectUintptr(image), uintptr(base.UnsafePointer(&imageRect)), api.PasBool(allResolutions))
}

func (m *TCustomImageList) ReplaceSliceCentered(index int32, imageWidth int32, image ICustomBitmap, allResolutions bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(51, m.Instance(), uintptr(index), uintptr(imageWidth), base.GetObjectUintptr(image), api.PasBool(allResolutions))
}

func (m *TCustomImageList) ReplaceIcon(index int32, icon ICustomIcon) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(52, m.Instance(), uintptr(index), base.GetObjectUintptr(icon))
}

func (m *TCustomImageList) ReplaceMasked(index int32, newImage ICustomBitmap, maskColor types.TColor, allResolutions bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(53, m.Instance(), uintptr(index), base.GetObjectUintptr(newImage), uintptr(maskColor), api.PasBool(allResolutions))
}

func (m *TCustomImageList) RegisterChanges(value IChangeLink) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(54, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomImageList) StretchDraw(canvas ICanvas, index int32, rect types.TRect, enabled bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(55, m.Instance(), base.GetObjectUintptr(canvas), uintptr(index), uintptr(base.UnsafePointer(&rect)), api.PasBool(enabled))
}

func (m *TCustomImageList) UnRegisterChanges(value IChangeLink) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(56, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomImageList) DeleteResolution(width int32) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(57, m.Instance(), uintptr(width))
}

func (m *TCustomImageList) HasOverlays() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageListAPI().SysCallN(58, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImageList) AllocBy() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(59, 0, m.Instance())
	return int32(r)
}

func (m *TCustomImageList) SetAllocBy(value int32) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(59, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImageList) BlendColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(60, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomImageList) SetBlendColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(60, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImageList) BkColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(61, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCustomImageList) SetBkColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(61, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImageList) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(62, m.Instance())
	return int32(r)
}

func (m *TCustomImageList) DrawingStyle() types.TDrawingStyle {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(63, 0, m.Instance())
	return types.TDrawingStyle(r)
}

func (m *TCustomImageList) SetDrawingStyle(value types.TDrawingStyle) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(63, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImageList) Handle() types.TLCLHandle {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(64, m.Instance())
	return types.TLCLHandle(r)
}

func (m *TCustomImageList) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(65, 0, m.Instance())
	return int32(r)
}

func (m *TCustomImageList) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(65, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImageList) HeightForPPI(imageWidth int32, pPI int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(66, m.Instance(), uintptr(imageWidth), uintptr(pPI))
	return int32(r)
}

func (m *TCustomImageList) HeightForWidth(width int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(67, m.Instance(), uintptr(width))
	return int32(r)
}

func (m *TCustomImageList) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(68, 0, m.Instance())
	return int32(r)
}

func (m *TCustomImageList) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(68, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImageList) WidthForPPI(imageWidth int32, pPI int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(69, m.Instance(), uintptr(imageWidth), uintptr(pPI))
	return int32(r)
}

func (m *TCustomImageList) SizeForPPI(imageWidth int32, pPI int32) (result types.TSize) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(70, m.Instance(), uintptr(imageWidth), uintptr(pPI), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomImageList) Masked() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageListAPI().SysCallN(71, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImageList) SetMasked(value bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(71, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImageList) ResolutionWithIntToCustomImageListResolution(imageWidth int32) ICustomImageListResolution {
	if !m.IsValid() {
		return nil
	}
	r := customImageListAPI().SysCallN(72, m.Instance(), uintptr(imageWidth))
	return AsCustomImageListResolution(r)
}

func (m *TCustomImageList) ResolutionByIndex(index int32) ICustomImageListResolution {
	if !m.IsValid() {
		return nil
	}
	r := customImageListAPI().SysCallN(73, m.Instance(), uintptr(index))
	return AsCustomImageListResolution(r)
}

func (m *TCustomImageList) ResolutionForPPI(imageWidth int32, pPI int32, canvasScaleFactor float64) IScaledImageListResolutionWrap {
	if !m.IsValid() {
		return nil
	}
	r := customImageListAPI().SysCallN(74, m.Instance(), uintptr(imageWidth), uintptr(pPI), uintptr(base.UnsafePointer(&canvasScaleFactor)))
	return AsScaledImageListResolutionWrap(r)
}

func (m *TCustomImageList) ResolutionCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(75, m.Instance())
	return int32(r)
}

func (m *TCustomImageList) Scaled() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageListAPI().SysCallN(76, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImageList) SetScaled(value bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(76, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImageList) ShareImages() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageListAPI().SysCallN(77, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImageList) SetShareImages(value bool) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(77, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomImageList) ImageType() types.TImageType {
	if !m.IsValid() {
		return 0
	}
	r := customImageListAPI().SysCallN(78, 0, m.Instance())
	return types.TImageType(r)
}

func (m *TCustomImageList) SetImageType(value types.TImageType) {
	if !m.IsValid() {
		return
	}
	customImageListAPI().SysCallN(78, 1, m.Instance(), uintptr(value))
}

func (m *TCustomImageList) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 79, customImageListAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomImageList) SetOnGetWidthForPPI(fn TCustomImageListGetWidthForPPI) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomImageListGetWidthForPPI(fn)
	base.SetEvent(m, 80, customImageListAPI(), api.MakeEventDataPtr(cb))
}

// CustomImageList  is static instance
var CustomImageList _CustomImageListClass

// _CustomImageListClass is class type defined by TCustomImageList
type _CustomImageListClass uintptr

func (_CustomImageListClass) ScaleImageWithCustomBitmapX2IntX2RGBAQuadArray(bitmap ICustomBitmap, mask ICustomBitmap, targetWidth int32, targetHeight int32, data *IRGBAQuadArray) {
	var dataPtr uintptr
	var dataCountPtr uintptr
	customImageListAPI().SysCallN(14, base.GetObjectUintptr(bitmap), base.GetObjectUintptr(mask), uintptr(targetWidth), uintptr(targetHeight), uintptr(base.UnsafePointer(&dataPtr)), uintptr(base.UnsafePointer(&dataCountPtr)))
	*data = NewRGBAQuadArray(int(dataCountPtr), dataPtr)
}

func (_CustomImageListClass) ScaleImageWithCustomBitmapX2RectIntX2RGBAQuadArray(bitmap ICustomBitmap, mask ICustomBitmap, sourceRect types.TRect, targetWidth int32, targetHeight int32, data *IRGBAQuadArray) {
	var dataPtr uintptr
	var dataCountPtr uintptr
	customImageListAPI().SysCallN(15, base.GetObjectUintptr(bitmap), base.GetObjectUintptr(mask), uintptr(base.UnsafePointer(&sourceRect)), uintptr(targetWidth), uintptr(targetHeight), uintptr(base.UnsafePointer(&dataPtr)), uintptr(base.UnsafePointer(&dataCountPtr)))
	*data = NewRGBAQuadArray(int(dataCountPtr), dataPtr)
}

func (_CustomImageListClass) ScaleImageWithHBITMAPX2RectIntX2RGBAQuadArray(bitmap types.HBitmap, mask types.HBitmap, sourceRect types.TRect, targetWidth int32, targetHeight int32, data *IRGBAQuadArray) {
	var dataPtr uintptr
	var dataCountPtr uintptr
	customImageListAPI().SysCallN(16, uintptr(bitmap), uintptr(mask), uintptr(base.UnsafePointer(&sourceRect)), uintptr(targetWidth), uintptr(targetHeight), uintptr(base.UnsafePointer(&dataPtr)), uintptr(base.UnsafePointer(&dataCountPtr)))
	*data = NewRGBAQuadArray(int(dataCountPtr), dataPtr)
}

// NewCustomImageList class constructor
func NewCustomImageList(owner IComponent) ICustomImageList {
	r := customImageListAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomImageList(r)
}

// NewCustomImageListSize class constructor
func NewCustomImageListSize(width int32, height int32) ICustomImageList {
	r := customImageListAPI().SysCallN(1, uintptr(width), uintptr(height))
	return AsCustomImageList(r)
}

func TCustomImageListClass() types.TClass {
	r := customImageListAPI().SysCallN(81)
	return types.TClass(r)
}

var (
	customImageListOnce   base.Once
	customImageListImport *imports.Imports = nil
)

func customImageListAPI() *imports.Imports {
	customImageListOnce.Do(func() {
		customImageListImport = api.NewDefaultImports()
		customImageListImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomImageList_Create", 0), // constructor NewCustomImageList
			/* 1 */ imports.NewTable("TCustomImageList_CreateSize", 0), // constructor NewCustomImageListSize
			/* 2 */ imports.NewTable("TCustomImageList_Add", 0), // function Add
			/* 3 */ imports.NewTable("TCustomImageList_AddSliced", 0), // function AddSliced
			/* 4 */ imports.NewTable("TCustomImageList_AddSlice", 0), // function AddSlice
			/* 5 */ imports.NewTable("TCustomImageList_AddSliceCentered", 0), // function AddSliceCentered
			/* 6 */ imports.NewTable("TCustomImageList_AddIcon", 0), // function AddIcon
			/* 7 */ imports.NewTable("TCustomImageList_AddMasked", 0), // function AddMasked
			/* 8 */ imports.NewTable("TCustomImageList_AddLazarusResource", 0), // function AddLazarusResource
			/* 9 */ imports.NewTable("TCustomImageList_AddResourceName", 0), // function AddResourceName
			/* 10 */ imports.NewTable("TCustomImageList_GetHotSpot", 0), // function GetHotSpot
			/* 11 */ imports.NewTable("TCustomImageList_FindResolution", 0), // function FindResolution
			/* 12 */ imports.NewTable("TCustomImageList_Resolutions", 0), // function Resolutions
			/* 13 */ imports.NewTable("TCustomImageList_ResolutionsDesc", 0), // function ResolutionsDesc
			/* 14 */ imports.NewTable("TCustomImageList_ScaleImageWithCustomBitmapX2IntX2RGBAQuadArray", 0), // static procedure ScaleImageWithCustomBitmapX2IntX2RGBAQuadArray
			/* 15 */ imports.NewTable("TCustomImageList_ScaleImageWithCustomBitmapX2RectIntX2RGBAQuadArray", 0), // static procedure ScaleImageWithCustomBitmapX2RectIntX2RGBAQuadArray
			/* 16 */ imports.NewTable("TCustomImageList_ScaleImageWithHBITMAPX2RectIntX2RGBAQuadArray", 0), // static procedure ScaleImageWithHBITMAPX2RectIntX2RGBAQuadArray
			/* 17 */ imports.NewTable("TCustomImageList_AssignTo", 0), // procedure AssignTo
			/* 18 */ imports.NewTable("TCustomImageList_WriteData", 0), // procedure WriteData
			/* 19 */ imports.NewTable("TCustomImageList_ReadData", 0), // procedure ReadData
			/* 20 */ imports.NewTable("TCustomImageList_WriteAdvData", 0), // procedure WriteAdvData
			/* 21 */ imports.NewTable("TCustomImageList_ReadAdvData", 0), // procedure ReadAdvData
			/* 22 */ imports.NewTable("TCustomImageList_BeginUpdate", 0), // procedure BeginUpdate
			/* 23 */ imports.NewTable("TCustomImageList_EndUpdate", 0), // procedure EndUpdate
			/* 24 */ imports.NewTable("TCustomImageList_AddImages", 0), // procedure AddImages
			/* 25 */ imports.NewTable("TCustomImageList_Change", 0), // procedure Change
			/* 26 */ imports.NewTable("TCustomImageList_Clear", 0), // procedure Clear
			/* 27 */ imports.NewTable("TCustomImageList_Delete", 0), // procedure Delete
			/* 28 */ imports.NewTable("TCustomImageList_DrawWithCanvasIntX3Bool", 0), // procedure DrawWithCanvasIntX3Bool
			/* 29 */ imports.NewTable("TCustomImageList_DrawWithCanvasIntX3GraphicsDrawEffect", 0), // procedure DrawWithCanvasIntX3GraphicsDrawEffect
			/* 30 */ imports.NewTable("TCustomImageList_DrawWithCanvasIntX3DrawingStyleImageTypeBool", 0), // procedure DrawWithCanvasIntX3DrawingStyleImageTypeBool
			/* 31 */ imports.NewTable("TCustomImageList_DrawWithCanvasIntX3DrawingStyleImageTypeGraphicsDrawEffect", 0), // procedure DrawWithCanvasIntX3DrawingStyleImageTypeGraphicsDrawEffect
			/* 32 */ imports.NewTable("TCustomImageList_DrawForPPIWithCanvasIntX5DoubleBool", 0), // procedure DrawForPPIWithCanvasIntX5DoubleBool
			/* 33 */ imports.NewTable("TCustomImageList_DrawForPPIWithCanvasIntX5DoubleGraphicsDrawEffect", 0), // procedure DrawForPPIWithCanvasIntX5DoubleGraphicsDrawEffect
			/* 34 */ imports.NewTable("TCustomImageList_DrawOverlayWithCanvasIntX3OverlayBool", 0), // procedure DrawOverlayWithCanvasIntX3OverlayBool
			/* 35 */ imports.NewTable("TCustomImageList_DrawOverlayWithCanvasIntX3OverlayGraphicsDrawEffect", 0), // procedure DrawOverlayWithCanvasIntX3OverlayGraphicsDrawEffect
			/* 36 */ imports.NewTable("TCustomImageList_DrawOverlayWithCanvasIntX3OverlayDrawingStyleImageTypeGraphicsDrawEffect", 0), // procedure DrawOverlayWithCanvasIntX3OverlayDrawingStyleImageTypeGraphicsDrawEffect
			/* 37 */ imports.NewTable("TCustomImageList_GetBitmapWithIntCustomBitmap", 0), // procedure GetBitmapWithIntCustomBitmap
			/* 38 */ imports.NewTable("TCustomImageList_GetBitmapWithIntCustomBitmapGraphicsDrawEffect", 0), // procedure GetBitmapWithIntCustomBitmapGraphicsDrawEffect
			/* 39 */ imports.NewTable("TCustomImageList_GetFullBitmap", 0), // procedure GetFullBitmap
			/* 40 */ imports.NewTable("TCustomImageList_GetFullRawImage", 0), // procedure GetFullRawImage
			/* 41 */ imports.NewTable("TCustomImageList_GetIconWithIntIconGraphicsDrawEffect", 0), // procedure GetIconWithIntIconGraphicsDrawEffect
			/* 42 */ imports.NewTable("TCustomImageList_GetIconWithIntIcon", 0), // procedure GetIconWithIntIcon
			/* 43 */ imports.NewTable("TCustomImageList_GetRawImage", 0), // procedure GetRawImage
			/* 44 */ imports.NewTable("TCustomImageList_Insert", 0), // procedure Insert
			/* 45 */ imports.NewTable("TCustomImageList_InsertIcon", 0), // procedure InsertIcon
			/* 46 */ imports.NewTable("TCustomImageList_InsertMasked", 0), // procedure InsertMasked
			/* 47 */ imports.NewTable("TCustomImageList_Move", 0), // procedure Move
			/* 48 */ imports.NewTable("TCustomImageList_Overlay", 0), // procedure Overlay
			/* 49 */ imports.NewTable("TCustomImageList_Replace", 0), // procedure Replace
			/* 50 */ imports.NewTable("TCustomImageList_ReplaceSlice", 0), // procedure ReplaceSlice
			/* 51 */ imports.NewTable("TCustomImageList_ReplaceSliceCentered", 0), // procedure ReplaceSliceCentered
			/* 52 */ imports.NewTable("TCustomImageList_ReplaceIcon", 0), // procedure ReplaceIcon
			/* 53 */ imports.NewTable("TCustomImageList_ReplaceMasked", 0), // procedure ReplaceMasked
			/* 54 */ imports.NewTable("TCustomImageList_RegisterChanges", 0), // procedure RegisterChanges
			/* 55 */ imports.NewTable("TCustomImageList_StretchDraw", 0), // procedure StretchDraw
			/* 56 */ imports.NewTable("TCustomImageList_UnRegisterChanges", 0), // procedure UnRegisterChanges
			/* 57 */ imports.NewTable("TCustomImageList_DeleteResolution", 0), // procedure DeleteResolution
			/* 58 */ imports.NewTable("TCustomImageList_HasOverlays", 0), // property HasOverlays
			/* 59 */ imports.NewTable("TCustomImageList_AllocBy", 0), // property AllocBy
			/* 60 */ imports.NewTable("TCustomImageList_BlendColor", 0), // property BlendColor
			/* 61 */ imports.NewTable("TCustomImageList_BkColor", 0), // property BkColor
			/* 62 */ imports.NewTable("TCustomImageList_Count", 0), // property Count
			/* 63 */ imports.NewTable("TCustomImageList_DrawingStyle", 0), // property DrawingStyle
			/* 64 */ imports.NewTable("TCustomImageList_Handle", 0), // property Handle
			/* 65 */ imports.NewTable("TCustomImageList_Height", 0), // property Height
			/* 66 */ imports.NewTable("TCustomImageList_HeightForPPI", 0), // property HeightForPPI
			/* 67 */ imports.NewTable("TCustomImageList_HeightForWidth", 0), // property HeightForWidth
			/* 68 */ imports.NewTable("TCustomImageList_Width", 0), // property Width
			/* 69 */ imports.NewTable("TCustomImageList_WidthForPPI", 0), // property WidthForPPI
			/* 70 */ imports.NewTable("TCustomImageList_SizeForPPI", 0), // property SizeForPPI
			/* 71 */ imports.NewTable("TCustomImageList_Masked", 0), // property Masked
			/* 72 */ imports.NewTable("TCustomImageList_ResolutionWithIntToCustomImageListResolution", 0), // property ResolutionWithIntToCustomImageListResolution
			/* 73 */ imports.NewTable("TCustomImageList_ResolutionByIndex", 0), // property ResolutionByIndex
			/* 74 */ imports.NewTable("TCustomImageList_ResolutionForPPI", 0), // property ResolutionForPPI
			/* 75 */ imports.NewTable("TCustomImageList_ResolutionCount", 0), // property ResolutionCount
			/* 76 */ imports.NewTable("TCustomImageList_Scaled", 0), // property Scaled
			/* 77 */ imports.NewTable("TCustomImageList_ShareImages", 0), // property ShareImages
			/* 78 */ imports.NewTable("TCustomImageList_ImageType", 0), // property ImageType
			/* 79 */ imports.NewTable("TCustomImageList_OnChange", 0), // event OnChange
			/* 80 */ imports.NewTable("TCustomImageList_OnGetWidthForPPI", 0), // event OnGetWidthForPPI
			/* 81 */ imports.NewTable("TCustomImageList_TClass", 0), // function TCustomImageListClass
		}
	})
	return customImageListImport
}
