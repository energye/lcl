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

// ICustomBitBtn Parent: ICustomButton
type ICustomBitBtn interface {
	ICustomButton
	DefaultCaption() bool                                     // property
	SetDefaultCaption(AValue bool)                            // property
	DisabledImageIndex() TImageIndex                          // property
	SetDisabledImageIndex(AValue TImageIndex)                 // property
	Glyph() IBitmap                                           // property
	SetGlyph(AValue IBitmap)                                  // property
	NumGlyphs() int32                                         // property
	SetNumGlyphs(AValue int32)                                // property
	HotImageIndex() TImageIndex                               // property
	SetHotImageIndex(AValue TImageIndex)                      // property
	Images() ICustomImageList                                 // property
	SetImages(AValue ICustomImageList)                        // property
	ImageIndex() TImageIndex                                  // property
	SetImageIndex(AValue TImageIndex)                         // property
	ImageWidth() int32                                        // property
	SetImageWidth(AValue int32)                               // property
	Kind() TBitBtnKind                                        // property
	SetKind(AValue TBitBtnKind)                               // property
	Layout() TButtonLayout                                    // property
	SetLayout(AValue TButtonLayout)                           // property
	Margin() int32                                            // property
	SetMargin(AValue int32)                                   // property
	PressedImageIndex() TImageIndex                           // property
	SetPressedImageIndex(AValue TImageIndex)                  // property
	Spacing() int32                                           // property
	SetSpacing(AValue int32)                                  // property
	GlyphShowMode() TGlyphShowMode                            // property
	SetGlyphShowMode(AValue TGlyphShowMode)                   // property
	CanShowGlyph(AWithShowMode bool) bool                     // function
	LoadGlyphFromResourceName(Instance THandle, AName string) // procedure
	LoadGlyphFromLazarusResource(AName string)                // procedure
	LoadGlyphFromStock(idButton int32)                        // procedure
	LoadGlyphFromResource(idButton TButtonImage)              // procedure
}

// TCustomBitBtn Parent: TCustomButton
type TCustomBitBtn struct {
	TCustomButton
}

func NewCustomBitBtn(TheOwner IComponent) ICustomBitBtn {
	r1 := customBitBtnImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsCustomBitBtn(r1)
}

func (m *TCustomBitBtn) DefaultCaption() bool {
	r1 := customBitBtnImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomBitBtn) SetDefaultCaption(AValue bool) {
	customBitBtnImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomBitBtn) DisabledImageIndex() TImageIndex {
	r1 := customBitBtnImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetDisabledImageIndex(AValue TImageIndex) {
	customBitBtnImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Glyph() IBitmap {
	r1 := customBitBtnImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCustomBitBtn) SetGlyph(AValue IBitmap) {
	customBitBtnImportAPI().SysCallN(5, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomBitBtn) NumGlyphs() int32 {
	r1 := customBitBtnImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetNumGlyphs(AValue int32) {
	customBitBtnImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) HotImageIndex() TImageIndex {
	r1 := customBitBtnImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetHotImageIndex(AValue TImageIndex) {
	customBitBtnImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Images() ICustomImageList {
	r1 := customBitBtnImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomBitBtn) SetImages(AValue ICustomImageList) {
	customBitBtnImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomBitBtn) ImageIndex() TImageIndex {
	r1 := customBitBtnImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetImageIndex(AValue TImageIndex) {
	customBitBtnImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) ImageWidth() int32 {
	r1 := customBitBtnImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetImageWidth(AValue int32) {
	customBitBtnImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Kind() TBitBtnKind {
	r1 := customBitBtnImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TBitBtnKind(r1)
}

func (m *TCustomBitBtn) SetKind(AValue TBitBtnKind) {
	customBitBtnImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Layout() TButtonLayout {
	r1 := customBitBtnImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TButtonLayout(r1)
}

func (m *TCustomBitBtn) SetLayout(AValue TButtonLayout) {
	customBitBtnImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Margin() int32 {
	r1 := customBitBtnImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetMargin(AValue int32) {
	customBitBtnImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) PressedImageIndex() TImageIndex {
	r1 := customBitBtnImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomBitBtn) SetPressedImageIndex(AValue TImageIndex) {
	customBitBtnImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) Spacing() int32 {
	r1 := customBitBtnImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomBitBtn) SetSpacing(AValue int32) {
	customBitBtnImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) GlyphShowMode() TGlyphShowMode {
	r1 := customBitBtnImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TGlyphShowMode(r1)
}

func (m *TCustomBitBtn) SetGlyphShowMode(AValue TGlyphShowMode) {
	customBitBtnImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomBitBtn) CanShowGlyph(AWithShowMode bool) bool {
	r1 := customBitBtnImportAPI().SysCallN(0, m.Instance(), PascalBool(AWithShowMode))
	return GoBool(r1)
}

func CustomBitBtnClass() TClass {
	ret := customBitBtnImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCustomBitBtn) LoadGlyphFromResourceName(Instance THandle, AName string) {
	customBitBtnImportAPI().SysCallN(15, m.Instance(), uintptr(Instance), PascalStr(AName))
}

func (m *TCustomBitBtn) LoadGlyphFromLazarusResource(AName string) {
	customBitBtnImportAPI().SysCallN(13, m.Instance(), PascalStr(AName))
}

func (m *TCustomBitBtn) LoadGlyphFromStock(idButton int32) {
	customBitBtnImportAPI().SysCallN(16, m.Instance(), uintptr(idButton))
}

func (m *TCustomBitBtn) LoadGlyphFromResource(idButton TButtonImage) {
	customBitBtnImportAPI().SysCallN(14, m.Instance(), uintptr(idButton))
}

var (
	customBitBtnImport       *imports.Imports = nil
	customBitBtnImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomBitBtn_CanShowGlyph", 0),
		/*1*/ imports.NewTable("CustomBitBtn_Class", 0),
		/*2*/ imports.NewTable("CustomBitBtn_Create", 0),
		/*3*/ imports.NewTable("CustomBitBtn_DefaultCaption", 0),
		/*4*/ imports.NewTable("CustomBitBtn_DisabledImageIndex", 0),
		/*5*/ imports.NewTable("CustomBitBtn_Glyph", 0),
		/*6*/ imports.NewTable("CustomBitBtn_GlyphShowMode", 0),
		/*7*/ imports.NewTable("CustomBitBtn_HotImageIndex", 0),
		/*8*/ imports.NewTable("CustomBitBtn_ImageIndex", 0),
		/*9*/ imports.NewTable("CustomBitBtn_ImageWidth", 0),
		/*10*/ imports.NewTable("CustomBitBtn_Images", 0),
		/*11*/ imports.NewTable("CustomBitBtn_Kind", 0),
		/*12*/ imports.NewTable("CustomBitBtn_Layout", 0),
		/*13*/ imports.NewTable("CustomBitBtn_LoadGlyphFromLazarusResource", 0),
		/*14*/ imports.NewTable("CustomBitBtn_LoadGlyphFromResource", 0),
		/*15*/ imports.NewTable("CustomBitBtn_LoadGlyphFromResourceName", 0),
		/*16*/ imports.NewTable("CustomBitBtn_LoadGlyphFromStock", 0),
		/*17*/ imports.NewTable("CustomBitBtn_Margin", 0),
		/*18*/ imports.NewTable("CustomBitBtn_NumGlyphs", 0),
		/*19*/ imports.NewTable("CustomBitBtn_PressedImageIndex", 0),
		/*20*/ imports.NewTable("CustomBitBtn_Spacing", 0),
	}
)

func customBitBtnImportAPI() *imports.Imports {
	if customBitBtnImport == nil {
		customBitBtnImport = NewDefaultImports()
		customBitBtnImport.SetImportTable(customBitBtnImportTables)
		customBitBtnImportTables = nil
	}
	return customBitBtnImport
}
