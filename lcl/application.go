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

// IApplication Parent: ICustomApplication
type IApplication interface {
	ICustomApplication
	GetControlAtMouse() IControl                                                     // function
	GetControlAtPos(P types.TPoint) IControl                                         // function
	BigIconHandle() types.HICON                                                      // function
	SmallIconHandle() types.HICON                                                    // function
	HelpCommand(command uint16, data uintptr) bool                                   // function
	HelpContext(context types.THelpContext) bool                                     // function
	HelpKeyword(keyword string) bool                                                 // function
	HelpShowTableOfContents() bool                                                   // function
	IsWaiting() bool                                                                 // function
	MessageBox(text uintptr, caption uintptr, flags int32) int32                     // function
	IsShortcut(message *types.TLMKey) bool                                           // function
	IsRightToLeft() bool                                                             // function
	IsRTLLang(lang string) bool                                                      // function
	Direction(lang string) types.TBiDiMode                                           // function
	ActivateHint(cursorPos types.TPoint, checkHintControlChange bool)                // procedure
	ControlDestroyed(control IControl)                                               // procedure
	BringToFront()                                                                   // procedure
	CreateForm(instanceClass IComponent, outReference *uintptr)                      // procedure
	UpdateMainForm(form IForm)                                                       // procedure
	RemoveAsyncCalls(anObject IObject)                                               // procedure
	ReleaseComponent(component IComponent)                                           // procedure
	HandleMessage()                                                                  // procedure
	ShowHelpForObject(sender IObject)                                                // procedure
	RemoveStayOnTop(systemTopAlso bool)                                              // procedure
	RestoreStayOnTop(systemTopAlso bool)                                             // procedure
	CancelHint()                                                                     // procedure
	HideHint()                                                                       // procedure
	HintMouseMessage(control IControl, message *types.TLMessage)                     // procedure
	Minimize()                                                                       // procedure
	ModalStarted()                                                                   // procedure
	ModalFinished()                                                                  // procedure
	Restore()                                                                        // procedure
	Notification(component IComponent, operation types.TOperation)                   // procedure
	ProcessMessages()                                                                // procedure
	Idle(wait bool)                                                                  // procedure
	Run()                                                                            // procedure
	DisableIdleHandler()                                                             // procedure
	EnableIdleHandler()                                                              // procedure
	NotifyUserInputHandler(sender IObject, msg uint32)                               // procedure
	NotifyKeyDownBeforeHandler(sender IObject, key *uint16, shift types.TShiftState) // procedure
	NotifyKeyDownHandler(sender IObject, key *uint16, shift types.TShiftState)       // procedure
	ControlKeyDown(sender IObject, key *uint16, shift types.TShiftState)             // procedure
	ControlKeyUp(sender IObject, key *uint16, shift types.TShiftState)               // procedure
	RemoveAllHandlersOfObject(anObject IObject)                                      // procedure
	DoBeforeMouseMessage(curMouseControl IControl)                                   // procedure
	IntfQueryEndSession(cancel *bool)                                                // procedure
	IntfEndSession()                                                                 // procedure
	IntfAppActivate(async bool)                                                      // procedure
	IntfAppDeactivate(async bool)                                                    // procedure
	IntfAppMinimize()                                                                // procedure
	IntfAppRestore()                                                                 // procedure
	IntfSettingsChange()                                                             // procedure
	IntfThemeOptionChange(themeServices IThemeServices, option types.TThemeOption)   // procedure
	// DoArrowKey
	//  on key down
	DoArrowKey(control IWinControl, key *uint16, shift types.TShiftState) // procedure
	DoTabKey(control IWinControl, key *uint16, shift types.TShiftState)   // procedure
	// DoEscapeKey
	//  on key up
	DoEscapeKey(control IWinControl, key *uint16, shift types.TShiftState) // procedure
	DoReturnKey(control IWinControl, key *uint16, shift types.TShiftState) // procedure
	Active() bool                                                          // property Active Getter
	ActiveFormHandle() types.HWND                                          // property ActiveFormHandle Getter
	ApplicationType() types.TApplicationType                               // property ApplicationType Getter
	SetApplicationType(value types.TApplicationType)                       // property ApplicationType Setter
	BidiMode() types.TBiDiMode                                             // property BidiMode Getter
	SetBidiMode(value types.TBiDiMode)                                     // property BidiMode Setter
	CaptureExceptions() bool                                               // property CaptureExceptions Getter
	SetCaptureExceptions(value bool)                                       // property CaptureExceptions Setter
	DoubleBuffered() types.TApplicationDoubleBuffered                      // property DoubleBuffered Getter
	SetDoubleBuffered(value types.TApplicationDoubleBuffered)              // property DoubleBuffered Setter
	ExtendedKeysSupport() bool                                             // property ExtendedKeysSupport Getter
	SetExtendedKeysSupport(value bool)                                     // property ExtendedKeysSupport Setter
	ExceptionDialog() types.TApplicationExceptionDlg                       // property ExceptionDialog Getter
	SetExceptionDialog(value types.TApplicationExceptionDlg)               // property ExceptionDialog Setter
	FindGlobalComponentEnabled() bool                                      // property FindGlobalComponentEnabled Getter
	SetFindGlobalComponentEnabled(value bool)                              // property FindGlobalComponentEnabled Setter
	Flags() types.TApplicationFlags                                        // property Flags Getter
	SetFlags(value types.TApplicationFlags)                                // property Flags Setter
	// Handle
	//  property HelpSystem : IHelpSystem read FHelpSystem;
	Handle() types.TLCLHandle                                      // property Handle Getter
	SetHandle(value types.TLCLHandle)                              // property Handle Setter
	Hint() string                                                  // property Hint Getter
	SetHint(value string)                                          // property Hint Setter
	HintColor() types.TColor                                       // property HintColor Getter
	SetHintColor(value types.TColor)                               // property HintColor Setter
	HintHidePause() int32                                          // property HintHidePause Getter
	SetHintHidePause(value int32)                                  // property HintHidePause Setter
	HintHidePausePerChar() int32                                   // property HintHidePausePerChar Getter
	SetHintHidePausePerChar(value int32)                           // property HintHidePausePerChar Setter
	HintPause() int32                                              // property HintPause Getter
	SetHintPause(value int32)                                      // property HintPause Setter
	HintShortCuts() bool                                           // property HintShortCuts Getter
	SetHintShortCuts(value bool)                                   // property HintShortCuts Setter
	HintShortPause() int32                                         // property HintShortPause Getter
	SetHintShortPause(value int32)                                 // property HintShortPause Setter
	Icon() IIcon                                                   // property Icon Getter
	SetIcon(value IIcon)                                           // property Icon Setter
	LayoutAdjustmentPolicy() types.TLayoutAdjustmentPolicy         // property LayoutAdjustmentPolicy Getter
	SetLayoutAdjustmentPolicy(value types.TLayoutAdjustmentPolicy) // property LayoutAdjustmentPolicy Setter
	Navigation() types.TApplicationNavigationOptions               // property Navigation Getter
	SetNavigation(value types.TApplicationNavigationOptions)       // property Navigation Setter
	MainForm() IForm                                               // property MainForm Getter
	MainFormHandle() types.HWND                                    // property MainFormHandle Getter
	MainFormOnTaskBar() bool                                       // property MainFormOnTaskBar Getter
	SetMainFormOnTaskBar(value bool)                               // property MainFormOnTaskBar Setter
	ModalLevel() int32                                             // property ModalLevel Getter
	MoveFormFocusToChildren() bool                                 // property MoveFormFocusToChildren Getter
	SetMoveFormFocusToChildren(value bool)                         // property MoveFormFocusToChildren Setter
	MouseControl() IControl                                        // property MouseControl Getter
	TaskBarBehavior() types.TTaskBarBehavior                       // property TaskBarBehavior Getter
	SetTaskBarBehavior(value types.TTaskBarBehavior)               // property TaskBarBehavior Setter
	UpdateFormatSettings() bool                                    // property UpdateFormatSettings Getter
	SetUpdateFormatSettings(value bool)                            // property UpdateFormatSettings Setter
	ShowButtonGlyphs() types.TApplicationShowGlyphs                // property ShowButtonGlyphs Getter
	SetShowButtonGlyphs(value types.TApplicationShowGlyphs)        // property ShowButtonGlyphs Setter
	ShowMenuGlyphs() types.TApplicationShowGlyphs                  // property ShowMenuGlyphs Getter
	SetShowMenuGlyphs(value types.TApplicationShowGlyphs)          // property ShowMenuGlyphs Setter
	ShowHint() bool                                                // property ShowHint Getter
	SetShowHint(value bool)                                        // property ShowHint Setter
	ShowMainForm() bool                                            // property ShowMainForm Getter
	SetShowMainForm(value bool)                                    // property ShowMainForm Setter
	Scaled() bool                                                  // property Scaled Getter
	SetScaled(value bool)                                          // property Scaled Setter
	SetOnActionExecute(fn TActionEvent)                            // property event
	SetOnActionUpdate(fn TActionEvent)                             // property event
	SetOnActivate(fn TNotifyEvent)                                 // property event
	SetOnDeactivate(fn TNotifyEvent)                               // property event
	SetOnGetMainFormHandle(fn TGetHandleEvent)                     // property event
	SetOnIdle(fn TIdleEvent)                                       // property event
	SetOnIdleEnd(fn TNotifyEvent)                                  // property event
	SetOnEndSession(fn TNotifyEvent)                               // property event
	SetOnQueryEndSession(fn TQueryEndSessionEvent)                 // property event
	SetOnMinimize(fn TNotifyEvent)                                 // property event
	SetOnMessageDialogFinished(fn TModalDialogFinished)            // property event
	SetOnModalBegin(fn TNotifyEvent)                               // property event
	SetOnModalEnd(fn TNotifyEvent)                                 // property event
	SetOnRestore(fn TNotifyEvent)                                  // property event
	SetOnDropFiles(fn TDropFilesEvent)                             // property event
	SetOnHelp(fn THelpEvent)                                       // property event
	SetOnHint(fn TNotifyEvent)                                     // property event
	SetOnShortcut(fn TShortcutEvent)                               // property event
	SetOnShowHint(fn TShowHintEvent)                               // property event
	SetOnUserInput(fn TOnUserInputEvent)                           // property event
	SetOnDestroy(fn TNotifyEvent)                                  // property event
}

type TApplication struct {
	TCustomApplication
}

func (m *TApplication) GetControlAtMouse() IControl {
	if !m.IsValid() {
		return nil
	}
	r := applicationAPI().SysCallN(1, m.Instance())
	return AsControl(r)
}

func (m *TApplication) GetControlAtPos(P types.TPoint) IControl {
	if !m.IsValid() {
		return nil
	}
	r := applicationAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&P)))
	return AsControl(r)
}

func (m *TApplication) BigIconHandle() types.HICON {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(3, m.Instance())
	return types.HICON(r)
}

func (m *TApplication) SmallIconHandle() types.HICON {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(4, m.Instance())
	return types.HICON(r)
}

func (m *TApplication) HelpCommand(command uint16, data uintptr) bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(5, m.Instance(), uintptr(command), uintptr(data))
	return api.GoBool(r)
}

func (m *TApplication) HelpContext(context types.THelpContext) bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(6, m.Instance(), uintptr(context))
	return api.GoBool(r)
}

func (m *TApplication) HelpKeyword(keyword string) bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(7, m.Instance(), api.PasStr(keyword))
	return api.GoBool(r)
}

func (m *TApplication) HelpShowTableOfContents() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(8, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) IsWaiting() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) MessageBox(text uintptr, caption uintptr, flags int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(10, m.Instance(), uintptr(text), uintptr(caption), uintptr(flags))
	return int32(r)
}

func (m *TApplication) IsShortcut(message *types.TLMKey) bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(11, m.Instance(), uintptr(base.UnsafePointer(message)))
	return api.GoBool(r)
}

func (m *TApplication) IsRightToLeft() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(12, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) IsRTLLang(lang string) bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(13, m.Instance(), api.PasStr(lang))
	return api.GoBool(r)
}

func (m *TApplication) Direction(lang string) types.TBiDiMode {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(14, m.Instance(), api.PasStr(lang))
	return types.TBiDiMode(r)
}

func (m *TApplication) ActivateHint(cursorPos types.TPoint, checkHintControlChange bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(15, m.Instance(), uintptr(base.UnsafePointer(&cursorPos)), api.PasBool(checkHintControlChange))
}

func (m *TApplication) ControlDestroyed(control IControl) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(control))
}

func (m *TApplication) BringToFront() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(17, m.Instance())
}

func (m *TApplication) CreateForm(instanceClass IComponent, outReference *uintptr) {
	if !m.IsValid() {
		return
	}
	var referencePtr uintptr
	applicationAPI().SysCallN(18, m.Instance(), base.GetObjectUintptr(instanceClass), uintptr(base.UnsafePointer(&referencePtr)))
	*outReference = uintptr(referencePtr)
}

func (m *TApplication) UpdateMainForm(form IForm) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(form))
}

func (m *TApplication) RemoveAsyncCalls(anObject IObject) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(anObject))
}

func (m *TApplication) ReleaseComponent(component IComponent) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(21, m.Instance(), base.GetObjectUintptr(component))
}

func (m *TApplication) HandleMessage() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(22, m.Instance())
}

func (m *TApplication) ShowHelpForObject(sender IObject) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(23, m.Instance(), base.GetObjectUintptr(sender))
}

func (m *TApplication) RemoveStayOnTop(systemTopAlso bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(24, m.Instance(), api.PasBool(systemTopAlso))
}

func (m *TApplication) RestoreStayOnTop(systemTopAlso bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(25, m.Instance(), api.PasBool(systemTopAlso))
}

func (m *TApplication) CancelHint() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(26, m.Instance())
}

func (m *TApplication) HideHint() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(27, m.Instance())
}

func (m *TApplication) HintMouseMessage(control IControl, message *types.TLMessage) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(28, m.Instance(), base.GetObjectUintptr(control), uintptr(base.UnsafePointer(message)))
}

func (m *TApplication) Minimize() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(29, m.Instance())
}

func (m *TApplication) ModalStarted() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(30, m.Instance())
}

func (m *TApplication) ModalFinished() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(31, m.Instance())
}

func (m *TApplication) Restore() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(32, m.Instance())
}

func (m *TApplication) Notification(component IComponent, operation types.TOperation) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(33, m.Instance(), base.GetObjectUintptr(component), uintptr(operation))
}

func (m *TApplication) ProcessMessages() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(34, m.Instance())
}

func (m *TApplication) Idle(wait bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(35, m.Instance(), api.PasBool(wait))
}

func (m *TApplication) Run() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(36, m.Instance())
}

func (m *TApplication) DisableIdleHandler() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(37, m.Instance())
}

func (m *TApplication) EnableIdleHandler() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(38, m.Instance())
}

func (m *TApplication) NotifyUserInputHandler(sender IObject, msg uint32) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(39, m.Instance(), base.GetObjectUintptr(sender), uintptr(msg))
}

func (m *TApplication) NotifyKeyDownBeforeHandler(sender IObject, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	applicationAPI().SysCallN(40, m.Instance(), base.GetObjectUintptr(sender), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TApplication) NotifyKeyDownHandler(sender IObject, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	applicationAPI().SysCallN(41, m.Instance(), base.GetObjectUintptr(sender), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TApplication) ControlKeyDown(sender IObject, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	applicationAPI().SysCallN(42, m.Instance(), base.GetObjectUintptr(sender), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TApplication) ControlKeyUp(sender IObject, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	applicationAPI().SysCallN(43, m.Instance(), base.GetObjectUintptr(sender), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TApplication) RemoveAllHandlersOfObject(anObject IObject) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(44, m.Instance(), base.GetObjectUintptr(anObject))
}

func (m *TApplication) DoBeforeMouseMessage(curMouseControl IControl) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(45, m.Instance(), base.GetObjectUintptr(curMouseControl))
}

func (m *TApplication) IntfQueryEndSession(cancel *bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(46, m.Instance(), uintptr(base.UnsafePointer(cancel)))
}

func (m *TApplication) IntfEndSession() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(47, m.Instance())
}

func (m *TApplication) IntfAppActivate(async bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(48, m.Instance(), api.PasBool(async))
}

func (m *TApplication) IntfAppDeactivate(async bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(49, m.Instance(), api.PasBool(async))
}

func (m *TApplication) IntfAppMinimize() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(50, m.Instance())
}

func (m *TApplication) IntfAppRestore() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(51, m.Instance())
}

func (m *TApplication) IntfSettingsChange() {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(52, m.Instance())
}

func (m *TApplication) IntfThemeOptionChange(themeServices IThemeServices, option types.TThemeOption) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(53, m.Instance(), base.GetObjectUintptr(themeServices), uintptr(option))
}

func (m *TApplication) DoArrowKey(control IWinControl, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	applicationAPI().SysCallN(54, m.Instance(), base.GetObjectUintptr(control), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TApplication) DoTabKey(control IWinControl, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	applicationAPI().SysCallN(55, m.Instance(), base.GetObjectUintptr(control), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TApplication) DoEscapeKey(control IWinControl, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	applicationAPI().SysCallN(56, m.Instance(), base.GetObjectUintptr(control), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TApplication) DoReturnKey(control IWinControl, key *uint16, shift types.TShiftState) {
	if !m.IsValid() {
		return
	}
	keyPtr := uintptr(*key)
	applicationAPI().SysCallN(57, m.Instance(), base.GetObjectUintptr(control), uintptr(base.UnsafePointer(&keyPtr)), uintptr(shift))
	*key = uint16(keyPtr)
}

func (m *TApplication) Active() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(58, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) ActiveFormHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(59, m.Instance())
	return types.HWND(r)
}

func (m *TApplication) ApplicationType() types.TApplicationType {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(60, 0, m.Instance())
	return types.TApplicationType(r)
}

func (m *TApplication) SetApplicationType(value types.TApplicationType) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(60, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) BidiMode() types.TBiDiMode {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(61, 0, m.Instance())
	return types.TBiDiMode(r)
}

func (m *TApplication) SetBidiMode(value types.TBiDiMode) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(61, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) CaptureExceptions() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(62, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetCaptureExceptions(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(62, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) DoubleBuffered() types.TApplicationDoubleBuffered {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(63, 0, m.Instance())
	return types.TApplicationDoubleBuffered(r)
}

func (m *TApplication) SetDoubleBuffered(value types.TApplicationDoubleBuffered) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(63, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) ExtendedKeysSupport() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(64, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetExtendedKeysSupport(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(64, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) ExceptionDialog() types.TApplicationExceptionDlg {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(65, 0, m.Instance())
	return types.TApplicationExceptionDlg(r)
}

func (m *TApplication) SetExceptionDialog(value types.TApplicationExceptionDlg) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(65, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) FindGlobalComponentEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(66, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetFindGlobalComponentEnabled(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(66, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) Flags() types.TApplicationFlags {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(67, 0, m.Instance())
	return types.TApplicationFlags(r)
}

func (m *TApplication) SetFlags(value types.TApplicationFlags) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(67, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) Handle() types.TLCLHandle {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(68, 0, m.Instance())
	return types.TLCLHandle(r)
}

func (m *TApplication) SetHandle(value types.TLCLHandle) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(68, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) Hint() string {
	if !m.IsValid() {
		return ""
	}
	r := applicationAPI().SysCallN(69, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TApplication) SetHint(value string) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(69, 1, m.Instance(), api.PasStr(value))
}

func (m *TApplication) HintColor() types.TColor {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(70, 0, m.Instance())
	return types.TColor(r)
}

func (m *TApplication) SetHintColor(value types.TColor) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(70, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) HintHidePause() int32 {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(71, 0, m.Instance())
	return int32(r)
}

func (m *TApplication) SetHintHidePause(value int32) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(71, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) HintHidePausePerChar() int32 {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(72, 0, m.Instance())
	return int32(r)
}

func (m *TApplication) SetHintHidePausePerChar(value int32) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(72, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) HintPause() int32 {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(73, 0, m.Instance())
	return int32(r)
}

func (m *TApplication) SetHintPause(value int32) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(73, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) HintShortCuts() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(74, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetHintShortCuts(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(74, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) HintShortPause() int32 {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(75, 0, m.Instance())
	return int32(r)
}

func (m *TApplication) SetHintShortPause(value int32) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(75, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) Icon() IIcon {
	if !m.IsValid() {
		return nil
	}
	r := applicationAPI().SysCallN(76, 0, m.Instance())
	return AsIcon(r)
}

func (m *TApplication) SetIcon(value IIcon) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(76, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TApplication) LayoutAdjustmentPolicy() types.TLayoutAdjustmentPolicy {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(77, 0, m.Instance())
	return types.TLayoutAdjustmentPolicy(r)
}

func (m *TApplication) SetLayoutAdjustmentPolicy(value types.TLayoutAdjustmentPolicy) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(77, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) Navigation() types.TApplicationNavigationOptions {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(78, 0, m.Instance())
	return types.TApplicationNavigationOptions(r)
}

func (m *TApplication) SetNavigation(value types.TApplicationNavigationOptions) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(78, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) MainForm() IForm {
	if !m.IsValid() {
		return nil
	}
	r := applicationAPI().SysCallN(79, m.Instance())
	return AsForm(r)
}

func (m *TApplication) MainFormHandle() types.HWND {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(80, m.Instance())
	return types.HWND(r)
}

func (m *TApplication) MainFormOnTaskBar() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(81, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetMainFormOnTaskBar(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(81, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) ModalLevel() int32 {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(82, m.Instance())
	return int32(r)
}

func (m *TApplication) MoveFormFocusToChildren() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(83, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetMoveFormFocusToChildren(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(83, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) MouseControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := applicationAPI().SysCallN(84, m.Instance())
	return AsControl(r)
}

func (m *TApplication) TaskBarBehavior() types.TTaskBarBehavior {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(85, 0, m.Instance())
	return types.TTaskBarBehavior(r)
}

func (m *TApplication) SetTaskBarBehavior(value types.TTaskBarBehavior) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(85, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) UpdateFormatSettings() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(86, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetUpdateFormatSettings(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(86, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) ShowButtonGlyphs() types.TApplicationShowGlyphs {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(87, 0, m.Instance())
	return types.TApplicationShowGlyphs(r)
}

func (m *TApplication) SetShowButtonGlyphs(value types.TApplicationShowGlyphs) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(87, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) ShowMenuGlyphs() types.TApplicationShowGlyphs {
	if !m.IsValid() {
		return 0
	}
	r := applicationAPI().SysCallN(88, 0, m.Instance())
	return types.TApplicationShowGlyphs(r)
}

func (m *TApplication) SetShowMenuGlyphs(value types.TApplicationShowGlyphs) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(88, 1, m.Instance(), uintptr(value))
}

func (m *TApplication) ShowHint() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(89, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetShowHint(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(89, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) ShowMainForm() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(90, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetShowMainForm(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(90, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) Scaled() bool {
	if !m.IsValid() {
		return false
	}
	r := applicationAPI().SysCallN(91, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TApplication) SetScaled(value bool) {
	if !m.IsValid() {
		return
	}
	applicationAPI().SysCallN(91, 1, m.Instance(), api.PasBool(value))
}

func (m *TApplication) SetOnActionExecute(fn TActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTActionEvent(fn)
	base.SetEvent(m, 92, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnActionUpdate(fn TActionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTActionEvent(fn)
	base.SetEvent(m, 93, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnActivate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 94, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnDeactivate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 95, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnGetMainFormHandle(fn TGetHandleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTGetHandleEvent(fn)
	base.SetEvent(m, 96, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnIdle(fn TIdleEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTIdleEvent(fn)
	base.SetEvent(m, 97, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnIdleEnd(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 98, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnEndSession(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 99, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnQueryEndSession(fn TQueryEndSessionEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTQueryEndSessionEvent(fn)
	base.SetEvent(m, 100, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnMinimize(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 101, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnMessageDialogFinished(fn TModalDialogFinished) {
	if !m.IsValid() {
		return
	}
	cb := makeTModalDialogFinished(fn)
	base.SetEvent(m, 102, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnModalBegin(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 103, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnModalEnd(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 104, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnRestore(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 105, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnDropFiles(fn TDropFilesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDropFilesEvent(fn)
	base.SetEvent(m, 106, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnHelp(fn THelpEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTHelpEvent(fn)
	base.SetEvent(m, 107, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnHint(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 108, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnShortcut(fn TShortcutEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTShortcutEvent(fn)
	base.SetEvent(m, 109, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnShowHint(fn TShowHintEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTShowHintEvent(fn)
	base.SetEvent(m, 110, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnUserInput(fn TOnUserInputEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTOnUserInputEvent(fn)
	base.SetEvent(m, 111, applicationAPI(), api.MakeEventDataPtr(cb))
}

func (m *TApplication) SetOnDestroy(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 112, applicationAPI(), api.MakeEventDataPtr(cb))
}

// NewApplication class constructor
func NewApplication(owner IComponent) IApplication {
	r := applicationAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsApplication(r)
}

func TApplicationClass() types.TClass {
	r := applicationAPI().SysCallN(113)
	return types.TClass(r)
}

var (
	applicationOnce   base.Once
	applicationImport *imports.Imports = nil
)

func applicationAPI() *imports.Imports {
	applicationOnce.Do(func() {
		applicationImport = api.NewDefaultImports()
		applicationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TApplication_Create", 0), // constructor NewApplication
			/* 1 */ imports.NewTable("TApplication_GetControlAtMouse", 0), // function GetControlAtMouse
			/* 2 */ imports.NewTable("TApplication_GetControlAtPos", 0), // function GetControlAtPos
			/* 3 */ imports.NewTable("TApplication_BigIconHandle", 0), // function BigIconHandle
			/* 4 */ imports.NewTable("TApplication_SmallIconHandle", 0), // function SmallIconHandle
			/* 5 */ imports.NewTable("TApplication_HelpCommand", 0), // function HelpCommand
			/* 6 */ imports.NewTable("TApplication_HelpContext", 0), // function HelpContext
			/* 7 */ imports.NewTable("TApplication_HelpKeyword", 0), // function HelpKeyword
			/* 8 */ imports.NewTable("TApplication_HelpShowTableOfContents", 0), // function HelpShowTableOfContents
			/* 9 */ imports.NewTable("TApplication_IsWaiting", 0), // function IsWaiting
			/* 10 */ imports.NewTable("TApplication_MessageBox", 0), // function MessageBox
			/* 11 */ imports.NewTable("TApplication_IsShortcut", 0), // function IsShortcut
			/* 12 */ imports.NewTable("TApplication_IsRightToLeft", 0), // function IsRightToLeft
			/* 13 */ imports.NewTable("TApplication_IsRTLLang", 0), // function IsRTLLang
			/* 14 */ imports.NewTable("TApplication_Direction", 0), // function Direction
			/* 15 */ imports.NewTable("TApplication_ActivateHint", 0), // procedure ActivateHint
			/* 16 */ imports.NewTable("TApplication_ControlDestroyed", 0), // procedure ControlDestroyed
			/* 17 */ imports.NewTable("TApplication_BringToFront", 0), // procedure BringToFront
			/* 18 */ imports.NewTable("TApplication_CreateForm", 0), // procedure CreateForm
			/* 19 */ imports.NewTable("TApplication_UpdateMainForm", 0), // procedure UpdateMainForm
			/* 20 */ imports.NewTable("TApplication_RemoveAsyncCalls", 0), // procedure RemoveAsyncCalls
			/* 21 */ imports.NewTable("TApplication_ReleaseComponent", 0), // procedure ReleaseComponent
			/* 22 */ imports.NewTable("TApplication_HandleMessage", 0), // procedure HandleMessage
			/* 23 */ imports.NewTable("TApplication_ShowHelpForObject", 0), // procedure ShowHelpForObject
			/* 24 */ imports.NewTable("TApplication_RemoveStayOnTop", 0), // procedure RemoveStayOnTop
			/* 25 */ imports.NewTable("TApplication_RestoreStayOnTop", 0), // procedure RestoreStayOnTop
			/* 26 */ imports.NewTable("TApplication_CancelHint", 0), // procedure CancelHint
			/* 27 */ imports.NewTable("TApplication_HideHint", 0), // procedure HideHint
			/* 28 */ imports.NewTable("TApplication_HintMouseMessage", 0), // procedure HintMouseMessage
			/* 29 */ imports.NewTable("TApplication_Minimize", 0), // procedure Minimize
			/* 30 */ imports.NewTable("TApplication_ModalStarted", 0), // procedure ModalStarted
			/* 31 */ imports.NewTable("TApplication_ModalFinished", 0), // procedure ModalFinished
			/* 32 */ imports.NewTable("TApplication_Restore", 0), // procedure Restore
			/* 33 */ imports.NewTable("TApplication_Notification", 0), // procedure Notification
			/* 34 */ imports.NewTable("TApplication_ProcessMessages", 0), // procedure ProcessMessages
			/* 35 */ imports.NewTable("TApplication_Idle", 0), // procedure Idle
			/* 36 */ imports.NewTable("TApplication_Run", 0), // procedure Run
			/* 37 */ imports.NewTable("TApplication_DisableIdleHandler", 0), // procedure DisableIdleHandler
			/* 38 */ imports.NewTable("TApplication_EnableIdleHandler", 0), // procedure EnableIdleHandler
			/* 39 */ imports.NewTable("TApplication_NotifyUserInputHandler", 0), // procedure NotifyUserInputHandler
			/* 40 */ imports.NewTable("TApplication_NotifyKeyDownBeforeHandler", 0), // procedure NotifyKeyDownBeforeHandler
			/* 41 */ imports.NewTable("TApplication_NotifyKeyDownHandler", 0), // procedure NotifyKeyDownHandler
			/* 42 */ imports.NewTable("TApplication_ControlKeyDown", 0), // procedure ControlKeyDown
			/* 43 */ imports.NewTable("TApplication_ControlKeyUp", 0), // procedure ControlKeyUp
			/* 44 */ imports.NewTable("TApplication_RemoveAllHandlersOfObject", 0), // procedure RemoveAllHandlersOfObject
			/* 45 */ imports.NewTable("TApplication_DoBeforeMouseMessage", 0), // procedure DoBeforeMouseMessage
			/* 46 */ imports.NewTable("TApplication_IntfQueryEndSession", 0), // procedure IntfQueryEndSession
			/* 47 */ imports.NewTable("TApplication_IntfEndSession", 0), // procedure IntfEndSession
			/* 48 */ imports.NewTable("TApplication_IntfAppActivate", 0), // procedure IntfAppActivate
			/* 49 */ imports.NewTable("TApplication_IntfAppDeactivate", 0), // procedure IntfAppDeactivate
			/* 50 */ imports.NewTable("TApplication_IntfAppMinimize", 0), // procedure IntfAppMinimize
			/* 51 */ imports.NewTable("TApplication_IntfAppRestore", 0), // procedure IntfAppRestore
			/* 52 */ imports.NewTable("TApplication_IntfSettingsChange", 0), // procedure IntfSettingsChange
			/* 53 */ imports.NewTable("TApplication_IntfThemeOptionChange", 0), // procedure IntfThemeOptionChange
			/* 54 */ imports.NewTable("TApplication_DoArrowKey", 0), // procedure DoArrowKey
			/* 55 */ imports.NewTable("TApplication_DoTabKey", 0), // procedure DoTabKey
			/* 56 */ imports.NewTable("TApplication_DoEscapeKey", 0), // procedure DoEscapeKey
			/* 57 */ imports.NewTable("TApplication_DoReturnKey", 0), // procedure DoReturnKey
			/* 58 */ imports.NewTable("TApplication_Active", 0), // property Active
			/* 59 */ imports.NewTable("TApplication_ActiveFormHandle", 0), // property ActiveFormHandle
			/* 60 */ imports.NewTable("TApplication_ApplicationType", 0), // property ApplicationType
			/* 61 */ imports.NewTable("TApplication_BidiMode", 0), // property BidiMode
			/* 62 */ imports.NewTable("TApplication_CaptureExceptions", 0), // property CaptureExceptions
			/* 63 */ imports.NewTable("TApplication_DoubleBuffered", 0), // property DoubleBuffered
			/* 64 */ imports.NewTable("TApplication_ExtendedKeysSupport", 0), // property ExtendedKeysSupport
			/* 65 */ imports.NewTable("TApplication_ExceptionDialog", 0), // property ExceptionDialog
			/* 66 */ imports.NewTable("TApplication_FindGlobalComponentEnabled", 0), // property FindGlobalComponentEnabled
			/* 67 */ imports.NewTable("TApplication_Flags", 0), // property Flags
			/* 68 */ imports.NewTable("TApplication_Handle", 0), // property Handle
			/* 69 */ imports.NewTable("TApplication_Hint", 0), // property Hint
			/* 70 */ imports.NewTable("TApplication_HintColor", 0), // property HintColor
			/* 71 */ imports.NewTable("TApplication_HintHidePause", 0), // property HintHidePause
			/* 72 */ imports.NewTable("TApplication_HintHidePausePerChar", 0), // property HintHidePausePerChar
			/* 73 */ imports.NewTable("TApplication_HintPause", 0), // property HintPause
			/* 74 */ imports.NewTable("TApplication_HintShortCuts", 0), // property HintShortCuts
			/* 75 */ imports.NewTable("TApplication_HintShortPause", 0), // property HintShortPause
			/* 76 */ imports.NewTable("TApplication_Icon", 0), // property Icon
			/* 77 */ imports.NewTable("TApplication_LayoutAdjustmentPolicy", 0), // property LayoutAdjustmentPolicy
			/* 78 */ imports.NewTable("TApplication_Navigation", 0), // property Navigation
			/* 79 */ imports.NewTable("TApplication_MainForm", 0), // property MainForm
			/* 80 */ imports.NewTable("TApplication_MainFormHandle", 0), // property MainFormHandle
			/* 81 */ imports.NewTable("TApplication_MainFormOnTaskBar", 0), // property MainFormOnTaskBar
			/* 82 */ imports.NewTable("TApplication_ModalLevel", 0), // property ModalLevel
			/* 83 */ imports.NewTable("TApplication_MoveFormFocusToChildren", 0), // property MoveFormFocusToChildren
			/* 84 */ imports.NewTable("TApplication_MouseControl", 0), // property MouseControl
			/* 85 */ imports.NewTable("TApplication_TaskBarBehavior", 0), // property TaskBarBehavior
			/* 86 */ imports.NewTable("TApplication_UpdateFormatSettings", 0), // property UpdateFormatSettings
			/* 87 */ imports.NewTable("TApplication_ShowButtonGlyphs", 0), // property ShowButtonGlyphs
			/* 88 */ imports.NewTable("TApplication_ShowMenuGlyphs", 0), // property ShowMenuGlyphs
			/* 89 */ imports.NewTable("TApplication_ShowHint", 0), // property ShowHint
			/* 90 */ imports.NewTable("TApplication_ShowMainForm", 0), // property ShowMainForm
			/* 91 */ imports.NewTable("TApplication_Scaled", 0), // property Scaled
			/* 92 */ imports.NewTable("TApplication_OnActionExecute", 0), // event OnActionExecute
			/* 93 */ imports.NewTable("TApplication_OnActionUpdate", 0), // event OnActionUpdate
			/* 94 */ imports.NewTable("TApplication_OnActivate", 0), // event OnActivate
			/* 95 */ imports.NewTable("TApplication_OnDeactivate", 0), // event OnDeactivate
			/* 96 */ imports.NewTable("TApplication_OnGetMainFormHandle", 0), // event OnGetMainFormHandle
			/* 97 */ imports.NewTable("TApplication_OnIdle", 0), // event OnIdle
			/* 98 */ imports.NewTable("TApplication_OnIdleEnd", 0), // event OnIdleEnd
			/* 99 */ imports.NewTable("TApplication_OnEndSession", 0), // event OnEndSession
			/* 100 */ imports.NewTable("TApplication_OnQueryEndSession", 0), // event OnQueryEndSession
			/* 101 */ imports.NewTable("TApplication_OnMinimize", 0), // event OnMinimize
			/* 102 */ imports.NewTable("TApplication_OnMessageDialogFinished", 0), // event OnMessageDialogFinished
			/* 103 */ imports.NewTable("TApplication_OnModalBegin", 0), // event OnModalBegin
			/* 104 */ imports.NewTable("TApplication_OnModalEnd", 0), // event OnModalEnd
			/* 105 */ imports.NewTable("TApplication_OnRestore", 0), // event OnRestore
			/* 106 */ imports.NewTable("TApplication_OnDropFiles", 0), // event OnDropFiles
			/* 107 */ imports.NewTable("TApplication_OnHelp", 0), // event OnHelp
			/* 108 */ imports.NewTable("TApplication_OnHint", 0), // event OnHint
			/* 109 */ imports.NewTable("TApplication_OnShortcut", 0), // event OnShortcut
			/* 110 */ imports.NewTable("TApplication_OnShowHint", 0), // event OnShowHint
			/* 111 */ imports.NewTable("TApplication_OnUserInput", 0), // event OnUserInput
			/* 112 */ imports.NewTable("TApplication_OnDestroy", 0), // event OnDestroy
			/* 113 */ imports.NewTable("TApplication_TClass", 0), // function TApplicationClass
		}
	})
	return applicationImport
}
