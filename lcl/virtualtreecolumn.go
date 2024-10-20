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

// IVirtualTreeColumn Parent: ICollectionItem
type IVirtualTreeColumn interface {
	ICollectionItem
	Left() int32                                   // property
	Owner() IVirtualTreeColumns                    // property
	Alignment() TAlignment                         // property
	SetAlignment(AValue TAlignment)                // property
	BiDiMode() TBiDiMode                           // property
	SetBiDiMode(AValue TBiDiMode)                  // property
	CaptionAlignment() TAlignment                  // property
	SetCaptionAlignment(AValue TAlignment)         // property
	CaptionText() string                           // property
	CheckType() TCheckType                         // property
	SetCheckType(AValue TCheckType)                // property
	CheckState() TCheckState                       // property
	SetCheckState(AValue TCheckState)              // property
	CheckBox() bool                                // property
	SetCheckBox(AValue bool)                       // property
	Color() TColor                                 // property
	SetColor(AValue TColor)                        // property
	DefaultSortDirection() TSortDirection          // property
	SetDefaultSortDirection(AValue TSortDirection) // property
	Hint() string                                  // property
	SetHint(AValue string)                         // property
	ImageIndex() TImageIndex                       // property
	SetImageIndex(AValue TImageIndex)              // property
	Layout() TVTHeaderColumnLayout                 // property
	SetLayout(AValue TVTHeaderColumnLayout)        // property
	Margin() int32                                 // property
	SetMargin(AValue int32)                        // property
	MaxWidth() int32                               // property
	SetMaxWidth(AValue int32)                      // property
	MinWidth() int32                               // property
	SetMinWidth(AValue int32)                      // property
	Options() TVTColumnOptions                     // property
	SetOptions(AValue TVTColumnOptions)            // property
	Position() TColumnPosition                     // property
	SetPosition(AValue TColumnPosition)            // property
	Spacing() int32                                // property
	SetSpacing(AValue int32)                       // property
	Style() TVirtualTreeColumnStyle                // property
	SetStyle(AValue TVirtualTreeColumnStyle)       // property
	Tag() uint32                                   // property
	SetTag(AValue uint32)                          // property
	Text() string                                  // property
	SetText(AValue string)                         // property
	Width() int32                                  // property
	SetWidth(AValue int32)                         // property
	GetRect() (resultRect TRect)                   // function
	UseRightToLeftReading() bool                   // function
	LoadFromStream(Stream IStream, Version int32)  // procedure
	ParentBiDiModeChanged()                        // procedure
	ParentColorChanged()                           // procedure
	RestoreLastWidth()                             // procedure
	SaveToStream(Stream IStream)                   // procedure
}

// TVirtualTreeColumn Parent: TCollectionItem
type TVirtualTreeColumn struct {
	TCollectionItem
}

func NewVirtualTreeColumn(Collection ICollection) IVirtualTreeColumn {
	r1 := virtualTreeColumnImportAPI().SysCallN(9, GetObjectUintptr(Collection))
	return AsVirtualTreeColumn(r1)
}

func (m *TVirtualTreeColumn) Left() int32 {
	r1 := virtualTreeColumnImportAPI().SysCallN(15, m.Instance())
	return int32(r1)
}

func (m *TVirtualTreeColumn) Owner() IVirtualTreeColumns {
	r1 := virtualTreeColumnImportAPI().SysCallN(21, m.Instance())
	return AsVirtualTreeColumns(r1)
}

func (m *TVirtualTreeColumn) Alignment() TAlignment {
	r1 := virtualTreeColumnImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TVirtualTreeColumn) SetAlignment(AValue TAlignment) {
	virtualTreeColumnImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) BiDiMode() TBiDiMode {
	r1 := virtualTreeColumnImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TVirtualTreeColumn) SetBiDiMode(AValue TBiDiMode) {
	virtualTreeColumnImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) CaptionAlignment() TAlignment {
	r1 := virtualTreeColumnImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TVirtualTreeColumn) SetCaptionAlignment(AValue TAlignment) {
	virtualTreeColumnImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) CaptionText() string {
	r1 := virtualTreeColumnImportAPI().SysCallN(3, m.Instance())
	return GoStr(r1)
}

func (m *TVirtualTreeColumn) CheckType() TCheckType {
	r1 := virtualTreeColumnImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TCheckType(r1)
}

func (m *TVirtualTreeColumn) SetCheckType(AValue TCheckType) {
	virtualTreeColumnImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) CheckState() TCheckState {
	r1 := virtualTreeColumnImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TCheckState(r1)
}

func (m *TVirtualTreeColumn) SetCheckState(AValue TCheckState) {
	virtualTreeColumnImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) CheckBox() bool {
	r1 := virtualTreeColumnImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TVirtualTreeColumn) SetCheckBox(AValue bool) {
	virtualTreeColumnImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TVirtualTreeColumn) Color() TColor {
	r1 := virtualTreeColumnImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVirtualTreeColumn) SetColor(AValue TColor) {
	virtualTreeColumnImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) DefaultSortDirection() TSortDirection {
	r1 := virtualTreeColumnImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TSortDirection(r1)
}

func (m *TVirtualTreeColumn) SetDefaultSortDirection(AValue TSortDirection) {
	virtualTreeColumnImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Hint() string {
	r1 := virtualTreeColumnImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TVirtualTreeColumn) SetHint(AValue string) {
	virtualTreeColumnImportAPI().SysCallN(12, 1, m.Instance(), PascalStr(AValue))
}

func (m *TVirtualTreeColumn) ImageIndex() TImageIndex {
	r1 := virtualTreeColumnImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TVirtualTreeColumn) SetImageIndex(AValue TImageIndex) {
	virtualTreeColumnImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Layout() TVTHeaderColumnLayout {
	r1 := virtualTreeColumnImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TVTHeaderColumnLayout(r1)
}

func (m *TVirtualTreeColumn) SetLayout(AValue TVTHeaderColumnLayout) {
	virtualTreeColumnImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Margin() int32 {
	r1 := virtualTreeColumnImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetMargin(AValue int32) {
	virtualTreeColumnImportAPI().SysCallN(17, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) MaxWidth() int32 {
	r1 := virtualTreeColumnImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetMaxWidth(AValue int32) {
	virtualTreeColumnImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) MinWidth() int32 {
	r1 := virtualTreeColumnImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetMinWidth(AValue int32) {
	virtualTreeColumnImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Options() TVTColumnOptions {
	r1 := virtualTreeColumnImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return TVTColumnOptions(r1)
}

func (m *TVirtualTreeColumn) SetOptions(AValue TVTColumnOptions) {
	virtualTreeColumnImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Position() TColumnPosition {
	r1 := virtualTreeColumnImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return TColumnPosition(r1)
}

func (m *TVirtualTreeColumn) SetPosition(AValue TColumnPosition) {
	virtualTreeColumnImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Spacing() int32 {
	r1 := virtualTreeColumnImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetSpacing(AValue int32) {
	virtualTreeColumnImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Style() TVirtualTreeColumnStyle {
	r1 := virtualTreeColumnImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return TVirtualTreeColumnStyle(r1)
}

func (m *TVirtualTreeColumn) SetStyle(AValue TVirtualTreeColumnStyle) {
	virtualTreeColumnImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Tag() uint32 {
	r1 := virtualTreeColumnImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TVirtualTreeColumn) SetTag(AValue uint32) {
	virtualTreeColumnImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Text() string {
	r1 := virtualTreeColumnImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TVirtualTreeColumn) SetText(AValue string) {
	virtualTreeColumnImportAPI().SysCallN(30, 1, m.Instance(), PascalStr(AValue))
}

func (m *TVirtualTreeColumn) Width() int32 {
	r1 := virtualTreeColumnImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetWidth(AValue int32) {
	virtualTreeColumnImportAPI().SysCallN(32, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) GetRect() (resultRect TRect) {
	virtualTreeColumnImportAPI().SysCallN(11, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TVirtualTreeColumn) UseRightToLeftReading() bool {
	r1 := virtualTreeColumnImportAPI().SysCallN(31, m.Instance())
	return GoBool(r1)
}

func VirtualTreeColumnClass() TClass {
	ret := virtualTreeColumnImportAPI().SysCallN(7)
	return TClass(ret)
}

func (m *TVirtualTreeColumn) LoadFromStream(Stream IStream, Version int32) {
	virtualTreeColumnImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(Stream), uintptr(Version))
}

func (m *TVirtualTreeColumn) ParentBiDiModeChanged() {
	virtualTreeColumnImportAPI().SysCallN(22, m.Instance())
}

func (m *TVirtualTreeColumn) ParentColorChanged() {
	virtualTreeColumnImportAPI().SysCallN(23, m.Instance())
}

func (m *TVirtualTreeColumn) RestoreLastWidth() {
	virtualTreeColumnImportAPI().SysCallN(25, m.Instance())
}

func (m *TVirtualTreeColumn) SaveToStream(Stream IStream) {
	virtualTreeColumnImportAPI().SysCallN(26, m.Instance(), GetObjectUintptr(Stream))
}

var (
	virtualTreeColumnImport       *imports.Imports = nil
	virtualTreeColumnImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("VirtualTreeColumn_Alignment", 0),
		/*1*/ imports.NewTable("VirtualTreeColumn_BiDiMode", 0),
		/*2*/ imports.NewTable("VirtualTreeColumn_CaptionAlignment", 0),
		/*3*/ imports.NewTable("VirtualTreeColumn_CaptionText", 0),
		/*4*/ imports.NewTable("VirtualTreeColumn_CheckBox", 0),
		/*5*/ imports.NewTable("VirtualTreeColumn_CheckState", 0),
		/*6*/ imports.NewTable("VirtualTreeColumn_CheckType", 0),
		/*7*/ imports.NewTable("VirtualTreeColumn_Class", 0),
		/*8*/ imports.NewTable("VirtualTreeColumn_Color", 0),
		/*9*/ imports.NewTable("VirtualTreeColumn_Create", 0),
		/*10*/ imports.NewTable("VirtualTreeColumn_DefaultSortDirection", 0),
		/*11*/ imports.NewTable("VirtualTreeColumn_GetRect", 0),
		/*12*/ imports.NewTable("VirtualTreeColumn_Hint", 0),
		/*13*/ imports.NewTable("VirtualTreeColumn_ImageIndex", 0),
		/*14*/ imports.NewTable("VirtualTreeColumn_Layout", 0),
		/*15*/ imports.NewTable("VirtualTreeColumn_Left", 0),
		/*16*/ imports.NewTable("VirtualTreeColumn_LoadFromStream", 0),
		/*17*/ imports.NewTable("VirtualTreeColumn_Margin", 0),
		/*18*/ imports.NewTable("VirtualTreeColumn_MaxWidth", 0),
		/*19*/ imports.NewTable("VirtualTreeColumn_MinWidth", 0),
		/*20*/ imports.NewTable("VirtualTreeColumn_Options", 0),
		/*21*/ imports.NewTable("VirtualTreeColumn_Owner", 0),
		/*22*/ imports.NewTable("VirtualTreeColumn_ParentBiDiModeChanged", 0),
		/*23*/ imports.NewTable("VirtualTreeColumn_ParentColorChanged", 0),
		/*24*/ imports.NewTable("VirtualTreeColumn_Position", 0),
		/*25*/ imports.NewTable("VirtualTreeColumn_RestoreLastWidth", 0),
		/*26*/ imports.NewTable("VirtualTreeColumn_SaveToStream", 0),
		/*27*/ imports.NewTable("VirtualTreeColumn_Spacing", 0),
		/*28*/ imports.NewTable("VirtualTreeColumn_Style", 0),
		/*29*/ imports.NewTable("VirtualTreeColumn_Tag", 0),
		/*30*/ imports.NewTable("VirtualTreeColumn_Text", 0),
		/*31*/ imports.NewTable("VirtualTreeColumn_UseRightToLeftReading", 0),
		/*32*/ imports.NewTable("VirtualTreeColumn_Width", 0),
	}
)

func virtualTreeColumnImportAPI() *imports.Imports {
	if virtualTreeColumnImport == nil {
		virtualTreeColumnImport = NewDefaultImports()
		virtualTreeColumnImport.SetImportTable(virtualTreeColumnImportTables)
		virtualTreeColumnImportTables = nil
	}
	return virtualTreeColumnImport
}
