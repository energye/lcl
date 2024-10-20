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

// ICustomAction Parent: IContainedAction
type ICustomAction interface {
	IContainedAction
	AutoCheck() bool                            // property
	SetAutoCheck(AValue bool)                   // property
	Caption() string                            // property
	SetCaption(AValue string)                   // property
	Checked() bool                              // property
	SetChecked(AValue bool)                     // property
	Grayed() bool                               // property
	SetGrayed(AValue bool)                      // property
	DisableIfNoHandler() bool                   // property
	SetDisableIfNoHandler(AValue bool)          // property
	Enabled() bool                              // property
	SetEnabled(AValue bool)                     // property
	GroupIndex() int32                          // property
	SetGroupIndex(AValue int32)                 // property
	HelpContext() THelpContext                  // property
	SetHelpContext(AValue THelpContext)         // property
	HelpKeyword() string                        // property
	SetHelpKeyword(AValue string)               // property
	HelpType() THelpType                        // property
	SetHelpType(AValue THelpType)               // property
	Hint() string                               // property
	SetHint(AValue string)                      // property
	ImageIndex() TImageIndex                    // property
	SetImageIndex(AValue TImageIndex)           // property
	SecondaryShortCuts() IShortCutList          // property
	SetSecondaryShortCuts(AValue IShortCutList) // property
	ShortCut() TShortCut                        // property
	SetShortCut(AValue TShortCut)               // property
	Visible() bool                              // property
	SetVisible(AValue bool)                     // property
	SetOnHint(fn THintEvent)                    // property event
}

// TCustomAction Parent: TContainedAction
type TCustomAction struct {
	TContainedAction
	hintPtr uintptr
}

func NewCustomAction(AOwner IComponent) ICustomAction {
	r1 := customActionImportAPI().SysCallN(4, GetObjectUintptr(AOwner))
	return AsCustomAction(r1)
}

func (m *TCustomAction) AutoCheck() bool {
	r1 := customActionImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetAutoCheck(AValue bool) {
	customActionImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) Caption() string {
	r1 := customActionImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAction) SetCaption(AValue string) {
	customActionImportAPI().SysCallN(1, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAction) Checked() bool {
	r1 := customActionImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetChecked(AValue bool) {
	customActionImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) Grayed() bool {
	r1 := customActionImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetGrayed(AValue bool) {
	customActionImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) DisableIfNoHandler() bool {
	r1 := customActionImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetDisableIfNoHandler(AValue bool) {
	customActionImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) Enabled() bool {
	r1 := customActionImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetEnabled(AValue bool) {
	customActionImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomAction) GroupIndex() int32 {
	r1 := customActionImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomAction) SetGroupIndex(AValue int32) {
	customActionImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) HelpContext() THelpContext {
	r1 := customActionImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TCustomAction) SetHelpContext(AValue THelpContext) {
	customActionImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) HelpKeyword() string {
	r1 := customActionImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAction) SetHelpKeyword(AValue string) {
	customActionImportAPI().SysCallN(10, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAction) HelpType() THelpType {
	r1 := customActionImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return THelpType(r1)
}

func (m *TCustomAction) SetHelpType(AValue THelpType) {
	customActionImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) Hint() string {
	r1 := customActionImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomAction) SetHint(AValue string) {
	customActionImportAPI().SysCallN(12, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomAction) ImageIndex() TImageIndex {
	r1 := customActionImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TCustomAction) SetImageIndex(AValue TImageIndex) {
	customActionImportAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) SecondaryShortCuts() IShortCutList {
	r1 := customActionImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return AsShortCutList(r1)
}

func (m *TCustomAction) SetSecondaryShortCuts(AValue IShortCutList) {
	customActionImportAPI().SysCallN(14, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomAction) ShortCut() TShortCut {
	r1 := customActionImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return TShortCut(r1)
}

func (m *TCustomAction) SetShortCut(AValue TShortCut) {
	customActionImportAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomAction) Visible() bool {
	r1 := customActionImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomAction) SetVisible(AValue bool) {
	customActionImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func CustomActionClass() TClass {
	ret := customActionImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TCustomAction) SetOnHint(fn THintEvent) {
	if m.hintPtr != 0 {
		RemoveEventElement(m.hintPtr)
	}
	m.hintPtr = MakeEventDataPtr(fn)
	customActionImportAPI().SysCallN(15, m.Instance(), m.hintPtr)
}

var (
	customActionImport       *imports.Imports = nil
	customActionImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomAction_AutoCheck", 0),
		/*1*/ imports.NewTable("CustomAction_Caption", 0),
		/*2*/ imports.NewTable("CustomAction_Checked", 0),
		/*3*/ imports.NewTable("CustomAction_Class", 0),
		/*4*/ imports.NewTable("CustomAction_Create", 0),
		/*5*/ imports.NewTable("CustomAction_DisableIfNoHandler", 0),
		/*6*/ imports.NewTable("CustomAction_Enabled", 0),
		/*7*/ imports.NewTable("CustomAction_Grayed", 0),
		/*8*/ imports.NewTable("CustomAction_GroupIndex", 0),
		/*9*/ imports.NewTable("CustomAction_HelpContext", 0),
		/*10*/ imports.NewTable("CustomAction_HelpKeyword", 0),
		/*11*/ imports.NewTable("CustomAction_HelpType", 0),
		/*12*/ imports.NewTable("CustomAction_Hint", 0),
		/*13*/ imports.NewTable("CustomAction_ImageIndex", 0),
		/*14*/ imports.NewTable("CustomAction_SecondaryShortCuts", 0),
		/*15*/ imports.NewTable("CustomAction_SetOnHint", 0),
		/*16*/ imports.NewTable("CustomAction_ShortCut", 0),
		/*17*/ imports.NewTable("CustomAction_Visible", 0),
	}
)

func customActionImportAPI() *imports.Imports {
	if customActionImport == nil {
		customActionImport = NewDefaultImports()
		customActionImport.SetImportTable(customActionImportTables)
		customActionImportTables = nil
	}
	return customActionImport
}
