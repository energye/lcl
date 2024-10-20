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

// IApplication Parent: ICustomApplication
type IApplication interface {
	ICustomApplication
	CreateForm(forms ...IForm) IForm
	Initialize()
	SetRunLoopReceived(proc uintptr)
	SetIconResId(id int)
	Active() bool                                                              // property
	ApplicationType() TApplicationType                                         // property
	SetApplicationType(AValue TApplicationType)                                // property
	BidiMode() TBiDiMode                                                       // property
	SetBidiMode(AValue TBiDiMode)                                              // property
	CaptureExceptions() bool                                                   // property
	SetCaptureExceptions(AValue bool)                                          // property
	DoubleBuffered() TApplicationDoubleBuffered                                // property
	SetDoubleBuffered(AValue TApplicationDoubleBuffered)                       // property
	ExtendedKeysSupport() bool                                                 // property
	SetExtendedKeysSupport(AValue bool)                                        // property
	ExceptionDialog() TApplicationExceptionDlg                                 // property
	SetExceptionDialog(AValue TApplicationExceptionDlg)                        // property
	FindGlobalComponentEnabled() bool                                          // property
	SetFindGlobalComponentEnabled(AValue bool)                                 // property
	Flags() TApplicationFlags                                                  // property
	SetFlags(AValue TApplicationFlags)                                         // property
	Handle() THandle                                                           // property
	SetHandle(AValue THandle)                                                  // property
	Hint() string                                                              // property
	SetHint(AValue string)                                                     // property
	HintColor() TColor                                                         // property
	SetHintColor(AValue TColor)                                                // property
	HintHidePause() int32                                                      // property
	SetHintHidePause(AValue int32)                                             // property
	HintHidePausePerChar() int32                                               // property
	SetHintHidePausePerChar(AValue int32)                                      // property
	HintPause() int32                                                          // property
	SetHintPause(AValue int32)                                                 // property
	HintShortCuts() bool                                                       // property
	SetHintShortCuts(AValue bool)                                              // property
	HintShortPause() int32                                                     // property
	SetHintShortPause(AValue int32)                                            // property
	Icon() IIcon                                                               // property
	SetIcon(AValue IIcon)                                                      // property
	LayoutAdjustmentPolicy() TLayoutAdjustmentPolicy                           // property
	SetLayoutAdjustmentPolicy(AValue TLayoutAdjustmentPolicy)                  // property
	Navigation() TApplicationNavigationOptions                                 // property
	SetNavigation(AValue TApplicationNavigationOptions)                        // property
	MainForm() IForm                                                           // property
	MainFormHandle() HWND                                                      // property
	MainFormOnTaskBar() bool                                                   // property
	SetMainFormOnTaskBar(AValue bool)                                          // property
	ModalLevel() int32                                                         // property
	MoveFormFocusToChildren() bool                                             // property
	SetMoveFormFocusToChildren(AValue bool)                                    // property
	MouseControl() IControl                                                    // property
	TaskBarBehavior() TTaskBarBehavior                                         // property
	SetTaskBarBehavior(AValue TTaskBarBehavior)                                // property
	UpdateFormatSettings() bool                                                // property
	SetUpdateFormatSettings(AValue bool)                                       // property
	ShowButtonGlyphs() TApplicationShowGlyphs                                  // property
	SetShowButtonGlyphs(AValue TApplicationShowGlyphs)                         // property
	ShowMenuGlyphs() TApplicationShowGlyphs                                    // property
	SetShowMenuGlyphs(AValue TApplicationShowGlyphs)                           // property
	ShowHint() bool                                                            // property
	SetShowHint(AValue bool)                                                   // property
	ShowMainForm() bool                                                        // property
	SetShowMainForm(AValue bool)                                               // property
	Scaled() bool                                                              // property
	SetScaled(AValue bool)                                                     // property
	GetControlAtMouse() IControl                                               // function
	GetControlAtPos(P *TPoint) IControl                                        // function
	BigIconHandle() HICON                                                      // function
	SmallIconHandle() HICON                                                    // function
	HelpCommand(Command Word, Data uint32) bool                                // function
	HelpContext(Context THelpContext) bool                                     // function
	HelpKeyword(Keyword string) bool                                           // function
	HelpShowTableOfContents() bool                                             // function
	IsWaiting() bool                                                           // function
	MessageBox(Text, Caption string, Flags int32) int32                        // function
	IsShortcut(Message *TLMKey) bool                                           // function
	IsRightToLeft() bool                                                       // function
	IsRTLLang(ALang string) bool                                               // function
	Direction(ALang string) TBiDiMode                                          // function
	ActivateHint(CursorPos *TPoint, CheckHintControlChange bool)               // procedure
	ControlDestroyed(AControl IControl)                                        // procedure
	BringToFront()                                                             // procedure
	UpdateMainForm(AForm IForm)                                                // procedure
	RemoveAsyncCalls(AnObject IObject)                                         // procedure
	ReleaseComponent(AComponent IComponent)                                    // procedure
	HandleMessage()                                                            // procedure
	ShowHelpForObject(Sender IObject)                                          // procedure
	RemoveStayOnTop(ASystemTopAlso bool)                                       // procedure
	RestoreStayOnTop(ASystemTopAlso bool)                                      // procedure
	CancelHint()                                                               // procedure
	HideHint()                                                                 // procedure
	HintMouseMessage(Control IControl, AMessage *TLMessage)                    // procedure
	Minimize()                                                                 // procedure
	ModalStarted()                                                             // procedure
	ModalFinished()                                                            // procedure
	Restore()                                                                  // procedure
	Notification(AComponent IComponent, Operation TOperation)                  // procedure
	ProcessMessages()                                                          // procedure
	Idle(Wait bool)                                                            // procedure
	DisableIdleHandler()                                                       // procedure
	EnableIdleHandler()                                                        // procedure
	NotifyUserInputHandler(Sender IObject, Msg uint32)                         // procedure
	NotifyKeyDownBeforeHandler(Sender IObject, Key *Word, Shift TShiftState)   // procedure
	NotifyKeyDownHandler(Sender IObject, Key *Word, Shift TShiftState)         // procedure
	ControlKeyDown(Sender IObject, Key *Word, Shift TShiftState)               // procedure
	ControlKeyUp(Sender IObject, Key *Word, Shift TShiftState)                 // procedure
	RemoveAllHandlersOfObject(AnObject IObject)                                // procedure
	DoBeforeMouseMessage(CurMouseControl IControl)                             // procedure
	IntfEndSession()                                                           // procedure
	IntfAppActivate(Async bool)                                                // procedure
	IntfAppDeactivate(Async bool)                                              // procedure
	IntfAppMinimize()                                                          // procedure
	IntfAppRestore()                                                           // procedure
	IntfSettingsChange()                                                       // procedure
	IntfThemeOptionChange(AThemeServices IThemeServices, AOption TThemeOption) // procedure
	DoArrowKey(AControl IWinControl, Key *Word, Shift TShiftState)             // procedure
	DoTabKey(AControl IWinControl, Key *Word, Shift TShiftState)               // procedure
	DoEscapeKey(AControl IWinControl, Key *Word, Shift TShiftState)            // procedure
	DoReturnKey(AControl IWinControl, Key *Word, Shift TShiftState)            // procedure
	SetOnActionExecute(fn TActionEvent)                                        // property event
	SetOnActionUpdate(fn TActionEvent)                                         // property event
	SetOnActivate(fn TNotifyEvent)                                             // property event
	SetOnDeactivate(fn TNotifyEvent)                                           // property event
	SetOnGetMainFormHandle(fn TGetHandleEvent)                                 // property event
	SetOnIdleEnd(fn TNotifyEvent)                                              // property event
	SetOnEndSession(fn TNotifyEvent)                                           // property event
	SetOnMinimize(fn TNotifyEvent)                                             // property event
	SetOnModalBegin(fn TNotifyEvent)                                           // property event
	SetOnModalEnd(fn TNotifyEvent)                                             // property event
	SetOnRestore(fn TNotifyEvent)                                              // property event
	SetOnDropFiles(fn TDropFilesEvent)                                         // property event
	SetOnHelp(fn THelpEvent)                                                   // property event
	SetOnHint(fn TNotifyEvent)                                                 // property event
	SetOnShortcut(fn TShortCutEvent)                                           // property event
	SetOnShowHint(fn TShowHintEvent)                                           // property event
	SetOnDestroy(fn TNotifyEvent)                                              // property event
	SetOnCircularException(fn TExceptionEvent)                                 // property event
}

// TApplication Parent: TCustomApplication
type TApplication struct {
	TCustomApplication
	actionExecutePtr     uintptr
	actionUpdatePtr      uintptr
	activatePtr          uintptr
	deactivatePtr        uintptr
	getMainFormHandlePtr uintptr
	idleEndPtr           uintptr
	endSessionPtr        uintptr
	minimizePtr          uintptr
	modalBeginPtr        uintptr
	modalEndPtr          uintptr
	restorePtr           uintptr
	dropFilesPtr         uintptr
	helpPtr              uintptr
	hintPtr              uintptr
	shortcutPtr          uintptr
	showHintPtr          uintptr
	destroyPtr           uintptr
	circularExceptionPtr uintptr
}

func NewApplication(AOwner IComponent) IApplication {
	r1 := applicationImportAPI().SysCallN(12, GetObjectUintptr(AOwner))
	return AsApplication(r1)
}

func (m *TApplication) Active() bool {
	r1 := applicationImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TApplication) ApplicationType() TApplicationType {
	r1 := applicationImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TApplicationType(r1)
}

func (m *TApplication) SetApplicationType(AValue TApplicationType) {
	applicationImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) BidiMode() TBiDiMode {
	r1 := applicationImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TApplication) SetBidiMode(AValue TBiDiMode) {
	applicationImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) CaptureExceptions() bool {
	r1 := applicationImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetCaptureExceptions(AValue bool) {
	applicationImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) DoubleBuffered() TApplicationDoubleBuffered {
	r1 := applicationImportAPI().SysCallN(20, 0, m.Instance(), 0)
	return TApplicationDoubleBuffered(r1)
}

func (m *TApplication) SetDoubleBuffered(AValue TApplicationDoubleBuffered) {
	applicationImportAPI().SysCallN(20, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) ExtendedKeysSupport() bool {
	r1 := applicationImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetExtendedKeysSupport(AValue bool) {
	applicationImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) ExceptionDialog() TApplicationExceptionDlg {
	r1 := applicationImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return TApplicationExceptionDlg(r1)
}

func (m *TApplication) SetExceptionDialog(AValue TApplicationExceptionDlg) {
	applicationImportAPI().SysCallN(22, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) FindGlobalComponentEnabled() bool {
	r1 := applicationImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetFindGlobalComponentEnabled(AValue bool) {
	applicationImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) Flags() TApplicationFlags {
	r1 := applicationImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return TApplicationFlags(r1)
}

func (m *TApplication) SetFlags(AValue TApplicationFlags) {
	applicationImportAPI().SysCallN(25, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) Handle() THandle {
	r1 := applicationImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return THandle(r1)
}

func (m *TApplication) SetHandle(AValue THandle) {
	applicationImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) Hint() string {
	r1 := applicationImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TApplication) SetHint(AValue string) {
	applicationImportAPI().SysCallN(35, 1, m.Instance(), PascalStr(AValue))
}

func (m *TApplication) HintColor() TColor {
	r1 := applicationImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TApplication) SetHintColor(AValue TColor) {
	applicationImportAPI().SysCallN(36, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) HintHidePause() int32 {
	r1 := applicationImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TApplication) SetHintHidePause(AValue int32) {
	applicationImportAPI().SysCallN(37, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) HintHidePausePerChar() int32 {
	r1 := applicationImportAPI().SysCallN(38, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TApplication) SetHintHidePausePerChar(AValue int32) {
	applicationImportAPI().SysCallN(38, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) HintPause() int32 {
	r1 := applicationImportAPI().SysCallN(40, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TApplication) SetHintPause(AValue int32) {
	applicationImportAPI().SysCallN(40, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) HintShortCuts() bool {
	r1 := applicationImportAPI().SysCallN(41, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetHintShortCuts(AValue bool) {
	applicationImportAPI().SysCallN(41, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) HintShortPause() int32 {
	r1 := applicationImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TApplication) SetHintShortPause(AValue int32) {
	applicationImportAPI().SysCallN(42, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) Icon() IIcon {
	r1 := applicationImportAPI().SysCallN(43, 0, m.Instance(), 0)
	return AsIcon(r1)
}

func (m *TApplication) SetIcon(AValue IIcon) {
	applicationImportAPI().SysCallN(43, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TApplication) LayoutAdjustmentPolicy() TLayoutAdjustmentPolicy {
	r1 := applicationImportAPI().SysCallN(56, 0, m.Instance(), 0)
	return TLayoutAdjustmentPolicy(r1)
}

func (m *TApplication) SetLayoutAdjustmentPolicy(AValue TLayoutAdjustmentPolicy) {
	applicationImportAPI().SysCallN(56, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) Navigation() TApplicationNavigationOptions {
	r1 := applicationImportAPI().SysCallN(67, 0, m.Instance(), 0)
	return TApplicationNavigationOptions(r1)
}

func (m *TApplication) SetNavigation(AValue TApplicationNavigationOptions) {
	applicationImportAPI().SysCallN(67, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) MainForm() IForm {
	r1 := applicationImportAPI().SysCallN(57, m.Instance())
	return AsForm(r1)
}

func (m *TApplication) MainFormHandle() HWND {
	r1 := applicationImportAPI().SysCallN(58, m.Instance())
	return HWND(r1)
}

func (m *TApplication) MainFormOnTaskBar() bool {
	r1 := applicationImportAPI().SysCallN(59, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetMainFormOnTaskBar(AValue bool) {
	applicationImportAPI().SysCallN(59, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) ModalLevel() int32 {
	r1 := applicationImportAPI().SysCallN(63, m.Instance())
	return int32(r1)
}

func (m *TApplication) MoveFormFocusToChildren() bool {
	r1 := applicationImportAPI().SysCallN(66, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetMoveFormFocusToChildren(AValue bool) {
	applicationImportAPI().SysCallN(66, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) MouseControl() IControl {
	r1 := applicationImportAPI().SysCallN(65, m.Instance())
	return AsControl(r1)
}

func (m *TApplication) TaskBarBehavior() TTaskBarBehavior {
	r1 := applicationImportAPI().SysCallN(104, 0, m.Instance(), 0)
	return TTaskBarBehavior(r1)
}

func (m *TApplication) SetTaskBarBehavior(AValue TTaskBarBehavior) {
	applicationImportAPI().SysCallN(104, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) UpdateFormatSettings() bool {
	r1 := applicationImportAPI().SysCallN(105, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetUpdateFormatSettings(AValue bool) {
	applicationImportAPI().SysCallN(105, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) ShowButtonGlyphs() TApplicationShowGlyphs {
	r1 := applicationImportAPI().SysCallN(98, 0, m.Instance(), 0)
	return TApplicationShowGlyphs(r1)
}

func (m *TApplication) SetShowButtonGlyphs(AValue TApplicationShowGlyphs) {
	applicationImportAPI().SysCallN(98, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) ShowMenuGlyphs() TApplicationShowGlyphs {
	r1 := applicationImportAPI().SysCallN(102, 0, m.Instance(), 0)
	return TApplicationShowGlyphs(r1)
}

func (m *TApplication) SetShowMenuGlyphs(AValue TApplicationShowGlyphs) {
	applicationImportAPI().SysCallN(102, 1, m.Instance(), uintptr(AValue))
}

func (m *TApplication) ShowHint() bool {
	r1 := applicationImportAPI().SysCallN(100, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetShowHint(AValue bool) {
	applicationImportAPI().SysCallN(100, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) ShowMainForm() bool {
	r1 := applicationImportAPI().SysCallN(101, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetShowMainForm(AValue bool) {
	applicationImportAPI().SysCallN(101, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) Scaled() bool {
	r1 := applicationImportAPI().SysCallN(79, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TApplication) SetScaled(AValue bool) {
	applicationImportAPI().SysCallN(79, 1, m.Instance(), PascalBool(AValue))
}

func (m *TApplication) GetControlAtMouse() IControl {
	r1 := applicationImportAPI().SysCallN(26, m.Instance())
	return AsControl(r1)
}

func (m *TApplication) GetControlAtPos(P *TPoint) IControl {
	r1 := applicationImportAPI().SysCallN(27, m.Instance(), uintptr(unsafePointer(P)))
	return AsControl(r1)
}

func (m *TApplication) BigIconHandle() HICON {
	r1 := applicationImportAPI().SysCallN(4, m.Instance())
	return HICON(r1)
}

func (m *TApplication) SmallIconHandle() HICON {
	r1 := applicationImportAPI().SysCallN(103, m.Instance())
	return HICON(r1)
}

func (m *TApplication) HelpCommand(Command Word, Data uint32) bool {
	r1 := applicationImportAPI().SysCallN(30, m.Instance(), uintptr(Command), uintptr(Data))
	return GoBool(r1)
}

func (m *TApplication) HelpContext(Context THelpContext) bool {
	r1 := applicationImportAPI().SysCallN(31, m.Instance(), uintptr(Context))
	return GoBool(r1)
}

func (m *TApplication) HelpKeyword(Keyword string) bool {
	r1 := applicationImportAPI().SysCallN(32, m.Instance(), PascalStr(Keyword))
	return GoBool(r1)
}

func (m *TApplication) HelpShowTableOfContents() bool {
	r1 := applicationImportAPI().SysCallN(33, m.Instance())
	return GoBool(r1)
}

func (m *TApplication) IsWaiting() bool {
	r1 := applicationImportAPI().SysCallN(55, m.Instance())
	return GoBool(r1)
}

func (m *TApplication) MessageBox(Text, Caption string, Flags int32) int32 {
	r1 := applicationImportAPI().SysCallN(60, m.Instance(), PascalStr(Text), PascalStr(Caption), uintptr(Flags))
	return int32(r1)
}

func (m *TApplication) IsShortcut(Message *TLMKey) bool {
	var result0 uintptr
	r1 := applicationImportAPI().SysCallN(54, m.Instance(), uintptr(unsafePointer(&result0)))
	*Message = *(*TLMKey)(getPointer(result0))
	return GoBool(r1)
}

func (m *TApplication) IsRightToLeft() bool {
	r1 := applicationImportAPI().SysCallN(53, m.Instance())
	return GoBool(r1)
}

func (m *TApplication) IsRTLLang(ALang string) bool {
	r1 := applicationImportAPI().SysCallN(52, m.Instance(), PascalStr(ALang))
	return GoBool(r1)
}

func (m *TApplication) Direction(ALang string) TBiDiMode {
	r1 := applicationImportAPI().SysCallN(13, m.Instance(), PascalStr(ALang))
	return TBiDiMode(r1)
}

func ApplicationClass() TClass {
	ret := applicationImportAPI().SysCallN(8)
	return TClass(ret)
}

func (m *TApplication) ActivateHint(CursorPos *TPoint, CheckHintControlChange bool) {
	applicationImportAPI().SysCallN(0, m.Instance(), uintptr(unsafePointer(CursorPos)), PascalBool(CheckHintControlChange))
}

func (m *TApplication) ControlDestroyed(AControl IControl) {
	applicationImportAPI().SysCallN(9, m.Instance(), GetObjectUintptr(AControl))
}

func (m *TApplication) BringToFront() {
	applicationImportAPI().SysCallN(5, m.Instance())
}

func (m *TApplication) UpdateMainForm(AForm IForm) {
	applicationImportAPI().SysCallN(106, m.Instance(), GetObjectUintptr(AForm))
}

func (m *TApplication) RemoveAsyncCalls(AnObject IObject) {
	applicationImportAPI().SysCallN(75, m.Instance(), GetObjectUintptr(AnObject))
}

func (m *TApplication) ReleaseComponent(AComponent IComponent) {
	applicationImportAPI().SysCallN(73, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TApplication) HandleMessage() {
	applicationImportAPI().SysCallN(29, m.Instance())
}

func (m *TApplication) ShowHelpForObject(Sender IObject) {
	applicationImportAPI().SysCallN(99, m.Instance(), GetObjectUintptr(Sender))
}

func (m *TApplication) RemoveStayOnTop(ASystemTopAlso bool) {
	applicationImportAPI().SysCallN(76, m.Instance(), PascalBool(ASystemTopAlso))
}

func (m *TApplication) RestoreStayOnTop(ASystemTopAlso bool) {
	applicationImportAPI().SysCallN(78, m.Instance(), PascalBool(ASystemTopAlso))
}

func (m *TApplication) CancelHint() {
	applicationImportAPI().SysCallN(6, m.Instance())
}

func (m *TApplication) HideHint() {
	applicationImportAPI().SysCallN(34, m.Instance())
}

func (m *TApplication) HintMouseMessage(Control IControl, AMessage *TLMessage) {
	var result1 uintptr
	applicationImportAPI().SysCallN(39, m.Instance(), GetObjectUintptr(Control), uintptr(unsafePointer(&result1)))
	*AMessage = *(*TLMessage)(getPointer(result1))
}

func (m *TApplication) Minimize() {
	applicationImportAPI().SysCallN(61, m.Instance())
}

func (m *TApplication) ModalStarted() {
	applicationImportAPI().SysCallN(64, m.Instance())
}

func (m *TApplication) ModalFinished() {
	applicationImportAPI().SysCallN(62, m.Instance())
}

func (m *TApplication) Restore() {
	applicationImportAPI().SysCallN(77, m.Instance())
}

func (m *TApplication) Notification(AComponent IComponent, Operation TOperation) {
	applicationImportAPI().SysCallN(68, m.Instance(), GetObjectUintptr(AComponent), uintptr(Operation))
}

func (m *TApplication) ProcessMessages() {
	applicationImportAPI().SysCallN(72, m.Instance())
}

func (m *TApplication) Idle(Wait bool) {
	applicationImportAPI().SysCallN(44, m.Instance(), PascalBool(Wait))
}

func (m *TApplication) DisableIdleHandler() {
	applicationImportAPI().SysCallN(14, m.Instance())
}

func (m *TApplication) EnableIdleHandler() {
	applicationImportAPI().SysCallN(21, m.Instance())
}

func (m *TApplication) NotifyUserInputHandler(Sender IObject, Msg uint32) {
	applicationImportAPI().SysCallN(71, m.Instance(), GetObjectUintptr(Sender), uintptr(Msg))
}

func (m *TApplication) NotifyKeyDownBeforeHandler(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	applicationImportAPI().SysCallN(69, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) NotifyKeyDownHandler(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	applicationImportAPI().SysCallN(70, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) ControlKeyDown(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	applicationImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) ControlKeyUp(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	applicationImportAPI().SysCallN(11, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) RemoveAllHandlersOfObject(AnObject IObject) {
	applicationImportAPI().SysCallN(74, m.Instance(), GetObjectUintptr(AnObject))
}

func (m *TApplication) DoBeforeMouseMessage(CurMouseControl IControl) {
	applicationImportAPI().SysCallN(16, m.Instance(), GetObjectUintptr(CurMouseControl))
}

func (m *TApplication) IntfEndSession() {
	applicationImportAPI().SysCallN(49, m.Instance())
}

func (m *TApplication) IntfAppActivate(Async bool) {
	applicationImportAPI().SysCallN(45, m.Instance(), PascalBool(Async))
}

func (m *TApplication) IntfAppDeactivate(Async bool) {
	applicationImportAPI().SysCallN(46, m.Instance(), PascalBool(Async))
}

func (m *TApplication) IntfAppMinimize() {
	applicationImportAPI().SysCallN(47, m.Instance())
}

func (m *TApplication) IntfAppRestore() {
	applicationImportAPI().SysCallN(48, m.Instance())
}

func (m *TApplication) IntfSettingsChange() {
	applicationImportAPI().SysCallN(50, m.Instance())
}

func (m *TApplication) IntfThemeOptionChange(AThemeServices IThemeServices, AOption TThemeOption) {
	applicationImportAPI().SysCallN(51, m.Instance(), GetObjectUintptr(AThemeServices), uintptr(AOption))
}

func (m *TApplication) DoArrowKey(AControl IWinControl, Key *Word, Shift TShiftState) {
	var result1 uintptr
	applicationImportAPI().SysCallN(15, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) DoTabKey(AControl IWinControl, Key *Word, Shift TShiftState) {
	var result1 uintptr
	applicationImportAPI().SysCallN(19, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) DoEscapeKey(AControl IWinControl, Key *Word, Shift TShiftState) {
	var result1 uintptr
	applicationImportAPI().SysCallN(17, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) DoReturnKey(AControl IWinControl, Key *Word, Shift TShiftState) {
	var result1 uintptr
	applicationImportAPI().SysCallN(18, m.Instance(), GetObjectUintptr(AControl), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TApplication) SetOnActionExecute(fn TActionEvent) {
	if m.actionExecutePtr != 0 {
		RemoveEventElement(m.actionExecutePtr)
	}
	m.actionExecutePtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(80, m.Instance(), m.actionExecutePtr)
}

func (m *TApplication) SetOnActionUpdate(fn TActionEvent) {
	if m.actionUpdatePtr != 0 {
		RemoveEventElement(m.actionUpdatePtr)
	}
	m.actionUpdatePtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(81, m.Instance(), m.actionUpdatePtr)
}

func (m *TApplication) SetOnActivate(fn TNotifyEvent) {
	if m.activatePtr != 0 {
		RemoveEventElement(m.activatePtr)
	}
	m.activatePtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(82, m.Instance(), m.activatePtr)
}

func (m *TApplication) SetOnDeactivate(fn TNotifyEvent) {
	if m.deactivatePtr != 0 {
		RemoveEventElement(m.deactivatePtr)
	}
	m.deactivatePtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(84, m.Instance(), m.deactivatePtr)
}

func (m *TApplication) SetOnGetMainFormHandle(fn TGetHandleEvent) {
	if m.getMainFormHandlePtr != 0 {
		RemoveEventElement(m.getMainFormHandlePtr)
	}
	m.getMainFormHandlePtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(88, m.Instance(), m.getMainFormHandlePtr)
}

func (m *TApplication) SetOnIdleEnd(fn TNotifyEvent) {
	if m.idleEndPtr != 0 {
		RemoveEventElement(m.idleEndPtr)
	}
	m.idleEndPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(91, m.Instance(), m.idleEndPtr)
}

func (m *TApplication) SetOnEndSession(fn TNotifyEvent) {
	if m.endSessionPtr != 0 {
		RemoveEventElement(m.endSessionPtr)
	}
	m.endSessionPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(87, m.Instance(), m.endSessionPtr)
}

func (m *TApplication) SetOnMinimize(fn TNotifyEvent) {
	if m.minimizePtr != 0 {
		RemoveEventElement(m.minimizePtr)
	}
	m.minimizePtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(92, m.Instance(), m.minimizePtr)
}

func (m *TApplication) SetOnModalBegin(fn TNotifyEvent) {
	if m.modalBeginPtr != 0 {
		RemoveEventElement(m.modalBeginPtr)
	}
	m.modalBeginPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(93, m.Instance(), m.modalBeginPtr)
}

func (m *TApplication) SetOnModalEnd(fn TNotifyEvent) {
	if m.modalEndPtr != 0 {
		RemoveEventElement(m.modalEndPtr)
	}
	m.modalEndPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(94, m.Instance(), m.modalEndPtr)
}

func (m *TApplication) SetOnRestore(fn TNotifyEvent) {
	if m.restorePtr != 0 {
		RemoveEventElement(m.restorePtr)
	}
	m.restorePtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(95, m.Instance(), m.restorePtr)
}

func (m *TApplication) SetOnDropFiles(fn TDropFilesEvent) {
	if m.dropFilesPtr != 0 {
		RemoveEventElement(m.dropFilesPtr)
	}
	m.dropFilesPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(86, m.Instance(), m.dropFilesPtr)
}

func (m *TApplication) SetOnHelp(fn THelpEvent) {
	if m.helpPtr != 0 {
		RemoveEventElement(m.helpPtr)
	}
	m.helpPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(89, m.Instance(), m.helpPtr)
}

func (m *TApplication) SetOnHint(fn TNotifyEvent) {
	if m.hintPtr != 0 {
		RemoveEventElement(m.hintPtr)
	}
	m.hintPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(90, m.Instance(), m.hintPtr)
}

func (m *TApplication) SetOnShortcut(fn TShortCutEvent) {
	if m.shortcutPtr != 0 {
		RemoveEventElement(m.shortcutPtr)
	}
	m.shortcutPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(96, m.Instance(), m.shortcutPtr)
}

func (m *TApplication) SetOnShowHint(fn TShowHintEvent) {
	if m.showHintPtr != 0 {
		RemoveEventElement(m.showHintPtr)
	}
	m.showHintPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(97, m.Instance(), m.showHintPtr)
}

func (m *TApplication) SetOnDestroy(fn TNotifyEvent) {
	if m.destroyPtr != 0 {
		RemoveEventElement(m.destroyPtr)
	}
	m.destroyPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(85, m.Instance(), m.destroyPtr)
}

func (m *TApplication) SetOnCircularException(fn TExceptionEvent) {
	if m.circularExceptionPtr != 0 {
		RemoveEventElement(m.circularExceptionPtr)
	}
	m.circularExceptionPtr = MakeEventDataPtr(fn)
	applicationImportAPI().SysCallN(83, m.Instance(), m.circularExceptionPtr)
}

var (
	applicationImport       *imports.Imports = nil
	applicationImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Application_ActivateHint", 0),
		/*1*/ imports.NewTable("Application_Active", 0),
		/*2*/ imports.NewTable("Application_ApplicationType", 0),
		/*3*/ imports.NewTable("Application_BidiMode", 0),
		/*4*/ imports.NewTable("Application_BigIconHandle", 0),
		/*5*/ imports.NewTable("Application_BringToFront", 0),
		/*6*/ imports.NewTable("Application_CancelHint", 0),
		/*7*/ imports.NewTable("Application_CaptureExceptions", 0),
		/*8*/ imports.NewTable("Application_Class", 0),
		/*9*/ imports.NewTable("Application_ControlDestroyed", 0),
		/*10*/ imports.NewTable("Application_ControlKeyDown", 0),
		/*11*/ imports.NewTable("Application_ControlKeyUp", 0),
		/*12*/ imports.NewTable("Application_Create", 0),
		/*13*/ imports.NewTable("Application_Direction", 0),
		/*14*/ imports.NewTable("Application_DisableIdleHandler", 0),
		/*15*/ imports.NewTable("Application_DoArrowKey", 0),
		/*16*/ imports.NewTable("Application_DoBeforeMouseMessage", 0),
		/*17*/ imports.NewTable("Application_DoEscapeKey", 0),
		/*18*/ imports.NewTable("Application_DoReturnKey", 0),
		/*19*/ imports.NewTable("Application_DoTabKey", 0),
		/*20*/ imports.NewTable("Application_DoubleBuffered", 0),
		/*21*/ imports.NewTable("Application_EnableIdleHandler", 0),
		/*22*/ imports.NewTable("Application_ExceptionDialog", 0),
		/*23*/ imports.NewTable("Application_ExtendedKeysSupport", 0),
		/*24*/ imports.NewTable("Application_FindGlobalComponentEnabled", 0),
		/*25*/ imports.NewTable("Application_Flags", 0),
		/*26*/ imports.NewTable("Application_GetControlAtMouse", 0),
		/*27*/ imports.NewTable("Application_GetControlAtPos", 0),
		/*28*/ imports.NewTable("Application_Handle", 0),
		/*29*/ imports.NewTable("Application_HandleMessage", 0),
		/*30*/ imports.NewTable("Application_HelpCommand", 0),
		/*31*/ imports.NewTable("Application_HelpContext", 0),
		/*32*/ imports.NewTable("Application_HelpKeyword", 0),
		/*33*/ imports.NewTable("Application_HelpShowTableOfContents", 0),
		/*34*/ imports.NewTable("Application_HideHint", 0),
		/*35*/ imports.NewTable("Application_Hint", 0),
		/*36*/ imports.NewTable("Application_HintColor", 0),
		/*37*/ imports.NewTable("Application_HintHidePause", 0),
		/*38*/ imports.NewTable("Application_HintHidePausePerChar", 0),
		/*39*/ imports.NewTable("Application_HintMouseMessage", 0),
		/*40*/ imports.NewTable("Application_HintPause", 0),
		/*41*/ imports.NewTable("Application_HintShortCuts", 0),
		/*42*/ imports.NewTable("Application_HintShortPause", 0),
		/*43*/ imports.NewTable("Application_Icon", 0),
		/*44*/ imports.NewTable("Application_Idle", 0),
		/*45*/ imports.NewTable("Application_IntfAppActivate", 0),
		/*46*/ imports.NewTable("Application_IntfAppDeactivate", 0),
		/*47*/ imports.NewTable("Application_IntfAppMinimize", 0),
		/*48*/ imports.NewTable("Application_IntfAppRestore", 0),
		/*49*/ imports.NewTable("Application_IntfEndSession", 0),
		/*50*/ imports.NewTable("Application_IntfSettingsChange", 0),
		/*51*/ imports.NewTable("Application_IntfThemeOptionChange", 0),
		/*52*/ imports.NewTable("Application_IsRTLLang", 0),
		/*53*/ imports.NewTable("Application_IsRightToLeft", 0),
		/*54*/ imports.NewTable("Application_IsShortcut", 0),
		/*55*/ imports.NewTable("Application_IsWaiting", 0),
		/*56*/ imports.NewTable("Application_LayoutAdjustmentPolicy", 0),
		/*57*/ imports.NewTable("Application_MainForm", 0),
		/*58*/ imports.NewTable("Application_MainFormHandle", 0),
		/*59*/ imports.NewTable("Application_MainFormOnTaskBar", 0),
		/*60*/ imports.NewTable("Application_MessageBox", 0),
		/*61*/ imports.NewTable("Application_Minimize", 0),
		/*62*/ imports.NewTable("Application_ModalFinished", 0),
		/*63*/ imports.NewTable("Application_ModalLevel", 0),
		/*64*/ imports.NewTable("Application_ModalStarted", 0),
		/*65*/ imports.NewTable("Application_MouseControl", 0),
		/*66*/ imports.NewTable("Application_MoveFormFocusToChildren", 0),
		/*67*/ imports.NewTable("Application_Navigation", 0),
		/*68*/ imports.NewTable("Application_Notification", 0),
		/*69*/ imports.NewTable("Application_NotifyKeyDownBeforeHandler", 0),
		/*70*/ imports.NewTable("Application_NotifyKeyDownHandler", 0),
		/*71*/ imports.NewTable("Application_NotifyUserInputHandler", 0),
		/*72*/ imports.NewTable("Application_ProcessMessages", 0),
		/*73*/ imports.NewTable("Application_ReleaseComponent", 0),
		/*74*/ imports.NewTable("Application_RemoveAllHandlersOfObject", 0),
		/*75*/ imports.NewTable("Application_RemoveAsyncCalls", 0),
		/*76*/ imports.NewTable("Application_RemoveStayOnTop", 0),
		/*77*/ imports.NewTable("Application_Restore", 0),
		/*78*/ imports.NewTable("Application_RestoreStayOnTop", 0),
		/*79*/ imports.NewTable("Application_Scaled", 0),
		/*80*/ imports.NewTable("Application_SetOnActionExecute", 0),
		/*81*/ imports.NewTable("Application_SetOnActionUpdate", 0),
		/*82*/ imports.NewTable("Application_SetOnActivate", 0),
		/*83*/ imports.NewTable("Application_SetOnCircularException", 0),
		/*84*/ imports.NewTable("Application_SetOnDeactivate", 0),
		/*85*/ imports.NewTable("Application_SetOnDestroy", 0),
		/*86*/ imports.NewTable("Application_SetOnDropFiles", 0),
		/*87*/ imports.NewTable("Application_SetOnEndSession", 0),
		/*88*/ imports.NewTable("Application_SetOnGetMainFormHandle", 0),
		/*89*/ imports.NewTable("Application_SetOnHelp", 0),
		/*90*/ imports.NewTable("Application_SetOnHint", 0),
		/*91*/ imports.NewTable("Application_SetOnIdleEnd", 0),
		/*92*/ imports.NewTable("Application_SetOnMinimize", 0),
		/*93*/ imports.NewTable("Application_SetOnModalBegin", 0),
		/*94*/ imports.NewTable("Application_SetOnModalEnd", 0),
		/*95*/ imports.NewTable("Application_SetOnRestore", 0),
		/*96*/ imports.NewTable("Application_SetOnShortcut", 0),
		/*97*/ imports.NewTable("Application_SetOnShowHint", 0),
		/*98*/ imports.NewTable("Application_ShowButtonGlyphs", 0),
		/*99*/ imports.NewTable("Application_ShowHelpForObject", 0),
		/*100*/ imports.NewTable("Application_ShowHint", 0),
		/*101*/ imports.NewTable("Application_ShowMainForm", 0),
		/*102*/ imports.NewTable("Application_ShowMenuGlyphs", 0),
		/*103*/ imports.NewTable("Application_SmallIconHandle", 0),
		/*104*/ imports.NewTable("Application_TaskBarBehavior", 0),
		/*105*/ imports.NewTable("Application_UpdateFormatSettings", 0),
		/*106*/ imports.NewTable("Application_UpdateMainForm", 0),
	}
)

func applicationImportAPI() *imports.Imports {
	if applicationImport == nil {
		applicationImport = NewDefaultImports()
		applicationImport.SetImportTable(applicationImportTables)
		applicationImportTables = nil
	}
	return applicationImport
}
