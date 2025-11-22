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

// IApplicationProperties Parent: ILCLComponent
type IApplicationProperties interface {
	ILCLComponent
	CaptureExceptions() bool                                 // property CaptureExceptions Getter
	SetCaptureExceptions(value bool)                         // property CaptureExceptions Setter
	ExceptionDialog() types.TApplicationExceptionDlg         // property ExceptionDialog Getter
	SetExceptionDialog(value types.TApplicationExceptionDlg) // property ExceptionDialog Setter
	HelpFile() string                                        // property HelpFile Getter
	SetHelpFile(value string)                                // property HelpFile Setter
	Hint() string                                            // property Hint Getter
	SetHint(value string)                                    // property Hint Setter
	HintColor() types.TColor                                 // property HintColor Getter
	SetHintColor(value types.TColor)                         // property HintColor Setter
	HintHidePause() int32                                    // property HintHidePause Getter
	SetHintHidePause(value int32)                            // property HintHidePause Setter
	HintPause() int32                                        // property HintPause Getter
	SetHintPause(value int32)                                // property HintPause Setter
	HintShortCuts() bool                                     // property HintShortCuts Getter
	SetHintShortCuts(value bool)                             // property HintShortCuts Setter
	HintShortPause() int32                                   // property HintShortPause Getter
	SetHintShortPause(value int32)                           // property HintShortPause Setter
	ShowButtonGlyphs() types.TApplicationShowGlyphs          // property ShowButtonGlyphs Getter
	SetShowButtonGlyphs(value types.TApplicationShowGlyphs)  // property ShowButtonGlyphs Setter
	ShowMenuGlyphs() types.TApplicationShowGlyphs            // property ShowMenuGlyphs Getter
	SetShowMenuGlyphs(value types.TApplicationShowGlyphs)    // property ShowMenuGlyphs Setter
	ShowHint() bool                                          // property ShowHint Getter
	SetShowHint(value bool)                                  // property ShowHint Setter
	ShowMainForm() bool                                      // property ShowMainForm Getter
	SetShowMainForm(value bool)                              // property ShowMainForm Setter
	Title() string                                           // property Title Getter
	SetTitle(value string)                                   // property Title Setter
	SetOnActivate(fn TNotifyEvent)                           // property event
	SetOnDeactivate(fn TNotifyEvent)                         // property event
	SetOnGetMainFormHandle(fn TGetHandleEvent)               // property event
	SetOnIdle(fn TIdleEvent)                                 // property event
	SetOnIdleEnd(fn TNotifyEvent)                            // property event
	SetOnEndSession(fn TNotifyEvent)                         // property event
	SetOnQueryEndSession(fn TQueryEndSessionEvent)           // property event
	SetOnMinimize(fn TNotifyEvent)                           // property event
	SetOnModalBegin(fn TNotifyEvent)                         // property event
	SetOnModalEnd(fn TNotifyEvent)                           // property event
	SetOnRestore(fn TNotifyEvent)                            // property event
	SetOnDropFiles(fn TDropFilesEvent)                       // property event
	SetOnHelp(fn THelpEvent)                                 // property event
	SetOnHint(fn TNotifyEvent)                               // property event
	SetOnShowHint(fn TShowHintEvent)                         // property event
	SetOnUserInput(fn TOnUserInputEvent)                     // property event
	SetOnActionExecute(fn TActionEvent)                      // property event
	SetOnActionUpdate(fn TActionEvent)                       // property event
}

type TApplicationProperties struct {
	TLCLComponent
}

func (m *TApplicationProperties) CaptureExceptions() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationPropertiesAPI().SysCallN(1, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplicationProperties) SetCaptureExceptions(value bool) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(1, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplicationProperties) ExceptionDialog() types.TApplicationExceptionDlg {
	if !m.IsValid() {
		return 0
	}
	r := applicationPropertiesAPI().SysCallN(2, 0, m.Instance())
	return types.TApplicationExceptionDlg(r)
}

func (m *TApplicationProperties) SetExceptionDialog(value types.TApplicationExceptionDlg) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(2, 1, m.Instance(), uintptr(value))
}

func (m *TApplicationProperties) HelpFile() string {
	if !m.IsValid() {
		return ""
	}
	r := applicationPropertiesAPI().SysCallN(3, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TApplicationProperties) SetHelpFile(value string) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(3, 1, m.Instance(), api.PasStr(value))
}

func (m *TApplicationProperties) Hint() string {
	if !m.IsValid() {
		return ""
	}
	r := applicationPropertiesAPI().SysCallN(4, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TApplicationProperties) SetHint(value string) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(4, 1, m.Instance(), api.PasStr(value))
}

func (m *TApplicationProperties) HintColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := applicationPropertiesAPI().SysCallN(5, 0, m.Instance())
	return types.TColor(r)
}

func (m *TApplicationProperties) SetHintColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(5, 1, m.Instance(), uintptr(value))
}

func (m *TApplicationProperties) HintHidePause() int32 {
	if !m.IsValid() {
		return 0
	}
	r := applicationPropertiesAPI().SysCallN(6, 0, m.Instance())
	return int32(r)
}

func (m *TApplicationProperties) SetHintHidePause(value int32) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(6, 1, m.Instance(), uintptr(value))
}

func (m *TApplicationProperties) HintPause() int32 {
	if !m.IsValid() {
		return 0
	}
	r := applicationPropertiesAPI().SysCallN(7, 0, m.Instance())
	return int32(r)
}

func (m *TApplicationProperties) SetHintPause(value int32) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(7, 1, m.Instance(), uintptr(value))
}

func (m *TApplicationProperties) HintShortCuts() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationPropertiesAPI().SysCallN(8, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplicationProperties) SetHintShortCuts(value bool) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(8, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplicationProperties) HintShortPause() int32 {
	if !m.IsValid() {
		return 0
	}
	r := applicationPropertiesAPI().SysCallN(9, 0, m.Instance())
	return int32(r)
}

func (m *TApplicationProperties) SetHintShortPause(value int32) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(9, 1, m.Instance(), uintptr(value))
}

func (m *TApplicationProperties) ShowButtonGlyphs() types.TApplicationShowGlyphs {
	if !m.IsValid() {
		return 0
	}
	r := applicationPropertiesAPI().SysCallN(10, 0, m.Instance())
	return types.TApplicationShowGlyphs(r)
}

func (m *TApplicationProperties) SetShowButtonGlyphs(value types.TApplicationShowGlyphs) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(10, 1, m.Instance(), uintptr(value))
}

func (m *TApplicationProperties) ShowMenuGlyphs() types.TApplicationShowGlyphs {
	if !m.IsValid() {
		return 0
	}
	r := applicationPropertiesAPI().SysCallN(11, 0, m.Instance())
	return types.TApplicationShowGlyphs(r)
}

func (m *TApplicationProperties) SetShowMenuGlyphs(value types.TApplicationShowGlyphs) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(11, 1, m.Instance(), uintptr(value))
}

func (m *TApplicationProperties) ShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationPropertiesAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplicationProperties) SetShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplicationProperties) ShowMainForm() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationPropertiesAPI().SysCallN(13, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplicationProperties) SetShowMainForm(value bool) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(13, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplicationProperties) Title() string {
	if !m.IsValid() {
		return ""
	}
	r := applicationPropertiesAPI().SysCallN(14, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TApplicationProperties) SetTitle(value string) {
	if !m.IsValid() {
		return
	}
	applicationPropertiesAPI().SysCallN(14, 1, m.Instance(), api.PasStr(value))
}

func (m *TApplicationProperties) SetOnActivate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 15, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnDeactivate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 16, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnGetMainFormHandle(fn TGetHandleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetHandleEvent(fn)
	base.SetEvent(m, 17, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnIdle(fn TIdleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTIdleEvent(fn)
	base.SetEvent(m, 18, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnIdleEnd(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 19, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnEndSession(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 20, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnQueryEndSession(fn TQueryEndSessionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTQueryEndSessionEvent(fn)
	base.SetEvent(m, 21, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnMinimize(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 22, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnModalBegin(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 23, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnModalEnd(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 24, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnRestore(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 25, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnDropFiles(fn TDropFilesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDropFilesEvent(fn)
	base.SetEvent(m, 26, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnHelp(fn THelpEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTHelpEvent(fn)
	base.SetEvent(m, 27, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnHint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 28, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnShowHint(fn TShowHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTShowHintEvent(fn)
	base.SetEvent(m, 29, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnUserInput(fn TOnUserInputEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUserInputEvent(fn)
	base.SetEvent(m, 30, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnActionExecute(fn TActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTActionEvent(fn)
	base.SetEvent(m, 31, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplicationProperties) SetOnActionUpdate(fn TActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTActionEvent(fn)
	base.SetEvent(m, 32, applicationPropertiesAPI(), api.MakeEventDataPtr(cb))
}

// NewApplicationProperties class constructor
func NewApplicationProperties(owner IComponent) IApplicationProperties {
	r := applicationPropertiesAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsApplicationProperties(r)
}

func TApplicationPropertiesClass() types.TClass {
	r := applicationPropertiesAPI().SysCallN(33)
	return types.TClass(r)
}

var (
	applicationPropertiesOnce   base.Once
	applicationPropertiesImport *imports.Imports = nil
)

func applicationPropertiesAPI() *imports.Imports {
	applicationPropertiesOnce.Do(func() {
		applicationPropertiesImport = api.NewDefaultImports()
		applicationPropertiesImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TApplicationProperties_Create", 0), // constructor NewApplicationProperties
			/* 1 */ imports.NewTable("TApplicationProperties_CaptureExceptions", 0), // property CaptureExceptions
			/* 2 */ imports.NewTable("TApplicationProperties_ExceptionDialog", 0), // property ExceptionDialog
			/* 3 */ imports.NewTable("TApplicationProperties_HelpFile", 0), // property HelpFile
			/* 4 */ imports.NewTable("TApplicationProperties_Hint", 0), // property Hint
			/* 5 */ imports.NewTable("TApplicationProperties_HintColor", 0), // property HintColor
			/* 6 */ imports.NewTable("TApplicationProperties_HintHidePause", 0), // property HintHidePause
			/* 7 */ imports.NewTable("TApplicationProperties_HintPause", 0), // property HintPause
			/* 8 */ imports.NewTable("TApplicationProperties_HintShortCuts", 0), // property HintShortCuts
			/* 9 */ imports.NewTable("TApplicationProperties_HintShortPause", 0), // property HintShortPause
			/* 10 */ imports.NewTable("TApplicationProperties_ShowButtonGlyphs", 0), // property ShowButtonGlyphs
			/* 11 */ imports.NewTable("TApplicationProperties_ShowMenuGlyphs", 0), // property ShowMenuGlyphs
			/* 12 */ imports.NewTable("TApplicationProperties_ShowHint", 0), // property ShowHint
			/* 13 */ imports.NewTable("TApplicationProperties_ShowMainForm", 0), // property ShowMainForm
			/* 14 */ imports.NewTable("TApplicationProperties_Title", 0), // property Title
			/* 15 */ imports.NewTable("TApplicationProperties_OnActivate", 0), // event OnActivate
			/* 16 */ imports.NewTable("TApplicationProperties_OnDeactivate", 0), // event OnDeactivate
			/* 17 */ imports.NewTable("TApplicationProperties_OnGetMainFormHandle", 0), // event OnGetMainFormHandle
			/* 18 */ imports.NewTable("TApplicationProperties_OnIdle", 0), // event OnIdle
			/* 19 */ imports.NewTable("TApplicationProperties_OnIdleEnd", 0), // event OnIdleEnd
			/* 20 */ imports.NewTable("TApplicationProperties_OnEndSession", 0), // event OnEndSession
			/* 21 */ imports.NewTable("TApplicationProperties_OnQueryEndSession", 0), // event OnQueryEndSession
			/* 22 */ imports.NewTable("TApplicationProperties_OnMinimize", 0), // event OnMinimize
			/* 23 */ imports.NewTable("TApplicationProperties_OnModalBegin", 0), // event OnModalBegin
			/* 24 */ imports.NewTable("TApplicationProperties_OnModalEnd", 0), // event OnModalEnd
			/* 25 */ imports.NewTable("TApplicationProperties_OnRestore", 0), // event OnRestore
			/* 26 */ imports.NewTable("TApplicationProperties_OnDropFiles", 0), // event OnDropFiles
			/* 27 */ imports.NewTable("TApplicationProperties_OnHelp", 0), // event OnHelp
			/* 28 */ imports.NewTable("TApplicationProperties_OnHint", 0), // event OnHint
			/* 29 */ imports.NewTable("TApplicationProperties_OnShowHint", 0), // event OnShowHint
			/* 30 */ imports.NewTable("TApplicationProperties_OnUserInput", 0), // event OnUserInput
			/* 31 */ imports.NewTable("TApplicationProperties_OnActionExecute", 0), // event OnActionExecute
			/* 32 */ imports.NewTable("TApplicationProperties_OnActionUpdate", 0), // event OnActionUpdate
			/* 33 */ imports.NewTable("TApplicationProperties_TClass", 0), // function TApplicationPropertiesClass
		}
	})
	return applicationPropertiesImport
}
