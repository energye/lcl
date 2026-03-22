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

// IFPCustomCanvas Parent: IPersistent
type IFPCustomCanvas interface {
	IPersistent
	Locked() bool                              // function
	CreateFont() IFPCustomFont                 // function
	CreatePen() IFPCustomPen                   // function
	CreateBrush() IFPCustomBrush               // function
	GetTextHeight(text string) int32           // function
	GetTextWidth(text string) int32            // function
	TextExtentWithStr(text string) types.TSize // function
	TextHeightWithStr(text string) int32       // function
	TextWidthWithStr(text string) int32        // function
	LockCanvas()                               // procedure
	UnlockCanvas()                             // procedure
	// TextOutWithIntX2Str
	//  using font
	TextOutWithIntX2Str(X int32, Y int32, text string) // procedure
	GetTextSize(text string, W *int32, H *int32)       // procedure
	// ArcWithIntX6
	//  using pen and brush
	ArcWithIntX6(left int32, top int32, right int32, bottom int32, angle16Deg int32, angle16DegLength int32)  // procedure
	ArcWithIntX8(left int32, top int32, right int32, bottom int32, sX int32, sY int32, eX int32, eY int32)    // procedure
	EllipseWithRect(bounds types.TRect)                                                                       // procedure
	EllipseWithIntX4(left int32, top int32, right int32, bottom int32)                                        // procedure
	EllipseC(X int32, Y int32, rx uint32, ry uint32)                                                          // procedure
	RadialPieWithIntX6(x1 int32, y1 int32, x2 int32, y2 int32, startAngle16Deg int32, angle16DegLength int32) // procedure
	PolyBezierWithPointIntBoolX2(points types.PPoint, numPts int32, filled bool, continuous bool)             // procedure
	RectangleWithRect(bounds types.TRect)                                                                     // procedure
	RectangleWithIntX4(left int32, top int32, right int32, bottom int32)                                      // procedure
	FillRectWithRect(rect types.TRect)                                                                        // procedure
	FillRectWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32)                                                 // procedure
	// FloodFillWithIntX2
	//  using brush
	FloodFillWithIntX2(X int32, Y int32) // procedure
	Clear()                              // procedure
	// MoveToWithIntX2
	//  using pen
	MoveToWithIntX2(X int32, Y int32)                     // procedure
	MoveToWithPoint(P types.TPoint)                       // procedure
	LineToWithIntX2(X int32, Y int32)                     // procedure
	LineToWithPoint(P types.TPoint)                       // procedure
	LineWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32) // procedure
	LineWithPointX2(p1 types.TPoint, p2 types.TPoint)     // procedure
	LineWithRect(points types.TRect)                      // procedure
	// CopyRectWithIntX2FPCCanvasRect
	//  other procedures
	CopyRectWithIntX2FPCCanvasRect(X int32, Y int32, canvas IFPCustomCanvas, sourceRect types.TRect) // procedure
	DrawWithIntX2FPCImage(X int32, Y int32, image IFPCustomImage)                                    // procedure
	StretchDrawWithIntX4FPCImage(X int32, Y int32, W int32, H int32, source IFPCustomImage)          // procedure
	Erase()                                                                                          // procedure
	DrawPixel(X int32, Y int32, newcolor TFPColor)                                                   // procedure
	// LockCount
	//  properties
	LockCount() int32                              // property LockCount Getter
	FontToFPCustomFont() IFPCustomFont             // property Font Getter
	SetFontToFPCustomFont(value IFPCustomFont)     // property Font Setter
	PenToFPCustomPen() IFPCustomPen                // property Pen Getter
	SetPenToFPCustomPen(value IFPCustomPen)        // property Pen Setter
	BrushToFPCustomBrush() IFPCustomBrush          // property Brush Getter
	SetBrushToFPCustomBrush(value IFPCustomBrush)  // property Brush Setter
	Interpolation() IFPCustomInterpolation         // property Interpolation Getter
	SetInterpolation(value IFPCustomInterpolation) // property Interpolation Setter
	Colors(X int32, Y int32) TFPColor              // property Colors Getter
	SetColors(X int32, Y int32, value TFPColor)    // property Colors Setter
	ClipRect() types.TRect                         // property ClipRect Getter
	SetClipRect(value types.TRect)                 // property ClipRect Setter
	ClipRegion() IFPCustomRegion                   // property ClipRegion Getter
	SetClipRegion(value IFPCustomRegion)           // property ClipRegion Setter
	Clipping() bool                                // property Clipping Getter
	SetClipping(value bool)                        // property Clipping Setter
	PenPos() types.TPoint                          // property PenPos Getter
	SetPenPos(value types.TPoint)                  // property PenPos Setter
	Height() int32                                 // property Height Getter
	SetHeight(value int32)                         // property Height Setter
	Width() int32                                  // property Width Getter
	SetWidth(value int32)                          // property Width Setter
	ManageResources() bool                         // property ManageResources Getter
	SetManageResources(value bool)                 // property ManageResources Setter
	DrawingMode() types.TFPDrawingMode             // property DrawingMode Getter
	SetDrawingMode(value types.TFPDrawingMode)     // property DrawingMode Setter
	SetOnCombineColors(fn TFPCanvasCombineColors)  // property event
}

type TFPCustomCanvas struct {
	TPersistent
}

func (m *TFPCustomCanvas) Locked() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomCanvasAPI().SysCallN(0, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCustomCanvas) CreateFont() IFPCustomFont {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomCanvasAPI().SysCallN(1, m.Instance())
	return AsFPCustomFont(r)
}

func (m *TFPCustomCanvas) CreatePen() IFPCustomPen {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomCanvasAPI().SysCallN(2, m.Instance())
	return AsFPCustomPen(r)
}

func (m *TFPCustomCanvas) CreateBrush() IFPCustomBrush {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomCanvasAPI().SysCallN(3, m.Instance())
	return AsFPCustomBrush(r)
}

func (m *TFPCustomCanvas) GetTextHeight(text string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomCanvasAPI().SysCallN(4, m.Instance(), api.PasStr(text))
	return int32(r)
}

func (m *TFPCustomCanvas) GetTextWidth(text string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomCanvasAPI().SysCallN(5, m.Instance(), api.PasStr(text))
	return int32(r)
}

func (m *TFPCustomCanvas) TextExtentWithStr(text string) (result types.TSize) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(6, m.Instance(), api.PasStr(text), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TFPCustomCanvas) TextHeightWithStr(text string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomCanvasAPI().SysCallN(7, m.Instance(), api.PasStr(text))
	return int32(r)
}

func (m *TFPCustomCanvas) TextWidthWithStr(text string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomCanvasAPI().SysCallN(8, m.Instance(), api.PasStr(text))
	return int32(r)
}

func (m *TFPCustomCanvas) LockCanvas() {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(9, m.Instance())
}

func (m *TFPCustomCanvas) UnlockCanvas() {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(10, m.Instance())
}

func (m *TFPCustomCanvas) TextOutWithIntX2Str(X int32, Y int32, text string) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(11, m.Instance(), uintptr(X), uintptr(Y), api.PasStr(text))
}

func (m *TFPCustomCanvas) GetTextSize(text string, W *int32, H *int32) {
	if !m.IsValid() {
		return
	}
	WPtr := uintptr(*W)
	HPtr := uintptr(*H)
	fPCustomCanvasAPI().SysCallN(12, m.Instance(), api.PasStr(text), uintptr(base.UnsafePointer(&WPtr)), uintptr(base.UnsafePointer(&HPtr)))
	*W = int32(WPtr)
	*H = int32(HPtr)
}

func (m *TFPCustomCanvas) ArcWithIntX6(left int32, top int32, right int32, bottom int32, angle16Deg int32, angle16DegLength int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(13, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(angle16Deg), uintptr(angle16DegLength))
}

func (m *TFPCustomCanvas) ArcWithIntX8(left int32, top int32, right int32, bottom int32, sX int32, sY int32, eX int32, eY int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(14, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(sX), uintptr(sY), uintptr(eX), uintptr(eY))
}

func (m *TFPCustomCanvas) EllipseWithRect(bounds types.TRect) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&bounds)))
}

func (m *TFPCustomCanvas) EllipseWithIntX4(left int32, top int32, right int32, bottom int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(16, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

func (m *TFPCustomCanvas) EllipseC(X int32, Y int32, rx uint32, ry uint32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(17, m.Instance(), uintptr(X), uintptr(Y), uintptr(rx), uintptr(ry))
}

func (m *TFPCustomCanvas) RadialPieWithIntX6(x1 int32, y1 int32, x2 int32, y2 int32, startAngle16Deg int32, angle16DegLength int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(18, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(startAngle16Deg), uintptr(angle16DegLength))
}

func (m *TFPCustomCanvas) PolyBezierWithPointIntBoolX2(points types.PPoint, numPts int32, filled bool, continuous bool) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(19, m.Instance(), uintptr(points), uintptr(numPts), api.PasBool(filled), api.PasBool(continuous))
}

func (m *TFPCustomCanvas) RectangleWithRect(bounds types.TRect) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(&bounds)))
}

func (m *TFPCustomCanvas) RectangleWithIntX4(left int32, top int32, right int32, bottom int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(21, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

func (m *TFPCustomCanvas) FillRectWithRect(rect types.TRect) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(22, m.Instance(), uintptr(base.UnsafePointer(&rect)))
}

func (m *TFPCustomCanvas) FillRectWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(23, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TFPCustomCanvas) FloodFillWithIntX2(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(24, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TFPCustomCanvas) Clear() {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(25, m.Instance())
}

func (m *TFPCustomCanvas) MoveToWithIntX2(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(26, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TFPCustomCanvas) MoveToWithPoint(P types.TPoint) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(27, m.Instance(), uintptr(base.UnsafePointer(&P)))
}

func (m *TFPCustomCanvas) LineToWithIntX2(X int32, Y int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(28, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TFPCustomCanvas) LineToWithPoint(P types.TPoint) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(29, m.Instance(), uintptr(base.UnsafePointer(&P)))
}

func (m *TFPCustomCanvas) LineWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(30, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TFPCustomCanvas) LineWithPointX2(p1 types.TPoint, p2 types.TPoint) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(31, m.Instance(), uintptr(base.UnsafePointer(&p1)), uintptr(base.UnsafePointer(&p2)))
}

func (m *TFPCustomCanvas) LineWithRect(points types.TRect) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(32, m.Instance(), uintptr(base.UnsafePointer(&points)))
}

func (m *TFPCustomCanvas) CopyRectWithIntX2FPCCanvasRect(X int32, Y int32, canvas IFPCustomCanvas, sourceRect types.TRect) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(33, m.Instance(), uintptr(X), uintptr(Y), base.GetObjectUintptr(canvas), uintptr(base.UnsafePointer(&sourceRect)))
}

func (m *TFPCustomCanvas) DrawWithIntX2FPCImage(X int32, Y int32, image IFPCustomImage) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(34, m.Instance(), uintptr(X), uintptr(Y), base.GetObjectUintptr(image))
}

func (m *TFPCustomCanvas) StretchDrawWithIntX4FPCImage(X int32, Y int32, W int32, H int32, source IFPCustomImage) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(35, m.Instance(), uintptr(X), uintptr(Y), uintptr(W), uintptr(H), base.GetObjectUintptr(source))
}

func (m *TFPCustomCanvas) Erase() {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(36, m.Instance())
}

func (m *TFPCustomCanvas) DrawPixel(X int32, Y int32, newcolor TFPColor) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(37, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(&newcolor)))
}

func (m *TFPCustomCanvas) LockCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomCanvasAPI().SysCallN(38, m.Instance())
	return int32(r)
}

func (m *TFPCustomCanvas) FontToFPCustomFont() IFPCustomFont {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomCanvasAPI().SysCallN(39, 0, m.Instance())
	return AsFPCustomFont(r)
}

func (m *TFPCustomCanvas) SetFontToFPCustomFont(value IFPCustomFont) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(39, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFPCustomCanvas) PenToFPCustomPen() IFPCustomPen {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomCanvasAPI().SysCallN(40, 0, m.Instance())
	return AsFPCustomPen(r)
}

func (m *TFPCustomCanvas) SetPenToFPCustomPen(value IFPCustomPen) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(40, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFPCustomCanvas) BrushToFPCustomBrush() IFPCustomBrush {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomCanvasAPI().SysCallN(41, 0, m.Instance())
	return AsFPCustomBrush(r)
}

func (m *TFPCustomCanvas) SetBrushToFPCustomBrush(value IFPCustomBrush) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(41, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFPCustomCanvas) Interpolation() IFPCustomInterpolation {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomCanvasAPI().SysCallN(42, 0, m.Instance())
	return AsFPCustomInterpolation(r)
}

func (m *TFPCustomCanvas) SetInterpolation(value IFPCustomInterpolation) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(42, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFPCustomCanvas) Colors(X int32, Y int32) (result TFPColor) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(43, 0, m.Instance(), uintptr(X), uintptr(Y), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TFPCustomCanvas) SetColors(X int32, Y int32, value TFPColor) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(43, 1, m.Instance(), uintptr(X), uintptr(Y), uintptr(base.UnsafePointer(&value)))
}

func (m *TFPCustomCanvas) ClipRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(44, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TFPCustomCanvas) SetClipRect(value types.TRect) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(44, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TFPCustomCanvas) ClipRegion() IFPCustomRegion {
	if !m.IsValid() {
		return nil
	}
	r := fPCustomCanvasAPI().SysCallN(45, 0, m.Instance())
	return AsFPCustomRegion(r)
}

func (m *TFPCustomCanvas) SetClipRegion(value IFPCustomRegion) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(45, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFPCustomCanvas) Clipping() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomCanvasAPI().SysCallN(46, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCustomCanvas) SetClipping(value bool) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(46, 1, m.Instance(), api.PasBool(value))
}

func (m *TFPCustomCanvas) PenPos() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(47, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TFPCustomCanvas) SetPenPos(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(47, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TFPCustomCanvas) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomCanvasAPI().SysCallN(48, 0, m.Instance())
	return int32(r)
}

func (m *TFPCustomCanvas) SetHeight(value int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(48, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomCanvas) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomCanvasAPI().SysCallN(49, 0, m.Instance())
	return int32(r)
}

func (m *TFPCustomCanvas) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(49, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomCanvas) ManageResources() bool {
	if !m.IsValid() {
		return false
	}
	r := fPCustomCanvasAPI().SysCallN(50, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFPCustomCanvas) SetManageResources(value bool) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(50, 1, m.Instance(), api.PasBool(value))
}

func (m *TFPCustomCanvas) DrawingMode() types.TFPDrawingMode {
	if !m.IsValid() {
		return 0
	}
	r := fPCustomCanvasAPI().SysCallN(51, 0, m.Instance())
	return types.TFPDrawingMode(r)
}

func (m *TFPCustomCanvas) SetDrawingMode(value types.TFPDrawingMode) {
	if !m.IsValid() {
		return
	}
	fPCustomCanvasAPI().SysCallN(51, 1, m.Instance(), uintptr(value))
}

func (m *TFPCustomCanvas) SetOnCombineColors(fn TFPCanvasCombineColors) {
	if !m.IsValid() {
		return
	}
	cb := makeTFPCanvasCombineColors(fn)
	base.SetEvent(m, 52, fPCustomCanvasAPI(), api.MakeEventDataPtr(cb))
}

var (
	fPCustomCanvasOnce   base.Once
	fPCustomCanvasImport *imports.Imports = nil
)

func fPCustomCanvasAPI() *imports.Imports {
	fPCustomCanvasOnce.Do(func() {
		fPCustomCanvasImport = api.NewDefaultImports()
		fPCustomCanvasImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFPCustomCanvas_Locked", 0), // function Locked
			/* 1 */ imports.NewTable("TFPCustomCanvas_CreateFont", 0), // function CreateFont
			/* 2 */ imports.NewTable("TFPCustomCanvas_CreatePen", 0), // function CreatePen
			/* 3 */ imports.NewTable("TFPCustomCanvas_CreateBrush", 0), // function CreateBrush
			/* 4 */ imports.NewTable("TFPCustomCanvas_GetTextHeight", 0), // function GetTextHeight
			/* 5 */ imports.NewTable("TFPCustomCanvas_GetTextWidth", 0), // function GetTextWidth
			/* 6 */ imports.NewTable("TFPCustomCanvas_TextExtentWithStr", 0), // function TextExtentWithStr
			/* 7 */ imports.NewTable("TFPCustomCanvas_TextHeightWithStr", 0), // function TextHeightWithStr
			/* 8 */ imports.NewTable("TFPCustomCanvas_TextWidthWithStr", 0), // function TextWidthWithStr
			/* 9 */ imports.NewTable("TFPCustomCanvas_LockCanvas", 0), // procedure LockCanvas
			/* 10 */ imports.NewTable("TFPCustomCanvas_UnlockCanvas", 0), // procedure UnlockCanvas
			/* 11 */ imports.NewTable("TFPCustomCanvas_TextOutWithIntX2Str", 0), // procedure TextOutWithIntX2Str
			/* 12 */ imports.NewTable("TFPCustomCanvas_GetTextSize", 0), // procedure GetTextSize
			/* 13 */ imports.NewTable("TFPCustomCanvas_ArcWithIntX6", 0), // procedure ArcWithIntX6
			/* 14 */ imports.NewTable("TFPCustomCanvas_ArcWithIntX8", 0), // procedure ArcWithIntX8
			/* 15 */ imports.NewTable("TFPCustomCanvas_EllipseWithRect", 0), // procedure EllipseWithRect
			/* 16 */ imports.NewTable("TFPCustomCanvas_EllipseWithIntX4", 0), // procedure EllipseWithIntX4
			/* 17 */ imports.NewTable("TFPCustomCanvas_EllipseC", 0), // procedure EllipseC
			/* 18 */ imports.NewTable("TFPCustomCanvas_RadialPieWithIntX6", 0), // procedure RadialPieWithIntX6
			/* 19 */ imports.NewTable("TFPCustomCanvas_PolyBezierWithPointIntBoolX2", 0), // procedure PolyBezierWithPointIntBoolX2
			/* 20 */ imports.NewTable("TFPCustomCanvas_RectangleWithRect", 0), // procedure RectangleWithRect
			/* 21 */ imports.NewTable("TFPCustomCanvas_RectangleWithIntX4", 0), // procedure RectangleWithIntX4
			/* 22 */ imports.NewTable("TFPCustomCanvas_FillRectWithRect", 0), // procedure FillRectWithRect
			/* 23 */ imports.NewTable("TFPCustomCanvas_FillRectWithIntX4", 0), // procedure FillRectWithIntX4
			/* 24 */ imports.NewTable("TFPCustomCanvas_FloodFillWithIntX2", 0), // procedure FloodFillWithIntX2
			/* 25 */ imports.NewTable("TFPCustomCanvas_Clear", 0), // procedure Clear
			/* 26 */ imports.NewTable("TFPCustomCanvas_MoveToWithIntX2", 0), // procedure MoveToWithIntX2
			/* 27 */ imports.NewTable("TFPCustomCanvas_MoveToWithPoint", 0), // procedure MoveToWithPoint
			/* 28 */ imports.NewTable("TFPCustomCanvas_LineToWithIntX2", 0), // procedure LineToWithIntX2
			/* 29 */ imports.NewTable("TFPCustomCanvas_LineToWithPoint", 0), // procedure LineToWithPoint
			/* 30 */ imports.NewTable("TFPCustomCanvas_LineWithIntX4", 0), // procedure LineWithIntX4
			/* 31 */ imports.NewTable("TFPCustomCanvas_LineWithPointX2", 0), // procedure LineWithPointX2
			/* 32 */ imports.NewTable("TFPCustomCanvas_LineWithRect", 0), // procedure LineWithRect
			/* 33 */ imports.NewTable("TFPCustomCanvas_CopyRectWithIntX2FPCCanvasRect", 0), // procedure CopyRectWithIntX2FPCCanvasRect
			/* 34 */ imports.NewTable("TFPCustomCanvas_DrawWithIntX2FPCImage", 0), // procedure DrawWithIntX2FPCImage
			/* 35 */ imports.NewTable("TFPCustomCanvas_StretchDrawWithIntX4FPCImage", 0), // procedure StretchDrawWithIntX4FPCImage
			/* 36 */ imports.NewTable("TFPCustomCanvas_Erase", 0), // procedure Erase
			/* 37 */ imports.NewTable("TFPCustomCanvas_DrawPixel", 0), // procedure DrawPixel
			/* 38 */ imports.NewTable("TFPCustomCanvas_LockCount", 0), // property LockCount
			/* 39 */ imports.NewTable("TFPCustomCanvas_FontToFPCustomFont", 0), // property FontToFPCustomFont
			/* 40 */ imports.NewTable("TFPCustomCanvas_PenToFPCustomPen", 0), // property PenToFPCustomPen
			/* 41 */ imports.NewTable("TFPCustomCanvas_BrushToFPCustomBrush", 0), // property BrushToFPCustomBrush
			/* 42 */ imports.NewTable("TFPCustomCanvas_Interpolation", 0), // property Interpolation
			/* 43 */ imports.NewTable("TFPCustomCanvas_Colors", 0), // property Colors
			/* 44 */ imports.NewTable("TFPCustomCanvas_ClipRect", 0), // property ClipRect
			/* 45 */ imports.NewTable("TFPCustomCanvas_ClipRegion", 0), // property ClipRegion
			/* 46 */ imports.NewTable("TFPCustomCanvas_Clipping", 0), // property Clipping
			/* 47 */ imports.NewTable("TFPCustomCanvas_PenPos", 0), // property PenPos
			/* 48 */ imports.NewTable("TFPCustomCanvas_Height", 0), // property Height
			/* 49 */ imports.NewTable("TFPCustomCanvas_Width", 0), // property Width
			/* 50 */ imports.NewTable("TFPCustomCanvas_ManageResources", 0), // property ManageResources
			/* 51 */ imports.NewTable("TFPCustomCanvas_DrawingMode", 0), // property DrawingMode
			/* 52 */ imports.NewTable("TFPCustomCanvas_OnCombineColors", 0), // event OnCombineColors
		}
	})
	return fPCustomCanvasImport
}
