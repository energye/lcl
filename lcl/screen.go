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

// IScreen Parent: ILCLComponent
type IScreen interface {
	ILCLComponent
	CustomFormIndex(form ICustomForm) int32                                                     // function
	FormIndex(form IForm) int32                                                                 // function
	CustomFormZIndex(form ICustomForm) int32                                                    // function
	GetCurrentModalForm() ICustomForm                                                           // function
	GetCurrentModalFormZIndex() int32                                                           // function
	CustomFormBelongsToActiveGroup(form ICustomForm) bool                                       // function
	FindNonDesignerForm(formName string) ICustomForm                                            // function
	FindForm(formName string) ICustomForm                                                       // function
	FindNonDesignerDataModule(dataModuleName string) IDataModule                                // function
	FindDataModule(dataModuleName string) IDataModule                                           // function
	DisableForms(skipForm ICustomForm, disabledList IList) IList                                // function
	MonitorFromPoint(point types.TPoint, monitorDefault types.TMonitorDefaultTo) IMonitor       // function
	MonitorFromRect(rect types.TRect, monitorDefault types.TMonitorDefaultTo) IMonitor          // function
	MonitorFromWindow(handle types.TLCLHandle, monitorDefault types.TMonitorDefaultTo) IMonitor // function
	MoveFormToFocusFront(customForm ICustomForm)                                                // procedure
	MoveFormToZFront(customForm ICustomForm)                                                    // procedure
	NewFormWasCreated(form ICustomForm)                                                         // procedure
	UpdateMonitors()                                                                            // procedure
	UpdateScreen()                                                                              // procedure
	EnableForms(formList *IList)                                                                // procedure
	BeginTempCursor(cursor types.TCursor)                                                       // procedure
	EndTempCursor(cursor types.TCursor)                                                         // procedure
	BeginWaitCursor()                                                                           // procedure
	EndWaitCursor()                                                                             // procedure
	BeginScreenCursor()                                                                         // procedure
	EndScreenCursor()                                                                           // procedure
	ActiveControl() IWinControl                                                                 // property ActiveControl Getter
	ActiveCustomForm() ICustomForm                                                              // property ActiveCustomForm Getter
	ActiveForm() IForm                                                                          // property ActiveForm Getter
	Cursor() types.TCursor                                                                      // property Cursor Getter
	SetCursor(value types.TCursor)                                                              // property Cursor Setter
	RealCursor() types.TCursor                                                                  // property RealCursor Getter
	Cursors(index int32) types.HCURSOR                                                          // property Cursors Getter
	SetCursors(index int32, value types.HCURSOR)                                                // property Cursors Setter
	CustomFormCount() int32                                                                     // property CustomFormCount Getter
	CustomForms(index int32) ICustomForm                                                        // property CustomForms Getter
	CustomFormZOrderCount() int32                                                               // property CustomFormZOrderCount Getter
	CustomFormsZOrdered(index int32) ICustomForm                                                // property CustomFormsZOrdered Getter
	DesktopLeft() int32                                                                         // property DesktopLeft Getter
	DesktopTop() int32                                                                          // property DesktopTop Getter
	DesktopHeight() int32                                                                       // property DesktopHeight Getter
	DesktopWidth() int32                                                                        // property DesktopWidth Getter
	DesktopRect() types.TRect                                                                   // property DesktopRect Getter
	FocusedForm() ICustomForm                                                                   // property FocusedForm Getter
	FormCount() int32                                                                           // property FormCount Getter
	Forms(index int32) IForm                                                                    // property Forms Getter
	DataModuleCount() int32                                                                     // property DataModuleCount Getter
	DataModules(index int32) IDataModule                                                        // property DataModules Getter
	HintFont() IFont                                                                            // property HintFont Getter
	SetHintFont(value IFont)                                                                    // property HintFont Setter
	IconFont() IFont                                                                            // property IconFont Getter
	SetIconFont(value IFont)                                                                    // property IconFont Setter
	MenuFont() IFont                                                                            // property MenuFont Getter
	SetMenuFont(value IFont)                                                                    // property MenuFont Setter
	SystemFont() IFont                                                                          // property SystemFont Getter
	SetSystemFont(value IFont)                                                                  // property SystemFont Setter
	Fonts() IStrings                                                                            // property Fonts Getter
	Height() int32                                                                              // property Height Getter
	MonitorCount() int32                                                                        // property MonitorCount Getter
	Monitors(index int32) IMonitor                                                              // property Monitors Getter
	PixelsPerInch() int32                                                                       // property PixelsPerInch Getter
	PrimaryMonitor() IMonitor                                                                   // property PrimaryMonitor Getter
	Width() int32                                                                               // property Width Getter
	WorkAreaRect() types.TRect                                                                  // property WorkAreaRect Getter
	WorkAreaHeight() int32                                                                      // property WorkAreaHeight Getter
	WorkAreaLeft() int32                                                                        // property WorkAreaLeft Getter
	WorkAreaTop() int32                                                                         // property WorkAreaTop Getter
	WorkAreaWidth() int32                                                                       // property WorkAreaWidth Getter
	SetOnActiveControlChange(fn TNotifyEvent)                                                   // property event
	SetOnActiveFormChange(fn TNotifyEvent)                                                      // property event
}

type TScreen struct {
	TLCLComponent
}

func (m *TScreen) CustomFormIndex(form ICustomForm) int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(1, m.Instance(), base.GetObjectUintptr(form))
	return int32(r)
}

func (m *TScreen) FormIndex(form IForm) int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(2, m.Instance(), base.GetObjectUintptr(form))
	return int32(r)
}

func (m *TScreen) CustomFormZIndex(form ICustomForm) int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(form))
	return int32(r)
}

func (m *TScreen) GetCurrentModalForm() ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(4, m.Instance())
	return AsCustomForm(r)
}

func (m *TScreen) GetCurrentModalFormZIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(5, m.Instance())
	return int32(r)
}

func (m *TScreen) CustomFormBelongsToActiveGroup(form ICustomForm) bool {
	if !m.IsValid() {
		return false
	}
	r := screenAPI().SysCallN(6, m.Instance(), base.GetObjectUintptr(form))
	return api.GoBool(r)
}

func (m *TScreen) FindNonDesignerForm(formName string) ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(7, m.Instance(), api.PasStr(formName))
	return AsCustomForm(r)
}

func (m *TScreen) FindForm(formName string) ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(8, m.Instance(), api.PasStr(formName))
	return AsCustomForm(r)
}

func (m *TScreen) FindNonDesignerDataModule(dataModuleName string) IDataModule {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(9, m.Instance(), api.PasStr(dataModuleName))
	return AsDataModule(r)
}

func (m *TScreen) FindDataModule(dataModuleName string) IDataModule {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(10, m.Instance(), api.PasStr(dataModuleName))
	return AsDataModule(r)
}

func (m *TScreen) DisableForms(skipForm ICustomForm, disabledList IList) IList {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(skipForm), base.GetObjectUintptr(disabledList))
	return AsList(r)
}

func (m *TScreen) MonitorFromPoint(point types.TPoint, monitorDefault types.TMonitorDefaultTo) IMonitor {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(12, m.Instance(), uintptr(base.UnsafePointer(&point)), uintptr(monitorDefault))
	return AsMonitor(r)
}

func (m *TScreen) MonitorFromRect(rect types.TRect, monitorDefault types.TMonitorDefaultTo) IMonitor {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(13, m.Instance(), uintptr(base.UnsafePointer(&rect)), uintptr(monitorDefault))
	return AsMonitor(r)
}

func (m *TScreen) MonitorFromWindow(handle types.TLCLHandle, monitorDefault types.TMonitorDefaultTo) IMonitor {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(14, m.Instance(), uintptr(handle), uintptr(monitorDefault))
	return AsMonitor(r)
}

func (m *TScreen) MoveFormToFocusFront(customForm ICustomForm) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(15, m.Instance(), base.GetObjectUintptr(customForm))
}

func (m *TScreen) MoveFormToZFront(customForm ICustomForm) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(16, m.Instance(), base.GetObjectUintptr(customForm))
}

func (m *TScreen) NewFormWasCreated(form ICustomForm) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(17, m.Instance(), base.GetObjectUintptr(form))
}

func (m *TScreen) UpdateMonitors() {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(18, m.Instance())
}

func (m *TScreen) UpdateScreen() {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(19, m.Instance())
}

func (m *TScreen) EnableForms(formList *IList) {
	if !m.IsValid() {
		return
	}
	formListPtr := base.GetObjectUintptr(*formList)
	screenAPI().SysCallN(20, m.Instance(), uintptr(base.UnsafePointer(&formListPtr)))
	*formList = AsList(formListPtr)
}

func (m *TScreen) BeginTempCursor(cursor types.TCursor) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(21, m.Instance(), uintptr(cursor))
}

func (m *TScreen) EndTempCursor(cursor types.TCursor) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(22, m.Instance(), uintptr(cursor))
}

func (m *TScreen) BeginWaitCursor() {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(23, m.Instance())
}

func (m *TScreen) EndWaitCursor() {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(24, m.Instance())
}

func (m *TScreen) BeginScreenCursor() {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(25, m.Instance())
}

func (m *TScreen) EndScreenCursor() {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(26, m.Instance())
}

func (m *TScreen) ActiveControl() IWinControl {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(27, m.Instance())
	return AsWinControl(r)
}

func (m *TScreen) ActiveCustomForm() ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(28, m.Instance())
	return AsCustomForm(r)
}

func (m *TScreen) ActiveForm() IForm {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(29, m.Instance())
	return AsForm(r)
}

func (m *TScreen) Cursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(30, 0, m.Instance())
	return types.TCursor(r)
}

func (m *TScreen) SetCursor(value types.TCursor) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(30, 1, m.Instance(), uintptr(value))
}

func (m *TScreen) RealCursor() types.TCursor {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(31, m.Instance())
	return types.TCursor(r)
}

func (m *TScreen) Cursors(index int32) types.HCURSOR {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(32, 0, m.Instance(), uintptr(index))
	return types.HCURSOR(r)
}

func (m *TScreen) SetCursors(index int32, value types.HCURSOR) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(32, 1, m.Instance(), uintptr(index), uintptr(value))
}

func (m *TScreen) CustomFormCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(33, m.Instance())
	return int32(r)
}

func (m *TScreen) CustomForms(index int32) ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(34, m.Instance(), uintptr(index))
	return AsCustomForm(r)
}

func (m *TScreen) CustomFormZOrderCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(35, m.Instance())
	return int32(r)
}

func (m *TScreen) CustomFormsZOrdered(index int32) ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(36, m.Instance(), uintptr(index))
	return AsCustomForm(r)
}

func (m *TScreen) DesktopLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(37, m.Instance())
	return int32(r)
}

func (m *TScreen) DesktopTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(38, m.Instance())
	return int32(r)
}

func (m *TScreen) DesktopHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(39, m.Instance())
	return int32(r)
}

func (m *TScreen) DesktopWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(40, m.Instance())
	return int32(r)
}

func (m *TScreen) DesktopRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(41, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TScreen) FocusedForm() ICustomForm {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(42, m.Instance())
	return AsCustomForm(r)
}

func (m *TScreen) FormCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(43, m.Instance())
	return int32(r)
}

func (m *TScreen) Forms(index int32) IForm {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(44, m.Instance(), uintptr(index))
	return AsForm(r)
}

func (m *TScreen) DataModuleCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(45, m.Instance())
	return int32(r)
}

func (m *TScreen) DataModules(index int32) IDataModule {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(46, m.Instance(), uintptr(index))
	return AsDataModule(r)
}

func (m *TScreen) HintFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(47, 0, m.Instance())
	return AsFont(r)
}

func (m *TScreen) SetHintFont(value IFont) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(47, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TScreen) IconFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(48, 0, m.Instance())
	return AsFont(r)
}

func (m *TScreen) SetIconFont(value IFont) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(48, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TScreen) MenuFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(49, 0, m.Instance())
	return AsFont(r)
}

func (m *TScreen) SetMenuFont(value IFont) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(49, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TScreen) SystemFont() IFont {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(50, 0, m.Instance())
	return AsFont(r)
}

func (m *TScreen) SetSystemFont(value IFont) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(50, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TScreen) Fonts() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(51, m.Instance())
	return AsStrings(r)
}

func (m *TScreen) Height() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(52, m.Instance())
	return int32(r)
}

func (m *TScreen) MonitorCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(53, m.Instance())
	return int32(r)
}

func (m *TScreen) Monitors(index int32) IMonitor {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(54, m.Instance(), uintptr(index))
	return AsMonitor(r)
}

func (m *TScreen) PixelsPerInch() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(55, m.Instance())
	return int32(r)
}

func (m *TScreen) PrimaryMonitor() IMonitor {
	if !m.IsValid() {
		return nil
	}
	r := screenAPI().SysCallN(56, m.Instance())
	return AsMonitor(r)
}

func (m *TScreen) Width() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(57, m.Instance())
	return int32(r)
}

func (m *TScreen) WorkAreaRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	screenAPI().SysCallN(58, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TScreen) WorkAreaHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(59, m.Instance())
	return int32(r)
}

func (m *TScreen) WorkAreaLeft() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(60, m.Instance())
	return int32(r)
}

func (m *TScreen) WorkAreaTop() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(61, m.Instance())
	return int32(r)
}

func (m *TScreen) WorkAreaWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := screenAPI().SysCallN(62, m.Instance())
	return int32(r)
}

func (m *TScreen) SetOnActiveControlChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 63, screenAPI(), api.MakeEventDataPtr(cb))
}

func (m *TScreen) SetOnActiveFormChange(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 64, screenAPI(), api.MakeEventDataPtr(cb))
}

// NewScreen class constructor
func NewScreen(owner IComponent) IScreen {
	r := screenAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsScreen(r)
}

func TScreenClass() types.TClass {
	r := screenAPI().SysCallN(65)
	return types.TClass(r)
}

var (
	screenOnce   base.Once
	screenImport *imports.Imports = nil
)

func screenAPI() *imports.Imports {
	screenOnce.Do(func() {
		screenImport = api.NewDefaultImports()
		screenImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TScreen_Create", 0), // constructor NewScreen
			/* 1 */ imports.NewTable("TScreen_CustomFormIndex", 0), // function CustomFormIndex
			/* 2 */ imports.NewTable("TScreen_FormIndex", 0), // function FormIndex
			/* 3 */ imports.NewTable("TScreen_CustomFormZIndex", 0), // function CustomFormZIndex
			/* 4 */ imports.NewTable("TScreen_GetCurrentModalForm", 0), // function GetCurrentModalForm
			/* 5 */ imports.NewTable("TScreen_GetCurrentModalFormZIndex", 0), // function GetCurrentModalFormZIndex
			/* 6 */ imports.NewTable("TScreen_CustomFormBelongsToActiveGroup", 0), // function CustomFormBelongsToActiveGroup
			/* 7 */ imports.NewTable("TScreen_FindNonDesignerForm", 0), // function FindNonDesignerForm
			/* 8 */ imports.NewTable("TScreen_FindForm", 0), // function FindForm
			/* 9 */ imports.NewTable("TScreen_FindNonDesignerDataModule", 0), // function FindNonDesignerDataModule
			/* 10 */ imports.NewTable("TScreen_FindDataModule", 0), // function FindDataModule
			/* 11 */ imports.NewTable("TScreen_DisableForms", 0), // function DisableForms
			/* 12 */ imports.NewTable("TScreen_MonitorFromPoint", 0), // function MonitorFromPoint
			/* 13 */ imports.NewTable("TScreen_MonitorFromRect", 0), // function MonitorFromRect
			/* 14 */ imports.NewTable("TScreen_MonitorFromWindow", 0), // function MonitorFromWindow
			/* 15 */ imports.NewTable("TScreen_MoveFormToFocusFront", 0), // procedure MoveFormToFocusFront
			/* 16 */ imports.NewTable("TScreen_MoveFormToZFront", 0), // procedure MoveFormToZFront
			/* 17 */ imports.NewTable("TScreen_NewFormWasCreated", 0), // procedure NewFormWasCreated
			/* 18 */ imports.NewTable("TScreen_UpdateMonitors", 0), // procedure UpdateMonitors
			/* 19 */ imports.NewTable("TScreen_UpdateScreen", 0), // procedure UpdateScreen
			/* 20 */ imports.NewTable("TScreen_EnableForms", 0), // procedure EnableForms
			/* 21 */ imports.NewTable("TScreen_BeginTempCursor", 0), // procedure BeginTempCursor
			/* 22 */ imports.NewTable("TScreen_EndTempCursor", 0), // procedure EndTempCursor
			/* 23 */ imports.NewTable("TScreen_BeginWaitCursor", 0), // procedure BeginWaitCursor
			/* 24 */ imports.NewTable("TScreen_EndWaitCursor", 0), // procedure EndWaitCursor
			/* 25 */ imports.NewTable("TScreen_BeginScreenCursor", 0), // procedure BeginScreenCursor
			/* 26 */ imports.NewTable("TScreen_EndScreenCursor", 0), // procedure EndScreenCursor
			/* 27 */ imports.NewTable("TScreen_ActiveControl", 0), // property ActiveControl
			/* 28 */ imports.NewTable("TScreen_ActiveCustomForm", 0), // property ActiveCustomForm
			/* 29 */ imports.NewTable("TScreen_ActiveForm", 0), // property ActiveForm
			/* 30 */ imports.NewTable("TScreen_Cursor", 0), // property Cursor
			/* 31 */ imports.NewTable("TScreen_RealCursor", 0), // property RealCursor
			/* 32 */ imports.NewTable("TScreen_Cursors", 0), // property Cursors
			/* 33 */ imports.NewTable("TScreen_CustomFormCount", 0), // property CustomFormCount
			/* 34 */ imports.NewTable("TScreen_CustomForms", 0), // property CustomForms
			/* 35 */ imports.NewTable("TScreen_CustomFormZOrderCount", 0), // property CustomFormZOrderCount
			/* 36 */ imports.NewTable("TScreen_CustomFormsZOrdered", 0), // property CustomFormsZOrdered
			/* 37 */ imports.NewTable("TScreen_DesktopLeft", 0), // property DesktopLeft
			/* 38 */ imports.NewTable("TScreen_DesktopTop", 0), // property DesktopTop
			/* 39 */ imports.NewTable("TScreen_DesktopHeight", 0), // property DesktopHeight
			/* 40 */ imports.NewTable("TScreen_DesktopWidth", 0), // property DesktopWidth
			/* 41 */ imports.NewTable("TScreen_DesktopRect", 0), // property DesktopRect
			/* 42 */ imports.NewTable("TScreen_FocusedForm", 0), // property FocusedForm
			/* 43 */ imports.NewTable("TScreen_FormCount", 0), // property FormCount
			/* 44 */ imports.NewTable("TScreen_Forms", 0), // property Forms
			/* 45 */ imports.NewTable("TScreen_DataModuleCount", 0), // property DataModuleCount
			/* 46 */ imports.NewTable("TScreen_DataModules", 0), // property DataModules
			/* 47 */ imports.NewTable("TScreen_HintFont", 0), // property HintFont
			/* 48 */ imports.NewTable("TScreen_IconFont", 0), // property IconFont
			/* 49 */ imports.NewTable("TScreen_MenuFont", 0), // property MenuFont
			/* 50 */ imports.NewTable("TScreen_SystemFont", 0), // property SystemFont
			/* 51 */ imports.NewTable("TScreen_Fonts", 0), // property Fonts
			/* 52 */ imports.NewTable("TScreen_Height", 0), // property Height
			/* 53 */ imports.NewTable("TScreen_MonitorCount", 0), // property MonitorCount
			/* 54 */ imports.NewTable("TScreen_Monitors", 0), // property Monitors
			/* 55 */ imports.NewTable("TScreen_PixelsPerInch", 0), // property PixelsPerInch
			/* 56 */ imports.NewTable("TScreen_PrimaryMonitor", 0), // property PrimaryMonitor
			/* 57 */ imports.NewTable("TScreen_Width", 0), // property Width
			/* 58 */ imports.NewTable("TScreen_WorkAreaRect", 0), // property WorkAreaRect
			/* 59 */ imports.NewTable("TScreen_WorkAreaHeight", 0), // property WorkAreaHeight
			/* 60 */ imports.NewTable("TScreen_WorkAreaLeft", 0), // property WorkAreaLeft
			/* 61 */ imports.NewTable("TScreen_WorkAreaTop", 0), // property WorkAreaTop
			/* 62 */ imports.NewTable("TScreen_WorkAreaWidth", 0), // property WorkAreaWidth
			/* 63 */ imports.NewTable("TScreen_OnActiveControlChange", 0), // event OnActiveControlChange
			/* 64 */ imports.NewTable("TScreen_OnActiveFormChange", 0), // event OnActiveFormChange
			/* 65 */ imports.NewTable("TScreen_TClass", 0), // function TScreenClass
		}
	})
	return screenImport
}
