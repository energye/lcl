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

// IDirectoryEdit Parent: ICustomEditButton
type IDirectoryEdit interface {
	ICustomEditButton
	RunDialog()                 // procedure
	AutoSelected() bool         // property AutoSelected Getter
	SetAutoSelected(value bool) // property AutoSelected Setter
	// Directory
	//  TDirectory properties.
	Directory() string                            // property Directory Getter
	SetDirectory(value string)                    // property Directory Setter
	RootDir() string                              // property RootDir Getter
	SetRootDir(value string)                      // property RootDir Setter
	DialogTitle() string                          // property DialogTitle Getter
	SetDialogTitle(value string)                  // property DialogTitle Setter
	DialogOptions() types.TOpenOptions            // property DialogOptions Getter
	SetDialogOptions(value types.TOpenOptions)    // property DialogOptions Setter
	ShowHidden() bool                             // property ShowHidden Getter
	SetShowHidden(value bool)                     // property ShowHidden Setter
	ButtonCaption() string                        // property ButtonCaption Getter
	SetButtonCaption(value string)                // property ButtonCaption Setter
	ButtonCursor() types.TCursor                  // property ButtonCursor Getter
	SetButtonCursor(value types.TCursor)          // property ButtonCursor Setter
	ButtonHint() string                           // property ButtonHint Getter
	SetButtonHint(value string)                   // property ButtonHint Setter
	ButtonOnlyWhenFocused() bool                  // property ButtonOnlyWhenFocused Getter
	SetButtonOnlyWhenFocused(value bool)          // property ButtonOnlyWhenFocused Setter
	ButtonWidth() int32                           // property ButtonWidth Getter
	SetButtonWidth(value int32)                   // property ButtonWidth Setter
	DirectInput() bool                            // property DirectInput Getter
	SetDirectInput(value bool)                    // property DirectInput Setter
	Glyph() IBitmap                               // property Glyph Getter
	SetGlyph(value IBitmap)                       // property Glyph Setter
	NumGlyphs() int32                             // property NumGlyphs Getter
	SetNumGlyphs(value int32)                     // property NumGlyphs Setter
	Images() ICustomImageList                     // property Images Getter
	SetImages(value ICustomImageList)             // property Images Setter
	ImageIndex() int32                            // property ImageIndex Getter
	SetImageIndex(value int32)                    // property ImageIndex Setter
	ImageWidth() int32                            // property ImageWidth Getter
	SetImageWidth(value int32)                    // property ImageWidth Setter
	Flat() bool                                   // property Flat Getter
	SetFlat(value bool)                           // property Flat Setter
	FocusOnButtonClick() bool                     // property FocusOnButtonClick Getter
	SetFocusOnButtonClick(value bool)             // property FocusOnButtonClick Setter
	AutoSelect() bool                             // property AutoSelect Getter
	SetAutoSelect(value bool)                     // property AutoSelect Setter
	DragCursor() types.TCursor                    // property DragCursor Getter
	SetDragCursor(value types.TCursor)            // property DragCursor Setter
	DragMode() types.TDragMode                    // property DragMode Getter
	SetDragMode(value types.TDragMode)            // property DragMode Setter
	Layout() types.TLeftRight                     // property Layout Getter
	SetLayout(value types.TLeftRight)             // property Layout Setter
	ParentFont() bool                             // property ParentFont Getter
	SetParentFont(value bool)                     // property ParentFont Setter
	ParentShowHint() bool                         // property ParentShowHint Getter
	SetParentShowHint(value bool)                 // property ParentShowHint Setter
	Spacing() int32                               // property Spacing Getter
	SetSpacing(value int32)                       // property Spacing Setter
	SetOnAcceptDirectory(fn TAcceptFileNameEvent) // property event
	SetOnButtonClick(fn TNotifyEvent)             // property event
}

type TDirectoryEdit struct {
	TCustomEditButton
}

func (m *TDirectoryEdit) RunDialog() {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(1, m.Instance())
}

func (m *TDirectoryEdit) AutoSelected() bool {
	if !m.IsValid() {
		return false
	}
	r := directoryEditAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDirectoryEdit) SetAutoSelected(value bool) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TDirectoryEdit) Directory() string {
	if !m.IsValid() {
		return ""
	}
	r := directoryEditAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDirectoryEdit) SetDirectory(value string) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TDirectoryEdit) RootDir() string {
	if !m.IsValid() {
		return ""
	}
	r := directoryEditAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDirectoryEdit) SetRootDir(value string) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TDirectoryEdit) DialogTitle() string {
	if !m.IsValid() {
		return ""
	}
	r := directoryEditAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDirectoryEdit) SetDialogTitle(value string) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TDirectoryEdit) DialogOptions() types.TOpenOptions {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(6, 0, m.Instance())
	return types.TOpenOptions(r)
}

func (m *TDirectoryEdit) SetDialogOptions(value types.TOpenOptions) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) ShowHidden() bool {
	if !m.IsValid() {
		return false
	}
	r := directoryEditAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDirectoryEdit) SetShowHidden(value bool) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TDirectoryEdit) ButtonCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := directoryEditAPI().SysCallN(8, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDirectoryEdit) SetButtonCaption(value string) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(8, 1, m.Instance(), api.PasStr(value))
}

func (m *TDirectoryEdit) ButtonCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(9, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TDirectoryEdit) SetButtonCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) ButtonHint() string {
	if !m.IsValid() {
		return ""
	}
	r := directoryEditAPI().SysCallN(10, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDirectoryEdit) SetButtonHint(value string) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(10, 1, m.Instance(), api.PasStr(value))
}

func (m *TDirectoryEdit) ButtonOnlyWhenFocused() bool {
	if !m.IsValid() {
		return false
	}
	r := directoryEditAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDirectoryEdit) SetButtonOnlyWhenFocused(value bool) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TDirectoryEdit) ButtonWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(12, 0, m.Instance())
	return int32(r)
}

func (m *TDirectoryEdit) SetButtonWidth(value int32) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) DirectInput() bool {
	if !m.IsValid() {
		return false
	}
	r := directoryEditAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDirectoryEdit) SetDirectInput(value bool) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TDirectoryEdit) Glyph() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := directoryEditAPI().SysCallN(14, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TDirectoryEdit) SetGlyph(value IBitmap) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDirectoryEdit) NumGlyphs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TDirectoryEdit) SetNumGlyphs(value int32) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := directoryEditAPI().SysCallN(16, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TDirectoryEdit) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(16, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDirectoryEdit) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TDirectoryEdit) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) ImageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TDirectoryEdit) SetImageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := directoryEditAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDirectoryEdit) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TDirectoryEdit) FocusOnButtonClick() bool {
	if !m.IsValid() {
		return false
	}
	r := directoryEditAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDirectoryEdit) SetFocusOnButtonClick(value bool) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TDirectoryEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := directoryEditAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDirectoryEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TDirectoryEdit) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(22, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TDirectoryEdit) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(23, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TDirectoryEdit) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) Layout() types.TLeftRight {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(24, 0, m.Instance())
	return types.TLeftRight(r)
}

func (m *TDirectoryEdit) SetLayout(value types.TLeftRight) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := directoryEditAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDirectoryEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TDirectoryEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := directoryEditAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDirectoryEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TDirectoryEdit) Spacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := directoryEditAPI().SysCallN(27, 0, m.Instance())
	return int32(r)
}

func (m *TDirectoryEdit) SetSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	directoryEditAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TDirectoryEdit) SetOnAcceptDirectory(fn TAcceptFileNameEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTAcceptFileNameEvent(fn)
	base.SetEvent(m, 28, directoryEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDirectoryEdit) SetOnButtonClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 29, directoryEditAPI(), api.MakeEventDataPtr(cb))
}

// NewDirectoryEdit class constructor
func NewDirectoryEdit(owner IComponent) IDirectoryEdit {
	r := directoryEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsDirectoryEdit(r)
}

func TDirectoryEditClass() types.TClass {
	r := directoryEditAPI().SysCallN(30)
	return types.TClass(r)
}

var (
	directoryEditOnce   base.Once
	directoryEditImport *imports.Imports = nil
)

func directoryEditAPI() *imports.Imports {
	directoryEditOnce.Do(func() {
		directoryEditImport = api.NewDefaultImports()
		directoryEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDirectoryEdit_Create", 0), // constructor NewDirectoryEdit
			/* 1 */ imports.NewTable("TDirectoryEdit_RunDialog", 0), // procedure RunDialog
			/* 2 */ imports.NewTable("TDirectoryEdit_AutoSelected", 0), // property AutoSelected
			/* 3 */ imports.NewTable("TDirectoryEdit_Directory", 0), // property Directory
			/* 4 */ imports.NewTable("TDirectoryEdit_RootDir", 0), // property RootDir
			/* 5 */ imports.NewTable("TDirectoryEdit_DialogTitle", 0), // property DialogTitle
			/* 6 */ imports.NewTable("TDirectoryEdit_DialogOptions", 0), // property DialogOptions
			/* 7 */ imports.NewTable("TDirectoryEdit_ShowHidden", 0), // property ShowHidden
			/* 8 */ imports.NewTable("TDirectoryEdit_ButtonCaption", 0), // property ButtonCaption
			/* 9 */ imports.NewTable("TDirectoryEdit_ButtonCursor", 0), // property ButtonCursor
			/* 10 */ imports.NewTable("TDirectoryEdit_ButtonHint", 0), // property ButtonHint
			/* 11 */ imports.NewTable("TDirectoryEdit_ButtonOnlyWhenFocused", 0), // property ButtonOnlyWhenFocused
			/* 12 */ imports.NewTable("TDirectoryEdit_ButtonWidth", 0), // property ButtonWidth
			/* 13 */ imports.NewTable("TDirectoryEdit_DirectInput", 0), // property DirectInput
			/* 14 */ imports.NewTable("TDirectoryEdit_Glyph", 0), // property Glyph
			/* 15 */ imports.NewTable("TDirectoryEdit_NumGlyphs", 0), // property NumGlyphs
			/* 16 */ imports.NewTable("TDirectoryEdit_Images", 0), // property Images
			/* 17 */ imports.NewTable("TDirectoryEdit_ImageIndex", 0), // property ImageIndex
			/* 18 */ imports.NewTable("TDirectoryEdit_ImageWidth", 0), // property ImageWidth
			/* 19 */ imports.NewTable("TDirectoryEdit_Flat", 0), // property Flat
			/* 20 */ imports.NewTable("TDirectoryEdit_FocusOnButtonClick", 0), // property FocusOnButtonClick
			/* 21 */ imports.NewTable("TDirectoryEdit_AutoSelect", 0), // property AutoSelect
			/* 22 */ imports.NewTable("TDirectoryEdit_DragCursor", 0), // property DragCursor
			/* 23 */ imports.NewTable("TDirectoryEdit_DragMode", 0), // property DragMode
			/* 24 */ imports.NewTable("TDirectoryEdit_Layout", 0), // property Layout
			/* 25 */ imports.NewTable("TDirectoryEdit_ParentFont", 0), // property ParentFont
			/* 26 */ imports.NewTable("TDirectoryEdit_ParentShowHint", 0), // property ParentShowHint
			/* 27 */ imports.NewTable("TDirectoryEdit_Spacing", 0), // property Spacing
			/* 28 */ imports.NewTable("TDirectoryEdit_OnAcceptDirectory", 0), // event OnAcceptDirectory
			/* 29 */ imports.NewTable("TDirectoryEdit_OnButtonClick", 0), // event OnButtonClick
			/* 30 */ imports.NewTable("TDirectoryEdit_TClass", 0), // function TDirectoryEditClass
		}
	})
	return directoryEditImport
}
