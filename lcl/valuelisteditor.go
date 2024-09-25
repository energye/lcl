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

// IValueListEditor Parent: ICustomStringGrid
type IValueListEditor interface {
	ICustomStringGrid
	Modified() bool                                      // property
	SetModified(AValue bool)                             // property
	Keys(Index int32) string                             // property
	SetKeys(Index int32, AValue string)                  // property
	Values(Key string) string                            // property
	SetValues(Key string, AValue string)                 // property
	ItemProps(AKeyOrIndex uintptr) IItemProp             // property
	SetItemProps(AKeyOrIndex uintptr, AValue IItemProp)  // property
	AlternateColor() TColor                              // property
	SetAlternateColor(AValue TColor)                     // property
	AutoEdit() bool                                      // property
	SetAutoEdit(AValue bool)                             // property
	DragCursor() TCursor                                 // property
	SetDragCursor(AValue TCursor)                        // property
	DragKind() TDragKind                                 // property
	SetDragKind(AValue TDragKind)                        // property
	DragMode() TDragMode                                 // property
	SetDragMode(AValue TDragMode)                        // property
	HeaderHotZones() TGridZoneSet                        // property
	SetHeaderHotZones(AValue TGridZoneSet)               // property
	HeaderPushZones() TGridZoneSet                       // property
	SetHeaderPushZones(AValue TGridZoneSet)              // property
	MouseWheelOption() TMouseWheelOption                 // property
	SetMouseWheelOption(AValue TMouseWheelOption)        // property
	ParentColor() bool                                   // property
	SetParentColor(AValue bool)                          // property
	ParentFont() bool                                    // property
	SetParentFont(AValue bool)                           // property
	TitleFont() IFont                                    // property
	SetTitleFont(AValue IFont)                           // property
	TitleImageList() IImageList                          // property
	SetTitleImageList(AValue IImageList)                 // property
	TitleStyle() TTitleStyle                             // property
	SetTitleStyle(AValue TTitleStyle)                    // property
	DisplayOptions() TDisplayOptions                     // property
	SetDisplayOptions(AValue TDisplayOptions)            // property
	DropDownRows() int32                                 // property
	SetDropDownRows(AValue int32)                        // property
	KeyOptions() TKeyOptions                             // property
	SetKeyOptions(AValue TKeyOptions)                    // property
	Strings() IValueListStrings                          // property
	SetStrings(AValue IValueListStrings)                 // property
	TitleCaptions() IStrings                             // property
	SetTitleCaptions(AValue IStrings)                    // property
	FindRow(KeyName string, OutRow *int32) bool          // function
	InsertRow(KeyName, Value string, Append bool) int32  // function
	IsEmptyRow() bool                                    // function
	IsEmptyRow1(aRow int32) bool                         // function
	RestoreCurrentRow() bool                             // function
	Sort(Index, IndxFrom, IndxTo int32)                  // procedure
	Sort1(ACol TVleSortCol)                              // procedure
	SetOnCheckboxToggled(fn TToggledCheckboxEvent)       // property event
	SetOnEditingDone(fn TNotifyEvent)                    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)             // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)      // property event
	SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) // property event
	SetOnGetPickList(fn TGetPickListEvent)               // property event
	SetOnStringsChange(fn TNotifyEvent)                  // property event
	SetOnStringsChanging(fn TNotifyEvent)                // property event
	SetOnValidate(fn TOnValidateEvent)                   // property event
}

// TValueListEditor Parent: TCustomStringGrid
type TValueListEditor struct {
	TCustomStringGrid
	checkboxToggledPtr    uintptr
	editingDonePtr        uintptr
	mouseWheelHorzPtr     uintptr
	mouseWheelLeftPtr     uintptr
	mouseWheelRightPtr    uintptr
	userCheckboxBitmapPtr uintptr
	getPickListPtr        uintptr
	stringsChangePtr      uintptr
	stringsChangingPtr    uintptr
	validatePtr           uintptr
}

func NewValueListEditor(AOwner IComponent) IValueListEditor {
	r1 := LCL().SysCallN(5940, GetObjectUintptr(AOwner))
	return AsValueListEditor(r1)
}

func (m *TValueListEditor) Modified() bool {
	r1 := LCL().SysCallN(5955, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetModified(AValue bool) {
	LCL().SysCallN(5955, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) Keys(Index int32) string {
	r1 := LCL().SysCallN(5954, 0, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TValueListEditor) SetKeys(Index int32, AValue string) {
	LCL().SysCallN(5954, 1, m.Instance(), uintptr(Index), PascalStr(AValue))
}

func (m *TValueListEditor) Values(Key string) string {
	r1 := LCL().SysCallN(5977, 0, m.Instance(), PascalStr(Key))
	return GoStr(r1)
}

func (m *TValueListEditor) SetValues(Key string, AValue string) {
	LCL().SysCallN(5977, 1, m.Instance(), PascalStr(Key), PascalStr(AValue))
}

func (m *TValueListEditor) ItemProps(AKeyOrIndex uintptr) IItemProp {
	r1 := LCL().SysCallN(5952, 0, m.Instance(), uintptr(AKeyOrIndex))
	return AsItemProp(r1)
}

func (m *TValueListEditor) SetItemProps(AKeyOrIndex uintptr, AValue IItemProp) {
	LCL().SysCallN(5952, 1, m.Instance(), uintptr(AKeyOrIndex), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) AlternateColor() TColor {
	r1 := LCL().SysCallN(5937, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TValueListEditor) SetAlternateColor(AValue TColor) {
	LCL().SysCallN(5937, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) AutoEdit() bool {
	r1 := LCL().SysCallN(5938, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetAutoEdit(AValue bool) {
	LCL().SysCallN(5938, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) DragCursor() TCursor {
	r1 := LCL().SysCallN(5942, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TValueListEditor) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5942, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DragKind() TDragKind {
	r1 := LCL().SysCallN(5943, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TValueListEditor) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5943, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DragMode() TDragMode {
	r1 := LCL().SysCallN(5944, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TValueListEditor) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5944, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) HeaderHotZones() TGridZoneSet {
	r1 := LCL().SysCallN(5947, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TValueListEditor) SetHeaderHotZones(AValue TGridZoneSet) {
	LCL().SysCallN(5947, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) HeaderPushZones() TGridZoneSet {
	r1 := LCL().SysCallN(5948, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TValueListEditor) SetHeaderPushZones(AValue TGridZoneSet) {
	LCL().SysCallN(5948, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) MouseWheelOption() TMouseWheelOption {
	r1 := LCL().SysCallN(5956, 0, m.Instance(), 0)
	return TMouseWheelOption(r1)
}

func (m *TValueListEditor) SetMouseWheelOption(AValue TMouseWheelOption) {
	LCL().SysCallN(5956, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) ParentColor() bool {
	r1 := LCL().SysCallN(5957, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetParentColor(AValue bool) {
	LCL().SysCallN(5957, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) ParentFont() bool {
	r1 := LCL().SysCallN(5958, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetParentFont(AValue bool) {
	LCL().SysCallN(5958, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) TitleFont() IFont {
	r1 := LCL().SysCallN(5974, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TValueListEditor) SetTitleFont(AValue IFont) {
	LCL().SysCallN(5974, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) TitleImageList() IImageList {
	r1 := LCL().SysCallN(5975, 0, m.Instance(), 0)
	return AsImageList(r1)
}

func (m *TValueListEditor) SetTitleImageList(AValue IImageList) {
	LCL().SysCallN(5975, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) TitleStyle() TTitleStyle {
	r1 := LCL().SysCallN(5976, 0, m.Instance(), 0)
	return TTitleStyle(r1)
}

func (m *TValueListEditor) SetTitleStyle(AValue TTitleStyle) {
	LCL().SysCallN(5976, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DisplayOptions() TDisplayOptions {
	r1 := LCL().SysCallN(5941, 0, m.Instance(), 0)
	return TDisplayOptions(r1)
}

func (m *TValueListEditor) SetDisplayOptions(AValue TDisplayOptions) {
	LCL().SysCallN(5941, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DropDownRows() int32 {
	r1 := LCL().SysCallN(5945, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TValueListEditor) SetDropDownRows(AValue int32) {
	LCL().SysCallN(5945, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) KeyOptions() TKeyOptions {
	r1 := LCL().SysCallN(5953, 0, m.Instance(), 0)
	return TKeyOptions(r1)
}

func (m *TValueListEditor) SetKeyOptions(AValue TKeyOptions) {
	LCL().SysCallN(5953, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) Strings() IValueListStrings {
	r1 := LCL().SysCallN(5972, 0, m.Instance(), 0)
	return AsValueListStrings(r1)
}

func (m *TValueListEditor) SetStrings(AValue IValueListStrings) {
	LCL().SysCallN(5972, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) TitleCaptions() IStrings {
	r1 := LCL().SysCallN(5973, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TValueListEditor) SetTitleCaptions(AValue IStrings) {
	LCL().SysCallN(5973, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) FindRow(KeyName string, OutRow *int32) bool {
	var result1 uintptr
	r1 := LCL().SysCallN(5946, m.Instance(), PascalStr(KeyName), uintptr(unsafePointer(&result1)))
	*OutRow = int32(result1)
	return GoBool(r1)
}

func (m *TValueListEditor) InsertRow(KeyName, Value string, Append bool) int32 {
	r1 := LCL().SysCallN(5949, m.Instance(), PascalStr(KeyName), PascalStr(Value), PascalBool(Append))
	return int32(r1)
}

func (m *TValueListEditor) IsEmptyRow() bool {
	r1 := LCL().SysCallN(5950, m.Instance())
	return GoBool(r1)
}

func (m *TValueListEditor) IsEmptyRow1(aRow int32) bool {
	r1 := LCL().SysCallN(5951, m.Instance(), uintptr(aRow))
	return GoBool(r1)
}

func (m *TValueListEditor) RestoreCurrentRow() bool {
	r1 := LCL().SysCallN(5959, m.Instance())
	return GoBool(r1)
}

func ValueListEditorClass() TClass {
	ret := LCL().SysCallN(5939)
	return TClass(ret)
}

func (m *TValueListEditor) Sort(Index, IndxFrom, IndxTo int32) {
	LCL().SysCallN(5970, m.Instance(), uintptr(Index), uintptr(IndxFrom), uintptr(IndxTo))
}

func (m *TValueListEditor) Sort1(ACol TVleSortCol) {
	LCL().SysCallN(5971, m.Instance(), uintptr(ACol))
}

func (m *TValueListEditor) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if m.checkboxToggledPtr != 0 {
		RemoveEventElement(m.checkboxToggledPtr)
	}
	m.checkboxToggledPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5960, m.Instance(), m.checkboxToggledPtr)
}

func (m *TValueListEditor) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5961, m.Instance(), m.editingDonePtr)
}

func (m *TValueListEditor) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5963, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TValueListEditor) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5964, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TValueListEditor) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5965, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TValueListEditor) SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) {
	if m.userCheckboxBitmapPtr != 0 {
		RemoveEventElement(m.userCheckboxBitmapPtr)
	}
	m.userCheckboxBitmapPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5968, m.Instance(), m.userCheckboxBitmapPtr)
}

func (m *TValueListEditor) SetOnGetPickList(fn TGetPickListEvent) {
	if m.getPickListPtr != 0 {
		RemoveEventElement(m.getPickListPtr)
	}
	m.getPickListPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5962, m.Instance(), m.getPickListPtr)
}

func (m *TValueListEditor) SetOnStringsChange(fn TNotifyEvent) {
	if m.stringsChangePtr != 0 {
		RemoveEventElement(m.stringsChangePtr)
	}
	m.stringsChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5966, m.Instance(), m.stringsChangePtr)
}

func (m *TValueListEditor) SetOnStringsChanging(fn TNotifyEvent) {
	if m.stringsChangingPtr != 0 {
		RemoveEventElement(m.stringsChangingPtr)
	}
	m.stringsChangingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5967, m.Instance(), m.stringsChangingPtr)
}

func (m *TValueListEditor) SetOnValidate(fn TOnValidateEvent) {
	if m.validatePtr != 0 {
		RemoveEventElement(m.validatePtr)
	}
	m.validatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5969, m.Instance(), m.validatePtr)
}
