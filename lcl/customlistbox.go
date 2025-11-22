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

// ICustomListBox Parent: IWinControl
type ICustomListBox interface {
	IWinControl
	GetIndexAtXY(X int32, Y int32) int32              // function
	GetIndexAtY(Y int32) int32                        // function
	GetSelectedText() string                          // function
	ItemAtPos(pos types.TPoint, existing bool) int32  // function
	ItemRect(index int32) types.TRect                 // function
	ItemVisible(index int32) bool                     // function
	ItemFullyVisible(index int32) bool                // function
	AddItem(item string, anObject IObject)            // procedure
	Click()                                           // procedure
	Clear()                                           // procedure
	ClearSelection()                                  // procedure
	LockSelectionChange()                             // procedure
	MakeCurrentVisible()                              // procedure
	MeasureItem(index int32, theHeight *int32)        // procedure
	SelectAll()                                       // procedure
	SelectRange(low int32, high int32, selected bool) // procedure
	DeleteSelected()                                  // procedure
	UnlockSelectionChange()                           // procedure
	// BorderStyle
	//  properties which are not supported by all descendents
	BorderStyle() types.TBorderStyle               // property BorderStyle Getter
	SetBorderStyle(value types.TBorderStyle)       // property BorderStyle Setter
	Canvas() ICanvas                               // property Canvas Getter
	ClickOnSelChange() bool                        // property ClickOnSelChange Getter
	SetClickOnSelChange(value bool)                // property ClickOnSelChange Setter
	Columns() int32                                // property Columns Getter
	SetColumns(value int32)                        // property Columns Setter
	Count() int32                                  // property Count Getter
	ExtendedSelect() bool                          // property ExtendedSelect Getter
	SetExtendedSelect(value bool)                  // property ExtendedSelect Setter
	IntegralHeight() bool                          // property IntegralHeight Getter
	SetIntegralHeight(value bool)                  // property IntegralHeight Setter
	ItemHeight() int32                             // property ItemHeight Getter
	SetItemHeight(value int32)                     // property ItemHeight Setter
	ItemIndex() int32                              // property ItemIndex Getter
	SetItemIndex(value int32)                      // property ItemIndex Setter
	Items() IStrings                               // property Items Getter
	SetItems(value IStrings)                       // property Items Setter
	MultiSelect() bool                             // property MultiSelect Getter
	SetMultiSelect(value bool)                     // property MultiSelect Setter
	Options() types.TListBoxOptions                // property Options Getter
	SetOptions(value types.TListBoxOptions)        // property Options Setter
	ParentColor() bool                             // property ParentColor Getter
	SetParentColor(value bool)                     // property ParentColor Setter
	ParentFont() bool                              // property ParentFont Getter
	SetParentFont(value bool)                      // property ParentFont Setter
	ParentShowHint() bool                          // property ParentShowHint Getter
	SetParentShowHint(value bool)                  // property ParentShowHint Setter
	ScrollWidth() int32                            // property ScrollWidth Getter
	SetScrollWidth(value int32)                    // property ScrollWidth Setter
	SelCount() int32                               // property SelCount Getter
	Selected(index int32) bool                     // property Selected Getter
	SetSelected(index int32, value bool)           // property Selected Setter
	Sorted() bool                                  // property Sorted Getter
	SetSorted(value bool)                          // property Sorted Setter
	Style() types.TListBoxStyle                    // property Style Getter
	SetStyle(value types.TListBoxStyle)            // property Style Setter
	TopIndex() int32                               // property TopIndex Getter
	SetTopIndex(value int32)                       // property TopIndex Setter
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDrawItem(fn TDrawItemEvent)               // property event
	SetOnMeasureItem(fn TMeasureItemEvent)         // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnSelectionChange(fn TSelectionChangeEvent) // property event
}

type TCustomListBox struct {
	TWinControl
}

func (m *TCustomListBox) GetIndexAtXY(X int32, Y int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(1, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r)
}

func (m *TCustomListBox) GetIndexAtY(Y int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(2, m.Instance(), uintptr(Y))
	return int32(r)
}

func (m *TCustomListBox) GetSelectedText() string {
	if !m.IsValid() {
		return ""
	}
	r := customListBoxAPI().SysCallN(3, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomListBox) ItemAtPos(pos types.TPoint, existing bool) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(4, m.Instance(), uintptr(base.UnsafePointer(&pos)), api.PasBool(existing))
	return int32(r)
}

func (m *TCustomListBox) ItemRect(index int32) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(5, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomListBox) ItemVisible(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(6, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomListBox) ItemFullyVisible(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(7, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomListBox) AddItem(item string, anObject IObject) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(8, m.Instance(), api.PasStr(item), base.GetObjectUintptr(anObject))
}

func (m *TCustomListBox) Click() {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(9, m.Instance())
}

func (m *TCustomListBox) Clear() {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(10, m.Instance())
}

func (m *TCustomListBox) ClearSelection() {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(11, m.Instance())
}

func (m *TCustomListBox) LockSelectionChange() {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(12, m.Instance())
}

func (m *TCustomListBox) MakeCurrentVisible() {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(13, m.Instance())
}

func (m *TCustomListBox) MeasureItem(index int32, theHeight *int32) {
	if !m.IsValid() {
		return
	}
	theHeightPtr := uintptr(*theHeight)
	customListBoxAPI().SysCallN(14, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&theHeightPtr)))
	*theHeight = int32(theHeightPtr)
}

func (m *TCustomListBox) SelectAll() {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(15, m.Instance())
}

func (m *TCustomListBox) SelectRange(low int32, high int32, selected bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(16, m.Instance(), uintptr(low), uintptr(high), api.PasBool(selected))
}

func (m *TCustomListBox) DeleteSelected() {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(17, m.Instance())
}

func (m *TCustomListBox) UnlockSelectionChange() {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(18, m.Instance())
}

func (m *TCustomListBox) BorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(19, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TCustomListBox) SetBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListBox) Canvas() ICanvas {
	if !m.IsValid() {
		return nil
	}
	r := customListBoxAPI().SysCallN(20, m.Instance())
	return AsCanvas(r)
}

func (m *TCustomListBox) ClickOnSelChange() bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(21, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListBox) SetClickOnSelChange(value bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(21, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListBox) Columns() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TCustomListBox) SetColumns(value int32) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListBox) Count() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(23, m.Instance())
	return int32(r)
}

func (m *TCustomListBox) ExtendedSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListBox) SetExtendedSelect(value bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListBox) IntegralHeight() bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListBox) SetIntegralHeight(value bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListBox) ItemHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(26, 0, m.Instance())
	return int32(r)
}

func (m *TCustomListBox) SetItemHeight(value int32) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListBox) ItemIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(27, 0, m.Instance())
	return int32(r)
}

func (m *TCustomListBox) SetItemIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListBox) Items() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := customListBoxAPI().SysCallN(28, 0, m.Instance())
	return AsStrings(r)
}

func (m *TCustomListBox) SetItems(value IStrings) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(28, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomListBox) MultiSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(29, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListBox) SetMultiSelect(value bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(29, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListBox) Options() types.TListBoxOptions {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(30, 0, m.Instance())
	return types.TListBoxOptions(r)
}

func (m *TCustomListBox) SetOptions(value types.TListBoxOptions) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListBox) ParentColor() bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(31, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListBox) SetParentColor(value bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(31, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListBox) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(32, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListBox) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(32, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListBox) ParentShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(33, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListBox) SetParentShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(33, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListBox) ScrollWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(34, 0, m.Instance())
	return int32(r)
}

func (m *TCustomListBox) SetScrollWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(34, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListBox) SelCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(35, m.Instance())
	return int32(r)
}

func (m *TCustomListBox) Selected(index int32) bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(36, 0, m.Instance(), uintptr(index))
	return api.GoBool(r)
}

func (m *TCustomListBox) SetSelected(index int32, value bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(36, 1, m.Instance(), uintptr(index), api.PasBool(value))
}

func (m *TCustomListBox) Sorted() bool {
	if !m.IsValid() {
		return false
	}
	r := customListBoxAPI().SysCallN(37, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListBox) SetSorted(value bool) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(37, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListBox) Style() types.TListBoxStyle {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(38, 0, m.Instance())
	return types.TListBoxStyle(r)
}

func (m *TCustomListBox) SetStyle(value types.TListBoxStyle) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(38, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListBox) TopIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListBoxAPI().SysCallN(39, 0, m.Instance())
	return int32(r)
}

func (m *TCustomListBox) SetTopIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customListBoxAPI().SysCallN(39, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListBox) SetOnDblClick(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 40, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnDrawItem(fn TDrawItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDrawItemEvent(fn)
	base.SetEvent(m, 41, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnMeasureItem(fn TMeasureItemEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMeasureItemEvent(fn)
	base.SetEvent(m, 42, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnMouseDown(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 43, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnMouseEnter(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 44, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnMouseLeave(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 45, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseMoveEvent(fn)
	base.SetEvent(m, 46, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnMouseUp(fn TMouseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseEvent(fn)
	base.SetEvent(m, 47, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelEvent(fn)
	base.SetEvent(m, 48, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 49, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTMouseWheelUpDownEvent(fn)
	base.SetEvent(m, 50, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomListBox) SetOnSelectionChange(fn TSelectionChangeEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTSelectionChangeEvent(fn)
	base.SetEvent(m, 51, customListBoxAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomListBox class constructor
func NewCustomListBox(theOwner IComponent) ICustomListBox {
	r := customListBoxAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomListBox(r)
}

func TCustomListBoxClass() types.TClass {
	r := customListBoxAPI().SysCallN(52)
	return types.TClass(r)
}

var (
	customListBoxOnce   base.Once
	customListBoxImport *imports.Imports = nil
)

func customListBoxAPI() *imports.Imports {
	customListBoxOnce.Do(func() {
		customListBoxImport = api.NewDefaultImports()
		customListBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomListBox_Create", 0), // constructor NewCustomListBox
			/* 1 */ imports.NewTable("TCustomListBox_GetIndexAtXY", 0), // function GetIndexAtXY
			/* 2 */ imports.NewTable("TCustomListBox_GetIndexAtY", 0), // function GetIndexAtY
			/* 3 */ imports.NewTable("TCustomListBox_GetSelectedText", 0), // function GetSelectedText
			/* 4 */ imports.NewTable("TCustomListBox_ItemAtPos", 0), // function ItemAtPos
			/* 5 */ imports.NewTable("TCustomListBox_ItemRect", 0), // function ItemRect
			/* 6 */ imports.NewTable("TCustomListBox_ItemVisible", 0), // function ItemVisible
			/* 7 */ imports.NewTable("TCustomListBox_ItemFullyVisible", 0), // function ItemFullyVisible
			/* 8 */ imports.NewTable("TCustomListBox_AddItem", 0), // procedure AddItem
			/* 9 */ imports.NewTable("TCustomListBox_Click", 0), // procedure Click
			/* 10 */ imports.NewTable("TCustomListBox_Clear", 0), // procedure Clear
			/* 11 */ imports.NewTable("TCustomListBox_ClearSelection", 0), // procedure ClearSelection
			/* 12 */ imports.NewTable("TCustomListBox_LockSelectionChange", 0), // procedure LockSelectionChange
			/* 13 */ imports.NewTable("TCustomListBox_MakeCurrentVisible", 0), // procedure MakeCurrentVisible
			/* 14 */ imports.NewTable("TCustomListBox_MeasureItem", 0), // procedure MeasureItem
			/* 15 */ imports.NewTable("TCustomListBox_SelectAll", 0), // procedure SelectAll
			/* 16 */ imports.NewTable("TCustomListBox_SelectRange", 0), // procedure SelectRange
			/* 17 */ imports.NewTable("TCustomListBox_DeleteSelected", 0), // procedure DeleteSelected
			/* 18 */ imports.NewTable("TCustomListBox_UnlockSelectionChange", 0), // procedure UnlockSelectionChange
			/* 19 */ imports.NewTable("TCustomListBox_BorderStyle", 0), // property BorderStyle
			/* 20 */ imports.NewTable("TCustomListBox_Canvas", 0), // property Canvas
			/* 21 */ imports.NewTable("TCustomListBox_ClickOnSelChange", 0), // property ClickOnSelChange
			/* 22 */ imports.NewTable("TCustomListBox_Columns", 0), // property Columns
			/* 23 */ imports.NewTable("TCustomListBox_Count", 0), // property Count
			/* 24 */ imports.NewTable("TCustomListBox_ExtendedSelect", 0), // property ExtendedSelect
			/* 25 */ imports.NewTable("TCustomListBox_IntegralHeight", 0), // property IntegralHeight
			/* 26 */ imports.NewTable("TCustomListBox_ItemHeight", 0), // property ItemHeight
			/* 27 */ imports.NewTable("TCustomListBox_ItemIndex", 0), // property ItemIndex
			/* 28 */ imports.NewTable("TCustomListBox_Items", 0), // property Items
			/* 29 */ imports.NewTable("TCustomListBox_MultiSelect", 0), // property MultiSelect
			/* 30 */ imports.NewTable("TCustomListBox_Options", 0), // property Options
			/* 31 */ imports.NewTable("TCustomListBox_ParentColor", 0), // property ParentColor
			/* 32 */ imports.NewTable("TCustomListBox_ParentFont", 0), // property ParentFont
			/* 33 */ imports.NewTable("TCustomListBox_ParentShowHint", 0), // property ParentShowHint
			/* 34 */ imports.NewTable("TCustomListBox_ScrollWidth", 0), // property ScrollWidth
			/* 35 */ imports.NewTable("TCustomListBox_SelCount", 0), // property SelCount
			/* 36 */ imports.NewTable("TCustomListBox_Selected", 0), // property Selected
			/* 37 */ imports.NewTable("TCustomListBox_Sorted", 0), // property Sorted
			/* 38 */ imports.NewTable("TCustomListBox_Style", 0), // property Style
			/* 39 */ imports.NewTable("TCustomListBox_TopIndex", 0), // property TopIndex
			/* 40 */ imports.NewTable("TCustomListBox_OnDblClick", 0), // event OnDblClick
			/* 41 */ imports.NewTable("TCustomListBox_OnDrawItem", 0), // event OnDrawItem
			/* 42 */ imports.NewTable("TCustomListBox_OnMeasureItem", 0), // event OnMeasureItem
			/* 43 */ imports.NewTable("TCustomListBox_OnMouseDown", 0), // event OnMouseDown
			/* 44 */ imports.NewTable("TCustomListBox_OnMouseEnter", 0), // event OnMouseEnter
			/* 45 */ imports.NewTable("TCustomListBox_OnMouseLeave", 0), // event OnMouseLeave
			/* 46 */ imports.NewTable("TCustomListBox_OnMouseMove", 0), // event OnMouseMove
			/* 47 */ imports.NewTable("TCustomListBox_OnMouseUp", 0), // event OnMouseUp
			/* 48 */ imports.NewTable("TCustomListBox_OnMouseWheel", 0), // event OnMouseWheel
			/* 49 */ imports.NewTable("TCustomListBox_OnMouseWheelDown", 0), // event OnMouseWheelDown
			/* 50 */ imports.NewTable("TCustomListBox_OnMouseWheelUp", 0), // event OnMouseWheelUp
			/* 51 */ imports.NewTable("TCustomListBox_OnSelectionChange", 0), // event OnSelectionChange
			/* 52 */ imports.NewTable("TCustomListBox_TClass", 0), // function TCustomListBoxClass
		}
	})
	return customListBoxImport
}
