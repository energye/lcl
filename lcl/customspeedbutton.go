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

// ICustomSpeedButton Parent: IGraphicControl
type ICustomSpeedButton interface {
	IGraphicControl
	FindDownButton() ICustomSpeedButton                               // function
	Click()                                                           // procedure
	LoadGlyphFromResourceName(instance types.TLCLHandle, name string) // procedure
	LoadGlyphFromLazarusResource(name string)                         // procedure
	Alignment() types.TAlignment                                      // property Alignment Getter
	SetAlignment(value types.TAlignment)                              // property Alignment Setter
	AllowAllUp() bool                                                 // property AllowAllUp Getter
	SetAllowAllUp(value bool)                                         // property AllowAllUp Setter
	DisabledImageIndex() int32                                        // property DisabledImageIndex Getter
	SetDisabledImageIndex(value int32)                                // property DisabledImageIndex Setter
	Down() bool                                                       // property Down Getter
	SetDown(value bool)                                               // property Down Setter
	Flat() bool                                                       // property Flat Getter
	SetFlat(value bool)                                               // property Flat Setter
	Glyph() IBitmap                                                   // property Glyph Getter
	SetGlyph(value IBitmap)                                           // property Glyph Setter
	GroupIndex() int32                                                // property GroupIndex Getter
	SetGroupIndex(value int32)                                        // property GroupIndex Setter
	HotImageIndex() int32                                             // property HotImageIndex Getter
	SetHotImageIndex(value int32)                                     // property HotImageIndex Setter
	Images() ICustomImageList                                         // property Images Getter
	SetImages(value ICustomImageList)                                 // property Images Setter
	ImageIndex() int32                                                // property ImageIndex Getter
	SetImageIndex(value int32)                                        // property ImageIndex Setter
	ImageWidth() int32                                                // property ImageWidth Getter
	SetImageWidth(value int32)                                        // property ImageWidth Setter
	Layout() types.TButtonLayout                                      // property Layout Getter
	SetLayout(value types.TButtonLayout)                              // property Layout Setter
	Margin() int32                                                    // property Margin Getter
	SetMargin(value int32)                                            // property Margin Setter
	NumGlyphs() int32                                                 // property NumGlyphs Getter
	SetNumGlyphs(value int32)                                         // property NumGlyphs Setter
	PressedImageIndex() int32                                         // property PressedImageIndex Getter
	SetPressedImageIndex(value int32)                                 // property PressedImageIndex Setter
	SelectedImageIndex() int32                                        // property SelectedImageIndex Getter
	SetSelectedImageIndex(value int32)                                // property SelectedImageIndex Setter
	ShowAccelChar() bool                                              // property ShowAccelChar Getter
	SetShowAccelChar(value bool)                                      // property ShowAccelChar Setter
	ShowCaption() bool                                                // property ShowCaption Getter
	SetShowCaption(value bool)                                        // property ShowCaption Setter
	Spacing() int32                                                   // property Spacing Getter
	SetSpacing(value int32)                                           // property Spacing Setter
	Transparent() bool                                                // property Transparent Getter
	SetTransparent(value bool)                                        // property Transparent Setter
}

type TCustomSpeedButton struct {
	TGraphicControl
}

func (m *TCustomSpeedButton) FindDownButton() ICustomSpeedButton {
	if !m.IsValid() {
		return nil
	}
	r := customSpeedButtonAPI().SysCallN(1, m.Instance())
	return AsCustomSpeedButton(r)
}

func (m *TCustomSpeedButton) Click() {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(2, m.Instance())
}

func (m *TCustomSpeedButton) LoadGlyphFromResourceName(instance types.TLCLHandle, name string) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(3, m.Instance(), uintptr(instance), api.PasStr(name))
}

func (m *TCustomSpeedButton) LoadGlyphFromLazarusResource(name string) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(4, m.Instance(), api.PasStr(name))
}

func (m *TCustomSpeedButton) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(5, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TCustomSpeedButton) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) AllowAllUp() bool {
	if !m.IsValid() {
		return false
	}
	r := customSpeedButtonAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSpeedButton) SetAllowAllUp(value bool) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSpeedButton) DisabledImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetDisabledImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) Down() bool {
	if !m.IsValid() {
		return false
	}
	r := customSpeedButtonAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSpeedButton) SetDown(value bool) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSpeedButton) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := customSpeedButtonAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSpeedButton) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSpeedButton) Glyph() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := customSpeedButtonAPI().SysCallN(10, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TCustomSpeedButton) SetGlyph(value IBitmap) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(10, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSpeedButton) GroupIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(11, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetGroupIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) HotImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetHotImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customSpeedButtonAPI().SysCallN(13, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomSpeedButton) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomSpeedButton) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) ImageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetImageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) Layout() types.TButtonLayout {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(16, 0, m.Instance())
	return types.TButtonLayout(r)
}

func (m *TCustomSpeedButton) SetLayout(value types.TButtonLayout) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) Margin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetMargin(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) NumGlyphs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetNumGlyphs(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) PressedImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(19, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetPressedImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) SelectedImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetSelectedImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) ShowAccelChar() bool {
	if !m.IsValid() {
		return false
	}
	r := customSpeedButtonAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSpeedButton) SetShowAccelChar(value bool) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSpeedButton) ShowCaption() bool {
	if !m.IsValid() {
		return false
	}
	r := customSpeedButtonAPI().SysCallN(22, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSpeedButton) SetShowCaption(value bool) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(22, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomSpeedButton) Spacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customSpeedButtonAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TCustomSpeedButton) SetSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TCustomSpeedButton) Transparent() bool {
	if !m.IsValid() {
		return false
	}
	r := customSpeedButtonAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomSpeedButton) SetTransparent(value bool) {
	if !m.IsValid() {
		return
	}
	customSpeedButtonAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

// NewCustomSpeedButton class constructor
func NewCustomSpeedButton(owner IComponent) ICustomSpeedButton {
	r := customSpeedButtonAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomSpeedButton(r)
}

func TCustomSpeedButtonClass() types.TClass {
	r := customSpeedButtonAPI().SysCallN(25)
	return types.TClass(r)
}

var (
	customSpeedButtonOnce   base.Once
	customSpeedButtonImport *imports.Imports = nil
)

func customSpeedButtonAPI() *imports.Imports {
	customSpeedButtonOnce.Do(func() {
		customSpeedButtonImport = api.NewDefaultImports()
		customSpeedButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomSpeedButton_Create", 0), // constructor NewCustomSpeedButton
			/* 1 */ imports.NewTable("TCustomSpeedButton_FindDownButton", 0), // function FindDownButton
			/* 2 */ imports.NewTable("TCustomSpeedButton_Click", 0), // procedure Click
			/* 3 */ imports.NewTable("TCustomSpeedButton_LoadGlyphFromResourceName", 0), // procedure LoadGlyphFromResourceName
			/* 4 */ imports.NewTable("TCustomSpeedButton_LoadGlyphFromLazarusResource", 0), // procedure LoadGlyphFromLazarusResource
			/* 5 */ imports.NewTable("TCustomSpeedButton_Alignment", 0), // property Alignment
			/* 6 */ imports.NewTable("TCustomSpeedButton_AllowAllUp", 0), // property AllowAllUp
			/* 7 */ imports.NewTable("TCustomSpeedButton_DisabledImageIndex", 0), // property DisabledImageIndex
			/* 8 */ imports.NewTable("TCustomSpeedButton_Down", 0), // property Down
			/* 9 */ imports.NewTable("TCustomSpeedButton_Flat", 0), // property Flat
			/* 10 */ imports.NewTable("TCustomSpeedButton_Glyph", 0), // property Glyph
			/* 11 */ imports.NewTable("TCustomSpeedButton_GroupIndex", 0), // property GroupIndex
			/* 12 */ imports.NewTable("TCustomSpeedButton_HotImageIndex", 0), // property HotImageIndex
			/* 13 */ imports.NewTable("TCustomSpeedButton_Images", 0), // property Images
			/* 14 */ imports.NewTable("TCustomSpeedButton_ImageIndex", 0), // property ImageIndex
			/* 15 */ imports.NewTable("TCustomSpeedButton_ImageWidth", 0), // property ImageWidth
			/* 16 */ imports.NewTable("TCustomSpeedButton_Layout", 0), // property Layout
			/* 17 */ imports.NewTable("TCustomSpeedButton_Margin", 0), // property Margin
			/* 18 */ imports.NewTable("TCustomSpeedButton_NumGlyphs", 0), // property NumGlyphs
			/* 19 */ imports.NewTable("TCustomSpeedButton_PressedImageIndex", 0), // property PressedImageIndex
			/* 20 */ imports.NewTable("TCustomSpeedButton_SelectedImageIndex", 0), // property SelectedImageIndex
			/* 21 */ imports.NewTable("TCustomSpeedButton_ShowAccelChar", 0), // property ShowAccelChar
			/* 22 */ imports.NewTable("TCustomSpeedButton_ShowCaption", 0), // property ShowCaption
			/* 23 */ imports.NewTable("TCustomSpeedButton_Spacing", 0), // property Spacing
			/* 24 */ imports.NewTable("TCustomSpeedButton_Transparent", 0), // property Transparent
			/* 25 */ imports.NewTable("TCustomSpeedButton_TClass", 0), // function TCustomSpeedButtonClass
		}
	})
	return customSpeedButtonImport
}
