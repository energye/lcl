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

// ILazSynTextArea Parent: ILazSynSurface
type ILazSynTextArea interface {
	ILazSynSurface
	ScreenColumnToXValue(col int32) int32                                                       // function
	RowColumnToPixels(rowCol types.TPoint) types.TPoint                                         // function
	PixelsToRowColumn(pixels types.TPoint, flags types.TSynCoordinateMappingFlags) types.TPoint // function
	FontChanged()                                                                               // procedure
	// Padding
	//  Settings controlled by SynEdit
	Padding(side types.TLazSynBorderSide) int32                 // property Padding Getter
	SetPadding(side types.TLazSynBorderSide, value int32)       // property Padding Setter
	ForegroundColor() types.TColor                              // property ForegroundColor Getter
	SetForegroundColor(value types.TColor)                      // property ForegroundColor Setter
	BackgroundColor() types.TColor                              // property BackgroundColor Getter
	SetBackgroundColor(value types.TColor)                      // property BackgroundColor Setter
	ExtraCharSpacing() int32                                    // property ExtraCharSpacing Getter
	SetExtraCharSpacing(value int32)                            // property ExtraCharSpacing Setter
	ExtraLineSpacing() int32                                    // property ExtraLineSpacing Getter
	SetExtraLineSpacing(value int32)                            // property ExtraLineSpacing Setter
	VisibleSpecialChars() types.TSynVisibleSpecialChars         // property VisibleSpecialChars Getter
	SetVisibleSpecialChars(value types.TSynVisibleSpecialChars) // property VisibleSpecialChars Setter
	RightEdgeColumn() int32                                     // property RightEdgeColumn Getter
	SetRightEdgeColumn(value int32)                             // property RightEdgeColumn Setter
	RightEdgeVisible() bool                                     // property RightEdgeVisible Getter
	SetRightEdgeVisible(value bool)                             // property RightEdgeVisible Setter
	RightEdgeColor() types.TColor                               // property RightEdgeColor Getter
	SetRightEdgeColor(value types.TColor)                       // property RightEdgeColor Setter
	TopLine() types.TLinePos                                    // property TopLine Getter
	SetTopLine(value types.TLinePos)                            // property TopLine Setter
	LeftChar() int32                                            // property LeftChar Getter
	SetLeftChar(value int32)                                    // property LeftChar Setter
	TheLinesView() ISynEditStrings                              // property TheLinesView Getter
	SetTheLinesView(value ISynEditStrings)                      // property TheLinesView Setter
	Highlighter() ISynCustomHighlighter                         // property Highlighter Getter
	SetHighlighter(value ISynCustomHighlighter)                 // property Highlighter Setter
	MarkupManager() ISynEditMarkupManager                       // property MarkupManager Getter
	SetMarkupManager(value ISynEditMarkupManager)               // property MarkupManager Setter
	TextDrawer() IheTextDrawer                                  // property TextDrawer Getter
	TextBounds() types.TRect                                    // property TextBounds Getter
	LineHeight() int32                                          // property LineHeight Getter
	CharWidth() int32                                           // property CharWidth Getter
	LinesInWindow() int32                                       // property LinesInWindow Getter
	CharsInWindow() int32                                       // property CharsInWindow Getter
}

type TLazSynTextArea struct {
	TLazSynSurface
}

func (m *TLazSynTextArea) ScreenColumnToXValue(col int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(1, m.Instance(), uintptr(col))
	return int32(r)
}

func (m *TLazSynTextArea) RowColumnToPixels(rowCol types.TPoint) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&rowCol)), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynTextArea) PixelsToRowColumn(pixels types.TPoint, flags types.TSynCoordinateMappingFlags) (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&pixels)), uintptr(flags), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynTextArea) FontChanged() {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(4, m.Instance())
}

func (m *TLazSynTextArea) Padding(side types.TLazSynBorderSide) int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(5, 0, m.Instance(), uintptr(side))
	return int32(r)
}

func (m *TLazSynTextArea) SetPadding(side types.TLazSynBorderSide, value int32) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(5, 1, m.Instance(), uintptr(side), uintptr(value))
}

func (m *TLazSynTextArea) ForegroundColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(6, 0, m.Instance())
	return types.TColor(r)
}

func (m *TLazSynTextArea) SetForegroundColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynTextArea) BackgroundColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(7, 0, m.Instance())
	return types.TColor(r)
}

func (m *TLazSynTextArea) SetBackgroundColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynTextArea) ExtraCharSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynTextArea) SetExtraCharSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynTextArea) ExtraLineSpacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynTextArea) SetExtraLineSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynTextArea) VisibleSpecialChars() types.TSynVisibleSpecialChars {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(10, 0, m.Instance())
	return types.TSynVisibleSpecialChars(r)
}

func (m *TLazSynTextArea) SetVisibleSpecialChars(value types.TSynVisibleSpecialChars) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynTextArea) RightEdgeColumn() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynTextArea) SetRightEdgeColumn(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynTextArea) RightEdgeVisible() bool {
	if !m.IsValid() {
		return false
	}
	r := lazSynTextAreaAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TLazSynTextArea) SetRightEdgeVisible(value bool) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TLazSynTextArea) RightEdgeColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(13, 0, m.Instance())
	return types.TColor(r)
}

func (m *TLazSynTextArea) SetRightEdgeColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynTextArea) TopLine() types.TLinePos {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(14, 0, m.Instance())
	return types.TLinePos(r)
}

func (m *TLazSynTextArea) SetTopLine(value types.TLinePos) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynTextArea) LeftChar() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TLazSynTextArea) SetLeftChar(value int32) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TLazSynTextArea) TheLinesView() ISynEditStrings {
	if !m.IsValid() {
		return nil
	}
	r := lazSynTextAreaAPI().SysCallN(16, 0, m.Instance())
	return AsSynEditStrings(r)
}

func (m *TLazSynTextArea) SetTheLinesView(value ISynEditStrings) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazSynTextArea) Highlighter() ISynCustomHighlighter {
	if !m.IsValid() {
		return nil
	}
	r := lazSynTextAreaAPI().SysCallN(17, 0, m.Instance())
	return AsSynCustomHighlighter(r)
}

func (m *TLazSynTextArea) SetHighlighter(value ISynCustomHighlighter) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(17, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazSynTextArea) MarkupManager() ISynEditMarkupManager {
	if !m.IsValid() {
		return nil
	}
	r := lazSynTextAreaAPI().SysCallN(18, 0, m.Instance())
	return AsSynEditMarkupManager(r)
}

func (m *TLazSynTextArea) SetMarkupManager(value ISynEditMarkupManager) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(18, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TLazSynTextArea) TextDrawer() IheTextDrawer {
	if !m.IsValid() {
		return nil
	}
	r := lazSynTextAreaAPI().SysCallN(19, m.Instance())
	return AsHeTextDrawer(r)
}

func (m *TLazSynTextArea) TextBounds() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	lazSynTextAreaAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TLazSynTextArea) LineHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(21, m.Instance())
	return int32(r)
}

func (m *TLazSynTextArea) CharWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(22, m.Instance())
	return int32(r)
}

func (m *TLazSynTextArea) LinesInWindow() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(23, m.Instance())
	return int32(r)
}

func (m *TLazSynTextArea) CharsInWindow() int32 {
	if !m.IsValid() {
		return 0
	}
	r := lazSynTextAreaAPI().SysCallN(24, m.Instance())
	return int32(r)
}

// NewLazSynTextArea class constructor
func NewLazSynTextArea(owner IWinControl, textDrawer IheTextDrawer) ILazSynTextArea {
	r := lazSynTextAreaAPI().SysCallN(0, base.GetObjectUintptr(owner), base.GetObjectUintptr(textDrawer))
	return AsLazSynTextArea(r)
}

var (
	lazSynTextAreaOnce   base.Once
	lazSynTextAreaImport *imports.Imports = nil
)

func lazSynTextAreaAPI() *imports.Imports {
	lazSynTextAreaOnce.Do(func() {
		lazSynTextAreaImport = api.NewDefaultImports()
		lazSynTextAreaImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TLazSynTextArea_Create", 0), // constructor NewLazSynTextArea
			/* 1 */ imports.NewTable("TLazSynTextArea_ScreenColumnToXValue", 0), // function ScreenColumnToXValue
			/* 2 */ imports.NewTable("TLazSynTextArea_RowColumnToPixels", 0), // function RowColumnToPixels
			/* 3 */ imports.NewTable("TLazSynTextArea_PixelsToRowColumn", 0), // function PixelsToRowColumn
			/* 4 */ imports.NewTable("TLazSynTextArea_FontChanged", 0), // procedure FontChanged
			/* 5 */ imports.NewTable("TLazSynTextArea_Padding", 0), // property Padding
			/* 6 */ imports.NewTable("TLazSynTextArea_ForegroundColor", 0), // property ForegroundColor
			/* 7 */ imports.NewTable("TLazSynTextArea_BackgroundColor", 0), // property BackgroundColor
			/* 8 */ imports.NewTable("TLazSynTextArea_ExtraCharSpacing", 0), // property ExtraCharSpacing
			/* 9 */ imports.NewTable("TLazSynTextArea_ExtraLineSpacing", 0), // property ExtraLineSpacing
			/* 10 */ imports.NewTable("TLazSynTextArea_VisibleSpecialChars", 0), // property VisibleSpecialChars
			/* 11 */ imports.NewTable("TLazSynTextArea_RightEdgeColumn", 0), // property RightEdgeColumn
			/* 12 */ imports.NewTable("TLazSynTextArea_RightEdgeVisible", 0), // property RightEdgeVisible
			/* 13 */ imports.NewTable("TLazSynTextArea_RightEdgeColor", 0), // property RightEdgeColor
			/* 14 */ imports.NewTable("TLazSynTextArea_TopLine", 0), // property TopLine
			/* 15 */ imports.NewTable("TLazSynTextArea_LeftChar", 0), // property LeftChar
			/* 16 */ imports.NewTable("TLazSynTextArea_TheLinesView", 0), // property TheLinesView
			/* 17 */ imports.NewTable("TLazSynTextArea_Highlighter", 0), // property Highlighter
			/* 18 */ imports.NewTable("TLazSynTextArea_MarkupManager", 0), // property MarkupManager
			/* 19 */ imports.NewTable("TLazSynTextArea_TextDrawer", 0), // property TextDrawer
			/* 20 */ imports.NewTable("TLazSynTextArea_TextBounds", 0), // property TextBounds
			/* 21 */ imports.NewTable("TLazSynTextArea_LineHeight", 0), // property LineHeight
			/* 22 */ imports.NewTable("TLazSynTextArea_CharWidth", 0), // property CharWidth
			/* 23 */ imports.NewTable("TLazSynTextArea_LinesInWindow", 0), // property LinesInWindow
			/* 24 */ imports.NewTable("TLazSynTextArea_CharsInWindow", 0), // property CharsInWindow
		}
	})
	return lazSynTextAreaImport
}
