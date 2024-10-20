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

// ICustomTaskDialog Parent: IComponent
type ICustomTaskDialog interface {
	IComponent
	Button() ITaskDialogButtonItem                    // property
	SetButton(AValue ITaskDialogButtonItem)           // property
	Buttons() ITaskDialogButtons                      // property
	SetButtons(AValue ITaskDialogButtons)             // property
	Caption() string                                  // property
	SetCaption(AValue string)                         // property
	CommonButtons() TTaskDialogCommonButtons          // property
	SetCommonButtons(AValue TTaskDialogCommonButtons) // property
	DefaultButton() TTaskDialogCommonButton           // property
	SetDefaultButton(AValue TTaskDialogCommonButton)  // property
	ExpandButtonCaption() string                      // property
	SetExpandButtonCaption(AValue string)             // property
	ExpandedText() string                             // property
	SetExpandedText(AValue string)                    // property
	Flags() TTaskDialogFlags                          // property
	SetFlags(AValue TTaskDialogFlags)                 // property
	FooterIcon() TTaskDialogIcon                      // property
	SetFooterIcon(AValue TTaskDialogIcon)             // property
	FooterText() string                               // property
	SetFooterText(AValue string)                      // property
	MainIcon() TTaskDialogIcon                        // property
	SetMainIcon(AValue TTaskDialogIcon)               // property
	ModalResult() TModalResult                        // property
	SetModalResult(AValue TModalResult)               // property
	RadioButton() ITaskDialogRadioButtonItem          // property
	RadioButtons() ITaskDialogButtons                 // property
	SetRadioButtons(AValue ITaskDialogButtons)        // property
	Text() string                                     // property
	SetText(AValue string)                            // property
	Title() string                                    // property
	SetTitle(AValue string)                           // property
	VerificationText() string                         // property
	SetVerificationText(AValue string)                // property
	Width() int32                                     // property
	SetWidth(AValue int32)                            // property
	Execute() bool                                    // function
	Execute1(ParentWnd HWND) bool                     // function
	SetOnButtonClicked(fn TTaskDlgClickEvent)         // property event
}

// TCustomTaskDialog Parent: TComponent
type TCustomTaskDialog struct {
	TComponent
	buttonClickedPtr uintptr
}

func NewCustomTaskDialog(AOwner IComponent) ICustomTaskDialog {
	r1 := customTaskDialogImportAPI().SysCallN(5, GetObjectUintptr(AOwner))
	return AsCustomTaskDialog(r1)
}

func (m *TCustomTaskDialog) Button() ITaskDialogButtonItem {
	r1 := customTaskDialogImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return AsTaskDialogButtonItem(r1)
}

func (m *TCustomTaskDialog) SetButton(AValue ITaskDialogButtonItem) {
	customTaskDialogImportAPI().SysCallN(0, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTaskDialog) Buttons() ITaskDialogButtons {
	r1 := customTaskDialogImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return AsTaskDialogButtons(r1)
}

func (m *TCustomTaskDialog) SetButtons(AValue ITaskDialogButtons) {
	customTaskDialogImportAPI().SysCallN(1, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTaskDialog) Caption() string {
	r1 := customTaskDialogImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetCaption(AValue string) {
	customTaskDialogImportAPI().SysCallN(2, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) CommonButtons() TTaskDialogCommonButtons {
	r1 := customTaskDialogImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TTaskDialogCommonButtons(r1)
}

func (m *TCustomTaskDialog) SetCommonButtons(AValue TTaskDialogCommonButtons) {
	customTaskDialogImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) DefaultButton() TTaskDialogCommonButton {
	r1 := customTaskDialogImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TTaskDialogCommonButton(r1)
}

func (m *TCustomTaskDialog) SetDefaultButton(AValue TTaskDialogCommonButton) {
	customTaskDialogImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) ExpandButtonCaption() string {
	r1 := customTaskDialogImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetExpandButtonCaption(AValue string) {
	customTaskDialogImportAPI().SysCallN(9, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) ExpandedText() string {
	r1 := customTaskDialogImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetExpandedText(AValue string) {
	customTaskDialogImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) Flags() TTaskDialogFlags {
	r1 := customTaskDialogImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TTaskDialogFlags(r1)
}

func (m *TCustomTaskDialog) SetFlags(AValue TTaskDialogFlags) {
	customTaskDialogImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) FooterIcon() TTaskDialogIcon {
	r1 := customTaskDialogImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TTaskDialogIcon(r1)
}

func (m *TCustomTaskDialog) SetFooterIcon(AValue TTaskDialogIcon) {
	customTaskDialogImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) FooterText() string {
	r1 := customTaskDialogImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetFooterText(AValue string) {
	customTaskDialogImportAPI().SysCallN(13, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) MainIcon() TTaskDialogIcon {
	r1 := customTaskDialogImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return TTaskDialogIcon(r1)
}

func (m *TCustomTaskDialog) SetMainIcon(AValue TTaskDialogIcon) {
	customTaskDialogImportAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) ModalResult() TModalResult {
	r1 := customTaskDialogImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TCustomTaskDialog) SetModalResult(AValue TModalResult) {
	customTaskDialogImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) RadioButton() ITaskDialogRadioButtonItem {
	r1 := customTaskDialogImportAPI().SysCallN(16, m.Instance())
	return AsTaskDialogRadioButtonItem(r1)
}

func (m *TCustomTaskDialog) RadioButtons() ITaskDialogButtons {
	r1 := customTaskDialogImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return AsTaskDialogButtons(r1)
}

func (m *TCustomTaskDialog) SetRadioButtons(AValue ITaskDialogButtons) {
	customTaskDialogImportAPI().SysCallN(17, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTaskDialog) Text() string {
	r1 := customTaskDialogImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetText(AValue string) {
	customTaskDialogImportAPI().SysCallN(19, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) Title() string {
	r1 := customTaskDialogImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetTitle(AValue string) {
	customTaskDialogImportAPI().SysCallN(20, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) VerificationText() string {
	r1 := customTaskDialogImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetVerificationText(AValue string) {
	customTaskDialogImportAPI().SysCallN(21, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) Width() int32 {
	r1 := customTaskDialogImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTaskDialog) SetWidth(AValue int32) {
	customTaskDialogImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) Execute() bool {
	r1 := customTaskDialogImportAPI().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTaskDialog) Execute1(ParentWnd HWND) bool {
	r1 := customTaskDialogImportAPI().SysCallN(8, m.Instance(), uintptr(ParentWnd))
	return GoBool(r1)
}

func CustomTaskDialogClass() TClass {
	ret := customTaskDialogImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TCustomTaskDialog) SetOnButtonClicked(fn TTaskDlgClickEvent) {
	if m.buttonClickedPtr != 0 {
		RemoveEventElement(m.buttonClickedPtr)
	}
	m.buttonClickedPtr = MakeEventDataPtr(fn)
	customTaskDialogImportAPI().SysCallN(18, m.Instance(), m.buttonClickedPtr)
}

var (
	customTaskDialogImport       *imports.Imports = nil
	customTaskDialogImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomTaskDialog_Button", 0),
		/*1*/ imports.NewTable("CustomTaskDialog_Buttons", 0),
		/*2*/ imports.NewTable("CustomTaskDialog_Caption", 0),
		/*3*/ imports.NewTable("CustomTaskDialog_Class", 0),
		/*4*/ imports.NewTable("CustomTaskDialog_CommonButtons", 0),
		/*5*/ imports.NewTable("CustomTaskDialog_Create", 0),
		/*6*/ imports.NewTable("CustomTaskDialog_DefaultButton", 0),
		/*7*/ imports.NewTable("CustomTaskDialog_Execute", 0),
		/*8*/ imports.NewTable("CustomTaskDialog_Execute1", 0),
		/*9*/ imports.NewTable("CustomTaskDialog_ExpandButtonCaption", 0),
		/*10*/ imports.NewTable("CustomTaskDialog_ExpandedText", 0),
		/*11*/ imports.NewTable("CustomTaskDialog_Flags", 0),
		/*12*/ imports.NewTable("CustomTaskDialog_FooterIcon", 0),
		/*13*/ imports.NewTable("CustomTaskDialog_FooterText", 0),
		/*14*/ imports.NewTable("CustomTaskDialog_MainIcon", 0),
		/*15*/ imports.NewTable("CustomTaskDialog_ModalResult", 0),
		/*16*/ imports.NewTable("CustomTaskDialog_RadioButton", 0),
		/*17*/ imports.NewTable("CustomTaskDialog_RadioButtons", 0),
		/*18*/ imports.NewTable("CustomTaskDialog_SetOnButtonClicked", 0),
		/*19*/ imports.NewTable("CustomTaskDialog_Text", 0),
		/*20*/ imports.NewTable("CustomTaskDialog_Title", 0),
		/*21*/ imports.NewTable("CustomTaskDialog_VerificationText", 0),
		/*22*/ imports.NewTable("CustomTaskDialog_Width", 0),
	}
)

func customTaskDialogImportAPI() *imports.Imports {
	if customTaskDialogImport == nil {
		customTaskDialogImport = NewDefaultImports()
		customTaskDialogImport.SetImportTable(customTaskDialogImportTables)
		customTaskDialogImportTables = nil
	}
	return customTaskDialogImport
}
