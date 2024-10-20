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

// ICustomImageListResolution Parent: ILCLReferenceComponent
type ICustomImageListResolution interface {
	ILCLReferenceComponent
	ImageList() ICustomImageList                                                                                                                                // property
	Width() int32                                                                                                                                               // property
	Height() int32                                                                                                                                              // property
	Count() int32                                                                                                                                               // property
	AutoCreatedInDesignTime() bool                                                                                                                              // property
	SetAutoCreatedInDesignTime(AValue bool)                                                                                                                     // property
	GetHotSpot() (resultPoint TPoint)                                                                                                                           // function
	GetBitmap(Index int32, Image ICustomBitmap)                                                                                                                 // procedure
	GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect)                                                                                   // procedure
	GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect)                                                                                              // procedure
	GetIcon1(Index int32, Image IIcon)                                                                                                                          // procedure
	GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect)                                                                                             // procedure
	GetRawImage(Index int32, OutImage *TRawImage)                                                                                                               // procedure
	Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool)                                                                                                  // procedure
	Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect)                                                                               // procedure
	Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool)                                             // procedure
	Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect)                           // procedure
	StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool)                                                                                        // procedure
	DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool)                                                                        // procedure
	DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect)                                                     // procedure
	DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) // procedure
}

// TCustomImageListResolution Parent: TLCLReferenceComponent
type TCustomImageListResolution struct {
	TLCLReferenceComponent
}

func NewCustomImageListResolution(TheOwner IComponent) ICustomImageListResolution {
	r1 := customImageListResolutionImportAPI().SysCallN(3, GetObjectUintptr(TheOwner))
	return AsCustomImageListResolution(r1)
}

func (m *TCustomImageListResolution) ImageList() ICustomImageList {
	r1 := customImageListResolutionImportAPI().SysCallN(19, m.Instance())
	return AsCustomImageList(r1)
}

func (m *TCustomImageListResolution) Width() int32 {
	r1 := customImageListResolutionImportAPI().SysCallN(21, m.Instance())
	return int32(r1)
}

func (m *TCustomImageListResolution) Height() int32 {
	r1 := customImageListResolutionImportAPI().SysCallN(18, m.Instance())
	return int32(r1)
}

func (m *TCustomImageListResolution) Count() int32 {
	r1 := customImageListResolutionImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TCustomImageListResolution) AutoCreatedInDesignTime() bool {
	r1 := customImageListResolutionImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImageListResolution) SetAutoCreatedInDesignTime(AValue bool) {
	customImageListResolutionImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImageListResolution) GetHotSpot() (resultPoint TPoint) {
	customImageListResolutionImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(&resultPoint)))
	return
}

func CustomImageListResolutionClass() TClass {
	ret := customImageListResolutionImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCustomImageListResolution) GetBitmap(Index int32, Image ICustomBitmap) {
	customImageListResolutionImportAPI().SysCallN(11, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageListResolution) GetBitmap1(Index int32, Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	customImageListResolutionImportAPI().SysCallN(12, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageListResolution) GetIcon(Index int32, Image IIcon, AEffect TGraphicsDrawEffect) {
	customImageListResolutionImportAPI().SysCallN(15, m.Instance(), uintptr(Index), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageListResolution) GetIcon1(Index int32, Image IIcon) {
	customImageListResolutionImportAPI().SysCallN(16, m.Instance(), uintptr(Index), GetObjectUintptr(Image))
}

func (m *TCustomImageListResolution) GetFullBitmap(Image ICustomBitmap, AEffect TGraphicsDrawEffect) {
	customImageListResolutionImportAPI().SysCallN(13, m.Instance(), GetObjectUintptr(Image), uintptr(AEffect))
}

func (m *TCustomImageListResolution) GetRawImage(Index int32, OutImage *TRawImage) {
	var result1 uintptr
	customImageListResolutionImportAPI().SysCallN(17, m.Instance(), uintptr(Index), uintptr(unsafePointer(&result1)))
	*OutImage = *(*TRawImage)(getPointer(result1))
}

func (m *TCustomImageListResolution) Draw(ACanvas ICanvas, AX, AY, AIndex int32, AEnabled bool) {
	customImageListResolutionImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), PascalBool(AEnabled))
}

func (m *TCustomImageListResolution) Draw1(ACanvas ICanvas, AX, AY, AIndex int32, ADrawEffect TGraphicsDrawEffect) {
	customImageListResolutionImportAPI().SysCallN(5, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawEffect))
}

func (m *TCustomImageListResolution) Draw2(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, AEnabled bool) {
	customImageListResolutionImportAPI().SysCallN(6, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), PascalBool(AEnabled))
}

func (m *TCustomImageListResolution) Draw3(ACanvas ICanvas, AX, AY, AIndex int32, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	customImageListResolutionImportAPI().SysCallN(7, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

func (m *TCustomImageListResolution) StretchDraw(Canvas ICanvas, Index int32, ARect *TRect, Enabled bool) {
	customImageListResolutionImportAPI().SysCallN(20, m.Instance(), GetObjectUintptr(Canvas), uintptr(Index), uintptr(unsafePointer(ARect)), PascalBool(Enabled))
}

func (m *TCustomImageListResolution) DrawOverlay(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, AEnabled bool) {
	customImageListResolutionImportAPI().SysCallN(8, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), PascalBool(AEnabled))
}

func (m *TCustomImageListResolution) DrawOverlay1(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawEffect TGraphicsDrawEffect) {
	customImageListResolutionImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawEffect))
}

func (m *TCustomImageListResolution) DrawOverlay2(ACanvas ICanvas, AX, AY, AIndex int32, AOverlay TOverlay, ADrawingStyle TDrawingStyle, AImageType TImageType, ADrawEffect TGraphicsDrawEffect) {
	customImageListResolutionImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(ACanvas), uintptr(AX), uintptr(AY), uintptr(AIndex), uintptr(AOverlay), uintptr(ADrawingStyle), uintptr(AImageType), uintptr(ADrawEffect))
}

var (
	customImageListResolutionImport       *imports.Imports = nil
	customImageListResolutionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomImageListResolution_AutoCreatedInDesignTime", 0),
		/*1*/ imports.NewTable("CustomImageListResolution_Class", 0),
		/*2*/ imports.NewTable("CustomImageListResolution_Count", 0),
		/*3*/ imports.NewTable("CustomImageListResolution_Create", 0),
		/*4*/ imports.NewTable("CustomImageListResolution_Draw", 0),
		/*5*/ imports.NewTable("CustomImageListResolution_Draw1", 0),
		/*6*/ imports.NewTable("CustomImageListResolution_Draw2", 0),
		/*7*/ imports.NewTable("CustomImageListResolution_Draw3", 0),
		/*8*/ imports.NewTable("CustomImageListResolution_DrawOverlay", 0),
		/*9*/ imports.NewTable("CustomImageListResolution_DrawOverlay1", 0),
		/*10*/ imports.NewTable("CustomImageListResolution_DrawOverlay2", 0),
		/*11*/ imports.NewTable("CustomImageListResolution_GetBitmap", 0),
		/*12*/ imports.NewTable("CustomImageListResolution_GetBitmap1", 0),
		/*13*/ imports.NewTable("CustomImageListResolution_GetFullBitmap", 0),
		/*14*/ imports.NewTable("CustomImageListResolution_GetHotSpot", 0),
		/*15*/ imports.NewTable("CustomImageListResolution_GetIcon", 0),
		/*16*/ imports.NewTable("CustomImageListResolution_GetIcon1", 0),
		/*17*/ imports.NewTable("CustomImageListResolution_GetRawImage", 0),
		/*18*/ imports.NewTable("CustomImageListResolution_Height", 0),
		/*19*/ imports.NewTable("CustomImageListResolution_ImageList", 0),
		/*20*/ imports.NewTable("CustomImageListResolution_StretchDraw", 0),
		/*21*/ imports.NewTable("CustomImageListResolution_Width", 0),
	}
)

func customImageListResolutionImportAPI() *imports.Imports {
	if customImageListResolutionImport == nil {
		customImageListResolutionImport = NewDefaultImports()
		customImageListResolutionImport.SetImportTable(customImageListResolutionImportTables)
		customImageListResolutionImportTables = nil
	}
	return customImageListResolutionImport
}
