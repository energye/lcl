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

// ICustomListView Parent: IWinControl
type ICustomListView interface {
	IWinControl
	AlphaSort() bool                                                                                               // function
	FindCaption(startIndex int32, value string, partial bool, inclusive bool, wrap bool, partStart bool) IListItem // function
	FindData(startIndex int32, value uintptr, inclusive bool, wrap bool) IListItem                                 // function
	GetHitTestInfoAt(X int32, Y int32) types.THitTests                                                             // function
	GetItemAt(X int32, Y int32) IListItem                                                                          // function
	// GetNearestItem
	//  GetNearestItem is used to locate a list item from a position specified in
	//  pixel coordinates relative to the top left corner of the list view.
	//  It starts looking at the position specified by the Point parameter,
	//  and moves in the direction indicated by the Direction parameter
	//  until it locates a list item.If no item is found Nil is returned.
	GetNearestItem(point types.TPoint, direction types.TSearchDirection) IListItem // function
	// GetNextItem
	//  Used to find the next list item after StartItem in the direction
	//  given by the Direction parameter.
	//  Only items in the state indicated by the States parameter are considered.
	//  If no item is found Nil is returned.
	GetNextItem(startItem IListItem, direction types.TSearchDirection, states types.TListItemStates) IListItem // function
	// GetNextSelected
	//  Returns True if it has, authoritivly, set AItem to the next selected Item.
	GetNextSelected(lVIndex int32, item *IListItem) bool // function
	IsEditing() bool                                     // function
	AddItem(item string, object IObject)                 // procedure
	Sort()                                               // procedure
	BeginUpdate()                                        // procedure
	Clear()                                              // procedure
	EndUpdate()                                          // procedure
	ClearSelection()                                     // procedure
	SelectAll()                                          // procedure
	BoundingRect() types.TRect                           // property BoundingRect Getter
	// BorderStyle
	//  properties which are not supported by all descendents
	BorderStyle() types.TBorderStyle                   // property BorderStyle Getter
	SetBorderStyle(value types.TBorderStyle)           // property BorderStyle Setter
	Canvas() ICanvas                                   // property Canvas Getter
	Checkboxes() bool                                  // property Checkboxes Getter
	SetCheckboxes(value bool)                          // property Checkboxes Setter
	Column(index int32) IListColumn                    // property Column Getter
	ColumnCount() int32                                // property ColumnCount Getter
	DropTarget() IListItem                             // property DropTarget Getter
	SetDropTarget(value IListItem)                     // property DropTarget Setter
	FlatScrollBars() bool                              // property FlatScrollBars Getter
	SetFlatScrollBars(value bool)                      // property FlatScrollBars Setter
	FullDrag() bool                                    // property FullDrag Getter
	SetFullDrag(value bool)                            // property FullDrag Setter
	GridLines() bool                                   // property GridLines Getter
	SetGridLines(value bool)                           // property GridLines Setter
	HotTrack() bool                                    // property HotTrack Getter
	SetHotTrack(value bool)                            // property HotTrack Setter
	HotTrackStyles() types.TListHotTrackStyles         // property HotTrackStyles Getter
	SetHotTrackStyles(value types.TListHotTrackStyles) // property HotTrackStyles Setter
	IconOptions() IIconOptions                         // property IconOptions Getter
	SetIconOptions(value IIconOptions)                 // property IconOptions Setter
	ItemFocused() IListItem                            // property ItemFocused Getter
	SetItemFocused(value IListItem)                    // property ItemFocused Setter
	ItemIndex() int32                                  // property ItemIndex Getter
	SetItemIndex(value int32)                          // property ItemIndex Setter
	Items() IListItems                                 // property Items Getter
	SetItems(value IListItems)                         // property Items Setter
	// MultiSelect
	//  MultiSelect and ReadOnly should be protected, but can't because Carbon Interface
	//  needs to access this property and it cannot cast to TListItem, because we have
	//  other classes descending from TCustomListItem which need to work too
	MultiSelect() bool                // property MultiSelect Getter
	SetMultiSelect(value bool)        // property MultiSelect Setter
	OwnerData() bool                  // property OwnerData Getter
	SetOwnerData(value bool)          // property OwnerData Setter
	ReadOnly() bool                   // property ReadOnly Getter
	SetReadOnly(value bool)           // property ReadOnly Setter
	RowSelect() bool                  // property RowSelect Getter
	SetRowSelect(value bool)          // property RowSelect Setter
	SelCount() int32                  // property SelCount Getter
	Selected() IListItem              // property Selected Getter
	SetSelected(value IListItem)      // property Selected Setter
	LastSelected() IListItem          // property LastSelected Getter
	TopItem() IListItem               // property TopItem Getter
	ViewOrigin() types.TPoint         // property ViewOrigin Getter
	SetViewOrigin(value types.TPoint) // property ViewOrigin Setter
	VisibleRowCount() int32           // property VisibleRowCount Getter
}

type TCustomListView struct {
	TWinControl
}

func (m *TCustomListView) AlphaSort() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(1, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) FindCaption(startIndex int32, value string, partial bool, inclusive bool, wrap bool, partStart bool) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(2, m.Instance(), uintptr(startIndex), api.PasStr(value), api.PasBool(partial), api.PasBool(inclusive), api.PasBool(wrap), api.PasBool(partStart))
	return AsListItem(r)
}

func (m *TCustomListView) FindData(startIndex int32, value uintptr, inclusive bool, wrap bool) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(3, m.Instance(), uintptr(startIndex), uintptr(value), api.PasBool(inclusive), api.PasBool(wrap))
	return AsListItem(r)
}

func (m *TCustomListView) GetHitTestInfoAt(X int32, Y int32) types.THitTests {
	if !m.IsValid() {
		return 0
	}
	r := customListViewAPI().SysCallN(4, m.Instance(), uintptr(X), uintptr(Y))
	return types.THitTests(r)
}

func (m *TCustomListView) GetItemAt(X int32, Y int32) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(5, m.Instance(), uintptr(X), uintptr(Y))
	return AsListItem(r)
}

func (m *TCustomListView) GetNearestItem(point types.TPoint, direction types.TSearchDirection) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(6, m.Instance(), uintptr(base.UnsafePointer(&point)), uintptr(direction))
	return AsListItem(r)
}

func (m *TCustomListView) GetNextItem(startItem IListItem, direction types.TSearchDirection, states types.TListItemStates) IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(7, m.Instance(), base.GetObjectUintptr(startItem), uintptr(direction), uintptr(states))
	return AsListItem(r)
}

func (m *TCustomListView) GetNextSelected(lVIndex int32, item *IListItem) bool {
	if !m.IsValid() {
		return false
	}
	itemPtr := base.GetObjectUintptr(*item)
	r := customListViewAPI().SysCallN(8, m.Instance(), uintptr(lVIndex), uintptr(base.UnsafePointer(&itemPtr)))
	*item = AsListItem(itemPtr)
	return api.GoBool(r)
}

func (m *TCustomListView) IsEditing() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(9, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) AddItem(item string, object IObject) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(10, m.Instance(), api.PasStr(item), base.GetObjectUintptr(object))
}

func (m *TCustomListView) Sort() {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(11, m.Instance())
}

func (m *TCustomListView) BeginUpdate() {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(12, m.Instance())
}

func (m *TCustomListView) Clear() {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(13, m.Instance())
}

func (m *TCustomListView) EndUpdate() {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(14, m.Instance())
}

func (m *TCustomListView) ClearSelection() {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(15, m.Instance())
}

func (m *TCustomListView) SelectAll() {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(16, m.Instance())
}

func (m *TCustomListView) BoundingRect() (result types.TRect) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(17, m.Instance(), uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomListView) BorderStyle() types.TBorderStyle {
	if !m.IsValid() {
		return 0
	}
	r := customListViewAPI().SysCallN(18, 0, m.Instance())
	return types.TBorderStyle(r)
}

func (m *TCustomListView) SetBorderStyle(value types.TBorderStyle) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(18, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListView) Canvas() ICanvas {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(19, m.Instance())
	return AsCanvas(r)
}

func (m *TCustomListView) Checkboxes() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(20, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) SetCheckboxes(value bool) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(20, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListView) Column(index int32) IListColumn {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(21, m.Instance(), uintptr(index))
	return AsListColumn(r)
}

func (m *TCustomListView) ColumnCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListViewAPI().SysCallN(22, m.Instance())
	return int32(r)
}

func (m *TCustomListView) DropTarget() IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(23, 0, m.Instance())
	return AsListItem(r)
}

func (m *TCustomListView) SetDropTarget(value IListItem) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(23, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomListView) FlatScrollBars() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(24, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) SetFlatScrollBars(value bool) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(24, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListView) FullDrag() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(25, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) SetFullDrag(value bool) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(25, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListView) GridLines() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(26, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) SetGridLines(value bool) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(26, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListView) HotTrack() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(27, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) SetHotTrack(value bool) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(27, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListView) HotTrackStyles() types.TListHotTrackStyles {
	if !m.IsValid() {
		return 0
	}
	r := customListViewAPI().SysCallN(28, 0, m.Instance())
	return types.TListHotTrackStyles(r)
}

func (m *TCustomListView) SetHotTrackStyles(value types.TListHotTrackStyles) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(28, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListView) IconOptions() IIconOptions {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(29, 0, m.Instance())
	return AsIconOptions(r)
}

func (m *TCustomListView) SetIconOptions(value IIconOptions) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(29, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomListView) ItemFocused() IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(30, 0, m.Instance())
	return AsListItem(r)
}

func (m *TCustomListView) SetItemFocused(value IListItem) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(30, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomListView) ItemIndex() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListViewAPI().SysCallN(31, 0, m.Instance())
	return int32(r)
}

func (m *TCustomListView) SetItemIndex(value int32) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(31, 1, m.Instance(), uintptr(value))
}

func (m *TCustomListView) Items() IListItems {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(32, 0, m.Instance())
	return AsListItems(r)
}

func (m *TCustomListView) SetItems(value IListItems) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(32, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomListView) MultiSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(33, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) SetMultiSelect(value bool) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(33, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListView) OwnerData() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(34, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) SetOwnerData(value bool) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(34, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListView) ReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(35, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) SetReadOnly(value bool) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(35, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListView) RowSelect() bool {
	if !m.IsValid() {
		return false
	}
	r := customListViewAPI().SysCallN(36, 0, m.Instance())
	return api.GoBool(r)
}

func (m *TCustomListView) SetRowSelect(value bool) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(36, 1, m.Instance(), api.PasBool(value))
}

func (m *TCustomListView) SelCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListViewAPI().SysCallN(37, m.Instance())
	return int32(r)
}

func (m *TCustomListView) Selected() IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(38, 0, m.Instance())
	return AsListItem(r)
}

func (m *TCustomListView) SetSelected(value IListItem) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(38, 1, m.Instance(), base.GetObjectUintptr(value))
}

func (m *TCustomListView) LastSelected() IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(39, m.Instance())
	return AsListItem(r)
}

func (m *TCustomListView) TopItem() IListItem {
	if !m.IsValid() {
		return nil
	}
	r := customListViewAPI().SysCallN(40, m.Instance())
	return AsListItem(r)
}

func (m *TCustomListView) ViewOrigin() (result types.TPoint) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(41, 0, m.Instance(), 0, uintptr(base.UnsafePointer(&result)))
	return
}

func (m *TCustomListView) SetViewOrigin(value types.TPoint) {
	if !m.IsValid() {
		return
	}
	customListViewAPI().SysCallN(41, 1, m.Instance(), uintptr(base.UnsafePointer(&value)))
}

func (m *TCustomListView) VisibleRowCount() int32 {
	if !m.IsValid() {
		return 0
	}
	r := customListViewAPI().SysCallN(42, m.Instance())
	return int32(r)
}

// NewCustomListView class constructor
func NewCustomListView(owner IComponent) ICustomListView {
	r := customListViewAPI().SysCallN(0, base.GetObjectUintptr(owner))
	return AsCustomListView(r)
}

func TCustomListViewClass() types.TClass {
	r := customListViewAPI().SysCallN(43)
	return types.TClass(r)
}

var (
	customListViewOnce   base.Once
	customListViewImport *imports.Imports = nil
)

func customListViewAPI() *imports.Imports {
	customListViewOnce.Do(func() {
		customListViewImport = api.NewDefaultImports()
		customListViewImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCustomListView_Create", 0), // constructor NewCustomListView
			/* 1 */ imports.NewTable("TCustomListView_AlphaSort", 0), // function AlphaSort
			/* 2 */ imports.NewTable("TCustomListView_FindCaption", 0), // function FindCaption
			/* 3 */ imports.NewTable("TCustomListView_FindData", 0), // function FindData
			/* 4 */ imports.NewTable("TCustomListView_GetHitTestInfoAt", 0), // function GetHitTestInfoAt
			/* 5 */ imports.NewTable("TCustomListView_GetItemAt", 0), // function GetItemAt
			/* 6 */ imports.NewTable("TCustomListView_GetNearestItem", 0), // function GetNearestItem
			/* 7 */ imports.NewTable("TCustomListView_GetNextItem", 0), // function GetNextItem
			/* 8 */ imports.NewTable("TCustomListView_GetNextSelected", 0), // function GetNextSelected
			/* 9 */ imports.NewTable("TCustomListView_IsEditing", 0), // function IsEditing
			/* 10 */ imports.NewTable("TCustomListView_AddItem", 0), // procedure AddItem
			/* 11 */ imports.NewTable("TCustomListView_Sort", 0), // procedure Sort
			/* 12 */ imports.NewTable("TCustomListView_BeginUpdate", 0), // procedure BeginUpdate
			/* 13 */ imports.NewTable("TCustomListView_Clear", 0), // procedure Clear
			/* 14 */ imports.NewTable("TCustomListView_EndUpdate", 0), // procedure EndUpdate
			/* 15 */ imports.NewTable("TCustomListView_ClearSelection", 0), // procedure ClearSelection
			/* 16 */ imports.NewTable("TCustomListView_SelectAll", 0), // procedure SelectAll
			/* 17 */ imports.NewTable("TCustomListView_BoundingRect", 0), // property BoundingRect
			/* 18 */ imports.NewTable("TCustomListView_BorderStyle", 0), // property BorderStyle
			/* 19 */ imports.NewTable("TCustomListView_Canvas", 0), // property Canvas
			/* 20 */ imports.NewTable("TCustomListView_Checkboxes", 0), // property Checkboxes
			/* 21 */ imports.NewTable("TCustomListView_Column", 0), // property Column
			/* 22 */ imports.NewTable("TCustomListView_ColumnCount", 0), // property ColumnCount
			/* 23 */ imports.NewTable("TCustomListView_DropTarget", 0), // property DropTarget
			/* 24 */ imports.NewTable("TCustomListView_FlatScrollBars", 0), // property FlatScrollBars
			/* 25 */ imports.NewTable("TCustomListView_FullDrag", 0), // property FullDrag
			/* 26 */ imports.NewTable("TCustomListView_GridLines", 0), // property GridLines
			/* 27 */ imports.NewTable("TCustomListView_HotTrack", 0), // property HotTrack
			/* 28 */ imports.NewTable("TCustomListView_HotTrackStyles", 0), // property HotTrackStyles
			/* 29 */ imports.NewTable("TCustomListView_IconOptions", 0), // property IconOptions
			/* 30 */ imports.NewTable("TCustomListView_ItemFocused", 0), // property ItemFocused
			/* 31 */ imports.NewTable("TCustomListView_ItemIndex", 0), // property ItemIndex
			/* 32 */ imports.NewTable("TCustomListView_Items", 0), // property Items
			/* 33 */ imports.NewTable("TCustomListView_MultiSelect", 0), // property MultiSelect
			/* 34 */ imports.NewTable("TCustomListView_OwnerData", 0), // property OwnerData
			/* 35 */ imports.NewTable("TCustomListView_ReadOnly", 0), // property ReadOnly
			/* 36 */ imports.NewTable("TCustomListView_RowSelect", 0), // property RowSelect
			/* 37 */ imports.NewTable("TCustomListView_SelCount", 0), // property SelCount
			/* 38 */ imports.NewTable("TCustomListView_Selected", 0), // property Selected
			/* 39 */ imports.NewTable("TCustomListView_LastSelected", 0), // property LastSelected
			/* 40 */ imports.NewTable("TCustomListView_TopItem", 0), // property TopItem
			/* 41 */ imports.NewTable("TCustomListView_ViewOrigin", 0), // property ViewOrigin
			/* 42 */ imports.NewTable("TCustomListView_VisibleRowCount", 0), // property VisibleRowCount
			/* 43 */ imports.NewTable("TCustomListView_TClass", 0), // function TCustomListViewClass
		}
	})
	return customListViewImport
}
