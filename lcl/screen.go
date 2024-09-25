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

// IScreen Parent: ILCLComponent
type IScreen interface {
	ILCLComponent
	ActiveControl() IWinControl                                                  // property
	ActiveCustomForm() ICustomForm                                               // property
	ActiveForm() IForm                                                           // property
	Cursor() TCursor                                                             // property
	SetCursor(AValue TCursor)                                                    // property
	RealCursor() TCursor                                                         // property
	Cursors(Index int32) HCURSOR                                                 // property
	SetCursors(Index int32, AValue HCURSOR)                                      // property
	CustomFormCount() int32                                                      // property
	CustomForms(Index int32) ICustomForm                                         // property
	CustomFormZOrderCount() int32                                                // property
	CustomFormsZOrdered(Index int32) ICustomForm                                 // property
	DesktopLeft() int32                                                          // property
	DesktopTop() int32                                                           // property
	DesktopHeight() int32                                                        // property
	DesktopWidth() int32                                                         // property
	DesktopRect() (resultRect TRect)                                             // property
	FocusedForm() ICustomForm                                                    // property
	FormCount() int32                                                            // property
	Forms(Index int32) IForm                                                     // property
	DataModuleCount() int32                                                      // property
	DataModules(Index int32) IDataModule                                         // property
	HintFont() IFont                                                             // property
	SetHintFont(AValue IFont)                                                    // property
	IconFont() IFont                                                             // property
	SetIconFont(AValue IFont)                                                    // property
	MenuFont() IFont                                                             // property
	SetMenuFont(AValue IFont)                                                    // property
	SystemFont() IFont                                                           // property
	SetSystemFont(AValue IFont)                                                  // property
	Fonts() IStrings                                                             // property
	Height() int32                                                               // property
	MonitorCount() int32                                                         // property
	Monitors(Index int32) IMonitor                                               // property
	PixelsPerInch() int32                                                        // property
	PrimaryMonitor() IMonitor                                                    // property
	Width() int32                                                                // property
	WorkAreaRect() (resultRect TRect)                                            // property
	WorkAreaHeight() int32                                                       // property
	WorkAreaLeft() int32                                                         // property
	WorkAreaTop() int32                                                          // property
	WorkAreaWidth() int32                                                        // property
	CustomFormIndex(AForm ICustomForm) int32                                     // function
	FormIndex(AForm IForm) int32                                                 // function
	CustomFormZIndex(AForm ICustomForm) int32                                    // function
	GetCurrentModalForm() ICustomForm                                            // function
	GetCurrentModalFormZIndex() int32                                            // function
	CustomFormBelongsToActiveGroup(AForm ICustomForm) bool                       // function
	FindNonDesignerForm(FormName string) ICustomForm                             // function
	FindForm(FormName string) ICustomForm                                        // function
	FindNonDesignerDataModule(DataModuleName string) IDataModule                 // function
	FindDataModule(DataModuleName string) IDataModule                            // function
	DisableForms(SkipForm ICustomForm, DisabledList IList) IList                 // function
	MonitorFromPoint(Point *TPoint, MonitorDefault TMonitorDefaultTo) IMonitor   // function
	MonitorFromRect(Rect *TRect, MonitorDefault TMonitorDefaultTo) IMonitor      // function
	MonitorFromWindow(Handle THandle, MonitorDefault TMonitorDefaultTo) IMonitor // function
	MoveFormToFocusFront(ACustomForm ICustomForm)                                // procedure
	MoveFormToZFront(ACustomForm ICustomForm)                                    // procedure
	NewFormWasCreated(AForm ICustomForm)                                         // procedure
	UpdateMonitors()                                                             // procedure
	UpdateScreen()                                                               // procedure
	EnableForms(AFormList *IList)                                                // procedure
	BeginTempCursor(aCursor TCursor)                                             // procedure
	EndTempCursor(aCursor TCursor)                                               // procedure
	BeginWaitCursor()                                                            // procedure
	EndWaitCursor()                                                              // procedure
	BeginScreenCursor()                                                          // procedure
	EndScreenCursor()                                                            // procedure
	SetOnActiveControlChange(fn TNotifyEvent)                                    // property event
	SetOnActiveFormChange(fn TNotifyEvent)                                       // property event
}

// TScreen Parent: TLCLComponent
type TScreen struct {
	TLCLComponent
	activeControlChangePtr uintptr
	activeFormChangePtr    uintptr
}

func NewScreen(AOwner IComponent) IScreen {
	r1 := LCL().SysCallN(4899, GetObjectUintptr(AOwner))
	return AsScreen(r1)
}

func (m *TScreen) ActiveControl() IWinControl {
	r1 := LCL().SysCallN(4892, m.Instance())
	return AsWinControl(r1)
}

func (m *TScreen) ActiveCustomForm() ICustomForm {
	r1 := LCL().SysCallN(4893, m.Instance())
	return AsCustomForm(r1)
}

func (m *TScreen) ActiveForm() IForm {
	r1 := LCL().SysCallN(4894, m.Instance())
	return AsForm(r1)
}

func (m *TScreen) Cursor() TCursor {
	r1 := LCL().SysCallN(4900, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TScreen) SetCursor(AValue TCursor) {
	LCL().SysCallN(4900, 1, m.Instance(), uintptr(AValue))
}

func (m *TScreen) RealCursor() TCursor {
	r1 := LCL().SysCallN(4946, m.Instance())
	return TCursor(r1)
}

func (m *TScreen) Cursors(Index int32) HCURSOR {
	r1 := LCL().SysCallN(4901, 0, m.Instance(), uintptr(Index))
	return HCURSOR(r1)
}

func (m *TScreen) SetCursors(Index int32, AValue HCURSOR) {
	LCL().SysCallN(4901, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TScreen) CustomFormCount() int32 {
	r1 := LCL().SysCallN(4903, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomForms(Index int32) ICustomForm {
	r1 := LCL().SysCallN(4907, m.Instance(), uintptr(Index))
	return AsCustomForm(r1)
}

func (m *TScreen) CustomFormZOrderCount() int32 {
	r1 := LCL().SysCallN(4906, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomFormsZOrdered(Index int32) ICustomForm {
	r1 := LCL().SysCallN(4908, m.Instance(), uintptr(Index))
	return AsCustomForm(r1)
}

func (m *TScreen) DesktopLeft() int32 {
	r1 := LCL().SysCallN(4912, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopTop() int32 {
	r1 := LCL().SysCallN(4914, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopHeight() int32 {
	r1 := LCL().SysCallN(4911, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopWidth() int32 {
	r1 := LCL().SysCallN(4915, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopRect() (resultRect TRect) {
	LCL().SysCallN(4913, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TScreen) FocusedForm() ICustomForm {
	r1 := LCL().SysCallN(4925, m.Instance())
	return AsCustomForm(r1)
}

func (m *TScreen) FormCount() int32 {
	r1 := LCL().SysCallN(4927, m.Instance())
	return int32(r1)
}

func (m *TScreen) Forms(Index int32) IForm {
	r1 := LCL().SysCallN(4929, m.Instance(), uintptr(Index))
	return AsForm(r1)
}

func (m *TScreen) DataModuleCount() int32 {
	r1 := LCL().SysCallN(4909, m.Instance())
	return int32(r1)
}

func (m *TScreen) DataModules(Index int32) IDataModule {
	r1 := LCL().SysCallN(4910, m.Instance(), uintptr(Index))
	return AsDataModule(r1)
}

func (m *TScreen) HintFont() IFont {
	r1 := LCL().SysCallN(4933, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetHintFont(AValue IFont) {
	LCL().SysCallN(4933, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) IconFont() IFont {
	r1 := LCL().SysCallN(4934, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetIconFont(AValue IFont) {
	LCL().SysCallN(4934, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) MenuFont() IFont {
	r1 := LCL().SysCallN(4935, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetMenuFont(AValue IFont) {
	LCL().SysCallN(4935, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) SystemFont() IFont {
	r1 := LCL().SysCallN(4949, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetSystemFont(AValue IFont) {
	LCL().SysCallN(4949, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) Fonts() IStrings {
	r1 := LCL().SysCallN(4926, m.Instance())
	return AsStrings(r1)
}

func (m *TScreen) Height() int32 {
	r1 := LCL().SysCallN(4932, m.Instance())
	return int32(r1)
}

func (m *TScreen) MonitorCount() int32 {
	r1 := LCL().SysCallN(4936, m.Instance())
	return int32(r1)
}

func (m *TScreen) Monitors(Index int32) IMonitor {
	r1 := LCL().SysCallN(4940, m.Instance(), uintptr(Index))
	return AsMonitor(r1)
}

func (m *TScreen) PixelsPerInch() int32 {
	r1 := LCL().SysCallN(4944, m.Instance())
	return int32(r1)
}

func (m *TScreen) PrimaryMonitor() IMonitor {
	r1 := LCL().SysCallN(4945, m.Instance())
	return AsMonitor(r1)
}

func (m *TScreen) Width() int32 {
	r1 := LCL().SysCallN(4952, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaRect() (resultRect TRect) {
	LCL().SysCallN(4955, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TScreen) WorkAreaHeight() int32 {
	r1 := LCL().SysCallN(4953, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaLeft() int32 {
	r1 := LCL().SysCallN(4954, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaTop() int32 {
	r1 := LCL().SysCallN(4956, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaWidth() int32 {
	r1 := LCL().SysCallN(4957, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomFormIndex(AForm ICustomForm) int32 {
	r1 := LCL().SysCallN(4904, m.Instance(), GetObjectUintptr(AForm))
	return int32(r1)
}

func (m *TScreen) FormIndex(AForm IForm) int32 {
	r1 := LCL().SysCallN(4928, m.Instance(), GetObjectUintptr(AForm))
	return int32(r1)
}

func (m *TScreen) CustomFormZIndex(AForm ICustomForm) int32 {
	r1 := LCL().SysCallN(4905, m.Instance(), GetObjectUintptr(AForm))
	return int32(r1)
}

func (m *TScreen) GetCurrentModalForm() ICustomForm {
	r1 := LCL().SysCallN(4930, m.Instance())
	return AsCustomForm(r1)
}

func (m *TScreen) GetCurrentModalFormZIndex() int32 {
	r1 := LCL().SysCallN(4931, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomFormBelongsToActiveGroup(AForm ICustomForm) bool {
	r1 := LCL().SysCallN(4902, m.Instance(), GetObjectUintptr(AForm))
	return GoBool(r1)
}

func (m *TScreen) FindNonDesignerForm(FormName string) ICustomForm {
	r1 := LCL().SysCallN(4924, m.Instance(), PascalStr(FormName))
	return AsCustomForm(r1)
}

func (m *TScreen) FindForm(FormName string) ICustomForm {
	r1 := LCL().SysCallN(4922, m.Instance(), PascalStr(FormName))
	return AsCustomForm(r1)
}

func (m *TScreen) FindNonDesignerDataModule(DataModuleName string) IDataModule {
	r1 := LCL().SysCallN(4923, m.Instance(), PascalStr(DataModuleName))
	return AsDataModule(r1)
}

func (m *TScreen) FindDataModule(DataModuleName string) IDataModule {
	r1 := LCL().SysCallN(4921, m.Instance(), PascalStr(DataModuleName))
	return AsDataModule(r1)
}

func (m *TScreen) DisableForms(SkipForm ICustomForm, DisabledList IList) IList {
	r1 := LCL().SysCallN(4916, m.Instance(), GetObjectUintptr(SkipForm), GetObjectUintptr(DisabledList))
	return AsList(r1)
}

func (m *TScreen) MonitorFromPoint(Point *TPoint, MonitorDefault TMonitorDefaultTo) IMonitor {
	r1 := LCL().SysCallN(4937, m.Instance(), uintptr(unsafePointer(Point)), uintptr(MonitorDefault))
	return AsMonitor(r1)
}

func (m *TScreen) MonitorFromRect(Rect *TRect, MonitorDefault TMonitorDefaultTo) IMonitor {
	r1 := LCL().SysCallN(4938, m.Instance(), uintptr(unsafePointer(Rect)), uintptr(MonitorDefault))
	return AsMonitor(r1)
}

func (m *TScreen) MonitorFromWindow(Handle THandle, MonitorDefault TMonitorDefaultTo) IMonitor {
	r1 := LCL().SysCallN(4939, m.Instance(), uintptr(Handle), uintptr(MonitorDefault))
	return AsMonitor(r1)
}

func ScreenClass() TClass {
	ret := LCL().SysCallN(4898)
	return TClass(ret)
}

func (m *TScreen) MoveFormToFocusFront(ACustomForm ICustomForm) {
	LCL().SysCallN(4941, m.Instance(), GetObjectUintptr(ACustomForm))
}

func (m *TScreen) MoveFormToZFront(ACustomForm ICustomForm) {
	LCL().SysCallN(4942, m.Instance(), GetObjectUintptr(ACustomForm))
}

func (m *TScreen) NewFormWasCreated(AForm ICustomForm) {
	LCL().SysCallN(4943, m.Instance(), GetObjectUintptr(AForm))
}

func (m *TScreen) UpdateMonitors() {
	LCL().SysCallN(4950, m.Instance())
}

func (m *TScreen) UpdateScreen() {
	LCL().SysCallN(4951, m.Instance())
}

func (m *TScreen) EnableForms(AFormList *IList) {
	var result0 uintptr
	LCL().SysCallN(4917, m.Instance(), uintptr(unsafePointer(&result0)))
	*AFormList = AsList(result0)
}

func (m *TScreen) BeginTempCursor(aCursor TCursor) {
	LCL().SysCallN(4896, m.Instance(), uintptr(aCursor))
}

func (m *TScreen) EndTempCursor(aCursor TCursor) {
	LCL().SysCallN(4919, m.Instance(), uintptr(aCursor))
}

func (m *TScreen) BeginWaitCursor() {
	LCL().SysCallN(4897, m.Instance())
}

func (m *TScreen) EndWaitCursor() {
	LCL().SysCallN(4920, m.Instance())
}

func (m *TScreen) BeginScreenCursor() {
	LCL().SysCallN(4895, m.Instance())
}

func (m *TScreen) EndScreenCursor() {
	LCL().SysCallN(4918, m.Instance())
}

func (m *TScreen) SetOnActiveControlChange(fn TNotifyEvent) {
	if m.activeControlChangePtr != 0 {
		RemoveEventElement(m.activeControlChangePtr)
	}
	m.activeControlChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4947, m.Instance(), m.activeControlChangePtr)
}

func (m *TScreen) SetOnActiveFormChange(fn TNotifyEvent) {
	if m.activeFormChangePtr != 0 {
		RemoveEventElement(m.activeFormChangePtr)
	}
	m.activeFormChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4948, m.Instance(), m.activeFormChangePtr)
}
