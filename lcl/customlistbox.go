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

// ICustomListBox Parent: IWinControl
type ICustomListBox interface {
	IWinControl
	BorderStyle() TBorderStyle                     // property
	SetBorderStyle(AValue TBorderStyle)            // property
	Canvas() ICanvas                               // property
	ClickOnSelChange() bool                        // property
	SetClickOnSelChange(AValue bool)               // property
	Columns() int32                                // property
	SetColumns(AValue int32)                       // property
	Count() int32                                  // property
	ExtendedSelect() bool                          // property
	SetExtendedSelect(AValue bool)                 // property
	IntegralHeight() bool                          // property
	SetIntegralHeight(AValue bool)                 // property
	ItemHeight() int32                             // property
	SetItemHeight(AValue int32)                    // property
	ItemIndex() int32                              // property
	SetItemIndex(AValue int32)                     // property
	Items() IStrings                               // property
	SetItems(AValue IStrings)                      // property
	MultiSelect() bool                             // property
	SetMultiSelect(AValue bool)                    // property
	Options() TListBoxOptions                      // property
	SetOptions(AValue TListBoxOptions)             // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	ScrollWidth() int32                            // property
	SetScrollWidth(AValue int32)                   // property
	SelCount() int32                               // property
	Selected(Index int32) bool                     // property
	SetSelected(Index int32, AValue bool)          // property
	Sorted() bool                                  // property
	SetSorted(AValue bool)                         // property
	Style() TListBoxStyle                          // property
	SetStyle(AValue TListBoxStyle)                 // property
	TopIndex() int32                               // property
	SetTopIndex(AValue int32)                      // property
	GetIndexAtXY(X, Y int32) int32                 // function
	GetIndexAtY(Y int32) int32                     // function
	GetSelectedText() string                       // function
	ItemAtPos(Pos *TPoint, Existing bool) int32    // function
	ItemRect(Index int32) (resultRect TRect)       // function
	ItemVisible(Index int32) bool                  // function
	ItemFullyVisible(Index int32) bool             // function
	AddItem(Item string, AnObject IObject)         // procedure
	Click()                                        // procedure
	Clear()                                        // procedure
	ClearSelection()                               // procedure
	LockSelectionChange()                          // procedure
	MakeCurrentVisible()                           // procedure
	MeasureItem(Index int32, TheHeight *int32)     // procedure
	SelectAll()                                    // procedure
	SelectRange(ALow, AHigh int32, ASelected bool) // procedure
	DeleteSelected()                               // procedure
	UnlockSelectionChange()                        // procedure
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

// TCustomListBox Parent: TWinControl
type TCustomListBox struct {
	TWinControl
	dblClickPtr        uintptr
	drawItemPtr        uintptr
	measureItemPtr     uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	selectionChangePtr uintptr
}

func NewCustomListBox(TheOwner IComponent) ICustomListBox {
	r1 := customListBoxImportAPI().SysCallN(10, GetObjectUintptr(TheOwner))
	return AsCustomListBox(r1)
}

func (m *TCustomListBox) BorderStyle() TBorderStyle {
	r1 := customListBoxImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomListBox) SetBorderStyle(AValue TBorderStyle) {
	customListBoxImportAPI().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) Canvas() ICanvas {
	r1 := customListBoxImportAPI().SysCallN(2, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomListBox) ClickOnSelChange() bool {
	r1 := customListBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetClickOnSelChange(AValue bool) {
	customListBoxImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) Columns() int32 {
	r1 := customListBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetColumns(AValue int32) {
	customListBoxImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) Count() int32 {
	r1 := customListBoxImportAPI().SysCallN(9, m.Instance())
	return int32(r1)
}

func (m *TCustomListBox) ExtendedSelect() bool {
	r1 := customListBoxImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetExtendedSelect(AValue bool) {
	customListBoxImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) IntegralHeight() bool {
	r1 := customListBoxImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetIntegralHeight(AValue bool) {
	customListBoxImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ItemHeight() int32 {
	r1 := customListBoxImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetItemHeight(AValue int32) {
	customListBoxImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) ItemIndex() int32 {
	r1 := customListBoxImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetItemIndex(AValue int32) {
	customListBoxImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) Items() IStrings {
	r1 := customListBoxImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomListBox) SetItems(AValue IStrings) {
	customListBoxImportAPI().SysCallN(23, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListBox) MultiSelect() bool {
	r1 := customListBoxImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetMultiSelect(AValue bool) {
	customListBoxImportAPI().SysCallN(27, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) Options() TListBoxOptions {
	r1 := customListBoxImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return TListBoxOptions(r1)
}

func (m *TCustomListBox) SetOptions(AValue TListBoxOptions) {
	customListBoxImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) ParentColor() bool {
	r1 := customListBoxImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetParentColor(AValue bool) {
	customListBoxImportAPI().SysCallN(29, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ParentFont() bool {
	r1 := customListBoxImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetParentFont(AValue bool) {
	customListBoxImportAPI().SysCallN(30, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ParentShowHint() bool {
	r1 := customListBoxImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetParentShowHint(AValue bool) {
	customListBoxImportAPI().SysCallN(31, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ScrollWidth() int32 {
	r1 := customListBoxImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetScrollWidth(AValue int32) {
	customListBoxImportAPI().SysCallN(32, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) SelCount() int32 {
	r1 := customListBoxImportAPI().SysCallN(33, m.Instance())
	return int32(r1)
}

func (m *TCustomListBox) Selected(Index int32) bool {
	r1 := customListBoxImportAPI().SysCallN(36, 0, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomListBox) SetSelected(Index int32, AValue bool) {
	customListBoxImportAPI().SysCallN(36, 1, m.Instance(), uintptr(Index), PascalBool(AValue))
}

func (m *TCustomListBox) Sorted() bool {
	r1 := customListBoxImportAPI().SysCallN(49, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetSorted(AValue bool) {
	customListBoxImportAPI().SysCallN(49, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) Style() TListBoxStyle {
	r1 := customListBoxImportAPI().SysCallN(50, 0, m.Instance(), 0)
	return TListBoxStyle(r1)
}

func (m *TCustomListBox) SetStyle(AValue TListBoxStyle) {
	customListBoxImportAPI().SysCallN(50, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) TopIndex() int32 {
	r1 := customListBoxImportAPI().SysCallN(51, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetTopIndex(AValue int32) {
	customListBoxImportAPI().SysCallN(51, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) GetIndexAtXY(X, Y int32) int32 {
	r1 := customListBoxImportAPI().SysCallN(13, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TCustomListBox) GetIndexAtY(Y int32) int32 {
	r1 := customListBoxImportAPI().SysCallN(14, m.Instance(), uintptr(Y))
	return int32(r1)
}

func (m *TCustomListBox) GetSelectedText() string {
	r1 := customListBoxImportAPI().SysCallN(15, m.Instance())
	return GoStr(r1)
}

func (m *TCustomListBox) ItemAtPos(Pos *TPoint, Existing bool) int32 {
	r1 := customListBoxImportAPI().SysCallN(17, m.Instance(), uintptr(unsafePointer(Pos)), PascalBool(Existing))
	return int32(r1)
}

func (m *TCustomListBox) ItemRect(Index int32) (resultRect TRect) {
	customListBoxImportAPI().SysCallN(21, m.Instance(), uintptr(Index), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCustomListBox) ItemVisible(Index int32) bool {
	r1 := customListBoxImportAPI().SysCallN(22, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomListBox) ItemFullyVisible(Index int32) bool {
	r1 := customListBoxImportAPI().SysCallN(18, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func CustomListBoxClass() TClass {
	ret := customListBoxImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TCustomListBox) AddItem(Item string, AnObject IObject) {
	customListBoxImportAPI().SysCallN(0, m.Instance(), PascalStr(Item), GetObjectUintptr(AnObject))
}

func (m *TCustomListBox) Click() {
	customListBoxImportAPI().SysCallN(6, m.Instance())
}

func (m *TCustomListBox) Clear() {
	customListBoxImportAPI().SysCallN(4, m.Instance())
}

func (m *TCustomListBox) ClearSelection() {
	customListBoxImportAPI().SysCallN(5, m.Instance())
}

func (m *TCustomListBox) LockSelectionChange() {
	customListBoxImportAPI().SysCallN(24, m.Instance())
}

func (m *TCustomListBox) MakeCurrentVisible() {
	customListBoxImportAPI().SysCallN(25, m.Instance())
}

func (m *TCustomListBox) MeasureItem(Index int32, TheHeight *int32) {
	var result1 uintptr
	customListBoxImportAPI().SysCallN(26, m.Instance(), uintptr(Index), uintptr(unsafePointer(&result1)))
	*TheHeight = int32(result1)
}

func (m *TCustomListBox) SelectAll() {
	customListBoxImportAPI().SysCallN(34, m.Instance())
}

func (m *TCustomListBox) SelectRange(ALow, AHigh int32, ASelected bool) {
	customListBoxImportAPI().SysCallN(35, m.Instance(), uintptr(ALow), uintptr(AHigh), PascalBool(ASelected))
}

func (m *TCustomListBox) DeleteSelected() {
	customListBoxImportAPI().SysCallN(11, m.Instance())
}

func (m *TCustomListBox) UnlockSelectionChange() {
	customListBoxImportAPI().SysCallN(52, m.Instance())
}

func (m *TCustomListBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(37, m.Instance(), m.dblClickPtr)
}

func (m *TCustomListBox) SetOnDrawItem(fn TDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(38, m.Instance(), m.drawItemPtr)
}

func (m *TCustomListBox) SetOnMeasureItem(fn TMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(39, m.Instance(), m.measureItemPtr)
}

func (m *TCustomListBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(40, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomListBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(41, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomListBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(42, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomListBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(43, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomListBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(44, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomListBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(45, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomListBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(46, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomListBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(47, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomListBox) SetOnSelectionChange(fn TSelectionChangeEvent) {
	if m.selectionChangePtr != 0 {
		RemoveEventElement(m.selectionChangePtr)
	}
	m.selectionChangePtr = MakeEventDataPtr(fn)
	customListBoxImportAPI().SysCallN(48, m.Instance(), m.selectionChangePtr)
}

var (
	customListBoxImport       *imports.Imports = nil
	customListBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomListBox_AddItem", 0),
		/*1*/ imports.NewTable("CustomListBox_BorderStyle", 0),
		/*2*/ imports.NewTable("CustomListBox_Canvas", 0),
		/*3*/ imports.NewTable("CustomListBox_Class", 0),
		/*4*/ imports.NewTable("CustomListBox_Clear", 0),
		/*5*/ imports.NewTable("CustomListBox_ClearSelection", 0),
		/*6*/ imports.NewTable("CustomListBox_Click", 0),
		/*7*/ imports.NewTable("CustomListBox_ClickOnSelChange", 0),
		/*8*/ imports.NewTable("CustomListBox_Columns", 0),
		/*9*/ imports.NewTable("CustomListBox_Count", 0),
		/*10*/ imports.NewTable("CustomListBox_Create", 0),
		/*11*/ imports.NewTable("CustomListBox_DeleteSelected", 0),
		/*12*/ imports.NewTable("CustomListBox_ExtendedSelect", 0),
		/*13*/ imports.NewTable("CustomListBox_GetIndexAtXY", 0),
		/*14*/ imports.NewTable("CustomListBox_GetIndexAtY", 0),
		/*15*/ imports.NewTable("CustomListBox_GetSelectedText", 0),
		/*16*/ imports.NewTable("CustomListBox_IntegralHeight", 0),
		/*17*/ imports.NewTable("CustomListBox_ItemAtPos", 0),
		/*18*/ imports.NewTable("CustomListBox_ItemFullyVisible", 0),
		/*19*/ imports.NewTable("CustomListBox_ItemHeight", 0),
		/*20*/ imports.NewTable("CustomListBox_ItemIndex", 0),
		/*21*/ imports.NewTable("CustomListBox_ItemRect", 0),
		/*22*/ imports.NewTable("CustomListBox_ItemVisible", 0),
		/*23*/ imports.NewTable("CustomListBox_Items", 0),
		/*24*/ imports.NewTable("CustomListBox_LockSelectionChange", 0),
		/*25*/ imports.NewTable("CustomListBox_MakeCurrentVisible", 0),
		/*26*/ imports.NewTable("CustomListBox_MeasureItem", 0),
		/*27*/ imports.NewTable("CustomListBox_MultiSelect", 0),
		/*28*/ imports.NewTable("CustomListBox_Options", 0),
		/*29*/ imports.NewTable("CustomListBox_ParentColor", 0),
		/*30*/ imports.NewTable("CustomListBox_ParentFont", 0),
		/*31*/ imports.NewTable("CustomListBox_ParentShowHint", 0),
		/*32*/ imports.NewTable("CustomListBox_ScrollWidth", 0),
		/*33*/ imports.NewTable("CustomListBox_SelCount", 0),
		/*34*/ imports.NewTable("CustomListBox_SelectAll", 0),
		/*35*/ imports.NewTable("CustomListBox_SelectRange", 0),
		/*36*/ imports.NewTable("CustomListBox_Selected", 0),
		/*37*/ imports.NewTable("CustomListBox_SetOnDblClick", 0),
		/*38*/ imports.NewTable("CustomListBox_SetOnDrawItem", 0),
		/*39*/ imports.NewTable("CustomListBox_SetOnMeasureItem", 0),
		/*40*/ imports.NewTable("CustomListBox_SetOnMouseDown", 0),
		/*41*/ imports.NewTable("CustomListBox_SetOnMouseEnter", 0),
		/*42*/ imports.NewTable("CustomListBox_SetOnMouseLeave", 0),
		/*43*/ imports.NewTable("CustomListBox_SetOnMouseMove", 0),
		/*44*/ imports.NewTable("CustomListBox_SetOnMouseUp", 0),
		/*45*/ imports.NewTable("CustomListBox_SetOnMouseWheel", 0),
		/*46*/ imports.NewTable("CustomListBox_SetOnMouseWheelDown", 0),
		/*47*/ imports.NewTable("CustomListBox_SetOnMouseWheelUp", 0),
		/*48*/ imports.NewTable("CustomListBox_SetOnSelectionChange", 0),
		/*49*/ imports.NewTable("CustomListBox_Sorted", 0),
		/*50*/ imports.NewTable("CustomListBox_Style", 0),
		/*51*/ imports.NewTable("CustomListBox_TopIndex", 0),
		/*52*/ imports.NewTable("CustomListBox_UnlockSelectionChange", 0),
	}
)

func customListBoxImportAPI() *imports.Imports {
	if customListBoxImport == nil {
		customListBoxImport = NewDefaultImports()
		customListBoxImport.SetImportTable(customListBoxImportTables)
		customListBoxImportTables = nil
	}
	return customListBoxImport
}
