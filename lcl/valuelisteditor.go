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
	r1 := valueListEditorImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsValueListEditor(r1)
}

func (m *TValueListEditor) Modified() bool {
	r1 := valueListEditorImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetModified(AValue bool) {
	valueListEditorImportAPI().SysCallN(18, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) Keys(Index int32) string {
	r1 := valueListEditorImportAPI().SysCallN(17, 0, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TValueListEditor) SetKeys(Index int32, AValue string) {
	valueListEditorImportAPI().SysCallN(17, 1, m.Instance(), uintptr(Index), PascalStr(AValue))
}

func (m *TValueListEditor) Values(Key string) string {
	r1 := valueListEditorImportAPI().SysCallN(40, 0, m.Instance(), PascalStr(Key))
	return GoStr(r1)
}

func (m *TValueListEditor) SetValues(Key string, AValue string) {
	valueListEditorImportAPI().SysCallN(40, 1, m.Instance(), PascalStr(Key), PascalStr(AValue))
}

func (m *TValueListEditor) ItemProps(AKeyOrIndex uintptr) IItemProp {
	r1 := valueListEditorImportAPI().SysCallN(15, 0, m.Instance(), uintptr(AKeyOrIndex))
	return AsItemProp(r1)
}

func (m *TValueListEditor) SetItemProps(AKeyOrIndex uintptr, AValue IItemProp) {
	valueListEditorImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AKeyOrIndex), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) AlternateColor() TColor {
	r1 := valueListEditorImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TValueListEditor) SetAlternateColor(AValue TColor) {
	valueListEditorImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) AutoEdit() bool {
	r1 := valueListEditorImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetAutoEdit(AValue bool) {
	valueListEditorImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) DragCursor() TCursor {
	r1 := valueListEditorImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TValueListEditor) SetDragCursor(AValue TCursor) {
	valueListEditorImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DragKind() TDragKind {
	r1 := valueListEditorImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TValueListEditor) SetDragKind(AValue TDragKind) {
	valueListEditorImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DragMode() TDragMode {
	r1 := valueListEditorImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TValueListEditor) SetDragMode(AValue TDragMode) {
	valueListEditorImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) HeaderHotZones() TGridZoneSet {
	r1 := valueListEditorImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TValueListEditor) SetHeaderHotZones(AValue TGridZoneSet) {
	valueListEditorImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) HeaderPushZones() TGridZoneSet {
	r1 := valueListEditorImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TValueListEditor) SetHeaderPushZones(AValue TGridZoneSet) {
	valueListEditorImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) MouseWheelOption() TMouseWheelOption {
	r1 := valueListEditorImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TMouseWheelOption(r1)
}

func (m *TValueListEditor) SetMouseWheelOption(AValue TMouseWheelOption) {
	valueListEditorImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) ParentColor() bool {
	r1 := valueListEditorImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetParentColor(AValue bool) {
	valueListEditorImportAPI().SysCallN(20, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) ParentFont() bool {
	r1 := valueListEditorImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TValueListEditor) SetParentFont(AValue bool) {
	valueListEditorImportAPI().SysCallN(21, 1, m.Instance(), PascalBool(AValue))
}

func (m *TValueListEditor) TitleFont() IFont {
	r1 := valueListEditorImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TValueListEditor) SetTitleFont(AValue IFont) {
	valueListEditorImportAPI().SysCallN(37, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) TitleImageList() IImageList {
	r1 := valueListEditorImportAPI().SysCallN(38, 0, m.Instance(), 0)
	return AsImageList(r1)
}

func (m *TValueListEditor) SetTitleImageList(AValue IImageList) {
	valueListEditorImportAPI().SysCallN(38, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) TitleStyle() TTitleStyle {
	r1 := valueListEditorImportAPI().SysCallN(39, 0, m.Instance(), 0)
	return TTitleStyle(r1)
}

func (m *TValueListEditor) SetTitleStyle(AValue TTitleStyle) {
	valueListEditorImportAPI().SysCallN(39, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DisplayOptions() TDisplayOptions {
	r1 := valueListEditorImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDisplayOptions(r1)
}

func (m *TValueListEditor) SetDisplayOptions(AValue TDisplayOptions) {
	valueListEditorImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) DropDownRows() int32 {
	r1 := valueListEditorImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TValueListEditor) SetDropDownRows(AValue int32) {
	valueListEditorImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) KeyOptions() TKeyOptions {
	r1 := valueListEditorImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return TKeyOptions(r1)
}

func (m *TValueListEditor) SetKeyOptions(AValue TKeyOptions) {
	valueListEditorImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TValueListEditor) Strings() IValueListStrings {
	r1 := valueListEditorImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return AsValueListStrings(r1)
}

func (m *TValueListEditor) SetStrings(AValue IValueListStrings) {
	valueListEditorImportAPI().SysCallN(35, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) TitleCaptions() IStrings {
	r1 := valueListEditorImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TValueListEditor) SetTitleCaptions(AValue IStrings) {
	valueListEditorImportAPI().SysCallN(36, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TValueListEditor) FindRow(KeyName string, OutRow *int32) bool {
	var result1 uintptr
	r1 := valueListEditorImportAPI().SysCallN(9, m.Instance(), PascalStr(KeyName), uintptr(unsafePointer(&result1)))
	*OutRow = int32(result1)
	return GoBool(r1)
}

func (m *TValueListEditor) InsertRow(KeyName, Value string, Append bool) int32 {
	r1 := valueListEditorImportAPI().SysCallN(12, m.Instance(), PascalStr(KeyName), PascalStr(Value), PascalBool(Append))
	return int32(r1)
}

func (m *TValueListEditor) IsEmptyRow() bool {
	r1 := valueListEditorImportAPI().SysCallN(13, m.Instance())
	return GoBool(r1)
}

func (m *TValueListEditor) IsEmptyRow1(aRow int32) bool {
	r1 := valueListEditorImportAPI().SysCallN(14, m.Instance(), uintptr(aRow))
	return GoBool(r1)
}

func (m *TValueListEditor) RestoreCurrentRow() bool {
	r1 := valueListEditorImportAPI().SysCallN(22, m.Instance())
	return GoBool(r1)
}

func ValueListEditorClass() TClass {
	ret := valueListEditorImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TValueListEditor) Sort(Index, IndxFrom, IndxTo int32) {
	valueListEditorImportAPI().SysCallN(33, m.Instance(), uintptr(Index), uintptr(IndxFrom), uintptr(IndxTo))
}

func (m *TValueListEditor) Sort1(ACol TVleSortCol) {
	valueListEditorImportAPI().SysCallN(34, m.Instance(), uintptr(ACol))
}

func (m *TValueListEditor) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if m.checkboxToggledPtr != 0 {
		RemoveEventElement(m.checkboxToggledPtr)
	}
	m.checkboxToggledPtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(23, m.Instance(), m.checkboxToggledPtr)
}

func (m *TValueListEditor) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(24, m.Instance(), m.editingDonePtr)
}

func (m *TValueListEditor) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(26, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TValueListEditor) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(27, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TValueListEditor) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(28, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TValueListEditor) SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) {
	if m.userCheckboxBitmapPtr != 0 {
		RemoveEventElement(m.userCheckboxBitmapPtr)
	}
	m.userCheckboxBitmapPtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(31, m.Instance(), m.userCheckboxBitmapPtr)
}

func (m *TValueListEditor) SetOnGetPickList(fn TGetPickListEvent) {
	if m.getPickListPtr != 0 {
		RemoveEventElement(m.getPickListPtr)
	}
	m.getPickListPtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(25, m.Instance(), m.getPickListPtr)
}

func (m *TValueListEditor) SetOnStringsChange(fn TNotifyEvent) {
	if m.stringsChangePtr != 0 {
		RemoveEventElement(m.stringsChangePtr)
	}
	m.stringsChangePtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(29, m.Instance(), m.stringsChangePtr)
}

func (m *TValueListEditor) SetOnStringsChanging(fn TNotifyEvent) {
	if m.stringsChangingPtr != 0 {
		RemoveEventElement(m.stringsChangingPtr)
	}
	m.stringsChangingPtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(30, m.Instance(), m.stringsChangingPtr)
}

func (m *TValueListEditor) SetOnValidate(fn TOnValidateEvent) {
	if m.validatePtr != 0 {
		RemoveEventElement(m.validatePtr)
	}
	m.validatePtr = MakeEventDataPtr(fn)
	valueListEditorImportAPI().SysCallN(32, m.Instance(), m.validatePtr)
}

var (
	valueListEditorImport       *imports.Imports = nil
	valueListEditorImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ValueListEditor_AlternateColor", 0),
		/*1*/ imports.NewTable("ValueListEditor_AutoEdit", 0),
		/*2*/ imports.NewTable("ValueListEditor_Class", 0),
		/*3*/ imports.NewTable("ValueListEditor_Create", 0),
		/*4*/ imports.NewTable("ValueListEditor_DisplayOptions", 0),
		/*5*/ imports.NewTable("ValueListEditor_DragCursor", 0),
		/*6*/ imports.NewTable("ValueListEditor_DragKind", 0),
		/*7*/ imports.NewTable("ValueListEditor_DragMode", 0),
		/*8*/ imports.NewTable("ValueListEditor_DropDownRows", 0),
		/*9*/ imports.NewTable("ValueListEditor_FindRow", 0),
		/*10*/ imports.NewTable("ValueListEditor_HeaderHotZones", 0),
		/*11*/ imports.NewTable("ValueListEditor_HeaderPushZones", 0),
		/*12*/ imports.NewTable("ValueListEditor_InsertRow", 0),
		/*13*/ imports.NewTable("ValueListEditor_IsEmptyRow", 0),
		/*14*/ imports.NewTable("ValueListEditor_IsEmptyRow1", 0),
		/*15*/ imports.NewTable("ValueListEditor_ItemProps", 0),
		/*16*/ imports.NewTable("ValueListEditor_KeyOptions", 0),
		/*17*/ imports.NewTable("ValueListEditor_Keys", 0),
		/*18*/ imports.NewTable("ValueListEditor_Modified", 0),
		/*19*/ imports.NewTable("ValueListEditor_MouseWheelOption", 0),
		/*20*/ imports.NewTable("ValueListEditor_ParentColor", 0),
		/*21*/ imports.NewTable("ValueListEditor_ParentFont", 0),
		/*22*/ imports.NewTable("ValueListEditor_RestoreCurrentRow", 0),
		/*23*/ imports.NewTable("ValueListEditor_SetOnCheckboxToggled", 0),
		/*24*/ imports.NewTable("ValueListEditor_SetOnEditingDone", 0),
		/*25*/ imports.NewTable("ValueListEditor_SetOnGetPickList", 0),
		/*26*/ imports.NewTable("ValueListEditor_SetOnMouseWheelHorz", 0),
		/*27*/ imports.NewTable("ValueListEditor_SetOnMouseWheelLeft", 0),
		/*28*/ imports.NewTable("ValueListEditor_SetOnMouseWheelRight", 0),
		/*29*/ imports.NewTable("ValueListEditor_SetOnStringsChange", 0),
		/*30*/ imports.NewTable("ValueListEditor_SetOnStringsChanging", 0),
		/*31*/ imports.NewTable("ValueListEditor_SetOnUserCheckboxBitmap", 0),
		/*32*/ imports.NewTable("ValueListEditor_SetOnValidate", 0),
		/*33*/ imports.NewTable("ValueListEditor_Sort", 0),
		/*34*/ imports.NewTable("ValueListEditor_Sort1", 0),
		/*35*/ imports.NewTable("ValueListEditor_Strings", 0),
		/*36*/ imports.NewTable("ValueListEditor_TitleCaptions", 0),
		/*37*/ imports.NewTable("ValueListEditor_TitleFont", 0),
		/*38*/ imports.NewTable("ValueListEditor_TitleImageList", 0),
		/*39*/ imports.NewTable("ValueListEditor_TitleStyle", 0),
		/*40*/ imports.NewTable("ValueListEditor_Values", 0),
	}
)

func valueListEditorImportAPI() *imports.Imports {
	if valueListEditorImport == nil {
		valueListEditorImport = NewDefaultImports()
		valueListEditorImport.SetImportTable(valueListEditorImportTables)
		valueListEditorImportTables = nil
	}
	return valueListEditorImport
}
