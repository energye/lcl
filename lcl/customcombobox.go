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

// ICustomComboBox Parent: IWinControl
type ICustomComboBox interface {
	IWinControl
	CharCase() TEditCharCase                                                                             // property
	SetCharCase(AValue TEditCharCase)                                                                    // property
	DroppedDown() bool                                                                                   // property
	SetDroppedDown(AValue bool)                                                                          // property
	AutoComplete() bool                                                                                  // property
	SetAutoComplete(AValue bool)                                                                         // property
	AutoCompleteText() TComboBoxAutoCompleteText                                                         // property
	SetAutoCompleteText(AValue TComboBoxAutoCompleteText)                                                // property
	AutoDropDown() bool                                                                                  // property
	SetAutoDropDown(AValue bool)                                                                         // property
	AutoSelect() bool                                                                                    // property
	SetAutoSelect(AValue bool)                                                                           // property
	AutoSelected() bool                                                                                  // property
	SetAutoSelected(AValue bool)                                                                         // property
	ArrowKeysTraverseList() bool                                                                         // property
	SetArrowKeysTraverseList(AValue bool)                                                                // property
	Canvas() ICanvas                                                                                     // property
	DropDownCount() int32                                                                                // property
	SetDropDownCount(AValue int32)                                                                       // property
	EmulatedTextHintStatus() TEmulatedTextHintStatus                                                     // property
	Items() IStrings                                                                                     // property
	SetItems(AValue IStrings)                                                                            // property
	ItemIndex() int32                                                                                    // property
	SetItemIndex(AValue int32)                                                                           // property
	ReadOnly() bool                                                                                      // property
	SetReadOnly(AValue bool)                                                                             // property
	SelLength() int32                                                                                    // property
	SetSelLength(AValue int32)                                                                           // property
	SelStart() int32                                                                                     // property
	SetSelStart(AValue int32)                                                                            // property
	SelText() string                                                                                     // property
	SetSelText(AValue string)                                                                            // property
	Style() TComboBoxStyle                                                                               // property
	SetStyle(AValue TComboBoxStyle)                                                                      // property
	Text() string                                                                                        // property
	SetText(AValue string)                                                                               // property
	TextHint() string                                                                                    // property
	SetTextHint(AValue string)                                                                           // property
	MatchListItem(AValue string) int32                                                                   // function
	IntfGetItems()                                                                                       // procedure
	AddItem(Item string, AnObject IObject)                                                               // procedure
	AddHistoryItem(Item string, MaxHistoryCount int32, SetAsText, CaseSensitive bool)                    // procedure
	AddHistoryItem1(Item string, AnObject IObject, MaxHistoryCount int32, SetAsText, CaseSensitive bool) // procedure
	Clear()                                                                                              // procedure
	ClearSelection()                                                                                     // procedure
	SelectAll()                                                                                          // procedure
}

// TCustomComboBox Parent: TWinControl
type TCustomComboBox struct {
	TWinControl
}

func NewCustomComboBox(TheOwner IComponent) ICustomComboBox {
	r1 := customComboBoxImportAPI().SysCallN(14, GetObjectUintptr(TheOwner))
	return AsCustomComboBox(r1)
}

func (m *TCustomComboBox) CharCase() TEditCharCase {
	r1 := customComboBoxImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TEditCharCase(r1)
}

func (m *TCustomComboBox) SetCharCase(AValue TEditCharCase) {
	customComboBoxImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) DroppedDown() bool {
	r1 := customComboBoxImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetDroppedDown(AValue bool) {
	customComboBoxImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoComplete() bool {
	r1 := customComboBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoComplete(AValue bool) {
	customComboBoxImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoCompleteText() TComboBoxAutoCompleteText {
	r1 := customComboBoxImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TComboBoxAutoCompleteText(r1)
}

func (m *TCustomComboBox) SetAutoCompleteText(AValue TComboBoxAutoCompleteText) {
	customComboBoxImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) AutoDropDown() bool {
	r1 := customComboBoxImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoDropDown(AValue bool) {
	customComboBoxImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoSelect() bool {
	r1 := customComboBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoSelect(AValue bool) {
	customComboBoxImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) AutoSelected() bool {
	r1 := customComboBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetAutoSelected(AValue bool) {
	customComboBoxImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) ArrowKeysTraverseList() bool {
	r1 := customComboBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetArrowKeysTraverseList(AValue bool) {
	customComboBoxImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) Canvas() ICanvas {
	r1 := customComboBoxImportAPI().SysCallN(9, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomComboBox) DropDownCount() int32 {
	r1 := customComboBoxImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetDropDownCount(AValue int32) {
	customComboBoxImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) EmulatedTextHintStatus() TEmulatedTextHintStatus {
	r1 := customComboBoxImportAPI().SysCallN(17, m.Instance())
	return TEmulatedTextHintStatus(r1)
}

func (m *TCustomComboBox) Items() IStrings {
	r1 := customComboBoxImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomComboBox) SetItems(AValue IStrings) {
	customComboBoxImportAPI().SysCallN(20, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomComboBox) ItemIndex() int32 {
	r1 := customComboBoxImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetItemIndex(AValue int32) {
	customComboBoxImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) ReadOnly() bool {
	r1 := customComboBoxImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomComboBox) SetReadOnly(AValue bool) {
	customComboBoxImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomComboBox) SelLength() int32 {
	r1 := customComboBoxImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetSelLength(AValue int32) {
	customComboBoxImportAPI().SysCallN(23, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) SelStart() int32 {
	r1 := customComboBoxImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomComboBox) SetSelStart(AValue int32) {
	customComboBoxImportAPI().SysCallN(24, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) SelText() string {
	r1 := customComboBoxImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomComboBox) SetSelText(AValue string) {
	customComboBoxImportAPI().SysCallN(25, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomComboBox) Style() TComboBoxStyle {
	r1 := customComboBoxImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return TComboBoxStyle(r1)
}

func (m *TCustomComboBox) SetStyle(AValue TComboBoxStyle) {
	customComboBoxImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomComboBox) Text() string {
	r1 := customComboBoxImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomComboBox) SetText(AValue string) {
	customComboBoxImportAPI().SysCallN(28, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomComboBox) TextHint() string {
	r1 := customComboBoxImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomComboBox) SetTextHint(AValue string) {
	customComboBoxImportAPI().SysCallN(29, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomComboBox) MatchListItem(AValue string) int32 {
	r1 := customComboBoxImportAPI().SysCallN(21, m.Instance(), PascalStr(AValue))
	return int32(r1)
}

func CustomComboBoxClass() TClass {
	ret := customComboBoxImportAPI().SysCallN(11)
	return TClass(ret)
}

func (m *TCustomComboBox) IntfGetItems() {
	customComboBoxImportAPI().SysCallN(18, m.Instance())
}

func (m *TCustomComboBox) AddItem(Item string, AnObject IObject) {
	customComboBoxImportAPI().SysCallN(2, m.Instance(), PascalStr(Item), GetObjectUintptr(AnObject))
}

func (m *TCustomComboBox) AddHistoryItem(Item string, MaxHistoryCount int32, SetAsText, CaseSensitive bool) {
	customComboBoxImportAPI().SysCallN(0, m.Instance(), PascalStr(Item), uintptr(MaxHistoryCount), PascalBool(SetAsText), PascalBool(CaseSensitive))
}

func (m *TCustomComboBox) AddHistoryItem1(Item string, AnObject IObject, MaxHistoryCount int32, SetAsText, CaseSensitive bool) {
	customComboBoxImportAPI().SysCallN(1, m.Instance(), PascalStr(Item), GetObjectUintptr(AnObject), uintptr(MaxHistoryCount), PascalBool(SetAsText), PascalBool(CaseSensitive))
}

func (m *TCustomComboBox) Clear() {
	customComboBoxImportAPI().SysCallN(12, m.Instance())
}

func (m *TCustomComboBox) ClearSelection() {
	customComboBoxImportAPI().SysCallN(13, m.Instance())
}

func (m *TCustomComboBox) SelectAll() {
	customComboBoxImportAPI().SysCallN(26, m.Instance())
}

var (
	customComboBoxImport       *imports.Imports = nil
	customComboBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomComboBox_AddHistoryItem", 0),
		/*1*/ imports.NewTable("CustomComboBox_AddHistoryItem1", 0),
		/*2*/ imports.NewTable("CustomComboBox_AddItem", 0),
		/*3*/ imports.NewTable("CustomComboBox_ArrowKeysTraverseList", 0),
		/*4*/ imports.NewTable("CustomComboBox_AutoComplete", 0),
		/*5*/ imports.NewTable("CustomComboBox_AutoCompleteText", 0),
		/*6*/ imports.NewTable("CustomComboBox_AutoDropDown", 0),
		/*7*/ imports.NewTable("CustomComboBox_AutoSelect", 0),
		/*8*/ imports.NewTable("CustomComboBox_AutoSelected", 0),
		/*9*/ imports.NewTable("CustomComboBox_Canvas", 0),
		/*10*/ imports.NewTable("CustomComboBox_CharCase", 0),
		/*11*/ imports.NewTable("CustomComboBox_Class", 0),
		/*12*/ imports.NewTable("CustomComboBox_Clear", 0),
		/*13*/ imports.NewTable("CustomComboBox_ClearSelection", 0),
		/*14*/ imports.NewTable("CustomComboBox_Create", 0),
		/*15*/ imports.NewTable("CustomComboBox_DropDownCount", 0),
		/*16*/ imports.NewTable("CustomComboBox_DroppedDown", 0),
		/*17*/ imports.NewTable("CustomComboBox_EmulatedTextHintStatus", 0),
		/*18*/ imports.NewTable("CustomComboBox_IntfGetItems", 0),
		/*19*/ imports.NewTable("CustomComboBox_ItemIndex", 0),
		/*20*/ imports.NewTable("CustomComboBox_Items", 0),
		/*21*/ imports.NewTable("CustomComboBox_MatchListItem", 0),
		/*22*/ imports.NewTable("CustomComboBox_ReadOnly", 0),
		/*23*/ imports.NewTable("CustomComboBox_SelLength", 0),
		/*24*/ imports.NewTable("CustomComboBox_SelStart", 0),
		/*25*/ imports.NewTable("CustomComboBox_SelText", 0),
		/*26*/ imports.NewTable("CustomComboBox_SelectAll", 0),
		/*27*/ imports.NewTable("CustomComboBox_Style", 0),
		/*28*/ imports.NewTable("CustomComboBox_Text", 0),
		/*29*/ imports.NewTable("CustomComboBox_TextHint", 0),
	}
)

func customComboBoxImportAPI() *imports.Imports {
	if customComboBoxImport == nil {
		customComboBoxImport = NewDefaultImports()
		customComboBoxImport.SetImportTable(customComboBoxImportTables)
		customComboBoxImportTables = nil
	}
	return customComboBoxImport
}
