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
	r1 := LCL().SysCallN(2324, GetObjectUintptr(TheOwner))
	return AsCustomTabControl(r1)
}

func (m *TCustomTabControl) HotTrack() bool {
	r1 := LCL().SysCallN(2331, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetHotTrack(AValue bool) {
	LCL().SysCallN(2331, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) Images() ICustomImageList {
	r1 := LCL().SysCallN(2332, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomTabControl) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(2332, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTabControl) ImagesWidth() int32 {
	r1 := LCL().SysCallN(2333, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTabControl) SetImagesWidth(AValue int32) {
	LCL().SysCallN(2333, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) MultiLine() bool {
	r1 := LCL().SysCallN(2335, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetMultiLine(AValue bool) {
	LCL().SysCallN(2335, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) MultiSelect() bool {
	r1 := LCL().SysCallN(2336, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetMultiSelect(AValue bool) {
	LCL().SysCallN(2336, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) Options() TCTabControlOptions {
	r1 := LCL().SysCallN(2337, 0, m.Instance(), 0)
	return TCTabControlOptions(r1)
}

func (m *TCustomTabControl) SetOptions(AValue TCTabControlOptions) {
	LCL().SysCallN(2337, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) OwnerDraw() bool {
	r1 := LCL().SysCallN(2338, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetOwnerDraw(AValue bool) {
	LCL().SysCallN(2338, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) Page(Index int32) ICustomPage {
	r1 := LCL().SysCallN(2339, m.Instance(), uintptr(Index))
	return AsCustomPage(r1)
}

func (m *TCustomTabControl) PageCount() int32 {
	r1 := LCL().SysCallN(2340, m.Instance())
	return int32(r1)
}

func (m *TCustomTabControl) PageIndex() int32 {
	r1 := LCL().SysCallN(2341, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTabControl) SetPageIndex(AValue int32) {
	LCL().SysCallN(2341, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) Pages() IStrings {
	r1 := LCL().SysCallN(2343, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomTabControl) SetPages(AValue IStrings) {
	LCL().SysCallN(2343, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTabControl) RaggedRight() bool {
	r1 := LCL().SysCallN(2344, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetRaggedRight(AValue bool) {
	LCL().SysCallN(2344, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) ScrollOpposite() bool {
	r1 := LCL().SysCallN(2345, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetScrollOpposite(AValue bool) {
	LCL().SysCallN(2345, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) ShowTabs() bool {
	r1 := LCL().SysCallN(2349, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTabControl) SetShowTabs(AValue bool) {
	LCL().SysCallN(2349, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTabControl) Style() TTabStyle {
	r1 := LCL().SysCallN(2350, 0, m.Instance(), 0)
	return TTabStyle(r1)
}

func (m *TCustomTabControl) SetStyle(AValue TTabStyle) {
	LCL().SysCallN(2350, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) TabHeight() SmallInt {
	r1 := LCL().SysCallN(2351, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TCustomTabControl) SetTabHeight(AValue SmallInt) {
	LCL().SysCallN(2351, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) TabPosition() TTabPosition {
	r1 := LCL().SysCallN(2352, 0, m.Instance(), 0)
	return TTabPosition(r1)
}

func (m *TCustomTabControl) SetTabPosition(AValue TTabPosition) {
	LCL().SysCallN(2352, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) TabWidth() SmallInt {
	r1 := LCL().SysCallN(2355, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TCustomTabControl) SetTabWidth(AValue SmallInt) {
	LCL().SysCallN(2355, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTabControl) TabRect(AIndex int32) (resultRect TRect) {
	LCL().SysCallN(2353, m.Instance(), uintptr(AIndex), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCustomTabControl) GetImageIndex(ThePageIndex int32) int32 {
	r1 := LCL().SysCallN(2328, m.Instance(), uintptr(ThePageIndex))
	return int32(r1)
}

func (m *TCustomTabControl) IndexOf(APage IPersistent) int32 {
	r1 := LCL().SysCallN(2334, m.Instance(), GetObjectUintptr(APage))
	return int32(r1)
}

func (m *TCustomTabControl) CustomPage(Index int32) ICustomPage {
	r1 := LCL().SysCallN(2325, m.Instance(), uintptr(Index))
	return AsCustomPage(r1)
}

func (m *TCustomTabControl) CanChangePageIndex() bool {
	r1 := LCL().SysCallN(2322, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTabControl) GetMinimumTabWidth() int32 {
	r1 := LCL().SysCallN(2330, m.Instance())
	return int32(r1)
}

func (m *TCustomTabControl) GetMinimumTabHeight() int32 {
	r1 := LCL().SysCallN(2329, m.Instance())
	return int32(r1)
}

func (m *TCustomTabControl) GetCapabilities() TCTabControlCapabilities {
	r1 := LCL().SysCallN(2327, m.Instance())
	return TCTabControlCapabilities(r1)
}

func (m *TCustomTabControl) TabToPageIndex(AIndex int32) int32 {
	r1 := LCL().SysCallN(2354, m.Instance(), uintptr(AIndex))
	return int32(r1)
}

func (m *TCustomTabControl) PageToTabIndex(AIndex int32) int32 {
	r1 := LCL().SysCallN(2342, m.Instance(), uintptr(AIndex))
	return int32(r1)
}

func CustomTabControlClass() TClass {
	ret := LCL().SysCallN(2323)
	return TClass(ret)
}

func (m *TCustomTabControl) DoCloseTabClicked(APage ICustomPage) {
	LCL().SysCallN(2326, m.Instance(), GetObjectUintptr(APage))
}

func (m *TCustomTabControl) SetOnChanging(fn TTabChangingEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2346, m.Instance(), m.changingPtr)
}

func (m *TCustomTabControl) SetOnCloseTabClicked(fn TNotifyEvent) {
	if m.closeTabClickedPtr != 0 {
		RemoveEventElement(m.closeTabClickedPtr)
	}
	m.closeTabClickedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2347, m.Instance(), m.closeTabClickedPtr)
}

func (m *TCustomTabControl) SetOnGetImageIndex(fn TTabGetImageEvent) {
	if m.getImageIndexPtr != 0 {
		RemoveEventElement(m.getImageIndexPtr)
	}
	m.getImageIndexPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2348, m.Instance(), m.getImageIndexPtr)
}
