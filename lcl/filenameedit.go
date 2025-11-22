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

// IFileNameEdit Parent: ICustomEditButton
type IFileNameEdit interface {
	ICustomEditButton
	RunDialog()                 // procedure
	AutoSelected() bool         // property AutoSelected Getter
	SetAutoSelected(value bool) // property AutoSelected Setter
	DialogFiles() IStrings      // property DialogFiles Getter
	// FileName
	//  TFileName properties.
	FileName() string                            // property FileName Getter
	SetFileName(value string)                    // property FileName Setter
	InitialDir() string                          // property InitialDir Getter
	SetInitialDir(value string)                  // property InitialDir Setter
	DialogKind() types.TDialogKind               // property DialogKind Getter
	SetDialogKind(value types.TDialogKind)       // property DialogKind Setter
	DialogTitle() string                         // property DialogTitle Getter
	SetDialogTitle(value string)                 // property DialogTitle Setter
	DialogOptions() types.TOpenOptions           // property DialogOptions Getter
	SetDialogOptions(value types.TOpenOptions)   // property DialogOptions Setter
	Filter() string                              // property Filter Getter
	SetFilter(value string)                      // property Filter Setter
	FilterIndex() int32                          // property FilterIndex Getter
	SetFilterIndex(value int32)                  // property FilterIndex Setter
	DefaultExt() string                          // property DefaultExt Getter
	SetDefaultExt(value string)                  // property DefaultExt Setter
	HideDirectories() bool                       // property HideDirectories Getter
	SetHideDirectories(value bool)               // property HideDirectories Setter
	ButtonCaption() string                       // property ButtonCaption Getter
	SetButtonCaption(value string)               // property ButtonCaption Setter
	ButtonCursor() types.TCursor                 // property ButtonCursor Getter
	SetButtonCursor(value types.TCursor)         // property ButtonCursor Setter
	ButtonHint() string                          // property ButtonHint Getter
	SetButtonHint(value string)                  // property ButtonHint Setter
	ButtonOnlyWhenFocused() bool                 // property ButtonOnlyWhenFocused Getter
	SetButtonOnlyWhenFocused(value bool)         // property ButtonOnlyWhenFocused Setter
	ButtonWidth() int32                          // property ButtonWidth Getter
	SetButtonWidth(value int32)                  // property ButtonWidth Setter
	DirectInput() bool                           // property DirectInput Getter
	SetDirectInput(value bool)                   // property DirectInput Setter
	Glyph() IBitmap                              // property Glyph Getter
	SetGlyph(value IBitmap)                      // property Glyph Setter
	NumGlyphs() int32                            // property NumGlyphs Getter
	SetNumGlyphs(value int32)                    // property NumGlyphs Setter
	Images() ICustomImageList                    // property Images Getter
	SetImages(value ICustomImageList)            // property Images Setter
	ImageIndex() int32                           // property ImageIndex Getter
	SetImageIndex(value int32)                   // property ImageIndex Setter
	ImageWidth() int32                           // property ImageWidth Getter
	SetImageWidth(value int32)                   // property ImageWidth Setter
	Flat() bool                                  // property Flat Getter
	SetFlat(value bool)                          // property Flat Setter
	FocusOnButtonClick() bool                    // property FocusOnButtonClick Getter
	SetFocusOnButtonClick(value bool)            // property FocusOnButtonClick Setter
	AutoSelect() bool                            // property AutoSelect Getter
	SetAutoSelect(value bool)                    // property AutoSelect Setter
	DragCursor() types.TCursor                   // property DragCursor Getter
	SetDragCursor(value types.TCursor)           // property DragCursor Setter
	DragMode() types.TDragMode                   // property DragMode Getter
	SetDragMode(value types.TDragMode)           // property DragMode Setter
	Layout() types.TLeftRight                    // property Layout Getter
	SetLayout(value types.TLeftRight)            // property Layout Setter
	ParentFont() bool                            // property ParentFont Getter
	SetParentFont(value bool)                    // property ParentFont Setter
	ParentShowHint() bool                        // property ParentShowHint Getter
	SetParentShowHint(value bool)                // property ParentShowHint Setter
	Spacing() int32                              // property Spacing Getter
	SetSpacing(value int32)                      // property Spacing Setter
	SetOnAcceptFileName(fn TAcceptFileNameEvent) // property event
	SetOnFolderChange(fn TNotifyEvent)           // property event
	SetOnButtonClick(fn TNotifyEvent)            // property event
}

type TFileNameEdit struct {
	TCustomEditButton
}

func (m *TFileNameEdit) RunDialog() {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(1, m.Instance())
}

func (m *TFileNameEdit) AutoSelected() bool {
	if !m.IsValid() {
		return false
	}
	r := fileNameEditAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFileNameEdit) SetAutoSelected(value bool) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TFileNameEdit) DialogFiles() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := fileNameEditAPI().SysCallN(3, m.Instance())
	return AsStrings(r)
}

func (m *TFileNameEdit) FileName() string {
	if !m.IsValid() {
		return ""
	}
	r := fileNameEditAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileNameEdit) SetFileName(value string) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileNameEdit) InitialDir() string {
	if !m.IsValid() {
		return ""
	}
	r := fileNameEditAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileNameEdit) SetInitialDir(value string) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileNameEdit) DialogKind() types.TDialogKind {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(6, 0, m.Instance())
	return types.TDialogKind(r)
}

func (m *TFileNameEdit) SetDialogKind(value types.TDialogKind) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) DialogTitle() string {
	if !m.IsValid() {
		return ""
	}
	r := fileNameEditAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileNameEdit) SetDialogTitle(value string) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileNameEdit) DialogOptions() types.TOpenOptions {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(8, 0, m.Instance())
	return types.TOpenOptions(r)
}

func (m *TFileNameEdit) SetDialogOptions(value types.TOpenOptions) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) Filter() string {
	if !m.IsValid() {
		return ""
	}
	r := fileNameEditAPI().SysCallN(9, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileNameEdit) SetFilter(value string) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileNameEdit) FilterIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TFileNameEdit) SetFilterIndex(value int32) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) DefaultExt() string {
	if !m.IsValid() {
		return ""
	}
	r := fileNameEditAPI().SysCallN(11, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileNameEdit) SetDefaultExt(value string) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(11, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileNameEdit) HideDirectories() bool {
	if !m.IsValid() {
		return false
	}
	r := fileNameEditAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFileNameEdit) SetHideDirectories(value bool) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TFileNameEdit) ButtonCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := fileNameEditAPI().SysCallN(13, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileNameEdit) SetButtonCaption(value string) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(13, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileNameEdit) ButtonCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(14, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TFileNameEdit) SetButtonCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) ButtonHint() string {
	if !m.IsValid() {
		return ""
	}
	r := fileNameEditAPI().SysCallN(15, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TFileNameEdit) SetButtonHint(value string) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(15, 1, m.Instance(), api.PasStr(value))
}

func (m *TFileNameEdit) ButtonOnlyWhenFocused() bool {
	if !m.IsValid() {
		return false
	}
	r := fileNameEditAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFileNameEdit) SetButtonOnlyWhenFocused(value bool) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TFileNameEdit) ButtonWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TFileNameEdit) SetButtonWidth(value int32) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) DirectInput() bool {
	if !m.IsValid() {
		return false
	}
	r := fileNameEditAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFileNameEdit) SetDirectInput(value bool) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TFileNameEdit) Glyph() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := fileNameEditAPI().SysCallN(19, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TFileNameEdit) SetGlyph(value IBitmap) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(19, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFileNameEdit) NumGlyphs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TFileNameEdit) SetNumGlyphs(value int32) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := fileNameEditAPI().SysCallN(21, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TFileNameEdit) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(21, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TFileNameEdit) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TFileNameEdit) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) ImageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TFileNameEdit) SetImageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := fileNameEditAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFileNameEdit) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TFileNameEdit) FocusOnButtonClick() bool {
	if !m.IsValid() {
		return false
	}
	r := fileNameEditAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFileNameEdit) SetFocusOnButtonClick(value bool) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TFileNameEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := fileNameEditAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFileNameEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TFileNameEdit) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(27, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TFileNameEdit) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(28, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TFileNameEdit) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) Layout() types.TLeftRight {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(29, 0, m.Instance())
	return types.TLeftRight(r)
}

func (m *TFileNameEdit) SetLayout(value types.TLeftRight) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := fileNameEditAPI().SysCallN(30, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFileNameEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(30, 1, m.Instance(), api.PasBool(value))
}

func (m *TFileNameEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := fileNameEditAPI().SysCallN(31, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TFileNameEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(31, 1, m.Instance(), api.PasBool(value))
}

func (m *TFileNameEdit) Spacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := fileNameEditAPI().SysCallN(32, 0, m.Instance())
	return int32(r)
}

func (m *TFileNameEdit) SetSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	fileNameEditAPI().SysCallN(32, 1, m.Instance(), uintptr(value))
}

func (m *TFileNameEdit) SetOnAcceptFileName(fn TAcceptFileNameEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTAcceptFileNameEvent(fn)
	base.SetEvent(m, 33, fileNameEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFileNameEdit) SetOnFolderChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 34, fileNameEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TFileNameEdit) SetOnButtonClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 35, fileNameEditAPI(), api.MakeEventDataPtr(cb))
}

// NewFileNameEdit class constructor
func NewFileNameEdit(owner IComponent) IFileNameEdit {
	r := fileNameEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsFileNameEdit(r)
}

func TFileNameEditClass() types.TClass {
	r := fileNameEditAPI().SysCallN(36)
	return types.TClass(r)
}

var (
	fileNameEditOnce   base.Once
	fileNameEditImport *imports.Imports = nil
)

func fileNameEditAPI() *imports.Imports {
	fileNameEditOnce.Do(func() {
		fileNameEditImport = api.NewDefaultImports()
		fileNameEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TFileNameEdit_Create", 0), // constructor NewFileNameEdit
			/* 1 */ imports.NewTable("TFileNameEdit_RunDialog", 0), // procedure RunDialog
			/* 2 */ imports.NewTable("TFileNameEdit_AutoSelected", 0), // property AutoSelected
			/* 3 */ imports.NewTable("TFileNameEdit_DialogFiles", 0), // property DialogFiles
			/* 4 */ imports.NewTable("TFileNameEdit_FileName", 0), // property FileName
			/* 5 */ imports.NewTable("TFileNameEdit_InitialDir", 0), // property InitialDir
			/* 6 */ imports.NewTable("TFileNameEdit_DialogKind", 0), // property DialogKind
			/* 7 */ imports.NewTable("TFileNameEdit_DialogTitle", 0), // property DialogTitle
			/* 8 */ imports.NewTable("TFileNameEdit_DialogOptions", 0), // property DialogOptions
			/* 9 */ imports.NewTable("TFileNameEdit_Filter", 0), // property Filter
			/* 10 */ imports.NewTable("TFileNameEdit_FilterIndex", 0), // property FilterIndex
			/* 11 */ imports.NewTable("TFileNameEdit_DefaultExt", 0), // property DefaultExt
			/* 12 */ imports.NewTable("TFileNameEdit_HideDirectories", 0), // property HideDirectories
			/* 13 */ imports.NewTable("TFileNameEdit_ButtonCaption", 0), // property ButtonCaption
			/* 14 */ imports.NewTable("TFileNameEdit_ButtonCursor", 0), // property ButtonCursor
			/* 15 */ imports.NewTable("TFileNameEdit_ButtonHint", 0), // property ButtonHint
			/* 16 */ imports.NewTable("TFileNameEdit_ButtonOnlyWhenFocused", 0), // property ButtonOnlyWhenFocused
			/* 17 */ imports.NewTable("TFileNameEdit_ButtonWidth", 0), // property ButtonWidth
			/* 18 */ imports.NewTable("TFileNameEdit_DirectInput", 0), // property DirectInput
			/* 19 */ imports.NewTable("TFileNameEdit_Glyph", 0), // property Glyph
			/* 20 */ imports.NewTable("TFileNameEdit_NumGlyphs", 0), // property NumGlyphs
			/* 21 */ imports.NewTable("TFileNameEdit_Images", 0), // property Images
			/* 22 */ imports.NewTable("TFileNameEdit_ImageIndex", 0), // property ImageIndex
			/* 23 */ imports.NewTable("TFileNameEdit_ImageWidth", 0), // property ImageWidth
			/* 24 */ imports.NewTable("TFileNameEdit_Flat", 0), // property Flat
			/* 25 */ imports.NewTable("TFileNameEdit_FocusOnButtonClick", 0), // property FocusOnButtonClick
			/* 26 */ imports.NewTable("TFileNameEdit_AutoSelect", 0), // property AutoSelect
			/* 27 */ imports.NewTable("TFileNameEdit_DragCursor", 0), // property DragCursor
			/* 28 */ imports.NewTable("TFileNameEdit_DragMode", 0), // property DragMode
			/* 29 */ imports.NewTable("TFileNameEdit_Layout", 0), // property Layout
			/* 30 */ imports.NewTable("TFileNameEdit_ParentFont", 0), // property ParentFont
			/* 31 */ imports.NewTable("TFileNameEdit_ParentShowHint", 0), // property ParentShowHint
			/* 32 */ imports.NewTable("TFileNameEdit_Spacing", 0), // property Spacing
			/* 33 */ imports.NewTable("TFileNameEdit_OnAcceptFileName", 0), // event OnAcceptFileName
			/* 34 */ imports.NewTable("TFileNameEdit_OnFolderChange", 0), // event OnFolderChange
			/* 35 */ imports.NewTable("TFileNameEdit_OnButtonClick", 0), // event OnButtonClick
			/* 36 */ imports.NewTable("TFileNameEdit_TClass", 0), // function TFileNameEditClass
		}
	})
	return fileNameEditImport
}
