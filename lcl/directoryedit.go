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

// IDirectoryEdit Parent: ICustomEditButton
type IDirectoryEdit interface {
	ICustomEditButton
	AutoSelected() bool                           // property
	SetAutoSelected(AValue bool)                  // property
	Directory() string                            // property
	SetDirectory(AValue string)                   // property
	RootDir() string                              // property
	SetRootDir(AValue string)                     // property
	DialogTitle() string                          // property
	SetDialogTitle(AValue string)                 // property
	DialogOptions() TOpenOptions                  // property
	SetDialogOptions(AValue TOpenOptions)         // property
	ShowHidden() bool                             // property
	SetShowHidden(AValue bool)                    // property
	ButtonCaption() string                        // property
	SetButtonCaption(AValue string)               // property
	ButtonCursor() TCursor                        // property
	SetButtonCursor(AValue TCursor)               // property
	ButtonHint() string                           // property
	SetButtonHint(AValue string)                  // property
	ButtonOnlyWhenFocused() bool                  // property
	SetButtonOnlyWhenFocused(AValue bool)         // property
	ButtonWidth() int32                           // property
	SetButtonWidth(AValue int32)                  // property
	DirectInput() bool                            // property
	SetDirectInput(AValue bool)                   // property
	Glyph() IBitmap                               // property
	SetGlyph(AValue IBitmap)                      // property
	NumGlyphs() int32                             // property
	SetNumGlyphs(AValue int32)                    // property
	Images() ICustomImageList                     // property
	SetImages(AValue ICustomImageList)            // property
	ImageIndex() TImageIndex                      // property
	SetImageIndex(AValue TImageIndex)             // property
	ImageWidth() int32                            // property
	SetImageWidth(AValue int32)                   // property
	Flat() bool                                   // property
	SetFlat(AValue bool)                          // property
	FocusOnButtonClick() bool                     // property
	SetFocusOnButtonClick(AValue bool)            // property
	AutoSelect() bool                             // property
	SetAutoSelect(AValue bool)                    // property
	DragCursor() TCursor                          // property
	SetDragCursor(AValue TCursor)                 // property
	DragMode() TDragMode                          // property
	SetDragMode(AValue TDragMode)                 // property
	Layout() TLeftRight                           // property
	SetLayout(AValue TLeftRight)                  // property
	ParentFont() bool                             // property
	SetParentFont(AValue bool)                    // property
	ParentShowHint() bool                         // property
	SetParentShowHint(AValue bool)                // property
	Spacing() int32                               // property
	SetSpacing(AValue int32)                      // property
	RunDialog()                                   // procedure
	SetOnAcceptDirectory(fn TAcceptFileNameEvent) // property event
	SetOnButtonClick(fn TNotifyEvent)             // property event
}

// TDirectoryEdit Parent: TCustomEditButton
type TDirectoryEdit struct {
	TCustomEditButton
	acceptDirectoryPtr uintptr
	buttonClickPtr     uintptr
}

func NewDirectoryEdit(AOwner IComponent) IDirectoryEdit {
	r1 := directoryEditImportAPI().SysCallN(8, GetObjectUintptr(AOwner))
	return AsDirectoryEdit(r1)
}

func (m *TDirectoryEdit) AutoSelected() bool {
	r1 := directoryEditImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetAutoSelected(AValue bool) {
	directoryEditImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Directory() string {
	r1 := directoryEditImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetDirectory(AValue string) {
	directoryEditImportAPI().SysCallN(12, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) RootDir() string {
	r1 := directoryEditImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetRootDir(AValue string) {
	directoryEditImportAPI().SysCallN(25, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) DialogTitle() string {
	r1 := directoryEditImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetDialogTitle(AValue string) {
	directoryEditImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) DialogOptions() TOpenOptions {
	r1 := directoryEditImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TOpenOptions(r1)
}

func (m *TDirectoryEdit) SetDialogOptions(AValue TOpenOptions) {
	directoryEditImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ShowHidden() bool {
	r1 := directoryEditImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetShowHidden(AValue bool) {
	directoryEditImportAPI().SysCallN(29, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ButtonCaption() string {
	r1 := directoryEditImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetButtonCaption(AValue string) {
	directoryEditImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) ButtonCursor() TCursor {
	r1 := directoryEditImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDirectoryEdit) SetButtonCursor(AValue TCursor) {
	directoryEditImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ButtonHint() string {
	r1 := directoryEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDirectoryEdit) SetButtonHint(AValue string) {
	directoryEditImportAPI().SysCallN(4, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDirectoryEdit) ButtonOnlyWhenFocused() bool {
	r1 := directoryEditImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetButtonOnlyWhenFocused(AValue bool) {
	directoryEditImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ButtonWidth() int32 {
	r1 := directoryEditImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetButtonWidth(AValue int32) {
	directoryEditImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) DirectInput() bool {
	r1 := directoryEditImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetDirectInput(AValue bool) {
	directoryEditImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Glyph() IBitmap {
	r1 := directoryEditImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TDirectoryEdit) SetGlyph(AValue IBitmap) {
	directoryEditImportAPI().SysCallN(17, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDirectoryEdit) NumGlyphs() int32 {
	r1 := directoryEditImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetNumGlyphs(AValue int32) {
	directoryEditImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Images() ICustomImageList {
	r1 := directoryEditImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TDirectoryEdit) SetImages(AValue ICustomImageList) {
	directoryEditImportAPI().SysCallN(20, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDirectoryEdit) ImageIndex() TImageIndex {
	r1 := directoryEditImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TDirectoryEdit) SetImageIndex(AValue TImageIndex) {
	directoryEditImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ImageWidth() int32 {
	r1 := directoryEditImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetImageWidth(AValue int32) {
	directoryEditImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Flat() bool {
	r1 := directoryEditImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetFlat(AValue bool) {
	directoryEditImportAPI().SysCallN(15, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) FocusOnButtonClick() bool {
	r1 := directoryEditImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetFocusOnButtonClick(AValue bool) {
	directoryEditImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) AutoSelect() bool {
	r1 := directoryEditImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetAutoSelect(AValue bool) {
	directoryEditImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) DragCursor() TCursor {
	r1 := directoryEditImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDirectoryEdit) SetDragCursor(AValue TCursor) {
	directoryEditImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) DragMode() TDragMode {
	r1 := directoryEditImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TDirectoryEdit) SetDragMode(AValue TDragMode) {
	directoryEditImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) Layout() TLeftRight {
	r1 := directoryEditImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return TLeftRight(r1)
}

func (m *TDirectoryEdit) SetLayout(AValue TLeftRight) {
	directoryEditImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TDirectoryEdit) ParentFont() bool {
	r1 := directoryEditImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetParentFont(AValue bool) {
	directoryEditImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) ParentShowHint() bool {
	r1 := directoryEditImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDirectoryEdit) SetParentShowHint(AValue bool) {
	directoryEditImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDirectoryEdit) Spacing() int32 {
	r1 := directoryEditImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDirectoryEdit) SetSpacing(AValue int32) {
	directoryEditImportAPI().SysCallN(30, 1, m.Instance(), uintptr(AValue))
}

func DirectoryEditClass() TClass {
	ret := directoryEditImportAPI().SysCallN(7)
	return TClass(ret)
}

func (m *TDirectoryEdit) RunDialog() {
	directoryEditImportAPI().SysCallN(26, m.Instance())
}

func (m *TDirectoryEdit) SetOnAcceptDirectory(fn TAcceptFileNameEvent) {
	if m.acceptDirectoryPtr != 0 {
		RemoveEventElement(m.acceptDirectoryPtr)
	}
	m.acceptDirectoryPtr = MakeEventDataPtr(fn)
	directoryEditImportAPI().SysCallN(27, m.Instance(), m.acceptDirectoryPtr)
}

func (m *TDirectoryEdit) SetOnButtonClick(fn TNotifyEvent) {
	if m.buttonClickPtr != 0 {
		RemoveEventElement(m.buttonClickPtr)
	}
	m.buttonClickPtr = MakeEventDataPtr(fn)
	directoryEditImportAPI().SysCallN(28, m.Instance(), m.buttonClickPtr)
}

var (
	directoryEditImport       *imports.Imports = nil
	directoryEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DirectoryEdit_AutoSelect", 0),
		/*1*/ imports.NewTable("DirectoryEdit_AutoSelected", 0),
		/*2*/ imports.NewTable("DirectoryEdit_ButtonCaption", 0),
		/*3*/ imports.NewTable("DirectoryEdit_ButtonCursor", 0),
		/*4*/ imports.NewTable("DirectoryEdit_ButtonHint", 0),
		/*5*/ imports.NewTable("DirectoryEdit_ButtonOnlyWhenFocused", 0),
		/*6*/ imports.NewTable("DirectoryEdit_ButtonWidth", 0),
		/*7*/ imports.NewTable("DirectoryEdit_Class", 0),
		/*8*/ imports.NewTable("DirectoryEdit_Create", 0),
		/*9*/ imports.NewTable("DirectoryEdit_DialogOptions", 0),
		/*10*/ imports.NewTable("DirectoryEdit_DialogTitle", 0),
		/*11*/ imports.NewTable("DirectoryEdit_DirectInput", 0),
		/*12*/ imports.NewTable("DirectoryEdit_Directory", 0),
		/*13*/ imports.NewTable("DirectoryEdit_DragCursor", 0),
		/*14*/ imports.NewTable("DirectoryEdit_DragMode", 0),
		/*15*/ imports.NewTable("DirectoryEdit_Flat", 0),
		/*16*/ imports.NewTable("DirectoryEdit_FocusOnButtonClick", 0),
		/*17*/ imports.NewTable("DirectoryEdit_Glyph", 0),
		/*18*/ imports.NewTable("DirectoryEdit_ImageIndex", 0),
		/*19*/ imports.NewTable("DirectoryEdit_ImageWidth", 0),
		/*20*/ imports.NewTable("DirectoryEdit_Images", 0),
		/*21*/ imports.NewTable("DirectoryEdit_Layout", 0),
		/*22*/ imports.NewTable("DirectoryEdit_NumGlyphs", 0),
		/*23*/ imports.NewTable("DirectoryEdit_ParentFont", 0),
		/*24*/ imports.NewTable("DirectoryEdit_ParentShowHint", 0),
		/*25*/ imports.NewTable("DirectoryEdit_RootDir", 0),
		/*26*/ imports.NewTable("DirectoryEdit_RunDialog", 0),
		/*27*/ imports.NewTable("DirectoryEdit_SetOnAcceptDirectory", 0),
		/*28*/ imports.NewTable("DirectoryEdit_SetOnButtonClick", 0),
		/*29*/ imports.NewTable("DirectoryEdit_ShowHidden", 0),
		/*30*/ imports.NewTable("DirectoryEdit_Spacing", 0),
	}
)

func directoryEditImportAPI() *imports.Imports {
	if directoryEditImport == nil {
		directoryEditImport = NewDefaultImports()
		directoryEditImport.SetImportTable(directoryEditImportTables)
		directoryEditImportTables = nil
	}
	return directoryEditImport
}
