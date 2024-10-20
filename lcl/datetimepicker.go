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

// IDateTimePicker Parent: ICustomDateTimePicker
type IDateTimePicker interface {
	ICustomDateTimePicker
	DateTime() TDateTime                           // property
	SetDateTime(AValue TDateTime)                  // property
	DroppedDown() bool                             // property
	ArrowShape() TArrowShape                       // property
	SetArrowShape(AValue TArrowShape)              // property
	ShowCheckBox() bool                            // property
	SetShowCheckBox(AValue bool)                   // property
	Checked() bool                                 // property
	SetChecked(AValue bool)                        // property
	CenturyFrom() Word                             // property
	SetCenturyFrom(AValue Word)                    // property
	DateDisplayOrder() TDateDisplayOrder           // property
	SetDateDisplayOrder(AValue TDateDisplayOrder)  // property
	MaxDate() TDate                                // property
	SetMaxDate(AValue TDate)                       // property
	MinDate() TDate                                // property
	SetMinDate(AValue TDate)                       // property
	ReadOnly() bool                                // property
	SetReadOnly(AValue bool)                       // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	DateSeparator() string                         // property
	SetDateSeparator(AValue string)                // property
	TrailingSeparator() bool                       // property
	SetTrailingSeparator(AValue bool)              // property
	TextForNullDate() string                       // property
	SetTextForNullDate(AValue string)              // property
	LeadingZeros() bool                            // property
	SetLeadingZeros(AValue bool)                   // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	NullInputAllowed() bool                        // property
	SetNullInputAllowed(AValue bool)               // property
	Kind() TDateTimeKind                           // property
	SetKind(AValue TDateTimeKind)                  // property
	TimeSeparator() string                         // property
	SetTimeSeparator(AValue string)                // property
	DecimalSeparator() string                      // property
	SetDecimalSeparator(AValue string)             // property
	TimeFormat() TTimeFormat                       // property
	SetTimeFormat(AValue TTimeFormat)              // property
	TimeDisplay() TTimeDisplay                     // property
	SetTimeDisplay(AValue TTimeDisplay)            // property
	DateMode() TDTDateMode                         // property
	SetDateMode(AValue TDTDateMode)                // property
	Date() TDate                                   // property
	SetDate(AValue TDate)                          // property
	Time() TTime                                   // property
	SetTime(AValue TTime)                          // property
	UseDefaultSeparators() bool                    // property
	SetUseDefaultSeparators(AValue bool)           // property
	Cascade() bool                                 // property
	SetCascade(AValue bool)                        // property
	AutoButtonSize() bool                          // property
	SetAutoButtonSize(AValue bool)                 // property
	AutoAdvance() bool                             // property
	SetAutoAdvance(AValue bool)                    // property
	HideDateTimeParts() TDateTimeParts             // property
	SetHideDateTimeParts(AValue TDateTimeParts)    // property
	MonthDisplay() TMonthDisplay                   // property
	SetMonthDisplay(AValue TMonthDisplay)          // property
	CustomMonthNames() IStrings                    // property
	SetCustomMonthNames(AValue IStrings)           // property
	ShowMonthNames() bool                          // property
	SetShowMonthNames(AValue bool)                 // property
	CalAlignment() TDTCalAlignment                 // property
	SetCalAlignment(AValue TDTCalAlignment)        // property
	Alignment() TAlignment                         // property
	SetAlignment(AValue TAlignment)                // property
	Options() TDateTimePickerOptions               // property
	SetOptions(AValue TDateTimePickerOptions)      // property
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnCheckBoxChange(fn TNotifyEvent)           // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnCloseUp(fn TNotifyEvent)                  // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
}

// TDateTimePicker Parent: TCustomDateTimePicker
type TDateTimePicker struct {
	TCustomDateTimePicker
	changePtr         uintptr
	checkBoxChangePtr uintptr
	dropDownPtr       uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	editingDonePtr    uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
}

func NewDateTimePicker(AOwner IComponent) IDateTimePicker {
	r1 := dateTimePickerImportAPI().SysCallN(9, GetObjectUintptr(AOwner))
	return AsDateTimePicker(r1)
}

func (m *TDateTimePicker) DateTime() TDateTime {
	r1 := dateTimePickerImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TDateTime(r1)
}

func (m *TDateTimePicker) SetDateTime(AValue TDateTime) {
	dateTimePickerImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) DroppedDown() bool {
	r1 := dateTimePickerImportAPI().SysCallN(17, m.Instance())
	return GoBool(r1)
}

func (m *TDateTimePicker) ArrowShape() TArrowShape {
	r1 := dateTimePickerImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TArrowShape(r1)
}

func (m *TDateTimePicker) SetArrowShape(AValue TArrowShape) {
	dateTimePickerImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) ShowCheckBox() bool {
	r1 := dateTimePickerImportAPI().SysCallN(45, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetShowCheckBox(AValue bool) {
	dateTimePickerImportAPI().SysCallN(45, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) Checked() bool {
	r1 := dateTimePickerImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetChecked(AValue bool) {
	dateTimePickerImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) CenturyFrom() Word {
	r1 := dateTimePickerImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return Word(r1)
}

func (m *TDateTimePicker) SetCenturyFrom(AValue Word) {
	dateTimePickerImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) DateDisplayOrder() TDateDisplayOrder {
	r1 := dateTimePickerImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TDateDisplayOrder(r1)
}

func (m *TDateTimePicker) SetDateDisplayOrder(AValue TDateDisplayOrder) {
	dateTimePickerImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) MaxDate() TDate {
	r1 := dateTimePickerImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return TDate(r1)
}

func (m *TDateTimePicker) SetMaxDate(AValue TDate) {
	dateTimePickerImportAPI().SysCallN(21, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) MinDate() TDate {
	r1 := dateTimePickerImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return TDate(r1)
}

func (m *TDateTimePicker) SetMinDate(AValue TDate) {
	dateTimePickerImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) ReadOnly() bool {
	r1 := dateTimePickerImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetReadOnly(AValue bool) {
	dateTimePickerImportAPI().SysCallN(29, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) ParentFont() bool {
	r1 := dateTimePickerImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetParentFont(AValue bool) {
	dateTimePickerImportAPI().SysCallN(27, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) ParentColor() bool {
	r1 := dateTimePickerImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetParentColor(AValue bool) {
	dateTimePickerImportAPI().SysCallN(26, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) DateSeparator() string {
	r1 := dateTimePickerImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDateTimePicker) SetDateSeparator(AValue string) {
	dateTimePickerImportAPI().SysCallN(14, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDateTimePicker) TrailingSeparator() bool {
	r1 := dateTimePickerImportAPI().SysCallN(52, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetTrailingSeparator(AValue bool) {
	dateTimePickerImportAPI().SysCallN(52, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) TextForNullDate() string {
	r1 := dateTimePickerImportAPI().SysCallN(47, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDateTimePicker) SetTextForNullDate(AValue string) {
	dateTimePickerImportAPI().SysCallN(47, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDateTimePicker) LeadingZeros() bool {
	r1 := dateTimePickerImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetLeadingZeros(AValue bool) {
	dateTimePickerImportAPI().SysCallN(20, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) ParentShowHint() bool {
	r1 := dateTimePickerImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetParentShowHint(AValue bool) {
	dateTimePickerImportAPI().SysCallN(28, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) NullInputAllowed() bool {
	r1 := dateTimePickerImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetNullInputAllowed(AValue bool) {
	dateTimePickerImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) Kind() TDateTimeKind {
	r1 := dateTimePickerImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TDateTimeKind(r1)
}

func (m *TDateTimePicker) SetKind(AValue TDateTimeKind) {
	dateTimePickerImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) TimeSeparator() string {
	r1 := dateTimePickerImportAPI().SysCallN(51, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDateTimePicker) SetTimeSeparator(AValue string) {
	dateTimePickerImportAPI().SysCallN(51, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDateTimePicker) DecimalSeparator() string {
	r1 := dateTimePickerImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDateTimePicker) SetDecimalSeparator(AValue string) {
	dateTimePickerImportAPI().SysCallN(16, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDateTimePicker) TimeFormat() TTimeFormat {
	r1 := dateTimePickerImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return TTimeFormat(r1)
}

func (m *TDateTimePicker) SetTimeFormat(AValue TTimeFormat) {
	dateTimePickerImportAPI().SysCallN(50, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) TimeDisplay() TTimeDisplay {
	r1 := dateTimePickerImportAPI().SysCallN(49, 0, m.Instance(), 0)
	return TTimeDisplay(r1)
}

func (m *TDateTimePicker) SetTimeDisplay(AValue TTimeDisplay) {
	dateTimePickerImportAPI().SysCallN(49, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) DateMode() TDTDateMode {
	r1 := dateTimePickerImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TDTDateMode(r1)
}

func (m *TDateTimePicker) SetDateMode(AValue TDTDateMode) {
	dateTimePickerImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) Date() TDate {
	r1 := dateTimePickerImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TDate(r1)
}

func (m *TDateTimePicker) SetDate(AValue TDate) {
	dateTimePickerImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) Time() TTime {
	r1 := dateTimePickerImportAPI().SysCallN(48, 0, m.Instance(), 0)
	return TTime(r1)
}

func (m *TDateTimePicker) SetTime(AValue TTime) {
	dateTimePickerImportAPI().SysCallN(48, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) UseDefaultSeparators() bool {
	r1 := dateTimePickerImportAPI().SysCallN(53, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetUseDefaultSeparators(AValue bool) {
	dateTimePickerImportAPI().SysCallN(53, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) Cascade() bool {
	r1 := dateTimePickerImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetCascade(AValue bool) {
	dateTimePickerImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) AutoButtonSize() bool {
	r1 := dateTimePickerImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetAutoButtonSize(AValue bool) {
	dateTimePickerImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) AutoAdvance() bool {
	r1 := dateTimePickerImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetAutoAdvance(AValue bool) {
	dateTimePickerImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) HideDateTimeParts() TDateTimeParts {
	r1 := dateTimePickerImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return TDateTimeParts(r1)
}

func (m *TDateTimePicker) SetHideDateTimeParts(AValue TDateTimeParts) {
	dateTimePickerImportAPI().SysCallN(18, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) MonthDisplay() TMonthDisplay {
	r1 := dateTimePickerImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return TMonthDisplay(r1)
}

func (m *TDateTimePicker) SetMonthDisplay(AValue TMonthDisplay) {
	dateTimePickerImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) CustomMonthNames() IStrings {
	r1 := dateTimePickerImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TDateTimePicker) SetCustomMonthNames(AValue IStrings) {
	dateTimePickerImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDateTimePicker) ShowMonthNames() bool {
	r1 := dateTimePickerImportAPI().SysCallN(46, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetShowMonthNames(AValue bool) {
	dateTimePickerImportAPI().SysCallN(46, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) CalAlignment() TDTCalAlignment {
	r1 := dateTimePickerImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDTCalAlignment(r1)
}

func (m *TDateTimePicker) SetCalAlignment(AValue TDTCalAlignment) {
	dateTimePickerImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) Alignment() TAlignment {
	r1 := dateTimePickerImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TDateTimePicker) SetAlignment(AValue TAlignment) {
	dateTimePickerImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) Options() TDateTimePickerOptions {
	r1 := dateTimePickerImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return TDateTimePickerOptions(r1)
}

func (m *TDateTimePicker) SetOptions(AValue TDateTimePickerOptions) {
	dateTimePickerImportAPI().SysCallN(25, 1, m.Instance(), uintptr(AValue))
}

func DateTimePickerClass() TClass {
	ret := dateTimePickerImportAPI().SysCallN(8)
	return TClass(ret)
}

func (m *TDateTimePicker) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(30, m.Instance(), m.changePtr)
}

func (m *TDateTimePicker) SetOnCheckBoxChange(fn TNotifyEvent) {
	if m.checkBoxChangePtr != 0 {
		RemoveEventElement(m.checkBoxChangePtr)
	}
	m.checkBoxChangePtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(31, m.Instance(), m.checkBoxChangePtr)
}

func (m *TDateTimePicker) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(35, m.Instance(), m.dropDownPtr)
}

func (m *TDateTimePicker) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(32, m.Instance(), m.closeUpPtr)
}

func (m *TDateTimePicker) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(33, m.Instance(), m.contextPopupPtr)
}

func (m *TDateTimePicker) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(34, m.Instance(), m.dblClickPtr)
}

func (m *TDateTimePicker) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(36, m.Instance(), m.editingDonePtr)
}

func (m *TDateTimePicker) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(37, m.Instance(), m.mouseDownPtr)
}

func (m *TDateTimePicker) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(38, m.Instance(), m.mouseEnterPtr)
}

func (m *TDateTimePicker) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(39, m.Instance(), m.mouseLeavePtr)
}

func (m *TDateTimePicker) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(40, m.Instance(), m.mouseMovePtr)
}

func (m *TDateTimePicker) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(41, m.Instance(), m.mouseUpPtr)
}

func (m *TDateTimePicker) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(42, m.Instance(), m.mouseWheelPtr)
}

func (m *TDateTimePicker) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(43, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TDateTimePicker) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	dateTimePickerImportAPI().SysCallN(44, m.Instance(), m.mouseWheelUpPtr)
}

var (
	dateTimePickerImport       *imports.Imports = nil
	dateTimePickerImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("DateTimePicker_Alignment", 0),
		/*1*/ imports.NewTable("DateTimePicker_ArrowShape", 0),
		/*2*/ imports.NewTable("DateTimePicker_AutoAdvance", 0),
		/*3*/ imports.NewTable("DateTimePicker_AutoButtonSize", 0),
		/*4*/ imports.NewTable("DateTimePicker_CalAlignment", 0),
		/*5*/ imports.NewTable("DateTimePicker_Cascade", 0),
		/*6*/ imports.NewTable("DateTimePicker_CenturyFrom", 0),
		/*7*/ imports.NewTable("DateTimePicker_Checked", 0),
		/*8*/ imports.NewTable("DateTimePicker_Class", 0),
		/*9*/ imports.NewTable("DateTimePicker_Create", 0),
		/*10*/ imports.NewTable("DateTimePicker_CustomMonthNames", 0),
		/*11*/ imports.NewTable("DateTimePicker_Date", 0),
		/*12*/ imports.NewTable("DateTimePicker_DateDisplayOrder", 0),
		/*13*/ imports.NewTable("DateTimePicker_DateMode", 0),
		/*14*/ imports.NewTable("DateTimePicker_DateSeparator", 0),
		/*15*/ imports.NewTable("DateTimePicker_DateTime", 0),
		/*16*/ imports.NewTable("DateTimePicker_DecimalSeparator", 0),
		/*17*/ imports.NewTable("DateTimePicker_DroppedDown", 0),
		/*18*/ imports.NewTable("DateTimePicker_HideDateTimeParts", 0),
		/*19*/ imports.NewTable("DateTimePicker_Kind", 0),
		/*20*/ imports.NewTable("DateTimePicker_LeadingZeros", 0),
		/*21*/ imports.NewTable("DateTimePicker_MaxDate", 0),
		/*22*/ imports.NewTable("DateTimePicker_MinDate", 0),
		/*23*/ imports.NewTable("DateTimePicker_MonthDisplay", 0),
		/*24*/ imports.NewTable("DateTimePicker_NullInputAllowed", 0),
		/*25*/ imports.NewTable("DateTimePicker_Options", 0),
		/*26*/ imports.NewTable("DateTimePicker_ParentColor", 0),
		/*27*/ imports.NewTable("DateTimePicker_ParentFont", 0),
		/*28*/ imports.NewTable("DateTimePicker_ParentShowHint", 0),
		/*29*/ imports.NewTable("DateTimePicker_ReadOnly", 0),
		/*30*/ imports.NewTable("DateTimePicker_SetOnChange", 0),
		/*31*/ imports.NewTable("DateTimePicker_SetOnCheckBoxChange", 0),
		/*32*/ imports.NewTable("DateTimePicker_SetOnCloseUp", 0),
		/*33*/ imports.NewTable("DateTimePicker_SetOnContextPopup", 0),
		/*34*/ imports.NewTable("DateTimePicker_SetOnDblClick", 0),
		/*35*/ imports.NewTable("DateTimePicker_SetOnDropDown", 0),
		/*36*/ imports.NewTable("DateTimePicker_SetOnEditingDone", 0),
		/*37*/ imports.NewTable("DateTimePicker_SetOnMouseDown", 0),
		/*38*/ imports.NewTable("DateTimePicker_SetOnMouseEnter", 0),
		/*39*/ imports.NewTable("DateTimePicker_SetOnMouseLeave", 0),
		/*40*/ imports.NewTable("DateTimePicker_SetOnMouseMove", 0),
		/*41*/ imports.NewTable("DateTimePicker_SetOnMouseUp", 0),
		/*42*/ imports.NewTable("DateTimePicker_SetOnMouseWheel", 0),
		/*43*/ imports.NewTable("DateTimePicker_SetOnMouseWheelDown", 0),
		/*44*/ imports.NewTable("DateTimePicker_SetOnMouseWheelUp", 0),
		/*45*/ imports.NewTable("DateTimePicker_ShowCheckBox", 0),
		/*46*/ imports.NewTable("DateTimePicker_ShowMonthNames", 0),
		/*47*/ imports.NewTable("DateTimePicker_TextForNullDate", 0),
		/*48*/ imports.NewTable("DateTimePicker_Time", 0),
		/*49*/ imports.NewTable("DateTimePicker_TimeDisplay", 0),
		/*50*/ imports.NewTable("DateTimePicker_TimeFormat", 0),
		/*51*/ imports.NewTable("DateTimePicker_TimeSeparator", 0),
		/*52*/ imports.NewTable("DateTimePicker_TrailingSeparator", 0),
		/*53*/ imports.NewTable("DateTimePicker_UseDefaultSeparators", 0),
	}
)

func dateTimePickerImportAPI() *imports.Imports {
	if dateTimePickerImport == nil {
		dateTimePickerImport = NewDefaultImports()
		dateTimePickerImport.SetImportTable(dateTimePickerImportTables)
		dateTimePickerImportTables = nil
	}
	return dateTimePickerImport
}
