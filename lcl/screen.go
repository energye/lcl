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
	r1 := screenImportAPI().SysCallN(7, GetObjectUintptr(AOwner))
	return AsScreen(r1)
}

func (m *TScreen) ActiveControl() IWinControl {
	r1 := screenImportAPI().SysCallN(0, m.Instance())
	return AsWinControl(r1)
}

func (m *TScreen) ActiveCustomForm() ICustomForm {
	r1 := screenImportAPI().SysCallN(1, m.Instance())
	return AsCustomForm(r1)
}

func (m *TScreen) ActiveForm() IForm {
	r1 := screenImportAPI().SysCallN(2, m.Instance())
	return AsForm(r1)
}

func (m *TScreen) Cursor() TCursor {
	r1 := screenImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TScreen) SetCursor(AValue TCursor) {
	screenImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TScreen) RealCursor() TCursor {
	r1 := screenImportAPI().SysCallN(54, m.Instance())
	return TCursor(r1)
}

func (m *TScreen) Cursors(Index int32) HCURSOR {
	r1 := screenImportAPI().SysCallN(9, 0, m.Instance(), uintptr(Index))
	return HCURSOR(r1)
}

func (m *TScreen) SetCursors(Index int32, AValue HCURSOR) {
	screenImportAPI().SysCallN(9, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TScreen) CustomFormCount() int32 {
	r1 := screenImportAPI().SysCallN(11, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomForms(Index int32) ICustomForm {
	r1 := screenImportAPI().SysCallN(15, m.Instance(), uintptr(Index))
	return AsCustomForm(r1)
}

func (m *TScreen) CustomFormZOrderCount() int32 {
	r1 := screenImportAPI().SysCallN(14, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomFormsZOrdered(Index int32) ICustomForm {
	r1 := screenImportAPI().SysCallN(16, m.Instance(), uintptr(Index))
	return AsCustomForm(r1)
}

func (m *TScreen) DesktopLeft() int32 {
	r1 := screenImportAPI().SysCallN(20, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopTop() int32 {
	r1 := screenImportAPI().SysCallN(22, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopHeight() int32 {
	r1 := screenImportAPI().SysCallN(19, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopWidth() int32 {
	r1 := screenImportAPI().SysCallN(23, m.Instance())
	return int32(r1)
}

func (m *TScreen) DesktopRect() (resultRect TRect) {
	screenImportAPI().SysCallN(21, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TScreen) FocusedForm() ICustomForm {
	r1 := screenImportAPI().SysCallN(33, m.Instance())
	return AsCustomForm(r1)
}

func (m *TScreen) FormCount() int32 {
	r1 := screenImportAPI().SysCallN(35, m.Instance())
	return int32(r1)
}

func (m *TScreen) Forms(Index int32) IForm {
	r1 := screenImportAPI().SysCallN(37, m.Instance(), uintptr(Index))
	return AsForm(r1)
}

func (m *TScreen) DataModuleCount() int32 {
	r1 := screenImportAPI().SysCallN(17, m.Instance())
	return int32(r1)
}

func (m *TScreen) DataModules(Index int32) IDataModule {
	r1 := screenImportAPI().SysCallN(18, m.Instance(), uintptr(Index))
	return AsDataModule(r1)
}

func (m *TScreen) HintFont() IFont {
	r1 := screenImportAPI().SysCallN(41, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetHintFont(AValue IFont) {
	screenImportAPI().SysCallN(41, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) IconFont() IFont {
	r1 := screenImportAPI().SysCallN(42, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetIconFont(AValue IFont) {
	screenImportAPI().SysCallN(42, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) MenuFont() IFont {
	r1 := screenImportAPI().SysCallN(43, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetMenuFont(AValue IFont) {
	screenImportAPI().SysCallN(43, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) SystemFont() IFont {
	r1 := screenImportAPI().SysCallN(57, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TScreen) SetSystemFont(AValue IFont) {
	screenImportAPI().SysCallN(57, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TScreen) Fonts() IStrings {
	r1 := screenImportAPI().SysCallN(34, m.Instance())
	return AsStrings(r1)
}

func (m *TScreen) Height() int32 {
	r1 := screenImportAPI().SysCallN(40, m.Instance())
	return int32(r1)
}

func (m *TScreen) MonitorCount() int32 {
	r1 := screenImportAPI().SysCallN(44, m.Instance())
	return int32(r1)
}

func (m *TScreen) Monitors(Index int32) IMonitor {
	r1 := screenImportAPI().SysCallN(48, m.Instance(), uintptr(Index))
	return AsMonitor(r1)
}

func (m *TScreen) PixelsPerInch() int32 {
	r1 := screenImportAPI().SysCallN(52, m.Instance())
	return int32(r1)
}

func (m *TScreen) PrimaryMonitor() IMonitor {
	r1 := screenImportAPI().SysCallN(53, m.Instance())
	return AsMonitor(r1)
}

func (m *TScreen) Width() int32 {
	r1 := screenImportAPI().SysCallN(60, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaRect() (resultRect TRect) {
	screenImportAPI().SysCallN(63, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TScreen) WorkAreaHeight() int32 {
	r1 := screenImportAPI().SysCallN(61, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaLeft() int32 {
	r1 := screenImportAPI().SysCallN(62, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaTop() int32 {
	r1 := screenImportAPI().SysCallN(64, m.Instance())
	return int32(r1)
}

func (m *TScreen) WorkAreaWidth() int32 {
	r1 := screenImportAPI().SysCallN(65, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomFormIndex(AForm ICustomForm) int32 {
	r1 := screenImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(AForm))
	return int32(r1)
}

func (m *TScreen) FormIndex(AForm IForm) int32 {
	r1 := screenImportAPI().SysCallN(36, m.Instance(), GetObjectUintptr(AForm))
	return int32(r1)
}

func (m *TScreen) CustomFormZIndex(AForm ICustomForm) int32 {
	r1 := screenImportAPI().SysCallN(13, m.Instance(), GetObjectUintptr(AForm))
	return int32(r1)
}

func (m *TScreen) GetCurrentModalForm() ICustomForm {
	r1 := screenImportAPI().SysCallN(38, m.Instance())
	return AsCustomForm(r1)
}

func (m *TScreen) GetCurrentModalFormZIndex() int32 {
	r1 := screenImportAPI().SysCallN(39, m.Instance())
	return int32(r1)
}

func (m *TScreen) CustomFormBelongsToActiveGroup(AForm ICustomForm) bool {
	r1 := screenImportAPI().SysCallN(10, m.Instance(), GetObjectUintptr(AForm))
	return GoBool(r1)
}

func (m *TScreen) FindNonDesignerForm(FormName string) ICustomForm {
	r1 := screenImportAPI().SysCallN(32, m.Instance(), PascalStr(FormName))
	return AsCustomForm(r1)
}

func (m *TScreen) FindForm(FormName string) ICustomForm {
	r1 := screenImportAPI().SysCallN(30, m.Instance(), PascalStr(FormName))
	return AsCustomForm(r1)
}

func (m *TScreen) FindNonDesignerDataModule(DataModuleName string) IDataModule {
	r1 := screenImportAPI().SysCallN(31, m.Instance(), PascalStr(DataModuleName))
	return AsDataModule(r1)
}

func (m *TScreen) FindDataModule(DataModuleName string) IDataModule {
	r1 := screenImportAPI().SysCallN(29, m.Instance(), PascalStr(DataModuleName))
	return AsDataModule(r1)
}

func (m *TScreen) DisableForms(SkipForm ICustomForm, DisabledList IList) IList {
	r1 := screenImportAPI().SysCallN(24, m.Instance(), GetObjectUintptr(SkipForm), GetObjectUintptr(DisabledList))
	return AsList(r1)
}

func (m *TScreen) MonitorFromPoint(Point *TPoint, MonitorDefault TMonitorDefaultTo) IMonitor {
	r1 := screenImportAPI().SysCallN(45, m.Instance(), uintptr(unsafePointer(Point)), uintptr(MonitorDefault))
	return AsMonitor(r1)
}

func (m *TScreen) MonitorFromRect(Rect *TRect, MonitorDefault TMonitorDefaultTo) IMonitor {
	r1 := screenImportAPI().SysCallN(46, m.Instance(), uintptr(unsafePointer(Rect)), uintptr(MonitorDefault))
	return AsMonitor(r1)
}

func (m *TScreen) MonitorFromWindow(Handle THandle, MonitorDefault TMonitorDefaultTo) IMonitor {
	r1 := screenImportAPI().SysCallN(47, m.Instance(), uintptr(Handle), uintptr(MonitorDefault))
	return AsMonitor(r1)
}

func ScreenClass() TClass {
	ret := screenImportAPI().SysCallN(6)
	return TClass(ret)
}

func (m *TScreen) MoveFormToFocusFront(ACustomForm ICustomForm) {
	screenImportAPI().SysCallN(49, m.Instance(), GetObjectUintptr(ACustomForm))
}

func (m *TScreen) MoveFormToZFront(ACustomForm ICustomForm) {
	screenImportAPI().SysCallN(50, m.Instance(), GetObjectUintptr(ACustomForm))
}

func (m *TScreen) NewFormWasCreated(AForm ICustomForm) {
	screenImportAPI().SysCallN(51, m.Instance(), GetObjectUintptr(AForm))
}

func (m *TScreen) UpdateMonitors() {
	screenImportAPI().SysCallN(58, m.Instance())
}

func (m *TScreen) UpdateScreen() {
	screenImportAPI().SysCallN(59, m.Instance())
}

func (m *TScreen) EnableForms(AFormList *IList) {
	var result0 uintptr
	screenImportAPI().SysCallN(25, m.Instance(), uintptr(unsafePointer(&result0)))
	*AFormList = AsList(result0)
}

func (m *TScreen) BeginTempCursor(aCursor TCursor) {
	screenImportAPI().SysCallN(4, m.Instance(), uintptr(aCursor))
}

func (m *TScreen) EndTempCursor(aCursor TCursor) {
	screenImportAPI().SysCallN(27, m.Instance(), uintptr(aCursor))
}

func (m *TScreen) BeginWaitCursor() {
	screenImportAPI().SysCallN(5, m.Instance())
}

func (m *TScreen) EndWaitCursor() {
	screenImportAPI().SysCallN(28, m.Instance())
}

func (m *TScreen) BeginScreenCursor() {
	screenImportAPI().SysCallN(3, m.Instance())
}

func (m *TScreen) EndScreenCursor() {
	screenImportAPI().SysCallN(26, m.Instance())
}

func (m *TScreen) SetOnActiveControlChange(fn TNotifyEvent) {
	if m.activeControlChangePtr != 0 {
		RemoveEventElement(m.activeControlChangePtr)
	}
	m.activeControlChangePtr = MakeEventDataPtr(fn)
	screenImportAPI().SysCallN(55, m.Instance(), m.activeControlChangePtr)
}

func (m *TScreen) SetOnActiveFormChange(fn TNotifyEvent) {
	if m.activeFormChangePtr != 0 {
		RemoveEventElement(m.activeFormChangePtr)
	}
	m.activeFormChangePtr = MakeEventDataPtr(fn)
	screenImportAPI().SysCallN(56, m.Instance(), m.activeFormChangePtr)
}

var (
	screenImport       *imports.Imports = nil
	screenImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Screen_ActiveControl", 0),
		/*1*/ imports.NewTable("Screen_ActiveCustomForm", 0),
		/*2*/ imports.NewTable("Screen_ActiveForm", 0),
		/*3*/ imports.NewTable("Screen_BeginScreenCursor", 0),
		/*4*/ imports.NewTable("Screen_BeginTempCursor", 0),
		/*5*/ imports.NewTable("Screen_BeginWaitCursor", 0),
		/*6*/ imports.NewTable("Screen_Class", 0),
		/*7*/ imports.NewTable("Screen_Create", 0),
		/*8*/ imports.NewTable("Screen_Cursor", 0),
		/*9*/ imports.NewTable("Screen_Cursors", 0),
		/*10*/ imports.NewTable("Screen_CustomFormBelongsToActiveGroup", 0),
		/*11*/ imports.NewTable("Screen_CustomFormCount", 0),
		/*12*/ imports.NewTable("Screen_CustomFormIndex", 0),
		/*13*/ imports.NewTable("Screen_CustomFormZIndex", 0),
		/*14*/ imports.NewTable("Screen_CustomFormZOrderCount", 0),
		/*15*/ imports.NewTable("Screen_CustomForms", 0),
		/*16*/ imports.NewTable("Screen_CustomFormsZOrdered", 0),
		/*17*/ imports.NewTable("Screen_DataModuleCount", 0),
		/*18*/ imports.NewTable("Screen_DataModules", 0),
		/*19*/ imports.NewTable("Screen_DesktopHeight", 0),
		/*20*/ imports.NewTable("Screen_DesktopLeft", 0),
		/*21*/ imports.NewTable("Screen_DesktopRect", 0),
		/*22*/ imports.NewTable("Screen_DesktopTop", 0),
		/*23*/ imports.NewTable("Screen_DesktopWidth", 0),
		/*24*/ imports.NewTable("Screen_DisableForms", 0),
		/*25*/ imports.NewTable("Screen_EnableForms", 0),
		/*26*/ imports.NewTable("Screen_EndScreenCursor", 0),
		/*27*/ imports.NewTable("Screen_EndTempCursor", 0),
		/*28*/ imports.NewTable("Screen_EndWaitCursor", 0),
		/*29*/ imports.NewTable("Screen_FindDataModule", 0),
		/*30*/ imports.NewTable("Screen_FindForm", 0),
		/*31*/ imports.NewTable("Screen_FindNonDesignerDataModule", 0),
		/*32*/ imports.NewTable("Screen_FindNonDesignerForm", 0),
		/*33*/ imports.NewTable("Screen_FocusedForm", 0),
		/*34*/ imports.NewTable("Screen_Fonts", 0),
		/*35*/ imports.NewTable("Screen_FormCount", 0),
		/*36*/ imports.NewTable("Screen_FormIndex", 0),
		/*37*/ imports.NewTable("Screen_Forms", 0),
		/*38*/ imports.NewTable("Screen_GetCurrentModalForm", 0),
		/*39*/ imports.NewTable("Screen_GetCurrentModalFormZIndex", 0),
		/*40*/ imports.NewTable("Screen_Height", 0),
		/*41*/ imports.NewTable("Screen_HintFont", 0),
		/*42*/ imports.NewTable("Screen_IconFont", 0),
		/*43*/ imports.NewTable("Screen_MenuFont", 0),
		/*44*/ imports.NewTable("Screen_MonitorCount", 0),
		/*45*/ imports.NewTable("Screen_MonitorFromPoint", 0),
		/*46*/ imports.NewTable("Screen_MonitorFromRect", 0),
		/*47*/ imports.NewTable("Screen_MonitorFromWindow", 0),
		/*48*/ imports.NewTable("Screen_Monitors", 0),
		/*49*/ imports.NewTable("Screen_MoveFormToFocusFront", 0),
		/*50*/ imports.NewTable("Screen_MoveFormToZFront", 0),
		/*51*/ imports.NewTable("Screen_NewFormWasCreated", 0),
		/*52*/ imports.NewTable("Screen_PixelsPerInch", 0),
		/*53*/ imports.NewTable("Screen_PrimaryMonitor", 0),
		/*54*/ imports.NewTable("Screen_RealCursor", 0),
		/*55*/ imports.NewTable("Screen_SetOnActiveControlChange", 0),
		/*56*/ imports.NewTable("Screen_SetOnActiveFormChange", 0),
		/*57*/ imports.NewTable("Screen_SystemFont", 0),
		/*58*/ imports.NewTable("Screen_UpdateMonitors", 0),
		/*59*/ imports.NewTable("Screen_UpdateScreen", 0),
		/*60*/ imports.NewTable("Screen_Width", 0),
		/*61*/ imports.NewTable("Screen_WorkAreaHeight", 0),
		/*62*/ imports.NewTable("Screen_WorkAreaLeft", 0),
		/*63*/ imports.NewTable("Screen_WorkAreaRect", 0),
		/*64*/ imports.NewTable("Screen_WorkAreaTop", 0),
		/*65*/ imports.NewTable("Screen_WorkAreaWidth", 0),
	}
)

func screenImportAPI() *imports.Imports {
	if screenImport == nil {
		screenImport = NewDefaultImports()
		screenImport.SetImportTable(screenImportTables)
		screenImportTables = nil
	}
	return screenImport
}
