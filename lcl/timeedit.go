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

// ITimeEdit Parent: ICustomEditButton
type ITimeEdit interface {
	ICustomEditButton
	ValidTimeFormat(format string) bool  // function
	Time() types.TDateTime               // property Time Getter
	SetTime(value types.TDateTime)       // property Time Setter
	Button() ISpeedButton                // property Button Getter
	DroppedDown() bool                   // property DroppedDown Getter
	DefaultNow() bool                    // property DefaultNow Getter
	SetDefaultNow(value bool)            // property DefaultNow Setter
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
	AutoSelect() bool                    // property AutoSelect Getter
	SetAutoSelect(value bool)            // property AutoSelect Setter
	DirectInput() bool                   // property DirectInput Getter
	SetDirectInput(value bool)           // property DirectInput Setter
	Glyph() IBitmap                      // property Glyph Getter
	SetGlyph(value IBitmap)              // property Glyph Setter
	NumGlyphs() int32                    // property NumGlyphs Getter
	SetNumGlyphs(value int32)            // property NumGlyphs Setter
	Images() ICustomImageList            // property Images Getter
	SetImages(value ICustomImageList)    // property Images Setter
	ImageIndex() int32                   // property ImageIndex Getter
	SetImageIndex(value int32)           // property ImageIndex Setter
	ImageWidth() int32                   // property ImageWidth Getter
	SetImageWidth(value int32)           // property ImageWidth Setter
	DragMode() types.TDragMode           // property DragMode Getter
	SetDragMode(value types.TDragMode)   // property DragMode Setter
	Flat() bool                          // property Flat Getter
	SetFlat(value bool)                  // property Flat Setter
	FocusOnButtonClick() bool            // property FocusOnButtonClick Getter
	SetFocusOnButtonClick(value bool)    // property FocusOnButtonClick Setter
	ParentFont() bool                    // property ParentFont Getter
	SetParentFont(value bool)            // property ParentFont Setter
	ParentShowHint() bool                // property ParentShowHint Getter
	SetParentShowHint(value bool)        // property ParentShowHint Setter
	SimpleLayout() bool                  // property SimpleLayout Getter
	SetSimpleLayout(value bool)          // property SimpleLayout Setter
	Spacing() int32                      // property Spacing Getter
	SetSpacing(value int32)              // property Spacing Setter
	TimeAMString() string                // property TimeAMString Getter
	SetTimeAMString(value string)        // property TimeAMString Setter
	TimeFormat() string                  // property TimeFormat Getter
	SetTimeFormat(value string)          // property TimeFormat Setter
	TimeSeparator() string               // property TimeSeparator Getter
	SetTimeSeparator(value string)       // property TimeSeparator Setter
	TimePMString() string                // property TimePMString Getter
	SetTimePMString(value string)        // property TimePMString Setter
	SetOnAcceptTime(fn TAcceptTimeEvent) // property event
	SetOnCustomTime(fn TCustomTimeEvent) // property event
	SetOnButtonClick(fn TNotifyEvent)    // property event
}

type TTimeEdit struct {
	TCustomEditButton
}

func (m *TTimeEdit) ValidTimeFormat(format string) bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(1, m.Instance(), api.PasStr(format))
	return api.GoBool(r)
}

func (m *TTimeEdit) Time() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(2, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TTimeEdit) SetTime(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(2, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TTimeEdit) Button() ISpeedButton {
	if !m.IsValid() {
		return nil
	}
	r := timeEditAPI().SysCallN(3, m.Instance())
	return AsSpeedButton(r)
}

func (m *TTimeEdit) DroppedDown() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(4, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) DefaultNow() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) SetDefaultNow(value bool) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TTimeEdit) ButtonCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := timeEditAPI().SysCallN(6, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTimeEdit) SetButtonCaption(value string) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(6, 1, m.Instance(), api.PasStr(value))
}

func (m *TTimeEdit) ButtonCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := timeEditAPI().SysCallN(7, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TTimeEdit) SetButtonCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TTimeEdit) ButtonHint() string {
	if !m.IsValid() {
		return ""
	}
	r := timeEditAPI().SysCallN(8, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTimeEdit) SetButtonHint(value string) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(8, 1, m.Instance(), api.PasStr(value))
}

func (m *TTimeEdit) ButtonOnlyWhenFocused() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(9, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) SetButtonOnlyWhenFocused(value bool) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(9, 1, m.Instance(), api.PasBool(value))
}

func (m *TTimeEdit) ButtonWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := timeEditAPI().SysCallN(10, 0, m.Instance())
	return int32(r)
}

func (m *TTimeEdit) SetButtonWidth(value int32) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TTimeEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TTimeEdit) DirectInput() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) SetDirectInput(value bool) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TTimeEdit) Glyph() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := timeEditAPI().SysCallN(13, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TTimeEdit) SetGlyph(value IBitmap) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TTimeEdit) NumGlyphs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := timeEditAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TTimeEdit) SetNumGlyphs(value int32) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TTimeEdit) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := timeEditAPI().SysCallN(15, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TTimeEdit) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(15, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TTimeEdit) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := timeEditAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TTimeEdit) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TTimeEdit) ImageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := timeEditAPI().SysCallN(17, 0, m.Instance())
	return int32(r)
}

func (m *TTimeEdit) SetImageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TTimeEdit) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := timeEditAPI().SysCallN(18, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TTimeEdit) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TTimeEdit) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(19, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(19, 1, m.Instance(), api.PasBool(value))
}

func (m *TTimeEdit) FocusOnButtonClick() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) SetFocusOnButtonClick(value bool) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TTimeEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TTimeEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(22, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(22, 1, m.Instance(), api.PasBool(value))
}

func (m *TTimeEdit) SimpleLayout() bool {
	if !m.IsValid() {
		return false
	}
	r := timeEditAPI().SysCallN(23, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TTimeEdit) SetSimpleLayout(value bool) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(23, 1, m.Instance(), api.PasBool(value))
}

func (m *TTimeEdit) Spacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := timeEditAPI().SysCallN(24, 0, m.Instance())
	return int32(r)
}

func (m *TTimeEdit) SetSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TTimeEdit) TimeAMString() string {
	if !m.IsValid() {
		return ""
	}
	r := timeEditAPI().SysCallN(25, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTimeEdit) SetTimeAMString(value string) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(25, 1, m.Instance(), api.PasStr(value))
}

func (m *TTimeEdit) TimeFormat() string {
	if !m.IsValid() {
		return ""
	}
	r := timeEditAPI().SysCallN(26, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTimeEdit) SetTimeFormat(value string) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(26, 1, m.Instance(), api.PasStr(value))
}

func (m *TTimeEdit) TimeSeparator() string {
	if !m.IsValid() {
		return ""
	}
	r := timeEditAPI().SysCallN(27, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTimeEdit) SetTimeSeparator(value string) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(27, 1, m.Instance(), api.PasStr(value))
}

func (m *TTimeEdit) TimePMString() string {
	if !m.IsValid() {
		return ""
	}
	r := timeEditAPI().SysCallN(28, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TTimeEdit) SetTimePMString(value string) {
	if !m.IsValid() {
		return
	}
	timeEditAPI().SysCallN(28, 1, m.Instance(), api.PasStr(value))
}

func (m *TTimeEdit) SetOnAcceptTime(fn TAcceptTimeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTAcceptTimeEvent(fn)
	base.SetEvent(m, 29, timeEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTimeEdit) SetOnCustomTime(fn TCustomTimeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomTimeEvent(fn)
	base.SetEvent(m, 30, timeEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TTimeEdit) SetOnButtonClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 31, timeEditAPI(), api.MakeEventDataPtr(cb))
}

// NewTimeEdit class constructor
func NewTimeEdit(owner IComponent) ITimeEdit {
	r := timeEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsTimeEdit(r)
}

func TTimeEditClass() types.TClass {
	r := timeEditAPI().SysCallN(32)
	return types.TClass(r)
}

var (
	timeEditOnce   base.Once
	timeEditImport *imports.Imports = nil
)

func timeEditAPI() *imports.Imports {
	timeEditOnce.Do(func() {
		timeEditImport = api.NewDefaultImports()
		timeEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTimeEdit_Create", 0), // constructor NewTimeEdit
			/* 1 */ imports.NewTable("TTimeEdit_ValidTimeFormat", 0), // function ValidTimeFormat
			/* 2 */ imports.NewTable("TTimeEdit_Time", 0), // property Time
			/* 3 */ imports.NewTable("TTimeEdit_Button", 0), // property Button
			/* 4 */ imports.NewTable("TTimeEdit_DroppedDown", 0), // property DroppedDown
			/* 5 */ imports.NewTable("TTimeEdit_DefaultNow", 0), // property DefaultNow
			/* 6 */ imports.NewTable("TTimeEdit_ButtonCaption", 0), // property ButtonCaption
			/* 7 */ imports.NewTable("TTimeEdit_ButtonCursor", 0), // property ButtonCursor
			/* 8 */ imports.NewTable("TTimeEdit_ButtonHint", 0), // property ButtonHint
			/* 9 */ imports.NewTable("TTimeEdit_ButtonOnlyWhenFocused", 0), // property ButtonOnlyWhenFocused
			/* 10 */ imports.NewTable("TTimeEdit_ButtonWidth", 0), // property ButtonWidth
			/* 11 */ imports.NewTable("TTimeEdit_AutoSelect", 0), // property AutoSelect
			/* 12 */ imports.NewTable("TTimeEdit_DirectInput", 0), // property DirectInput
			/* 13 */ imports.NewTable("TTimeEdit_Glyph", 0), // property Glyph
			/* 14 */ imports.NewTable("TTimeEdit_NumGlyphs", 0), // property NumGlyphs
			/* 15 */ imports.NewTable("TTimeEdit_Images", 0), // property Images
			/* 16 */ imports.NewTable("TTimeEdit_ImageIndex", 0), // property ImageIndex
			/* 17 */ imports.NewTable("TTimeEdit_ImageWidth", 0), // property ImageWidth
			/* 18 */ imports.NewTable("TTimeEdit_DragMode", 0), // property DragMode
			/* 19 */ imports.NewTable("TTimeEdit_Flat", 0), // property Flat
			/* 20 */ imports.NewTable("TTimeEdit_FocusOnButtonClick", 0), // property FocusOnButtonClick
			/* 21 */ imports.NewTable("TTimeEdit_ParentFont", 0), // property ParentFont
			/* 22 */ imports.NewTable("TTimeEdit_ParentShowHint", 0), // property ParentShowHint
			/* 23 */ imports.NewTable("TTimeEdit_SimpleLayout", 0), // property SimpleLayout
			/* 24 */ imports.NewTable("TTimeEdit_Spacing", 0), // property Spacing
			/* 25 */ imports.NewTable("TTimeEdit_TimeAMString", 0), // property TimeAMString
			/* 26 */ imports.NewTable("TTimeEdit_TimeFormat", 0), // property TimeFormat
			/* 27 */ imports.NewTable("TTimeEdit_TimeSeparator", 0), // property TimeSeparator
			/* 28 */ imports.NewTable("TTimeEdit_TimePMString", 0), // property TimePMString
			/* 29 */ imports.NewTable("TTimeEdit_OnAcceptTime", 0), // event OnAcceptTime
			/* 30 */ imports.NewTable("TTimeEdit_OnCustomTime", 0), // event OnCustomTime
			/* 31 */ imports.NewTable("TTimeEdit_OnButtonClick", 0), // event OnButtonClick
			/* 32 */ imports.NewTable("TTimeEdit_TClass", 0), // function TTimeEditClass
		}
	})
	return timeEditImport
}
