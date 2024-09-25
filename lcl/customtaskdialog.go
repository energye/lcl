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
	r1 := LCL().SysCallN(2361, GetObjectUintptr(AOwner))
	return AsCustomTaskDialog(r1)
}

func (m *TCustomTaskDialog) Button() ITaskDialogButtonItem {
	r1 := LCL().SysCallN(2356, 0, m.Instance(), 0)
	return AsTaskDialogButtonItem(r1)
}

func (m *TCustomTaskDialog) SetButton(AValue ITaskDialogButtonItem) {
	LCL().SysCallN(2356, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTaskDialog) Buttons() ITaskDialogButtons {
	r1 := LCL().SysCallN(2357, 0, m.Instance(), 0)
	return AsTaskDialogButtons(r1)
}

func (m *TCustomTaskDialog) SetButtons(AValue ITaskDialogButtons) {
	LCL().SysCallN(2357, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTaskDialog) Caption() string {
	r1 := LCL().SysCallN(2358, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetCaption(AValue string) {
	LCL().SysCallN(2358, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) CommonButtons() TTaskDialogCommonButtons {
	r1 := LCL().SysCallN(2360, 0, m.Instance(), 0)
	return TTaskDialogCommonButtons(r1)
}

func (m *TCustomTaskDialog) SetCommonButtons(AValue TTaskDialogCommonButtons) {
	LCL().SysCallN(2360, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) DefaultButton() TTaskDialogCommonButton {
	r1 := LCL().SysCallN(2362, 0, m.Instance(), 0)
	return TTaskDialogCommonButton(r1)
}

func (m *TCustomTaskDialog) SetDefaultButton(AValue TTaskDialogCommonButton) {
	LCL().SysCallN(2362, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) ExpandButtonCaption() string {
	r1 := LCL().SysCallN(2365, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetExpandButtonCaption(AValue string) {
	LCL().SysCallN(2365, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) ExpandedText() string {
	r1 := LCL().SysCallN(2366, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetExpandedText(AValue string) {
	LCL().SysCallN(2366, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) Flags() TTaskDialogFlags {
	r1 := LCL().SysCallN(2367, 0, m.Instance(), 0)
	return TTaskDialogFlags(r1)
}

func (m *TCustomTaskDialog) SetFlags(AValue TTaskDialogFlags) {
	LCL().SysCallN(2367, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) FooterIcon() TTaskDialogIcon {
	r1 := LCL().SysCallN(2368, 0, m.Instance(), 0)
	return TTaskDialogIcon(r1)
}

func (m *TCustomTaskDialog) SetFooterIcon(AValue TTaskDialogIcon) {
	LCL().SysCallN(2368, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) FooterText() string {
	r1 := LCL().SysCallN(2369, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetFooterText(AValue string) {
	LCL().SysCallN(2369, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) MainIcon() TTaskDialogIcon {
	r1 := LCL().SysCallN(2370, 0, m.Instance(), 0)
	return TTaskDialogIcon(r1)
}

func (m *TCustomTaskDialog) SetMainIcon(AValue TTaskDialogIcon) {
	LCL().SysCallN(2370, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) ModalResult() TModalResult {
	r1 := LCL().SysCallN(2371, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TCustomTaskDialog) SetModalResult(AValue TModalResult) {
	LCL().SysCallN(2371, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) RadioButton() ITaskDialogRadioButtonItem {
	r1 := LCL().SysCallN(2372, m.Instance())
	return AsTaskDialogRadioButtonItem(r1)
}

func (m *TCustomTaskDialog) RadioButtons() ITaskDialogButtons {
	r1 := LCL().SysCallN(2373, 0, m.Instance(), 0)
	return AsTaskDialogButtons(r1)
}

func (m *TCustomTaskDialog) SetRadioButtons(AValue ITaskDialogButtons) {
	LCL().SysCallN(2373, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTaskDialog) Text() string {
	r1 := LCL().SysCallN(2375, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetText(AValue string) {
	LCL().SysCallN(2375, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) Title() string {
	r1 := LCL().SysCallN(2376, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetTitle(AValue string) {
	LCL().SysCallN(2376, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) VerificationText() string {
	r1 := LCL().SysCallN(2377, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomTaskDialog) SetVerificationText(AValue string) {
	LCL().SysCallN(2377, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomTaskDialog) Width() int32 {
	r1 := LCL().SysCallN(2378, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTaskDialog) SetWidth(AValue int32) {
	LCL().SysCallN(2378, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTaskDialog) Execute() bool {
	r1 := LCL().SysCallN(2363, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTaskDialog) Execute1(ParentWnd HWND) bool {
	r1 := LCL().SysCallN(2364, m.Instance(), uintptr(ParentWnd))
	return GoBool(r1)
}

func CustomTaskDialogClass() TClass {
	ret := LCL().SysCallN(2359)
	return TClass(ret)
}

func (m *TCustomTaskDialog) SetOnButtonClicked(fn TTaskDlgClickEvent) {
	if m.buttonClickedPtr != 0 {
		RemoveEventElement(m.buttonClickedPtr)
	}
	m.buttonClickedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2374, m.Instance(), m.buttonClickedPtr)
}
