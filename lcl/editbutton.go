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

// IEditButton Parent: ICustomEditButton
type IEditButton interface {
	ICustomEditButton
	AutoSelected() bool                  // property AutoSelected Getter
	SetAutoSelected(value bool)          // property AutoSelected Setter
	Button() ISpeedButton                // property Button Getter
	Edit() IEbEdit                       // property Edit Getter
	AutoSelect() bool                    // property AutoSelect Getter
	SetAutoSelect(value bool)            // property AutoSelect Setter
	ButtonCaption() string               // property ButtonCaption Getter
	SetButtonCaption(value string)       // property ButtonCaption Setter
	ButtonCursor() types.TCursor         // property ButtonCursor Getter
	SetButtonCursor(value types.TCursor) // property ButtonCursor Setter
	ButtonHint() string                  // property ButtonHint Getter
	SetButtonHint(value string)          // property ButtonHint Setter
	ButtonOnlyWhenFocused() bool         // property ButtonOnlyWhenFocused Getter
	SetButtonOnlyWhenFocused(value bool) // property ButtonOnlyWhenFocused Setter
	ButtonWidth() int32                  // property ButtonWidth Getter
	SetButtonWidth(value int32)          // property ButtonWidth Setter
	DirectInput() bool                   // property DirectInput Getter
	SetDirectInput(value bool)           // property DirectInput Setter
	Flat() bool                          // property Flat Getter
	SetFlat(value bool)                  // property Flat Setter
	FocusOnButtonClick() bool            // property FocusOnButtonClick Getter
	SetFocusOnButtonClick(value bool)    // property FocusOnButtonClick Setter
	Glyph() IBitmap                      // property Glyph Getter
	SetGlyph(value IBitmap)              // property Glyph Setter
	Images() ICustomImageList            // property Images Getter
	SetImages(value ICustomImageList)    // property Images Setter
	ImageIndex() int32                   // property ImageIndex Getter
	SetImageIndex(value int32)           // property ImageIndex Setter
	ImageWidth() int32                   // property ImageWidth Getter
	SetImageWidth(value int32)           // property ImageWidth Setter
	Layout() types.TLeftRight            // property Layout Getter
	SetLayout(value types.TLeftRight)    // property Layout Setter
	NumGlyphs() int32                    // property NumGlyphs Getter
	SetNumGlyphs(value int32)            // property NumGlyphs Setter
	ParentFont() bool                    // property ParentFont Getter
	SetParentFont(value bool)            // property ParentFont Setter
	ParentShowHint() bool                // property ParentShowHint Getter
	SetParentShowHint(value bool)        // property ParentShowHint Setter
	Spacing() int32                      // property Spacing Getter
	SetSpacing(value int32)              // property Spacing Setter
	SetOnButtonClick(fn TNotifyEvent)    // property event
}

type TEditButton struct {
	TCustomEditButton
}

func (m *TEditButton) AutoSelected() bool {
	if !m.IsValid() {
		return false
	}
	r := editButtonAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEditButton) SetAutoSelected(value bool) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TEditButton) Button() ISpeedButton {
	if !m.IsValid() {
		return nil
	}
	r := editButtonAPI().SysCallN(2, m.Instance())
	return AsSpeedButton(r)
}

func (m *TEditButton) Edit() IEbEdit {
	if !m.IsValid() {
		return nil
	}
	r := editButtonAPI().SysCallN(3, m.Instance())
	return AsEbEdit(r)
}

func (m *TEditButton) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := editButtonAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEditButton) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TEditButton) ButtonCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := editButtonAPI().SysCallN(5, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TEditButton) SetButtonCaption(value string) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(5, 1, m.Instance(), api.PasStr(value))
}

func (m *TEditButton) ButtonCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := editButtonAPI().SysCallN(6, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TEditButton) SetButtonCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TEditButton) ButtonHint() string {
	if !m.IsValid() {
		return ""
	}
	r := editButtonAPI().SysCallN(7, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TEditButton) SetButtonHint(value string) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(7, 1, m.Instance(), api.PasStr(value))
}

func (m *TEditButton) ButtonOnlyWhenFocused() bool {
	if !m.IsValid() {
		return false
	}
	r := editButtonAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEditButton) SetButtonOnlyWhenFocused(value bool) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TEditButton) ButtonWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := editButtonAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TEditButton) SetButtonWidth(value int32) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TEditButton) DirectInput() bool {
	if !m.IsValid() {
		return false
	}
	r := editButtonAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEditButton) SetDirectInput(value bool) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TEditButton) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := editButtonAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEditButton) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TEditButton) FocusOnButtonClick() bool {
	if !m.IsValid() {
		return false
	}
	r := editButtonAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEditButton) SetFocusOnButtonClick(value bool) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TEditButton) Glyph() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := editButtonAPI().SysCallN(13, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TEditButton) SetGlyph(value IBitmap) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TEditButton) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := editButtonAPI().SysCallN(14, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TEditButton) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TEditButton) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := editButtonAPI().SysCallN(15, 0, m.Instance())
	return int32(r)
}

func (m *TEditButton) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TEditButton) ImageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := editButtonAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TEditButton) SetImageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TEditButton) Layout() types.TLeftRight {
	if !m.IsValid() {
		return 0
	}
	r := editButtonAPI().SysCallN(17, 0, m.Instance())
	return types.TLeftRight(r)
}

func (m *TEditButton) SetLayout(value types.TLeftRight) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TEditButton) NumGlyphs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := editButtonAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TEditButton) SetNumGlyphs(value int32) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TEditButton) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := editButtonAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEditButton) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TEditButton) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := editButtonAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TEditButton) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TEditButton) Spacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := editButtonAPI().SysCallN(21, 0, m.Instance())
	return int32(r)
}

func (m *TEditButton) SetSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	editButtonAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TEditButton) SetOnButtonClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 22, editButtonAPI(), api.MakeEventDataPtr(cb))
}

// NewEditButton class constructor
func NewEditButton(owner IComponent) IEditButton {
	r := editButtonAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsEditButton(r)
}

func TEditButtonClass() types.TClass {
	r := editButtonAPI().SysCallN(23)
	return types.TClass(r)
}

var (
	editButtonOnce   base.Once
	editButtonImport *imports.Imports = nil
)

func editButtonAPI() *imports.Imports {
	editButtonOnce.Do(func() {
		editButtonImport = api.NewDefaultImports()
		editButtonImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TEditButton_Create", 0), // constructor NewEditButton
			/* 1 */ imports.NewTable("TEditButton_AutoSelected", 0), // property AutoSelected
			/* 2 */ imports.NewTable("TEditButton_Button", 0), // property Button
			/* 3 */ imports.NewTable("TEditButton_Edit", 0), // property Edit
			/* 4 */ imports.NewTable("TEditButton_AutoSelect", 0), // property AutoSelect
			/* 5 */ imports.NewTable("TEditButton_ButtonCaption", 0), // property ButtonCaption
			/* 6 */ imports.NewTable("TEditButton_ButtonCursor", 0), // property ButtonCursor
			/* 7 */ imports.NewTable("TEditButton_ButtonHint", 0), // property ButtonHint
			/* 8 */ imports.NewTable("TEditButton_ButtonOnlyWhenFocused", 0), // property ButtonOnlyWhenFocused
			/* 9 */ imports.NewTable("TEditButton_ButtonWidth", 0), // property ButtonWidth
			/* 10 */ imports.NewTable("TEditButton_DirectInput", 0), // property DirectInput
			/* 11 */ imports.NewTable("TEditButton_Flat", 0), // property Flat
			/* 12 */ imports.NewTable("TEditButton_FocusOnButtonClick", 0), // property FocusOnButtonClick
			/* 13 */ imports.NewTable("TEditButton_Glyph", 0), // property Glyph
			/* 14 */ imports.NewTable("TEditButton_Images", 0), // property Images
			/* 15 */ imports.NewTable("TEditButton_ImageIndex", 0), // property ImageIndex
			/* 16 */ imports.NewTable("TEditButton_ImageWidth", 0), // property ImageWidth
			/* 17 */ imports.NewTable("TEditButton_Layout", 0), // property Layout
			/* 18 */ imports.NewTable("TEditButton_NumGlyphs", 0), // property NumGlyphs
			/* 19 */ imports.NewTable("TEditButton_ParentFont", 0), // property ParentFont
			/* 20 */ imports.NewTable("TEditButton_ParentShowHint", 0), // property ParentShowHint
			/* 21 */ imports.NewTable("TEditButton_Spacing", 0), // property Spacing
			/* 22 */ imports.NewTable("TEditButton_OnButtonClick", 0), // event OnButtonClick
			/* 23 */ imports.NewTable("TEditButton_TClass", 0), // function TEditButtonClass
		}
	})
	return editButtonImport
}
