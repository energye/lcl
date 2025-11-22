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

// ICustomTabControl Parent: IWinControl
type ICustomTabControl interface {
	IWinControl
	TabRect(index int32) types.TRect                 // function
	GetImageIndex(thePageIndex int32) int32          // function
	IndexOf(page IPersistent) int32                  // function
	CustomPage(index int32) ICustomPage              // function
	CanChangePageIndex() bool                        // function
	GetMinimumTabWidth() int32                       // function
	GetMinimumTabHeight() int32                      // function
	GetCapabilities() types.TCTabControlCapabilities // function
	TabToPageIndex(index int32) int32                // function
	PageToTabIndex(index int32) int32                // function
	DoCloseTabClicked(page ICustomPage)              // procedure
	HotTrack() bool                                  // property HotTrack Getter
	SetHotTrack(value bool)                          // property HotTrack Setter
	Images() ICustomImageList                        // property Images Getter
	SetImages(value ICustomImageList)                // property Images Setter
	ImagesWidth() int32                              // property ImagesWidth Getter
	SetImagesWidth(value int32)                      // property ImagesWidth Setter
	MultiLine() bool                                 // property MultiLine Getter
	SetMultiLine(value bool)                         // property MultiLine Setter
	MultiSelect() bool                               // property MultiSelect Getter
	SetMultiSelect(value bool)                       // property MultiSelect Setter
	Options() types.TCTabControlOptions              // property Options Getter
	SetOptions(value types.TCTabControlOptions)      // property Options Setter
	OwnerDraw() bool                                 // property OwnerDraw Getter
	SetOwnerDraw(value bool)                         // property OwnerDraw Setter
	Page(index int32) ICustomPage                    // property Page Getter
	PageCount() int32                                // property PageCount Getter
	PageIndex() int32                                // property PageIndex Getter
	SetPageIndex(value int32)                        // property PageIndex Setter
	// PagesToStrings
	//  property PageList: TList read FPageList; - iff paged
	PagesToStrings() IStrings                // property Pages Getter
	SetPages(value IStrings)                 // property Pages Setter
	RaggedRight() bool                       // property RaggedRight Getter
	SetRaggedRight(value bool)               // property RaggedRight Setter
	ScrollOpposite() bool                    // property ScrollOpposite Getter
	SetScrollOpposite(value bool)            // property ScrollOpposite Setter
	ShowTabs() bool                          // property ShowTabs Getter
	SetShowTabs(value bool)                  // property ShowTabs Setter
	Style() types.TTabStyle                  // property Style Getter
	SetStyle(value types.TTabStyle)          // property Style Setter
	TabHeight() types.SmallInt               // property TabHeight Getter
	SetTabHeight(value types.SmallInt)       // property TabHeight Setter
	TabPosition() types.TTabPosition         // property TabPosition Getter
	SetTabPosition(value types.TTabPosition) // property TabPosition Setter
	TabWidth() types.SmallInt                // property TabWidth Getter
	SetTabWidth(value types.SmallInt)        // property TabWidth Setter
	SetOnChanging(fn TTabChangingEvent)      // property event
	SetOnCloseTabClicked(fn TNotifyEvent)    // property event
	SetOnGetImageIndex(fn TTabGetImageEvent) // property event
}

type TCustomTabControl struct {
	TWinControl
}

func (m *TCustomTabControl) TabRect(index int32) (result types.TRect) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(1, m.Instance(), uintptr(index), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomTabControl) GetImageIndex(thePageIndex int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(2, m.Instance(), uintptr(thePageIndex))
	return int32(r)
}

func (m *TCustomTabControl) IndexOf(page IPersistent) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(3, m.Instance(), base.GetObjectUintptr(page))
	return int32(r)
}

func (m *TCustomTabControl) CustomPage(index int32) ICustomPage {
	if !m.IsValid() {
		return nil
	}
	r := customTabControlAPI().SysCallN(4, m.Instance(), uintptr(index))
	return AsCustomPage(r)
}

func (m *TCustomTabControl) CanChangePageIndex() bool {
	if !m.IsValid() {
		return false
	}
	r := customTabControlAPI().SysCallN(5, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTabControl) GetMinimumTabWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(6, m.Instance())
	return int32(r)
}

func (m *TCustomTabControl) GetMinimumTabHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(7, m.Instance())
	return int32(r)
}

func (m *TCustomTabControl) GetCapabilities() types.TCTabControlCapabilities {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(8, m.Instance())
	return types.TCTabControlCapabilities(r)
}

func (m *TCustomTabControl) TabToPageIndex(index int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(9, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TCustomTabControl) PageToTabIndex(index int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(10, m.Instance(), uintptr(index))
	return int32(r)
}

func (m *TCustomTabControl) DoCloseTabClicked(page ICustomPage) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(11, m.Instance(), base.GetObjectUintptr(page))
}

func (m *TCustomTabControl) HotTrack() bool {
	if !m.IsValid() {
		return false
	}
	r := customTabControlAPI().SysCallN(12, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTabControl) SetHotTrack(value bool) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(12, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTabControl) Images() ICustomImageList {
	if !m.IsValid() {
		return nil
	}
	r := customTabControlAPI().SysCallN(13, 0, m.Instance())
	return AsCustomImageList(r)
}

func (m *TCustomTabControl) SetImages(value ICustomImageList) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(13, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTabControl) ImagesWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(14, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTabControl) SetImagesWidth(value int32) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(14, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTabControl) MultiLine() bool {
	if !m.IsValid() {
		return false
	}
	r := customTabControlAPI().SysCallN(15, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTabControl) SetMultiLine(value bool) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(15, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTabControl) MultiSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := customTabControlAPI().SysCallN(16, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTabControl) SetMultiSelect(value bool) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(16, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTabControl) Options() types.TCTabControlOptions {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(17, 0, m.Instance())
	return types.TCTabControlOptions(r)
}

func (m *TCustomTabControl) SetOptions(value types.TCTabControlOptions) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(17, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTabControl) OwnerDraw() bool {
	if !m.IsValid() {
		return false
	}
	r := customTabControlAPI().SysCallN(18, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTabControl) SetOwnerDraw(value bool) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(18, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTabControl) Page(index int32) ICustomPage {
	if !m.IsValid() {
		return nil
	}
	r := customTabControlAPI().SysCallN(19, m.Instance(), uintptr(index))
	return AsCustomPage(r)
}

func (m *TCustomTabControl) PageCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(20, m.Instance())
	return int32(r)
}

func (m *TCustomTabControl) PageIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(21, 0, m.Instance())
	return int32(r)
}

func (m *TCustomTabControl) SetPageIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(21, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTabControl) PagesToStrings() IStrings {
	if !m.IsValid() {
		return nil
	}
	r := customTabControlAPI().SysCallN(22, 0, m.Instance())
	return AsStrings(r)
}

func (m *TCustomTabControl) SetPages(value IStrings) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(22, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomTabControl) RaggedRight() bool {
	if !m.IsValid() {
		return false
	}
	r := customTabControlAPI().SysCallN(23, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTabControl) SetRaggedRight(value bool) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(23, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTabControl) ScrollOpposite() bool {
	if !m.IsValid() {
		return false
	}
	r := customTabControlAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTabControl) SetScrollOpposite(value bool) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTabControl) ShowTabs() bool {
	if !m.IsValid() {
		return false
	}
	r := customTabControlAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomTabControl) SetShowTabs(value bool) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomTabControl) Style() types.TTabStyle {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(26, 0, m.Instance())
	return types.TTabStyle(r)
}

func (m *TCustomTabControl) SetStyle(value types.TTabStyle) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(26, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTabControl) TabHeight() types.SmallInt {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(27, 0, m.Instance())
	return types.SmallInt(r)
}

func (m *TCustomTabControl) SetTabHeight(value types.SmallInt) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(27, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTabControl) TabPosition() types.TTabPosition {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(28, 0, m.Instance())
	return types.TTabPosition(r)
}

func (m *TCustomTabControl) SetTabPosition(value types.TTabPosition) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTabControl) TabWidth() types.SmallInt {
	if !m.IsValid() {
		return 0
	}
	r := customTabControlAPI().SysCallN(29, 0, m.Instance())
	return types.SmallInt(r)
}

func (m *TCustomTabControl) SetTabWidth(value types.SmallInt) {
	if !m.IsValid() {
		return
	}
	customTabControlAPI().SysCallN(29, 1, m.Instance(), uintptr(value))
}

func (m *TCustomTabControl) SetOnChanging(fn TTabChangingEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTabChangingEvent(fn)
	base.SetEvent(m, 30, customTabControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTabControl) SetOnCloseTabClicked(fn TNotifyEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTNotifyEvent(fn)
	base.SetEvent(m, 31, customTabControlAPI(), api.MakeEventDataPtr(cb))
}

func (m *TCustomTabControl) SetOnGetImageIndex(fn TTabGetImageEvent) {
	if !m.IsValid() {
		return
	}
	cb := makeTTabGetImageEvent(fn)
	base.SetEvent(m, 32, customTabControlAPI(), api.MakeEventDataPtr(cb))
}

// NewCustomTabControl class constructor
func NewCustomTabControl(theOwner IComponent) ICustomTabControl {
	r := customTabControlAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsCustomTabControl(r)
}

func TCustomTabControlClass() types.TClass {
	r := customTabControlAPI().SysCallN(33)
	return types.TClass(r)
}

var (
	customTabControlOnce   base.Once
	customTabControlImport *imports.Imports = nil
)

func customTabControlAPI() *imports.Imports {
	customTabControlOnce.Do(func() {
		customTabControlImport = api.NewDefaultImports()
		customTabControlImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomTabControl_Create", 0), // constructor NewCustomTabControl
			/* 1 */ imports.NewTable("TCustomTabControl_TabRect", 0), // function TabRect
			/* 2 */ imports.NewTable("TCustomTabControl_GetImageIndex", 0), // function GetImageIndex
			/* 3 */ imports.NewTable("TCustomTabControl_IndexOf", 0), // function IndexOf
			/* 4 */ imports.NewTable("TCustomTabControl_CustomPage", 0), // function CustomPage
			/* 5 */ imports.NewTable("TCustomTabControl_CanChangePageIndex", 0), // function CanChangePageIndex
			/* 6 */ imports.NewTable("TCustomTabControl_GetMinimumTabWidth", 0), // function GetMinimumTabWidth
			/* 7 */ imports.NewTable("TCustomTabControl_GetMinimumTabHeight", 0), // function GetMinimumTabHeight
			/* 8 */ imports.NewTable("TCustomTabControl_GetCapabilities", 0), // function GetCapabilities
			/* 9 */ imports.NewTable("TCustomTabControl_TabToPageIndex", 0), // function TabToPageIndex
			/* 10 */ imports.NewTable("TCustomTabControl_PageToTabIndex", 0), // function PageToTabIndex
			/* 11 */ imports.NewTable("TCustomTabControl_DoCloseTabClicked", 0), // procedure DoCloseTabClicked
			/* 12 */ imports.NewTable("TCustomTabControl_HotTrack", 0), // property HotTrack
			/* 13 */ imports.NewTable("TCustomTabControl_Images", 0), // property Images
			/* 14 */ imports.NewTable("TCustomTabControl_ImagesWidth", 0), // property ImagesWidth
			/* 15 */ imports.NewTable("TCustomTabControl_MultiLine", 0), // property MultiLine
			/* 16 */ imports.NewTable("TCustomTabControl_MultiSelect", 0), // property MultiSelect
			/* 17 */ imports.NewTable("TCustomTabControl_Options", 0), // property Options
			/* 18 */ imports.NewTable("TCustomTabControl_OwnerDraw", 0), // property OwnerDraw
			/* 19 */ imports.NewTable("TCustomTabControl_Page", 0), // property Page
			/* 20 */ imports.NewTable("TCustomTabControl_PageCount", 0), // property PageCount
			/* 21 */ imports.NewTable("TCustomTabControl_PageIndex", 0), // property PageIndex
			/* 22 */ imports.NewTable("TCustomTabControl_PagesToStrings", 0), // property PagesToStrings
			/* 23 */ imports.NewTable("TCustomTabControl_RaggedRight", 0), // property RaggedRight
			/* 24 */ imports.NewTable("TCustomTabControl_ScrollOpposite", 0), // property ScrollOpposite
			/* 25 */ imports.NewTable("TCustomTabControl_ShowTabs", 0), // property ShowTabs
			/* 26 */ imports.NewTable("TCustomTabControl_Style", 0), // property Style
			/* 27 */ imports.NewTable("TCustomTabControl_TabHeight", 0), // property TabHeight
			/* 28 */ imports.NewTable("TCustomTabControl_TabPosition", 0), // property TabPosition
			/* 29 */ imports.NewTable("TCustomTabControl_TabWidth", 0), // property TabWidth
			/* 30 */ imports.NewTable("TCustomTabControl_OnChanging", 0), // event OnChanging
			/* 31 */ imports.NewTable("TCustomTabControl_OnCloseTabClicked", 0), // event OnCloseTabClicked
			/* 32 */ imports.NewTable("TCustomTabControl_OnGetImageIndex", 0), // event OnGetImageIndex
			/* 33 */ imports.NewTable("TCustomTabControl_TClass", 0), // function TCustomTabControlClass
		}
	})
	return customTabControlImport
}
