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

// ICustomComboBox Parent: IWinControl
type ICustomComboBox interface {
	IWinControl
	MatchListItem(value string) int32                                                                                                 // function
	IntfGetItems()                                                                                                                    // procedure
	AddItem(item string, anObject IObject)                                                                                            // procedure
	AddHistoryItemWithStringIntBoolX2(item string, maxHistoryCount int32, setAsText bool, caseSensitive bool)                         // procedure
	AddHistoryItemWithStringObjectIntBoolX2(item string, anObject IObject, maxHistoryCount int32, setAsText bool, caseSensitive bool) // procedure
	Clear()                                                                                                                           // procedure
	ClearSelection()                                                                                                                  // procedure
	SelectAll()                                                                                                                       // procedure
	CharCase() types.TEditCharCase                                                                                                    // property CharCase Getter
	SetCharCase(value types.TEditCharCase)                                                                                            // property CharCase Setter
	DroppedDown() bool                                                                                                                // property DroppedDown Getter
	SetDroppedDown(value bool)                                                                                                        // property DroppedDown Setter
	AutoComplete() bool                                                                                                               // property AutoComplete Getter
	SetAutoComplete(value bool)                                                                                                       // property AutoComplete Setter
	AutoCompleteText() types.TComboBoxAutoCompleteText                                                                                // property AutoCompleteText Getter
	SetAutoCompleteText(value types.TComboBoxAutoCompleteText)                                                                        // property AutoCompleteText Setter
	AutoDropDown() bool                                                                                                               // property AutoDropDown Getter
	SetAutoDropDown(value bool)                                                                                                       // property AutoDropDown Setter
	AutoSelect() bool                                                                                                                 // property AutoSelect Getter
	SetAutoSelect(value bool)                                                                                                         // property AutoSelect Setter
	AutoSelected() bool                                                                                                               // property AutoSelected Getter
	SetAutoSelected(value bool)                                                                                                       // property AutoSelected Setter
	ArrowKeysTraverseList() bool                                                                                                      // property ArrowKeysTraverseList Getter
	SetArrowKeysTraverseList(value bool)                                                                                              // property ArrowKeysTraverseList Setter
	Canvas() ICanvas                                                                                                                  // property Canvas Getter
	DropDownCount() int32                                                                                                             // property DropDownCount Getter
	SetDropDownCount(value int32)                                                                                                     // property DropDownCount Setter
	EmulatedTextHintStatus() types.TEmulatedTextHintStatus                                                                            // property EmulatedTextHintStatus Getter
	Items() IStrings                                                                                                                  // property Items Getter
	SetItems(value IStrings)                                                                                                          // property Items Setter
	ItemIndex() int32                                                                                                                 // property ItemIndex Getter
	SetItemIndex(value int32)                                                                                                         // property ItemIndex Setter
	ReadOnly() bool                                                                                                                   // property ReadOnly Getter
	SetReadOnly(value bool)                                                                                                           // property ReadOnly Setter
	SelLength() int32                                                                                                                 // property SelLength Getter
	SetSelLength(value int32)                                                                                                         // property SelLength Setter
	SelStart() int32                                                                                                                  // property SelStart Getter
	SetSelStart(value int32)                                                                                                          // property SelStart Setter
	SelText() string                                                                                                                  // property SelText Getter
	SetSelText(value string)                                                                                                          // property SelText Setter
	Style() types.TComboBoxStyle                                                                                                      // property Style Getter
	SetStyle(value types.TComboBoxStyle)                                                                                              // property Style Setter
	Text() string                                                                                                                     // property Text Getter
	SetText(value string)                                                                                                             // property Text Setter
	TextHint() string                                                                                                                 // property TextHint Getter
	SetTextHint(value string)                                                                                                         // property TextHint Setter
}

type TCustomComboBox struct {
	TWinControl
}

func (m *TCustomComboBox) MatchListItem(value string) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxAPI().SysCallN(1, m.Instance(), api.PasStr(value))
	return int32(r)
}

func (m *TCustomComboBox) IntfGetItems() {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(2, m.Instance())
}

func (m *TCustomComboBox) AddItem(item string, anObject IObject) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(3, m.Instance(), api.PasStr(item), base.GetObjectUintptr(anObject))
}

func (m *TCustomComboBox) AddHistoryItemWithStringIntBoolX2(item string, maxHistoryCount int32, setAsText bool, caseSensitive bool) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(4, m.Instance(), api.PasStr(item), uintptr(maxHistoryCount), api.PasBool(setAsText), api.PasBool(caseSensitive))
}

func (m *TCustomComboBox) AddHistoryItemWithStringObjectIntBoolX2(item string, anObject IObject, maxHistoryCount int32, setAsText bool, caseSensitive bool) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(5, m.Instance(), api.PasStr(item), base.GetObjectUintptr(anObject), uintptr(maxHistoryCount), api.PasBool(setAsText), api.PasBool(caseSensitive))
}

func (m *TCustomComboBox) Clear() {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(6, m.Instance())
}

func (m *TCustomComboBox) ClearSelection() {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(7, m.Instance())
}

func (m *TCustomComboBox) SelectAll() {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(8, m.Instance())
}

func (m *TCustomComboBox) CharCase() types.TEditCharCase {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxAPI().SysCallN(9, 0, m.Instance())
	return types.TEditCharCase(r)
}

func (m *TCustomComboBox) SetCharCase(value types.TEditCharCase) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBox) DroppedDown() bool {
	if !m.IsValid() {
		return false
	}
	r := customComboBoxAPI().SysCallN(10, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomComboBox) SetDroppedDown(value bool) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(10, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomComboBox) AutoComplete() bool {
	if !m.IsValid() {
		return false
	}
	r := customComboBoxAPI().SysCallN(11, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomComboBox) SetAutoComplete(value bool) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(11, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomComboBox) AutoCompleteText() types.TComboBoxAutoCompleteText {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxAPI().SysCallN(12, 0, m.Instance())
	return types.TComboBoxAutoCompleteText(r)
}

func (m *TCustomComboBox) SetAutoCompleteText(value types.TComboBoxAutoCompleteText) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(12, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBox) AutoDropDown() bool {
	if !m.IsValid() {
		return false
	}
	r := customComboBoxAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomComboBox) SetAutoDropDown(value bool) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomComboBox) AutoSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := customComboBoxAPI().SysCallN(14, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomComboBox) SetAutoSelect(value bool) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(14, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomComboBox) AutoSelected() bool {
	if !m.IsValid() {
		return false
	}
	r := customComboBoxAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomComboBox) SetAutoSelected(value bool) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomComboBox) ArrowKeysTraverseList() bool {
	if !m.IsValid() {
		return false
	}
	r := customComboBoxAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomComboBox) SetArrowKeysTraverseList(value bool) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomComboBox) Canvas() ICanvas {
	if !m.IsValid() {
		return nil
	}
	r := customComboBoxAPI().SysCallN(17, m.Instance())
	return AsCanvas(r)
}

func (m *TCustomComboBox) DropDownCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxAPI().SysCallN(18, 0, m.Instance())
	return int32(r)
}

func (m *TCustomComboBox) SetDropDownCount(value int32) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBox) EmulatedTextHintStatus() types.TEmulatedTextHintStatus {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxAPI().SysCallN(19, m.Instance())
	return types.TEmulatedTextHintStatus(r)
}

func (m *TCustomComboBox) Items() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := customComboBoxAPI().SysCallN(20, 0, m.Instance())
	return AsStrings(r)
}

func (m *TCustomComboBox) SetItems(value IStrings) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(20, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomComboBox) ItemIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxAPI().SysCallN(21, 0, m.Instance())
	return int32(r)
}

func (m *TCustomComboBox) SetItemIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBox) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := customComboBoxAPI().SysCallN(22, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomComboBox) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(22, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomComboBox) SelLength() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxAPI().SysCallN(23, 0, m.Instance())
	return int32(r)
}

func (m *TCustomComboBox) SetSelLength(value int32) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(23, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBox) SelStart() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxAPI().SysCallN(24, 0, m.Instance())
	return int32(r)
}

func (m *TCustomComboBox) SetSelStart(value int32) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(24, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBox) SelText() string {
	if !m.IsValid() {
		return ""
	}
	r := customComboBoxAPI().SysCallN(25, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomComboBox) SetSelText(value string) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(25, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomComboBox) Style() types.TComboBoxStyle {
	if !m.IsValid() {
		return 0
	}
	r := customComboBoxAPI().SysCallN(26, 0, m.Instance())
	return types.TComboBoxStyle(r)
}

func (m *TCustomComboBox) SetStyle(value types.TComboBoxStyle) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TCustomComboBox) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := customComboBoxAPI().SysCallN(27, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomComboBox) SetText(value string) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(27, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomComboBox) TextHint() string {
	if !m.IsValid() {
		return ""
	}
	r := customComboBoxAPI().SysCallN(28, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomComboBox) SetTextHint(value string) {
	if !m.IsValid() {
		return
	}
	customComboBoxAPI().SysCallN(28, 1, m.Instance(), api.PasStr(value))
}

// NewCustomComboBox class constructor
func NewCustomComboBox(theOwner IComponent) ICustomComboBox {
	r := customComboBoxAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomComboBox(r)
}

func TCustomComboBoxClass() types.TClass {
	r := customComboBoxAPI().SysCallN(29)
	return types.TClass(r)
}

var (
	customComboBoxOnce   base.Once
	customComboBoxImport *imports.Imports = nil
)

func customComboBoxAPI() *imports.Imports {
	customComboBoxOnce.Do(func() {
		customComboBoxImport = api.NewDefaultImports()
		customComboBoxImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomComboBox_Create", 0), // constructor NewCustomComboBox
			/* 1 */ imports.NewTable("TCustomComboBox_MatchListItem", 0), // function MatchListItem
			/* 2 */ imports.NewTable("TCustomComboBox_IntfGetItems", 0), // procedure IntfGetItems
			/* 3 */ imports.NewTable("TCustomComboBox_AddItem", 0), // procedure AddItem
			/* 4 */ imports.NewTable("TCustomComboBox_AddHistoryItemWithStringIntBoolX2", 0), // procedure AddHistoryItemWithStringIntBoolX2
			/* 5 */ imports.NewTable("TCustomComboBox_AddHistoryItemWithStringObjectIntBoolX2", 0), // procedure AddHistoryItemWithStringObjectIntBoolX2
			/* 6 */ imports.NewTable("TCustomComboBox_Clear", 0), // procedure Clear
			/* 7 */ imports.NewTable("TCustomComboBox_ClearSelection", 0), // procedure ClearSelection
			/* 8 */ imports.NewTable("TCustomComboBox_SelectAll", 0), // procedure SelectAll
			/* 9 */ imports.NewTable("TCustomComboBox_CharCase", 0), // property CharCase
			/* 10 */ imports.NewTable("TCustomComboBox_DroppedDown", 0), // property DroppedDown
			/* 11 */ imports.NewTable("TCustomComboBox_AutoComplete", 0), // property AutoComplete
			/* 12 */ imports.NewTable("TCustomComboBox_AutoCompleteText", 0), // property AutoCompleteText
			/* 13 */ imports.NewTable("TCustomComboBox_AutoDropDown", 0), // property AutoDropDown
			/* 14 */ imports.NewTable("TCustomComboBox_AutoSelect", 0), // property AutoSelect
			/* 15 */ imports.NewTable("TCustomComboBox_AutoSelected", 0), // property AutoSelected
			/* 16 */ imports.NewTable("TCustomComboBox_ArrowKeysTraverseList", 0), // property ArrowKeysTraverseList
			/* 17 */ imports.NewTable("TCustomComboBox_Canvas", 0), // property Canvas
			/* 18 */ imports.NewTable("TCustomComboBox_DropDownCount", 0), // property DropDownCount
			/* 19 */ imports.NewTable("TCustomComboBox_EmulatedTextHintStatus", 0), // property EmulatedTextHintStatus
			/* 20 */ imports.NewTable("TCustomComboBox_Items", 0), // property Items
			/* 21 */ imports.NewTable("TCustomComboBox_ItemIndex", 0), // property ItemIndex
			/* 22 */ imports.NewTable("TCustomComboBox_ReadOnly", 0), // property ReadOnly
			/* 23 */ imports.NewTable("TCustomComboBox_SelLength", 0), // property SelLength
			/* 24 */ imports.NewTable("TCustomComboBox_SelStart", 0), // property SelStart
			/* 25 */ imports.NewTable("TCustomComboBox_SelText", 0), // property SelText
			/* 26 */ imports.NewTable("TCustomComboBox_Style", 0), // property Style
			/* 27 */ imports.NewTable("TCustomComboBox_Text", 0), // property Text
			/* 28 */ imports.NewTable("TCustomComboBox_TextHint", 0), // property TextHint
			/* 29 */ imports.NewTable("TCustomComboBox_TClass", 0), // function TCustomComboBoxClass
		}
	})
	return customComboBoxImport
}
