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

// ICustomBitBtn Parent: ICustomButton
type ICustomBitBtn interface {
	ICustomButton
	CanShowGlyph(withShowMode bool) bool                              // function
	LoadGlyphFromResourceName(instance types.TLCLHandle, name string) // procedure
	LoadGlyphFromLazarusResource(name string)                         // procedure
	LoadGlyphFromStock(idButton int32)                                // procedure
	LoadGlyphFromResource(idButton types.TButtonImage)                // procedure
	DefaultCaption() bool                                             // property DefaultCaption Getter
	SetDefaultCaption(value bool)                                     // property DefaultCaption Setter
	DisabledImageIndex() int32                                        // property DisabledImageIndex Getter
	SetDisabledImageIndex(value int32)                                // property DisabledImageIndex Setter
	Glyph() IBitmap                                                   // property Glyph Getter
	SetGlyph(value IBitmap)                                           // property Glyph Setter
	NumGlyphs() int32                                                 // property NumGlyphs Getter
	SetNumGlyphs(value int32)                                         // property NumGlyphs Setter
	HotImageIndex() int32                                             // property HotImageIndex Getter
	SetHotImageIndex(value int32)                                     // property HotImageIndex Setter
	Images() ICustomImageList                                         // property Images Getter
	SetImages(value ICustomImageList)                                 // property Images Setter
	ImageIndex() int32                                                // property ImageIndex Getter
	SetImageIndex(value int32)                                        // property ImageIndex Setter
	ImageWidth() int32                                                // property ImageWidth Getter
	SetImageWidth(value int32)                                        // property ImageWidth Setter
	Kind() types.TBitBtnKind                                          // property Kind Getter
	SetKind(value types.TBitBtnKind)                                  // property Kind Setter
	Layout() types.TButtonLayout                                      // property Layout Getter
	SetLayout(value types.TButtonLayout)                              // property Layout Setter
	Margin() int32                                                    // property Margin Getter
	SetMargin(value int32)                                            // property Margin Setter
	PressedImageIndex() int32                                         // property PressedImageIndex Getter
	SetPressedImageIndex(value int32)                                 // property PressedImageIndex Setter
	// Spacing
	//  property SelectedImageIndex: TImageIndex index bsExclusive read GetImageIndex write SetImageIndex default -1;
	Spacing() int32                              // property Spacing Getter
	SetSpacing(value int32)                      // property Spacing Setter
	GlyphShowMode() types.TGlyphShowMode         // property GlyphShowMode Getter
	SetGlyphShowMode(value types.TGlyphShowMode) // property GlyphShowMode Setter
}

type TCustomBitBtn struct {
	TCustomButton
}

func (m *TCustomBitBtn) CanShowGlyph(withShowMode bool) bool {
	if !m.IsValid() {
		return false
	}
	r := customBitBtnAPI().SysCallN(1, m.Instance(), api.PasBool(withShowMode))
	return api.GoBool(r)
}

func (m *TCustomBitBtn) LoadGlyphFromResourceName(instance types.TLCLHandle, name string) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(2, m.Instance(), uintptr(instance), api.PasStr(name))
}

func (m *TCustomBitBtn) LoadGlyphFromLazarusResource(name string) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(3, m.Instance(), api.PasStr(name))
}

func (m *TCustomBitBtn) LoadGlyphFromStock(idButton int32) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(4, m.Instance(), uintptr(idButton))
}

func (m *TCustomBitBtn) LoadGlyphFromResource(idButton types.TButtonImage) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(5, m.Instance(), uintptr(idButton))
}

func (m *TCustomBitBtn) DefaultCaption() bool {
	if !m.IsValid() {
		return false
	}
	r := customBitBtnAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomBitBtn) SetDefaultCaption(value bool) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomBitBtn) DisabledImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TCustomBitBtn) SetDisabledImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) Glyph() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := customBitBtnAPI().SysCallN(8, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TCustomBitBtn) SetGlyph(value IBitmap) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomBitBtn) NumGlyphs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TCustomBitBtn) SetNumGlyphs(value int32) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) HotImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TCustomBitBtn) SetHotImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customBitBtnAPI().SysCallN(11, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomBitBtn) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomBitBtn) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TCustomBitBtn) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) ImageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TCustomBitBtn) SetImageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) Kind() types.TBitBtnKind {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(14, 0, m.Instance())
	return types.TBitBtnKind(r)
}

func (m *TCustomBitBtn) SetKind(value types.TBitBtnKind) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) Layout() types.TButtonLayout {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(15, 0, m.Instance())
	return types.TButtonLayout(r)
}

func (m *TCustomBitBtn) SetLayout(value types.TButtonLayout) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) Margin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TCustomBitBtn) SetMargin(value int32) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) PressedImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TCustomBitBtn) SetPressedImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) Spacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TCustomBitBtn) SetSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TCustomBitBtn) GlyphShowMode() types.TGlyphShowMode {
	if !m.IsValid() {
		return 0
	}
	r := customBitBtnAPI().SysCallN(19, 0, m.Instance())
	return types.TGlyphShowMode(r)
}

func (m *TCustomBitBtn) SetGlyphShowMode(value types.TGlyphShowMode) {
	if !m.IsValid() {
		return
	}
	customBitBtnAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

// NewCustomBitBtn class constructor
func NewCustomBitBtn(theOwner IComponent) ICustomBitBtn {
	r := customBitBtnAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomBitBtn(r)
}

func TCustomBitBtnClass() types.TClass {
	r := customBitBtnAPI().SysCallN(20)
	return types.TClass(r)
}

var (
	customBitBtnOnce   base.Once
	customBitBtnImport *imports.Imports = nil
)

func customBitBtnAPI() *imports.Imports {
	customBitBtnOnce.Do(func() {
		customBitBtnImport = api.NewDefaultImports()
		customBitBtnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomBitBtn_Create", 0), // constructor NewCustomBitBtn
			/* 1 */ imports.NewTable("TCustomBitBtn_CanShowGlyph", 0), // function CanShowGlyph
			/* 2 */ imports.NewTable("TCustomBitBtn_LoadGlyphFromResourceName", 0), // procedure LoadGlyphFromResourceName
			/* 3 */ imports.NewTable("TCustomBitBtn_LoadGlyphFromLazarusResource", 0), // procedure LoadGlyphFromLazarusResource
			/* 4 */ imports.NewTable("TCustomBitBtn_LoadGlyphFromStock", 0), // procedure LoadGlyphFromStock
			/* 5 */ imports.NewTable("TCustomBitBtn_LoadGlyphFromResource", 0), // procedure LoadGlyphFromResource
			/* 6 */ imports.NewTable("TCustomBitBtn_DefaultCaption", 0), // property DefaultCaption
			/* 7 */ imports.NewTable("TCustomBitBtn_DisabledImageIndex", 0), // property DisabledImageIndex
			/* 8 */ imports.NewTable("TCustomBitBtn_Glyph", 0), // property Glyph
			/* 9 */ imports.NewTable("TCustomBitBtn_NumGlyphs", 0), // property NumGlyphs
			/* 10 */ imports.NewTable("TCustomBitBtn_HotImageIndex", 0), // property HotImageIndex
			/* 11 */ imports.NewTable("TCustomBitBtn_Images", 0), // property Images
			/* 12 */ imports.NewTable("TCustomBitBtn_ImageIndex", 0), // property ImageIndex
			/* 13 */ imports.NewTable("TCustomBitBtn_ImageWidth", 0), // property ImageWidth
			/* 14 */ imports.NewTable("TCustomBitBtn_Kind", 0), // property Kind
			/* 15 */ imports.NewTable("TCustomBitBtn_Layout", 0), // property Layout
			/* 16 */ imports.NewTable("TCustomBitBtn_Margin", 0), // property Margin
			/* 17 */ imports.NewTable("TCustomBitBtn_PressedImageIndex", 0), // property PressedImageIndex
			/* 18 */ imports.NewTable("TCustomBitBtn_Spacing", 0), // property Spacing
			/* 19 */ imports.NewTable("TCustomBitBtn_GlyphShowMode", 0), // property GlyphShowMode
			/* 20 */ imports.NewTable("TCustomBitBtn_TClass", 0), // function TCustomBitBtnClass
		}
	})
	return customBitBtnImport
}
