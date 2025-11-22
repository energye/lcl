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

// IScaledImageListResolutionWrap Parent: IObject
type IScaledImageListResolutionWrap interface {
	IObject
	Width() int32                                                                                                                                                                                             // function
	Height() int32                                                                                                                                                                                            // function
	Size() types.TSize                                                                                                                                                                                        // function
	Resolution() ICustomImageListResolution                                                                                                                                                                   // function
	Count() int32                                                                                                                                                                                             // function
	Valid() bool                                                                                                                                                                                              // function
	GetBitmapWithIntCustomBitmap(index int32, image ICustomBitmap)                                                                                                                                            // procedure
	GetBitmapWithIntCustomBitmapGraphicsDrawEffect(index int32, image ICustomBitmap, effect types.TGraphicsDrawEffect)                                                                                        // procedure
	DrawXY(canvas ICanvas, X int32, Y int32, index int32, enabled bool)                                                                                                                                       // procedure
	DrawXYEffect(canvas ICanvas, X int32, Y int32, index int32, drawEffect types.TGraphicsDrawEffect)                                                                                                         // procedure
	DrawXYStyleType(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, enabled bool)                                                                // procedure
	DrawXYStyleTypeEffect(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect)                                  // procedure
	StretchDraw(canvas ICanvas, index int32, rect types.TRect, enabled bool)                                                                                                                                  // procedure
	DrawOverlayXYOverlay(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, enabled bool)                                                                                                 // procedure
	DrawOverlayXYOverlayEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawEffect types.TGraphicsDrawEffect)                                                                   // procedure
	DrawOverlayXYOverlayTypeEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect) // procedure
}

type TScaledImageListResolutionWrap struct {
	TObject
}

func (m *TScaledImageListResolutionWrap) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := scaledImageListResolutionWrapAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TScaledImageListResolutionWrap) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := scaledImageListResolutionWrapAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TScaledImageListResolutionWrap) Size() (result types.TSize) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TScaledImageListResolutionWrap) Resolution() ICustomImageListResolution {
	if !m.IsValid() {
		return nil
	}
	r := scaledImageListResolutionWrapAPI().SysCallN(4, m.Instance())
	return AsCustomImageListResolution(r)
}

func (m *TScaledImageListResolutionWrap) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := scaledImageListResolutionWrapAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TScaledImageListResolutionWrap) Valid() bool {
	if !m.IsValid() {
		return false
	}
	r := scaledImageListResolutionWrapAPI().SysCallN(6, m.Instance())
	return api.GoBool(r)
}

func (m *TScaledImageListResolutionWrap) GetBitmapWithIntCustomBitmap(index int32, image ICustomBitmap) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(7, m.Instance(), uintptr(index), base.GetObjectUintptr(image))
}

func (m *TScaledImageListResolutionWrap) GetBitmapWithIntCustomBitmapGraphicsDrawEffect(index int32, image ICustomBitmap, effect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(8, m.Instance(), uintptr(index), base.GetObjectUintptr(image), uintptr(effect))
}

func (m *TScaledImageListResolutionWrap) DrawXY(canvas ICanvas, X int32, Y int32, index int32, enabled bool) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(9, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), api.PasBool(enabled))
}

func (m *TScaledImageListResolutionWrap) DrawXYEffect(canvas ICanvas, X int32, Y int32, index int32, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(10, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(drawEffect))
}

func (m *TScaledImageListResolutionWrap) DrawXYStyleType(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, enabled bool) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(drawingStyle), uintptr(imageType), api.PasBool(enabled))
}

func (m *TScaledImageListResolutionWrap) DrawXYStyleTypeEffect(canvas ICanvas, X int32, Y int32, index int32, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(12, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(drawingStyle), uintptr(imageType), uintptr(drawEffect))
}

func (m *TScaledImageListResolutionWrap) StretchDraw(canvas ICanvas, index int32, rect types.TRect, enabled bool) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(canvas), uintptr(index), uintptr(base.UnsafePointer(&rect)), api.PasBool(enabled))
}

func (m *TScaledImageListResolutionWrap) DrawOverlayXYOverlay(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, enabled bool) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(14, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(overlay), api.PasBool(enabled))
}

func (m *TScaledImageListResolutionWrap) DrawOverlayXYOverlayEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(overlay), uintptr(drawEffect))
}

func (m *TScaledImageListResolutionWrap) DrawOverlayXYOverlayTypeEffect(canvas ICanvas, X int32, Y int32, index int32, overlay types.TOverlay, drawingStyle types.TDrawingStyle, imageType types.TImageType, drawEffect types.TGraphicsDrawEffect) {
	if !m.IsValid() {
		return
	}
	scaledImageListResolutionWrapAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(canvas), uintptr(X), uintptr(Y), uintptr(index), uintptr(overlay), uintptr(drawingStyle), uintptr(imageType), uintptr(drawEffect))
}

// NewScaledImageListResolutionWrapWithCustomImageListResolutionDouble class constructor
func NewScaledImageListResolutionWrapWithCustomImageListResolutionDouble(resolution ICustomImageListResolution, scaleFactor float64) IScaledImageListResolutionWrap {
	r := scaledImageListResolutionWrapAPI().SysCallN(0, base.GetObjectUintptr(resolution), uintptr(base.UnsafePointer(&scaleFactor)))
	return AsScaledImageListResolutionWrap(r)
}

var (
	scaledImageListResolutionWrapOnce   base.Once
	scaledImageListResolutionWrapImport *imports.Imports = nil
)

func scaledImageListResolutionWrapAPI() *imports.Imports {
	scaledImageListResolutionWrapOnce.Do(func() {
		scaledImageListResolutionWrapImport = api.NewDefaultImports()
		scaledImageListResolutionWrapImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TScaledImageListResolutionWrap_CreateWithCustomImageListResolutionDouble", 0), // constructor NewScaledImageListResolutionWrapWithCustomImageListResolutionDouble
			/* 1 */ imports.NewTable("TScaledImageListResolutionWrap_Width", 0), // function Width
			/* 2 */ imports.NewTable("TScaledImageListResolutionWrap_Height", 0), // function Height
			/* 3 */ imports.NewTable("TScaledImageListResolutionWrap_Size", 0), // function Size
			/* 4 */ imports.NewTable("TScaledImageListResolutionWrap_Resolution", 0), // function Resolution
			/* 5 */ imports.NewTable("TScaledImageListResolutionWrap_Count", 0), // function Count
			/* 6 */ imports.NewTable("TScaledImageListResolutionWrap_Valid", 0), // function Valid
			/* 7 */ imports.NewTable("TScaledImageListResolutionWrap_GetBitmapWithIntCustomBitmap", 0), // procedure GetBitmapWithIntCustomBitmap
			/* 8 */ imports.NewTable("TScaledImageListResolutionWrap_GetBitmapWithIntCustomBitmapGraphicsDrawEffect", 0), // procedure GetBitmapWithIntCustomBitmapGraphicsDrawEffect
			/* 9 */ imports.NewTable("TScaledImageListResolutionWrap_DrawXY", 0), // procedure DrawXY
			/* 10 */ imports.NewTable("TScaledImageListResolutionWrap_DrawXYEffect", 0), // procedure DrawXYEffect
			/* 11 */ imports.NewTable("TScaledImageListResolutionWrap_DrawXYStyleType", 0), // procedure DrawXYStyleType
			/* 12 */ imports.NewTable("TScaledImageListResolutionWrap_DrawXYStyleTypeEffect", 0), // procedure DrawXYStyleTypeEffect
			/* 13 */ imports.NewTable("TScaledImageListResolutionWrap_StretchDraw", 0), // procedure StretchDraw
			/* 14 */ imports.NewTable("TScaledImageListResolutionWrap_DrawOverlayXYOverlay", 0), // procedure DrawOverlayXYOverlay
			/* 15 */ imports.NewTable("TScaledImageListResolutionWrap_DrawOverlayXYOverlayEffect", 0), // procedure DrawOverlayXYOverlayEffect
			/* 16 */ imports.NewTable("TScaledImageListResolutionWrap_DrawOverlayXYOverlayTypeEffect", 0), // procedure DrawOverlayXYOverlayTypeEffect
		}
	})
	return scaledImageListResolutionWrapImport
}
