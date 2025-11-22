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

// ICanvas Parent: IFPCustomCanvas
type ICanvas interface {
	IFPCustomCanvas
	TryLock() bool                                          // function
	GetTextMetrics(outTM *TLCLTextMetric) bool              // function
	TextExtentWithString(text string) types.TSize           // function
	TextHeightWithString(text string) int32                 // function
	TextWidthWithString(text string) int32                  // function
	TextFitInfo(text string, maxWidth int32) int32          // function
	HandleAllocated() bool                                  // function
	GetUpdatedHandle(reqState types.TCanvasState) types.HDC // function
	Lock()                                                  // procedure
	Unlock()                                                // procedure
	Refresh()                                               // procedure
	Changing()                                              // procedure
	Changed()                                               // procedure
	SaveHandleState()                                       // procedure
	RestoreHandleState()                                    // procedure
	// ArcWithIntX6
	//  extra drawing methods (there are more in the ancestor TFPCustomCanvas)
	ArcWithIntX6(left int32, top int32, right int32, bottom int32, angle16Deg int32, angle16DegLength int32)                    // procedure
	ArcWithIntX8(left int32, top int32, right int32, bottom int32, sX int32, sY int32, eX int32, eY int32)                      // procedure
	ArcTo(left int32, top int32, right int32, bottom int32, sX int32, sY int32, eX int32, eY int32)                             // procedure
	AngleArc(X int32, Y int32, radius uint32, startAngle float32, sweepAngle float32)                                           // procedure
	BrushCopy(destRect types.TRect, bitmap IBitmap, sourceRect types.TRect, transparentColor types.TColor)                      // procedure
	ChordWithIntX6(x1 int32, y1 int32, x2 int32, y2 int32, angle16Deg int32, angle16DegLength int32)                            // procedure
	ChordWithIntX8(x1 int32, y1 int32, x2 int32, y2 int32, sX int32, sY int32, eX int32, eY int32)                              // procedure
	CopyRectWithRectX2Canvas(dest types.TRect, srcCanvas ICanvas, source types.TRect)                                           // procedure
	DrawWithIntX2Graphic(X int32, Y int32, srcGraphic IGraphic)                                                                 // procedure
	DrawFocusRect(rect types.TRect)                                                                                             // procedure
	StretchDrawWithRectGraphic(destRect types.TRect, srcGraphic IGraphic)                                                       // procedure
	EllipseWithRect(rect types.TRect)                                                                                           // procedure
	EllipseWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32)                                                                    // procedure
	FillRectWithRect(rect types.TRect)                                                                                          // procedure
	FillRectWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32)                                                                   // procedure
	FloodFillWithIntX2ColorFillStyle(X int32, Y int32, fillColor types.TColor, fillStyle types.TFillStyle)                      // procedure
	Frame3dWithRectIntGraphicsBevelCut(rect *types.TRect, frameWidth int32, style types.TGraphicsBevelCut)                      // procedure
	Frame3DWithRectColorX2Int(rect *types.TRect, topColor types.TColor, bottomColor types.TColor, frameWidth int32)             // procedure
	FrameWithRect(rect types.TRect)                                                                                             // procedure
	FrameWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32)                                                                      // procedure
	FrameRectWithRect(rect types.TRect)                                                                                         // procedure
	FrameRectWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32)                                                                  // procedure
	GradientFill(rect types.TRect, start types.TColor, stop types.TColor, direction types.TGradientDirection)                   // procedure
	RadialPieWithIntX6(x1 int32, y1 int32, x2 int32, y2 int32, startAngle16Deg int32, angle16DegLength int32)                   // procedure
	Pie(ellipseX1 int32, ellipseY1 int32, ellipseX2 int32, ellipseY2 int32, startX int32, startY int32, endX int32, endY int32) // procedure
	PolyBezierWithPointIntBoolX2(points types.PPoint, numPts int32, filled bool, continuous bool)                               // procedure
	Polygon(points types.PPoint, numPts int32, winding bool)                                                                    // procedure
	Polyline(points types.PPoint, numPts int32)                                                                                 // procedure
	RectangleWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32)                                                                  // procedure
	RectangleWithRect(rect types.TRect)                                                                                         // procedure
	RoundRectWithIntX6(x1 int32, y1 int32, x2 int32, y2 int32, rX int32, rY int32)                                              // procedure
	RoundRectWithRectIntX2(rect types.TRect, rX int32, rY int32)                                                                // procedure
	TextOutWithIntX2String(X int32, Y int32, text string)                                                                       // procedure
	TextRectWithRectIntX2String(rect types.TRect, X int32, Y int32, text string)                                                // procedure
	TextRectWithRectIntX2StringTextStyle(rect types.TRect, X int32, Y int32, text string, style TTextStyle)                     // procedure
	Pixels(X int32, Y int32) types.TColor                                                                                       // property Pixels Getter
	SetPixels(X int32, Y int32, value types.TColor)                                                                             // property Pixels Setter
	Handle() types.HDC                                                                                                          // property Handle Getter
	SetHandle(value types.HDC)                                                                                                  // property Handle Setter
	TextStyle() TTextStyle                                                                                                      // property TextStyle Getter
	SetTextStyle(value TTextStyle)                                                                                              // property TextStyle Setter
	AntialiasingMode() types.TAntialiasingMode                                                                                  // property AntialiasingMode Getter
	SetAntialiasingMode(value types.TAntialiasingMode)                                                                          // property AntialiasingMode Setter
	AutoRedraw() bool                                                                                                           // property AutoRedraw Getter
	SetAutoRedraw(value bool)                                                                                                   // property AutoRedraw Setter
	BrushToBrush() IBrush                                                                                                       // property Brush Getter
	SetBrushToBrush(value IBrush)                                                                                               // property Brush Setter
	CopyMode() types.TCopyMode                                                                                                  // property CopyMode Getter
	SetCopyMode(value types.TCopyMode)                                                                                          // property CopyMode Setter
	FontToFont() IFont                                                                                                          // property Font Getter
	SetFontToFont(value IFont)                                                                                                  // property Font Setter
	PenToPen() IPen                                                                                                             // property Pen Getter
	SetPenToPen(value IPen)                                                                                                     // property Pen Setter
	Region() IRegion                                                                                                            // property Region Getter
	SetRegion(value IRegion)                                                                                                    // property Region Setter
	SetOnChange(fn TNotifyEvent)                                                                                                // property event
	SetOnChanging(fn TNotifyEvent)                                                                                              // property event
}

type TCanvas struct {
	TFPCustomCanvas
}

func (m *TCanvas) TryLock() bool {
	if !m.IsValid() {
		return false
	}
	r := canvasAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCanvas) GetTextMetrics(outTM *TLCLTextMetric) bool {
	if !m.IsValid() {
		return false
	}
	r := canvasAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(outTM)))
	return api.GoBool(r)
}

func (m *TCanvas) TextExtentWithString(text string) (result types.TSize) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(3, m.Instance(), api.PasStr(text), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCanvas) TextHeightWithString(text string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := canvasAPI().SysCallN(4, m.Instance(), api.PasStr(text))
	return int32(r)
}

func (m *TCanvas) TextWidthWithString(text string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := canvasAPI().SysCallN(5, m.Instance(), api.PasStr(text))
	return int32(r)
}

func (m *TCanvas) TextFitInfo(text string, maxWidth int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := canvasAPI().SysCallN(6, m.Instance(), api.PasStr(text), uintptr(maxWidth))
	return int32(r)
}

func (m *TCanvas) HandleAllocated() bool {
	if !m.IsValid() {
		return false
	}
	r := canvasAPI().SysCallN(7, m.Instance())
	return api.GoBool(r)
}

func (m *TCanvas) GetUpdatedHandle(reqState types.TCanvasState) types.HDC {
	if !m.IsValid() {
		return 0
	}
	r := canvasAPI().SysCallN(8, m.Instance(), uintptr(reqState))
	return types.HDC(r)
}

func (m *TCanvas) Lock() {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(9, m.Instance())
}

func (m *TCanvas) Unlock() {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(10, m.Instance())
}

func (m *TCanvas) Refresh() {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(11, m.Instance())
}

func (m *TCanvas) Changing() {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(12, m.Instance())
}

func (m *TCanvas) Changed() {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(13, m.Instance())
}

func (m *TCanvas) SaveHandleState() {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(14, m.Instance())
}

func (m *TCanvas) RestoreHandleState() {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(15, m.Instance())
}

func (m *TCanvas) ArcWithIntX6(left int32, top int32, right int32, bottom int32, angle16Deg int32, angle16DegLength int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(16, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(angle16Deg), uintptr(angle16DegLength))
}

func (m *TCanvas) ArcWithIntX8(left int32, top int32, right int32, bottom int32, sX int32, sY int32, eX int32, eY int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(17, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(sX), uintptr(sY), uintptr(eX), uintptr(eY))
}

func (m *TCanvas) ArcTo(left int32, top int32, right int32, bottom int32, sX int32, sY int32, eX int32, eY int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(18, m.Instance(), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(sX), uintptr(sY), uintptr(eX), uintptr(eY))
}

func (m *TCanvas) AngleArc(X int32, Y int32, radius uint32, startAngle float32, sweepAngle float32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(19, m.Instance(), uintptr(X), uintptr(Y), uintptr(radius), uintptr(base.UnsafePointer(&startAngle)), uintptr(base.UnsafePointer(&sweepAngle)))
}

func (m *TCanvas) BrushCopy(destRect types.TRect, bitmap IBitmap, sourceRect types.TRect, transparentColor types.TColor) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(&destRect)), base.GetObjectUintptr(bitmap), uintptr(base.UnsafePointer(&sourceRect)), uintptr(transparentColor))
}

func (m *TCanvas) ChordWithIntX6(x1 int32, y1 int32, x2 int32, y2 int32, angle16Deg int32, angle16DegLength int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(21, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(angle16Deg), uintptr(angle16DegLength))
}

func (m *TCanvas) ChordWithIntX8(x1 int32, y1 int32, x2 int32, y2 int32, sX int32, sY int32, eX int32, eY int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(22, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(sX), uintptr(sY), uintptr(eX), uintptr(eY))
}

func (m *TCanvas) CopyRectWithRectX2Canvas(dest types.TRect, srcCanvas ICanvas, source types.TRect) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(23, m.Instance(), uintptr(base.UnsafePointer(&dest)), base.GetObjectUintptr(srcCanvas), uintptr(base.UnsafePointer(&source)))
}

func (m *TCanvas) DrawWithIntX2Graphic(X int32, Y int32, srcGraphic IGraphic) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(24, m.Instance(), uintptr(X), uintptr(Y), base.GetObjectUintptr(srcGraphic))
}

func (m *TCanvas) DrawFocusRect(rect types.TRect) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(25, m.Instance(), uintptr(base.UnsafePointer(&rect)))
}

func (m *TCanvas) StretchDrawWithRectGraphic(destRect types.TRect, srcGraphic IGraphic) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(26, m.Instance(), uintptr(base.UnsafePointer(&destRect)), base.GetObjectUintptr(srcGraphic))
}

func (m *TCanvas) EllipseWithRect(rect types.TRect) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(27, m.Instance(), uintptr(base.UnsafePointer(&rect)))
}

func (m *TCanvas) EllipseWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(28, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TCanvas) FillRectWithRect(rect types.TRect) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(29, m.Instance(), uintptr(base.UnsafePointer(&rect)))
}

func (m *TCanvas) FillRectWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(30, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TCanvas) FloodFillWithIntX2ColorFillStyle(X int32, Y int32, fillColor types.TColor, fillStyle types.TFillStyle) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(31, m.Instance(), uintptr(X), uintptr(Y), uintptr(fillColor), uintptr(fillStyle))
}

func (m *TCanvas) Frame3dWithRectIntGraphicsBevelCut(rect *types.TRect, frameWidth int32, style types.TGraphicsBevelCut) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(32, m.Instance(), uintptr(base.UnsafePointer(rect)), uintptr(frameWidth), uintptr(style))
}

func (m *TCanvas) Frame3DWithRectColorX2Int(rect *types.TRect, topColor types.TColor, bottomColor types.TColor, frameWidth int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(33, m.Instance(), uintptr(base.UnsafePointer(rect)), uintptr(topColor), uintptr(bottomColor), uintptr(frameWidth))
}

func (m *TCanvas) FrameWithRect(rect types.TRect) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(34, m.Instance(), uintptr(base.UnsafePointer(&rect)))
}

func (m *TCanvas) FrameWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(35, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TCanvas) FrameRectWithRect(rect types.TRect) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(36, m.Instance(), uintptr(base.UnsafePointer(&rect)))
}

func (m *TCanvas) FrameRectWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(37, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TCanvas) GradientFill(rect types.TRect, start types.TColor, stop types.TColor, direction types.TGradientDirection) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(38, m.Instance(), uintptr(base.UnsafePointer(&rect)), uintptr(start), uintptr(stop), uintptr(direction))
}

func (m *TCanvas) RadialPieWithIntX6(x1 int32, y1 int32, x2 int32, y2 int32, startAngle16Deg int32, angle16DegLength int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(39, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(startAngle16Deg), uintptr(angle16DegLength))
}

func (m *TCanvas) Pie(ellipseX1 int32, ellipseY1 int32, ellipseX2 int32, ellipseY2 int32, startX int32, startY int32, endX int32, endY int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(40, m.Instance(), uintptr(ellipseX1), uintptr(ellipseY1), uintptr(ellipseX2), uintptr(ellipseY2), uintptr(startX), uintptr(startY), uintptr(endX), uintptr(endY))
}

func (m *TCanvas) PolyBezierWithPointIntBoolX2(points types.PPoint, numPts int32, filled bool, continuous bool) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(41, m.Instance(), uintptr(points), uintptr(numPts), api.PasBool(filled), api.PasBool(continuous))
}

func (m *TCanvas) Polygon(points types.PPoint, numPts int32, winding bool) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(42, m.Instance(), uintptr(points), uintptr(numPts), api.PasBool(winding))
}

func (m *TCanvas) Polyline(points types.PPoint, numPts int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(43, m.Instance(), uintptr(points), uintptr(numPts))
}

func (m *TCanvas) RectangleWithIntX4(x1 int32, y1 int32, x2 int32, y2 int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(44, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
}

func (m *TCanvas) RectangleWithRect(rect types.TRect) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(45, m.Instance(), uintptr(base.UnsafePointer(&rect)))
}

func (m *TCanvas) RoundRectWithIntX6(x1 int32, y1 int32, x2 int32, y2 int32, rX int32, rY int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(46, m.Instance(), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(rX), uintptr(rY))
}

func (m *TCanvas) RoundRectWithRectIntX2(rect types.TRect, rX int32, rY int32) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(47, m.Instance(), uintptr(base.UnsafePointer(&rect)), uintptr(rX), uintptr(rY))
}

func (m *TCanvas) TextOutWithIntX2String(X int32, Y int32, text string) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(48, m.Instance(), uintptr(X), uintptr(Y), api.PasStr(text))
}

func (m *TCanvas) TextRectWithRectIntX2String(rect types.TRect, X int32, Y int32, text string) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(49, m.Instance(), uintptr(base.UnsafePointer(&rect)), uintptr(X), uintptr(Y), api.PasStr(text))
}

func (m *TCanvas) TextRectWithRectIntX2StringTextStyle(rect types.TRect, X int32, Y int32, text string, style TTextStyle) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(50, m.Instance(), uintptr(base.UnsafePointer(&rect)), uintptr(X), uintptr(Y), api.PasStr(text), uintptr(base.UnsafePointer(&style)))
}

func (m *TCanvas) Pixels(X int32, Y int32) types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := canvasAPI().SysCallN(51, 0, m.Instance(), uintptr(X), uintptr(Y))
	return types.TColor(r)
}

func (m *TCanvas) SetPixels(X int32, Y int32, value types.TColor) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(51, 1, m.Instance(), uintptr(X), uintptr(Y), uintptr(value))
}

func (m *TCanvas) Handle() types.HDC {
	if !m.IsValid() {
		return 0
	}
	r := canvasAPI().SysCallN(52, 0, m.Instance())
	return types.HDC(r)
}

func (m *TCanvas) SetHandle(value types.HDC) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(52, 1, m.Instance(), uintptr(value))
}

func (m *TCanvas) TextStyle() (result TTextStyle) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(53, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCanvas) SetTextStyle(value TTextStyle) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(53, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCanvas) AntialiasingMode() types.TAntialiasingMode {
	if !m.IsValid() {
		return 0
	}
	r := canvasAPI().SysCallN(54, 0, m.Instance())
	return types.TAntialiasingMode(r)
}

func (m *TCanvas) SetAntialiasingMode(value types.TAntialiasingMode) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(54, 1, m.Instance(), uintptr(value))
}

func (m *TCanvas) AutoRedraw() bool {
	if !m.IsValid() {
		return false
	}
	r := canvasAPI().SysCallN(55, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCanvas) SetAutoRedraw(value bool) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(55, 1, m.Instance(), api.PasBool(value))
}

func (m *TCanvas) BrushToBrush() IBrush {
	if !m.IsValid() {
		return nil
	}
	r := canvasAPI().SysCallN(56, 0, m.Instance())
	return AsBrush(r)
}

func (m *TCanvas) SetBrushToBrush(value IBrush) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(56, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCanvas) CopyMode() types.TCopyMode {
	if !m.IsValid() {
		return 0
	}
	r := canvasAPI().SysCallN(57, 0, m.Instance())
	return types.TCopyMode(r)
}

func (m *TCanvas) SetCopyMode(value types.TCopyMode) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(57, 1, m.Instance(), uintptr(value))
}

func (m *TCanvas) FontToFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := canvasAPI().SysCallN(58, 0, m.Instance())
	return AsFont(r)
}

func (m *TCanvas) SetFontToFont(value IFont) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(58, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCanvas) PenToPen() IPen {
	if !m.IsValid() {
		return nil
	}
	r := canvasAPI().SysCallN(59, 0, m.Instance())
	return AsPen(r)
}

func (m *TCanvas) SetPenToPen(value IPen) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(59, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCanvas) Region() IRegion {
	if !m.IsValid() {
		return nil
	}
	r := canvasAPI().SysCallN(60, 0, m.Instance())
	return AsRegion(r)
}

func (m *TCanvas) SetRegion(value IRegion) {
	if !m.IsValid() {
		return
	}
	canvasAPI().SysCallN(60, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCanvas) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 61, canvasAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCanvas) SetOnChanging(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 62, canvasAPI(), api.MakeEventDataPtr(cb))
}

// NewCanvas class constructor
func NewCanvas() ICanvas {
	r := canvasAPI().SysCallN(0)
	return AsCanvas(r)
}

var (
	canvasOnce   base.Once
	canvasImport *imports.Imports = nil
)

func canvasAPI() *imports.Imports {
	canvasOnce.Do(func() {
		canvasImport = api.NewDefaultImports()
		canvasImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCanvas_Create", 0), // constructor NewCanvas
			/* 1 */ imports.NewTable("TCanvas_TryLock", 0), // function TryLock
			/* 2 */ imports.NewTable("TCanvas_GetTextMetrics", 0), // function GetTextMetrics
			/* 3 */ imports.NewTable("TCanvas_TextExtentWithString", 0), // function TextExtentWithString
			/* 4 */ imports.NewTable("TCanvas_TextHeightWithString", 0), // function TextHeightWithString
			/* 5 */ imports.NewTable("TCanvas_TextWidthWithString", 0), // function TextWidthWithString
			/* 6 */ imports.NewTable("TCanvas_TextFitInfo", 0), // function TextFitInfo
			/* 7 */ imports.NewTable("TCanvas_HandleAllocated", 0), // function HandleAllocated
			/* 8 */ imports.NewTable("TCanvas_GetUpdatedHandle", 0), // function GetUpdatedHandle
			/* 9 */ imports.NewTable("TCanvas_Lock", 0), // procedure Lock
			/* 10 */ imports.NewTable("TCanvas_Unlock", 0), // procedure Unlock
			/* 11 */ imports.NewTable("TCanvas_Refresh", 0), // procedure Refresh
			/* 12 */ imports.NewTable("TCanvas_Changing", 0), // procedure Changing
			/* 13 */ imports.NewTable("TCanvas_Changed", 0), // procedure Changed
			/* 14 */ imports.NewTable("TCanvas_SaveHandleState", 0), // procedure SaveHandleState
			/* 15 */ imports.NewTable("TCanvas_RestoreHandleState", 0), // procedure RestoreHandleState
			/* 16 */ imports.NewTable("TCanvas_ArcWithIntX6", 0), // procedure ArcWithIntX6
			/* 17 */ imports.NewTable("TCanvas_ArcWithIntX8", 0), // procedure ArcWithIntX8
			/* 18 */ imports.NewTable("TCanvas_ArcTo", 0), // procedure ArcTo
			/* 19 */ imports.NewTable("TCanvas_AngleArc", 0), // procedure AngleArc
			/* 20 */ imports.NewTable("TCanvas_BrushCopy", 0), // procedure BrushCopy
			/* 21 */ imports.NewTable("TCanvas_ChordWithIntX6", 0), // procedure ChordWithIntX6
			/* 22 */ imports.NewTable("TCanvas_ChordWithIntX8", 0), // procedure ChordWithIntX8
			/* 23 */ imports.NewTable("TCanvas_CopyRectWithRectX2Canvas", 0), // procedure CopyRectWithRectX2Canvas
			/* 24 */ imports.NewTable("TCanvas_DrawWithIntX2Graphic", 0), // procedure DrawWithIntX2Graphic
			/* 25 */ imports.NewTable("TCanvas_DrawFocusRect", 0), // procedure DrawFocusRect
			/* 26 */ imports.NewTable("TCanvas_StretchDrawWithRectGraphic", 0), // procedure StretchDrawWithRectGraphic
			/* 27 */ imports.NewTable("TCanvas_EllipseWithRect", 0), // procedure EllipseWithRect
			/* 28 */ imports.NewTable("TCanvas_EllipseWithIntX4", 0), // procedure EllipseWithIntX4
			/* 29 */ imports.NewTable("TCanvas_FillRectWithRect", 0), // procedure FillRectWithRect
			/* 30 */ imports.NewTable("TCanvas_FillRectWithIntX4", 0), // procedure FillRectWithIntX4
			/* 31 */ imports.NewTable("TCanvas_FloodFillWithIntX2ColorFillStyle", 0), // procedure FloodFillWithIntX2ColorFillStyle
			/* 32 */ imports.NewTable("TCanvas_Frame3dWithRectIntGraphicsBevelCut", 0), // procedure Frame3dWithRectIntGraphicsBevelCut
			/* 33 */ imports.NewTable("TCanvas_Frame3DWithRectColorX2Int", 0), // procedure Frame3DWithRectColorX2Int
			/* 34 */ imports.NewTable("TCanvas_FrameWithRect", 0), // procedure FrameWithRect
			/* 35 */ imports.NewTable("TCanvas_FrameWithIntX4", 0), // procedure FrameWithIntX4
			/* 36 */ imports.NewTable("TCanvas_FrameRectWithRect", 0), // procedure FrameRectWithRect
			/* 37 */ imports.NewTable("TCanvas_FrameRectWithIntX4", 0), // procedure FrameRectWithIntX4
			/* 38 */ imports.NewTable("TCanvas_GradientFill", 0), // procedure GradientFill
			/* 39 */ imports.NewTable("TCanvas_RadialPieWithIntX6", 0), // procedure RadialPieWithIntX6
			/* 40 */ imports.NewTable("TCanvas_Pie", 0), // procedure Pie
			/* 41 */ imports.NewTable("TCanvas_PolyBezierWithPointIntBoolX2", 0), // procedure PolyBezierWithPointIntBoolX2
			/* 42 */ imports.NewTable("TCanvas_Polygon", 0), // procedure Polygon
			/* 43 */ imports.NewTable("TCanvas_Polyline", 0), // procedure Polyline
			/* 44 */ imports.NewTable("TCanvas_RectangleWithIntX4", 0), // procedure RectangleWithIntX4
			/* 45 */ imports.NewTable("TCanvas_RectangleWithRect", 0), // procedure RectangleWithRect
			/* 46 */ imports.NewTable("TCanvas_RoundRectWithIntX6", 0), // procedure RoundRectWithIntX6
			/* 47 */ imports.NewTable("TCanvas_RoundRectWithRectIntX2", 0), // procedure RoundRectWithRectIntX2
			/* 48 */ imports.NewTable("TCanvas_TextOutWithIntX2String", 0), // procedure TextOutWithIntX2String
			/* 49 */ imports.NewTable("TCanvas_TextRectWithRectIntX2String", 0), // procedure TextRectWithRectIntX2String
			/* 50 */ imports.NewTable("TCanvas_TextRectWithRectIntX2StringTextStyle", 0), // procedure TextRectWithRectIntX2StringTextStyle
			/* 51 */ imports.NewTable("TCanvas_Pixels", 0), // property Pixels
			/* 52 */ imports.NewTable("TCanvas_Handle", 0), // property Handle
			/* 53 */ imports.NewTable("TCanvas_TextStyle", 0), // property TextStyle
			/* 54 */ imports.NewTable("TCanvas_AntialiasingMode", 0), // property AntialiasingMode
			/* 55 */ imports.NewTable("TCanvas_AutoRedraw", 0), // property AutoRedraw
			/* 56 */ imports.NewTable("TCanvas_BrushToBrush", 0), // property BrushToBrush
			/* 57 */ imports.NewTable("TCanvas_CopyMode", 0), // property CopyMode
			/* 58 */ imports.NewTable("TCanvas_FontToFont", 0), // property FontToFont
			/* 59 */ imports.NewTable("TCanvas_PenToPen", 0), // property PenToPen
			/* 60 */ imports.NewTable("TCanvas_Region", 0), // property Region
			/* 61 */ imports.NewTable("TCanvas_OnChange", 0), // event OnChange
			/* 62 */ imports.NewTable("TCanvas_OnChanging", 0), // event OnChanging
		}
	})
	return canvasImport
}
