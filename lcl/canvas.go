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

// ICanvas Parent: IFPCustomCanvas
type ICanvas interface {
	IFPCustomCanvas
	TextRect2(aRect *TRect, text string, textFormat TTextFormat)
	Pixels(X, Y int32) TColor                                                                  // property
	SetPixels(X, Y int32, AValue TColor)                                                       // property
	Handle() HDC                                                                               // property
	SetHandle(AValue HDC)                                                                      // property
	TextStyle() (resultTextStyle TTextStyle)                                                   // property
	SetTextStyle(AValue *TTextStyle)                                                           // property
	AntialiasingMode() TAntialiasingMode                                                       // property
	SetAntialiasingMode(AValue TAntialiasingMode)                                              // property
	AutoRedraw() bool                                                                          // property
	SetAutoRedraw(AValue bool)                                                                 // property
	BrushForBrush() IBrush                                                                     // property
	SetBrushForBrush(AValue IBrush)                                                            // property
	CopyMode() TCopyMode                                                                       // property
	SetCopyMode(AValue TCopyMode)                                                              // property
	FontForFont() IFont                                                                        // property
	SetFontForFont(AValue IFont)                                                               // property
	PenForPen() IPen                                                                           // property
	SetPenForPen(AValue IPen)                                                                  // property
	Region() IRegion                                                                           // property
	SetRegion(AValue IRegion)                                                                  // property
	TryLock() bool                                                                             // function
	GetTextMetrics(OutTM *TLCLTextMetric) bool                                                 // function
	TextFitInfo(Text string, MaxWidth int32) int32                                             // function
	HandleAllocated() bool                                                                     // function
	GetUpdatedHandle(ReqState TCanvasState) HDC                                                // function
	Lock()                                                                                     // procedure
	Unlock()                                                                                   // procedure
	Refresh()                                                                                  // procedure
	Changing()                                                                                 // procedure
	Changed()                                                                                  // procedure
	SaveHandleState()                                                                          // procedure
	RestoreHandleState()                                                                       // procedure
	ArcTo(ALeft, ATop, ARight, ABottom, SX, SY, EX, EY int32)                                  // procedure
	AngleArc(X, Y int32, Radius uint32, StartAngle, SweepAngle float32)                        // procedure
	BrushCopy(ADestRect *TRect, ABitmap IBitmap, ASourceRect *TRect, ATransparentColor TColor) // procedure
	Chord(x1, y1, x2, y2, Angle16Deg, Angle16DegLength int32)                                  // procedure
	Chord1(x1, y1, x2, y2, SX, SY, EX, EY int32)                                               // procedure
	CopyRectForCanvas(Dest *TRect, SrcCanvas ICanvas, Source *TRect)                           // procedure
	DrawForGraphic(X, Y int32, SrcGraphic IGraphic)                                            // procedure
	DrawFocusRect(ARect *TRect)                                                                // procedure
	StretchDrawForGraphic(DestRect *TRect, SrcGraphic IGraphic)                                // procedure
	FloodFillForInteger(X, Y int32, FillColor TColor, FillStyle TFillStyle)                    // procedure
	Frame3d(ARect *TRect, FrameWidth int32, Style TGraphicsBevelCut)                           // procedure
	Frame3D1(ARect *TRect, TopColor, BottomColor TColor, FrameWidth int32)                     // procedure
	Frame(ARect *TRect)                                                                        // procedure
	Frame1(X1, Y1, X2, Y2 int32)                                                               // procedure
	FrameRect(ARect *TRect)                                                                    // procedure
	FrameRect1(X1, Y1, X2, Y2 int32)                                                           // procedure
	GradientFill(ARect *TRect, AStart, AStop TColor, ADirection TGradientDirection)            // procedure
	Pie(EllipseX1, EllipseY1, EllipseX2, EllipseY2, StartX, StartY, EndX, EndY int32)          // procedure
	Polygon(Points []TPoint, Winding bool)                                                     // procedure
	Polyline(Points []TPoint)                                                                  // procedure
	RoundRect(X1, Y1, X2, Y2 int32, RX, RY int32)                                              // procedure
	RoundRect1(Rect *TRect, RX, RY int32)                                                      // procedure
	TextRect(ARect *TRect, X, Y int32, Text string)                                            // procedure
	TextRect1(ARect *TRect, X, Y int32, Text string, Style *TTextStyle)                        // procedure
	SetOnChange(fn TNotifyEvent)                                                               // property event
	SetOnChanging(fn TNotifyEvent)                                                             // property event
}

// TCanvas Parent: TFPCustomCanvas
type TCanvas struct {
	TFPCustomCanvas
	changePtr   uintptr
	changingPtr uintptr
}

func NewCanvas() ICanvas {
	r1 := canvasImportAPI().SysCallN(13)
	return AsCanvas(r1)
}

func (m *TCanvas) Pixels(X, Y int32) TColor {
	r1 := canvasImportAPI().SysCallN(32, 0, m.Instance(), uintptr(X), uintptr(Y))
	return TColor(r1)
}

func (m *TCanvas) SetPixels(X, Y int32, AValue TColor) {
	canvasImportAPI().SysCallN(32, 1, m.Instance(), uintptr(X), uintptr(Y), uintptr(AValue))
}

func (m *TCanvas) Handle() HDC {
	r1 := canvasImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return HDC(r1)
}

func (m *TCanvas) SetHandle(AValue HDC) {
	canvasImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func (m *TCanvas) TextStyle() (resultTextStyle TTextStyle) {
	canvasImportAPI().SysCallN(47, 0, m.Instance(), uintptr(unsafePointer(&resultTextStyle)), uintptr(unsafePointer(&resultTextStyle)))
	return
}

func (m *TCanvas) SetTextStyle(AValue *TTextStyle) {
	canvasImportAPI().SysCallN(47, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCanvas) AntialiasingMode() TAntialiasingMode {
	r1 := canvasImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TAntialiasingMode(r1)
}

func (m *TCanvas) SetAntialiasingMode(AValue TAntialiasingMode) {
	canvasImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCanvas) AutoRedraw() bool {
	r1 := canvasImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCanvas) SetAutoRedraw(AValue bool) {
	canvasImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCanvas) BrushForBrush() IBrush {
	r1 := canvasImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsBrush(r1)
}

func (m *TCanvas) SetBrushForBrush(AValue IBrush) {
	canvasImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCanvas) CopyMode() TCopyMode {
	r1 := canvasImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TCopyMode(r1)
}

func (m *TCanvas) SetCopyMode(AValue TCopyMode) {
	canvasImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCanvas) FontForFont() IFont {
	r1 := canvasImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TCanvas) SetFontForFont(AValue IFont) {
	canvasImportAPI().SysCallN(17, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCanvas) PenForPen() IPen {
	r1 := canvasImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return AsPen(r1)
}

func (m *TCanvas) SetPenForPen(AValue IPen) {
	canvasImportAPI().SysCallN(30, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCanvas) Region() IRegion {
	r1 := canvasImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return AsRegion(r1)
}

func (m *TCanvas) SetRegion(AValue IRegion) {
	canvasImportAPI().SysCallN(36, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCanvas) TryLock() bool {
	r1 := canvasImportAPI().SysCallN(48, m.Instance())
	return GoBool(r1)
}

func (m *TCanvas) GetTextMetrics(OutTM *TLCLTextMetric) bool {
	var result0 uintptr
	r1 := canvasImportAPI().SysCallN(24, m.Instance(), uintptr(unsafePointer(&result0)))
	*OutTM = *(*TLCLTextMetric)(getPointer(result0))
	return GoBool(r1)
}

func (m *TCanvas) TextFitInfo(Text string, MaxWidth int32) int32 {
	r1 := canvasImportAPI().SysCallN(44, m.Instance(), PascalStr(Text), uintptr(MaxWidth))
	return int32(r1)
}

func (m *TCanvas) HandleAllocated() bool {
	r1 := canvasImportAPI().SysCallN(28, m.Instance())
	return GoBool(r1)
}

func (m *TCanvas) GetUpdatedHandle(ReqState TCanvasState) HDC {
	r1 := canvasImportAPI().SysCallN(25, m.Instance(), uintptr(ReqState))
	return HDC(r1)
}

func CanvasClass() TClass {
	ret := canvasImportAPI().SysCallN(10)
	return TClass(ret)
}

func (m *TCanvas) Lock() {
	canvasImportAPI().SysCallN(29, m.Instance())
}

func (m *TCanvas) Unlock() {
	canvasImportAPI().SysCallN(49, m.Instance())
}

func (m *TCanvas) Refresh() {
	canvasImportAPI().SysCallN(35, m.Instance())
}

func (m *TCanvas) Changing() {
	canvasImportAPI().SysCallN(7, m.Instance())
}

func (m *TCanvas) Changed() {
	canvasImportAPI().SysCallN(6, m.Instance())
}

func (m *TCanvas) SaveHandleState() {
	canvasImportAPI().SysCallN(40, m.Instance())
}

func (m *TCanvas) RestoreHandleState() {
	canvasImportAPI().SysCallN(37, m.Instance())
}

func (m *TCanvas) ArcTo(ALeft, ATop, ARight, ABottom, SX, SY, EX, EY int32) {
	canvasImportAPI().SysCallN(2, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(ARight), uintptr(ABottom), uintptr(SX), uintptr(SY), uintptr(EX), uintptr(EY))
}

func (m *TCanvas) AngleArc(X, Y int32, Radius uint32, StartAngle, SweepAngle float32) {
	canvasImportAPI().SysCallN(0, m.Instance(), uintptr(X), uintptr(Y), uintptr(Radius), uintptr(StartAngle), uintptr(SweepAngle))
}

func (m *TCanvas) BrushCopy(ADestRect *TRect, ABitmap IBitmap, ASourceRect *TRect, ATransparentColor TColor) {
	canvasImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(ADestRect)), GetObjectUintptr(ABitmap), uintptr(unsafePointer(ASourceRect)), uintptr(ATransparentColor))
}

func (m *TCanvas) Chord(x1, y1, x2, y2, Angle16Deg, Angle16DegLength int32) {
	canvasImportAPI().SysCallN(8, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(Angle16Deg), uintptr(Angle16DegLength))
}

func (m *TCanvas) Chord1(x1, y1, x2, y2, SX, SY, EX, EY int32) {
	canvasImportAPI().SysCallN(9, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(SX), uintptr(SY), uintptr(EX), uintptr(EY))
}

func (m *TCanvas) CopyRectForCanvas(Dest *TRect, SrcCanvas ICanvas, Source *TRect) {
	canvasImportAPI().SysCallN(12, m.Instance(), uintptr(unsafePointer(Dest)), GetObjectUintptr(SrcCanvas), uintptr(unsafePointer(Source)))
}

func (m *TCanvas) DrawForGraphic(X, Y int32, SrcGraphic IGraphic) {
	canvasImportAPI().SysCallN(15, m.Instance(), uintptr(X), uintptr(Y), GetObjectUintptr(SrcGraphic))
}

func (m *TCanvas) DrawFocusRect(ARect *TRect) {
	canvasImportAPI().SysCallN(14, m.Instance(), uintptr(unsafePointer(ARect)))
}

func (m *TCanvas) StretchDrawForGraphic(DestRect *TRect, SrcGraphic IGraphic) {
	canvasImportAPI().SysCallN(43, m.Instance(), uintptr(unsafePointer(DestRect)), GetObjectUintptr(SrcGraphic))
}

func (m *TCanvas) FloodFillForInteger(X, Y int32, FillColor TColor, FillStyle TFillStyle) {
	canvasImportAPI().SysCallN(16, m.Instance(), uintptr(X), uintptr(Y), uintptr(FillColor), uintptr(FillStyle))
}

func (m *TCanvas) Frame3d(ARect *TRect, FrameWidth int32, Style TGraphicsBevelCut) {
	var result0 uintptr
	canvasImportAPI().SysCallN(21, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(FrameWidth), uintptr(Style))
	*ARect = *(*TRect)(getPointer(result0))
}

func (m *TCanvas) Frame3D1(ARect *TRect, TopColor, BottomColor TColor, FrameWidth int32) {
	var result0 uintptr
	canvasImportAPI().SysCallN(20, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(TopColor), uintptr(BottomColor), uintptr(FrameWidth))
	*ARect = *(*TRect)(getPointer(result0))
}

func (m *TCanvas) Frame(ARect *TRect) {
	canvasImportAPI().SysCallN(18, m.Instance(), uintptr(unsafePointer(ARect)))
}

func (m *TCanvas) Frame1(X1, Y1, X2, Y2 int32) {
	canvasImportAPI().SysCallN(19, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

func (m *TCanvas) FrameRect(ARect *TRect) {
	canvasImportAPI().SysCallN(22, m.Instance(), uintptr(unsafePointer(ARect)))
}

func (m *TCanvas) FrameRect1(X1, Y1, X2, Y2 int32) {
	canvasImportAPI().SysCallN(23, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}

func (m *TCanvas) GradientFill(ARect *TRect, AStart, AStop TColor, ADirection TGradientDirection) {
	canvasImportAPI().SysCallN(26, m.Instance(), uintptr(unsafePointer(ARect)), uintptr(AStart), uintptr(AStop), uintptr(ADirection))
}

func (m *TCanvas) Pie(EllipseX1, EllipseY1, EllipseX2, EllipseY2, StartX, StartY, EndX, EndY int32) {
	canvasImportAPI().SysCallN(31, m.Instance(), uintptr(EllipseX1), uintptr(EllipseY1), uintptr(EllipseX2), uintptr(EllipseY2), uintptr(StartX), uintptr(StartY), uintptr(EndX), uintptr(EndY))
}

func (m *TCanvas) Polygon(Points []TPoint, Winding bool) {
	sysCallPoint(canvasImportAPI(), 33, m.Instance(), Points, PascalBool(Winding))
}

func (m *TCanvas) Polyline(Points []TPoint) {
	sysCallPoint(canvasImportAPI(), 34, m.Instance(), Points)
}

func (m *TCanvas) RoundRect(X1, Y1, X2, Y2 int32, RX, RY int32) {
	canvasImportAPI().SysCallN(38, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(RX), uintptr(RY))
}

func (m *TCanvas) RoundRect1(Rect *TRect, RX, RY int32) {
	canvasImportAPI().SysCallN(39, m.Instance(), uintptr(unsafePointer(Rect)), uintptr(RX), uintptr(RY))
}

func (m *TCanvas) TextRect(ARect *TRect, X, Y int32, Text string) {
	canvasImportAPI().SysCallN(45, m.Instance(), uintptr(unsafePointer(ARect)), uintptr(X), uintptr(Y), PascalStr(Text))
}

func (m *TCanvas) TextRect1(ARect *TRect, X, Y int32, Text string, Style *TTextStyle) {
	canvasImportAPI().SysCallN(46, m.Instance(), uintptr(unsafePointer(ARect)), uintptr(X), uintptr(Y), PascalStr(Text), uintptr(unsafePointer(Style)))
}

func (m *TCanvas) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	canvasImportAPI().SysCallN(41, m.Instance(), m.changePtr)
}

func (m *TCanvas) SetOnChanging(fn TNotifyEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	canvasImportAPI().SysCallN(42, m.Instance(), m.changingPtr)
}

var (
	canvasImport       *imports.Imports = nil
	canvasImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Canvas_AngleArc", 0),
		/*1*/ imports.NewTable("Canvas_AntialiasingMode", 0),
		/*2*/ imports.NewTable("Canvas_ArcTo", 0),
		/*3*/ imports.NewTable("Canvas_AutoRedraw", 0),
		/*4*/ imports.NewTable("Canvas_BrushCopy", 0),
		/*5*/ imports.NewTable("Canvas_BrushForBrush", 0),
		/*6*/ imports.NewTable("Canvas_Changed", 0),
		/*7*/ imports.NewTable("Canvas_Changing", 0),
		/*8*/ imports.NewTable("Canvas_Chord", 0),
		/*9*/ imports.NewTable("Canvas_Chord1", 0),
		/*10*/ imports.NewTable("Canvas_Class", 0),
		/*11*/ imports.NewTable("Canvas_CopyMode", 0),
		/*12*/ imports.NewTable("Canvas_CopyRectForCanvas", 0),
		/*13*/ imports.NewTable("Canvas_Create", 0),
		/*14*/ imports.NewTable("Canvas_DrawFocusRect", 0),
		/*15*/ imports.NewTable("Canvas_DrawForGraphic", 0),
		/*16*/ imports.NewTable("Canvas_FloodFillForInteger", 0),
		/*17*/ imports.NewTable("Canvas_FontForFont", 0),
		/*18*/ imports.NewTable("Canvas_Frame", 0),
		/*19*/ imports.NewTable("Canvas_Frame1", 0),
		/*20*/ imports.NewTable("Canvas_Frame3D1", 0),
		/*21*/ imports.NewTable("Canvas_Frame3d", 0),
		/*22*/ imports.NewTable("Canvas_FrameRect", 0),
		/*23*/ imports.NewTable("Canvas_FrameRect1", 0),
		/*24*/ imports.NewTable("Canvas_GetTextMetrics", 0),
		/*25*/ imports.NewTable("Canvas_GetUpdatedHandle", 0),
		/*26*/ imports.NewTable("Canvas_GradientFill", 0),
		/*27*/ imports.NewTable("Canvas_Handle", 0),
		/*28*/ imports.NewTable("Canvas_HandleAllocated", 0),
		/*29*/ imports.NewTable("Canvas_Lock", 0),
		/*30*/ imports.NewTable("Canvas_PenForPen", 0),
		/*31*/ imports.NewTable("Canvas_Pie", 0),
		/*32*/ imports.NewTable("Canvas_Pixels", 0),
		/*33*/ imports.NewTable("Canvas_Polygon", 0),
		/*34*/ imports.NewTable("Canvas_Polyline", 0),
		/*35*/ imports.NewTable("Canvas_Refresh", 0),
		/*36*/ imports.NewTable("Canvas_Region", 0),
		/*37*/ imports.NewTable("Canvas_RestoreHandleState", 0),
		/*38*/ imports.NewTable("Canvas_RoundRect", 0),
		/*39*/ imports.NewTable("Canvas_RoundRect1", 0),
		/*40*/ imports.NewTable("Canvas_SaveHandleState", 0),
		/*41*/ imports.NewTable("Canvas_SetOnChange", 0),
		/*42*/ imports.NewTable("Canvas_SetOnChanging", 0),
		/*43*/ imports.NewTable("Canvas_StretchDrawForGraphic", 0),
		/*44*/ imports.NewTable("Canvas_TextFitInfo", 0),
		/*45*/ imports.NewTable("Canvas_TextRect", 0),
		/*46*/ imports.NewTable("Canvas_TextRect1", 0),
		/*47*/ imports.NewTable("Canvas_TextStyle", 0),
		/*48*/ imports.NewTable("Canvas_TryLock", 0),
		/*49*/ imports.NewTable("Canvas_Unlock", 0),
	}
)

func canvasImportAPI() *imports.Imports {
	if canvasImport == nil {
		canvasImport = NewDefaultImports()
		canvasImport.SetImportTable(canvasImportTables)
		canvasImportTables = nil
	}
	return canvasImport
}
