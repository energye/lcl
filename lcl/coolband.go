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

// ICoolBand Parent: ICollectionItem
type ICoolBand interface {
	ICollectionItem
	Height() int32                      // property
	Left() int32                        // property
	Right() int32                       // property
	Top() int32                         // property
	Bitmap() IBitmap                    // property
	SetBitmap(AValue IBitmap)           // property
	BorderStyle() TBorderStyle          // property
	SetBorderStyle(AValue TBorderStyle) // property
	Break() bool                        // property
	SetBreak(AValue bool)               // property
	Color() TColor                      // property
	SetColor(AValue TColor)             // property
	Control() IControl                  // property
	SetControl(AValue IControl)         // property
	FixedBackground() bool              // property
	SetFixedBackground(AValue bool)     // property
	FixedSize() bool                    // property
	SetFixedSize(AValue bool)           // property
	HorizontalOnly() bool               // property
	SetHorizontalOnly(AValue bool)      // property
	ImageIndex() TImageIndex            // property
	SetImageIndex(AValue TImageIndex)   // property
	MinHeight() int32                   // property
	SetMinHeight(AValue int32)          // property
	MinWidth() int32                    // property
	SetMinWidth(AValue int32)           // property
	ParentColor() bool                  // property
	SetParentColor(AValue bool)         // property
	ParentBitmap() bool                 // property
	SetParentBitmap(AValue bool)        // property
	Text() string                       // property
	SetText(AValue string)              // property
	Visible() bool                      // property
	SetVisible(AValue bool)             // property
	Width() int32                       // property
	SetWidth(AValue int32)              // property
	AutosizeWidth()                     // procedure
	InvalidateCoolBar(Sender IObject)   // procedure
}

// TCoolBand Parent: TCollectionItem
type TCoolBand struct {
	TCollectionItem
}

func NewCoolBand(aCollection ICollection) ICoolBand {
	r1 := coolBandImportAPI().SysCallN(7, GetObjectUintptr(aCollection))
	return AsCoolBand(r1)
}

func (m *TCoolBand) Height() int32 {
	r1 := coolBandImportAPI().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Left() int32 {
	r1 := coolBandImportAPI().SysCallN(14, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Right() int32 {
	r1 := coolBandImportAPI().SysCallN(19, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Top() int32 {
	r1 := coolBandImportAPI().SysCallN(21, m.Instance())
	return int32(r1)
}

func (m *TCoolBand) Bitmap() IBitmap {
	r1 := coolBandImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCoolBand) SetBitmap(AValue IBitmap) {
	coolBandImportAPI().SysCallN(1, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCoolBand) BorderStyle() TBorderStyle {
	r1 := coolBandImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCoolBand) SetBorderStyle(AValue TBorderStyle) {
	coolBandImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) Break() bool {
	r1 := coolBandImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetBreak(AValue bool) {
	coolBandImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) Color() TColor {
	r1 := coolBandImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCoolBand) SetColor(AValue TColor) {
	coolBandImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) Control() IControl {
	r1 := coolBandImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCoolBand) SetControl(AValue IControl) {
	coolBandImportAPI().SysCallN(6, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCoolBand) FixedBackground() bool {
	r1 := coolBandImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetFixedBackground(AValue bool) {
	coolBandImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) FixedSize() bool {
	r1 := coolBandImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetFixedSize(AValue bool) {
	coolBandImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) HorizontalOnly() bool {
	r1 := coolBandImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetHorizontalOnly(AValue bool) {
	coolBandImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) ImageIndex() TImageIndex {
	r1 := coolBandImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCoolBand) SetImageIndex(AValue TImageIndex) {
	coolBandImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) MinHeight() int32 {
	r1 := coolBandImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoolBand) SetMinHeight(AValue int32) {
	coolBandImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) MinWidth() int32 {
	r1 := coolBandImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoolBand) SetMinWidth(AValue int32) {
	coolBandImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBand) ParentColor() bool {
	r1 := coolBandImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetParentColor(AValue bool) {
	coolBandImportAPI().SysCallN(18, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) ParentBitmap() bool {
	r1 := coolBandImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetParentBitmap(AValue bool) {
	coolBandImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) Text() string {
	r1 := coolBandImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoolBand) SetText(AValue string) {
	coolBandImportAPI().SysCallN(20, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoolBand) Visible() bool {
	r1 := coolBandImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBand) SetVisible(AValue bool) {
	coolBandImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBand) Width() int32 {
	r1 := coolBandImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoolBand) SetWidth(AValue int32) {
	coolBandImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func CoolBandClass() TClass {
	ret := coolBandImportAPI().SysCallN(4)
	return TClass(ret)
}

func (m *TCoolBand) AutosizeWidth() {
	coolBandImportAPI().SysCallN(0, m.Instance())
}

func (m *TCoolBand) InvalidateCoolBar(Sender IObject) {
	coolBandImportAPI().SysCallN(13, m.Instance(), GetObjectUintptr(Sender))
}

var (
	coolBandImport       *imports.Imports = nil
	coolBandImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoolBand_AutosizeWidth", 0),
		/*1*/ imports.NewTable("CoolBand_Bitmap", 0),
		/*2*/ imports.NewTable("CoolBand_BorderStyle", 0),
		/*3*/ imports.NewTable("CoolBand_Break", 0),
		/*4*/ imports.NewTable("CoolBand_Class", 0),
		/*5*/ imports.NewTable("CoolBand_Color", 0),
		/*6*/ imports.NewTable("CoolBand_Control", 0),
		/*7*/ imports.NewTable("CoolBand_Create", 0),
		/*8*/ imports.NewTable("CoolBand_FixedBackground", 0),
		/*9*/ imports.NewTable("CoolBand_FixedSize", 0),
		/*10*/ imports.NewTable("CoolBand_Height", 0),
		/*11*/ imports.NewTable("CoolBand_HorizontalOnly", 0),
		/*12*/ imports.NewTable("CoolBand_ImageIndex", 0),
		/*13*/ imports.NewTable("CoolBand_InvalidateCoolBar", 0),
		/*14*/ imports.NewTable("CoolBand_Left", 0),
		/*15*/ imports.NewTable("CoolBand_MinHeight", 0),
		/*16*/ imports.NewTable("CoolBand_MinWidth", 0),
		/*17*/ imports.NewTable("CoolBand_ParentBitmap", 0),
		/*18*/ imports.NewTable("CoolBand_ParentColor", 0),
		/*19*/ imports.NewTable("CoolBand_Right", 0),
		/*20*/ imports.NewTable("CoolBand_Text", 0),
		/*21*/ imports.NewTable("CoolBand_Top", 0),
		/*22*/ imports.NewTable("CoolBand_Visible", 0),
		/*23*/ imports.NewTable("CoolBand_Width", 0),
	}
)

func coolBandImportAPI() *imports.Imports {
	if coolBandImport == nil {
		coolBandImport = NewDefaultImports()
		coolBandImport.SetImportTable(coolBandImportTables)
		coolBandImportTables = nil
	}
	return coolBandImport
}
