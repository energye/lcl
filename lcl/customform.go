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

// ICustomForm Parent: ICustomDesignControl
type ICustomForm interface {
	ICustomDesignControl
	Active() bool                                                                // property
	ActiveControl() IWinControl                                                  // property
	SetActiveControl(AValue IWinControl)                                         // property
	ActiveDefaultControl() IControl                                              // property
	SetActiveDefaultControl(AValue IControl)                                     // property
	AllowDropFiles() bool                                                        // property
	SetAllowDropFiles(AValue bool)                                               // property
	AlphaBlend() bool                                                            // property
	SetAlphaBlend(AValue bool)                                                   // property
	AlphaBlendValue() Byte                                                       // property
	SetAlphaBlendValue(AValue Byte)                                              // property
	AutoScroll() bool                                                            // property
	SetAutoScroll(AValue bool)                                                   // property
	BorderIcons() TBorderIcons                                                   // property
	SetBorderIcons(AValue TBorderIcons)                                          // property
	BorderStyleForFormBorderStyle() TFormBorderStyle                             // property
	SetBorderStyleForFormBorderStyle(AValue TFormBorderStyle)                    // property
	CancelControl() IControl                                                     // property
	SetCancelControl(AValue IControl)                                            // property
	DefaultControl() IControl                                                    // property
	SetDefaultControl(AValue IControl)                                           // property
	DefaultMonitor() TDefaultMonitor                                             // property
	SetDefaultMonitor(AValue TDefaultMonitor)                                    // property
	Designer() IIDesigner                                                        // property
	SetDesigner(AValue IIDesigner)                                               // property
	EffectiveShowInTaskBar() TShowInTaskBar                                      // property
	FormState() TFormState                                                       // property
	FormStyle() TFormStyle                                                       // property
	SetFormStyle(AValue TFormStyle)                                              // property
	HelpFile() string                                                            // property
	SetHelpFile(AValue string)                                                   // property
	Icon() IIcon                                                                 // property
	SetIcon(AValue IIcon)                                                        // property
	KeyPreview() bool                                                            // property
	SetKeyPreview(AValue bool)                                                   // property
	MDIChildren(I int32) ICustomForm                                             // property
	Menu() IMainMenu                                                             // property
	SetMenu(AValue IMainMenu)                                                    // property
	ModalResult() TModalResult                                                   // property
	SetModalResult(AValue TModalResult)                                          // property
	Monitor() IMonitor                                                           // property
	LastActiveControl() IWinControl                                              // property
	PopupMode() TPopupMode                                                       // property
	SetPopupMode(AValue TPopupMode)                                              // property
	PopupParent() ICustomForm                                                    // property
	SetPopupParent(AValue ICustomForm)                                           // property
	SnapOptions() IWindowMagnetOptions                                           // property
	SetSnapOptions(AValue IWindowMagnetOptions)                                  // property
	ScreenSnap() bool                                                            // property
	SetScreenSnap(AValue bool)                                                   // property
	SnapBuffer() int32                                                           // property
	SetSnapBuffer(AValue int32)                                                  // property
	ParentFont() bool                                                            // property
	SetParentFont(AValue bool)                                                   // property
	Position() TPosition                                                         // property
	SetPosition(AValue TPosition)                                                // property
	RestoredLeft() int32                                                         // property
	RestoredTop() int32                                                          // property
	RestoredWidth() int32                                                        // property
	RestoredHeight() int32                                                       // property
	ShowInTaskBar() TShowInTaskBar                                               // property
	SetShowInTaskBar(AValue TShowInTaskBar)                                      // property
	WindowState() TWindowState                                                   // property
	SetWindowState(AValue TWindowState)                                          // property
	BigIconHandle() HICON                                                        // function
	CloseQuery() bool                                                            // function
	GetFormImage() IBitmap                                                       // function
	GetRolesForControl(AControl IControl) TControlRolesForForm                   // function
	GetRealPopupParent() ICustomForm                                             // function
	IsShortcut(Message *TLMKey) bool                                             // function
	SetFocusedControl(Control IWinControl) bool                                  // function
	ShowModal() int32                                                            // function
	SmallIconHandle() HICON                                                      // function
	WantChildKey(Child IControl, Message *TLMessage) bool                        // function
	ActiveMDIChild() ICustomForm                                                 // function
	GetMDIChildren(AIndex int32) ICustomForm                                     // function
	MDIChildCount() int32                                                        // function
	Close()                                                                      // procedure
	DefocusControl(Control IWinControl, Removing bool)                           // procedure
	DestroyWnd()                                                                 // procedure
	EnsureVisible(AMoveToTop bool)                                               // procedure
	FocusControl(WinControl IWinControl)                                         // procedure
	GetFormImage1(ABitmap ICustomBitmap)                                         // procedure
	IntfHelp(AComponent IComponent)                                              // procedure
	MakeFullyVisible(AMonitor IMonitor, UseWorkarea bool)                        // procedure
	Release()                                                                    // procedure
	SetRestoredBounds(ALeft, ATop, AWidth, AHeight int32, ADefaultPosition bool) // procedure
	ShowOnTop()                                                                  // procedure
	AutoScale()                                                                  // procedure
	SetOnActivate(fn TNotifyEvent)                                               // property event
	SetOnClose(fn TCloseEvent)                                                   // property event
	SetOnCloseQuery(fn TCloseQueryEvent)                                         // property event
	SetOnCreate(fn TNotifyEvent)                                                 // property event
	SetOnDeactivate(fn TNotifyEvent)                                             // property event
	SetOnDestroy(fn TNotifyEvent)                                                // property event
	SetOnDropFiles(fn TDropFilesEvent)                                           // property event
	SetOnHelp(fn THelpEvent)                                                     // property event
	SetOnHide(fn TNotifyEvent)                                                   // property event
	SetOnShortcut(fn TShortCutEvent)                                             // property event
	SetOnShow(fn TNotifyEvent)                                                   // property event
	SetOnShowModalFinished(fn TModalDialogFinished)                              // property event
	SetOnWindowStateChange(fn TNotifyEvent)                                      // property event
}

// TCustomForm Parent: TCustomDesignControl
type TCustomForm struct {
	TCustomDesignControl
	activatePtr          uintptr
	closePtr             uintptr
	closeQueryPtr        uintptr
	createPtr            uintptr
	deactivatePtr        uintptr
	destroyPtr           uintptr
	dropFilesPtr         uintptr
	helpPtr              uintptr
	hidePtr              uintptr
	shortcutPtr          uintptr
	showPtr              uintptr
	showModalFinishedPtr uintptr
	windowStateChangePtr uintptr
}

func NewCustomForm(AOwner IComponent) ICustomForm {
	r1 := customFormImportAPI().SysCallN(16, GetObjectUintptr(AOwner))
	return AsCustomForm(r1)
}

func NewCustomFormNew(AOwner IComponent, Num int32) ICustomForm {
	r1 := customFormImportAPI().SysCallN(17, GetObjectUintptr(AOwner), uintptr(Num))
	return AsCustomForm(r1)
}

func (m *TCustomForm) Active() bool {
	r1 := customFormImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCustomForm) ActiveControl() IWinControl {
	r1 := customFormImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TCustomForm) SetActiveControl(AValue IWinControl) {
	customFormImportAPI().SysCallN(1, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) ActiveDefaultControl() IControl {
	r1 := customFormImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomForm) SetActiveDefaultControl(AValue IControl) {
	customFormImportAPI().SysCallN(2, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) AllowDropFiles() bool {
	r1 := customFormImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetAllowDropFiles(AValue bool) {
	customFormImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) AlphaBlend() bool {
	r1 := customFormImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetAlphaBlend(AValue bool) {
	customFormImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) AlphaBlendValue() Byte {
	r1 := customFormImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TCustomForm) SetAlphaBlendValue(AValue Byte) {
	customFormImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) AutoScroll() bool {
	r1 := customFormImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetAutoScroll(AValue bool) {
	customFormImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) BorderIcons() TBorderIcons {
	r1 := customFormImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TBorderIcons(r1)
}

func (m *TCustomForm) SetBorderIcons(AValue TBorderIcons) {
	customFormImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) BorderStyleForFormBorderStyle() TFormBorderStyle {
	r1 := customFormImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TFormBorderStyle(r1)
}

func (m *TCustomForm) SetBorderStyleForFormBorderStyle(AValue TFormBorderStyle) {
	customFormImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) CancelControl() IControl {
	r1 := customFormImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomForm) SetCancelControl(AValue IControl) {
	customFormImportAPI().SysCallN(12, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) DefaultControl() IControl {
	r1 := customFormImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomForm) SetDefaultControl(AValue IControl) {
	customFormImportAPI().SysCallN(18, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) DefaultMonitor() TDefaultMonitor {
	r1 := customFormImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return TDefaultMonitor(r1)
}

func (m *TCustomForm) SetDefaultMonitor(AValue TDefaultMonitor) {
	customFormImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) Designer() IIDesigner {
	r1 := customFormImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return AsIDesigner(r1)
}

func (m *TCustomForm) SetDesigner(AValue IIDesigner) {
	customFormImportAPI().SysCallN(21, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) EffectiveShowInTaskBar() TShowInTaskBar {
	r1 := customFormImportAPI().SysCallN(23, m.Instance())
	return TShowInTaskBar(r1)
}

func (m *TCustomForm) FormState() TFormState {
	r1 := customFormImportAPI().SysCallN(26, m.Instance())
	return TFormState(r1)
}

func (m *TCustomForm) FormStyle() TFormStyle {
	r1 := customFormImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return TFormStyle(r1)
}

func (m *TCustomForm) SetFormStyle(AValue TFormStyle) {
	customFormImportAPI().SysCallN(27, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) HelpFile() string {
	r1 := customFormImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomForm) SetHelpFile(AValue string) {
	customFormImportAPI().SysCallN(33, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomForm) Icon() IIcon {
	r1 := customFormImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return AsIcon(r1)
}

func (m *TCustomForm) SetIcon(AValue IIcon) {
	customFormImportAPI().SysCallN(34, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) KeyPreview() bool {
	r1 := customFormImportAPI().SysCallN(37, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetKeyPreview(AValue bool) {
	customFormImportAPI().SysCallN(37, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) MDIChildren(I int32) ICustomForm {
	r1 := customFormImportAPI().SysCallN(40, m.Instance(), uintptr(I))
	return AsCustomForm(r1)
}

func (m *TCustomForm) Menu() IMainMenu {
	r1 := customFormImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return AsMainMenu(r1)
}

func (m *TCustomForm) SetMenu(AValue IMainMenu) {
	customFormImportAPI().SysCallN(42, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) ModalResult() TModalResult {
	r1 := customFormImportAPI().SysCallN(43, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TCustomForm) SetModalResult(AValue TModalResult) {
	customFormImportAPI().SysCallN(43, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) Monitor() IMonitor {
	r1 := customFormImportAPI().SysCallN(44, m.Instance())
	return AsMonitor(r1)
}

func (m *TCustomForm) LastActiveControl() IWinControl {
	r1 := customFormImportAPI().SysCallN(38, m.Instance())
	return AsWinControl(r1)
}

func (m *TCustomForm) PopupMode() TPopupMode {
	r1 := customFormImportAPI().SysCallN(46, 0, m.Instance(), 0)
	return TPopupMode(r1)
}

func (m *TCustomForm) SetPopupMode(AValue TPopupMode) {
	customFormImportAPI().SysCallN(46, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) PopupParent() ICustomForm {
	r1 := customFormImportAPI().SysCallN(47, 0, m.Instance(), 0)
	return AsCustomForm(r1)
}

func (m *TCustomForm) SetPopupParent(AValue ICustomForm) {
	customFormImportAPI().SysCallN(47, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) SnapOptions() IWindowMagnetOptions {
	r1 := customFormImportAPI().SysCallN(75, 0, m.Instance(), 0)
	return AsWindowMagnetOptions(r1)
}

func (m *TCustomForm) SetSnapOptions(AValue IWindowMagnetOptions) {
	customFormImportAPI().SysCallN(75, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomForm) ScreenSnap() bool {
	r1 := customFormImportAPI().SysCallN(54, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetScreenSnap(AValue bool) {
	customFormImportAPI().SysCallN(54, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) SnapBuffer() int32 {
	r1 := customFormImportAPI().SysCallN(74, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomForm) SetSnapBuffer(AValue int32) {
	customFormImportAPI().SysCallN(74, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) ParentFont() bool {
	r1 := customFormImportAPI().SysCallN(45, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomForm) SetParentFont(AValue bool) {
	customFormImportAPI().SysCallN(45, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomForm) Position() TPosition {
	r1 := customFormImportAPI().SysCallN(48, 0, m.Instance(), 0)
	return TPosition(r1)
}

func (m *TCustomForm) SetPosition(AValue TPosition) {
	customFormImportAPI().SysCallN(48, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) RestoredLeft() int32 {
	r1 := customFormImportAPI().SysCallN(51, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) RestoredTop() int32 {
	r1 := customFormImportAPI().SysCallN(52, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) RestoredWidth() int32 {
	r1 := customFormImportAPI().SysCallN(53, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) RestoredHeight() int32 {
	r1 := customFormImportAPI().SysCallN(50, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) ShowInTaskBar() TShowInTaskBar {
	r1 := customFormImportAPI().SysCallN(70, 0, m.Instance(), 0)
	return TShowInTaskBar(r1)
}

func (m *TCustomForm) SetShowInTaskBar(AValue TShowInTaskBar) {
	customFormImportAPI().SysCallN(70, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) WindowState() TWindowState {
	r1 := customFormImportAPI().SysCallN(77, 0, m.Instance(), 0)
	return TWindowState(r1)
}

func (m *TCustomForm) SetWindowState(AValue TWindowState) {
	customFormImportAPI().SysCallN(77, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomForm) BigIconHandle() HICON {
	r1 := customFormImportAPI().SysCallN(9, m.Instance())
	return HICON(r1)
}

func (m *TCustomForm) CloseQuery() bool {
	r1 := customFormImportAPI().SysCallN(15, m.Instance())
	return GoBool(r1)
}

func (m *TCustomForm) GetFormImage() IBitmap {
	r1 := customFormImportAPI().SysCallN(28, m.Instance())
	return AsBitmap(r1)
}

func (m *TCustomForm) GetRolesForControl(AControl IControl) TControlRolesForForm {
	r1 := customFormImportAPI().SysCallN(32, m.Instance(), GetObjectUintptr(AControl))
	return TControlRolesForForm(r1)
}

func (m *TCustomForm) GetRealPopupParent() ICustomForm {
	r1 := customFormImportAPI().SysCallN(31, m.Instance())
	return AsCustomForm(r1)
}

func (m *TCustomForm) IsShortcut(Message *TLMKey) bool {
	var result0 uintptr
	r1 := customFormImportAPI().SysCallN(36, m.Instance(), uintptr(unsafePointer(&result0)))
	*Message = *(*TLMKey)(getPointer(result0))
	return GoBool(r1)
}

func (m *TCustomForm) SetFocusedControl(Control IWinControl) bool {
	r1 := customFormImportAPI().SysCallN(55, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TCustomForm) ShowModal() int32 {
	r1 := customFormImportAPI().SysCallN(71, m.Instance())
	return int32(r1)
}

func (m *TCustomForm) SmallIconHandle() HICON {
	r1 := customFormImportAPI().SysCallN(73, m.Instance())
	return HICON(r1)
}

func (m *TCustomForm) WantChildKey(Child IControl, Message *TLMessage) bool {
	var result1 uintptr
	r1 := customFormImportAPI().SysCallN(76, m.Instance(), GetObjectUintptr(Child), uintptr(unsafePointer(&result1)))
	*Message = *(*TLMessage)(getPointer(result1))
	return GoBool(r1)
}

func (m *TCustomForm) ActiveMDIChild() ICustomForm {
	r1 := customFormImportAPI().SysCallN(3, m.Instance())
	return AsCustomForm(r1)
}

func (m *TCustomForm) GetMDIChildren(AIndex int32) ICustomForm {
	r1 := customFormImportAPI().SysCallN(30, m.Instance(), uintptr(AIndex))
	return AsCustomForm(r1)
}

func (m *TCustomForm) MDIChildCount() int32 {
	r1 := customFormImportAPI().SysCallN(39, m.Instance())
	return int32(r1)
}

func CustomFormClass() TClass {
	ret := customFormImportAPI().SysCallN(13)
	return TClass(ret)
}

func (m *TCustomForm) Close() {
	customFormImportAPI().SysCallN(14, m.Instance())
}

func (m *TCustomForm) DefocusControl(Control IWinControl, Removing bool) {
	customFormImportAPI().SysCallN(20, m.Instance(), GetObjectUintptr(Control), PascalBool(Removing))
}

func (m *TCustomForm) DestroyWnd() {
	customFormImportAPI().SysCallN(22, m.Instance())
}

func (m *TCustomForm) EnsureVisible(AMoveToTop bool) {
	customFormImportAPI().SysCallN(24, m.Instance(), PascalBool(AMoveToTop))
}

func (m *TCustomForm) FocusControl(WinControl IWinControl) {
	customFormImportAPI().SysCallN(25, m.Instance(), GetObjectUintptr(WinControl))
}

func (m *TCustomForm) GetFormImage1(ABitmap ICustomBitmap) {
	customFormImportAPI().SysCallN(29, m.Instance(), GetObjectUintptr(ABitmap))
}

func (m *TCustomForm) IntfHelp(AComponent IComponent) {
	customFormImportAPI().SysCallN(35, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TCustomForm) MakeFullyVisible(AMonitor IMonitor, UseWorkarea bool) {
	customFormImportAPI().SysCallN(41, m.Instance(), GetObjectUintptr(AMonitor), PascalBool(UseWorkarea))
}

func (m *TCustomForm) Release() {
	customFormImportAPI().SysCallN(49, m.Instance())
}

func (m *TCustomForm) SetRestoredBounds(ALeft, ATop, AWidth, AHeight int32, ADefaultPosition bool) {
	customFormImportAPI().SysCallN(69, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight), PascalBool(ADefaultPosition))
}

func (m *TCustomForm) ShowOnTop() {
	customFormImportAPI().SysCallN(72, m.Instance())
}

func (m *TCustomForm) AutoScale() {
	customFormImportAPI().SysCallN(7, m.Instance())
}

func (m *TCustomForm) SetOnActivate(fn TNotifyEvent) {
	if m.activatePtr != 0 {
		RemoveEventElement(m.activatePtr)
	}
	m.activatePtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(56, m.Instance(), m.activatePtr)
}

func (m *TCustomForm) SetOnClose(fn TCloseEvent) {
	if m.closePtr != 0 {
		RemoveEventElement(m.closePtr)
	}
	m.closePtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(57, m.Instance(), m.closePtr)
}

func (m *TCustomForm) SetOnCloseQuery(fn TCloseQueryEvent) {
	if m.closeQueryPtr != 0 {
		RemoveEventElement(m.closeQueryPtr)
	}
	m.closeQueryPtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(58, m.Instance(), m.closeQueryPtr)
}

func (m *TCustomForm) SetOnCreate(fn TNotifyEvent) {
	if m.createPtr != 0 {
		RemoveEventElement(m.createPtr)
	}
	m.createPtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(59, m.Instance(), m.createPtr)
}

func (m *TCustomForm) SetOnDeactivate(fn TNotifyEvent) {
	if m.deactivatePtr != 0 {
		RemoveEventElement(m.deactivatePtr)
	}
	m.deactivatePtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(60, m.Instance(), m.deactivatePtr)
}

func (m *TCustomForm) SetOnDestroy(fn TNotifyEvent) {
	if m.destroyPtr != 0 {
		RemoveEventElement(m.destroyPtr)
	}
	m.destroyPtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(61, m.Instance(), m.destroyPtr)
}

func (m *TCustomForm) SetOnDropFiles(fn TDropFilesEvent) {
	if m.dropFilesPtr != 0 {
		RemoveEventElement(m.dropFilesPtr)
	}
	m.dropFilesPtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(62, m.Instance(), m.dropFilesPtr)
}

func (m *TCustomForm) SetOnHelp(fn THelpEvent) {
	if m.helpPtr != 0 {
		RemoveEventElement(m.helpPtr)
	}
	m.helpPtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(63, m.Instance(), m.helpPtr)
}

func (m *TCustomForm) SetOnHide(fn TNotifyEvent) {
	if m.hidePtr != 0 {
		RemoveEventElement(m.hidePtr)
	}
	m.hidePtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(64, m.Instance(), m.hidePtr)
}

func (m *TCustomForm) SetOnShortcut(fn TShortCutEvent) {
	if m.shortcutPtr != 0 {
		RemoveEventElement(m.shortcutPtr)
	}
	m.shortcutPtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(65, m.Instance(), m.shortcutPtr)
}

func (m *TCustomForm) SetOnShow(fn TNotifyEvent) {
	if m.showPtr != 0 {
		RemoveEventElement(m.showPtr)
	}
	m.showPtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(66, m.Instance(), m.showPtr)
}

func (m *TCustomForm) SetOnShowModalFinished(fn TModalDialogFinished) {
	if m.showModalFinishedPtr != 0 {
		RemoveEventElement(m.showModalFinishedPtr)
	}
	m.showModalFinishedPtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(67, m.Instance(), m.showModalFinishedPtr)
}

func (m *TCustomForm) SetOnWindowStateChange(fn TNotifyEvent) {
	if m.windowStateChangePtr != 0 {
		RemoveEventElement(m.windowStateChangePtr)
	}
	m.windowStateChangePtr = MakeEventDataPtr(fn)
	customFormImportAPI().SysCallN(68, m.Instance(), m.windowStateChangePtr)
}

var (
	customFormImport       *imports.Imports = nil
	customFormImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomForm_Active", 0),
		/*1*/ imports.NewTable("CustomForm_ActiveControl", 0),
		/*2*/ imports.NewTable("CustomForm_ActiveDefaultControl", 0),
		/*3*/ imports.NewTable("CustomForm_ActiveMDIChild", 0),
		/*4*/ imports.NewTable("CustomForm_AllowDropFiles", 0),
		/*5*/ imports.NewTable("CustomForm_AlphaBlend", 0),
		/*6*/ imports.NewTable("CustomForm_AlphaBlendValue", 0),
		/*7*/ imports.NewTable("CustomForm_AutoScale", 0),
		/*8*/ imports.NewTable("CustomForm_AutoScroll", 0),
		/*9*/ imports.NewTable("CustomForm_BigIconHandle", 0),
		/*10*/ imports.NewTable("CustomForm_BorderIcons", 0),
		/*11*/ imports.NewTable("CustomForm_BorderStyleForFormBorderStyle", 0),
		/*12*/ imports.NewTable("CustomForm_CancelControl", 0),
		/*13*/ imports.NewTable("CustomForm_Class", 0),
		/*14*/ imports.NewTable("CustomForm_Close", 0),
		/*15*/ imports.NewTable("CustomForm_CloseQuery", 0),
		/*16*/ imports.NewTable("CustomForm_Create", 0),
		/*17*/ imports.NewTable("CustomForm_CreateNew", 0),
		/*18*/ imports.NewTable("CustomForm_DefaultControl", 0),
		/*19*/ imports.NewTable("CustomForm_DefaultMonitor", 0),
		/*20*/ imports.NewTable("CustomForm_DefocusControl", 0),
		/*21*/ imports.NewTable("CustomForm_Designer", 0),
		/*22*/ imports.NewTable("CustomForm_DestroyWnd", 0),
		/*23*/ imports.NewTable("CustomForm_EffectiveShowInTaskBar", 0),
		/*24*/ imports.NewTable("CustomForm_EnsureVisible", 0),
		/*25*/ imports.NewTable("CustomForm_FocusControl", 0),
		/*26*/ imports.NewTable("CustomForm_FormState", 0),
		/*27*/ imports.NewTable("CustomForm_FormStyle", 0),
		/*28*/ imports.NewTable("CustomForm_GetFormImage", 0),
		/*29*/ imports.NewTable("CustomForm_GetFormImage1", 0),
		/*30*/ imports.NewTable("CustomForm_GetMDIChildren", 0),
		/*31*/ imports.NewTable("CustomForm_GetRealPopupParent", 0),
		/*32*/ imports.NewTable("CustomForm_GetRolesForControl", 0),
		/*33*/ imports.NewTable("CustomForm_HelpFile", 0),
		/*34*/ imports.NewTable("CustomForm_Icon", 0),
		/*35*/ imports.NewTable("CustomForm_IntfHelp", 0),
		/*36*/ imports.NewTable("CustomForm_IsShortcut", 0),
		/*37*/ imports.NewTable("CustomForm_KeyPreview", 0),
		/*38*/ imports.NewTable("CustomForm_LastActiveControl", 0),
		/*39*/ imports.NewTable("CustomForm_MDIChildCount", 0),
		/*40*/ imports.NewTable("CustomForm_MDIChildren", 0),
		/*41*/ imports.NewTable("CustomForm_MakeFullyVisible", 0),
		/*42*/ imports.NewTable("CustomForm_Menu", 0),
		/*43*/ imports.NewTable("CustomForm_ModalResult", 0),
		/*44*/ imports.NewTable("CustomForm_Monitor", 0),
		/*45*/ imports.NewTable("CustomForm_ParentFont", 0),
		/*46*/ imports.NewTable("CustomForm_PopupMode", 0),
		/*47*/ imports.NewTable("CustomForm_PopupParent", 0),
		/*48*/ imports.NewTable("CustomForm_Position", 0),
		/*49*/ imports.NewTable("CustomForm_Release", 0),
		/*50*/ imports.NewTable("CustomForm_RestoredHeight", 0),
		/*51*/ imports.NewTable("CustomForm_RestoredLeft", 0),
		/*52*/ imports.NewTable("CustomForm_RestoredTop", 0),
		/*53*/ imports.NewTable("CustomForm_RestoredWidth", 0),
		/*54*/ imports.NewTable("CustomForm_ScreenSnap", 0),
		/*55*/ imports.NewTable("CustomForm_SetFocusedControl", 0),
		/*56*/ imports.NewTable("CustomForm_SetOnActivate", 0),
		/*57*/ imports.NewTable("CustomForm_SetOnClose", 0),
		/*58*/ imports.NewTable("CustomForm_SetOnCloseQuery", 0),
		/*59*/ imports.NewTable("CustomForm_SetOnCreate", 0),
		/*60*/ imports.NewTable("CustomForm_SetOnDeactivate", 0),
		/*61*/ imports.NewTable("CustomForm_SetOnDestroy", 0),
		/*62*/ imports.NewTable("CustomForm_SetOnDropFiles", 0),
		/*63*/ imports.NewTable("CustomForm_SetOnHelp", 0),
		/*64*/ imports.NewTable("CustomForm_SetOnHide", 0),
		/*65*/ imports.NewTable("CustomForm_SetOnShortcut", 0),
		/*66*/ imports.NewTable("CustomForm_SetOnShow", 0),
		/*67*/ imports.NewTable("CustomForm_SetOnShowModalFinished", 0),
		/*68*/ imports.NewTable("CustomForm_SetOnWindowStateChange", 0),
		/*69*/ imports.NewTable("CustomForm_SetRestoredBounds", 0),
		/*70*/ imports.NewTable("CustomForm_ShowInTaskBar", 0),
		/*71*/ imports.NewTable("CustomForm_ShowModal", 0),
		/*72*/ imports.NewTable("CustomForm_ShowOnTop", 0),
		/*73*/ imports.NewTable("CustomForm_SmallIconHandle", 0),
		/*74*/ imports.NewTable("CustomForm_SnapBuffer", 0),
		/*75*/ imports.NewTable("CustomForm_SnapOptions", 0),
		/*76*/ imports.NewTable("CustomForm_WantChildKey", 0),
		/*77*/ imports.NewTable("CustomForm_WindowState", 0),
	}
)

func customFormImportAPI() *imports.Imports {
	if customFormImport == nil {
		customFormImport = NewDefaultImports()
		customFormImport.SetImportTable(customFormImportTables)
		customFormImportTables = nil
	}
	return customFormImport
}
