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

// ICustomImageListResolution Parent: ILCLReferenceComponent
type ICustomImageListResolution interface {
	ILCLReferenceComponent
	GetHotSpot() types.TPoint                                                                                                                                                                                                                           // function
	FillDescription(outDesc *IRawImageDescriptionWrap)                                                                                                                                                                                                  // procedure
	GetBitmapWithIntCustomBitmap(index int32, image ICustomBitmap)                                                                                                                                                                                      // procedure
	GetBitmapWithIntCustomBitmapGraphicsDrawEffect(index int32, image ICustomBitmap, effect types.TGraphicsDrawEffect)                                                                                                                                  // procedure
	GetIconWithIntIconGraphicsDrawEffect(index int32, image IIcon, effect types.TGraphicsDrawEffect)                                                                                                                                                    // procedure
	GetIconWithIntIcon(index int32, image IIcon)                                                                                                                                                                                                        // procedure
	GetFullBitmap(image ICustomBitmap, effect types.TGraphicsDrawEffect)                                                                                                                                                                                // procedure
	GetRawImage(index int32, outImage *IRawImageWrap)                                                                                                                                                                                                   // procedure
	DrawWithCanvasIntX3Bool(canvas ICanvas, X int32, Y int32, index int32, enabled bool)                                                                                                                                                                // procedure
	DrawWithCanvasIntX3GraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, drawEffect types.TGraphicsDrawEffect)                                                                                                                          // procedure
	DrawWithCanvasIntX3DrawingStyleImageTypeBool(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, enabled bool)                                                                             // procedure
	DrawWithCanvasIntX3DrawingStyleImageTypeGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect)                                       // procedure
	StretchDraw(canvas ICanvas, index int32, rect types.TRect, enabled bool)                                                                                                                                                                            // procedure
	DrawOverlayWithCanvasIntX3OverlayBool(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, enabled bool)                                                                                                                          // procedure
	DrawOverlayWithCanvasIntX3OverlayGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawEffect types.TGraphicsDrawEffect)                                                                                    // procedure
	DrawOverlayWithCanvasIntX3OverlayDrawingStyleImageTypeGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect) // procedure
	ImageList() ICustomImageList                                                                                                                                                                                                                        // property ImageList Getter
	Width() int32                                                                                                                                                                                                                                       // property Width Getter
	Height() int32                                                                                                                                                                                                                                      // property Height Getter
	Count() int32                                                                                                                                                                                                                                       // property Count Getter
	AutoCreatedInDesignTime() bool                                                                                                                                                                                                                      // property AutoCreatedInDesignTime Getter
	SetAutoCreatedInDesignTime(value bool)                                                                                                                                                                                                              // property AutoCreatedInDesignTime Setter
}

type TCustomImageListResolution struct {
	TLCLReferenceComponent
}

func (m *TCustomImageListResolution) GetHotSpot() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomImageListResolution) FillDescription(outDesc *IRawImageDescriptionWrap) {
	if !m.IsValid() {
		return
	}
	var descPtr uintptr
	customImageListResolutionAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&descPtr)))
	*outDesc = AsRawImageDescriptionWrap(descPtr)
}

func (m *TCustomImageListResolution) GetBitmapWithIntCustomBitmap(index int32, image ICustomBitmap) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(3, m.Instance(), uintptr(index), base.GetObjectUintptr(image))
}

func (m *TCustomImageListResolution) GetBitmapWithIntCustomBitmapGraphicsDrawEffect(index int32, image ICustomBitmap, effect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(4, m.Instance(), uintptr(index), base.GetObjectUintptr(image), uintptr(effect))
}

func (m *TCustomImageListResolution) GetIconWithIntIconGraphicsDrawEffect(index int32, image IIcon, effect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(5, m.Instance(), uintptr(index), base.GetObjectUintptr(image), uintptr(effect))
}

func (m *TCustomImageListResolution) GetIconWithIntIcon(index int32, image IIcon) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(6, m.Instance(), uintptr(index), base.GetObjectUintptr(image))
}

func (m *TCustomImageListResolution) GetFullBitmap(image ICustomBitmap, effect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(image), uintptr(effect))
}

func (m *TCustomImageListResolution) GetRawImage(index int32, outImage *IRawImageWrap) {
	if !m.IsValid() {
		return
	}
	var imagePtr uintptr
	customImageListResolutionAPI().SysCallN(8, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&imagePtr)))
	*outImage = AsRawImageWrap(imagePtr)
}

func (m *TCustomImageListResolution) DrawWithCanvasIntX3Bool(canvas ICanvas, X int32, Y int32, index int32, enabled bool) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), api.PasBool(enabled))
}

func (m *TCustomImageListResolution) DrawWithCanvasIntX3GraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(drawEffect))
}

func (m *TCustomImageListResolution) DrawWithCanvasIntX3DrawingStyleImageTypeBool(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, enabled bool) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(drawingStyle), uintptr(imageType), api.PasBool(enabled))
}

func (m *TCustomImageListResolution) DrawWithCanvasIntX3DrawingStyleImageTypeGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(drawingStyle), uintptr(imageType), uintptr(drawEffect))
}

func (m *TCustomImageListResolution) StretchDraw(canvas ICanvas, index int32, rect types.TRect, enabled bool) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(canvas), uintptr(index), uintptr(base.UnsafePointer(&rect)), api.PasBool(enabled))
}

func (m *TCustomImageListResolution) DrawOverlayWithCanvasIntX3OverlayBool(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, enabled bool) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(overlay), api.PasBool(enabled))
}

func (m *TCustomImageListResolution) DrawOverlayWithCanvasIntX3OverlayGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(overlay), uintptr(drawEffect))
}

func (m *TCustomImageListResolution) DrawOverlayWithCanvasIntX3OverlayDrawingStyleImageTypeGraphicsDrawEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(overlay), uintptr(drawingStyle), uintptr(imageType), uintptr(drawEffect))
}

func (m *TCustomImageListResolution) ImageList() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customImageListResolutionAPI().SysCallN(17, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomImageListResolution) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListResolutionAPI().SysCallN(18, m.Instance())
	return int32(r)
}

func (m *TCustomImageListResolution) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListResolutionAPI().SysCallN(19, m.Instance())
	return int32(r)
}

func (m *TCustomImageListResolution) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customImageListResolutionAPI().SysCallN(20, m.Instance())
	return int32(r)
}

func (m *TCustomImageListResolution) AutoCreatedInDesignTime() bool {
	if !m.IsValid() {
		return false
	}
	r := customImageListResolutionAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomImageListResolution) SetAutoCreatedInDesignTime(value bool) {
	if !m.IsValid() {
		return
	}
	customImageListResolutionAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

// NewCustomImageListResolution class constructor
func NewCustomImageListResolution(owner IComponent) ICustomImageListResolution {
	r := customImageListResolutionAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomImageListResolution(r)
}

func TCustomImageListResolutionClass() types.TClass {
	r := customImageListResolutionAPI().SysCallN(22)
	return types.TClass(r)
}

var (
	customImageListResolutionOnce   base.Once
	customImageListResolutionImport *imports.Imports = nil
)

func customImageListResolutionAPI() *imports.Imports {
	customImageListResolutionOnce.Do(func() {
		customImageListResolutionImport = api.NewDefaultImports()
		customImageListResolutionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomImageListResolution_Create", 0), // constructor NewCustomImageListResolution
			/* 1 */ imports.NewTable("TCustomImageListResolution_GetHotSpot", 0), // function GetHotSpot
			/* 2 */ imports.NewTable("TCustomImageListResolution_FillDescription", 0), // procedure FillDescription
			/* 3 */ imports.NewTable("TCustomImageListResolution_GetBitmapWithIntCustomBitmap", 0), // procedure GetBitmapWithIntCustomBitmap
			/* 4 */ imports.NewTable("TCustomImageListResolution_GetBitmapWithIntCustomBitmapGraphicsDrawEffect", 0), // procedure GetBitmapWithIntCustomBitmapGraphicsDrawEffect
			/* 5 */ imports.NewTable("TCustomImageListResolution_GetIconWithIntIconGraphicsDrawEffect", 0), // procedure GetIconWithIntIconGraphicsDrawEffect
			/* 6 */ imports.NewTable("TCustomImageListResolution_GetIconWithIntIcon", 0), // procedure GetIconWithIntIcon
			/* 7 */ imports.NewTable("TCustomImageListResolution_GetFullBitmap", 0), // procedure GetFullBitmap
			/* 8 */ imports.NewTable("TCustomImageListResolution_GetRawImage", 0), // procedure GetRawImage
			/* 9 */ imports.NewTable("TCustomImageListResolution_DrawWithCanvasIntX3Bool", 0), // procedure DrawWithCanvasIntX3Bool
			/* 10 */ imports.NewTable("TCustomImageListResolution_DrawWithCanvasIntX3GraphicsDrawEffect", 0), // procedure DrawWithCanvasIntX3GraphicsDrawEffect
			/* 11 */ imports.NewTable("TCustomImageListResolution_DrawWithCanvasIntX3DrawingStyleImageTypeBool", 0), // procedure DrawWithCanvasIntX3DrawingStyleImageTypeBool
			/* 12 */ imports.NewTable("TCustomImageListResolution_DrawWithCanvasIntX3DrawingStyleImageTypeGraphicsDrawEffect", 0), // procedure DrawWithCanvasIntX3DrawingStyleImageTypeGraphicsDrawEffect
			/* 13 */ imports.NewTable("TCustomImageListResolution_StretchDraw", 0), // procedure StretchDraw
			/* 14 */ imports.NewTable("TCustomImageListResolution_DrawOverlayWithCanvasIntX3OverlayBool", 0), // procedure DrawOverlayWithCanvasIntX3OverlayBool
			/* 15 */ imports.NewTable("TCustomImageListResolution_DrawOverlayWithCanvasIntX3OverlayGraphicsDrawEffect", 0), // procedure DrawOverlayWithCanvasIntX3OverlayGraphicsDrawEffect
			/* 16 */ imports.NewTable("TCustomImageListResolution_DrawOverlayWithCanvasIntX3OverlayDrawingStyleImageTypeGraphicsDrawEffect", 0), // procedure DrawOverlayWithCanvasIntX3OverlayDrawingStyleImageTypeGraphicsDrawEffect
			/* 17 */ imports.NewTable("TCustomImageListResolution_ImageList", 0), // property ImageList
			/* 18 */ imports.NewTable("TCustomImageListResolution_Width", 0), // property Width
			/* 19 */ imports.NewTable("TCustomImageListResolution_Height", 0), // property Height
			/* 20 */ imports.NewTable("TCustomImageListResolution_Count", 0), // property Count
			/* 21 */ imports.NewTable("TCustomImageListResolution_AutoCreatedInDesignTime", 0), // property AutoCreatedInDesignTime
			/* 22 */ imports.NewTable("TCustomImageListResolution_TClass", 0), // function TCustomImageListResolutionClass
		}
	})
	return customImageListResolutionImport
}
