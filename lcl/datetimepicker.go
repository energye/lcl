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

// IDateTimePicker Parent: ICustomDateTimePicker
type IDateTimePicker interface {
	ICustomDateTimePicker
	DateTime() types.TDateTime                         // property DateTime Getter
	SetDateTime(value types.TDateTime)                 // property DateTime Setter
	DroppedDown() bool                                 // property DroppedDown Getter
	ArrowShape() types.TArrowShape                     // property ArrowShape Getter
	SetArrowShape(value types.TArrowShape)             // property ArrowShape Setter
	ShowCheckBox() bool                                // property ShowCheckBox Getter
	SetShowCheckBox(value bool)                        // property ShowCheckBox Setter
	Checked() bool                                     // property Checked Getter
	SetChecked(value bool)                             // property Checked Setter
	CenturyFrom() uint16                               // property CenturyFrom Getter
	SetCenturyFrom(value uint16)                       // property CenturyFrom Setter
	DateDisplayOrder() types.TDateDisplayOrder         // property DateDisplayOrder Getter
	SetDateDisplayOrder(value types.TDateDisplayOrder) // property DateDisplayOrder Setter
	MaxDate() types.TDateTime                          // property MaxDate Getter
	SetMaxDate(value types.TDateTime)                  // property MaxDate Setter
	MinDate() types.TDateTime                          // property MinDate Getter
	SetMinDate(value types.TDateTime)                  // property MinDate Setter
	ReadOnly() bool                                    // property ReadOnly Getter
	SetReadOnly(value bool)                            // property ReadOnly Setter
	ParentFont() bool                                  // property ParentFont Getter
	SetParentFont(value bool)                          // property ParentFont Setter
	ParentColor() bool                                 // property ParentColor Getter
	SetParentColor(value bool)                         // property ParentColor Setter
	DateSeparator() string                             // property DateSeparator Getter
	SetDateSeparator(value string)                     // property DateSeparator Setter
	TrailingSeparator() bool                           // property TrailingSeparator Getter
	SetTrailingSeparator(value bool)                   // property TrailingSeparator Setter
	TextForNullDate() string                           // property TextForNullDate Getter
	SetTextForNullDate(value string)                   // property TextForNullDate Setter
	LeadingZeros() bool                                // property LeadingZeros Getter
	SetLeadingZeros(value bool)                        // property LeadingZeros Setter
	ParentShowHint() bool                              // property ParentShowHint Getter
	SetParentShowHint(value bool)                      // property ParentShowHint Setter
	NullInputAllowed() bool                            // property NullInputAllowed Getter
	SetNullInputAllowed(value bool)                    // property NullInputAllowed Setter
	Kind() types.TDateTimeKind                         // property Kind Getter
	SetKind(value types.TDateTimeKind)                 // property Kind Setter
	TimeSeparator() string                             // property TimeSeparator Getter
	SetTimeSeparator(value string)                     // property TimeSeparator Setter
	DecimalSeparator() string                          // property DecimalSeparator Getter
	SetDecimalSeparator(value string)                  // property DecimalSeparator Setter
	TimeFormat() types.TTimeFormat                     // property TimeFormat Getter
	SetTimeFormat(value types.TTimeFormat)             // property TimeFormat Setter
	TimeDisplay() types.TTimeDisplay                   // property TimeDisplay Getter
	SetTimeDisplay(value types.TTimeDisplay)           // property TimeDisplay Setter
	DateMode() types.TDTDateMode                       // property DateMode Getter
	SetDateMode(value types.TDTDateMode)               // property DateMode Setter
	Date() types.TDateTime                             // property Date Getter
	SetDate(value types.TDateTime)                     // property Date Setter
	Time() types.TDateTime                             // property Time Getter
	SetTime(value types.TDateTime)                     // property Time Setter
	UseDefaultSeparators() bool                        // property UseDefaultSeparators Getter
	SetUseDefaultSeparators(value bool)                // property UseDefaultSeparators Setter
	Cascade() bool                                     // property Cascade Getter
	SetCascade(value bool)                             // property Cascade Setter
	AutoButtonSize() bool                              // property AutoButtonSize Getter
	SetAutoButtonSize(value bool)                      // property AutoButtonSize Setter
	AutoAdvance() bool                                 // property AutoAdvance Getter
	SetAutoAdvance(value bool)                         // property AutoAdvance Setter
	HideDateTimeParts() types.TDateTimeParts           // property HideDateTimeParts Getter
	SetHideDateTimeParts(value types.TDateTimeParts)   // property HideDateTimeParts Setter
	MonthDisplay() types.TMonthDisplay                 // property MonthDisplay Getter
	SetMonthDisplay(value types.TMonthDisplay)         // property MonthDisplay Setter
	CustomMonthNames() IStrings                        // property CustomMonthNames Getter
	SetCustomMonthNames(value IStrings)                // property CustomMonthNames Setter
	ShowMonthNames() bool                              // property ShowMonthNames Getter
	SetShowMonthNames(value bool)                      // property ShowMonthNames Setter
	CalAlignment() types.TDTCalAlignment               // property CalAlignment Getter
	SetCalAlignment(value types.TDTCalAlignment)       // property CalAlignment Setter
	Alignment() types.TAlignment                       // property Alignment Getter
	SetAlignment(value types.TAlignment)               // property Alignment Setter
	Options() types.TDateTimePickerOptions             // property Options Getter
	SetOptions(value types.TDateTimePickerOptions)     // property Options Setter
	SetOnChange(fn TNotifyEvent)                       // property event
	SetOnCheckBoxChange(fn TNotifyEvent)               // property event
	SetOnDropDown(fn TNotifyEvent)                     // property event
	SetOnCloseUp(fn TNotifyEvent)                      // property event
	SetOnContextPopup(fn TContextPopupEvent)           // property event
	SetOnDblClick(fn TNotifyEvent)                     // property event
	SetOnEditingDone(fn TNotifyEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                     // property event
	SetOnMouseEnter(fn TNotifyEvent)                   // property event
	SetOnMouseLeave(fn TNotifyEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                       // property event
	SetOnMouseWheel(fn TMouseWheelEvent)               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)       // property event
}

type TDateTimePicker struct {
	TCustomDateTimePicker
}

func (m *TDateTimePicker) DateTime() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(1, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDateTimePicker) SetDateTime(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(1, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDateTimePicker) DroppedDown() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) ArrowShape() types.TArrowShape {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(3, 0, m.Instance())
	return types.TArrowShape(r)
}

func (m *TDateTimePicker) SetArrowShape(value types.TArrowShape) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(3, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) ShowCheckBox() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetShowCheckBox(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) Checked() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetChecked(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) CenturyFrom() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(6, 0, m.Instance())
	return uint16(r)
}

func (m *TDateTimePicker) SetCenturyFrom(value uint16) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) DateDisplayOrder() types.TDateDisplayOrder {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(7, 0, m.Instance())
	return types.TDateDisplayOrder(r)
}

func (m *TDateTimePicker) SetDateDisplayOrder(value types.TDateDisplayOrder) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) MaxDate() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(8, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDateTimePicker) SetMaxDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(8, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDateTimePicker) MinDate() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(9, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDateTimePicker) SetMinDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(9, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDateTimePicker) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) DateSeparator() string {
	if !m.IsValid() {
		return ""
	}
	r := dateTimePickerAPI().SysCallN(13, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDateTimePicker) SetDateSeparator(value string) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(13, 1, m.Instance(), api.PasStr(value))
}

func (m *TDateTimePicker) TrailingSeparator() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetTrailingSeparator(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) TextForNullDate() string {
	if !m.IsValid() {
		return ""
	}
	r := dateTimePickerAPI().SysCallN(15, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDateTimePicker) SetTextForNullDate(value string) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(15, 1, m.Instance(), api.PasStr(value))
}

func (m *TDateTimePicker) LeadingZeros() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetLeadingZeros(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) NullInputAllowed() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetNullInputAllowed(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) Kind() types.TDateTimeKind {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(19, 0, m.Instance())
	return types.TDateTimeKind(r)
}

func (m *TDateTimePicker) SetKind(value types.TDateTimeKind) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) TimeSeparator() string {
	if !m.IsValid() {
		return ""
	}
	r := dateTimePickerAPI().SysCallN(20, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDateTimePicker) SetTimeSeparator(value string) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(20, 1, m.Instance(), api.PasStr(value))
}

func (m *TDateTimePicker) DecimalSeparator() string {
	if !m.IsValid() {
		return ""
	}
	r := dateTimePickerAPI().SysCallN(21, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TDateTimePicker) SetDecimalSeparator(value string) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(21, 1, m.Instance(), api.PasStr(value))
}

func (m *TDateTimePicker) TimeFormat() types.TTimeFormat {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(22, 0, m.Instance())
	return types.TTimeFormat(r)
}

func (m *TDateTimePicker) SetTimeFormat(value types.TTimeFormat) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) TimeDisplay() types.TTimeDisplay {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(23, 0, m.Instance())
	return types.TTimeDisplay(r)
}

func (m *TDateTimePicker) SetTimeDisplay(value types.TTimeDisplay) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) DateMode() types.TDTDateMode {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(24, 0, m.Instance())
	return types.TDTDateMode(r)
}

func (m *TDateTimePicker) SetDateMode(value types.TDTDateMode) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) Date() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(25, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDateTimePicker) SetDate(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(25, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDateTimePicker) Time() (result types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(26, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TDateTimePicker) SetTime(value types.TDateTime) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(26, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TDateTimePicker) UseDefaultSeparators() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(27, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetUseDefaultSeparators(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(27, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) Cascade() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(28, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetCascade(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(28, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) AutoButtonSize() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(29, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetAutoButtonSize(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(29, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) AutoAdvance() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(30, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetAutoAdvance(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(30, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) HideDateTimeParts() types.TDateTimeParts {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(31, 0, m.Instance())
	return types.TDateTimeParts(r)
}

func (m *TDateTimePicker) SetHideDateTimeParts(value types.TDateTimeParts) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) MonthDisplay() types.TMonthDisplay {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(32, 0, m.Instance())
	return types.TMonthDisplay(r)
}

func (m *TDateTimePicker) SetMonthDisplay(value types.TMonthDisplay) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(32, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) CustomMonthNames() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := dateTimePickerAPI().SysCallN(33, 0, m.Instance())
	return AsStrings(r)
}

func (m *TDateTimePicker) SetCustomMonthNames(value IStrings) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(33, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TDateTimePicker) ShowMonthNames() bool {
	if !m.IsValid() {
		return false
	}
	r := dateTimePickerAPI().SysCallN(34, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TDateTimePicker) SetShowMonthNames(value bool) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(34, 1, m.Instance(), api.PasBool(value))
}

func (m *TDateTimePicker) CalAlignment() types.TDTCalAlignment {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(35, 0, m.Instance())
	return types.TDTCalAlignment(r)
}

func (m *TDateTimePicker) SetCalAlignment(value types.TDTCalAlignment) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(35, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) Alignment() types.TAlignment {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(36, 0, m.Instance())
	return types.TAlignment(r)
}

func (m *TDateTimePicker) SetAlignment(value types.TAlignment) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) Options() types.TDateTimePickerOptions {
	if !m.IsValid() {
		return 0
	}
	r := dateTimePickerAPI().SysCallN(37, 0, m.Instance())
	return types.TDateTimePickerOptions(r)
}

func (m *TDateTimePicker) SetOptions(value types.TDateTimePickerOptions) {
	if !m.IsValid() {
		return
	}
	dateTimePickerAPI().SysCallN(37, 1, m.Instance(), uintptr(value))
}

func (m *TDateTimePicker) SetOnChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 38, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnCheckBoxChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 39, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnDropDown(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 40, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnCloseUp(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 41, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnContextPopup(fn TContextPopupEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTContextPopupEvent(fn)
	base.SetEvent(m, 42, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 43, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 44, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 45, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 46, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 47, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 48, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 49, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 50, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 51, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

func (m *TDateTimePicker) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 52, dateTimePickerAPI(), api.MakeEventDataPtr(cb))
}

// NewDateTimePicker class constructor
func NewDateTimePicker(owner IComponent) IDateTimePicker {
	r := dateTimePickerAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsDateTimePicker(r)
}

func TDateTimePickerClass() types.TClass {
	r := dateTimePickerAPI().SysCallN(53)
	return types.TClass(r)
}

var (
	dateTimePickerOnce   base.Once
	dateTimePickerImport *imports.Imports = nil
)

func dateTimePickerAPI() *imports.Imports {
	dateTimePickerOnce.Do(func() {
		dateTimePickerImport = api.NewDefaultImports()
		dateTimePickerImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TDateTimePicker_Create", 0), // constructor NewDateTimePicker
			/* 1 */ imports.NewTable("TDateTimePicker_DateTime", 0), // property DateTime
			/* 2 */ imports.NewTable("TDateTimePicker_DroppedDown", 0), // property DroppedDown
			/* 3 */ imports.NewTable("TDateTimePicker_ArrowShape", 0), // property ArrowShape
			/* 4 */ imports.NewTable("TDateTimePicker_ShowCheckBox", 0), // property ShowCheckBox
			/* 5 */ imports.NewTable("TDateTimePicker_Checked", 0), // property Checked
			/* 6 */ imports.NewTable("TDateTimePicker_CenturyFrom", 0), // property CenturyFrom
			/* 7 */ imports.NewTable("TDateTimePicker_DateDisplayOrder", 0), // property DateDisplayOrder
			/* 8 */ imports.NewTable("TDateTimePicker_MaxDate", 0), // property MaxDate
			/* 9 */ imports.NewTable("TDateTimePicker_MinDate", 0), // property MinDate
			/* 10 */ imports.NewTable("TDateTimePicker_ReadOnly", 0), // property ReadOnly
			/* 11 */ imports.NewTable("TDateTimePicker_ParentFont", 0), // property ParentFont
			/* 12 */ imports.NewTable("TDateTimePicker_ParentColor", 0), // property ParentColor
			/* 13 */ imports.NewTable("TDateTimePicker_DateSeparator", 0), // property DateSeparator
			/* 14 */ imports.NewTable("TDateTimePicker_TrailingSeparator", 0), // property TrailingSeparator
			/* 15 */ imports.NewTable("TDateTimePicker_TextForNullDate", 0), // property TextForNullDate
			/* 16 */ imports.NewTable("TDateTimePicker_LeadingZeros", 0), // property LeadingZeros
			/* 17 */ imports.NewTable("TDateTimePicker_ParentShowHint", 0), // property ParentShowHint
			/* 18 */ imports.NewTable("TDateTimePicker_NullInputAllowed", 0), // property NullInputAllowed
			/* 19 */ imports.NewTable("TDateTimePicker_Kind", 0), // property Kind
			/* 20 */ imports.NewTable("TDateTimePicker_TimeSeparator", 0), // property TimeSeparator
			/* 21 */ imports.NewTable("TDateTimePicker_DecimalSeparator", 0), // property DecimalSeparator
			/* 22 */ imports.NewTable("TDateTimePicker_TimeFormat", 0), // property TimeFormat
			/* 23 */ imports.NewTable("TDateTimePicker_TimeDisplay", 0), // property TimeDisplay
			/* 24 */ imports.NewTable("TDateTimePicker_DateMode", 0), // property DateMode
			/* 25 */ imports.NewTable("TDateTimePicker_Date", 0), // property Date
			/* 26 */ imports.NewTable("TDateTimePicker_Time", 0), // property Time
			/* 27 */ imports.NewTable("TDateTimePicker_UseDefaultSeparators", 0), // property UseDefaultSeparators
			/* 28 */ imports.NewTable("TDateTimePicker_Cascade", 0), // property Cascade
			/* 29 */ imports.NewTable("TDateTimePicker_AutoButtonSize", 0), // property AutoButtonSize
			/* 30 */ imports.NewTable("TDateTimePicker_AutoAdvance", 0), // property AutoAdvance
			/* 31 */ imports.NewTable("TDateTimePicker_HideDateTimeParts", 0), // property HideDateTimeParts
			/* 32 */ imports.NewTable("TDateTimePicker_MonthDisplay", 0), // property MonthDisplay
			/* 33 */ imports.NewTable("TDateTimePicker_CustomMonthNames", 0), // property CustomMonthNames
			/* 34 */ imports.NewTable("TDateTimePicker_ShowMonthNames", 0), // property ShowMonthNames
			/* 35 */ imports.NewTable("TDateTimePicker_CalAlignment", 0), // property CalAlignment
			/* 36 */ imports.NewTable("TDateTimePicker_Alignment", 0), // property Alignment
			/* 37 */ imports.NewTable("TDateTimePicker_Options", 0), // property Options
			/* 38 */ imports.NewTable("TDateTimePicker_OnChange", 0), // event OnChange
			/* 39 */ imports.NewTable("TDateTimePicker_OnCheckBoxChange", 0), // event OnCheckBoxChange
			/* 40 */ imports.NewTable("TDateTimePicker_OnDropDown", 0), // event OnDropDown
			/* 41 */ imports.NewTable("TDateTimePicker_OnCloseUp", 0), // event OnCloseUp
			/* 42 */ imports.NewTable("TDateTimePicker_OnContextPopup", 0), // event OnContextPopup
			/* 43 */ imports.NewTable("TDateTimePicker_OnDblClick", 0), // event OnDblClick
			/* 44 */ imports.NewTable("TDateTimePicker_OnEditingDone", 0), // event OnEditingDone
			/* 45 */ imports.NewTable("TDateTimePicker_OnMouseDown", 0), // event OnMouseDown
			/* 46 */ imports.NewTable("TDateTimePicker_OnMouseEnter", 0), // event OnMouseEnter
			/* 47 */ imports.NewTable("TDateTimePicker_OnMouseLeave", 0), // event OnMouseLeave
			/* 48 */ imports.NewTable("TDateTimePicker_OnMouseMove", 0), // event OnMouseMove
			/* 49 */ imports.NewTable("TDateTimePicker_OnMouseUp", 0), // event OnMouseUp
			/* 50 */ imports.NewTable("TDateTimePicker_OnMouseWheel", 0), // event OnMouseWheel
			/* 51 */ imports.NewTable("TDateTimePicker_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 52 */ imports.NewTable("TDateTimePicker_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 53 */ imports.NewTable("TDateTimePicker_TClass", 0), // function TDateTimePickerClass
		}
	})
	return dateTimePickerImport
}
