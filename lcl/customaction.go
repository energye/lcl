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

// ICustomAction Parent: IContainedAction
type ICustomAction interface {
	IContainedAction
	DoHint(hintStr *string) bool               // function
	AutoCheck() bool                           // property AutoCheck Getter
	SetAutoCheck(value bool)                   // property AutoCheck Setter
	Caption() string                           // property Caption Getter
	SetCaption(value string)                   // property Caption Setter
	Checked() bool                             // property Checked Getter
	SetChecked(value bool)                     // property Checked Setter
	Grayed() bool                              // property Grayed Getter
	SetGrayed(value bool)                      // property Grayed Setter
	DisableIfNoHandler() bool                  // property DisableIfNoHandler Getter
	SetDisableIfNoHandler(value bool)          // property DisableIfNoHandler Setter
	Enabled() bool                             // property Enabled Getter
	SetEnabled(value bool)                     // property Enabled Setter
	GroupIndex() int32                         // property GroupIndex Getter
	SetGroupIndex(value int32)                 // property GroupIndex Setter
	HelpContext() types.THelpContext           // property HelpContext Getter
	SetHelpContext(value types.THelpContext)   // property HelpContext Setter
	HelpKeyword() string                       // property HelpKeyword Getter
	SetHelpKeyword(value string)               // property HelpKeyword Setter
	HelpType() types.THelpType                 // property HelpType Getter
	SetHelpType(value types.THelpType)         // property HelpType Setter
	Hint() string                              // property Hint Getter
	SetHint(value string)                      // property Hint Setter
	ImageIndex() int32                         // property ImageIndex Getter
	SetImageIndex(value int32)                 // property ImageIndex Setter
	SecondaryShortCuts() IShortCutList         // property SecondaryShortCuts Getter
	SetSecondaryShortCuts(value IShortCutList) // property SecondaryShortCuts Setter
	ShortCut() types.TShortCut                 // property ShortCut Getter
	SetShortCut(value types.TShortCut)         // property ShortCut Setter
	Visible() bool                             // property Visible Getter
	SetVisible(value bool)                     // property Visible Setter
	SetOnHint(fn THintEvent)                   // property event
}

type TCustomAction struct {
	TContainedAction
}

func (m *TCustomAction) DoHint(hintStr *string) bool {
	if !m.IsValid() {
		return false
	}
	hintStrPtr := api.PasStr(*hintStr)
	r := customActionAPI().SysCallN(1, m.Instance(), uintptr(base.UnsafePointer(&hintStrPtr)))
	*hintStr = api.GoStr(hintStrPtr)
	return api.GoBool(r)
}

func (m *TCustomAction) AutoCheck() bool {
	if !m.IsValid() {
		return false
	}
	r := customActionAPI().SysCallN(2, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAction) SetAutoCheck(value bool) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(2, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAction) Caption() string {
	if !m.IsValid() {
		return ""
	}
	r := customActionAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomAction) SetCaption(value string) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomAction) Checked() bool {
	if !m.IsValid() {
		return false
	}
	r := customActionAPI().SysCallN(4, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAction) SetChecked(value bool) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(4, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAction) Grayed() bool {
	if !m.IsValid() {
		return false
	}
	r := customActionAPI().SysCallN(5, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAction) SetGrayed(value bool) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(5, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAction) DisableIfNoHandler() bool {
	if !m.IsValid() {
		return false
	}
	r := customActionAPI().SysCallN(6, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAction) SetDisableIfNoHandler(value bool) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(6, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAction) Enabled() bool {
	if !m.IsValid() {
		return false
	}
	r := customActionAPI().SysCallN(7, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAction) SetEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(7, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAction) GroupIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customActionAPI().SysCallN(8, 0, m.Instance())
	return int32(r)
}

func (m *TCustomAction) SetGroupIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(8, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAction) HelpContext() types.THelpContext {
	if !m.IsValid() {
		return 0
	}
	r := customActionAPI().SysCallN(9, 0, m.Instance())
	return types.THelpContext(r)
}

func (m *TCustomAction) SetHelpContext(value types.THelpContext) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAction) HelpKeyword() string {
	if !m.IsValid() {
		return ""
	}
	r := customActionAPI().SysCallN(10, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomAction) SetHelpKeyword(value string) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(10, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomAction) HelpType() types.THelpType {
	if !m.IsValid() {
		return 0
	}
	r := customActionAPI().SysCallN(11, 0, m.Instance())
	return types.THelpType(r)
}

func (m *TCustomAction) SetHelpType(value types.THelpType) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAction) Hint() string {
	if !m.IsValid() {
		return ""
	}
	r := customActionAPI().SysCallN(12, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomAction) SetHint(value string) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(12, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomAction) ImageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customActionAPI().SysCallN(13, 0, m.Instance())
	return int32(r)
}

func (m *TCustomAction) SetImageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(13, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAction) SecondaryShortCuts() IShortCutList {
	if !m.IsValid() {
		return nil
	}
	r := customActionAPI().SysCallN(14, 0, m.Instance())
	return AsShortCutList(r)
}

func (m *TCustomAction) SetSecondaryShortCuts(value IShortCutList) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(14, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomAction) ShortCut() types.TShortCut {
	if !m.IsValid() {
		return 0
	}
	r := customActionAPI().SysCallN(15, 0, m.Instance())
	return types.TShortCut(r)
}

func (m *TCustomAction) SetShortCut(value types.TShortCut) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(15, 1, m.Instance(), uintptr(value))
}

func (m *TCustomAction) Visible() bool {
	if !m.IsValid() {
		return false
	}
	r := customActionAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomAction) SetVisible(value bool) {
	if !m.IsValid() {
		return
	}
	customActionAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomAction) SetOnHint(fn THintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTHintEvent(fn)
	base.SetEvent(m, 17, customActionAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomAction class constructor
func NewCustomAction(owner IComponent) ICustomAction {
	r := customActionAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomAction(r)
}

func TCustomActionClass() types.TClass {
	r := customActionAPI().SysCallN(18)
	return types.TClass(r)
}

var (
	customActionOnce   base.Once
	customActionImport *imports.Imports = nil
)

func customActionAPI() *imports.Imports {
	customActionOnce.Do(func() {
		customActionImport = api.NewDefaultImports()
		customActionImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomAction_Create", 0), // constructor NewCustomAction
			/* 1 */ imports.NewTable("TCustomAction_DoHint", 0), // function DoHint
			/* 2 */ imports.NewTable("TCustomAction_AutoCheck", 0), // property AutoCheck
			/* 3 */ imports.NewTable("TCustomAction_Caption", 0), // property Caption
			/* 4 */ imports.NewTable("TCustomAction_Checked", 0), // property Checked
			/* 5 */ imports.NewTable("TCustomAction_Grayed", 0), // property Grayed
			/* 6 */ imports.NewTable("TCustomAction_DisableIfNoHandler", 0), // property DisableIfNoHandler
			/* 7 */ imports.NewTable("TCustomAction_Enabled", 0), // property Enabled
			/* 8 */ imports.NewTable("TCustomAction_GroupIndex", 0), // property GroupIndex
			/* 9 */ imports.NewTable("TCustomAction_HelpContext", 0), // property HelpContext
			/* 10 */ imports.NewTable("TCustomAction_HelpKeyword", 0), // property HelpKeyword
			/* 11 */ imports.NewTable("TCustomAction_HelpType", 0), // property HelpType
			/* 12 */ imports.NewTable("TCustomAction_Hint", 0), // property Hint
			/* 13 */ imports.NewTable("TCustomAction_ImageIndex", 0), // property ImageIndex
			/* 14 */ imports.NewTable("TCustomAction_SecondaryShortCuts", 0), // property SecondaryShortCuts
			/* 15 */ imports.NewTable("TCustomAction_ShortCut", 0), // property ShortCut
			/* 16 */ imports.NewTable("TCustomAction_Visible", 0), // property Visible
			/* 17 */ imports.NewTable("TCustomAction_OnHint", 0), // event OnHint
			/* 18 */ imports.NewTable("TCustomAction_TClass", 0), // function TCustomActionClass
		}
	})
	return customActionImport
}
