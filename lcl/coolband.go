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

// ICoolBand Parent: ICollectionItem
type ICoolBand interface {
	ICollectionItem
	AutosizeWidth()                          // procedure
	InvalidateCoolBar(sender IObject)        // procedure
	Height() int32                           // property Height Getter
	Left() int32                             // property Left Getter
	Right() int32                            // property Right Getter
	Top() int32                              // property Top Getter
	Bitmap() IBitmap                         // property Bitmap Getter
	SetBitmap(value IBitmap)                 // property Bitmap Setter
	BorderStyle() types.TBorderStyle         // property BorderStyle Getter
	SetBorderStyle(value types.TBorderStyle) // property BorderStyle Setter
	Break() bool                             // property Break Getter
	SetBreak(value bool)                     // property Break Setter
	Color() types.TColor                     // property Color Getter
	SetColor(value types.TColor)             // property Color Setter
	Control() IControl                       // property Control Getter
	SetControl(value IControl)               // property Control Setter
	FixedBackground() bool                   // property FixedBackground Getter
	SetFixedBackground(value bool)           // property FixedBackground Setter
	FixedSize() bool                         // property FixedSize Getter
	SetFixedSize(value bool)                 // property FixedSize Setter
	HorizontalOnly() bool                    // property HorizontalOnly Getter
	SetHorizontalOnly(value bool)            // property HorizontalOnly Setter
	ImageIndex() int32                       // property ImageIndex Getter
	SetImageIndex(value int32)               // property ImageIndex Setter
	MinHeight() int32                        // property MinHeight Getter
	SetMinHeight(value int32)                // property MinHeight Setter
	MinWidth() int32                         // property MinWidth Getter
	SetMinWidth(value int32)                 // property MinWidth Setter
	ParentColor() bool                       // property ParentColor Getter
	SetParentColor(value bool)               // property ParentColor Setter
	ParentBitmap() bool                      // property ParentBitmap Getter
	SetParentBitmap(value bool)              // property ParentBitmap Setter
	Text() string                            // property Text Getter
	SetText(value string)                    // property Text Setter
	Visible() bool                           // property Visible Getter
	SetVisible(value bool)                   // property Visible Setter
	Width() int32                            // property Width Getter
	SetWidth(value int32)                    // property Width Setter
}

type TCoolBand struct {
	TCollectionItem
}

func (m *TCoolBand) AutosizeWidth() {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(1, m.Instance())
}

func (m *TCoolBand) InvalidateCoolBar(sender IObject) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(sender))
}

func (m *TCoolBand) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(3, m.Instance())
	return int32(r)
}

func (m *TCoolBand) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(4, m.Instance())
	return int32(r)
}

func (m *TCoolBand) Right() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TCoolBand) Top() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCoolBand) Bitmap() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := coolBandAPI().SysCallN(7, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TCoolBand) SetBitmap(value IBitmap) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoolBand) BorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(8, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TCoolBand) SetBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCoolBand) Break() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBandAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBand) SetBreak(value bool) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBand) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(10, 0, m.Instance())
	return types.TColor(r)
}

func (m *TCoolBand) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TCoolBand) Control() IControl {
	if !m.IsValid() {
		return nil
	}
	r := coolBandAPI().SysCallN(11, 0, m.Instance())
	return AsControl(r)
}

func (m *TCoolBand) SetControl(value IControl) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(11, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCoolBand) FixedBackground() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBandAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBand) SetFixedBackground(value bool) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBand) FixedSize() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBandAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBand) SetFixedSize(value bool) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBand) HorizontalOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBandAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBand) SetHorizontalOnly(value bool) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBand) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TCoolBand) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCoolBand) MinHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TCoolBand) SetMinHeight(value int32) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TCoolBand) MinWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TCoolBand) SetMinWidth(value int32) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCoolBand) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBandAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBand) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBand) ParentBitmap() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBandAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBand) SetParentBitmap(value bool) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBand) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := coolBandAPI().SysCallN(20, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCoolBand) SetText(value string) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(20, 1, m.Instance(), api.PasStr(value))
}

func (m *TCoolBand) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := coolBandAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCoolBand) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TCoolBand) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := coolBandAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TCoolBand) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	coolBandAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

// NewCoolBand class constructor
func NewCoolBand(collection ICollection) ICoolBand {
	r := coolBandAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsCoolBand(r)
}

var (
	coolBandOnce   base.Once
	coolBandImport *imports.Imports = nil
)

func coolBandAPI() *imports.Imports {
	coolBandOnce.Do(func() {
		coolBandImport = api.NewDefaultImports()
		coolBandImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoolBand_Create", 0), // constructor NewCoolBand
			/* 1 */ imports.NewTable("TCoolBand_AutosizeWidth", 0), // procedure AutosizeWidth
			/* 2 */ imports.NewTable("TCoolBand_InvalidateCoolBar", 0), // procedure InvalidateCoolBar
			/* 3 */ imports.NewTable("TCoolBand_Height", 0), // property Height
			/* 4 */ imports.NewTable("TCoolBand_Left", 0), // property Left
			/* 5 */ imports.NewTable("TCoolBand_Right", 0), // property Right
			/* 6 */ imports.NewTable("TCoolBand_Top", 0), // property Top
			/* 7 */ imports.NewTable("TCoolBand_Bitmap", 0), // property Bitmap
			/* 8 */ imports.NewTable("TCoolBand_BorderStyle", 0), // property BorderStyle
			/* 9 */ imports.NewTable("TCoolBand_Break", 0), // property Break
			/* 10 */ imports.NewTable("TCoolBand_Color", 0), // property Color
			/* 11 */ imports.NewTable("TCoolBand_Control", 0), // property Control
			/* 12 */ imports.NewTable("TCoolBand_FixedBackground", 0), // property FixedBackground
			/* 13 */ imports.NewTable("TCoolBand_FixedSize", 0), // property FixedSize
			/* 14 */ imports.NewTable("TCoolBand_HorizontalOnly", 0), // property HorizontalOnly
			/* 15 */ imports.NewTable("TCoolBand_ImageIndex", 0), // property ImageIndex
			/* 16 */ imports.NewTable("TCoolBand_MinHeight", 0), // property MinHeight
			/* 17 */ imports.NewTable("TCoolBand_MinWidth", 0), // property MinWidth
			/* 18 */ imports.NewTable("TCoolBand_ParentColor", 0), // property ParentColor
			/* 19 */ imports.NewTable("TCoolBand_ParentBitmap", 0), // property ParentBitmap
			/* 20 */ imports.NewTable("TCoolBand_Text", 0), // property Text
			/* 21 */ imports.NewTable("TCoolBand_Visible", 0), // property Visible
			/* 22 */ imports.NewTable("TCoolBand_Width", 0), // property Width
		}
	})
	return coolBandImport
}
