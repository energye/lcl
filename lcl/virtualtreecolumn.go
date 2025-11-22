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

// IVirtualTreeColumn Parent: ICollectionItem
type IVirtualTreeColumn interface {
	ICollectionItem
	GetRect() types.TRect                               // function
	UseRightToLeftReading() bool                        // function
	LoadFromStream(stream IStream, version int32)       // procedure
	ParentBiDiModeChanged()                             // procedure
	ParentColorChanged()                                // procedure
	RestoreLastWidth()                                  // procedure
	SaveToStream(stream IStream)                        // procedure
	Left() int32                                        // property Left Getter
	Owner() IVirtualTreeColumns                         // property Owner Getter
	Alignment() types.TAlignment                        // property Alignment Getter
	SetAlignment(value types.TAlignment)                // property Alignment Setter
	BiDiMode() types.TBiDiMode                          // property BiDiMode Getter
	SetBiDiMode(value types.TBiDiMode)                  // property BiDiMode Setter
	CaptionAlignment() types.TAlignment                 // property CaptionAlignment Getter
	SetCaptionAlignment(value types.TAlignment)         // property CaptionAlignment Setter
	CaptionText() string                                // property CaptionText Getter
	CheckType() types.TCheckType                        // property CheckType Getter
	SetCheckType(value types.TCheckType)                // property CheckType Setter
	CheckState() types.TCheckState                      // property CheckState Getter
	SetCheckState(value types.TCheckState)              // property CheckState Setter
	CheckBox() bool                                     // property CheckBox Getter
	SetCheckBox(value bool)                             // property CheckBox Setter
	Color() types.TColor                                // property Color Getter
	SetColor(value types.TColor)                        // property Color Setter
	DefaultSortDirection() types.TSortDirection         // property DefaultSortDirection Getter
	SetDefaultSortDirection(value types.TSortDirection) // property DefaultSortDirection Setter
	Hint() string                                       // property Hint Getter
	SetHint(value string)                               // property Hint Setter
	ImageIndex() int32                                  // property ImageIndex Getter
	SetImageIndex(value int32)                          // property ImageIndex Setter
	Layout() types.TVTHeaderColumnLayout                // property Layout Getter
	SetLayout(value types.TVTHeaderColumnLayout)        // property Layout Setter
	Margin() int32                                      // property Margin Getter
	SetMargin(value int32)                              // property Margin Setter
	MaxWidth() int32                                    // property MaxWidth Getter
	SetMaxWidth(value int32)                            // property MaxWidth Setter
	MinWidth() int32                                    // property MinWidth Getter
	SetMinWidth(value int32)                            // property MinWidth Setter
	Options() types.TVTColumnOptions                    // property Options Getter
	SetOptions(value types.TVTColumnOptions)            // property Options Setter
	Position() uint32                                   // property Position Getter
	SetPosition(value uint32)                           // property Position Setter
	Spacing() int32                                     // property Spacing Getter
	SetSpacing(value int32)                             // property Spacing Setter
	Style() types.TVirtualTreeColumnStyle               // property Style Getter
	SetStyle(value types.TVirtualTreeColumnStyle)       // property Style Setter
	Tag() uint32                                        // property Tag Getter
	SetTag(value uint32)                                // property Tag Setter
	Text() string                                       // property Text Getter
	SetText(value string)                               // property Text Setter
	Width() int32                                       // property Width Getter
	SetWidth(value int32)                               // property Width Setter
}

type TVirtualTreeColumn struct {
	TCollectionItem
}

func (m *TVirtualTreeColumn) GetRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TVirtualTreeColumn) UseRightToLeftReading() bool {
	if !m.IsValid() {
		return false
	}
	r := virtualTreeColumnAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TVirtualTreeColumn) LoadFromStream(stream IStream, version int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(stream), uintptr(version))
}

func (m *TVirtualTreeColumn) ParentBiDiModeChanged() {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(4, m.Instance())
}

func (m *TVirtualTreeColumn) ParentColorChanged() {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(5, m.Instance())
}

func (m *TVirtualTreeColumn) RestoreLastWidth() {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(6, m.Instance())
}

func (m *TVirtualTreeColumn) SaveToStream(stream IStream) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(stream))
}

func (m *TVirtualTreeColumn) Left() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(8, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumn) Owner() IVirtualTreeColumns {
	if !m.IsValid() {
		return nil
	}
	r := virtualTreeColumnAPI().SysCallN(9, m.Instance())
	return AsVirtualTreeColumns(r)
}

func (m *TVirtualTreeColumn) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(10, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TVirtualTreeColumn) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) BiDiMode() types.TBiDiMode {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(11, 0, m.Instance())
	return types.TBiDiMode(r)
}

func (m *TVirtualTreeColumn) SetBiDiMode(value types.TBiDiMode) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) CaptionAlignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(12, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TVirtualTreeColumn) SetCaptionAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) CaptionText() string {
	if !m.IsValid() {
		return ""
	}
	r := virtualTreeColumnAPI().SysCallN(13, m.Instance())
	return api.GoStr(r)
}

func (m *TVirtualTreeColumn) CheckType() types.TCheckType {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(14, 0, m.Instance())
	return types.TCheckType(r)
}

func (m *TVirtualTreeColumn) SetCheckType(value types.TCheckType) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) CheckState() types.TCheckState {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(15, 0, m.Instance())
	return types.TCheckState(r)
}

func (m *TVirtualTreeColumn) SetCheckState(value types.TCheckState) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) CheckBox() bool {
	if !m.IsValid() {
		return false
	}
	r := virtualTreeColumnAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TVirtualTreeColumn) SetCheckBox(value bool) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TVirtualTreeColumn) Color() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(17, 0, m.Instance())
	return types.TColor(r)
}

func (m *TVirtualTreeColumn) SetColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) DefaultSortDirection() types.TSortDirection {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(18, 0, m.Instance())
	return types.TSortDirection(r)
}

func (m *TVirtualTreeColumn) SetDefaultSortDirection(value types.TSortDirection) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) Hint() string {
	if !m.IsValid() {
		return ""
	}
	r := virtualTreeColumnAPI().SysCallN(19, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TVirtualTreeColumn) SetHint(value string) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(19, 1, m.Instance(), api.PasStr(value))
}

func (m *TVirtualTreeColumn) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumn) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) Layout() types.TVTHeaderColumnLayout {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(21, 0, m.Instance())
	return types.TVTHeaderColumnLayout(r)
}

func (m *TVirtualTreeColumn) SetLayout(value types.TVTHeaderColumnLayout) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) Margin() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumn) SetMargin(value int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) MaxWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumn) SetMaxWidth(value int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) MinWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(24, 0, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumn) SetMinWidth(value int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) Options() types.TVTColumnOptions {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(25, 0, m.Instance())
	return types.TVTColumnOptions(r)
}

func (m *TVirtualTreeColumn) SetOptions(value types.TVTColumnOptions) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(25, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) Position() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(26, 0, m.Instance())
	return uint32(r)
}

func (m *TVirtualTreeColumn) SetPosition(value uint32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) Spacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(27, 0, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumn) SetSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) Style() types.TVirtualTreeColumnStyle {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(28, 0, m.Instance())
	return types.TVirtualTreeColumnStyle(r)
}

func (m *TVirtualTreeColumn) SetStyle(value types.TVirtualTreeColumnStyle) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) Tag() uint32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(29, 0, m.Instance())
	return uint32(r)
}

func (m *TVirtualTreeColumn) SetTag(value uint32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TVirtualTreeColumn) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := virtualTreeColumnAPI().SysCallN(30, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TVirtualTreeColumn) SetText(value string) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(30, 1, m.Instance(), api.PasStr(value))
}

func (m *TVirtualTreeColumn) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := virtualTreeColumnAPI().SysCallN(31, 0, m.Instance())
	return int32(r)
}

func (m *TVirtualTreeColumn) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	virtualTreeColumnAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

// NewVirtualTreeColumn class constructor
func NewVirtualTreeColumn(collection ICollection) IVirtualTreeColumn {
	r := virtualTreeColumnAPI().SysCallN(0, base.GetObjectUintptr(collection))
	return AsVirtualTreeColumn(r)
}

var (
	virtualTreeColumnOnce   base.Once
	virtualTreeColumnImport *imports.Imports = nil
)

func virtualTreeColumnAPI() *imports.Imports {
	virtualTreeColumnOnce.Do(func() {
		virtualTreeColumnImport = api.NewDefaultImports()
		virtualTreeColumnImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TVirtualTreeColumn_Create", 0), // constructor NewVirtualTreeColumn
			/* 1 */ imports.NewTable("TVirtualTreeColumn_GetRect", 0), // function GetRect
			/* 2 */ imports.NewTable("TVirtualTreeColumn_UseRightToLeftReading", 0), // function UseRightToLeftReading
			/* 3 */ imports.NewTable("TVirtualTreeColumn_LoadFromStream", 0), // procedure LoadFromStream
			/* 4 */ imports.NewTable("TVirtualTreeColumn_ParentBiDiModeChanged", 0), // procedure ParentBiDiModeChanged
			/* 5 */ imports.NewTable("TVirtualTreeColumn_ParentColorChanged", 0), // procedure ParentColorChanged
			/* 6 */ imports.NewTable("TVirtualTreeColumn_RestoreLastWidth", 0), // procedure RestoreLastWidth
			/* 7 */ imports.NewTable("TVirtualTreeColumn_SaveToStream", 0), // procedure SaveToStream
			/* 8 */ imports.NewTable("TVirtualTreeColumn_Left", 0), // property Left
			/* 9 */ imports.NewTable("TVirtualTreeColumn_Owner", 0), // property Owner
			/* 10 */ imports.NewTable("TVirtualTreeColumn_Alignment", 0), // property Alignment
			/* 11 */ imports.NewTable("TVirtualTreeColumn_BiDiMode", 0), // property BiDiMode
			/* 12 */ imports.NewTable("TVirtualTreeColumn_CaptionAlignment", 0), // property CaptionAlignment
			/* 13 */ imports.NewTable("TVirtualTreeColumn_CaptionText", 0), // property CaptionText
			/* 14 */ imports.NewTable("TVirtualTreeColumn_CheckType", 0), // property CheckType
			/* 15 */ imports.NewTable("TVirtualTreeColumn_CheckState", 0), // property CheckState
			/* 16 */ imports.NewTable("TVirtualTreeColumn_CheckBox", 0), // property CheckBox
			/* 17 */ imports.NewTable("TVirtualTreeColumn_Color", 0), // property Color
			/* 18 */ imports.NewTable("TVirtualTreeColumn_DefaultSortDirection", 0), // property DefaultSortDirection
			/* 19 */ imports.NewTable("TVirtualTreeColumn_Hint", 0), // property Hint
			/* 20 */ imports.NewTable("TVirtualTreeColumn_ImageIndex", 0), // property ImageIndex
			/* 21 */ imports.NewTable("TVirtualTreeColumn_Layout", 0), // property Layout
			/* 22 */ imports.NewTable("TVirtualTreeColumn_Margin", 0), // property Margin
			/* 23 */ imports.NewTable("TVirtualTreeColumn_MaxWidth", 0), // property MaxWidth
			/* 24 */ imports.NewTable("TVirtualTreeColumn_MinWidth", 0), // property MinWidth
			/* 25 */ imports.NewTable("TVirtualTreeColumn_Options", 0), // property Options
			/* 26 */ imports.NewTable("TVirtualTreeColumn_Position", 0), // property Position
			/* 27 */ imports.NewTable("TVirtualTreeColumn_Spacing", 0), // property Spacing
			/* 28 */ imports.NewTable("TVirtualTreeColumn_Style", 0), // property Style
			/* 29 */ imports.NewTable("TVirtualTreeColumn_Tag", 0), // property Tag
			/* 30 */ imports.NewTable("TVirtualTreeColumn_Text", 0), // property Text
			/* 31 */ imports.NewTable("TVirtualTreeColumn_Width", 0), // property Width
		}
	})
	return virtualTreeColumnImport
}
