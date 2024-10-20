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

// ICustomTabControl Parent: IWinControl
type ICustomTabControl interface {
	IWinControl
	HotTrack() bool                            // property
	SetHotTrack(AValue bool)                   // property
	Images() ICustomImageList                  // property
	SetImages(AValue ICustomImageList)         // property
	ImagesWidth() int32                        // property
	SetImagesWidth(AValue int32)               // property
	MultiLine() bool                           // property
	SetMultiLine(AValue bool)                  // property
	MultiSelect() bool                         // property
	SetMultiSelect(AValue bool)                // property
	Options() TCTabControlOptions              // property
	SetOptions(AValue TCTabControlOptions)     // property
	OwnerDraw() bool                           // property
	SetOwnerDraw(AValue bool)                  // property
	Page(Index int32) ICustomPage              // property
	PageCount() int32                          // property
	PageIndex() int32                          // property
	SetPageIndex(AValue int32)                 // property
	Pages() IStrings                           // property
	SetPages(AValue IStrings)                  // property
	RaggedRight() bool                         // property
	SetRaggedRight(AValue bool)                // property
	ScrollOpposite() bool                      // property
	SetScrollOpposite(AValue bool)             // property
	ShowTabs() bool                            // property
	SetShowTabs(AValue bool)                   // property
	Style() TTabStyle                          // property
	SetStyle(AValue TTabStyle)                 // property
	TabHeight() SmallInt                       // property
	SetTabHeight(AValue SmallInt)              // property
	TabPosition() TTabPosition                 // property
	SetTabPosition(AValue TTabPosition)        // property
	TabWidth() SmallInt                        // property
	SetTabWidth(AValue SmallInt)               // property
	TabRect(AIndex int32) (resultRect TRect)   // function
	GetImageIndex(ThePageIndex int32) int32    // function
	IndexOf(APage IPersistent) int32           // function
	CustomPage(Index int32) ICustomPage        // function
	CanChangePageIndex() bool                  // function
	GetMinimumTabWidth() int32                 // function
	GetMinimumTabHeight() int32                // function
	GetCapabilities() TCTabControlCapabilities // function
	TabToPageIndex(AIndex int32) int32         // function
	PageToTabIndex(AIndex int32) int32         // function
	DoCloseTabClicked(APage ICustomPage)       // procedure
	SetOnChanging(fn TTabChangingEvent)        // property event
	SetOnCloseTabClicked(fn TNotifyEvent)      // property event
	SetOnGetImageIndex(fn TTabGetImageEvent)   // property event
}

// TCustomTabControl Parent: TWinControl
type TCustomTabControl struct {
	TWinControl
	changingPtr        uintptr
	closeTabClickedPtr uintptr
	getImageIndexPtr   uintptr
}

func NewCustomTabControl(TheOwner IComponent) ICustomTabControl {
	r1 := customTabControlImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsCustomTabControl(r1)
}

func (m *TCustomTabControl) HotTrack() bool {
	r1 := customTabControlImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetHotTrack(AValue bool) {
	customTabControlImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) Images() ICustomImageList {
	r1 := customTabControlImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomTabControl) SetImages(AValue ICustomImageList) {
	customTabControlImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTabControl) ImagesWidth() int32 {
	r1 := customTabControlImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTabControl) SetImagesWidth(AValue int32) {
	customTabControlImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) MultiLine() bool {
	r1 := customTabControlImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetMultiLine(AValue bool) {
	customTabControlImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) MultiSelect() bool {
	r1 := customTabControlImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetMultiSelect(AValue bool) {
	customTabControlImportAPI().SysCallN(14, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) Options() TCTabControlOptions {
	r1 := customTabControlImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return TCTabControlOptions(r1)
}

func (m *TCustomTabControl) SetOptions(AValue TCTabControlOptions) {
	customTabControlImportAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) OwnerDraw() bool {
	r1 := customTabControlImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetOwnerDraw(AValue bool) {
	customTabControlImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) Page(Index int32) ICustomPage {
	r1 := customTabControlImportAPI().SysCallN(17, m.Instance(), uintptr(Index))
	return AsCustomPage(r1)
}

func (m *TCustomTabControl) PageCount() int32 {
	r1 := customTabControlImportAPI().SysCallN(18, m.Instance())
	return int32(r1)
}

func (m *TCustomTabControl) PageIndex() int32 {
	r1 := customTabControlImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTabControl) SetPageIndex(AValue int32) {
	customTabControlImportAPI().SysCallN(19, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) Pages() IStrings {
	r1 := customTabControlImportAPI().SysCallN(21, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomTabControl) SetPages(AValue IStrings) {
	customTabControlImportAPI().SysCallN(21, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTabControl) RaggedRight() bool {
	r1 := customTabControlImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetRaggedRight(AValue bool) {
	customTabControlImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) ScrollOpposite() bool {
	r1 := customTabControlImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetScrollOpposite(AValue bool) {
	customTabControlImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) ShowTabs() bool {
	r1 := customTabControlImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetShowTabs(AValue bool) {
	customTabControlImportAPI().SysCallN(27, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) Style() TTabStyle {
	r1 := customTabControlImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return TTabStyle(r1)
}

func (m *TCustomTabControl) SetStyle(AValue TTabStyle) {
	customTabControlImportAPI().SysCallN(28, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) TabHeight() SmallInt {
	r1 := customTabControlImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TCustomTabControl) SetTabHeight(AValue SmallInt) {
	customTabControlImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) TabPosition() TTabPosition {
	r1 := customTabControlImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return TTabPosition(r1)
}

func (m *TCustomTabControl) SetTabPosition(AValue TTabPosition) {
	customTabControlImportAPI().SysCallN(30, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) TabWidth() SmallInt {
	r1 := customTabControlImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TCustomTabControl) SetTabWidth(AValue SmallInt) {
	customTabControlImportAPI().SysCallN(33, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) TabRect(AIndex int32) (resultRect TRect) {
	customTabControlImportAPI().SysCallN(31, m.Instance(), uintptr(AIndex), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCustomTabControl) GetImageIndex(ThePageIndex int32) int32 {
	r1 := customTabControlImportAPI().SysCallN(6, m.Instance(), uintptr(ThePageIndex))
	return int32(r1)
}

func (m *TCustomTabControl) IndexOf(APage IPersistent) int32 {
	r1 := customTabControlImportAPI().SysCallN(12, m.Instance(), GetObjectUintptr(APage))
	return int32(r1)
}

func (m *TCustomTabControl) CustomPage(Index int32) ICustomPage {
	r1 := customTabControlImportAPI().SysCallN(3, m.Instance(), uintptr(Index))
	return AsCustomPage(r1)
}

func (m *TCustomTabControl) CanChangePageIndex() bool {
	r1 := customTabControlImportAPI().SysCallN(0, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTabControl) GetMinimumTabWidth() int32 {
	r1 := customTabControlImportAPI().SysCallN(8, m.Instance())
	return int32(r1)
}

func (m *TCustomTabControl) GetMinimumTabHeight() int32 {
	r1 := customTabControlImportAPI().SysCallN(7, m.Instance())
	return int32(r1)
}

func (m *TCustomTabControl) GetCapabilities() TCTabControlCapabilities {
	r1 := customTabControlImportAPI().SysCallN(5, m.Instance())
	return TCTabControlCapabilities(r1)
}

func (m *TCustomTabControl) TabToPageIndex(AIndex int32) int32 {
	r1 := customTabControlImportAPI().SysCallN(32, m.Instance(), uintptr(AIndex))
	return int32(r1)
}

func (m *TCustomTabControl) PageToTabIndex(AIndex int32) int32 {
	r1 := customTabControlImportAPI().SysCallN(20, m.Instance(), uintptr(AIndex))
	return int32(r1)
}

func CustomTabControlClass() TClass {
	ret := customTabControlImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCustomTabControl) DoCloseTabClicked(APage ICustomPage) {
	customTabControlImportAPI().SysCallN(4, m.Instance(), GetObjectUintptr(APage))
}

func (m *TCustomTabControl) SetOnChanging(fn TTabChangingEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	customTabControlImportAPI().SysCallN(24, m.Instance(), m.changingPtr)
}

func (m *TCustomTabControl) SetOnCloseTabClicked(fn TNotifyEvent) {
	if m.closeTabClickedPtr != 0 {
		RemoveEventElement(m.closeTabClickedPtr)
	}
	m.closeTabClickedPtr = MakeEventDataPtr(fn)
	customTabControlImportAPI().SysCallN(25, m.Instance(), m.closeTabClickedPtr)
}

func (m *TCustomTabControl) SetOnGetImageIndex(fn TTabGetImageEvent) {
	if m.getImageIndexPtr != 0 {
		RemoveEventElement(m.getImageIndexPtr)
	}
	m.getImageIndexPtr = MakeEventDataPtr(fn)
	customTabControlImportAPI().SysCallN(26, m.Instance(), m.getImageIndexPtr)
}

var (
	customTabControlImport       *imports.Imports = nil
	customTabControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomTabControl_CanChangePageIndex", 0),
		/*1*/ imports.NewTable("CustomTabControl_Class", 0),
		/*2*/ imports.NewTable("CustomTabControl_Create", 0),
		/*3*/ imports.NewTable("CustomTabControl_CustomPage", 0),
		/*4*/ imports.NewTable("CustomTabControl_DoCloseTabClicked", 0),
		/*5*/ imports.NewTable("CustomTabControl_GetCapabilities", 0),
		/*6*/ imports.NewTable("CustomTabControl_GetImageIndex", 0),
		/*7*/ imports.NewTable("CustomTabControl_GetMinimumTabHeight", 0),
		/*8*/ imports.NewTable("CustomTabControl_GetMinimumTabWidth", 0),
		/*9*/ imports.NewTable("CustomTabControl_HotTrack", 0),
		/*10*/ imports.NewTable("CustomTabControl_Images", 0),
		/*11*/ imports.NewTable("CustomTabControl_ImagesWidth", 0),
		/*12*/ imports.NewTable("CustomTabControl_IndexOf", 0),
		/*13*/ imports.NewTable("CustomTabControl_MultiLine", 0),
		/*14*/ imports.NewTable("CustomTabControl_MultiSelect", 0),
		/*15*/ imports.NewTable("CustomTabControl_Options", 0),
		/*16*/ imports.NewTable("CustomTabControl_OwnerDraw", 0),
		/*17*/ imports.NewTable("CustomTabControl_Page", 0),
		/*18*/ imports.NewTable("CustomTabControl_PageCount", 0),
		/*19*/ imports.NewTable("CustomTabControl_PageIndex", 0),
		/*20*/ imports.NewTable("CustomTabControl_PageToTabIndex", 0),
		/*21*/ imports.NewTable("CustomTabControl_Pages", 0),
		/*22*/ imports.NewTable("CustomTabControl_RaggedRight", 0),
		/*23*/ imports.NewTable("CustomTabControl_ScrollOpposite", 0),
		/*24*/ imports.NewTable("CustomTabControl_SetOnChanging", 0),
		/*25*/ imports.NewTable("CustomTabControl_SetOnCloseTabClicked", 0),
		/*26*/ imports.NewTable("CustomTabControl_SetOnGetImageIndex", 0),
		/*27*/ imports.NewTable("CustomTabControl_ShowTabs", 0),
		/*28*/ imports.NewTable("CustomTabControl_Style", 0),
		/*29*/ imports.NewTable("CustomTabControl_TabHeight", 0),
		/*30*/ imports.NewTable("CustomTabControl_TabPosition", 0),
		/*31*/ imports.NewTable("CustomTabControl_TabRect", 0),
		/*32*/ imports.NewTable("CustomTabControl_TabToPageIndex", 0),
		/*33*/ imports.NewTable("CustomTabControl_TabWidth", 0),
	}
)

func customTabControlImportAPI() *imports.Imports {
	if customTabControlImport == nil {
		customTabControlImport = NewDefaultImports()
		customTabControlImport.SetImportTable(customTabControlImportTables)
		customTabControlImportTables = nil
	}
	return customTabControlImport
}
