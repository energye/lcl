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

// ICustomSpeedButton Parent: IGraphicControl
type ICustomSpeedButton interface {
	IGraphicControl
	Alignment() TAlignment                                    // property
	SetAlignment(AValue TAlignment)                           // property
	AllowAllUp() bool                                         // property
	SetAllowAllUp(AValue bool)                                // property
	DisabledImageIndex() TImageIndex                          // property
	SetDisabledImageIndex(AValue TImageIndex)                 // property
	Down() bool                                               // property
	SetDown(AValue bool)                                      // property
	Flat() bool                                               // property
	SetFlat(AValue bool)                                      // property
	Glyph() IBitmap                                           // property
	SetGlyph(AValue IBitmap)                                  // property
	GroupIndex() int32                                        // property
	SetGroupIndex(AValue int32)                               // property
	HotImageIndex() TImageIndex                               // property
	SetHotImageIndex(AValue TImageIndex)                      // property
	Images() ICustomImageList                                 // property
	SetImages(AValue ICustomImageList)                        // property
	ImageIndex() TImageIndex                                  // property
	SetImageIndex(AValue TImageIndex)                         // property
	ImageWidth() int32                                        // property
	SetImageWidth(AValue int32)                               // property
	Layout() TButtonLayout                                    // property
	SetLayout(AValue TButtonLayout)                           // property
	Margin() int32                                            // property
	SetMargin(AValue int32)                                   // property
	NumGlyphs() int32                                         // property
	SetNumGlyphs(AValue int32)                                // property
	PressedImageIndex() TImageIndex                           // property
	SetPressedImageIndex(AValue TImageIndex)                  // property
	SelectedImageIndex() TImageIndex                          // property
	SetSelectedImageIndex(AValue TImageIndex)                 // property
	ShowAccelChar() bool                                      // property
	SetShowAccelChar(AValue bool)                             // property
	ShowCaption() bool                                        // property
	SetShowCaption(AValue bool)                               // property
	Spacing() int32                                           // property
	SetSpacing(AValue int32)                                  // property
	Transparent() bool                                        // property
	SetTransparent(AValue bool)                               // property
	FindDownButton() ICustomSpeedButton                       // function
	Click()                                                   // procedure
	LoadGlyphFromResourceName(Instance THandle, AName string) // procedure
	LoadGlyphFromLazarusResource(AName string)                // procedure
}

// TCustomSpeedButton Parent: TGraphicControl
type TCustomSpeedButton struct {
	TGraphicControl
}

func NewCustomSpeedButton(AOwner IComponent) ICustomSpeedButton {
	r1 := customSpeedButtonImportAPI().SysCallN(4, GetObjectUintptr(AOwner))
	return AsCustomSpeedButton(r1)
}

func (m *TCustomSpeedButton) Alignment() TAlignment {
	r1 := customSpeedButtonImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TCustomSpeedButton) SetAlignment(AValue TAlignment) {
	customSpeedButtonImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) AllowAllUp() bool {
	r1 := customSpeedButtonImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetAllowAllUp(AValue bool) {
	customSpeedButtonImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) DisabledImageIndex() TImageIndex {
	r1 := customSpeedButtonImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetDisabledImageIndex(AValue TImageIndex) {
	customSpeedButtonImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Down() bool {
	r1 := customSpeedButtonImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetDown(AValue bool) {
	customSpeedButtonImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) Flat() bool {
	r1 := customSpeedButtonImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetFlat(AValue bool) {
	customSpeedButtonImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) Glyph() IBitmap {
	r1 := customSpeedButtonImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCustomSpeedButton) SetGlyph(AValue IBitmap) {
	customSpeedButtonImportAPI().SysCallN(9, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomSpeedButton) GroupIndex() int32 {
	r1 := customSpeedButtonImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetGroupIndex(AValue int32) {
	customSpeedButtonImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) HotImageIndex() TImageIndex {
	r1 := customSpeedButtonImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetHotImageIndex(AValue TImageIndex) {
	customSpeedButtonImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Images() ICustomImageList {
	r1 := customSpeedButtonImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomSpeedButton) SetImages(AValue ICustomImageList) {
	customSpeedButtonImportAPI().SysCallN(14, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomSpeedButton) ImageIndex() TImageIndex {
	r1 := customSpeedButtonImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetImageIndex(AValue TImageIndex) {
	customSpeedButtonImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) ImageWidth() int32 {
	r1 := customSpeedButtonImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetImageWidth(AValue int32) {
	customSpeedButtonImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Layout() TButtonLayout {
	r1 := customSpeedButtonImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TButtonLayout(r1)
}

func (m *TCustomSpeedButton) SetLayout(AValue TButtonLayout) {
	customSpeedButtonImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Margin() int32 {
	r1 := customSpeedButtonImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetMargin(AValue int32) {
	customSpeedButtonImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) NumGlyphs() int32 {
	r1 := customSpeedButtonImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetNumGlyphs(AValue int32) {
	customSpeedButtonImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) PressedImageIndex() TImageIndex {
	r1 := customSpeedButtonImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetPressedImageIndex(AValue TImageIndex) {
	customSpeedButtonImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) SelectedImageIndex() TImageIndex {
	r1 := customSpeedButtonImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomSpeedButton) SetSelectedImageIndex(AValue TImageIndex) {
	customSpeedButtonImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) ShowAccelChar() bool {
	r1 := customSpeedButtonImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetShowAccelChar(AValue bool) {
	customSpeedButtonImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) ShowCaption() bool {
	r1 := customSpeedButtonImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetShowCaption(AValue bool) {
	customSpeedButtonImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) Spacing() int32 {
	r1 := customSpeedButtonImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSpeedButton) SetSpacing(AValue int32) {
	customSpeedButtonImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSpeedButton) Transparent() bool {
	r1 := customSpeedButtonImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSpeedButton) SetTransparent(AValue bool) {
	customSpeedButtonImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSpeedButton) FindDownButton() ICustomSpeedButton {
	r1 := customSpeedButtonImportAPI().SysCallN(7, m.Instance())
	return AsCustomSpeedButton(r1)
}

func CustomSpeedButtonClass() TClass {
	ret := customSpeedButtonImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TCustomSpeedButton) Click() {
	customSpeedButtonImportAPI().SysCallN(3, m.Instance())
}

func (m *TCustomSpeedButton) LoadGlyphFromResourceName(Instance THandle, AName string) {
	customSpeedButtonImportAPI().SysCallN(17, m.Instance(), uintptr(Instance), PascalStr(AName))
}

func (m *TCustomSpeedButton) LoadGlyphFromLazarusResource(AName string) {
	customSpeedButtonImportAPI().SysCallN(16, m.Instance(), PascalStr(AName))
}

var (
	customSpeedButtonImport       *imports.Imports = nil
	customSpeedButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomSpeedButton_Alignment", 0),
		/*1*/ imports.NewTable("CustomSpeedButton_AllowAllUp", 0),
		/*2*/ imports.NewTable("CustomSpeedButton_Class", 0),
		/*3*/ imports.NewTable("CustomSpeedButton_Click", 0),
		/*4*/ imports.NewTable("CustomSpeedButton_Create", 0),
		/*5*/ imports.NewTable("CustomSpeedButton_DisabledImageIndex", 0),
		/*6*/ imports.NewTable("CustomSpeedButton_Down", 0),
		/*7*/ imports.NewTable("CustomSpeedButton_FindDownButton", 0),
		/*8*/ imports.NewTable("CustomSpeedButton_Flat", 0),
		/*9*/ imports.NewTable("CustomSpeedButton_Glyph", 0),
		/*10*/ imports.NewTable("CustomSpeedButton_GroupIndex", 0),
		/*11*/ imports.NewTable("CustomSpeedButton_HotImageIndex", 0),
		/*12*/ imports.NewTable("CustomSpeedButton_ImageIndex", 0),
		/*13*/ imports.NewTable("CustomSpeedButton_ImageWidth", 0),
		/*14*/ imports.NewTable("CustomSpeedButton_Images", 0),
		/*15*/ imports.NewTable("CustomSpeedButton_Layout", 0),
		/*16*/ imports.NewTable("CustomSpeedButton_LoadGlyphFromLazarusResource", 0),
		/*17*/ imports.NewTable("CustomSpeedButton_LoadGlyphFromResourceName", 0),
		/*18*/ imports.NewTable("CustomSpeedButton_Margin", 0),
		/*19*/ imports.NewTable("CustomSpeedButton_NumGlyphs", 0),
		/*20*/ imports.NewTable("CustomSpeedButton_PressedImageIndex", 0),
		/*21*/ imports.NewTable("CustomSpeedButton_SelectedImageIndex", 0),
		/*22*/ imports.NewTable("CustomSpeedButton_ShowAccelChar", 0),
		/*23*/ imports.NewTable("CustomSpeedButton_ShowCaption", 0),
		/*24*/ imports.NewTable("CustomSpeedButton_Spacing", 0),
		/*25*/ imports.NewTable("CustomSpeedButton_Transparent", 0),
	}
)

func customSpeedButtonImportAPI() *imports.Imports {
	if customSpeedButtonImport == nil {
		customSpeedButtonImport = NewDefaultImports()
		customSpeedButtonImport.SetImportTable(customSpeedButtonImportTables)
		customSpeedButtonImportTables = nil
	}
	return customSpeedButtonImport
}
