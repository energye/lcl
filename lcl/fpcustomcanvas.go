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

// IFPCustomCanvas Is Abstract Class Parent: IPersistent
type IFPCustomCanvas interface {
	IPersistent
	LockCount() int32                                                     // property
	Font() IFPCustomFont                                                  // property
	SetFont(AValue IFPCustomFont)                                         // property
	Pen() IFPCustomPen                                                    // property
	SetPen(AValue IFPCustomPen)                                           // property
	Brush() IFPCustomBrush                                                // property
	SetBrush(AValue IFPCustomBrush)                                       // property
	Colors(x, y int32) (resultFPColor TFPColor)                           // property
	SetColors(x, y int32, AValue *TFPColor)                               // property
	ClipRect() (resultRect TRect)                                         // property
	SetClipRect(AValue *TRect)                                            // property
	ClipRegion() IFPCustomRegion                                          // property
	SetClipRegion(AValue IFPCustomRegion)                                 // property
	Clipping() bool                                                       // property
	SetClipping(AValue bool)                                              // property
	PenPos() (resultPoint TPoint)                                         // property
	SetPenPos(AValue *TPoint)                                             // property
	Height() int32                                                        // property
	SetHeight(AValue int32)                                               // property
	Width() int32                                                         // property
	SetWidth(AValue int32)                                                // property
	ManageResources() bool                                                // property
	SetManageResources(AValue bool)                                       // property
	DrawingMode() TFPDrawingMode                                          // property
	SetDrawingMode(AValue TFPDrawingMode)                                 // property
	Locked() bool                                                         // function
	CreateFont() IFPCustomFont                                            // function
	CreatePen() IFPCustomPen                                              // function
	CreateBrush() IFPCustomBrush                                          // function
	GetTextHeight(text string) int32                                      // function
	GetTextWidth(text string) int32                                       // function
	TextExtent(Text string) (resultSize TSize)                            // function
	TextHeight(Text string) int32                                         // function
	TextWidth(Text string) int32                                          // function
	GetTextHeight1(text string) int32                                     // function
	GetTextWidth1(text string) int32                                      // function
	TextExtent1(Text string) (resultSize TSize)                           // function
	TextHeight1(Text string) int32                                        // function
	TextWidth1(Text string) int32                                         // function
	LockCanvas()                                                          // procedure
	UnlockCanvas()                                                        // procedure
	TextOut(x, y int32, text string)                                      // procedure
	GetTextSize(text string, w, h *int32)                                 // procedure
	TextOut1(x, y int32, text string)                                     // procedure
	GetTextSize1(text string, w, h *int32)                                // procedure
	Arc(ALeft, ATop, ARight, ABottom, Angle16Deg, Angle16DegLength int32) // procedure
	Arc1(ALeft, ATop, ARight, ABottom, SX, SY, EX, EY int32)              // procedure
	Ellipse(Bounds *TRect)                                                // procedure
	Ellipse1(left, top, right, bottom int32)                              // procedure
	EllipseC(x, y int32, rx, ry uint32)                                   // procedure
	RadialPie(x1, y1, x2, y2, StartAngle16Deg, Angle16DegLength int32)    // procedure
	PolyBezier(Points []TPoint, Filled bool, Continuous bool)             // procedure
	Rectangle(Bounds *TRect)                                              // procedure
	Rectangle1(left, top, right, bottom int32)                            // procedure
	FillRect(ARect *TRect)                                                // procedure
	FillRect1(X1, Y1, X2, Y2 int32)                                       // procedure
	FloodFill(x, y int32)                                                 // procedure
	Clear()                                                               // procedure
	MoveTo(x, y int32)                                                    // procedure
	MoveTo1(p *TPoint)                                                    // procedure
	LineTo(x, y int32)                                                    // procedure
	LineTo1(p *TPoint)                                                    // procedure
	Line(x1, y1, x2, y2 int32)                                            // procedure
	Line1(p1, p2 *TPoint)                                                 // procedure
	Line2(points *TRect)                                                  // procedure
	CopyRect(x, y int32, canvas IFPCustomCanvas, SourceRect *TRect)       // procedure
	Draw(x, y int32, image IFPCustomImage)                                // procedure
	StretchDraw(x, y, w, h int32, source IFPCustomImage)                  // procedure
	Erase()                                                               // procedure
	DrawPixel(x, y int32, newcolor *TFPColor)                             // procedure
	SetOnCombineColors(fn TFPCanvasCombineColors)                         // property event
}

// TFPCustomCanvas Is Abstract Class Parent: TPersistent
type TFPCustomCanvas struct {
	TPersistent
	combineColorsPtr uintptr
}

func (m *TFPCustomCanvas) LockCount() int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(37, m.Instance())
	return int32(r1)
}

func (m *TFPCustomCanvas) Font() IFPCustomFont {
	r1 := fPCustomCanvasImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return AsFPCustomFont(r1)
}

func (m *TFPCustomCanvas) SetFont(AValue IFPCustomFont) {
	fPCustomCanvasImportAPI().SysCallN(23, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Pen() IFPCustomPen {
	r1 := fPCustomCanvasImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return AsFPCustomPen(r1)
}

func (m *TFPCustomCanvas) SetPen(AValue IFPCustomPen) {
	fPCustomCanvasImportAPI().SysCallN(42, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Brush() IFPCustomBrush {
	r1 := fPCustomCanvasImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return AsFPCustomBrush(r1)
}

func (m *TFPCustomCanvas) SetBrush(AValue IFPCustomBrush) {
	fPCustomCanvasImportAPI().SysCallN(2, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Colors(x, y int32) (resultFPColor TFPColor) {
	r1 := fPCustomCanvasImportAPI().SysCallN(8, 0, m.Instance(), uintptr(x), uintptr(y))
	return *(*TFPColor)(getPointer(r1))
}

func (m *TFPCustomCanvas) SetColors(x, y int32, AValue *TFPColor) {
	fPCustomCanvasImportAPI().SysCallN(8, 1, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(AValue)))
}

func (m *TFPCustomCanvas) ClipRect() (resultRect TRect) {
	fPCustomCanvasImportAPI().SysCallN(5, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TFPCustomCanvas) SetClipRect(AValue *TRect) {
	fPCustomCanvasImportAPI().SysCallN(5, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TFPCustomCanvas) ClipRegion() IFPCustomRegion {
	r1 := fPCustomCanvasImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return AsFPCustomRegion(r1)
}

func (m *TFPCustomCanvas) SetClipRegion(AValue IFPCustomRegion) {
	fPCustomCanvasImportAPI().SysCallN(6, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFPCustomCanvas) Clipping() bool {
	r1 := fPCustomCanvasImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomCanvas) SetClipping(AValue bool) {
	fPCustomCanvasImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomCanvas) PenPos() (resultPoint TPoint) {
	fPCustomCanvasImportAPI().SysCallN(43, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TFPCustomCanvas) SetPenPos(AValue *TPoint) {
	fPCustomCanvasImportAPI().SysCallN(43, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TFPCustomCanvas) Height() int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomCanvas) SetHeight(AValue int32) {
	fPCustomCanvasImportAPI().SysCallN(30, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomCanvas) Width() int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(59, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFPCustomCanvas) SetWidth(AValue int32) {
	fPCustomCanvasImportAPI().SysCallN(59, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomCanvas) ManageResources() bool {
	r1 := fPCustomCanvasImportAPI().SysCallN(39, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFPCustomCanvas) SetManageResources(AValue bool) {
	fPCustomCanvasImportAPI().SysCallN(39, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFPCustomCanvas) DrawingMode() TFPDrawingMode {
	r1 := fPCustomCanvasImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TFPDrawingMode(r1)
}

func (m *TFPCustomCanvas) SetDrawingMode(AValue TFPDrawingMode) {
	fPCustomCanvasImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TFPCustomCanvas) Locked() bool {
	r1 := fPCustomCanvasImportAPI().SysCallN(38, m.Instance())
	return GoBool(r1)
}

func (m *TFPCustomCanvas) CreateFont() IFPCustomFont {
	r1 := fPCustomCanvasImportAPI().SysCallN(11, m.Instance())
	return AsFPCustomFont(r1)
}

func (m *TFPCustomCanvas) CreatePen() IFPCustomPen {
	r1 := fPCustomCanvasImportAPI().SysCallN(12, m.Instance())
	return AsFPCustomPen(r1)
}

func (m *TFPCustomCanvas) CreateBrush() IFPCustomBrush {
	r1 := fPCustomCanvasImportAPI().SysCallN(10, m.Instance())
	return AsFPCustomBrush(r1)
}

func (m *TFPCustomCanvas) GetTextHeight(text string) int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(24, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) GetTextWidth(text string) int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(28, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextExtent(Text string) (resultSize TSize) {
	fPCustomCanvasImportAPI().SysCallN(50, m.Instance(), PascalStr(Text), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TFPCustomCanvas) TextHeight(Text string) int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(52, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextWidth(Text string) int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(56, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func (m *TFPCustomCanvas) GetTextHeight1(text string) int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(25, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) GetTextWidth1(text string) int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(29, m.Instance(), PascalStr(text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextExtent1(Text string) (resultSize TSize) {
	fPCustomCanvasImportAPI().SysCallN(51, m.Instance(), PascalStr(Text), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TFPCustomCanvas) TextHeight1(Text string) int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(53, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func (m *TFPCustomCanvas) TextWidth1(Text string) int32 {
	r1 := fPCustomCanvasImportAPI().SysCallN(57, m.Instance(), PascalStr(Text))
	return int32(r1)
}

func FPCustomCanvasClass() TClass {
	ret := fPCustomCanvasImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TFPCustomCanvas) LockCanvas() {
	fPCustomCanvasImportAPI().SysCallN(36, m.Instance())
}

func (m *TFPCustomCanvas) UnlockCanvas() {
	fPCustomCanvasImportAPI().SysCallN(58, m.Instance())
}

func (m *TFPCustomCanvas) TextOut(x, y int32, text string) {
	fPCustomCanvasImportAPI().SysCallN(54, m.Instance(), uintptr(x), uintptr(y), PascalStr(text))
}

func (m *TFPCustomCanvas) GetTextSize(text string, w, h *int32) {
	var result1 uintptr
	var result2 uintptr
	fPCustomCanvasImportAPI().SysCallN(26, m.Instance(), PascalStr(text), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*w = int32(result1)
	*h = int32(result2)
}

func (m *TFPCustomCanvas) TextOut1(x, y int32, text string) {
	fPCustomCanvasImportAPI().SysCallN(55, m.Instance(), uintptr(x), uintptr(y), PascalStr(text))
}

func (m *TFPCustomCanvas) GetTextSize1(text string, w, h *int32) {
	var result1 uintptr
	var result2 uintptr
	fPCustomCanvasImportAPI().SysCallN(27, m.Instance(), PascalStr(text), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*w = int32(result1)
	*h = int32(result2)
}

func (m *TFPCustomCanvas) Arc(ALeft, ATop, ARight, ABottom, Angle16Deg, Angle16DegLength int32) {
	fPCustomCanvasImportAPI().SysCallN(0, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(ARight), uintptr(ABottom), uintptr(Angle16Deg), uintptr(Angle16DegLength))
}

func (m *TFPCustomCanvas) Arc1(ALeft, ATop, ARight, ABottom, SX, SY, EX, EY int32) {
	fPCustomCanvasImportAPI().SysCallN(1, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(ARight), uintptr(ABottom), uintptr(SX), uintptr(SY), uintptr(EX), uintptr(EY))
}

func (m *TFPCustomCanvas) Ellipse(Bounds *TRect) {
	fPCustomCanvasImportAPI().SysCallN(16, m.Instance(), uintptr(unsafePointer(Bounds)))
}

func (m *TFPCustomCanvas) Ellipse1(left, top, right, bottom int32) {
	fPCustomCanvasImportAPI().SysCallN(17, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

func (m *TFPCustomCanvas) EllipseC(x, y int32, rx, ry uint32) {
	fPCustomCanvasImportAPI().SysCallN(18, m.Instance(), uintptr(x), uintptr(y), uintptr(rx), uintptr(ry))
}

func (m *TFPCustomCanvas) RadialPie(x1, y1, x2, y2, StartAngle16Deg, Angle16DegLength int32) {
	fPCustomCanvasImportAPI().SysCallN(45, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(StartAngle16Deg), uintptr(Angle16DegLength))
}

func (m *TFPCustomCanvas) PolyBezier(Points []TPoint, Filled bool, Continuous bool) {
	sysCallPoint(fPCustomCanvasImportAPI(), 44, m.Instance(), Points, PascalBool(Filled), PascalBool(Continuous))
}

func (m *TFPCustomCanvas) Rectangle(Bounds *TRect) {
	fPCustomCanvasImportAPI().SysCallN(46, m.Instance(), uintptr(unsafePointer(Bounds)))
}

func (m *TFPCustomCanvas) Rectangle1(left, top, right, bottom int32) {
	fPCustomCanvasImportAPI().SysCallN(47, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

func (m *TFPCustomCanvas) FillRect(ARect *TRect) {
	fPCustomCanvasImportAPI().SysCallN(20, m.Instance(), uintptr(unsafePointer(ARect)))
}

func (m *TFPCustomCanvas) FillRect1(X1, Y1, X2, Y2 int32) {
	fPCustomCanvasImportAPI().SysCallN(21, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

func (m *TFPCustomCanvas) FloodFill(x, y int32) {
	fPCustomCanvasImportAPI().SysCallN(22, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TFPCustomCanvas) Clear() {
	fPCustomCanvasImportAPI().SysCallN(4, m.Instance())
}

func (m *TFPCustomCanvas) MoveTo(x, y int32) {
	fPCustomCanvasImportAPI().SysCallN(40, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TFPCustomCanvas) MoveTo1(p *TPoint) {
	fPCustomCanvasImportAPI().SysCallN(41, m.Instance(), uintptr(unsafePointer(p)))
}

func (m *TFPCustomCanvas) LineTo(x, y int32) {
	fPCustomCanvasImportAPI().SysCallN(34, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TFPCustomCanvas) LineTo1(p *TPoint) {
	fPCustomCanvasImportAPI().SysCallN(35, m.Instance(), uintptr(unsafePointer(p)))
}

func (m *TFPCustomCanvas) Line(x1, y1, x2, y2 int32) {
	fPCustomCanvasImportAPI().SysCallN(31, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TFPCustomCanvas) Line1(p1, p2 *TPoint) {
	fPCustomCanvasImportAPI().SysCallN(32, m.Instance(), uintptr(unsafePointer(p1)), uintptr(unsafePointer(p2)))
}

func (m *TFPCustomCanvas) Line2(points *TRect) {
	fPCustomCanvasImportAPI().SysCallN(33, m.Instance(), uintptr(unsafePointer(points)))
}

func (m *TFPCustomCanvas) CopyRect(x, y int32, canvas IFPCustomCanvas, SourceRect *TRect) {
	fPCustomCanvasImportAPI().SysCallN(9, m.Instance(), uintptr(x), uintptr(y), GetObjectUintptr(canvas), uintptr(unsafePointer(SourceRect)))
}

func (m *TFPCustomCanvas) Draw(x, y int32, image IFPCustomImage) {
	fPCustomCanvasImportAPI().SysCallN(13, m.Instance(), uintptr(x), uintptr(y), GetObjectUintptr(image))
}

func (m *TFPCustomCanvas) StretchDraw(x, y, w, h int32, source IFPCustomImage) {
	fPCustomCanvasImportAPI().SysCallN(49, m.Instance(), uintptr(x), uintptr(y), uintptr(w), uintptr(h), GetObjectUintptr(source))
}

func (m *TFPCustomCanvas) Erase() {
	fPCustomCanvasImportAPI().SysCallN(19, m.Instance())
}

func (m *TFPCustomCanvas) DrawPixel(x, y int32, newcolor *TFPColor) {
	fPCustomCanvasImportAPI().SysCallN(14, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(newcolor)))
}

func (m *TFPCustomCanvas) SetOnCombineColors(fn TFPCanvasCombineColors) {
	if m.combineColorsPtr != 0 {
		RemoveEventElement(m.combineColorsPtr)
	}
	m.combineColorsPtr = MakeEventDataPtr(fn)
	fPCustomCanvasImportAPI().SysCallN(48, m.Instance(), m.combineColorsPtr)
}

var (
	fPCustomCanvasImport       *imports.Imports = nil
	fPCustomCanvasImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FPCustomCanvas_Arc", 0),
		/*1*/ imports.NewTable("FPCustomCanvas_Arc1", 0),
		/*2*/ imports.NewTable("FPCustomCanvas_Brush", 0),
		/*3*/ imports.NewTable("FPCustomCanvas_Class", 0),
		/*4*/ imports.NewTable("FPCustomCanvas_Clear", 0),
		/*5*/ imports.NewTable("FPCustomCanvas_ClipRect", 0),
		/*6*/ imports.NewTable("FPCustomCanvas_ClipRegion", 0),
		/*7*/ imports.NewTable("FPCustomCanvas_Clipping", 0),
		/*8*/ imports.NewTable("FPCustomCanvas_Colors", 0),
		/*9*/ imports.NewTable("FPCustomCanvas_CopyRect", 0),
		/*10*/ imports.NewTable("FPCustomCanvas_CreateBrush", 0),
		/*11*/ imports.NewTable("FPCustomCanvas_CreateFont", 0),
		/*12*/ imports.NewTable("FPCustomCanvas_CreatePen", 0),
		/*13*/ imports.NewTable("FPCustomCanvas_Draw", 0),
		/*14*/ imports.NewTable("FPCustomCanvas_DrawPixel", 0),
		/*15*/ imports.NewTable("FPCustomCanvas_DrawingMode", 0),
		/*16*/ imports.NewTable("FPCustomCanvas_Ellipse", 0),
		/*17*/ imports.NewTable("FPCustomCanvas_Ellipse1", 0),
		/*18*/ imports.NewTable("FPCustomCanvas_EllipseC", 0),
		/*19*/ imports.NewTable("FPCustomCanvas_Erase", 0),
		/*20*/ imports.NewTable("FPCustomCanvas_FillRect", 0),
		/*21*/ imports.NewTable("FPCustomCanvas_FillRect1", 0),
		/*22*/ imports.NewTable("FPCustomCanvas_FloodFill", 0),
		/*23*/ imports.NewTable("FPCustomCanvas_Font", 0),
		/*24*/ imports.NewTable("FPCustomCanvas_GetTextHeight", 0),
		/*25*/ imports.NewTable("FPCustomCanvas_GetTextHeight1", 0),
		/*26*/ imports.NewTable("FPCustomCanvas_GetTextSize", 0),
		/*27*/ imports.NewTable("FPCustomCanvas_GetTextSize1", 0),
		/*28*/ imports.NewTable("FPCustomCanvas_GetTextWidth", 0),
		/*29*/ imports.NewTable("FPCustomCanvas_GetTextWidth1", 0),
		/*30*/ imports.NewTable("FPCustomCanvas_Height", 0),
		/*31*/ imports.NewTable("FPCustomCanvas_Line", 0),
		/*32*/ imports.NewTable("FPCustomCanvas_Line1", 0),
		/*33*/ imports.NewTable("FPCustomCanvas_Line2", 0),
		/*34*/ imports.NewTable("FPCustomCanvas_LineTo", 0),
		/*35*/ imports.NewTable("FPCustomCanvas_LineTo1", 0),
		/*36*/ imports.NewTable("FPCustomCanvas_LockCanvas", 0),
		/*37*/ imports.NewTable("FPCustomCanvas_LockCount", 0),
		/*38*/ imports.NewTable("FPCustomCanvas_Locked", 0),
		/*39*/ imports.NewTable("FPCustomCanvas_ManageResources", 0),
		/*40*/ imports.NewTable("FPCustomCanvas_MoveTo", 0),
		/*41*/ imports.NewTable("FPCustomCanvas_MoveTo1", 0),
		/*42*/ imports.NewTable("FPCustomCanvas_Pen", 0),
		/*43*/ imports.NewTable("FPCustomCanvas_PenPos", 0),
		/*44*/ imports.NewTable("FPCustomCanvas_PolyBezier", 0),
		/*45*/ imports.NewTable("FPCustomCanvas_RadialPie", 0),
		/*46*/ imports.NewTable("FPCustomCanvas_Rectangle", 0),
		/*47*/ imports.NewTable("FPCustomCanvas_Rectangle1", 0),
		/*48*/ imports.NewTable("FPCustomCanvas_SetOnCombineColors", 0),
		/*49*/ imports.NewTable("FPCustomCanvas_StretchDraw", 0),
		/*50*/ imports.NewTable("FPCustomCanvas_TextExtent", 0),
		/*51*/ imports.NewTable("FPCustomCanvas_TextExtent1", 0),
		/*52*/ imports.NewTable("FPCustomCanvas_TextHeight", 0),
		/*53*/ imports.NewTable("FPCustomCanvas_TextHeight1", 0),
		/*54*/ imports.NewTable("FPCustomCanvas_TextOut", 0),
		/*55*/ imports.NewTable("FPCustomCanvas_TextOut1", 0),
		/*56*/ imports.NewTable("FPCustomCanvas_TextWidth", 0),
		/*57*/ imports.NewTable("FPCustomCanvas_TextWidth1", 0),
		/*58*/ imports.NewTable("FPCustomCanvas_UnlockCanvas", 0),
		/*59*/ imports.NewTable("FPCustomCanvas_Width", 0),
	}
)

func fPCustomCanvasImportAPI() *imports.Imports {
	if fPCustomCanvasImport == nil {
		fPCustomCanvasImport = NewDefaultImports()
		fPCustomCanvasImport.SetImportTable(fPCustomCanvasImportTables)
		fPCustomCanvasImportTables = nil
	}
	return fPCustomCanvasImport
}
