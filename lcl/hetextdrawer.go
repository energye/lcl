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

// IheTextDrawer Parent: IObject
type IheTextDrawer interface {
	IObject
	GetCharWidth() int32                                                                                                // function
	GetCharHeight() int32                                                                                               // function
	NeedsEto() bool                                                                                                     // function
	BeginDrawing(dC types.HDC)                                                                                          // procedure
	EndDrawing()                                                                                                        // procedure
	TextOut(X int32, Y int32, text uintptr, length int32)                                                               // procedure
	ExtTextOut(X int32, Y int32, fuOptions types.UINT, rect types.TRect, text uintptr, length int32, frameBottom int32) // procedure
	NewTextOut(X int32, Y int32, fuOptions types.UINT, rect types.TRect, text uintptr, length int32, anEto IEtoBuffer)  // procedure
	DrawFrame(rect types.TRect)                                                                                         // procedure
	ForceNextTokenWithEto()                                                                                             // procedure
	DrawLine(X int32, Y int32, x2 int32, y2 int32, color types.TColor)                                                  // procedure
	FillRect(rect types.TRect)                                                                                          // procedure
	SetBaseFontWithFont(value IFont)                                                                                    // procedure
	SetBaseStyleWithFontStyles(value types.TFontStyles)                                                                 // procedure
	SetStyleWithFontStyles(value types.TFontStyles)                                                                     // procedure
	SetForeColorWithColor(value types.TColor)                                                                           // procedure
	SetBackColorWithColor(value types.TColor)                                                                           // procedure
	SetFrameColorWithLazSynBorderSideColor(side types.TLazSynBorderSide, value types.TColor)                            // procedure
	SetFrameColorWithColor(value types.TColor)                                                                          // procedure
	SetFrameStyleWithLazSynBorderSideSynLineStyle(side types.TLazSynBorderSide, value types.TSynLineStyle)              // procedure
	// SetCharExtraWithInt
	//  procedure SetFrameStyle(AValue: TSynLineStyle); virtual; overload;
	SetCharExtraWithInt(value int32)                                                                         // procedure
	ReleaseTemporaryResources()                                                                              // procedure
	Eto() IEtoBuffer                                                                                         // property Eto Getter
	CharWidth() int32                                                                                        // property CharWidth Getter
	CharHeight() int32                                                                                       // property CharHeight Getter
	SetBaseFontToFont(value IFont)                                                                           // property BaseFont Setter
	SetBaseStyleToFontStyles(value types.TFontStyles)                                                        // property BaseStyle Setter
	SetForeColorToColor(value types.TColor)                                                                  // property ForeColor Setter
	BackColor() types.TColor                                                                                 // property BackColor Getter
	SetBackColorToColor(value types.TColor)                                                                  // property BackColor Setter
	SetFrameColorWithLazSynBorderSideToColor(side types.TLazSynBorderSide, value types.TColor)               // property FrameColor Setter
	SetFrameStyleWithLazSynBorderSideToSynLineStyle(side types.TLazSynBorderSide, value types.TSynLineStyle) // property FrameStyle Setter
	SetStyleToFontStyles(value types.TFontStyles)                                                            // property Style Setter
	CharExtra() int32                                                                                        // property CharExtra Getter
	SetCharExtraToInt(value int32)                                                                           // property CharExtra Setter
	UseUTF8() bool                                                                                           // property UseUTF8 Getter
	MonoSpace() bool                                                                                         // property MonoSpace Getter
	StockDC() types.HDC                                                                                      // property StockDC Getter
}

type TheTextDrawer struct {
	TObject
}

func (m *TheTextDrawer) GetCharWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := heTextDrawerAPI().SysCallN(1, m.Instance())
	return int32(r)
}

func (m *TheTextDrawer) GetCharHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := heTextDrawerAPI().SysCallN(2, m.Instance())
	return int32(r)
}

func (m *TheTextDrawer) NeedsEto() bool {
	if !m.IsValid() {
		return false
	}
	r := heTextDrawerAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TheTextDrawer) BeginDrawing(dC types.HDC) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(4, m.Instance(), uintptr(dC))
}

func (m *TheTextDrawer) EndDrawing() {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(5, m.Instance())
}

func (m *TheTextDrawer) TextOut(X int32, Y int32, text uintptr, length int32) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(6, m.Instance(), uintptr(X), uintptr(Y), uintptr(text), uintptr(length))
}

func (m *TheTextDrawer) ExtTextOut(X int32, Y int32, fuOptions types.UINT, rect types.TRect, text uintptr, length int32, frameBottom int32) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(7, m.Instance(), uintptr(X), uintptr(Y), uintptr(fuOptions), uintptr(base.UnsafePointer(&rect)), uintptr(text), uintptr(length), uintptr(frameBottom))
}

func (m *TheTextDrawer) NewTextOut(X int32, Y int32, fuOptions types.UINT, rect types.TRect, text uintptr, length int32, anEto IEtoBuffer) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(8, m.Instance(), uintptr(X), uintptr(Y), uintptr(fuOptions), uintptr(base.UnsafePointer(&rect)), uintptr(text), uintptr(length), base.GetObjectUintptr(anEto))
}

func (m *TheTextDrawer) DrawFrame(rect types.TRect) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(9, m.Instance(), uintptr(base.UnsafePointer(&rect)))
}

func (m *TheTextDrawer) ForceNextTokenWithEto() {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(10, m.Instance())
}

func (m *TheTextDrawer) DrawLine(X int32, Y int32, x2 int32, y2 int32, color types.TColor) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(11, m.Instance(), uintptr(X), uintptr(Y), uintptr(x2), uintptr(y2), uintptr(color))
}

func (m *TheTextDrawer) FillRect(rect types.TRect) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&rect)))
}

func (m *TheTextDrawer) SetBaseFontWithFont(value IFont) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(13, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TheTextDrawer) SetBaseStyleWithFontStyles(value types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(14, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) SetStyleWithFontStyles(value types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(15, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) SetForeColorWithColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(16, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) SetBackColorWithColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(17, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) SetFrameColorWithLazSynBorderSideColor(side types.TLazSynBorderSide, value types.TColor) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(18, m.Instance(), uintptr(side), uintptr(value))
}

func (m *TheTextDrawer) SetFrameColorWithColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(19, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) SetFrameStyleWithLazSynBorderSideSynLineStyle(side types.TLazSynBorderSide, value types.TSynLineStyle) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(20, m.Instance(), uintptr(side), uintptr(value))
}

func (m *TheTextDrawer) SetCharExtraWithInt(value int32) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(21, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) ReleaseTemporaryResources() {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(22, m.Instance())
}

func (m *TheTextDrawer) Eto() IEtoBuffer {
	if !m.IsValid() {
		return nil
	}
	r := heTextDrawerAPI().SysCallN(23, m.Instance())
	return AsEtoBuffer(r)
}

func (m *TheTextDrawer) CharWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := heTextDrawerAPI().SysCallN(24, m.Instance())
	return int32(r)
}

func (m *TheTextDrawer) CharHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := heTextDrawerAPI().SysCallN(25, m.Instance())
	return int32(r)
}

func (m *TheTextDrawer) SetBaseFontToFont(value IFont) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(26, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TheTextDrawer) SetBaseStyleToFontStyles(value types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(27, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) SetForeColorToColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(28, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) BackColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := heTextDrawerAPI().SysCallN(29, 0, m.Instance())
	return types.TColor(r)
}

func (m *TheTextDrawer) SetBackColorToColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) SetFrameColorWithLazSynBorderSideToColor(side types.TLazSynBorderSide, value types.TColor) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(30, m.Instance(), uintptr(side), uintptr(value))
}

func (m *TheTextDrawer) SetFrameStyleWithLazSynBorderSideToSynLineStyle(side types.TLazSynBorderSide, value types.TSynLineStyle) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(31, m.Instance(), uintptr(side), uintptr(value))
}

func (m *TheTextDrawer) SetStyleToFontStyles(value types.TFontStyles) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(32, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) CharExtra() int32 {
	if !m.IsValid() {
		return 0
	}
	r := heTextDrawerAPI().SysCallN(33, 0, m.Instance())
	return int32(r)
}

func (m *TheTextDrawer) SetCharExtraToInt(value int32) {
	if !m.IsValid() {
		return
	}
	heTextDrawerAPI().SysCallN(33, 1, m.Instance(), uintptr(value))
}

func (m *TheTextDrawer) UseUTF8() bool {
	if !m.IsValid() {
		return false
	}
	r := heTextDrawerAPI().SysCallN(34, m.Instance())
	return api.GoBool(r)
}

func (m *TheTextDrawer) MonoSpace() bool {
	if !m.IsValid() {
		return false
	}
	r := heTextDrawerAPI().SysCallN(35, m.Instance())
	return api.GoBool(r)
}

func (m *TheTextDrawer) StockDC() types.HDC {
	if !m.IsValid() {
		return 0
	}
	r := heTextDrawerAPI().SysCallN(36, m.Instance())
	return types.HDC(r)
}

// NewHeTextDrawer class constructor
func NewHeTextDrawer(calcExtentBaseStyle types.TFontStyles, baseFont IFont) IheTextDrawer {
	r := heTextDrawerAPI().SysCallN(0, uintptr(calcExtentBaseStyle), base.GetObjectUintptr(baseFont))
	return AsHeTextDrawer(r)
}

var (
	heTextDrawerOnce   base.Once
	heTextDrawerImport *imports.Imports = nil
)

func heTextDrawerAPI() *imports.Imports {
	heTextDrawerOnce.Do(func() {
		heTextDrawerImport = api.NewDefaultImports()
		heTextDrawerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TheTextDrawer_Create", 0), // constructor NewHeTextDrawer
			/* 1 */ imports.NewTable("TheTextDrawer_GetCharWidth", 0), // function GetCharWidth
			/* 2 */ imports.NewTable("TheTextDrawer_GetCharHeight", 0), // function GetCharHeight
			/* 3 */ imports.NewTable("TheTextDrawer_NeedsEto", 0), // function NeedsEto
			/* 4 */ imports.NewTable("TheTextDrawer_BeginDrawing", 0), // procedure BeginDrawing
			/* 5 */ imports.NewTable("TheTextDrawer_EndDrawing", 0), // procedure EndDrawing
			/* 6 */ imports.NewTable("TheTextDrawer_TextOut", 0), // procedure TextOut
			/* 7 */ imports.NewTable("TheTextDrawer_ExtTextOut", 0), // procedure ExtTextOut
			/* 8 */ imports.NewTable("TheTextDrawer_NewTextOut", 0), // procedure NewTextOut
			/* 9 */ imports.NewTable("TheTextDrawer_DrawFrame", 0), // procedure DrawFrame
			/* 10 */ imports.NewTable("TheTextDrawer_ForceNextTokenWithEto", 0), // procedure ForceNextTokenWithEto
			/* 11 */ imports.NewTable("TheTextDrawer_DrawLine", 0), // procedure DrawLine
			/* 12 */ imports.NewTable("TheTextDrawer_FillRect", 0), // procedure FillRect
			/* 13 */ imports.NewTable("TheTextDrawer_SetBaseFontWithFont", 0), // procedure SetBaseFontWithFont
			/* 14 */ imports.NewTable("TheTextDrawer_SetBaseStyleWithFontStyles", 0), // procedure SetBaseStyleWithFontStyles
			/* 15 */ imports.NewTable("TheTextDrawer_SetStyleWithFontStyles", 0), // procedure SetStyleWithFontStyles
			/* 16 */ imports.NewTable("TheTextDrawer_SetForeColorWithColor", 0), // procedure SetForeColorWithColor
			/* 17 */ imports.NewTable("TheTextDrawer_SetBackColorWithColor", 0), // procedure SetBackColorWithColor
			/* 18 */ imports.NewTable("TheTextDrawer_SetFrameColorWithLazSynBorderSideColor", 0), // procedure SetFrameColorWithLazSynBorderSideColor
			/* 19 */ imports.NewTable("TheTextDrawer_SetFrameColorWithColor", 0), // procedure SetFrameColorWithColor
			/* 20 */ imports.NewTable("TheTextDrawer_SetFrameStyleWithLazSynBorderSideSynLineStyle", 0), // procedure SetFrameStyleWithLazSynBorderSideSynLineStyle
			/* 21 */ imports.NewTable("TheTextDrawer_SetCharExtraWithInt", 0), // procedure SetCharExtraWithInt
			/* 22 */ imports.NewTable("TheTextDrawer_ReleaseTemporaryResources", 0), // procedure ReleaseTemporaryResources
			/* 23 */ imports.NewTable("TheTextDrawer_Eto", 0), // property Eto
			/* 24 */ imports.NewTable("TheTextDrawer_CharWidth", 0), // property CharWidth
			/* 25 */ imports.NewTable("TheTextDrawer_CharHeight", 0), // property CharHeight
			/* 26 */ imports.NewTable("TheTextDrawer_BaseFont", 0), // property BaseFont
			/* 27 */ imports.NewTable("TheTextDrawer_BaseStyle", 0), // property BaseStyle
			/* 28 */ imports.NewTable("TheTextDrawer_ForeColor", 0), // property ForeColor
			/* 29 */ imports.NewTable("TheTextDrawer_BackColor", 0), // property BackColor
			/* 30 */ imports.NewTable("TheTextDrawer_FrameColor", 0), // property FrameColor
			/* 31 */ imports.NewTable("TheTextDrawer_FrameStyle", 0), // property FrameStyle
			/* 32 */ imports.NewTable("TheTextDrawer_Style", 0), // property Style
			/* 33 */ imports.NewTable("TheTextDrawer_CharExtra", 0), // property CharExtra
			/* 34 */ imports.NewTable("TheTextDrawer_UseUTF8", 0), // property UseUTF8
			/* 35 */ imports.NewTable("TheTextDrawer_MonoSpace", 0), // property MonoSpace
			/* 36 */ imports.NewTable("TheTextDrawer_StockDC", 0), // property StockDC
		}
	})
	return heTextDrawerImport
}
