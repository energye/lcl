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

// IEditButton Parent: ICustomEditButton
type IEditButton interface {
	ICustomEditButton
	AutoSelected() bool                   // property
	SetAutoSelected(AValue bool)          // property
	Button() ISpeedButton                 // property
	Edit() IEbEdit                        // property
	AutoSelect() bool                     // property
	SetAutoSelect(AValue bool)            // property
	ButtonCaption() string                // property
	SetButtonCaption(AValue string)       // property
	ButtonCursor() TCursor                // property
	SetButtonCursor(AValue TCursor)       // property
	ButtonHint() string                   // property
	SetButtonHint(AValue string)          // property
	ButtonOnlyWhenFocused() bool          // property
	SetButtonOnlyWhenFocused(AValue bool) // property
	ButtonWidth() int32                   // property
	SetButtonWidth(AValue int32)          // property
	DirectInput() bool                    // property
	SetDirectInput(AValue bool)           // property
	Flat() bool                           // property
	SetFlat(AValue bool)                  // property
	FocusOnButtonClick() bool             // property
	SetFocusOnButtonClick(AValue bool)    // property
	Glyph() IBitmap                       // property
	SetGlyph(AValue IBitmap)              // property
	Images() ICustomImageList             // property
	SetImages(AValue ICustomImageList)    // property
	ImageIndex() TImageIndex              // property
	SetImageIndex(AValue TImageIndex)     // property
	ImageWidth() int32                    // property
	SetImageWidth(AValue int32)           // property
	Layout() TLeftRight                   // property
	SetLayout(AValue TLeftRight)          // property
	NumGlyphs() int32                     // property
	SetNumGlyphs(AValue int32)            // property
	ParentFont() bool                     // property
	SetParentFont(AValue bool)            // property
	ParentShowHint() bool                 // property
	SetParentShowHint(AValue bool)        // property
	Spacing() int32                       // property
	SetSpacing(AValue int32)              // property
	SetOnButtonClick(fn TNotifyEvent)     // property event
}

// TEditButton Parent: TCustomEditButton
type TEditButton struct {
	TCustomEditButton
	buttonClickPtr uintptr
}

func NewEditButton(AOwner IComponent) IEditButton {
	r1 := editButtonImportAPI().SysCallN(9, GetObjectUintptr(AOwner))
	return AsEditButton(r1)
}

func (m *TEditButton) AutoSelected() bool {
	r1 := editButtonImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetAutoSelected(AValue bool) {
	editButtonImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) Button() ISpeedButton {
	r1 := editButtonImportAPI().SysCallN(2, m.Instance())
	return AsSpeedButton(r1)
}

func (m *TEditButton) Edit() IEbEdit {
	r1 := editButtonImportAPI().SysCallN(11, m.Instance())
	return AsEbEdit(r1)
}

func (m *TEditButton) AutoSelect() bool {
	r1 := editButtonImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetAutoSelect(AValue bool) {
	editButtonImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) ButtonCaption() string {
	r1 := editButtonImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TEditButton) SetButtonCaption(AValue string) {
	editButtonImportAPI().SysCallN(3, 1, m.Instance(), PascalStr(AValue))
}

func (m *TEditButton) ButtonCursor() TCursor {
	r1 := editButtonImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TEditButton) SetButtonCursor(AValue TCursor) {
	editButtonImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) ButtonHint() string {
	r1 := editButtonImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TEditButton) SetButtonHint(AValue string) {
	editButtonImportAPI().SysCallN(5, 1, m.Instance(), PascalStr(AValue))
}

func (m *TEditButton) ButtonOnlyWhenFocused() bool {
	r1 := editButtonImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetButtonOnlyWhenFocused(AValue bool) {
	editButtonImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) ButtonWidth() int32 {
	r1 := editButtonImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TEditButton) SetButtonWidth(AValue int32) {
	editButtonImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) DirectInput() bool {
	r1 := editButtonImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetDirectInput(AValue bool) {
	editButtonImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) Flat() bool {
	r1 := editButtonImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetFlat(AValue bool) {
	editButtonImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) FocusOnButtonClick() bool {
	r1 := editButtonImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetFocusOnButtonClick(AValue bool) {
	editButtonImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) Glyph() IBitmap {
	r1 := editButtonImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TEditButton) SetGlyph(AValue IBitmap) {
	editButtonImportAPI().SysCallN(14, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TEditButton) Images() ICustomImageList {
	r1 := editButtonImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TEditButton) SetImages(AValue ICustomImageList) {
	editButtonImportAPI().SysCallN(17, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TEditButton) ImageIndex() TImageIndex {
	r1 := editButtonImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TEditButton) SetImageIndex(AValue TImageIndex) {
	editButtonImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) ImageWidth() int32 {
	r1 := editButtonImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TEditButton) SetImageWidth(AValue int32) {
	editButtonImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) Layout() TLeftRight {
	r1 := editButtonImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TLeftRight(r1)
}

func (m *TEditButton) SetLayout(AValue TLeftRight) {
	editButtonImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) NumGlyphs() int32 {
	r1 := editButtonImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TEditButton) SetNumGlyphs(AValue int32) {
	editButtonImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TEditButton) ParentFont() bool {
	r1 := editButtonImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetParentFont(AValue bool) {
	editButtonImportAPI().SysCallN(20, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) ParentShowHint() bool {
	r1 := editButtonImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEditButton) SetParentShowHint(AValue bool) {
	editButtonImportAPI().SysCallN(21, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEditButton) Spacing() int32 {
	r1 := editButtonImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TEditButton) SetSpacing(AValue int32) {
	editButtonImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func EditButtonClass() TClass {
	ret := editButtonImportAPI().SysCallN(8)
	return TClass(ret)
}

func (m *TEditButton) SetOnButtonClick(fn TNotifyEvent) {
	if m.buttonClickPtr != 0 {
		RemoveEventElement(m.buttonClickPtr)
	}
	m.buttonClickPtr = MakeEventDataPtr(fn)
	editButtonImportAPI().SysCallN(22, m.Instance(), m.buttonClickPtr)
}

var (
	editButtonImport       *imports.Imports = nil
	editButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("EditButton_AutoSelect", 0),
		/*1*/ imports.NewTable("EditButton_AutoSelected", 0),
		/*2*/ imports.NewTable("EditButton_Button", 0),
		/*3*/ imports.NewTable("EditButton_ButtonCaption", 0),
		/*4*/ imports.NewTable("EditButton_ButtonCursor", 0),
		/*5*/ imports.NewTable("EditButton_ButtonHint", 0),
		/*6*/ imports.NewTable("EditButton_ButtonOnlyWhenFocused", 0),
		/*7*/ imports.NewTable("EditButton_ButtonWidth", 0),
		/*8*/ imports.NewTable("EditButton_Class", 0),
		/*9*/ imports.NewTable("EditButton_Create", 0),
		/*10*/ imports.NewTable("EditButton_DirectInput", 0),
		/*11*/ imports.NewTable("EditButton_Edit", 0),
		/*12*/ imports.NewTable("EditButton_Flat", 0),
		/*13*/ imports.NewTable("EditButton_FocusOnButtonClick", 0),
		/*14*/ imports.NewTable("EditButton_Glyph", 0),
		/*15*/ imports.NewTable("EditButton_ImageIndex", 0),
		/*16*/ imports.NewTable("EditButton_ImageWidth", 0),
		/*17*/ imports.NewTable("EditButton_Images", 0),
		/*18*/ imports.NewTable("EditButton_Layout", 0),
		/*19*/ imports.NewTable("EditButton_NumGlyphs", 0),
		/*20*/ imports.NewTable("EditButton_ParentFont", 0),
		/*21*/ imports.NewTable("EditButton_ParentShowHint", 0),
		/*22*/ imports.NewTable("EditButton_SetOnButtonClick", 0),
		/*23*/ imports.NewTable("EditButton_Spacing", 0),
	}
)

func editButtonImportAPI() *imports.Imports {
	if editButtonImport == nil {
		editButtonImport = NewDefaultImports()
		editButtonImport.SetImportTable(editButtonImportTables)
		editButtonImportTables = nil
	}
	return editButtonImport
}
