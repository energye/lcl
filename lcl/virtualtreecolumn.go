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
	r1 := LCL().SysCallN(5989, GetObjectUintptr(Collection))
	return AsVirtualTreeColumn(r1)
}

func (m *TVirtualTreeColumn) Left() int32 {
	r1 := LCL().SysCallN(5995, m.Instance())
	return int32(r1)
}

func (m *TVirtualTreeColumn) Owner() IVirtualTreeColumns {
	r1 := LCL().SysCallN(6001, m.Instance())
	return AsVirtualTreeColumns(r1)
}

func (m *TVirtualTreeColumn) Alignment() TAlignment {
	r1 := LCL().SysCallN(5980, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TVirtualTreeColumn) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(5980, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) BiDiMode() TBiDiMode {
	r1 := LCL().SysCallN(5981, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TVirtualTreeColumn) SetBiDiMode(AValue TBiDiMode) {
	LCL().SysCallN(5981, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) CaptionAlignment() TAlignment {
	r1 := LCL().SysCallN(5982, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TVirtualTreeColumn) SetCaptionAlignment(AValue TAlignment) {
	LCL().SysCallN(5982, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) CaptionText() string {
	r1 := LCL().SysCallN(5983, m.Instance())
	return GoStr(r1)
}

func (m *TVirtualTreeColumn) CheckType() TCheckType {
	r1 := LCL().SysCallN(5986, 0, m.Instance(), 0)
	return TCheckType(r1)
}

func (m *TVirtualTreeColumn) SetCheckType(AValue TCheckType) {
	LCL().SysCallN(5986, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) CheckState() TCheckState {
	r1 := LCL().SysCallN(5985, 0, m.Instance(), 0)
	return TCheckState(r1)
}

func (m *TVirtualTreeColumn) SetCheckState(AValue TCheckState) {
	LCL().SysCallN(5985, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) CheckBox() bool {
	r1 := LCL().SysCallN(5984, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TVirtualTreeColumn) SetCheckBox(AValue bool) {
	LCL().SysCallN(5984, 1, m.Instance(), PascalBool(AValue))
}

func (m *TVirtualTreeColumn) Color() TColor {
	r1 := LCL().SysCallN(5988, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TVirtualTreeColumn) SetColor(AValue TColor) {
	LCL().SysCallN(5988, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) DefaultSortDirection() TSortDirection {
	r1 := LCL().SysCallN(5990, 0, m.Instance(), 0)
	return TSortDirection(r1)
}

func (m *TVirtualTreeColumn) SetDefaultSortDirection(AValue TSortDirection) {
	LCL().SysCallN(5990, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Hint() string {
	r1 := LCL().SysCallN(5992, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TVirtualTreeColumn) SetHint(AValue string) {
	LCL().SysCallN(5992, 1, m.Instance(), PascalStr(AValue))
}

func (m *TVirtualTreeColumn) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(5993, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TVirtualTreeColumn) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(5993, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Layout() TVTHeaderColumnLayout {
	r1 := LCL().SysCallN(5994, 0, m.Instance(), 0)
	return TVTHeaderColumnLayout(r1)
}

func (m *TVirtualTreeColumn) SetLayout(AValue TVTHeaderColumnLayout) {
	LCL().SysCallN(5994, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Margin() int32 {
	r1 := LCL().SysCallN(5997, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetMargin(AValue int32) {
	LCL().SysCallN(5997, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) MaxWidth() int32 {
	r1 := LCL().SysCallN(5998, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetMaxWidth(AValue int32) {
	LCL().SysCallN(5998, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) MinWidth() int32 {
	r1 := LCL().SysCallN(5999, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetMinWidth(AValue int32) {
	LCL().SysCallN(5999, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Options() TVTColumnOptions {
	r1 := LCL().SysCallN(6000, 0, m.Instance(), 0)
	return TVTColumnOptions(r1)
}

func (m *TVirtualTreeColumn) SetOptions(AValue TVTColumnOptions) {
	LCL().SysCallN(6000, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Position() TColumnPosition {
	r1 := LCL().SysCallN(6004, 0, m.Instance(), 0)
	return TColumnPosition(r1)
}

func (m *TVirtualTreeColumn) SetPosition(AValue TColumnPosition) {
	LCL().SysCallN(6004, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Spacing() int32 {
	r1 := LCL().SysCallN(6007, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetSpacing(AValue int32) {
	LCL().SysCallN(6007, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Style() TVirtualTreeColumnStyle {
	r1 := LCL().SysCallN(6008, 0, m.Instance(), 0)
	return TVirtualTreeColumnStyle(r1)
}

func (m *TVirtualTreeColumn) SetStyle(AValue TVirtualTreeColumnStyle) {
	LCL().SysCallN(6008, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Tag() uint32 {
	r1 := LCL().SysCallN(6009, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TVirtualTreeColumn) SetTag(AValue uint32) {
	LCL().SysCallN(6009, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) Text() string {
	r1 := LCL().SysCallN(6010, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TVirtualTreeColumn) SetText(AValue string) {
	LCL().SysCallN(6010, 1, m.Instance(), PascalStr(AValue))
}

func (m *TVirtualTreeColumn) Width() int32 {
	r1 := LCL().SysCallN(6012, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TVirtualTreeColumn) SetWidth(AValue int32) {
	LCL().SysCallN(6012, 1, m.Instance(), uintptr(AValue))
}

func (m *TVirtualTreeColumn) GetRect() (resultRect TRect) {
	LCL().SysCallN(5991, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TVirtualTreeColumn) UseRightToLeftReading() bool {
	r1 := LCL().SysCallN(6011, m.Instance())
	return GoBool(r1)
}

func VirtualTreeColumnClass() TClass {
	ret := LCL().SysCallN(5987)
	return TClass(ret)
}

func (m *TVirtualTreeColumn) LoadFromStream(Stream IStream, Version int32) {
	LCL().SysCallN(5996, m.Instance(), GetObjectUintptr(Stream), uintptr(Version))
}

func (m *TVirtualTreeColumn) ParentBiDiModeChanged() {
	LCL().SysCallN(6002, m.Instance())
}

func (m *TVirtualTreeColumn) ParentColorChanged() {
	LCL().SysCallN(6003, m.Instance())
}

func (m *TVirtualTreeColumn) RestoreLastWidth() {
	LCL().SysCallN(6005, m.Instance())
}

func (m *TVirtualTreeColumn) SaveToStream(Stream IStream) {
	LCL().SysCallN(6006, m.Instance(), GetObjectUintptr(Stream))
}
