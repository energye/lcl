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

// ICustomListView Parent: IWinControl
type ICustomListView interface {
	IWinControl
	BoundingRect() (resultRect TRect)                                                                    // property
	BorderStyle() TBorderStyle                                                                           // property
	SetBorderStyle(AValue TBorderStyle)                                                                  // property
	Canvas() ICanvas                                                                                     // property
	Checkboxes() bool                                                                                    // property
	SetCheckboxes(AValue bool)                                                                           // property
	Column(AIndex int32) IListColumn                                                                     // property
	ColumnCount() int32                                                                                  // property
	DropTarget() IListItem                                                                               // property
	SetDropTarget(AValue IListItem)                                                                      // property
	FlatScrollBars() bool                                                                                // property
	SetFlatScrollBars(AValue bool)                                                                       // property
	FullDrag() bool                                                                                      // property
	SetFullDrag(AValue bool)                                                                             // property
	GridLines() bool                                                                                     // property
	SetGridLines(AValue bool)                                                                            // property
	HotTrack() bool                                                                                      // property
	SetHotTrack(AValue bool)                                                                             // property
	HotTrackStyles() TListHotTrackStyles                                                                 // property
	SetHotTrackStyles(AValue TListHotTrackStyles)                                                        // property
	IconOptions() IIconOptions                                                                           // property
	SetIconOptions(AValue IIconOptions)                                                                  // property
	ItemFocused() IListItem                                                                              // property
	SetItemFocused(AValue IListItem)                                                                     // property
	ItemIndex() int32                                                                                    // property
	SetItemIndex(AValue int32)                                                                           // property
	Items() IListItems                                                                                   // property
	SetItems(AValue IListItems)                                                                          // property
	MultiSelect() bool                                                                                   // property
	SetMultiSelect(AValue bool)                                                                          // property
	OwnerData() bool                                                                                     // property
	SetOwnerData(AValue bool)                                                                            // property
	ReadOnly() bool                                                                                      // property
	SetReadOnly(AValue bool)                                                                             // property
	RowSelect() bool                                                                                     // property
	SetRowSelect(AValue bool)                                                                            // property
	SelCount() int32                                                                                     // property
	Selected() IListItem                                                                                 // property
	SetSelected(AValue IListItem)                                                                        // property
	LastSelected() IListItem                                                                             // property
	TopItem() IListItem                                                                                  // property
	ViewOrigin() (resultPoint TPoint)                                                                    // property
	SetViewOrigin(AValue *TPoint)                                                                        // property
	VisibleRowCount() int32                                                                              // property
	AlphaSort() bool                                                                                     // function
	CustomSort(fn TLVCompare, AOptionalParam uint32) bool                                                // function
	FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem // function
	FindData(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem                            // function
	GetHitTestInfoAt(X, Y int32) THitTests                                                               // function
	GetItemAt(x, y int32) IListItem                                                                      // function
	GetNearestItem(APoint *TPoint, Direction TSearchDirection) IListItem                                 // function
	GetNextItem(StartItem IListItem, Direction TSearchDirection, States TListItemStates) IListItem       // function
	IsEditing() bool                                                                                     // function
	AddItem(Item string, AObject IObject)                                                                // procedure
	Sort()                                                                                               // procedure
	BeginUpdate()                                                                                        // procedure
	Clear()                                                                                              // procedure
	EndUpdate()                                                                                          // procedure
	ClearSelection()                                                                                     // procedure
	SelectAll()                                                                                          // procedure
}

// TCustomListView Parent: TWinControl
type TCustomListView struct {
	TWinControl
	customSortPtr uintptr
}

func NewCustomListView(AOwner IComponent) ICustomListView {
	r1 := customListViewImportAPI().SysCallN(12, GetObjectUintptr(AOwner))
	return AsCustomListView(r1)
}

func (m *TCustomListView) BoundingRect() (resultRect TRect) {
	customListViewImportAPI().SysCallN(4, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCustomListView) BorderStyle() TBorderStyle {
	r1 := customListViewImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomListView) SetBorderStyle(AValue TBorderStyle) {
	customListViewImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListView) Canvas() ICanvas {
	r1 := customListViewImportAPI().SysCallN(5, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomListView) Checkboxes() bool {
	r1 := customListViewImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetCheckboxes(AValue bool) {
	customListViewImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) Column(AIndex int32) IListColumn {
	r1 := customListViewImportAPI().SysCallN(10, m.Instance(), uintptr(AIndex))
	return AsListColumn(r1)
}

func (m *TCustomListView) ColumnCount() int32 {
	r1 := customListViewImportAPI().SysCallN(11, m.Instance())
	return int32(r1)
}

func (m *TCustomListView) DropTarget() IListItem {
	r1 := customListViewImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return AsListItem(r1)
}

func (m *TCustomListView) SetDropTarget(AValue IListItem) {
	customListViewImportAPI().SysCallN(14, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) FlatScrollBars() bool {
	r1 := customListViewImportAPI().SysCallN(18, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetFlatScrollBars(AValue bool) {
	customListViewImportAPI().SysCallN(18, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) FullDrag() bool {
	r1 := customListViewImportAPI().SysCallN(19, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetFullDrag(AValue bool) {
	customListViewImportAPI().SysCallN(19, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) GridLines() bool {
	r1 := customListViewImportAPI().SysCallN(24, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetGridLines(AValue bool) {
	customListViewImportAPI().SysCallN(24, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) HotTrack() bool {
	r1 := customListViewImportAPI().SysCallN(25, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetHotTrack(AValue bool) {
	customListViewImportAPI().SysCallN(25, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) HotTrackStyles() TListHotTrackStyles {
	r1 := customListViewImportAPI().SysCallN(26, 0, m.Instance(), 0)
	return TListHotTrackStyles(r1)
}

func (m *TCustomListView) SetHotTrackStyles(AValue TListHotTrackStyles) {
	customListViewImportAPI().SysCallN(26, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListView) IconOptions() IIconOptions {
	r1 := customListViewImportAPI().SysCallN(27, 0, m.Instance(), 0)
	return AsIconOptions(r1)
}

func (m *TCustomListView) SetIconOptions(AValue IIconOptions) {
	customListViewImportAPI().SysCallN(27, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) ItemFocused() IListItem {
	r1 := customListViewImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return AsListItem(r1)
}

func (m *TCustomListView) SetItemFocused(AValue IListItem) {
	customListViewImportAPI().SysCallN(29, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) ItemIndex() int32 {
	r1 := customListViewImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListView) SetItemIndex(AValue int32) {
	customListViewImportAPI().SysCallN(30, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListView) Items() IListItems {
	r1 := customListViewImportAPI().SysCallN(31, 0, m.Instance(), 0)
	return AsListItems(r1)
}

func (m *TCustomListView) SetItems(AValue IListItems) {
	customListViewImportAPI().SysCallN(31, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) MultiSelect() bool {
	r1 := customListViewImportAPI().SysCallN(33, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetMultiSelect(AValue bool) {
	customListViewImportAPI().SysCallN(33, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) OwnerData() bool {
	r1 := customListViewImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetOwnerData(AValue bool) {
	customListViewImportAPI().SysCallN(34, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) ReadOnly() bool {
	r1 := customListViewImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetReadOnly(AValue bool) {
	customListViewImportAPI().SysCallN(35, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) RowSelect() bool {
	r1 := customListViewImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetRowSelect(AValue bool) {
	customListViewImportAPI().SysCallN(36, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) SelCount() int32 {
	r1 := customListViewImportAPI().SysCallN(37, m.Instance())
	return int32(r1)
}

func (m *TCustomListView) Selected() IListItem {
	r1 := customListViewImportAPI().SysCallN(39, 0, m.Instance(), 0)
	return AsListItem(r1)
}

func (m *TCustomListView) SetSelected(AValue IListItem) {
	customListViewImportAPI().SysCallN(39, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) LastSelected() IListItem {
	r1 := customListViewImportAPI().SysCallN(32, m.Instance())
	return AsListItem(r1)
}

func (m *TCustomListView) TopItem() IListItem {
	r1 := customListViewImportAPI().SysCallN(41, m.Instance())
	return AsListItem(r1)
}

func (m *TCustomListView) ViewOrigin() (resultPoint TPoint) {
	customListViewImportAPI().SysCallN(42, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomListView) SetViewOrigin(AValue *TPoint) {
	customListViewImportAPI().SysCallN(42, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomListView) VisibleRowCount() int32 {
	r1 := customListViewImportAPI().SysCallN(43, m.Instance())
	return int32(r1)
}

func (m *TCustomListView) AlphaSort() bool {
	r1 := customListViewImportAPI().SysCallN(1, m.Instance())
	return GoBool(r1)
}

func (m *TCustomListView) CustomSort(fn TLVCompare, AOptionalParam uint32) bool {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	r1 := customListViewImportAPI().SysCallN(13, m.Instance(), m.customSortPtr, uintptr(AOptionalParam))
	return GoBool(r1)
}

func (m *TCustomListView) FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem {
	r1 := customListViewImportAPI().SysCallN(16, m.Instance(), uintptr(StartIndex), PascalStr(Value), PascalBool(Partial), PascalBool(Inclusive), PascalBool(Wrap), PascalBool(PartStart))
	return AsListItem(r1)
}

func (m *TCustomListView) FindData(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem {
	r1 := customListViewImportAPI().SysCallN(17, m.Instance(), uintptr(StartIndex), uintptr(Value), PascalBool(Inclusive), PascalBool(Wrap))
	return AsListItem(r1)
}

func (m *TCustomListView) GetHitTestInfoAt(X, Y int32) THitTests {
	r1 := customListViewImportAPI().SysCallN(20, m.Instance(), uintptr(X), uintptr(Y))
	return THitTests(r1)
}

func (m *TCustomListView) GetItemAt(x, y int32) IListItem {
	r1 := customListViewImportAPI().SysCallN(21, m.Instance(), uintptr(x), uintptr(y))
	return AsListItem(r1)
}

func (m *TCustomListView) GetNearestItem(APoint *TPoint, Direction TSearchDirection) IListItem {
	r1 := customListViewImportAPI().SysCallN(22, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(Direction))
	return AsListItem(r1)
}

func (m *TCustomListView) GetNextItem(StartItem IListItem, Direction TSearchDirection, States TListItemStates) IListItem {
	r1 := customListViewImportAPI().SysCallN(23, m.Instance(), GetObjectUintptr(StartItem), uintptr(Direction), uintptr(States))
	return AsListItem(r1)
}

func (m *TCustomListView) IsEditing() bool {
	r1 := customListViewImportAPI().SysCallN(28, m.Instance())
	return GoBool(r1)
}

func CustomListViewClass() TClass {
	ret := customListViewImportAPI().SysCallN(7)
	return TClass(ret)
}

func (m *TCustomListView) AddItem(Item string, AObject IObject) {
	customListViewImportAPI().SysCallN(0, m.Instance(), PascalStr(Item), GetObjectUintptr(AObject))
}

func (m *TCustomListView) Sort() {
	customListViewImportAPI().SysCallN(40, m.Instance())
}

func (m *TCustomListView) BeginUpdate() {
	customListViewImportAPI().SysCallN(2, m.Instance())
}

func (m *TCustomListView) Clear() {
	customListViewImportAPI().SysCallN(8, m.Instance())
}

func (m *TCustomListView) EndUpdate() {
	customListViewImportAPI().SysCallN(15, m.Instance())
}

func (m *TCustomListView) ClearSelection() {
	customListViewImportAPI().SysCallN(9, m.Instance())
}

func (m *TCustomListView) SelectAll() {
	customListViewImportAPI().SysCallN(38, m.Instance())
}

var (
	customListViewImport       *imports.Imports = nil
	customListViewImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomListView_AddItem", 0),
		/*1*/ imports.NewTable("CustomListView_AlphaSort", 0),
		/*2*/ imports.NewTable("CustomListView_BeginUpdate", 0),
		/*3*/ imports.NewTable("CustomListView_BorderStyle", 0),
		/*4*/ imports.NewTable("CustomListView_BoundingRect", 0),
		/*5*/ imports.NewTable("CustomListView_Canvas", 0),
		/*6*/ imports.NewTable("CustomListView_Checkboxes", 0),
		/*7*/ imports.NewTable("CustomListView_Class", 0),
		/*8*/ imports.NewTable("CustomListView_Clear", 0),
		/*9*/ imports.NewTable("CustomListView_ClearSelection", 0),
		/*10*/ imports.NewTable("CustomListView_Column", 0),
		/*11*/ imports.NewTable("CustomListView_ColumnCount", 0),
		/*12*/ imports.NewTable("CustomListView_Create", 0),
		/*13*/ imports.NewTable("CustomListView_CustomSort", 0),
		/*14*/ imports.NewTable("CustomListView_DropTarget", 0),
		/*15*/ imports.NewTable("CustomListView_EndUpdate", 0),
		/*16*/ imports.NewTable("CustomListView_FindCaption", 0),
		/*17*/ imports.NewTable("CustomListView_FindData", 0),
		/*18*/ imports.NewTable("CustomListView_FlatScrollBars", 0),
		/*19*/ imports.NewTable("CustomListView_FullDrag", 0),
		/*20*/ imports.NewTable("CustomListView_GetHitTestInfoAt", 0),
		/*21*/ imports.NewTable("CustomListView_GetItemAt", 0),
		/*22*/ imports.NewTable("CustomListView_GetNearestItem", 0),
		/*23*/ imports.NewTable("CustomListView_GetNextItem", 0),
		/*24*/ imports.NewTable("CustomListView_GridLines", 0),
		/*25*/ imports.NewTable("CustomListView_HotTrack", 0),
		/*26*/ imports.NewTable("CustomListView_HotTrackStyles", 0),
		/*27*/ imports.NewTable("CustomListView_IconOptions", 0),
		/*28*/ imports.NewTable("CustomListView_IsEditing", 0),
		/*29*/ imports.NewTable("CustomListView_ItemFocused", 0),
		/*30*/ imports.NewTable("CustomListView_ItemIndex", 0),
		/*31*/ imports.NewTable("CustomListView_Items", 0),
		/*32*/ imports.NewTable("CustomListView_LastSelected", 0),
		/*33*/ imports.NewTable("CustomListView_MultiSelect", 0),
		/*34*/ imports.NewTable("CustomListView_OwnerData", 0),
		/*35*/ imports.NewTable("CustomListView_ReadOnly", 0),
		/*36*/ imports.NewTable("CustomListView_RowSelect", 0),
		/*37*/ imports.NewTable("CustomListView_SelCount", 0),
		/*38*/ imports.NewTable("CustomListView_SelectAll", 0),
		/*39*/ imports.NewTable("CustomListView_Selected", 0),
		/*40*/ imports.NewTable("CustomListView_Sort", 0),
		/*41*/ imports.NewTable("CustomListView_TopItem", 0),
		/*42*/ imports.NewTable("CustomListView_ViewOrigin", 0),
		/*43*/ imports.NewTable("CustomListView_VisibleRowCount", 0),
	}
)

func customListViewImportAPI() *imports.Imports {
	if customListViewImport == nil {
		customListViewImport = NewDefaultImports()
		customListViewImport.SetImportTable(customListViewImportTables)
		customListViewImportTables = nil
	}
	return customListViewImport
}
