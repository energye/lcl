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

// ISynSelectedColor Parent: ISynHighlighterAttributesModifier
type ISynSelectedColor interface {
	ISynHighlighterAttributesModifier
	GetModifiedStyle(style types.TFontStyles) types.TFontStyles // function
	// SetFrameBoundsPhys
	//  boundaries of the frame
	SetFrameBoundsPhys(start int32, end int32)                                                                                                            // procedure
	SetFrameBoundsLog(start int32, end int32, startOffs int32, endOffs int32)                                                                             // procedure
	ModifyColors(foreground *types.TColor, background *types.TColor, frameColor *types.TColor, style *types.TFontStyles, frameStyle *types.TSynLineStyle) // procedure
	StartX() TLazSynDisplayTokenBound                                                                                                                     // property StartX Getter
	SetStartX(value TLazSynDisplayTokenBound)                                                                                                             // property StartX Setter
	EndX() TLazSynDisplayTokenBound                                                                                                                       // property EndX Getter
	SetEndX(value TLazSynDisplayTokenBound)                                                                                                               // property EndX Setter
}

type TSynSelectedColor struct {
	TSynHighlighterAttributesModifier
}

func (m *TSynSelectedColor) GetModifiedStyle(style types.TFontStyles) types.TFontStyles {
	if !m.IsValid() {
		return 0
	}
	r := synSelectedColorAPI().SysCallN(2, m.Instance(), uintptr(style))
	return types.TFontStyles(r)
}

func (m *TSynSelectedColor) SetFrameBoundsPhys(start int32, end int32) {
	if !m.IsValid() {
		return
	}
	synSelectedColorAPI().SysCallN(3, m.Instance(), uintptr(start), uintptr(end))
}

func (m *TSynSelectedColor) SetFrameBoundsLog(start int32, end int32, startOffs int32, endOffs int32) {
	if !m.IsValid() {
		return
	}
	synSelectedColorAPI().SysCallN(4, m.Instance(), uintptr(start), uintptr(end), uintptr(startOffs), uintptr(endOffs))
}

func (m *TSynSelectedColor) ModifyColors(foreground *types.TColor, background *types.TColor, frameColor *types.TColor, style *types.TFontStyles, frameStyle *types.TSynLineStyle) {
	if !m.IsValid() {
		return
	}
	foregroundPtr := uintptr(*foreground)
	backgroundPtr := uintptr(*background)
	frameColorPtr := uintptr(*frameColor)
	stylePtr := uintptr(*style)
	frameStylePtr := uintptr(*frameStyle)
	synSelectedColorAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&foregroundPtr)), uintptr(base.UnsafePointer(&backgroundPtr)), uintptr(base.UnsafePointer(&frameColorPtr)), uintptr(base.UnsafePointer(&stylePtr)), uintptr(base.UnsafePointer(&frameStylePtr)))
	*foreground = types.TColor(foregroundPtr)
	*background = types.TColor(backgroundPtr)
	*frameColor = types.TColor(frameColorPtr)
	*style = types.TFontStyles(stylePtr)
	*frameStyle = types.TSynLineStyle(frameStylePtr)
}

func (m *TSynSelectedColor) StartX() (result TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorAPI().SysCallN(6, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynSelectedColor) SetStartX(value TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorAPI().SysCallN(6, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TSynSelectedColor) EndX() (result TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorAPI().SysCallN(7, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TSynSelectedColor) SetEndX(value TLazSynDisplayTokenBound) {
	if !m.IsValid() {
		return
	}
	synSelectedColorAPI().SysCallN(7, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

// NewSynSelectedColor class constructor
func NewSynSelectedColor() ISynSelectedColor {
	r := synSelectedColorAPI().SysCallN(0)
	return AsSynSelectedColor(r)
}

// NewSynSelectedColorWithStrX2 class constructor
func NewSynSelectedColorWithStrX2(caption string, storedName string) ISynSelectedColor {
	r := synSelectedColorAPI().SysCallN(1, api.PasStr(caption), api.PasStr(storedName))
	return AsSynSelectedColor(r)
}

var (
	synSelectedColorOnce   base.Once
	synSelectedColorImport *imports.Imports = nil
)

func synSelectedColorAPI() *imports.Imports {
	synSelectedColorOnce.Do(func() {
		synSelectedColorImport = api.NewDefaultImports()
		synSelectedColorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TSynSelectedColor_Create", 0), // constructor NewSynSelectedColor
			/* 1 */ imports.NewTable("TSynSelectedColor_CreateWithStrX2", 0), // constructor NewSynSelectedColorWithStrX2
			/* 2 */ imports.NewTable("TSynSelectedColor_GetModifiedStyle", 0), // function GetModifiedStyle
			/* 3 */ imports.NewTable("TSynSelectedColor_SetFrameBoundsPhys", 0), // procedure SetFrameBoundsPhys
			/* 4 */ imports.NewTable("TSynSelectedColor_SetFrameBoundsLog", 0), // procedure SetFrameBoundsLog
			/* 5 */ imports.NewTable("TSynSelectedColor_ModifyColors", 0), // procedure ModifyColors
			/* 6 */ imports.NewTable("TSynSelectedColor_StartX", 0), // property StartX
			/* 7 */ imports.NewTable("TSynSelectedColor_EndX", 0), // property EndX
		}
	})
	return synSelectedColorImport
}
