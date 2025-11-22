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

// ICustomTaskDialog Parent: ILCLComponent
type ICustomTaskDialog interface {
	ILCLComponent
	ButtonIDToModalResult(buttonID int32) types.TModalResult // function
	ExecuteToBool() bool                                     // function
	ExecuteWithHWND(parentWnd types.HWND) bool               // function
	Button() ITaskDialogButtonItem                           // property Button Getter
	SetButton(value ITaskDialogButtonItem)                   // property Button Setter
	Buttons() ITaskDialogButtons                             // property Buttons Getter
	SetButtons(value ITaskDialogButtons)                     // property Buttons Setter
	Caption() string                                         // property Caption Getter
	SetCaption(value string)                                 // property Caption Setter
	CustomFooterIcon() IIcon                                 // property CustomFooterIcon Getter
	SetCustomFooterIcon(value IIcon)                         // property CustomFooterIcon Setter
	CustomMainIcon() IIcon                                   // property CustomMainIcon Getter
	SetCustomMainIcon(value IIcon)                           // property CustomMainIcon Setter
	CommonButtons() types.TTaskDialogCommonButtons           // property CommonButtons Getter
	SetCommonButtons(value types.TTaskDialogCommonButtons)   // property CommonButtons Setter
	CollapseButtonCaption() string                           // property CollapseButtonCaption Getter
	SetCollapseButtonCaption(value string)                   // property CollapseButtonCaption Setter
	DefaultButton() types.TTaskDialogCommonButton            // property DefaultButton Getter
	SetDefaultButton(value types.TTaskDialogCommonButton)    // property DefaultButton Setter
	ExpandButtonCaption() string                             // property ExpandButtonCaption Getter
	SetExpandButtonCaption(value string)                     // property ExpandButtonCaption Setter
	ExpandedText() string                                    // property ExpandedText Getter
	SetExpandedText(value string)                            // property ExpandedText Setter
	Flags() types.TTaskDialogFlags                           // property Flags Getter
	SetFlags(value types.TTaskDialogFlags)                   // property Flags Setter
	FooterIcon() types.TTaskDialogIcon                       // property FooterIcon Getter
	SetFooterIcon(value types.TTaskDialogIcon)               // property FooterIcon Setter
	FooterText() string                                      // property FooterText Getter
	SetFooterText(value string)                              // property FooterText Setter
	MainIcon() types.TTaskDialogIcon                         // property MainIcon Getter
	SetMainIcon(value types.TTaskDialogIcon)                 // property MainIcon Setter
	Handle() types.THandle                                   // property Handle Getter
	ModalResult() types.TModalResult                         // property ModalResult Getter
	SetModalResult(value types.TModalResult)                 // property ModalResult Setter
	ProgressBar() ITaskDialogProgressBar                     // property ProgressBar Getter
	SetProgressBar(value ITaskDialogProgressBar)             // property ProgressBar Setter
	QueryChoices() IStrings                                  // property QueryChoices Getter
	SetQueryChoices(value IStrings)                          // property QueryChoices Setter
	QueryItemIndex() int32                                   // property QueryItemIndex Getter
	SetQueryItemIndex(value int32)                           // property QueryItemIndex Setter
	QueryResult() string                                     // property QueryResult Getter
	SetQueryResult(value string)                             // property QueryResult Setter
	RadioButton() ITaskDialogRadioButtonItem                 // property RadioButton Getter
	RadioButtons() ITaskDialogButtons                        // property RadioButtons Getter
	SetRadioButtons(value ITaskDialogButtons)                // property RadioButtons Setter
	SimpleQuery() string                                     // property SimpleQuery Getter
	SetSimpleQuery(value string)                             // property SimpleQuery Setter
	SimpleQueryPasswordChar() uint16                         // property SimpleQueryPasswordChar Getter
	SetSimpleQueryPasswordChar(value uint16)                 // property SimpleQueryPasswordChar Setter
	Text() string                                            // property Text Getter
	SetText(value string)                                    // property Text Setter
	Title() string                                           // property Title Getter
	SetTitle(value string)                                   // property Title Setter
	VerificationText() string                                // property VerificationText Getter
	SetVerificationText(value string)                        // property VerificationText Setter
	Width() int32                                            // property Width Getter
	SetWidth(value int32)                                    // property Width Setter
	URL() string                                             // property URL Getter
	Expanded() bool                                          // property Expanded Getter
	SetOnButtonClicked(fn TTaskDlgClickEvent)                // property event
	SetOnDialogConstructed(fn TNotifyEvent)                  // property event
	SetOnDialogCreated(fn TNotifyEvent)                      // property event
	SetOnDialogDestroyed(fn TNotifyEvent)                    // property event
	SetOnVerificationClicked(fn TNotifyEvent)                // property event
	SetOnExpanded(fn TNotifyEvent)                           // property event
	SetOnTimer(fn TTaskDlgTimerEvent)                        // property event
	SetOnRadioButtonClicked(fn TNotifyEvent)                 // property event
	SetOnHyperlinkClicked(fn TNotifyEvent)                   // property event
	SetOnNavigated(fn TNotifyEvent)                          // property event
	SetOnHelp(fn TNotifyEvent)                               // property event
}

type TCustomTaskDialog struct {
	TLCLComponent
}

func (m *TCustomTaskDialog) ButtonIDToModalResult(buttonID int32) types.TModalResult {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(1, m.Instance(), uintptr(buttonID))
	return types.TModalResult(r)
}

func (m *TCustomTaskDialog) ExecuteToBool() bool {
	if !m.IsValid() {
		return false
	}
	r := customTaskDialogAPI().SysCallN(2, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTaskDialog) ExecuteWithHWND(parentWnd types.HWND) bool {
	if !m.IsValid() {
		return false
	}
	r := customTaskDialogAPI().SysCallN(3, m.Instance(), uintptr(parentWnd))
	return api.GoBool(r)
}

func (m *TCustomTaskDialog) Button() ITaskDialogButtonItem {
	if !m.IsValid() {
		return nil
	}
	r := customTaskDialogAPI().SysCallN(4, 0, m.Instance())
	return AsTaskDialogButtonItem(r)
}

func (m *TCustomTaskDialog) SetButton(value ITaskDialogButtonItem) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(4, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTaskDialog) Buttons() ITaskDialogButtons {
	if !m.IsValid() {
		return nil
	}
	r := customTaskDialogAPI().SysCallN(5, 0, m.Instance())
	return AsTaskDialogButtons(r)
}

func (m *TCustomTaskDialog) SetButtons(value ITaskDialogButtons) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(5, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTaskDialog) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(6, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetCaption(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(6, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) CustomFooterIcon() IIcon {
	if !m.IsValid() {
		return nil
	}
	r := customTaskDialogAPI().SysCallN(7, 0, m.Instance())
	return AsIcon(r)
}

func (m *TCustomTaskDialog) SetCustomFooterIcon(value IIcon) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(7, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTaskDialog) CustomMainIcon() IIcon {
	if !m.IsValid() {
		return nil
	}
	r := customTaskDialogAPI().SysCallN(8, 0, m.Instance())
	return AsIcon(r)
}

func (m *TCustomTaskDialog) SetCustomMainIcon(value IIcon) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(8, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTaskDialog) CommonButtons() types.TTaskDialogCommonButtons {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(9, 0, m.Instance())
	return types.TTaskDialogCommonButtons(r)
}

func (m *TCustomTaskDialog) SetCommonButtons(value types.TTaskDialogCommonButtons) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTaskDialog) CollapseButtonCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(10, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetCollapseButtonCaption(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(10, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) DefaultButton() types.TTaskDialogCommonButton {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(11, 0, m.Instance())
	return types.TTaskDialogCommonButton(r)
}

func (m *TCustomTaskDialog) SetDefaultButton(value types.TTaskDialogCommonButton) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTaskDialog) ExpandButtonCaption() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(12, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetExpandButtonCaption(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(12, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) ExpandedText() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(13, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetExpandedText(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(13, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) Flags() types.TTaskDialogFlags {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(14, 0, m.Instance())
	return types.TTaskDialogFlags(r)
}

func (m *TCustomTaskDialog) SetFlags(value types.TTaskDialogFlags) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTaskDialog) FooterIcon() types.TTaskDialogIcon {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(15, 0, m.Instance())
	return types.TTaskDialogIcon(r)
}

func (m *TCustomTaskDialog) SetFooterIcon(value types.TTaskDialogIcon) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTaskDialog) FooterText() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(16, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetFooterText(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(16, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) MainIcon() types.TTaskDialogIcon {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(17, 0, m.Instance())
	return types.TTaskDialogIcon(r)
}

func (m *TCustomTaskDialog) SetMainIcon(value types.TTaskDialogIcon) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTaskDialog) Handle() types.THandle {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(18, m.Instance())
	return types.THandle(r)
}

func (m *TCustomTaskDialog) ModalResult() types.TModalResult {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(19, 0, m.Instance())
	return types.TModalResult(r)
}

func (m *TCustomTaskDialog) SetModalResult(value types.TModalResult) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(19, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTaskDialog) ProgressBar() ITaskDialogProgressBar {
	if !m.IsValid() {
		return nil
	}
	r := customTaskDialogAPI().SysCallN(20, 0, m.Instance())
	return AsTaskDialogProgressBar(r)
}

func (m *TCustomTaskDialog) SetProgressBar(value ITaskDialogProgressBar) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(20, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTaskDialog) QueryChoices() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := customTaskDialogAPI().SysCallN(21, 0, m.Instance())
	return AsStrings(r)
}

func (m *TCustomTaskDialog) SetQueryChoices(value IStrings) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(21, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTaskDialog) QueryItemIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(22, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTaskDialog) SetQueryItemIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(22, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTaskDialog) QueryResult() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(23, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetQueryResult(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(23, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) RadioButton() ITaskDialogRadioButtonItem {
	if !m.IsValid() {
		return nil
	}
	r := customTaskDialogAPI().SysCallN(24, m.Instance())
	return AsTaskDialogRadioButtonItem(r)
}

func (m *TCustomTaskDialog) RadioButtons() ITaskDialogButtons {
	if !m.IsValid() {
		return nil
	}
	r := customTaskDialogAPI().SysCallN(25, 0, m.Instance())
	return AsTaskDialogButtons(r)
}

func (m *TCustomTaskDialog) SetRadioButtons(value ITaskDialogButtons) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(25, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTaskDialog) SimpleQuery() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(26, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetSimpleQuery(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(26, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) SimpleQueryPasswordChar() uint16 {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(27, 0, m.Instance())
	return uint16(r)
}

func (m *TCustomTaskDialog) SetSimpleQueryPasswordChar(value uint16) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTaskDialog) Text() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(28, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetText(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(28, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(29, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(29, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) VerificationText() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(30, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) SetVerificationText(value string) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(30, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomTaskDialog) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTaskDialogAPI().SysCallN(31, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTaskDialog) SetWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customTaskDialogAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTaskDialog) URL() string {
	if !m.IsValid() {
		return ""
	}
	r := customTaskDialogAPI().SysCallN(32, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomTaskDialog) Expanded() bool {
	if !m.IsValid() {
		return false
	}
	r := customTaskDialogAPI().SysCallN(33, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTaskDialog) SetOnButtonClicked(fn TTaskDlgClickEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTaskDlgClickEvent(fn)
	base.SetEvent(m, 34, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnDialogConstructed(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 35, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnDialogCreated(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 36, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnDialogDestroyed(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 37, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnVerificationClicked(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 38, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnExpanded(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 39, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnTimer(fn TTaskDlgTimerEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTaskDlgTimerEvent(fn)
	base.SetEvent(m, 40, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnRadioButtonClicked(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 41, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnHyperlinkClicked(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 42, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnNavigated(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 43, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTaskDialog) SetOnHelp(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 44, customTaskDialogAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomTaskDialog class constructor
func NewCustomTaskDialog(owner IComponent) ICustomTaskDialog {
	r := customTaskDialogAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomTaskDialog(r)
}

func TCustomTaskDialogClass() types.TClass {
	r := customTaskDialogAPI().SysCallN(45)
	return types.TClass(r)
}

var (
	customTaskDialogOnce   base.Once
	customTaskDialogImport *imports.Imports = nil
)

func customTaskDialogAPI() *imports.Imports {
	customTaskDialogOnce.Do(func() {
		customTaskDialogImport = api.NewDefaultImports()
		customTaskDialogImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomTaskDialog_Create", 0), // constructor NewCustomTaskDialog
			/* 1 */ imports.NewTable("TCustomTaskDialog_ButtonIDToModalResult", 0), // function ButtonIDToModalResult
			/* 2 */ imports.NewTable("TCustomTaskDialog_ExecuteToBool", 0), // function ExecuteToBool
			/* 3 */ imports.NewTable("TCustomTaskDialog_ExecuteWithHWND", 0), // function ExecuteWithHWND
			/* 4 */ imports.NewTable("TCustomTaskDialog_Button", 0), // property Button
			/* 5 */ imports.NewTable("TCustomTaskDialog_Buttons", 0), // property Buttons
			/* 6 */ imports.NewTable("TCustomTaskDialog_Caption", 0), // property Caption
			/* 7 */ imports.NewTable("TCustomTaskDialog_CustomFooterIcon", 0), // property CustomFooterIcon
			/* 8 */ imports.NewTable("TCustomTaskDialog_CustomMainIcon", 0), // property CustomMainIcon
			/* 9 */ imports.NewTable("TCustomTaskDialog_CommonButtons", 0), // property CommonButtons
			/* 10 */ imports.NewTable("TCustomTaskDialog_CollapseButtonCaption", 0), // property CollapseButtonCaption
			/* 11 */ imports.NewTable("TCustomTaskDialog_DefaultButton", 0), // property DefaultButton
			/* 12 */ imports.NewTable("TCustomTaskDialog_ExpandButtonCaption", 0), // property ExpandButtonCaption
			/* 13 */ imports.NewTable("TCustomTaskDialog_ExpandedText", 0), // property ExpandedText
			/* 14 */ imports.NewTable("TCustomTaskDialog_Flags", 0), // property Flags
			/* 15 */ imports.NewTable("TCustomTaskDialog_FooterIcon", 0), // property FooterIcon
			/* 16 */ imports.NewTable("TCustomTaskDialog_FooterText", 0), // property FooterText
			/* 17 */ imports.NewTable("TCustomTaskDialog_MainIcon", 0), // property MainIcon
			/* 18 */ imports.NewTable("TCustomTaskDialog_Handle", 0), // property Handle
			/* 19 */ imports.NewTable("TCustomTaskDialog_ModalResult", 0), // property ModalResult
			/* 20 */ imports.NewTable("TCustomTaskDialog_ProgressBar", 0), // property ProgressBar
			/* 21 */ imports.NewTable("TCustomTaskDialog_QueryChoices", 0), // property QueryChoices
			/* 22 */ imports.NewTable("TCustomTaskDialog_QueryItemIndex", 0), // property QueryItemIndex
			/* 23 */ imports.NewTable("TCustomTaskDialog_QueryResult", 0), // property QueryResult
			/* 24 */ imports.NewTable("TCustomTaskDialog_RadioButton", 0), // property RadioButton
			/* 25 */ imports.NewTable("TCustomTaskDialog_RadioButtons", 0), // property RadioButtons
			/* 26 */ imports.NewTable("TCustomTaskDialog_SimpleQuery", 0), // property SimpleQuery
			/* 27 */ imports.NewTable("TCustomTaskDialog_SimpleQueryPasswordChar", 0), // property SimpleQueryPasswordChar
			/* 28 */ imports.NewTable("TCustomTaskDialog_Text", 0), // property Text
			/* 29 */ imports.NewTable("TCustomTaskDialog_Title", 0), // property Title
			/* 30 */ imports.NewTable("TCustomTaskDialog_VerificationText", 0), // property VerificationText
			/* 31 */ imports.NewTable("TCustomTaskDialog_Width", 0), // property Width
			/* 32 */ imports.NewTable("TCustomTaskDialog_URL", 0), // property URL
			/* 33 */ imports.NewTable("TCustomTaskDialog_Expanded", 0), // property Expanded
			/* 34 */ imports.NewTable("TCustomTaskDialog_OnButtonClicked", 0), // event OnButtonClicked
			/* 35 */ imports.NewTable("TCustomTaskDialog_OnDialogConstructed", 0), // event OnDialogConstructed
			/* 36 */ imports.NewTable("TCustomTaskDialog_OnDialogCreated", 0), // event OnDialogCreated
			/* 37 */ imports.NewTable("TCustomTaskDialog_OnDialogDestroyed", 0), // event OnDialogDestroyed
			/* 38 */ imports.NewTable("TCustomTaskDialog_OnVerificationClicked", 0), // event OnVerificationClicked
			/* 39 */ imports.NewTable("TCustomTaskDialog_OnExpanded", 0), // event OnExpanded
			/* 40 */ imports.NewTable("TCustomTaskDialog_OnTimer", 0), // event OnTimer
			/* 41 */ imports.NewTable("TCustomTaskDialog_OnRadioButtonClicked", 0), // event OnRadioButtonClicked
			/* 42 */ imports.NewTable("TCustomTaskDialog_OnHyperlinkClicked", 0), // event OnHyperlinkClicked
			/* 43 */ imports.NewTable("TCustomTaskDialog_OnNavigated", 0), // event OnNavigated
			/* 44 */ imports.NewTable("TCustomTaskDialog_OnHelp", 0), // event OnHelp
			/* 45 */ imports.NewTable("TCustomTaskDialog_TClass", 0), // function TCustomTaskDialogClass
		}
	})
	return customTaskDialogImport
}
