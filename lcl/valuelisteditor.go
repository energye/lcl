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

// IValueListEditor Parent: ICustomStringGrid
type IValueListEditor interface {
	ICustomStringGrid
	FindRow(keyName string, outRow *int32) bool                            // function
	InsertRow(keyName string, value string, append bool) int32             // function
	IsEmptyRowToBool() bool                                                // function
	IsEmptyRowWithInt(row int32) bool                                      // function
	RestoreCurrentRow() bool                                               // function
	Clear()                                                                // procedure
	DeleteColRowWithBoolInt(isColumn bool, index int32)                    // procedure
	InsertColRowWithBoolInt(isColumn bool, index int32)                    // procedure
	MoveColRowWithBoolIntX2(isColumn bool, fromIndex int32, toIndex int32) // procedure
	SortWithIntX3(index int32, indxFrom int32, indxTo int32)               // procedure
	SortWithVleSortCol(col types.TVleSortCol)                              // procedure
	Modified() bool                                                        // property Modified Getter
	SetModified(value bool)                                                // property Modified Setter
	Keys(index int32) string                                               // property Keys Getter
	SetKeys(index int32, value string)                                     // property Keys Setter
	Values(key string) string                                              // property Values Getter
	SetValues(key string, value string)                                    // property Values Setter
	ItemProps(keyOrIndex types.Variant) IItemProp                          // property ItemProps Getter
	SetItemProps(keyOrIndex types.Variant, value IItemProp)                // property ItemProps Setter
	AlternateColor() types.TColor                                          // property AlternateColor Getter
	SetAlternateColor(value types.TColor)                                  // property AlternateColor Setter
	AutoEdit() bool                                                        // property AutoEdit Getter
	SetAutoEdit(value bool)                                                // property AutoEdit Setter
	DragCursor() types.TCursor                                             // property DragCursor Getter
	SetDragCursor(value types.TCursor)                                     // property DragCursor Setter
	DragKind() types.TDragKind                                             // property DragKind Getter
	SetDragKind(value types.TDragKind)                                     // property DragKind Setter
	DragMode() types.TDragMode                                             // property DragMode Getter
	SetDragMode(value types.TDragMode)                                     // property DragMode Setter
	HeaderHotZones() types.TGridZoneSet                                    // property HeaderHotZones Getter
	SetHeaderHotZones(value types.TGridZoneSet)                            // property HeaderHotZones Setter
	HeaderPushZones() types.TGridZoneSet                                   // property HeaderPushZones Getter
	SetHeaderPushZones(value types.TGridZoneSet)                           // property HeaderPushZones Setter
	MouseWheelOption() types.TMouseWheelOption                             // property MouseWheelOption Getter
	SetMouseWheelOption(value types.TMouseWheelOption)                     // property MouseWheelOption Setter
	ParentColor() bool                                                     // property ParentColor Getter
	SetParentColor(value bool)                                             // property ParentColor Setter
	ParentFont() bool                                                      // property ParentFont Getter
	SetParentFont(value bool)                                              // property ParentFont Setter
	TitleFont() IFont                                                      // property TitleFont Getter
	SetTitleFont(value IFont)                                              // property TitleFont Setter
	TitleImageList() ICustomImageList                                      // property TitleImageList Getter
	SetTitleImageList(value ICustomImageList)                              // property TitleImageList Setter
	TitleStyle() types.TTitleStyle                                         // property TitleStyle Getter
	SetTitleStyle(value types.TTitleStyle)                                 // property TitleStyle Setter
	// DisplayOptions
	//  Compatible with Delphi TValueListEditor:
	DisplayOptions() types.TDisplayOptions               // property DisplayOptions Getter
	SetDisplayOptions(value types.TDisplayOptions)       // property DisplayOptions Setter
	DropDownRows() int32                                 // property DropDownRows Getter
	SetDropDownRows(value int32)                         // property DropDownRows Setter
	KeyOptions() types.TKeyOptions                       // property KeyOptions Getter
	SetKeyOptions(value types.TKeyOptions)               // property KeyOptions Setter
	Strings() IValueListStrings                          // property Strings Getter
	SetStrings(value IValueListStrings)                  // property Strings Setter
	TitleCaptions() IStrings                             // property TitleCaptions Getter
	SetTitleCaptions(value IStrings)                     // property TitleCaptions Setter
	SetOnCheckboxToggled(fn TToggledCheckboxEvent)       // property event
	SetOnEditingDone(fn TNotifyEvent)                    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)             // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)      // property event
	SetOnUserCheckboxBitmap(fn TUserCheckboxBitmapEvent) // property event
	SetOnGetPickList(fn TGetPickListEvent)               // property event
	SetOnStringsChange(fn TNotifyEvent)                  // property event
	SetOnStringsChanging(fn TNotifyEvent)                // property event
	SetOnValidate(fn TOnValidateEvent)                   // property event
}

type TValueListEditor struct {
	TCustomStringGrid
}

func (m *TValueListEditor) FindRow(keyName string, outRow *int32) bool {
	if !m.IsValid() {
		return false
	}
	var rowPtr uintptr
	r := valueListEditorAPI().SysCallN(1, m.Instance(), api.PasStr(keyName), uintptr(base.UnsafePointer(&rowPtr)))
	*outRow = int32(rowPtr)
	return api.GoBool(r)
}

func (m *TValueListEditor) InsertRow(keyName string, value string, append bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(2, m.Instance(), api.PasStr(keyName), api.PasStr(value), api.PasBool(append))
	return int32(r)
}

func (m *TValueListEditor) IsEmptyRowToBool() bool {
	if !m.IsValid() {
		return false
	}
	r := valueListEditorAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TValueListEditor) IsEmptyRowWithInt(row int32) bool {
	if !m.IsValid() {
		return false
	}
	r := valueListEditorAPI().SysCallN(4, m.Instance(), uintptr(row))
	return api.GoBool(r)
}

func (m *TValueListEditor) RestoreCurrentRow() bool {
	if !m.IsValid() {
		return false
	}
	r := valueListEditorAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TValueListEditor) Clear() {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(6, m.Instance())
}

func (m *TValueListEditor) DeleteColRowWithBoolInt(isColumn bool, index int32) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(7, m.Instance(), api.PasBool(isColumn), uintptr(index))
}

func (m *TValueListEditor) InsertColRowWithBoolInt(isColumn bool, index int32) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(8, m.Instance(), api.PasBool(isColumn), uintptr(index))
}

func (m *TValueListEditor) MoveColRowWithBoolIntX2(isColumn bool, fromIndex int32, toIndex int32) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(9, m.Instance(), api.PasBool(isColumn), uintptr(fromIndex), uintptr(toIndex))
}

func (m *TValueListEditor) SortWithIntX3(index int32, indxFrom int32, indxTo int32) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(10, m.Instance(), uintptr(index), uintptr(indxFrom), uintptr(indxTo))
}

func (m *TValueListEditor) SortWithVleSortCol(col types.TVleSortCol) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(11, m.Instance(), uintptr(col))
}

func (m *TValueListEditor) Modified() bool {
	if !m.IsValid() {
		return false
	}
	r := valueListEditorAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TValueListEditor) SetModified(value bool) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TValueListEditor) Keys(index int32) string {
	if !m.IsValid() {
		return ""
	}
	r := valueListEditorAPI().SysCallN(13, 0, m.Instance(), uintptr(index))
	return api.GoStr(r)
}

func (m *TValueListEditor) SetKeys(index int32, value string) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(13, 1, m.Instance(), uintptr(index), api.PasStr(value))
}

func (m *TValueListEditor) Values(key string) string {
	if !m.IsValid() {
		return ""
	}
	r := valueListEditorAPI().SysCallN(14, 0, m.Instance(), api.PasStr(key))
	return api.GoStr(r)
}

func (m *TValueListEditor) SetValues(key string, value string) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(14, 1, m.Instance(), api.PasStr(key), api.PasStr(value))
}

func (m *TValueListEditor) ItemProps(keyOrIndex types.Variant) IItemProp {
	if !m.IsValid() {
		return nil
	}
	r := valueListEditorAPI().SysCallN(15, 0, m.Instance(), uintptr(keyOrIndex))
	return AsItemProp(r)
}

func (m *TValueListEditor) SetItemProps(keyOrIndex types.Variant, value IItemProp) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(15, 1, m.Instance(), uintptr(keyOrIndex), base.GetObjectUintptr(value))
}

func (m *TValueListEditor) AlternateColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(16, 0, m.Instance())
	return types.TColor(r)
}

func (m *TValueListEditor) SetAlternateColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(16, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) AutoEdit() bool {
	if !m.IsValid() {
		return false
	}
	r := valueListEditorAPI().SysCallN(17, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TValueListEditor) SetAutoEdit(value bool) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(17, 1, m.Instance(), api.PasBool(value))
}

func (m *TValueListEditor) DragCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(18, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TValueListEditor) SetDragCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) DragKind() types.TDragKind {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(19, 0, m.Instance())
	return types.TDragKind(r)
}

func (m *TValueListEditor) SetDragKind(value types.TDragKind) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) DragMode() types.TDragMode {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(20, 0, m.Instance())
	return types.TDragMode(r)
}

func (m *TValueListEditor) SetDragMode(value types.TDragMode) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(20, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) HeaderHotZones() types.TGridZoneSet {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(21, 0, m.Instance())
	return types.TGridZoneSet(r)
}

func (m *TValueListEditor) SetHeaderHotZones(value types.TGridZoneSet) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) HeaderPushZones() types.TGridZoneSet {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(22, 0, m.Instance())
	return types.TGridZoneSet(r)
}

func (m *TValueListEditor) SetHeaderPushZones(value types.TGridZoneSet) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) MouseWheelOption() types.TMouseWheelOption {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(23, 0, m.Instance())
	return types.TMouseWheelOption(r)
}

func (m *TValueListEditor) SetMouseWheelOption(value types.TMouseWheelOption) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := valueListEditorAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TValueListEditor) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TValueListEditor) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := valueListEditorAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TValueListEditor) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TValueListEditor) TitleFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := valueListEditorAPI().SysCallN(26, 0, m.Instance())
	return AsFont(r)
}

func (m *TValueListEditor) SetTitleFont(value IFont) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(26, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TValueListEditor) TitleImageList() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := valueListEditorAPI().SysCallN(27, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TValueListEditor) SetTitleImageList(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(27, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TValueListEditor) TitleStyle() types.TTitleStyle {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(28, 0, m.Instance())
	return types.TTitleStyle(r)
}

func (m *TValueListEditor) SetTitleStyle(value types.TTitleStyle) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) DisplayOptions() types.TDisplayOptions {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(29, 0, m.Instance())
	return types.TDisplayOptions(r)
}

func (m *TValueListEditor) SetDisplayOptions(value types.TDisplayOptions) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) DropDownRows() int32 {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(30, 0, m.Instance())
	return int32(r)
}

func (m *TValueListEditor) SetDropDownRows(value int32) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) KeyOptions() types.TKeyOptions {
	if !m.IsValid() {
		return 0
	}
	r := valueListEditorAPI().SysCallN(31, 0, m.Instance())
	return types.TKeyOptions(r)
}

func (m *TValueListEditor) SetKeyOptions(value types.TKeyOptions) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TValueListEditor) Strings() IValueListStrings {
	if !m.IsValid() {
		return nil
	}
	r := valueListEditorAPI().SysCallN(32, 0, m.Instance())
	return AsValueListStrings(r)
}

func (m *TValueListEditor) SetStrings(value IValueListStrings) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(32, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TValueListEditor) TitleCaptions() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := valueListEditorAPI().SysCallN(33, 0, m.Instance())
	return AsStrings(r)
}

func (m *TValueListEditor) SetTitleCaptions(value IStrings) {
	if !m.IsValid() {
		return
	}
	valueListEditorAPI().SysCallN(33, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TValueListEditor) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTToggledCheckboxEvent(fn)
	base.SetEvent(m, 34, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TValueListEditor) SetOnEditingDone(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 35, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TValueListEditor) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 36, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TValueListEditor) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 37, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TValueListEditor) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 38, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TValueListEditor) SetOnUserCheckboxBitmap(fn TUserCheckboxBitmapEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTUserCheckboxBitmapEvent(fn)
	base.SetEvent(m, 39, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TValueListEditor) SetOnGetPickList(fn TGetPickListEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetPickListEvent(fn)
	base.SetEvent(m, 40, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TValueListEditor) SetOnStringsChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 41, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TValueListEditor) SetOnStringsChanging(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 42, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

func (m *TValueListEditor) SetOnValidate(fn TOnValidateEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnValidateEvent(fn)
	base.SetEvent(m, 43, valueListEditorAPI(), api.MakeEventDataPtr(cb))
}

// NewValueListEditor class constructor
func NewValueListEditor(owner IComponent) IValueListEditor {
	r := valueListEditorAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsValueListEditor(r)
}

func TValueListEditorClass() types.TClass {
	r := valueListEditorAPI().SysCallN(44)
	return types.TClass(r)
}

var (
	valueListEditorOnce   base.Once
	valueListEditorImport *imports.Imports = nil
)

func valueListEditorAPI() *imports.Imports {
	valueListEditorOnce.Do(func() {
		valueListEditorImport = api.NewDefaultImports()
		valueListEditorImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TValueListEditor_Create", 0), // constructor NewValueListEditor
			/* 1 */ imports.NewTable("TValueListEditor_FindRow", 0), // function FindRow
			/* 2 */ imports.NewTable("TValueListEditor_InsertRow", 0), // function InsertRow
			/* 3 */ imports.NewTable("TValueListEditor_IsEmptyRowToBool", 0), // function IsEmptyRowToBool
			/* 4 */ imports.NewTable("TValueListEditor_IsEmptyRowWithInt", 0), // function IsEmptyRowWithInt
			/* 5 */ imports.NewTable("TValueListEditor_RestoreCurrentRow", 0), // function RestoreCurrentRow
			/* 6 */ imports.NewTable("TValueListEditor_Clear", 0), // procedure Clear
			/* 7 */ imports.NewTable("TValueListEditor_DeleteColRowWithBoolInt", 0), // procedure DeleteColRowWithBoolInt
			/* 8 */ imports.NewTable("TValueListEditor_InsertColRowWithBoolInt", 0), // procedure InsertColRowWithBoolInt
			/* 9 */ imports.NewTable("TValueListEditor_MoveColRowWithBoolIntX2", 0), // procedure MoveColRowWithBoolIntX2
			/* 10 */ imports.NewTable("TValueListEditor_SortWithIntX3", 0), // procedure SortWithIntX3
			/* 11 */ imports.NewTable("TValueListEditor_SortWithVleSortCol", 0), // procedure SortWithVleSortCol
			/* 12 */ imports.NewTable("TValueListEditor_Modified", 0), // property Modified
			/* 13 */ imports.NewTable("TValueListEditor_Keys", 0), // property Keys
			/* 14 */ imports.NewTable("TValueListEditor_Values", 0), // property Values
			/* 15 */ imports.NewTable("TValueListEditor_ItemProps", 0), // property ItemProps
			/* 16 */ imports.NewTable("TValueListEditor_AlternateColor", 0), // property AlternateColor
			/* 17 */ imports.NewTable("TValueListEditor_AutoEdit", 0), // property AutoEdit
			/* 18 */ imports.NewTable("TValueListEditor_DragCursor", 0), // property DragCursor
			/* 19 */ imports.NewTable("TValueListEditor_DragKind", 0), // property DragKind
			/* 20 */ imports.NewTable("TValueListEditor_DragMode", 0), // property DragMode
			/* 21 */ imports.NewTable("TValueListEditor_HeaderHotZones", 0), // property HeaderHotZones
			/* 22 */ imports.NewTable("TValueListEditor_HeaderPushZones", 0), // property HeaderPushZones
			/* 23 */ imports.NewTable("TValueListEditor_MouseWheelOption", 0), // property MouseWheelOption
			/* 24 */ imports.NewTable("TValueListEditor_ParentColor", 0), // property ParentColor
			/* 25 */ imports.NewTable("TValueListEditor_ParentFont", 0), // property ParentFont
			/* 26 */ imports.NewTable("TValueListEditor_TitleFont", 0), // property TitleFont
			/* 27 */ imports.NewTable("TValueListEditor_TitleImageList", 0), // property TitleImageList
			/* 28 */ imports.NewTable("TValueListEditor_TitleStyle", 0), // property TitleStyle
			/* 29 */ imports.NewTable("TValueListEditor_DisplayOptions", 0), // property DisplayOptions
			/* 30 */ imports.NewTable("TValueListEditor_DropDownRows", 0), // property DropDownRows
			/* 31 */ imports.NewTable("TValueListEditor_KeyOptions", 0), // property KeyOptions
			/* 32 */ imports.NewTable("TValueListEditor_Strings", 0), // property Strings
			/* 33 */ imports.NewTable("TValueListEditor_TitleCaptions", 0), // property TitleCaptions
			/* 34 */ imports.NewTable("TValueListEditor_OnCheckboxToggled", 0), // event OnCheckboxToggled
			/* 35 */ imports.NewTable("TValueListEditor_OnEditingDone", 0), // event OnEditingDone
			/* 36 */ imports.NewTable("TValueListEditor_OnMouseWheelHorz", 0), // event OnMouseWheelHorz
			/* 37 */ imports.NewTable("TValueListEditor_OnMouseWheelLeft", 0), // event OnMouseWheelLeft
			/* 38 */ imports.NewTable("TValueListEditor_OnMouseWheelRight", 0), // event OnMouseWheelRight
			/* 39 */ imports.NewTable("TValueListEditor_OnUserCheckboxBitmap", 0), // event OnUserCheckboxBitmap
			/* 40 */ imports.NewTable("TValueListEditor_OnGetPickList", 0), // event OnGetPickList
			/* 41 */ imports.NewTable("TValueListEditor_OnStringsChange", 0), // event OnStringsChange
			/* 42 */ imports.NewTable("TValueListEditor_OnStringsChanging", 0), // event OnStringsChanging
			/* 43 */ imports.NewTable("TValueListEditor_OnValidate", 0), // event OnValidate
			/* 44 */ imports.NewTable("TValueListEditor_TClass", 0), // function TValueListEditorClass
		}
	})
	return valueListEditorImport
}
