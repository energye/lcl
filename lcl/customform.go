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

// ICustomForm Parent: ICustomDesignControl
type ICustomForm interface {
	ICustomDesignControl
	BigIconHandle() types.HICON                                     // function
	CloseQuery() bool                                               // function
	GetFormImageToBitmap() IBitmap                                  // function
	GetRolesForControl(control IControl) types.TControlRolesForForm // function
	GetRealPopupParent() ICustomForm                                // function
	IsShortcut(message *types.TLMKey) bool                          // function
	SetFocusedControl(control IWinControl) bool                     // function
	ShowModal() int32                                               // function
	SmallIconHandle() types.HICON                                   // function
	WantChildKey(child IControl, message *types.TLMessage) bool     // function
	// ActiveMDIChild
	//  MDI implementation
	ActiveMDIChild() ICustomForm                                                                    // function
	GetMDIChildren(index int32) ICustomForm                                                         // function
	MDIChildCount() int32                                                                           // function
	Close()                                                                                         // procedure
	DefocusControl(control IWinControl, removing bool)                                              // procedure
	DestroyWnd()                                                                                    // procedure
	EnsureVisible(moveToTop bool)                                                                   // procedure
	FocusControl(winControl IWinControl)                                                            // procedure
	GetFormImageWithCustomBitmap(bitmap ICustomBitmap)                                              // procedure
	Hide()                                                                                          // procedure
	IntfHelp(component IComponent)                                                                  // procedure
	MakeFullyVisible(monitor IMonitor, useWorkarea bool)                                            // procedure
	Release()                                                                                       // procedure
	SetRestoredBounds(left int32, top int32, width int32, height int32, constADefaultPosition bool) // procedure
	Show()                                                                                          // procedure
	ShowOnTop()                                                                                     // procedure
	AutoScale()                                                                                     // procedure
	Active() bool                                                                                   // property Active Getter
	ActiveControl() IWinControl                                                                     // property ActiveControl Getter
	SetActiveControl(value IWinControl)                                                             // property ActiveControl Setter
	ActiveDefaultControl() IControl                                                                 // property ActiveDefaultControl Getter
	SetActiveDefaultControl(value IControl)                                                         // property ActiveDefaultControl Setter
	AllowDropFiles() bool                                                                           // property AllowDropFiles Getter
	SetAllowDropFiles(value bool)                                                                   // property AllowDropFiles Setter
	AlphaBlend() bool                                                                               // property AlphaBlend Getter
	SetAlphaBlend(value bool)                                                                       // property AlphaBlend Setter
	AlphaBlendValue() byte                                                                          // property AlphaBlendValue Getter
	SetAlphaBlendValue(value byte)                                                                  // property AlphaBlendValue Setter
	AutoScroll() bool                                                                               // property AutoScroll Getter
	SetAutoScroll(value bool)                                                                       // property AutoScroll Setter
	BorderIcons() types.TBorderIcons                                                                // property BorderIcons Getter
	SetBorderIcons(value types.TBorderIcons)                                                        // property BorderIcons Setter
	BorderStyleToFormBorderStyle() types.TFormBorderStyle                                           // property BorderStyle Getter
	SetBorderStyleToFormBorderStyle(value types.TFormBorderStyle)                                   // property BorderStyle Setter
	CancelControl() IControl                                                                        // property CancelControl Getter
	SetCancelControl(value IControl)                                                                // property CancelControl Setter
	DefaultControl() IControl                                                                       // property DefaultControl Getter
	SetDefaultControl(value IControl)                                                               // property DefaultControl Setter
	DefaultMonitor() types.TDefaultMonitor                                                          // property DefaultMonitor Getter
	SetDefaultMonitor(value types.TDefaultMonitor)                                                  // property DefaultMonitor Setter
	Designer() IIDesigner                                                                           // property Designer Getter
	SetDesigner(value IIDesigner)                                                                   // property Designer Setter
	EffectiveShowInTaskBar() types.TShowInTaskBar                                                   // property EffectiveShowInTaskBar Getter
	FormState() types.TFormState                                                                    // property FormState Getter
	FormStyle() types.TFormStyle                                                                    // property FormStyle Getter
	SetFormStyle(value types.TFormStyle)                                                            // property FormStyle Setter
	HelpFile() string                                                                               // property HelpFile Getter
	SetHelpFile(value string)                                                                       // property HelpFile Setter
	Icon() IIcon                                                                                    // property Icon Getter
	SetIcon(value IIcon)                                                                            // property Icon Setter
	KeyPreview() bool                                                                               // property KeyPreview Getter
	SetKeyPreview(value bool)                                                                       // property KeyPreview Setter
	MDIChildren(I int32) ICustomForm                                                                // property MDIChildren Getter
	Menu() IMainMenu                                                                                // property Menu Getter
	SetMenu(value IMainMenu)                                                                        // property Menu Setter
	ModalResult() types.TModalResult                                                                // property ModalResult Getter
	SetModalResult(value types.TModalResult)                                                        // property ModalResult Setter
	Monitor() IMonitor                                                                              // property Monitor Getter
	LastActiveControl() IWinControl                                                                 // property LastActiveControl Getter
	PopupMode() types.TPopupMode                                                                    // property PopupMode Getter
	SetPopupMode(value types.TPopupMode)                                                            // property PopupMode Setter
	PopupParent() ICustomForm                                                                       // property PopupParent Getter
	SetPopupParent(value ICustomForm)                                                               // property PopupParent Setter
	SnapOptions() IWindowMagnetOptions                                                              // property SnapOptions Getter
	SetSnapOptions(value IWindowMagnetOptions)                                                      // property SnapOptions Setter
	ScreenSnap() bool                                                                               // property ScreenSnap Getter
	SetScreenSnap(value bool)                                                                       // property ScreenSnap Setter
	SnapBuffer() int32                                                                              // property SnapBuffer Getter
	SetSnapBuffer(value int32)                                                                      // property SnapBuffer Setter
	ParentFont() bool                                                                               // property ParentFont Getter
	SetParentFont(value bool)                                                                       // property ParentFont Setter
	Position() types.TPosition                                                                      // property Position Getter
	SetPosition(value types.TPosition)                                                              // property Position Setter
	RestoredLeft() int32                                                                            // property RestoredLeft Getter
	RestoredTop() int32                                                                             // property RestoredTop Getter
	RestoredWidth() int32                                                                           // property RestoredWidth Getter
	RestoredHeight() int32                                                                          // property RestoredHeight Getter
	ShowInTaskBar() types.TShowInTaskBar                                                            // property ShowInTaskBar Getter
	SetShowInTaskBar(value types.TShowInTaskBar)                                                    // property ShowInTaskBar Setter
	WindowState() types.TWindowState                                                                // property WindowState Getter
	SetWindowState(value types.TWindowState)                                                        // property WindowState Setter
	SetOnActivate(fn TNotifyEvent)                                                                  // property event
	SetOnClose(fn TCloseEvent)                                                                      // property event
	SetOnCloseQuery(fn TCloseQueryEvent)                                                            // property event
	SetOnCreate(fn TNotifyEvent)                                                                    // property event
	SetOnDeactivate(fn TNotifyEvent)                                                                // property event
	SetOnDestroy(fn TNotifyEvent)                                                                   // property event
	SetOnDropFiles(fn TDropFilesEvent)                                                              // property event
	SetOnHelp(fn THelpEvent)                                                                        // property event
	SetOnHide(fn TNotifyEvent)                                                                      // property event
	SetOnShortcut(fn TShortcutEvent)                                                                // property event
	SetOnShow(fn TNotifyEvent)                                                                      // property event
	SetOnShowModalFinished(fn TModalDialogFinished)                                                 // property event
	SetOnWindowStateChange(fn TNotifyEvent)                                                         // property event
}

type TCustomForm struct {
	TCustomDesignControl
}

func (m *TCustomForm) BigIconHandle() types.HICON {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(2, m.Instance())
	return types.HICON(r)
}

func (m *TCustomForm) CloseQuery() bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(3, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomForm) GetFormImageToBitmap() IBitmap {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(4, m.Instance())
	return AsBitmap(r)
}

func (m *TCustomForm) GetRolesForControl(control IControl) types.TControlRolesForForm {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(5, m.Instance(), base.GetObjectUintptr(control))
	return types.TControlRolesForForm(r)
}

func (m *TCustomForm) GetRealPopupParent() ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(6, m.Instance())
	return AsCustomForm(r)
}

func (m *TCustomForm) IsShortcut(message *types.TLMKey) bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(message)))
	return api.GoBool(r)
}

func (m *TCustomForm) SetFocusedControl(control IWinControl) bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(8, m.Instance(), base.GetObjectUintptr(control))
	return api.GoBool(r)
}

func (m *TCustomForm) ShowModal() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(9, m.Instance())
	return int32(r)
}

func (m *TCustomForm) SmallIconHandle() types.HICON {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(10, m.Instance())
	return types.HICON(r)
}

func (m *TCustomForm) WantChildKey(child IControl, message *types.TLMessage) bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(child), uintptr(base.UnsafePointer(message)))
	return api.GoBool(r)
}

func (m *TCustomForm) ActiveMDIChild() ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(12, m.Instance())
	return AsCustomForm(r)
}

func (m *TCustomForm) GetMDIChildren(index int32) ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(13, m.Instance(), uintptr(index))
	return AsCustomForm(r)
}

func (m *TCustomForm) MDIChildCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(14, m.Instance())
	return int32(r)
}

func (m *TCustomForm) Close() {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(15, m.Instance())
}

func (m *TCustomForm) DefocusControl(control IWinControl, removing bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(control), api.PasBool(removing))
}

func (m *TCustomForm) DestroyWnd() {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(17, m.Instance())
}

func (m *TCustomForm) EnsureVisible(moveToTop bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(18, m.Instance(), api.PasBool(moveToTop))
}

func (m *TCustomForm) FocusControl(winControl IWinControl) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(19, m.Instance(), base.GetObjectUintptr(winControl))
}

func (m *TCustomForm) GetFormImageWithCustomBitmap(bitmap ICustomBitmap) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(20, m.Instance(), base.GetObjectUintptr(bitmap))
}

func (m *TCustomForm) Hide() {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(21, m.Instance())
}

func (m *TCustomForm) IntfHelp(component IComponent) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(22, m.Instance(), base.GetObjectUintptr(component))
}

func (m *TCustomForm) MakeFullyVisible(monitor IMonitor, useWorkarea bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(23, m.Instance(), base.GetObjectUintptr(monitor), api.PasBool(useWorkarea))
}

func (m *TCustomForm) Release() {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(24, m.Instance())
}

func (m *TCustomForm) SetRestoredBounds(left int32, top int32, width int32, height int32, constADefaultPosition bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(25, m.Instance(), uintptr(left), uintptr(top), uintptr(width), uintptr(height), api.PasBool(constADefaultPosition))
}

func (m *TCustomForm) Show() {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(26, m.Instance())
}

func (m *TCustomForm) ShowOnTop() {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(27, m.Instance())
}

func (m *TCustomForm) AutoScale() {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(28, m.Instance())
}

func (m *TCustomForm) Active() bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(29, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomForm) ActiveControl() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(30, 0, m.Instance())
	return AsWinControl(r)
}

func (m *TCustomForm) SetActiveControl(value IWinControl) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(30, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomForm) ActiveDefaultControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(31, 0, m.Instance())
	return AsControl(r)
}

func (m *TCustomForm) SetActiveDefaultControl(value IControl) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(31, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomForm) AllowDropFiles() bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(32, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomForm) SetAllowDropFiles(value bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(32, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomForm) AlphaBlend() bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(33, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomForm) SetAlphaBlend(value bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(33, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomForm) AlphaBlendValue() byte {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(34, 0, m.Instance())
	return byte(r)
}

func (m *TCustomForm) SetAlphaBlendValue(value byte) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(34, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) AutoScroll() bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(35, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomForm) SetAutoScroll(value bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(35, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomForm) BorderIcons() types.TBorderIcons {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(36, 0, m.Instance())
	return types.TBorderIcons(r)
}

func (m *TCustomForm) SetBorderIcons(value types.TBorderIcons) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(36, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) BorderStyleToFormBorderStyle() types.TFormBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(37, 0, m.Instance())
	return types.TFormBorderStyle(r)
}

func (m *TCustomForm) SetBorderStyleToFormBorderStyle(value types.TFormBorderStyle) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(37, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) CancelControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(38, 0, m.Instance())
	return AsControl(r)
}

func (m *TCustomForm) SetCancelControl(value IControl) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(38, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomForm) DefaultControl() IControl {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(39, 0, m.Instance())
	return AsControl(r)
}

func (m *TCustomForm) SetDefaultControl(value IControl) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(39, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomForm) DefaultMonitor() types.TDefaultMonitor {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(40, 0, m.Instance())
	return types.TDefaultMonitor(r)
}

func (m *TCustomForm) SetDefaultMonitor(value types.TDefaultMonitor) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(40, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) Designer() IIDesigner {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(41, 0, m.Instance())
	return AsIDesigner(r)
}

func (m *TCustomForm) SetDesigner(value IIDesigner) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(41, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomForm) EffectiveShowInTaskBar() types.TShowInTaskBar {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(42, m.Instance())
	return types.TShowInTaskBar(r)
}

func (m *TCustomForm) FormState() types.TFormState {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(43, m.Instance())
	return types.TFormState(r)
}

func (m *TCustomForm) FormStyle() types.TFormStyle {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(44, 0, m.Instance())
	return types.TFormStyle(r)
}

func (m *TCustomForm) SetFormStyle(value types.TFormStyle) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(44, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) HelpFile() string {
	if !m.IsValid() {
		return ""
	}
	r := customFormAPI().SysCallN(45, 0, m.Instance())
	return api.GoStr(r)
}

func (m *TCustomForm) SetHelpFile(value string) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(45, 1, m.Instance(), api.PasStr(value))
}

func (m *TCustomForm) Icon() IIcon {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(46, 0, m.Instance())
	return AsIcon(r)
}

func (m *TCustomForm) SetIcon(value IIcon) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(46, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomForm) KeyPreview() bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(47, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomForm) SetKeyPreview(value bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(47, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomForm) MDIChildren(I int32) ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(48, m.Instance(), uintptr(I))
	return AsCustomForm(r)
}

func (m *TCustomForm) Menu() IMainMenu {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(49, 0, m.Instance())
	return AsMainMenu(r)
}

func (m *TCustomForm) SetMenu(value IMainMenu) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(49, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomForm) ModalResult() types.TModalResult {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(50, 0, m.Instance())
	return types.TModalResult(r)
}

func (m *TCustomForm) SetModalResult(value types.TModalResult) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(50, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) Monitor() IMonitor {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(51, m.Instance())
	return AsMonitor(r)
}

func (m *TCustomForm) LastActiveControl() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(52, m.Instance())
	return AsWinControl(r)
}

func (m *TCustomForm) PopupMode() types.TPopupMode {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(53, 0, m.Instance())
	return types.TPopupMode(r)
}

func (m *TCustomForm) SetPopupMode(value types.TPopupMode) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(53, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) PopupParent() ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(54, 0, m.Instance())
	return AsCustomForm(r)
}

func (m *TCustomForm) SetPopupParent(value ICustomForm) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(54, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomForm) SnapOptions() IWindowMagnetOptions {
	if !m.IsValid() {
		return nil
	}
	r := customFormAPI().SysCallN(55, 0, m.Instance())
	return AsWindowMagnetOptions(r)
}

func (m *TCustomForm) SetSnapOptions(value IWindowMagnetOptions) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(55, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomForm) ScreenSnap() bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(56, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomForm) SetScreenSnap(value bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(56, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomForm) SnapBuffer() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(57, 0, m.Instance())
	return int32(r)
}

func (m *TCustomForm) SetSnapBuffer(value int32) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(57, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) ParentFont() bool {
	if !m.IsValid() {
		return false
	}
	r := customFormAPI().SysCallN(58, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomForm) SetParentFont(value bool) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(58, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomForm) Position() types.TPosition {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(59, 0, m.Instance())
	return types.TPosition(r)
}

func (m *TCustomForm) SetPosition(value types.TPosition) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(59, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) RestoredLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(60, m.Instance())
	return int32(r)
}

func (m *TCustomForm) RestoredTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(61, m.Instance())
	return int32(r)
}

func (m *TCustomForm) RestoredWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(62, m.Instance())
	return int32(r)
}

func (m *TCustomForm) RestoredHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(63, m.Instance())
	return int32(r)
}

func (m *TCustomForm) ShowInTaskBar() types.TShowInTaskBar {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(64, 0, m.Instance())
	return types.TShowInTaskBar(r)
}

func (m *TCustomForm) SetShowInTaskBar(value types.TShowInTaskBar) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(64, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) WindowState() types.TWindowState {
	if !m.IsValid() {
		return 0
	}
	r := customFormAPI().SysCallN(65, 0, m.Instance())
	return types.TWindowState(r)
}

func (m *TCustomForm) SetWindowState(value types.TWindowState) {
	if !m.IsValid() {
		return
	}
	customFormAPI().SysCallN(65, 1, m.Instance(), uintptr(value))
}

func (m *TCustomForm) SetOnActivate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 66, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnClose(fn TCloseEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCloseEvent(fn)
	base.SetEvent(m, 67, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnCloseQuery(fn TCloseQueryEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTCloseQueryEvent(fn)
	base.SetEvent(m, 68, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnCreate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 69, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnDeactivate(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 70, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnDestroy(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 71, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnDropFiles(fn TDropFilesEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTDropFilesEvent(fn)
	base.SetEvent(m, 72, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnHelp(fn THelpEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTHelpEvent(fn)
	base.SetEvent(m, 73, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnHide(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 74, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnShortcut(fn TShortcutEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTShortcutEvent(fn)
	base.SetEvent(m, 75, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnShow(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 76, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnShowModalFinished(fn TModalDialogFinished) {
	if !m.IsValid() {
		return
	}
	cb := makeTModalDialogFinished(fn)
	base.SetEvent(m, 77, customFormAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomForm) SetOnWindowStateChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 78, customFormAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomForm class constructor
func NewCustomForm(owner IComponent) ICustomForm {
	r := customFormAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomForm(r)
}

// NewCustomFormNew class constructor
func NewCustomFormNew(owner IComponent, num int32) ICustomForm {
	r := customFormAPI().SysCallN(1, base.GetObjectUintptr(owner), uintptr(num))
	return AsCustomForm(r)
}

func TCustomFormClass() types.TClass {
	r := customFormAPI().SysCallN(79)
	return types.TClass(r)
}

var (
	customFormOnce   base.Once
	customFormImport *imports.Imports = nil
)

func customFormAPI() *imports.Imports {
	customFormOnce.Do(func() {
		customFormImport = api.NewDefaultImports()
		customFormImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomForm_Create", 0), // constructor NewCustomForm
			/* 1 */ imports.NewTable("TCustomForm_CreateNew", 0), // constructor NewCustomFormNew
			/* 2 */ imports.NewTable("TCustomForm_BigIconHandle", 0), // function BigIconHandle
			/* 3 */ imports.NewTable("TCustomForm_CloseQuery", 0), // function CloseQuery
			/* 4 */ imports.NewTable("TCustomForm_GetFormImageToBitmap", 0), // function GetFormImageToBitmap
			/* 5 */ imports.NewTable("TCustomForm_GetRolesForControl", 0), // function GetRolesForControl
			/* 6 */ imports.NewTable("TCustomForm_GetRealPopupParent", 0), // function GetRealPopupParent
			/* 7 */ imports.NewTable("TCustomForm_IsShortcut", 0), // function IsShortcut
			/* 8 */ imports.NewTable("TCustomForm_SetFocusedControl", 0), // function SetFocusedControl
			/* 9 */ imports.NewTable("TCustomForm_ShowModal", 0), // function ShowModal
			/* 10 */ imports.NewTable("TCustomForm_SmallIconHandle", 0), // function SmallIconHandle
			/* 11 */ imports.NewTable("TCustomForm_WantChildKey", 0), // function WantChildKey
			/* 12 */ imports.NewTable("TCustomForm_ActiveMDIChild", 0), // function ActiveMDIChild
			/* 13 */ imports.NewTable("TCustomForm_GetMDIChildren", 0), // function GetMDIChildren
			/* 14 */ imports.NewTable("TCustomForm_MDIChildCount", 0), // function MDIChildCount
			/* 15 */ imports.NewTable("TCustomForm_Close", 0), // procedure Close
			/* 16 */ imports.NewTable("TCustomForm_DefocusControl", 0), // procedure DefocusControl
			/* 17 */ imports.NewTable("TCustomForm_DestroyWnd", 0), // procedure DestroyWnd
			/* 18 */ imports.NewTable("TCustomForm_EnsureVisible", 0), // procedure EnsureVisible
			/* 19 */ imports.NewTable("TCustomForm_FocusControl", 0), // procedure FocusControl
			/* 20 */ imports.NewTable("TCustomForm_GetFormImageWithCustomBitmap", 0), // procedure GetFormImageWithCustomBitmap
			/* 21 */ imports.NewTable("TCustomForm_Hide", 0), // procedure Hide
			/* 22 */ imports.NewTable("TCustomForm_IntfHelp", 0), // procedure IntfHelp
			/* 23 */ imports.NewTable("TCustomForm_MakeFullyVisible", 0), // procedure MakeFullyVisible
			/* 24 */ imports.NewTable("TCustomForm_Release", 0), // procedure Release
			/* 25 */ imports.NewTable("TCustomForm_SetRestoredBounds", 0), // procedure SetRestoredBounds
			/* 26 */ imports.NewTable("TCustomForm_Show", 0), // procedure Show
			/* 27 */ imports.NewTable("TCustomForm_ShowOnTop", 0), // procedure ShowOnTop
			/* 28 */ imports.NewTable("TCustomForm_AutoScale", 0), // procedure AutoScale
			/* 29 */ imports.NewTable("TCustomForm_Active", 0), // property Active
			/* 30 */ imports.NewTable("TCustomForm_ActiveControl", 0), // property ActiveControl
			/* 31 */ imports.NewTable("TCustomForm_ActiveDefaultControl", 0), // property ActiveDefaultControl
			/* 32 */ imports.NewTable("TCustomForm_AllowDropFiles", 0), // property AllowDropFiles
			/* 33 */ imports.NewTable("TCustomForm_AlphaBlend", 0), // property AlphaBlend
			/* 34 */ imports.NewTable("TCustomForm_AlphaBlendValue", 0), // property AlphaBlendValue
			/* 35 */ imports.NewTable("TCustomForm_AutoScroll", 0), // property AutoScroll
			/* 36 */ imports.NewTable("TCustomForm_BorderIcons", 0), // property BorderIcons
			/* 37 */ imports.NewTable("TCustomForm_BorderStyleToFormBorderStyle", 0), // property BorderStyleToFormBorderStyle
			/* 38 */ imports.NewTable("TCustomForm_CancelControl", 0), // property CancelControl
			/* 39 */ imports.NewTable("TCustomForm_DefaultControl", 0), // property DefaultControl
			/* 40 */ imports.NewTable("TCustomForm_DefaultMonitor", 0), // property DefaultMonitor
			/* 41 */ imports.NewTable("TCustomForm_Designer", 0), // property Designer
			/* 42 */ imports.NewTable("TCustomForm_EffectiveShowInTaskBar", 0), // property EffectiveShowInTaskBar
			/* 43 */ imports.NewTable("TCustomForm_FormState", 0), // property FormState
			/* 44 */ imports.NewTable("TCustomForm_FormStyle", 0), // property FormStyle
			/* 45 */ imports.NewTable("TCustomForm_HelpFile", 0), // property HelpFile
			/* 46 */ imports.NewTable("TCustomForm_Icon", 0), // property Icon
			/* 47 */ imports.NewTable("TCustomForm_KeyPreview", 0), // property KeyPreview
			/* 48 */ imports.NewTable("TCustomForm_MDIChildren", 0), // property MDIChildren
			/* 49 */ imports.NewTable("TCustomForm_Menu", 0), // property Menu
			/* 50 */ imports.NewTable("TCustomForm_ModalResult", 0), // property ModalResult
			/* 51 */ imports.NewTable("TCustomForm_Monitor", 0), // property Monitor
			/* 52 */ imports.NewTable("TCustomForm_LastActiveControl", 0), // property LastActiveControl
			/* 53 */ imports.NewTable("TCustomForm_PopupMode", 0), // property PopupMode
			/* 54 */ imports.NewTable("TCustomForm_PopupParent", 0), // property PopupParent
			/* 55 */ imports.NewTable("TCustomForm_SnapOptions", 0), // property SnapOptions
			/* 56 */ imports.NewTable("TCustomForm_ScreenSnap", 0), // property ScreenSnap
			/* 57 */ imports.NewTable("TCustomForm_SnapBuffer", 0), // property SnapBuffer
			/* 58 */ imports.NewTable("TCustomForm_ParentFont", 0), // property ParentFont
			/* 59 */ imports.NewTable("TCustomForm_Position", 0), // property Position
			/* 60 */ imports.NewTable("TCustomForm_RestoredLeft", 0), // property RestoredLeft
			/* 61 */ imports.NewTable("TCustomForm_RestoredTop", 0), // property RestoredTop
			/* 62 */ imports.NewTable("TCustomForm_RestoredWidth", 0), // property RestoredWidth
			/* 63 */ imports.NewTable("TCustomForm_RestoredHeight", 0), // property RestoredHeight
			/* 64 */ imports.NewTable("TCustomForm_ShowInTaskBar", 0), // property ShowInTaskBar
			/* 65 */ imports.NewTable("TCustomForm_WindowState", 0), // property WindowState
			/* 66 */ imports.NewTable("TCustomForm_OnActivate", 0), // event OnActivate
			/* 67 */ imports.NewTable("TCustomForm_OnClose", 0), // event OnClose
			/* 68 */ imports.NewTable("TCustomForm_OnCloseQuery", 0), // event OnCloseQuery
			/* 69 */ imports.NewTable("TCustomForm_OnCreate", 0), // event OnCreate
			/* 70 */ imports.NewTable("TCustomForm_OnDeactivate", 0), // event OnDeactivate
			/* 71 */ imports.NewTable("TCustomForm_OnDestroy", 0), // event OnDestroy
			/* 72 */ imports.NewTable("TCustomForm_OnDropFiles", 0), // event OnDropFiles
			/* 73 */ imports.NewTable("TCustomForm_OnHelp", 0), // event OnHelp
			/* 74 */ imports.NewTable("TCustomForm_OnHide", 0), // event OnHide
			/* 75 */ imports.NewTable("TCustomForm_OnShortcut", 0), // event OnShortcut
			/* 76 */ imports.NewTable("TCustomForm_OnShow", 0), // event OnShow
			/* 77 */ imports.NewTable("TCustomForm_OnShowModalFinished", 0), // event OnShowModalFinished
			/* 78 */ imports.NewTable("TCustomForm_OnWindowStateChange", 0), // event OnWindowStateChange
			/* 79 */ imports.NewTable("TCustomForm_TClass", 0), // function TCustomFormClass
		}
	})
	return customFormImport
}
