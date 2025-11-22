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

// IDateEdit Parent: ICustomEditButton
type IDateEdit interface {
	ICustomEditButton
	GetDateFormat() string                                   // function
	AutoSelected() bool                                      // property AutoSelected Getter
	SetAutoSelected(value bool)                              // property AutoSelected Setter
	Date() types.TDateTime                                   // property Date Getter
	SetDate(value types.TDateTime)                           // property Date Setter
	Button() ISpeedButton                                    // property Button Getter
	DroppedDown() bool                                       // property DroppedDown Getter
	CalendarDisplaySettings() types.TDisplaySettings         // property CalendarDisplaySettings Getter
	SetCalendarDisplaySettings(value types.TDisplaySettings) // property CalendarDisplaySettings Setter
	DefaultToday() bool                                      // property DefaultToday Getter
	SetDefaultToday(value bool)                              // property DefaultToday Setter
	DateOrder() types.TDateOrder                             // property DateOrder Getter
	SetDateOrder(value types.TDateOrder)                     // property DateOrder Setter
	DateFormat() string                                      // property DateFormat Getter
	SetDateFormat(value string)                              // property DateFormat Setter
	MinDate() types.TDateTime                                // property MinDate Getter
	SetMinDate(value types.TDateTime)                        // property MinDate Setter
	MaxDate() types.TDateTime                                // property MaxDate Getter
	SetMaxDate(value types.TDateTime)                        // property MaxDate Setter
	ButtonOnlyWhenFocused() bool                             // property ButtonOnlyWhenFocused Getter
	SetButtonOnlyWhenFocused(value bool)                     // property ButtonOnlyWhenFocused Setter
	ButtonCaption() string                                   // property ButtonCaption Getter
	SetButtonCaption(value string)                           // property ButtonCaption Setter
	ButtonCursor() types.TCursor                             // property ButtonCursor Getter
	SetButtonCursor(value types.TCursor)                     // property ButtonCursor Setter
	ButtonHint() string                                      // property ButtonHint Getter
	SetButtonHint(value string)                              // property ButtonHint Setter
	ButtonWidth() int32                                      // property ButtonWidth Getter
	SetButtonWidth(value int32)                              // property ButtonWidth Setter
	AutoSelect() bool                                        // property AutoSelect Getter
	SetAutoSelect(value bool)                                // property AutoSelect Setter
	DirectInput() bool                                       // property DirectInput Getter
	SetDirectInput(value bool)                               // property DirectInput Setter
	Glyph() IBitmap                                          // property Glyph Getter
	SetGlyph(value IBitmap)                                  // property Glyph Setter
	NumGlyphs() int32                                        // property NumGlyphs Getter
	SetNumGlyphs(value int32)                                // property NumGlyphs Setter
	Images() ICustomImageList                                // property Images Getter
	SetImages(value ICustomImageList)                        // property Images Setter
	ImageIndex() int32                                       // property ImageIndex Getter
	SetImageIndex(value int32)                               // property ImageIndex Setter
	ImageWidth() int32                                       // property ImageWidth Getter
	SetImageWidth(value int32)                               // property ImageWidth Setter
	DragMode() types.TDragMode                               // property DragMode Getter
	SetDragMode(value types.TDragMode)                       // property DragMode Setter
	Flat() bool                                              // property Flat Getter
	SetFlat(value bool)                                      // property Flat Setter
	FocusOnButtonClick() bool                                // property FocusOnButtonClick Getter
	SetFocusOnButtonClick(value bool)                        // property FocusOnButtonClick Setter
	Layout() types.TLeftRight                                // property Layout Getter
	SetLayout(value types.TLeftRight)                        // property Layout Setter
	ParentFont() bool                                        // property ParentFont Getter
	SetParentFont(value bool)                                // property ParentFont Setter
	ParentShowHint() bool                                    // property ParentShowHint Getter
	SetParentShowHint(value bool)                            // property ParentShowHint Setter
	Spacing() int32                                          // property Spacing Getter
	SetSpacing(value int32)                                  // property Spacing Setter
	SetOnAcceptDate(fn TAcceptDateEvent)                     // property event
	SetOnCustomDate(fn TCustomDateEvent)                     // property event
	SetOnDateRangeCheck(fn TDateRangeCheckEvent)             // property event
	SetOnButtonClick(fn TNotifyEvent)                        // property event
}

type TDateEdit struct {
	TCustomEditButton
}

func (m *TDateEdit) GetDateFormat() string {
	if !m.IsValid() {
		return ""
	}
	r := dateEditAPI().SysCallN(1, m.Instance())
	return api.GoStr(r)
}

func (m *TDateEdit) AutoSelected() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) SetAutoSelected(value bool) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateEdit) Date() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(3, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDateEdit) SetDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(3, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDateEdit) Button() ISpeedButton {
	if !m.IsValid() {
		return nil
	}
	r := dateEditAPI().SysCallN(4, m.Instance())
	return AsSpeedButton(r)
}

func (m *TDateEdit) DroppedDown() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) CalendarDisplaySettings() types.TDisplaySettings {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(6, 0, m.Instance())
	return types.TDisplaySettings(r)
}

func (m *TDateEdit) SetCalendarDisplaySettings(value types.TDisplaySettings) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) DefaultToday() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) SetDefaultToday(value bool) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateEdit) DateOrder() types.TDateOrder {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(8, 0, m.Instance())
	return types.TDateOrder(r)
}

func (m *TDateEdit) SetDateOrder(value types.TDateOrder) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) DateFormat() string {
	if !m.IsValid() {
		return ""
	}
	r := dateEditAPI().SysCallN(9, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDateEdit) SetDateFormat(value string) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(9, 1, m.Instance(), api.PasStr(value))
}

func (m *TDateEdit) MinDate() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(10, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDateEdit) SetMinDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(10, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDateEdit) MaxDate() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(11, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDateEdit) SetMaxDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(11, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDateEdit) ButtonOnlyWhenFocused() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) SetButtonOnlyWhenFocused(value bool) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateEdit) ButtonCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := dateEditAPI().SysCallN(13, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDateEdit) SetButtonCaption(value string) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(13, 1, m.Instance(), api.PasStr(value))
}

func (m *TDateEdit) ButtonCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(14, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TDateEdit) SetButtonCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) ButtonHint() string {
	if !m.IsValid() {
		return ""
	}
	r := dateEditAPI().SysCallN(15, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDateEdit) SetButtonHint(value string) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(15, 1, m.Instance(), api.PasStr(value))
}

func (m *TDateEdit) ButtonWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(16, 0, m.Instance())
	return int32(r)
}

func (m *TDateEdit) SetButtonWidth(value int32) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateEdit) DirectInput() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) SetDirectInput(value bool) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateEdit) Glyph() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := dateEditAPI().SysCallN(19, 0, m.Instance())
	return AsBitmap(r)
}

func (m *TDateEdit) SetGlyph(value IBitmap) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(19, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDateEdit) NumGlyphs() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(20, 0, m.Instance())
	return int32(r)
}

func (m *TDateEdit) SetNumGlyphs(value int32) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := dateEditAPI().SysCallN(21, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TDateEdit) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(21, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDateEdit) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TDateEdit) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) ImageWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TDateEdit) SetImageWidth(value int32) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(24, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TDateEdit) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) Flat() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) SetFlat(value bool) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateEdit) FocusOnButtonClick() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) SetFocusOnButtonClick(value bool) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateEdit) Layout() types.TLeftRight {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(27, 0, m.Instance())
	return types.TLeftRight(r)
}

func (m *TDateEdit) SetLayout(value types.TLeftRight) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(28, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(28, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateEdit) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := dateEditAPI().SysCallN(29, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateEdit) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(29, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateEdit) Spacing() int32 {
	if !m.IsValid() {
		return 0
	}
	r := dateEditAPI().SysCallN(30, 0, m.Instance())
	return int32(r)
}

func (m *TDateEdit) SetSpacing(value int32) {
	if !m.IsValid() {
		return
	}
	dateEditAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

func (m *TDateEdit) SetOnAcceptDate(fn TAcceptDateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTAcceptDateEvent(fn)
	base.SetEvent(m, 31, dateEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateEdit) SetOnCustomDate(fn TCustomDateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCustomDateEvent(fn)
	base.SetEvent(m, 32, dateEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateEdit) SetOnDateRangeCheck(fn TDateRangeCheckEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDateRangeCheckEvent(fn)
	base.SetEvent(m, 33, dateEditAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateEdit) SetOnButtonClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 34, dateEditAPI(), api.MakeEventDataPtr(cb))
}

// NewDateEdit class constructor
func NewDateEdit(owner IComponent) IDateEdit {
	r := dateEditAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsDateEdit(r)
}

func TDateEditClass() types.TClass {
	r := dateEditAPI().SysCallN(35)
	return types.TClass(r)
}

var (
	dateEditOnce   base.Once
	dateEditImport *imports.Imports = nil
)

func dateEditAPI() *imports.Imports {
	dateEditOnce.Do(func() {
		dateEditImport = api.NewDefaultImports()
		dateEditImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDateEdit_Create", 0), // constructor NewDateEdit
			/* 1 */ imports.NewTable("TDateEdit_GetDateFormat", 0), // function GetDateFormat
			/* 2 */ imports.NewTable("TDateEdit_AutoSelected", 0), // property AutoSelected
			/* 3 */ imports.NewTable("TDateEdit_Date", 0), // property Date
			/* 4 */ imports.NewTable("TDateEdit_Button", 0), // property Button
			/* 5 */ imports.NewTable("TDateEdit_DroppedDown", 0), // property DroppedDown
			/* 6 */ imports.NewTable("TDateEdit_CalendarDisplaySettings", 0), // property CalendarDisplaySettings
			/* 7 */ imports.NewTable("TDateEdit_DefaultToday", 0), // property DefaultToday
			/* 8 */ imports.NewTable("TDateEdit_DateOrder", 0), // property DateOrder
			/* 9 */ imports.NewTable("TDateEdit_DateFormat", 0), // property DateFormat
			/* 10 */ imports.NewTable("TDateEdit_MinDate", 0), // property MinDate
			/* 11 */ imports.NewTable("TDateEdit_MaxDate", 0), // property MaxDate
			/* 12 */ imports.NewTable("TDateEdit_ButtonOnlyWhenFocused", 0), // property ButtonOnlyWhenFocused
			/* 13 */ imports.NewTable("TDateEdit_ButtonCaption", 0), // property ButtonCaption
			/* 14 */ imports.NewTable("TDateEdit_ButtonCursor", 0), // property ButtonCursor
			/* 15 */ imports.NewTable("TDateEdit_ButtonHint", 0), // property ButtonHint
			/* 16 */ imports.NewTable("TDateEdit_ButtonWidth", 0), // property ButtonWidth
			/* 17 */ imports.NewTable("TDateEdit_AutoSelect", 0), // property AutoSelect
			/* 18 */ imports.NewTable("TDateEdit_DirectInput", 0), // property DirectInput
			/* 19 */ imports.NewTable("TDateEdit_Glyph", 0), // property Glyph
			/* 20 */ imports.NewTable("TDateEdit_NumGlyphs", 0), // property NumGlyphs
			/* 21 */ imports.NewTable("TDateEdit_Images", 0), // property Images
			/* 22 */ imports.NewTable("TDateEdit_ImageIndex", 0), // property ImageIndex
			/* 23 */ imports.NewTable("TDateEdit_ImageWidth", 0), // property ImageWidth
			/* 24 */ imports.NewTable("TDateEdit_DragMode", 0), // property DragMode
			/* 25 */ imports.NewTable("TDateEdit_Flat", 0), // property Flat
			/* 26 */ imports.NewTable("TDateEdit_FocusOnButtonClick", 0), // property FocusOnButtonClick
			/* 27 */ imports.NewTable("TDateEdit_Layout", 0), // property Layout
			/* 28 */ imports.NewTable("TDateEdit_ParentFont", 0), // property ParentFont
			/* 29 */ imports.NewTable("TDateEdit_ParentShowHint", 0), // property ParentShowHint
			/* 30 */ imports.NewTable("TDateEdit_Spacing", 0), // property Spacing
			/* 31 */ imports.NewTable("TDateEdit_OnAcceptDate", 0), // event OnAcceptDate
			/* 32 */ imports.NewTable("TDateEdit_OnCustomDate", 0), // event OnCustomDate
			/* 33 */ imports.NewTable("TDateEdit_OnDateRangeCheck", 0), // event OnDateRangeCheck
			/* 34 */ imports.NewTable("TDateEdit_OnButtonClick", 0), // event OnButtonClick
			/* 35 */ imports.NewTable("TDateEdit_TClass", 0), // function TDateEditClass
		}
	})
	return dateEditImport
}
